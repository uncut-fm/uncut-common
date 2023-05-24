package config

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	"errors"
	"fmt"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"gopkg.in/yaml.v3"
	"os"
)

func LoadConfigsFromSecretManager(ctx context.Context, client *secretmanager.Client, configStruct interface{}) error {
	secretName, err := getSecretNameByConfigStruct(configStruct)
	if err != nil {
		return err
	}

	return loadConfigFromSecretManager(ctx, client, secretName, configStruct)
}

func getSecretNameByConfigStruct(configStruct interface{}) (string, error) {
	switch configStruct.(type) {
	case *Common:
		return os.Getenv(CommonSecretEnvVar), nil
	case *ManagementConfigs:
		return os.Getenv(ManagementSecretEnvVar), nil
	case *BackofficeConfigs:
		return os.Getenv(BackofficeSecretEnvVar), nil
	case *AuthConfigs:
		return os.Getenv(AuthSecretEnvVar), nil
	case *TranscoderConfigs:
		return os.Getenv(TranscoderSecretEnvVar), nil
	case *Web3EventsConfigs:
		return os.Getenv(Web3EventsSecretEnvVar), nil
	case *GraphConfigs:
		return os.Getenv(GraphSecretEnvVar), nil
	default:
		return "", errors.New("unsupported configStruct")
	}
}

func loadConfigFromSecretManager(ctx context.Context, client *secretmanager.Client, name string, configStruct interface{}) error {
	secretData, err := getSecretFromSecretManager(ctx, client, name)
	if err != nil {
		return err
	}

	return unmarshallYamlIntoStruct(secretData, configStruct)
}

func getSecretFromSecretManager(ctx context.Context, client *secretmanager.Client, name string) ([]byte, error) {
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to access secret version: %v", err)
	}

	return result.Payload.Data, nil
}

func unmarshallYamlIntoStruct(data []byte, dstStruct interface{}) error {
	err := yaml.Unmarshal(data, dstStruct)
	return err
}

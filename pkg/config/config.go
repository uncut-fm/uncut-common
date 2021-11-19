package config

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	"fmt"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"gopkg.in/yaml.v3"
)

const (
	CommonSecretEnvVar     = "COMMON_SECRET_ID"
	ManagementSecretEnvVar = "MANAGEMENT_SECRET_ID"
)

type Configuration struct {
	Common     Common            `yaml:"common"`
	Management ManagementConfigs `yaml:"management"`
}

type Common struct {
	ProjectID   string `yaml:"project_id"`
	Environment string `yaml:"environment"`
}

type ManagementConfigs struct {
	PlatformConfigs DBConfigs `yaml:"platform_configs"`
}

type DBConfigs struct {
	Host                  string `yaml:"host"`
	Port                  string `yaml:"port"`
	DBName                string `yaml:"db_name"`
	User                  string `yaml:"user"`
	Password              string `yaml:"password"`
	MaxOpenConnections    int    `yaml:"max_open_conns"`
	MaxIdleConnections    int    `yaml:"max_idle_conns"`
	ConnectionMaxLifetime string `yaml:"conn_max_lifetime"`
}

func LoadConfigsFromSecretManager(ctx context.Context, client *secretmanager.Client, secretName string, dstStruct interface{}) error {
	return loadConfigFromSecretManager(ctx, client, secretName, dstStruct)
}

func loadConfigFromSecretManager(ctx context.Context, client *secretmanager.Client, name string, dstStruct interface{}) error {
	secretData, err := getSecretFromSecretManager(ctx, client, name)
	if err != nil {
		return err
	}

	return unmarshallYamlIntoStruct(secretData, dstStruct)
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

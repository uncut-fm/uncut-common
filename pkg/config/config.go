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

const (
	CommonSecretEnvVar     = "COMMON_SECRET_ID"
	ManagementSecretEnvVar = "MANAGEMENT_SECRET_ID"
	BackofficeSecretEnvVar = "BACKOFFICE_SECRET_ID"
	AuthSecretEnvVar       = "AUTH_SECRET_ID"
)

type Configuration struct {
	Common     Common            `yaml:"common"`
	Management ManagementConfigs `yaml:"management"`
	Backoffice BackofficeConfigs `yaml:"backoffice"`
	Auth       AuthConfigs       `yaml:"auth"`
}

type Common struct {
	ProjectID   string     `yaml:"project_id"`
	Environment string     `yaml:"environment"`
	JWT         JWTConfigs `yaml:"jwt"`
}

type ManagementConfigs struct {
	PlatformDB DBConfigs     `yaml:"platform_db"`
	Server     ServerConfigs `yaml:"server"`
}

type BackofficeConfigs struct {
	PlatformDB       DBConfigs      `yaml:"platform_db"`
	Server           ServerConfigs  `yaml:"server"`
	Twitter          TwitterConfigs `yaml:"twitter"`
	SearchAPIUrl     string         `yaml:"search_api_url"`
	GcpStorageBucket string         `yaml:"gcp_storage_bucket"`
	AdminToken       string         `yaml:"admin_token"`
}

type AuthConfigs struct {
	AuthDB         DBConfigs      `yaml:"auth_db"`
	Server         ServerConfigs  `yaml:"server"`
	OauthProviders OauthProviders `yaml:"oauth_providers"`
	AdminToken     string         `yaml:"admin_token"`
	BaseURL        string         `yaml:"base_url"`
	MagicLinkKey   string         `yaml:"magic_link_key"`
}

type ServerConfigs struct {
	Port string `yaml:"port"`
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

type TwitterConfigs struct {
	ConsumerKey    string `yaml:"consumer_key"`
	ConsumerSecret string `yaml:"consumer_secret"`
}

type GoogleOauthConfigs struct {
	ClientKey string `yaml:"client_key"`
	Secret    string `yaml:"secret"`
}

type RedisConfigs struct {
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	DB       int    `yaml:"db"`
	Password string `yaml:"password"`
}

type JWTConfigs struct {
	SigningMethod   string `yaml:"signing_method"`
	AccessSecret    string `yaml:"access_secret"`
	AccessDuration  int    `yaml:"access_duration"`
	RefreshDuration int    `yaml:"refresh_duration"`
}

type OauthProviders struct {
	Twitter TwitterConfigs     `yaml:"twitter"`
	Google  GoogleOauthConfigs `yaml:"google"`
}

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

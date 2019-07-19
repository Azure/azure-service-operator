package config

import (
	"fmt"
	"os"

	"github.com/Azure/go-autorest/autorest/azure"
)

const (
	CloudNameEnvVar         = "CLOUD_NAME"
	SubscriptionIDEnvVar    = "SUBSCRIPTION_ID"
	ClientIDEnvVar          = "CLIENT_ID"
	ClientSecretEnvVar      = "CLIENT_SECRET"
	TenantIDEnvVar          = "TENANT_ID"
	UseAADPodIdentityEnvVar = "USE_AAD_POD_IDENTITY"
)

// GetEnvVar returns the value of the environment variable
func GetEnvVar(envVarName string) string {
	v, found := os.LookupEnv(envVarName)
	if !found {
		panic(fmt.Sprintf("%s must be set", envVarName))
	}
	return v
}

// CloudName returns the cloud name
func CloudName() string {
	return GetEnvVar(CloudNameEnvVar)
}

// GetSubscriptionID returns the subscription ID
func SubscriptionID() string {
	return GetEnvVar(SubscriptionIDEnvVar)
}

// GetClientID returns the client ID
func ClientID() string {
	return GetEnvVar(ClientIDEnvVar)
}

// GetClientSecret returns the client secret
func ClientSecret() string {
	return GetEnvVar(ClientSecretEnvVar)
}

// GetTenantID returns the tenant ID
func TenantID() string {
	return GetEnvVar(TenantIDEnvVar)
}

// UseAADPodIdentity returns whether AAD Pod Identity is used
func UseAADPodIdentity() bool {
	return GetEnvVar(UseAADPodIdentityEnvVar) == "true"
}

// Environment() returns an `azure.Environment{...}` for the current cloud.
func Environment() azure.Environment {
	cloudName := CloudName()
	env, err := azure.EnvironmentFromName(cloudName)
	if err != nil {
		panic(fmt.Sprintf(
			"invalid cloud name '%s' specified, cannot continue\n", cloudName))
	}
	return env
}

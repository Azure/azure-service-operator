// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/gobuffalo/envy"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

type ConfigRequirementType int

const (
	RequireClientID ConfigRequirementType = iota
	RequireClientSecret
	RequireTenantID
	RequireSubscriptionID
	OptionalClientID
)

// ParseEnvironment loads a sibling `.env` file then looks through all environment
// variables to set global configuration.
func ParseEnvironment() error {
	azcloud := os.Getenv("AZURE_CLOUD_ENV")
	envy.Load()

	if azcloud == "" {
		azcloud = "AzurePublicCloud"
	}

	allowed := []string{
		"AzurePublicCloud",
		"AzureUSGovernmentCloud",
		"AzureChinaCloud",
		"AzureGermanCloud",
	}

	if !helpers.ContainsString(allowed, azcloud) {
		return fmt.Errorf("Invalid Cloud chosen: AZURE_CLOUD_ENV set to '%s'", azcloud)
	}

	var err error
	cloudName = azcloud

	azureEnv, _ := azure.EnvironmentFromName(azcloud) // shouldn't fail
	authorizationServerURL = azureEnv.ActiveDirectoryEndpoint
	baseURI = azureEnv.ResourceManagerEndpoint // BaseURI()

	locationDefault = envy.Get("AZURE_LOCATION_DEFAULT", "westus2")                 // DefaultLocation()
	useDeviceFlow = ParseBoolFromEnvironment("AZURE_USE_DEVICEFLOW", false)         // UseDeviceFlow()
	creds.useManagedIdentity = ParseBoolFromEnvironment("AZURE_USE_MI", false)      // UseManagedIdentity()
	keepResources = ParseBoolFromEnvironment("AZURE_SAMPLES_KEEP_RESOURCES", false) // KeepResources()
	creds.operatorKeyvault = envy.Get("AZURE_OPERATOR_KEYVAULT", "")                // operatorKeyvault()
	testResourcePrefix = envy.Get("TEST_RESOURCE_PREFIX", "t-"+helpers.RandomString(6))
	podNamespace, err = envy.MustGet("POD_NAMESPACE")
	if err != nil {
		return errors.Wrapf(err, "couldn't get POD_NAMESPACE env variable")
	}
	targetNamespaces = ParseStringListFromEnvironment("AZURE_TARGET_NAMESPACES")
	purgeDeletedKeyVaultSecrets = ParseBoolFromEnvironment("PURGE_DELETED_KEYVAULT_SECRETS", false)
	recoverSoftDeletedKeyVaultSecrets = ParseBoolFromEnvironment("RECOVER_SOFT_DELETED_KEYVAULT_SECRETS", true)

	operatorMode, err = ParseOperatorMode(envy.Get("AZURE_OPERATOR_MODE", OperatorModeBoth.String()))
	if err != nil {
		return errors.Wrap(err, "reading AZURE_OPERATOR_MODE")
	}

	if !operatorMode.IncludesWatchers() && len(targetNamespaces) > 0 {
		return errors.Errorf("mode must include watchers to specify target namespaces")
	}

	secretNamingVersionInt, err := ParseIntFromEnvironment("AZURE_SECRET_NAMING_VERSION")
	if err != nil {
		return errors.Wrapf(err, "couldn't get AZURE_SECRET_NAMING_VERSION env variable")
	}

	// If this isn't set, default to version 2
	if secretNamingVersionInt == 1 {
		secretNamingVersion = secrets.SecretNamingV1
	} else if secretNamingVersionInt == 0 || secretNamingVersionInt == 2 {
		secretNamingVersion = secrets.SecretNamingV2
	} else {
		return errors.Errorf("secret naming version must be one of 0, 1, or 2 but was %d", secretNamingVersionInt)
	}

	for _, requirement := range GetExpectedConfigurationVariables() {
		switch requirement {
		case RequireClientID:
			creds.clientID, err = envy.MustGet("AZURE_CLIENT_ID") // ClientID()
			if err != nil {
				return fmt.Errorf("expected env vars not provided (AZURE_CLIENT_ID): %s\n", err)
			}
		case OptionalClientID:
			creds.clientID = envy.Get("AZURE_CLIENT_ID", "")
		case RequireClientSecret:
			creds.clientSecret, err = envy.MustGet("AZURE_CLIENT_SECRET") // ClientSecret()
			if err != nil {
				return fmt.Errorf("expected env vars not provided (AZURE_CLIENT_SECRET): %s\n", err)
			}
		case RequireTenantID:
			creds.tenantID, err = envy.MustGet("AZURE_TENANT_ID") // TenantID()
			if err != nil {
				return fmt.Errorf("expected env vars not provided (AZURE_TENANT_ID): %s\n", err)
			}
		case RequireSubscriptionID:
			creds.subscriptionID, err = envy.MustGet("AZURE_SUBSCRIPTION_ID") // SubscriptionID()
			if err != nil {
				return fmt.Errorf("expected env vars not provided (AZURE_SUBSCRIPTION_ID): %s\n", err)
			}
		}

	}

	return nil
}

func GetExpectedConfigurationVariables() []ConfigRequirementType {
	if useDeviceFlow {
		// Device flow required Configs
		return []ConfigRequirementType{RequireClientID, RequireTenantID, RequireSubscriptionID}
	}
	if creds.useManagedIdentity {
		// Managed Service Identity required Configs
		return []ConfigRequirementType{RequireTenantID, RequireSubscriptionID, OptionalClientID}
	}
	// Default required Configs (service principal)
	return []ConfigRequirementType{RequireClientID, RequireClientSecret, RequireTenantID, RequireSubscriptionID}
}

func ParseBoolFromEnvironment(variable string, defaultValue bool) bool {
	defaultValueStr := fmt.Sprintf("%t", defaultValue)
	env := envy.Get(variable, defaultValueStr)
	value, err := strconv.ParseBool(env)
	if err != nil {
		log.Printf("WARNING: invalid input value specified for %q, expected bool, actual: %q. Disabling\n", variable, env)
		value = false
	}
	return value
}

func ParseIntFromEnvironment(variable string) (int, error) {
	env := envy.Get(variable, "0")
	value, err := strconv.Atoi(env)
	if err != nil {
		return 0, errors.Wrapf(err, "invalid input value specified for %q, expected int, actual: %q", variable, env)
	}
	return value, nil
}

func ParseStringListFromEnvironment(variable string) []string {
	env := envy.Get(variable, "")
	if len(strings.TrimSpace(env)) == 0 {
		return nil
	}
	items := strings.Split(env, ",")
	// Remove any whitespace used to separate items.
	for i, item := range items {
		items[i] = strings.TrimSpace(item)
	}
	return items
}

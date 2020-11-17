// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/gobuffalo/envy"
	"github.com/pkg/errors"
)

type ConfigRequirementType int

const (
	RequireClientID ConfigRequirementType = iota
	RequireClientSecret
	RequireTenantID
	RequireSubscriptionID
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

	locationDefault = envy.Get("AZURE_LOCATION_DEFAULT", "westus2")          // DefaultLocation()
	useDeviceFlow = ParseBoolFromEnvironment("AZURE_USE_DEVICEFLOW")         // UseDeviceFlow()
	creds.useManagedIdentity = ParseBoolFromEnvironment("AZURE_USE_MI")      // UseManagedIdentity()
	keepResources = ParseBoolFromEnvironment("AZURE_SAMPLES_KEEP_RESOURCES") // KeepResources()
	creds.operatorKeyvault = envy.Get("AZURE_OPERATOR_KEYVAULT", "")         // operatorKeyvault()
	testResourcePrefix = envy.Get("TEST_RESOURCE_PREFIX", "t-"+helpers.RandomString(6))
	podNamespace, err = envy.MustGet("POD_NAMESPACE")
	if err != nil {
		return errors.Wrapf(err, "couldn't get POD_NAMESPACE env variable")
	}

	for _, requirement := range GetRequiredConfigs() {
		switch requirement {
		case RequireClientID:
			creds.clientID, err = envy.MustGet("AZURE_CLIENT_ID") // ClientID()
			if err != nil {
				return fmt.Errorf("expected env vars not provided (AZURE_CLIENT_ID): %s\n", err)
			}
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

func GetRequiredConfigs() []ConfigRequirementType {
	if useDeviceFlow {
		// Device flow required Configs
		return []ConfigRequirementType{RequireClientID, RequireTenantID, RequireSubscriptionID}
	}
	if creds.useManagedIdentity {
		// Managed Service Identity required Configs
		return []ConfigRequirementType{RequireTenantID, RequireSubscriptionID}
	}
	// Default required Configs
	return []ConfigRequirementType{RequireClientID, RequireClientSecret, RequireTenantID, RequireSubscriptionID}
}

func ParseBoolFromEnvironment(variable string) bool {
	value, err := strconv.ParseBool(envy.Get(variable, "0"))
	if err != nil {
		log.Printf("WARNING: invalid input value specified for bool %v: \"%v\", disabling\n", variable, value)
		value = false
	}
	return value
}

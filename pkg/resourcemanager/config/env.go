package config

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/gobuffalo/envy"
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
	envy.Load()
	azureEnv, _ := azure.EnvironmentFromName("AzurePublicCloud") // shouldn't fail
	authorizationServerURL = azureEnv.ActiveDirectoryEndpoint

	// AZURE_GROUP_NAME and `config.GroupName()` are deprecated.
	// Use AZURE_BASE_GROUP_NAME and `config.GenerateGroupName()` instead.
	groupName = envy.Get("AZURE_GROUP_NAME", "azure-go-samples")    // GroupName()
	baseGroupName = envy.Get("AZURE_BASE_GROUP_NAME", groupName)    // BaseGroupName()
	locationDefault = envy.Get("AZURE_LOCATION_DEFAULT", "westus2") // DefaultLocation()

	useDeviceFlow = ParseBoolFromEnvironment("AZURE_USE_DEVICEFLOW")         // UseDeviceFlow()
	useMSI = ParseBoolFromEnvironment("AZURE_USE_MSI")                       // UseMSI()
	keepResources = ParseBoolFromEnvironment("AZURE_SAMPLES_KEEP_RESOURCES") // KeepResources()

	var err error

	for _, requirement := range GetRequiredConfigs() {
		switch requirement {
		case RequireClientID:
			clientID, err = envy.MustGet("AZURE_CLIENT_ID") // ClientID()
			if err != nil {
				return fmt.Errorf("expected env vars not provided (AZURE_CLIENT_ID): %s\n", err)
			}
		case RequireClientSecret:
			clientSecret, err = envy.MustGet("AZURE_CLIENT_SECRET") // ClientSecret()
			if err != nil {
				return fmt.Errorf("expected env vars not provided (AZURE_CLIENT_SECRET): %s\n", err)
			}
		case RequireTenantID:
			tenantID, err = envy.MustGet("AZURE_TENANT_ID") // TenantID()
			if err != nil {
				return fmt.Errorf("expected env vars not provided (AZURE_TENANT_ID): %s\n", err)
			}
		case RequireSubscriptionID:
			subscriptionID, err = envy.MustGet("AZURE_SUBSCRIPTION_ID") // SubscriptionID()
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
	if useMSI {
		// Managed Service Identity required Configs
		return []ConfigRequirementType{RequireTenantID, RequireSubscriptionID}
	}
	// Default required Configs
	return []ConfigRequirementType{RequireClientID, RequireClientSecret, RequireTenantID, RequireSubscriptionID}
}

func ParseBoolFromEnvironment(variable_name string) bool {
	value, err := strconv.ParseBool(envy.Get(variable_name, "0"))
	if err != nil {
		log.Printf("WARNING: invalid input value specified for bool %v: \"%v\", disabling\n", variable_name, value)
		value = false
	}
	return value
}

package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/gobuffalo/envy"
)

// LoadSettings loads blah
func LoadSettings() error {

	fmt.Println()

	azureEnv, _ := azure.EnvironmentFromName("AzurePublicCloud") // shouldn't fail
	authorizationServerURL = azureEnv.ActiveDirectoryEndpoint

	// AZURE_GROUP_NAME and `config.GroupName()` are deprecated.
	// Use AZURE_BASE_GROUP_NAME and `config.GenerateGroupName()` instead.
	groupName = os.Getenv("AZURE_GROUP_NAME")
	baseGroupName = os.Getenv("AZURE_BASE_GROUP_NAME")

	locationDefault = os.Getenv("AZURE_LOCATION_DEFAULT")

	// these must be provided by environment
	// clientID
	clientID = os.Getenv("AZURE_CLIENT_ID")
	if len(clientID) == 0 {
		return fmt.Errorf("expected env vars not provided")
	}

	// clientSecret
	clientSecret = os.Getenv("AZURE_CLIENT_SECRET")
	if len(clientSecret) == 0 { // don't need a secret for device flow
		return fmt.Errorf("expected env vars not provided")
	}

	// tenantID (AAD)
	tenantID = os.Getenv("AZURE_TENANT_ID")
	if len(tenantID) == 0 {
		return fmt.Errorf("expected env vars not provided")
	}

	// subscriptionID (ARM)
	subscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")
	if len(subscriptionID) == 0 {
		return fmt.Errorf("expected env vars not provided")
	}

	return nil
}

// ParseEnvironment loads a sibling `.env` file then looks through all environment
// variables to set global configuration.
func ParseEnvironment() error {
	envy.Load()
	azureEnv, _ := azure.EnvironmentFromName("AzurePublicCloud") // shouldn't fail
	authorizationServerURL = azureEnv.ActiveDirectoryEndpoint

	// AZURE_GROUP_NAME and `config.GroupName()` are deprecated.
	// Use AZURE_BASE_GROUP_NAME and `config.GenerateGroupName()` instead.
	groupName = envy.Get("AZURE_GROUP_NAME", "azure-go-samples")
	baseGroupName = envy.Get("AZURE_BASE_GROUP_NAME", groupName)

	locationDefault = envy.Get("AZURE_LOCATION_DEFAULT", "westus2")

	var err error
	useDeviceFlow, err = strconv.ParseBool(envy.Get("AZURE_USE_DEVICEFLOW", "0"))
	if err != nil {
		log.Printf("invalid value specified for AZURE_USE_DEVICEFLOW, disabling\n")
		useDeviceFlow = false
	}
	keepResources, err = strconv.ParseBool(envy.Get("AZURE_SAMPLES_KEEP_RESOURCES", "0"))
	if err != nil {
		log.Printf("invalid value specified for AZURE_SAMPLES_KEEP_RESOURCES, discarding\n")
		keepResources = false
	}

	// these must be provided by environment
	// clientID
	clientID, err = envy.MustGet("AZURE_CLIENT_ID")
	if err != nil {
		return fmt.Errorf("expected env vars not provided: %s\n", err)
	}

	// clientSecret
	clientSecret, err = envy.MustGet("AZURE_CLIENT_SECRET")
	if err != nil && useDeviceFlow != true { // don't need a secret for device flow
		return fmt.Errorf("expected env vars not provided: %s\n", err)
	}

	// tenantID (AAD)
	tenantID, err = envy.MustGet("AZURE_TENANT_ID")
	if err != nil {
		return fmt.Errorf("expected env vars not provided: %s\n", err)
	}

	// subscriptionID (ARM)
	subscriptionID, err = envy.MustGet("AZURE_SUBSCRIPTION_ID")
	if err != nil {
		return fmt.Errorf("expected env vars not provided: %s\n", err)
	}

	return nil
}

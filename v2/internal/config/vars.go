// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/pkg/errors"
)

const (
	// #nosec
	ClientSecretVar            = "AZURE_CLIENT_SECRET"
	SubscriptionIDVar          = "AZURE_SUBSCRIPTION_ID"
	TenantIDVar                = "AZURE_TENANT_ID"
	ClientIDVar                = "AZURE_CLIENT_ID"
	AzureFederatedTokenFileVar = "AZURE_FEDERATED_TOKEN_FILE"
	targetNamespacesVar        = "AZURE_TARGET_NAMESPACES"
	operatorModeVar            = "AZURE_OPERATOR_MODE"
	syncPeriodVar              = "AZURE_SYNC_PERIOD"
	resourceManagerEndpointVar = "AZURE_RESOURCE_MANAGER_ENDPOINT"
	resourceManagerAudienceVar = "AZURE_RESOURCE_MANAGER_AUDIENCE"
	azureAuthorityHostVar      = "AZURE_AUTHORITY_HOST"
	podNamespaceVar            = "POD_NAMESPACE"
	useWorkloadIdentityAuth    = "USE_WORKLOAD_IDENTITY_AUTH"

	// TODO: These values are used for single operator multitenancy tests. We can get rid of them once we have Managed Identity support.
	AzureClientIDMultitenantVar = "AZURE_CLIENT_ID_MULTITENANT"
	// #nosec
	AzureClientSecretMultitenantVar = "AZURE_CLIENT_SECRET_MULTITENANT"
)

// These are hardcoded because the init function that initializes them in azcore isn't in /cloud it's in /arm which
// we don't import.

var DefaultEndpoint = "https://management.azure.com"
var DefaultAudience = "https://management.core.windows.net/"
var DefaultAADAuthorityHost = "https://login.microsoftonline.com/"

// NOTE: Changes to documentation or available values here should be documented in Helm values.yaml as well

// Values stores configuration values that are set for the operator.
type Values struct {
	// SubscriptionID is the Azure subscription the operator will use
	// for ARM communication.
	SubscriptionID string

	// TenantID is the Azure tenantID the operator will use
	// for ARM communication.
	TenantID string

	// ClientID is the Azure clientID the operator will use
	// for ARM communication.
	ClientID string

	// PodNamespace is the namespace the operator pods are running in.
	PodNamespace string

	// OperatorMode determines whether the operator should run
	// watchers, webhooks or both.
	OperatorMode OperatorMode

	// TargetNamespaces lists the namespaces the operator will watch
	// for Azure resources (if the mode includes running watchers). If
	// it's empty the operator will watch all namespaces.
	TargetNamespaces []string

	// SyncPeriod is the frequency at which resources are re-reconciled with Azure
	// when there have been no triggering changes in the Kubernetes resources. This sync
	// exists to detect and correct changes that happened in Azure that Kubernetes is not
	// aware about. BE VERY CAREFUL setting this value low - even a modest number of resources
	// can cause subscription level throttling if they are re-synced frequently.
	// If nil, no sync is performed. Durations are specified as "1h", "15m", or "60s". See
	// https://pkg.go.dev/time#ParseDuration for more details.
	//
	// This can be set to nil by specifying empty string for AZURE_SYNC_PERIOD explicitly in
	// the config.
	SyncPeriod *time.Duration

	// ResourceManagerEndpoint is the Azure Resource Manager endpoint.
	// If not specified, the default is the Public cloud resource manager endpoint.
	// See https://docs.microsoft.com/cli/azure/manage-clouds-azure-cli#list-available-clouds for details
	// about how to find available resource manager endpoints for your cloud. Note that the resource manager
	// endpoint is referred to as "resourceManager" in the Azure CLI.
	ResourceManagerEndpoint string

	// ResourceManagerAudience is the Azure Resource Manager AAD audience.
	// If not specified, the default is the Public cloud resource manager audience https://management.core.windows.net/.
	// See https://docs.microsoft.com/cli/azure/manage-clouds-azure-cli#list-available-clouds for details
	// about how to find available resource manager audiences for your cloud. Note that the resource manager
	// audience is referred to as "activeDirectoryResourceId" in the Azure CLI.
	ResourceManagerAudience string

	// AzureAuthorityHost is the URL of the AAD authority. If not specified, the default
	// is the AAD URL for the public cloud: https://login.microsoftonline.com/. See
	// https://docs.microsoft.com/azure/active-directory/develop/authentication-national-cloud
	AzureAuthorityHost string

	// UseWorkloadIdentityAuth boolean is used to determine if we're using Workload Identity authentication for global credential
	UseWorkloadIdentityAuth bool
}

var _ fmt.Stringer = Values{}

// Returns the configuration as a string
func (v Values) String() string {
	return fmt.Sprintf(
		"SubscriptionID:%s/TenantID:%s/ClientID:%s/PodNamespace:%s/OperatorMode:%s/TargetNamespaces:%s/SyncPeriod:%s/ResourceManagerEndpoint:%s/ResourceManagerAudience:%s/AzureAuthorityHost:%s/UseWorkloadIdentityAuth:%t",
		v.SubscriptionID,
		v.TenantID,
		v.ClientID,
		v.PodNamespace,
		v.OperatorMode,
		strings.Join(v.TargetNamespaces, "|"),
		v.SyncPeriod,
		v.ResourceManagerEndpoint,
		v.ResourceManagerAudience,
		v.AzureAuthorityHost,
		v.UseWorkloadIdentityAuth)
}

// Cloud returns the cloud the configuration is using
func (v Values) Cloud() cloud.Configuration {

	// Special handling if we've got all the defaults just return the official public cloud
	// configuration
	hasDefaultAzureAuthorityHost := v.AzureAuthorityHost == "" || v.AzureAuthorityHost == DefaultAADAuthorityHost
	hasDefaultResourceManagerEndpoint := v.ResourceManagerEndpoint == "" || v.ResourceManagerEndpoint == DefaultEndpoint
	hasDefaultResourceManagerAudience := v.ResourceManagerAudience == "" || v.ResourceManagerAudience == DefaultAudience

	if hasDefaultResourceManagerEndpoint && hasDefaultResourceManagerAudience && hasDefaultAzureAuthorityHost {
		return cloud.AzurePublic
	}

	// We default here too to more easily support empty Values objects
	azureAuthorityHost := v.AzureAuthorityHost
	resourceManagerEndpoint := v.ResourceManagerEndpoint
	resourceManagerAudience := v.ResourceManagerAudience
	if azureAuthorityHost == "" {
		azureAuthorityHost = DefaultAADAuthorityHost
	}
	if resourceManagerAudience == "" {
		resourceManagerAudience = DefaultAudience
	}
	if resourceManagerEndpoint == "" {
		resourceManagerEndpoint = DefaultEndpoint
	}

	return cloud.Configuration{
		ActiveDirectoryAuthorityHost: azureAuthorityHost,
		Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
			cloud.ResourceManager: {
				Endpoint: resourceManagerEndpoint,
				Audience: resourceManagerAudience,
			},
		},
	}
}

// ReadFromEnvironment loads configuration values from the AZURE_*
// environment variables.
func ReadFromEnvironment() (Values, error) {
	var result Values
	modeValue := os.Getenv(operatorModeVar)
	if modeValue == "" {
		result.OperatorMode = OperatorModeBoth
	} else {
		mode, err := ParseOperatorMode(modeValue)
		if err != nil {
			return Values{}, err
		}
		result.OperatorMode = mode
	}

	var err error

	result.SubscriptionID = os.Getenv(SubscriptionIDVar)
	result.PodNamespace = os.Getenv(podNamespaceVar)
	result.TargetNamespaces = parseTargetNamespaces(os.Getenv(targetNamespacesVar))
	result.SyncPeriod, err = parseSyncPeriod()
	result.ResourceManagerEndpoint = envOrDefault(resourceManagerEndpointVar, DefaultEndpoint)
	result.ResourceManagerAudience = envOrDefault(resourceManagerAudienceVar, DefaultAudience)
	result.AzureAuthorityHost = envOrDefault(azureAuthorityHostVar, DefaultAADAuthorityHost)
	result.ClientID = os.Getenv(ClientIDVar)
	result.TenantID = os.Getenv(TenantIDVar)

	// Ignoring error here, as any other value or empty value means we should default to false
	result.UseWorkloadIdentityAuth, _ = strconv.ParseBool(os.Getenv(useWorkloadIdentityAuth))

	if err != nil {
		return result, errors.Wrapf(err, "parsing %q", syncPeriodVar)
	}

	// Not calling validate here to support using from tests where we
	// don't require consistent settings.
	return result, nil
}

// ReadAndValidate loads the configuration values and checks that
// they're consistent.
func ReadAndValidate() (Values, error) {
	result, err := ReadFromEnvironment()
	if err != nil {
		return Values{}, err
	}
	err = result.Validate()
	if err != nil {
		return Values{}, err
	}
	return result, nil
}

// Validate checks whether the configuration settings are consistent.
func (v Values) Validate() error {
	if v.SubscriptionID == "" {
		return errors.Errorf("missing value for %s", SubscriptionIDVar)
	}
	if v.PodNamespace == "" {
		return errors.Errorf("missing value for %s", podNamespaceVar)
	}
	if !v.OperatorMode.IncludesWatchers() && len(v.TargetNamespaces) > 0 {
		return errors.Errorf("%s must include watchers to specify target namespaces", targetNamespacesVar)
	}
	return nil
}

// parseTargetNamespaces splits a comma-separated string into a slice
// of strings with spaces trimmed.
func parseTargetNamespaces(fromEnv string) []string {
	if len(strings.TrimSpace(fromEnv)) == 0 {
		return nil
	}
	items := strings.Split(fromEnv, ",")
	// Remove any whitespace used to separate items.
	for i, item := range items {
		items[i] = strings.TrimSpace(item)
	}
	return items
}

// parseSyncPeriod parses the sync period from the environment
func parseSyncPeriod() (*time.Duration, error) {
	syncPeriodStr := envOrDefault(syncPeriodVar, "1h")
	if syncPeriodStr == "" {
		return nil, nil
	}

	syncPeriod, err := time.ParseDuration(syncPeriodStr)
	if err != nil {
		return nil, err
	}
	return &syncPeriod, nil
}

// envOrDefault returns the value of the specified env variable or the default value if
// the env variable was not set.
func envOrDefault(env string, def string) string {
	result, specified := os.LookupEnv(env)
	if !specified {
		return def
	}

	return result
}

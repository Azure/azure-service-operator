/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package creds

import (
	"encoding/json"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/google/uuid"
	"github.com/rotisserie/eris"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/pkg/common/config"
)

// this is shared between tests because
// instantiating it requires HTTP calls
var (
	cachedCreds azcore.TokenCredential
	cachedIds   AzureIDs
)

const (
	TestBillingIDVar = "TEST_BILLING_ID"

	// TestResourceGroupTagsVar names the environment variable holding the tags the tenant used for
	// recording requires by policy on every resource group, as a JSON object:
	// {"Owner":"someone","Project":"something"}
	TestResourceGroupTagsVar = "TEST_RESOURCE_GROUP_TAGS"
)

type AzureIDs struct {
	SubscriptionID   string
	TenantID         string
	BillingInvoiceID string
	EntraAppID       string

	// ResourceGroupTags are added to the resource groups tests create, and redacted from recordings
	// again, so that a recording made in a tenant with such a policy is not tied to that tenant
	ResourceGroupTags map[string]string
}

// getCredentials returns the token credential authentication modes supported by
// the test framework.
// We primarily support two modes of authentication:
// - EnvironmentCredential
// - CLICredential
// We don't use NewDefaultAzureCredential because it puts CLI credentials last
// which can cause issues when trying to do CLI auth from clients such as Virtual DevBoxes (which have a UMI
// that gets preferred over the CLI credentials).
func getCredentials() (*azidentity.ChainedTokenCredential, error) {
	var result []azcore.TokenCredential
	var errs []error
	cliCred, err := azidentity.NewAzureCLICredential(nil)
	if err != nil {
		errs = append(errs, err)
	} else {
		result = append(result, cliCred)
	}

	envCred, err := azidentity.NewEnvironmentCredential(nil)
	if err != nil {
		errs = append(errs, err)
	} else {
		result = append(result, envCred)
	}

	if len(result) > 0 {
		var chained *azidentity.ChainedTokenCredential
		chained, err = azidentity.NewChainedTokenCredential(result, nil)
		if err != nil {
			return nil, err
		}
		return chained, nil
	} else {
		return nil, kerrors.NewAggregate(errs)
	}
}

func GetCreds() (azcore.TokenCredential, AzureIDs, error) {
	if cachedCreds != nil {
		return cachedCreds, cachedIds, nil
	}

	creds, err := getCredentials()
	if err != nil {
		return nil, AzureIDs{}, eris.Wrapf(err, "creating credentials")
	}

	// Read AZURE_SUBSCRIPTION_ID
	subscriptionID := os.Getenv(config.AzureSubscriptionID)
	if subscriptionID == "" {
		return nil, AzureIDs{}, eris.Errorf("required environment variable %q was not supplied", config.AzureSubscriptionID)
	}

	// Read AZURE_TENANT_ID
	tenantID := os.Getenv(config.AzureTenantID)
	if tenantID == "" {
		return nil, AzureIDs{}, eris.Errorf("required environment variable %q was not supplied", config.AzureTenantID)
	}

	// Read AZURE_ENTRA_APP_ID (optional; only used if Entra resources are used)
	entraID := os.Getenv(config.EntraAppID)

	// This is test specific and doesn't have a corresponding config entry. It's also optional as it's only required for
	// a small number of tests. Those tests will check for it explicitly
	billingInvoiceID := os.Getenv(TestBillingIDVar)

	resourceGroupTags, err := ResourceGroupTagsFromEnvironment()
	if err != nil {
		return nil, AzureIDs{}, err
	}

	ids := AzureIDs{
		SubscriptionID:    subscriptionID,
		TenantID:          tenantID,
		BillingInvoiceID:  billingInvoiceID,
		EntraAppID:        entraID,
		ResourceGroupTags: resourceGroupTags,
	}

	cachedCreds = creds
	cachedIds = ids
	return creds, ids, nil
}

// ResourceGroupTagsFromEnvironment returns the tags the tenant used for recording requires by policy on
// every resource group. Only recording needs them, so replaying without them set is expected.
//
// The tags are read as JSON so that a value can contain anything a tag can, a comma and significant
// whitespace included: a value that doesn't survive the round trip would be sent to Azure in one form and
// looked for in the recording in another, leaving it unredacted.
func ResourceGroupTagsFromEnvironment() (map[string]string, error) {
	setting := os.Getenv(TestResourceGroupTagsVar)
	if setting == "" {
		return nil, nil
	}

	var tags map[string]string
	err := json.Unmarshal([]byte(setting), &tags)
	if err != nil {
		return nil, eris.Wrapf(err, "reading %s, which must be a JSON object of tag names to values", TestResourceGroupTagsVar)
	}
	if tags == nil {
		return nil, eris.Errorf("reading %s, which must be a JSON object of tag names to values", TestResourceGroupTagsVar)
	}

	return tags, nil
}

func DummyAzureIDs() AzureIDs {
	// Replaying doesn't need credentials, but the tests still tag the resource groups they create when
	// the variable is set, so those tags have to be removed on the way out as they were when the
	// recording was made. A malformed setting is reported when recording; ignore it here.
	resourceGroupTags, _ := ResourceGroupTagsFromEnvironment()

	return AzureIDs{
		SubscriptionID:    uuid.Nil.String(),
		TenantID:          uuid.Nil.String(),
		BillingInvoiceID:  DummyBillingID,
		ResourceGroupTags: resourceGroupTags,
	}
}

// newScopedCredentialSecret is the internal factory used to create credential secrets
func newScopedCredentialSecret(subscriptionID, tenantID, name, namespace string) *v1.Secret {
	secretData := make(map[string][]byte)

	secretData[config.AzureTenantID] = []byte(tenantID)
	secretData[config.AzureSubscriptionID] = []byte(subscriptionID)

	return &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: secretData,
	}
}

func NewScopedServicePrincipalSecret(
	subscriptionID string,
	tenantID string,
	clientID string,
	clientSecret string,
	name string,
	namespace string,
) *v1.Secret {
	secret := newScopedCredentialSecret(subscriptionID, tenantID, name, namespace)

	secret.Data[config.AzureClientID] = []byte(clientID)
	secret.Data[config.AzureClientSecret] = []byte(clientSecret)

	return secret
}

func NewScopedManagedIdentitySecret(
	subscriptionID string,
	tenantID string,
	clientID string,
	name string,
	namespace string,
) *v1.Secret {
	secret := newScopedCredentialSecret(subscriptionID, tenantID, name, namespace)

	secret.Data[config.AzureClientID] = []byte(clientID)

	return secret
}

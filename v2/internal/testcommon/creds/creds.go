/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package creds

import (
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

const TestBillingIDVar = "TEST_BILLING_ID"

type AzureIDs struct {
	SubscriptionID   string
	TenantID         string
	BillingInvoiceID string
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

	subscriptionID := os.Getenv(config.AzureSubscriptionID)
	if subscriptionID == "" {
		return nil, AzureIDs{}, eris.Errorf("required environment variable %q was not supplied", config.AzureSubscriptionID)
	}

	tenantID := os.Getenv(config.AzureTenantID)
	if tenantID == "" {
		return nil, AzureIDs{}, eris.Errorf("required environment variable %q was not supplied", config.AzureTenantID)
	}

	// This is test specific and doesn't have a corresponding config entry. It's also optional as it's only required for
	// a small number of tests. Those tests will check for it explicitly
	billingInvoiceID := os.Getenv(TestBillingIDVar)

	ids := AzureIDs{
		SubscriptionID:   subscriptionID,
		TenantID:         tenantID,
		BillingInvoiceID: billingInvoiceID,
	}

	cachedCreds = creds
	cachedIds = ids
	return creds, ids, nil
}

func DummyAzureIDs() AzureIDs {
	return AzureIDs{
		SubscriptionID:   uuid.Nil.String(),
		TenantID:         uuid.Nil.String(),
		BillingInvoiceID: DummyBillingID,
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

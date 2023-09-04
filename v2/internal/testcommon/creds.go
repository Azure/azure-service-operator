/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/v2/pkg/common/config"
)

// this is shared between tests because
// instantiating it requires HTTP calls
var cachedCreds azcore.TokenCredential
var cachedIds AzureIDs

const TestBillingIDVar = "TEST_BILLING_ID"

type AzureIDs struct {
	subscriptionID   string
	tenantID         string
	billingInvoiceID string
}

func getCreds() (azcore.TokenCredential, AzureIDs, error) {

	if cachedCreds != nil {
		return cachedCreds, cachedIds, nil
	}

	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, AzureIDs{}, errors.Wrapf(err, "creating default credential")
	}

	subscriptionID := os.Getenv(config.AzureSubscriptionID)
	if subscriptionID == "" {
		return nil, AzureIDs{}, errors.Errorf("required environment variable %q was not supplied", config.AzureSubscriptionID)
	}

	tenantID := os.Getenv(config.AzureTenantID)
	if tenantID == "" {
		return nil, AzureIDs{}, errors.Errorf("required environment variable %q was not supplied", config.AzureTenantID)
	}

	// This is test specific and doesn't have a corresponding config entry. It's also optional as it's only required for
	// a small number of tests. Those tests will check for it explicitly
	billingInvoiceId := os.Getenv(TestBillingIDVar)

	ids := AzureIDs{
		subscriptionID:   subscriptionID,
		tenantID:         tenantID,
		billingInvoiceID: billingInvoiceId,
	}

	cachedCreds = creds
	cachedIds = ids
	return creds, ids, nil
}

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

func NewScopedServicePrincipalSecret(subscriptionID, tenantID, clientID, clientSecret, name, namespace string) *v1.Secret {
	secret := newScopedCredentialSecret(subscriptionID, tenantID, name, namespace)

	secret.Data[config.AzureClientID] = []byte(clientID)
	secret.Data[config.AzureClientSecret] = []byte(clientSecret)

	return secret
}

func NewScopedManagedIdentitySecret(subscriptionID, tenantID, clientID, name, namespace string) *v1.Secret {
	secret := newScopedCredentialSecret(subscriptionID, tenantID, name, namespace)

	secret.Data[config.AzureClientID] = []byte(clientID)

	return secret
}

func NewScopedServicePrincipalCertificateSecret(subscriptionID, tenantID, clientID, clientCert, name, namespace string) *v1.Secret {
	secret := newScopedCredentialSecret(subscriptionID, tenantID, name, namespace)

	secret.Data[config.AzureClientID] = []byte(clientID)
	secret.Data[config.AzureClientCertificate] = []byte(clientCert)

	return secret
}

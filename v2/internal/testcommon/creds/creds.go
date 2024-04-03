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
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

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

func GetCreds() (azcore.TokenCredential, AzureIDs, error) {
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
		SubscriptionID:   subscriptionID,
		TenantID:         tenantID,
		BillingInvoiceID: billingInvoiceId,
	}

	cachedCreds = creds
	cachedIds = ids
	return creds, ids, nil
}

func DummyAzureIDs() AzureIDs {
	return AzureIDs{
		SubscriptionID:   uuid.Nil.String(),
		TenantID:         uuid.Nil.String(),
		BillingInvoiceID: DummyBillingId,
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

func NewScopedServicePrincipalCertificateSecret(
	subscriptionID string,
	tenantID string,
	clientID string,
	clientCert string,
	name string,
	namespace string,
) *v1.Secret {
	secret := newScopedCredentialSecret(subscriptionID, tenantID, name, namespace)

	secret.Data[config.AzureClientID] = []byte(clientID)
	secret.Data[config.AzureClientCertificate] = []byte(clientCert)

	return secret
}

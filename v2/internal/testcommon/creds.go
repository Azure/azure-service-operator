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

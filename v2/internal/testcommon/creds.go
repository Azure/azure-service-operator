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

	"github.com/Azure/azure-service-operator/v2/internal/config"
)

// this is shared between tests because
// instantiating it requires HTTP calls
var cachedCreds azcore.TokenCredential
var cachedSubID AzureIDs

type AzureIDs struct {
	subscriptionID string
	tenantID       string
}

func getCreds() (azcore.TokenCredential, AzureIDs, error) {

	if cachedCreds != nil {
		return cachedCreds, cachedSubID, nil
	}

	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, AzureIDs{}, errors.Wrapf(err, "creating default credential")
	}

	subscriptionID := os.Getenv(config.SubscriptionIDVar)
	if subscriptionID == "" {
		return nil, AzureIDs{}, errors.Errorf("required environment variable %q was not supplied", config.SubscriptionIDVar)
	}

	tenantID := os.Getenv(config.TenantIDVar)
	if tenantID == "" {
		return nil, AzureIDs{}, errors.Errorf("required environment variable %q was not supplied", config.TenantIDVar)
	}

	ids := AzureIDs{
		subscriptionID: subscriptionID,
		tenantID:       tenantID,
	}

	cachedCreds = creds
	cachedSubID = ids
	return creds, ids, nil
}

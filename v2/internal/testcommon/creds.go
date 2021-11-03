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
var cachedSubID string

func getCreds() (azcore.TokenCredential, string, error) {

	if cachedCreds != nil {
		return cachedCreds, cachedSubID, nil
	}

	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, "", errors.Wrapf(err, "creating default credential")
	}

	subscriptionID := os.Getenv(config.SubscriptionIDVar)
	if subscriptionID == "" {
		return nil, "", errors.Errorf("required environment variable %q was not supplied", config.SubscriptionIDVar)
	}

	cachedCreds = creds
	cachedSubID = subscriptionID
	return creds, subscriptionID, nil
}

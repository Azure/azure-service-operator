/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"os"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/internal/armclient"
)

// this is shared between tests because
// instantiating it requires HTTP calls
var (
	cachedAuthorizer autorest.Authorizer
	cachedSubID      string
)

func getAuthorizer() (autorest.Authorizer, string, error) {
	if cachedAuthorizer != nil {
		return cachedAuthorizer, cachedSubID, nil
	}

	authorizer, err := armclient.AuthorizerFromEnvironment()
	if err != nil {
		return nil, "", errors.Wrapf(err, "creating authorizer")
	}

	subscriptionID := os.Getenv(auth.SubscriptionID)
	if subscriptionID == "" {
		return nil, "", errors.Wrapf(err, "required environment variable %q was not supplied", auth.SubscriptionID)
	}

	cachedAuthorizer = authorizer
	cachedSubID = subscriptionID
	return authorizer, subscriptionID, nil
}

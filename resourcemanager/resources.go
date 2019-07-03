// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package resources

import (
	"context"
	"net/http"
	"net/url"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"
	"github.com/Azure/go-autorest/autorest"

	"Telstra.Dx.AzureOperator/resourcemanager/config"
	"Telstra.Dx.AzureOperator/resourcemanager/iam"
)

func getResourcesClient() resources.Client {
	resourcesClient := resources.NewClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	resourcesClient.Authorizer = a
	resourcesClient.AddToUserAgent(config.UserAgent())
	return resourcesClient
}

// WithAPIVersion returns a prepare decorator that changes the request's query for api-version
// This can be set up as a client's RequestInspector.
func WithAPIVersion(apiVersion string) autorest.PrepareDecorator {
	return func(p autorest.Preparer) autorest.Preparer {
		return autorest.PreparerFunc(func(r *http.Request) (*http.Request, error) {
			r, err := p.Prepare(r)
			if err == nil {
				v := r.URL.Query()
				d, err := url.QueryUnescape(apiVersion)
				if err != nil {
					return r, err
				}
				v.Set("api-version", d)
				r.URL.RawQuery = v.Encode()
			}
			return r, err
		})
	}
}

// GetResource gets a resource, the generic way.
// The API version parameter overrides the API version in
// the SDK, this is needed because not all resources are
// supported on all API versions.
func GetResource(ctx context.Context, resourceProvider, resourceType, resourceName, apiVersion string) (resources.GenericResource, error) {
	resourcesClient := getResourcesClient()
	resourcesClient.RequestInspector = WithAPIVersion(apiVersion)

	return resourcesClient.Get(
		ctx,
		config.GroupName(),
		resourceProvider,
		"",
		resourceType,
		resourceName,
	)
}

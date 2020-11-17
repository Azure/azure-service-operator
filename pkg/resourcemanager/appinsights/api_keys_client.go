// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package appinsights

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/appinsights/mgmt/2015-05-01/insights"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"k8s.io/apimachinery/pkg/runtime"
)

type InsightsAPIKeysClient struct {
	Creds        config.Credentials
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAPIKeyClient(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) *InsightsAPIKeysClient {
	return &InsightsAPIKeysClient{
		Creds:        creds,
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

func getApiKeysClient(creds config.Credentials) (insights.APIKeysClient, error) {
	insightsClient := insights.NewAPIKeysClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		insightsClient = insights.APIKeysClient{}
		return insights.APIKeysClient{}, err
	}

	insightsClient.Authorizer = a
	insightsClient.AddToUserAgent(config.UserAgent())

	return insightsClient, err
}

func (c *InsightsAPIKeysClient) CreateKey(ctx context.Context, resourceGroup, insightsaccount, name string, read, write, authSDK bool) (insights.ApplicationInsightsComponentAPIKey, error) {
	apiKey := insights.ApplicationInsightsComponentAPIKey{}

	client, err := getApiKeysClient(c.Creds)
	if err != nil {
		return apiKey, err
	}

	readIds := []string{
		fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/microsoft.insights/components/%s/api", c.Creds.SubscriptionID(), resourceGroup, insightsaccount),
		fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/microsoft.insights/components/%s/draft", c.Creds.SubscriptionID(), resourceGroup, insightsaccount),
		fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/microsoft.insights/components/%s/extendqueries", c.Creds.SubscriptionID(), resourceGroup, insightsaccount),
		fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/microsoft.insights/components/%s/search", c.Creds.SubscriptionID(), resourceGroup, insightsaccount),
		fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/microsoft.insights/components/%s/aggregate", c.Creds.SubscriptionID(), resourceGroup, insightsaccount),
	}

	writeIds := []string{
		fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/microsoft.insights/components/%s/annotations", c.Creds.SubscriptionID(), resourceGroup, insightsaccount),
	}

	authSDKControl := []string{fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/microsoft.insights/components/%s/agentconfig", c.Creds.SubscriptionID(), resourceGroup, insightsaccount)}

	keyprops := insights.APIKeyRequest{
		Name: &name,
	}

	if read {
		keyprops.LinkedReadProperties = &readIds
	}

	if write {
		keyprops.LinkedWriteProperties = &writeIds
	}

	if authSDK {
		if keyprops.LinkedReadProperties == nil {
			keyprops.LinkedReadProperties = &authSDKControl
		} else {
			combined := append(*keyprops.LinkedReadProperties, authSDKControl...)
			keyprops.LinkedReadProperties = &combined
		}
	}

	apiKey, err = client.Create(
		ctx,
		resourceGroup,
		insightsaccount,
		keyprops,
	)
	if err != nil {
		return apiKey, err
	}

	return apiKey, nil
}

func (c *InsightsAPIKeysClient) DeleteKey(ctx context.Context, resourceGroup, insightsaccount, name string) error {
	client, err := getApiKeysClient(c.Creds)
	if err != nil {
		return err
	}

	_, err = client.Delete(ctx, resourceGroup, insightsaccount, name)
	if err != nil {
		return err
	}
	return nil
}

func (c *InsightsAPIKeysClient) GetKey(ctx context.Context, resourceGroup, insightsaccount, name string) (insights.ApplicationInsightsComponentAPIKey, error) {
	result := insights.ApplicationInsightsComponentAPIKey{}
	client, err := getApiKeysClient(c.Creds)
	if err != nil {
		return result, err
	}

	result, err = client.Get(ctx, resourceGroup, insightsaccount, name)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (c *InsightsAPIKeysClient) ListKeys(ctx context.Context, resourceGroup, insightsaccount string) (insights.ApplicationInsightsComponentAPIKeyListResult, error) {
	result := insights.ApplicationInsightsComponentAPIKeyListResult{}
	client, err := getApiKeysClient(c.Creds)
	if err != nil {
		return result, err
	}

	result, err = client.List(ctx, resourceGroup, insightsaccount)
	if err != nil {
		return result, err
	}

	return result, nil
}

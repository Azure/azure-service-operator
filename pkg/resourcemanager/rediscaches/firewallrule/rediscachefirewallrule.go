// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package rediscachefirewallrules

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

// AzureRedisCacheFirewallRuleManager creates a new AzureRedisCacheFirewallRuleManager
type AzureRedisCacheFirewallRuleManager struct{}

// NewAzureRedisCacheFirewallRuleManager creates a new AzureRedisCacheFirewallRuleManager
func NewAzureRedisCacheFirewallRuleManager() *AzureRedisCacheFirewallRuleManager {
	return &AzureRedisCacheFirewallRuleManager{}
}

// getRedisCacheFirewallRuleClient retrieves a firewallrules client
func getRedisCacheFirewallRuleClient() (redis.FirewallRulesClient, error) {
	firewallRulesClient := redis.NewFirewallRulesClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return redis.FirewallRulesClient{}, err
	}
	firewallRulesClient.Authorizer = a
	firewallRulesClient.AddToUserAgent(config.UserAgent())
	return firewallRulesClient, nil
}

// CreateRedisCacheFirewallRule creates a new RedisCacheFirewallRule
func (r *AzureRedisCacheFirewallRuleManager) CreateRedisCacheFirewallRule(ctx context.Context, instance azurev1alpha1.RedisCacheFirewallRule) (result redis.FirewallRule, err error) {

	firewallRuleClient, err := getRedisCacheFirewallRuleClient()
	if err != nil {
		return redis.FirewallRule{}, err
	}

	resourceGroup := instance.Spec.resourceGroup
	redisCacheName := instance.Spec.CacheName
	firewallRuleName := instance.ObjectMeta.Name

	firewallRuleParameters := redis.FirewallRuleCreateParameters{
		FirewallRuleProperties: &redis.FirewallRuleProperties{
			StartIP: to.StringPtr(instance.Spec.Properties.StartIP),
			EndIP:   to.StringPtr(instance.Spec.Properties.EndIP),
		},
	}
	future, err := firewallRuleClient.CreateOrUpdate(
		ctx,
		resourceGroup,
		redisCacheName,
		firewallRuleName,
		firewallRuleParameters,
	)
	if err != nil {
		return redis.FirewallRule{}, err
	}

	return future, err
}

// Get gets a single firewall rule in a specified redis cache
func (r *AzureRedisCacheFirewallRuleManager) Get(ctx context.Context, resourceGroup string, redisCacheName string, firewallRuleName string) (result redis.FirewallRule, err error) {

	firewallRuleClient, err := getRedisCacheFirewallRuleClient()
	if err != nil {
		return redis.FirewallRule{}, err
	}

	return firewallRuleClient.Get(ctx, resourceGroup, redisCacheName, firewallRuleName)
}

// DeleteRedisCacheFirewallRule deletes a redis firewall rule
func (r *AzureRedisCacheFirewallRuleManager) DeleteRedisCacheFirewallRule(ctx context.Context, resourceGroup string, redisCacheName string, firewallRuleName string) (result autorest.Response, err error) {
	result = autorest.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}
	firewallRuleClient, err := getRedisCacheFirewallRuleClient()
	if err != nil {
		return result, err
	}
	future, err := firewallRuleClient.Delete(ctx, resourceGroup, redisCacheName, firewallRuleName)
	if err != nil {
		return result, nil
	}
	return future, err
}

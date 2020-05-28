// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package rediscachefirewallrules

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/go-autorest/autorest"
)

// RedisCacheFirewallRuleManager for RedisCacheFirewallRule
type RedisCacheFirewallRuleManager interface {
	// CreateRedisCacheFirewallRule creates a new RedisCacheFirewallRule
	CreateRedisCacheFirewallRule(ctx context.Context, instance azurev1alpha1.RedisCacheFirewallRule) (result redis.FirewallRule, err error)

	// Get gets a single firewall rule in a specified redis cache
	Get(ctx context.Context, resourceGroup string, redisCacheName string, firewallRuleName string) (err error)

	// DeleteRedisCacheFirewallRule deletes a server firewall rule
	DeleteRedisCacheFirewallRule(ctx context.Context, resourceGroup string, redisCacheName string, firewallRuleName string) (result autorest.Response, err error)

	// also embed async client methods
	resourcemanager.ARMClient
}

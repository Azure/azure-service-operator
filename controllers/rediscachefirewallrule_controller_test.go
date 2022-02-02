// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || rediscache
// +build all rediscache

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestRedisCacheFirewallRuleControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	var rgName string
	var redisCache string
	var redisCacheFirewallRule string

	rgName = GenerateTestResourceNameWithRandom("rcfwr-rg", 10)
	redisCache = GenerateTestResourceNameWithRandom("rediscache", 10)
	redisCacheFirewallRule = GenerateTestResourceNameWithRandom("rediscachefirewallrule", 10)

	redisCacheFirewallRuleInstance := &azurev1alpha1.RedisCacheFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      redisCacheFirewallRule,
			Namespace: "default",
		},
		Spec: azurev1alpha1.RedisCacheFirewallRuleSpec{
			ResourceGroup: rgName,
			CacheName:     redisCache,
			Properties: azurev1alpha1.RedisCacheFirewallRuleProperties{
				StartIP: "0.0.0.0",
				EndIP:   "0.0.0.0",
			},
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, redisCacheFirewallRuleInstance, errhelp.ResourceGroupNotFoundErrorCode, false)
	EnsureDelete(ctx, t, tc, redisCacheFirewallRuleInstance)
}

func TestRedisCacheFirewallRuleNoRedisCache(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	var rgName string
	var redisCache string
	var redisCacheFirewallRule string

	rgName = tc.resourceGroupName
	redisCache = GenerateTestResourceNameWithRandom("rediscache", 10)
	redisCacheFirewallRule = GenerateTestResourceNameWithRandom("rediscachefirewallrule", 10)

	redisCacheFirewallRuleInstance := &azurev1alpha1.RedisCacheFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      redisCacheFirewallRule,
			Namespace: "default",
		},
		Spec: azurev1alpha1.RedisCacheFirewallRuleSpec{
			ResourceGroup: rgName,
			CacheName:     redisCache,
			Properties: azurev1alpha1.RedisCacheFirewallRuleProperties{
				StartIP: "0.0.0.0",
				EndIP:   "0.0.0.0",
			},
		},
	}
	EnsureInstanceWithResult(ctx, t, tc, redisCacheFirewallRuleInstance, errhelp.ResourceNotFound, false)
	EnsureDelete(ctx, t, tc, redisCacheFirewallRuleInstance)
}

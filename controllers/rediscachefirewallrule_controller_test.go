// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all resourcegroup

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestRedisCacheFirewallRuleHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	var rgName string
	var redisCache string
	var redisCacheFirewallRule string

	rgName := tc.ResourceGroupName
	redisCache := GenerateTestResourceNameWithRandom("redisache", 10)
	redisCacheFirewallRule := GenerateTestResourceNameWithRandom("rediscachefirewallrule", 10)

	redisCacheFirewallRuleInstance := &azurev1alpha1.RedisCacheFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      redisCacheFirewallRule,
			Namespace: "default",
		},
		Spec: azurev1alpha1.RedisCacheFirewallRuleSpec{
			ResourceGroupName: rgName,
			RedisCache:        redisCache,
			Properties: azurev1alpha1.RedisCacheFirewallRuleProperties{
				StartIP: "0.0.0.0",
				EndIP:   "0.0.0.0",
			},
		},
	}

	// create rcfwr
	EnsureInstance(ctx, t, tc, redisCacheFirewallRuleInstance)

	// delete rcfwr
	EnsureDelete(ctx, t, tc, redisCacheFirewallRuleInstance)
}

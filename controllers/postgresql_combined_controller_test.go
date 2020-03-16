// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all psql psqldatabase

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestPSQLDatabaseController(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	var rgName string
	var rgLocation string
	var postgreSQLServerName string
	var postgreSQLServerInstance *azurev1alpha1.PostgreSQLServer
	var postgreSQLServerNamespacedName types.NamespacedName
	var err error

	// Add any setup steps that needs to be executed before each test
	rgName = tc.resourceGroupName
	rgLocation = tc.resourceGroupLocation

	postgreSQLServerName = GenerateTestResourceNameWithRandom("psql-srv", 10)

	// Create the PostgreSQLServer object and expect the Reconcile to be created
	postgreSQLServerInstance = &azurev1alpha1.PostgreSQLServer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      postgreSQLServerName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.PostgreSQLServerSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			Sku: azurev1alpha1.PSQLSku{
				Name:     "B_Gen5_2",
				Tier:     azurev1alpha1.SkuTier("Basic"),
				Family:   "Gen5",
				Size:     "51200",
				Capacity: 2,
			},
			ServerVersion:  azurev1alpha1.ServerVersion("10"),
			SSLEnforcement: azurev1alpha1.SslEnforcementEnumEnabled,
		},
	}

	err = tc.k8sClient.Create(ctx, postgreSQLServerInstance)
	assert.Equal(nil, err, "create postgreSQLServerInstance in k8s")

	postgreSQLServerNamespacedName = types.NamespacedName{Name: postgreSQLServerName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, postgreSQLServerNamespacedName, postgreSQLServerInstance)
		return HasFinalizer(postgreSQLServerInstance, finalizerName)
	}, tc.timeout, tc.retry, "wait for postgreSQLserver to have finlizer")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, postgreSQLServerNamespacedName, postgreSQLServerInstance)
		return postgreSQLServerInstance.Status.Provisioned
	}, tc.timeout, tc.retry, "wait for postgreSQLserver to be provisioned")

	postgreSQLDatabaseName := GenerateTestResourceNameWithRandom("psql-db", 10)

	// Create the PostgreSQLDatabase object and expect the Reconcile to be created
	postgreSQLDatabaseInstance := &azurev1alpha1.PostgreSQLDatabase{
		ObjectMeta: metav1.ObjectMeta{
			Name:      postgreSQLDatabaseName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.PostgreSQLDatabaseSpec{
			ResourceGroup: rgName,
			Server:        postgreSQLServerName,
		},
	}

	err = tc.k8sClient.Create(ctx, postgreSQLDatabaseInstance)
	assert.Equal(nil, err, "create postgreSQLDatabaseInstance in k8s")

	postgreSQLDatabaseNamespacedName := types.NamespacedName{Name: postgreSQLDatabaseName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, postgreSQLDatabaseNamespacedName, postgreSQLDatabaseInstance)
		return HasFinalizer(postgreSQLDatabaseInstance, finalizerName)
	}, tc.timeout, tc.retry, "wait for postgreSQLDBInstance to have finalizer")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, postgreSQLDatabaseNamespacedName, postgreSQLDatabaseInstance)
		return postgreSQLDatabaseInstance.Status.Provisioned
	}, tc.timeout, tc.retry, "wait for postgreSQLDBInstance to be provisioned")

	err = tc.k8sClient.Delete(ctx, postgreSQLDatabaseInstance)
	assert.Equal(nil, err, "delete postgreSQLDatabaseInstance in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, postgreSQLDatabaseNamespacedName, postgreSQLDatabaseInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for postgreSQLDBInstance to be gone from k8s")

	// Test firewall rule -------------------------------

	postgreSQLFirewallRuleName := GenerateTestResourceNameWithRandom("psql-fwrule", 10)

	// Create the PostgreSQLFirewallRule object and expect the Reconcile to be created
	postgreSQLFirewallRuleInstance := &azurev1alpha1.PostgreSQLFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      postgreSQLFirewallRuleName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.PostgreSQLFirewallRuleSpec{
			ResourceGroup:  rgName,
			Server:         postgreSQLServerName,
			StartIPAddress: "0.0.0.0",
			EndIPAddress:   "0.0.0.0",
		},
	}

	err = tc.k8sClient.Create(ctx, postgreSQLFirewallRuleInstance)
	assert.Equal(nil, err, "create postgreSQLFirewallRuleInstance in k8s")

	postgreSQLFirewallRuleNamespacedName := types.NamespacedName{Name: postgreSQLFirewallRuleName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, postgreSQLFirewallRuleNamespacedName, postgreSQLFirewallRuleInstance)
		return HasFinalizer(postgreSQLFirewallRuleInstance, finalizerName)
	}, tc.timeout, tc.retry, "wait for postgreSQLFirewallRuleInstance to have finalizer")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, postgreSQLFirewallRuleNamespacedName, postgreSQLFirewallRuleInstance)
		return postgreSQLFirewallRuleInstance.Status.Provisioned
	}, tc.timeout, tc.retry, "wait for postgreSQLFirewallRuleInstance to be provisioned")

	err = tc.k8sClient.Delete(ctx, postgreSQLFirewallRuleInstance)
	assert.Equal(nil, err, "delete postgreSQLFirewallRuleInstance in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, postgreSQLFirewallRuleNamespacedName, postgreSQLFirewallRuleInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for postgreSQLFirewallRuleInstance to be gone from k8s")

	// Add any teardown steps that needs to be executed after each test
	err = tc.k8sClient.Delete(ctx, postgreSQLServerInstance)
	assert.Equal(nil, err, "delete postgreSQLServerInstance in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, postgreSQLServerNamespacedName, postgreSQLServerInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for postgreSQLServerInstance to be gone from k8s")

}

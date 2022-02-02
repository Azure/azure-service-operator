// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || mysql
// +build all mysql

package controllers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
)

func TestMySQLServerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := GenerateTestResourceNameWithRandom("mysqlsrv-rg", 10)
	rgLocation := tc.resourceGroupLocation
	mySQLServerName := GenerateTestResourceNameWithRandom("mysql-srv", 10)

	// Create the mySQLServer object and expect the Reconcile to be created
	mySQLServerInstance := v1alpha2.NewDefaultMySQLServer(mySQLServerName, rgName, rgLocation)

	EnsureInstanceWithResult(ctx, t, tc, mySQLServerInstance, errhelp.ResourceGroupNotFoundErrorCode, false)
	EnsureDelete(ctx, t, tc, mySQLServerInstance)
}

func TestMySQLServerBadLocation(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgLocation := GenerateTestResourceNameWithRandom("mysqlsrv-rg", 10)
	rgName := tc.resourceGroupName
	mySQLServerName := GenerateTestResourceNameWithRandom("mysql-srv", 10)

	// Create the mySQLServer object and expect the Reconcile to be created
	mySQLServerInstance := v1alpha2.NewDefaultMySQLServer(mySQLServerName, rgName, rgLocation)

	EnsureInstanceWithResult(ctx, t, tc, mySQLServerInstance, errhelp.InvalidResourceLocation, false)
	EnsureDelete(ctx, t, tc, mySQLServerInstance)
}

func TestMySQLServerMissingUserSpecifiedSecret(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := require.New(t)

	// Add any setup steps that needs to be executed before each test
	mySQLServerName := GenerateTestResourceNameWithRandom("mysql-srv", 10)

	mySQLServerInstance := v1alpha2.NewDefaultMySQLServer(mySQLServerName, tc.resourceGroupName, tc.resourceGroupLocation)
	mySQLServerInstance.Spec.AdminSecret = "doesntexist"

	EnsureInstanceWithResult(ctx, t, tc, mySQLServerInstance, "Failed to get AdminSecret", false)

	// Confirm that the resource is not done reconciling
	assert.Equal(false, mySQLServerInstance.Status.FailedProvisioning)
	assert.Equal(false, mySQLServerInstance.Status.Provisioned)
	// The fact that provisioning is false is an artifact of where we set provisioning in the reconcile loop
	// and also how we're inconsistent with what exactly provisioning means in the context of these resources.
	// Changing it is a larger change though so asserting it is false for now.
	assert.Equal(false, mySQLServerInstance.Status.Provisioning)

	EnsureDelete(ctx, t, tc, mySQLServerInstance)
}

func TestMySQLServerUserSpecifiedSecretMissingPassword(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := require.New(t)

	// Add any setup steps that needs to be executed before each test
	mySQLServerName := GenerateTestResourceNameWithRandom("mysql-srv", 10)
	secretName := GenerateTestResourceNameWithRandom("mysqlserversecret", 10)
	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: "default",
		},
		StringData: map[string]string{
			"username": "testuser",
		},
	}
	err := tc.k8sClient.Create(ctx, secret)
	assert.NoError(err)

	mySQLServerInstance := v1alpha2.NewDefaultMySQLServer(mySQLServerName, tc.resourceGroupName, tc.resourceGroupLocation)
	mySQLServerInstance.Spec.AdminSecret = secretName

	EnsureInstanceWithResult(ctx, t, tc, mySQLServerInstance, "is missing required \"password\" field", false)
	EnsureDelete(ctx, t, tc, mySQLServerInstance)
}

func TestMySQLServerUserSpecifiedSecretMissingUsername(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := require.New(t)

	// Add any setup steps that needs to be executed before each test
	mySQLServerName := GenerateTestResourceNameWithRandom("mysql-srv", 10)
	secretName := GenerateTestResourceNameWithRandom("mysqlserversecret", 10)
	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: "default",
		},
		StringData: map[string]string{
			"password": "testpassword",
		},
	}
	err := tc.k8sClient.Create(ctx, secret)
	assert.NoError(err)

	mySQLServerInstance := v1alpha2.NewDefaultMySQLServer(mySQLServerName, tc.resourceGroupName, tc.resourceGroupLocation)
	mySQLServerInstance.Spec.AdminSecret = secretName

	EnsureInstanceWithResult(ctx, t, tc, mySQLServerInstance, "is missing required \"username\" field", false)
	EnsureDelete(ctx, t, tc, mySQLServerInstance)
}

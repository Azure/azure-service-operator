// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all azuresqlserver fog

package controllers

import (
	"context"
	"strings"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestAzureSqlFailoverGroupControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	var rgName string
	var rgLocation1 string
	var sqlServerOneName string
	var sqlServerTwoName string
	var sqlDatabaseName string
	var err error

	// Add any setup steps that needs to be executed before each test
	rgName = tc.resourceGroupName
	rgLocation1 = "westus2"
	sqlServerOneName = GenerateTestResourceNameWithRandom("sqlfog-srvone", 10)
	sqlServerTwoName = GenerateTestResourceNameWithRandom("sqlfog-srvtwo", 10)
	sqlDatabaseName = GenerateTestResourceNameWithRandom("sqldb", 10)
	sqlFailoverGroupName := GenerateTestResourceNameWithRandom("sqlfog-dev", 10)

	// Create the SqlFailoverGroup object and expect the Reconcile to be created
	sqlFailoverGroupInstance := &azurev1alpha1.AzureSqlFailoverGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlFailoverGroupName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSqlFailoverGroupSpec{
			Location:                     rgLocation1,
			ResourceGroup:                GenerateTestResourceNameWithRandom("rg-fake", 10),
			Server:                       sqlServerOneName,
			FailoverPolicy:               "automatic",
			FailoverGracePeriod:          30,
			SecondaryServer:              sqlServerTwoName,
			SecondaryServerResourceGroup: rgName,
			DatabaseList:                 []string{sqlDatabaseName},
		},
	}

	err = tc.k8sClient.Create(ctx, sqlFailoverGroupInstance)
	assert.Equal(nil, err, "create failovergroup in k8s")

	sqlFailoverGroupNamespacedName := types.NamespacedName{Name: sqlFailoverGroupName, Namespace: "default"}

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlFailoverGroupNamespacedName, sqlFailoverGroupInstance)
		return strings.Contains(sqlFailoverGroupInstance.Status.Message, errhelp.ResourceGroupNotFoundErrorCode)
	}, tc.timeout, tc.retry, "wait for rg not found error to clear")

	err = tc.k8sClient.Delete(ctx, sqlFailoverGroupInstance)
	assert.Equal(nil, err, "delete failovergroup in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlFailoverGroupNamespacedName, sqlFailoverGroupInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for failovergroup to be gone from k8s")
}

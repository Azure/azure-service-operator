// +build all azuresqlserver fog

package controllers

import (
	"context"
	"strings"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestAzureSqlFailoverGroupControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover()
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
	sqlServerOneName = "t-sqlfog-srvone" + helpers.RandomString(10)
	sqlServerTwoName = "t-sqlfog-srvtwo" + helpers.RandomString(10)
	sqlDatabaseName = "t-sqldb" + helpers.RandomString(10)

	randomName := helpers.RandomString(10)
	sqlFailoverGroupName := "t-sqlfog-dev-" + randomName

	// Create the SqlFailoverGroup object and expect the Reconcile to be created
	sqlFailoverGroupInstance := &azurev1alpha1.AzureSqlFailoverGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlFailoverGroupName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSqlFailoverGroupSpec{
			Location:                     rgLocation1,
			ResourceGroup:                "t-rg-fake-" + helpers.RandomString(10),
			Server:                       sqlServerOneName,
			FailoverPolicy:               "automatic",
			FailoverGracePeriod:          30,
			SecondaryServerName:          sqlServerTwoName,
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
	}, tc.timeout, tc.retry, "wait for rg not found error")

	err = tc.k8sClient.Delete(ctx, sqlFailoverGroupInstance)
	assert.Equal(nil, err, "delete failovergroup in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlFailoverGroupNamespacedName, sqlFailoverGroupInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for failovergroup to be gone from k8s")
}

// +build all azuresqlserver azuresqlservercombined testaction

package controllers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)


func TestAzureSqlTestFuncs(t *testing.T) {
	t.Parallel()
	defer PanicRecover()
	ctx := context.Background()
	assert := assert.New(t)

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	sqlServerName := "t-sqlserver-dev-" + helpers.RandomString(10)
	rgLocation := "westus2"

	sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}

	// Create the SqlServer object and expect the Reconcile to be created
	sqlServerInstance := azurev1alpha1.NewAzureSQLServer(sqlServerNamespacedName, rgName, rgLocation)
	err := tc.k8sClient.Create(ctx, sqlServerInstance)
	assert.Equal(nil, err, "create server in k8s")

	EnsureInstance(ctx, t, tc, sqlServerInstance)

}

func RunSQLActionHappy(t *testing.T) {
	//t.Parallel()
	defer PanicRecover()
	ctx := context.Background()
	assert := assert.New(t)

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	sqlServerName := "t-sqldb-test-srv" + helpers.RandomString(10)

	sqlActionName := "t-azuresqlaction-dev-" + helpers.RandomString(10)

	// Create the Sql Action object and expect the Reconcile to be created
	sqlActionInstance := &azurev1alpha1.AzureSqlAction{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlActionName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSqlActionSpec{
			ActionName:    rgLocation,
			ServerName:    sqlServerName,
			ResourceGroup: rgName,
		},
	}

	//Get SQL Database credentials to compare after rollover

	err := tc.k8sClient.Create(ctx, sqlActionInstance)
	assert.Equal(nil, err, "create sqlaction in k8s")

	sqlActionInstanceNamespacedName := types.NamespacedName{Name: sqlActionName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlActionInstanceNamespacedName, sqlActionInstance)
		return sqlActionInstance.IsSubmitted()
	}, tc.timeout, tc.retry, "wait for sql action to be submitted")

	// TODO Check SQL Database credentials

	// TODO Assert credentials are not the same as previous

}

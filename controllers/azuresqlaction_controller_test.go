// +build all azuresqlserver azuresqlaction

package controllers

import (
	"context"
	"github.com/stretchr/testify/assert"


	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestAzureSqlActionController(t *testing.T) {
	t.Parallel()
	defer PanicRecover()
	ctx := context.Background()
	assert := assert.New(t)

		// Add any setup steps that needs to be executed before each test
		rgName := tc.resourceGroupName
		rgLocation := tc.resourceGroupLocation
		sqlServerName := "t-sqldb-test-srv" + helpers.RandomString(10)

		// Create the SQL servers
		// Create the SqlServer object and expect the Reconcile to be created
		sqlServerInstance := &azurev1alpha1.AzureSqlServer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      sqlServerName,
				Namespace: "default",
			},
			Spec: azurev1alpha1.AzureSqlServerSpec{
				Location:      rgLocation,
				ResourceGroup: rgName,
			},
		}

		err := tc.k8sClient.Create(ctx, sqlServerInstance)
		assert.Equal(nil, err, "create sqlserver in k8s")

		sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}

		// Check to make sure the SQL server is provisioned before moving ahead
		assert.Eventually(func() bool {
			_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlServerInstance)
			return sqlServerInstance.Status.Provisioned
		}, tc.timeout, tc.retry, "wait for sql server to provision")
	

			sqlActionName := "t-azuresqlaction-dev-" + helpers.RandomString(10)

			var err error

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

			err = tc.k8sClient.Create(ctx, sqlActionInstance)
			assert.Equal(nil, err, "create sqlaction in k8s")


			sqlActionInstanceNamespacedName := types.NamespacedName{Name: sqlActionName, Namespace: "default"}

			assert.Eventually(func() bool {
				_ = tc.k8sClient.Get(ctx, sqlActionInstanceNamespacedName, sqlActionInstance)
				return sqlActionInstance.IsSubmitted()
			}, tc.timeout, tc.retry, "wait for sql action to be submitted")

			// TODO Check SQL Database credentials

			// TODO Assert credentials are not the same as previous

			err = tc.k8sClient.Delete(ctx, sqlActionInstance)
			assert.Equal(nil, err, "delete sql action in k8s")

			assert.Eventually(func() bool {
				_ = tc.k8sClient.Get(ctx, sqlActionInstanceNamespacedName, sqlActionInstance)
				return helpers.IsBeingDeleted(sqlActionInstance)
			}, tc.timeout, tc.retry, "wait for sql action to be submitted")

			sqlActionName := "t-azuresqlaction-dev-" + helpers.RandomString(10)
			invalidSqlServerName := "404sqlserver" + helpers.RandomString(10)

			var err error

			// Create the Sql Action object and expect the Reconcile to be created
			sqlActionInstance := &azurev1alpha1.AzureSqlAction{
				ObjectMeta: metav1.ObjectMeta{
					Name:      sqlActionName,
					Namespace: "default",
				},
				Spec: azurev1alpha1.AzureSqlActionSpec{
					ActionName:    rgLocation,
					ServerName:    invalidSqlServerName,
					ResourceGroup: rgName,
				},
			}

			err = tc.k8sClient.Create(ctx, sqlActionInstance)
			assert.Equal(nil, err, "create sqlaction in k8s")

			sqlActionInstanceNamespacedName := types.NamespacedName{Name: sqlActionName, Namespace: "default"}

			assert.Eventually(func() bool {
				_ = tc.k8sClient.Get(ctx, sqlActionInstanceNamespacedName, sqlActionInstance)
				return helpers.HasFinalizer(sqlActionInstance, AzureSQLDatabaseFinalizerName)
			}, tc.timeout, tc.retry, "wait for sql action to be submitted")


			err = tc.k8sClient.Delete(ctx, sqlActionInstance)
			assert.Equal(nil, err, "delete sql action in k8s")

		// Add any teardown steps that needs to be executed after each test
		// delete the sql servers from K8s
		sqlServerInstance := &azurev1alpha1.AzureSqlServer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      sqlServerName,
				Namespace: "default",
			},
			Spec: azurev1alpha1.AzureSqlServerSpec{
				Location:      tc.resourceGroupLocation,
				ResourceGroup: tc.resourceGroupName,
			},
		}
		sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}

		_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlServerInstance)
		_ = tc.k8sClient.Delete(ctx, sqlServerInstance)

}

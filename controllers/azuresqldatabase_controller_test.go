// +build all azuresqldatabase

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestAzureSqlDatabaseControllerHappyPath(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	PanicRecover()
	ctx := context.Background()

	var rgName string
	var rgLocation string
	var sqlServerName string
	var err error

	// Add any setup steps that needs to be executed before each test
	rgName = tc.resourceGroupName
	rgLocation = tc.resourceGroupLocation
	sqlServerName = "t-sqldb-test-srv" + helpers.RandomString(10)

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

	err = tc.k8sClient.Create(ctx, sqlServerInstance)
	Expect(err).NotTo(HaveOccurred())

	sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}

	// Check to make sure the SQL server is provisioned before moving ahead
	Eventually(func() string {
		_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlServerInstance)
		return sqlServerInstance.Status.Message
	}, tc.timeout, tc.retry,
	).Should(ContainSubstring("successfully provisioned"))

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	randomName := helpers.RandomString(10)
	sqlDatabaseName := "t-sqldatabase-dev-" + randomName

	// Create the SqlDatabase object and expect the Reconcile to be created
	sqlDatabaseInstance := &azurev1alpha1.AzureSqlDatabase{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlDatabaseName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSqlDatabaseSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			Server:        sqlServerName,
			Edition:       0,
		},
	}

	err = tc.k8sClient.Create(ctx, sqlDatabaseInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	sqlDatabaseNamespacedName := types.NamespacedName{Name: sqlDatabaseName, Namespace: "default"}

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return helpers.HasFinalizer(sqlDatabaseInstance, AzureSQLDatabaseFinalizerName)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	Eventually(func() string {
		_ = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return sqlDatabaseInstance.Status.Message
	}, tc.timeout, tc.retry,
	).Should(ContainSubstring("successfully provisioned"))

	err = tc.k8sClient.Delete(ctx, sqlDatabaseInstance)
	//Expect(err).NotTo(HaveOccurred())  //Commenting as this call is async and returns an asyncopincomplete error

	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	// Add any teardown steps that needs to be executed after each test
	// delete the sql servers from K8s
	err = tc.k8sClient.Delete(ctx, sqlServerInstance)

	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlServerInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

}

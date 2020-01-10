// +build all azuresqlserver azuresqldatabase

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestAzureSqlDatabaseControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	defer PanicRecover()
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgLocation := tc.resourceGroupLocation
	sqlServerName := "t-sqldb-test-srv" + helpers.RandomString(10)

	sqlDatabaseName := "t-sqldatabase-dev-" + helpers.RandomString(10)

	// Create the SqlDatabase object and expect the Reconcile to be created
	sqlDatabaseInstance := &azurev1alpha1.AzureSqlDatabase{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlDatabaseName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSqlDatabaseSpec{
			Location:      rgLocation,
			ResourceGroup: "t-rg-test-srv" + helpers.RandomString(10),
			Server:        sqlServerName,
			Edition:       0,
		},
	}

	err := tc.k8sClient.Create(ctx, sqlDatabaseInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	sqlDatabaseNamespacedName := types.NamespacedName{Name: sqlDatabaseName, Namespace: "default"}

	Eventually(func() string {
		_ = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return sqlDatabaseInstance.Status.Message
	}, tc.timeout, tc.retry,
	).Should(ContainSubstring(errhelp.ResourceGroupNotFoundErrorCode))

	err = tc.k8sClient.Delete(ctx, sqlDatabaseInstance)
	//Expect(err).NotTo(HaveOccurred())  //Commenting as this call is async and returns an asyncopincomplete error

	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

}

func TestAzureSqlDatabaseControllerNoServer(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	defer PanicRecover()
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	sqlServerName := "t-sqldb-test-srv" + helpers.RandomString(10)
	sqlDatabaseName := "t-sqldatabase-dev-" + helpers.RandomString(10)

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

	err := tc.k8sClient.Create(ctx, sqlDatabaseInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	sqlDatabaseNamespacedName := types.NamespacedName{Name: sqlDatabaseName, Namespace: "default"}

	Eventually(func() string {
		_ = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return sqlDatabaseInstance.Status.Message
	}, tc.timeout, tc.retry,
	).Should(ContainSubstring(errhelp.ParentNotFoundErrorCode))

	err = tc.k8sClient.Delete(ctx, sqlDatabaseInstance)
	//Expect(err).NotTo(HaveOccurred())  //Commenting as this call is async and returns an asyncopincomplete error

	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

}

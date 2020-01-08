// +build all fog

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

func TestAzureSqlFailoverGroupControllerHappyPath(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	PanicRecover()
	ctx := context.Background()

	var rgName string
	var rgLocation1 string
	var rgLocation2 string
	var sqlServerOneName string
	var sqlServerTwoName string
	var sqlDatabaseName string
	var err error

	// Add any setup steps that needs to be executed before each test
	rgName = tc.resourceGroupName
	rgLocation1 = "westus2"
	rgLocation2 = "southcentralus"
	sqlServerOneName = "t-sqlfog-srvone" + helpers.RandomString(10)
	sqlServerTwoName = "t-sqlfog-srvtwo" + helpers.RandomString(10)
	sqlDatabaseName = "t-sqldb" + helpers.RandomString(10)

	// Create the SQL servers
	// Create the first SqlServer object and expect the Reconcile to be created
	sqlServerInstance1 := &azurev1alpha1.AzureSqlServer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlServerOneName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSqlServerSpec{
			Location:      rgLocation1,
			ResourceGroup: rgName,
		},
	}

	err = tc.k8sClient.Create(ctx, sqlServerInstance1)
	Expect(err).NotTo(HaveOccurred())

	// Create the second SqlServer object and expect the Reconcile to be created
	sqlServerInstance2 := &azurev1alpha1.AzureSqlServer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlServerTwoName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSqlServerSpec{
			Location:      rgLocation2,
			ResourceGroup: rgName,
		},
	}

	err = tc.k8sClient.Create(ctx, sqlServerInstance2)
	Expect(err).NotTo(HaveOccurred())

	sqlServerNamespacedName1 := types.NamespacedName{Name: sqlServerOneName, Namespace: "default"}
	// Check to make sure the SQL server is provisioned before moving ahead
	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName1, sqlServerInstance1)
		return sqlServerInstance1.Status.Provisioned
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	sqlServerNamespacedName2 := types.NamespacedName{Name: sqlServerTwoName, Namespace: "default"}
	// Check to make sure the SQL server is provisioned before moving ahead
	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName2, sqlServerInstance2)
		return sqlServerInstance2.Status.Provisioned
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	//Create the SQL database on the first SQL server
	sqlDatabaseInstance := &azurev1alpha1.AzureSqlDatabase{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlDatabaseName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSqlDatabaseSpec{
			Location:      rgLocation1,
			ResourceGroup: rgName,
			Server:        sqlServerOneName,
			Edition:       0,
		},
	}

	err = tc.k8sClient.Create(ctx, sqlDatabaseInstance)

	sqlDatabaseNamespacedName := types.NamespacedName{Name: sqlDatabaseName, Namespace: "default"}

	// Check to make sure the SQL database is provisioned before moving ahead
	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return sqlDatabaseInstance.Status.Provisioned
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

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
			ResourceGroup:                rgName,
			Server:                       sqlServerOneName,
			FailoverPolicy:               "automatic",
			FailoverGracePeriod:          30,
			SecondaryServerName:          sqlServerTwoName,
			SecondaryServerResourceGroup: rgName,
			DatabaseList:                 []string{sqlDatabaseName},
		},
	}

	err = tc.k8sClient.Create(context.Background(), sqlFailoverGroupInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	sqlFailoverGroupNamespacedName := types.NamespacedName{Name: sqlFailoverGroupName, Namespace: "default"}

	Eventually(func() bool {
		_ = tc.k8sClient.Get(context.Background(), sqlFailoverGroupNamespacedName, sqlFailoverGroupInstance)
		return helpers.HasFinalizer(sqlFailoverGroupInstance, azureSQLFailoverGroupFinalizerName)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	Eventually(func() bool {
		_ = tc.k8sClient.Get(context.Background(), sqlFailoverGroupNamespacedName, sqlFailoverGroupInstance)

		return sqlFailoverGroupInstance.Status.Provisioned
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	err = tc.k8sClient.Delete(context.Background(), sqlFailoverGroupInstance)

	Eventually(func() bool {
		_ = tc.k8sClient.Get(context.Background(), sqlFailoverGroupNamespacedName, sqlFailoverGroupInstance)
		return helpers.IsBeingDeleted(sqlFailoverGroupInstance)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	//AfterEach
	// Add any teardown steps that needs to be executed after each test

	err = tc.k8sClient.Delete(context.Background(), sqlDatabaseInstance)

	Eventually(func() bool {
		_ = tc.k8sClient.Get(context.Background(), sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return helpers.IsBeingDeleted(sqlDatabaseInstance)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	// delete the sql servers from K8s.
	//_ = tc.k8sClient.Get(context.Background(), sqlServerNamespacedName, sqlServerInstance1)
	err = tc.k8sClient.Delete(context.Background(), sqlServerInstance1)

	Eventually(func() bool {
		_ = tc.k8sClient.Get(context.Background(), sqlServerNamespacedName1, sqlServerInstance1)
		return helpers.IsBeingDeleted(sqlServerInstance1)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	// Delete the SQL server two
	err = tc.k8sClient.Delete(context.Background(), sqlServerInstance2)

	Eventually(func() bool {
		_ = tc.k8sClient.Get(context.Background(), sqlServerNamespacedName2, sqlServerInstance2)
		return helpers.IsBeingDeleted(sqlServerInstance2)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

}

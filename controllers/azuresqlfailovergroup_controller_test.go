// +build all azuresqlserver fog

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

func TestAzureSqlFailoverGroupControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	PanicRecover()
	ctx := context.Background()

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
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	sqlFailoverGroupNamespacedName := types.NamespacedName{Name: sqlFailoverGroupName, Namespace: "default"}

	Eventually(func() string {
		_ = tc.k8sClient.Get(ctx, sqlFailoverGroupNamespacedName, sqlFailoverGroupInstance)
		return sqlFailoverGroupInstance.Status.Message
	}, tc.timeout, tc.retry,
	).Should(ContainSubstring(errhelp.ResourceGroupNotFoundErrorCode))

	err = tc.k8sClient.Delete(ctx, sqlFailoverGroupInstance)

	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlFailoverGroupNamespacedName, sqlFailoverGroupInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())
}

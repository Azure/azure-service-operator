// +build all psql psqlserver

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestPSQLServerController(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	PanicRecover()
	ctx := context.Background()

	var rgName string
	var rgLocation string
	var postgreSQLServerName string

	// Add any setup steps that needs to be executed before each test
	rgName = tc.resourceGroupName
	rgLocation = tc.resourceGroupLocation

	postgreSQLServerName = "t-psql-srv-" + helpers.RandomString(10)

	// Create the PostgreSQLServer object and expect the Reconcile to be created
	postgreSQLServerInstance := &azurev1alpha1.PostgreSQLServer{
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

	err := tc.k8sClient.Create(ctx, postgreSQLServerInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	postgreSQLServerNamespacedName := types.NamespacedName{Name: postgreSQLServerName, Namespace: "default"}

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, postgreSQLServerNamespacedName, postgreSQLServerInstance)
		return helpers.HasFinalizer(postgreSQLServerInstance, finalizerName)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, postgreSQLServerNamespacedName, postgreSQLServerInstance)
		return postgreSQLServerInstance.Status.Provisioned
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	err = tc.k8sClient.Delete(ctx, postgreSQLServerInstance)
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, postgreSQLServerNamespacedName, postgreSQLServerInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

}

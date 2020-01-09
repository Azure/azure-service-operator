// +build all eventhub eventhubnamespace

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

func TestEventHubNamespaceControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	PanicRecover()
	ctx := context.Background()

	var rgLocation string
	rgLocation = tc.resourceGroupLocation

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	// setting this rg name tells the mocks to set a proper error
	resourceGroupName := "gone"
	eventhubNamespaceName := "t-ns-dev-eh-" + helpers.RandomString(10)

	// Create the EventHubNamespace object and expect the Reconcile to be created
	eventhubNamespaceInstance := &azurev1alpha1.EventhubNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubNamespaceName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubNamespaceSpec{
			Location:      rgLocation,
			ResourceGroup: resourceGroupName,
		},
	}

	err := tc.k8sClient.Create(ctx, eventhubNamespaceInstance)
	Expect(err).NotTo(HaveOccurred())

	eventhubNamespacedName := types.NamespacedName{Name: eventhubNamespaceName, Namespace: "default"}

	Eventually(func() string {
		_ = tc.k8sClient.Get(ctx, eventhubNamespacedName, eventhubNamespaceInstance)
		return eventhubNamespaceInstance.Status.Message
	}, tc.timeout, tc.retry,
	).Should(ContainSubstring("ResourceGroupNotFound"))

	err = tc.k8sClient.Delete(ctx, eventhubNamespaceInstance)
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, eventhubNamespacedName, eventhubNamespaceInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

}

func TestEventHubNamespaceControllerHappy(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	PanicRecover()
	ctx := context.Background()

	var rgName string = tc.resourceGroupName
	var rgLocation string = tc.resourceGroupLocation
	eventhubNamespaceName := "t-ns-dev-eh-" + helpers.RandomString(10)

	// Create the Eventhub namespace object and expect the Reconcile to be created
	eventhubNamespaceInstance := &azurev1alpha1.EventhubNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubNamespaceName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubNamespaceSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
		},
	}

	err := tc.k8sClient.Create(ctx, eventhubNamespaceInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	eventhubNamespacedName := types.NamespacedName{Name: eventhubNamespaceName, Namespace: "default"}

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, eventhubNamespacedName, eventhubNamespaceInstance)
		return eventhubNamespaceInstance.HasFinalizer(finalizerName)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, eventhubNamespacedName, eventhubNamespaceInstance)
		return eventhubNamespaceInstance.Status.Provisioned
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	err = tc.k8sClient.Delete(ctx, eventhubNamespaceInstance)
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, eventhubNamespacedName, eventhubNamespaceInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

}

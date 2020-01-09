// +build all blobcontainer

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	s "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestBlobContainerControlleNoResourceGroup(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	PanicRecover()
	ctx := context.Background()

	var rgLocation string
	var saName string
	var containerAccessLevel s.PublicAccess

	// Add any setup steps that needs to be executed before each test
	rgLocation = tc.resourceGroupLocation
	saName = tc.storageAccountName
	containerAccessLevel = s.PublicAccessContainer

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	blobContainerName := "bc-" + helpers.RandomString(10)
	resourceGroupName := "rg-" + helpers.RandomString(10)

	var err error

	blobContainerInstance := &azurev1alpha1.BlobContainer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      blobContainerName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.BlobContainerSpec{
			Location:      rgLocation,
			ResourceGroup: resourceGroupName,
			AccountName:   saName,
			AccessLevel:   containerAccessLevel,
		},
	}

	err = tc.k8sClient.Create(ctx, blobContainerInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	blobContainerNamespacedName := types.NamespacedName{Name: blobContainerName, Namespace: "default"}
	Eventually(func() string {
		_ = tc.k8sClient.Get(ctx, blobContainerNamespacedName, blobContainerInstance)
		// @todo check the content of MEssage instead, this check returns a value that defaults to false
		return blobContainerInstance.Status.Message
	}, tc.timeout, tc.retry,
	).Should(ContainSubstring("ResourceGroupNotFound"))

	err = tc.k8sClient.Delete(ctx, blobContainerInstance)
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, blobContainerNamespacedName, blobContainerInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

}

func TestTestBlobContainerControllerNoStorageAccount(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	PanicRecover()
	ctx := context.Background()

	var rgLocation string
	var rgName string
	var containerAccessLevel s.PublicAccess

	// Add any setup steps that needs to be executed before each test
	rgLocation = tc.resourceGroupLocation
	rgName = tc.resourceGroupName
	containerAccessLevel = s.PublicAccessContainer

	blobContainerName := "bc-" + helpers.RandomString(10)
	storageAccountName := "sa-" + helpers.RandomString(10)

	var err error

	blobContainerInstance := &azurev1alpha1.BlobContainer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      blobContainerName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.BlobContainerSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			AccountName:   storageAccountName,
			AccessLevel:   containerAccessLevel,
		},
	}

	err = tc.k8sClient.Create(ctx, blobContainerInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	blobContainerNamespacedName := types.NamespacedName{Name: blobContainerName, Namespace: "default"}
	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, blobContainerNamespacedName, blobContainerInstance)
		//@todo check Message instead
		return blobContainerInstance.Status.Provisioned
	}, tc.timeout, tc.retry,
	).Should(BeFalse())

	err = tc.k8sClient.Delete(ctx, blobContainerInstance)
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, blobContainerNamespacedName, blobContainerInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

}

func TestTestBlobContainerControllerHappyPath(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	PanicRecover()
	ctx := context.Background()

	var rgLocation string
	var rgName string
	var saName string
	var containerAccessLevel s.PublicAccess

	// Add any setup steps that needs to be executed before each test
	rgLocation = tc.resourceGroupLocation
	rgName = tc.resourceGroupName
	saName = tc.storageAccountName
	containerAccessLevel = s.PublicAccessContainer

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	blobContainerName := "bc-" + helpers.RandomString(10)

	var err error

	blobContainerInstance := &azurev1alpha1.BlobContainer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      blobContainerName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.BlobContainerSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			AccountName:   saName,
			AccessLevel:   containerAccessLevel,
		},
	}

	err = tc.k8sClient.Create(ctx, blobContainerInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	blobContainerNamespacedName := types.NamespacedName{Name: blobContainerName, Namespace: "default"}

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, blobContainerNamespacedName, blobContainerInstance)
		return blobContainerInstance.HasFinalizer(blobContainerFinalizerName)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, blobContainerNamespacedName, blobContainerInstance)
		return blobContainerInstance.Status.Provisioned
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	err = tc.k8sClient.Delete(ctx, blobContainerInstance)
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, blobContainerNamespacedName, blobContainerInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

}

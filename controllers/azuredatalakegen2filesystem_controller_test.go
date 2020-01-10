// +build all adlsgen2

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

func TestADLSFilesystemControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	defer PanicRecover()

	ctx := context.Background()
	var saName string = tc.storageAccountName

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	fileSystemName := "adls-filesystem-" + helpers.RandomString(10)
	resouceGroupName := "rg-" + helpers.RandomString(10)

	var err error

	fileSystemInstance := &azurev1alpha1.AzureDataLakeGen2FileSystem{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fileSystemName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureDataLakeGen2FileSystemSpec{
			StorageAccountName: saName,
			ResourceGroupName:  resouceGroupName,
		},
	}

	err = tc.k8sClient.Create(ctx, fileSystemInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	// @todo update this to check something other than status.Provisioned
	fileSystemNamespacedName := types.NamespacedName{Name: fileSystemName, Namespace: "default"}
	Eventually(func() bool {
		_ = tc.k8sClient.Get(context.Background(), fileSystemNamespacedName, fileSystemInstance)
		return fileSystemInstance.Status.Provisioned
	}, tc.timeout, tc.retry,
	).Should(BeFalse())

	// Delete should still appear successful
	err = tc.k8sClient.Delete(ctx, fileSystemInstance)
	Expect(err).NotTo(HaveOccurred())
}

func TestADLSFilesystemControllerNoStorageAccount(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	defer PanicRecover()

	var rgName string = tc.resourceGroupName
	fileSystemName := "adls-filesystem-" + helpers.RandomString(10)
	storageAccountName := "sa-" + helpers.RandomString(10)

	var err error

	fileSystemInstance := &azurev1alpha1.AzureDataLakeGen2FileSystem{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fileSystemName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureDataLakeGen2FileSystemSpec{
			StorageAccountName: storageAccountName,
			ResourceGroupName:  rgName,
		},
	}

	err = tc.k8sClient.Create(context.Background(), fileSystemInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	fileSystemNamespacedName := types.NamespacedName{Name: fileSystemName, Namespace: "default"}
	Eventually(func() bool {
		_ = tc.k8sClient.Get(context.Background(), fileSystemNamespacedName, fileSystemInstance)
		// @todo this should check the content of Status.Message instead as the default value of Provisioned is actually False
		// so this is unreliable
		return fileSystemInstance.Status.Provisioned
	}, tc.timeout, tc.retry,
	).Should(BeFalse())

	// Delete should still appear successful
	err = tc.k8sClient.Delete(context.Background(), fileSystemInstance)
	Expect(err).NotTo(HaveOccurred())
}

// func TestADLSFilesystemControllerHappyPath(t *testing.T) {
// 	t.Parallel()
// 	RegisterTestingT(t)
// 	defer PanicRecover()

// 	var rgName string = tc.resourceGroupName
// 	var saName string = tc.storageAccountName

// 	fileSystemName := "adls-filesystem-" + helpers.RandomString(10)

// 	var err error

// 	fileSystemInstance := &azurev1alpha1.AzureDataLakeGen2FileSystem{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      fileSystemName,
// 			Namespace: "default",
// 		},
// 		Spec: azurev1alpha1.AzureDataLakeGen2FileSystemSpec{
// 			StorageAccountName: saName,
// 			ResourceGroupName:  rgName,
// 		},
// 	}

// 	err = tc.k8sClient.Create(context.Background(), fileSystemInstance)
// 	Expect(apierrors.IsInvalid(err)).To(Equal(false))
// 	Expect(err).NotTo(HaveOccurred())

// 	fileSystemNamespacedName := types.NamespacedName{Name: fileSystemName, Namespace: "default"}

// 	Eventually(func() bool {
// 		_ = tc.k8sClient.Get(context.Background(), fileSystemNamespacedName, fileSystemInstance)
// 		return fileSystemInstance.HasFinalizer(fileSystemFinalizerName)
// 	}, tc.timeout, tc.retry,
// 	).Should(BeTrue())

// 	Eventually(func() bool {
// 		_ = tc.k8sClient.Get(context.Background(), fileSystemNamespacedName, fileSystemInstance)
// 		// return fileSystemInstance.Status.Provisioned
// 		// this is wrong @todo fix
// 		return fileSystemInstance.IsSubmitted()
// 	}, tc.timeout, tc.retry,
// 	).Should(BeTrue())

// 	err = tc.k8sClient.Delete(context.Background(), fileSystemInstance)
// 	Expect(err).NotTo(HaveOccurred())

// 	Eventually(func() bool {
// 		err = tc.k8sClient.Get(context.Background(), fileSystemNamespacedName, fileSystemInstance)
// 		//return apierrors.IsNotFound(err)
// 		return fileSystemInstance.IsBeingDeleted()
// 	}, tc.timeout, tc.retry,
// 	).Should(BeTrue())

// }

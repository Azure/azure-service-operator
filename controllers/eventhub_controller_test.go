// +build all eventhub

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

func TestEventHubControlleNoNamespace(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	PanicRecover()
	ctx := context.Background()

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	eventhubName := "t-eh-" + helpers.RandomString(10)

	// Create the EventHub object and expect the Reconcile to be created
	eventhubInstance := &azurev1alpha1.Eventhub{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubSpec{
			Location:      "westus",
			Namespace:     "t-ns-dev-eh-" + helpers.RandomString(10),
			ResourceGroup: "t-rg-dev-eh-" + helpers.RandomString(10),
			Properties: azurev1alpha1.EventhubProperties{
				MessageRetentionInDays: 7,
				PartitionCount:         2,
			},
		},
	}

	err := tc.k8sClient.Create(ctx, eventhubInstance)
	Expect(err).NotTo(HaveOccurred())

	eventhubNamespacedName := types.NamespacedName{Name: eventhubName, Namespace: "default"}

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, eventhubNamespacedName, eventhubInstance)
		return eventhubInstance.IsSubmitted()
	}, tc.timeout, tc.retry,
	).Should(BeFalse())

}

func TestEventHubControlleCeateAndDelete(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	PanicRecover()
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	ehnName := tc.eventhubNamespaceName
	eventhubName := "t-eh-" + helpers.RandomString(10)

	// Create the EventHub object and expect the Reconcile to be created
	eventhubInstance := &azurev1alpha1.Eventhub{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubSpec{
			Location:      "westus",
			Namespace:     ehnName,
			ResourceGroup: rgName,
			Properties: azurev1alpha1.EventhubProperties{
				MessageRetentionInDays: 7,
				PartitionCount:         2,
			},
			AuthorizationRule: azurev1alpha1.EventhubAuthorizationRule{
				Name:   "RootManageSharedAccessKey",
				Rights: []string{"Listen"},
			},
		},
	}

	err := tc.k8sClient.Create(ctx, eventhubInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	eventhubNamespacedName := types.NamespacedName{Name: eventhubName, Namespace: "default"}

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, eventhubNamespacedName, eventhubInstance)
		return eventhubInstance.HasFinalizer(finalizerName)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, eventhubNamespacedName, eventhubInstance)
		return eventhubInstance.IsSubmitted()
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	err = tc.k8sClient.Delete(ctx, eventhubInstance)
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, eventhubNamespacedName, eventhubInstance)
		return eventhubInstance.IsBeingDeleted()
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

}

func TestEventHubControlleCeateAndDeleteCustomSecret(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	PanicRecover()
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	ehnName := tc.eventhubNamespaceName
	eventhubName := "t-eh-" + helpers.RandomString(10)
	secretName := "secret-" + eventhubName

	// Create the EventHub object and expect the Reconcile to be created
	eventhubInstance := &azurev1alpha1.Eventhub{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventhubName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubSpec{
			Location:      rgLocation,
			Namespace:     ehnName,
			ResourceGroup: rgName,
			Properties: azurev1alpha1.EventhubProperties{
				MessageRetentionInDays: 7,
				PartitionCount:         2,
			},
			AuthorizationRule: azurev1alpha1.EventhubAuthorizationRule{
				Name:   "RootManageSharedAccessKey",
				Rights: []string{"Listen"},
			},
			SecretName: secretName,
		},
	}

	err := tc.k8sClient.Create(ctx, eventhubInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	eventhubNamespacedName := types.NamespacedName{Name: eventhubName, Namespace: "default"}

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, eventhubNamespacedName, eventhubInstance)
		return eventhubInstance.HasFinalizer(finalizerName)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, eventhubNamespacedName, eventhubInstance)
		return eventhubInstance.IsSubmitted()
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	err = tc.k8sClient.Delete(ctx, eventhubInstance)
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, eventhubNamespacedName, eventhubInstance)
		return eventhubInstance.IsBeingDeleted()
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

}

func TestEventHubControlleCeateAndDeleteCapture(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	PanicRecover()
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	ehnName := tc.eventhubNamespaceName
	saName := tc.storageAccountName
	bcName := tc.blobContainerName
	eventHubName := "t-eh-" + helpers.RandomString(10)

	// Create the EventHub object and expect the Reconcile to be created
	eventHubInstance := &azurev1alpha1.Eventhub{
		ObjectMeta: metav1.ObjectMeta{
			Name:      eventHubName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.EventhubSpec{
			Location:      rgLocation,
			Namespace:     ehnName,
			ResourceGroup: rgName,
			Properties: azurev1alpha1.EventhubProperties{
				MessageRetentionInDays: 7,
				PartitionCount:         2,
				CaptureDescription: azurev1alpha1.CaptureDescription{
					Destination: azurev1alpha1.Destination{
						ArchiveNameFormat: "{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}",
						BlobContainer:     bcName,
						Name:              "EventHubArchive.AzureBlockBlob",
						StorageAccount: azurev1alpha1.StorageAccount{
							ResourceGroup: rgName,
							AccountName:   saName,
						},
					},
					Enabled:           true,
					SizeLimitInBytes:  524288000,
					IntervalInSeconds: 300,
				},
			},
		},
	}

	err := tc.k8sClient.Create(ctx, eventHubInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	eventHubNamespacedName := types.NamespacedName{Name: eventHubName, Namespace: "default"}

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, eventHubNamespacedName, eventHubInstance)
		return eventHubInstance.HasFinalizer(finalizerName)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, eventHubNamespacedName, eventHubInstance)
		return eventHubInstance.IsSubmitted()
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	Eventually(func() bool {
		hub, _ := tc.eventhubClient.GetHub(ctx, rgName, ehnName, eventHubName)
		if hub.Properties == nil || hub.CaptureDescription == nil || hub.CaptureDescription.Enabled == nil {
			return false
		}
		return *hub.CaptureDescription.Enabled
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	err = tc.k8sClient.Delete(ctx, eventHubInstance)
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, eventHubNamespacedName, eventHubInstance)
		return eventHubInstance.IsBeingDeleted()
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

}

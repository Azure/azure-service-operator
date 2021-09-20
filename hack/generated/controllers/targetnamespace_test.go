// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers_test

import (
	"context"
	"os"
	"testing"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	resources "github.com/Azure/azure-service-operator/hack/generated/apis/microsoft.resources/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/hack/generated/controllers"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/reconcilers"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/testcommon"

	. "github.com/onsi/gomega"
)

func TestTargetNamespaces(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	createNamespaces(tc, "watched", "unwatched")
	configuredNamespaces := os.Getenv("AZURE_TARGET_NAMESPACES")
	podNamespace := os.Getenv("POD_NAMESPACE")

	// We can't check for operator namespace
	tc.Expect(podNamespace).ToNot(Equal(""))

	standardSpec := resources.ResourceGroupSpec{
		Location: tc.AzureRegion,
		Tags:     testcommon.CreateTestResourceGroupDefaultTags(),
	}
	// Create resource groups in these two namespaces - we can't
	// easily use the test context for this because that's geared to
	// creating things in one test namespace, while we need fixed
	// namespaces.
	rgDefault := resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tc.Namer.GenerateName("rg"),
			Namespace: "default",
		},
		Spec: standardSpec,
	}
	tc.CreateResourceGroupAndWait(&rgDefault)
	// Check that the instance is annotated with the operator namespace.
	checkNamespaceAnnotation(tc, &rgDefault, podNamespace)

	// The watched namespace is also reconciled.
	rgWatched := resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tc.Namer.GenerateName("rg"),
			Namespace: "watched",
		},
		Spec: standardSpec,
	}
	tc.CreateResourceGroupAndWait(&rgWatched)
	checkNamespaceAnnotation(tc, &rgWatched, podNamespace)

	// But the unwatched namespace isn't...
	unwatchedName := tc.Namer.GenerateName("rg")
	rgUnwatched := resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      unwatchedName,
			Namespace: "unwatched",
		},
		Spec: standardSpec,
	}
	_, err := tc.CreateResourceGroup(&rgUnwatched)
	tc.Expect(err).ToNot(HaveOccurred())

	names := types.NamespacedName{Name: unwatchedName, Namespace: "unwatched"}

	gotFinalizer := func(g Gomega) bool {
		var instance resources.ResourceGroup
		ctx := context.Background()
		err := tc.KubeClient.Get(ctx, names, &instance)
		g.Expect(err).NotTo(HaveOccurred())
		res, err := meta.Accessor(&instance)
		g.Expect(err).NotTo(HaveOccurred())
		return HasFinalizer(res, finalizerName)
	}

	gotNamespaceAnnotation := func(g Gomega) bool {
		var instance resources.ResourceGroup
		ctx := context.Background()
		err := tc.KubeClient.Get(ctx, names, &instance)
		g.Expect(err).NotTo(HaveOccurred())
		res, err := meta.Accessor(&instance)
		g.Expect(err).NotTo(HaveOccurred())
		return res.GetAnnotations()[controllers.NamespaceAnnotation] == podNamespace
	}

	if configuredNamespaces == "" {
		t.Log("**** all namespaces mode")
		// The operator should be watching all namespaces.
		tc.G.Eventually(
			gotFinalizer,
			timeoutFast,
			retry,
		).Should(
			BeTrue(),
			"instance in some namespace never got a finalizer",
		)
		// And there should also be a namespace annotation.
		tc.G.Eventually(
			gotNamespaceAnnotation,
			timeoutFast,
			retry,
		).Should(
			BeTrue(),
			"instance in some namespace never got an operator namespace annotation",
		)
	} else {
		t.Log("**** restricted namespaces mode")
		// We can tell that the resource isn't being reconciled if it
		// never gets a finalizer.
		tc.G.Consistently(
			gotFinalizer,
			20*time.Second,
			time.Second,
		).Should(
			BeFalse(),
			"instance in unwatched namespace got finalizer",
		)
		// There also shouldn't be a namespace annotation.
		checkNoNamespaceAnnotation(tc, &rgUnwatched)
	}
}

func checkNamespaceAnnotation(tc testcommon.KubePerTestContext, instance metav1.Object, expected string) {
	res, err := meta.Accessor(instance)
	namespace := res.GetNamespace()
	tc.Expect(err).ToNot(HaveOccurred(), namespace)
	actual, found := res.GetAnnotations()[controllers.NamespaceAnnotation]
	tc.Expect(found).To(BeTrue(), namespace)
	tc.Expect(actual).To(Equal(expected), namespace)
}

func checkNoNamespaceAnnotation(tc testcommon.KubePerTestContext, instance metav1.Object) {
	res, err := meta.Accessor(instance)
	namespace := res.GetNamespace()
	tc.Expect(err).ToNot(HaveOccurred(), namespace)
	_, found := res.GetAnnotations()[controllers.NamespaceAnnotation]
	tc.Expect(found).To(BeFalse(), namespace)
}

// HasFinalizer accepts a metav1 object and returns true if the the
// object has the provided finalizer.
func HasFinalizer(o metav1.Object, finalizer string) bool {
	f := o.GetFinalizers()
	for _, e := range f {
		if e == finalizer {
			return true
		}
	}
	return false
}

func createNamespaces(tc testcommon.KubePerTestContext, names ...string) {
	for _, name := range names {
		err := tc.KubeClient.Create(tc.Ctx, &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
		})
		tc.Expect(err).NotTo(HaveOccurred())
	}
}

const finalizerName = reconcilers.GenericControllerFinalizer

const (
	timeoutFast = time.Minute * 3
	retry       = time.Second * 3
)

// }

// func TestOperatorNamespacePreventsReconciling(t *testing.T) {
// 	t.Parallel()
// 	defer PanicRecover(t)
// 	ctx := context.Background()

// 	// If a resource has a different operator's namespace it won't be
// 	// reconciled.
// 	notMine := v1alpha1.StorageAccount{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      "storageacct" + helpers.RandomString(6),
// 			Namespace: "default",
// 			Annotations: map[string]string{
// 				namespaceAnnotation: "hard-times",
// 			},
// 		},
// 		Spec: v1alpha1.StorageAccountSpec{
// 			Kind:          "BlobStorage",
// 			Location:      tc.resourceGroupLocation,
// 			ResourceGroup: tc.resourceGroupName,
// 			Sku: v1alpha1.StorageAccountSku{
// 				Name: "Standard_LRS",
// 			},
// 			AccessTier:             "Hot",
// 			EnableHTTPSTrafficOnly: to.BoolPtr(true),
// 		},
// 	}

// 	require := require.New(t)
// 	err := tc.k8sClient.Create(ctx, &notMine)
// 	tc.Expect(err).NotTo(HaveOccurred())
// 	defer EnsureDelete(ctx, t, tc, &notMine)

// 	names := types.NamespacedName{
// 		Name:      notMine.ObjectMeta.Name,
// 		Namespace: "default",
// 	}

// 	gotFinalizer := func() bool {
// 		var instance v1alpha1.StorageAccount
// 		err := tc.k8sClient.Get(ctx, names, &instance)
// 		tc.Expect(err).NotTo(HaveOccurred())
// 		res, err := meta.Accessor(&instance)
// 		tc.Expect(err).NotTo(HaveOccurred())
// 		return HasFinalizer(res, finalizerName)
// 	}

// 	require.Never(
// 		gotFinalizer,
// 		20*time.Second,
// 		time.Second,
// 		"instance claimed by some other operator got finalizer",
// 	)

// 	var events corev1.EventList
// 	err = tc.k8sClient.List(ctx, &events, &client.ListOptions{
// 		FieldSelector: fields.ParseSelectorOrDie("involvedObject.name=" + notMine.ObjectMeta.Name),
// 		Namespace:     "default",
// 	})
// 	tc.Expect(err).NotTo(HaveOccurred())
// 	require.Len(events.Items, 1)
// 	event := events.Items[0]
// 	require.Equal(event.Type, "Warning")
// 	require.Equal(event.Reason, "Overlap")
// 	require.Equal(event.Message, `Operators in podNamespace and "hard-times" are both configured to manage this resource`)

// 	// But an instance that I've claimed gets reconciled fine.
// 	mine := notMine
// 	mine.ObjectMeta = metav1.ObjectMeta{
// 		Name:      "storaceacct" + helpers.RandomString(6),
// 		Namespace: "default",
// 		Annotations: map[string]string{
// 			namespaceAnnotation: podNamespace,
// 		},
// 	}
// 	EnsureInstance(ctx, t, tc, &mine)
// 	EnsureDelete(ctx, t, tc, &mine)
// }

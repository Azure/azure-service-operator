// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/controllers"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

const finalizerName = controllers.GenericControllerFinalizer

func TestTargetNamespaces(t *testing.T) {
	t.Parallel()
	podNamespace := "some-operator"
	tc := globalTestContext.ForTestWithConfig(t, config.Values{
		OperatorMode:     config.OperatorModeBoth,
		PodNamespace:     podNamespace,
		TargetNamespaces: []string{"default", "watched"},
	})

	err := tc.CreateTestNamespaces("watched", "unwatched")
	tc.Expect(err).ToNot(HaveOccurred())

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
	unwatchedNamespace := "unwatched"
	rgUnwatched := resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      unwatchedName,
			Namespace: unwatchedNamespace,
		},
		Spec: standardSpec,
	}
	_, err = tc.CreateResourceGroup(&rgUnwatched)
	tc.Expect(err).ToNot(HaveOccurred())

	// We can tell that the resource isn't being reconciled if it
	// never gets a finalizer.
	checkNeverGetsFinalizer(tc, &rgUnwatched, "instance in unwatched namespace got finalizer")

	// There also shouldn't be a namespace annotation.
	checkNoNamespaceAnnotation(tc, &rgUnwatched)
}

func checkNamespaceAnnotation(tc *testcommon.KubePerTestContext, instance metav1.Object, expected string) {
	tc.T.Helper()
	res, err := meta.Accessor(instance)
	namespace := res.GetNamespace()
	tc.Expect(err).ToNot(HaveOccurred(), namespace)
	actual, found := res.GetAnnotations()[controllers.NamespaceAnnotation]
	tc.Expect(found).To(BeTrue(), namespace)
	tc.Expect(actual).To(Equal(expected), namespace)
}

func checkNoNamespaceAnnotation(tc *testcommon.KubePerTestContext, instance metav1.Object) {
	tc.T.Helper()
	res, err := meta.Accessor(instance)
	namespace := res.GetNamespace()
	tc.Expect(err).ToNot(HaveOccurred(), namespace)
	_, found := res.GetAnnotations()[controllers.NamespaceAnnotation]
	tc.Expect(found).To(BeFalse(), namespace)
}

func checkNeverGetsFinalizer(tc *testcommon.KubePerTestContext, original metav1.Object, message string) {
	tc.T.Helper()
	res, err := meta.Accessor(original)
	tc.Expect(err).ToNot(HaveOccurred())

	name := types.NamespacedName{
		Namespace: res.GetNamespace(),
		Name:      res.GetName(),
	}

	gotFinalizer := func(g Gomega) bool {
		var instance resources.ResourceGroup
		tc.GetResource(name, &instance)
		res, err := meta.Accessor(&instance)
		g.Expect(err).NotTo(HaveOccurred())
		return hasFinalizer(res, finalizerName)
	}

	tc.G.Consistently(
		gotFinalizer,
		20*time.Second,
		time.Second,
	).Should(
		BeFalse(),
		message,
	)
}

// hasFinalizer accepts a metav1 object and returns true if the object has the provided finalizer.
func hasFinalizer(o metav1.Object, finalizer string) bool {
	f := o.GetFinalizers()
	for _, e := range f {
		if e == finalizer {
			return true
		}
	}
	return false
}

func TestOperatorNamespacePreventsReconciling(t *testing.T) {
	t.Parallel()
	podNamespace := "this-operator"
	tc := globalTestContext.ForTestWithConfig(t, config.Values{
		OperatorMode: config.OperatorModeBoth,
		PodNamespace: podNamespace,
	})

	// If a resource has a different operator's namespace it won't be
	// reconciled.
	notMine := resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tc.Namer.GenerateName("rg"),
			Namespace: tc.Namespace,
			Annotations: map[string]string{
				controllers.NamespaceAnnotation: "some-other-operator",
			},
		},
		Spec: resources.ResourceGroupSpec{
			Location: tc.AzureRegion,
			Tags:     testcommon.CreateTestResourceGroupDefaultTags(),
		},
	}
	_, err := tc.CreateResourceGroup(&notMine)
	tc.Expect(err).NotTo(HaveOccurred())

	checkNeverGetsFinalizer(tc, &notMine, "instance claimed by some other operator got finalizer")

	var events corev1.EventList
	tc.ListResources(&events, &client.ListOptions{
		FieldSelector: fields.ParseSelectorOrDie("involvedObject.name=" + notMine.ObjectMeta.Name),
		Namespace:     tc.Namespace,
	})
	tc.Expect(events.Items).To(HaveLen(1))
	event := events.Items[0]
	tc.Expect(event.Type).To(Equal("Warning"))
	tc.Expect(event.Reason).To(Equal("Overlap"))
	tc.Expect(event.Message).To(Equal(
		fmt.Sprintf(`Operators in %q and "some-other-operator" are both configured to manage this resource`, podNamespace),
	))

	// But an instance that I've claimed gets reconciled fine.
	mine := resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tc.Namer.GenerateName("rg"),
			Namespace: tc.Namespace,
			Annotations: map[string]string{
				controllers.NamespaceAnnotation: podNamespace,
			},
		},
		Spec: resources.ResourceGroupSpec{
			Location: tc.AzureRegion,
			Tags:     testcommon.CreateTestResourceGroupDefaultTags(),
		},
	}
	tc.CreateResourceGroupAndWait(&mine)
}

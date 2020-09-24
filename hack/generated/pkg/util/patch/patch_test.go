/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package patch

import (
	"context"
	"testing"

	apierrors "k8s.io/apimachinery/pkg/api/errors"

	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	storage "github.com/Azure/k8s-infra/hack/generated/apis/microsoft.storage/v20190401"
)

func TestHelperUnstructuredPatch(t *testing.T) {
	g := NewWithT(t)
	ctx := context.TODO()

	obj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ResourceGroup",
			"apiVersion": "microsoft.resources.infra.azure.com/v1",
			"metadata": map[string]interface{}{
				"name":      "test-foo",
				"namespace": "default",
			},
			"status": map[string]interface{}{
				"provisioningState": "Succeeded",
			},
		},
	}

	s := runtime.NewScheme()
	fakeClient := fake.NewFakeClientWithScheme(s)
	g.Expect(fakeClient.Create(ctx, obj)).To(Succeed())

	h, err := NewHelper(obj, fakeClient)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(h.Patch(ctx, obj)).To(Succeed())

	// Make sure that the status has been preserved.
	succeeded, err := IsSucceeded(obj)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(succeeded).To(BeTrue())

	// Make sure that the object has been patched properly.
	afterObj := obj.DeepCopy()
	g.Expect(fakeClient.Get(ctx, client.ObjectKey{Namespace: "default", Name: "test-foo"}, afterObj)).To(Succeed())
}

func TestPatchNotFound(t *testing.T) {
	g := NewWithT(t)
	ctx := context.TODO()

	s := runtime.NewScheme()
	fakeClient := fake.NewFakeClientWithScheme(s)
	g.Expect(storage.AddToScheme(s)).To(Succeed())

	obj := &storage.StorageAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-group",
			Namespace: "test-namespace",
		},
	}

	h, err := NewHelper(obj, fakeClient)
	g.Expect(err).NotTo(HaveOccurred())

	// Make a modification so that there is something to patch
	obj.Spec.AzureName = "test"

	err = h.Patch(ctx, obj)
	g.Expect(err).To(HaveOccurred())
	g.Expect(apierrors.IsNotFound(err)).To(BeTrue())
}

func TestHelperPatch(t *testing.T) {

	tests := []struct {
		name    string
		before  runtime.Object
		after   runtime.Object
		wantErr bool
	}{
		{
			name: "Only remove finalizer update",
			before: &storage.StorageAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
					Finalizers: []string{
						"TestFinalizer",
					},
				},
			},
			after: &storage.StorageAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
				},
			},
		},
		{
			name: "Only status update",
			before: &storage.StorageAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
				},
			},
			after: &storage.StorageAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
				},
				Status: storage.StorageAccount_Status{
					Location: "westus",
				},
			},
		},
		{
			name: "Only spec update",
			before: &storage.StorageAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
				},
			},
			after: &storage.StorageAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
				},
				Spec: storage.StorageAccounts_Spec{
					Location: "westus",
				},
			},
		},
		{
			name: "Both spec and status update",
			before: &storage.StorageAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
				},
			},
			after: &storage.StorageAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
				},
				Spec: storage.StorageAccounts_Spec{
					Location: "westus",
				},
				Status: storage.StorageAccount_Status{
					Location: "westus",
				},
			},
		},
		{
			name: "Only add finalizer update",
			before: &storage.StorageAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
				},
			},
			after: &storage.StorageAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
					Finalizers: []string{
						"TestFinalizer",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)

			ctx := context.Background()

			s := runtime.NewScheme()
			g.Expect(storage.AddToScheme(s)).To(Succeed())
			fakeClient := fake.NewFakeClientWithScheme(s)

			beforeCopy := tt.before.DeepCopyObject()
			g.Expect(fakeClient.Create(ctx, beforeCopy)).To(Succeed())

			h, err := NewHelper(beforeCopy, fakeClient)
			g.Expect(err).NotTo(HaveOccurred())

			// Set the resource version of the "after" target to be whatever
			// "before" has -- this value is assigned automatically by Kubernetes
			// and must match for Patch to succeeded. Since our objective here
			// is to test the PatchHelper, the resource version must match (non-matching
			// resource versions are trivially "correct" because nothing is
			// supposed to happen - this is enforced by controller-runtime + Kubernetes so we
			// don't need to test it)
			beforeMeta := beforeCopy.(*storage.StorageAccount)
			afterMeta := tt.after.(*storage.StorageAccount)
			afterMeta.ResourceVersion = beforeMeta.ResourceVersion

			afterCopy := tt.after.DeepCopyObject()

			err = h.Patch(ctx, afterCopy)
			if tt.wantErr {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).NotTo(HaveOccurred())
			}

			g.Expect(afterCopy).To(Equal(tt.after))
		})
	}
}

// IsSucceeded returns true if the Status.ProvisioningState == Succeeded
func IsSucceeded(obj *unstructured.Unstructured) (bool, error) {
	ready, found, err := unstructured.NestedString(obj.Object, "status", "provisioningState")
	if err != nil {
		return false, errors.Wrapf(err, "failed to determine %v %q readiness",
			obj.GroupVersionKind(), obj.GetName())
	}
	return ready == "Succeeded" && found, nil
}

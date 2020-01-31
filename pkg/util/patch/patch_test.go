/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package patch

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/Azure/k8s-infra/apis"
	"github.com/Azure/k8s-infra/apis/microsoft.resources/v1"
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
	fakeClient := fake.NewFakeClientWithScheme(scheme.Scheme)
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

func TestHelperPatch(t *testing.T) {

	tests := []struct {
		name    string
		before  runtime.Object
		after   runtime.Object
		wantErr bool
	}{
		{
			name: "Only remove finalizer update",
			before: &v1.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
					Finalizers: []string{
						apis.AzureInfraFinalizer,
					},
				},
			},
			after: &v1.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
				},
			},
		},
		{
			name: "Only status update",
			before: &v1.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
				},
			},
			after: &v1.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
				},
				Status: v1.ResourceGroupStatus{
					ProvisioningState: "Succeeded",
				},
			},
		},
		{
			name: "Only spec update",
			before: &v1.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
				},
			},
			after: &v1.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
				},
				Spec: v1.ResourceGroupSpec{
					Location:  "westus2",
					ManagedBy: "foo",
				},
			},
		},
		{
			name: "Both spec and status update",
			before: &v1.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
				},
			},
			after: &v1.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
				},
				Spec: v1.ResourceGroupSpec{
					Location:  "westus2",
					ManagedBy: "foo",
				},
				Status: v1.ResourceGroupStatus{
					ProvisioningState: "Succeeded",
				},
			},
		},
		{
			name: "Only add finalizer update",
			before: &v1.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
				},
			},
			after: &v1.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-group",
					Namespace: "test-namespace",
					Finalizers: []string{
						apis.AzureInfraFinalizer,
					},
				},
			},
		},
		{
			name: "Only add ownerReferences to an unstructured object",
			before: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "ResourceGroup",
					"apiVersion": "microsoft.resources.infra.azure.com/v1",
					"metadata": map[string]interface{}{
						"name":      "test-group",
						"namespace": "default",
					},
				},
			},
			after: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "ResourceGroup",
					"apiVersion": "microsoft.resources.infra.azure.com/v1",
					"metadata": map[string]interface{}{
						"name":      "test-group",
						"namespace": "default",
						"ownerReferences": []interface{}{
							map[string]interface{}{
								"kind":       "TestOwner",
								"apiVersion": "test.infra.azure.com/v1",
								"name":       "test",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)

			g.Expect(v1.AddToScheme(scheme.Scheme)).To(Succeed())

			ctx := context.Background()
			fakeClient := fake.NewFakeClientWithScheme(scheme.Scheme)

			beforeCopy := tt.before.DeepCopyObject()
			g.Expect(fakeClient.Create(ctx, beforeCopy)).To(Succeed())

			h, err := NewHelper(beforeCopy, fakeClient)
			g.Expect(err).NotTo(HaveOccurred())

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

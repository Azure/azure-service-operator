/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "github.com/Azure/k8s-infra/apis/microsoft.resources/v1"
	"github.com/Azure/k8s-infra/pkg/zips"
)

type (
	ApplierMock struct {
		mock.Mock
	}
)

func (am *ApplierMock) Apply(ctx context.Context, res zips.Resource) (zips.Resource, error) {
	args := am.Called(ctx, res)
	return args.Get(0).(zips.Resource), args.Error(1)
}

func (am *ApplierMock) Delete(ctx context.Context, res zips.Resource) error {
	args := am.Called(ctx, res)
	return args.Error(0)
}

var _ = Describe("GenericReconciler", func() {
	BeforeEach(func() {})
	AfterEach(func() {})

	Context("Reconcile a generic Azure Resource", func() {
		It("should call Apply for ResourceGroup resource with finalizer", func() {

			ctx := context.Background()
			applier := new(ApplierMock)

			nn := client.ObjectKey{
				Namespace: "default",
				Name:      "foo",
			}

			instance := &v1.ResourceGroup{
				TypeMeta: metav1.TypeMeta{
					Kind:       "ResourceGroup",
					APIVersion: v1.GroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      nn.Name,
					Namespace: nn.Namespace,
				},
				Spec: v1.ResourceGroupSpec{
					Location:   "westus2",
					APIVersion: "2019-10-01",
				},
			}

			Expect(k8sClient.Create(ctx, instance)).To(Succeed())
			resBefore := zips.Resource{
				Name:              nn.Name,
				Type:              "Microsoft.Resources/resourceGroups",
				Location:          "westus2",
				APIVersion:        "2019-10-01",
				ProvisioningState: "Applying",
			}

			resAfter := zips.Resource{
				Name:              nn.Name,
				Type:              "Microsoft.Resources/resourceGroups",
				Location:          "westus2",
				APIVersion:        "2019-10-01",
				ID:                "/subscriptions/bar/providers/Microsoft.Resources/resourceGroup/foo",
				ProvisioningState: "Succeeded",
			}

			// setup the applier call with the projected resource
			applier.On("Apply", mock.Anything, resBefore).Return(resAfter, nil)
			reconciler := getReconciler(instance, mgr, k8sClient, applier)
			result, err := reconciler.Reconcile(ctrl.Request{
				NamespacedName: nn,
			})
			Expect(err).To(BeNil())
			Expect(result.RequeueAfter).To(BeZero())
			Expect(k8sClient.Get(ctx, nn, instance)).ToNot(HaveOccurred())
			Expect(instance.Status.ProvisioningState).To(Equal("Succeeded"))
			Expect(instance.ObjectMeta.Finalizers).To(ContainElement("infra.azure.com/finalizer"))
		})

		It("should delete a resource", func() {
			ctx := context.Background()
			applier := new(ApplierMock)
			nn := client.ObjectKey{
				Namespace: "default",
				Name:      "foo1",
			}

			instance := &v1.ResourceGroup{
				TypeMeta: metav1.TypeMeta{
					Kind:       "ResourceGroup",
					APIVersion: v1.GroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      nn.Name,
					Namespace: nn.Namespace,
				},
				Spec: v1.ResourceGroupSpec{
					Location:   "westus2",
					APIVersion: "2019-10-01",
				},
			}

			createResourceGroup(ctx, instance, applier)
			Expect(k8sClient.Get(ctx, nn, instance)).ToNot(HaveOccurred())
			deleteResourceGroup(ctx, instance, applier)
			Expect(k8sClient.Get(ctx, nn, instance)).To(HaveOccurred())
		})
	})
})

func createResourceGroup(ctx context.Context, obj *v1.ResourceGroup, applier *ApplierMock) {
	nn := client.ObjectKey{
		Name:      obj.ObjectMeta.Name,
		Namespace: obj.ObjectMeta.Namespace,
	}

	Expect(k8sClient.Create(ctx, obj)).To(Succeed())
	resBefore := zips.Resource{
		Name:              nn.Name,
		Type:              obj.ToResource().Type,
		Location:          obj.Spec.Location,
		APIVersion:        obj.ToResource().APIVersion,
		ProvisioningState: "Applying",
	}

	resAfter := zips.Resource{
		Type:              obj.ToResource().Type,
		Location:          obj.Spec.Location,
		APIVersion:        obj.ToResource().APIVersion,
		ID:                "/subscriptions/bar/providers/Microsoft.Resources/resourceGroup/foo",
		ProvisioningState: "Succeeded",
	}

	// setup the applier call with the projected resource
	applier.On("Apply", mock.Anything, resBefore).Return(resAfter, nil)
	reconciler := getReconciler(obj, mgr, k8sClient, applier)
	result, err := reconciler.Reconcile(ctrl.Request{
		NamespacedName: nn,
	})
	Expect(err).To(BeNil())
	Expect(result.RequeueAfter).To(BeZero())
}

func deleteResourceGroup(ctx context.Context, obj *v1.ResourceGroup, applier *ApplierMock) {
	Expect(k8sClient.Delete(ctx, obj)).ToNot(HaveOccurred())
	nn := client.ObjectKey{
		Name:      obj.ObjectMeta.Name,
		Namespace: obj.ObjectMeta.Namespace,
	}

	resBefore := zips.Resource{
		Name:              nn.Name,
		Type:              obj.ToResource().Type,
		Location:          obj.Spec.Location,
		APIVersion:        obj.ToResource().APIVersion,
		ProvisioningState: "Deleting",
		ID:                "/subscriptions/bar/providers/Microsoft.Resources/resourceGroup/foo",
	}

	applier.On("Delete", mock.Anything, resBefore).Return(nil)
	reconciler := getReconciler(obj, mgr, k8sClient, applier)
	result, err := reconciler.Reconcile(ctrl.Request{
		NamespacedName: nn,
	})
	Expect(err).To(BeNil())
	Expect(result.RequeueAfter).To(BeZero())
}

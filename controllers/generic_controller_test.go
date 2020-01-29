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

			instance := &v1.ResourceGroup{
				TypeMeta: metav1.TypeMeta{
					Kind:       "ResourceGroup",
					APIVersion: v1.GroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:       "foo",
					Namespace:  "default",
					Finalizers: []string{"infra.azure.com/finalizer"},
				},
				Spec: v1.ResourceGroupSpec{
					Location:   "westus2",
					APIVersion: "2019-10-01",
				},
			}

			Expect(k8sClient.Create(ctx, instance)).To(Succeed())
			resBefore := zips.Resource{
				Name:              "foo",
				Type:              "Microsoft.Resources/resourceGroups",
				Location:          "westus2",
				APIVersion:        "2019-10-01",
				ProvisioningState: "Applying",
			}

			resAfter := zips.Resource{
				Name:       "foo",
				Type:       "Microsoft.Resources/resourceGroups",
				Location:   "westus2",
				APIVersion: "2019-10-01",
				ID:         "/subscriptions/bar/providers/Microsoft.Resources/resourceGroup/foo",
			}

			// setup the applier call with the projected resource
			applier.On("Apply", mock.Anything, resBefore).Return(resAfter, nil)

			reconciler := getReconciler(instance, mgr, k8sClient, applier)
			result, err := reconciler.Reconcile(ctrl.Request{
				NamespacedName: client.ObjectKey{
					Namespace: instance.Namespace,
					Name:      instance.Name,
				},
			})
			Expect(err).To(BeNil())
			Expect(result.RequeueAfter).To(BeZero())
		})
	})
})

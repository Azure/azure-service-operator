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
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/Azure/k8s-infra/api/azmetav1"
	protov1 "github.com/Azure/k8s-infra/api/v1alpha1"
	"github.com/Azure/k8s-infra/pkg/zips"
)

type (
	ApplierMock struct {
		mock.Mock
	}
)

func (am *ApplierMock) Apply(ctx context.Context, res zips.Resource) error {
	args := am.Called(ctx, res)
	return args.Error(0)
}

func (am *ApplierMock) Delete(ctx context.Context, resID string) error {
	args := am.Called(ctx, resID)
	return args.Error(0)
}

var _ = Describe("GenericReconciler", func() {
	BeforeEach(func() {})
	AfterEach(func() {})

	Context("Reconcile a generic Azure Resource", func() {
		It("should call Apply for ResourceGroup resource with finalizer", func() {

			ctx := context.Background()
			applier := new(ApplierMock)
			scheme, err := protov1.SchemeBuilder.Build()
			Expect(err).To(BeNil())
			reconciler := &GenericReconciler{
				Client:  k8sClient,
				Log:     log.Log,
				Applier: applier,
				Scheme:  scheme,
			}

			instance := &protov1.ResourceGroup{
				TypeMeta: metav1.TypeMeta{
					Kind:       "ResourceGroup",
					APIVersion: protov1.GroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: protov1.ResourceGroupSpec{
					TrackedResourceSpec: azmetav1.TrackedResourceSpec{
						Location: "westus2",
					},
				},
			}
			AddFinalizer(instance, "infra.azure.com/finalizer")
			Expect(k8sClient.Create(ctx, instance)).To(Succeed())

			// setup the applier call with the projected resource
			applier.On("Apply", mock.Anything, zips.Resource{
				Name:       "foo",
				Type:       "Microsoft.Resources/resourceGroup",
				Location:   "westus2",
				APIVersion: "2018-05-01",
			}).Return(nil)

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

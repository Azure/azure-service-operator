/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"

	azcorev1 "github.com/Azure/k8s-infra/apis/core/v1"
	microsoftnetworkv1 "github.com/Azure/k8s-infra/apis/microsoft.network/v1"
	microsoftresourcesv1 "github.com/Azure/k8s-infra/apis/microsoft.resources/v1"
	"github.com/Azure/k8s-infra/internal/test"
	"github.com/Azure/k8s-infra/pkg/util/ownerutil"
	"github.com/Azure/k8s-infra/pkg/xform"
	"github.com/Azure/k8s-infra/pkg/zips"
)

type (
	ApplierMock struct {
		mock.Mock
	}
)

func (am *ApplierMock) Apply(ctx context.Context, res *zips.Resource) (*zips.Resource, error) {
	args := am.Called(ctx, res)
	return args.Get(0).(*zips.Resource), args.Error(1)
}

func (am *ApplierMock) DeleteApply(ctx context.Context, deploymentID string) error {
	args := am.Called(ctx, deploymentID)
	return args.Error(0)
}

func (am *ApplierMock) BeginDelete(ctx context.Context, res *zips.Resource) (*zips.Resource, error) {
	args := am.Called(ctx, res)
	return args.Get(0).(*zips.Resource), args.Error(1)
}

func (am *ApplierMock) GetResource(ctx context.Context, res *zips.Resource) (*zips.Resource, error) {
	args := am.Called(ctx, res)
	return args.Get(0).(*zips.Resource), args.Error(1)
}

func (am *ApplierMock) HeadResource(ctx context.Context, res *zips.Resource) (bool, error) {
	args := am.Called(ctx, res)
	return args.Bool(0), args.Error(1)
}

var _ = Describe("GenericReconciler", func() {
	BeforeEach(func() {})
	AfterEach(func() {})

	Context("Reconcile a generic Azure Resource", func() {
		It("should call Apply for ResourceGroup resource with finalizer", func() {

			ctx := context.Background()
			applier := new(ApplierMock)
			randomName := test.RandomName("foo", 10)

			nn := client.ObjectKey{
				Namespace: "default",
				Name:      randomName,
			}

			instance := &microsoftresourcesv1.ResourceGroup{
				TypeMeta: metav1.TypeMeta{
					Kind:       "ResourceGroup",
					APIVersion: microsoftresourcesv1.GroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      nn.Name,
					Namespace: nn.Namespace,
				},
				Spec: microsoftresourcesv1.ResourceGroupSpec{
					Location:   "westus2",
					APIVersion: "2019-10-01",
				},
			}

			Expect(k8sClient.Create(ctx, instance)).To(Succeed())
			resBefore := zips.Resource{
				Name:       nn.Name,
				Type:       "Microsoft.Resources/resourceGroups",
				Location:   "westus2",
				APIVersion: "2019-10-01",
			}

			resAfter := zips.Resource{
				Name:              nn.Name,
				Type:              "Microsoft.Resources/resourceGroups",
				Location:          "westus2",
				APIVersion:        "2019-10-01",
				ID:                "/subscriptions/bar/providers/Microsoft.Resources/resourceGroup/foo",
				ProvisioningState: "Accepted",
			}

			// setup the applier call with the projected resource
			applier.On("Apply", mock.Anything, &resBefore).Return(&resAfter, nil)
			gvk, err := apiutil.GVKForObject(instance, mgr.GetScheme())
			Expect(err).ToNot(HaveOccurred())
			gr := buildGenericReconciler(gvk, applier)
			result, err := gr.Reconcile(ctrl.Request{
				NamespacedName: nn,
			})
			Expect(err).To(BeNil())
			Expect(result.RequeueAfter).To(Equal(5 * time.Second))
			Expect(k8sClient.Get(ctx, nn, instance)).ToNot(HaveOccurred())
			Expect(instance.Status.ProvisioningState).To(Equal("Accepted"))
			Expect(instance.ObjectMeta.Finalizers).To(ContainElement("infra.azure.com/finalizer"))
			Expect(instance.ObjectMeta.Annotations).To(HaveKey(ResourceSigAnnotationKey))
		})

		It("should delete a resource", func() {
			ctx := context.Background()
			applier := new(ApplierMock)
			nn := client.ObjectKey{
				Namespace: "default",
				Name:      "foo1",
			}

			instance := &microsoftresourcesv1.ResourceGroup{
				TypeMeta: metav1.TypeMeta{
					Kind:       "ResourceGroup",
					APIVersion: microsoftresourcesv1.GroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      nn.Name,
					Namespace: nn.Namespace,
				},
				Spec: microsoftresourcesv1.ResourceGroupSpec{
					Location:   "westus2",
					APIVersion: "2019-10-01",
				},
			}

			createAndReconcileResourceGroup(ctx, instance, applier)
			Expect(k8sClient.Get(ctx, nn, instance)).ToNot(HaveOccurred())
			deleteResourceGroup(ctx, instance, applier)
			Expect(k8sClient.Get(ctx, nn, instance)).To(HaveOccurred())
		})

		It("should requeue if resource group is not succeeded", func() {
			ctx := context.Background()
			applier := new(ApplierMock)
			randomName := test.RandomName("foo", 10)
			nn := client.ObjectKey{
				Namespace: "default",
				Name:      randomName,
			}

			instance := &microsoftresourcesv1.ResourceGroup{
				TypeMeta: metav1.TypeMeta{
					Kind:       "ResourceGroup",
					APIVersion: microsoftresourcesv1.GroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      nn.Name,
					Namespace: nn.Namespace,
				},
				Spec: microsoftresourcesv1.ResourceGroupSpec{
					Location:   "westus2",
					APIVersion: "2019-10-01",
				},
			}

			createAndReconcileResourceGroup(ctx, instance, applier)
			Expect(k8sClient.Get(ctx, nn, instance)).ToNot(HaveOccurred())
			deleteResourceGroup(ctx, instance, applier)
			Expect(k8sClient.Get(ctx, nn, instance)).To(HaveOccurred())
		})

		It("should create a virtual network, but requeue in 30 seconds because the rg is not there", func() {
			ctx := context.Background()
			applier := new(ApplierMock)
			randomName := test.RandomName("foo", 10)
			nn := client.ObjectKey{
				Namespace: "default",
				Name:      randomName,
			}

			subnet := microsoftnetworkv1.Subnet{
				ObjectMeta: metav1.ObjectMeta{
					Name: "subnet1",
				},
				Spec: microsoftnetworkv1.SubnetSpec{
					Properties: microsoftnetworkv1.SubnetProperties{
						AddressPrefixes: []string{
							"10.0.0.0/28",
							"10.1.0.0/28",
						},
					},
				},
			}

			vnetSpecProps := microsoftnetworkv1.VirtualNetworkSpecProperties{
				AddressSpace: &microsoftnetworkv1.AddressSpaceSpec{
					AddressPrefixes: []string{
						"10.0.0.0/16",
					},
				},
				SubnetRefs: []azcorev1.KnownTypeReference{
					{
						Name: subnet.Name,
					},
				},
			}

			group := createResourceGroupByName(ctx, "test-group1")
			obj := &microsoftnetworkv1.VirtualNetwork{
				TypeMeta: metav1.TypeMeta{
					Kind:       "VirtualNetwork",
					APIVersion: microsoftnetworkv1.GroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      nn.Name,
					Namespace: nn.Namespace,
				},
				Spec: microsoftnetworkv1.VirtualNetworkSpec{
					Location:   "westus2",
					APIVersion: "2019-10-01",
					ResourceGroupRef: &azcorev1.KnownTypeReference{
						Namespace: group.Namespace,
						Name:      group.Name,
					},
					Properties: &vnetSpecProps,
				},
			}

			Expect(k8sClient.Create(ctx, obj)).To(Succeed())
			gvk, err := apiutil.GVKForObject(obj, mgr.GetScheme())
			Expect(err).ToNot(HaveOccurred())
			gr := buildGenericReconciler(gvk, applier)
			result, err := gr.Reconcile(ctrl.Request{
				NamespacedName: nn,
			})
			Expect(err).To(BeNil())
			Expect(result.RequeueAfter).To(Equal(30 * time.Second))
		})

		It("should create resource group, routeTable and route out of order", func() {
			ctx := context.Background()
			group := createResourceGroupByNameAndStatus(ctx, "test-group2", microsoftresourcesv1.ResourceGroupStatus{
				ID:                "rg-id",
				ProvisioningState: string(zips.SucceededProvisioningState),
			})

			applier := new(ApplierMock)
			randomName := test.RandomName("foo", 10)
			nn := client.ObjectKey{
				Namespace: "default",
				Name:      randomName,
			}

			route := &microsoftnetworkv1.Route{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Route",
					APIVersion: microsoftnetworkv1.GroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      nn.Name + "-route1",
					Namespace: nn.Namespace,
				},
				Spec: microsoftnetworkv1.RouteSpec{
					APIVersion: "2019-10-01",
					Properties: &microsoftnetworkv1.RouteSpecProperties{
						AddressPrefix:    "10.0.0.0/24",
						NextHopIPAddress: "10.0.0.1",
						NextHopType:      "VnetLocal",
					},
				},
			}

			rt := &microsoftnetworkv1.RouteTable{
				TypeMeta: metav1.TypeMeta{
					Kind:       "RouteTable",
					APIVersion: microsoftnetworkv1.GroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      nn.Name + "-routetable",
					Namespace: nn.Namespace,
				},
				Spec: microsoftnetworkv1.RouteTableSpec{
					Location:   "westus2",
					APIVersion: "2019-10-01",
					ResourceGroupRef: &azcorev1.KnownTypeReference{
						Namespace: group.Namespace,
						Name:      group.Name,
					},
					Properties: &microsoftnetworkv1.RouteTableSpecProperties{
						DisableBGPRoutePropagation: false,
						RouteRefs: []azcorev1.KnownTypeReference{
							{
								Name:      route.Name,
								Namespace: route.Namespace,
							},
						},
					},
				},
			}

			// reconcile the route first, because it depends on the routeTable resource
			Expect(k8sClient.Create(ctx, route)).To(Succeed())
			gvk, err := apiutil.GVKForObject(route, mgr.GetScheme())
			Expect(err).ToNot(HaveOccurred())
			gr := buildGenericReconciler(gvk, applier)
			result, err := gr.Reconcile(ctrl.Request{
				NamespacedName: types.NamespacedName{
					Namespace: route.Namespace,
					Name:      route.Name,
				},
			})
			Expect(err).To(BeNil())
			Expect(result.RequeueAfter).To(Equal(30 * time.Second)) // requeue after 30 seconds b/c owner count is greater than or equal to 1, required for the sub resource to apply

			// create routetable, but not in succeeded state
			Expect(k8sClient.Create(ctx, rt)).To(Succeed())

			// update with a routeTable owner reference
			route.OwnerReferences = ownerutil.EnsureOwnerRef(route.OwnerReferences, metav1.OwnerReference{
				APIVersion: microsoftnetworkv1.GroupVersion.String(),
				Kind:       "RouteTable",
				Name:       rt.Name,
				UID:        rt.UID,
			})
			Expect(k8sClient.Update(ctx, route)).To(Succeed())
			result, err = gr.Reconcile(ctrl.Request{
				NamespacedName: types.NamespacedName{
					Namespace: route.Namespace,
					Name:      route.Name,
				},
			})
			Expect(err).To(BeNil())
			Expect(route.OwnerReferences).To(HaveLen(1))
			Expect(result.RequeueAfter).To(Equal(30 * time.Second)) // requeue after 30 seconds b/c owner(s) is not in succeeded state

			// update the routeTable to succeeded, should apply the route
			rt.Status.ProvisioningState = string(zips.SucceededProvisioningState)
			Expect(k8sClient.Status().Update(ctx, rt)).To(Succeed())

			// now that the owner is in succeeded state, we expect apply will be called with the sub resource route
			bits, err := json.Marshal(route.Spec.Properties)
			Expect(err).ToNot(HaveOccurred())
			beforeRouteResource := &zips.Resource{
				Name:          fmt.Sprintf("%s/%s", rt.Name, route.Name),
				Type:          "Microsoft.Network/routeTables/routes",
				APIVersion:    "2019-10-01",
				ResourceGroup: rt.GetResourceGroupObjectRef().Name,
				Properties:    bits,
			}
			afterRouteResource := *beforeRouteResource
			afterRouteResource.ProvisioningState = zips.SucceededProvisioningState
			applier.On("Apply", mock.Anything, beforeRouteResource).Return(&afterRouteResource, nil)
			result, err = gr.Reconcile(ctrl.Request{
				NamespacedName: types.NamespacedName{
					Namespace: route.Namespace,
					Name:      route.Name,
				},
			})
			Expect(err).To(BeNil())
			Expect(result.RequeueAfter).To(Equal(0 * time.Second)) // should proceed with apply
		})
	})
})

func createResourceGroupByName(ctx context.Context, name string) *microsoftresourcesv1.ResourceGroup {
	return createResourceGroupByNameAndStatus(ctx, name, microsoftresourcesv1.ResourceGroupStatus{})
}

func createResourceGroupByNameAndStatus(ctx context.Context, name string, status microsoftresourcesv1.ResourceGroupStatus) *microsoftresourcesv1.ResourceGroup {
	group := &microsoftresourcesv1.ResourceGroup{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ResourceGroup",
			APIVersion: microsoftresourcesv1.GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: "default",
		},
		Spec: microsoftresourcesv1.ResourceGroupSpec{
			Location:   "westus2",
			APIVersion: "2019-10-01",
		},
		Status: status,
	}
	Expect(k8sClient.Create(ctx, group)).To(Succeed())
	return group
}

func createAndReconcileResourceGroup(ctx context.Context, obj *microsoftresourcesv1.ResourceGroup, applier *ApplierMock) {
	nn := client.ObjectKey{
		Name:      obj.ObjectMeta.Name,
		Namespace: obj.ObjectMeta.Namespace,
	}

	Expect(k8sClient.Create(ctx, obj)).To(Succeed())
	resBefore := zips.Resource{
		Name:       nn.Name,
		Type:       obj.ResourceType(),
		Location:   obj.Spec.Location,
		APIVersion: obj.Spec.APIVersion,
	}

	resAfter := zips.Resource{
		Type:              obj.ResourceType(),
		Location:          obj.Spec.Location,
		APIVersion:        obj.Spec.APIVersion,
		ID:                "/subscriptions/bar/providers/Microsoft.Resources/resourceGroup/foo",
		ProvisioningState: zips.SucceededProvisioningState, // short cutting with succeeded rather than Accepted -> Succeeded
	}

	// setup the applier call with the projected resource
	applier.On("Apply", mock.Anything, &resBefore).Return(&resAfter, nil)
	gvk, err := apiutil.GVKForObject(obj, mgr.GetScheme())
	Expect(err).ToNot(HaveOccurred())
	gr := buildGenericReconciler(gvk, applier)
	result, err := gr.Reconcile(ctrl.Request{
		NamespacedName: nn,
	})
	Expect(err).To(BeNil())
	Expect(result.RequeueAfter).To(BeZero())
}

func deleteResourceGroup(ctx context.Context, obj *microsoftresourcesv1.ResourceGroup, applier *ApplierMock) {
	Expect(k8sClient.Delete(ctx, obj)).ToNot(HaveOccurred())
	nn := client.ObjectKey{
		Name:      obj.ObjectMeta.Name,
		Namespace: obj.ObjectMeta.Namespace,
	}

	resBefore := zips.Resource{
		Name:              nn.Name,
		Type:              obj.ResourceType(),
		Location:          obj.Spec.Location,
		APIVersion:        obj.Spec.APIVersion,
		ProvisioningState: zips.SucceededProvisioningState,
		ID:                "/subscriptions/bar/providers/Microsoft.Resources/resourceGroup/foo",
	}

	resAfter := zips.Resource{
		Name:              nn.Name,
		Type:              obj.ResourceType(),
		Location:          obj.Spec.Location,
		APIVersion:        obj.Spec.APIVersion,
		ProvisioningState: zips.DeletingProvisioningState,
		ID:                "/subscriptions/bar/providers/Microsoft.Resources/resourceGroup/foo",
	}

	applier.On("BeginDelete", mock.Anything, &resBefore).Return(&resAfter, nil)
	gvk, err := apiutil.GVKForObject(obj, mgr.GetScheme())
	Expect(err).ToNot(HaveOccurred())
	gr := buildGenericReconciler(gvk, applier)
	result, err := gr.Reconcile(ctrl.Request{
		NamespacedName: nn,
	})
	Expect(err).To(BeNil())
	Expect(result.RequeueAfter).To(Equal(5 * time.Second))

	applier.On("HeadResource", mock.Anything, &resAfter).Return(false, nil)
	result, err = gr.Reconcile(ctrl.Request{
		NamespacedName: nn,
	})
	Expect(err).To(BeNil())
	Expect(result.RequeueAfter).To(BeZero())
}

func buildGenericReconciler(gvk schema.GroupVersionKind, applier *ApplierMock) *GenericReconciler {
	return &GenericReconciler{
		GVK:       gvk,
		Client:    k8sClient,
		Applier:   applier,
		Scheme:    mgr.GetScheme(),
		Log:       ctrl.Log.WithName("test-controller"),
		Name:      "test-controller",
		Recorder:  record.NewFakeRecorder(10),
		Converter: xform.NewARMConverter(mgr.GetClient(), mgr.GetScheme()),
	}
}

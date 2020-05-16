/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package xform

import (
	"testing"

	"github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	azcorev1 "github.com/Azure/k8s-infra/apis/core/v1"
	"github.com/Azure/k8s-infra/internal/test"
)

func TestGetTypeReferenceData(t *testing.T) {
	randomName := test.RandomName("foo", 10)
	nn := &client.ObjectKey{
		Namespace: "default",
		Name:      randomName,
	}

	rt := newLocalRouteTable(nn)
	route := newRoute(nn)
	rt.Spec.Properties.RouteRefs = append(rt.Spec.Properties.RouteRefs, azcorev1.KnownTypeReference{
		Name:      route.Name,
		Namespace: route.Namespace,
	})

	refsData, err := GetTypeReferenceData(rt)
	g := gomega.NewGomegaWithT(t)
	g.Expect(err).ToNot(gomega.HaveOccurred())
	g.Expect(refsData).ToNot(gomega.BeEmpty())
	g.Expect(refsData).To(gomega.HaveLen(4))
	g.Expect(refsData).To(gomega.Equal([]TypeReferenceLocation{
		{
			JSONFieldName:     "embeddedBazzRef",
			TemplateFieldName: "embeddedBazz",
			Path:              []string{"spec", "properties"},
			Group:             "microsoft.network.infra.azure.com",
			Kind:              "Bazz",
			IsSlice:           false,
		},
		{
			Path:              []string{"spec", "properties"},
			TemplateFieldName: "routes",
			JSONFieldName:     "routeRefs",
			Group:             "microsoft.network.infra.azure.com",
			Kind:              "Route",
			IsSlice:           true,
		},
		{
			JSONFieldName:     "blahRefs",
			TemplateFieldName: "blahs",
			Path:              []string{"spec", "properties", "foo"},
			Group:             "microsoft.network.infra.azure.com",
			Kind:              "Blah",
			IsSlice:           true,
		},
		{
			JSONFieldName:     "bazzRef",
			TemplateFieldName: "bazz",
			Path:              []string{"spec", "properties", "foo"},
			Group:             "microsoft.network.infra.azure.com",
			Kind:              "Bazz",
			IsSlice:           false,
		},
	}))
}

func TestTypeReferenceLocation_JSONFields(t *testing.T) {
	trl := &TypeReferenceLocation{
		JSONFieldName:     "foo",
		TemplateFieldName: "bar",
		Path:              []string{"hello", "world"},
		Group:             "group",
		Kind:              "kind",
		IsSlice:           false,
	}
	g := gomega.NewGomegaWithT(t)
	g.Expect(trl.JSONFields()).To(gomega.Equal([]string{"hello", "world", "foo"}))
	g.Expect(trl.Path).To(gomega.Equal([]string{"hello", "world"}))
}

func TestTypeReferenceLocation_TemplateFields(t *testing.T) {
	trl := &TypeReferenceLocation{
		JSONFieldName:     "foo",
		TemplateFieldName: "bar",
		Path:              []string{"hello", "world"},
		Group:             "group",
		Kind:              "kind",
		IsSlice:           false,
	}
	g := gomega.NewGomegaWithT(t)
	g.Expect(trl.TemplateFields()).To(gomega.Equal([]string{"hello", "world", "bar"}))
	g.Expect(trl.Path).To(gomega.Equal([]string{"hello", "world"}))
}

func newLocalRouteTable(nn *client.ObjectKey) *RouteTable {
	return &RouteTable{
		TypeMeta: metav1.TypeMeta{
			Kind: "RouteTable",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      nn.Name,
			Namespace: nn.Namespace,
		},
		Spec: RouteTableSpec{
			Location:   "westus2",
			APIVersion: "2019-11-01",
			ResourceGroupRef: &azcorev1.KnownTypeReference{
				Name:      nn.Name,
				Namespace: nn.Namespace,
			},
			Properties: &RouteTableSpecProperties{},
		},
	}
}

type (
	Foo struct {
		BlahRefs []azcorev1.KnownTypeReference `json:"blahRefs,omitempty" group:"microsoft.network.infra.azure.com" kind:"Blah"`
		BazzRef  *azcorev1.KnownTypeReference  `json:"bazzRef,omitempty" group:"microsoft.network.infra.azure.com" kind:"Bazz"`
	}

	Embedded struct {
		EmbeddedBazzRef *azcorev1.KnownTypeReference `json:"embeddedBazzRef,omitempty" group:"microsoft.network.infra.azure.com" kind:"Bazz"`
	}

	// RouteTableSpecProperties are the resource specific properties
	RouteTableSpecProperties struct {
		*Embedded
		DisableBGPRoutePropagation bool                          `json:"disableBgpRoutePropagation,omitempty"`
		RouteRefs                  []azcorev1.KnownTypeReference `json:"routeRefs,omitempty" group:"microsoft.network.infra.azure.com" kind:"Route"`
		Foo                        *Foo                          `json:"foo,omitempty"`
	}

	// RouteTableSpec defines the desired state of RouteTable
	RouteTableSpec struct {
		// +k8s:conversion-gen=false
		APIVersion string `json:"apiVersion"`

		// ResourceGroupRef is the Azure Resource Group the VirtualNetwork resides within
		// +kubebuilder:validation:Required
		ResourceGroupRef *azcorev1.KnownTypeReference `json:"resourceGroupRef" group:"microsoft.resources.infra.azure.com" kind:"ResourceGroup"`

		// Location of the VNET in Azure
		// +kubebuilder:validation:Required
		Location string `json:"location"`

		// Tags are user defined key value pairs
		// +optional
		Tags map[string]string `json:"tags,omitempty"`

		// Properties of the Virtual Network
		Properties *RouteTableSpecProperties `json:"properties,omitempty"`
	}

	// RouteTableStatus defines the observed state of RouteTable
	RouteTableStatus struct {
		ID string `json:"id,omitempty"`
		// +k8s:conversion-gen=false
		DeploymentID      string `json:"deploymentId,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// RouteTable is the Schema for the routetables API
	RouteTable struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   RouteTableSpec   `json:"spec,omitempty"`
		Status RouteTableStatus `json:"status,omitempty"`
	}
)

func (rt *RouteTable) ResourceType() string {
	return "Microsoft.Network/routeTables"
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTable) DeepCopyInto(out *RouteTable) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTable.
func (in *RouteTable) DeepCopy() *RouteTable {
	if in == nil {
		return nil
	}
	out := new(RouteTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteTable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableSpec) DeepCopyInto(out *RouteTableSpec) {
	*out = *in
	if in.ResourceGroupRef != nil {
		in, out := &in.ResourceGroupRef, &out.ResourceGroupRef
		*out = new(azcorev1.KnownTypeReference)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(RouteTableSpecProperties)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableSpec.
func (in *RouteTableSpec) DeepCopy() *RouteTableSpec {
	if in == nil {
		return nil
	}
	out := new(RouteTableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableSpecProperties) DeepCopyInto(out *RouteTableSpecProperties) {
	*out = *in
	if in.RouteRefs != nil {
		in, out := &in.RouteRefs, &out.RouteRefs
		*out = make([]azcorev1.KnownTypeReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableSpecProperties.
func (in *RouteTableSpecProperties) DeepCopy() *RouteTableSpecProperties {
	if in == nil {
		return nil
	}
	out := new(RouteTableSpecProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTableStatus) DeepCopyInto(out *RouteTableStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTableStatus.
func (in *RouteTableStatus) DeepCopy() *RouteTableStatus {
	if in == nil {
		return nil
	}
	out := new(RouteTableStatus)
	in.DeepCopyInto(out)
	return out
}

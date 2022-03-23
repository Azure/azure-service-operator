/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reflecthelpers_test

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	//nolint:staticcheck // ignoring deprecation (SA1019) to unblock CI builds
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"

	. "github.com/onsi/gomega"
)

type ResourceWithReferences struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceWithReferencesSpec `json:"spec,omitempty"`
}

var _ client.Object = &ResourceWithReferences{}

func (in *ResourceWithReferences) DeepCopyInto(out *ResourceWithReferences) {
	// TODO: This isn't a full impelmentation, but that's ok we don't need it
	*out = *in
	out.TypeMeta = in.TypeMeta
}

func (in *ResourceWithReferences) DeepCopy() *ResourceWithReferences {
	if in == nil {
		return nil
	}
	out := new(ResourceWithReferences)
	in.DeepCopyInto(out)
	return out
}

func (in *ResourceWithReferences) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

type ResourceWithReferencesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceWithReferences `json:"items"`
}

func (in *ResourceWithReferencesList) DeepCopyInto(out *ResourceWithReferencesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceWithReferences, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

func (in *ResourceWithReferencesList) DeepCopy() *ResourceWithReferencesList {
	if in == nil {
		return nil
	}
	out := new(ResourceWithReferencesList)
	in.DeepCopyInto(out)
	return out
}

func (in *ResourceWithReferencesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

type ResourceWithReferencesSpec struct {
	Owner genruntime.KnownResourceReference `json:"owner"`

	AzureName string `json:"azureName"`

	Ref *ResourceReference `json:"ref,omitempty"`

	Secret *genruntime.SecretReference `json:"secret,omitempty"`

	Location string `json:"location,omitempty"`
}

type ResourceReference struct {
	Reference genruntime.ResourceReference `armReference:"Id" json:"reference"`
}

func Test_FindReferences(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ref := genruntime.ResourceReference{ARMID: "test"}

	res := ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName: "azureName",
			Location:  "westus",
			Owner: genruntime.KnownResourceReference{
				Name: "myrg",
			},
			Ref: &ResourceReference{
				Reference: ref,
			},
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-group",
			Namespace: "test-namespace",
		},
	}

	refs, err := reflecthelpers.FindResourceReferences(res)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(refs).To(HaveLen(1))
	g.Expect(refs).To(HaveKey(ref))
}

func Test_FindSecrets(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ref := genruntime.SecretReference{Name: "foo", Key: "key"}

	res := ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName: "azureName",
			Location:  "westus",
			Owner: genruntime.KnownResourceReference{
				Name: "myrg",
			},
			Secret: &ref,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-group",
			Namespace: "test-namespace",
		},
	}

	refs, err := reflecthelpers.FindSecretReferences(res)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(refs).To(HaveLen(1))
	g.Expect(refs).To(HaveKey(ref))
}

func Test_GetObjectListItems(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	res := ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName: "azureName",
			Location:  "westus",
			Owner: genruntime.KnownResourceReference{
				Name: "myrg",
			},
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-group",
			Namespace: "test-namespace",
		},
	}

	list := &ResourceWithReferencesList{
		Items: []ResourceWithReferences{
			res,
		},
	}

	items, err := reflecthelpers.GetObjectListItems(list)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(items).To(HaveLen(1))
	g.Expect(items[0].GetName()).To(Equal("test-group"))
}

func Test_SetObjectListItems(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	res := &ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName: "azureName",
			Location:  "westus",
			Owner: genruntime.KnownResourceReference{
				Name: "myrg",
			},
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-group",
			Namespace: "test-namespace",
		},
	}

	list := &ResourceWithReferencesList{}

	itemList := []client.Object{res}
	err := reflecthelpers.SetObjectListItems(list, itemList)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(list.Items).To(HaveLen(1))
	g.Expect(list.Items[0].GetName()).To(Equal("test-group"))
}

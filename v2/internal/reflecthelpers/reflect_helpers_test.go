/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reflecthelpers_test

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ResourceWithReferences struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceWithReferencesSpec   `json:"spec,omitempty"`
	Status            ResourceWithReferencesStatus `json:"status,omitempty"`
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

	Ref      *ResourceReference                      `json:"ref,omitempty"`
	RefSlice []genruntime.ResourceReference          `armReference:"RefSlice" json:"refSlice,omitempty"`
	RefMap   map[string]genruntime.ResourceReference `armReference:"RefMap" json:"refMap,omitempty"`

	Secret *genruntime.SecretReference `json:"secret,omitempty"`

	PropertyWithTag           *string                        `optionalConfigMapPair:"PropertyWithTag" json:"propertyWithTag"`
	PropertyWithTagFromConfig *genruntime.ConfigMapReference `optionalConfigMapPair:"PropertyWithTag" json:"propertyWithTagFromConfig"`

	Location string `json:"location,omitempty"`
}

var _ genruntime.ARMTransformer = &ResourceWithReferencesSpec{}

func (in *ResourceWithReferencesSpec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	panic("not expected to be called")
}

func (in *ResourceWithReferencesSpec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	panic("not expected to be called")
}

func (in *ResourceWithReferencesSpec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, input interface{}) error {
	panic("not expected to be called")
}

type ResourceWithReferencesStatus struct {
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
}

type ResourceReference struct {
	Reference genruntime.ResourceReference `armReference:"Id" json:"reference"`
}

type ProvisioningState string

const (
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateFailed    ProvisioningState = "Failed"
)

func Test_FindReferences(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ref1 := genruntime.ResourceReference{ARMID: "test1"}
	ref2 := genruntime.ResourceReference{ARMID: "test2"}
	ref3 := genruntime.ResourceReference{ARMID: "test3"}
	ref4 := genruntime.ResourceReference{ARMID: "test4"}
	ref5 := genruntime.ResourceReference{ARMID: "test5"}

	res := ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName: "azureName",
			Location:  "westus",
			Owner: genruntime.KnownResourceReference{
				Name: "myrg",
			},
			Ref: &ResourceReference{
				Reference: ref1,
			},
			RefSlice: []genruntime.ResourceReference{
				ref2,
				ref3,
			},
			RefMap: map[string]genruntime.ResourceReference{
				"a": ref4,
				"b": ref5,
			},
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-group",
			Namespace: "test-namespace",
		},
	}

	refs, err := reflecthelpers.FindResourceReferences(res)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(refs).To(HaveLen(5))
	g.Expect(refs).To(HaveKey(ref1))
	g.Expect(refs).To(HaveKey(ref2))
	g.Expect(refs).To(HaveKey(ref3))
	g.Expect(refs).To(HaveKey(ref4))
	g.Expect(refs).To(HaveKey(ref5))
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

func Test_FindPropertiesWithTag(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ref1 := genruntime.ResourceReference{ARMID: "test1"}
	ref2 := genruntime.ResourceReference{ARMID: "test2"}
	ref3 := genruntime.ResourceReference{ARMID: "test3"}
	ref4 := genruntime.ResourceReference{ARMID: "test4"}
	ref5 := genruntime.ResourceReference{ARMID: "test5"}

	res := ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName: "azureName",
			Location:  "westus",
			Owner: genruntime.KnownResourceReference{
				Name: "myrg",
			},
			Ref: &ResourceReference{
				Reference: ref1,
			},
			RefSlice: []genruntime.ResourceReference{
				ref2,
				ref3,
			},
			RefMap: map[string]genruntime.ResourceReference{
				"a": ref4,
				"b": ref5,
			},
			PropertyWithTag: to.Ptr("hello"),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-group",
			Namespace: "test-namespace",
		},
	}

	results, err := reflecthelpers.FindPropertiesWithTag(res, "optionalConfigMapPair")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(results).To(HaveLen(2))
	g.Expect(results).To(HaveKey("Spec.PropertyWithTag"))
	g.Expect(results["Spec.PropertyWithTag"]).To(Equal([]any{to.Ptr("hello")}))
	g.Expect(results).To(HaveKey("Spec.PropertyWithTagFromConfig"))
	g.Expect(results["Spec.PropertyWithTagFromConfig"]).To(Equal([]any{(*genruntime.ConfigMapReference)(nil)}))

	// Now try finding all the JSON tags
	results, err = reflecthelpers.FindPropertiesWithTag(res.Spec, "json")
	g.Expect(err).ToNot(HaveOccurred())
	// This is the number of properties and child properties on this object. It's fragile to structural changes
	// in the object so may need to be changed in the future
	g.Expect(results).To(HaveLen(24))
}

func Test_FindOptionalConfigMapReferences(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ref1 := genruntime.ResourceReference{ARMID: "test1"}
	ref2 := genruntime.ResourceReference{ARMID: "test2"}
	ref3 := genruntime.ResourceReference{ARMID: "test3"}
	ref4 := genruntime.ResourceReference{ARMID: "test4"}
	ref5 := genruntime.ResourceReference{ARMID: "test5"}

	res := ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName: "azureName",
			Location:  "westus",
			Owner: genruntime.KnownResourceReference{
				Name: "myrg",
			},
			Ref: &ResourceReference{
				Reference: ref1,
			},
			RefSlice: []genruntime.ResourceReference{
				ref2,
				ref3,
			},
			RefMap: map[string]genruntime.ResourceReference{
				"a": ref4,
				"b": ref5,
			},
			PropertyWithTag: to.Ptr("hello"),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-group",
			Namespace: "test-namespace",
		},
	}

	results, err := reflecthelpers.FindOptionalConfigMapReferences(res)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(results).To(HaveLen(1))
	g.Expect(results[0].Name).To(Equal("Spec.PropertyWithTag"))
	g.Expect(results[0].Value).To(Equal(to.Ptr("hello")))
	g.Expect(results[0].RefName).To(Equal("Spec.PropertyWithTagFromConfig"))
	g.Expect(results[0].Ref).To(Equal((*genruntime.ConfigMapReference)(nil)))
}

// defaultResourceReferencesName exists to showcase an example where ReflectVisitor is used to modify the object in question
func defaultResourceReferencesName(transformer genruntime.ARMTransformer, name string) error {
	visitor := reflecthelpers.NewReflectVisitor()
	visitor.VisitStruct = func(this *reflecthelpers.ReflectVisitor, it reflect.Value, ctx interface{}) error {
		if it.Type() == reflect.TypeOf(genruntime.ResourceReference{}) {
			if it.CanInterface() {
				reference := it.Interface().(genruntime.ResourceReference)
				if reference.Name == "" {
					// Cannot do assignment on the reference variable as it is a copy
					f := it.FieldByName("Name")
					if !f.CanSet() {
						return eris.New("cannot set 'Name' field of 'genruntime.ResourceReference'")
					}
					f.SetString(name)
				}
			} else {
				// This should be impossible given how the visitor works
				return eris.New("genruntime.ResourceReference field was unexpectedly nil")
			}
			return nil
		}

		return reflecthelpers.IdentityVisitStruct(this, it, ctx)
	}

	err := visitor.Visit(transformer, nil)
	if err != nil {
		return eris.Wrap(err, "defaulting genruntime.ResourceReference")
	}

	return nil
}

func Test_CanUseReflectVisitorToModifyResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ref := genruntime.ResourceReference{Group: "microsoft.keyvault", Kind: "keyvault"}

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

	err := defaultResourceReferencesName(&res.Spec, "myname")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(res.Spec.Ref.Reference.Name).To(Equal("myname"))
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

func Test_SetProperty_TargetingStringProperty_MakesChange(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferencesSpec{
		AzureName: "azureName",
	}

	newValue := "newName"
	err := reflecthelpers.SetProperty(subject, "AzureName", newValue)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(subject.AzureName).To(Equal(newValue))
}

func Test_SetProperty_TargetingMultipleProperties_MakesChanges(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferencesSpec{
		AzureName:       "azureName",
		Location:        "westus",
		PropertyWithTag: to.Ptr("hello"),
	}

	name := "dorothy"
	location := "land-of-oz"
	property := to.Ptr("flying-house")

	g.Expect(reflecthelpers.SetProperty(subject, "AzureName", name)).To(Succeed())
	g.Expect(reflecthelpers.SetProperty(subject, "Location", location)).To(Succeed())
	g.Expect(reflecthelpers.SetProperty(subject, "PropertyWithTag", property)).To(Succeed())

	g.Expect(subject.AzureName).To(Equal(name))
	g.Expect(subject.Location).To(Equal(location))
	g.Expect(subject.PropertyWithTag).To(Equal(property))
}

func Test_SetProperty_TargetingNestedStringProperty_MakesChange(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName: "azureName",
		},
	}

	newValue := "newName"
	err := reflecthelpers.SetProperty(subject, "Spec.AzureName", newValue)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(subject.Spec.AzureName).To(Equal(newValue))
}

func Test_SetProperty_TargetingMultipleNestedProperties_MakesChange(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName:       "azureName",
			Location:        "westus",
			PropertyWithTag: to.Ptr("hello"),
		},
	}

	name := "dorothy"
	location := "land-of-oz"
	property := to.Ptr("flying-house")

	g.Expect(reflecthelpers.SetProperty(subject, "Spec.AzureName", name)).To(Succeed())
	g.Expect(reflecthelpers.SetProperty(subject, "Spec.Location", location)).To(Succeed())
	g.Expect(reflecthelpers.SetProperty(subject, "Spec.PropertyWithTag", property)).To(Succeed())

	g.Expect(subject.Spec.AzureName).To(Equal(name))
	g.Expect(subject.Spec.Location).To(Equal(location))
	g.Expect(subject.Spec.PropertyWithTag).To(Equal(property))
}

func Test_SetProperty_TargetingNestedNestedStringProperty_MakesChange(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName: "azureName",
		},
	}

	newValue := "newName"
	err := reflecthelpers.SetProperty(subject, "Spec.Owner.Name", newValue)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(subject.Spec.Owner.Name).To(Equal(newValue))
}

func Test_SetProperty_TargetingMultipleNestedNestedProperties_MakesChanges(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferences{}

	name := "lock"
	key := "key"

	g.Expect(reflecthelpers.SetProperty(subject, "Spec.Secret.Name", name)).To(Succeed())
	g.Expect(reflecthelpers.SetProperty(subject, "Spec.Secret.Key", key)).To(Succeed())

	g.Expect(subject.Spec.Secret.Name).To(Equal(name))
	g.Expect(subject.Spec.Secret.Key).To(Equal(key))
}

func Test_SetProperty_TargetingUnknownProperty_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferencesSpec{
		AzureName: "azureName",
	}

	err := reflecthelpers.SetProperty(subject, "UnknownProperty", "newValue")
	g.Expect(err).To(HaveOccurred())
	g.Expect(err).To(MatchError(ContainSubstring("UnknownProperty")))
}

func Test_SetProperty_TargetingUnknownNestedProperty_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferencesSpec{}

	err := reflecthelpers.SetProperty(subject, "Owner.UnknownProperty", "newValue")
	g.Expect(err).To(HaveOccurred())

	g.Expect(err).To(MatchError(ContainSubstring("Owner")))
	g.Expect(err).To(MatchError(ContainSubstring("UnknownProperty")))
}

func Test_SetProperty_WhenValueOfWrongType_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName: "azureName",
		},
	}

	err := reflecthelpers.SetProperty(subject, "Spec.AzureName", make([]int, 0, 10))
	g.Expect(err).To(HaveOccurred())
	g.Expect(err).To(MatchError(ContainSubstring("Spec")))
	g.Expect(err).To(MatchError(ContainSubstring("AzureName")))
	g.Expect(err).To(MatchError(ContainSubstring("kind []")))
	g.Expect(err).To(MatchError(ContainSubstring("not compatible")))
}

func Test_SetProperty_WhenValueOfCompatibleType_ModifiesValue(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferences{
		Status: ResourceWithReferencesStatus{
			ProvisioningState: to.Ptr(ProvisioningStateSucceeded),
		},
	}

	err := reflecthelpers.SetProperty(subject, "Status.ProvisioningState", to.Ptr(string(ProvisioningStateFailed)))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(*subject.Status.ProvisioningState).To(Equal(ProvisioningStateFailed))
}

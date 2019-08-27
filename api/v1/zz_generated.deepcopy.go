// +build !ignore_autogenerated

/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// autogenerated by controller-gen object, do not modify manually

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaptureDescription) DeepCopyInto(out *CaptureDescription) {
	*out = *in
	out.Destination = in.Destination
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaptureDescription.
func (in *CaptureDescription) DeepCopy() *CaptureDescription {
	if in == nil {
		return nil
	}
	out := new(CaptureDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsumerGroup) DeepCopyInto(out *ConsumerGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsumerGroup.
func (in *ConsumerGroup) DeepCopy() *ConsumerGroup {
	if in == nil {
		return nil
	}
	out := new(ConsumerGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsumerGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsumerGroupList) DeepCopyInto(out *ConsumerGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConsumerGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsumerGroupList.
func (in *ConsumerGroupList) DeepCopy() *ConsumerGroupList {
	if in == nil {
		return nil
	}
	out := new(ConsumerGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsumerGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsumerGroupSpec) DeepCopyInto(out *ConsumerGroupSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsumerGroupSpec.
func (in *ConsumerGroupSpec) DeepCopy() *ConsumerGroupSpec {
	if in == nil {
		return nil
	}
	out := new(ConsumerGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsumerGroupStatus) DeepCopyInto(out *ConsumerGroupStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsumerGroupStatus.
func (in *ConsumerGroupStatus) DeepCopy() *ConsumerGroupStatus {
	if in == nil {
		return nil
	}
	out := new(ConsumerGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Destination) DeepCopyInto(out *Destination) {
	*out = *in
	out.StorageAccount = in.StorageAccount
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Destination.
func (in *Destination) DeepCopy() *Destination {
	if in == nil {
		return nil
	}
	out := new(Destination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Eventhub) DeepCopyInto(out *Eventhub) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Eventhub.
func (in *Eventhub) DeepCopy() *Eventhub {
	if in == nil {
		return nil
	}
	out := new(Eventhub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Eventhub) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventhubAuthorizationRule) DeepCopyInto(out *EventhubAuthorizationRule) {
	*out = *in
	if in.Rights != nil {
		in, out := &in.Rights, &out.Rights
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventhubAuthorizationRule.
func (in *EventhubAuthorizationRule) DeepCopy() *EventhubAuthorizationRule {
	if in == nil {
		return nil
	}
	out := new(EventhubAuthorizationRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventhubList) DeepCopyInto(out *EventhubList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Eventhub, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventhubList.
func (in *EventhubList) DeepCopy() *EventhubList {
	if in == nil {
		return nil
	}
	out := new(EventhubList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventhubList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventhubNamespace) DeepCopyInto(out *EventhubNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventhubNamespace.
func (in *EventhubNamespace) DeepCopy() *EventhubNamespace {
	if in == nil {
		return nil
	}
	out := new(EventhubNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventhubNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventhubNamespaceList) DeepCopyInto(out *EventhubNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventhubNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventhubNamespaceList.
func (in *EventhubNamespaceList) DeepCopy() *EventhubNamespaceList {
	if in == nil {
		return nil
	}
	out := new(EventhubNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventhubNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventhubNamespaceProperties) DeepCopyInto(out *EventhubNamespaceProperties) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventhubNamespaceProperties.
func (in *EventhubNamespaceProperties) DeepCopy() *EventhubNamespaceProperties {
	if in == nil {
		return nil
	}
	out := new(EventhubNamespaceProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventhubNamespaceResource) DeepCopyInto(out *EventhubNamespaceResource) {
	*out = *in
	out.Sku = in.Sku
	out.Properties = in.Properties
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventhubNamespaceResource.
func (in *EventhubNamespaceResource) DeepCopy() *EventhubNamespaceResource {
	if in == nil {
		return nil
	}
	out := new(EventhubNamespaceResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventhubNamespaceSku) DeepCopyInto(out *EventhubNamespaceSku) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventhubNamespaceSku.
func (in *EventhubNamespaceSku) DeepCopy() *EventhubNamespaceSku {
	if in == nil {
		return nil
	}
	out := new(EventhubNamespaceSku)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventhubNamespaceSpec) DeepCopyInto(out *EventhubNamespaceSpec) {
	*out = *in
	out.Sku = in.Sku
	out.Properties = in.Properties
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventhubNamespaceSpec.
func (in *EventhubNamespaceSpec) DeepCopy() *EventhubNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(EventhubNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventhubNamespaceStatus) DeepCopyInto(out *EventhubNamespaceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventhubNamespaceStatus.
func (in *EventhubNamespaceStatus) DeepCopy() *EventhubNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(EventhubNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventhubProperties) DeepCopyInto(out *EventhubProperties) {
	*out = *in
	out.CaptureDescription = in.CaptureDescription
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventhubProperties.
func (in *EventhubProperties) DeepCopy() *EventhubProperties {
	if in == nil {
		return nil
	}
	out := new(EventhubProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventhubSpec) DeepCopyInto(out *EventhubSpec) {
	*out = *in
	out.Properties = in.Properties
	in.AuthorizationRule.DeepCopyInto(&out.AuthorizationRule)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventhubSpec.
func (in *EventhubSpec) DeepCopy() *EventhubSpec {
	if in == nil {
		return nil
	}
	out := new(EventhubSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventhubStatus) DeepCopyInto(out *EventhubStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventhubStatus.
func (in *EventhubStatus) DeepCopy() *EventhubStatus {
	if in == nil {
		return nil
	}
	out := new(EventhubStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVault) DeepCopyInto(out *KeyVault) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVault.
func (in *KeyVault) DeepCopy() *KeyVault {
	if in == nil {
		return nil
	}
	out := new(KeyVault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeyVault) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultList) DeepCopyInto(out *KeyVaultList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeyVault, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultList.
func (in *KeyVaultList) DeepCopy() *KeyVaultList {
	if in == nil {
		return nil
	}
	out := new(KeyVaultList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeyVaultList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultSpec) DeepCopyInto(out *KeyVaultSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultSpec.
func (in *KeyVaultSpec) DeepCopy() *KeyVaultSpec {
	if in == nil {
		return nil
	}
	out := new(KeyVaultSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyVaultStatus) DeepCopyInto(out *KeyVaultStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyVaultStatus.
func (in *KeyVaultStatus) DeepCopy() *KeyVaultStatus {
	if in == nil {
		return nil
	}
	out := new(KeyVaultStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroup) DeepCopyInto(out *ResourceGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroup.
func (in *ResourceGroup) DeepCopy() *ResourceGroup {
	if in == nil {
		return nil
	}
	out := new(ResourceGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupList) DeepCopyInto(out *ResourceGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupList.
func (in *ResourceGroupList) DeepCopy() *ResourceGroupList {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupSpec) DeepCopyInto(out *ResourceGroupSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupSpec.
func (in *ResourceGroupSpec) DeepCopy() *ResourceGroupSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupStatus) DeepCopyInto(out *ResourceGroupStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupStatus.
func (in *ResourceGroupStatus) DeepCopy() *ResourceGroupStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageAccount) DeepCopyInto(out *StorageAccount) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageAccount.
func (in *StorageAccount) DeepCopy() *StorageAccount {
	if in == nil {
		return nil
	}
	out := new(StorageAccount)
	in.DeepCopyInto(out)
	return out
}

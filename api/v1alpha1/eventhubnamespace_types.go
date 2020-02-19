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

package v1alpha1

import (
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EventhubNamespaceSpec defines the desired state of EventhubNamespace
type EventhubNamespaceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location      string                      `json:"location"`
	Sku           EventhubNamespaceSku        `json:"sku,omitempty"`
	Properties    EventhubNamespaceProperties `json:"properties,omitempty"`
	ResourceGroup string                      `json:"resourceGroup,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// EventhubNamespace is the Schema for the eventhubnamespaces API
type EventhubNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EventhubNamespaceSpec `json:"spec,omitempty"`
	Status ASOStatus             `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventhubNamespaceList contains a list of EventhubNamespace
type EventhubNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventhubNamespace `json:"items"`
}

// EventhubNamespaceSku defines the sku
type EventhubNamespaceSku struct {
	Name     string `json:"name,omitempty"`     //allowedValues "Basic", "Standard"
	Tier     string `json:"tier,omitempty"`     //allowedValues "Basic", "Standard"
	Capacity int32  `json:"capacity,omitempty"` //allowedValues 1, 2, 4
}

//EventhubNamespaceProperties defines the namespace properties
type EventhubNamespaceProperties struct {
	IsAutoInflateEnabled   bool  `json:"isAutoInflateEnabled,omitempty"`
	MaximumThroughputUnits int32 `json:"maximumThroughputUnits,omitempty"`
	KafkaEnabled           bool  `json:"kafkaEnabled,omitempty"`
}

func init() {
	SchemeBuilder.Register(&EventhubNamespace{}, &EventhubNamespaceList{})
}

func (eventhubNamespace *EventhubNamespace) IsBeingDeleted() bool {
	return !eventhubNamespace.ObjectMeta.DeletionTimestamp.IsZero()
}

func (eventhubNamespace *EventhubNamespace) IsSubmitted() bool {
	return eventhubNamespace.Status.Provisioning || eventhubNamespace.Status.Provisioned

}

func (eventhubNamespace *EventhubNamespace) HasFinalizer(finalizerName string) bool {
	return helpers.ContainsString(eventhubNamespace.ObjectMeta.Finalizers, finalizerName)
}

func (eventhubNamespace *EventhubNamespace) AddFinalizer(finalizerName string) {
	eventhubNamespace.ObjectMeta.Finalizers = append(eventhubNamespace.ObjectMeta.Finalizers, finalizerName)
}

func (eventhubNamespace *EventhubNamespace) RemoveFinalizer(finalizerName string) {
	eventhubNamespace.ObjectMeta.Finalizers = helpers.RemoveString(eventhubNamespace.ObjectMeta.Finalizers, finalizerName)
}

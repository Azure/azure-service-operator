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

package v1

import (
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

//EventhubNamespaceResource defines the namespace
type EventhubNamespaceResource struct {
	Name              string                      `json:"name"`
	Location          string                      `json:"location"`
	Sku               EventhubNamespaceSku        `json:"sku,omitempty"`
	Properties        EventhubNamespaceProperties `json:"properties,omitempty"`
	ResourceGroupName string                      `json:"resourcegroup,omitempty"`
}

// EventhubSpec defines the desired state of Eventhub
type EventhubSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location          string                    `json:"location"`
	Namespace         string                    `json:"namespace,omitempty"`
	Properties        EventhubProperties        `json:"properties,omitempty"`
	ResourceGroup     string                    `json:"resourcegroup,omitempty"`
	AuthorizationRule EventhubAuthorizationRule `json:"authorizationrule,omitempty"`
}

// EventhubStatus defines the observed state of Eventhub
type EventhubStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Provisioning bool `json:"provisioning,omitempty"`
	Provisioned  bool `json:"provisioned,omitempty"`
}

//EventhubAuthorizationRule defines the name and rights of the access policy
type EventhubAuthorizationRule struct {
	// Name - Name of AuthorizationRule for eventhub
	Name string `json:"name,omitempty"`
	// Rights - Rights set on the AuthorizationRule
	Rights []string `json:"rights,omitempty"`
}

//EventhubProperties defines the namespace properties
type EventhubProperties struct {
	// +kubebuilder:validation:Maximum=7
	// +kubebuilder:validation:Minimum=1
	// MessageRetentionInDays - Number of days to retain the events for this Event Hub, value should be 1 to 7 days
	MessageRetentionInDays int32 `json:"messageretentionindays,omitempty"`
	// +kubebuilder:validation:Maximum=32
	// +kubebuilder:validation:Minimum=1
	// PartitionCount - Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.
	PartitionCount int32 `json:"partitioncount,omitempty"`
}

// +kubebuilder:object:root=true

// Eventhub is the Schema for the eventhubs API
type Eventhub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EventhubSpec   `json:"spec,omitempty"`
	Status EventhubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventhubList contains a list of Eventhub
type EventhubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Eventhub `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Eventhub{}, &EventhubList{})
}

func (eventhub *Eventhub) IsBeingDeleted() bool {
	return !eventhub.ObjectMeta.DeletionTimestamp.IsZero()
}

func (eventhub *Eventhub) IsSubmitted() bool {
	return eventhub.Status.Provisioning || eventhub.Status.Provisioned
}

func (eventhub *Eventhub) HasFinalizer(finalizerName string) bool {
	return helpers.ContainsString(eventhub.ObjectMeta.Finalizers, finalizerName)
}

func (eventhub *Eventhub) AddFinalizer(finalizerName string) {
	eventhub.ObjectMeta.Finalizers = append(eventhub.ObjectMeta.Finalizers, finalizerName)
}

func (eventhub *Eventhub) RemoveFinalizer(finalizerName string) {
	eventhub.ObjectMeta.Finalizers = helpers.RemoveString(eventhub.ObjectMeta.Finalizers, finalizerName)
}

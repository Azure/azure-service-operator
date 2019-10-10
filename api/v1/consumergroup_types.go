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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConsumerGroupSpec defines the desired state of ConsumerGroup
type ConsumerGroupSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ResourceGroupName string `json:"resourceGroup,omitempty"`
	NamespaceName     string `json:"namespace,omitempty"`
	EventhubName      string `json:"eventHub,omitempty"`
}

// ConsumerGroupStatus defines the observed state of ConsumerGroup
type ConsumerGroupStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Provisioning bool `json:"provisioning,omitempty"`
	Provisioned  bool `json:"provisioned,omitempty"`
}

// +kubebuilder:object:root=true

// ConsumerGroup is the Schema for the consumergroups API
type ConsumerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConsumerGroupSpec   `json:"spec,omitempty"`
	Status ConsumerGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConsumerGroupList contains a list of ConsumerGroup
type ConsumerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConsumerGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConsumerGroup{}, &ConsumerGroupList{})
}

func init() {
	SchemeBuilder.Register(&ConsumerGroup{}, &ConsumerGroupList{})
}

func (consumerGroup *ConsumerGroup) IsBeingDeleted() bool {
	return !consumerGroup.ObjectMeta.DeletionTimestamp.IsZero()
}

func (consumerGroup *ConsumerGroup) IsSubmitted() bool {
	return consumerGroup.Status.Provisioning || consumerGroup.Status.Provisioned

}

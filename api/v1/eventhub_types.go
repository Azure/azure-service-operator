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

// AzureSpec blah
type ResourceGroupSpec struct {
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
}

// EventHubResourceSpec blah
type EventHubResourceSpec struct {
	Location string `json:"location,omitempty"`
	NSName   string `json:"nSName,omitempty"`
	HubName  string `json:"hubName,omitempty"`
}

// EventhubSpec defines the desired state of Eventhub
type EventhubSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ResourceGroupSpec    ResourceGroupSpec    `json:"resourceGroup,omitempty"`
	EventHubResourceSpec EventHubResourceSpec `json:"eventHub,omitempty"`
}

// EventhubStatus defines the observed state of Eventhub
type EventhubStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
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

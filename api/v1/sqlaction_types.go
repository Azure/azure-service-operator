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

// SqlActionSpec defines the desired state of SqlAction
type SqlActionSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ResourceGroup string `json:"resourcegroup"`
	ActionName    string `json:"actionname"`
	ServerName    string `json:"servername"`
}

// SqlActionStatus defines the observed state of SqlAction
type SqlActionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Provisioning bool   `json:"provisioning,omitempty"`
	Provisioned  bool   `json:"provisioned,omitempty"`
	State        string `json:"state,omitempty"`
}

// +kubebuilder:object:root=true

// SqlAction is the Schema for the sqlactions API
type SqlAction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SqlActionSpec   `json:"spec,omitempty"`
	Status SqlActionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SqlActionList contains a list of SqlAction
type SqlActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlAction `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SqlAction{}, &SqlActionList{})
}

func (s *SqlAction) IsSubmitted() bool {
	return s.Status.Provisioned || s.Status.Provisioning
}

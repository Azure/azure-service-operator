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

// SqlUserSpec defines the desired state of SqlUser
type SqlUserSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location      string   `json:"location"`
	ResourceGroup string   `json:"resourcegroup,omitempty"`
	Server        string   `json:"server"`
	DbName        string   `json:"dbname"`
	Roles         []string `json:"roles"`
}

// SqlUserStatus defines the observed state of SqlUser
type SqlUserStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Provisioning bool `json:"provisioning,omitempty"`
	Provisioned  bool `json:"provisioned,omitempty"`
}

// +kubebuilder:object:root=true

// SqlUser is the Schema for the sqlusers API
type SqlUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SqlUserSpec   `json:"spec,omitempty"`
	Status SqlUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SqlUserList contains a list of SqlUser
type SqlUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SqlUser{}, &SqlUserList{})
}

// IsSubmitted checks if sqluser is provisioning
func (s *SqlUser) IsSubmitted() bool {
	return s.Status.Provisioning || s.Status.Provisioned
}

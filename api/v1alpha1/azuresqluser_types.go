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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureSQLUserSpec defines the desired state of SqlUser
type AzureSQLUserSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Server        string   `json:"server"`
	DbName        string   `json:"dbName"`
	ResourceGroup string   `json:"resourceGroup,omitempty"`
	AdminSecret   string   `json:"adminSecret,omitempty"`
	Roles         []string `json:"roles"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AzureSQLUser is the Schema for the sqlusers API
type AzureSQLUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureSQLUserSpec `json:"spec,omitempty"`
	Status ASOStatus        `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureSQLUserList contains a list of SqlUser
type AzureSQLUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureSQLUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureSQLUser{}, &AzureSQLUserList{})
}

// IsSubmitted checks if sqluser is provisioning
func (s *AzureSQLUser) IsSubmitted() bool {
	return s.Status.Provisioning || s.Status.Provisioned
}

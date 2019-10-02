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

// KeyVaultSpec defines the desired state of KeyVault
type KeyVaultSpec struct {
	Location          string `json:"location"`
	ResourceGroupName string `json:"resourceGroup"`
}

// KeyVaultStatus defines the observed state of KeyVault
type KeyVaultStatus struct {
	Provisioning bool `json:"provisioning,omitempty"`
	Provisioned  bool `json:"provisioned,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KeyVault is the Schema for the keyvaults API
type KeyVault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeyVaultSpec   `json:"spec,omitempty"`
	Status KeyVaultStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KeyVaultList contains a list of KeyVault
type KeyVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyVault `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeyVault{}, &KeyVaultList{})
}

func (keyVault *KeyVault) IsSubmitted() bool {
	return keyVault.Status.Provisioning || keyVault.Status.Provisioned
}

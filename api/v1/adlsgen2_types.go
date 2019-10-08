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

// AdlsGen2Spec defines the desired state of AdlsGen2
type AdlsGen2Spec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location			string `json:"location,omitempty"`
	ResourceGroupName 	string `json:"resourceGroup"`
}

// AdlsGen2Status defines the observed state of AdlsGen2
type AdlsGen2Status struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Provisioning	bool `json:"provisioning,omitempty"`
	Provisioned 	bool `json:"provisioned,omitempty"`
}

// +kubebuilder:object:root=true

// AdlsGen2 is the Schema for the adlsgen2s API
type AdlsGen2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AdlsGen2Spec   `json:"spec,omitempty"`
	Status AdlsGen2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AdlsGen2List contains a list of AdlsGen2
type AdlsGen2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdlsGen2 `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdlsGen2{}, &AdlsGen2List{})
}

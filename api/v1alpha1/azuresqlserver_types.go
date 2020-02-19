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
	"k8s.io/apimachinery/pkg/types"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureSqlServerSpec defines the desired state of AzureSqlServer
type AzureSqlServerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location      string `json:"location"`
	ResourceGroup string `json:"resourcegroup,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// AzureSqlServer is the Schema for the azuresqlservers API
type AzureSqlServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureSqlServerSpec `json:"spec,omitempty"`
	Status ASOStatus          `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureSqlServerList contains a list of AzureSqlServer
type AzureSqlServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureSqlServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureSqlServer{}, &AzureSqlServerList{})
}

func (s *AzureSqlServer) IsSubmitted() bool {
	return s.Status.Provisioned || s.Status.Provisioning
}

// NewAzureSQLServer returns a simple server struct filled with passed in values
func NewAzureSQLServer(names types.NamespacedName, resourceGroup, region string) *AzureSqlServer {
	return &AzureSqlServer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      names.Name,
			Namespace: names.Namespace,
		},
		Spec: AzureSqlServerSpec{
			Location:      region,
			ResourceGroup: resourceGroup,
		},
	}
}

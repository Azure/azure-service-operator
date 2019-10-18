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
	sql "github.com/Azure/azure-service-operator/pkg/resourcemanager/sqlclient"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureSQLFailoverGroupSpec defines the desired state of AzureSQLFailoverGroup
type AzureSQLFailoverGroupSpec struct {
	// Important: Run "make" to regenerate code after modifying this file
	Location                     string                              `json:"location"`
	ResourceGroup                string                              `json:"resourcegroup,omitempty"`
	Server                       string                              `json:"server"`
	FailoverPolicy               sql.ReadWriteEndpointFailoverPolicy `json:"failoverpolicy"`
	FailoverGracePeriod          int32                               `json:"failovergraceperiod"`
	SecondaryServerName          string                              `json:"secondaryserver"`
	SecondaryServerResourceGroup string                              `json:"secondaryserverresourcegroup"`
	DatabaseList                 []string                            `json:"databaselist"`
}

// AzureSQLFailoverGroupStatus defines the observed state of AzureSQLFailoverGroup
type AzureSQLFailoverGroupStatus struct {
	// Important: Run "make" to regenerate code after modifying this file
	Provisioning bool   `json:"provisioning,omitempty"`
	Provisioned  bool   `json:"provisioned,omitempty"`
	State        string `json:"state,omitempty"`
	Message      string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// AzureSQLFailoverGroup is the Schema for the azuresqlfailovergroups API
type AzureSQLFailoverGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureSQLFailoverGroupSpec   `json:"spec,omitempty"`
	Status AzureSQLFailoverGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureSQLFailoverGroupList contains a list of AzureSQLFailoverGroup
type AzureSQLFailoverGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureSQLFailoverGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureSQLFailoverGroup{}, &AzureSQLFailoverGroupList{})
}

func (s *AzureSQLFailoverGroup) IsSubmitted() bool {
	return s.Status.Provisioned || s.Status.Provisioning
}

func (s *AzureSQLFailoverGroup) IsProvisioned() bool {
	return s.Status.Provisioned
}

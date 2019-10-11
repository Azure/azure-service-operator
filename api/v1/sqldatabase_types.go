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
	sql "github.com/Azure/azure-service-operator/pkg/resourcemanager/sqlclient"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SqlDatabaseSpec defines the desired state of SqlDatabase
type SqlDatabaseSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location      string        `json:"location"`
	ResourceGroup string        `json:"resourcegroup,omitempty"`
	Server        string        `json:"server"`
	Edition       sql.DBEdition `json:"edition"`
}

// SqlDatabaseStatus defines the observed state of SqlDatabase
type SqlDatabaseStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Provisioning bool `json:"provisioning,omitempty"`
	Provisioned  bool `json:"provisioned,omitempty"`
}

// +kubebuilder:object:root=true
// SqlDatabase is the Schema for the sqldatabases API
type SqlDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SqlDatabaseSpec   `json:"spec,omitempty"`
	Status SqlDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SqlDatabaseList contains a list of SqlDatabase
type SqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SqlDatabase{}, &SqlDatabaseList{})
}

func (s *SqlDatabase) IsSubmitted() bool {
	return s.Status.Provisioned
}

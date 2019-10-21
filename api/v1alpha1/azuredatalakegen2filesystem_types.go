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

// AzureDataLakeGen2FileSystemSpec defines the desired state of AzureDataLakeGen2FileSystem
type AzureDataLakeGen2FileSystemSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	StorageAccountName string `json:"storageAccountName,omitempty"`
	ResourceGroupName  string `json:"resourceGroup"`
}

// AzureDataLakeGen2FileSystemStatus defines the observed state of AzureDataLakeGen2FileSystem
type AzureDataLakeGen2FileSystemStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Provisioning bool `json:"provisioning,omitempty"`
	Provisioned  bool `json:"provisioned,omitempty"`
}

// +kubebuilder:object:root=true

// AzureDataLakeGen2FileSystem is the Schema for the azuredatalakegen2filesystems API
type AzureDataLakeGen2FileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureDataLakeGen2FileSystemSpec   `json:"spec,omitempty"`
	Status AzureDataLakeGen2FileSystemStatus `json:"status,omitempty"`
	Output AzureDataLakeGen2FileSystemOutput `json:"output,omitempty"`
}

// AzureDataLakeGen2FileSystemOutput is the object that contains the output from creating and AdlsGen2 object
type AzureDataLakeGen2FileSystemOutput struct {
	AzureDataLakeGen2FileSystemName string `json:"AzureDataLakeGen2FileSystemName,omitempty"`
	Key1                            string `json:"key1,omitempty"`
	Key2                            string `json:"key2,omitempty"`
	ConnectionString1               string `json:"connectionString1,omitempty"`
	ConnectionString2               string `json:"connectionString2,omitempty"`
}

// +kubebuilder:object:root=true

// AzureDataLakeGen2FileSystemList contains a list of AzureDataLakeGen2FileSystem
type AzureDataLakeGen2FileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureDataLakeGen2FileSystem `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureDataLakeGen2FileSystem{}, &AzureDataLakeGen2FileSystemList{})
}

// IsSubmitted checks to see if resource has been successfully submitted for creation
func (adlsGen2FileSystem *AzureDataLakeGen2FileSystem) IsSubmitted() bool {
	return adlsGen2FileSystem.Status.Provisioning || adlsGen2FileSystem.Status.Provisioned
}
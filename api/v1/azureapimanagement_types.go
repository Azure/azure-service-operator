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
	helpers "github.com/Azure/azure-service-operator/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// APIManagementSpec defines the desired state of APIManagement
type APIManagementSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// APIManagementStatus defines the observed state of APIManagement
type APIManagementStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Provisioning bool `json:"provisioning,omitempty"`
	Provisioned  bool `json:"provisioned,omitempty"`
}

// +kubebuilder:object:root=true

// APIManagement is the Schema for the apimanagements API
type APIManagement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   APIManagementSpec   `json:"spec,omitempty"`
	Status APIManagementStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIManagementList contains a list of APIManagement
type APIManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIManagement `json:"items"`
}

func init() {
	SchemeBuilder.Register(&APIManagement{}, &APIManagementList{})
}

func (apimanagement *APIManagement) IsBeingDeleted() bool {
	return !apimanagement.ObjectMeta.DeletionTimestamp.IsZero()
}

func (apimanagement *APIManagement) IsSubmitted() bool {
	return apimanagement.Status.Provisioning || apimanagement.Status.Provisioned
}

func (apimanagement *APIManagement) HasFinalizer(finalizerName string) bool {
	return helpers.ContainsString(apimanagement.ObjectMeta.Finalizers, finalizerName)
}

func (apimanagement *APIManagement) AddFinalizer(finalizerName string) {
	apimanagement.ObjectMeta.Finalizers = append(apimanagement.ObjectMeta.Finalizers, finalizerName)
}

func (apimanagement *APIManagement) RemoveFinalizer(finalizerName string) {
	apimanagement.ObjectMeta.Finalizers = helpers.RemoveString(apimanagement.ObjectMeta.Finalizers, finalizerName)
}

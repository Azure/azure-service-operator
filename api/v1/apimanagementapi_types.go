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
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ApiManagementAPISpec defines the desired state of ApiManagementAPI
type ApiManagementAPISpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// ApiManagementAPIStatus defines the observed state of ApiManagementAPI
type ApiManagementAPIStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Provisioning bool `json:"provisioning,omitempty"`
	Provisioned  bool `json:"provisioned,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementAPI is the Schema for the apimanagementapis API
type ApiManagementAPI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApiManagementAPISpec   `json:"spec,omitempty"`
	Status ApiManagementAPIStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementAPIList contains a list of ApiManagementAPI
type ApiManagementAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiManagementAPI `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApiManagementAPI{}, &ApiManagementAPIList{})
}

func (apimanagementapi *ApiManagementAPI) IsBeingDeleted() bool {
	return !apimanagementapi.ObjectMeta.DeletionTimestamp.IsZero()
}

func (apimanagementapi *ApiManagementAPI) IsSubmitted() bool {
	return apimanagementapi.Status.Provisioning || apimanagementapi.Status.Provisioned
}

func (apimanagementapi *ApiManagementAPI) HasFinalizer(finalizerName string) bool {
	return helpers.ContainsString(apimanagementapi.ObjectMeta.Finalizers, finalizerName)
}

func (apimanagementapi *ApiManagementAPI) AddFinalizer(finalizerName string) {
	apimanagementapi.ObjectMeta.Finalizers = append(apimanagementapi.ObjectMeta.Finalizers, finalizerName)
}

func (apimanagementapi *ApiManagementAPI) RemoveFinalizer(finalizerName string) {
	apimanagementapi.ObjectMeta.Finalizers = helpers.RemoveString(apimanagementapi.ObjectMeta.Finalizers, finalizerName)
}

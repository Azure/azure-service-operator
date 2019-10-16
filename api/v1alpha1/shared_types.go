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
	"github.com/Azure/azure-service-operator/pkg/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ProvisionState enumerates the values for provisioning state.
// +kubebuilder:validation:Enum=Pending;Provisioning;Verifying;Succeeded;Failed
type ProvisionState string

const (
	Pending      ProvisionState = "Pending"
	Provisioning ProvisionState = "Provisioning"
	Verifying    ProvisionState = "Verifying"
	Succeeded    ProvisionState = "Succeeded"
	Failed       ProvisionState = "Failed"
)

// ResourceStatus defines the observed state of ResourceGroup
type ResourceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Provisioning   bool           `json:"provisioning,omitempty"`
	Provisioned    bool           `json:"provisioned,omitempty"`
	ProvisionState ProvisionState `json:"provisionState,omitempty"`
}

type ResourceBaseState struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status ResourceStatus `json:"status,omitempty"`
}

func (resourceGroup *ResourceBaseState) IsBeingDeleted() bool {
	return !resourceGroup.ObjectMeta.DeletionTimestamp.IsZero()
}

func (resourceGroup *ResourceBaseState) IsSubmitted() bool {
	return resourceGroup.Status.Provisioning || resourceGroup.Status.Provisioned

}

func (resourceGroup *ResourceBaseState) HasFinalizer(finalizerName string) bool {
	return helpers.ContainsString(resourceGroup.ObjectMeta.Finalizers, finalizerName)
}

func (resourceGroup *ResourceBaseState) AddFinalizer(finalizerName string) {
	resourceGroup.ObjectMeta.Finalizers = append(resourceGroup.ObjectMeta.Finalizers, finalizerName)
}

func (resourceGroup *ResourceBaseState) RemoveFinalizer(finalizerName string) {
	resourceGroup.ObjectMeta.Finalizers = helpers.RemoveString(resourceGroup.ObjectMeta.Finalizers, finalizerName)
}

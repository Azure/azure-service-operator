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

// ProvisionState enumerates the values for provisioning state.
// +kubebuilder:validation:Enum=Pending;Creating;Updating;Verifying;PostProvisioning;Succeeded;Failed;Terminating
type ProvisionState string

const (
	Pending     ProvisionState = "Pending"
	Creating    ProvisionState = "Creating"
	Updating    ProvisionState = "Updating"
	Verifying   ProvisionState = "Verifying"
	PostProvisioning   ProvisionState = "PostProvisioning"
	Succeeded   ProvisionState = "Succeeded"
	Failed      ProvisionState = "Failed"
	Terminating ProvisionState = "Terminating"
)

// ResourceStatus defines the observed state of any resources
type ResourceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ProvisionState ProvisionState `json:"provisionState,omitempty"`
	// Deprecated fields - to be removed
	Provisioning bool `json:"provisioning,omitempty"`
	Provisioned  bool `json:"provisioned,omitempty"`
}

func (s ProvisionState) IsPending() bool     { return s == Pending }
func (s ProvisionState) IsCreating() bool    { return s == Creating }
func (s ProvisionState) IsUpdating() bool    { return s == Updating }
func (s ProvisionState) IsVerifying() bool   { return s == Verifying }
func (s ProvisionState) IsPostProvisioning() bool   { return s == PostProvisioning }
func (s ProvisionState) IsSucceeded() bool   { return s == Succeeded }
func (s ProvisionState) IsFailed() bool      { return s == Failed }
func (s ProvisionState) IsTerminating() bool { return s == Terminating }

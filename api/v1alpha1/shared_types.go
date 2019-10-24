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

func (s *ASOStatus) IsPending() bool     { return ProvisionState(s.State) == Pending }
func (s *ASOStatus) IsCreating() bool    { return ProvisionState(s.State) == Creating }
func (s *ASOStatus) IsUpdating() bool    { return ProvisionState(s.State) == Updating }
func (s *ASOStatus) IsVerifying() bool   { return ProvisionState(s.State) == Verifying }
func (s *ASOStatus) IsPostProvisioning() bool   { return ProvisionState(s.State) == PostProvisioning }
func (s *ASOStatus) IsSucceeded() bool   { return ProvisionState(s.State) == Succeeded }
func (s *ASOStatus) IsFailed() bool      { return ProvisionState(s.State) == Failed }
func (s *ASOStatus) IsTerminating() bool { return ProvisionState(s.State) == Terminating }

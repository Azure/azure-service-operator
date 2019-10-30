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

package reconciler

type ProvisionState string

const (
	Pending          ProvisionState = "Pending"
	Creating         ProvisionState = "Creating"
	Updating         ProvisionState = "Updating"
	Verifying        ProvisionState = "Verifying"
	PostProvisioning ProvisionState = "PostProvisioning"
	Succeeded        ProvisionState = "Succeeded"
	Recreating       ProvisionState = "Recreating"
	Failed           ProvisionState = "Failed"
	Terminating      ProvisionState = "Terminating"
)

func (s *Status) IsPending() bool          { return s.State == Pending }
func (s *Status) IsCreating() bool         { return s.State == Creating }
func (s *Status) IsUpdating() bool         { return s.State == Updating }
func (s *Status) IsVerifying() bool        { return s.State == Verifying }
func (s *Status) IsPostProvisioning() bool { return s.State == PostProvisioning }
func (s *Status) IsSucceeded() bool        { return s.State == Succeeded }
func (s *Status) IsRecreating() bool       { return s.State == Terminating }
func (s *Status) IsFailed() bool           { return s.State == Failed }
func (s *Status) IsTerminating() bool      { return s.State == Terminating }

type Status struct {
	State         ProvisionState
	Message       string
	StatusPayload interface{}
}

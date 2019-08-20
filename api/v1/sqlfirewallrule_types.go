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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SqlFirewallRuleSpec defines the desired state of SqlFirewallRule
type SqlFirewallRuleSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location          string `json:"location"`
	ResourceGroupName string `json:"resourcegroup,omitempty"`
}

// SqlFirewallRuleStatus defines the observed state of SqlFirewallRule
type SqlFirewallRuleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Provisioning bool `json:"provisioning,omitempty"`
	Provisioned  bool `json:"provisioned,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// SqlFirewallRule is the Schema for the sqlfirewallrules API
type SqlFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SqlFirewallRuleSpec   `json:"spec,omitempty"`
	Status SqlFirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SqlFirewallRuleList contains a list of SqlFirewallRule
type SqlFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlFirewallRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SqlFirewallRule{}, &SqlFirewallRuleList{})
}

func (s *SqlFirewallRule) IsSubmitted() bool {
	return s.Status.Provisioning || s.Status.Provisioned
}

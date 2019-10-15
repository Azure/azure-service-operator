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

// AzureSqlFirewallRuleSpec defines the desired state of AzureSqlFirewallRule
type AzureSqlFirewallRuleSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ResourceGroup  string `json:"resourcegroup,omitempty"`
	Server         string `json:"server"`
	StartIPAddress string `json:"startipaddress,omitempty"`
	EndIPAddress   string `json:"endipaddress,omitempty"`
}

// AzureSqlFirewallRuleStatus defines the observed state of AzureSqlFirewallRule
type AzureSqlFirewallRuleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Provisioning bool `json:"provisioning,omitempty"`
	Provisioned  bool `json:"provisioned,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// AzureSqlFirewallRule is the Schema for the azuresqlfirewallrules API
type AzureSqlFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureSqlFirewallRuleSpec   `json:"spec,omitempty"`
	Status AzureSqlFirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureSqlFirewallRuleList contains a list of AzureSqlFirewallRule
type AzureSqlFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureSqlFirewallRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureSqlFirewallRule{}, &AzureSqlFirewallRuleList{})
}

func (s *AzureSqlFirewallRule) IsSubmitted() bool {
	return s.Status.Provisioning || s.Status.Provisioned
}

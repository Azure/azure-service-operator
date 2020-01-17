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

// KeyVaultSpec defines the desired state of KeyVault
type KeyVaultSpec struct {
	Location        string               `json:"location"`
	ResourceGroup   string               `json:"resourceGroup"`
	NetworkPolicies *NetworkRuleSet      `json:"networkPolicies,omitempty"`
	AccessPolicies  *[]AccessPolicyEntry `json:"accessPolicies,omitempty"`
}

type NetworkRuleSet struct {
	// Bypass - Tells what traffic can bypass network rules. This can be 'AzureServices' or 'None'.  If not specified the default is 'AzureServices'. Possible values include: 'AzureServices', 'None'
	Bypass string `json:"bypass,omitempty"`
	// DefaultAction - The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after the bypass property has been evaluated. Possible values include: 'Allow', 'Deny'
	DefaultAction string `json:"defaultAction,omitempty"`
	// IPRules - The list of IP address rules.
	IPRules *[]string `json:"ipRules,omitempty"`
	// VirtualNetworkRules - The list of virtual network rules.
	VirtualNetworkRules *[]string `json:"virtualNetworkRules,omitempty"`
}

type AccessPolicyEntry struct {
	// TenantID - The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantID string `json:"tenantId,omitempty"`
	// ObjectID - The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault. The object ID must be unique for the list of access policies.
	ObjectID string `json:"objectId,omitempty"`
	// ApplicationID -  Application ID of the client making request on behalf of a principal
	ApplicationID string `json:"applicationId,omitempty"`
	// Permissions - Permissions the identity has for keys, secrets and certificates.
	Permissions *Permissions `json:"permissions,omitempty"`
}

type Permissions struct {
	Keys         *[]string `json:"keys,omitempty"`
	Secrets      *[]string `json:"secrets,omitempty"`
	Certificates *[]string `json:"certificates,omitempty"`
	Storage      *[]string `json:"storage,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KeyVault is the Schema for the keyvaults API
type KeyVault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeyVaultSpec `json:"spec,omitempty"`
	Status ASOStatus    `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KeyVaultList contains a list of KeyVault
type KeyVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyVault `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeyVault{}, &KeyVaultList{})
}

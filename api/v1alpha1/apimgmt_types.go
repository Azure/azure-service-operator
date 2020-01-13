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

// ApimgmtSpec defines the desired state of Apimgmt
type ApimgmtSpec struct {
	Location             string   `json:"location"`
	ResourceGroup        string   `json:"resourceGroup"`
	Description          string   `json:"description,omitempty"`
	DisplayName          string   `json:"displayName,omitempty"`
	Format               string   `json:"format,omitempty"`
	APIType              string   `json:"type,omitempty"`
	APIRevision          string   `json:"apiRevision,omitempty"`
	APIVersion           string   `json:"apiVersion,omitempty"`
	IsCurrent            bool     `json:"isCurrent,omitempty"`
	Path                 string   `json:"path,omitempty"`
	Protocols            []string `json:"protocols,omitempty"`
	ServiceURL           string   `json:"serviceUrl,omitempty"`
	SubscriptionRequired bool     `json:"subscriptionRequired,omitempty"`
}

// ApimgmtStatus defines the observed state of Apimgmt
type ApimgmtStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Apimgmt is the Schema for the apimgmts API
type Apimgmt struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApimgmtSpec `json:"spec,omitempty"`
	Status ASOStatus   `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApimgmtList contains a list of Apimgmt
type ApimgmtList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apimgmt `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Apimgmt{}, &ApimgmtList{})
}

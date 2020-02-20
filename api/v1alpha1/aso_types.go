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

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// AzureServiceOperatorsStatus (ASOStatus) defines the observed state of resource actions
type ASOStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Provisioning   bool   `json:"provisioning,omitempty"`
	Provisioned    bool   `json:"provisioned,omitempty"`
	State          string `json:"state,omitempty"`
	Message        string `json:"message,omitempty"`
	SpecHash       uint64 `json:"specHash,omitempty"`
	ContainsUpdate bool   `json:"containsUpdate,omitempty"`
}

// StatusedObject used to unmarshall runtime.Object when we need Status
type StatusedObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status ASOStatus `json:"status,omitempty"`
}

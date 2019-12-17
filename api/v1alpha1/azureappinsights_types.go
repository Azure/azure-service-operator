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
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AppInsightsSpec is the k8s specification for Azure Application Insights
type AppInsightsSpec struct {
	Location      string `json:"location"`
	Namespace     string `json:"namespace,omitempty"`
	ResourceGroup string `json:"resourceGroup,omitempty"`
}

// AppInsights is the operator reflection type for Azure Application Insights
type AppInsights struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppInsightsSpec `json:"spec,omitempty"`
	Status ASOStatus       `json:"status,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Eventhub{}, &EventhubList{})
}

// IsBeingDeleted determines if the AppInsights operator is being deleted
func (appinsights *AppInsights) IsBeingDeleted() bool {
	return !appinsights.ObjectMeta.DeletionTimestamp.IsZero()
}

// IsSubmitted determines if the AppInsights operator has been submitted to the k8s control plane
func (appinsights *AppInsights) IsSubmitted() bool {
	return appinsights.Status.Provisioning || appinsights.Status.Provisioned
}

// HasFinalizer determines if the AppInsights operator has a k8s finalizer
func (appinsights *AppInsights) HasFinalizer(finalizerName string) bool {
	return helpers.ContainsString(appinsights.ObjectMeta.Finalizers, finalizerName)
}

// AddFinalizer adds a finalizer for the AppInsights operator
func (appinsights *AppInsights) AddFinalizer(finalizerName string) {
	appinsights.ObjectMeta.Finalizers = append(appinsights.ObjectMeta.Finalizers, finalizerName)
}

// RemoveFinalizer removes the finalizer for the AppInsights operator
func (appinsights *AppInsights) RemoveFinalizer(finalizerName string) {
	appinsights.ObjectMeta.Finalizers = helpers.RemoveString(appinsights.ObjectMeta.Finalizers, finalizerName)
}

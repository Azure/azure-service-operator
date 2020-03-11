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

// APIMgmtSpec defines the desired state of APIMgmt
type APIMgmtSpec struct {
	Location      string        `json:"location"`
	ResourceGroup string        `json:"resourceGroup"`
	APIService    string        `json:"apiService"`
	APIId         string        `json:"apiId"`
	Properties    APIProperties `json:"properties"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type APIMgmtAPI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   APIMgmtSpec `json:"spec,omitempty"`
	Status ASOStatus   `json:"status,omitempty"`
}

type APIVersionSet struct {
	// ID - Identifier for existing API Version Set. Omit this value to create a new Version Set.
	ID string `json:"id,omitempty"`
	// Name - The display Name of the API Version Set.
	Name string `json:"name,omitempty"`
	// Description - Description of API Version Set.
	Description string `json:"description,omitempty"`
}

type APIProperties struct {
	/* Format - Format of the Content in which the API is getting imported. Possible values include:
	'WadlXML', 'WadlLinkJSON', 'SwaggerJSON', 'SwaggerLinkJSON', 'Wsdl', 'WsdlLink', 'Openapi', 'Openapijson', 'OpenapiLink' */
	Format string `json:"format,omitempty"`
	// SourceAPIID - API identifier of the source API.
	SourceAPIID string `json:"sourceApiId,omitempty"`
	// DisplayName - API name. Must be 1 to 300 characters long.
	DisplayName string `json:"displayName,omitempty"`
	// ServiceURL - Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
	ServiceURL string `json:"serviceUrl,omitempty"`
	/* Path - Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance.
	It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API. */
	Path string `json:"path,omitempty"`
	// Protocols - Describes on which protocols the operations in this API can be invoked.
	Protocols []string `json:"protocols,omitempty"`
	// Description - Description of the API. May include HTML formatting tags.
	Description string `json:"description,omitempty"`
	// APIRevision - Describes the Revision of the Api. If no value is provided, default revision 1 is created
	APIRevision string `json:"apiRevision,omitempty"`
	// APIVersion - Indicates the Version identifier of the API if the API is versioned
	APIVersion string `json:"apiVersion,omitempty"`
	// IsCurrent - Indicates if API revision is current api revision.
	IsCurrent bool `json:"isCurrent,omitempty"`
	// IsOnline - READ-ONLY; Indicates if API revision is accessible via the gateway.
	IsOnline bool `json:"isOnline,omitempty"`
	// APIRevisionDescription - Description of the Api Revision.
	APIRevisionDescription string `json:"apiRevisionDescription,omitempty"`
	// APIVersionDescription - Description of the Api Version.
	APIVersionDescription string `json:"apiVersionDescription,omitempty"`
	// APIVersionSet - APIVersionSetContractDetails an API Version Set contains the common configuration for a set of API versions.
	APIVersionSet APIVersionSet `json:"apiVersionSets,omitempty"`
	// APIVersionSetID - A resource identifier for the related ApiVersionSet.
	APIVersionSetID string `json:"apiVersionSetId,omitempty"`
	// SubscriptionRequired - Specifies whether an API or Product subscription is required for accessing the API.
	SubscriptionRequired bool `json:"subscriptionRequired,omitempty"`
}

// +kubebuilder:object:root=true
// APIMgmtList contains a list of APIMgmt
type APIMgmtAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIMgmtAPI `json:"items"`
}

func init() {
	SchemeBuilder.Register(&APIMgmtAPI{}, &APIMgmtAPIList{})
}

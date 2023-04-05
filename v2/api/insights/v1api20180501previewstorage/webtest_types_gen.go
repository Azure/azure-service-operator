// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180501previewstorage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=insights.azure.com,resources=webtests,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=insights.azure.com,resources={webtests/status,webtests/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20180501preview.Webtest
// Generator information:
// - Generated from: /applicationinsights/resource-manager/Microsoft.Insights/preview/2018-05-01-preview/webTests_API.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}
type Webtest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Webtest_Spec   `json:"spec,omitempty"`
	Status            Webtest_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Webtest{}

// GetConditions returns the conditions of the resource
func (webtest *Webtest) GetConditions() conditions.Conditions {
	return webtest.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (webtest *Webtest) SetConditions(conditions conditions.Conditions) {
	webtest.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Webtest{}

// AzureName returns the Azure name of the resource
func (webtest *Webtest) AzureName() string {
	return webtest.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-05-01-preview"
func (webtest Webtest) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (webtest *Webtest) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (webtest *Webtest) GetSpec() genruntime.ConvertibleSpec {
	return &webtest.Spec
}

// GetStatus returns the status of this resource
func (webtest *Webtest) GetStatus() genruntime.ConvertibleStatus {
	return &webtest.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Insights/webtests"
func (webtest *Webtest) GetType() string {
	return "Microsoft.Insights/webtests"
}

// NewEmptyStatus returns a new empty (blank) status
func (webtest *Webtest) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Webtest_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (webtest *Webtest) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(webtest.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  webtest.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (webtest *Webtest) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Webtest_STATUS); ok {
		webtest.Status = *st
		return nil
	}

	// Convert status to required version
	var st Webtest_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	webtest.Status = st
	return nil
}

// Hub marks that this Webtest is the hub type for conversion
func (webtest *Webtest) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (webtest *Webtest) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: webtest.Spec.OriginalVersion,
		Kind:    "Webtest",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20180501preview.Webtest
// Generator information:
// - Generated from: /applicationinsights/resource-manager/Microsoft.Insights/preview/2018-05-01-preview/webTests_API.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}
type WebtestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Webtest `json:"items"`
}

// Storage version of v1api20180501preview.APIVersion
// +kubebuilder:validation:Enum={"2018-05-01-preview"}
type APIVersion string

const APIVersion_Value = APIVersion("2018-05-01-preview")

// Storage version of v1api20180501preview.Webtest_Spec
type Webtest_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                           `json:"azureName,omitempty"`
	Configuration   *WebTestProperties_Configuration `json:"Configuration,omitempty"`
	Description     *string                          `json:"Description,omitempty"`
	Enabled         *bool                            `json:"Enabled,omitempty"`
	Frequency       *int                             `json:"Frequency,omitempty"`
	Kind            *string                          `json:"Kind,omitempty"`
	Location        *string                          `json:"location,omitempty"`
	Locations       []WebTestGeolocation             `json:"Locations,omitempty"`
	Name            *string                          `json:"Name,omitempty"`
	OriginalVersion string                           `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner              *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag        genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Request            *WebTestProperties_Request         `json:"Request,omitempty"`
	RetryEnabled       *bool                              `json:"RetryEnabled,omitempty"`
	SyntheticMonitorId *string                            `json:"SyntheticMonitorId,omitempty"`
	Tags               map[string]string                  `json:"tags,omitempty"`
	Timeout            *int                               `json:"Timeout,omitempty"`
	ValidationRules    *WebTestProperties_ValidationRules `json:"ValidationRules,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Webtest_Spec{}

// ConvertSpecFrom populates our Webtest_Spec from the provided source
func (webtest *Webtest_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == webtest {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(webtest)
}

// ConvertSpecTo populates the provided destination from our Webtest_Spec
func (webtest *Webtest_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == webtest {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(webtest)
}

// Storage version of v1api20180501preview.Webtest_STATUS
type Webtest_STATUS struct {
	Conditions         []conditions.Condition                    `json:"conditions,omitempty"`
	Configuration      *WebTestProperties_Configuration_STATUS   `json:"Configuration,omitempty"`
	Description        *string                                   `json:"Description,omitempty"`
	Enabled            *bool                                     `json:"Enabled,omitempty"`
	Frequency          *int                                      `json:"Frequency,omitempty"`
	Id                 *string                                   `json:"id,omitempty"`
	Kind               *string                                   `json:"Kind,omitempty"`
	Location           *string                                   `json:"location,omitempty"`
	Locations          []WebTestGeolocation_STATUS               `json:"Locations,omitempty"`
	Name               *string                                   `json:"name,omitempty"`
	PropertiesName     *string                                   `json:"properties_name,omitempty"`
	PropertyBag        genruntime.PropertyBag                    `json:"$propertyBag,omitempty"`
	ProvisioningState  *string                                   `json:"provisioningState,omitempty"`
	Request            *WebTestProperties_Request_STATUS         `json:"Request,omitempty"`
	RetryEnabled       *bool                                     `json:"RetryEnabled,omitempty"`
	SyntheticMonitorId *string                                   `json:"SyntheticMonitorId,omitempty"`
	Tags               map[string]string                         `json:"tags,omitempty"`
	Timeout            *int                                      `json:"Timeout,omitempty"`
	Type               *string                                   `json:"type,omitempty"`
	ValidationRules    *WebTestProperties_ValidationRules_STATUS `json:"ValidationRules,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Webtest_STATUS{}

// ConvertStatusFrom populates our Webtest_STATUS from the provided source
func (webtest *Webtest_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == webtest {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(webtest)
}

// ConvertStatusTo populates the provided destination from our Webtest_STATUS
func (webtest *Webtest_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == webtest {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(webtest)
}

// Storage version of v1api20180501preview.WebTestGeolocation
// Geo-physical location to run a WebTest from. You must specify one or more locations for the test to run from.
type WebTestGeolocation struct {
	Id          *string                `json:"Id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20180501preview.WebTestGeolocation_STATUS
// Geo-physical location to run a WebTest from. You must specify one or more locations for the test to run from.
type WebTestGeolocation_STATUS struct {
	Id          *string                `json:"Id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20180501preview.WebTestProperties_Configuration
type WebTestProperties_Configuration struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	WebTest     *string                `json:"WebTest,omitempty"`
}

// Storage version of v1api20180501preview.WebTestProperties_Configuration_STATUS
type WebTestProperties_Configuration_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	WebTest     *string                `json:"WebTest,omitempty"`
}

// Storage version of v1api20180501preview.WebTestProperties_Request
type WebTestProperties_Request struct {
	FollowRedirects        *bool                  `json:"FollowRedirects,omitempty"`
	Headers                []HeaderField          `json:"Headers,omitempty"`
	HttpVerb               *string                `json:"HttpVerb,omitempty"`
	ParseDependentRequests *bool                  `json:"ParseDependentRequests,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequestBody            *string                `json:"RequestBody,omitempty"`
	RequestUrl             *string                `json:"RequestUrl,omitempty"`
}

// Storage version of v1api20180501preview.WebTestProperties_Request_STATUS
type WebTestProperties_Request_STATUS struct {
	FollowRedirects        *bool                  `json:"FollowRedirects,omitempty"`
	Headers                []HeaderField_STATUS   `json:"Headers,omitempty"`
	HttpVerb               *string                `json:"HttpVerb,omitempty"`
	ParseDependentRequests *bool                  `json:"ParseDependentRequests,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequestBody            *string                `json:"RequestBody,omitempty"`
	RequestUrl             *string                `json:"RequestUrl,omitempty"`
}

// Storage version of v1api20180501preview.WebTestProperties_ValidationRules
type WebTestProperties_ValidationRules struct {
	ContentValidation             *WebTestProperties_ValidationRules_ContentValidation `json:"ContentValidation,omitempty"`
	ExpectedHttpStatusCode        *int                                                 `json:"ExpectedHttpStatusCode,omitempty"`
	IgnoreHttpsStatusCode         *bool                                                `json:"IgnoreHttpsStatusCode,omitempty"`
	PropertyBag                   genruntime.PropertyBag                               `json:"$propertyBag,omitempty"`
	SSLCertRemainingLifetimeCheck *int                                                 `json:"SSLCertRemainingLifetimeCheck,omitempty"`
	SSLCheck                      *bool                                                `json:"SSLCheck,omitempty"`
}

// Storage version of v1api20180501preview.WebTestProperties_ValidationRules_STATUS
type WebTestProperties_ValidationRules_STATUS struct {
	ContentValidation             *WebTestProperties_ValidationRules_ContentValidation_STATUS `json:"ContentValidation,omitempty"`
	ExpectedHttpStatusCode        *int                                                        `json:"ExpectedHttpStatusCode,omitempty"`
	IgnoreHttpsStatusCode         *bool                                                       `json:"IgnoreHttpsStatusCode,omitempty"`
	PropertyBag                   genruntime.PropertyBag                                      `json:"$propertyBag,omitempty"`
	SSLCertRemainingLifetimeCheck *int                                                        `json:"SSLCertRemainingLifetimeCheck,omitempty"`
	SSLCheck                      *bool                                                       `json:"SSLCheck,omitempty"`
}

// Storage version of v1api20180501preview.HeaderField
// A header to add to the WebTest.
type HeaderField struct {
	Key         *string                `json:"key,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1api20180501preview.HeaderField_STATUS
// A header to add to the WebTest.
type HeaderField_STATUS struct {
	Key         *string                `json:"key,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1api20180501preview.WebTestProperties_ValidationRules_ContentValidation
type WebTestProperties_ValidationRules_ContentValidation struct {
	ContentMatch    *string                `json:"ContentMatch,omitempty"`
	IgnoreCase      *bool                  `json:"IgnoreCase,omitempty"`
	PassIfTextFound *bool                  `json:"PassIfTextFound,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20180501preview.WebTestProperties_ValidationRules_ContentValidation_STATUS
type WebTestProperties_ValidationRules_ContentValidation_STATUS struct {
	ContentMatch    *string                `json:"ContentMatch,omitempty"`
	IgnoreCase      *bool                  `json:"IgnoreCase,omitempty"`
	PassIfTextFound *bool                  `json:"PassIfTextFound,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Webtest{}, &WebtestList{})
}

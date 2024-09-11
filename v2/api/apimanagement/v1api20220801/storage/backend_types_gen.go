// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=apimanagement.azure.com,resources=backends,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apimanagement.azure.com,resources={backends/status,backends/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220801.Backend
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimbackends.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends/{backendId}
type Backend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Service_Backend_Spec   `json:"spec,omitempty"`
	Status            Service_Backend_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Backend{}

// GetConditions returns the conditions of the resource
func (backend *Backend) GetConditions() conditions.Conditions {
	return backend.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (backend *Backend) SetConditions(conditions conditions.Conditions) {
	backend.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Backend{}

// AzureName returns the Azure name of the resource
func (backend *Backend) AzureName() string {
	return backend.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-08-01"
func (backend Backend) GetAPIVersion() string {
	return "2022-08-01"
}

// GetResourceScope returns the scope of the resource
func (backend *Backend) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (backend *Backend) GetSpec() genruntime.ConvertibleSpec {
	return &backend.Spec
}

// GetStatus returns the status of this resource
func (backend *Backend) GetStatus() genruntime.ConvertibleStatus {
	return &backend.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (backend *Backend) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationHead,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/backends"
func (backend *Backend) GetType() string {
	return "Microsoft.ApiManagement/service/backends"
}

// NewEmptyStatus returns a new empty (blank) status
func (backend *Backend) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Service_Backend_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (backend *Backend) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(backend.Spec)
	return backend.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (backend *Backend) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Service_Backend_STATUS); ok {
		backend.Status = *st
		return nil
	}

	// Convert status to required version
	var st Service_Backend_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	backend.Status = st
	return nil
}

// Hub marks that this Backend is the hub type for conversion
func (backend *Backend) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (backend *Backend) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: backend.Spec.OriginalVersion,
		Kind:    "Backend",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220801.Backend
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimbackends.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends/{backendId}
type BackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Backend `json:"items"`
}

// Storage version of v1api20220801.Service_Backend_Spec
type Service_Backend_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                      `json:"azureName,omitempty"`
	Credentials     *BackendCredentialsContract `json:"credentials,omitempty"`
	Description     *string                     `json:"description,omitempty"`
	OriginalVersion string                      `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/Service resource
	Owner       *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"Service"`
	Properties  *BackendProperties                 `json:"properties,omitempty"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Protocol    *string                            `json:"protocol,omitempty"`
	Proxy       *BackendProxyContract              `json:"proxy,omitempty"`

	// ResourceReference: Management Uri of the Resource in External System. This URL can be the Arm Resource Id of Logic Apps,
	// Function Apps or API Apps.
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
	Title             *string                       `json:"title,omitempty"`
	Tls               *BackendTlsProperties         `json:"tls,omitempty"`
	Url               *string                       `json:"url,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Service_Backend_Spec{}

// ConvertSpecFrom populates our Service_Backend_Spec from the provided source
func (backend *Service_Backend_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == backend {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(backend)
}

// ConvertSpecTo populates the provided destination from our Service_Backend_Spec
func (backend *Service_Backend_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == backend {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(backend)
}

// Storage version of v1api20220801.Service_Backend_STATUS
type Service_Backend_STATUS struct {
	Conditions  []conditions.Condition             `json:"conditions,omitempty"`
	Credentials *BackendCredentialsContract_STATUS `json:"credentials,omitempty"`
	Description *string                            `json:"description,omitempty"`
	Id          *string                            `json:"id,omitempty"`
	Name        *string                            `json:"name,omitempty"`
	Properties  *BackendProperties_STATUS          `json:"properties,omitempty"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Protocol    *string                            `json:"protocol,omitempty"`
	Proxy       *BackendProxyContract_STATUS       `json:"proxy,omitempty"`
	ResourceId  *string                            `json:"resourceId,omitempty"`
	Title       *string                            `json:"title,omitempty"`
	Tls         *BackendTlsProperties_STATUS       `json:"tls,omitempty"`
	Type        *string                            `json:"type,omitempty"`
	Url         *string                            `json:"url,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Service_Backend_STATUS{}

// ConvertStatusFrom populates our Service_Backend_STATUS from the provided source
func (backend *Service_Backend_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == backend {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(backend)
}

// ConvertStatusTo populates the provided destination from our Service_Backend_STATUS
func (backend *Service_Backend_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == backend {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(backend)
}

// Storage version of v1api20220801.BackendCredentialsContract
// Details of the Credentials used to connect to Backend.
type BackendCredentialsContract struct {
	Authorization  *BackendAuthorizationHeaderCredentials `json:"authorization,omitempty"`
	Certificate    []string                               `json:"certificate,omitempty"`
	CertificateIds []string                               `json:"certificateIds,omitempty"`
	Header         map[string][]string                    `json:"header,omitempty"`
	PropertyBag    genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	Query          map[string][]string                    `json:"query,omitempty"`
}

// Storage version of v1api20220801.BackendCredentialsContract_STATUS
// Details of the Credentials used to connect to Backend.
type BackendCredentialsContract_STATUS struct {
	Authorization  *BackendAuthorizationHeaderCredentials_STATUS `json:"authorization,omitempty"`
	Certificate    []string                                      `json:"certificate,omitempty"`
	CertificateIds []string                                      `json:"certificateIds,omitempty"`
	Header         map[string][]string                           `json:"header,omitempty"`
	PropertyBag    genruntime.PropertyBag                        `json:"$propertyBag,omitempty"`
	Query          map[string][]string                           `json:"query,omitempty"`
}

// Storage version of v1api20220801.BackendProperties
// Properties specific to the Backend Type.
type BackendProperties struct {
	PropertyBag          genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	ServiceFabricCluster *BackendServiceFabricClusterProperties `json:"serviceFabricCluster,omitempty"`
}

// Storage version of v1api20220801.BackendProperties_STATUS
// Properties specific to the Backend Type.
type BackendProperties_STATUS struct {
	PropertyBag          genruntime.PropertyBag                        `json:"$propertyBag,omitempty"`
	ServiceFabricCluster *BackendServiceFabricClusterProperties_STATUS `json:"serviceFabricCluster,omitempty"`
}

// Storage version of v1api20220801.BackendProxyContract
// Details of the Backend WebProxy Server to use in the Request to Backend.
type BackendProxyContract struct {
	Password    *genruntime.SecretReference `json:"password,omitempty"`
	PropertyBag genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	Url         *string                     `json:"url,omitempty"`
	Username    *string                     `json:"username,omitempty"`
}

// Storage version of v1api20220801.BackendProxyContract_STATUS
// Details of the Backend WebProxy Server to use in the Request to Backend.
type BackendProxyContract_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Url         *string                `json:"url,omitempty"`
	Username    *string                `json:"username,omitempty"`
}

// Storage version of v1api20220801.BackendTlsProperties
// Properties controlling TLS Certificate Validation.
type BackendTlsProperties struct {
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ValidateCertificateChain *bool                  `json:"validateCertificateChain,omitempty"`
	ValidateCertificateName  *bool                  `json:"validateCertificateName,omitempty"`
}

// Storage version of v1api20220801.BackendTlsProperties_STATUS
// Properties controlling TLS Certificate Validation.
type BackendTlsProperties_STATUS struct {
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ValidateCertificateChain *bool                  `json:"validateCertificateChain,omitempty"`
	ValidateCertificateName  *bool                  `json:"validateCertificateName,omitempty"`
}

// Storage version of v1api20220801.BackendAuthorizationHeaderCredentials
// Authorization header information.
type BackendAuthorizationHeaderCredentials struct {
	Parameter   *string                `json:"parameter,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Scheme      *string                `json:"scheme,omitempty"`
}

// Storage version of v1api20220801.BackendAuthorizationHeaderCredentials_STATUS
// Authorization header information.
type BackendAuthorizationHeaderCredentials_STATUS struct {
	Parameter   *string                `json:"parameter,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Scheme      *string                `json:"scheme,omitempty"`
}

// Storage version of v1api20220801.BackendServiceFabricClusterProperties
// Properties of the Service Fabric Type Backend.
type BackendServiceFabricClusterProperties struct {
	ClientCertificateId           *string                `json:"clientCertificateId,omitempty"`
	ClientCertificatethumbprint   *string                `json:"clientCertificatethumbprint,omitempty"`
	ManagementEndpoints           []string               `json:"managementEndpoints,omitempty"`
	MaxPartitionResolutionRetries *int                   `json:"maxPartitionResolutionRetries,omitempty"`
	PropertyBag                   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ServerCertificateThumbprints  []string               `json:"serverCertificateThumbprints,omitempty"`
	ServerX509Names               []X509CertificateName  `json:"serverX509Names,omitempty"`
}

// Storage version of v1api20220801.BackendServiceFabricClusterProperties_STATUS
// Properties of the Service Fabric Type Backend.
type BackendServiceFabricClusterProperties_STATUS struct {
	ClientCertificateId           *string                      `json:"clientCertificateId,omitempty"`
	ClientCertificatethumbprint   *string                      `json:"clientCertificatethumbprint,omitempty"`
	ManagementEndpoints           []string                     `json:"managementEndpoints,omitempty"`
	MaxPartitionResolutionRetries *int                         `json:"maxPartitionResolutionRetries,omitempty"`
	PropertyBag                   genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	ServerCertificateThumbprints  []string                     `json:"serverCertificateThumbprints,omitempty"`
	ServerX509Names               []X509CertificateName_STATUS `json:"serverX509Names,omitempty"`
}

// Storage version of v1api20220801.X509CertificateName
// Properties of server X509Names.
type X509CertificateName struct {
	IssuerCertificateThumbprint *string                `json:"issuerCertificateThumbprint,omitempty"`
	Name                        *string                `json:"name,omitempty"`
	PropertyBag                 genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220801.X509CertificateName_STATUS
// Properties of server X509Names.
type X509CertificateName_STATUS struct {
	IssuerCertificateThumbprint *string                `json:"issuerCertificateThumbprint,omitempty"`
	Name                        *string                `json:"name,omitempty"`
	PropertyBag                 genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Backend{}, &BackendList{})
}

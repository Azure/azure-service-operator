// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101storage

import (
	"context"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +kubebuilder:rbac:groups=sql.azure.com,resources=servers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sql.azure.com,resources={servers/status,servers/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20211101.Server
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/Servers.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Server_Spec   `json:"spec,omitempty"`
	Status            Server_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Server{}

// GetConditions returns the conditions of the resource
func (server *Server) GetConditions() conditions.Conditions {
	return server.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (server *Server) SetConditions(conditions conditions.Conditions) {
	server.Status.Conditions = conditions
}

var _ genruntime.KubernetesExporter = &Server{}

// ExportKubernetesResources defines a resource which can create other resources in Kubernetes.
func (server *Server) ExportKubernetesResources(_ context.Context, _ genruntime.MetaObject, _ *genericarmclient.GenericClient, _ logr.Logger) ([]client.Object, error) {
	collector := configmaps.NewCollector(server.Namespace)
	if server.Spec.OperatorSpec != nil && server.Spec.OperatorSpec.ConfigMaps != nil {
		if server.Status.FullyQualifiedDomainName != nil {
			collector.AddValue(server.Spec.OperatorSpec.ConfigMaps.FullyQualifiedDomainName, *server.Status.FullyQualifiedDomainName)
		}
	}
	result, err := collector.Values()
	if err != nil {
		return nil, err
	}
	return configmaps.SliceToClientObjectSlice(result), nil
}

var _ genruntime.KubernetesResource = &Server{}

// AzureName returns the Azure name of the resource
func (server *Server) AzureName() string {
	return server.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (server Server) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (server *Server) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (server *Server) GetSpec() genruntime.ConvertibleSpec {
	return &server.Spec
}

// GetStatus returns the status of this resource
func (server *Server) GetStatus() genruntime.ConvertibleStatus {
	return &server.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers"
func (server *Server) GetType() string {
	return "Microsoft.Sql/servers"
}

// NewEmptyStatus returns a new empty (blank) status
func (server *Server) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Server_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (server *Server) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(server.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  server.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (server *Server) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Server_STATUS); ok {
		server.Status = *st
		return nil
	}

	// Convert status to required version
	var st Server_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	server.Status = st
	return nil
}

// Hub marks that this Server is the hub type for conversion
func (server *Server) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (server *Server) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: server.Spec.OriginalVersion,
		Kind:    "Server",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20211101.Server
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/Servers.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

// Storage version of v1beta20211101.APIVersion
// +kubebuilder:validation:Enum={"2021-11-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2021-11-01")

// Storage version of v1beta20211101.Server_Spec
type Server_Spec struct {
	AdministratorLogin         *string                      `json:"administratorLogin,omitempty"`
	AdministratorLoginPassword *genruntime.SecretReference  `json:"administratorLoginPassword,omitempty"`
	Administrators             *ServerExternalAdministrator `json:"administrators,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName         string              `json:"azureName,omitempty"`
	FederatedClientId *string             `json:"federatedClientId,omitempty"`
	Identity          *ResourceIdentity   `json:"identity,omitempty"`
	KeyId             *string             `json:"keyId,omitempty"`
	Location          *string             `json:"location,omitempty"`
	MinimalTlsVersion *string             `json:"minimalTlsVersion,omitempty"`
	OperatorSpec      *ServerOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion   string              `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`

	// PrimaryUserAssignedIdentityReference: The resource id of a user assigned identity to be used by default.
	PrimaryUserAssignedIdentityReference *genruntime.ResourceReference `armReference:"PrimaryUserAssignedIdentityId" json:"primaryUserAssignedIdentityReference,omitempty"`
	PropertyBag                          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	PublicNetworkAccess                  *string                       `json:"publicNetworkAccess,omitempty"`
	RestrictOutboundNetworkAccess        *string                       `json:"restrictOutboundNetworkAccess,omitempty"`
	Tags                                 map[string]string             `json:"tags,omitempty"`
	Version                              *string                       `json:"version,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Server_Spec{}

// ConvertSpecFrom populates our Server_Spec from the provided source
func (server *Server_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == server {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(server)
}

// ConvertSpecTo populates the provided destination from our Server_Spec
func (server *Server_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == server {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(server)
}

// Storage version of v1beta20211101.Server_STATUS
// An Azure SQL Database server.
type Server_STATUS struct {
	AdministratorLogin            *string                                  `json:"administratorLogin,omitempty"`
	Administrators                *ServerExternalAdministrator_STATUS      `json:"administrators,omitempty"`
	Conditions                    []conditions.Condition                   `json:"conditions,omitempty"`
	FederatedClientId             *string                                  `json:"federatedClientId,omitempty"`
	FullyQualifiedDomainName      *string                                  `json:"fullyQualifiedDomainName,omitempty"`
	Id                            *string                                  `json:"id,omitempty"`
	Identity                      *ResourceIdentity_STATUS                 `json:"identity,omitempty"`
	KeyId                         *string                                  `json:"keyId,omitempty"`
	Kind                          *string                                  `json:"kind,omitempty"`
	Location                      *string                                  `json:"location,omitempty"`
	MinimalTlsVersion             *string                                  `json:"minimalTlsVersion,omitempty"`
	Name                          *string                                  `json:"name,omitempty"`
	PrimaryUserAssignedIdentityId *string                                  `json:"primaryUserAssignedIdentityId,omitempty"`
	PrivateEndpointConnections    []ServerPrivateEndpointConnection_STATUS `json:"privateEndpointConnections,omitempty"`
	PropertyBag                   genruntime.PropertyBag                   `json:"$propertyBag,omitempty"`
	PublicNetworkAccess           *string                                  `json:"publicNetworkAccess,omitempty"`
	RestrictOutboundNetworkAccess *string                                  `json:"restrictOutboundNetworkAccess,omitempty"`
	State                         *string                                  `json:"state,omitempty"`
	Tags                          map[string]string                        `json:"tags,omitempty"`
	Type                          *string                                  `json:"type,omitempty"`
	Version                       *string                                  `json:"version,omitempty"`
	WorkspaceFeature              *string                                  `json:"workspaceFeature,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Server_STATUS{}

// ConvertStatusFrom populates our Server_STATUS from the provided source
func (server *Server_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == server {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(server)
}

// ConvertStatusTo populates the provided destination from our Server_STATUS
func (server *Server_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == server {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(server)
}

// Storage version of v1beta20211101.ResourceIdentity
// Azure Active Directory identity configuration for a resource.
type ResourceIdentity struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1beta20211101.ResourceIdentity_STATUS
// Azure Active Directory identity configuration for a resource.
type ResourceIdentity_STATUS struct {
	PrincipalId            *string                        `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	TenantId               *string                        `json:"tenantId,omitempty"`
	Type                   *string                        `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserIdentity_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1beta20211101.ServerExternalAdministrator
// Properties of a active directory administrator.
type ServerExternalAdministrator struct {
	AdministratorType         *string                `json:"administratorType,omitempty"`
	AzureADOnlyAuthentication *bool                  `json:"azureADOnlyAuthentication,omitempty"`
	Login                     *string                `json:"login,omitempty"`
	PrincipalType             *string                `json:"principalType,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Sid                       *string                `json:"sid,omitempty"`
	TenantId                  *string                `json:"tenantId,omitempty"`
}

// Storage version of v1beta20211101.ServerExternalAdministrator_STATUS
// Properties of a active directory administrator.
type ServerExternalAdministrator_STATUS struct {
	AdministratorType         *string                `json:"administratorType,omitempty"`
	AzureADOnlyAuthentication *bool                  `json:"azureADOnlyAuthentication,omitempty"`
	Login                     *string                `json:"login,omitempty"`
	PrincipalType             *string                `json:"principalType,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Sid                       *string                `json:"sid,omitempty"`
	TenantId                  *string                `json:"tenantId,omitempty"`
}

// Storage version of v1beta20211101.ServerOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type ServerOperatorSpec struct {
	ConfigMaps  *ServerOperatorConfigMaps `json:"configMaps,omitempty"`
	PropertyBag genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211101.ServerPrivateEndpointConnection_STATUS
// A private endpoint connection under a server
type ServerPrivateEndpointConnection_STATUS struct {
	Id          *string                                     `json:"id,omitempty"`
	Properties  *PrivateEndpointConnectionProperties_STATUS `json:"properties,omitempty"`
	PropertyBag genruntime.PropertyBag                      `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211101.PrivateEndpointConnectionProperties_STATUS
// Properties of a private endpoint connection.
type PrivateEndpointConnectionProperties_STATUS struct {
	GroupIds                          []string                                          `json:"groupIds,omitempty"`
	PrivateEndpoint                   *PrivateEndpointProperty_STATUS                   `json:"privateEndpoint,omitempty"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateProperty_STATUS `json:"privateLinkServiceConnectionState,omitempty"`
	PropertyBag                       genruntime.PropertyBag                            `json:"$propertyBag,omitempty"`
	ProvisioningState                 *string                                           `json:"provisioningState,omitempty"`
}

// Storage version of v1beta20211101.ServerOperatorConfigMaps
type ServerOperatorConfigMaps struct {
	FullyQualifiedDomainName *genruntime.ConfigMapDestination `json:"fullyQualifiedDomainName,omitempty"`
	PropertyBag              genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211101.UserIdentity_STATUS
// Azure Active Directory identity configuration for a resource.
type UserIdentity_STATUS struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211101.PrivateEndpointProperty_STATUS
type PrivateEndpointProperty_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211101.PrivateLinkServiceConnectionStateProperty_STATUS
type PrivateLinkServiceConnectionStateProperty_STATUS struct {
	ActionsRequired *string                `json:"actionsRequired,omitempty"`
	Description     *string                `json:"description,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status          *string                `json:"status,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}

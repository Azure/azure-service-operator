// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=dbformariadb.azure.com,resources=servers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=dbformariadb.azure.com,resources={servers/status,servers/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20180601.Server
// Generated from: https://schema.management.azure.com/schemas/2018-06-01/Microsoft.DBforMariaDB.json#/resourceDefinitions/servers
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_Spec  `json:"spec,omitempty"`
	Status            Server_Status `json:"status,omitempty"`
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

var _ genruntime.KubernetesResource = &Server{}

// AzureName returns the Azure name of the resource
func (server *Server) AzureName() string {
	return server.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-06-01"
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

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMariaDB/servers"
func (server *Server) GetType() string {
	return "Microsoft.DBforMariaDB/servers"
}

// NewEmptyStatus returns a new empty (blank) status
func (server *Server) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Server_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
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
	if st, ok := status.(*Server_Status); ok {
		server.Status = *st
		return nil
	}

	// Convert status to required version
	var st Server_Status
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
// Storage version of v1beta20180601.Server
// Generated from: https://schema.management.azure.com/schemas/2018-06-01/Microsoft.DBforMariaDB.json#/resourceDefinitions/servers
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

// Storage version of v1beta20180601.Server_Status
type Server_Status struct {
	AdministratorLogin         *string                                  `json:"administratorLogin,omitempty"`
	Conditions                 []conditions.Condition                   `json:"conditions,omitempty"`
	EarliestRestoreDate        *string                                  `json:"earliestRestoreDate,omitempty"`
	FullyQualifiedDomainName   *string                                  `json:"fullyQualifiedDomainName,omitempty"`
	Id                         *string                                  `json:"id,omitempty"`
	Location                   *string                                  `json:"location,omitempty"`
	MasterServerId             *string                                  `json:"masterServerId,omitempty"`
	MinimalTlsVersion          *string                                  `json:"minimalTlsVersion,omitempty"`
	Name                       *string                                  `json:"name,omitempty"`
	PrivateEndpointConnections []ServerPrivateEndpointConnection_Status `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag                   `json:"$propertyBag,omitempty"`
	PublicNetworkAccess        *string                                  `json:"publicNetworkAccess,omitempty"`
	ReplicaCapacity            *int                                     `json:"replicaCapacity,omitempty"`
	ReplicationRole            *string                                  `json:"replicationRole,omitempty"`
	Sku                        *Sku_Status                              `json:"sku,omitempty"`
	SslEnforcement             *string                                  `json:"sslEnforcement,omitempty"`
	StorageProfile             *StorageProfile_Status                   `json:"storageProfile,omitempty"`
	Tags                       map[string]string                        `json:"tags,omitempty"`
	Type                       *string                                  `json:"type,omitempty"`
	UserVisibleState           *string                                  `json:"userVisibleState,omitempty"`
	Version                    *string                                  `json:"version,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Server_Status{}

// ConvertStatusFrom populates our Server_Status from the provided source
func (server *Server_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == server {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(server)
}

// ConvertStatusTo populates the provided destination from our Server_Status
func (server *Server_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == server {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(server)
}

// Storage version of v1beta20180601.Servers_Spec
type Servers_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string              `json:"azureName,omitempty"`
	Location        *string             `json:"location,omitempty"`
	OperatorSpec    *ServerOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string              `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	Properties  *ServerPropertiesForCreate         `json:"properties,omitempty"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Sku         *Sku                               `json:"sku,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Servers_Spec{}

// ConvertSpecFrom populates our Servers_Spec from the provided source
func (servers *Servers_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == servers {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(servers)
}

// ConvertSpecTo populates the provided destination from our Servers_Spec
func (servers *Servers_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == servers {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(servers)
}

// Storage version of v1beta20180601.ServerOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type ServerOperatorSpec struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Secrets     *ServerOperatorSecrets `json:"secrets,omitempty"`
}

// Storage version of v1beta20180601.ServerPrivateEndpointConnection_Status
type ServerPrivateEndpointConnection_Status struct {
	Id          *string                                           `json:"id,omitempty"`
	Properties  *ServerPrivateEndpointConnectionProperties_Status `json:"properties,omitempty"`
	PropertyBag genruntime.PropertyBag                            `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20180601.ServerPropertiesForCreate
// Generated from: https://schema.management.azure.com/schemas/2018-06-01/Microsoft.DBforMariaDB.json#/definitions/ServerPropertiesForCreate
type ServerPropertiesForCreate struct {
	PropertyBag                      genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	ServerPropertiesForDefaultCreate *ServerPropertiesForDefaultCreate `json:"serverPropertiesForDefaultCreate,omitempty"`
	ServerPropertiesForGeoRestore    *ServerPropertiesForGeoRestore    `json:"serverPropertiesForGeoRestore,omitempty"`
	ServerPropertiesForReplica       *ServerPropertiesForReplica       `json:"serverPropertiesForReplica,omitempty"`
	ServerPropertiesForRestore       *ServerPropertiesForRestore       `json:"serverPropertiesForRestore,omitempty"`
}

// Storage version of v1beta20180601.Sku
// Generated from: https://schema.management.azure.com/schemas/2018-06-01/Microsoft.DBforMariaDB.json#/definitions/Sku
type Sku struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Family      *string                `json:"family,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Size        *string                `json:"size,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1beta20180601.Sku_Status
type Sku_Status struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Family      *string                `json:"family,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Size        *string                `json:"size,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1beta20180601.StorageProfile_Status
type StorageProfile_Status struct {
	BackupRetentionDays *int                   `json:"backupRetentionDays,omitempty"`
	GeoRedundantBackup  *string                `json:"geoRedundantBackup,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StorageAutogrow     *string                `json:"storageAutogrow,omitempty"`
	StorageMB           *int                   `json:"storageMB,omitempty"`
}

// Storage version of v1beta20180601.ServerOperatorSecrets
type ServerOperatorSecrets struct {
	FullyQualifiedDomainName *genruntime.SecretDestination `json:"fullyQualifiedDomainName,omitempty"`
	PropertyBag              genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20180601.ServerPrivateEndpointConnectionProperties_Status
type ServerPrivateEndpointConnectionProperties_Status struct {
	PrivateEndpoint                   *PrivateEndpointProperty_Status                         `json:"privateEndpoint,omitempty"`
	PrivateLinkServiceConnectionState *ServerPrivateLinkServiceConnectionStateProperty_Status `json:"privateLinkServiceConnectionState,omitempty"`
	PropertyBag                       genruntime.PropertyBag                                  `json:"$propertyBag,omitempty"`
	ProvisioningState                 *string                                                 `json:"provisioningState,omitempty"`
}

// Storage version of v1beta20180601.ServerPropertiesForDefaultCreate
type ServerPropertiesForDefaultCreate struct {
	AdministratorLogin         *string                     `json:"administratorLogin,omitempty"`
	AdministratorLoginPassword *genruntime.SecretReference `json:"administratorLoginPassword,omitempty"`
	CreateMode                 *string                     `json:"createMode,omitempty"`
	MinimalTlsVersion          *string                     `json:"minimalTlsVersion,omitempty"`
	PropertyBag                genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	PublicNetworkAccess        *string                     `json:"publicNetworkAccess,omitempty"`
	SslEnforcement             *string                     `json:"sslEnforcement,omitempty"`
	StorageProfile             *StorageProfile             `json:"storageProfile,omitempty"`
	Version                    *string                     `json:"version,omitempty"`
}

// Storage version of v1beta20180601.ServerPropertiesForGeoRestore
type ServerPropertiesForGeoRestore struct {
	CreateMode          *string                `json:"createMode,omitempty"`
	MinimalTlsVersion   *string                `json:"minimalTlsVersion,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublicNetworkAccess *string                `json:"publicNetworkAccess,omitempty"`
	SourceServerId      *string                `json:"sourceServerId,omitempty"`
	SslEnforcement      *string                `json:"sslEnforcement,omitempty"`
	StorageProfile      *StorageProfile        `json:"storageProfile,omitempty"`
	Version             *string                `json:"version,omitempty"`
}

// Storage version of v1beta20180601.ServerPropertiesForReplica
type ServerPropertiesForReplica struct {
	CreateMode          *string                `json:"createMode,omitempty"`
	MinimalTlsVersion   *string                `json:"minimalTlsVersion,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublicNetworkAccess *string                `json:"publicNetworkAccess,omitempty"`
	SourceServerId      *string                `json:"sourceServerId,omitempty"`
	SslEnforcement      *string                `json:"sslEnforcement,omitempty"`
	StorageProfile      *StorageProfile        `json:"storageProfile,omitempty"`
	Version             *string                `json:"version,omitempty"`
}

// Storage version of v1beta20180601.ServerPropertiesForRestore
type ServerPropertiesForRestore struct {
	CreateMode          *string                `json:"createMode,omitempty"`
	MinimalTlsVersion   *string                `json:"minimalTlsVersion,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublicNetworkAccess *string                `json:"publicNetworkAccess,omitempty"`
	RestorePointInTime  *string                `json:"restorePointInTime,omitempty"`
	SourceServerId      *string                `json:"sourceServerId,omitempty"`
	SslEnforcement      *string                `json:"sslEnforcement,omitempty"`
	StorageProfile      *StorageProfile        `json:"storageProfile,omitempty"`
	Version             *string                `json:"version,omitempty"`
}

// Storage version of v1beta20180601.PrivateEndpointProperty_Status
type PrivateEndpointProperty_Status struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20180601.ServerPrivateLinkServiceConnectionStateProperty_Status
type ServerPrivateLinkServiceConnectionStateProperty_Status struct {
	ActionsRequired *string                `json:"actionsRequired,omitempty"`
	Description     *string                `json:"description,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status          *string                `json:"status,omitempty"`
}

// Storage version of v1beta20180601.StorageProfile
// Generated from: https://schema.management.azure.com/schemas/2018-06-01/Microsoft.DBforMariaDB.json#/definitions/StorageProfile
type StorageProfile struct {
	BackupRetentionDays *int                   `json:"backupRetentionDays,omitempty"`
	GeoRedundantBackup  *string                `json:"geoRedundantBackup,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StorageAutogrow     *string                `json:"storageAutogrow,omitempty"`
	StorageMB           *int                   `json:"storageMB,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201storage

import (
	"fmt"
	v20201201s "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20201201storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1alpha1api20201201.RedisLinkedServer
// Deprecated version of RedisLinkedServer. Use v1beta20201201.RedisLinkedServer instead
type RedisLinkedServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisLinkedServer_Spec                 `json:"spec,omitempty"`
	Status            RedisLinkedServerWithProperties_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisLinkedServer{}

// GetConditions returns the conditions of the resource
func (server *RedisLinkedServer) GetConditions() conditions.Conditions {
	return server.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (server *RedisLinkedServer) SetConditions(conditions conditions.Conditions) {
	server.Status.Conditions = conditions
}

var _ conversion.Convertible = &RedisLinkedServer{}

// ConvertFrom populates our RedisLinkedServer from the provided hub RedisLinkedServer
func (server *RedisLinkedServer) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20201201s.RedisLinkedServer)
	if !ok {
		return fmt.Errorf("expected cache/v1beta20201201storage/RedisLinkedServer but received %T instead", hub)
	}

	return server.AssignPropertiesFromRedisLinkedServer(source)
}

// ConvertTo populates the provided hub RedisLinkedServer from our RedisLinkedServer
func (server *RedisLinkedServer) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20201201s.RedisLinkedServer)
	if !ok {
		return fmt.Errorf("expected cache/v1beta20201201storage/RedisLinkedServer but received %T instead", hub)
	}

	return server.AssignPropertiesToRedisLinkedServer(destination)
}

var _ genruntime.KubernetesResource = &RedisLinkedServer{}

// AzureName returns the Azure name of the resource
func (server *RedisLinkedServer) AzureName() string {
	return server.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (server RedisLinkedServer) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceKind returns the kind of the resource
func (server *RedisLinkedServer) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (server *RedisLinkedServer) GetSpec() genruntime.ConvertibleSpec {
	return &server.Spec
}

// GetStatus returns the status of this resource
func (server *RedisLinkedServer) GetStatus() genruntime.ConvertibleStatus {
	return &server.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/linkedServers"
func (server *RedisLinkedServer) GetType() string {
	return "Microsoft.Cache/redis/linkedServers"
}

// NewEmptyStatus returns a new empty (blank) status
func (server *RedisLinkedServer) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RedisLinkedServerWithProperties_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (server *RedisLinkedServer) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(server.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  server.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (server *RedisLinkedServer) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RedisLinkedServerWithProperties_STATUS); ok {
		server.Status = *st
		return nil
	}

	// Convert status to required version
	var st RedisLinkedServerWithProperties_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	server.Status = st
	return nil
}

// AssignPropertiesFromRedisLinkedServer populates our RedisLinkedServer from the provided source RedisLinkedServer
func (server *RedisLinkedServer) AssignPropertiesFromRedisLinkedServer(source *v20201201s.RedisLinkedServer) error {

	// ObjectMeta
	server.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RedisLinkedServer_Spec
	err := spec.AssignPropertiesFromRedisLinkedServer_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisLinkedServer_Spec() to populate field Spec")
	}
	server.Spec = spec

	// Status
	var status RedisLinkedServerWithProperties_STATUS
	err = status.AssignPropertiesFromRedisLinkedServerWithProperties_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisLinkedServerWithProperties_STATUS() to populate field Status")
	}
	server.Status = status

	// No error
	return nil
}

// AssignPropertiesToRedisLinkedServer populates the provided destination RedisLinkedServer from our RedisLinkedServer
func (server *RedisLinkedServer) AssignPropertiesToRedisLinkedServer(destination *v20201201s.RedisLinkedServer) error {

	// ObjectMeta
	destination.ObjectMeta = *server.ObjectMeta.DeepCopy()

	// Spec
	var spec v20201201s.RedisLinkedServer_Spec
	err := server.Spec.AssignPropertiesToRedisLinkedServer_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisLinkedServer_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20201201s.RedisLinkedServerWithProperties_STATUS
	err = server.Status.AssignPropertiesToRedisLinkedServerWithProperties_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisLinkedServerWithProperties_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (server *RedisLinkedServer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: server.Spec.OriginalVersion,
		Kind:    "RedisLinkedServer",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20201201.RedisLinkedServer
// Deprecated version of RedisLinkedServer. Use v1beta20201201.RedisLinkedServer instead
type RedisLinkedServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisLinkedServer `json:"items"`
}

// Storage version of v1alpha1api20201201.RedisLinkedServerWithProperties_STATUS
// Deprecated version of RedisLinkedServerWithProperties_STATUS. Use v1beta20201201.RedisLinkedServerWithProperties_STATUS instead
type RedisLinkedServerWithProperties_STATUS struct {
	Conditions               []conditions.Condition `json:"conditions,omitempty"`
	Id                       *string                `json:"id,omitempty"`
	LinkedRedisCacheId       *string                `json:"linkedRedisCacheId,omitempty"`
	LinkedRedisCacheLocation *string                `json:"linkedRedisCacheLocation,omitempty"`
	Name                     *string                `json:"name,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState        *string                `json:"provisioningState,omitempty"`
	ServerRole               *string                `json:"serverRole,omitempty"`
	Type                     *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisLinkedServerWithProperties_STATUS{}

// ConvertStatusFrom populates our RedisLinkedServerWithProperties_STATUS from the provided source
func (properties *RedisLinkedServerWithProperties_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20201201s.RedisLinkedServerWithProperties_STATUS)
	if ok {
		// Populate our instance from source
		return properties.AssignPropertiesFromRedisLinkedServerWithProperties_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20201201s.RedisLinkedServerWithProperties_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = properties.AssignPropertiesFromRedisLinkedServerWithProperties_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RedisLinkedServerWithProperties_STATUS
func (properties *RedisLinkedServerWithProperties_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20201201s.RedisLinkedServerWithProperties_STATUS)
	if ok {
		// Populate destination from our instance
		return properties.AssignPropertiesToRedisLinkedServerWithProperties_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20201201s.RedisLinkedServerWithProperties_STATUS{}
	err := properties.AssignPropertiesToRedisLinkedServerWithProperties_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignPropertiesFromRedisLinkedServerWithProperties_STATUS populates our RedisLinkedServerWithProperties_STATUS from the provided source RedisLinkedServerWithProperties_STATUS
func (properties *RedisLinkedServerWithProperties_STATUS) AssignPropertiesFromRedisLinkedServerWithProperties_STATUS(source *v20201201s.RedisLinkedServerWithProperties_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	properties.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	properties.Id = genruntime.ClonePointerToString(source.Id)

	// LinkedRedisCacheId
	properties.LinkedRedisCacheId = genruntime.ClonePointerToString(source.LinkedRedisCacheId)

	// LinkedRedisCacheLocation
	properties.LinkedRedisCacheLocation = genruntime.ClonePointerToString(source.LinkedRedisCacheLocation)

	// Name
	properties.Name = genruntime.ClonePointerToString(source.Name)

	// ProvisioningState
	properties.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// ServerRole
	properties.ServerRole = genruntime.ClonePointerToString(source.ServerRole)

	// Type
	properties.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		properties.PropertyBag = propertyBag
	} else {
		properties.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToRedisLinkedServerWithProperties_STATUS populates the provided destination RedisLinkedServerWithProperties_STATUS from our RedisLinkedServerWithProperties_STATUS
func (properties *RedisLinkedServerWithProperties_STATUS) AssignPropertiesToRedisLinkedServerWithProperties_STATUS(destination *v20201201s.RedisLinkedServerWithProperties_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(properties.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(properties.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(properties.Id)

	// LinkedRedisCacheId
	destination.LinkedRedisCacheId = genruntime.ClonePointerToString(properties.LinkedRedisCacheId)

	// LinkedRedisCacheLocation
	destination.LinkedRedisCacheLocation = genruntime.ClonePointerToString(properties.LinkedRedisCacheLocation)

	// Name
	destination.Name = genruntime.ClonePointerToString(properties.Name)

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(properties.ProvisioningState)

	// ServerRole
	destination.ServerRole = genruntime.ClonePointerToString(properties.ServerRole)

	// Type
	destination.Type = genruntime.ClonePointerToString(properties.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20201201.RedisLinkedServer_Spec
type RedisLinkedServer_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                string  `json:"azureName,omitempty"`
	LinkedRedisCacheLocation *string `json:"linkedRedisCacheLocation,omitempty"`

	// +kubebuilder:validation:Required
	LinkedRedisCacheReference *genruntime.ResourceReference `armReference:"LinkedRedisCacheId" json:"linkedRedisCacheReference,omitempty"`
	OriginalVersion           string                        `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cache.azure.com/Redis resource
	Owner       *genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner,omitempty" kind:"Redis"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ServerRole  *string                            `json:"serverRole,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RedisLinkedServer_Spec{}

// ConvertSpecFrom populates our RedisLinkedServer_Spec from the provided source
func (server *RedisLinkedServer_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20201201s.RedisLinkedServer_Spec)
	if ok {
		// Populate our instance from source
		return server.AssignPropertiesFromRedisLinkedServer_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20201201s.RedisLinkedServer_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = server.AssignPropertiesFromRedisLinkedServer_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RedisLinkedServer_Spec
func (server *RedisLinkedServer_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20201201s.RedisLinkedServer_Spec)
	if ok {
		// Populate destination from our instance
		return server.AssignPropertiesToRedisLinkedServer_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20201201s.RedisLinkedServer_Spec{}
	err := server.AssignPropertiesToRedisLinkedServer_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignPropertiesFromRedisLinkedServer_Spec populates our RedisLinkedServer_Spec from the provided source RedisLinkedServer_Spec
func (server *RedisLinkedServer_Spec) AssignPropertiesFromRedisLinkedServer_Spec(source *v20201201s.RedisLinkedServer_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	server.AzureName = source.AzureName

	// LinkedRedisCacheLocation
	server.LinkedRedisCacheLocation = genruntime.ClonePointerToString(source.LinkedRedisCacheLocation)

	// LinkedRedisCacheReference
	if source.LinkedRedisCacheReference != nil {
		linkedRedisCacheReference := source.LinkedRedisCacheReference.Copy()
		server.LinkedRedisCacheReference = &linkedRedisCacheReference
	} else {
		server.LinkedRedisCacheReference = nil
	}

	// OriginalVersion
	server.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		server.Owner = &owner
	} else {
		server.Owner = nil
	}

	// ServerRole
	server.ServerRole = genruntime.ClonePointerToString(source.ServerRole)

	// Update the property bag
	if len(propertyBag) > 0 {
		server.PropertyBag = propertyBag
	} else {
		server.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToRedisLinkedServer_Spec populates the provided destination RedisLinkedServer_Spec from our RedisLinkedServer_Spec
func (server *RedisLinkedServer_Spec) AssignPropertiesToRedisLinkedServer_Spec(destination *v20201201s.RedisLinkedServer_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(server.PropertyBag)

	// AzureName
	destination.AzureName = server.AzureName

	// LinkedRedisCacheLocation
	destination.LinkedRedisCacheLocation = genruntime.ClonePointerToString(server.LinkedRedisCacheLocation)

	// LinkedRedisCacheReference
	if server.LinkedRedisCacheReference != nil {
		linkedRedisCacheReference := server.LinkedRedisCacheReference.Copy()
		destination.LinkedRedisCacheReference = &linkedRedisCacheReference
	} else {
		destination.LinkedRedisCacheReference = nil
	}

	// OriginalVersion
	destination.OriginalVersion = server.OriginalVersion

	// Owner
	if server.Owner != nil {
		owner := server.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// ServerRole
	destination.ServerRole = genruntime.ClonePointerToString(server.ServerRole)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&RedisLinkedServer{}, &RedisLinkedServerList{})
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210301storage

import (
	"fmt"
	v20210301s "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20210301storage"
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
// Storage version of v1alpha1api20210301.RedisEnterpriseDatabase
// Deprecated version of RedisEnterpriseDatabase. Use v1beta20210301.RedisEnterpriseDatabase instead
type RedisEnterpriseDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisEnterpriseDatabase_Spec   `json:"spec,omitempty"`
	Status            RedisEnterpriseDatabase_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisEnterpriseDatabase{}

// GetConditions returns the conditions of the resource
func (database *RedisEnterpriseDatabase) GetConditions() conditions.Conditions {
	return database.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (database *RedisEnterpriseDatabase) SetConditions(conditions conditions.Conditions) {
	database.Status.Conditions = conditions
}

var _ conversion.Convertible = &RedisEnterpriseDatabase{}

// ConvertFrom populates our RedisEnterpriseDatabase from the provided hub RedisEnterpriseDatabase
func (database *RedisEnterpriseDatabase) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210301s.RedisEnterpriseDatabase)
	if !ok {
		return fmt.Errorf("expected cache/v1beta20210301storage/RedisEnterpriseDatabase but received %T instead", hub)
	}

	return database.AssignPropertiesFromRedisEnterpriseDatabase(source)
}

// ConvertTo populates the provided hub RedisEnterpriseDatabase from our RedisEnterpriseDatabase
func (database *RedisEnterpriseDatabase) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210301s.RedisEnterpriseDatabase)
	if !ok {
		return fmt.Errorf("expected cache/v1beta20210301storage/RedisEnterpriseDatabase but received %T instead", hub)
	}

	return database.AssignPropertiesToRedisEnterpriseDatabase(destination)
}

var _ genruntime.KubernetesResource = &RedisEnterpriseDatabase{}

// AzureName returns the Azure name of the resource
func (database *RedisEnterpriseDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-03-01"
func (database RedisEnterpriseDatabase) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (database *RedisEnterpriseDatabase) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (database *RedisEnterpriseDatabase) GetSpec() genruntime.ConvertibleSpec {
	return &database.Spec
}

// GetStatus returns the status of this resource
func (database *RedisEnterpriseDatabase) GetStatus() genruntime.ConvertibleStatus {
	return &database.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redisEnterprise/databases"
func (database *RedisEnterpriseDatabase) GetType() string {
	return "Microsoft.Cache/redisEnterprise/databases"
}

// NewEmptyStatus returns a new empty (blank) status
func (database *RedisEnterpriseDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RedisEnterpriseDatabase_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (database *RedisEnterpriseDatabase) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  database.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (database *RedisEnterpriseDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RedisEnterpriseDatabase_STATUS); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st RedisEnterpriseDatabase_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// AssignPropertiesFromRedisEnterpriseDatabase populates our RedisEnterpriseDatabase from the provided source RedisEnterpriseDatabase
func (database *RedisEnterpriseDatabase) AssignPropertiesFromRedisEnterpriseDatabase(source *v20210301s.RedisEnterpriseDatabase) error {

	// ObjectMeta
	database.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RedisEnterpriseDatabase_Spec
	err := spec.AssignPropertiesFromRedisEnterpriseDatabase_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisEnterpriseDatabase_Spec() to populate field Spec")
	}
	database.Spec = spec

	// Status
	var status RedisEnterpriseDatabase_STATUS
	err = status.AssignPropertiesFromRedisEnterpriseDatabase_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisEnterpriseDatabase_STATUS() to populate field Status")
	}
	database.Status = status

	// No error
	return nil
}

// AssignPropertiesToRedisEnterpriseDatabase populates the provided destination RedisEnterpriseDatabase from our RedisEnterpriseDatabase
func (database *RedisEnterpriseDatabase) AssignPropertiesToRedisEnterpriseDatabase(destination *v20210301s.RedisEnterpriseDatabase) error {

	// ObjectMeta
	destination.ObjectMeta = *database.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210301s.RedisEnterpriseDatabase_Spec
	err := database.Spec.AssignPropertiesToRedisEnterpriseDatabase_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisEnterpriseDatabase_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210301s.RedisEnterpriseDatabase_STATUS
	err = database.Status.AssignPropertiesToRedisEnterpriseDatabase_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisEnterpriseDatabase_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *RedisEnterpriseDatabase) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion,
		Kind:    "RedisEnterpriseDatabase",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20210301.RedisEnterpriseDatabase
// Deprecated version of RedisEnterpriseDatabase. Use v1beta20210301.RedisEnterpriseDatabase instead
type RedisEnterpriseDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisEnterpriseDatabase `json:"items"`
}

// Storage version of v1alpha1api20210301.RedisEnterpriseDatabase_Spec
type RedisEnterpriseDatabase_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName        string   `json:"azureName,omitempty"`
	ClientProtocol   *string  `json:"clientProtocol,omitempty"`
	ClusteringPolicy *string  `json:"clusteringPolicy,omitempty"`
	EvictionPolicy   *string  `json:"evictionPolicy,omitempty"`
	Modules          []Module `json:"modules,omitempty"`
	OriginalVersion  string   `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cache.azure.com/RedisEnterprise resource
	Owner       *genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner,omitempty" kind:"RedisEnterprise"`
	Persistence *Persistence                       `json:"persistence,omitempty"`
	Port        *int                               `json:"port,omitempty"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RedisEnterpriseDatabase_Spec{}

// ConvertSpecFrom populates our RedisEnterpriseDatabase_Spec from the provided source
func (database *RedisEnterpriseDatabase_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210301s.RedisEnterpriseDatabase_Spec)
	if ok {
		// Populate our instance from source
		return database.AssignPropertiesFromRedisEnterpriseDatabase_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210301s.RedisEnterpriseDatabase_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = database.AssignPropertiesFromRedisEnterpriseDatabase_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RedisEnterpriseDatabase_Spec
func (database *RedisEnterpriseDatabase_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210301s.RedisEnterpriseDatabase_Spec)
	if ok {
		// Populate destination from our instance
		return database.AssignPropertiesToRedisEnterpriseDatabase_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210301s.RedisEnterpriseDatabase_Spec{}
	err := database.AssignPropertiesToRedisEnterpriseDatabase_Spec(dst)
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

// AssignPropertiesFromRedisEnterpriseDatabase_Spec populates our RedisEnterpriseDatabase_Spec from the provided source RedisEnterpriseDatabase_Spec
func (database *RedisEnterpriseDatabase_Spec) AssignPropertiesFromRedisEnterpriseDatabase_Spec(source *v20210301s.RedisEnterpriseDatabase_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	database.AzureName = source.AzureName

	// ClientProtocol
	database.ClientProtocol = genruntime.ClonePointerToString(source.ClientProtocol)

	// ClusteringPolicy
	database.ClusteringPolicy = genruntime.ClonePointerToString(source.ClusteringPolicy)

	// EvictionPolicy
	database.EvictionPolicy = genruntime.ClonePointerToString(source.EvictionPolicy)

	// Modules
	if source.Modules != nil {
		moduleList := make([]Module, len(source.Modules))
		for moduleIndex, moduleItem := range source.Modules {
			// Shadow the loop variable to avoid aliasing
			moduleItem := moduleItem
			var module Module
			err := module.AssignPropertiesFromModule(&moduleItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromModule() to populate field Modules")
			}
			moduleList[moduleIndex] = module
		}
		database.Modules = moduleList
	} else {
		database.Modules = nil
	}

	// OriginalVersion
	database.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		database.Owner = &owner
	} else {
		database.Owner = nil
	}

	// Persistence
	if source.Persistence != nil {
		var persistence Persistence
		err := persistence.AssignPropertiesFromPersistence(source.Persistence)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromPersistence() to populate field Persistence")
		}
		database.Persistence = &persistence
	} else {
		database.Persistence = nil
	}

	// Port
	database.Port = genruntime.ClonePointerToInt(source.Port)

	// Update the property bag
	if len(propertyBag) > 0 {
		database.PropertyBag = propertyBag
	} else {
		database.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToRedisEnterpriseDatabase_Spec populates the provided destination RedisEnterpriseDatabase_Spec from our RedisEnterpriseDatabase_Spec
func (database *RedisEnterpriseDatabase_Spec) AssignPropertiesToRedisEnterpriseDatabase_Spec(destination *v20210301s.RedisEnterpriseDatabase_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(database.PropertyBag)

	// AzureName
	destination.AzureName = database.AzureName

	// ClientProtocol
	destination.ClientProtocol = genruntime.ClonePointerToString(database.ClientProtocol)

	// ClusteringPolicy
	destination.ClusteringPolicy = genruntime.ClonePointerToString(database.ClusteringPolicy)

	// EvictionPolicy
	destination.EvictionPolicy = genruntime.ClonePointerToString(database.EvictionPolicy)

	// Modules
	if database.Modules != nil {
		moduleList := make([]v20210301s.Module, len(database.Modules))
		for moduleIndex, moduleItem := range database.Modules {
			// Shadow the loop variable to avoid aliasing
			moduleItem := moduleItem
			var module v20210301s.Module
			err := moduleItem.AssignPropertiesToModule(&module)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToModule() to populate field Modules")
			}
			moduleList[moduleIndex] = module
		}
		destination.Modules = moduleList
	} else {
		destination.Modules = nil
	}

	// OriginalVersion
	destination.OriginalVersion = database.OriginalVersion

	// Owner
	if database.Owner != nil {
		owner := database.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Persistence
	if database.Persistence != nil {
		var persistence v20210301s.Persistence
		err := database.Persistence.AssignPropertiesToPersistence(&persistence)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToPersistence() to populate field Persistence")
		}
		destination.Persistence = &persistence
	} else {
		destination.Persistence = nil
	}

	// Port
	destination.Port = genruntime.ClonePointerToInt(database.Port)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210301.RedisEnterpriseDatabase_STATUS
// Deprecated version of RedisEnterpriseDatabase_STATUS. Use v1beta20210301.RedisEnterpriseDatabase_STATUS instead
type RedisEnterpriseDatabase_STATUS struct {
	ClientProtocol    *string                `json:"clientProtocol,omitempty"`
	ClusteringPolicy  *string                `json:"clusteringPolicy,omitempty"`
	Conditions        []conditions.Condition `json:"conditions,omitempty"`
	EvictionPolicy    *string                `json:"evictionPolicy,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Modules           []Module_STATUS        `json:"modules,omitempty"`
	Name              *string                `json:"name,omitempty"`
	Persistence       *Persistence_STATUS    `json:"persistence,omitempty"`
	Port              *int                   `json:"port,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	ResourceState     *string                `json:"resourceState,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisEnterpriseDatabase_STATUS{}

// ConvertStatusFrom populates our RedisEnterpriseDatabase_STATUS from the provided source
func (database *RedisEnterpriseDatabase_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210301s.RedisEnterpriseDatabase_STATUS)
	if ok {
		// Populate our instance from source
		return database.AssignPropertiesFromRedisEnterpriseDatabase_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210301s.RedisEnterpriseDatabase_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = database.AssignPropertiesFromRedisEnterpriseDatabase_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RedisEnterpriseDatabase_STATUS
func (database *RedisEnterpriseDatabase_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210301s.RedisEnterpriseDatabase_STATUS)
	if ok {
		// Populate destination from our instance
		return database.AssignPropertiesToRedisEnterpriseDatabase_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210301s.RedisEnterpriseDatabase_STATUS{}
	err := database.AssignPropertiesToRedisEnterpriseDatabase_STATUS(dst)
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

// AssignPropertiesFromRedisEnterpriseDatabase_STATUS populates our RedisEnterpriseDatabase_STATUS from the provided source RedisEnterpriseDatabase_STATUS
func (database *RedisEnterpriseDatabase_STATUS) AssignPropertiesFromRedisEnterpriseDatabase_STATUS(source *v20210301s.RedisEnterpriseDatabase_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ClientProtocol
	database.ClientProtocol = genruntime.ClonePointerToString(source.ClientProtocol)

	// ClusteringPolicy
	database.ClusteringPolicy = genruntime.ClonePointerToString(source.ClusteringPolicy)

	// Conditions
	database.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// EvictionPolicy
	database.EvictionPolicy = genruntime.ClonePointerToString(source.EvictionPolicy)

	// Id
	database.Id = genruntime.ClonePointerToString(source.Id)

	// Modules
	if source.Modules != nil {
		moduleList := make([]Module_STATUS, len(source.Modules))
		for moduleIndex, moduleItem := range source.Modules {
			// Shadow the loop variable to avoid aliasing
			moduleItem := moduleItem
			var module Module_STATUS
			err := module.AssignPropertiesFromModule_STATUS(&moduleItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromModule_STATUS() to populate field Modules")
			}
			moduleList[moduleIndex] = module
		}
		database.Modules = moduleList
	} else {
		database.Modules = nil
	}

	// Name
	database.Name = genruntime.ClonePointerToString(source.Name)

	// Persistence
	if source.Persistence != nil {
		var persistence Persistence_STATUS
		err := persistence.AssignPropertiesFromPersistence_STATUS(source.Persistence)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromPersistence_STATUS() to populate field Persistence")
		}
		database.Persistence = &persistence
	} else {
		database.Persistence = nil
	}

	// Port
	database.Port = genruntime.ClonePointerToInt(source.Port)

	// ProvisioningState
	database.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// ResourceState
	database.ResourceState = genruntime.ClonePointerToString(source.ResourceState)

	// Type
	database.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		database.PropertyBag = propertyBag
	} else {
		database.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToRedisEnterpriseDatabase_STATUS populates the provided destination RedisEnterpriseDatabase_STATUS from our RedisEnterpriseDatabase_STATUS
func (database *RedisEnterpriseDatabase_STATUS) AssignPropertiesToRedisEnterpriseDatabase_STATUS(destination *v20210301s.RedisEnterpriseDatabase_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(database.PropertyBag)

	// ClientProtocol
	destination.ClientProtocol = genruntime.ClonePointerToString(database.ClientProtocol)

	// ClusteringPolicy
	destination.ClusteringPolicy = genruntime.ClonePointerToString(database.ClusteringPolicy)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(database.Conditions)

	// EvictionPolicy
	destination.EvictionPolicy = genruntime.ClonePointerToString(database.EvictionPolicy)

	// Id
	destination.Id = genruntime.ClonePointerToString(database.Id)

	// Modules
	if database.Modules != nil {
		moduleList := make([]v20210301s.Module_STATUS, len(database.Modules))
		for moduleIndex, moduleItem := range database.Modules {
			// Shadow the loop variable to avoid aliasing
			moduleItem := moduleItem
			var module v20210301s.Module_STATUS
			err := moduleItem.AssignPropertiesToModule_STATUS(&module)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToModule_STATUS() to populate field Modules")
			}
			moduleList[moduleIndex] = module
		}
		destination.Modules = moduleList
	} else {
		destination.Modules = nil
	}

	// Name
	destination.Name = genruntime.ClonePointerToString(database.Name)

	// Persistence
	if database.Persistence != nil {
		var persistence v20210301s.Persistence_STATUS
		err := database.Persistence.AssignPropertiesToPersistence_STATUS(&persistence)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToPersistence_STATUS() to populate field Persistence")
		}
		destination.Persistence = &persistence
	} else {
		destination.Persistence = nil
	}

	// Port
	destination.Port = genruntime.ClonePointerToInt(database.Port)

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(database.ProvisioningState)

	// ResourceState
	destination.ResourceState = genruntime.ClonePointerToString(database.ResourceState)

	// Type
	destination.Type = genruntime.ClonePointerToString(database.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210301.Module
// Deprecated version of Module. Use v1beta20210301.Module instead
type Module struct {
	Args        *string                `json:"args,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignPropertiesFromModule populates our Module from the provided source Module
func (module *Module) AssignPropertiesFromModule(source *v20210301s.Module) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Args
	module.Args = genruntime.ClonePointerToString(source.Args)

	// Name
	module.Name = genruntime.ClonePointerToString(source.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		module.PropertyBag = propertyBag
	} else {
		module.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToModule populates the provided destination Module from our Module
func (module *Module) AssignPropertiesToModule(destination *v20210301s.Module) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(module.PropertyBag)

	// Args
	destination.Args = genruntime.ClonePointerToString(module.Args)

	// Name
	destination.Name = genruntime.ClonePointerToString(module.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210301.Module_STATUS
// Deprecated version of Module_STATUS. Use v1beta20210301.Module_STATUS instead
type Module_STATUS struct {
	Args        *string                `json:"args,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Version     *string                `json:"version,omitempty"`
}

// AssignPropertiesFromModule_STATUS populates our Module_STATUS from the provided source Module_STATUS
func (module *Module_STATUS) AssignPropertiesFromModule_STATUS(source *v20210301s.Module_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Args
	module.Args = genruntime.ClonePointerToString(source.Args)

	// Name
	module.Name = genruntime.ClonePointerToString(source.Name)

	// Version
	module.Version = genruntime.ClonePointerToString(source.Version)

	// Update the property bag
	if len(propertyBag) > 0 {
		module.PropertyBag = propertyBag
	} else {
		module.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToModule_STATUS populates the provided destination Module_STATUS from our Module_STATUS
func (module *Module_STATUS) AssignPropertiesToModule_STATUS(destination *v20210301s.Module_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(module.PropertyBag)

	// Args
	destination.Args = genruntime.ClonePointerToString(module.Args)

	// Name
	destination.Name = genruntime.ClonePointerToString(module.Name)

	// Version
	destination.Version = genruntime.ClonePointerToString(module.Version)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210301.Persistence
// Deprecated version of Persistence. Use v1beta20210301.Persistence instead
type Persistence struct {
	AofEnabled   *bool                  `json:"aofEnabled,omitempty"`
	AofFrequency *string                `json:"aofFrequency,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RdbEnabled   *bool                  `json:"rdbEnabled,omitempty"`
	RdbFrequency *string                `json:"rdbFrequency,omitempty"`
}

// AssignPropertiesFromPersistence populates our Persistence from the provided source Persistence
func (persistence *Persistence) AssignPropertiesFromPersistence(source *v20210301s.Persistence) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AofEnabled
	if source.AofEnabled != nil {
		aofEnabled := *source.AofEnabled
		persistence.AofEnabled = &aofEnabled
	} else {
		persistence.AofEnabled = nil
	}

	// AofFrequency
	persistence.AofFrequency = genruntime.ClonePointerToString(source.AofFrequency)

	// RdbEnabled
	if source.RdbEnabled != nil {
		rdbEnabled := *source.RdbEnabled
		persistence.RdbEnabled = &rdbEnabled
	} else {
		persistence.RdbEnabled = nil
	}

	// RdbFrequency
	persistence.RdbFrequency = genruntime.ClonePointerToString(source.RdbFrequency)

	// Update the property bag
	if len(propertyBag) > 0 {
		persistence.PropertyBag = propertyBag
	} else {
		persistence.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToPersistence populates the provided destination Persistence from our Persistence
func (persistence *Persistence) AssignPropertiesToPersistence(destination *v20210301s.Persistence) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(persistence.PropertyBag)

	// AofEnabled
	if persistence.AofEnabled != nil {
		aofEnabled := *persistence.AofEnabled
		destination.AofEnabled = &aofEnabled
	} else {
		destination.AofEnabled = nil
	}

	// AofFrequency
	destination.AofFrequency = genruntime.ClonePointerToString(persistence.AofFrequency)

	// RdbEnabled
	if persistence.RdbEnabled != nil {
		rdbEnabled := *persistence.RdbEnabled
		destination.RdbEnabled = &rdbEnabled
	} else {
		destination.RdbEnabled = nil
	}

	// RdbFrequency
	destination.RdbFrequency = genruntime.ClonePointerToString(persistence.RdbFrequency)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210301.Persistence_STATUS
// Deprecated version of Persistence_STATUS. Use v1beta20210301.Persistence_STATUS instead
type Persistence_STATUS struct {
	AofEnabled   *bool                  `json:"aofEnabled,omitempty"`
	AofFrequency *string                `json:"aofFrequency,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RdbEnabled   *bool                  `json:"rdbEnabled,omitempty"`
	RdbFrequency *string                `json:"rdbFrequency,omitempty"`
}

// AssignPropertiesFromPersistence_STATUS populates our Persistence_STATUS from the provided source Persistence_STATUS
func (persistence *Persistence_STATUS) AssignPropertiesFromPersistence_STATUS(source *v20210301s.Persistence_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AofEnabled
	if source.AofEnabled != nil {
		aofEnabled := *source.AofEnabled
		persistence.AofEnabled = &aofEnabled
	} else {
		persistence.AofEnabled = nil
	}

	// AofFrequency
	persistence.AofFrequency = genruntime.ClonePointerToString(source.AofFrequency)

	// RdbEnabled
	if source.RdbEnabled != nil {
		rdbEnabled := *source.RdbEnabled
		persistence.RdbEnabled = &rdbEnabled
	} else {
		persistence.RdbEnabled = nil
	}

	// RdbFrequency
	persistence.RdbFrequency = genruntime.ClonePointerToString(source.RdbFrequency)

	// Update the property bag
	if len(propertyBag) > 0 {
		persistence.PropertyBag = propertyBag
	} else {
		persistence.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToPersistence_STATUS populates the provided destination Persistence_STATUS from our Persistence_STATUS
func (persistence *Persistence_STATUS) AssignPropertiesToPersistence_STATUS(destination *v20210301s.Persistence_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(persistence.PropertyBag)

	// AofEnabled
	if persistence.AofEnabled != nil {
		aofEnabled := *persistence.AofEnabled
		destination.AofEnabled = &aofEnabled
	} else {
		destination.AofEnabled = nil
	}

	// AofFrequency
	destination.AofFrequency = genruntime.ClonePointerToString(persistence.AofFrequency)

	// RdbEnabled
	if persistence.RdbEnabled != nil {
		rdbEnabled := *persistence.RdbEnabled
		destination.RdbEnabled = &rdbEnabled
	} else {
		destination.RdbEnabled = nil
	}

	// RdbFrequency
	destination.RdbFrequency = genruntime.ClonePointerToString(persistence.RdbFrequency)

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
	SchemeBuilder.Register(&RedisEnterpriseDatabase{}, &RedisEnterpriseDatabaseList{})
}

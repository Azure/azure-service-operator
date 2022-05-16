// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

import (
	"fmt"
	v20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515storage"
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
// Storage version of v1alpha1api20210515.SqlDatabaseContainerTrigger
// Deprecated version of SqlDatabaseContainerTrigger. Use v1beta20210515.SqlDatabaseContainerTrigger instead
type SqlDatabaseContainerTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabasesContainersTriggers_Spec `json:"spec,omitempty"`
	Status            SqlTriggerGetResults_Status                         `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerTrigger{}

// GetConditions returns the conditions of the resource
func (trigger *SqlDatabaseContainerTrigger) GetConditions() conditions.Conditions {
	return trigger.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (trigger *SqlDatabaseContainerTrigger) SetConditions(conditions conditions.Conditions) {
	trigger.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabaseContainerTrigger{}

// ConvertFrom populates our SqlDatabaseContainerTrigger from the provided hub SqlDatabaseContainerTrigger
func (trigger *SqlDatabaseContainerTrigger) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210515s.SqlDatabaseContainerTrigger)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/SqlDatabaseContainerTrigger but received %T instead", hub)
	}

	return trigger.AssignPropertiesFromSqlDatabaseContainerTrigger(source)
}

// ConvertTo populates the provided hub SqlDatabaseContainerTrigger from our SqlDatabaseContainerTrigger
func (trigger *SqlDatabaseContainerTrigger) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210515s.SqlDatabaseContainerTrigger)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/SqlDatabaseContainerTrigger but received %T instead", hub)
	}

	return trigger.AssignPropertiesToSqlDatabaseContainerTrigger(destination)
}

var _ genruntime.KubernetesResource = &SqlDatabaseContainerTrigger{}

// AzureName returns the Azure name of the resource
func (trigger *SqlDatabaseContainerTrigger) AzureName() string {
	return trigger.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (trigger SqlDatabaseContainerTrigger) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (trigger *SqlDatabaseContainerTrigger) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (trigger *SqlDatabaseContainerTrigger) GetSpec() genruntime.ConvertibleSpec {
	return &trigger.Spec
}

// GetStatus returns the status of this resource
func (trigger *SqlDatabaseContainerTrigger) GetStatus() genruntime.ConvertibleStatus {
	return &trigger.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/triggers"
func (trigger *SqlDatabaseContainerTrigger) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/triggers"
}

// NewEmptyStatus returns a new empty (blank) status
func (trigger *SqlDatabaseContainerTrigger) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlTriggerGetResults_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (trigger *SqlDatabaseContainerTrigger) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(trigger.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  trigger.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (trigger *SqlDatabaseContainerTrigger) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlTriggerGetResults_Status); ok {
		trigger.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlTriggerGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	trigger.Status = st
	return nil
}

// AssignPropertiesFromSqlDatabaseContainerTrigger populates our SqlDatabaseContainerTrigger from the provided source SqlDatabaseContainerTrigger
func (trigger *SqlDatabaseContainerTrigger) AssignPropertiesFromSqlDatabaseContainerTrigger(source *v20210515s.SqlDatabaseContainerTrigger) error {

	// ObjectMeta
	trigger.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccountsSqlDatabasesContainersTriggers_Spec
	err := spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersTriggersSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersTriggersSpec() to populate field Spec")
	}
	trigger.Spec = spec

	// Status
	var status SqlTriggerGetResults_Status
	err = status.AssignPropertiesFromSqlTriggerGetResultsStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromSqlTriggerGetResultsStatus() to populate field Status")
	}
	trigger.Status = status

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseContainerTrigger populates the provided destination SqlDatabaseContainerTrigger from our SqlDatabaseContainerTrigger
func (trigger *SqlDatabaseContainerTrigger) AssignPropertiesToSqlDatabaseContainerTrigger(destination *v20210515s.SqlDatabaseContainerTrigger) error {

	// ObjectMeta
	destination.ObjectMeta = *trigger.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210515s.DatabaseAccountsSqlDatabasesContainersTriggers_Spec
	err := trigger.Spec.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersTriggersSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsSqlDatabasesContainersTriggersSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210515s.SqlTriggerGetResults_Status
	err = trigger.Status.AssignPropertiesToSqlTriggerGetResultsStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToSqlTriggerGetResultsStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (trigger *SqlDatabaseContainerTrigger) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: trigger.Spec.OriginalVersion,
		Kind:    "SqlDatabaseContainerTrigger",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20210515.SqlDatabaseContainerTrigger
// Deprecated version of SqlDatabaseContainerTrigger. Use v1beta20210515.SqlDatabaseContainerTrigger instead
type SqlDatabaseContainerTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerTrigger `json:"items"`
}

// Storage version of v1alpha1api20210515.DatabaseAccountsSqlDatabasesContainersTriggers_Spec
type DatabaseAccountsSqlDatabasesContainersTriggers_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string               `json:"azureName,omitempty"`
	Location        *string              `json:"location,omitempty"`
	Options         *CreateUpdateOptions `json:"options,omitempty"`
	OriginalVersion string               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/SqlDatabaseContainer resource
	Owner       *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabaseContainer"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *SqlTriggerResource                `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabasesContainersTriggers_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabasesContainersTriggers_Spec from the provided source
func (triggers *DatabaseAccountsSqlDatabasesContainersTriggers_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210515s.DatabaseAccountsSqlDatabasesContainersTriggers_Spec)
	if ok {
		// Populate our instance from source
		return triggers.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersTriggersSpec(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccountsSqlDatabasesContainersTriggers_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = triggers.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersTriggersSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabasesContainersTriggers_Spec
func (triggers *DatabaseAccountsSqlDatabasesContainersTriggers_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210515s.DatabaseAccountsSqlDatabasesContainersTriggers_Spec)
	if ok {
		// Populate destination from our instance
		return triggers.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersTriggersSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccountsSqlDatabasesContainersTriggers_Spec{}
	err := triggers.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersTriggersSpec(dst)
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

// AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersTriggersSpec populates our DatabaseAccountsSqlDatabasesContainersTriggers_Spec from the provided source DatabaseAccountsSqlDatabasesContainersTriggers_Spec
func (triggers *DatabaseAccountsSqlDatabasesContainersTriggers_Spec) AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersTriggersSpec(source *v20210515s.DatabaseAccountsSqlDatabasesContainersTriggers_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	triggers.AzureName = source.AzureName

	// Location
	triggers.Location = genruntime.ClonePointerToString(source.Location)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions
		err := option.AssignPropertiesFromCreateUpdateOptions(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCreateUpdateOptions() to populate field Options")
		}
		triggers.Options = &option
	} else {
		triggers.Options = nil
	}

	// OriginalVersion
	triggers.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		triggers.Owner = &owner
	} else {
		triggers.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlTriggerResource
		err := resource.AssignPropertiesFromSqlTriggerResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlTriggerResource() to populate field Resource")
		}
		triggers.Resource = &resource
	} else {
		triggers.Resource = nil
	}

	// Tags
	triggers.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		triggers.PropertyBag = propertyBag
	} else {
		triggers.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabasesContainersTriggersSpec populates the provided destination DatabaseAccountsSqlDatabasesContainersTriggers_Spec from our DatabaseAccountsSqlDatabasesContainersTriggers_Spec
func (triggers *DatabaseAccountsSqlDatabasesContainersTriggers_Spec) AssignPropertiesToDatabaseAccountsSqlDatabasesContainersTriggersSpec(destination *v20210515s.DatabaseAccountsSqlDatabasesContainersTriggers_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(triggers.PropertyBag)

	// AzureName
	destination.AzureName = triggers.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(triggers.Location)

	// Options
	if triggers.Options != nil {
		var option v20210515s.CreateUpdateOptions
		err := triggers.Options.AssignPropertiesToCreateUpdateOptions(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCreateUpdateOptions() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// OriginalVersion
	destination.OriginalVersion = triggers.OriginalVersion

	// Owner
	if triggers.Owner != nil {
		owner := triggers.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if triggers.Resource != nil {
		var resource v20210515s.SqlTriggerResource
		err := triggers.Resource.AssignPropertiesToSqlTriggerResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlTriggerResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(triggers.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.SqlTriggerGetResults_Status
// Deprecated version of SqlTriggerGetResults_Status. Use v1beta20210515.SqlTriggerGetResults_Status instead
type SqlTriggerGetResults_Status struct {
	Conditions  []conditions.Condition                   `json:"conditions,omitempty"`
	Id          *string                                  `json:"id,omitempty"`
	Location    *string                                  `json:"location,omitempty"`
	Name        *string                                  `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag                   `json:"$propertyBag,omitempty"`
	Resource    *SqlTriggerGetProperties_Status_Resource `json:"resource,omitempty"`
	Tags        map[string]string                        `json:"tags,omitempty"`
	Type        *string                                  `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlTriggerGetResults_Status{}

// ConvertStatusFrom populates our SqlTriggerGetResults_Status from the provided source
func (results *SqlTriggerGetResults_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210515s.SqlTriggerGetResults_Status)
	if ok {
		// Populate our instance from source
		return results.AssignPropertiesFromSqlTriggerGetResultsStatus(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.SqlTriggerGetResults_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = results.AssignPropertiesFromSqlTriggerGetResultsStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SqlTriggerGetResults_Status
func (results *SqlTriggerGetResults_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210515s.SqlTriggerGetResults_Status)
	if ok {
		// Populate destination from our instance
		return results.AssignPropertiesToSqlTriggerGetResultsStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.SqlTriggerGetResults_Status{}
	err := results.AssignPropertiesToSqlTriggerGetResultsStatus(dst)
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

// AssignPropertiesFromSqlTriggerGetResultsStatus populates our SqlTriggerGetResults_Status from the provided source SqlTriggerGetResults_Status
func (results *SqlTriggerGetResults_Status) AssignPropertiesFromSqlTriggerGetResultsStatus(source *v20210515s.SqlTriggerGetResults_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	results.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	results.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	results.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	results.Name = genruntime.ClonePointerToString(source.Name)

	// Resource
	if source.Resource != nil {
		var resource SqlTriggerGetProperties_Status_Resource
		err := resource.AssignPropertiesFromSqlTriggerGetPropertiesStatusResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlTriggerGetPropertiesStatusResource() to populate field Resource")
		}
		results.Resource = &resource
	} else {
		results.Resource = nil
	}

	// Tags
	results.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	results.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		results.PropertyBag = propertyBag
	} else {
		results.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSqlTriggerGetResultsStatus populates the provided destination SqlTriggerGetResults_Status from our SqlTriggerGetResults_Status
func (results *SqlTriggerGetResults_Status) AssignPropertiesToSqlTriggerGetResultsStatus(destination *v20210515s.SqlTriggerGetResults_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(results.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(results.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(results.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(results.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(results.Name)

	// Resource
	if results.Resource != nil {
		var resource v20210515s.SqlTriggerGetProperties_Status_Resource
		err := results.Resource.AssignPropertiesToSqlTriggerGetPropertiesStatusResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlTriggerGetPropertiesStatusResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(results.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(results.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.SqlTriggerGetProperties_Status_Resource
// Deprecated version of SqlTriggerGetProperties_Status_Resource. Use v1beta20210515.SqlTriggerGetProperties_Status_Resource instead
type SqlTriggerGetProperties_Status_Resource struct {
	Body             *string                `json:"body,omitempty"`
	Etag             *string                `json:"_etag,omitempty"`
	Id               *string                `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rid              *string                `json:"_rid,omitempty"`
	TriggerOperation *string                `json:"triggerOperation,omitempty"`
	TriggerType      *string                `json:"triggerType,omitempty"`
	Ts               *float64               `json:"_ts,omitempty"`
}

// AssignPropertiesFromSqlTriggerGetPropertiesStatusResource populates our SqlTriggerGetProperties_Status_Resource from the provided source SqlTriggerGetProperties_Status_Resource
func (resource *SqlTriggerGetProperties_Status_Resource) AssignPropertiesFromSqlTriggerGetPropertiesStatusResource(source *v20210515s.SqlTriggerGetProperties_Status_Resource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Body
	resource.Body = genruntime.ClonePointerToString(source.Body)

	// Etag
	resource.Etag = genruntime.ClonePointerToString(source.Etag)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// Rid
	resource.Rid = genruntime.ClonePointerToString(source.Rid)

	// TriggerOperation
	resource.TriggerOperation = genruntime.ClonePointerToString(source.TriggerOperation)

	// TriggerType
	resource.TriggerType = genruntime.ClonePointerToString(source.TriggerType)

	// Ts
	if source.Ts != nil {
		t := *source.Ts
		resource.Ts = &t
	} else {
		resource.Ts = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSqlTriggerGetPropertiesStatusResource populates the provided destination SqlTriggerGetProperties_Status_Resource from our SqlTriggerGetProperties_Status_Resource
func (resource *SqlTriggerGetProperties_Status_Resource) AssignPropertiesToSqlTriggerGetPropertiesStatusResource(destination *v20210515s.SqlTriggerGetProperties_Status_Resource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Body
	destination.Body = genruntime.ClonePointerToString(resource.Body)

	// Etag
	destination.Etag = genruntime.ClonePointerToString(resource.Etag)

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

	// Rid
	destination.Rid = genruntime.ClonePointerToString(resource.Rid)

	// TriggerOperation
	destination.TriggerOperation = genruntime.ClonePointerToString(resource.TriggerOperation)

	// TriggerType
	destination.TriggerType = genruntime.ClonePointerToString(resource.TriggerType)

	// Ts
	if resource.Ts != nil {
		t := *resource.Ts
		destination.Ts = &t
	} else {
		destination.Ts = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.SqlTriggerResource
// Deprecated version of SqlTriggerResource. Use v1beta20210515.SqlTriggerResource instead
type SqlTriggerResource struct {
	Body             *string                `json:"body,omitempty"`
	Id               *string                `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TriggerOperation *string                `json:"triggerOperation,omitempty"`
	TriggerType      *string                `json:"triggerType,omitempty"`
}

// AssignPropertiesFromSqlTriggerResource populates our SqlTriggerResource from the provided source SqlTriggerResource
func (resource *SqlTriggerResource) AssignPropertiesFromSqlTriggerResource(source *v20210515s.SqlTriggerResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Body
	resource.Body = genruntime.ClonePointerToString(source.Body)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// TriggerOperation
	resource.TriggerOperation = genruntime.ClonePointerToString(source.TriggerOperation)

	// TriggerType
	resource.TriggerType = genruntime.ClonePointerToString(source.TriggerType)

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSqlTriggerResource populates the provided destination SqlTriggerResource from our SqlTriggerResource
func (resource *SqlTriggerResource) AssignPropertiesToSqlTriggerResource(destination *v20210515s.SqlTriggerResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Body
	destination.Body = genruntime.ClonePointerToString(resource.Body)

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

	// TriggerOperation
	destination.TriggerOperation = genruntime.ClonePointerToString(resource.TriggerOperation)

	// TriggerType
	destination.TriggerType = genruntime.ClonePointerToString(resource.TriggerType)

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
	SchemeBuilder.Register(&SqlDatabaseContainerTrigger{}, &SqlDatabaseContainerTriggerList{})
}

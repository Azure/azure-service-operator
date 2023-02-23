// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import (
	"fmt"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/routeTable.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}
type RouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTable_Spec   `json:"spec,omitempty"`
	Status            RouteTable_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RouteTable{}

// GetConditions returns the conditions of the resource
func (table *RouteTable) GetConditions() conditions.Conditions {
	return table.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (table *RouteTable) SetConditions(conditions conditions.Conditions) {
	table.Status.Conditions = conditions
}

var _ conversion.Convertible = &RouteTable{}

// ConvertFrom populates our RouteTable from the provided hub RouteTable
func (table *RouteTable) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20201101s.RouteTable)
	if !ok {
		return fmt.Errorf("expected network/v1beta20201101storage/RouteTable but received %T instead", hub)
	}

	return table.AssignProperties_From_RouteTable(source)
}

// ConvertTo populates the provided hub RouteTable from our RouteTable
func (table *RouteTable) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20201101s.RouteTable)
	if !ok {
		return fmt.Errorf("expected network/v1beta20201101storage/RouteTable but received %T instead", hub)
	}

	return table.AssignProperties_To_RouteTable(destination)
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1beta20201101-routetable,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=routetables,verbs=create;update,versions=v1beta20201101,name=default.v1beta20201101.routetables.network.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &RouteTable{}

// Default applies defaults to the RouteTable resource
func (table *RouteTable) Default() {
	table.defaultImpl()
	var temp any = table
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (table *RouteTable) defaultAzureName() {
	if table.Spec.AzureName == "" {
		table.Spec.AzureName = table.Name
	}
}

// defaultImpl applies the code generated defaults to the RouteTable resource
func (table *RouteTable) defaultImpl() { table.defaultAzureName() }

var _ genruntime.KubernetesResource = &RouteTable{}

// AzureName returns the Azure name of the resource
func (table *RouteTable) AzureName() string {
	return table.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (table RouteTable) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (table *RouteTable) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (table *RouteTable) GetSpec() genruntime.ConvertibleSpec {
	return &table.Spec
}

// GetStatus returns the status of this resource
func (table *RouteTable) GetStatus() genruntime.ConvertibleStatus {
	return &table.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/routeTables"
func (table *RouteTable) GetType() string {
	return "Microsoft.Network/routeTables"
}

// NewEmptyStatus returns a new empty (blank) status
func (table *RouteTable) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RouteTable_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (table *RouteTable) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(table.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  table.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (table *RouteTable) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RouteTable_STATUS); ok {
		table.Status = *st
		return nil
	}

	// Convert status to required version
	var st RouteTable_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	table.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1beta20201101-routetable,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=routetables,verbs=create;update,versions=v1beta20201101,name=validate.v1beta20201101.routetables.network.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &RouteTable{}

// ValidateCreate validates the creation of the resource
func (table *RouteTable) ValidateCreate() error {
	validations := table.createValidations()
	var temp any = table
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (table *RouteTable) ValidateDelete() error {
	validations := table.deleteValidations()
	var temp any = table
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (table *RouteTable) ValidateUpdate(old runtime.Object) error {
	validations := table.updateValidations()
	var temp any = table
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (table *RouteTable) createValidations() []func() error {
	return []func() error{table.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (table *RouteTable) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (table *RouteTable) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return table.validateResourceReferences()
		},
		table.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (table *RouteTable) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&table.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (table *RouteTable) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*RouteTable)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, table)
}

// AssignProperties_From_RouteTable populates our RouteTable from the provided source RouteTable
func (table *RouteTable) AssignProperties_From_RouteTable(source *v20201101s.RouteTable) error {

	// ObjectMeta
	table.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RouteTable_Spec
	err := spec.AssignProperties_From_RouteTable_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_RouteTable_Spec() to populate field Spec")
	}
	table.Spec = spec

	// Status
	var status RouteTable_STATUS
	err = status.AssignProperties_From_RouteTable_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_RouteTable_STATUS() to populate field Status")
	}
	table.Status = status

	// No error
	return nil
}

// AssignProperties_To_RouteTable populates the provided destination RouteTable from our RouteTable
func (table *RouteTable) AssignProperties_To_RouteTable(destination *v20201101s.RouteTable) error {

	// ObjectMeta
	destination.ObjectMeta = *table.ObjectMeta.DeepCopy()

	// Spec
	var spec v20201101s.RouteTable_Spec
	err := table.Spec.AssignProperties_To_RouteTable_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_RouteTable_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20201101s.RouteTable_STATUS
	err = table.Status.AssignProperties_To_RouteTable_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_RouteTable_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (table *RouteTable) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: table.Spec.OriginalVersion(),
		Kind:    "RouteTable",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/routeTable.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTable `json:"items"`
}

type RouteTable_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// DisableBgpRoutePropagation: Whether to disable the routes learned by BGP on that route table. True means disable.
	DisableBgpRoutePropagation *bool `json:"disableBgpRoutePropagation,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &RouteTable_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (table *RouteTable_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if table == nil {
		return nil, nil
	}
	result := &RouteTable_Spec_ARM{}

	// Set property ‘Location’:
	if table.Location != nil {
		location := *table.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if table.DisableBgpRoutePropagation != nil {
		result.Properties = &RouteTablePropertiesFormat_ARM{}
	}
	if table.DisableBgpRoutePropagation != nil {
		disableBgpRoutePropagation := *table.DisableBgpRoutePropagation
		result.Properties.DisableBgpRoutePropagation = &disableBgpRoutePropagation
	}

	// Set property ‘Tags’:
	if table.Tags != nil {
		result.Tags = make(map[string]string, len(table.Tags))
		for key, value := range table.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (table *RouteTable_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RouteTable_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (table *RouteTable_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RouteTable_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RouteTable_Spec_ARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	table.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘DisableBgpRoutePropagation’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DisableBgpRoutePropagation != nil {
			disableBgpRoutePropagation := *typedInput.Properties.DisableBgpRoutePropagation
			table.DisableBgpRoutePropagation = &disableBgpRoutePropagation
		}
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		table.Location = &location
	}

	// Set property ‘Owner’:
	table.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		table.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			table.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &RouteTable_Spec{}

// ConvertSpecFrom populates our RouteTable_Spec from the provided source
func (table *RouteTable_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20201101s.RouteTable_Spec)
	if ok {
		// Populate our instance from source
		return table.AssignProperties_From_RouteTable_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20201101s.RouteTable_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = table.AssignProperties_From_RouteTable_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RouteTable_Spec
func (table *RouteTable_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20201101s.RouteTable_Spec)
	if ok {
		// Populate destination from our instance
		return table.AssignProperties_To_RouteTable_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20201101s.RouteTable_Spec{}
	err := table.AssignProperties_To_RouteTable_Spec(dst)
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

// AssignProperties_From_RouteTable_Spec populates our RouteTable_Spec from the provided source RouteTable_Spec
func (table *RouteTable_Spec) AssignProperties_From_RouteTable_Spec(source *v20201101s.RouteTable_Spec) error {

	// AzureName
	table.AzureName = source.AzureName

	// DisableBgpRoutePropagation
	if source.DisableBgpRoutePropagation != nil {
		disableBgpRoutePropagation := *source.DisableBgpRoutePropagation
		table.DisableBgpRoutePropagation = &disableBgpRoutePropagation
	} else {
		table.DisableBgpRoutePropagation = nil
	}

	// Location
	table.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		table.Owner = &owner
	} else {
		table.Owner = nil
	}

	// Tags
	table.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignProperties_To_RouteTable_Spec populates the provided destination RouteTable_Spec from our RouteTable_Spec
func (table *RouteTable_Spec) AssignProperties_To_RouteTable_Spec(destination *v20201101s.RouteTable_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = table.AzureName

	// DisableBgpRoutePropagation
	if table.DisableBgpRoutePropagation != nil {
		disableBgpRoutePropagation := *table.DisableBgpRoutePropagation
		destination.DisableBgpRoutePropagation = &disableBgpRoutePropagation
	} else {
		destination.DisableBgpRoutePropagation = nil
	}

	// Location
	destination.Location = genruntime.ClonePointerToString(table.Location)

	// OriginalVersion
	destination.OriginalVersion = table.OriginalVersion()

	// Owner
	if table.Owner != nil {
		owner := table.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(table.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (table *RouteTable_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (table *RouteTable_Spec) SetAzureName(azureName string) { table.AzureName = azureName }

// Route table resource.
type RouteTable_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// DisableBgpRoutePropagation: Whether to disable the routes learned by BGP on that route table. True means disable.
	DisableBgpRoutePropagation *bool `json:"disableBgpRoutePropagation,omitempty"`

	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// ProvisioningState: The provisioning state of the route table resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ResourceGuid: The resource GUID property of the route table.
	ResourceGuid *string `json:"resourceGuid,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RouteTable_STATUS{}

// ConvertStatusFrom populates our RouteTable_STATUS from the provided source
func (table *RouteTable_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20201101s.RouteTable_STATUS)
	if ok {
		// Populate our instance from source
		return table.AssignProperties_From_RouteTable_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20201101s.RouteTable_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = table.AssignProperties_From_RouteTable_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RouteTable_STATUS
func (table *RouteTable_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20201101s.RouteTable_STATUS)
	if ok {
		// Populate destination from our instance
		return table.AssignProperties_To_RouteTable_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20201101s.RouteTable_STATUS{}
	err := table.AssignProperties_To_RouteTable_STATUS(dst)
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

var _ genruntime.FromARMConverter = &RouteTable_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (table *RouteTable_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RouteTable_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (table *RouteTable_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RouteTable_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RouteTable_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘DisableBgpRoutePropagation’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DisableBgpRoutePropagation != nil {
			disableBgpRoutePropagation := *typedInput.Properties.DisableBgpRoutePropagation
			table.DisableBgpRoutePropagation = &disableBgpRoutePropagation
		}
	}

	// Set property ‘Etag’:
	if typedInput.Etag != nil {
		etag := *typedInput.Etag
		table.Etag = &etag
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		table.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		table.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		table.Name = &name
	}

	// Set property ‘ProvisioningState’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ProvisioningState != nil {
			provisioningState := *typedInput.Properties.ProvisioningState
			table.ProvisioningState = &provisioningState
		}
	}

	// Set property ‘ResourceGuid’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ResourceGuid != nil {
			resourceGuid := *typedInput.Properties.ResourceGuid
			table.ResourceGuid = &resourceGuid
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		table.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			table.Tags[key] = value
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		table.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_RouteTable_STATUS populates our RouteTable_STATUS from the provided source RouteTable_STATUS
func (table *RouteTable_STATUS) AssignProperties_From_RouteTable_STATUS(source *v20201101s.RouteTable_STATUS) error {

	// Conditions
	table.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// DisableBgpRoutePropagation
	if source.DisableBgpRoutePropagation != nil {
		disableBgpRoutePropagation := *source.DisableBgpRoutePropagation
		table.DisableBgpRoutePropagation = &disableBgpRoutePropagation
	} else {
		table.DisableBgpRoutePropagation = nil
	}

	// Etag
	table.Etag = genruntime.ClonePointerToString(source.Etag)

	// Id
	table.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	table.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	table.Name = genruntime.ClonePointerToString(source.Name)

	// ProvisioningState
	if source.ProvisioningState != nil {
		provisioningState := ProvisioningState_STATUS(*source.ProvisioningState)
		table.ProvisioningState = &provisioningState
	} else {
		table.ProvisioningState = nil
	}

	// ResourceGuid
	table.ResourceGuid = genruntime.ClonePointerToString(source.ResourceGuid)

	// Tags
	table.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	table.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_RouteTable_STATUS populates the provided destination RouteTable_STATUS from our RouteTable_STATUS
func (table *RouteTable_STATUS) AssignProperties_To_RouteTable_STATUS(destination *v20201101s.RouteTable_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(table.Conditions)

	// DisableBgpRoutePropagation
	if table.DisableBgpRoutePropagation != nil {
		disableBgpRoutePropagation := *table.DisableBgpRoutePropagation
		destination.DisableBgpRoutePropagation = &disableBgpRoutePropagation
	} else {
		destination.DisableBgpRoutePropagation = nil
	}

	// Etag
	destination.Etag = genruntime.ClonePointerToString(table.Etag)

	// Id
	destination.Id = genruntime.ClonePointerToString(table.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(table.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(table.Name)

	// ProvisioningState
	if table.ProvisioningState != nil {
		provisioningState := string(*table.ProvisioningState)
		destination.ProvisioningState = &provisioningState
	} else {
		destination.ProvisioningState = nil
	}

	// ResourceGuid
	destination.ResourceGuid = genruntime.ClonePointerToString(table.ResourceGuid)

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(table.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(table.Type)

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
	SchemeBuilder.Register(&RouteTable{}, &RouteTableList{})
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210701

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1api20210701/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
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
// - Generated from: /machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/machineLearningServices.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/connections/{connectionName}
type WorkspacesConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Workspaces_Connection_Spec   `json:"spec,omitempty"`
	Status            Workspaces_Connection_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &WorkspacesConnection{}

// GetConditions returns the conditions of the resource
func (connection *WorkspacesConnection) GetConditions() conditions.Conditions {
	return connection.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (connection *WorkspacesConnection) SetConditions(conditions conditions.Conditions) {
	connection.Status.Conditions = conditions
}

var _ conversion.Convertible = &WorkspacesConnection{}

// ConvertFrom populates our WorkspacesConnection from the provided hub WorkspacesConnection
func (connection *WorkspacesConnection) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.WorkspacesConnection)
	if !ok {
		return fmt.Errorf("expected machinelearningservices/v1api20210701/storage/WorkspacesConnection but received %T instead", hub)
	}

	return connection.AssignProperties_From_WorkspacesConnection(source)
}

// ConvertTo populates the provided hub WorkspacesConnection from our WorkspacesConnection
func (connection *WorkspacesConnection) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.WorkspacesConnection)
	if !ok {
		return fmt.Errorf("expected machinelearningservices/v1api20210701/storage/WorkspacesConnection but received %T instead", hub)
	}

	return connection.AssignProperties_To_WorkspacesConnection(destination)
}

// +kubebuilder:webhook:path=/mutate-machinelearningservices-azure-com-v1api20210701-workspacesconnection,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=machinelearningservices.azure.com,resources=workspacesconnections,verbs=create;update,versions=v1api20210701,name=default.v1api20210701.workspacesconnections.machinelearningservices.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &WorkspacesConnection{}

// Default applies defaults to the WorkspacesConnection resource
func (connection *WorkspacesConnection) Default() {
	connection.defaultImpl()
	var temp any = connection
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (connection *WorkspacesConnection) defaultAzureName() {
	if connection.Spec.AzureName == "" {
		connection.Spec.AzureName = connection.Name
	}
}

// defaultImpl applies the code generated defaults to the WorkspacesConnection resource
func (connection *WorkspacesConnection) defaultImpl() { connection.defaultAzureName() }

var _ genruntime.ImportableResource = &WorkspacesConnection{}

// InitializeSpec initializes the spec for this resource from the given status
func (connection *WorkspacesConnection) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Workspaces_Connection_STATUS); ok {
		return connection.Spec.Initialize_From_Workspaces_Connection_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Workspaces_Connection_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &WorkspacesConnection{}

// AzureName returns the Azure name of the resource
func (connection *WorkspacesConnection) AzureName() string {
	return connection.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-07-01"
func (connection WorkspacesConnection) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (connection *WorkspacesConnection) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (connection *WorkspacesConnection) GetSpec() genruntime.ConvertibleSpec {
	return &connection.Spec
}

// GetStatus returns the status of this resource
func (connection *WorkspacesConnection) GetStatus() genruntime.ConvertibleStatus {
	return &connection.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (connection *WorkspacesConnection) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.MachineLearningServices/workspaces/connections"
func (connection *WorkspacesConnection) GetType() string {
	return "Microsoft.MachineLearningServices/workspaces/connections"
}

// NewEmptyStatus returns a new empty (blank) status
func (connection *WorkspacesConnection) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Workspaces_Connection_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (connection *WorkspacesConnection) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(connection.Spec)
	return connection.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (connection *WorkspacesConnection) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Workspaces_Connection_STATUS); ok {
		connection.Status = *st
		return nil
	}

	// Convert status to required version
	var st Workspaces_Connection_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	connection.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-machinelearningservices-azure-com-v1api20210701-workspacesconnection,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=machinelearningservices.azure.com,resources=workspacesconnections,verbs=create;update,versions=v1api20210701,name=validate.v1api20210701.workspacesconnections.machinelearningservices.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &WorkspacesConnection{}

// ValidateCreate validates the creation of the resource
func (connection *WorkspacesConnection) ValidateCreate() (admission.Warnings, error) {
	validations := connection.createValidations()
	var temp any = connection
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (connection *WorkspacesConnection) ValidateDelete() (admission.Warnings, error) {
	validations := connection.deleteValidations()
	var temp any = connection
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (connection *WorkspacesConnection) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := connection.updateValidations()
	var temp any = connection
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (connection *WorkspacesConnection) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){connection.validateResourceReferences, connection.validateOwnerReference}
}

// deleteValidations validates the deletion of the resource
func (connection *WorkspacesConnection) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (connection *WorkspacesConnection) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return connection.validateResourceReferences()
		},
		connection.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return connection.validateOwnerReference()
		},
	}
}

// validateOwnerReference validates the owner field
func (connection *WorkspacesConnection) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(connection)
}

// validateResourceReferences validates all resource references
func (connection *WorkspacesConnection) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&connection.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (connection *WorkspacesConnection) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*WorkspacesConnection)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, connection)
}

// AssignProperties_From_WorkspacesConnection populates our WorkspacesConnection from the provided source WorkspacesConnection
func (connection *WorkspacesConnection) AssignProperties_From_WorkspacesConnection(source *storage.WorkspacesConnection) error {

	// ObjectMeta
	connection.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Workspaces_Connection_Spec
	err := spec.AssignProperties_From_Workspaces_Connection_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Workspaces_Connection_Spec() to populate field Spec")
	}
	connection.Spec = spec

	// Status
	var status Workspaces_Connection_STATUS
	err = status.AssignProperties_From_Workspaces_Connection_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Workspaces_Connection_STATUS() to populate field Status")
	}
	connection.Status = status

	// No error
	return nil
}

// AssignProperties_To_WorkspacesConnection populates the provided destination WorkspacesConnection from our WorkspacesConnection
func (connection *WorkspacesConnection) AssignProperties_To_WorkspacesConnection(destination *storage.WorkspacesConnection) error {

	// ObjectMeta
	destination.ObjectMeta = *connection.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.Workspaces_Connection_Spec
	err := connection.Spec.AssignProperties_To_Workspaces_Connection_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Workspaces_Connection_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.Workspaces_Connection_STATUS
	err = connection.Status.AssignProperties_To_Workspaces_Connection_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Workspaces_Connection_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (connection *WorkspacesConnection) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: connection.Spec.OriginalVersion(),
		Kind:    "WorkspacesConnection",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/machineLearningServices.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/connections/{connectionName}
type WorkspacesConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspacesConnection `json:"items"`
}

type Workspaces_Connection_Spec struct {
	// AuthType: Authorization type of the workspace connection.
	AuthType *string `json:"authType,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// Category: Category of the workspace connection.
	Category *string `json:"category,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a machinelearningservices.azure.com/Workspace resource
	Owner *genruntime.KnownResourceReference `group:"machinelearningservices.azure.com" json:"owner,omitempty" kind:"Workspace"`

	// Target: Target of the workspace connection.
	Target *string `json:"target,omitempty"`

	// Value: Value details of the workspace connection.
	Value *string `json:"value,omitempty"`

	// ValueFormat: format for the workspace connection value
	ValueFormat *WorkspaceConnectionProps_ValueFormat `json:"valueFormat,omitempty"`
}

var _ genruntime.ARMTransformer = &Workspaces_Connection_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (connection *Workspaces_Connection_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if connection == nil {
		return nil, nil
	}
	result := &Workspaces_Connection_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if connection.AuthType != nil ||
		connection.Category != nil ||
		connection.Target != nil ||
		connection.Value != nil ||
		connection.ValueFormat != nil {
		result.Properties = &WorkspaceConnectionProps_ARM{}
	}
	if connection.AuthType != nil {
		authType := *connection.AuthType
		result.Properties.AuthType = &authType
	}
	if connection.Category != nil {
		category := *connection.Category
		result.Properties.Category = &category
	}
	if connection.Target != nil {
		target := *connection.Target
		result.Properties.Target = &target
	}
	if connection.Value != nil {
		value := *connection.Value
		result.Properties.Value = &value
	}
	if connection.ValueFormat != nil {
		var temp string
		temp = string(*connection.ValueFormat)
		valueFormat := WorkspaceConnectionProps_ValueFormat_ARM(temp)
		result.Properties.ValueFormat = &valueFormat
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (connection *Workspaces_Connection_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Workspaces_Connection_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (connection *Workspaces_Connection_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Workspaces_Connection_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Workspaces_Connection_Spec_ARM, got %T", armInput)
	}

	// Set property "AuthType":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AuthType != nil {
			authType := *typedInput.Properties.AuthType
			connection.AuthType = &authType
		}
	}

	// Set property "AzureName":
	connection.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "Category":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Category != nil {
			category := *typedInput.Properties.Category
			connection.Category = &category
		}
	}

	// Set property "Owner":
	connection.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "Target":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Target != nil {
			target := *typedInput.Properties.Target
			connection.Target = &target
		}
	}

	// Set property "Value":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Value != nil {
			value := *typedInput.Properties.Value
			connection.Value = &value
		}
	}

	// Set property "ValueFormat":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ValueFormat != nil {
			var temp string
			temp = string(*typedInput.Properties.ValueFormat)
			valueFormat := WorkspaceConnectionProps_ValueFormat(temp)
			connection.ValueFormat = &valueFormat
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Workspaces_Connection_Spec{}

// ConvertSpecFrom populates our Workspaces_Connection_Spec from the provided source
func (connection *Workspaces_Connection_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.Workspaces_Connection_Spec)
	if ok {
		// Populate our instance from source
		return connection.AssignProperties_From_Workspaces_Connection_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.Workspaces_Connection_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = connection.AssignProperties_From_Workspaces_Connection_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Workspaces_Connection_Spec
func (connection *Workspaces_Connection_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.Workspaces_Connection_Spec)
	if ok {
		// Populate destination from our instance
		return connection.AssignProperties_To_Workspaces_Connection_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Workspaces_Connection_Spec{}
	err := connection.AssignProperties_To_Workspaces_Connection_Spec(dst)
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

// AssignProperties_From_Workspaces_Connection_Spec populates our Workspaces_Connection_Spec from the provided source Workspaces_Connection_Spec
func (connection *Workspaces_Connection_Spec) AssignProperties_From_Workspaces_Connection_Spec(source *storage.Workspaces_Connection_Spec) error {

	// AuthType
	connection.AuthType = genruntime.ClonePointerToString(source.AuthType)

	// AzureName
	connection.AzureName = source.AzureName

	// Category
	connection.Category = genruntime.ClonePointerToString(source.Category)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		connection.Owner = &owner
	} else {
		connection.Owner = nil
	}

	// Target
	connection.Target = genruntime.ClonePointerToString(source.Target)

	// Value
	connection.Value = genruntime.ClonePointerToString(source.Value)

	// ValueFormat
	if source.ValueFormat != nil {
		valueFormat := *source.ValueFormat
		valueFormatTemp := genruntime.ToEnum(valueFormat, workspaceConnectionProps_ValueFormat_Values)
		connection.ValueFormat = &valueFormatTemp
	} else {
		connection.ValueFormat = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Workspaces_Connection_Spec populates the provided destination Workspaces_Connection_Spec from our Workspaces_Connection_Spec
func (connection *Workspaces_Connection_Spec) AssignProperties_To_Workspaces_Connection_Spec(destination *storage.Workspaces_Connection_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AuthType
	destination.AuthType = genruntime.ClonePointerToString(connection.AuthType)

	// AzureName
	destination.AzureName = connection.AzureName

	// Category
	destination.Category = genruntime.ClonePointerToString(connection.Category)

	// OriginalVersion
	destination.OriginalVersion = connection.OriginalVersion()

	// Owner
	if connection.Owner != nil {
		owner := connection.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Target
	destination.Target = genruntime.ClonePointerToString(connection.Target)

	// Value
	destination.Value = genruntime.ClonePointerToString(connection.Value)

	// ValueFormat
	if connection.ValueFormat != nil {
		valueFormat := string(*connection.ValueFormat)
		destination.ValueFormat = &valueFormat
	} else {
		destination.ValueFormat = nil
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

// Initialize_From_Workspaces_Connection_STATUS populates our Workspaces_Connection_Spec from the provided source Workspaces_Connection_STATUS
func (connection *Workspaces_Connection_Spec) Initialize_From_Workspaces_Connection_STATUS(source *Workspaces_Connection_STATUS) error {

	// AuthType
	connection.AuthType = genruntime.ClonePointerToString(source.AuthType)

	// Category
	connection.Category = genruntime.ClonePointerToString(source.Category)

	// Target
	connection.Target = genruntime.ClonePointerToString(source.Target)

	// Value
	connection.Value = genruntime.ClonePointerToString(source.Value)

	// ValueFormat
	if source.ValueFormat != nil {
		valueFormat := genruntime.ToEnum(string(*source.ValueFormat), workspaceConnectionProps_ValueFormat_Values)
		connection.ValueFormat = &valueFormat
	} else {
		connection.ValueFormat = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (connection *Workspaces_Connection_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (connection *Workspaces_Connection_Spec) SetAzureName(azureName string) {
	connection.AzureName = azureName
}

type Workspaces_Connection_STATUS struct {
	// AuthType: Authorization type of the workspace connection.
	AuthType *string `json:"authType,omitempty"`

	// Category: Category of the workspace connection.
	Category *string `json:"category,omitempty"`

	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: ResourceId of the workspace connection.
	Id *string `json:"id,omitempty"`

	// Name: Friendly name of the workspace connection.
	Name *string `json:"name,omitempty"`

	// Target: Target of the workspace connection.
	Target *string `json:"target,omitempty"`

	// Type: Resource type of workspace connection.
	Type *string `json:"type,omitempty"`

	// Value: Value details of the workspace connection.
	Value *string `json:"value,omitempty"`

	// ValueFormat: format for the workspace connection value
	ValueFormat *WorkspaceConnectionProps_ValueFormat_STATUS `json:"valueFormat,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Workspaces_Connection_STATUS{}

// ConvertStatusFrom populates our Workspaces_Connection_STATUS from the provided source
func (connection *Workspaces_Connection_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.Workspaces_Connection_STATUS)
	if ok {
		// Populate our instance from source
		return connection.AssignProperties_From_Workspaces_Connection_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.Workspaces_Connection_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = connection.AssignProperties_From_Workspaces_Connection_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Workspaces_Connection_STATUS
func (connection *Workspaces_Connection_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.Workspaces_Connection_STATUS)
	if ok {
		// Populate destination from our instance
		return connection.AssignProperties_To_Workspaces_Connection_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Workspaces_Connection_STATUS{}
	err := connection.AssignProperties_To_Workspaces_Connection_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Workspaces_Connection_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (connection *Workspaces_Connection_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Workspaces_Connection_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (connection *Workspaces_Connection_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Workspaces_Connection_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Workspaces_Connection_STATUS_ARM, got %T", armInput)
	}

	// Set property "AuthType":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AuthType != nil {
			authType := *typedInput.Properties.AuthType
			connection.AuthType = &authType
		}
	}

	// Set property "Category":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Category != nil {
			category := *typedInput.Properties.Category
			connection.Category = &category
		}
	}

	// no assignment for property "Conditions"

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		connection.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		connection.Name = &name
	}

	// Set property "Target":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Target != nil {
			target := *typedInput.Properties.Target
			connection.Target = &target
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		connection.Type = &typeVar
	}

	// Set property "Value":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Value != nil {
			value := *typedInput.Properties.Value
			connection.Value = &value
		}
	}

	// Set property "ValueFormat":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ValueFormat != nil {
			var temp string
			temp = string(*typedInput.Properties.ValueFormat)
			valueFormat := WorkspaceConnectionProps_ValueFormat_STATUS(temp)
			connection.ValueFormat = &valueFormat
		}
	}

	// No error
	return nil
}

// AssignProperties_From_Workspaces_Connection_STATUS populates our Workspaces_Connection_STATUS from the provided source Workspaces_Connection_STATUS
func (connection *Workspaces_Connection_STATUS) AssignProperties_From_Workspaces_Connection_STATUS(source *storage.Workspaces_Connection_STATUS) error {

	// AuthType
	connection.AuthType = genruntime.ClonePointerToString(source.AuthType)

	// Category
	connection.Category = genruntime.ClonePointerToString(source.Category)

	// Conditions
	connection.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	connection.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	connection.Name = genruntime.ClonePointerToString(source.Name)

	// Target
	connection.Target = genruntime.ClonePointerToString(source.Target)

	// Type
	connection.Type = genruntime.ClonePointerToString(source.Type)

	// Value
	connection.Value = genruntime.ClonePointerToString(source.Value)

	// ValueFormat
	if source.ValueFormat != nil {
		valueFormat := *source.ValueFormat
		valueFormatTemp := genruntime.ToEnum(valueFormat, workspaceConnectionProps_ValueFormat_STATUS_Values)
		connection.ValueFormat = &valueFormatTemp
	} else {
		connection.ValueFormat = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Workspaces_Connection_STATUS populates the provided destination Workspaces_Connection_STATUS from our Workspaces_Connection_STATUS
func (connection *Workspaces_Connection_STATUS) AssignProperties_To_Workspaces_Connection_STATUS(destination *storage.Workspaces_Connection_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AuthType
	destination.AuthType = genruntime.ClonePointerToString(connection.AuthType)

	// Category
	destination.Category = genruntime.ClonePointerToString(connection.Category)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(connection.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(connection.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(connection.Name)

	// Target
	destination.Target = genruntime.ClonePointerToString(connection.Target)

	// Type
	destination.Type = genruntime.ClonePointerToString(connection.Type)

	// Value
	destination.Value = genruntime.ClonePointerToString(connection.Value)

	// ValueFormat
	if connection.ValueFormat != nil {
		valueFormat := string(*connection.ValueFormat)
		destination.ValueFormat = &valueFormat
	} else {
		destination.ValueFormat = nil
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

// +kubebuilder:validation:Enum={"JSON"}
type WorkspaceConnectionProps_ValueFormat string

const WorkspaceConnectionProps_ValueFormat_JSON = WorkspaceConnectionProps_ValueFormat("JSON")

// Mapping from string to WorkspaceConnectionProps_ValueFormat
var workspaceConnectionProps_ValueFormat_Values = map[string]WorkspaceConnectionProps_ValueFormat{
	"json": WorkspaceConnectionProps_ValueFormat_JSON,
}

type WorkspaceConnectionProps_ValueFormat_STATUS string

const WorkspaceConnectionProps_ValueFormat_STATUS_JSON = WorkspaceConnectionProps_ValueFormat_STATUS("JSON")

// Mapping from string to WorkspaceConnectionProps_ValueFormat_STATUS
var workspaceConnectionProps_ValueFormat_STATUS_Values = map[string]WorkspaceConnectionProps_ValueFormat_STATUS{
	"json": WorkspaceConnectionProps_ValueFormat_STATUS_JSON,
}

func init() {
	SchemeBuilder.Register(&WorkspacesConnection{}, &WorkspacesConnectionList{})
}

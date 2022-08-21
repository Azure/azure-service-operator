// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210701

import (
	"fmt"
	v20210701s "github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1beta20210701storage"
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
// - Generated from: /machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/machineLearningServices.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/connections/{connectionName}
type WorkspacesConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspacesConnection_Spec   `json:"spec,omitempty"`
	Status            WorkspacesConnection_STATUS `json:"status,omitempty"`
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
	source, ok := hub.(*v20210701s.WorkspacesConnection)
	if !ok {
		return fmt.Errorf("expected machinelearningservices/v1beta20210701storage/WorkspacesConnection but received %T instead", hub)
	}

	return connection.AssignPropertiesFromWorkspacesConnection(source)
}

// ConvertTo populates the provided hub WorkspacesConnection from our WorkspacesConnection
func (connection *WorkspacesConnection) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210701s.WorkspacesConnection)
	if !ok {
		return fmt.Errorf("expected machinelearningservices/v1beta20210701storage/WorkspacesConnection but received %T instead", hub)
	}

	return connection.AssignPropertiesToWorkspacesConnection(destination)
}

// +kubebuilder:webhook:path=/mutate-machinelearningservices-azure-com-v1beta20210701-workspacesconnection,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=machinelearningservices.azure.com,resources=workspacesconnections,verbs=create;update,versions=v1beta20210701,name=default.v1beta20210701.workspacesconnections.machinelearningservices.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &WorkspacesConnection{}

// Default applies defaults to the WorkspacesConnection resource
func (connection *WorkspacesConnection) Default() {
	connection.defaultImpl()
	var temp interface{} = connection
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

// GetType returns the ARM Type of the resource. This is always "Microsoft.MachineLearningServices/workspaces/connections"
func (connection *WorkspacesConnection) GetType() string {
	return "Microsoft.MachineLearningServices/workspaces/connections"
}

// NewEmptyStatus returns a new empty (blank) status
func (connection *WorkspacesConnection) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &WorkspacesConnection_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (connection *WorkspacesConnection) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(connection.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  connection.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (connection *WorkspacesConnection) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*WorkspacesConnection_STATUS); ok {
		connection.Status = *st
		return nil
	}

	// Convert status to required version
	var st WorkspacesConnection_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	connection.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-machinelearningservices-azure-com-v1beta20210701-workspacesconnection,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=machinelearningservices.azure.com,resources=workspacesconnections,verbs=create;update,versions=v1beta20210701,name=validate.v1beta20210701.workspacesconnections.machinelearningservices.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &WorkspacesConnection{}

// ValidateCreate validates the creation of the resource
func (connection *WorkspacesConnection) ValidateCreate() error {
	validations := connection.createValidations()
	var temp interface{} = connection
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
func (connection *WorkspacesConnection) ValidateDelete() error {
	validations := connection.deleteValidations()
	var temp interface{} = connection
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
func (connection *WorkspacesConnection) ValidateUpdate(old runtime.Object) error {
	validations := connection.updateValidations()
	var temp interface{} = connection
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
func (connection *WorkspacesConnection) createValidations() []func() error {
	return []func() error{connection.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (connection *WorkspacesConnection) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (connection *WorkspacesConnection) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return connection.validateResourceReferences()
		},
		connection.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (connection *WorkspacesConnection) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&connection.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (connection *WorkspacesConnection) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*WorkspacesConnection)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, connection)
}

// AssignPropertiesFromWorkspacesConnection populates our WorkspacesConnection from the provided source WorkspacesConnection
func (connection *WorkspacesConnection) AssignPropertiesFromWorkspacesConnection(source *v20210701s.WorkspacesConnection) error {

	// ObjectMeta
	connection.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec WorkspacesConnection_Spec
	err := spec.AssignPropertiesFromWorkspacesConnection_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromWorkspacesConnection_Spec() to populate field Spec")
	}
	connection.Spec = spec

	// Status
	var status WorkspacesConnection_STATUS
	err = status.AssignPropertiesFromWorkspacesConnection_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromWorkspacesConnection_STATUS() to populate field Status")
	}
	connection.Status = status

	// No error
	return nil
}

// AssignPropertiesToWorkspacesConnection populates the provided destination WorkspacesConnection from our WorkspacesConnection
func (connection *WorkspacesConnection) AssignPropertiesToWorkspacesConnection(destination *v20210701s.WorkspacesConnection) error {

	// ObjectMeta
	destination.ObjectMeta = *connection.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210701s.WorkspacesConnection_Spec
	err := connection.Spec.AssignPropertiesToWorkspacesConnection_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToWorkspacesConnection_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210701s.WorkspacesConnection_STATUS
	err = connection.Status.AssignPropertiesToWorkspacesConnection_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToWorkspacesConnection_STATUS() to populate field Status")
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

type WorkspacesConnection_Spec struct {
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

var _ genruntime.ARMTransformer = &WorkspacesConnection_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (connection *WorkspacesConnection_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if connection == nil {
		return nil, nil
	}
	result := &WorkspacesConnection_SpecARM{}

	// Set property ‘AzureName’:
	result.AzureName = connection.AzureName

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if connection.AuthType != nil ||
		connection.Category != nil ||
		connection.Target != nil ||
		connection.Value != nil ||
		connection.ValueFormat != nil {
		result.Properties = &WorkspaceConnectionPropsARM{}
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
		valueFormat := *connection.ValueFormat
		result.Properties.ValueFormat = &valueFormat
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (connection *WorkspacesConnection_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &WorkspacesConnection_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (connection *WorkspacesConnection_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(WorkspacesConnection_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected WorkspacesConnection_SpecARM, got %T", armInput)
	}

	// Set property ‘AuthType’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AuthType != nil {
			authType := *typedInput.Properties.AuthType
			connection.AuthType = &authType
		}
	}

	// Set property ‘AzureName’:
	connection.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Category’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Category != nil {
			category := *typedInput.Properties.Category
			connection.Category = &category
		}
	}

	// Set property ‘Owner’:
	connection.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Target’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Target != nil {
			target := *typedInput.Properties.Target
			connection.Target = &target
		}
	}

	// Set property ‘Value’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Value != nil {
			value := *typedInput.Properties.Value
			connection.Value = &value
		}
	}

	// Set property ‘ValueFormat’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ValueFormat != nil {
			valueFormat := *typedInput.Properties.ValueFormat
			connection.ValueFormat = &valueFormat
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &WorkspacesConnection_Spec{}

// ConvertSpecFrom populates our WorkspacesConnection_Spec from the provided source
func (connection *WorkspacesConnection_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210701s.WorkspacesConnection_Spec)
	if ok {
		// Populate our instance from source
		return connection.AssignPropertiesFromWorkspacesConnection_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210701s.WorkspacesConnection_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = connection.AssignPropertiesFromWorkspacesConnection_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our WorkspacesConnection_Spec
func (connection *WorkspacesConnection_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210701s.WorkspacesConnection_Spec)
	if ok {
		// Populate destination from our instance
		return connection.AssignPropertiesToWorkspacesConnection_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210701s.WorkspacesConnection_Spec{}
	err := connection.AssignPropertiesToWorkspacesConnection_Spec(dst)
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

// AssignPropertiesFromWorkspacesConnection_Spec populates our WorkspacesConnection_Spec from the provided source WorkspacesConnection_Spec
func (connection *WorkspacesConnection_Spec) AssignPropertiesFromWorkspacesConnection_Spec(source *v20210701s.WorkspacesConnection_Spec) error {

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
		valueFormat := WorkspaceConnectionProps_ValueFormat(*source.ValueFormat)
		connection.ValueFormat = &valueFormat
	} else {
		connection.ValueFormat = nil
	}

	// No error
	return nil
}

// AssignPropertiesToWorkspacesConnection_Spec populates the provided destination WorkspacesConnection_Spec from our WorkspacesConnection_Spec
func (connection *WorkspacesConnection_Spec) AssignPropertiesToWorkspacesConnection_Spec(destination *v20210701s.WorkspacesConnection_Spec) error {
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

// OriginalVersion returns the original API version used to create the resource.
func (connection *WorkspacesConnection_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (connection *WorkspacesConnection_Spec) SetAzureName(azureName string) {
	connection.AzureName = azureName
}

type WorkspacesConnection_STATUS struct {
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

var _ genruntime.ConvertibleStatus = &WorkspacesConnection_STATUS{}

// ConvertStatusFrom populates our WorkspacesConnection_STATUS from the provided source
func (connection *WorkspacesConnection_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210701s.WorkspacesConnection_STATUS)
	if ok {
		// Populate our instance from source
		return connection.AssignPropertiesFromWorkspacesConnection_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210701s.WorkspacesConnection_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = connection.AssignPropertiesFromWorkspacesConnection_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our WorkspacesConnection_STATUS
func (connection *WorkspacesConnection_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210701s.WorkspacesConnection_STATUS)
	if ok {
		// Populate destination from our instance
		return connection.AssignPropertiesToWorkspacesConnection_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210701s.WorkspacesConnection_STATUS{}
	err := connection.AssignPropertiesToWorkspacesConnection_STATUS(dst)
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

var _ genruntime.FromARMConverter = &WorkspacesConnection_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (connection *WorkspacesConnection_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &WorkspacesConnection_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (connection *WorkspacesConnection_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(WorkspacesConnection_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected WorkspacesConnection_STATUSARM, got %T", armInput)
	}

	// Set property ‘AuthType’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AuthType != nil {
			authType := *typedInput.Properties.AuthType
			connection.AuthType = &authType
		}
	}

	// Set property ‘Category’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Category != nil {
			category := *typedInput.Properties.Category
			connection.Category = &category
		}
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		connection.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		connection.Name = &name
	}

	// Set property ‘Target’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Target != nil {
			target := *typedInput.Properties.Target
			connection.Target = &target
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		connection.Type = &typeVar
	}

	// Set property ‘Value’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Value != nil {
			value := *typedInput.Properties.Value
			connection.Value = &value
		}
	}

	// Set property ‘ValueFormat’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ValueFormat != nil {
			valueFormat := *typedInput.Properties.ValueFormat
			connection.ValueFormat = &valueFormat
		}
	}

	// No error
	return nil
}

// AssignPropertiesFromWorkspacesConnection_STATUS populates our WorkspacesConnection_STATUS from the provided source WorkspacesConnection_STATUS
func (connection *WorkspacesConnection_STATUS) AssignPropertiesFromWorkspacesConnection_STATUS(source *v20210701s.WorkspacesConnection_STATUS) error {

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
		valueFormat := WorkspaceConnectionProps_ValueFormat_STATUS(*source.ValueFormat)
		connection.ValueFormat = &valueFormat
	} else {
		connection.ValueFormat = nil
	}

	// No error
	return nil
}

// AssignPropertiesToWorkspacesConnection_STATUS populates the provided destination WorkspacesConnection_STATUS from our WorkspacesConnection_STATUS
func (connection *WorkspacesConnection_STATUS) AssignPropertiesToWorkspacesConnection_STATUS(destination *v20210701s.WorkspacesConnection_STATUS) error {
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

type WorkspaceConnectionProps_ValueFormat_STATUS string

const WorkspaceConnectionProps_ValueFormat_JSON_STATUS = WorkspaceConnectionProps_ValueFormat_STATUS("JSON")

func init() {
	SchemeBuilder.Register(&WorkspacesConnection{}, &WorkspacesConnectionList{})
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import (
	"fmt"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1beta20211101storage"
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
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/DatabaseAdvancedThreatProtectionSettings.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advancedThreatProtectionSettings/Default
type ServersDatabasesAdvancedThreatProtectionSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_Databases_AdvancedThreatProtectionSetting_Spec   `json:"spec,omitempty"`
	Status            Servers_Databases_AdvancedThreatProtectionSetting_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersDatabasesAdvancedThreatProtectionSetting{}

// GetConditions returns the conditions of the resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ conversion.Convertible = &ServersDatabasesAdvancedThreatProtectionSetting{}

// ConvertFrom populates our ServersDatabasesAdvancedThreatProtectionSetting from the provided hub ServersDatabasesAdvancedThreatProtectionSetting
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20211101s.ServersDatabasesAdvancedThreatProtectionSetting)
	if !ok {
		return fmt.Errorf("expected sql/v1beta20211101storage/ServersDatabasesAdvancedThreatProtectionSetting but received %T instead", hub)
	}

	return setting.AssignProperties_From_ServersDatabasesAdvancedThreatProtectionSetting(source)
}

// ConvertTo populates the provided hub ServersDatabasesAdvancedThreatProtectionSetting from our ServersDatabasesAdvancedThreatProtectionSetting
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20211101s.ServersDatabasesAdvancedThreatProtectionSetting)
	if !ok {
		return fmt.Errorf("expected sql/v1beta20211101storage/ServersDatabasesAdvancedThreatProtectionSetting but received %T instead", hub)
	}

	return setting.AssignProperties_To_ServersDatabasesAdvancedThreatProtectionSetting(destination)
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1beta20211101-serversdatabasesadvancedthreatprotectionsetting,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversdatabasesadvancedthreatprotectionsettings,verbs=create;update,versions=v1beta20211101,name=default.v1beta20211101.serversdatabasesadvancedthreatprotectionsettings.sql.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &ServersDatabasesAdvancedThreatProtectionSetting{}

// Default applies defaults to the ServersDatabasesAdvancedThreatProtectionSetting resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) Default() {
	setting.defaultImpl()
	var temp any = setting
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the ServersDatabasesAdvancedThreatProtectionSetting resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) defaultImpl() {}

var _ genruntime.ImportableResource = &ServersDatabasesAdvancedThreatProtectionSetting{}

// InitializeSpec initializes the spec for this resource from the given status
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Servers_Databases_AdvancedThreatProtectionSetting_STATUS); ok {
		return setting.Spec.Initialize_From_Servers_Databases_AdvancedThreatProtectionSetting_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Servers_Databases_AdvancedThreatProtectionSetting_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &ServersDatabasesAdvancedThreatProtectionSetting{}

// AzureName returns the Azure name of the resource (always "Default")
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) AzureName() string {
	return "Default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (setting ServersDatabasesAdvancedThreatProtectionSetting) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/databases/advancedThreatProtectionSettings"
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) GetType() string {
	return "Microsoft.Sql/servers/databases/advancedThreatProtectionSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Servers_Databases_AdvancedThreatProtectionSetting_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  setting.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Servers_Databases_AdvancedThreatProtectionSetting_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_Databases_AdvancedThreatProtectionSetting_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1beta20211101-serversdatabasesadvancedthreatprotectionsetting,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversdatabasesadvancedthreatprotectionsettings,verbs=create;update,versions=v1beta20211101,name=validate.v1beta20211101.serversdatabasesadvancedthreatprotectionsettings.sql.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &ServersDatabasesAdvancedThreatProtectionSetting{}

// ValidateCreate validates the creation of the resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) ValidateCreate() error {
	validations := setting.createValidations()
	var temp any = setting
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
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) ValidateDelete() error {
	validations := setting.deleteValidations()
	var temp any = setting
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
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) ValidateUpdate(old runtime.Object) error {
	validations := setting.updateValidations()
	var temp any = setting
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
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) createValidations() []func() error {
	return []func() error{setting.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return setting.validateResourceReferences()
		},
		setting.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&setting.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*ServersDatabasesAdvancedThreatProtectionSetting)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, setting)
}

// AssignProperties_From_ServersDatabasesAdvancedThreatProtectionSetting populates our ServersDatabasesAdvancedThreatProtectionSetting from the provided source ServersDatabasesAdvancedThreatProtectionSetting
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) AssignProperties_From_ServersDatabasesAdvancedThreatProtectionSetting(source *v20211101s.ServersDatabasesAdvancedThreatProtectionSetting) error {

	// ObjectMeta
	setting.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Servers_Databases_AdvancedThreatProtectionSetting_Spec
	err := spec.AssignProperties_From_Servers_Databases_AdvancedThreatProtectionSetting_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_Databases_AdvancedThreatProtectionSetting_Spec() to populate field Spec")
	}
	setting.Spec = spec

	// Status
	var status Servers_Databases_AdvancedThreatProtectionSetting_STATUS
	err = status.AssignProperties_From_Servers_Databases_AdvancedThreatProtectionSetting_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_Databases_AdvancedThreatProtectionSetting_STATUS() to populate field Status")
	}
	setting.Status = status

	// No error
	return nil
}

// AssignProperties_To_ServersDatabasesAdvancedThreatProtectionSetting populates the provided destination ServersDatabasesAdvancedThreatProtectionSetting from our ServersDatabasesAdvancedThreatProtectionSetting
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) AssignProperties_To_ServersDatabasesAdvancedThreatProtectionSetting(destination *v20211101s.ServersDatabasesAdvancedThreatProtectionSetting) error {

	// ObjectMeta
	destination.ObjectMeta = *setting.ObjectMeta.DeepCopy()

	// Spec
	var spec v20211101s.Servers_Databases_AdvancedThreatProtectionSetting_Spec
	err := setting.Spec.AssignProperties_To_Servers_Databases_AdvancedThreatProtectionSetting_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_Databases_AdvancedThreatProtectionSetting_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20211101s.Servers_Databases_AdvancedThreatProtectionSetting_STATUS
	err = setting.Status.AssignProperties_To_Servers_Databases_AdvancedThreatProtectionSetting_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_Databases_AdvancedThreatProtectionSetting_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion(),
		Kind:    "ServersDatabasesAdvancedThreatProtectionSetting",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/DatabaseAdvancedThreatProtectionSettings.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advancedThreatProtectionSettings/Default
type ServersDatabasesAdvancedThreatProtectionSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersDatabasesAdvancedThreatProtectionSetting `json:"items"`
}

type Servers_Databases_AdvancedThreatProtectionSetting_Spec struct {
	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/ServersDatabase resource
	Owner *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"ServersDatabase"`

	// +kubebuilder:validation:Required
	// State: Specifies the state of the Advanced Threat Protection, whether it is enabled or disabled or a state has not been
	// applied yet on the specific database or server.
	State *AdvancedThreatProtectionProperties_State `json:"state,omitempty"`
}

var _ genruntime.ARMTransformer = &Servers_Databases_AdvancedThreatProtectionSetting_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (setting *Servers_Databases_AdvancedThreatProtectionSetting_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if setting == nil {
		return nil, nil
	}
	result := &Servers_Databases_AdvancedThreatProtectionSetting_Spec_ARM{}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if setting.State != nil {
		result.Properties = &AdvancedThreatProtectionProperties_ARM{}
	}
	if setting.State != nil {
		state := *setting.State
		result.Properties.State = &state
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (setting *Servers_Databases_AdvancedThreatProtectionSetting_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_Databases_AdvancedThreatProtectionSetting_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (setting *Servers_Databases_AdvancedThreatProtectionSetting_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_Databases_AdvancedThreatProtectionSetting_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_Databases_AdvancedThreatProtectionSetting_Spec_ARM, got %T", armInput)
	}

	// Set property ‘Owner’:
	setting.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// Set property ‘State’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.State != nil {
			state := *typedInput.Properties.State
			setting.State = &state
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Servers_Databases_AdvancedThreatProtectionSetting_Spec{}

// ConvertSpecFrom populates our Servers_Databases_AdvancedThreatProtectionSetting_Spec from the provided source
func (setting *Servers_Databases_AdvancedThreatProtectionSetting_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20211101s.Servers_Databases_AdvancedThreatProtectionSetting_Spec)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_Servers_Databases_AdvancedThreatProtectionSetting_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Servers_Databases_AdvancedThreatProtectionSetting_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_Servers_Databases_AdvancedThreatProtectionSetting_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Servers_Databases_AdvancedThreatProtectionSetting_Spec
func (setting *Servers_Databases_AdvancedThreatProtectionSetting_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20211101s.Servers_Databases_AdvancedThreatProtectionSetting_Spec)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_Servers_Databases_AdvancedThreatProtectionSetting_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Servers_Databases_AdvancedThreatProtectionSetting_Spec{}
	err := setting.AssignProperties_To_Servers_Databases_AdvancedThreatProtectionSetting_Spec(dst)
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

// AssignProperties_From_Servers_Databases_AdvancedThreatProtectionSetting_Spec populates our Servers_Databases_AdvancedThreatProtectionSetting_Spec from the provided source Servers_Databases_AdvancedThreatProtectionSetting_Spec
func (setting *Servers_Databases_AdvancedThreatProtectionSetting_Spec) AssignProperties_From_Servers_Databases_AdvancedThreatProtectionSetting_Spec(source *v20211101s.Servers_Databases_AdvancedThreatProtectionSetting_Spec) error {

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		setting.Owner = &owner
	} else {
		setting.Owner = nil
	}

	// State
	if source.State != nil {
		state := AdvancedThreatProtectionProperties_State(*source.State)
		setting.State = &state
	} else {
		setting.State = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Servers_Databases_AdvancedThreatProtectionSetting_Spec populates the provided destination Servers_Databases_AdvancedThreatProtectionSetting_Spec from our Servers_Databases_AdvancedThreatProtectionSetting_Spec
func (setting *Servers_Databases_AdvancedThreatProtectionSetting_Spec) AssignProperties_To_Servers_Databases_AdvancedThreatProtectionSetting_Spec(destination *v20211101s.Servers_Databases_AdvancedThreatProtectionSetting_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// OriginalVersion
	destination.OriginalVersion = setting.OriginalVersion()

	// Owner
	if setting.Owner != nil {
		owner := setting.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// State
	if setting.State != nil {
		state := string(*setting.State)
		destination.State = &state
	} else {
		destination.State = nil
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

// Initialize_From_Servers_Databases_AdvancedThreatProtectionSetting_STATUS populates our Servers_Databases_AdvancedThreatProtectionSetting_Spec from the provided source Servers_Databases_AdvancedThreatProtectionSetting_STATUS
func (setting *Servers_Databases_AdvancedThreatProtectionSetting_Spec) Initialize_From_Servers_Databases_AdvancedThreatProtectionSetting_STATUS(source *Servers_Databases_AdvancedThreatProtectionSetting_STATUS) error {

	// State
	if source.State != nil {
		state := AdvancedThreatProtectionProperties_State(*source.State)
		setting.State = &state
	} else {
		setting.State = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (setting *Servers_Databases_AdvancedThreatProtectionSetting_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type Servers_Databases_AdvancedThreatProtectionSetting_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// CreationTime: Specifies the UTC creation time of the policy.
	CreationTime *string `json:"creationTime,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// State: Specifies the state of the Advanced Threat Protection, whether it is enabled or disabled or a state has not been
	// applied yet on the specific database or server.
	State *AdvancedThreatProtectionProperties_State_STATUS `json:"state,omitempty"`

	// SystemData: SystemData of AdvancedThreatProtectionResource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_Databases_AdvancedThreatProtectionSetting_STATUS{}

// ConvertStatusFrom populates our Servers_Databases_AdvancedThreatProtectionSetting_STATUS from the provided source
func (setting *Servers_Databases_AdvancedThreatProtectionSetting_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20211101s.Servers_Databases_AdvancedThreatProtectionSetting_STATUS)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_Servers_Databases_AdvancedThreatProtectionSetting_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Servers_Databases_AdvancedThreatProtectionSetting_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_Servers_Databases_AdvancedThreatProtectionSetting_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Servers_Databases_AdvancedThreatProtectionSetting_STATUS
func (setting *Servers_Databases_AdvancedThreatProtectionSetting_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20211101s.Servers_Databases_AdvancedThreatProtectionSetting_STATUS)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_Servers_Databases_AdvancedThreatProtectionSetting_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Servers_Databases_AdvancedThreatProtectionSetting_STATUS{}
	err := setting.AssignProperties_To_Servers_Databases_AdvancedThreatProtectionSetting_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Servers_Databases_AdvancedThreatProtectionSetting_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (setting *Servers_Databases_AdvancedThreatProtectionSetting_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_Databases_AdvancedThreatProtectionSetting_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (setting *Servers_Databases_AdvancedThreatProtectionSetting_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_Databases_AdvancedThreatProtectionSetting_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_Databases_AdvancedThreatProtectionSetting_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘CreationTime’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.CreationTime != nil {
			creationTime := *typedInput.Properties.CreationTime
			setting.CreationTime = &creationTime
		}
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		setting.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		setting.Name = &name
	}

	// Set property ‘State’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.State != nil {
			state := *typedInput.Properties.State
			setting.State = &state
		}
	}

	// Set property ‘SystemData’:
	if typedInput.SystemData != nil {
		var systemData1 SystemData_STATUS
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		setting.SystemData = &systemData
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		setting.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Servers_Databases_AdvancedThreatProtectionSetting_STATUS populates our Servers_Databases_AdvancedThreatProtectionSetting_STATUS from the provided source Servers_Databases_AdvancedThreatProtectionSetting_STATUS
func (setting *Servers_Databases_AdvancedThreatProtectionSetting_STATUS) AssignProperties_From_Servers_Databases_AdvancedThreatProtectionSetting_STATUS(source *v20211101s.Servers_Databases_AdvancedThreatProtectionSetting_STATUS) error {

	// Conditions
	setting.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// CreationTime
	setting.CreationTime = genruntime.ClonePointerToString(source.CreationTime)

	// Id
	setting.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	setting.Name = genruntime.ClonePointerToString(source.Name)

	// State
	if source.State != nil {
		state := AdvancedThreatProtectionProperties_State_STATUS(*source.State)
		setting.State = &state
	} else {
		setting.State = nil
	}

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		setting.SystemData = &systemDatum
	} else {
		setting.SystemData = nil
	}

	// Type
	setting.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_Servers_Databases_AdvancedThreatProtectionSetting_STATUS populates the provided destination Servers_Databases_AdvancedThreatProtectionSetting_STATUS from our Servers_Databases_AdvancedThreatProtectionSetting_STATUS
func (setting *Servers_Databases_AdvancedThreatProtectionSetting_STATUS) AssignProperties_To_Servers_Databases_AdvancedThreatProtectionSetting_STATUS(destination *v20211101s.Servers_Databases_AdvancedThreatProtectionSetting_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(setting.Conditions)

	// CreationTime
	destination.CreationTime = genruntime.ClonePointerToString(setting.CreationTime)

	// Id
	destination.Id = genruntime.ClonePointerToString(setting.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(setting.Name)

	// State
	if setting.State != nil {
		state := string(*setting.State)
		destination.State = &state
	} else {
		destination.State = nil
	}

	// SystemData
	if setting.SystemData != nil {
		var systemDatum v20211101s.SystemData_STATUS
		err := setting.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(setting.Type)

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
	SchemeBuilder.Register(&ServersDatabasesAdvancedThreatProtectionSetting{}, &ServersDatabasesAdvancedThreatProtectionSettingList{})
}

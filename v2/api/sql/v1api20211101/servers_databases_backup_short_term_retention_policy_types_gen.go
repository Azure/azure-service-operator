// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"fmt"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101storage"
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
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/BackupShortTermRetentionPolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/backupShortTermRetentionPolicies/default
type ServersDatabasesBackupShortTermRetentionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_Databases_BackupShortTermRetentionPolicy_Spec   `json:"spec,omitempty"`
	Status            Servers_Databases_BackupShortTermRetentionPolicy_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersDatabasesBackupShortTermRetentionPolicy{}

// GetConditions returns the conditions of the resource
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) GetConditions() conditions.Conditions {
	return policy.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) SetConditions(conditions conditions.Conditions) {
	policy.Status.Conditions = conditions
}

var _ conversion.Convertible = &ServersDatabasesBackupShortTermRetentionPolicy{}

// ConvertFrom populates our ServersDatabasesBackupShortTermRetentionPolicy from the provided hub ServersDatabasesBackupShortTermRetentionPolicy
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20211101s.ServersDatabasesBackupShortTermRetentionPolicy)
	if !ok {
		return fmt.Errorf("expected sql/v1api20211101storage/ServersDatabasesBackupShortTermRetentionPolicy but received %T instead", hub)
	}

	return policy.AssignProperties_From_ServersDatabasesBackupShortTermRetentionPolicy(source)
}

// ConvertTo populates the provided hub ServersDatabasesBackupShortTermRetentionPolicy from our ServersDatabasesBackupShortTermRetentionPolicy
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20211101s.ServersDatabasesBackupShortTermRetentionPolicy)
	if !ok {
		return fmt.Errorf("expected sql/v1api20211101storage/ServersDatabasesBackupShortTermRetentionPolicy but received %T instead", hub)
	}

	return policy.AssignProperties_To_ServersDatabasesBackupShortTermRetentionPolicy(destination)
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1api20211101-serversdatabasesbackupshorttermretentionpolicy,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversdatabasesbackupshorttermretentionpolicies,verbs=create;update,versions=v1api20211101,name=default.v1api20211101.serversdatabasesbackupshorttermretentionpolicies.sql.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &ServersDatabasesBackupShortTermRetentionPolicy{}

// Default applies defaults to the ServersDatabasesBackupShortTermRetentionPolicy resource
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) Default() {
	policy.defaultImpl()
	var temp any = policy
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the ServersDatabasesBackupShortTermRetentionPolicy resource
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) defaultImpl() {}

var _ genruntime.ImportableResource = &ServersDatabasesBackupShortTermRetentionPolicy{}

// InitializeSpec initializes the spec for this resource from the given status
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Servers_Databases_BackupShortTermRetentionPolicy_STATUS); ok {
		return policy.Spec.Initialize_From_Servers_Databases_BackupShortTermRetentionPolicy_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Servers_Databases_BackupShortTermRetentionPolicy_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &ServersDatabasesBackupShortTermRetentionPolicy{}

// AzureName returns the Azure name of the resource (always "default")
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (policy ServersDatabasesBackupShortTermRetentionPolicy) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) GetSpec() genruntime.ConvertibleSpec {
	return &policy.Spec
}

// GetStatus returns the status of this resource
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) GetStatus() genruntime.ConvertibleStatus {
	return &policy.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/databases/backupShortTermRetentionPolicies"
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) GetType() string {
	return "Microsoft.Sql/servers/databases/backupShortTermRetentionPolicies"
}

// NewEmptyStatus returns a new empty (blank) status
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Servers_Databases_BackupShortTermRetentionPolicy_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(policy.Spec)
	return policy.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Servers_Databases_BackupShortTermRetentionPolicy_STATUS); ok {
		policy.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_Databases_BackupShortTermRetentionPolicy_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	policy.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1api20211101-serversdatabasesbackupshorttermretentionpolicy,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversdatabasesbackupshorttermretentionpolicies,verbs=create;update,versions=v1api20211101,name=validate.v1api20211101.serversdatabasesbackupshorttermretentionpolicies.sql.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &ServersDatabasesBackupShortTermRetentionPolicy{}

// ValidateCreate validates the creation of the resource
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) ValidateCreate() (admission.Warnings, error) {
	validations := policy.createValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) ValidateDelete() (admission.Warnings, error) {
	validations := policy.deleteValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := policy.updateValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){policy.validateResourceReferences, policy.validateOwnerReference}
}

// deleteValidations validates the deletion of the resource
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return policy.validateResourceReferences()
		},
		policy.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return policy.validateOwnerReference()
		},
	}
}

// validateOwnerReference validates the owner field
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(policy)
}

// validateResourceReferences validates all resource references
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&policy.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*ServersDatabasesBackupShortTermRetentionPolicy)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, policy)
}

// AssignProperties_From_ServersDatabasesBackupShortTermRetentionPolicy populates our ServersDatabasesBackupShortTermRetentionPolicy from the provided source ServersDatabasesBackupShortTermRetentionPolicy
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) AssignProperties_From_ServersDatabasesBackupShortTermRetentionPolicy(source *v20211101s.ServersDatabasesBackupShortTermRetentionPolicy) error {

	// ObjectMeta
	policy.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Servers_Databases_BackupShortTermRetentionPolicy_Spec
	err := spec.AssignProperties_From_Servers_Databases_BackupShortTermRetentionPolicy_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_Databases_BackupShortTermRetentionPolicy_Spec() to populate field Spec")
	}
	policy.Spec = spec

	// Status
	var status Servers_Databases_BackupShortTermRetentionPolicy_STATUS
	err = status.AssignProperties_From_Servers_Databases_BackupShortTermRetentionPolicy_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_Databases_BackupShortTermRetentionPolicy_STATUS() to populate field Status")
	}
	policy.Status = status

	// No error
	return nil
}

// AssignProperties_To_ServersDatabasesBackupShortTermRetentionPolicy populates the provided destination ServersDatabasesBackupShortTermRetentionPolicy from our ServersDatabasesBackupShortTermRetentionPolicy
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) AssignProperties_To_ServersDatabasesBackupShortTermRetentionPolicy(destination *v20211101s.ServersDatabasesBackupShortTermRetentionPolicy) error {

	// ObjectMeta
	destination.ObjectMeta = *policy.ObjectMeta.DeepCopy()

	// Spec
	var spec v20211101s.Servers_Databases_BackupShortTermRetentionPolicy_Spec
	err := policy.Spec.AssignProperties_To_Servers_Databases_BackupShortTermRetentionPolicy_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_Databases_BackupShortTermRetentionPolicy_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20211101s.Servers_Databases_BackupShortTermRetentionPolicy_STATUS
	err = policy.Status.AssignProperties_To_Servers_Databases_BackupShortTermRetentionPolicy_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_Databases_BackupShortTermRetentionPolicy_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (policy *ServersDatabasesBackupShortTermRetentionPolicy) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: policy.Spec.OriginalVersion(),
		Kind:    "ServersDatabasesBackupShortTermRetentionPolicy",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/BackupShortTermRetentionPolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/backupShortTermRetentionPolicies/default
type ServersDatabasesBackupShortTermRetentionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersDatabasesBackupShortTermRetentionPolicy `json:"items"`
}

type Servers_Databases_BackupShortTermRetentionPolicy_Spec struct {
	// DiffBackupIntervalInHours: The differential backup interval in hours. This is how many interval hours between each
	// differential backup will be supported. This is only applicable to live databases but not dropped databases.
	DiffBackupIntervalInHours *BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours `json:"diffBackupIntervalInHours,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/ServersDatabase resource
	Owner *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"ServersDatabase"`

	// RetentionDays: The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
	RetentionDays *int `json:"retentionDays,omitempty"`
}

var _ genruntime.ARMTransformer = &Servers_Databases_BackupShortTermRetentionPolicy_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (policy *Servers_Databases_BackupShortTermRetentionPolicy_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if policy == nil {
		return nil, nil
	}
	result := &Servers_Databases_BackupShortTermRetentionPolicy_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if policy.DiffBackupIntervalInHours != nil || policy.RetentionDays != nil {
		result.Properties = &BackupShortTermRetentionPolicyProperties_ARM{}
	}
	if policy.DiffBackupIntervalInHours != nil {
		diffBackupIntervalInHours := *policy.DiffBackupIntervalInHours
		result.Properties.DiffBackupIntervalInHours = &diffBackupIntervalInHours
	}
	if policy.RetentionDays != nil {
		retentionDays := *policy.RetentionDays
		result.Properties.RetentionDays = &retentionDays
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (policy *Servers_Databases_BackupShortTermRetentionPolicy_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_Databases_BackupShortTermRetentionPolicy_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (policy *Servers_Databases_BackupShortTermRetentionPolicy_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_Databases_BackupShortTermRetentionPolicy_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_Databases_BackupShortTermRetentionPolicy_Spec_ARM, got %T", armInput)
	}

	// Set property "DiffBackupIntervalInHours":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DiffBackupIntervalInHours != nil {
			diffBackupIntervalInHours := *typedInput.Properties.DiffBackupIntervalInHours
			policy.DiffBackupIntervalInHours = &diffBackupIntervalInHours
		}
	}

	// Set property "Owner":
	policy.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "RetentionDays":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.RetentionDays != nil {
			retentionDays := *typedInput.Properties.RetentionDays
			policy.RetentionDays = &retentionDays
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Servers_Databases_BackupShortTermRetentionPolicy_Spec{}

// ConvertSpecFrom populates our Servers_Databases_BackupShortTermRetentionPolicy_Spec from the provided source
func (policy *Servers_Databases_BackupShortTermRetentionPolicy_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20211101s.Servers_Databases_BackupShortTermRetentionPolicy_Spec)
	if ok {
		// Populate our instance from source
		return policy.AssignProperties_From_Servers_Databases_BackupShortTermRetentionPolicy_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Servers_Databases_BackupShortTermRetentionPolicy_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_Servers_Databases_BackupShortTermRetentionPolicy_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Servers_Databases_BackupShortTermRetentionPolicy_Spec
func (policy *Servers_Databases_BackupShortTermRetentionPolicy_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20211101s.Servers_Databases_BackupShortTermRetentionPolicy_Spec)
	if ok {
		// Populate destination from our instance
		return policy.AssignProperties_To_Servers_Databases_BackupShortTermRetentionPolicy_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Servers_Databases_BackupShortTermRetentionPolicy_Spec{}
	err := policy.AssignProperties_To_Servers_Databases_BackupShortTermRetentionPolicy_Spec(dst)
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

// AssignProperties_From_Servers_Databases_BackupShortTermRetentionPolicy_Spec populates our Servers_Databases_BackupShortTermRetentionPolicy_Spec from the provided source Servers_Databases_BackupShortTermRetentionPolicy_Spec
func (policy *Servers_Databases_BackupShortTermRetentionPolicy_Spec) AssignProperties_From_Servers_Databases_BackupShortTermRetentionPolicy_Spec(source *v20211101s.Servers_Databases_BackupShortTermRetentionPolicy_Spec) error {

	// DiffBackupIntervalInHours
	if source.DiffBackupIntervalInHours != nil {
		diffBackupIntervalInHour := BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours(*source.DiffBackupIntervalInHours)
		policy.DiffBackupIntervalInHours = &diffBackupIntervalInHour
	} else {
		policy.DiffBackupIntervalInHours = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		policy.Owner = &owner
	} else {
		policy.Owner = nil
	}

	// RetentionDays
	policy.RetentionDays = genruntime.ClonePointerToInt(source.RetentionDays)

	// No error
	return nil
}

// AssignProperties_To_Servers_Databases_BackupShortTermRetentionPolicy_Spec populates the provided destination Servers_Databases_BackupShortTermRetentionPolicy_Spec from our Servers_Databases_BackupShortTermRetentionPolicy_Spec
func (policy *Servers_Databases_BackupShortTermRetentionPolicy_Spec) AssignProperties_To_Servers_Databases_BackupShortTermRetentionPolicy_Spec(destination *v20211101s.Servers_Databases_BackupShortTermRetentionPolicy_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// DiffBackupIntervalInHours
	if policy.DiffBackupIntervalInHours != nil {
		diffBackupIntervalInHour := int(*policy.DiffBackupIntervalInHours)
		destination.DiffBackupIntervalInHours = &diffBackupIntervalInHour
	} else {
		destination.DiffBackupIntervalInHours = nil
	}

	// OriginalVersion
	destination.OriginalVersion = policy.OriginalVersion()

	// Owner
	if policy.Owner != nil {
		owner := policy.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// RetentionDays
	destination.RetentionDays = genruntime.ClonePointerToInt(policy.RetentionDays)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_Servers_Databases_BackupShortTermRetentionPolicy_STATUS populates our Servers_Databases_BackupShortTermRetentionPolicy_Spec from the provided source Servers_Databases_BackupShortTermRetentionPolicy_STATUS
func (policy *Servers_Databases_BackupShortTermRetentionPolicy_Spec) Initialize_From_Servers_Databases_BackupShortTermRetentionPolicy_STATUS(source *Servers_Databases_BackupShortTermRetentionPolicy_STATUS) error {

	// DiffBackupIntervalInHours
	if source.DiffBackupIntervalInHours != nil {
		diffBackupIntervalInHour := BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours(*source.DiffBackupIntervalInHours)
		policy.DiffBackupIntervalInHours = &diffBackupIntervalInHour
	} else {
		policy.DiffBackupIntervalInHours = nil
	}

	// RetentionDays
	policy.RetentionDays = genruntime.ClonePointerToInt(source.RetentionDays)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (policy *Servers_Databases_BackupShortTermRetentionPolicy_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type Servers_Databases_BackupShortTermRetentionPolicy_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// DiffBackupIntervalInHours: The differential backup interval in hours. This is how many interval hours between each
	// differential backup will be supported. This is only applicable to live databases but not dropped databases.
	DiffBackupIntervalInHours *BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_STATUS `json:"diffBackupIntervalInHours,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// RetentionDays: The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
	RetentionDays *int `json:"retentionDays,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_Databases_BackupShortTermRetentionPolicy_STATUS{}

// ConvertStatusFrom populates our Servers_Databases_BackupShortTermRetentionPolicy_STATUS from the provided source
func (policy *Servers_Databases_BackupShortTermRetentionPolicy_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20211101s.Servers_Databases_BackupShortTermRetentionPolicy_STATUS)
	if ok {
		// Populate our instance from source
		return policy.AssignProperties_From_Servers_Databases_BackupShortTermRetentionPolicy_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Servers_Databases_BackupShortTermRetentionPolicy_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_Servers_Databases_BackupShortTermRetentionPolicy_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Servers_Databases_BackupShortTermRetentionPolicy_STATUS
func (policy *Servers_Databases_BackupShortTermRetentionPolicy_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20211101s.Servers_Databases_BackupShortTermRetentionPolicy_STATUS)
	if ok {
		// Populate destination from our instance
		return policy.AssignProperties_To_Servers_Databases_BackupShortTermRetentionPolicy_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Servers_Databases_BackupShortTermRetentionPolicy_STATUS{}
	err := policy.AssignProperties_To_Servers_Databases_BackupShortTermRetentionPolicy_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Servers_Databases_BackupShortTermRetentionPolicy_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (policy *Servers_Databases_BackupShortTermRetentionPolicy_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (policy *Servers_Databases_BackupShortTermRetentionPolicy_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "DiffBackupIntervalInHours":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DiffBackupIntervalInHours != nil {
			diffBackupIntervalInHours := *typedInput.Properties.DiffBackupIntervalInHours
			policy.DiffBackupIntervalInHours = &diffBackupIntervalInHours
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		policy.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		policy.Name = &name
	}

	// Set property "RetentionDays":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.RetentionDays != nil {
			retentionDays := *typedInput.Properties.RetentionDays
			policy.RetentionDays = &retentionDays
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		policy.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Servers_Databases_BackupShortTermRetentionPolicy_STATUS populates our Servers_Databases_BackupShortTermRetentionPolicy_STATUS from the provided source Servers_Databases_BackupShortTermRetentionPolicy_STATUS
func (policy *Servers_Databases_BackupShortTermRetentionPolicy_STATUS) AssignProperties_From_Servers_Databases_BackupShortTermRetentionPolicy_STATUS(source *v20211101s.Servers_Databases_BackupShortTermRetentionPolicy_STATUS) error {

	// Conditions
	policy.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// DiffBackupIntervalInHours
	if source.DiffBackupIntervalInHours != nil {
		diffBackupIntervalInHour := BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_STATUS(*source.DiffBackupIntervalInHours)
		policy.DiffBackupIntervalInHours = &diffBackupIntervalInHour
	} else {
		policy.DiffBackupIntervalInHours = nil
	}

	// Id
	policy.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	policy.Name = genruntime.ClonePointerToString(source.Name)

	// RetentionDays
	policy.RetentionDays = genruntime.ClonePointerToInt(source.RetentionDays)

	// Type
	policy.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_Servers_Databases_BackupShortTermRetentionPolicy_STATUS populates the provided destination Servers_Databases_BackupShortTermRetentionPolicy_STATUS from our Servers_Databases_BackupShortTermRetentionPolicy_STATUS
func (policy *Servers_Databases_BackupShortTermRetentionPolicy_STATUS) AssignProperties_To_Servers_Databases_BackupShortTermRetentionPolicy_STATUS(destination *v20211101s.Servers_Databases_BackupShortTermRetentionPolicy_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(policy.Conditions)

	// DiffBackupIntervalInHours
	if policy.DiffBackupIntervalInHours != nil {
		diffBackupIntervalInHour := int(*policy.DiffBackupIntervalInHours)
		destination.DiffBackupIntervalInHours = &diffBackupIntervalInHour
	} else {
		destination.DiffBackupIntervalInHours = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(policy.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(policy.Name)

	// RetentionDays
	destination.RetentionDays = genruntime.ClonePointerToInt(policy.RetentionDays)

	// Type
	destination.Type = genruntime.ClonePointerToString(policy.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={12,24}
type BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours int

const (
	BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_12 = BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours(12)
	BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_24 = BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours(24)
)

type BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_STATUS int

const (
	BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_STATUS_12 = BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_STATUS(12)
	BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_STATUS_24 = BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_STATUS(24)
)

func init() {
	SchemeBuilder.Register(&ServersDatabasesBackupShortTermRetentionPolicy{}, &ServersDatabasesBackupShortTermRetentionPolicyList{})
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240815

import (
	"fmt"
	arm "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/arm"
	storage "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
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
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-08-15/rbac.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlRoleAssignments/{roleAssignmentId}
type SqlRoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlRoleAssignment_Spec   `json:"spec,omitempty"`
	Status            SqlRoleAssignment_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlRoleAssignment{}

// GetConditions returns the conditions of the resource
func (assignment *SqlRoleAssignment) GetConditions() conditions.Conditions {
	return assignment.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (assignment *SqlRoleAssignment) SetConditions(conditions conditions.Conditions) {
	assignment.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlRoleAssignment{}

// ConvertFrom populates our SqlRoleAssignment from the provided hub SqlRoleAssignment
func (assignment *SqlRoleAssignment) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.SqlRoleAssignment)
	if !ok {
		return fmt.Errorf("expected documentdb/v1api20240815/storage/SqlRoleAssignment but received %T instead", hub)
	}

	return assignment.AssignProperties_From_SqlRoleAssignment(source)
}

// ConvertTo populates the provided hub SqlRoleAssignment from our SqlRoleAssignment
func (assignment *SqlRoleAssignment) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.SqlRoleAssignment)
	if !ok {
		return fmt.Errorf("expected documentdb/v1api20240815/storage/SqlRoleAssignment but received %T instead", hub)
	}

	return assignment.AssignProperties_To_SqlRoleAssignment(destination)
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1api20240815-sqlroleassignment,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqlroleassignments,verbs=create;update,versions=v1api20240815,name=default.v1api20240815.sqlroleassignments.documentdb.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &SqlRoleAssignment{}

// Default applies defaults to the SqlRoleAssignment resource
func (assignment *SqlRoleAssignment) Default() {
	assignment.defaultImpl()
	var temp any = assignment
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the SqlRoleAssignment resource
func (assignment *SqlRoleAssignment) defaultImpl() {}

var _ configmaps.Exporter = &SqlRoleAssignment{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (assignment *SqlRoleAssignment) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if assignment.Spec.OperatorSpec == nil {
		return nil
	}
	return assignment.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &SqlRoleAssignment{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (assignment *SqlRoleAssignment) SecretDestinationExpressions() []*core.DestinationExpression {
	if assignment.Spec.OperatorSpec == nil {
		return nil
	}
	return assignment.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.ImportableResource = &SqlRoleAssignment{}

// InitializeSpec initializes the spec for this resource from the given status
func (assignment *SqlRoleAssignment) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*SqlRoleAssignment_STATUS); ok {
		return assignment.Spec.Initialize_From_SqlRoleAssignment_STATUS(s)
	}

	return fmt.Errorf("expected Status of type SqlRoleAssignment_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &SqlRoleAssignment{}

// AzureName returns the Azure name of the resource
func (assignment *SqlRoleAssignment) AzureName() string {
	return assignment.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-08-15"
func (assignment SqlRoleAssignment) GetAPIVersion() string {
	return "2024-08-15"
}

// GetResourceScope returns the scope of the resource
func (assignment *SqlRoleAssignment) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (assignment *SqlRoleAssignment) GetSpec() genruntime.ConvertibleSpec {
	return &assignment.Spec
}

// GetStatus returns the status of this resource
func (assignment *SqlRoleAssignment) GetStatus() genruntime.ConvertibleStatus {
	return &assignment.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (assignment *SqlRoleAssignment) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlRoleAssignments"
func (assignment *SqlRoleAssignment) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlRoleAssignments"
}

// NewEmptyStatus returns a new empty (blank) status
func (assignment *SqlRoleAssignment) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlRoleAssignment_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (assignment *SqlRoleAssignment) Owner() *genruntime.ResourceReference {
	if assignment.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(assignment.Spec)
	return assignment.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (assignment *SqlRoleAssignment) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlRoleAssignment_STATUS); ok {
		assignment.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlRoleAssignment_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	assignment.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1api20240815-sqlroleassignment,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqlroleassignments,verbs=create;update,versions=v1api20240815,name=validate.v1api20240815.sqlroleassignments.documentdb.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &SqlRoleAssignment{}

// ValidateCreate validates the creation of the resource
func (assignment *SqlRoleAssignment) ValidateCreate() (admission.Warnings, error) {
	validations := assignment.createValidations()
	var temp any = assignment
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (assignment *SqlRoleAssignment) ValidateDelete() (admission.Warnings, error) {
	validations := assignment.deleteValidations()
	var temp any = assignment
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (assignment *SqlRoleAssignment) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := assignment.updateValidations()
	var temp any = assignment
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (assignment *SqlRoleAssignment) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){assignment.validateResourceReferences, assignment.validateOwnerReference, assignment.validateSecretDestinations, assignment.validateConfigMapDestinations, assignment.validateOptionalConfigMapReferences}
}

// deleteValidations validates the deletion of the resource
func (assignment *SqlRoleAssignment) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (assignment *SqlRoleAssignment) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return assignment.validateResourceReferences()
		},
		assignment.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return assignment.validateOwnerReference()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return assignment.validateSecretDestinations()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return assignment.validateConfigMapDestinations()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return assignment.validateOptionalConfigMapReferences()
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (assignment *SqlRoleAssignment) validateConfigMapDestinations() (admission.Warnings, error) {
	if assignment.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(assignment, nil, assignment.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOptionalConfigMapReferences validates all optional configmap reference pairs to ensure that at most 1 is set
func (assignment *SqlRoleAssignment) validateOptionalConfigMapReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindOptionalConfigMapReferences(&assignment.Spec)
	if err != nil {
		return nil, err
	}
	return configmaps.ValidateOptionalReferences(refs)
}

// validateOwnerReference validates the owner field
func (assignment *SqlRoleAssignment) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(assignment)
}

// validateResourceReferences validates all resource references
func (assignment *SqlRoleAssignment) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&assignment.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (assignment *SqlRoleAssignment) validateSecretDestinations() (admission.Warnings, error) {
	if assignment.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(assignment, nil, assignment.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (assignment *SqlRoleAssignment) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*SqlRoleAssignment)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, assignment)
}

// AssignProperties_From_SqlRoleAssignment populates our SqlRoleAssignment from the provided source SqlRoleAssignment
func (assignment *SqlRoleAssignment) AssignProperties_From_SqlRoleAssignment(source *storage.SqlRoleAssignment) error {

	// ObjectMeta
	assignment.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec SqlRoleAssignment_Spec
	err := spec.AssignProperties_From_SqlRoleAssignment_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_SqlRoleAssignment_Spec() to populate field Spec")
	}
	assignment.Spec = spec

	// Status
	var status SqlRoleAssignment_STATUS
	err = status.AssignProperties_From_SqlRoleAssignment_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_SqlRoleAssignment_STATUS() to populate field Status")
	}
	assignment.Status = status

	// No error
	return nil
}

// AssignProperties_To_SqlRoleAssignment populates the provided destination SqlRoleAssignment from our SqlRoleAssignment
func (assignment *SqlRoleAssignment) AssignProperties_To_SqlRoleAssignment(destination *storage.SqlRoleAssignment) error {

	// ObjectMeta
	destination.ObjectMeta = *assignment.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.SqlRoleAssignment_Spec
	err := assignment.Spec.AssignProperties_To_SqlRoleAssignment_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_SqlRoleAssignment_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.SqlRoleAssignment_STATUS
	err = assignment.Status.AssignProperties_To_SqlRoleAssignment_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_SqlRoleAssignment_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (assignment *SqlRoleAssignment) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: assignment.Spec.OriginalVersion(),
		Kind:    "SqlRoleAssignment",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-08-15/rbac.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlRoleAssignments/{roleAssignmentId}
type SqlRoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlRoleAssignment `json:"items"`
}

type SqlRoleAssignment_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
	// passed directly to Azure
	OperatorSpec *SqlRoleAssignmentOperatorSpec `json:"operatorSpec,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/DatabaseAccount resource
	Owner *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"DatabaseAccount"`

	// PrincipalId: The unique identifier for the associated AAD principal in the AAD graph to which access is being granted
	// through this Role Assignment. Tenant ID for the principal is inferred using the tenant associated with the subscription.
	PrincipalId *string `json:"principalId,omitempty" optionalConfigMapPair:"PrincipalId"`

	// PrincipalIdFromConfig: The unique identifier for the associated AAD principal in the AAD graph to which access is being
	// granted through this Role Assignment. Tenant ID for the principal is inferred using the tenant associated with the
	// subscription.
	PrincipalIdFromConfig *genruntime.ConfigMapReference `json:"principalIdFromConfig,omitempty" optionalConfigMapPair:"PrincipalId"`

	// RoleDefinitionId: The unique identifier for the associated Role Definition.
	RoleDefinitionId *string `json:"roleDefinitionId,omitempty"`

	// Scope: The data plane resource path for which access is being granted through this Role Assignment.
	Scope *string `json:"scope,omitempty"`
}

var _ genruntime.ARMTransformer = &SqlRoleAssignment_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (assignment *SqlRoleAssignment_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if assignment == nil {
		return nil, nil
	}
	result := &arm.SqlRoleAssignment_Spec{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if assignment.PrincipalId != nil ||
		assignment.PrincipalIdFromConfig != nil ||
		assignment.RoleDefinitionId != nil ||
		assignment.Scope != nil {
		result.Properties = &arm.SqlRoleAssignmentResource{}
	}
	if assignment.PrincipalId != nil {
		principalId := *assignment.PrincipalId
		result.Properties.PrincipalId = &principalId
	}
	if assignment.PrincipalIdFromConfig != nil {
		principalIdValue, err := resolved.ResolvedConfigMaps.Lookup(*assignment.PrincipalIdFromConfig)
		if err != nil {
			return nil, eris.Wrap(err, "looking up configmap for property PrincipalId")
		}
		principalId := principalIdValue
		result.Properties.PrincipalId = &principalId
	}
	if assignment.RoleDefinitionId != nil {
		roleDefinitionId := *assignment.RoleDefinitionId
		result.Properties.RoleDefinitionId = &roleDefinitionId
	}
	if assignment.Scope != nil {
		scope := *assignment.Scope
		result.Properties.Scope = &scope
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (assignment *SqlRoleAssignment_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.SqlRoleAssignment_Spec{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (assignment *SqlRoleAssignment_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.SqlRoleAssignment_Spec)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.SqlRoleAssignment_Spec, got %T", armInput)
	}

	// Set property "AzureName":
	assignment.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// no assignment for property "OperatorSpec"

	// Set property "Owner":
	assignment.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "PrincipalId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.PrincipalId != nil {
			principalId := *typedInput.Properties.PrincipalId
			assignment.PrincipalId = &principalId
		}
	}

	// no assignment for property "PrincipalIdFromConfig"

	// Set property "RoleDefinitionId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.RoleDefinitionId != nil {
			roleDefinitionId := *typedInput.Properties.RoleDefinitionId
			assignment.RoleDefinitionId = &roleDefinitionId
		}
	}

	// Set property "Scope":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Scope != nil {
			scope := *typedInput.Properties.Scope
			assignment.Scope = &scope
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &SqlRoleAssignment_Spec{}

// ConvertSpecFrom populates our SqlRoleAssignment_Spec from the provided source
func (assignment *SqlRoleAssignment_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.SqlRoleAssignment_Spec)
	if ok {
		// Populate our instance from source
		return assignment.AssignProperties_From_SqlRoleAssignment_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.SqlRoleAssignment_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = assignment.AssignProperties_From_SqlRoleAssignment_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our SqlRoleAssignment_Spec
func (assignment *SqlRoleAssignment_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.SqlRoleAssignment_Spec)
	if ok {
		// Populate destination from our instance
		return assignment.AssignProperties_To_SqlRoleAssignment_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.SqlRoleAssignment_Spec{}
	err := assignment.AssignProperties_To_SqlRoleAssignment_Spec(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_SqlRoleAssignment_Spec populates our SqlRoleAssignment_Spec from the provided source SqlRoleAssignment_Spec
func (assignment *SqlRoleAssignment_Spec) AssignProperties_From_SqlRoleAssignment_Spec(source *storage.SqlRoleAssignment_Spec) error {

	// AzureName
	assignment.AzureName = source.AzureName

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec SqlRoleAssignmentOperatorSpec
		err := operatorSpec.AssignProperties_From_SqlRoleAssignmentOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_SqlRoleAssignmentOperatorSpec() to populate field OperatorSpec")
		}
		assignment.OperatorSpec = &operatorSpec
	} else {
		assignment.OperatorSpec = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		assignment.Owner = &owner
	} else {
		assignment.Owner = nil
	}

	// PrincipalId
	assignment.PrincipalId = genruntime.ClonePointerToString(source.PrincipalId)

	// PrincipalIdFromConfig
	if source.PrincipalIdFromConfig != nil {
		principalIdFromConfig := source.PrincipalIdFromConfig.Copy()
		assignment.PrincipalIdFromConfig = &principalIdFromConfig
	} else {
		assignment.PrincipalIdFromConfig = nil
	}

	// RoleDefinitionId
	assignment.RoleDefinitionId = genruntime.ClonePointerToString(source.RoleDefinitionId)

	// Scope
	assignment.Scope = genruntime.ClonePointerToString(source.Scope)

	// No error
	return nil
}

// AssignProperties_To_SqlRoleAssignment_Spec populates the provided destination SqlRoleAssignment_Spec from our SqlRoleAssignment_Spec
func (assignment *SqlRoleAssignment_Spec) AssignProperties_To_SqlRoleAssignment_Spec(destination *storage.SqlRoleAssignment_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = assignment.AzureName

	// OperatorSpec
	if assignment.OperatorSpec != nil {
		var operatorSpec storage.SqlRoleAssignmentOperatorSpec
		err := assignment.OperatorSpec.AssignProperties_To_SqlRoleAssignmentOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_SqlRoleAssignmentOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = assignment.OriginalVersion()

	// Owner
	if assignment.Owner != nil {
		owner := assignment.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// PrincipalId
	destination.PrincipalId = genruntime.ClonePointerToString(assignment.PrincipalId)

	// PrincipalIdFromConfig
	if assignment.PrincipalIdFromConfig != nil {
		principalIdFromConfig := assignment.PrincipalIdFromConfig.Copy()
		destination.PrincipalIdFromConfig = &principalIdFromConfig
	} else {
		destination.PrincipalIdFromConfig = nil
	}

	// RoleDefinitionId
	destination.RoleDefinitionId = genruntime.ClonePointerToString(assignment.RoleDefinitionId)

	// Scope
	destination.Scope = genruntime.ClonePointerToString(assignment.Scope)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_SqlRoleAssignment_STATUS populates our SqlRoleAssignment_Spec from the provided source SqlRoleAssignment_STATUS
func (assignment *SqlRoleAssignment_Spec) Initialize_From_SqlRoleAssignment_STATUS(source *SqlRoleAssignment_STATUS) error {

	// PrincipalId
	assignment.PrincipalId = genruntime.ClonePointerToString(source.PrincipalId)

	// RoleDefinitionId
	assignment.RoleDefinitionId = genruntime.ClonePointerToString(source.RoleDefinitionId)

	// Scope
	assignment.Scope = genruntime.ClonePointerToString(source.Scope)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (assignment *SqlRoleAssignment_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (assignment *SqlRoleAssignment_Spec) SetAzureName(azureName string) {
	assignment.AzureName = azureName
}

type SqlRoleAssignment_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: The unique resource identifier of the database account.
	Id *string `json:"id,omitempty"`

	// Name: The name of the database account.
	Name *string `json:"name,omitempty"`

	// PrincipalId: The unique identifier for the associated AAD principal in the AAD graph to which access is being granted
	// through this Role Assignment. Tenant ID for the principal is inferred using the tenant associated with the subscription.
	PrincipalId *string `json:"principalId,omitempty"`

	// RoleDefinitionId: The unique identifier for the associated Role Definition.
	RoleDefinitionId *string `json:"roleDefinitionId,omitempty"`

	// Scope: The data plane resource path for which access is being granted through this Role Assignment.
	Scope *string `json:"scope,omitempty"`

	// Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlRoleAssignment_STATUS{}

// ConvertStatusFrom populates our SqlRoleAssignment_STATUS from the provided source
func (assignment *SqlRoleAssignment_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.SqlRoleAssignment_STATUS)
	if ok {
		// Populate our instance from source
		return assignment.AssignProperties_From_SqlRoleAssignment_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.SqlRoleAssignment_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = assignment.AssignProperties_From_SqlRoleAssignment_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SqlRoleAssignment_STATUS
func (assignment *SqlRoleAssignment_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.SqlRoleAssignment_STATUS)
	if ok {
		// Populate destination from our instance
		return assignment.AssignProperties_To_SqlRoleAssignment_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.SqlRoleAssignment_STATUS{}
	err := assignment.AssignProperties_To_SqlRoleAssignment_STATUS(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &SqlRoleAssignment_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (assignment *SqlRoleAssignment_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.SqlRoleAssignment_STATUS{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (assignment *SqlRoleAssignment_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.SqlRoleAssignment_STATUS)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.SqlRoleAssignment_STATUS, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		assignment.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		assignment.Name = &name
	}

	// Set property "PrincipalId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.PrincipalId != nil {
			principalId := *typedInput.Properties.PrincipalId
			assignment.PrincipalId = &principalId
		}
	}

	// Set property "RoleDefinitionId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.RoleDefinitionId != nil {
			roleDefinitionId := *typedInput.Properties.RoleDefinitionId
			assignment.RoleDefinitionId = &roleDefinitionId
		}
	}

	// Set property "Scope":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Scope != nil {
			scope := *typedInput.Properties.Scope
			assignment.Scope = &scope
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		assignment.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_SqlRoleAssignment_STATUS populates our SqlRoleAssignment_STATUS from the provided source SqlRoleAssignment_STATUS
func (assignment *SqlRoleAssignment_STATUS) AssignProperties_From_SqlRoleAssignment_STATUS(source *storage.SqlRoleAssignment_STATUS) error {

	// Conditions
	assignment.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	assignment.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	assignment.Name = genruntime.ClonePointerToString(source.Name)

	// PrincipalId
	assignment.PrincipalId = genruntime.ClonePointerToString(source.PrincipalId)

	// RoleDefinitionId
	assignment.RoleDefinitionId = genruntime.ClonePointerToString(source.RoleDefinitionId)

	// Scope
	assignment.Scope = genruntime.ClonePointerToString(source.Scope)

	// Type
	assignment.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_SqlRoleAssignment_STATUS populates the provided destination SqlRoleAssignment_STATUS from our SqlRoleAssignment_STATUS
func (assignment *SqlRoleAssignment_STATUS) AssignProperties_To_SqlRoleAssignment_STATUS(destination *storage.SqlRoleAssignment_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(assignment.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(assignment.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(assignment.Name)

	// PrincipalId
	destination.PrincipalId = genruntime.ClonePointerToString(assignment.PrincipalId)

	// RoleDefinitionId
	destination.RoleDefinitionId = genruntime.ClonePointerToString(assignment.RoleDefinitionId)

	// Scope
	destination.Scope = genruntime.ClonePointerToString(assignment.Scope)

	// Type
	destination.Type = genruntime.ClonePointerToString(assignment.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type SqlRoleAssignmentOperatorSpec struct {
	// ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`

	// SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).
	SecretExpressions []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_SqlRoleAssignmentOperatorSpec populates our SqlRoleAssignmentOperatorSpec from the provided source SqlRoleAssignmentOperatorSpec
func (operator *SqlRoleAssignmentOperatorSpec) AssignProperties_From_SqlRoleAssignmentOperatorSpec(source *storage.SqlRoleAssignmentOperatorSpec) error {

	// ConfigMapExpressions
	if source.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(source.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range source.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		operator.ConfigMapExpressions = configMapExpressionList
	} else {
		operator.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if source.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(source.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range source.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		operator.SecretExpressions = secretExpressionList
	} else {
		operator.SecretExpressions = nil
	}

	// No error
	return nil
}

// AssignProperties_To_SqlRoleAssignmentOperatorSpec populates the provided destination SqlRoleAssignmentOperatorSpec from our SqlRoleAssignmentOperatorSpec
func (operator *SqlRoleAssignmentOperatorSpec) AssignProperties_To_SqlRoleAssignmentOperatorSpec(destination *storage.SqlRoleAssignmentOperatorSpec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// ConfigMapExpressions
	if operator.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(operator.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range operator.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		destination.ConfigMapExpressions = configMapExpressionList
	} else {
		destination.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if operator.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(operator.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range operator.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		destination.SecretExpressions = secretExpressionList
	} else {
		destination.SecretExpressions = nil
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

func init() {
	SchemeBuilder.Register(&SqlRoleAssignment{}, &SqlRoleAssignmentList{})
}

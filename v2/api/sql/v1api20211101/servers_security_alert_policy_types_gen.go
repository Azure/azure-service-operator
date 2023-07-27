// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"fmt"
	v1api20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101storage"
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
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/ServerSecurityAlertPolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/securityAlertPolicies/Default
type ServersSecurityAlertPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_SecurityAlertPolicy_Spec   `json:"spec,omitempty"`
	Status            Servers_SecurityAlertPolicy_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersSecurityAlertPolicy{}

// GetConditions returns the conditions of the resource
func (policy *ServersSecurityAlertPolicy) GetConditions() conditions.Conditions {
	return policy.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (policy *ServersSecurityAlertPolicy) SetConditions(conditions conditions.Conditions) {
	policy.Status.Conditions = conditions
}

var _ conversion.Convertible = &ServersSecurityAlertPolicy{}

// ConvertFrom populates our ServersSecurityAlertPolicy from the provided hub ServersSecurityAlertPolicy
func (policy *ServersSecurityAlertPolicy) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1api20211101s.ServersSecurityAlertPolicy)
	if !ok {
		return fmt.Errorf("expected sql/v1api20211101storage/ServersSecurityAlertPolicy but received %T instead", hub)
	}

	return policy.AssignProperties_From_ServersSecurityAlertPolicy(source)
}

// ConvertTo populates the provided hub ServersSecurityAlertPolicy from our ServersSecurityAlertPolicy
func (policy *ServersSecurityAlertPolicy) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1api20211101s.ServersSecurityAlertPolicy)
	if !ok {
		return fmt.Errorf("expected sql/v1api20211101storage/ServersSecurityAlertPolicy but received %T instead", hub)
	}

	return policy.AssignProperties_To_ServersSecurityAlertPolicy(destination)
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1api20211101-serverssecurityalertpolicy,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serverssecurityalertpolicies,verbs=create;update,versions=v1api20211101,name=default.v1api20211101.serverssecurityalertpolicies.sql.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &ServersSecurityAlertPolicy{}

// Default applies defaults to the ServersSecurityAlertPolicy resource
func (policy *ServersSecurityAlertPolicy) Default() {
	policy.defaultImpl()
	var temp any = policy
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the ServersSecurityAlertPolicy resource
func (policy *ServersSecurityAlertPolicy) defaultImpl() {}

var _ genruntime.ImportableResource = &ServersSecurityAlertPolicy{}

// InitializeSpec initializes the spec for this resource from the given status
func (policy *ServersSecurityAlertPolicy) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Servers_SecurityAlertPolicy_STATUS); ok {
		return policy.Spec.Initialize_From_Servers_SecurityAlertPolicy_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Servers_SecurityAlertPolicy_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &ServersSecurityAlertPolicy{}

// AzureName returns the Azure name of the resource (always "Default")
func (policy *ServersSecurityAlertPolicy) AzureName() string {
	return "Default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (policy ServersSecurityAlertPolicy) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (policy *ServersSecurityAlertPolicy) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (policy *ServersSecurityAlertPolicy) GetSpec() genruntime.ConvertibleSpec {
	return &policy.Spec
}

// GetStatus returns the status of this resource
func (policy *ServersSecurityAlertPolicy) GetStatus() genruntime.ConvertibleStatus {
	return &policy.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/securityAlertPolicies"
func (policy *ServersSecurityAlertPolicy) GetType() string {
	return "Microsoft.Sql/servers/securityAlertPolicies"
}

// NewEmptyStatus returns a new empty (blank) status
func (policy *ServersSecurityAlertPolicy) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Servers_SecurityAlertPolicy_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (policy *ServersSecurityAlertPolicy) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(policy.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  policy.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (policy *ServersSecurityAlertPolicy) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Servers_SecurityAlertPolicy_STATUS); ok {
		policy.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_SecurityAlertPolicy_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	policy.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1api20211101-serverssecurityalertpolicy,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serverssecurityalertpolicies,verbs=create;update,versions=v1api20211101,name=validate.v1api20211101.serverssecurityalertpolicies.sql.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &ServersSecurityAlertPolicy{}

// ValidateCreate validates the creation of the resource
func (policy *ServersSecurityAlertPolicy) ValidateCreate() (admission.Warnings, error) {
	validations := policy.createValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (policy *ServersSecurityAlertPolicy) ValidateDelete() (admission.Warnings, error) {
	validations := policy.deleteValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (policy *ServersSecurityAlertPolicy) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := policy.updateValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (policy *ServersSecurityAlertPolicy) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){policy.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (policy *ServersSecurityAlertPolicy) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (policy *ServersSecurityAlertPolicy) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return policy.validateResourceReferences()
		},
		policy.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (policy *ServersSecurityAlertPolicy) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&policy.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (policy *ServersSecurityAlertPolicy) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*ServersSecurityAlertPolicy)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, policy)
}

// AssignProperties_From_ServersSecurityAlertPolicy populates our ServersSecurityAlertPolicy from the provided source ServersSecurityAlertPolicy
func (policy *ServersSecurityAlertPolicy) AssignProperties_From_ServersSecurityAlertPolicy(source *v1api20211101s.ServersSecurityAlertPolicy) error {

	// ObjectMeta
	policy.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Servers_SecurityAlertPolicy_Spec
	err := spec.AssignProperties_From_Servers_SecurityAlertPolicy_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_SecurityAlertPolicy_Spec() to populate field Spec")
	}
	policy.Spec = spec

	// Status
	var status Servers_SecurityAlertPolicy_STATUS
	err = status.AssignProperties_From_Servers_SecurityAlertPolicy_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_SecurityAlertPolicy_STATUS() to populate field Status")
	}
	policy.Status = status

	// No error
	return nil
}

// AssignProperties_To_ServersSecurityAlertPolicy populates the provided destination ServersSecurityAlertPolicy from our ServersSecurityAlertPolicy
func (policy *ServersSecurityAlertPolicy) AssignProperties_To_ServersSecurityAlertPolicy(destination *v1api20211101s.ServersSecurityAlertPolicy) error {

	// ObjectMeta
	destination.ObjectMeta = *policy.ObjectMeta.DeepCopy()

	// Spec
	var spec v1api20211101s.Servers_SecurityAlertPolicy_Spec
	err := policy.Spec.AssignProperties_To_Servers_SecurityAlertPolicy_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_SecurityAlertPolicy_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1api20211101s.Servers_SecurityAlertPolicy_STATUS
	err = policy.Status.AssignProperties_To_Servers_SecurityAlertPolicy_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_SecurityAlertPolicy_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (policy *ServersSecurityAlertPolicy) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: policy.Spec.OriginalVersion(),
		Kind:    "ServersSecurityAlertPolicy",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/ServerSecurityAlertPolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/securityAlertPolicies/Default
type ServersSecurityAlertPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersSecurityAlertPolicy `json:"items"`
}

type Servers_SecurityAlertPolicy_Spec struct {
	// DisabledAlerts: Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection,
	// Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action, Brute_Force
	DisabledAlerts []string `json:"disabledAlerts,omitempty"`

	// EmailAccountAdmins: Specifies that the alert is sent to the account administrators.
	EmailAccountAdmins *bool `json:"emailAccountAdmins,omitempty"`

	// EmailAddresses: Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses []string `json:"emailAddresses,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/Server resource
	Owner *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"Server"`

	// RetentionDays: Specifies the number of days to keep in the Threat Detection audit logs.
	RetentionDays *int `json:"retentionDays,omitempty"`

	// +kubebuilder:validation:Required
	// State: Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the
	// specific database.
	State *ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State `json:"state,omitempty"`

	// StorageAccountAccessKey: Specifies the identifier key of the Threat Detection audit storage account.
	StorageAccountAccessKey *genruntime.SecretReference `json:"storageAccountAccessKey,omitempty"`

	// StorageEndpoint: Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage
	// will hold all Threat Detection audit logs.
	StorageEndpoint *string `json:"storageEndpoint,omitempty"`
}

var _ genruntime.ARMTransformer = &Servers_SecurityAlertPolicy_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (policy *Servers_SecurityAlertPolicy_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if policy == nil {
		return nil, nil
	}
	result := &Servers_SecurityAlertPolicy_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if policy.DisabledAlerts != nil ||
		policy.EmailAccountAdmins != nil ||
		policy.EmailAddresses != nil ||
		policy.RetentionDays != nil ||
		policy.State != nil ||
		policy.StorageAccountAccessKey != nil ||
		policy.StorageEndpoint != nil {
		result.Properties = &ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_ARM{}
	}
	for _, item := range policy.DisabledAlerts {
		result.Properties.DisabledAlerts = append(result.Properties.DisabledAlerts, item)
	}
	if policy.EmailAccountAdmins != nil {
		emailAccountAdmins := *policy.EmailAccountAdmins
		result.Properties.EmailAccountAdmins = &emailAccountAdmins
	}
	for _, item := range policy.EmailAddresses {
		result.Properties.EmailAddresses = append(result.Properties.EmailAddresses, item)
	}
	if policy.RetentionDays != nil {
		retentionDays := *policy.RetentionDays
		result.Properties.RetentionDays = &retentionDays
	}
	if policy.State != nil {
		state := *policy.State
		result.Properties.State = &state
	}
	if policy.StorageAccountAccessKey != nil {
		storageAccountAccessKeySecret, err := resolved.ResolvedSecrets.Lookup(*policy.StorageAccountAccessKey)
		if err != nil {
			return nil, errors.Wrap(err, "looking up secret for property StorageAccountAccessKey")
		}
		storageAccountAccessKey := storageAccountAccessKeySecret
		result.Properties.StorageAccountAccessKey = &storageAccountAccessKey
	}
	if policy.StorageEndpoint != nil {
		storageEndpoint := *policy.StorageEndpoint
		result.Properties.StorageEndpoint = &storageEndpoint
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (policy *Servers_SecurityAlertPolicy_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_SecurityAlertPolicy_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (policy *Servers_SecurityAlertPolicy_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_SecurityAlertPolicy_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_SecurityAlertPolicy_Spec_ARM, got %T", armInput)
	}

	// Set property "DisabledAlerts":
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.DisabledAlerts {
			policy.DisabledAlerts = append(policy.DisabledAlerts, item)
		}
	}

	// Set property "EmailAccountAdmins":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.EmailAccountAdmins != nil {
			emailAccountAdmins := *typedInput.Properties.EmailAccountAdmins
			policy.EmailAccountAdmins = &emailAccountAdmins
		}
	}

	// Set property "EmailAddresses":
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.EmailAddresses {
			policy.EmailAddresses = append(policy.EmailAddresses, item)
		}
	}

	// Set property "Owner":
	policy.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// Set property "RetentionDays":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.RetentionDays != nil {
			retentionDays := *typedInput.Properties.RetentionDays
			policy.RetentionDays = &retentionDays
		}
	}

	// Set property "State":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.State != nil {
			state := *typedInput.Properties.State
			policy.State = &state
		}
	}

	// no assignment for property "StorageAccountAccessKey"

	// Set property "StorageEndpoint":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StorageEndpoint != nil {
			storageEndpoint := *typedInput.Properties.StorageEndpoint
			policy.StorageEndpoint = &storageEndpoint
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Servers_SecurityAlertPolicy_Spec{}

// ConvertSpecFrom populates our Servers_SecurityAlertPolicy_Spec from the provided source
func (policy *Servers_SecurityAlertPolicy_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1api20211101s.Servers_SecurityAlertPolicy_Spec)
	if ok {
		// Populate our instance from source
		return policy.AssignProperties_From_Servers_SecurityAlertPolicy_Spec(src)
	}

	// Convert to an intermediate form
	src = &v1api20211101s.Servers_SecurityAlertPolicy_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_Servers_SecurityAlertPolicy_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Servers_SecurityAlertPolicy_Spec
func (policy *Servers_SecurityAlertPolicy_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1api20211101s.Servers_SecurityAlertPolicy_Spec)
	if ok {
		// Populate destination from our instance
		return policy.AssignProperties_To_Servers_SecurityAlertPolicy_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20211101s.Servers_SecurityAlertPolicy_Spec{}
	err := policy.AssignProperties_To_Servers_SecurityAlertPolicy_Spec(dst)
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

// AssignProperties_From_Servers_SecurityAlertPolicy_Spec populates our Servers_SecurityAlertPolicy_Spec from the provided source Servers_SecurityAlertPolicy_Spec
func (policy *Servers_SecurityAlertPolicy_Spec) AssignProperties_From_Servers_SecurityAlertPolicy_Spec(source *v1api20211101s.Servers_SecurityAlertPolicy_Spec) error {

	// DisabledAlerts
	policy.DisabledAlerts = genruntime.CloneSliceOfString(source.DisabledAlerts)

	// EmailAccountAdmins
	if source.EmailAccountAdmins != nil {
		emailAccountAdmin := *source.EmailAccountAdmins
		policy.EmailAccountAdmins = &emailAccountAdmin
	} else {
		policy.EmailAccountAdmins = nil
	}

	// EmailAddresses
	policy.EmailAddresses = genruntime.CloneSliceOfString(source.EmailAddresses)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		policy.Owner = &owner
	} else {
		policy.Owner = nil
	}

	// RetentionDays
	policy.RetentionDays = genruntime.ClonePointerToInt(source.RetentionDays)

	// State
	if source.State != nil {
		state := ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State(*source.State)
		policy.State = &state
	} else {
		policy.State = nil
	}

	// StorageAccountAccessKey
	if source.StorageAccountAccessKey != nil {
		storageAccountAccessKey := source.StorageAccountAccessKey.Copy()
		policy.StorageAccountAccessKey = &storageAccountAccessKey
	} else {
		policy.StorageAccountAccessKey = nil
	}

	// StorageEndpoint
	policy.StorageEndpoint = genruntime.ClonePointerToString(source.StorageEndpoint)

	// No error
	return nil
}

// AssignProperties_To_Servers_SecurityAlertPolicy_Spec populates the provided destination Servers_SecurityAlertPolicy_Spec from our Servers_SecurityAlertPolicy_Spec
func (policy *Servers_SecurityAlertPolicy_Spec) AssignProperties_To_Servers_SecurityAlertPolicy_Spec(destination *v1api20211101s.Servers_SecurityAlertPolicy_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// DisabledAlerts
	destination.DisabledAlerts = genruntime.CloneSliceOfString(policy.DisabledAlerts)

	// EmailAccountAdmins
	if policy.EmailAccountAdmins != nil {
		emailAccountAdmin := *policy.EmailAccountAdmins
		destination.EmailAccountAdmins = &emailAccountAdmin
	} else {
		destination.EmailAccountAdmins = nil
	}

	// EmailAddresses
	destination.EmailAddresses = genruntime.CloneSliceOfString(policy.EmailAddresses)

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

	// State
	if policy.State != nil {
		state := string(*policy.State)
		destination.State = &state
	} else {
		destination.State = nil
	}

	// StorageAccountAccessKey
	if policy.StorageAccountAccessKey != nil {
		storageAccountAccessKey := policy.StorageAccountAccessKey.Copy()
		destination.StorageAccountAccessKey = &storageAccountAccessKey
	} else {
		destination.StorageAccountAccessKey = nil
	}

	// StorageEndpoint
	destination.StorageEndpoint = genruntime.ClonePointerToString(policy.StorageEndpoint)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_Servers_SecurityAlertPolicy_STATUS populates our Servers_SecurityAlertPolicy_Spec from the provided source Servers_SecurityAlertPolicy_STATUS
func (policy *Servers_SecurityAlertPolicy_Spec) Initialize_From_Servers_SecurityAlertPolicy_STATUS(source *Servers_SecurityAlertPolicy_STATUS) error {

	// DisabledAlerts
	policy.DisabledAlerts = genruntime.CloneSliceOfString(source.DisabledAlerts)

	// EmailAccountAdmins
	if source.EmailAccountAdmins != nil {
		emailAccountAdmin := *source.EmailAccountAdmins
		policy.EmailAccountAdmins = &emailAccountAdmin
	} else {
		policy.EmailAccountAdmins = nil
	}

	// EmailAddresses
	policy.EmailAddresses = genruntime.CloneSliceOfString(source.EmailAddresses)

	// RetentionDays
	policy.RetentionDays = genruntime.ClonePointerToInt(source.RetentionDays)

	// State
	if source.State != nil {
		state := ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State(*source.State)
		policy.State = &state
	} else {
		policy.State = nil
	}

	// StorageEndpoint
	policy.StorageEndpoint = genruntime.ClonePointerToString(source.StorageEndpoint)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (policy *Servers_SecurityAlertPolicy_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type Servers_SecurityAlertPolicy_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// CreationTime: Specifies the UTC creation time of the policy.
	CreationTime *string `json:"creationTime,omitempty"`

	// DisabledAlerts: Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection,
	// Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action, Brute_Force
	DisabledAlerts []string `json:"disabledAlerts,omitempty"`

	// EmailAccountAdmins: Specifies that the alert is sent to the account administrators.
	EmailAccountAdmins *bool `json:"emailAccountAdmins,omitempty"`

	// EmailAddresses: Specifies an array of e-mail addresses to which the alert is sent.
	EmailAddresses []string `json:"emailAddresses,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// RetentionDays: Specifies the number of days to keep in the Threat Detection audit logs.
	RetentionDays *int `json:"retentionDays,omitempty"`

	// State: Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the
	// specific database.
	State *ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_STATUS `json:"state,omitempty"`

	// StorageEndpoint: Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage
	// will hold all Threat Detection audit logs.
	StorageEndpoint *string `json:"storageEndpoint,omitempty"`

	// SystemData: SystemData of SecurityAlertPolicyResource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_SecurityAlertPolicy_STATUS{}

// ConvertStatusFrom populates our Servers_SecurityAlertPolicy_STATUS from the provided source
func (policy *Servers_SecurityAlertPolicy_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1api20211101s.Servers_SecurityAlertPolicy_STATUS)
	if ok {
		// Populate our instance from source
		return policy.AssignProperties_From_Servers_SecurityAlertPolicy_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v1api20211101s.Servers_SecurityAlertPolicy_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_Servers_SecurityAlertPolicy_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Servers_SecurityAlertPolicy_STATUS
func (policy *Servers_SecurityAlertPolicy_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1api20211101s.Servers_SecurityAlertPolicy_STATUS)
	if ok {
		// Populate destination from our instance
		return policy.AssignProperties_To_Servers_SecurityAlertPolicy_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20211101s.Servers_SecurityAlertPolicy_STATUS{}
	err := policy.AssignProperties_To_Servers_SecurityAlertPolicy_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Servers_SecurityAlertPolicy_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (policy *Servers_SecurityAlertPolicy_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_SecurityAlertPolicy_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (policy *Servers_SecurityAlertPolicy_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_SecurityAlertPolicy_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_SecurityAlertPolicy_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "CreationTime":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.CreationTime != nil {
			creationTime := *typedInput.Properties.CreationTime
			policy.CreationTime = &creationTime
		}
	}

	// Set property "DisabledAlerts":
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.DisabledAlerts {
			policy.DisabledAlerts = append(policy.DisabledAlerts, item)
		}
	}

	// Set property "EmailAccountAdmins":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.EmailAccountAdmins != nil {
			emailAccountAdmins := *typedInput.Properties.EmailAccountAdmins
			policy.EmailAccountAdmins = &emailAccountAdmins
		}
	}

	// Set property "EmailAddresses":
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.EmailAddresses {
			policy.EmailAddresses = append(policy.EmailAddresses, item)
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

	// Set property "State":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.State != nil {
			state := *typedInput.Properties.State
			policy.State = &state
		}
	}

	// Set property "StorageEndpoint":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StorageEndpoint != nil {
			storageEndpoint := *typedInput.Properties.StorageEndpoint
			policy.StorageEndpoint = &storageEndpoint
		}
	}

	// Set property "SystemData":
	if typedInput.SystemData != nil {
		var systemData1 SystemData_STATUS
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		policy.SystemData = &systemData
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		policy.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Servers_SecurityAlertPolicy_STATUS populates our Servers_SecurityAlertPolicy_STATUS from the provided source Servers_SecurityAlertPolicy_STATUS
func (policy *Servers_SecurityAlertPolicy_STATUS) AssignProperties_From_Servers_SecurityAlertPolicy_STATUS(source *v1api20211101s.Servers_SecurityAlertPolicy_STATUS) error {

	// Conditions
	policy.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// CreationTime
	policy.CreationTime = genruntime.ClonePointerToString(source.CreationTime)

	// DisabledAlerts
	policy.DisabledAlerts = genruntime.CloneSliceOfString(source.DisabledAlerts)

	// EmailAccountAdmins
	if source.EmailAccountAdmins != nil {
		emailAccountAdmin := *source.EmailAccountAdmins
		policy.EmailAccountAdmins = &emailAccountAdmin
	} else {
		policy.EmailAccountAdmins = nil
	}

	// EmailAddresses
	policy.EmailAddresses = genruntime.CloneSliceOfString(source.EmailAddresses)

	// Id
	policy.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	policy.Name = genruntime.ClonePointerToString(source.Name)

	// RetentionDays
	policy.RetentionDays = genruntime.ClonePointerToInt(source.RetentionDays)

	// State
	if source.State != nil {
		state := ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_STATUS(*source.State)
		policy.State = &state
	} else {
		policy.State = nil
	}

	// StorageEndpoint
	policy.StorageEndpoint = genruntime.ClonePointerToString(source.StorageEndpoint)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		policy.SystemData = &systemDatum
	} else {
		policy.SystemData = nil
	}

	// Type
	policy.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_Servers_SecurityAlertPolicy_STATUS populates the provided destination Servers_SecurityAlertPolicy_STATUS from our Servers_SecurityAlertPolicy_STATUS
func (policy *Servers_SecurityAlertPolicy_STATUS) AssignProperties_To_Servers_SecurityAlertPolicy_STATUS(destination *v1api20211101s.Servers_SecurityAlertPolicy_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(policy.Conditions)

	// CreationTime
	destination.CreationTime = genruntime.ClonePointerToString(policy.CreationTime)

	// DisabledAlerts
	destination.DisabledAlerts = genruntime.CloneSliceOfString(policy.DisabledAlerts)

	// EmailAccountAdmins
	if policy.EmailAccountAdmins != nil {
		emailAccountAdmin := *policy.EmailAccountAdmins
		destination.EmailAccountAdmins = &emailAccountAdmin
	} else {
		destination.EmailAccountAdmins = nil
	}

	// EmailAddresses
	destination.EmailAddresses = genruntime.CloneSliceOfString(policy.EmailAddresses)

	// Id
	destination.Id = genruntime.ClonePointerToString(policy.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(policy.Name)

	// RetentionDays
	destination.RetentionDays = genruntime.ClonePointerToInt(policy.RetentionDays)

	// State
	if policy.State != nil {
		state := string(*policy.State)
		destination.State = &state
	} else {
		destination.State = nil
	}

	// StorageEndpoint
	destination.StorageEndpoint = genruntime.ClonePointerToString(policy.StorageEndpoint)

	// SystemData
	if policy.SystemData != nil {
		var systemDatum v1api20211101s.SystemData_STATUS
		err := policy.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

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

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State string

const (
	ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_Disabled = ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State("Disabled")
	ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_Enabled  = ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State("Enabled")
)

type ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_STATUS string

const (
	ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_STATUS_Disabled = ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_STATUS("Disabled")
	ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_STATUS_Enabled  = ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_STATUS("Enabled")
)

func init() {
	SchemeBuilder.Register(&ServersSecurityAlertPolicy{}, &ServersSecurityAlertPolicyList{})
}

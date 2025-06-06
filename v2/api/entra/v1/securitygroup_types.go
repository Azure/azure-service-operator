// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// +kubebuilder:rbac:groups=entra.azure.com,resources=securitygroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=entra.azure.com,resources={securitygroups/status,users/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// +kubebuilder:storageversion
// SecurityGroup is an Entra Security Group.
type SecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupSpec   `json:"spec,omitempty"`
	Status            SecurityGroupStatus `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SecurityGroup{}

// GetConditions returns the conditions of the resource
func (group *SecurityGroup) GetConditions() conditions.Conditions {
	return group.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (group *SecurityGroup) SetConditions(conditions conditions.Conditions) {
	group.Status.Conditions = conditions
}

// +kubebuilder:webhook:path=/mutate-entra-azure-com-v1-user,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=entra.azure.com,resources=users,verbs=create;update,versions=v1,name=default.v1.securitygroup.entra.azure.com,admissionReviewVersions=v1

/*
var _ admission.Defaulter = &SecurityGroup{}

// Default applies defaults to the FlexibleServer resource
func (group *SecurityGroup) Default() {
	group.defaultImpl()
	var temp interface{} = group
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}
*/

// // defaultAzureName defaults the Azure name of the resource to the Kubernetes name
// func (group *SecurityGroup) defaultAzureName() {
// 	if group.Spec.AzureName == "" {
// 		group.Spec.AzureName = group.Name
// 	}
// }

/*
// defaultImpl applies the code generated defaults to the FlexibleServer resource
func (group *SecurityGroup) defaultImpl() { group.defaultAzureName() }
*/

// +kubebuilder:webhook:path=/validate-entra-azure-com-v1-user,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=entra.azure.com,resources=users,verbs=create;update,versions=v1,name=validate.v1.securitygroup.entra.azure.com,admissionReviewVersions=v1

/*
var _ admission.Validator = &SecurityGroup{}

// ValidateCreate validates the creation of the resource
func (group *SecurityGroup) ValidateCreate() (admission.Warnings, error) {
	validations := group.createValidations()
	var temp interface{} = group
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (group *SecurityGroup) ValidateDelete() (admission.Warnings, error) {
	validations := group.deleteValidations()
	var temp interface{} = group
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (group *SecurityGroup) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := group.updateValidations()
	var temp interface{} = group
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}
*/

// // createValidations validates the creation of the resource
// func (group *SecurityGroup) createValidations() []func() (admission.Warnings, error) {
// 	return nil
// }

// // deleteValidations validates the deletion of the resource
// func (group *SecurityGroup) deleteValidations() []func() (admission.Warnings, error) {
// 	return nil
// }

// // updateValidations validates the update of the resource
// func (group *SecurityGroup) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
// 	return nil
// }

var _ conversion.Hub = &SecurityGroup{}

// Hub marks that this userSpec is the hub type for conversion
func (user *SecurityGroup) Hub() {}

// +kubebuilder:object:root=true
type SecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroup `json:"items"`
}

type SecurityGroupSpec struct {
	// DisplayName: The display name of the group.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName,omitempty"`

	// MailNickname: The email address of the group.
	// +kubebuilder:validation:Required
	MailNickname *string `json:"groupEmailAddress,omitempty"`

	// Description: The description of the group.
	Description *string `json:"description,omitempty"`

	// MembershipType: The membership type of the group.
	MembershipType *SecurityGroupMembershipType `json:"membershipType,omitempty"`

	// OperatorSpec: The operator specific configuration for the resource.
	OperatorSpec *SecurityGroupOperatorSpec `json:"operatorSpec,omitempty"`
}

// OriginalVersion returns the original API version used to create the resource.
func (spec *SecurityGroupSpec) OriginalVersion() string {
	return GroupVersion.Version
}

// AssignToGroup configures the provided instance with the details of the group
func (spec *SecurityGroupSpec) AssignToGroup(model models.Groupable) {
	model.SetSecurityEnabled(to.Ptr(true))
	model.SetDisplayName(spec.DisplayName)

	if spec.MailNickname != nil {
		model.SetMailNickname(spec.MailNickname)
	}

	if spec.Description != nil {
		model.SetDescription(spec.Description)
	}

	groupTypes := []string{"Unified"}
	if spec.MembershipType != nil && *spec.MembershipType == SecurityGroupMembershipTypeDynamic {
		groupTypes = append(groupTypes, "DynamicMembership")
	}

	model.SetGroupTypes(groupTypes)

	// This is a security group, not a mail distribution group
	model.SetMailEnabled(to.Ptr(false))
}

type SecurityGroupStatus struct {
	// EntraID: The GUID identifing the resource in Entra
	EntraID *string `json:"entraID,omitempty"`

	// DisplayName: The display name of the group.
	DisplayName *string `json:"displayName,omitempty"`

	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// +kubebuilder:validation:Required
	// MailNickname: The email address of the group.
	MailNickname *string `json:"groupEmailAddress,omitempty"`

	// Description: The description of the group.
	Description *string `json:"description,omitempty"`
}

func (status *SecurityGroupStatus) AssignFromGroup(model models.Groupable) {
	if id := model.GetId(); id != nil {
		status.EntraID = id
	}

	if name := model.GetDisplayName(); name != nil {
		status.DisplayName = name
	}

	if mailNickname := model.GetMailNickname(); mailNickname != nil {
		status.MailNickname = mailNickname
	}

	if description := model.GetDescription(); description != nil {
		status.Description = description
	}
}

// +kubebuilder:validation:Enum={"assigned","enabled"}
type SecurityGroupMembershipType string

const (
	SecurityGroupMembershipTypeAssigned SecurityGroupMembershipType = "assigned"
	SecurityGroupMembershipTypeDynamic  SecurityGroupMembershipType = "dynamic"
)

type SecurityGroupOperatorSpec struct {
	// CreationMode: Specifies how ASO will try to create the resource.
	// If not specified, defaults to "AdoptOrCreate".
	// +kubebuilder:default=AdoptOrCreate
	// +kubebuilder:validation:Enum=AdoptOrCreate;AlwaysCreate
	CreationMode *CreationMode `json:"creationMode,omitempty"`
}

// CreationAllowed checks if the creation mode allows ASO to create a new security group.
func (spec *SecurityGroupOperatorSpec) CreationAllowed() bool {
	if spec.CreationMode == nil {
		// Default is AdoptOrCreate
		return true
	}

	return spec.CreationMode.AllowsCreation()
}

// AllowsAdoption checks if the creation mode allows ASO to adopt an existing security group.
func (spec *SecurityGroupOperatorSpec) AdoptionAllowed() bool {
	if spec.CreationMode == nil {
		// Default is AdoptOrCreate
		return true
	}

	return spec.CreationMode.AllowsAdoption()
}

func init() {
	SchemeBuilder.Register(&SecurityGroup{}, &SecurityGroupList{})
}

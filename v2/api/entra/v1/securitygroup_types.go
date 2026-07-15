// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1

import (
	"fmt"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	asoentra "github.com/Azure/azure-service-operator/v2/api/entra"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/google/uuid"
)

// +kubebuilder:rbac:groups=entra.azure.com,resources=securitygroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=entra.azure.com,resources={securitygroups/status,users/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories={azure,entra,aad}
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

var _ conversion.Hub = &SecurityGroup{}

// Hub marks that this userSpec is the hub type for conversion
func (user *SecurityGroup) Hub() {}

// +kubebuilder:object:root=true
type SecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroup `json:"items"`
}

// +kubebuilder:validation:XValidation:rule="size(self.owners) + size(self.members) <= 20",message="the combined number of owners and members must not exceed 20"
type SecurityGroupSpec struct {
	// DisplayName: The display name of the group.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName,omitempty"`

	// MailNickname: The email address of the group, specified either as a mail nickname (`mygroup`)
	// or as a full email address (`mygroup@mydomain.com`).
	// +kubebuilder:validation:Required
	MailNickname *string `json:"mailNickname,omitempty"`

	// Description: The description of the group.
	Description *string `json:"description,omitempty"`

	// MembershipType: The membership type of the group.
	MembershipType *SecurityGroupMembershipType `json:"membershipType,omitempty"`

	// OperatorSpec: The operator specific configuration for the resource.
	OperatorSpec *SecurityGroupOperatorSpec `json:"operatorSpec,omitempty"`

	// IsAssignableToRole: Indicates whether the group can be assigned to a role.
	IsAssignableToRole *bool `json:"isAssignableToRole,omitempty"`

	// Owners: Directory objects (users, service principals, groups) to assign as owners of the security
	// group at creation time. Applied during the initial POST to Microsoft Graph via `owners@odata.bind`
	// and then used as desired owner state for later reconciliation.
	// Required when ASO authenticates with an app-only token and the calling principal lacks Group.ReadWrite.All —
	// otherwise the created group has no owners and is unmanageable.
	// +kubebuilder:validation:MaxItems=20
	// +kubebuilder:validation:XValidation:rule="size(self.filter(x, has(x.objectID)).map(x, x.objectID).distinct()) == size(self.filter(x, has(x.objectID)))",message="owners must be unique"
	// +kubebuilder:validation:XValidation:rule="size(self.filter(x, has(x.objectIDFromConfig)).map(x, x.objectIDFromConfig).distinct()) == size(self.filter(x, has(x.objectIDFromConfig)))",message="owners must be unique"
	Owners []SecurityGroupMemberReference `json:"owners,omitempty"`

	// Members: Directory objects (users, service principals, groups) to assign as members of the security
	// group at creation time. Applied during the initial POST to Microsoft Graph via `members@odata.bind`
	// and then used as desired member state for later reconciliation.
	// +kubebuilder:validation:MaxItems=20
	// +kubebuilder:validation:XValidation:rule="size(self.filter(x, has(x.objectID)).map(x, x.objectID).distinct()) == size(self.filter(x, has(x.objectID)))",message="members must be unique"
	// +kubebuilder:validation:XValidation:rule="size(self.filter(x, has(x.objectIDFromConfig)).map(x, x.objectIDFromConfig).distinct()) == size(self.filter(x, has(x.objectIDFromConfig)))",message="members must be unique"
	Members []SecurityGroupMemberReference `json:"members,omitempty"`
}

// SecurityGroupMemberReference is a reference to a directory object (user, service principal, or group) by its
// Entra Object ID.
// +kubebuilder:validation:XValidation:rule="has(self.objectID) != has(self.objectIDFromConfig)",message="exactly one of objectID or objectIDFromConfig must be specified"
type SecurityGroupMemberReference struct {
	// ObjectID: The Entra Object ID (GUID) of the directory object.
	// +kubebuilder:validation:Pattern="^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	ObjectID *string `json:"objectID,omitempty" optionalConfigMapPair:"ObjectID"`

	// ObjectIDFromConfig: Reference to a configmap value containing the Entra Object ID.
	ObjectIDFromConfig *genruntime.ConfigMapReference `json:"objectIDFromConfig,omitempty" optionalConfigMapPair:"ObjectID"`
}

// ResolveObjectID returns the Entra Object ID for this reference, looking it up via the resolved
// config maps when ObjectIDFromConfig is set. Returns an error if neither ObjectID nor
// ObjectIDFromConfig is populated, or if the configmap lookup fails.
func (ref *SecurityGroupMemberReference) ResolveObjectID(
	resolved genruntime.Resolved[genruntime.ConfigMapReference, string],
) (string, error) {
	switch {
	case ref.ObjectID != nil:
		return *ref.ObjectID, nil

	case ref.ObjectIDFromConfig != nil:
		val, err := resolved.Lookup(*ref.ObjectIDFromConfig)
		if err != nil {
			return "", eris.Wrapf(err, "failed resolving objectIDFromConfig")
		}

		// Check that the resolved value is a valid GUID by trying to parse it
		if _, err = uuid.Parse(val); err != nil {
			return "", eris.Wrapf(err, "invalid GUID resolved from objectIDFromConfig")
		}

		return val, nil

	default:
		return "", eris.New("missing objectID or objectIDFromConfig")
	}
}

// OriginalVersion returns the original API version used to create the resource.
func (spec *SecurityGroupSpec) OriginalVersion() string {
	return GroupVersion.Version
}

// AssignODataBindOnCreate sets the `owners@odata.bind` and `members@odata.bind` additional data
// on the group model. These annotations are only valid during the initial POST to Microsoft Graph;
// they must NOT be included in PATCH requests.
// The typed setters (SetOwners/SetMembers) serialize as nested objects which Graph rejects on create —
// the @odata.bind annotation is the only working shape for setting owners/members inline at creation time.
func (spec *SecurityGroupSpec) AssignODataBindOnCreate(
	model models.Groupable,
	resolved genruntime.Resolved[genruntime.ConfigMapReference, string],
) error {
	additionalData := model.GetAdditionalData()
	if additionalData == nil {
		additionalData = make(map[string]any)
	}

	if len(spec.Owners) > 0 {
		owners, err := graphDirectoryURIs(spec.Owners, "owners", resolved)
		if err != nil {
			return err
		}

		additionalData["owners@odata.bind"] = owners
	}

	if len(spec.Members) > 0 {
		members, err := graphDirectoryURIs(spec.Members, "members", resolved)
		if err != nil {
			return err
		}

		additionalData["members@odata.bind"] = members
	}

	model.SetAdditionalData(additionalData)
	return nil
}

// ResolveOwnerObjectIDs resolves and de-duplicates owner object IDs.
func (spec *SecurityGroupSpec) ResolveOwnerObjectIDs(resolved genruntime.Resolved[genruntime.ConfigMapReference, string]) ([]string, error) {
	return resolveMemberObjectIDs(spec.Owners, "owners", resolved)
}

// ResolveMemberObjectIDs resolves and de-duplicates member object IDs.
func (spec *SecurityGroupSpec) ResolveMemberObjectIDs(resolved genruntime.Resolved[genruntime.ConfigMapReference, string]) ([]string, error) {
	return resolveMemberObjectIDs(spec.Members, "members", resolved)
}

func resolveMemberObjectIDs(
	references []SecurityGroupMemberReference,
	field string,
	resolved genruntime.Resolved[genruntime.ConfigMapReference, string],
) ([]string, error) {
	seen := set.Make[string]()
	result := make([]string, 0, len(references))
	for i, ref := range references {
		id, err := ref.ResolveObjectID(resolved)
		if err != nil {
			return nil, eris.Wrapf(err, "%s[%d]", field, i)
		}

		if seen.Contains(id) {
			continue
		}

		seen.Add(id)
		result = append(result, id)
	}

	return result, nil
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

	// Set the membership type
	membershipType := SecurityGroupMembershipTypeAssigned
	if spec.MembershipType != nil {
		membershipType = *spec.MembershipType
	}

	var groupTypes []string
	switch membershipType {
	case SecurityGroupMembershipTypeAssigned:
	// Empty list means assigned membership
	case SecurityGroupMembershipTypeAssignedM365:
		groupTypes = []string{"Unified"}
	case SecurityGroupMembershipTypeDynamic:
		groupTypes = []string{"DynamicMembership"}
	case SecurityGroupMembershipTypeDynamicM365:
		groupTypes = []string{"Unified", "DynamicMembership"}
	}
	model.SetGroupTypes(groupTypes)

	// Set isAssignableToRole
	if spec.IsAssignableToRole != nil {
		model.SetIsAssignableToRole(spec.IsAssignableToRole)
	}

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

// +kubebuilder:validation:Enum={"assigned","enabled","assignedm365","enabledm365"}
// +kubebuilder:default=AdoptOrCreate
type SecurityGroupMembershipType string

const (
	// SecurityGroupMembershipTypeAssigned indicates that the group is a security group with assigned members.
	SecurityGroupMembershipTypeAssigned SecurityGroupMembershipType = "assigned"
	// SecurityGroupMembershipTypeDynamic indicates that the group is a security group with dynamic membership.
	SecurityGroupMembershipTypeDynamic SecurityGroupMembershipType = "dynamic"
	// SecurityGroupMembershipTypeAssigned indicates that the group is a Microsoft 365 security group with assigned members.
	SecurityGroupMembershipTypeAssignedM365 SecurityGroupMembershipType = "assignedm365"
	// SecurityGroupMembershipTypeDynamic indicates that the group is a Microsoft 365 security group with dynamic membership.
	SecurityGroupMembershipTypeDynamicM365 SecurityGroupMembershipType = "dynamicm365"
)

type SecurityGroupOperatorSpec struct {
	// CreationMode: Specifies how ASO will try to create the resource.
	// Specify "AlwaysCreate" to always create a new security group when first reconciled.
	// Or specify "AdoptOrCreate" to first try to adopt an existing security group with the same display name.
	// If multiple security groups with the same display name are found, the resource condition will show an error.
	// If not specified, defaults to "AdoptOrCreate".
	CreationMode *CreationMode `json:"creationMode,omitempty"`

	// ConfigMaps specifies any config maps that should be created by the operator.
	ConfigMaps *SecurityGroupOperatorConfigMaps `json:"configmaps,omitempty"`
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

type SecurityGroupOperatorConfigMaps struct {
	// EntraID: The Entra ID of the group.
	EntraID *genruntime.ConfigMapDestination `json:"entraID,omitempty"`
}

// graphDirectoryURIs resolves a slice of member references to Microsoft Graph directory object URIs.
// fieldName (e.g. "owners" or "members") is used to provide context in error messages.
func graphDirectoryURIs(
	references []SecurityGroupMemberReference,
	fieldName string,
	resolved genruntime.Resolved[genruntime.ConfigMapReference, string],
) ([]string, error) {
	uris := make([]string, 0, len(references))
	for i, ref := range references {
		id, err := ref.ResolveObjectID(resolved)
		if err != nil {
			return nil, fmt.Errorf("%s[%d]: %w", fieldName, i, err)
		}

		uris = append(uris, asoentra.DirectoryObjectRefURI(id))
	}

	return uris, nil
}

func init() {
	SchemeBuilder.Register(&SecurityGroup{}, &SecurityGroupList{})
}

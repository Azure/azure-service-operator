// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// +kubebuilder:rbac:groups=dbforpostgresql.azure.com,resources=users,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=dbforpostgresql.azure.com,resources={users/status,users/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories={azure,dbforpostgresql}
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// +kubebuilder:storageversion
// User is a postgresql user.
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:Required
	Spec   UserSpec   `json:"spec,omitempty"`
	Status UserStatus `json:"status,omitempty"`
}

var _ conditions.Conditioner = &User{}

// GetConditions returns the conditions of the resource
func (user *User) GetConditions() conditions.Conditions {
	return user.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (user *User) SetConditions(conditions conditions.Conditions) {
	user.Status.Conditions = conditions
}

var _ genruntime.ARMOwned = &User{}

// AzureName returns the Azure name of the resource
func (user *User) AzureName() string {
	return user.Spec.AzureName
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (user *User) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(user.Spec)
	return user.Spec.Owner.AsResourceReference(group, kind)
}

var _ conversion.Hub = &User{}

// Hub marks that this userSpec is the hub type for conversion
func (user *User) Hub() {}

// +kubebuilder:object:root=true
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// +kubebuilder:validation:XValidation:rule="has(self.localUser) != has(self.aadUser)",message="exactly one of spec.localUser or spec.aadUser must be set"
// +kubebuilder:validation:XValidation:rule="!has(oldSelf.localUser) || has(self.localUser)",message="cannot change from local user to AAD user"
// +kubebuilder:validation:XValidation:rule="!has(oldSelf.aadUser) || has(self.aadUser)",message="cannot change from AAD user to local user"
type UserSpec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbforpostgresql.azure.com/FlexibleServer resource
	Owner *genruntime.KubernetesOwnerReference `group:"dbforpostgresql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`

	// Roles is the set of roles granted to the user upon creation.
	// The Azure Database for PostgreSQL server is created with 3 default roles defined: azure_pg_admin, azure_superuser,
	// and your server admin user (this last is a role w/ login permission, commonly called a User).
	Roles []string `json:"roles,omitempty"`

	// +kubebuilder:default={login: true}
	// RoleOptions defines additional attributes of the user role. You can read more about these attributes
	// at https://www.postgresql.org/docs/current/role-attributes.html.
	RoleOptions *RoleOptionsSpec `json:"roleOptions,omitempty"`

	// DatabasePrivileges are the database-level privileges granted to the user, keyed by database name.
	// Privileges include the following: CONNECT, CREATE, TEMPORARY (or TEMP) and ALL (or ALL PRIVILEGES), which
	// is expanded to CONNECT, CREATE and TEMPORARY. Privileges must be upper case.
	// When set, the user ends up with exactly these privileges on each listed database, and privileges
	// granted directly to the user on any database not listed are revoked. When omitted, database
	// privileges are not managed. Databases owned by the user are never managed (an owner holds every
	// privilege on its database).
	// An empty map is kept (not dropped as omitted) and revokes every privilege granted directly to the user.
	// You can read more about these privileges at https://www.postgresql.org/docs/current/ddl-priv.html.
	DatabasePrivileges map[string][]DatabasePrivilege `json:"databasePrivileges,omitzero"`

	// LocalUser contains details for creating a standard (non-aad) postgresql User.
	// Exactly one of LocalUser and AADUser must be set.
	LocalUser *LocalUserSpec `json:"localUser,omitempty"`

	// AADUser contains details for creating a Microsoft Entra ID (AAD) user: a role bound to an Entra principal (a managed
	// identity, a service principal, a user or a group), which signs in with an Entra token instead of a password. The
	// operator signs in as the server's Entra administrator using its own credential.
	// Exactly one of LocalUser and AADUser must be set.
	AADUser *AADUserSpec `json:"aadUser,omitempty"`
}

// OriginalVersion returns the original API version used to create the resource.
func (userSpec *UserSpec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (userSpec *UserSpec) SetAzureName(azureName string) { userSpec.AzureName = azureName }

// DatabasePrivilege is a privilege that can be granted on a PostgreSQL database.
// +kubebuilder:validation:Enum={"CONNECT","CREATE","TEMPORARY","TEMP","ALL","ALL PRIVILEGES"}
type DatabasePrivilege string

type LocalUserSpec struct {
	// +kubebuilder:validation:Required
	// ServerAdminUsername is the user name of the Server administrator
	ServerAdminUsername string `json:"serverAdminUsername,omitempty"`

	// +kubebuilder:validation:Required
	// ServerAdminPassword is a reference to a secret containing the servers administrator password
	ServerAdminPassword *genruntime.SecretReference `json:"serverAdminPassword,omitempty"`

	// +kubebuilder:validation:Required
	// Password is the password to use for the user
	Password *genruntime.SecretReference `json:"password,omitempty"`
}

// +kubebuilder:validation:XValidation:rule="(has(self.objectId) && self.objectId != \"\") == (has(self.principalType) && self.principalType != \"\")",message="spec.aadUser.objectId and spec.aadUser.principalType must be set together"
// +kubebuilder:validation:XValidation:rule="(has(self.objectId) ? self.objectId : \"\") == (has(oldSelf.objectId) ? oldSelf.objectId : \"\")",message="cannot change AAD user objectId"
// +kubebuilder:validation:XValidation:rule="(has(self.principalType) ? self.principalType : \"\") == (has(oldSelf.principalType) ? oldSelf.principalType : \"\")",message="cannot change AAD user principalType"
type AADUserSpec struct {
	// +kubebuilder:validation:Required
	// ServerAdminUsername is the name of the server's Microsoft Entra administrator that the operator's identity signs in
	// as. If the administrator was configured with Azure Service Operator, this should match the Administrator's
	// $.spec.principalName field. If the administrator is a group, this is the group name, not the name of the identity
	// that is a member of the group.
	ServerAdminUsername string `json:"serverAdminUsername,omitempty"`

	// ObjectId is the Microsoft Entra object ID of the principal. When set, the role is created with
	// pgaadauth_create_principal_with_oid, so the AzureName may differ from the principal's display name, and
	// PrincipalType must be set too. When omitted, the role is created with pgaadauth_create_principal, which binds it to
	// the principal whose display name is the AzureName.
	ObjectId string `json:"objectId,omitempty"`

	// +kubebuilder:validation:Enum=user;group;service
	// PrincipalType is the kind of principal ObjectId names: user, group or service (a managed identity or service
	// principal). Required when ObjectId is set.
	PrincipalType string `json:"principalType,omitempty"`
}

type RoleOptionsSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=true
	// WITH LOGIN or NOLOGIN
	Login bool `json:"login,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	// WITH CREATEROLE or NOCREATEROLE
	CreateRole bool `json:"createRole,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	// WITH CREATEDB or NOCREATEDB
	CreateDb bool `json:"createDb,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	// WITH REPLICATION or NOREPLICATION
	Replication bool `json:"replication,omitempty"`
}

type UserStatus struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}

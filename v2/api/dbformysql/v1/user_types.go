// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// +kubebuilder:rbac:groups=dbformysql.azure.com,resources=users,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=dbformysql.azure.com,resources={users/status,users/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// +kubebuilder:storageversion
// User is a MySQL user
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserSpec   `json:"spec,omitempty"`
	Status            UserStatus `json:"status,omitempty"`
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

// +kubebuilder:webhook:path=/mutate-dbformysql-azure-com-v1-user,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformysql.azure.com,resources=users,verbs=create;update,versions=v1,name=default.v1.users.dbformysql.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &User{}

// Default applies defaults to the FlexibleServer resource
func (user *User) Default() {
	user.defaultImpl()
	var temp interface{} = user
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (user *User) defaultAzureName() {
	if user.Spec.AzureName == "" {
		user.Spec.AzureName = user.Name
	}
}

// defaultImpl applies the code generated defaults to the FlexibleServer resource
func (user *User) defaultImpl() { user.defaultAzureName() }

var _ genruntime.ARMOwned = &User{}

// AzureName returns the Azure name of the resource
func (user *User) AzureName() string {
	return user.Spec.AzureName
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (user *User) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(user.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  user.Spec.Owner.Name,
	}
}

// +kubebuilder:webhook:path=/validate-dbformysql-azure-com-v1-user,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformysql.azure.com,resources=users,verbs=create;update,versions=v1,name=validate.v1.users.dbformysql.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &User{}

// ValidateCreate validates the creation of the resource
func (user *User) ValidateCreate() error {
	validations := user.createValidations()
	var temp interface{} = user
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
func (user *User) ValidateDelete() error {
	validations := user.deleteValidations()
	var temp interface{} = user
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
func (user *User) ValidateUpdate(old runtime.Object) error {
	validations := user.updateValidations()
	var temp interface{} = user
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
func (user *User) createValidations() []func() error {
	return nil
}

// deleteValidations validates the deletion of the resource
func (user *User) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (user *User) updateValidations() []func(old runtime.Object) error {
	return nil
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

type UserSpec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	//Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	//controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	//reference to a dbformysql.azure.com/FlexibleServer resource
	Owner *genruntime.KnownResourceReference `group:"dbformysql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`

	// Hostname is the host the user will connect from. If omitted, the default is to allow connection from any hostname.
	Hostname string `json:"hostname,omitempty"`

	// The server-level roles assigned to the user.
	// Privileges include the following: RELOAD, PROCESS, SHOW
	// DATABASES, REPLICATION SLAVE, REPLICATION CLIENT, CREATE USER
	Privileges []string `json:"privileges,omitempty"`

	// The database-level roles assigned to the user (keyed by
	// database name). Privileges include the following: SELECT,
	// INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX,
	// ALTER, CREATE TEMPORARY TABLES, LOCK TABLES, EXECUTE, CREATE
	// VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, EVENT, TRIGGER
	DatabasePrivileges map[string][]string `json:"databasePrivileges,omitempty"`

	// TODO: Note this is required right now but will move to be optional in the future when we have AAD support
	// +kubebuilder:validation:Required
	// LocalUser contains details for creating a standard (non-aad) MySQL User
	LocalUser *LocalUserSpec `json:"localUser,omitempty"`
}

// OriginalVersion returns the original API version used to create the resource.
func (userSpec *UserSpec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (userSpec *UserSpec) SetAzureName(azureName string) { userSpec.AzureName = azureName }

//var _ genruntime.ConvertibleSpec = &UserSpec{}
//
//// ConvertSpecFrom populates our ConfigurationStore_Spec from the provided source
//func (userSpec *UserSpec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
//	if source == userSpec {
//		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
//	}
//
//	return source.ConvertSpecTo(userSpec)
//}
//
//// ConvertSpecTo populates the provided destination from our ConfigurationStore_Spec
//func (userSpec *UserSpec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
//	if destination == userSpec {
//		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
//	}
//
//	return destination.ConvertSpecFrom(userSpec)
//}

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

type UserStatus struct {
	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`
}

// TODO: Where are these (and the spec flavors) called? Or are these just placehodlers for if/when there's a newer version?
//var _ genruntime.ConvertibleStatus = &UserStatus{}
//
//// ConvertStatusFrom populates our ConfigurationStore_STATUS from the provided source
//func (userStatus *UserStatus) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
//	if source == userStatus {
//		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
//	}
//
//	return source.ConvertStatusTo(userStatus)
//}
//
//// ConvertStatusTo populates the provided destination from our ConfigurationStore_STATUS
//func (userStatus *UserStatus) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
//	if destination == userStatus {
//		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
//	}
//
//	return destination.ConvertStatusFrom(userStatus)
//}

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}

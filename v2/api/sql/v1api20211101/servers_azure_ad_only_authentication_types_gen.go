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
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/ServerAzureADOnlyAuthentications.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/azureADOnlyAuthentications/Default
type ServersAzureADOnlyAuthentication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_AzureADOnlyAuthentication_Spec   `json:"spec,omitempty"`
	Status            Servers_AzureADOnlyAuthentication_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersAzureADOnlyAuthentication{}

// GetConditions returns the conditions of the resource
func (authentication *ServersAzureADOnlyAuthentication) GetConditions() conditions.Conditions {
	return authentication.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (authentication *ServersAzureADOnlyAuthentication) SetConditions(conditions conditions.Conditions) {
	authentication.Status.Conditions = conditions
}

var _ conversion.Convertible = &ServersAzureADOnlyAuthentication{}

// ConvertFrom populates our ServersAzureADOnlyAuthentication from the provided hub ServersAzureADOnlyAuthentication
func (authentication *ServersAzureADOnlyAuthentication) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1api20211101s.ServersAzureADOnlyAuthentication)
	if !ok {
		return fmt.Errorf("expected sql/v1api20211101storage/ServersAzureADOnlyAuthentication but received %T instead", hub)
	}

	return authentication.AssignProperties_From_ServersAzureADOnlyAuthentication(source)
}

// ConvertTo populates the provided hub ServersAzureADOnlyAuthentication from our ServersAzureADOnlyAuthentication
func (authentication *ServersAzureADOnlyAuthentication) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1api20211101s.ServersAzureADOnlyAuthentication)
	if !ok {
		return fmt.Errorf("expected sql/v1api20211101storage/ServersAzureADOnlyAuthentication but received %T instead", hub)
	}

	return authentication.AssignProperties_To_ServersAzureADOnlyAuthentication(destination)
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1api20211101-serversazureadonlyauthentication,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversazureadonlyauthentications,verbs=create;update,versions=v1api20211101,name=default.v1api20211101.serversazureadonlyauthentications.sql.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &ServersAzureADOnlyAuthentication{}

// Default applies defaults to the ServersAzureADOnlyAuthentication resource
func (authentication *ServersAzureADOnlyAuthentication) Default() {
	authentication.defaultImpl()
	var temp any = authentication
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the ServersAzureADOnlyAuthentication resource
func (authentication *ServersAzureADOnlyAuthentication) defaultImpl() {}

var _ genruntime.ImportableResource = &ServersAzureADOnlyAuthentication{}

// InitializeSpec initializes the spec for this resource from the given status
func (authentication *ServersAzureADOnlyAuthentication) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Servers_AzureADOnlyAuthentication_STATUS); ok {
		return authentication.Spec.Initialize_From_Servers_AzureADOnlyAuthentication_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Servers_AzureADOnlyAuthentication_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &ServersAzureADOnlyAuthentication{}

// AzureName returns the Azure name of the resource (always "Default")
func (authentication *ServersAzureADOnlyAuthentication) AzureName() string {
	return "Default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (authentication ServersAzureADOnlyAuthentication) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (authentication *ServersAzureADOnlyAuthentication) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (authentication *ServersAzureADOnlyAuthentication) GetSpec() genruntime.ConvertibleSpec {
	return &authentication.Spec
}

// GetStatus returns the status of this resource
func (authentication *ServersAzureADOnlyAuthentication) GetStatus() genruntime.ConvertibleStatus {
	return &authentication.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/azureADOnlyAuthentications"
func (authentication *ServersAzureADOnlyAuthentication) GetType() string {
	return "Microsoft.Sql/servers/azureADOnlyAuthentications"
}

// NewEmptyStatus returns a new empty (blank) status
func (authentication *ServersAzureADOnlyAuthentication) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Servers_AzureADOnlyAuthentication_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (authentication *ServersAzureADOnlyAuthentication) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(authentication.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  authentication.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (authentication *ServersAzureADOnlyAuthentication) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Servers_AzureADOnlyAuthentication_STATUS); ok {
		authentication.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_AzureADOnlyAuthentication_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	authentication.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1api20211101-serversazureadonlyauthentication,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversazureadonlyauthentications,verbs=create;update,versions=v1api20211101,name=validate.v1api20211101.serversazureadonlyauthentications.sql.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &ServersAzureADOnlyAuthentication{}

// ValidateCreate validates the creation of the resource
func (authentication *ServersAzureADOnlyAuthentication) ValidateCreate() error {
	validations := authentication.createValidations()
	var temp any = authentication
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
func (authentication *ServersAzureADOnlyAuthentication) ValidateDelete() error {
	validations := authentication.deleteValidations()
	var temp any = authentication
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
func (authentication *ServersAzureADOnlyAuthentication) ValidateUpdate(old runtime.Object) error {
	validations := authentication.updateValidations()
	var temp any = authentication
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
func (authentication *ServersAzureADOnlyAuthentication) createValidations() []func() error {
	return []func() error{authentication.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (authentication *ServersAzureADOnlyAuthentication) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (authentication *ServersAzureADOnlyAuthentication) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return authentication.validateResourceReferences()
		},
		authentication.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (authentication *ServersAzureADOnlyAuthentication) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&authentication.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (authentication *ServersAzureADOnlyAuthentication) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*ServersAzureADOnlyAuthentication)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, authentication)
}

// AssignProperties_From_ServersAzureADOnlyAuthentication populates our ServersAzureADOnlyAuthentication from the provided source ServersAzureADOnlyAuthentication
func (authentication *ServersAzureADOnlyAuthentication) AssignProperties_From_ServersAzureADOnlyAuthentication(source *v1api20211101s.ServersAzureADOnlyAuthentication) error {

	// ObjectMeta
	authentication.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Servers_AzureADOnlyAuthentication_Spec
	err := spec.AssignProperties_From_Servers_AzureADOnlyAuthentication_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_AzureADOnlyAuthentication_Spec() to populate field Spec")
	}
	authentication.Spec = spec

	// Status
	var status Servers_AzureADOnlyAuthentication_STATUS
	err = status.AssignProperties_From_Servers_AzureADOnlyAuthentication_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_AzureADOnlyAuthentication_STATUS() to populate field Status")
	}
	authentication.Status = status

	// No error
	return nil
}

// AssignProperties_To_ServersAzureADOnlyAuthentication populates the provided destination ServersAzureADOnlyAuthentication from our ServersAzureADOnlyAuthentication
func (authentication *ServersAzureADOnlyAuthentication) AssignProperties_To_ServersAzureADOnlyAuthentication(destination *v1api20211101s.ServersAzureADOnlyAuthentication) error {

	// ObjectMeta
	destination.ObjectMeta = *authentication.ObjectMeta.DeepCopy()

	// Spec
	var spec v1api20211101s.Servers_AzureADOnlyAuthentication_Spec
	err := authentication.Spec.AssignProperties_To_Servers_AzureADOnlyAuthentication_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_AzureADOnlyAuthentication_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1api20211101s.Servers_AzureADOnlyAuthentication_STATUS
	err = authentication.Status.AssignProperties_To_Servers_AzureADOnlyAuthentication_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_AzureADOnlyAuthentication_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (authentication *ServersAzureADOnlyAuthentication) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: authentication.Spec.OriginalVersion(),
		Kind:    "ServersAzureADOnlyAuthentication",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/ServerAzureADOnlyAuthentications.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/azureADOnlyAuthentications/Default
type ServersAzureADOnlyAuthenticationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersAzureADOnlyAuthentication `json:"items"`
}

type Servers_AzureADOnlyAuthentication_Spec struct {
	// +kubebuilder:validation:Required
	// AzureADOnlyAuthentication: Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication *bool `json:"azureADOnlyAuthentication,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/Server resource
	Owner *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"Server"`
}

var _ genruntime.ARMTransformer = &Servers_AzureADOnlyAuthentication_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (authentication *Servers_AzureADOnlyAuthentication_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if authentication == nil {
		return nil, nil
	}
	result := &Servers_AzureADOnlyAuthentication_Spec_ARM{}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if authentication.AzureADOnlyAuthentication != nil {
		result.Properties = &AzureADOnlyAuthProperties_ARM{}
	}
	if authentication.AzureADOnlyAuthentication != nil {
		azureADOnlyAuthentication := *authentication.AzureADOnlyAuthentication
		result.Properties.AzureADOnlyAuthentication = &azureADOnlyAuthentication
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (authentication *Servers_AzureADOnlyAuthentication_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_AzureADOnlyAuthentication_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (authentication *Servers_AzureADOnlyAuthentication_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_AzureADOnlyAuthentication_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_AzureADOnlyAuthentication_Spec_ARM, got %T", armInput)
	}

	// Set property ‘AzureADOnlyAuthentication’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AzureADOnlyAuthentication != nil {
			azureADOnlyAuthentication := *typedInput.Properties.AzureADOnlyAuthentication
			authentication.AzureADOnlyAuthentication = &azureADOnlyAuthentication
		}
	}

	// Set property ‘Owner’:
	authentication.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Servers_AzureADOnlyAuthentication_Spec{}

// ConvertSpecFrom populates our Servers_AzureADOnlyAuthentication_Spec from the provided source
func (authentication *Servers_AzureADOnlyAuthentication_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1api20211101s.Servers_AzureADOnlyAuthentication_Spec)
	if ok {
		// Populate our instance from source
		return authentication.AssignProperties_From_Servers_AzureADOnlyAuthentication_Spec(src)
	}

	// Convert to an intermediate form
	src = &v1api20211101s.Servers_AzureADOnlyAuthentication_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = authentication.AssignProperties_From_Servers_AzureADOnlyAuthentication_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Servers_AzureADOnlyAuthentication_Spec
func (authentication *Servers_AzureADOnlyAuthentication_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1api20211101s.Servers_AzureADOnlyAuthentication_Spec)
	if ok {
		// Populate destination from our instance
		return authentication.AssignProperties_To_Servers_AzureADOnlyAuthentication_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20211101s.Servers_AzureADOnlyAuthentication_Spec{}
	err := authentication.AssignProperties_To_Servers_AzureADOnlyAuthentication_Spec(dst)
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

// AssignProperties_From_Servers_AzureADOnlyAuthentication_Spec populates our Servers_AzureADOnlyAuthentication_Spec from the provided source Servers_AzureADOnlyAuthentication_Spec
func (authentication *Servers_AzureADOnlyAuthentication_Spec) AssignProperties_From_Servers_AzureADOnlyAuthentication_Spec(source *v1api20211101s.Servers_AzureADOnlyAuthentication_Spec) error {

	// AzureADOnlyAuthentication
	if source.AzureADOnlyAuthentication != nil {
		azureADOnlyAuthentication := *source.AzureADOnlyAuthentication
		authentication.AzureADOnlyAuthentication = &azureADOnlyAuthentication
	} else {
		authentication.AzureADOnlyAuthentication = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		authentication.Owner = &owner
	} else {
		authentication.Owner = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Servers_AzureADOnlyAuthentication_Spec populates the provided destination Servers_AzureADOnlyAuthentication_Spec from our Servers_AzureADOnlyAuthentication_Spec
func (authentication *Servers_AzureADOnlyAuthentication_Spec) AssignProperties_To_Servers_AzureADOnlyAuthentication_Spec(destination *v1api20211101s.Servers_AzureADOnlyAuthentication_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureADOnlyAuthentication
	if authentication.AzureADOnlyAuthentication != nil {
		azureADOnlyAuthentication := *authentication.AzureADOnlyAuthentication
		destination.AzureADOnlyAuthentication = &azureADOnlyAuthentication
	} else {
		destination.AzureADOnlyAuthentication = nil
	}

	// OriginalVersion
	destination.OriginalVersion = authentication.OriginalVersion()

	// Owner
	if authentication.Owner != nil {
		owner := authentication.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
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

// Initialize_From_Servers_AzureADOnlyAuthentication_STATUS populates our Servers_AzureADOnlyAuthentication_Spec from the provided source Servers_AzureADOnlyAuthentication_STATUS
func (authentication *Servers_AzureADOnlyAuthentication_Spec) Initialize_From_Servers_AzureADOnlyAuthentication_STATUS(source *Servers_AzureADOnlyAuthentication_STATUS) error {

	// AzureADOnlyAuthentication
	if source.AzureADOnlyAuthentication != nil {
		azureADOnlyAuthentication := *source.AzureADOnlyAuthentication
		authentication.AzureADOnlyAuthentication = &azureADOnlyAuthentication
	} else {
		authentication.AzureADOnlyAuthentication = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (authentication *Servers_AzureADOnlyAuthentication_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type Servers_AzureADOnlyAuthentication_STATUS struct {
	// AzureADOnlyAuthentication: Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication *bool `json:"azureADOnlyAuthentication,omitempty"`

	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_AzureADOnlyAuthentication_STATUS{}

// ConvertStatusFrom populates our Servers_AzureADOnlyAuthentication_STATUS from the provided source
func (authentication *Servers_AzureADOnlyAuthentication_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1api20211101s.Servers_AzureADOnlyAuthentication_STATUS)
	if ok {
		// Populate our instance from source
		return authentication.AssignProperties_From_Servers_AzureADOnlyAuthentication_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v1api20211101s.Servers_AzureADOnlyAuthentication_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = authentication.AssignProperties_From_Servers_AzureADOnlyAuthentication_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Servers_AzureADOnlyAuthentication_STATUS
func (authentication *Servers_AzureADOnlyAuthentication_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1api20211101s.Servers_AzureADOnlyAuthentication_STATUS)
	if ok {
		// Populate destination from our instance
		return authentication.AssignProperties_To_Servers_AzureADOnlyAuthentication_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20211101s.Servers_AzureADOnlyAuthentication_STATUS{}
	err := authentication.AssignProperties_To_Servers_AzureADOnlyAuthentication_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Servers_AzureADOnlyAuthentication_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (authentication *Servers_AzureADOnlyAuthentication_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_AzureADOnlyAuthentication_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (authentication *Servers_AzureADOnlyAuthentication_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_AzureADOnlyAuthentication_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_AzureADOnlyAuthentication_STATUS_ARM, got %T", armInput)
	}

	// Set property ‘AzureADOnlyAuthentication’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AzureADOnlyAuthentication != nil {
			azureADOnlyAuthentication := *typedInput.Properties.AzureADOnlyAuthentication
			authentication.AzureADOnlyAuthentication = &azureADOnlyAuthentication
		}
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		authentication.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		authentication.Name = &name
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		authentication.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Servers_AzureADOnlyAuthentication_STATUS populates our Servers_AzureADOnlyAuthentication_STATUS from the provided source Servers_AzureADOnlyAuthentication_STATUS
func (authentication *Servers_AzureADOnlyAuthentication_STATUS) AssignProperties_From_Servers_AzureADOnlyAuthentication_STATUS(source *v1api20211101s.Servers_AzureADOnlyAuthentication_STATUS) error {

	// AzureADOnlyAuthentication
	if source.AzureADOnlyAuthentication != nil {
		azureADOnlyAuthentication := *source.AzureADOnlyAuthentication
		authentication.AzureADOnlyAuthentication = &azureADOnlyAuthentication
	} else {
		authentication.AzureADOnlyAuthentication = nil
	}

	// Conditions
	authentication.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	authentication.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	authentication.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	authentication.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_Servers_AzureADOnlyAuthentication_STATUS populates the provided destination Servers_AzureADOnlyAuthentication_STATUS from our Servers_AzureADOnlyAuthentication_STATUS
func (authentication *Servers_AzureADOnlyAuthentication_STATUS) AssignProperties_To_Servers_AzureADOnlyAuthentication_STATUS(destination *v1api20211101s.Servers_AzureADOnlyAuthentication_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureADOnlyAuthentication
	if authentication.AzureADOnlyAuthentication != nil {
		azureADOnlyAuthentication := *authentication.AzureADOnlyAuthentication
		destination.AzureADOnlyAuthentication = &azureADOnlyAuthentication
	} else {
		destination.AzureADOnlyAuthentication = nil
	}

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(authentication.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(authentication.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(authentication.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(authentication.Type)

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
	SchemeBuilder.Register(&ServersAzureADOnlyAuthentication{}, &ServersAzureADOnlyAuthenticationList{})
}

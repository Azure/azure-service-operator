// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901

import (
	"fmt"
	v1api20220901s "github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901storage"
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
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2022-09-01/table.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/default
type StorageAccountsTableService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccounts_TableService_Spec   `json:"spec,omitempty"`
	Status            StorageAccounts_TableService_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsTableService{}

// GetConditions returns the conditions of the resource
func (service *StorageAccountsTableService) GetConditions() conditions.Conditions {
	return service.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (service *StorageAccountsTableService) SetConditions(conditions conditions.Conditions) {
	service.Status.Conditions = conditions
}

var _ conversion.Convertible = &StorageAccountsTableService{}

// ConvertFrom populates our StorageAccountsTableService from the provided hub StorageAccountsTableService
func (service *StorageAccountsTableService) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1api20220901s.StorageAccountsTableService)
	if !ok {
		return fmt.Errorf("expected storage/v1api20220901storage/StorageAccountsTableService but received %T instead", hub)
	}

	return service.AssignProperties_From_StorageAccountsTableService(source)
}

// ConvertTo populates the provided hub StorageAccountsTableService from our StorageAccountsTableService
func (service *StorageAccountsTableService) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1api20220901s.StorageAccountsTableService)
	if !ok {
		return fmt.Errorf("expected storage/v1api20220901storage/StorageAccountsTableService but received %T instead", hub)
	}

	return service.AssignProperties_To_StorageAccountsTableService(destination)
}

// +kubebuilder:webhook:path=/mutate-storage-azure-com-v1api20220901-storageaccountstableservice,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountstableservices,verbs=create;update,versions=v1api20220901,name=default.v1api20220901.storageaccountstableservices.storage.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &StorageAccountsTableService{}

// Default applies defaults to the StorageAccountsTableService resource
func (service *StorageAccountsTableService) Default() {
	service.defaultImpl()
	var temp any = service
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the StorageAccountsTableService resource
func (service *StorageAccountsTableService) defaultImpl() {}

var _ genruntime.ImportableResource = &StorageAccountsTableService{}

// InitializeSpec initializes the spec for this resource from the given status
func (service *StorageAccountsTableService) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*StorageAccounts_TableService_STATUS); ok {
		return service.Spec.Initialize_From_StorageAccounts_TableService_STATUS(s)
	}

	return fmt.Errorf("expected Status of type StorageAccounts_TableService_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &StorageAccountsTableService{}

// AzureName returns the Azure name of the resource (always "default")
func (service *StorageAccountsTableService) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-09-01"
func (service StorageAccountsTableService) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (service *StorageAccountsTableService) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (service *StorageAccountsTableService) GetSpec() genruntime.ConvertibleSpec {
	return &service.Spec
}

// GetStatus returns the status of this resource
func (service *StorageAccountsTableService) GetStatus() genruntime.ConvertibleStatus {
	return &service.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/tableServices"
func (service *StorageAccountsTableService) GetType() string {
	return "Microsoft.Storage/storageAccounts/tableServices"
}

// NewEmptyStatus returns a new empty (blank) status
func (service *StorageAccountsTableService) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &StorageAccounts_TableService_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (service *StorageAccountsTableService) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(service.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  service.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (service *StorageAccountsTableService) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*StorageAccounts_TableService_STATUS); ok {
		service.Status = *st
		return nil
	}

	// Convert status to required version
	var st StorageAccounts_TableService_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	service.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-storage-azure-com-v1api20220901-storageaccountstableservice,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountstableservices,verbs=create;update,versions=v1api20220901,name=validate.v1api20220901.storageaccountstableservices.storage.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &StorageAccountsTableService{}

// ValidateCreate validates the creation of the resource
func (service *StorageAccountsTableService) ValidateCreate() error {
	validations := service.createValidations()
	var temp any = service
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
func (service *StorageAccountsTableService) ValidateDelete() error {
	validations := service.deleteValidations()
	var temp any = service
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
func (service *StorageAccountsTableService) ValidateUpdate(old runtime.Object) error {
	validations := service.updateValidations()
	var temp any = service
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
func (service *StorageAccountsTableService) createValidations() []func() error {
	return []func() error{service.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (service *StorageAccountsTableService) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (service *StorageAccountsTableService) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return service.validateResourceReferences()
		},
		service.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (service *StorageAccountsTableService) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&service.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (service *StorageAccountsTableService) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*StorageAccountsTableService)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, service)
}

// AssignProperties_From_StorageAccountsTableService populates our StorageAccountsTableService from the provided source StorageAccountsTableService
func (service *StorageAccountsTableService) AssignProperties_From_StorageAccountsTableService(source *v1api20220901s.StorageAccountsTableService) error {

	// ObjectMeta
	service.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec StorageAccounts_TableService_Spec
	err := spec.AssignProperties_From_StorageAccounts_TableService_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_StorageAccounts_TableService_Spec() to populate field Spec")
	}
	service.Spec = spec

	// Status
	var status StorageAccounts_TableService_STATUS
	err = status.AssignProperties_From_StorageAccounts_TableService_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_StorageAccounts_TableService_STATUS() to populate field Status")
	}
	service.Status = status

	// No error
	return nil
}

// AssignProperties_To_StorageAccountsTableService populates the provided destination StorageAccountsTableService from our StorageAccountsTableService
func (service *StorageAccountsTableService) AssignProperties_To_StorageAccountsTableService(destination *v1api20220901s.StorageAccountsTableService) error {

	// ObjectMeta
	destination.ObjectMeta = *service.ObjectMeta.DeepCopy()

	// Spec
	var spec v1api20220901s.StorageAccounts_TableService_Spec
	err := service.Spec.AssignProperties_To_StorageAccounts_TableService_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_StorageAccounts_TableService_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1api20220901s.StorageAccounts_TableService_STATUS
	err = service.Status.AssignProperties_To_StorageAccounts_TableService_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_StorageAccounts_TableService_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (service *StorageAccountsTableService) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: service.Spec.OriginalVersion(),
		Kind:    "StorageAccountsTableService",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2022-09-01/table.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/default
type StorageAccountsTableServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsTableService `json:"items"`
}

type StorageAccounts_TableService_Spec struct {
	// Cors: Specifies CORS rules for the Table service. You can include up to five CorsRule elements in the request. If no
	// CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the
	// Table service.
	Cors *CorsRules `json:"cors,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a storage.azure.com/StorageAccount resource
	Owner *genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner,omitempty" kind:"StorageAccount"`
}

var _ genruntime.ARMTransformer = &StorageAccounts_TableService_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (service *StorageAccounts_TableService_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if service == nil {
		return nil, nil
	}
	result := &StorageAccounts_TableService_Spec_ARM{}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if service.Cors != nil {
		result.Properties = &StorageAccounts_TableService_Properties_Spec_ARM{}
	}
	if service.Cors != nil {
		cors_ARM, err := (*service.Cors).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		cors := *cors_ARM.(*CorsRules_ARM)
		result.Properties.Cors = &cors
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (service *StorageAccounts_TableService_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &StorageAccounts_TableService_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (service *StorageAccounts_TableService_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(StorageAccounts_TableService_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected StorageAccounts_TableService_Spec_ARM, got %T", armInput)
	}

	// Set property ‘Cors’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Cors != nil {
			var cors1 CorsRules
			err := cors1.PopulateFromARM(owner, *typedInput.Properties.Cors)
			if err != nil {
				return err
			}
			cors := cors1
			service.Cors = &cors
		}
	}

	// Set property ‘Owner’:
	service.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &StorageAccounts_TableService_Spec{}

// ConvertSpecFrom populates our StorageAccounts_TableService_Spec from the provided source
func (service *StorageAccounts_TableService_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1api20220901s.StorageAccounts_TableService_Spec)
	if ok {
		// Populate our instance from source
		return service.AssignProperties_From_StorageAccounts_TableService_Spec(src)
	}

	// Convert to an intermediate form
	src = &v1api20220901s.StorageAccounts_TableService_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = service.AssignProperties_From_StorageAccounts_TableService_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our StorageAccounts_TableService_Spec
func (service *StorageAccounts_TableService_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1api20220901s.StorageAccounts_TableService_Spec)
	if ok {
		// Populate destination from our instance
		return service.AssignProperties_To_StorageAccounts_TableService_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20220901s.StorageAccounts_TableService_Spec{}
	err := service.AssignProperties_To_StorageAccounts_TableService_Spec(dst)
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

// AssignProperties_From_StorageAccounts_TableService_Spec populates our StorageAccounts_TableService_Spec from the provided source StorageAccounts_TableService_Spec
func (service *StorageAccounts_TableService_Spec) AssignProperties_From_StorageAccounts_TableService_Spec(source *v1api20220901s.StorageAccounts_TableService_Spec) error {

	// Cors
	if source.Cors != nil {
		var cor CorsRules
		err := cor.AssignProperties_From_CorsRules(source.Cors)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_CorsRules() to populate field Cors")
		}
		service.Cors = &cor
	} else {
		service.Cors = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		service.Owner = &owner
	} else {
		service.Owner = nil
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccounts_TableService_Spec populates the provided destination StorageAccounts_TableService_Spec from our StorageAccounts_TableService_Spec
func (service *StorageAccounts_TableService_Spec) AssignProperties_To_StorageAccounts_TableService_Spec(destination *v1api20220901s.StorageAccounts_TableService_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Cors
	if service.Cors != nil {
		var cor v1api20220901s.CorsRules
		err := service.Cors.AssignProperties_To_CorsRules(&cor)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_CorsRules() to populate field Cors")
		}
		destination.Cors = &cor
	} else {
		destination.Cors = nil
	}

	// OriginalVersion
	destination.OriginalVersion = service.OriginalVersion()

	// Owner
	if service.Owner != nil {
		owner := service.Owner.Copy()
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

// Initialize_From_StorageAccounts_TableService_STATUS populates our StorageAccounts_TableService_Spec from the provided source StorageAccounts_TableService_STATUS
func (service *StorageAccounts_TableService_Spec) Initialize_From_StorageAccounts_TableService_STATUS(source *StorageAccounts_TableService_STATUS) error {

	// Cors
	if source.Cors != nil {
		var cor CorsRules
		err := cor.Initialize_From_CorsRules_STATUS(source.Cors)
		if err != nil {
			return errors.Wrap(err, "calling Initialize_From_CorsRules_STATUS() to populate field Cors")
		}
		service.Cors = &cor
	} else {
		service.Cors = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (service *StorageAccounts_TableService_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type StorageAccounts_TableService_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Cors: Specifies CORS rules for the Table service. You can include up to five CorsRule elements in the request. If no
	// CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the
	// Table service.
	Cors *CorsRules_STATUS `json:"cors,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageAccounts_TableService_STATUS{}

// ConvertStatusFrom populates our StorageAccounts_TableService_STATUS from the provided source
func (service *StorageAccounts_TableService_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1api20220901s.StorageAccounts_TableService_STATUS)
	if ok {
		// Populate our instance from source
		return service.AssignProperties_From_StorageAccounts_TableService_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v1api20220901s.StorageAccounts_TableService_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = service.AssignProperties_From_StorageAccounts_TableService_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our StorageAccounts_TableService_STATUS
func (service *StorageAccounts_TableService_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1api20220901s.StorageAccounts_TableService_STATUS)
	if ok {
		// Populate destination from our instance
		return service.AssignProperties_To_StorageAccounts_TableService_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20220901s.StorageAccounts_TableService_STATUS{}
	err := service.AssignProperties_To_StorageAccounts_TableService_STATUS(dst)
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

var _ genruntime.FromARMConverter = &StorageAccounts_TableService_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (service *StorageAccounts_TableService_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &StorageAccounts_TableService_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (service *StorageAccounts_TableService_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(StorageAccounts_TableService_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected StorageAccounts_TableService_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Cors’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Cors != nil {
			var cors1 CorsRules_STATUS
			err := cors1.PopulateFromARM(owner, *typedInput.Properties.Cors)
			if err != nil {
				return err
			}
			cors := cors1
			service.Cors = &cors
		}
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		service.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		service.Name = &name
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		service.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_StorageAccounts_TableService_STATUS populates our StorageAccounts_TableService_STATUS from the provided source StorageAccounts_TableService_STATUS
func (service *StorageAccounts_TableService_STATUS) AssignProperties_From_StorageAccounts_TableService_STATUS(source *v1api20220901s.StorageAccounts_TableService_STATUS) error {

	// Conditions
	service.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Cors
	if source.Cors != nil {
		var cor CorsRules_STATUS
		err := cor.AssignProperties_From_CorsRules_STATUS(source.Cors)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_CorsRules_STATUS() to populate field Cors")
		}
		service.Cors = &cor
	} else {
		service.Cors = nil
	}

	// Id
	service.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	service.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	service.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_StorageAccounts_TableService_STATUS populates the provided destination StorageAccounts_TableService_STATUS from our StorageAccounts_TableService_STATUS
func (service *StorageAccounts_TableService_STATUS) AssignProperties_To_StorageAccounts_TableService_STATUS(destination *v1api20220901s.StorageAccounts_TableService_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(service.Conditions)

	// Cors
	if service.Cors != nil {
		var cor v1api20220901s.CorsRules_STATUS
		err := service.Cors.AssignProperties_To_CorsRules_STATUS(&cor)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_CorsRules_STATUS() to populate field Cors")
		}
		destination.Cors = &cor
	} else {
		destination.Cors = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(service.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(service.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(service.Type)

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
	SchemeBuilder.Register(&StorageAccountsTableService{}, &StorageAccountsTableServiceList{})
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901/storage"
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
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2022-09-01/queue.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default
type StorageAccountsQueueService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccounts_QueueService_Spec   `json:"spec,omitempty"`
	Status            StorageAccounts_QueueService_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsQueueService{}

// GetConditions returns the conditions of the resource
func (service *StorageAccountsQueueService) GetConditions() conditions.Conditions {
	return service.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (service *StorageAccountsQueueService) SetConditions(conditions conditions.Conditions) {
	service.Status.Conditions = conditions
}

var _ conversion.Convertible = &StorageAccountsQueueService{}

// ConvertFrom populates our StorageAccountsQueueService from the provided hub StorageAccountsQueueService
func (service *StorageAccountsQueueService) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source storage.StorageAccountsQueueService

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = service.AssignProperties_From_StorageAccountsQueueService(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to service")
	}

	return nil
}

// ConvertTo populates the provided hub StorageAccountsQueueService from our StorageAccountsQueueService
func (service *StorageAccountsQueueService) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination storage.StorageAccountsQueueService
	err := service.AssignProperties_To_StorageAccountsQueueService(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from service")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-storage-azure-com-v1api20220901-storageaccountsqueueservice,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountsqueueservices,verbs=create;update,versions=v1api20220901,name=default.v1api20220901.storageaccountsqueueservices.storage.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &StorageAccountsQueueService{}

// Default applies defaults to the StorageAccountsQueueService resource
func (service *StorageAccountsQueueService) Default() {
	service.defaultImpl()
	var temp any = service
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the StorageAccountsQueueService resource
func (service *StorageAccountsQueueService) defaultImpl() {}

var _ genruntime.KubernetesResource = &StorageAccountsQueueService{}

// AzureName returns the Azure name of the resource (always "default")
func (service *StorageAccountsQueueService) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-09-01"
func (service StorageAccountsQueueService) GetAPIVersion() string {
	return "2022-09-01"
}

// GetResourceScope returns the scope of the resource
func (service *StorageAccountsQueueService) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (service *StorageAccountsQueueService) GetSpec() genruntime.ConvertibleSpec {
	return &service.Spec
}

// GetStatus returns the status of this resource
func (service *StorageAccountsQueueService) GetStatus() genruntime.ConvertibleStatus {
	return &service.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (service *StorageAccountsQueueService) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/queueServices"
func (service *StorageAccountsQueueService) GetType() string {
	return "Microsoft.Storage/storageAccounts/queueServices"
}

// NewEmptyStatus returns a new empty (blank) status
func (service *StorageAccountsQueueService) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &StorageAccounts_QueueService_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (service *StorageAccountsQueueService) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(service.Spec)
	return service.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (service *StorageAccountsQueueService) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*StorageAccounts_QueueService_STATUS); ok {
		service.Status = *st
		return nil
	}

	// Convert status to required version
	var st StorageAccounts_QueueService_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	service.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-storage-azure-com-v1api20220901-storageaccountsqueueservice,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountsqueueservices,verbs=create;update,versions=v1api20220901,name=validate.v1api20220901.storageaccountsqueueservices.storage.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &StorageAccountsQueueService{}

// ValidateCreate validates the creation of the resource
func (service *StorageAccountsQueueService) ValidateCreate() (admission.Warnings, error) {
	validations := service.createValidations()
	var temp any = service
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (service *StorageAccountsQueueService) ValidateDelete() (admission.Warnings, error) {
	validations := service.deleteValidations()
	var temp any = service
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (service *StorageAccountsQueueService) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := service.updateValidations()
	var temp any = service
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (service *StorageAccountsQueueService) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){service.validateResourceReferences, service.validateOwnerReference}
}

// deleteValidations validates the deletion of the resource
func (service *StorageAccountsQueueService) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (service *StorageAccountsQueueService) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return service.validateResourceReferences()
		},
		service.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return service.validateOwnerReference()
		},
	}
}

// validateOwnerReference validates the owner field
func (service *StorageAccountsQueueService) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(service)
}

// validateResourceReferences validates all resource references
func (service *StorageAccountsQueueService) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&service.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (service *StorageAccountsQueueService) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*StorageAccountsQueueService)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, service)
}

// AssignProperties_From_StorageAccountsQueueService populates our StorageAccountsQueueService from the provided source StorageAccountsQueueService
func (service *StorageAccountsQueueService) AssignProperties_From_StorageAccountsQueueService(source *storage.StorageAccountsQueueService) error {

	// ObjectMeta
	service.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec StorageAccounts_QueueService_Spec
	err := spec.AssignProperties_From_StorageAccounts_QueueService_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_StorageAccounts_QueueService_Spec() to populate field Spec")
	}
	service.Spec = spec

	// Status
	var status StorageAccounts_QueueService_STATUS
	err = status.AssignProperties_From_StorageAccounts_QueueService_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_StorageAccounts_QueueService_STATUS() to populate field Status")
	}
	service.Status = status

	// No error
	return nil
}

// AssignProperties_To_StorageAccountsQueueService populates the provided destination StorageAccountsQueueService from our StorageAccountsQueueService
func (service *StorageAccountsQueueService) AssignProperties_To_StorageAccountsQueueService(destination *storage.StorageAccountsQueueService) error {

	// ObjectMeta
	destination.ObjectMeta = *service.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.StorageAccounts_QueueService_Spec
	err := service.Spec.AssignProperties_To_StorageAccounts_QueueService_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_StorageAccounts_QueueService_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.StorageAccounts_QueueService_STATUS
	err = service.Status.AssignProperties_To_StorageAccounts_QueueService_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_StorageAccounts_QueueService_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (service *StorageAccountsQueueService) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: service.Spec.OriginalVersion(),
		Kind:    "StorageAccountsQueueService",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2022-09-01/queue.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default
type StorageAccountsQueueServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsQueueService `json:"items"`
}

type StorageAccounts_QueueService_Spec struct {
	// Cors: Specifies CORS rules for the Queue service. You can include up to five CorsRule elements in the request. If no
	// CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the
	// Queue service.
	Cors *CorsRules `json:"cors,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a storage.azure.com/StorageAccount resource
	Owner *genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner,omitempty" kind:"StorageAccount"`
}

var _ genruntime.ARMTransformer = &StorageAccounts_QueueService_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (service *StorageAccounts_QueueService_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if service == nil {
		return nil, nil
	}
	result := &StorageAccounts_QueueService_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if service.Cors != nil {
		result.Properties = &StorageAccounts_QueueService_Properties_Spec_ARM{}
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
func (service *StorageAccounts_QueueService_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &StorageAccounts_QueueService_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (service *StorageAccounts_QueueService_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(StorageAccounts_QueueService_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected StorageAccounts_QueueService_Spec_ARM, got %T", armInput)
	}

	// Set property "Cors":
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

	// Set property "Owner":
	service.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &StorageAccounts_QueueService_Spec{}

// ConvertSpecFrom populates our StorageAccounts_QueueService_Spec from the provided source
func (service *StorageAccounts_QueueService_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.StorageAccounts_QueueService_Spec)
	if ok {
		// Populate our instance from source
		return service.AssignProperties_From_StorageAccounts_QueueService_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.StorageAccounts_QueueService_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = service.AssignProperties_From_StorageAccounts_QueueService_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our StorageAccounts_QueueService_Spec
func (service *StorageAccounts_QueueService_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.StorageAccounts_QueueService_Spec)
	if ok {
		// Populate destination from our instance
		return service.AssignProperties_To_StorageAccounts_QueueService_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.StorageAccounts_QueueService_Spec{}
	err := service.AssignProperties_To_StorageAccounts_QueueService_Spec(dst)
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

// AssignProperties_From_StorageAccounts_QueueService_Spec populates our StorageAccounts_QueueService_Spec from the provided source StorageAccounts_QueueService_Spec
func (service *StorageAccounts_QueueService_Spec) AssignProperties_From_StorageAccounts_QueueService_Spec(source *storage.StorageAccounts_QueueService_Spec) error {

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

// AssignProperties_To_StorageAccounts_QueueService_Spec populates the provided destination StorageAccounts_QueueService_Spec from our StorageAccounts_QueueService_Spec
func (service *StorageAccounts_QueueService_Spec) AssignProperties_To_StorageAccounts_QueueService_Spec(destination *storage.StorageAccounts_QueueService_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Cors
	if service.Cors != nil {
		var cor storage.CorsRules
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

// OriginalVersion returns the original API version used to create the resource.
func (service *StorageAccounts_QueueService_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type StorageAccounts_QueueService_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Cors: Specifies CORS rules for the Queue service. You can include up to five CorsRule elements in the request. If no
	// CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the
	// Queue service.
	Cors *CorsRules_STATUS `json:"cors,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageAccounts_QueueService_STATUS{}

// ConvertStatusFrom populates our StorageAccounts_QueueService_STATUS from the provided source
func (service *StorageAccounts_QueueService_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.StorageAccounts_QueueService_STATUS)
	if ok {
		// Populate our instance from source
		return service.AssignProperties_From_StorageAccounts_QueueService_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.StorageAccounts_QueueService_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = service.AssignProperties_From_StorageAccounts_QueueService_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our StorageAccounts_QueueService_STATUS
func (service *StorageAccounts_QueueService_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.StorageAccounts_QueueService_STATUS)
	if ok {
		// Populate destination from our instance
		return service.AssignProperties_To_StorageAccounts_QueueService_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.StorageAccounts_QueueService_STATUS{}
	err := service.AssignProperties_To_StorageAccounts_QueueService_STATUS(dst)
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

var _ genruntime.FromARMConverter = &StorageAccounts_QueueService_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (service *StorageAccounts_QueueService_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &StorageAccounts_QueueService_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (service *StorageAccounts_QueueService_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(StorageAccounts_QueueService_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected StorageAccounts_QueueService_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "Cors":
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

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		service.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		service.Name = &name
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		service.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_StorageAccounts_QueueService_STATUS populates our StorageAccounts_QueueService_STATUS from the provided source StorageAccounts_QueueService_STATUS
func (service *StorageAccounts_QueueService_STATUS) AssignProperties_From_StorageAccounts_QueueService_STATUS(source *storage.StorageAccounts_QueueService_STATUS) error {

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

// AssignProperties_To_StorageAccounts_QueueService_STATUS populates the provided destination StorageAccounts_QueueService_STATUS from our StorageAccounts_QueueService_STATUS
func (service *StorageAccounts_QueueService_STATUS) AssignProperties_To_StorageAccounts_QueueService_STATUS(destination *storage.StorageAccounts_QueueService_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(service.Conditions)

	// Cors
	if service.Cors != nil {
		var cor storage.CorsRules_STATUS
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
	SchemeBuilder.Register(&StorageAccountsQueueService{}, &StorageAccountsQueueServiceList{})
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210401storage

import (
	"fmt"
	v20220901s "github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20210401.StorageAccountsQueueService
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2021-04-01/queue.json
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
	source, ok := hub.(*v20220901s.StorageAccountsQueueService)
	if !ok {
		return fmt.Errorf("expected storage/v1api20220901storage/StorageAccountsQueueService but received %T instead", hub)
	}

	return service.AssignProperties_From_StorageAccountsQueueService(source)
}

// ConvertTo populates the provided hub StorageAccountsQueueService from our StorageAccountsQueueService
func (service *StorageAccountsQueueService) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20220901s.StorageAccountsQueueService)
	if !ok {
		return fmt.Errorf("expected storage/v1api20220901storage/StorageAccountsQueueService but received %T instead", hub)
	}

	return service.AssignProperties_To_StorageAccountsQueueService(destination)
}

var _ genruntime.KubernetesResource = &StorageAccountsQueueService{}

// AzureName returns the Azure name of the resource (always "default")
func (service *StorageAccountsQueueService) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (service StorageAccountsQueueService) GetAPIVersion() string {
	return string(APIVersion_Value)
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

// AssignProperties_From_StorageAccountsQueueService populates our StorageAccountsQueueService from the provided source StorageAccountsQueueService
func (service *StorageAccountsQueueService) AssignProperties_From_StorageAccountsQueueService(source *v20220901s.StorageAccountsQueueService) error {

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

	// Invoke the augmentConversionForStorageAccountsQueueService interface (if implemented) to customize the conversion
	var serviceAsAny any = service
	if augmentedService, ok := serviceAsAny.(augmentConversionForStorageAccountsQueueService); ok {
		err := augmentedService.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccountsQueueService populates the provided destination StorageAccountsQueueService from our StorageAccountsQueueService
func (service *StorageAccountsQueueService) AssignProperties_To_StorageAccountsQueueService(destination *v20220901s.StorageAccountsQueueService) error {

	// ObjectMeta
	destination.ObjectMeta = *service.ObjectMeta.DeepCopy()

	// Spec
	var spec v20220901s.StorageAccounts_QueueService_Spec
	err := service.Spec.AssignProperties_To_StorageAccounts_QueueService_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_StorageAccounts_QueueService_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20220901s.StorageAccounts_QueueService_STATUS
	err = service.Status.AssignProperties_To_StorageAccounts_QueueService_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_StorageAccounts_QueueService_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForStorageAccountsQueueService interface (if implemented) to customize the conversion
	var serviceAsAny any = service
	if augmentedService, ok := serviceAsAny.(augmentConversionForStorageAccountsQueueService); ok {
		err := augmentedService.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (service *StorageAccountsQueueService) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: service.Spec.OriginalVersion,
		Kind:    "StorageAccountsQueueService",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210401.StorageAccountsQueueService
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2021-04-01/queue.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default
type StorageAccountsQueueServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsQueueService `json:"items"`
}

type augmentConversionForStorageAccountsQueueService interface {
	AssignPropertiesFrom(src *v20220901s.StorageAccountsQueueService) error
	AssignPropertiesTo(dst *v20220901s.StorageAccountsQueueService) error
}

// Storage version of v1api20210401.StorageAccounts_QueueService_Spec
type StorageAccounts_QueueService_Spec struct {
	Cors            *CorsRules `json:"cors,omitempty"`
	OriginalVersion string     `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a storage.azure.com/StorageAccount resource
	Owner       *genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner,omitempty" kind:"StorageAccount"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &StorageAccounts_QueueService_Spec{}

// ConvertSpecFrom populates our StorageAccounts_QueueService_Spec from the provided source
func (service *StorageAccounts_QueueService_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20220901s.StorageAccounts_QueueService_Spec)
	if ok {
		// Populate our instance from source
		return service.AssignProperties_From_StorageAccounts_QueueService_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20220901s.StorageAccounts_QueueService_Spec{}
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
	dst, ok := destination.(*v20220901s.StorageAccounts_QueueService_Spec)
	if ok {
		// Populate destination from our instance
		return service.AssignProperties_To_StorageAccounts_QueueService_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20220901s.StorageAccounts_QueueService_Spec{}
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
func (service *StorageAccounts_QueueService_Spec) AssignProperties_From_StorageAccounts_QueueService_Spec(source *v20220901s.StorageAccounts_QueueService_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

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

	// OriginalVersion
	service.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		service.Owner = &owner
	} else {
		service.Owner = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		service.PropertyBag = propertyBag
	} else {
		service.PropertyBag = nil
	}

	// Invoke the augmentConversionForStorageAccounts_QueueService_Spec interface (if implemented) to customize the conversion
	var serviceAsAny any = service
	if augmentedService, ok := serviceAsAny.(augmentConversionForStorageAccounts_QueueService_Spec); ok {
		err := augmentedService.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccounts_QueueService_Spec populates the provided destination StorageAccounts_QueueService_Spec from our StorageAccounts_QueueService_Spec
func (service *StorageAccounts_QueueService_Spec) AssignProperties_To_StorageAccounts_QueueService_Spec(destination *v20220901s.StorageAccounts_QueueService_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(service.PropertyBag)

	// Cors
	if service.Cors != nil {
		var cor v20220901s.CorsRules
		err := service.Cors.AssignProperties_To_CorsRules(&cor)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_CorsRules() to populate field Cors")
		}
		destination.Cors = &cor
	} else {
		destination.Cors = nil
	}

	// OriginalVersion
	destination.OriginalVersion = service.OriginalVersion

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

	// Invoke the augmentConversionForStorageAccounts_QueueService_Spec interface (if implemented) to customize the conversion
	var serviceAsAny any = service
	if augmentedService, ok := serviceAsAny.(augmentConversionForStorageAccounts_QueueService_Spec); ok {
		err := augmentedService.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20210401.StorageAccounts_QueueService_STATUS
type StorageAccounts_QueueService_STATUS struct {
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Cors        *CorsRules_STATUS      `json:"cors,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageAccounts_QueueService_STATUS{}

// ConvertStatusFrom populates our StorageAccounts_QueueService_STATUS from the provided source
func (service *StorageAccounts_QueueService_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20220901s.StorageAccounts_QueueService_STATUS)
	if ok {
		// Populate our instance from source
		return service.AssignProperties_From_StorageAccounts_QueueService_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20220901s.StorageAccounts_QueueService_STATUS{}
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
	dst, ok := destination.(*v20220901s.StorageAccounts_QueueService_STATUS)
	if ok {
		// Populate destination from our instance
		return service.AssignProperties_To_StorageAccounts_QueueService_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20220901s.StorageAccounts_QueueService_STATUS{}
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

// AssignProperties_From_StorageAccounts_QueueService_STATUS populates our StorageAccounts_QueueService_STATUS from the provided source StorageAccounts_QueueService_STATUS
func (service *StorageAccounts_QueueService_STATUS) AssignProperties_From_StorageAccounts_QueueService_STATUS(source *v20220901s.StorageAccounts_QueueService_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

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

	// Update the property bag
	if len(propertyBag) > 0 {
		service.PropertyBag = propertyBag
	} else {
		service.PropertyBag = nil
	}

	// Invoke the augmentConversionForStorageAccounts_QueueService_STATUS interface (if implemented) to customize the conversion
	var serviceAsAny any = service
	if augmentedService, ok := serviceAsAny.(augmentConversionForStorageAccounts_QueueService_STATUS); ok {
		err := augmentedService.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccounts_QueueService_STATUS populates the provided destination StorageAccounts_QueueService_STATUS from our StorageAccounts_QueueService_STATUS
func (service *StorageAccounts_QueueService_STATUS) AssignProperties_To_StorageAccounts_QueueService_STATUS(destination *v20220901s.StorageAccounts_QueueService_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(service.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(service.Conditions)

	// Cors
	if service.Cors != nil {
		var cor v20220901s.CorsRules_STATUS
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

	// Invoke the augmentConversionForStorageAccounts_QueueService_STATUS interface (if implemented) to customize the conversion
	var serviceAsAny any = service
	if augmentedService, ok := serviceAsAny.(augmentConversionForStorageAccounts_QueueService_STATUS); ok {
		err := augmentedService.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForStorageAccounts_QueueService_Spec interface {
	AssignPropertiesFrom(src *v20220901s.StorageAccounts_QueueService_Spec) error
	AssignPropertiesTo(dst *v20220901s.StorageAccounts_QueueService_Spec) error
}

type augmentConversionForStorageAccounts_QueueService_STATUS interface {
	AssignPropertiesFrom(src *v20220901s.StorageAccounts_QueueService_STATUS) error
	AssignPropertiesTo(dst *v20220901s.StorageAccounts_QueueService_STATUS) error
}

func init() {
	SchemeBuilder.Register(&StorageAccountsQueueService{}, &StorageAccountsQueueServiceList{})
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
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
// Storage version of v1api20220901.StorageAccountsTableService
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2022-09-01/table.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/default
type StorageAccountsTableService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccountsTableService_Spec   `json:"spec,omitempty"`
	Status            StorageAccountsTableService_STATUS `json:"status,omitempty"`
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
	source, ok := hub.(*storage.StorageAccountsTableService)
	if !ok {
		return fmt.Errorf("expected storage/v1api20230101/storage/StorageAccountsTableService but received %T instead", hub)
	}

	return service.AssignProperties_From_StorageAccountsTableService(source)
}

// ConvertTo populates the provided hub StorageAccountsTableService from our StorageAccountsTableService
func (service *StorageAccountsTableService) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.StorageAccountsTableService)
	if !ok {
		return fmt.Errorf("expected storage/v1api20230101/storage/StorageAccountsTableService but received %T instead", hub)
	}

	return service.AssignProperties_To_StorageAccountsTableService(destination)
}

var _ configmaps.Exporter = &StorageAccountsTableService{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (service *StorageAccountsTableService) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if service.Spec.OperatorSpec == nil {
		return nil
	}
	return service.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &StorageAccountsTableService{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (service *StorageAccountsTableService) SecretDestinationExpressions() []*core.DestinationExpression {
	if service.Spec.OperatorSpec == nil {
		return nil
	}
	return service.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &StorageAccountsTableService{}

// AzureName returns the Azure name of the resource (always "default")
func (service *StorageAccountsTableService) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-09-01"
func (service StorageAccountsTableService) GetAPIVersion() string {
	return "2022-09-01"
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

// GetSupportedOperations returns the operations supported by the resource
func (service *StorageAccountsTableService) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/tableServices"
func (service *StorageAccountsTableService) GetType() string {
	return "Microsoft.Storage/storageAccounts/tableServices"
}

// NewEmptyStatus returns a new empty (blank) status
func (service *StorageAccountsTableService) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &StorageAccountsTableService_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (service *StorageAccountsTableService) Owner() *genruntime.ResourceReference {
	if service.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(service.Spec)
	return service.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (service *StorageAccountsTableService) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*StorageAccountsTableService_STATUS); ok {
		service.Status = *st
		return nil
	}

	// Convert status to required version
	var st StorageAccountsTableService_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	service.Status = st
	return nil
}

// AssignProperties_From_StorageAccountsTableService populates our StorageAccountsTableService from the provided source StorageAccountsTableService
func (service *StorageAccountsTableService) AssignProperties_From_StorageAccountsTableService(source *storage.StorageAccountsTableService) error {

	// ObjectMeta
	service.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec StorageAccountsTableService_Spec
	err := spec.AssignProperties_From_StorageAccountsTableService_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_StorageAccountsTableService_Spec() to populate field Spec")
	}
	service.Spec = spec

	// Status
	var status StorageAccountsTableService_STATUS
	err = status.AssignProperties_From_StorageAccountsTableService_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_StorageAccountsTableService_STATUS() to populate field Status")
	}
	service.Status = status

	// Invoke the augmentConversionForStorageAccountsTableService interface (if implemented) to customize the conversion
	var serviceAsAny any = service
	if augmentedService, ok := serviceAsAny.(augmentConversionForStorageAccountsTableService); ok {
		err := augmentedService.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccountsTableService populates the provided destination StorageAccountsTableService from our StorageAccountsTableService
func (service *StorageAccountsTableService) AssignProperties_To_StorageAccountsTableService(destination *storage.StorageAccountsTableService) error {

	// ObjectMeta
	destination.ObjectMeta = *service.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.StorageAccountsTableService_Spec
	err := service.Spec.AssignProperties_To_StorageAccountsTableService_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_StorageAccountsTableService_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.StorageAccountsTableService_STATUS
	err = service.Status.AssignProperties_To_StorageAccountsTableService_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_StorageAccountsTableService_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForStorageAccountsTableService interface (if implemented) to customize the conversion
	var serviceAsAny any = service
	if augmentedService, ok := serviceAsAny.(augmentConversionForStorageAccountsTableService); ok {
		err := augmentedService.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (service *StorageAccountsTableService) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: service.Spec.OriginalVersion,
		Kind:    "StorageAccountsTableService",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220901.StorageAccountsTableService
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2022-09-01/table.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/default
type StorageAccountsTableServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsTableService `json:"items"`
}

type augmentConversionForStorageAccountsTableService interface {
	AssignPropertiesFrom(src *storage.StorageAccountsTableService) error
	AssignPropertiesTo(dst *storage.StorageAccountsTableService) error
}

// Storage version of v1api20220901.StorageAccountsTableService_Spec
type StorageAccountsTableService_Spec struct {
	Cors            *CorsRules                               `json:"cors,omitempty"`
	OperatorSpec    *StorageAccountsTableServiceOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                                   `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a storage.azure.com/StorageAccount resource
	Owner       *genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner,omitempty" kind:"StorageAccount"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &StorageAccountsTableService_Spec{}

// ConvertSpecFrom populates our StorageAccountsTableService_Spec from the provided source
func (service *StorageAccountsTableService_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.StorageAccountsTableService_Spec)
	if ok {
		// Populate our instance from source
		return service.AssignProperties_From_StorageAccountsTableService_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.StorageAccountsTableService_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = service.AssignProperties_From_StorageAccountsTableService_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our StorageAccountsTableService_Spec
func (service *StorageAccountsTableService_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.StorageAccountsTableService_Spec)
	if ok {
		// Populate destination from our instance
		return service.AssignProperties_To_StorageAccountsTableService_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.StorageAccountsTableService_Spec{}
	err := service.AssignProperties_To_StorageAccountsTableService_Spec(dst)
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

// AssignProperties_From_StorageAccountsTableService_Spec populates our StorageAccountsTableService_Spec from the provided source StorageAccountsTableService_Spec
func (service *StorageAccountsTableService_Spec) AssignProperties_From_StorageAccountsTableService_Spec(source *storage.StorageAccountsTableService_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Cors
	if source.Cors != nil {
		var cor CorsRules
		err := cor.AssignProperties_From_CorsRules(source.Cors)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_CorsRules() to populate field Cors")
		}
		service.Cors = &cor
	} else {
		service.Cors = nil
	}

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec StorageAccountsTableServiceOperatorSpec
		err := operatorSpec.AssignProperties_From_StorageAccountsTableServiceOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_StorageAccountsTableServiceOperatorSpec() to populate field OperatorSpec")
		}
		service.OperatorSpec = &operatorSpec
	} else {
		service.OperatorSpec = nil
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

	// Invoke the augmentConversionForStorageAccountsTableService_Spec interface (if implemented) to customize the conversion
	var serviceAsAny any = service
	if augmentedService, ok := serviceAsAny.(augmentConversionForStorageAccountsTableService_Spec); ok {
		err := augmentedService.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccountsTableService_Spec populates the provided destination StorageAccountsTableService_Spec from our StorageAccountsTableService_Spec
func (service *StorageAccountsTableService_Spec) AssignProperties_To_StorageAccountsTableService_Spec(destination *storage.StorageAccountsTableService_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(service.PropertyBag)

	// Cors
	if service.Cors != nil {
		var cor storage.CorsRules
		err := service.Cors.AssignProperties_To_CorsRules(&cor)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_CorsRules() to populate field Cors")
		}
		destination.Cors = &cor
	} else {
		destination.Cors = nil
	}

	// OperatorSpec
	if service.OperatorSpec != nil {
		var operatorSpec storage.StorageAccountsTableServiceOperatorSpec
		err := service.OperatorSpec.AssignProperties_To_StorageAccountsTableServiceOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_StorageAccountsTableServiceOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
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

	// Invoke the augmentConversionForStorageAccountsTableService_Spec interface (if implemented) to customize the conversion
	var serviceAsAny any = service
	if augmentedService, ok := serviceAsAny.(augmentConversionForStorageAccountsTableService_Spec); ok {
		err := augmentedService.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20220901.StorageAccountsTableService_STATUS
type StorageAccountsTableService_STATUS struct {
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Cors        *CorsRules_STATUS      `json:"cors,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageAccountsTableService_STATUS{}

// ConvertStatusFrom populates our StorageAccountsTableService_STATUS from the provided source
func (service *StorageAccountsTableService_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.StorageAccountsTableService_STATUS)
	if ok {
		// Populate our instance from source
		return service.AssignProperties_From_StorageAccountsTableService_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.StorageAccountsTableService_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = service.AssignProperties_From_StorageAccountsTableService_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our StorageAccountsTableService_STATUS
func (service *StorageAccountsTableService_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.StorageAccountsTableService_STATUS)
	if ok {
		// Populate destination from our instance
		return service.AssignProperties_To_StorageAccountsTableService_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.StorageAccountsTableService_STATUS{}
	err := service.AssignProperties_To_StorageAccountsTableService_STATUS(dst)
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

// AssignProperties_From_StorageAccountsTableService_STATUS populates our StorageAccountsTableService_STATUS from the provided source StorageAccountsTableService_STATUS
func (service *StorageAccountsTableService_STATUS) AssignProperties_From_StorageAccountsTableService_STATUS(source *storage.StorageAccountsTableService_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	service.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Cors
	if source.Cors != nil {
		var cor CorsRules_STATUS
		err := cor.AssignProperties_From_CorsRules_STATUS(source.Cors)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_CorsRules_STATUS() to populate field Cors")
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

	// Invoke the augmentConversionForStorageAccountsTableService_STATUS interface (if implemented) to customize the conversion
	var serviceAsAny any = service
	if augmentedService, ok := serviceAsAny.(augmentConversionForStorageAccountsTableService_STATUS); ok {
		err := augmentedService.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccountsTableService_STATUS populates the provided destination StorageAccountsTableService_STATUS from our StorageAccountsTableService_STATUS
func (service *StorageAccountsTableService_STATUS) AssignProperties_To_StorageAccountsTableService_STATUS(destination *storage.StorageAccountsTableService_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(service.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(service.Conditions)

	// Cors
	if service.Cors != nil {
		var cor storage.CorsRules_STATUS
		err := service.Cors.AssignProperties_To_CorsRules_STATUS(&cor)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_CorsRules_STATUS() to populate field Cors")
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

	// Invoke the augmentConversionForStorageAccountsTableService_STATUS interface (if implemented) to customize the conversion
	var serviceAsAny any = service
	if augmentedService, ok := serviceAsAny.(augmentConversionForStorageAccountsTableService_STATUS); ok {
		err := augmentedService.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForStorageAccountsTableService_Spec interface {
	AssignPropertiesFrom(src *storage.StorageAccountsTableService_Spec) error
	AssignPropertiesTo(dst *storage.StorageAccountsTableService_Spec) error
}

type augmentConversionForStorageAccountsTableService_STATUS interface {
	AssignPropertiesFrom(src *storage.StorageAccountsTableService_STATUS) error
	AssignPropertiesTo(dst *storage.StorageAccountsTableService_STATUS) error
}

// Storage version of v1api20220901.StorageAccountsTableServiceOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type StorageAccountsTableServiceOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_StorageAccountsTableServiceOperatorSpec populates our StorageAccountsTableServiceOperatorSpec from the provided source StorageAccountsTableServiceOperatorSpec
func (operator *StorageAccountsTableServiceOperatorSpec) AssignProperties_From_StorageAccountsTableServiceOperatorSpec(source *storage.StorageAccountsTableServiceOperatorSpec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

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

	// Update the property bag
	if len(propertyBag) > 0 {
		operator.PropertyBag = propertyBag
	} else {
		operator.PropertyBag = nil
	}

	// Invoke the augmentConversionForStorageAccountsTableServiceOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForStorageAccountsTableServiceOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccountsTableServiceOperatorSpec populates the provided destination StorageAccountsTableServiceOperatorSpec from our StorageAccountsTableServiceOperatorSpec
func (operator *StorageAccountsTableServiceOperatorSpec) AssignProperties_To_StorageAccountsTableServiceOperatorSpec(destination *storage.StorageAccountsTableServiceOperatorSpec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(operator.PropertyBag)

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

	// Invoke the augmentConversionForStorageAccountsTableServiceOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForStorageAccountsTableServiceOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForStorageAccountsTableServiceOperatorSpec interface {
	AssignPropertiesFrom(src *storage.StorageAccountsTableServiceOperatorSpec) error
	AssignPropertiesTo(dst *storage.StorageAccountsTableServiceOperatorSpec) error
}

func init() {
	SchemeBuilder.Register(&StorageAccountsTableService{}, &StorageAccountsTableServiceList{})
}

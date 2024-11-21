// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901/storage"
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
// Storage version of v1api20210401.StorageAccountsQueueServicesQueue
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2021-04-01/queue.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues/{queueName}
type StorageAccountsQueueServicesQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccountsQueueServicesQueue_Spec   `json:"spec,omitempty"`
	Status            StorageAccountsQueueServicesQueue_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsQueueServicesQueue{}

// GetConditions returns the conditions of the resource
func (queue *StorageAccountsQueueServicesQueue) GetConditions() conditions.Conditions {
	return queue.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (queue *StorageAccountsQueueServicesQueue) SetConditions(conditions conditions.Conditions) {
	queue.Status.Conditions = conditions
}

var _ conversion.Convertible = &StorageAccountsQueueServicesQueue{}

// ConvertFrom populates our StorageAccountsQueueServicesQueue from the provided hub StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source storage.StorageAccountsQueueServicesQueue

	err := source.ConvertFrom(hub)
	if err != nil {
		return eris.Wrap(err, "converting from hub to source")
	}

	err = queue.AssignProperties_From_StorageAccountsQueueServicesQueue(&source)
	if err != nil {
		return eris.Wrap(err, "converting from source to queue")
	}

	return nil
}

// ConvertTo populates the provided hub StorageAccountsQueueServicesQueue from our StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination storage.StorageAccountsQueueServicesQueue
	err := queue.AssignProperties_To_StorageAccountsQueueServicesQueue(&destination)
	if err != nil {
		return eris.Wrap(err, "converting to destination from queue")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return eris.Wrap(err, "converting from destination to hub")
	}

	return nil
}

var _ configmaps.Exporter = &StorageAccountsQueueServicesQueue{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (queue *StorageAccountsQueueServicesQueue) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if queue.Spec.OperatorSpec == nil {
		return nil
	}
	return queue.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &StorageAccountsQueueServicesQueue{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (queue *StorageAccountsQueueServicesQueue) SecretDestinationExpressions() []*core.DestinationExpression {
	if queue.Spec.OperatorSpec == nil {
		return nil
	}
	return queue.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &StorageAccountsQueueServicesQueue{}

// AzureName returns the Azure name of the resource
func (queue *StorageAccountsQueueServicesQueue) AzureName() string {
	return queue.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (queue StorageAccountsQueueServicesQueue) GetAPIVersion() string {
	return "2021-04-01"
}

// GetResourceScope returns the scope of the resource
func (queue *StorageAccountsQueueServicesQueue) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (queue *StorageAccountsQueueServicesQueue) GetSpec() genruntime.ConvertibleSpec {
	return &queue.Spec
}

// GetStatus returns the status of this resource
func (queue *StorageAccountsQueueServicesQueue) GetStatus() genruntime.ConvertibleStatus {
	return &queue.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (queue *StorageAccountsQueueServicesQueue) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/queueServices/queues"
func (queue *StorageAccountsQueueServicesQueue) GetType() string {
	return "Microsoft.Storage/storageAccounts/queueServices/queues"
}

// NewEmptyStatus returns a new empty (blank) status
func (queue *StorageAccountsQueueServicesQueue) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &StorageAccountsQueueServicesQueue_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (queue *StorageAccountsQueueServicesQueue) Owner() *genruntime.ResourceReference {
	if queue.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(queue.Spec)
	return queue.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (queue *StorageAccountsQueueServicesQueue) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*StorageAccountsQueueServicesQueue_STATUS); ok {
		queue.Status = *st
		return nil
	}

	// Convert status to required version
	var st StorageAccountsQueueServicesQueue_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	queue.Status = st
	return nil
}

// AssignProperties_From_StorageAccountsQueueServicesQueue populates our StorageAccountsQueueServicesQueue from the provided source StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) AssignProperties_From_StorageAccountsQueueServicesQueue(source *storage.StorageAccountsQueueServicesQueue) error {

	// ObjectMeta
	queue.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec StorageAccountsQueueServicesQueue_Spec
	err := spec.AssignProperties_From_StorageAccountsQueueServicesQueue_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_StorageAccountsQueueServicesQueue_Spec() to populate field Spec")
	}
	queue.Spec = spec

	// Status
	var status StorageAccountsQueueServicesQueue_STATUS
	err = status.AssignProperties_From_StorageAccountsQueueServicesQueue_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_StorageAccountsQueueServicesQueue_STATUS() to populate field Status")
	}
	queue.Status = status

	// Invoke the augmentConversionForStorageAccountsQueueServicesQueue interface (if implemented) to customize the conversion
	var queueAsAny any = queue
	if augmentedQueue, ok := queueAsAny.(augmentConversionForStorageAccountsQueueServicesQueue); ok {
		err := augmentedQueue.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccountsQueueServicesQueue populates the provided destination StorageAccountsQueueServicesQueue from our StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) AssignProperties_To_StorageAccountsQueueServicesQueue(destination *storage.StorageAccountsQueueServicesQueue) error {

	// ObjectMeta
	destination.ObjectMeta = *queue.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.StorageAccountsQueueServicesQueue_Spec
	err := queue.Spec.AssignProperties_To_StorageAccountsQueueServicesQueue_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_StorageAccountsQueueServicesQueue_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.StorageAccountsQueueServicesQueue_STATUS
	err = queue.Status.AssignProperties_To_StorageAccountsQueueServicesQueue_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_StorageAccountsQueueServicesQueue_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForStorageAccountsQueueServicesQueue interface (if implemented) to customize the conversion
	var queueAsAny any = queue
	if augmentedQueue, ok := queueAsAny.(augmentConversionForStorageAccountsQueueServicesQueue); ok {
		err := augmentedQueue.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (queue *StorageAccountsQueueServicesQueue) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: queue.Spec.OriginalVersion,
		Kind:    "StorageAccountsQueueServicesQueue",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210401.StorageAccountsQueueServicesQueue
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2021-04-01/queue.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues/{queueName}
type StorageAccountsQueueServicesQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsQueueServicesQueue `json:"items"`
}

type augmentConversionForStorageAccountsQueueServicesQueue interface {
	AssignPropertiesFrom(src *storage.StorageAccountsQueueServicesQueue) error
	AssignPropertiesTo(dst *storage.StorageAccountsQueueServicesQueue) error
}

// Storage version of v1api20210401.StorageAccountsQueueServicesQueue_Spec
type StorageAccountsQueueServicesQueue_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                                         `json:"azureName,omitempty"`
	Metadata        map[string]string                              `json:"metadata,omitempty"`
	OperatorSpec    *StorageAccountsQueueServicesQueueOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                                         `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a storage.azure.com/StorageAccountsQueueService resource
	Owner       *genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner,omitempty" kind:"StorageAccountsQueueService"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &StorageAccountsQueueServicesQueue_Spec{}

// ConvertSpecFrom populates our StorageAccountsQueueServicesQueue_Spec from the provided source
func (queue *StorageAccountsQueueServicesQueue_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.StorageAccountsQueueServicesQueue_Spec)
	if ok {
		// Populate our instance from source
		return queue.AssignProperties_From_StorageAccountsQueueServicesQueue_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.StorageAccountsQueueServicesQueue_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = queue.AssignProperties_From_StorageAccountsQueueServicesQueue_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our StorageAccountsQueueServicesQueue_Spec
func (queue *StorageAccountsQueueServicesQueue_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.StorageAccountsQueueServicesQueue_Spec)
	if ok {
		// Populate destination from our instance
		return queue.AssignProperties_To_StorageAccountsQueueServicesQueue_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.StorageAccountsQueueServicesQueue_Spec{}
	err := queue.AssignProperties_To_StorageAccountsQueueServicesQueue_Spec(dst)
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

// AssignProperties_From_StorageAccountsQueueServicesQueue_Spec populates our StorageAccountsQueueServicesQueue_Spec from the provided source StorageAccountsQueueServicesQueue_Spec
func (queue *StorageAccountsQueueServicesQueue_Spec) AssignProperties_From_StorageAccountsQueueServicesQueue_Spec(source *storage.StorageAccountsQueueServicesQueue_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	queue.AzureName = source.AzureName

	// Metadata
	queue.Metadata = genruntime.CloneMapOfStringToString(source.Metadata)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec StorageAccountsQueueServicesQueueOperatorSpec
		err := operatorSpec.AssignProperties_From_StorageAccountsQueueServicesQueueOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_StorageAccountsQueueServicesQueueOperatorSpec() to populate field OperatorSpec")
		}
		queue.OperatorSpec = &operatorSpec
	} else {
		queue.OperatorSpec = nil
	}

	// OriginalVersion
	queue.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		queue.Owner = &owner
	} else {
		queue.Owner = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		queue.PropertyBag = propertyBag
	} else {
		queue.PropertyBag = nil
	}

	// Invoke the augmentConversionForStorageAccountsQueueServicesQueue_Spec interface (if implemented) to customize the conversion
	var queueAsAny any = queue
	if augmentedQueue, ok := queueAsAny.(augmentConversionForStorageAccountsQueueServicesQueue_Spec); ok {
		err := augmentedQueue.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccountsQueueServicesQueue_Spec populates the provided destination StorageAccountsQueueServicesQueue_Spec from our StorageAccountsQueueServicesQueue_Spec
func (queue *StorageAccountsQueueServicesQueue_Spec) AssignProperties_To_StorageAccountsQueueServicesQueue_Spec(destination *storage.StorageAccountsQueueServicesQueue_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(queue.PropertyBag)

	// AzureName
	destination.AzureName = queue.AzureName

	// Metadata
	destination.Metadata = genruntime.CloneMapOfStringToString(queue.Metadata)

	// OperatorSpec
	if queue.OperatorSpec != nil {
		var operatorSpec storage.StorageAccountsQueueServicesQueueOperatorSpec
		err := queue.OperatorSpec.AssignProperties_To_StorageAccountsQueueServicesQueueOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_StorageAccountsQueueServicesQueueOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = queue.OriginalVersion

	// Owner
	if queue.Owner != nil {
		owner := queue.Owner.Copy()
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

	// Invoke the augmentConversionForStorageAccountsQueueServicesQueue_Spec interface (if implemented) to customize the conversion
	var queueAsAny any = queue
	if augmentedQueue, ok := queueAsAny.(augmentConversionForStorageAccountsQueueServicesQueue_Spec); ok {
		err := augmentedQueue.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20210401.StorageAccountsQueueServicesQueue_STATUS
type StorageAccountsQueueServicesQueue_STATUS struct {
	ApproximateMessageCount *int                   `json:"approximateMessageCount,omitempty"`
	Conditions              []conditions.Condition `json:"conditions,omitempty"`
	Id                      *string                `json:"id,omitempty"`
	Metadata                map[string]string      `json:"metadata,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type                    *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageAccountsQueueServicesQueue_STATUS{}

// ConvertStatusFrom populates our StorageAccountsQueueServicesQueue_STATUS from the provided source
func (queue *StorageAccountsQueueServicesQueue_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.StorageAccountsQueueServicesQueue_STATUS)
	if ok {
		// Populate our instance from source
		return queue.AssignProperties_From_StorageAccountsQueueServicesQueue_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.StorageAccountsQueueServicesQueue_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = queue.AssignProperties_From_StorageAccountsQueueServicesQueue_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our StorageAccountsQueueServicesQueue_STATUS
func (queue *StorageAccountsQueueServicesQueue_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.StorageAccountsQueueServicesQueue_STATUS)
	if ok {
		// Populate destination from our instance
		return queue.AssignProperties_To_StorageAccountsQueueServicesQueue_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.StorageAccountsQueueServicesQueue_STATUS{}
	err := queue.AssignProperties_To_StorageAccountsQueueServicesQueue_STATUS(dst)
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

// AssignProperties_From_StorageAccountsQueueServicesQueue_STATUS populates our StorageAccountsQueueServicesQueue_STATUS from the provided source StorageAccountsQueueServicesQueue_STATUS
func (queue *StorageAccountsQueueServicesQueue_STATUS) AssignProperties_From_StorageAccountsQueueServicesQueue_STATUS(source *storage.StorageAccountsQueueServicesQueue_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ApproximateMessageCount
	queue.ApproximateMessageCount = genruntime.ClonePointerToInt(source.ApproximateMessageCount)

	// Conditions
	queue.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	queue.Id = genruntime.ClonePointerToString(source.Id)

	// Metadata
	queue.Metadata = genruntime.CloneMapOfStringToString(source.Metadata)

	// Name
	queue.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	queue.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		queue.PropertyBag = propertyBag
	} else {
		queue.PropertyBag = nil
	}

	// Invoke the augmentConversionForStorageAccountsQueueServicesQueue_STATUS interface (if implemented) to customize the conversion
	var queueAsAny any = queue
	if augmentedQueue, ok := queueAsAny.(augmentConversionForStorageAccountsQueueServicesQueue_STATUS); ok {
		err := augmentedQueue.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccountsQueueServicesQueue_STATUS populates the provided destination StorageAccountsQueueServicesQueue_STATUS from our StorageAccountsQueueServicesQueue_STATUS
func (queue *StorageAccountsQueueServicesQueue_STATUS) AssignProperties_To_StorageAccountsQueueServicesQueue_STATUS(destination *storage.StorageAccountsQueueServicesQueue_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(queue.PropertyBag)

	// ApproximateMessageCount
	destination.ApproximateMessageCount = genruntime.ClonePointerToInt(queue.ApproximateMessageCount)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(queue.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(queue.Id)

	// Metadata
	destination.Metadata = genruntime.CloneMapOfStringToString(queue.Metadata)

	// Name
	destination.Name = genruntime.ClonePointerToString(queue.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(queue.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForStorageAccountsQueueServicesQueue_STATUS interface (if implemented) to customize the conversion
	var queueAsAny any = queue
	if augmentedQueue, ok := queueAsAny.(augmentConversionForStorageAccountsQueueServicesQueue_STATUS); ok {
		err := augmentedQueue.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForStorageAccountsQueueServicesQueue_Spec interface {
	AssignPropertiesFrom(src *storage.StorageAccountsQueueServicesQueue_Spec) error
	AssignPropertiesTo(dst *storage.StorageAccountsQueueServicesQueue_Spec) error
}

type augmentConversionForStorageAccountsQueueServicesQueue_STATUS interface {
	AssignPropertiesFrom(src *storage.StorageAccountsQueueServicesQueue_STATUS) error
	AssignPropertiesTo(dst *storage.StorageAccountsQueueServicesQueue_STATUS) error
}

// Storage version of v1api20210401.StorageAccountsQueueServicesQueueOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type StorageAccountsQueueServicesQueueOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_StorageAccountsQueueServicesQueueOperatorSpec populates our StorageAccountsQueueServicesQueueOperatorSpec from the provided source StorageAccountsQueueServicesQueueOperatorSpec
func (operator *StorageAccountsQueueServicesQueueOperatorSpec) AssignProperties_From_StorageAccountsQueueServicesQueueOperatorSpec(source *storage.StorageAccountsQueueServicesQueueOperatorSpec) error {
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

	// Invoke the augmentConversionForStorageAccountsQueueServicesQueueOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForStorageAccountsQueueServicesQueueOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccountsQueueServicesQueueOperatorSpec populates the provided destination StorageAccountsQueueServicesQueueOperatorSpec from our StorageAccountsQueueServicesQueueOperatorSpec
func (operator *StorageAccountsQueueServicesQueueOperatorSpec) AssignProperties_To_StorageAccountsQueueServicesQueueOperatorSpec(destination *storage.StorageAccountsQueueServicesQueueOperatorSpec) error {
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

	// Invoke the augmentConversionForStorageAccountsQueueServicesQueueOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForStorageAccountsQueueServicesQueueOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForStorageAccountsQueueServicesQueueOperatorSpec interface {
	AssignPropertiesFrom(src *storage.StorageAccountsQueueServicesQueueOperatorSpec) error
	AssignPropertiesTo(dst *storage.StorageAccountsQueueServicesQueueOperatorSpec) error
}

func init() {
	SchemeBuilder.Register(&StorageAccountsQueueServicesQueue{}, &StorageAccountsQueueServicesQueueList{})
}

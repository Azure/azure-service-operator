// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

import (
	"fmt"
	v20210401s "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401storage"
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
	source, ok := hub.(*v20210401s.StorageAccountsQueueServicesQueue)
	if !ok {
		return fmt.Errorf("expected storage/v1beta20210401storage/StorageAccountsQueueServicesQueue but received %T instead", hub)
	}

	return queue.AssignPropertiesFromStorageAccountsQueueServicesQueue(source)
}

// ConvertTo populates the provided hub StorageAccountsQueueServicesQueue from our StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210401s.StorageAccountsQueueServicesQueue)
	if !ok {
		return fmt.Errorf("expected storage/v1beta20210401storage/StorageAccountsQueueServicesQueue but received %T instead", hub)
	}

	return queue.AssignPropertiesToStorageAccountsQueueServicesQueue(destination)
}

// +kubebuilder:webhook:path=/mutate-storage-azure-com-v1beta20210401-storageaccountsqueueservicesqueue,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountsqueueservicesqueues,verbs=create;update,versions=v1beta20210401,name=default.v1beta20210401.storageaccountsqueueservicesqueues.storage.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &StorageAccountsQueueServicesQueue{}

// Default applies defaults to the StorageAccountsQueueServicesQueue resource
func (queue *StorageAccountsQueueServicesQueue) Default() {
	queue.defaultImpl()
	var temp interface{} = queue
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (queue *StorageAccountsQueueServicesQueue) defaultAzureName() {
	if queue.Spec.AzureName == "" {
		queue.Spec.AzureName = queue.Name
	}
}

// defaultImpl applies the code generated defaults to the StorageAccountsQueueServicesQueue resource
func (queue *StorageAccountsQueueServicesQueue) defaultImpl() { queue.defaultAzureName() }

var _ genruntime.KubernetesResource = &StorageAccountsQueueServicesQueue{}

// AzureName returns the Azure name of the resource
func (queue *StorageAccountsQueueServicesQueue) AzureName() string {
	return queue.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (queue StorageAccountsQueueServicesQueue) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceKind returns the kind of the resource
func (queue *StorageAccountsQueueServicesQueue) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (queue *StorageAccountsQueueServicesQueue) GetSpec() genruntime.ConvertibleSpec {
	return &queue.Spec
}

// GetStatus returns the status of this resource
func (queue *StorageAccountsQueueServicesQueue) GetStatus() genruntime.ConvertibleStatus {
	return &queue.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/queueServices/queues"
func (queue *StorageAccountsQueueServicesQueue) GetType() string {
	return "Microsoft.Storage/storageAccounts/queueServices/queues"
}

// NewEmptyStatus returns a new empty (blank) status
func (queue *StorageAccountsQueueServicesQueue) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &StorageAccountsQueueServicesQueue_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (queue *StorageAccountsQueueServicesQueue) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(queue.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  queue.Spec.Owner.Name,
	}
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
		return errors.Wrap(err, "failed to convert status")
	}

	queue.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-storage-azure-com-v1beta20210401-storageaccountsqueueservicesqueue,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountsqueueservicesqueues,verbs=create;update,versions=v1beta20210401,name=validate.v1beta20210401.storageaccountsqueueservicesqueues.storage.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &StorageAccountsQueueServicesQueue{}

// ValidateCreate validates the creation of the resource
func (queue *StorageAccountsQueueServicesQueue) ValidateCreate() error {
	validations := queue.createValidations()
	var temp interface{} = queue
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
func (queue *StorageAccountsQueueServicesQueue) ValidateDelete() error {
	validations := queue.deleteValidations()
	var temp interface{} = queue
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
func (queue *StorageAccountsQueueServicesQueue) ValidateUpdate(old runtime.Object) error {
	validations := queue.updateValidations()
	var temp interface{} = queue
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
func (queue *StorageAccountsQueueServicesQueue) createValidations() []func() error {
	return []func() error{queue.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (queue *StorageAccountsQueueServicesQueue) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (queue *StorageAccountsQueueServicesQueue) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return queue.validateResourceReferences()
		},
		queue.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (queue *StorageAccountsQueueServicesQueue) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&queue.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (queue *StorageAccountsQueueServicesQueue) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*StorageAccountsQueueServicesQueue)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, queue)
}

// AssignPropertiesFromStorageAccountsQueueServicesQueue populates our StorageAccountsQueueServicesQueue from the provided source StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) AssignPropertiesFromStorageAccountsQueueServicesQueue(source *v20210401s.StorageAccountsQueueServicesQueue) error {

	// ObjectMeta
	queue.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec StorageAccountsQueueServicesQueue_Spec
	err := spec.AssignPropertiesFromStorageAccountsQueueServicesQueue_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromStorageAccountsQueueServicesQueue_Spec() to populate field Spec")
	}
	queue.Spec = spec

	// Status
	var status StorageAccountsQueueServicesQueue_STATUS
	err = status.AssignPropertiesFromStorageAccountsQueueServicesQueue_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromStorageAccountsQueueServicesQueue_STATUS() to populate field Status")
	}
	queue.Status = status

	// No error
	return nil
}

// AssignPropertiesToStorageAccountsQueueServicesQueue populates the provided destination StorageAccountsQueueServicesQueue from our StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) AssignPropertiesToStorageAccountsQueueServicesQueue(destination *v20210401s.StorageAccountsQueueServicesQueue) error {

	// ObjectMeta
	destination.ObjectMeta = *queue.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210401s.StorageAccountsQueueServicesQueue_Spec
	err := queue.Spec.AssignPropertiesToStorageAccountsQueueServicesQueue_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToStorageAccountsQueueServicesQueue_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210401s.StorageAccountsQueueServicesQueue_STATUS
	err = queue.Status.AssignPropertiesToStorageAccountsQueueServicesQueue_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToStorageAccountsQueueServicesQueue_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (queue *StorageAccountsQueueServicesQueue) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: queue.Spec.OriginalVersion(),
		Kind:    "StorageAccountsQueueServicesQueue",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2021-04-01/queue.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues/{queueName}
type StorageAccountsQueueServicesQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsQueueServicesQueue `json:"items"`
}

type StorageAccountsQueueServicesQueue_STATUS struct {
	// ApproximateMessageCount: Integer indicating an approximate number of messages in the queue. This number is not lower
	// than the actual number of messages in the queue, but could be higher.
	ApproximateMessageCount *int `json:"approximateMessageCount,omitempty"`

	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Metadata: A name-value pair that represents queue metadata.
	Metadata map[string]string `json:"metadata,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageAccountsQueueServicesQueue_STATUS{}

// ConvertStatusFrom populates our StorageAccountsQueueServicesQueue_STATUS from the provided source
func (queue *StorageAccountsQueueServicesQueue_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210401s.StorageAccountsQueueServicesQueue_STATUS)
	if ok {
		// Populate our instance from source
		return queue.AssignPropertiesFromStorageAccountsQueueServicesQueue_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210401s.StorageAccountsQueueServicesQueue_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = queue.AssignPropertiesFromStorageAccountsQueueServicesQueue_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our StorageAccountsQueueServicesQueue_STATUS
func (queue *StorageAccountsQueueServicesQueue_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210401s.StorageAccountsQueueServicesQueue_STATUS)
	if ok {
		// Populate destination from our instance
		return queue.AssignPropertiesToStorageAccountsQueueServicesQueue_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210401s.StorageAccountsQueueServicesQueue_STATUS{}
	err := queue.AssignPropertiesToStorageAccountsQueueServicesQueue_STATUS(dst)
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

var _ genruntime.FromARMConverter = &StorageAccountsQueueServicesQueue_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (queue *StorageAccountsQueueServicesQueue_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &StorageAccountsQueueServicesQueue_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (queue *StorageAccountsQueueServicesQueue_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(StorageAccountsQueueServicesQueue_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected StorageAccountsQueueServicesQueue_STATUSARM, got %T", armInput)
	}

	// Set property ‘ApproximateMessageCount’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ApproximateMessageCount != nil {
			approximateMessageCount := *typedInput.Properties.ApproximateMessageCount
			queue.ApproximateMessageCount = &approximateMessageCount
		}
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		queue.Id = &id
	}

	// Set property ‘Metadata’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Metadata != nil {
			queue.Metadata = make(map[string]string)
			for key, value := range typedInput.Properties.Metadata {
				queue.Metadata[key] = value
			}
		}
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		queue.Name = &name
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		queue.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromStorageAccountsQueueServicesQueue_STATUS populates our StorageAccountsQueueServicesQueue_STATUS from the provided source StorageAccountsQueueServicesQueue_STATUS
func (queue *StorageAccountsQueueServicesQueue_STATUS) AssignPropertiesFromStorageAccountsQueueServicesQueue_STATUS(source *v20210401s.StorageAccountsQueueServicesQueue_STATUS) error {

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

	// No error
	return nil
}

// AssignPropertiesToStorageAccountsQueueServicesQueue_STATUS populates the provided destination StorageAccountsQueueServicesQueue_STATUS from our StorageAccountsQueueServicesQueue_STATUS
func (queue *StorageAccountsQueueServicesQueue_STATUS) AssignPropertiesToStorageAccountsQueueServicesQueue_STATUS(destination *v20210401s.StorageAccountsQueueServicesQueue_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

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

	// No error
	return nil
}

type StorageAccountsQueueServicesQueue_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// Metadata: A name-value pair that represents queue metadata.
	Metadata map[string]string `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a storage.azure.com/StorageAccountsQueueService resource
	Owner *genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner,omitempty" kind:"StorageAccountsQueueService"`
}

var _ genruntime.ARMTransformer = &StorageAccountsQueueServicesQueue_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (queue *StorageAccountsQueueServicesQueue_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if queue == nil {
		return nil, nil
	}
	result := &StorageAccountsQueueServicesQueue_SpecARM{}

	// Set property ‘AzureName’:
	result.AzureName = queue.AzureName

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if queue.Metadata != nil {
		result.Properties = &QueuePropertiesARM{}
	}
	if queue.Metadata != nil {
		result.Properties.Metadata = make(map[string]string)
		for key, value := range queue.Metadata {
			result.Properties.Metadata[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (queue *StorageAccountsQueueServicesQueue_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &StorageAccountsQueueServicesQueue_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (queue *StorageAccountsQueueServicesQueue_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(StorageAccountsQueueServicesQueue_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected StorageAccountsQueueServicesQueue_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	queue.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Metadata’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Metadata != nil {
			queue.Metadata = make(map[string]string)
			for key, value := range typedInput.Properties.Metadata {
				queue.Metadata[key] = value
			}
		}
	}

	// Set property ‘Owner’:
	queue.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &StorageAccountsQueueServicesQueue_Spec{}

// ConvertSpecFrom populates our StorageAccountsQueueServicesQueue_Spec from the provided source
func (queue *StorageAccountsQueueServicesQueue_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210401s.StorageAccountsQueueServicesQueue_Spec)
	if ok {
		// Populate our instance from source
		return queue.AssignPropertiesFromStorageAccountsQueueServicesQueue_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210401s.StorageAccountsQueueServicesQueue_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = queue.AssignPropertiesFromStorageAccountsQueueServicesQueue_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our StorageAccountsQueueServicesQueue_Spec
func (queue *StorageAccountsQueueServicesQueue_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210401s.StorageAccountsQueueServicesQueue_Spec)
	if ok {
		// Populate destination from our instance
		return queue.AssignPropertiesToStorageAccountsQueueServicesQueue_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210401s.StorageAccountsQueueServicesQueue_Spec{}
	err := queue.AssignPropertiesToStorageAccountsQueueServicesQueue_Spec(dst)
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

// AssignPropertiesFromStorageAccountsQueueServicesQueue_Spec populates our StorageAccountsQueueServicesQueue_Spec from the provided source StorageAccountsQueueServicesQueue_Spec
func (queue *StorageAccountsQueueServicesQueue_Spec) AssignPropertiesFromStorageAccountsQueueServicesQueue_Spec(source *v20210401s.StorageAccountsQueueServicesQueue_Spec) error {

	// AzureName
	queue.AzureName = source.AzureName

	// Metadata
	queue.Metadata = genruntime.CloneMapOfStringToString(source.Metadata)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		queue.Owner = &owner
	} else {
		queue.Owner = nil
	}

	// No error
	return nil
}

// AssignPropertiesToStorageAccountsQueueServicesQueue_Spec populates the provided destination StorageAccountsQueueServicesQueue_Spec from our StorageAccountsQueueServicesQueue_Spec
func (queue *StorageAccountsQueueServicesQueue_Spec) AssignPropertiesToStorageAccountsQueueServicesQueue_Spec(destination *v20210401s.StorageAccountsQueueServicesQueue_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = queue.AzureName

	// Metadata
	destination.Metadata = genruntime.CloneMapOfStringToString(queue.Metadata)

	// OriginalVersion
	destination.OriginalVersion = queue.OriginalVersion()

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

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (queue *StorageAccountsQueueServicesQueue_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (queue *StorageAccountsQueueServicesQueue_Spec) SetAzureName(azureName string) {
	queue.AzureName = azureName
}

func init() {
	SchemeBuilder.Register(&StorageAccountsQueueServicesQueue{}, &StorageAccountsQueueServicesQueueList{})
}

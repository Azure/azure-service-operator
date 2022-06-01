// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=storage.azure.com,resources=storageaccountsqueueservicesqueues,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=storage.azure.com,resources={storageaccountsqueueservicesqueues/status,storageaccountsqueueservicesqueues/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210401.StorageAccountsQueueServicesQueue
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

// Hub marks that this StorageAccountsQueueServicesQueue is the hub type for conversion
func (queue *StorageAccountsQueueServicesQueue) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (queue *StorageAccountsQueueServicesQueue) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: queue.Spec.OriginalVersion,
		Kind:    "StorageAccountsQueueServicesQueue",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210401.StorageAccountsQueueServicesQueue
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2021-04-01/queue.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues/{queueName}
type StorageAccountsQueueServicesQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsQueueServicesQueue `json:"items"`
}

// Storage version of v1beta20210401.StorageAccountsQueueServicesQueue_STATUS
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
	if source == queue {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(queue)
}

// ConvertStatusTo populates the provided destination from our StorageAccountsQueueServicesQueue_STATUS
func (queue *StorageAccountsQueueServicesQueue_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == queue {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(queue)
}

// Storage version of v1beta20210401.StorageAccountsQueueServicesQueue_Spec
type StorageAccountsQueueServicesQueue_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string            `json:"azureName,omitempty"`
	Metadata        map[string]string `json:"metadata,omitempty"`
	OriginalVersion string            `json:"originalVersion,omitempty"`

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
	if source == queue {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(queue)
}

// ConvertSpecTo populates the provided destination from our StorageAccountsQueueServicesQueue_Spec
func (queue *StorageAccountsQueueServicesQueue_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == queue {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(queue)
}

func init() {
	SchemeBuilder.Register(&StorageAccountsQueueServicesQueue{}, &StorageAccountsQueueServicesQueueList{})
}

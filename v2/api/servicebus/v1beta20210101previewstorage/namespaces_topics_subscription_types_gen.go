// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101previewstorage

import (
	v20210101ps "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20210101previewstorage"
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
// Storage version of v1beta20210101preview.NamespacesTopicsSubscription
// Deprecated version of NamespacesTopicsSubscription. Use v1api20210101preview.NamespacesTopicsSubscription instead
type NamespacesTopicsSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Namespaces_Topics_Subscription_Spec   `json:"spec,omitempty"`
	Status            Namespaces_Topics_Subscription_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesTopicsSubscription{}

// GetConditions returns the conditions of the resource
func (subscription *NamespacesTopicsSubscription) GetConditions() conditions.Conditions {
	return subscription.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (subscription *NamespacesTopicsSubscription) SetConditions(conditions conditions.Conditions) {
	subscription.Status.Conditions = conditions
}

var _ conversion.Convertible = &NamespacesTopicsSubscription{}

// ConvertFrom populates our NamespacesTopicsSubscription from the provided hub NamespacesTopicsSubscription
func (subscription *NamespacesTopicsSubscription) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source v20210101ps.NamespacesTopicsSubscription

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = subscription.AssignProperties_From_NamespacesTopicsSubscription(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to subscription")
	}

	return nil
}

// ConvertTo populates the provided hub NamespacesTopicsSubscription from our NamespacesTopicsSubscription
func (subscription *NamespacesTopicsSubscription) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination v20210101ps.NamespacesTopicsSubscription
	err := subscription.AssignProperties_To_NamespacesTopicsSubscription(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from subscription")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

var _ genruntime.KubernetesResource = &NamespacesTopicsSubscription{}

// AzureName returns the Azure name of the resource
func (subscription *NamespacesTopicsSubscription) AzureName() string {
	return subscription.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01-preview"
func (subscription NamespacesTopicsSubscription) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (subscription *NamespacesTopicsSubscription) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (subscription *NamespacesTopicsSubscription) GetSpec() genruntime.ConvertibleSpec {
	return &subscription.Spec
}

// GetStatus returns the status of this resource
func (subscription *NamespacesTopicsSubscription) GetStatus() genruntime.ConvertibleStatus {
	return &subscription.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (subscription *NamespacesTopicsSubscription) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/topics/subscriptions"
func (subscription *NamespacesTopicsSubscription) GetType() string {
	return "Microsoft.ServiceBus/namespaces/topics/subscriptions"
}

// NewEmptyStatus returns a new empty (blank) status
func (subscription *NamespacesTopicsSubscription) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Namespaces_Topics_Subscription_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (subscription *NamespacesTopicsSubscription) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(subscription.Spec)
	return subscription.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (subscription *NamespacesTopicsSubscription) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Namespaces_Topics_Subscription_STATUS); ok {
		subscription.Status = *st
		return nil
	}

	// Convert status to required version
	var st Namespaces_Topics_Subscription_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	subscription.Status = st
	return nil
}

// AssignProperties_From_NamespacesTopicsSubscription populates our NamespacesTopicsSubscription from the provided source NamespacesTopicsSubscription
func (subscription *NamespacesTopicsSubscription) AssignProperties_From_NamespacesTopicsSubscription(source *v20210101ps.NamespacesTopicsSubscription) error {

	// ObjectMeta
	subscription.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Namespaces_Topics_Subscription_Spec
	err := spec.AssignProperties_From_Namespaces_Topics_Subscription_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Namespaces_Topics_Subscription_Spec() to populate field Spec")
	}
	subscription.Spec = spec

	// Status
	var status Namespaces_Topics_Subscription_STATUS
	err = status.AssignProperties_From_Namespaces_Topics_Subscription_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Namespaces_Topics_Subscription_STATUS() to populate field Status")
	}
	subscription.Status = status

	// Invoke the augmentConversionForNamespacesTopicsSubscription interface (if implemented) to customize the conversion
	var subscriptionAsAny any = subscription
	if augmentedSubscription, ok := subscriptionAsAny.(augmentConversionForNamespacesTopicsSubscription); ok {
		err := augmentedSubscription.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_NamespacesTopicsSubscription populates the provided destination NamespacesTopicsSubscription from our NamespacesTopicsSubscription
func (subscription *NamespacesTopicsSubscription) AssignProperties_To_NamespacesTopicsSubscription(destination *v20210101ps.NamespacesTopicsSubscription) error {

	// ObjectMeta
	destination.ObjectMeta = *subscription.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210101ps.Namespaces_Topics_Subscription_Spec
	err := subscription.Spec.AssignProperties_To_Namespaces_Topics_Subscription_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Namespaces_Topics_Subscription_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210101ps.Namespaces_Topics_Subscription_STATUS
	err = subscription.Status.AssignProperties_To_Namespaces_Topics_Subscription_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Namespaces_Topics_Subscription_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForNamespacesTopicsSubscription interface (if implemented) to customize the conversion
	var subscriptionAsAny any = subscription
	if augmentedSubscription, ok := subscriptionAsAny.(augmentConversionForNamespacesTopicsSubscription); ok {
		err := augmentedSubscription.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (subscription *NamespacesTopicsSubscription) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: subscription.Spec.OriginalVersion,
		Kind:    "NamespacesTopicsSubscription",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210101preview.NamespacesTopicsSubscription
// Deprecated version of NamespacesTopicsSubscription. Use v1api20210101preview.NamespacesTopicsSubscription instead
type NamespacesTopicsSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesTopicsSubscription `json:"items"`
}

type augmentConversionForNamespacesTopicsSubscription interface {
	AssignPropertiesFrom(src *v20210101ps.NamespacesTopicsSubscription) error
	AssignPropertiesTo(dst *v20210101ps.NamespacesTopicsSubscription) error
}

// Storage version of v1beta20210101preview.Namespaces_Topics_Subscription_Spec
type Namespaces_Topics_Subscription_Spec struct {
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty"`

	// +kubebuilder:validation:MaxLength=50
	// +kubebuilder:validation:MinLength=1
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                                 string  `json:"azureName,omitempty"`
	DeadLetteringOnFilterEvaluationExceptions *bool   `json:"deadLetteringOnFilterEvaluationExceptions,omitempty"`
	DeadLetteringOnMessageExpiration          *bool   `json:"deadLetteringOnMessageExpiration,omitempty"`
	DefaultMessageTimeToLive                  *string `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow       *string `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations                   *bool   `json:"enableBatchedOperations,omitempty"`
	ForwardDeadLetteredMessagesTo             *string `json:"forwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                                 *string `json:"forwardTo,omitempty"`
	LockDuration                              *string `json:"lockDuration,omitempty"`
	MaxDeliveryCount                          *int    `json:"maxDeliveryCount,omitempty"`
	OriginalVersion                           string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a servicebus.azure.com/NamespacesTopic resource
	Owner           *genruntime.KnownResourceReference `group:"servicebus.azure.com" json:"owner,omitempty" kind:"NamespacesTopic"`
	PropertyBag     genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RequiresSession *bool                              `json:"requiresSession,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Namespaces_Topics_Subscription_Spec{}

// ConvertSpecFrom populates our Namespaces_Topics_Subscription_Spec from the provided source
func (subscription *Namespaces_Topics_Subscription_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210101ps.Namespaces_Topics_Subscription_Spec)
	if ok {
		// Populate our instance from source
		return subscription.AssignProperties_From_Namespaces_Topics_Subscription_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210101ps.Namespaces_Topics_Subscription_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = subscription.AssignProperties_From_Namespaces_Topics_Subscription_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Namespaces_Topics_Subscription_Spec
func (subscription *Namespaces_Topics_Subscription_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210101ps.Namespaces_Topics_Subscription_Spec)
	if ok {
		// Populate destination from our instance
		return subscription.AssignProperties_To_Namespaces_Topics_Subscription_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210101ps.Namespaces_Topics_Subscription_Spec{}
	err := subscription.AssignProperties_To_Namespaces_Topics_Subscription_Spec(dst)
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

// AssignProperties_From_Namespaces_Topics_Subscription_Spec populates our Namespaces_Topics_Subscription_Spec from the provided source Namespaces_Topics_Subscription_Spec
func (subscription *Namespaces_Topics_Subscription_Spec) AssignProperties_From_Namespaces_Topics_Subscription_Spec(source *v20210101ps.Namespaces_Topics_Subscription_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AutoDeleteOnIdle
	subscription.AutoDeleteOnIdle = genruntime.ClonePointerToString(source.AutoDeleteOnIdle)

	// AzureName
	subscription.AzureName = source.AzureName

	// DeadLetteringOnFilterEvaluationExceptions
	if source.DeadLetteringOnFilterEvaluationExceptions != nil {
		deadLetteringOnFilterEvaluationException := *source.DeadLetteringOnFilterEvaluationExceptions
		subscription.DeadLetteringOnFilterEvaluationExceptions = &deadLetteringOnFilterEvaluationException
	} else {
		subscription.DeadLetteringOnFilterEvaluationExceptions = nil
	}

	// DeadLetteringOnMessageExpiration
	if source.DeadLetteringOnMessageExpiration != nil {
		deadLetteringOnMessageExpiration := *source.DeadLetteringOnMessageExpiration
		subscription.DeadLetteringOnMessageExpiration = &deadLetteringOnMessageExpiration
	} else {
		subscription.DeadLetteringOnMessageExpiration = nil
	}

	// DefaultMessageTimeToLive
	subscription.DefaultMessageTimeToLive = genruntime.ClonePointerToString(source.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	subscription.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(source.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if source.EnableBatchedOperations != nil {
		enableBatchedOperation := *source.EnableBatchedOperations
		subscription.EnableBatchedOperations = &enableBatchedOperation
	} else {
		subscription.EnableBatchedOperations = nil
	}

	// ForwardDeadLetteredMessagesTo
	subscription.ForwardDeadLetteredMessagesTo = genruntime.ClonePointerToString(source.ForwardDeadLetteredMessagesTo)

	// ForwardTo
	subscription.ForwardTo = genruntime.ClonePointerToString(source.ForwardTo)

	// LockDuration
	subscription.LockDuration = genruntime.ClonePointerToString(source.LockDuration)

	// MaxDeliveryCount
	subscription.MaxDeliveryCount = genruntime.ClonePointerToInt(source.MaxDeliveryCount)

	// OriginalVersion
	subscription.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		subscription.Owner = &owner
	} else {
		subscription.Owner = nil
	}

	// RequiresSession
	if source.RequiresSession != nil {
		requiresSession := *source.RequiresSession
		subscription.RequiresSession = &requiresSession
	} else {
		subscription.RequiresSession = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		subscription.PropertyBag = propertyBag
	} else {
		subscription.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespaces_Topics_Subscription_Spec interface (if implemented) to customize the conversion
	var subscriptionAsAny any = subscription
	if augmentedSubscription, ok := subscriptionAsAny.(augmentConversionForNamespaces_Topics_Subscription_Spec); ok {
		err := augmentedSubscription.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Namespaces_Topics_Subscription_Spec populates the provided destination Namespaces_Topics_Subscription_Spec from our Namespaces_Topics_Subscription_Spec
func (subscription *Namespaces_Topics_Subscription_Spec) AssignProperties_To_Namespaces_Topics_Subscription_Spec(destination *v20210101ps.Namespaces_Topics_Subscription_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(subscription.PropertyBag)

	// AutoDeleteOnIdle
	destination.AutoDeleteOnIdle = genruntime.ClonePointerToString(subscription.AutoDeleteOnIdle)

	// AzureName
	destination.AzureName = subscription.AzureName

	// DeadLetteringOnFilterEvaluationExceptions
	if subscription.DeadLetteringOnFilterEvaluationExceptions != nil {
		deadLetteringOnFilterEvaluationException := *subscription.DeadLetteringOnFilterEvaluationExceptions
		destination.DeadLetteringOnFilterEvaluationExceptions = &deadLetteringOnFilterEvaluationException
	} else {
		destination.DeadLetteringOnFilterEvaluationExceptions = nil
	}

	// DeadLetteringOnMessageExpiration
	if subscription.DeadLetteringOnMessageExpiration != nil {
		deadLetteringOnMessageExpiration := *subscription.DeadLetteringOnMessageExpiration
		destination.DeadLetteringOnMessageExpiration = &deadLetteringOnMessageExpiration
	} else {
		destination.DeadLetteringOnMessageExpiration = nil
	}

	// DefaultMessageTimeToLive
	destination.DefaultMessageTimeToLive = genruntime.ClonePointerToString(subscription.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	destination.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(subscription.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if subscription.EnableBatchedOperations != nil {
		enableBatchedOperation := *subscription.EnableBatchedOperations
		destination.EnableBatchedOperations = &enableBatchedOperation
	} else {
		destination.EnableBatchedOperations = nil
	}

	// ForwardDeadLetteredMessagesTo
	destination.ForwardDeadLetteredMessagesTo = genruntime.ClonePointerToString(subscription.ForwardDeadLetteredMessagesTo)

	// ForwardTo
	destination.ForwardTo = genruntime.ClonePointerToString(subscription.ForwardTo)

	// LockDuration
	destination.LockDuration = genruntime.ClonePointerToString(subscription.LockDuration)

	// MaxDeliveryCount
	destination.MaxDeliveryCount = genruntime.ClonePointerToInt(subscription.MaxDeliveryCount)

	// OriginalVersion
	destination.OriginalVersion = subscription.OriginalVersion

	// Owner
	if subscription.Owner != nil {
		owner := subscription.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// RequiresSession
	if subscription.RequiresSession != nil {
		requiresSession := *subscription.RequiresSession
		destination.RequiresSession = &requiresSession
	} else {
		destination.RequiresSession = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespaces_Topics_Subscription_Spec interface (if implemented) to customize the conversion
	var subscriptionAsAny any = subscription
	if augmentedSubscription, ok := subscriptionAsAny.(augmentConversionForNamespaces_Topics_Subscription_Spec); ok {
		err := augmentedSubscription.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1beta20210101preview.Namespaces_Topics_Subscription_STATUS
// Deprecated version of Namespaces_Topics_Subscription_STATUS. Use v1api20210101preview.Namespaces_Topics_Subscription_STATUS instead
type Namespaces_Topics_Subscription_STATUS struct {
	AccessedAt                                *string                     `json:"accessedAt,omitempty"`
	AutoDeleteOnIdle                          *string                     `json:"autoDeleteOnIdle,omitempty"`
	Conditions                                []conditions.Condition      `json:"conditions,omitempty"`
	CountDetails                              *MessageCountDetails_STATUS `json:"countDetails,omitempty"`
	CreatedAt                                 *string                     `json:"createdAt,omitempty"`
	DeadLetteringOnFilterEvaluationExceptions *bool                       `json:"deadLetteringOnFilterEvaluationExceptions,omitempty"`
	DeadLetteringOnMessageExpiration          *bool                       `json:"deadLetteringOnMessageExpiration,omitempty"`
	DefaultMessageTimeToLive                  *string                     `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow       *string                     `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations                   *bool                       `json:"enableBatchedOperations,omitempty"`
	ForwardDeadLetteredMessagesTo             *string                     `json:"forwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                                 *string                     `json:"forwardTo,omitempty"`
	Id                                        *string                     `json:"id,omitempty"`
	LockDuration                              *string                     `json:"lockDuration,omitempty"`
	MaxDeliveryCount                          *int                        `json:"maxDeliveryCount,omitempty"`
	MessageCount                              *int                        `json:"messageCount,omitempty"`
	Name                                      *string                     `json:"name,omitempty"`
	PropertyBag                               genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	RequiresSession                           *bool                       `json:"requiresSession,omitempty"`
	Status                                    *string                     `json:"status,omitempty"`
	SystemData                                *SystemData_STATUS          `json:"systemData,omitempty"`
	Type                                      *string                     `json:"type,omitempty"`
	UpdatedAt                                 *string                     `json:"updatedAt,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Namespaces_Topics_Subscription_STATUS{}

// ConvertStatusFrom populates our Namespaces_Topics_Subscription_STATUS from the provided source
func (subscription *Namespaces_Topics_Subscription_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210101ps.Namespaces_Topics_Subscription_STATUS)
	if ok {
		// Populate our instance from source
		return subscription.AssignProperties_From_Namespaces_Topics_Subscription_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210101ps.Namespaces_Topics_Subscription_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = subscription.AssignProperties_From_Namespaces_Topics_Subscription_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Namespaces_Topics_Subscription_STATUS
func (subscription *Namespaces_Topics_Subscription_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210101ps.Namespaces_Topics_Subscription_STATUS)
	if ok {
		// Populate destination from our instance
		return subscription.AssignProperties_To_Namespaces_Topics_Subscription_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210101ps.Namespaces_Topics_Subscription_STATUS{}
	err := subscription.AssignProperties_To_Namespaces_Topics_Subscription_STATUS(dst)
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

// AssignProperties_From_Namespaces_Topics_Subscription_STATUS populates our Namespaces_Topics_Subscription_STATUS from the provided source Namespaces_Topics_Subscription_STATUS
func (subscription *Namespaces_Topics_Subscription_STATUS) AssignProperties_From_Namespaces_Topics_Subscription_STATUS(source *v20210101ps.Namespaces_Topics_Subscription_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AccessedAt
	subscription.AccessedAt = genruntime.ClonePointerToString(source.AccessedAt)

	// AutoDeleteOnIdle
	subscription.AutoDeleteOnIdle = genruntime.ClonePointerToString(source.AutoDeleteOnIdle)

	// Conditions
	subscription.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// CountDetails
	if source.CountDetails != nil {
		var countDetail MessageCountDetails_STATUS
		err := countDetail.AssignProperties_From_MessageCountDetails_STATUS(source.CountDetails)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_MessageCountDetails_STATUS() to populate field CountDetails")
		}
		subscription.CountDetails = &countDetail
	} else {
		subscription.CountDetails = nil
	}

	// CreatedAt
	subscription.CreatedAt = genruntime.ClonePointerToString(source.CreatedAt)

	// DeadLetteringOnFilterEvaluationExceptions
	if source.DeadLetteringOnFilterEvaluationExceptions != nil {
		deadLetteringOnFilterEvaluationException := *source.DeadLetteringOnFilterEvaluationExceptions
		subscription.DeadLetteringOnFilterEvaluationExceptions = &deadLetteringOnFilterEvaluationException
	} else {
		subscription.DeadLetteringOnFilterEvaluationExceptions = nil
	}

	// DeadLetteringOnMessageExpiration
	if source.DeadLetteringOnMessageExpiration != nil {
		deadLetteringOnMessageExpiration := *source.DeadLetteringOnMessageExpiration
		subscription.DeadLetteringOnMessageExpiration = &deadLetteringOnMessageExpiration
	} else {
		subscription.DeadLetteringOnMessageExpiration = nil
	}

	// DefaultMessageTimeToLive
	subscription.DefaultMessageTimeToLive = genruntime.ClonePointerToString(source.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	subscription.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(source.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if source.EnableBatchedOperations != nil {
		enableBatchedOperation := *source.EnableBatchedOperations
		subscription.EnableBatchedOperations = &enableBatchedOperation
	} else {
		subscription.EnableBatchedOperations = nil
	}

	// ForwardDeadLetteredMessagesTo
	subscription.ForwardDeadLetteredMessagesTo = genruntime.ClonePointerToString(source.ForwardDeadLetteredMessagesTo)

	// ForwardTo
	subscription.ForwardTo = genruntime.ClonePointerToString(source.ForwardTo)

	// Id
	subscription.Id = genruntime.ClonePointerToString(source.Id)

	// LockDuration
	subscription.LockDuration = genruntime.ClonePointerToString(source.LockDuration)

	// MaxDeliveryCount
	subscription.MaxDeliveryCount = genruntime.ClonePointerToInt(source.MaxDeliveryCount)

	// MessageCount
	subscription.MessageCount = genruntime.ClonePointerToInt(source.MessageCount)

	// Name
	subscription.Name = genruntime.ClonePointerToString(source.Name)

	// RequiresSession
	if source.RequiresSession != nil {
		requiresSession := *source.RequiresSession
		subscription.RequiresSession = &requiresSession
	} else {
		subscription.RequiresSession = nil
	}

	// Status
	subscription.Status = genruntime.ClonePointerToString(source.Status)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		subscription.SystemData = &systemDatum
	} else {
		subscription.SystemData = nil
	}

	// Type
	subscription.Type = genruntime.ClonePointerToString(source.Type)

	// UpdatedAt
	subscription.UpdatedAt = genruntime.ClonePointerToString(source.UpdatedAt)

	// Update the property bag
	if len(propertyBag) > 0 {
		subscription.PropertyBag = propertyBag
	} else {
		subscription.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespaces_Topics_Subscription_STATUS interface (if implemented) to customize the conversion
	var subscriptionAsAny any = subscription
	if augmentedSubscription, ok := subscriptionAsAny.(augmentConversionForNamespaces_Topics_Subscription_STATUS); ok {
		err := augmentedSubscription.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Namespaces_Topics_Subscription_STATUS populates the provided destination Namespaces_Topics_Subscription_STATUS from our Namespaces_Topics_Subscription_STATUS
func (subscription *Namespaces_Topics_Subscription_STATUS) AssignProperties_To_Namespaces_Topics_Subscription_STATUS(destination *v20210101ps.Namespaces_Topics_Subscription_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(subscription.PropertyBag)

	// AccessedAt
	destination.AccessedAt = genruntime.ClonePointerToString(subscription.AccessedAt)

	// AutoDeleteOnIdle
	destination.AutoDeleteOnIdle = genruntime.ClonePointerToString(subscription.AutoDeleteOnIdle)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(subscription.Conditions)

	// CountDetails
	if subscription.CountDetails != nil {
		var countDetail v20210101ps.MessageCountDetails_STATUS
		err := subscription.CountDetails.AssignProperties_To_MessageCountDetails_STATUS(&countDetail)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_MessageCountDetails_STATUS() to populate field CountDetails")
		}
		destination.CountDetails = &countDetail
	} else {
		destination.CountDetails = nil
	}

	// CreatedAt
	destination.CreatedAt = genruntime.ClonePointerToString(subscription.CreatedAt)

	// DeadLetteringOnFilterEvaluationExceptions
	if subscription.DeadLetteringOnFilterEvaluationExceptions != nil {
		deadLetteringOnFilterEvaluationException := *subscription.DeadLetteringOnFilterEvaluationExceptions
		destination.DeadLetteringOnFilterEvaluationExceptions = &deadLetteringOnFilterEvaluationException
	} else {
		destination.DeadLetteringOnFilterEvaluationExceptions = nil
	}

	// DeadLetteringOnMessageExpiration
	if subscription.DeadLetteringOnMessageExpiration != nil {
		deadLetteringOnMessageExpiration := *subscription.DeadLetteringOnMessageExpiration
		destination.DeadLetteringOnMessageExpiration = &deadLetteringOnMessageExpiration
	} else {
		destination.DeadLetteringOnMessageExpiration = nil
	}

	// DefaultMessageTimeToLive
	destination.DefaultMessageTimeToLive = genruntime.ClonePointerToString(subscription.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	destination.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(subscription.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if subscription.EnableBatchedOperations != nil {
		enableBatchedOperation := *subscription.EnableBatchedOperations
		destination.EnableBatchedOperations = &enableBatchedOperation
	} else {
		destination.EnableBatchedOperations = nil
	}

	// ForwardDeadLetteredMessagesTo
	destination.ForwardDeadLetteredMessagesTo = genruntime.ClonePointerToString(subscription.ForwardDeadLetteredMessagesTo)

	// ForwardTo
	destination.ForwardTo = genruntime.ClonePointerToString(subscription.ForwardTo)

	// Id
	destination.Id = genruntime.ClonePointerToString(subscription.Id)

	// LockDuration
	destination.LockDuration = genruntime.ClonePointerToString(subscription.LockDuration)

	// MaxDeliveryCount
	destination.MaxDeliveryCount = genruntime.ClonePointerToInt(subscription.MaxDeliveryCount)

	// MessageCount
	destination.MessageCount = genruntime.ClonePointerToInt(subscription.MessageCount)

	// Name
	destination.Name = genruntime.ClonePointerToString(subscription.Name)

	// RequiresSession
	if subscription.RequiresSession != nil {
		requiresSession := *subscription.RequiresSession
		destination.RequiresSession = &requiresSession
	} else {
		destination.RequiresSession = nil
	}

	// Status
	destination.Status = genruntime.ClonePointerToString(subscription.Status)

	// SystemData
	if subscription.SystemData != nil {
		var systemDatum v20210101ps.SystemData_STATUS
		err := subscription.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(subscription.Type)

	// UpdatedAt
	destination.UpdatedAt = genruntime.ClonePointerToString(subscription.UpdatedAt)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespaces_Topics_Subscription_STATUS interface (if implemented) to customize the conversion
	var subscriptionAsAny any = subscription
	if augmentedSubscription, ok := subscriptionAsAny.(augmentConversionForNamespaces_Topics_Subscription_STATUS); ok {
		err := augmentedSubscription.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForNamespaces_Topics_Subscription_Spec interface {
	AssignPropertiesFrom(src *v20210101ps.Namespaces_Topics_Subscription_Spec) error
	AssignPropertiesTo(dst *v20210101ps.Namespaces_Topics_Subscription_Spec) error
}

type augmentConversionForNamespaces_Topics_Subscription_STATUS interface {
	AssignPropertiesFrom(src *v20210101ps.Namespaces_Topics_Subscription_STATUS) error
	AssignPropertiesTo(dst *v20210101ps.Namespaces_Topics_Subscription_STATUS) error
}

func init() {
	SchemeBuilder.Register(&NamespacesTopicsSubscription{}, &NamespacesTopicsSubscriptionList{})
}

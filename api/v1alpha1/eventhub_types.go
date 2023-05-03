// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EventhubSpec defines the desired state of Eventhub
type EventhubSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location   string             `json:"location"`
	Namespace  string             `json:"namespace,omitempty"`
	Properties EventhubProperties `json:"properties,omitempty"`
	// +kubebuilder:validation:Pattern=^[-\w\._\(\)]+$
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	ResourceGroup     string                    `json:"resourceGroup"`
	AuthorizationRule EventhubAuthorizationRule `json:"authorizationRule,omitempty"`
	// SecretName - Used to specify the name of the secret. Defaults to Event Hub name if omitted.
	SecretName             string `json:"secretName,omitempty"`
	KeyVaultToStoreSecrets string `json:"keyVaultToStoreSecrets,omitempty"`
}

// EventhubAuthorizationRule defines the name and rights of the access policy
type EventhubAuthorizationRule struct {
	// Name - Name of AuthorizationRule for eventhub
	Name string `json:"name,omitempty"`
	// Rights - Rights set on the AuthorizationRule
	Rights []string `json:"rights,omitempty"`
}

// EventHubStorageAccount contains details of the eventhub storage account
type EventHubStorageAccount struct {
	// ResourceGroup - Name of the storage account resource group
	// +kubebuilder:validation:Pattern=^[-\w\._\(\)]+$
	// +kubebuilder:validation:MinLength=1
	ResourceGroup string `json:"resourceGroup,omitempty"`
	// AccountName - Name of the storage account
	// +kubebuilder:validation:MaxLength=24
	// +kubebuilder:validation:MinLength=3
	// +kubebuilder:validation:Pattern=^[a-z0-9]+$
	AccountName string `json:"accountName,omitempty"`
}

// Destination for capture (blob storage etc)
type Destination struct {
	// ArchiveNameFormat - Blob naming convention for archive, e.g. {Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}. Here all the parameters (Namespace,EventHub .. etc) are mandatory irrespective of order
	ArchiveNameFormat string `json:"archiveNameFormat,omitempty"`
	// BlobContainer - Blob container Name
	BlobContainer string `json:"blobContainer,omitempty"`
	// Name - Name for capture destination
	// +kubebuilder:validation:Enum=EventHubArchive.AzureBlockBlob;EventHubArchive.AzureDataLake
	Name string `json:"name,omitempty"`
	// StorageAccount - Details of the storage account
	StorageAccount EventHubStorageAccount `json:"storageAccount,omitempty"`
}

// CaptureDescription defines the properties required for eventhub capture
type CaptureDescription struct {
	// Destination - Resource id of the storage account to be used to create the blobs
	Destination Destination `json:"destination,omitempty"`
	// Enabled - indicates whether capture is enabled
	Enabled bool `json:"enabled"`
	// SizeLimitInBytes - The size window defines the amount of data built up in your Event Hub before an capture operation
	// +kubebuilder:validation:Maximum=524288000
	// +kubebuilder:validation:Minimum=10485760
	SizeLimitInBytes int32 `json:"sizeLimitInBytes,omitempty"`
	// IntervalInSeconds - The time window allows you to set the frequency with which the capture to Azure Blobs will happen
	// +kubebuilder:validation:Maximum=900
	// +kubebuilder:validation:Minimum=60
	IntervalInSeconds int32 `json:"intervalInSeconds,omitempty"`
}

// EventhubProperties defines the namespace properties
type EventhubProperties struct {
	// +kubebuilder:validation:Maximum=7
	// +kubebuilder:validation:Minimum=1
	// MessageRetentionInDays - Number of days to retain the events for this Event Hub, value should be 1 to 7 days
	MessageRetentionInDays int32 `json:"messageRetentionInDays,omitempty"`
	// +kubebuilder:validation:Maximum=32
	// +kubebuilder:validation:Minimum=2
	// PartitionCount - Number of partitions created for the Event Hub, allowed values are from 2 to 32 partitions.
	PartitionCount int32 `json:"partitionCount,omitempty"`
	// CaptureDescription - Details specifying EventHub capture to persistent storage
	CaptureDescription CaptureDescription `json:"captureDescription,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// Eventhub is the Schema for the eventhubs API
// +kubebuilder:resource:shortName=eh
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type Eventhub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EventhubSpec `json:"spec,omitempty"`
	Status ASOStatus    `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventhubList contains a list of Eventhub
type EventhubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Eventhub `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Eventhub{}, &EventhubList{})
}

func (eventhub *Eventhub) IsSubmitted() bool {
	return eventhub.Status.Provisioning || eventhub.Status.Provisioned
}

func (eventhub *Eventhub) HasFinalizer(finalizerName string) bool {
	return helpers.ContainsString(eventhub.ObjectMeta.Finalizers, finalizerName)
}

func (eventhub *Eventhub) AddFinalizer(finalizerName string) {
	eventhub.ObjectMeta.Finalizers = append(eventhub.ObjectMeta.Finalizers, finalizerName)
}

func (eventhub *Eventhub) RemoveFinalizer(finalizerName string) {
	eventhub.ObjectMeta.Finalizers = helpers.RemoveString(eventhub.ObjectMeta.Finalizers, finalizerName)
}

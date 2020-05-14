// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// ASOStatus (AzureServiceOperatorsStatus) defines the observed state of resource actions
type ASOStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Provisioning       bool         `json:"provisioning,omitempty"`
	Provisioned        bool         `json:"provisioned,omitempty"`
	State              string       `json:"state,omitempty"`
	Message            string       `json:"message,omitempty"`
	ResourceId         string       `json:"resourceId,omitempty"`
	PollingURL         string       `json:"pollingUrl,omitempty"`
	SpecHash           string       `json:"specHash,omitempty"`
	ContainsUpdate     bool         `json:"containsUpdate,omitempty"`
	RequestedAt        *metav1.Time `json:"requested,omitempty"`
	CompletedAt        *metav1.Time `json:"completed,omitempty"`
	FailedProvisioning bool         `json:"failedProvisioning,omitempty"`
	FlattenedSecrets   bool         `json:"flattenedSecrets,omitempty"`
}

// GenericSpec is a struct to help get the KeyVaultName from the Spec
type GenericSpec struct {
	KeyVaultToStoreSecrets string `json:"keyVaultToStoreSecrets,omitempty"`
}

// GenericResource is a struct to help get a generic resource to extract keyvault name
type GenericResource struct {
	Spec GenericSpec `json:"spec,omitempty"`
}

// // StatusedObject used to unmarshall runtime.Object when we need Status
// type StatusedObject struct {
// 	metav1.TypeMeta   `json:",inline"`
// 	metav1.ObjectMeta `json:"metadata,omitempty"`

// 	Status ASOStatus `json:"status,omitempty"`
// }

// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:validation:Enum=rollallkeys;rollprimarykey;rollsecondarykey;rebootallnodes;rebootprimarynode;rebootsecondarynode
type RedisCacheActionName string

const (
	RedisCacheActionNameRollAllKeys         RedisCacheActionName = "rollallkeys"
	RedisCacheActionNameRollPrimaryKey      RedisCacheActionName = "rollprimarykey"
	RedisCacheActionNameRollSecondaryKey    RedisCacheActionName = "rollsecondarykey"
	RedisCacheActionNameRebootAllNodes      RedisCacheActionName = "rebootallnodes"
	RedisCacheActionNameRebootPrimaryNode   RedisCacheActionName = "rebootprimarynode"
	RedisCacheActionNameRebootSecondaryNode RedisCacheActionName = "rebootsecondarynode"
)

// RedisCacheActionSpec defines the desired state of RedisCacheAction
type RedisCacheActionSpec struct {
	// +kubebuilder:validation:Pattern=^[-\w\._\(\)]+$
	// +kubebuilder:validation:MinLength:1
	// +kubebuilder:validation:Required
	ResourceGroup          string               `json:"resourceGroup"`
	CacheName              string               `json:"cacheName"`
	ActionName             RedisCacheActionName `json:"actionName"`
	SecretName             string               `json:"secretName,omitempty"`
	KeyVaultToStoreSecrets string               `json:"keyVaultToStoreSecrets,omitempty"`
	ShardID                *int32               `json:"shardID,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RedisCacheAction is the Schema for the rediscacheactions API
// +kubebuilder:resource:shortName=rca
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type RedisCacheAction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedisCacheActionSpec `json:"spec,omitempty"`
	Status ASOStatus            `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisCacheActionList contains a list of RedisCacheAction
type RedisCacheActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisCacheAction `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RedisCacheAction{}, &RedisCacheActionList{})
}

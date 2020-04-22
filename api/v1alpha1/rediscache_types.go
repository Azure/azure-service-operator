// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RedisCacheSpec defines the desired state of RedisCache
type RedisCacheSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:MinLength=0

	Location               string               `json:"location,omitempty"`
	ResourceGroupName      string               `json:"resourceGroup"`
	Properties             RedisCacheProperties `json:"properties,omitempty"`
	SecretName             string               `json:"secretName,omitempty"`
	KeyVaultToStoreSecrets string               `json:"keyVaultToStoreSecrets,omitempty"`
}

// RedisCacheProperties the properties of the Redis Cache.
type RedisCacheProperties struct {
	Sku              RedisCacheSku     `json:"sku,omitempty"`
	EnableNonSslPort bool              `json:"enableNonSslPort,omitempty"`
	SubnetID         string            `json:"subnetId,omitempty"`
	StaticIP         string            `json:"staticIp,omitempty"`
	Configuration    map[string]string `json:"configuration,omitempty"`
}

// RedisCacheSku the SKU of the Redis Cache.
type RedisCacheSku struct {
	// Name - The SKU name. Required for account creation; optional for update.
	// Possible values include: 'StandardLRS', 'StandardGRS', 'StandardRAGRS', 'StandardZRS', 'PremiumLRS', 'PremiumZRS', 'StandardGZRS', 'StandardRAGZRS'
	Name RedisCacheSkuName `json:"name,omitempty"`

	Family RedisCacheSkuFamily `json:"family,omitempty"`

	Capacity int32 `json:"capacity,omitempty"`
}

type RedisCacheSkuName string

const (
	Basic    RedisCacheSkuName = "Basic"
	Premium  RedisCacheSkuName = "Premium"
	Standard RedisCacheSkuName = "Standard"
)

type RedisCacheSkuFamily string

const (
	C RedisCacheSkuFamily = "C"
	P RedisCacheSkuFamily = "P"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RedisCache is the Schema for the rediscaches API
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type RedisCache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedisCacheSpec `json:"spec,omitempty"`
	Status ASOStatus      `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisCacheList contains a list of RedisCache
type RedisCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisCache `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RedisCache{}, &RedisCacheList{})
}

func (redisCache *RedisCache) IsSubmitted() bool {
	return redisCache.Status.Provisioning || redisCache.Status.Provisioned
}

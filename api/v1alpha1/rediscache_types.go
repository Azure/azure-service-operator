/*
MIT License

Copyright (c) Microsoft Corporation. All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE
*/

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

	Location          string               `json:"location,omitempty"`
	ResourceGroupName string               `json:"resourceGroup"`
	Properties        RedisCacheProperties `json:"properties,omitempty"`
}

// RedisCacheProperties the properties of the Redis Cache.
type RedisCacheProperties struct {
	Sku RedisCacheSku `json:"sku,omitempty"`

	EnableNonSslPort bool `json:"enableNonSslPort,omitempty"`
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

type RedisCacheOutput struct {
	RedisCacheName string `json:"redisCacheName,omitempty"`
	PrimaryKey     string `json:"primaryKey,omitempty"`
	SecondaryKey   string `json:"secondaryKey,omitempty"`
}

// StorageAdditionalResources holds the additional resources
type RedisCacheAdditionalResources struct {
	Secrets []string `json:"secrets,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RedisCache is the Schema for the rediscaches API
type RedisCache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec                RedisCacheSpec                `json:"spec,omitempty"`
	Status              ASOStatus              `json:"status,omitempty"`
	Output              RedisCacheOutput              `json:"output,omitempty"`
	AdditionalResources RedisCacheAdditionalResources `json:"additionalResources,omitempty"`
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

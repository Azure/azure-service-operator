// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package extensions

import (
	cache "github.com/Azure/azure-service-operator/v2/api/cache/v1alpha1api20210301"
	"github.com/Azure/azure-service-operator/v2/api/cache/v1alpha1api20210301storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type RedisEnterpriseExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *RedisEnterpriseExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&cache.RedisEnterprise{},
		&v1alpha1api20210301storage.RedisEnterprise{}}
}

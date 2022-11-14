// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210301

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of RedisEnterprise_Database_Spec. Use v1beta20210301.RedisEnterprise_Database_Spec instead
type RedisEnterprise_Database_Spec_ARM struct {
	Name       string                  `json:"name,omitempty"`
	Properties *DatabaseProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RedisEnterprise_Database_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-03-01"
func (database RedisEnterprise_Database_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (database *RedisEnterprise_Database_Spec_ARM) GetName() string {
	return database.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redisEnterprise/databases"
func (database *RedisEnterprise_Database_Spec_ARM) GetType() string {
	return "Microsoft.Cache/redisEnterprise/databases"
}

// Deprecated version of DatabaseProperties. Use v1beta20210301.DatabaseProperties instead
type DatabaseProperties_ARM struct {
	ClientProtocol   *DatabaseProperties_ClientProtocol   `json:"clientProtocol,omitempty"`
	ClusteringPolicy *DatabaseProperties_ClusteringPolicy `json:"clusteringPolicy,omitempty"`
	EvictionPolicy   *DatabaseProperties_EvictionPolicy   `json:"evictionPolicy,omitempty"`
	Modules          []Module_ARM                         `json:"modules,omitempty"`
	Persistence      *Persistence_ARM                     `json:"persistence,omitempty"`
	Port             *int                                 `json:"port,omitempty"`
}

// Deprecated version of Module. Use v1beta20210301.Module instead
type Module_ARM struct {
	Args *string `json:"args,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Deprecated version of Persistence. Use v1beta20210301.Persistence instead
type Persistence_ARM struct {
	AofEnabled   *bool                     `json:"aofEnabled,omitempty"`
	AofFrequency *Persistence_AofFrequency `json:"aofFrequency,omitempty"`
	RdbEnabled   *bool                     `json:"rdbEnabled,omitempty"`
	RdbFrequency *Persistence_RdbFrequency `json:"rdbFrequency,omitempty"`
}

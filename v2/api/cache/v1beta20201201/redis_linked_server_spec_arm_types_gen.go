// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Redis_LinkedServer_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties required to create a linked server.
	Properties *RedisLinkedServerCreateProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Redis_LinkedServer_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (server Redis_LinkedServer_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (server *Redis_LinkedServer_Spec_ARM) GetName() string {
	return server.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/linkedServers"
func (server *Redis_LinkedServer_Spec_ARM) GetType() string {
	return "Microsoft.Cache/redis/linkedServers"
}

// Create properties for a linked server
type RedisLinkedServerCreateProperties_ARM struct {
	LinkedRedisCacheId *string `json:"linkedRedisCacheId,omitempty"`

	// LinkedRedisCacheLocation: Location of the linked redis cache.
	LinkedRedisCacheLocation *string `json:"linkedRedisCacheLocation,omitempty"`

	// ServerRole: Role of the linked server.
	ServerRole *RedisLinkedServerCreateProperties_ServerRole `json:"serverRole,omitempty"`
}

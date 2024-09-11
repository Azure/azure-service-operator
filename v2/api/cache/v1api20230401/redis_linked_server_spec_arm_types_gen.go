// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Redis_LinkedServer_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties required to create a linked server.
	Properties *RedisLinkedServerCreateProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Redis_LinkedServer_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-04-01"
func (server Redis_LinkedServer_Spec_ARM) GetAPIVersion() string {
	return "2023-04-01"
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
	ServerRole *RedisLinkedServerCreateProperties_ServerRole_ARM `json:"serverRole,omitempty"`
}

// +kubebuilder:validation:Enum={"Primary","Secondary"}
type RedisLinkedServerCreateProperties_ServerRole_ARM string

const (
	RedisLinkedServerCreateProperties_ServerRole_ARM_Primary   = RedisLinkedServerCreateProperties_ServerRole_ARM("Primary")
	RedisLinkedServerCreateProperties_ServerRole_ARM_Secondary = RedisLinkedServerCreateProperties_ServerRole_ARM("Secondary")
)

// Mapping from string to RedisLinkedServerCreateProperties_ServerRole_ARM
var redisLinkedServerCreateProperties_ServerRole_ARM_Values = map[string]RedisLinkedServerCreateProperties_ServerRole_ARM{
	"primary":   RedisLinkedServerCreateProperties_ServerRole_ARM_Primary,
	"secondary": RedisLinkedServerCreateProperties_ServerRole_ARM_Secondary,
}

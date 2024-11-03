// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RedisLinkedServer_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties required to create a linked server.
	Properties *RedisLinkedServerCreateProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RedisLinkedServer_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (server RedisLinkedServer_Spec) GetAPIVersion() string {
	return "2020-12-01"
}

// GetName returns the Name of the resource
func (server *RedisLinkedServer_Spec) GetName() string {
	return server.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/linkedServers"
func (server *RedisLinkedServer_Spec) GetType() string {
	return "Microsoft.Cache/redis/linkedServers"
}

// Create properties for a linked server
type RedisLinkedServerCreateProperties struct {
	LinkedRedisCacheId *string `json:"linkedRedisCacheId,omitempty"`

	// LinkedRedisCacheLocation: Location of the linked redis cache.
	LinkedRedisCacheLocation *string `json:"linkedRedisCacheLocation,omitempty"`

	// ServerRole: Role of the linked server.
	ServerRole *RedisLinkedServerCreateProperties_ServerRole `json:"serverRole,omitempty"`
}

// +kubebuilder:validation:Enum={"Primary","Secondary"}
type RedisLinkedServerCreateProperties_ServerRole string

const (
	RedisLinkedServerCreateProperties_ServerRole_Primary   = RedisLinkedServerCreateProperties_ServerRole("Primary")
	RedisLinkedServerCreateProperties_ServerRole_Secondary = RedisLinkedServerCreateProperties_ServerRole("Secondary")
)

// Mapping from string to RedisLinkedServerCreateProperties_ServerRole
var redisLinkedServerCreateProperties_ServerRole_Values = map[string]RedisLinkedServerCreateProperties_ServerRole{
	"primary":   RedisLinkedServerCreateProperties_ServerRole_Primary,
	"secondary": RedisLinkedServerCreateProperties_ServerRole_Secondary,
}

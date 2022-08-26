// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

<<<<<<<< HEAD:v2/api/cache/v1beta20201201/redis_linked_server__spec_arm_types_gen.go
type RedisLinkedServer_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`
	Name      string `json:"name,omitempty"`
========
type Redis_LinkedServers_SpecARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`
>>>>>>>> main:v2/api/cache/v1beta20201201/redis_linked_servers_spec_arm_types_gen.go

	// Properties: Properties required to create a linked server.
	Properties *RedisLinkedServerCreatePropertiesARM `json:"properties,omitempty"`
}

<<<<<<<< HEAD:v2/api/cache/v1beta20201201/redis_linked_server__spec_arm_types_gen.go
var _ genruntime.ARMResourceSpec = &RedisLinkedServer_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (server RedisLinkedServer_SpecARM) GetAPIVersion() string {
========
var _ genruntime.ARMResourceSpec = &Redis_LinkedServers_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (servers Redis_LinkedServers_SpecARM) GetAPIVersion() string {
>>>>>>>> main:v2/api/cache/v1beta20201201/redis_linked_servers_spec_arm_types_gen.go
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
<<<<<<<< HEAD:v2/api/cache/v1beta20201201/redis_linked_server__spec_arm_types_gen.go
func (server *RedisLinkedServer_SpecARM) GetName() string {
	return server.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/linkedServers"
func (server *RedisLinkedServer_SpecARM) GetType() string {
========
func (servers *Redis_LinkedServers_SpecARM) GetName() string {
	return servers.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/linkedServers"
func (servers *Redis_LinkedServers_SpecARM) GetType() string {
>>>>>>>> main:v2/api/cache/v1beta20201201/redis_linked_servers_spec_arm_types_gen.go
	return "Microsoft.Cache/redis/linkedServers"
}

type RedisLinkedServerCreatePropertiesARM struct {
	LinkedRedisCacheId *string `json:"linkedRedisCacheId,omitempty"`

	// LinkedRedisCacheLocation: Location of the linked redis cache.
	LinkedRedisCacheLocation *string `json:"linkedRedisCacheLocation,omitempty"`

	// ServerRole: Role of the linked server.
	ServerRole *RedisLinkedServerCreateProperties_ServerRole `json:"serverRole,omitempty"`
}

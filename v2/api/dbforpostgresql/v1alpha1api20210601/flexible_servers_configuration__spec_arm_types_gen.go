// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

<<<<<<<< HEAD:v2/api/dbforpostgresql/v1alpha1api20210601/flexible_servers_configuration__spec_arm_types_gen.go
// Deprecated version of FlexibleServersConfiguration_Spec. Use v1beta20210601.FlexibleServersConfiguration_Spec instead
type FlexibleServersConfiguration_SpecARM struct {
	AzureName  string                      `json:"azureName,omitempty"`
========
// Deprecated version of FlexibleServers_Configurations_Spec. Use v1beta20210601.FlexibleServers_Configurations_Spec instead
type FlexibleServers_Configurations_SpecARM struct {
	Location   *string                     `json:"location,omitempty"`
>>>>>>>> main:v2/api/dbforpostgresql/v1alpha1api20210601/flexible_servers_configurations_spec_arm_types_gen.go
	Name       string                      `json:"name,omitempty"`
	Properties *ConfigurationPropertiesARM `json:"properties,omitempty"`
}

<<<<<<<< HEAD:v2/api/dbforpostgresql/v1alpha1api20210601/flexible_servers_configuration__spec_arm_types_gen.go
var _ genruntime.ARMResourceSpec = &FlexibleServersConfiguration_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (configuration FlexibleServersConfiguration_SpecARM) GetAPIVersion() string {
========
var _ genruntime.ARMResourceSpec = &FlexibleServers_Configurations_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (configurations FlexibleServers_Configurations_SpecARM) GetAPIVersion() string {
>>>>>>>> main:v2/api/dbforpostgresql/v1alpha1api20210601/flexible_servers_configurations_spec_arm_types_gen.go
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
<<<<<<<< HEAD:v2/api/dbforpostgresql/v1alpha1api20210601/flexible_servers_configuration__spec_arm_types_gen.go
func (configuration *FlexibleServersConfiguration_SpecARM) GetName() string {
	return configuration.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
func (configuration *FlexibleServersConfiguration_SpecARM) GetType() string {
========
func (configurations *FlexibleServers_Configurations_SpecARM) GetName() string {
	return configurations.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
func (configurations *FlexibleServers_Configurations_SpecARM) GetType() string {
>>>>>>>> main:v2/api/dbforpostgresql/v1alpha1api20210601/flexible_servers_configurations_spec_arm_types_gen.go
	return "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
}

// Deprecated version of ConfigurationProperties. Use v1beta20210601.ConfigurationProperties instead
type ConfigurationPropertiesARM struct {
	Source *string `json:"source,omitempty"`
	Value  *string `json:"value,omitempty"`
}

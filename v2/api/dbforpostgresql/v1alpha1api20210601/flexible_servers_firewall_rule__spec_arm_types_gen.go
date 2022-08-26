// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

<<<<<<<< HEAD:v2/api/dbforpostgresql/v1alpha1api20210601/flexible_servers_firewall_rule__spec_arm_types_gen.go
// Deprecated version of FlexibleServersFirewallRule_Spec. Use v1beta20210601.FlexibleServersFirewallRule_Spec instead
type FlexibleServersFirewallRule_SpecARM struct {
	AzureName  string                     `json:"azureName,omitempty"`
========
// Deprecated version of FlexibleServers_FirewallRules_Spec. Use v1beta20210601.FlexibleServers_FirewallRules_Spec instead
type FlexibleServers_FirewallRules_SpecARM struct {
	Location   *string                    `json:"location,omitempty"`
>>>>>>>> main:v2/api/dbforpostgresql/v1alpha1api20210601/flexible_servers_firewall_rules_spec_arm_types_gen.go
	Name       string                     `json:"name,omitempty"`
	Properties *FirewallRulePropertiesARM `json:"properties,omitempty"`
}

<<<<<<<< HEAD:v2/api/dbforpostgresql/v1alpha1api20210601/flexible_servers_firewall_rule__spec_arm_types_gen.go
var _ genruntime.ARMResourceSpec = &FlexibleServersFirewallRule_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (rule FlexibleServersFirewallRule_SpecARM) GetAPIVersion() string {
========
var _ genruntime.ARMResourceSpec = &FlexibleServers_FirewallRules_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (rules FlexibleServers_FirewallRules_SpecARM) GetAPIVersion() string {
>>>>>>>> main:v2/api/dbforpostgresql/v1alpha1api20210601/flexible_servers_firewall_rules_spec_arm_types_gen.go
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
<<<<<<<< HEAD:v2/api/dbforpostgresql/v1alpha1api20210601/flexible_servers_firewall_rule__spec_arm_types_gen.go
func (rule *FlexibleServersFirewallRule_SpecARM) GetName() string {
	return rule.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/firewallRules"
func (rule *FlexibleServersFirewallRule_SpecARM) GetType() string {
========
func (rules *FlexibleServers_FirewallRules_SpecARM) GetName() string {
	return rules.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/firewallRules"
func (rules *FlexibleServers_FirewallRules_SpecARM) GetType() string {
>>>>>>>> main:v2/api/dbforpostgresql/v1alpha1api20210601/flexible_servers_firewall_rules_spec_arm_types_gen.go
	return "Microsoft.DBforPostgreSQL/flexibleServers/firewallRules"
}

// Deprecated version of FirewallRuleProperties. Use v1beta20210601.FirewallRuleProperties instead
type FirewallRulePropertiesARM struct {
	EndIpAddress   *string `json:"endIpAddress,omitempty"`
	StartIpAddress *string `json:"startIpAddress,omitempty"`
}

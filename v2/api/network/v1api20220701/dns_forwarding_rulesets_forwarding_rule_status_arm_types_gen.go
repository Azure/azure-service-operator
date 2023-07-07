// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

type DnsForwardingRulesets_ForwardingRule_STATUS_ARM struct {
	// Etag: ETag of the forwarding rule.
	Etag *string `json:"etag,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the forwarding rule.
	Properties *ForwardingRuleProperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Represents the properties of a forwarding rule within a DNS forwarding ruleset.
type ForwardingRuleProperties_STATUS_ARM struct {
	// DomainName: The domain name for the forwarding rule.
	DomainName *string `json:"domainName,omitempty"`

	// ForwardingRuleState: The state of forwarding rule.
	ForwardingRuleState *ForwardingRuleProperties_ForwardingRuleState_STATUS `json:"forwardingRuleState,omitempty"`

	// Metadata: Metadata attached to the forwarding rule.
	Metadata map[string]string `json:"metadata,omitempty"`

	// ProvisioningState: The current provisioning state of the forwarding rule. This is a read-only property and any attempt
	// to set this value will be ignored.
	ProvisioningState *DnsresolverProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// TargetDnsServers: DNS servers to forward the DNS query to.
	TargetDnsServers []TargetDnsServer_STATUS_ARM `json:"targetDnsServers,omitempty"`
}

// Describes a server to forward the DNS queries to.
type TargetDnsServer_STATUS_ARM struct {
	// IpAddress: DNS server IP address.
	IpAddress *string `json:"ipAddress,omitempty"`

	// Port: DNS server port.
	Port *int `json:"port,omitempty"`
}
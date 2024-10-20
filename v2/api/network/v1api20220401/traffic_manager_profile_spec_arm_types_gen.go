// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type TrafficManagerProfile_Spec_ARM struct {
	// Location: The Azure Region where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name string `json:"name,omitempty"`

	// Properties: The properties of the Traffic Manager profile.
	Properties *ProfileProperties_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &TrafficManagerProfile_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-04-01"
func (profile TrafficManagerProfile_Spec_ARM) GetAPIVersion() string {
	return "2022-04-01"
}

// GetName returns the Name of the resource
func (profile *TrafficManagerProfile_Spec_ARM) GetName() string {
	return profile.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/trafficmanagerprofiles"
func (profile *TrafficManagerProfile_Spec_ARM) GetType() string {
	return "Microsoft.Network/trafficmanagerprofiles"
}

// Class representing the Traffic Manager profile properties.
type ProfileProperties_ARM struct {
	// AllowedEndpointRecordTypes: The list of allowed endpoint record types.
	AllowedEndpointRecordTypes []AllowedEndpointRecordType_ARM `json:"allowedEndpointRecordTypes,omitempty"`

	// DnsConfig: The DNS settings of the Traffic Manager profile.
	DnsConfig *DnsConfig_ARM `json:"dnsConfig,omitempty"`

	// MaxReturn: Maximum number of endpoints to be returned for MultiValue routing type.
	MaxReturn *int `json:"maxReturn,omitempty"`

	// MonitorConfig: The endpoint monitoring settings of the Traffic Manager profile.
	MonitorConfig *MonitorConfig_ARM `json:"monitorConfig,omitempty"`

	// ProfileStatus: The status of the Traffic Manager profile.
	ProfileStatus *ProfileProperties_ProfileStatus_ARM `json:"profileStatus,omitempty"`

	// TrafficRoutingMethod: The traffic routing method of the Traffic Manager profile.
	TrafficRoutingMethod *ProfileProperties_TrafficRoutingMethod_ARM `json:"trafficRoutingMethod,omitempty"`

	// TrafficViewEnrollmentStatus: Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile.
	// Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
	TrafficViewEnrollmentStatus *ProfileProperties_TrafficViewEnrollmentStatus_ARM `json:"trafficViewEnrollmentStatus,omitempty"`
}

// The allowed type DNS record types for this profile.
// +kubebuilder:validation:Enum={"Any","DomainName","IPv4Address","IPv6Address"}
type AllowedEndpointRecordType_ARM string

const (
	AllowedEndpointRecordType_ARM_Any         = AllowedEndpointRecordType_ARM("Any")
	AllowedEndpointRecordType_ARM_DomainName  = AllowedEndpointRecordType_ARM("DomainName")
	AllowedEndpointRecordType_ARM_IPv4Address = AllowedEndpointRecordType_ARM("IPv4Address")
	AllowedEndpointRecordType_ARM_IPv6Address = AllowedEndpointRecordType_ARM("IPv6Address")
)

// Mapping from string to AllowedEndpointRecordType_ARM
var allowedEndpointRecordType_ARM_Values = map[string]AllowedEndpointRecordType_ARM{
	"any":         AllowedEndpointRecordType_ARM_Any,
	"domainname":  AllowedEndpointRecordType_ARM_DomainName,
	"ipv4address": AllowedEndpointRecordType_ARM_IPv4Address,
	"ipv6address": AllowedEndpointRecordType_ARM_IPv6Address,
}

// Class containing DNS settings in a Traffic Manager profile.
type DnsConfig_ARM struct {
	// RelativeName: The relative DNS name provided by this Traffic Manager profile. This value is combined with the DNS domain
	// name used by Azure Traffic Manager to form the fully-qualified domain name (FQDN) of the profile.
	RelativeName *string `json:"relativeName,omitempty"`

	// Ttl: The DNS Time-To-Live (TTL), in seconds. This informs the local DNS resolvers and DNS clients how long to cache DNS
	// responses provided by this Traffic Manager profile.
	Ttl *int `json:"ttl,omitempty"`
}

// Class containing endpoint monitoring settings in a Traffic Manager profile.
type MonitorConfig_ARM struct {
	// CustomHeaders: List of custom headers.
	CustomHeaders []MonitorConfig_CustomHeaders_ARM `json:"customHeaders,omitempty"`

	// ExpectedStatusCodeRanges: List of expected status code ranges.
	ExpectedStatusCodeRanges []MonitorConfig_ExpectedStatusCodeRanges_ARM `json:"expectedStatusCodeRanges,omitempty"`

	// IntervalInSeconds: The monitor interval for endpoints in this profile. This is the interval at which Traffic Manager
	// will check the health of each endpoint in this profile.
	IntervalInSeconds *int `json:"intervalInSeconds,omitempty"`

	// Path: The path relative to the endpoint domain name used to probe for endpoint health.
	Path *string `json:"path,omitempty"`

	// Port: The TCP port used to probe for endpoint health.
	Port *int `json:"port,omitempty"`

	// ProfileMonitorStatus: The profile-level monitoring status of the Traffic Manager profile.
	ProfileMonitorStatus *MonitorConfig_ProfileMonitorStatus_ARM `json:"profileMonitorStatus,omitempty"`

	// Protocol: The protocol (HTTP, HTTPS or TCP) used to probe for endpoint health.
	Protocol *MonitorConfig_Protocol_ARM `json:"protocol,omitempty"`

	// TimeoutInSeconds: The monitor timeout for endpoints in this profile. This is the time that Traffic Manager allows
	// endpoints in this profile to response to the health check.
	TimeoutInSeconds *int `json:"timeoutInSeconds,omitempty"`

	// ToleratedNumberOfFailures: The number of consecutive failed health check that Traffic Manager tolerates before declaring
	// an endpoint in this profile Degraded after the next failed health check.
	ToleratedNumberOfFailures *int `json:"toleratedNumberOfFailures,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ProfileProperties_ProfileStatus_ARM string

const (
	ProfileProperties_ProfileStatus_ARM_Disabled = ProfileProperties_ProfileStatus_ARM("Disabled")
	ProfileProperties_ProfileStatus_ARM_Enabled  = ProfileProperties_ProfileStatus_ARM("Enabled")
)

// Mapping from string to ProfileProperties_ProfileStatus_ARM
var profileProperties_ProfileStatus_ARM_Values = map[string]ProfileProperties_ProfileStatus_ARM{
	"disabled": ProfileProperties_ProfileStatus_ARM_Disabled,
	"enabled":  ProfileProperties_ProfileStatus_ARM_Enabled,
}

// +kubebuilder:validation:Enum={"Geographic","MultiValue","Performance","Priority","Subnet","Weighted"}
type ProfileProperties_TrafficRoutingMethod_ARM string

const (
	ProfileProperties_TrafficRoutingMethod_ARM_Geographic  = ProfileProperties_TrafficRoutingMethod_ARM("Geographic")
	ProfileProperties_TrafficRoutingMethod_ARM_MultiValue  = ProfileProperties_TrafficRoutingMethod_ARM("MultiValue")
	ProfileProperties_TrafficRoutingMethod_ARM_Performance = ProfileProperties_TrafficRoutingMethod_ARM("Performance")
	ProfileProperties_TrafficRoutingMethod_ARM_Priority    = ProfileProperties_TrafficRoutingMethod_ARM("Priority")
	ProfileProperties_TrafficRoutingMethod_ARM_Subnet      = ProfileProperties_TrafficRoutingMethod_ARM("Subnet")
	ProfileProperties_TrafficRoutingMethod_ARM_Weighted    = ProfileProperties_TrafficRoutingMethod_ARM("Weighted")
)

// Mapping from string to ProfileProperties_TrafficRoutingMethod_ARM
var profileProperties_TrafficRoutingMethod_ARM_Values = map[string]ProfileProperties_TrafficRoutingMethod_ARM{
	"geographic":  ProfileProperties_TrafficRoutingMethod_ARM_Geographic,
	"multivalue":  ProfileProperties_TrafficRoutingMethod_ARM_MultiValue,
	"performance": ProfileProperties_TrafficRoutingMethod_ARM_Performance,
	"priority":    ProfileProperties_TrafficRoutingMethod_ARM_Priority,
	"subnet":      ProfileProperties_TrafficRoutingMethod_ARM_Subnet,
	"weighted":    ProfileProperties_TrafficRoutingMethod_ARM_Weighted,
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ProfileProperties_TrafficViewEnrollmentStatus_ARM string

const (
	ProfileProperties_TrafficViewEnrollmentStatus_ARM_Disabled = ProfileProperties_TrafficViewEnrollmentStatus_ARM("Disabled")
	ProfileProperties_TrafficViewEnrollmentStatus_ARM_Enabled  = ProfileProperties_TrafficViewEnrollmentStatus_ARM("Enabled")
)

// Mapping from string to ProfileProperties_TrafficViewEnrollmentStatus_ARM
var profileProperties_TrafficViewEnrollmentStatus_ARM_Values = map[string]ProfileProperties_TrafficViewEnrollmentStatus_ARM{
	"disabled": ProfileProperties_TrafficViewEnrollmentStatus_ARM_Disabled,
	"enabled":  ProfileProperties_TrafficViewEnrollmentStatus_ARM_Enabled,
}

type MonitorConfig_CustomHeaders_ARM struct {
	// Name: Header name.
	Name *string `json:"name,omitempty"`

	// Value: Header value.
	Value *string `json:"value,omitempty"`
}

type MonitorConfig_ExpectedStatusCodeRanges_ARM struct {
	// Max: Max status code.
	Max *int `json:"max,omitempty"`

	// Min: Min status code.
	Min *int `json:"min,omitempty"`
}

// +kubebuilder:validation:Enum={"CheckingEndpoints","Degraded","Disabled","Inactive","Online"}
type MonitorConfig_ProfileMonitorStatus_ARM string

const (
	MonitorConfig_ProfileMonitorStatus_ARM_CheckingEndpoints = MonitorConfig_ProfileMonitorStatus_ARM("CheckingEndpoints")
	MonitorConfig_ProfileMonitorStatus_ARM_Degraded          = MonitorConfig_ProfileMonitorStatus_ARM("Degraded")
	MonitorConfig_ProfileMonitorStatus_ARM_Disabled          = MonitorConfig_ProfileMonitorStatus_ARM("Disabled")
	MonitorConfig_ProfileMonitorStatus_ARM_Inactive          = MonitorConfig_ProfileMonitorStatus_ARM("Inactive")
	MonitorConfig_ProfileMonitorStatus_ARM_Online            = MonitorConfig_ProfileMonitorStatus_ARM("Online")
)

// Mapping from string to MonitorConfig_ProfileMonitorStatus_ARM
var monitorConfig_ProfileMonitorStatus_ARM_Values = map[string]MonitorConfig_ProfileMonitorStatus_ARM{
	"checkingendpoints": MonitorConfig_ProfileMonitorStatus_ARM_CheckingEndpoints,
	"degraded":          MonitorConfig_ProfileMonitorStatus_ARM_Degraded,
	"disabled":          MonitorConfig_ProfileMonitorStatus_ARM_Disabled,
	"inactive":          MonitorConfig_ProfileMonitorStatus_ARM_Inactive,
	"online":            MonitorConfig_ProfileMonitorStatus_ARM_Online,
}

// +kubebuilder:validation:Enum={"HTTP","HTTPS","TCP"}
type MonitorConfig_Protocol_ARM string

const (
	MonitorConfig_Protocol_ARM_HTTP  = MonitorConfig_Protocol_ARM("HTTP")
	MonitorConfig_Protocol_ARM_HTTPS = MonitorConfig_Protocol_ARM("HTTPS")
	MonitorConfig_Protocol_ARM_TCP   = MonitorConfig_Protocol_ARM("TCP")
)

// Mapping from string to MonitorConfig_Protocol_ARM
var monitorConfig_Protocol_ARM_Values = map[string]MonitorConfig_Protocol_ARM{
	"http":  MonitorConfig_Protocol_ARM_HTTP,
	"https": MonitorConfig_Protocol_ARM_HTTPS,
	"tcp":   MonitorConfig_Protocol_ARM_TCP,
}

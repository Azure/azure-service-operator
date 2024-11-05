// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type AfdOrigin_STATUS struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: The JSON object that contains the properties of the origin.
	Properties *AFDOriginProperties_STATUS `json:"properties,omitempty"`

	// SystemData: Read only system data
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// The JSON object that contains the properties of the origin.
type AFDOriginProperties_STATUS struct {
	// AzureOrigin: Resource reference to the Azure origin resource.
	AzureOrigin      *ResourceReference_STATUS                    `json:"azureOrigin,omitempty"`
	DeploymentStatus *AFDOriginProperties_DeploymentStatus_STATUS `json:"deploymentStatus,omitempty"`

	// EnabledState: Whether to enable health probes to be made against backends defined under backendPools. Health probes can
	// only be disabled if there is a single enabled backend in single enabled backend pool.
	EnabledState *AFDOriginProperties_EnabledState_STATUS `json:"enabledState,omitempty"`

	// EnforceCertificateNameCheck: Whether to enable certificate name check at origin level
	EnforceCertificateNameCheck *bool `json:"enforceCertificateNameCheck,omitempty"`

	// HostName: The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be
	// unique across all origins in an endpoint.
	HostName *string `json:"hostName,omitempty"`

	// HttpPort: The value of the HTTP port. Must be between 1 and 65535.
	HttpPort *int `json:"httpPort,omitempty"`

	// HttpsPort: The value of the HTTPS port. Must be between 1 and 65535.
	HttpsPort *int `json:"httpsPort,omitempty"`

	// OriginGroupName: The name of the origin group which contains this origin.
	OriginGroupName *string `json:"originGroupName,omitempty"`

	// OriginHostHeader: The host header value sent to the origin with each request. If you leave this blank, the request
	// hostname determines this value. Azure Front Door origins, such as Web Apps, Blob Storage, and Cloud Services require
	// this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
	OriginHostHeader *string `json:"originHostHeader,omitempty"`

	// Priority: Priority of origin in given origin group for load balancing. Higher priorities will not be used for load
	// balancing if any lower priority origin is healthy.Must be between 1 and 5
	Priority *int `json:"priority,omitempty"`

	// ProvisioningState: Provisioning status
	ProvisioningState *AFDOriginProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// SharedPrivateLinkResource: The properties of the private link resource for private origin.
	SharedPrivateLinkResource *SharedPrivateLinkResourceProperties_STATUS `json:"sharedPrivateLinkResource,omitempty"`

	// Weight: Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
	Weight *int `json:"weight,omitempty"`
}

type AFDOriginProperties_DeploymentStatus_STATUS string

const (
	AFDOriginProperties_DeploymentStatus_STATUS_Failed     = AFDOriginProperties_DeploymentStatus_STATUS("Failed")
	AFDOriginProperties_DeploymentStatus_STATUS_InProgress = AFDOriginProperties_DeploymentStatus_STATUS("InProgress")
	AFDOriginProperties_DeploymentStatus_STATUS_NotStarted = AFDOriginProperties_DeploymentStatus_STATUS("NotStarted")
	AFDOriginProperties_DeploymentStatus_STATUS_Succeeded  = AFDOriginProperties_DeploymentStatus_STATUS("Succeeded")
)

// Mapping from string to AFDOriginProperties_DeploymentStatus_STATUS
var aFDOriginProperties_DeploymentStatus_STATUS_Values = map[string]AFDOriginProperties_DeploymentStatus_STATUS{
	"failed":     AFDOriginProperties_DeploymentStatus_STATUS_Failed,
	"inprogress": AFDOriginProperties_DeploymentStatus_STATUS_InProgress,
	"notstarted": AFDOriginProperties_DeploymentStatus_STATUS_NotStarted,
	"succeeded":  AFDOriginProperties_DeploymentStatus_STATUS_Succeeded,
}

type AFDOriginProperties_EnabledState_STATUS string

const (
	AFDOriginProperties_EnabledState_STATUS_Disabled = AFDOriginProperties_EnabledState_STATUS("Disabled")
	AFDOriginProperties_EnabledState_STATUS_Enabled  = AFDOriginProperties_EnabledState_STATUS("Enabled")
)

// Mapping from string to AFDOriginProperties_EnabledState_STATUS
var aFDOriginProperties_EnabledState_STATUS_Values = map[string]AFDOriginProperties_EnabledState_STATUS{
	"disabled": AFDOriginProperties_EnabledState_STATUS_Disabled,
	"enabled":  AFDOriginProperties_EnabledState_STATUS_Enabled,
}

type AFDOriginProperties_ProvisioningState_STATUS string

const (
	AFDOriginProperties_ProvisioningState_STATUS_Creating  = AFDOriginProperties_ProvisioningState_STATUS("Creating")
	AFDOriginProperties_ProvisioningState_STATUS_Deleting  = AFDOriginProperties_ProvisioningState_STATUS("Deleting")
	AFDOriginProperties_ProvisioningState_STATUS_Failed    = AFDOriginProperties_ProvisioningState_STATUS("Failed")
	AFDOriginProperties_ProvisioningState_STATUS_Succeeded = AFDOriginProperties_ProvisioningState_STATUS("Succeeded")
	AFDOriginProperties_ProvisioningState_STATUS_Updating  = AFDOriginProperties_ProvisioningState_STATUS("Updating")
)

// Mapping from string to AFDOriginProperties_ProvisioningState_STATUS
var aFDOriginProperties_ProvisioningState_STATUS_Values = map[string]AFDOriginProperties_ProvisioningState_STATUS{
	"creating":  AFDOriginProperties_ProvisioningState_STATUS_Creating,
	"deleting":  AFDOriginProperties_ProvisioningState_STATUS_Deleting,
	"failed":    AFDOriginProperties_ProvisioningState_STATUS_Failed,
	"succeeded": AFDOriginProperties_ProvisioningState_STATUS_Succeeded,
	"updating":  AFDOriginProperties_ProvisioningState_STATUS_Updating,
}

// Describes the properties of an existing Shared Private Link Resource to use when connecting to a private origin.
type SharedPrivateLinkResourceProperties_STATUS struct {
	// GroupId: The group id from the provider of resource the shared private link resource is for.
	GroupId *string `json:"groupId,omitempty"`

	// PrivateLink: The resource id of the resource the shared private link resource is for.
	PrivateLink *ResourceReference_STATUS `json:"privateLink,omitempty"`

	// PrivateLinkLocation: The location of the shared private link resource
	PrivateLinkLocation *string `json:"privateLinkLocation,omitempty"`

	// RequestMessage: The request message for requesting approval of the shared private link resource.
	RequestMessage *string `json:"requestMessage,omitempty"`

	// Status: Status of the shared private link resource. Can be Pending, Approved, Rejected, Disconnected, or Timeout.
	Status *SharedPrivateLinkResourceProperties_Status_STATUS `json:"status,omitempty"`
}

type SharedPrivateLinkResourceProperties_Status_STATUS string

const (
	SharedPrivateLinkResourceProperties_Status_STATUS_Approved     = SharedPrivateLinkResourceProperties_Status_STATUS("Approved")
	SharedPrivateLinkResourceProperties_Status_STATUS_Disconnected = SharedPrivateLinkResourceProperties_Status_STATUS("Disconnected")
	SharedPrivateLinkResourceProperties_Status_STATUS_Pending      = SharedPrivateLinkResourceProperties_Status_STATUS("Pending")
	SharedPrivateLinkResourceProperties_Status_STATUS_Rejected     = SharedPrivateLinkResourceProperties_Status_STATUS("Rejected")
	SharedPrivateLinkResourceProperties_Status_STATUS_Timeout      = SharedPrivateLinkResourceProperties_Status_STATUS("Timeout")
)

// Mapping from string to SharedPrivateLinkResourceProperties_Status_STATUS
var sharedPrivateLinkResourceProperties_Status_STATUS_Values = map[string]SharedPrivateLinkResourceProperties_Status_STATUS{
	"approved":     SharedPrivateLinkResourceProperties_Status_STATUS_Approved,
	"disconnected": SharedPrivateLinkResourceProperties_Status_STATUS_Disconnected,
	"pending":      SharedPrivateLinkResourceProperties_Status_STATUS_Pending,
	"rejected":     SharedPrivateLinkResourceProperties_Status_STATUS_Rejected,
	"timeout":      SharedPrivateLinkResourceProperties_Status_STATUS_Timeout,
}

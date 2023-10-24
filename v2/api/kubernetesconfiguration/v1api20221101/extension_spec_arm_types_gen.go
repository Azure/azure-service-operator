// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20221101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Extension_Spec_ARM struct {
	// Identity: Identity of the Extension resource
	Identity *Identity_ARM `json:"identity,omitempty"`
	Name     string        `json:"name,omitempty"`

	// Plan: The plan information.
	Plan *Plan_ARM `json:"plan,omitempty"`

	// Properties: Properties of an Extension resource
	Properties *Extension_Properties_Spec_ARM `json:"properties,omitempty"`

	// SystemData: Top level metadata
	// https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
	SystemData *SystemData_ARM `json:"systemData,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Extension_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-11-01"
func (extension Extension_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (extension *Extension_Spec_ARM) GetName() string {
	return extension.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.KubernetesConfiguration/extensions"
func (extension *Extension_Spec_ARM) GetType() string {
	return "Microsoft.KubernetesConfiguration/extensions"
}

type Extension_Properties_Spec_ARM struct {
	// AksAssignedIdentity: Identity of the Extension resource in an AKS cluster
	AksAssignedIdentity *Extension_Properties_AksAssignedIdentity_Spec_ARM `json:"aksAssignedIdentity,omitempty"`

	// AutoUpgradeMinorVersion: Flag to note if this extension participates in auto upgrade of minor version, or not.
	AutoUpgradeMinorVersion *bool `json:"autoUpgradeMinorVersion,omitempty"`

	// ConfigurationSettings: Configuration settings, as name-value pairs for configuring this extension.
	ConfigurationSettings map[string]string `json:"configurationSettings,omitempty"`

	// ExtensionType: Type of the Extension, of which this resource is an instance of.  It must be one of the Extension Types
	// registered with Microsoft.KubernetesConfiguration by the Extension publisher.
	ExtensionType *string `json:"extensionType,omitempty"`

	// ReleaseTrain: ReleaseTrain this extension participates in for auto-upgrade (e.g. Stable, Preview, etc.) - only if
	// autoUpgradeMinorVersion is 'true'.
	ReleaseTrain *string `json:"releaseTrain,omitempty"`

	// Scope: Scope at which the extension is installed.
	Scope *Scope_ARM `json:"scope,omitempty"`

	// Statuses: Status from this extension.
	Statuses []ExtensionStatus_ARM `json:"statuses,omitempty"`

	// Version: User-specified version of the extension for this extension to 'pin'. To use 'version', autoUpgradeMinorVersion
	// must be 'false'.
	Version *string `json:"version,omitempty"`
}

// Identity for the resource.
type Identity_ARM struct {
	// Type: The identity type.
	Type *Identity_Type `json:"type,omitempty"`
}

// Plan for the resource.
type Plan_ARM struct {
	// Name: A user defined name of the 3rd Party Artifact that is being procured.
	Name *string `json:"name,omitempty"`

	// Product: The 3rd Party artifact that is being procured. E.g. NewRelic. Product maps to the OfferID specified for the
	// artifact at the time of Data Market onboarding.
	Product *string `json:"product,omitempty"`

	// PromotionCode: A publisher provided promotion code as provisioned in Data Market for the said product/artifact.
	PromotionCode *string `json:"promotionCode,omitempty"`

	// Publisher: The publisher of the 3rd Party Artifact that is being bought. E.g. NewRelic
	Publisher *string `json:"publisher,omitempty"`

	// Version: The version of the desired product/artifact.
	Version *string `json:"version,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_ARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType `json:"lastModifiedByType,omitempty"`
}

type Extension_Properties_AksAssignedIdentity_Spec_ARM struct {
	// Type: The identity type.
	Type *Extension_Properties_AksAssignedIdentity_Type_Spec `json:"type,omitempty"`
}

// Status from the extension.
type ExtensionStatus_ARM struct {
	// Code: Status code provided by the Extension
	Code *string `json:"code,omitempty"`

	// DisplayStatus: Short description of status of the extension.
	DisplayStatus *string `json:"displayStatus,omitempty"`

	// Level: Level of the status.
	Level *ExtensionStatus_Level `json:"level,omitempty"`

	// Message: Detailed message of the status from the Extension.
	Message *string `json:"message,omitempty"`

	// Time: DateLiteral (per ISO8601) noting the time of installation status.
	Time *string `json:"time,omitempty"`
}

// +kubebuilder:validation:Enum={"SystemAssigned"}
type Identity_Type string

const Identity_Type_SystemAssigned = Identity_Type("SystemAssigned")

// Scope of the extension. It can be either Cluster or Namespace; but not both.
type Scope_ARM struct {
	// Cluster: Specifies that the scope of the extension is Cluster
	Cluster *ScopeCluster_ARM `json:"cluster,omitempty"`

	// Namespace: Specifies that the scope of the extension is Namespace
	Namespace *ScopeNamespace_ARM `json:"namespace,omitempty"`
}

// +kubebuilder:validation:Enum={"Application","Key","ManagedIdentity","User"}
type SystemData_CreatedByType string

const (
	SystemData_CreatedByType_Application     = SystemData_CreatedByType("Application")
	SystemData_CreatedByType_Key             = SystemData_CreatedByType("Key")
	SystemData_CreatedByType_ManagedIdentity = SystemData_CreatedByType("ManagedIdentity")
	SystemData_CreatedByType_User            = SystemData_CreatedByType("User")
)

// +kubebuilder:validation:Enum={"Application","Key","ManagedIdentity","User"}
type SystemData_LastModifiedByType string

const (
	SystemData_LastModifiedByType_Application     = SystemData_LastModifiedByType("Application")
	SystemData_LastModifiedByType_Key             = SystemData_LastModifiedByType("Key")
	SystemData_LastModifiedByType_ManagedIdentity = SystemData_LastModifiedByType("ManagedIdentity")
	SystemData_LastModifiedByType_User            = SystemData_LastModifiedByType("User")
)

// Specifies that the scope of the extension is Cluster
type ScopeCluster_ARM struct {
	// ReleaseNamespace: Namespace where the extension Release must be placed, for a Cluster scoped extension.  If this
	// namespace does not exist, it will be created
	ReleaseNamespace *string `json:"releaseNamespace,omitempty"`
}

// Specifies that the scope of the extension is Namespace
type ScopeNamespace_ARM struct {
	// TargetNamespace: Namespace where the extension will be created for an Namespace scoped extension.  If this namespace
	// does not exist, it will be created
	TargetNamespace *string `json:"targetNamespace,omitempty"`
}

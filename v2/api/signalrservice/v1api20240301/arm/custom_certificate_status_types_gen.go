// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type CustomCertificate_STATUS struct {
	// Id: Fully qualified resource ID for the resource. E.g.
	// "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Custom certificate properties.
	Properties *CustomCertificateProperties_STATUS `json:"properties,omitempty"`

	// SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Custom certificate properties.
type CustomCertificateProperties_STATUS struct {
	// KeyVaultBaseUri: Base uri of the KeyVault that stores certificate.
	KeyVaultBaseUri *string `json:"keyVaultBaseUri,omitempty"`

	// KeyVaultSecretName: Certificate secret name.
	KeyVaultSecretName *string `json:"keyVaultSecretName,omitempty"`

	// KeyVaultSecretVersion: Certificate secret version.
	KeyVaultSecretVersion *string `json:"keyVaultSecretVersion,omitempty"`

	// ProvisioningState: Provisioning state of the resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

// Provisioning state of the resource.
type ProvisioningState_STATUS string

const (
	ProvisioningState_STATUS_Canceled  = ProvisioningState_STATUS("Canceled")
	ProvisioningState_STATUS_Creating  = ProvisioningState_STATUS("Creating")
	ProvisioningState_STATUS_Deleting  = ProvisioningState_STATUS("Deleting")
	ProvisioningState_STATUS_Failed    = ProvisioningState_STATUS("Failed")
	ProvisioningState_STATUS_Moving    = ProvisioningState_STATUS("Moving")
	ProvisioningState_STATUS_Running   = ProvisioningState_STATUS("Running")
	ProvisioningState_STATUS_Succeeded = ProvisioningState_STATUS("Succeeded")
	ProvisioningState_STATUS_Unknown   = ProvisioningState_STATUS("Unknown")
	ProvisioningState_STATUS_Updating  = ProvisioningState_STATUS("Updating")
)

// Mapping from string to ProvisioningState_STATUS
var provisioningState_STATUS_Values = map[string]ProvisioningState_STATUS{
	"canceled":  ProvisioningState_STATUS_Canceled,
	"creating":  ProvisioningState_STATUS_Creating,
	"deleting":  ProvisioningState_STATUS_Deleting,
	"failed":    ProvisioningState_STATUS_Failed,
	"moving":    ProvisioningState_STATUS_Moving,
	"running":   ProvisioningState_STATUS_Running,
	"succeeded": ProvisioningState_STATUS_Succeeded,
	"unknown":   ProvisioningState_STATUS_Unknown,
	"updating":  ProvisioningState_STATUS_Updating,
}

type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_STATUS_Application     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_STATUS_Key             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_STATUS_ManagedIdentity = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_STATUS_User            = SystemData_CreatedByType_STATUS("User")
)

// Mapping from string to SystemData_CreatedByType_STATUS
var systemData_CreatedByType_STATUS_Values = map[string]SystemData_CreatedByType_STATUS{
	"application":     SystemData_CreatedByType_STATUS_Application,
	"key":             SystemData_CreatedByType_STATUS_Key,
	"managedidentity": SystemData_CreatedByType_STATUS_ManagedIdentity,
	"user":            SystemData_CreatedByType_STATUS_User,
}

type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_STATUS_Application     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_STATUS_Key             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_STATUS_ManagedIdentity = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_User            = SystemData_LastModifiedByType_STATUS("User")
)

// Mapping from string to SystemData_LastModifiedByType_STATUS
var systemData_LastModifiedByType_STATUS_Values = map[string]SystemData_LastModifiedByType_STATUS{
	"application":     SystemData_LastModifiedByType_STATUS_Application,
	"key":             SystemData_LastModifiedByType_STATUS_Key,
	"managedidentity": SystemData_LastModifiedByType_STATUS_ManagedIdentity,
	"user":            SystemData_LastModifiedByType_STATUS_User,
}

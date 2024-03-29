// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180601

import (
	"encoding/json"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// Factory resource type.
type Factory_STATUS_ARM struct {
	AdditionalProperties map[string]v1.JSON `json:"additionalProperties,omitempty"`

	// ETag: Etag identifies change in the resource.
	ETag *string `json:"eTag,omitempty"`

	// Id: The resource identifier.
	Id *string `json:"id,omitempty"`

	// Identity: Managed service identity of the factory.
	Identity *FactoryIdentity_STATUS_ARM `json:"identity,omitempty"`

	// Location: The resource location.
	Location *string `json:"location,omitempty"`

	// Name: The resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the factory.
	Properties *FactoryProperties_STATUS_ARM `json:"properties,omitempty"`

	// Tags: The resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The resource type.
	Type *string `json:"type,omitempty"`
}

// Identity properties of the factory resource.
type FactoryIdentity_STATUS_ARM struct {
	// PrincipalId: The principal id of the identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The client tenant id of the identity.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The identity type.
	Type *FactoryIdentity_Type_STATUS `json:"type,omitempty"`

	// UserAssignedIdentities: List of user assigned identities for the factory.
	UserAssignedIdentities map[string]v1.JSON `json:"userAssignedIdentities,omitempty"`
}

// Factory resource properties.
type FactoryProperties_STATUS_ARM struct {
	// CreateTime: Time the factory was created in ISO8601 format.
	CreateTime *string `json:"createTime,omitempty"`

	// Encryption: Properties to enable Customer Managed Key for the factory.
	Encryption *EncryptionConfiguration_STATUS_ARM `json:"encryption,omitempty"`

	// GlobalParameters: List of parameters for factory.
	GlobalParameters map[string]GlobalParameterSpecification_STATUS_ARM `json:"globalParameters,omitempty"`

	// ProvisioningState: Factory provisioning state, example Succeeded.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for the data factory.
	PublicNetworkAccess *FactoryProperties_PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`

	// PurviewConfiguration: Purview information of the factory.
	PurviewConfiguration *PurviewConfiguration_STATUS_ARM `json:"purviewConfiguration,omitempty"`

	// RepoConfiguration: Git repo information of the factory.
	RepoConfiguration *FactoryRepoConfiguration_STATUS_ARM `json:"repoConfiguration,omitempty"`

	// Version: Version of the factory.
	Version *string `json:"version,omitempty"`
}

// Definition of CMK for the factory.
type EncryptionConfiguration_STATUS_ARM struct {
	// Identity: User assigned identity to use to authenticate to customer's key vault. If not provided Managed Service
	// Identity will be used.
	Identity *CMKIdentityDefinition_STATUS_ARM `json:"identity,omitempty"`

	// KeyName: The name of the key in Azure Key Vault to use as Customer Managed Key.
	KeyName *string `json:"keyName,omitempty"`

	// KeyVersion: The version of the key used for CMK. If not provided, latest version will be used.
	KeyVersion *string `json:"keyVersion,omitempty"`

	// VaultBaseUrl: The url of the Azure Key Vault used for CMK.
	VaultBaseUrl *string `json:"vaultBaseUrl,omitempty"`
}

type FactoryIdentity_Type_STATUS string

const (
	FactoryIdentity_Type_STATUS_SystemAssigned             = FactoryIdentity_Type_STATUS("SystemAssigned")
	FactoryIdentity_Type_STATUS_SystemAssignedUserAssigned = FactoryIdentity_Type_STATUS("SystemAssigned,UserAssigned")
	FactoryIdentity_Type_STATUS_UserAssigned               = FactoryIdentity_Type_STATUS("UserAssigned")
)

// Mapping from string to FactoryIdentity_Type_STATUS
var factoryIdentity_Type_STATUS_Values = map[string]FactoryIdentity_Type_STATUS{
	"systemassigned":              FactoryIdentity_Type_STATUS_SystemAssigned,
	"systemassigned,userassigned": FactoryIdentity_Type_STATUS_SystemAssignedUserAssigned,
	"userassigned":                FactoryIdentity_Type_STATUS_UserAssigned,
}

type FactoryRepoConfiguration_STATUS_ARM struct {
	// FactoryGitHub: Mutually exclusive with all other properties
	FactoryGitHub *FactoryGitHubConfiguration_STATUS_ARM `json:"factoryGitHubConfiguration,omitempty"`

	// FactoryVSTS: Mutually exclusive with all other properties
	FactoryVSTS *FactoryVSTSConfiguration_STATUS_ARM `json:"factoryVSTSConfiguration,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because FactoryRepoConfiguration_STATUS_ARM represents a discriminated union (JSON OneOf)
func (configuration FactoryRepoConfiguration_STATUS_ARM) MarshalJSON() ([]byte, error) {
	if configuration.FactoryGitHub != nil {
		return json.Marshal(configuration.FactoryGitHub)
	}
	if configuration.FactoryVSTS != nil {
		return json.Marshal(configuration.FactoryVSTS)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the FactoryRepoConfiguration_STATUS_ARM
func (configuration *FactoryRepoConfiguration_STATUS_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["type"]
	if discriminator == "FactoryGitHubConfiguration" {
		configuration.FactoryGitHub = &FactoryGitHubConfiguration_STATUS_ARM{}
		return json.Unmarshal(data, configuration.FactoryGitHub)
	}
	if discriminator == "FactoryVSTSConfiguration" {
		configuration.FactoryVSTS = &FactoryVSTSConfiguration_STATUS_ARM{}
		return json.Unmarshal(data, configuration.FactoryVSTS)
	}

	// No error
	return nil
}

// Definition of a single parameter for an entity.
type GlobalParameterSpecification_STATUS_ARM struct {
	// Type: Global Parameter type.
	Type *GlobalParameterSpecification_Type_STATUS `json:"type,omitempty"`

	// Value: Value of parameter.
	Value map[string]v1.JSON `json:"value,omitempty"`
}

// Purview configuration.
type PurviewConfiguration_STATUS_ARM struct {
	// PurviewResourceId: Purview resource id.
	PurviewResourceId *string `json:"purviewResourceId,omitempty"`
}

// Managed Identity used for CMK.
type CMKIdentityDefinition_STATUS_ARM struct {
	// UserAssignedIdentity: The resource id of the user assigned identity to authenticate to customer's key vault.
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}

type FactoryGitHubConfiguration_STATUS_ARM struct {
	// AccountName: Account name.
	AccountName *string `json:"accountName,omitempty"`

	// ClientId: GitHub bring your own app client id.
	ClientId *string `json:"clientId,omitempty"`

	// ClientSecret: GitHub bring your own app client secret information.
	ClientSecret *GitHubClientSecret_STATUS_ARM `json:"clientSecret,omitempty"`

	// CollaborationBranch: Collaboration branch.
	CollaborationBranch *string `json:"collaborationBranch,omitempty"`

	// DisablePublish: Disable manual publish operation in ADF studio to favor automated publish.
	DisablePublish *bool `json:"disablePublish,omitempty"`

	// HostName: GitHub Enterprise host name. For example: `https://github.mydomain.com`
	HostName *string `json:"hostName,omitempty"`

	// LastCommitId: Last commit id.
	LastCommitId *string `json:"lastCommitId,omitempty"`

	// RepositoryName: Repository name.
	RepositoryName *string `json:"repositoryName,omitempty"`

	// RootFolder: Root folder.
	RootFolder *string `json:"rootFolder,omitempty"`

	// Type: Type of repo configuration.
	Type FactoryGitHubConfiguration_Type_STATUS `json:"type,omitempty"`
}

type FactoryVSTSConfiguration_STATUS_ARM struct {
	// AccountName: Account name.
	AccountName *string `json:"accountName,omitempty"`

	// CollaborationBranch: Collaboration branch.
	CollaborationBranch *string `json:"collaborationBranch,omitempty"`

	// DisablePublish: Disable manual publish operation in ADF studio to favor automated publish.
	DisablePublish *bool `json:"disablePublish,omitempty"`

	// LastCommitId: Last commit id.
	LastCommitId *string `json:"lastCommitId,omitempty"`

	// ProjectName: VSTS project name.
	ProjectName *string `json:"projectName,omitempty"`

	// RepositoryName: Repository name.
	RepositoryName *string `json:"repositoryName,omitempty"`

	// RootFolder: Root folder.
	RootFolder *string `json:"rootFolder,omitempty"`

	// TenantId: VSTS tenant id.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: Type of repo configuration.
	Type FactoryVSTSConfiguration_Type_STATUS `json:"type,omitempty"`
}

// Client secret information for factory's bring your own app repository configuration.
type GitHubClientSecret_STATUS_ARM struct {
	// ByoaSecretAkvUrl: Bring your own app client secret AKV URL.
	ByoaSecretAkvUrl *string `json:"byoaSecretAkvUrl,omitempty"`

	// ByoaSecretName: Bring your own app client secret name in AKV.
	ByoaSecretName *string `json:"byoaSecretName,omitempty"`
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180601

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Factory_Spec_ARM struct {
	AdditionalProperties map[string]v1.JSON `json:"additionalProperties,omitempty"`

	// Identity: Managed service identity of the factory.
	Identity *FactoryIdentity_ARM `json:"identity,omitempty"`

	// Location: The resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the factory.
	Properties *FactoryProperties_ARM `json:"properties,omitempty"`

	// Tags: The resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Factory_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-06-01"
func (factory Factory_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (factory *Factory_Spec_ARM) GetName() string {
	return factory.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DataFactory/factories"
func (factory *Factory_Spec_ARM) GetType() string {
	return "Microsoft.DataFactory/factories"
}

// Identity properties of the factory resource.
type FactoryIdentity_ARM struct {
	// Type: The identity type.
	Type                   *FactoryIdentity_Type                      `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

// Factory resource properties.
type FactoryProperties_ARM struct {
	// Encryption: Properties to enable Customer Managed Key for the factory.
	Encryption *EncryptionConfiguration_ARM `json:"encryption,omitempty"`

	// GlobalParameters: List of parameters for factory.
	GlobalParameters map[string]GlobalParameterSpecification_ARM `json:"globalParameters,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for the data factory.
	PublicNetworkAccess *FactoryProperties_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// PurviewConfiguration: Purview information of the factory.
	PurviewConfiguration *PurviewConfiguration_ARM `json:"purviewConfiguration,omitempty"`

	// RepoConfiguration: Git repo information of the factory.
	RepoConfiguration *FactoryRepoConfiguration_ARM `json:"repoConfiguration,omitempty"`
}

// Definition of CMK for the factory.
type EncryptionConfiguration_ARM struct {
	// Identity: User assigned identity to use to authenticate to customer's key vault. If not provided Managed Service
	// Identity will be used.
	Identity *CMKIdentityDefinition_ARM `json:"identity,omitempty"`

	// KeyName: The name of the key in Azure Key Vault to use as Customer Managed Key.
	KeyName *string `json:"keyName,omitempty"`

	// KeyVersion: The version of the key used for CMK. If not provided, latest version will be used.
	KeyVersion *string `json:"keyVersion,omitempty"`

	// VaultBaseUrl: The url of the Azure Key Vault used for CMK.
	VaultBaseUrl *string `json:"vaultBaseUrl,omitempty"`
}

// +kubebuilder:validation:Enum={"SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type FactoryIdentity_Type string

const (
	FactoryIdentity_Type_SystemAssigned             = FactoryIdentity_Type("SystemAssigned")
	FactoryIdentity_Type_SystemAssignedUserAssigned = FactoryIdentity_Type("SystemAssigned,UserAssigned")
	FactoryIdentity_Type_UserAssigned               = FactoryIdentity_Type("UserAssigned")
)

type FactoryRepoConfiguration_ARM struct {
	// FactoryGitHub: Mutually exclusive with all other properties
	FactoryGitHub *FactoryGitHubConfiguration_ARM `json:"factoryGitHubConfiguration,omitempty"`

	// FactoryVSTS: Mutually exclusive with all other properties
	FactoryVSTS *FactoryVSTSConfiguration_ARM `json:"factoryVSTSConfiguration,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because FactoryRepoConfiguration_ARM represents a discriminated union (JSON OneOf)
func (configuration FactoryRepoConfiguration_ARM) MarshalJSON() ([]byte, error) {
	if configuration.FactoryGitHub != nil {
		return json.Marshal(configuration.FactoryGitHub)
	}
	if configuration.FactoryVSTS != nil {
		return json.Marshal(configuration.FactoryVSTS)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the FactoryRepoConfiguration_ARM
func (configuration *FactoryRepoConfiguration_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["type"]
	if discriminator == "FactoryGitHubConfiguration" {
		configuration.FactoryGitHub = &FactoryGitHubConfiguration_ARM{}
		return json.Unmarshal(data, configuration.FactoryGitHub)
	}
	if discriminator == "FactoryVSTSConfiguration" {
		configuration.FactoryVSTS = &FactoryVSTSConfiguration_ARM{}
		return json.Unmarshal(data, configuration.FactoryVSTS)
	}

	// No error
	return nil
}

// Definition of a single parameter for an entity.
type GlobalParameterSpecification_ARM struct {
	// Type: Global Parameter type.
	Type *GlobalParameterSpecification_Type `json:"type,omitempty"`

	// Value: Value of parameter.
	Value map[string]v1.JSON `json:"value,omitempty"`
}

// Purview configuration.
type PurviewConfiguration_ARM struct {
	PurviewResourceId *string `json:"purviewResourceId,omitempty"`
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}

// Managed Identity used for CMK.
type CMKIdentityDefinition_ARM struct {
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}

type FactoryGitHubConfiguration_ARM struct {
	// AccountName: Account name.
	AccountName *string `json:"accountName,omitempty"`

	// ClientId: GitHub bring your own app client id.
	ClientId *string `json:"clientId,omitempty"`

	// ClientSecret: GitHub bring your own app client secret information.
	ClientSecret *GitHubClientSecret_ARM `json:"clientSecret,omitempty"`

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
	Type FactoryGitHubConfiguration_Type `json:"type,omitempty"`
}

type FactoryVSTSConfiguration_ARM struct {
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
	Type FactoryVSTSConfiguration_Type `json:"type,omitempty"`
}

// Client secret information for factory's bring your own app repository configuration.
type GitHubClientSecret_ARM struct {
	// ByoaSecretAkvUrl: Bring your own app client secret AKV URL.
	ByoaSecretAkvUrl *string `json:"byoaSecretAkvUrl,omitempty"`

	// ByoaSecretName: Bring your own app client secret name in AKV.
	ByoaSecretName *string `json:"byoaSecretName,omitempty"`
}

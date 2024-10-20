// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

// The Flux Configuration object returned in Get & Put response.
type FluxConfiguration_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties to create a Flux Configuration resource
	Properties *FluxConfiguration_Properties_STATUS_ARM `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type FluxConfiguration_Properties_STATUS_ARM struct {
	// AzureBlob: Parameters to reconcile to the AzureBlob source kind type.
	AzureBlob *AzureBlobDefinition_STATUS_ARM `json:"azureBlob,omitempty"`

	// Bucket: Parameters to reconcile to the Bucket source kind type.
	Bucket *BucketDefinition_STATUS_ARM `json:"bucket,omitempty"`

	// ComplianceState: Combined status of the Flux Kubernetes resources created by the fluxConfiguration or created by the
	// managed objects.
	ComplianceState *FluxComplianceStateDefinition_STATUS_ARM `json:"complianceState,omitempty"`

	// ConfigurationProtectedSettings: Key-value pairs of protected configuration settings for the configuration
	ConfigurationProtectedSettings map[string]string `json:"configurationProtectedSettings,omitempty"`

	// ErrorMessage: Error message returned to the user in the case of provisioning failure.
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// GitRepository: Parameters to reconcile to the GitRepository source kind type.
	GitRepository *GitRepositoryDefinition_STATUS_ARM `json:"gitRepository,omitempty"`

	// Kustomizations: Array of kustomizations used to reconcile the artifact pulled by the source type on the cluster.
	Kustomizations map[string]KustomizationDefinition_STATUS_ARM `json:"kustomizations,omitempty"`

	// Namespace: The namespace to which this configuration is installed to. Maximum of 253 lower case alphanumeric characters,
	// hyphen and period only.
	Namespace *string `json:"namespace,omitempty"`

	// ProvisioningState: Status of the creation of the fluxConfiguration.
	ProvisioningState *ProvisioningStateDefinition_STATUS_ARM `json:"provisioningState,omitempty"`

	// ReconciliationWaitDuration: Maximum duration to wait for flux configuration reconciliation. E.g PT1H, PT5M, P1D
	ReconciliationWaitDuration *string `json:"reconciliationWaitDuration,omitempty"`

	// RepositoryPublicKey: Public Key associated with this fluxConfiguration (either generated within the cluster or provided
	// by the user).
	RepositoryPublicKey *string `json:"repositoryPublicKey,omitempty"`

	// Scope: Scope at which the operator will be installed.
	Scope *ScopeDefinition_STATUS_ARM `json:"scope,omitempty"`

	// SourceKind: Source Kind to pull the configuration data from.
	SourceKind *SourceKindDefinition_STATUS_ARM `json:"sourceKind,omitempty"`

	// SourceSyncedCommitId: Branch and/or SHA of the source commit synced with the cluster.
	SourceSyncedCommitId *string `json:"sourceSyncedCommitId,omitempty"`

	// SourceUpdatedAt: Datetime the fluxConfiguration synced its source on the cluster.
	SourceUpdatedAt *string `json:"sourceUpdatedAt,omitempty"`

	// StatusUpdatedAt: Datetime the fluxConfiguration synced its status on the cluster with Azure.
	StatusUpdatedAt *string `json:"statusUpdatedAt,omitempty"`

	// Statuses: Statuses of the Flux Kubernetes resources created by the fluxConfiguration or created by the managed objects
	// provisioned by the fluxConfiguration.
	Statuses []ObjectStatusDefinition_STATUS_ARM `json:"statuses,omitempty"`

	// Suspend: Whether this configuration should suspend its reconciliation of its kustomizations and sources.
	Suspend *bool `json:"suspend,omitempty"`

	// WaitForReconciliation: Whether flux configuration deployment should wait for cluster to reconcile the kustomizations.
	WaitForReconciliation *bool `json:"waitForReconciliation,omitempty"`
}

// Parameters to reconcile to the AzureBlob source kind type.
type AzureBlobDefinition_STATUS_ARM struct {
	// ContainerName: The Azure Blob container name to sync from the url endpoint for the flux configuration.
	ContainerName *string `json:"containerName,omitempty"`

	// LocalAuthRef: Name of a local secret on the Kubernetes cluster to use as the authentication secret rather than the
	// managed or user-provided configuration secrets.
	LocalAuthRef *string `json:"localAuthRef,omitempty"`

	// ManagedIdentity: Parameters to authenticate using a Managed Identity.
	ManagedIdentity *ManagedIdentityDefinition_STATUS_ARM `json:"managedIdentity,omitempty"`

	// ServicePrincipal: Parameters to authenticate using Service Principal.
	ServicePrincipal *ServicePrincipalDefinition_STATUS_ARM `json:"servicePrincipal,omitempty"`

	// SyncIntervalInSeconds: The interval at which to re-reconcile the cluster Azure Blob source with the remote.
	SyncIntervalInSeconds *int `json:"syncIntervalInSeconds,omitempty"`

	// TimeoutInSeconds: The maximum time to attempt to reconcile the cluster Azure Blob source with the remote.
	TimeoutInSeconds *int `json:"timeoutInSeconds,omitempty"`

	// Url: The URL to sync for the flux configuration Azure Blob storage account.
	Url *string `json:"url,omitempty"`
}

// Parameters to reconcile to the Bucket source kind type.
type BucketDefinition_STATUS_ARM struct {
	// BucketName: The bucket name to sync from the url endpoint for the flux configuration.
	BucketName *string `json:"bucketName,omitempty"`

	// Insecure: Specify whether to use insecure communication when puling data from the S3 bucket.
	Insecure *bool `json:"insecure,omitempty"`

	// LocalAuthRef: Name of a local secret on the Kubernetes cluster to use as the authentication secret rather than the
	// managed or user-provided configuration secrets.
	LocalAuthRef *string `json:"localAuthRef,omitempty"`

	// SyncIntervalInSeconds: The interval at which to re-reconcile the cluster bucket source with the remote.
	SyncIntervalInSeconds *int `json:"syncIntervalInSeconds,omitempty"`

	// TimeoutInSeconds: The maximum time to attempt to reconcile the cluster bucket source with the remote.
	TimeoutInSeconds *int `json:"timeoutInSeconds,omitempty"`

	// Url: The URL to sync for the flux configuration S3 bucket.
	Url *string `json:"url,omitempty"`
}

// Compliance state of the cluster object.
type FluxComplianceStateDefinition_STATUS_ARM string

const (
	FluxComplianceStateDefinition_STATUS_ARM_Compliant    = FluxComplianceStateDefinition_STATUS_ARM("Compliant")
	FluxComplianceStateDefinition_STATUS_ARM_NonCompliant = FluxComplianceStateDefinition_STATUS_ARM("Non-Compliant")
	FluxComplianceStateDefinition_STATUS_ARM_Pending      = FluxComplianceStateDefinition_STATUS_ARM("Pending")
	FluxComplianceStateDefinition_STATUS_ARM_Suspended    = FluxComplianceStateDefinition_STATUS_ARM("Suspended")
	FluxComplianceStateDefinition_STATUS_ARM_Unknown      = FluxComplianceStateDefinition_STATUS_ARM("Unknown")
)

// Mapping from string to FluxComplianceStateDefinition_STATUS_ARM
var fluxComplianceStateDefinition_STATUS_ARM_Values = map[string]FluxComplianceStateDefinition_STATUS_ARM{
	"compliant":     FluxComplianceStateDefinition_STATUS_ARM_Compliant,
	"non-compliant": FluxComplianceStateDefinition_STATUS_ARM_NonCompliant,
	"pending":       FluxComplianceStateDefinition_STATUS_ARM_Pending,
	"suspended":     FluxComplianceStateDefinition_STATUS_ARM_Suspended,
	"unknown":       FluxComplianceStateDefinition_STATUS_ARM_Unknown,
}

// Parameters to reconcile to the GitRepository source kind type.
type GitRepositoryDefinition_STATUS_ARM struct {
	// HttpsUser: Plaintext HTTPS username used to access private git repositories over HTTPS
	HttpsUser *string `json:"httpsUser,omitempty"`

	// LocalAuthRef: Name of a local secret on the Kubernetes cluster to use as the authentication secret rather than the
	// managed or user-provided configuration secrets.
	LocalAuthRef *string `json:"localAuthRef,omitempty"`

	// RepositoryRef: The source reference for the GitRepository object.
	RepositoryRef *RepositoryRefDefinition_STATUS_ARM `json:"repositoryRef,omitempty"`

	// SshKnownHosts: Base64-encoded known_hosts value containing public SSH keys required to access private git repositories
	// over SSH
	SshKnownHosts *string `json:"sshKnownHosts,omitempty"`

	// SyncIntervalInSeconds: The interval at which to re-reconcile the cluster git repository source with the remote.
	SyncIntervalInSeconds *int `json:"syncIntervalInSeconds,omitempty"`

	// TimeoutInSeconds: The maximum time to attempt to reconcile the cluster git repository source with the remote.
	TimeoutInSeconds *int `json:"timeoutInSeconds,omitempty"`

	// Url: The URL to sync for the flux configuration git repository.
	Url *string `json:"url,omitempty"`
}

// The Kustomization defining how to reconcile the artifact pulled by the source type on the cluster.
type KustomizationDefinition_STATUS_ARM struct {
	// DependsOn: Specifies other Kustomizations that this Kustomization depends on. This Kustomization will not reconcile
	// until all dependencies have completed their reconciliation.
	DependsOn []string `json:"dependsOn,omitempty"`

	// Force: Enable/disable re-creating Kubernetes resources on the cluster when patching fails due to an immutable field
	// change.
	Force *bool `json:"force,omitempty"`

	// Name: Name of the Kustomization, matching the key in the Kustomizations object map.
	Name *string `json:"name,omitempty"`

	// Path: The path in the source reference to reconcile on the cluster.
	Path *string `json:"path,omitempty"`

	// PostBuild: Used for variable substitution for this Kustomization after kustomize build.
	PostBuild *PostBuildDefinition_STATUS_ARM `json:"postBuild,omitempty"`

	// Prune: Enable/disable garbage collections of Kubernetes objects created by this Kustomization.
	Prune *bool `json:"prune,omitempty"`

	// RetryIntervalInSeconds: The interval at which to re-reconcile the Kustomization on the cluster in the event of failure
	// on reconciliation.
	RetryIntervalInSeconds *int `json:"retryIntervalInSeconds,omitempty"`

	// SyncIntervalInSeconds: The interval at which to re-reconcile the Kustomization on the cluster.
	SyncIntervalInSeconds *int `json:"syncIntervalInSeconds,omitempty"`

	// TimeoutInSeconds: The maximum time to attempt to reconcile the Kustomization on the cluster.
	TimeoutInSeconds *int `json:"timeoutInSeconds,omitempty"`

	// Wait: Enable/disable health check for all Kubernetes objects created by this Kustomization.
	Wait *bool `json:"wait,omitempty"`
}

// Statuses of objects deployed by the user-specified kustomizations from the git repository.
type ObjectStatusDefinition_STATUS_ARM struct {
	// AppliedBy: Object reference to the Kustomization that applied this object
	AppliedBy *ObjectReferenceDefinition_STATUS_ARM `json:"appliedBy,omitempty"`

	// ComplianceState: Compliance state of the applied object showing whether the applied object has come into a ready state
	// on the cluster.
	ComplianceState *FluxComplianceStateDefinition_STATUS_ARM `json:"complianceState,omitempty"`

	// HelmReleaseProperties: Additional properties that are provided from objects of the HelmRelease kind
	HelmReleaseProperties *HelmReleasePropertiesDefinition_STATUS_ARM `json:"helmReleaseProperties,omitempty"`

	// Kind: Kind of the applied object
	Kind *string `json:"kind,omitempty"`

	// Name: Name of the applied object
	Name *string `json:"name,omitempty"`

	// Namespace: Namespace of the applied object
	Namespace *string `json:"namespace,omitempty"`

	// StatusConditions: List of Kubernetes object status conditions present on the cluster
	StatusConditions []ObjectStatusConditionDefinition_STATUS_ARM `json:"statusConditions,omitempty"`
}

// Scope at which the configuration will be installed.
type ScopeDefinition_STATUS_ARM string

const (
	ScopeDefinition_STATUS_ARM_Cluster   = ScopeDefinition_STATUS_ARM("cluster")
	ScopeDefinition_STATUS_ARM_Namespace = ScopeDefinition_STATUS_ARM("namespace")
)

// Mapping from string to ScopeDefinition_STATUS_ARM
var scopeDefinition_STATUS_ARM_Values = map[string]ScopeDefinition_STATUS_ARM{
	"cluster":   ScopeDefinition_STATUS_ARM_Cluster,
	"namespace": ScopeDefinition_STATUS_ARM_Namespace,
}

// Source Kind to pull the configuration data from.
type SourceKindDefinition_STATUS_ARM string

const (
	SourceKindDefinition_STATUS_ARM_AzureBlob     = SourceKindDefinition_STATUS_ARM("AzureBlob")
	SourceKindDefinition_STATUS_ARM_Bucket        = SourceKindDefinition_STATUS_ARM("Bucket")
	SourceKindDefinition_STATUS_ARM_GitRepository = SourceKindDefinition_STATUS_ARM("GitRepository")
)

// Mapping from string to SourceKindDefinition_STATUS_ARM
var sourceKindDefinition_STATUS_ARM_Values = map[string]SourceKindDefinition_STATUS_ARM{
	"azureblob":     SourceKindDefinition_STATUS_ARM_AzureBlob,
	"bucket":        SourceKindDefinition_STATUS_ARM_Bucket,
	"gitrepository": SourceKindDefinition_STATUS_ARM_GitRepository,
}

// Properties for HelmRelease objects
type HelmReleasePropertiesDefinition_STATUS_ARM struct {
	// FailureCount: Total number of times that the HelmRelease failed to install or upgrade
	FailureCount *int `json:"failureCount,omitempty"`

	// HelmChartRef: The reference to the HelmChart object used as the source to this HelmRelease
	HelmChartRef *ObjectReferenceDefinition_STATUS_ARM `json:"helmChartRef,omitempty"`

	// InstallFailureCount: Number of times that the HelmRelease failed to install
	InstallFailureCount *int `json:"installFailureCount,omitempty"`

	// LastRevisionApplied: The revision number of the last released object change
	LastRevisionApplied *int `json:"lastRevisionApplied,omitempty"`

	// UpgradeFailureCount: Number of times that the HelmRelease failed to upgrade
	UpgradeFailureCount *int `json:"upgradeFailureCount,omitempty"`
}

// Parameters to authenticate using a Managed Identity.
type ManagedIdentityDefinition_STATUS_ARM struct {
	// ClientId: The client Id for authenticating a Managed Identity.
	ClientId *string `json:"clientId,omitempty"`
}

// Object reference to a Kubernetes object on a cluster
type ObjectReferenceDefinition_STATUS_ARM struct {
	// Name: Name of the object
	Name *string `json:"name,omitempty"`

	// Namespace: Namespace of the object
	Namespace *string `json:"namespace,omitempty"`
}

// Status condition of Kubernetes object
type ObjectStatusConditionDefinition_STATUS_ARM struct {
	// LastTransitionTime: Last time this status condition has changed
	LastTransitionTime *string `json:"lastTransitionTime,omitempty"`

	// Message: A more verbose description of the object status condition
	Message *string `json:"message,omitempty"`

	// Reason: Reason for the specified status condition type status
	Reason *string `json:"reason,omitempty"`

	// Status: Status of the Kubernetes object condition type
	Status *string `json:"status,omitempty"`

	// Type: Object status condition type for this object
	Type *string `json:"type,omitempty"`
}

// The postBuild definitions defining variable substitutions for this Kustomization after kustomize build.
type PostBuildDefinition_STATUS_ARM struct {
	// Substitute: Key/value pairs holding the variables to be substituted in this Kustomization.
	Substitute map[string]string `json:"substitute,omitempty"`

	// SubstituteFrom: Array of ConfigMaps/Secrets from which the variables are substituted for this Kustomization.
	SubstituteFrom []SubstituteFromDefinition_STATUS_ARM `json:"substituteFrom,omitempty"`
}

// The source reference for the GitRepository object.
type RepositoryRefDefinition_STATUS_ARM struct {
	// Branch: The git repository branch name to checkout.
	Branch *string `json:"branch,omitempty"`

	// Commit: The commit SHA to checkout. This value must be combined with the branch name to be valid. This takes precedence
	// over semver.
	Commit *string `json:"commit,omitempty"`

	// Semver: The semver range used to match against git repository tags. This takes precedence over tag.
	Semver *string `json:"semver,omitempty"`

	// Tag: The git repository tag name to checkout. This takes precedence over branch.
	Tag *string `json:"tag,omitempty"`
}

// Parameters to authenticate using Service Principal.
type ServicePrincipalDefinition_STATUS_ARM struct {
	// ClientCertificateSendChain: Specifies whether to include x5c header in client claims when acquiring a token to enable
	// subject name / issuer based authentication for the Client Certificate
	ClientCertificateSendChain *bool `json:"clientCertificateSendChain,omitempty"`

	// ClientId: The client Id for authenticating a Service Principal.
	ClientId *string `json:"clientId,omitempty"`

	// TenantId: The tenant Id for authenticating a Service Principal
	TenantId *string `json:"tenantId,omitempty"`
}

// Array of ConfigMaps/Secrets from which the variables are substituted for this Kustomization.
type SubstituteFromDefinition_STATUS_ARM struct {
	// Kind: Define whether it is ConfigMap or Secret that holds the variables to be used in substitution.
	Kind *string `json:"kind,omitempty"`

	// Name: Name of the ConfigMap/Secret that holds the variables to be used in substitution.
	Name *string `json:"name,omitempty"`

	// Optional: Set to True to proceed without ConfigMap/Secret, if it is not present.
	Optional *bool `json:"optional,omitempty"`
}

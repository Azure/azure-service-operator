// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package common

// =================== Config variable constants ===================

const (
	// #nosec
	ClientSecretVar      = "AZURE_CLIENT_SECRET"
	SubscriptionIDVar    = "AZURE_SUBSCRIPTION_ID"
	TenantIDVar          = "AZURE_TENANT_ID"
	ClientIDVar          = "AZURE_CLIENT_ID"
	ClientCertificateVar = "AZURE_CLIENT_CERTIFICATE"
	// #nosec
	ClientCertificatePasswordVar = "AZURE_CLIENT_CERTIFICATE_PASSWORD"
	TargetNamespacesVar          = "AZURE_TARGET_NAMESPACES"
	OperatorModeVar              = "AZURE_OPERATOR_MODE"
	SyncPeriodVar                = "AZURE_SYNC_PERIOD"
	ResourceManagerEndpointVar   = "AZURE_RESOURCE_MANAGER_ENDPOINT"
	ResourceManagerAudienceVar   = "AZURE_RESOURCE_MANAGER_AUDIENCE"
	AzureAuthorityHostVar        = "AZURE_AUTHORITY_HOST"
	PodNamespaceVar              = "POD_NAMESPACE"
	UseWorkloadIdentityAuth      = "USE_WORKLOAD_IDENTITY_AUTH"
	// #nosec
	FederatedTokenFilePath = "/var/run/secrets/tokens/azure-identity"
)

// =================== MultiTenant Constants ===================

const (
	PerResourceSecretAnnotation = "serviceoperator.azure.com/credential-from"
)

// =================== Reconcile Policy Constants ===================

// ReconcilePolicyAnnotation describes the reconcile policy for the resource in question.
// A reconcile policy describes what action (if any) the operator is allowed to take when
// reconciling the resource.
// If no reconcile policy is specified, the default is "run"
const ReconcilePolicyAnnotation = "serviceoperator.azure.com/reconcile-policy"

const (
	// ReconcilePolicyManage instructs the operator to manage the resource in question.
	// This includes issuing PUTs to update it and DELETE's to delete it from Azure if deleted in Kubernetes.
	// This is the default policy when no policy is specified.
	ReconcilePolicyManage = "manage"

	// ReconcilePolicySkip instructs the operator to skip all reconciliation actions. This includes creating
	// the resource.
	ReconcilePolicySkip = "skip"

	// ReconcilePolicyDetachOnDelete instructs the operator to skip deletion of resources in Azure. This allows
	// deletion of the resource in Kubernetes to go through but does not delete the underlying Azure resource.
	ReconcilePolicyDetachOnDelete = "detach-on-delete"
)

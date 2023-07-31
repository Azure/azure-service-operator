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

// =================== Annotation Constants ===================

const (
	PerResourceSecretAnnotation = "serviceoperator.azure.com/credential-from"
)

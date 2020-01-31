package apis

const (
	// AzureInfraFinalizer is the finalizer label added to Azure Resources which need to reaped before K8s deletion
	AzureInfraFinalizer = "infra.azure.com/finalizer"
)

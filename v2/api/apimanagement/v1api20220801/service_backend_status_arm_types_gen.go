// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

type Service_Backend_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Backend entity contract properties.
	Properties *BackendContractProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Parameters supplied to the Create Backend operation.
type BackendContractProperties_STATUS_ARM struct {
	// Credentials: Backend Credentials Contract Properties
	Credentials *BackendCredentialsContract_STATUS_ARM `json:"credentials,omitempty"`

	// Description: Backend Description.
	Description *string `json:"description,omitempty"`

	// Properties: Backend Properties contract
	Properties *BackendProperties_STATUS_ARM `json:"properties,omitempty"`

	// Protocol: Backend communication protocol.
	Protocol *BackendContractProperties_Protocol_STATUS `json:"protocol,omitempty"`

	// Proxy: Backend gateway Contract Properties
	Proxy *BackendProxyContract_STATUS_ARM `json:"proxy,omitempty"`

	// ResourceId: Management Uri of the Resource in External System. This URL can be the Arm Resource Id of Logic Apps,
	// Function Apps or API Apps.
	ResourceId *string `json:"resourceId,omitempty"`

	// Title: Backend Title.
	Title *string `json:"title,omitempty"`

	// Tls: Backend TLS Properties
	Tls *BackendTlsProperties_STATUS_ARM `json:"tls,omitempty"`

	// Url: Runtime Url of the Backend.
	Url *string `json:"url,omitempty"`
}

// Details of the Credentials used to connect to Backend.
type BackendCredentialsContract_STATUS_ARM struct {
	// Authorization: Authorization header authentication
	Authorization *BackendAuthorizationHeaderCredentials_STATUS_ARM `json:"authorization,omitempty"`

	// Certificate: List of Client Certificate Thumbprints. Will be ignored if certificatesIds are provided.
	Certificate []string `json:"certificate,omitempty"`

	// CertificateIds: List of Client Certificate Ids.
	CertificateIds []string `json:"certificateIds,omitempty"`

	// Header: Header Parameter description.
	Header map[string][]string `json:"header,omitempty"`

	// Query: Query Parameter description.
	Query map[string][]string `json:"query,omitempty"`
}

// Properties specific to the Backend Type.
type BackendProperties_STATUS_ARM struct {
	// ServiceFabricCluster: Backend Service Fabric Cluster Properties
	ServiceFabricCluster *BackendServiceFabricClusterProperties_STATUS_ARM `json:"serviceFabricCluster,omitempty"`
}

// Details of the Backend WebProxy Server to use in the Request to Backend.
type BackendProxyContract_STATUS_ARM struct {
	// Url: WebProxy Server AbsoluteUri property which includes the entire URI stored in the Uri instance, including all
	// fragments and query strings.
	Url *string `json:"url,omitempty"`

	// Username: Username to connect to the WebProxy server
	Username *string `json:"username,omitempty"`
}

// Properties controlling TLS Certificate Validation.
type BackendTlsProperties_STATUS_ARM struct {
	// ValidateCertificateChain: Flag indicating whether SSL certificate chain validation should be done when using self-signed
	// certificates for this backend host.
	ValidateCertificateChain *bool `json:"validateCertificateChain,omitempty"`

	// ValidateCertificateName: Flag indicating whether SSL certificate name validation should be done when using self-signed
	// certificates for this backend host.
	ValidateCertificateName *bool `json:"validateCertificateName,omitempty"`
}

// Authorization header information.
type BackendAuthorizationHeaderCredentials_STATUS_ARM struct {
	// Parameter: Authentication Parameter value.
	Parameter *string `json:"parameter,omitempty"`

	// Scheme: Authentication Scheme name.
	Scheme *string `json:"scheme,omitempty"`
}

// Properties of the Service Fabric Type Backend.
type BackendServiceFabricClusterProperties_STATUS_ARM struct {
	// ClientCertificateId: The client certificate id for the management endpoint.
	ClientCertificateId *string `json:"clientCertificateId,omitempty"`

	// ClientCertificatethumbprint: The client certificate thumbprint for the management endpoint. Will be ignored if
	// certificatesIds are provided
	ClientCertificatethumbprint *string `json:"clientCertificatethumbprint,omitempty"`

	// ManagementEndpoints: The cluster management endpoint.
	ManagementEndpoints []string `json:"managementEndpoints,omitempty"`

	// MaxPartitionResolutionRetries: Maximum number of retries while attempting resolve the partition.
	MaxPartitionResolutionRetries *int `json:"maxPartitionResolutionRetries,omitempty"`

	// ServerCertificateThumbprints: Thumbprints of certificates cluster management service uses for tls communication
	ServerCertificateThumbprints []string `json:"serverCertificateThumbprints,omitempty"`

	// ServerX509Names: Server X509 Certificate Names Collection
	ServerX509Names []X509CertificateName_STATUS_ARM `json:"serverX509Names,omitempty"`
}

// Properties of server X509Names.
type X509CertificateName_STATUS_ARM struct {
	// IssuerCertificateThumbprint: Thumbprint for the Issuer of the Certificate.
	IssuerCertificateThumbprint *string `json:"issuerCertificateThumbprint,omitempty"`

	// Name: Common Name of the Certificate.
	Name *string `json:"name,omitempty"`
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20220301

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Sites_SpecARM struct {
	// ExtendedLocation: Extended Location.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`

	// Identity: Managed service identity.
	Identity *ManagedServiceIdentityARM `json:"identity,omitempty"`

	// Kind: Kind of resource.
	Kind *string `json:"kind,omitempty"`

	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: Unique name of the app to create or update. To create or update a deployment slot, use the {slot} parameter.
	Name string `json:"name,omitempty"`

	// Properties: Site resource specific properties
	Properties *Sites_Spec_PropertiesARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Sites_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-03-01"
func (sites Sites_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (sites *Sites_SpecARM) GetName() string {
	return sites.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Web/sites"
func (sites *Sites_SpecARM) GetType() string {
	return "Microsoft.Web/sites"
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/ManagedServiceIdentity
type ManagedServiceIdentityARM struct {
	// Type: Type of managed service identity.
	Type *ManagedServiceIdentityType `json:"type,omitempty"`

	// UserAssignedIdentities: The list of user assigned identities associated with the resource. The user identity dictionary
	// key references will be ARM resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}
	UserAssignedIdentities map[string]v1.JSON `json:"userAssignedIdentities,omitempty"`
}

type Sites_Spec_PropertiesARM struct {
	// ClientAffinityEnabled: <code>true</code> to enable client affinity; <code>false</code> to stop sending session affinity
	// cookies, which route client requests in the same session to the same instance. Default is <code>true</code>.
	ClientAffinityEnabled *bool `json:"clientAffinityEnabled,omitempty"`

	// ClientCertEnabled: <code>true</code> to enable client certificate authentication (TLS mutual authentication); otherwise,
	// <code>false</code>. Default is <code>false</code>.
	ClientCertEnabled *bool `json:"clientCertEnabled,omitempty"`

	// ClientCertExclusionPaths: client certificate authentication comma-separated exclusion paths
	ClientCertExclusionPaths *string `json:"clientCertExclusionPaths,omitempty"`

	// ClientCertMode: This composes with ClientCertEnabled setting.
	// - ClientCertEnabled: false means ClientCert is ignored.
	// - ClientCertEnabled: true and ClientCertMode: Required means ClientCert is required.
	// - ClientCertEnabled: true and ClientCertMode: Optional means ClientCert is optional or accepted.
	ClientCertMode *SitesSpecPropertiesClientCertMode `json:"clientCertMode,omitempty"`

	// CloningInfo: Information needed for cloning operation.
	CloningInfo *CloningInfoARM `json:"cloningInfo,omitempty"`

	// ContainerSize: Size of the function container.
	ContainerSize *int `json:"containerSize,omitempty"`

	// CustomDomainVerificationId: Unique identifier that verifies the custom domains assigned to the app. Customer will add
	// this id to a txt record for verification.
	CustomDomainVerificationId *string `json:"customDomainVerificationId,omitempty"`

	// DailyMemoryTimeQuota: Maximum allowed daily memory-time quota (applicable on dynamic apps only).
	DailyMemoryTimeQuota *int `json:"dailyMemoryTimeQuota,omitempty"`

	// Enabled: <code>true</code> if the app is enabled; otherwise, <code>false</code>. Setting this value to false disables
	// the app (takes the app offline).
	Enabled *bool `json:"enabled,omitempty"`

	// HostNameSslStates: Hostname SSL states are used to manage the SSL bindings for app's hostnames.
	HostNameSslStates []HostNameSslStateARM `json:"hostNameSslStates,omitempty"`

	// HostNamesDisabled: <code>true</code> to disable the public hostnames of the app; otherwise, <code>false</code>.
	// If <code>true</code>, the app is only accessible via API management process.
	HostNamesDisabled *bool `json:"hostNamesDisabled,omitempty"`

	// HostingEnvironmentProfile: Specification for an App Service Environment to use for this resource.
	HostingEnvironmentProfile *HostingEnvironmentProfileARM `json:"hostingEnvironmentProfile,omitempty"`

	// HttpsOnly: HttpsOnly: configures a web site to accept only https requests. Issues redirect for
	// http requests
	HttpsOnly *bool `json:"httpsOnly,omitempty"`

	// HyperV: Hyper-V sandbox.
	HyperV *bool `json:"hyperV,omitempty"`

	// IsXenon: Obsolete: Hyper-V sandbox.
	IsXenon *bool `json:"isXenon,omitempty"`

	// KeyVaultReferenceIdentity: Identity to use for Key Vault Reference authentication.
	KeyVaultReferenceIdentity *string `json:"keyVaultReferenceIdentity,omitempty"`

	// PublicNetworkAccess: Property to allow or block all public traffic. Allowed Values: 'Enabled', 'Disabled' or an empty
	// string.
	PublicNetworkAccess *string `json:"publicNetworkAccess,omitempty"`

	// RedundancyMode: Site redundancy mode.
	RedundancyMode *SitesSpecPropertiesRedundancyMode `json:"redundancyMode,omitempty"`

	// Reserved: <code>true</code> if reserved; otherwise, <code>false</code>.
	Reserved *bool `json:"reserved,omitempty"`

	// ScmSiteAlsoStopped: <code>true</code> to stop SCM (KUDU) site when the app is stopped; otherwise, <code>false</code>.
	// The default is <code>false</code>.
	ScmSiteAlsoStopped *bool   `json:"scmSiteAlsoStopped,omitempty"`
	ServerFarmId       *string `json:"serverFarmId,omitempty"`

	// SiteConfig: Configuration of an App Service app.
	SiteConfig *Sites_Spec_Properties_SiteConfigARM `json:"siteConfig,omitempty"`

	// StorageAccountRequired: Checks if Customer provided storage account is required
	StorageAccountRequired *bool   `json:"storageAccountRequired,omitempty"`
	VirtualNetworkSubnetId *string `json:"virtualNetworkSubnetId,omitempty"`

	// VnetContentShareEnabled: To enable accessing content over virtual network
	VnetContentShareEnabled *bool `json:"vnetContentShareEnabled,omitempty"`

	// VnetImagePullEnabled: To enable pulling image over Virtual Network
	VnetImagePullEnabled *bool `json:"vnetImagePullEnabled,omitempty"`

	// VnetRouteAllEnabled: Virtual Network Route All enabled. This causes all outbound traffic to have Virtual Network
	// Security Groups and User Defined Routes applied.
	VnetRouteAllEnabled *bool `json:"vnetRouteAllEnabled,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/CloningInfo
type CloningInfoARM struct {
	// AppSettingsOverrides: Application setting overrides for cloned app. If specified, these settings override the settings
	// cloned
	// from source app. Otherwise, application settings from source app are retained.
	AppSettingsOverrides map[string]string `json:"appSettingsOverrides,omitempty"`

	// CloneCustomHostNames: <code>true</code> to clone custom hostnames from source app; otherwise, <code>false</code>.
	CloneCustomHostNames *bool `json:"cloneCustomHostNames,omitempty"`

	// CloneSourceControl: <code>true</code> to clone source control from source app; otherwise, <code>false</code>.
	CloneSourceControl *bool `json:"cloneSourceControl,omitempty"`

	// ConfigureLoadBalancing: <code>true</code> to configure load balancing for source and destination app.
	ConfigureLoadBalancing *bool `json:"configureLoadBalancing,omitempty"`

	// CorrelationId: Correlation ID of cloning operation. This ID ties multiple cloning operations
	// together to use the same snapshot.
	CorrelationId *string `json:"correlationId,omitempty"`

	// HostingEnvironment: App Service Environment.
	HostingEnvironment *string `json:"hostingEnvironment,omitempty"`

	// Overwrite: <code>true</code> to overwrite destination app; otherwise, <code>false</code>.
	Overwrite      *bool   `json:"overwrite,omitempty"`
	SourceWebAppId *string `json:"sourceWebAppId,omitempty"`

	// SourceWebAppLocation: Location of source app ex: West US or North Europe
	SourceWebAppLocation    *string `json:"sourceWebAppLocation,omitempty"`
	TrafficManagerProfileId *string `json:"trafficManagerProfileId,omitempty"`

	// TrafficManagerProfileName: Name of Traffic Manager profile to create. This is only needed if Traffic Manager profile
	// does not already exist.
	TrafficManagerProfileName *string `json:"trafficManagerProfileName,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/HostNameSslState
type HostNameSslStateARM struct {
	// HostType: Indicates whether the hostname is a standard or repository hostname.
	HostType *HostNameSslStateHostType `json:"hostType,omitempty"`

	// Name: Hostname.
	Name *string `json:"name,omitempty"`

	// SslState: SSL type.
	SslState *HostNameSslStateSslState `json:"sslState,omitempty"`

	// Thumbprint: SSL certificate thumbprint.
	Thumbprint *string `json:"thumbprint,omitempty"`

	// ToUpdate: Set to <code>true</code> to update existing hostname.
	ToUpdate *bool `json:"toUpdate,omitempty"`

	// VirtualIP: Virtual IP address assigned to the hostname if IP based SSL is enabled.
	VirtualIP *string `json:"virtualIP,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityType_None                       = ManagedServiceIdentityType("None")
	ManagedServiceIdentityType_SystemAssigned             = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityType_SystemAssignedUserAssigned = ManagedServiceIdentityType("SystemAssigned, UserAssigned")
	ManagedServiceIdentityType_UserAssigned               = ManagedServiceIdentityType("UserAssigned")
)

type Sites_Spec_Properties_SiteConfigARM struct {
	// AcrUseManagedIdentityCreds: Flag to use Managed Identity Creds for ACR pull
	AcrUseManagedIdentityCreds *bool `json:"acrUseManagedIdentityCreds,omitempty"`

	// AcrUserManagedIdentityID: If using user managed identity, the user managed identity ClientId
	AcrUserManagedIdentityID *string `json:"acrUserManagedIdentityID,omitempty"`

	// AlwaysOn: <code>true</code> if Always On is enabled; otherwise, <code>false</code>.
	AlwaysOn *bool `json:"alwaysOn,omitempty"`

	// ApiDefinition: Information about the formal API definition for the app.
	ApiDefinition *ApiDefinitionInfoARM `json:"apiDefinition,omitempty"`

	// ApiManagementConfig: Azure API management (APIM) configuration linked to the app.
	ApiManagementConfig *ApiManagementConfigARM `json:"apiManagementConfig,omitempty"`

	// AppCommandLine: App command line to launch.
	AppCommandLine *string `json:"appCommandLine,omitempty"`

	// AppSettings: Application settings.
	AppSettings []NameValuePairARM `json:"appSettings,omitempty"`

	// AutoHealEnabled: <code>true</code> if Auto Heal is enabled; otherwise, <code>false</code>.
	AutoHealEnabled *bool `json:"autoHealEnabled,omitempty"`

	// AutoHealRules: Rules that can be defined for auto-heal.
	AutoHealRules *AutoHealRulesARM `json:"autoHealRules,omitempty"`

	// AutoSwapSlotName: Auto-swap slot name.
	AutoSwapSlotName *string `json:"autoSwapSlotName,omitempty"`

	// AzureStorageAccounts: List of Azure Storage Accounts.
	AzureStorageAccounts map[string]Sites_Spec_Properties_SiteConfig_AzureStorageAccountsARM `json:"azureStorageAccounts,omitempty"`

	// ConnectionStrings: Connection strings.
	ConnectionStrings []ConnStringInfoARM `json:"connectionStrings,omitempty"`

	// Cors: Cross-Origin Resource Sharing (CORS) settings for the app.
	Cors *CorsSettingsARM `json:"cors,omitempty"`

	// DefaultDocuments: Default documents.
	DefaultDocuments []string `json:"defaultDocuments,omitempty"`

	// DetailedErrorLoggingEnabled: <code>true</code> if detailed error logging is enabled; otherwise, <code>false</code>.
	DetailedErrorLoggingEnabled *bool `json:"detailedErrorLoggingEnabled,omitempty"`

	// DocumentRoot: Document root.
	DocumentRoot *string `json:"documentRoot,omitempty"`

	// Experiments: Routing rules in production experiments.
	Experiments *ExperimentsARM `json:"experiments,omitempty"`

	// FtpsState: State of FTP / FTPS service.
	FtpsState *SitesSpecPropertiesSiteConfigFtpsState `json:"ftpsState,omitempty"`

	// FunctionAppScaleLimit: Maximum number of workers that a site can scale out to.
	// This setting only applies to the Consumption and Elastic Premium Plans
	FunctionAppScaleLimit *int `json:"functionAppScaleLimit,omitempty"`

	// FunctionsRuntimeScaleMonitoringEnabled: Gets or sets a value indicating whether functions runtime scale monitoring is
	// enabled. When enabled,
	// the ScaleController will not monitor event sources directly, but will instead call to the
	// runtime to get scale status.
	FunctionsRuntimeScaleMonitoringEnabled *bool `json:"functionsRuntimeScaleMonitoringEnabled,omitempty"`

	// HandlerMappings: Handler mappings.
	HandlerMappings []HandlerMappingARM `json:"handlerMappings,omitempty"`

	// HealthCheckPath: Health check path
	HealthCheckPath *string `json:"healthCheckPath,omitempty"`

	// Http20Enabled: Http20Enabled: configures a web site to allow clients to connect over http2.0
	Http20Enabled *bool `json:"http20Enabled,omitempty"`

	// HttpLoggingEnabled: <code>true</code> if HTTP logging is enabled; otherwise, <code>false</code>.
	HttpLoggingEnabled *bool `json:"httpLoggingEnabled,omitempty"`

	// IpSecurityRestrictions: IP security restrictions for main.
	IpSecurityRestrictions []IpSecurityRestrictionARM `json:"ipSecurityRestrictions,omitempty"`

	// JavaContainer: Java container.
	JavaContainer *string `json:"javaContainer,omitempty"`

	// JavaContainerVersion: Java container version.
	JavaContainerVersion *string `json:"javaContainerVersion,omitempty"`

	// JavaVersion: Java version.
	JavaVersion *string `json:"javaVersion,omitempty"`

	// KeyVaultReferenceIdentity: Identity to use for Key Vault Reference authentication.
	KeyVaultReferenceIdentity *string `json:"keyVaultReferenceIdentity,omitempty"`

	// Limits: Metric limits set on an app.
	Limits *SiteLimitsARM `json:"limits,omitempty"`

	// LinuxFxVersion: Linux App Framework and version
	LinuxFxVersion *string `json:"linuxFxVersion,omitempty"`

	// LoadBalancing: Site load balancing.
	LoadBalancing *SitesSpecPropertiesSiteConfigLoadBalancing `json:"loadBalancing,omitempty"`

	// LocalMySqlEnabled: <code>true</code> to enable local MySQL; otherwise, <code>false</code>.
	LocalMySqlEnabled *bool `json:"localMySqlEnabled,omitempty"`

	// LogsDirectorySizeLimit: HTTP logs directory size limit.
	LogsDirectorySizeLimit *int `json:"logsDirectorySizeLimit,omitempty"`

	// ManagedPipelineMode: Managed pipeline mode.
	ManagedPipelineMode *SitesSpecPropertiesSiteConfigManagedPipelineMode `json:"managedPipelineMode,omitempty"`

	// ManagedServiceIdentityId: Managed Service Identity Id
	ManagedServiceIdentityId *int `json:"managedServiceIdentityId,omitempty"`

	// MinTlsVersion: MinTlsVersion: configures the minimum version of TLS required for SSL requests.
	MinTlsVersion *SitesSpecPropertiesSiteConfigMinTlsVersion `json:"minTlsVersion,omitempty"`

	// MinimumElasticInstanceCount: Number of minimum instance count for a site
	// This setting only applies to the Elastic Plans
	MinimumElasticInstanceCount *int `json:"minimumElasticInstanceCount,omitempty"`

	// NetFrameworkVersion: .NET Framework version.
	NetFrameworkVersion *string `json:"netFrameworkVersion,omitempty"`

	// NodeVersion: Version of Node.js.
	NodeVersion *string `json:"nodeVersion,omitempty"`

	// NumberOfWorkers: Number of workers.
	NumberOfWorkers *int `json:"numberOfWorkers,omitempty"`

	// PhpVersion: Version of PHP.
	PhpVersion *string `json:"phpVersion,omitempty"`

	// PowerShellVersion: Version of PowerShell.
	PowerShellVersion *string `json:"powerShellVersion,omitempty"`

	// PreWarmedInstanceCount: Number of preWarmed instances.
	// This setting only applies to the Consumption and Elastic Plans
	PreWarmedInstanceCount *int `json:"preWarmedInstanceCount,omitempty"`

	// PublicNetworkAccess: Property to allow or block all public traffic.
	PublicNetworkAccess *string `json:"publicNetworkAccess,omitempty"`

	// PublishingUsername: Publishing user name.
	PublishingUsername *string `json:"publishingUsername,omitempty"`

	// Push: Push settings for the App.
	Push *Sites_Spec_Properties_SiteConfig_PushARM `json:"push,omitempty"`

	// PythonVersion: Version of Python.
	PythonVersion *string `json:"pythonVersion,omitempty"`

	// RemoteDebuggingEnabled: <code>true</code> if remote debugging is enabled; otherwise, <code>false</code>.
	RemoteDebuggingEnabled *bool `json:"remoteDebuggingEnabled,omitempty"`

	// RemoteDebuggingVersion: Remote debugging version.
	RemoteDebuggingVersion *string `json:"remoteDebuggingVersion,omitempty"`

	// RequestTracingEnabled: <code>true</code> if request tracing is enabled; otherwise, <code>false</code>.
	RequestTracingEnabled *bool `json:"requestTracingEnabled,omitempty"`

	// RequestTracingExpirationTime: Request tracing expiration time.
	RequestTracingExpirationTime *string `json:"requestTracingExpirationTime,omitempty"`

	// ScmIpSecurityRestrictions: IP security restrictions for scm.
	ScmIpSecurityRestrictions []IpSecurityRestrictionARM `json:"scmIpSecurityRestrictions,omitempty"`

	// ScmIpSecurityRestrictionsUseMain: IP security restrictions for scm to use main.
	ScmIpSecurityRestrictionsUseMain *bool `json:"scmIpSecurityRestrictionsUseMain,omitempty"`

	// ScmMinTlsVersion: ScmMinTlsVersion: configures the minimum version of TLS required for SSL requests for SCM site.
	ScmMinTlsVersion *SitesSpecPropertiesSiteConfigScmMinTlsVersion `json:"scmMinTlsVersion,omitempty"`

	// ScmType: SCM type.
	ScmType *SitesSpecPropertiesSiteConfigScmType `json:"scmType,omitempty"`

	// TracingOptions: Tracing options.
	TracingOptions *string `json:"tracingOptions,omitempty"`

	// Use32BitWorkerProcess: <code>true</code> to use 32-bit worker process; otherwise, <code>false</code>.
	Use32BitWorkerProcess *bool `json:"use32BitWorkerProcess,omitempty"`

	// VirtualApplications: Virtual applications.
	VirtualApplications []VirtualApplicationARM `json:"virtualApplications,omitempty"`

	// VnetName: Virtual Network name.
	VnetName *string `json:"vnetName,omitempty"`

	// VnetPrivatePortsCount: The number of private ports assigned to this app. These will be assigned dynamically on runtime.
	VnetPrivatePortsCount *int `json:"vnetPrivatePortsCount,omitempty"`

	// VnetRouteAllEnabled: Virtual Network Route All enabled. This causes all outbound traffic to have Virtual Network
	// Security Groups and User Defined Routes applied.
	VnetRouteAllEnabled *bool `json:"vnetRouteAllEnabled,omitempty"`

	// WebSocketsEnabled: <code>true</code> if WebSocket is enabled; otherwise, <code>false</code>.
	WebSocketsEnabled *bool `json:"webSocketsEnabled,omitempty"`

	// WebsiteTimeZone: Sets the time zone a site uses for generating timestamps. Compatible with Linux and Windows App
	// Service. Setting the WEBSITE_TIME_ZONE app setting takes precedence over this config. For Linux, expects tz database
	// values https://www.iana.org/time-zones (for a quick reference see
	// https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). For Windows, expects one of the time zones listed under
	// HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones
	WebsiteTimeZone *string `json:"websiteTimeZone,omitempty"`

	// WindowsFxVersion: Xenon App Framework and version
	WindowsFxVersion *string `json:"windowsFxVersion,omitempty"`

	// XManagedServiceIdentityId: Explicit Managed Service Identity Id
	XManagedServiceIdentityId *int `json:"xManagedServiceIdentityId,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/ApiDefinitionInfo
type ApiDefinitionInfoARM struct {
	// Url: The URL of the API definition.
	Url *string `json:"url,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/ApiManagementConfig
type ApiManagementConfigARM struct {
	Id *string `json:"id,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/AutoHealRules
type AutoHealRulesARM struct {
	// Actions: Actions which to take by the auto-heal module when a rule is triggered.
	Actions *AutoHealActionsARM `json:"actions,omitempty"`

	// Triggers: Triggers for auto-heal.
	Triggers *AutoHealTriggersARM `json:"triggers,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/ConnStringInfo
type ConnStringInfoARM struct {
	// ConnectionString: Connection string value.
	ConnectionString *string `json:"connectionString,omitempty"`

	// Name: Name of connection string.
	Name *string `json:"name,omitempty"`

	// Type: Type of database.
	Type *ConnStringInfoType `json:"type,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/CorsSettings
type CorsSettingsARM struct {
	// AllowedOrigins: Gets or sets the list of origins that should be allowed to make cross-origin
	// calls (for example: http://example.com:12345). Use "*" to allow all.
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`

	// SupportCredentials: Gets or sets whether CORS requests with credentials are allowed. See
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS#Requests_with_credentials
	// for more details.
	SupportCredentials *bool `json:"supportCredentials,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/Experiments
type ExperimentsARM struct {
	// RampUpRules: List of ramp-up rules.
	RampUpRules []RampUpRuleARM `json:"rampUpRules,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/HandlerMapping
type HandlerMappingARM struct {
	// Arguments: Command-line arguments to be passed to the script processor.
	Arguments *string `json:"arguments,omitempty"`

	// Extension: Requests with this extension will be handled using the specified FastCGI application.
	Extension *string `json:"extension,omitempty"`

	// ScriptProcessor: The absolute path to the FastCGI application.
	ScriptProcessor *string `json:"scriptProcessor,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/IpSecurityRestriction
type IpSecurityRestrictionARM struct {
	// Action: Allow or Deny access for this IP range.
	Action *string `json:"action,omitempty"`

	// Description: IP restriction rule description.
	Description *string `json:"description,omitempty"`

	// Headers: IP restriction rule headers.
	// X-Forwarded-Host (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-Host#Examples).
	// The matching logic is ..
	// - If the property is null or empty (default), all hosts(or lack of) are allowed.
	// - A value is compared using ordinal-ignore-case (excluding port number).
	// - Subdomain wildcards are permitted but don't match the root domain. For example, *.contoso.com matches the subdomain
	// foo.contoso.com
	// but not the root domain contoso.com or multi-level foo.bar.contoso.com
	// - Unicode host names are allowed but are converted to Punycode for matching.
	// X-Forwarded-For (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For#Examples).
	// The matching logic is ..
	// - If the property is null or empty (default), any forwarded-for chains (or lack of) are allowed.
	// - If any address (excluding port number) in the chain (comma separated) matches the CIDR defined by the property.
	// X-Azure-FDID and X-FD-HealthProbe.
	// The matching logic is exact match.
	Headers map[string][]string `json:"headers,omitempty"`

	// IpAddress: IP address the security restriction is valid for.
	// It can be in form of pure ipv4 address (required SubnetMask property) or
	// CIDR notation such as ipv4/mask (leading bit match). For CIDR,
	// SubnetMask property must not be specified.
	IpAddress *string `json:"ipAddress,omitempty"`

	// Name: IP restriction rule name.
	Name *string `json:"name,omitempty"`

	// Priority: Priority of IP restriction rule.
	Priority *int `json:"priority,omitempty"`

	// SubnetMask: Subnet mask for the range of IP addresses the restriction is valid for.
	SubnetMask *string `json:"subnetMask,omitempty"`

	// SubnetTrafficTag: (internal) Subnet traffic tag
	SubnetTrafficTag *int `json:"subnetTrafficTag,omitempty"`

	// Tag: Defines what this IP filter will be used for. This is to support IP filtering on proxies.
	Tag                  *IpSecurityRestrictionTag `json:"tag,omitempty"`
	VnetSubnetResourceId *string                   `json:"vnetSubnetResourceId,omitempty"`

	// VnetTrafficTag: (internal) Vnet traffic tag
	VnetTrafficTag *int `json:"vnetTrafficTag,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/NameValuePair
type NameValuePairARM struct {
	// Name: Pair name.
	Name *string `json:"name,omitempty"`

	// Value: Pair value.
	Value *string `json:"value,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/SiteLimits
type SiteLimitsARM struct {
	// MaxDiskSizeInMb: Maximum allowed disk size usage in MB.
	MaxDiskSizeInMb *int `json:"maxDiskSizeInMb,omitempty"`

	// MaxMemoryInMb: Maximum allowed memory usage in MB.
	MaxMemoryInMb *int `json:"maxMemoryInMb,omitempty"`

	// MaxPercentageCpu: Maximum allowed CPU usage percentage.
	MaxPercentageCpu *float64 `json:"maxPercentageCpu,omitempty"`
}

type Sites_Spec_Properties_SiteConfig_AzureStorageAccountsARM struct {
	// AccessKey: Access key for the storage account.
	AccessKey *string `json:"accessKey,omitempty"`

	// AccountName: Name of the storage account.
	AccountName *string `json:"accountName,omitempty"`

	// MountPath: Path to mount the storage within the site's runtime environment.
	MountPath *string `json:"mountPath,omitempty"`

	// ShareName: Name of the file share (container name, for Blob storage).
	ShareName *string `json:"shareName,omitempty"`

	// Type: Type of storage.
	Type *SitesSpecPropertiesSiteConfigAzureStorageAccountsType `json:"type,omitempty"`
}

type Sites_Spec_Properties_SiteConfig_PushARM struct {
	// Kind: Kind of resource.
	Kind *string `json:"kind,omitempty"`

	// Properties: PushSettings resource specific properties
	Properties *PushSettingsPropertiesARM `json:"properties,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/VirtualApplication
type VirtualApplicationARM struct {
	// PhysicalPath: Physical path.
	PhysicalPath *string `json:"physicalPath,omitempty"`

	// PreloadEnabled: <code>true</code> if preloading is enabled; otherwise, <code>false</code>.
	PreloadEnabled *bool `json:"preloadEnabled,omitempty"`

	// VirtualDirectories: Virtual directories for virtual application.
	VirtualDirectories []VirtualDirectoryARM `json:"virtualDirectories,omitempty"`

	// VirtualPath: Virtual path.
	VirtualPath *string `json:"virtualPath,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/AutoHealActions
type AutoHealActionsARM struct {
	// ActionType: Predefined action to be taken.
	ActionType *AutoHealActionsActionType `json:"actionType,omitempty"`

	// CustomAction: Custom action to be executed
	// when an auto heal rule is triggered.
	CustomAction *AutoHealCustomActionARM `json:"customAction,omitempty"`

	// MinProcessExecutionTime: Minimum time the process must execute
	// before taking the action
	MinProcessExecutionTime *string `json:"minProcessExecutionTime,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/AutoHealTriggers
type AutoHealTriggersARM struct {
	// PrivateBytesInKB: A rule based on private bytes.
	PrivateBytesInKB *int `json:"privateBytesInKB,omitempty"`

	// Requests: Trigger based on total requests.
	Requests *RequestsBasedTriggerARM `json:"requests,omitempty"`

	// SlowRequests: Trigger based on request execution time.
	SlowRequests *SlowRequestsBasedTriggerARM `json:"slowRequests,omitempty"`

	// SlowRequestsWithPath: A rule based on multiple Slow Requests Rule with path
	SlowRequestsWithPath []SlowRequestsBasedTriggerARM `json:"slowRequestsWithPath,omitempty"`

	// StatusCodes: A rule based on status codes.
	StatusCodes []StatusCodesBasedTriggerARM `json:"statusCodes,omitempty"`

	// StatusCodesRange: A rule based on status codes ranges.
	StatusCodesRange []StatusCodesRangeBasedTriggerARM `json:"statusCodesRange,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/PushSettingsProperties
type PushSettingsPropertiesARM struct {
	// DynamicTagsJson: Gets or sets a JSON string containing a list of dynamic tags that will be evaluated from user claims in
	// the push registration endpoint.
	DynamicTagsJson *string `json:"dynamicTagsJson,omitempty"`

	// IsPushEnabled: Gets or sets a flag indicating whether the Push endpoint is enabled.
	IsPushEnabled *bool `json:"isPushEnabled,omitempty"`

	// TagWhitelistJson: Gets or sets a JSON string containing a list of tags that are whitelisted for use by the push
	// registration endpoint.
	TagWhitelistJson *string `json:"tagWhitelistJson,omitempty"`

	// TagsRequiringAuth: Gets or sets a JSON string containing a list of tags that require user authentication to be used in
	// the push registration endpoint.
	// Tags can consist of alphanumeric characters and the following:
	// '_', '@', '#', '.', ':', '-'.
	// Validation should be performed at the PushRequestHandler.
	TagsRequiringAuth *string `json:"tagsRequiringAuth,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/RampUpRule
type RampUpRuleARM struct {
	// ActionHostName: Hostname of a slot to which the traffic will be redirected if decided to. E.g.
	// myapp-stage.azurewebsites.net.
	ActionHostName *string `json:"actionHostName,omitempty"`

	// ChangeDecisionCallbackUrl: Custom decision algorithm can be provided in TiPCallback site extension which URL can be
	// specified. See TiPCallback site extension for the scaffold and contracts.
	// https://www.siteextensions.net/packages/TiPCallback/
	ChangeDecisionCallbackUrl *string `json:"changeDecisionCallbackUrl,omitempty"`

	// ChangeIntervalInMinutes: Specifies interval in minutes to reevaluate ReroutePercentage.
	ChangeIntervalInMinutes *int `json:"changeIntervalInMinutes,omitempty"`

	// ChangeStep: In auto ramp up scenario this is the step to add/remove from <code>ReroutePercentage</code> until it reaches
	// \n<code>MinReroutePercentage</code> or
	// <code>MaxReroutePercentage</code>. Site metrics are checked every N minutes specified in
	// <code>ChangeIntervalInMinutes</code>.\nCustom decision algorithm
	// can be provided in TiPCallback site extension which URL can be specified in <code>ChangeDecisionCallbackUrl</code>.
	ChangeStep *float64 `json:"changeStep,omitempty"`

	// MaxReroutePercentage: Specifies upper boundary below which ReroutePercentage will stay.
	MaxReroutePercentage *float64 `json:"maxReroutePercentage,omitempty"`

	// MinReroutePercentage: Specifies lower boundary above which ReroutePercentage will stay.
	MinReroutePercentage *float64 `json:"minReroutePercentage,omitempty"`

	// Name: Name of the routing rule. The recommended name would be to point to the slot which will receive the traffic in the
	// experiment.
	Name *string `json:"name,omitempty"`

	// ReroutePercentage: Percentage of the traffic which will be redirected to <code>ActionHostName</code>.
	ReroutePercentage *float64 `json:"reroutePercentage,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/VirtualDirectory
type VirtualDirectoryARM struct {
	// PhysicalPath: Physical path.
	PhysicalPath *string `json:"physicalPath,omitempty"`

	// VirtualPath: Path to virtual application.
	VirtualPath *string `json:"virtualPath,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/AutoHealCustomAction
type AutoHealCustomActionARM struct {
	// Exe: Executable to be run.
	Exe *string `json:"exe,omitempty"`

	// Parameters: Parameters for the executable.
	Parameters *string `json:"parameters,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/RequestsBasedTrigger
type RequestsBasedTriggerARM struct {
	// Count: Request Count.
	Count *int `json:"count,omitempty"`

	// TimeInterval: Time interval.
	TimeInterval *string `json:"timeInterval,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/SlowRequestsBasedTrigger
type SlowRequestsBasedTriggerARM struct {
	// Count: Request Count.
	Count *int `json:"count,omitempty"`

	// Path: Request Path.
	Path *string `json:"path,omitempty"`

	// TimeInterval: Time interval.
	TimeInterval *string `json:"timeInterval,omitempty"`

	// TimeTaken: Time taken.
	TimeTaken *string `json:"timeTaken,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/StatusCodesBasedTrigger
type StatusCodesBasedTriggerARM struct {
	// Count: Request Count.
	Count *int `json:"count,omitempty"`

	// Path: Request Path
	Path *string `json:"path,omitempty"`

	// Status: HTTP status code.
	Status *int `json:"status,omitempty"`

	// SubStatus: Request Sub Status.
	SubStatus *int `json:"subStatus,omitempty"`

	// TimeInterval: Time interval.
	TimeInterval *string `json:"timeInterval,omitempty"`

	// Win32Status: Win32 error code.
	Win32Status *int `json:"win32Status,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2022-03-01/Microsoft.Web.json#/definitions/StatusCodesRangeBasedTrigger
type StatusCodesRangeBasedTriggerARM struct {
	// Count: Request Count.
	Count *int    `json:"count,omitempty"`
	Path  *string `json:"path,omitempty"`

	// StatusCodes: HTTP status code.
	StatusCodes *string `json:"statusCodes,omitempty"`

	// TimeInterval: Time interval.
	TimeInterval *string `json:"timeInterval,omitempty"`
}

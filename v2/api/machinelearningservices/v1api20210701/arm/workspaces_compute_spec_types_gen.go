// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type WorkspacesCompute_Spec struct {
	// Identity: The identity of the resource.
	Identity *Identity `json:"identity,omitempty"`

	// Location: Specifies the location of the resource.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Compute properties
	Properties *Compute `json:"properties,omitempty"`

	// Sku: The sku of the workspace.
	Sku *Sku `json:"sku,omitempty"`

	// SystemData: System data
	SystemData *SystemData `json:"systemData,omitempty"`

	// Tags: Contains resource tags defined as key/value pairs.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &WorkspacesCompute_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-07-01"
func (compute WorkspacesCompute_Spec) GetAPIVersion() string {
	return "2021-07-01"
}

// GetName returns the Name of the resource
func (compute *WorkspacesCompute_Spec) GetName() string {
	return compute.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.MachineLearningServices/workspaces/computes"
func (compute *WorkspacesCompute_Spec) GetType() string {
	return "Microsoft.MachineLearningServices/workspaces/computes"
}

type Compute struct {
	// AKS: Mutually exclusive with all other properties
	AKS *AKS `json:"aks,omitempty"`

	// AmlCompute: Mutually exclusive with all other properties
	AmlCompute *AmlCompute `json:"amlCompute,omitempty"`

	// ComputeInstance: Mutually exclusive with all other properties
	ComputeInstance *ComputeInstance `json:"computeInstance,omitempty"`

	// DataFactory: Mutually exclusive with all other properties
	DataFactory *DataFactory `json:"dataFactory,omitempty"`

	// DataLakeAnalytics: Mutually exclusive with all other properties
	DataLakeAnalytics *DataLakeAnalytics `json:"dataLakeAnalytics,omitempty"`

	// Databricks: Mutually exclusive with all other properties
	Databricks *Databricks `json:"databricks,omitempty"`

	// HDInsight: Mutually exclusive with all other properties
	HDInsight *HDInsight `json:"hdInsight,omitempty"`

	// Kubernetes: Mutually exclusive with all other properties
	Kubernetes *Kubernetes `json:"kubernetes,omitempty"`

	// SynapseSpark: Mutually exclusive with all other properties
	SynapseSpark *SynapseSpark `json:"synapseSpark,omitempty"`

	// VirtualMachine: Mutually exclusive with all other properties
	VirtualMachine *VirtualMachine `json:"virtualMachine,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because Compute represents a discriminated union (JSON OneOf)
func (compute Compute) MarshalJSON() ([]byte, error) {
	if compute.AKS != nil {
		return json.Marshal(compute.AKS)
	}

	if compute.AmlCompute != nil {
		return json.Marshal(compute.AmlCompute)
	}

	if compute.ComputeInstance != nil {
		return json.Marshal(compute.ComputeInstance)
	}

	if compute.DataFactory != nil {
		return json.Marshal(compute.DataFactory)
	}

	if compute.DataLakeAnalytics != nil {
		return json.Marshal(compute.DataLakeAnalytics)
	}

	if compute.Databricks != nil {
		return json.Marshal(compute.Databricks)
	}

	if compute.HDInsight != nil {
		return json.Marshal(compute.HDInsight)
	}

	if compute.Kubernetes != nil {
		return json.Marshal(compute.Kubernetes)
	}

	if compute.SynapseSpark != nil {
		return json.Marshal(compute.SynapseSpark)
	}

	if compute.VirtualMachine != nil {
		return json.Marshal(compute.VirtualMachine)
	}

	return nil, nil
}

// UnmarshalJSON unmarshals the Compute
func (compute *Compute) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["computeType"]
	if discriminator == "AKS" {
		compute.AKS = &AKS{}
		return json.Unmarshal(data, compute.AKS)
	}
	if discriminator == "AmlCompute" {
		compute.AmlCompute = &AmlCompute{}
		return json.Unmarshal(data, compute.AmlCompute)
	}
	if discriminator == "ComputeInstance" {
		compute.ComputeInstance = &ComputeInstance{}
		return json.Unmarshal(data, compute.ComputeInstance)
	}
	if discriminator == "DataFactory" {
		compute.DataFactory = &DataFactory{}
		return json.Unmarshal(data, compute.DataFactory)
	}
	if discriminator == "DataLakeAnalytics" {
		compute.DataLakeAnalytics = &DataLakeAnalytics{}
		return json.Unmarshal(data, compute.DataLakeAnalytics)
	}
	if discriminator == "Databricks" {
		compute.Databricks = &Databricks{}
		return json.Unmarshal(data, compute.Databricks)
	}
	if discriminator == "HDInsight" {
		compute.HDInsight = &HDInsight{}
		return json.Unmarshal(data, compute.HDInsight)
	}
	if discriminator == "Kubernetes" {
		compute.Kubernetes = &Kubernetes{}
		return json.Unmarshal(data, compute.Kubernetes)
	}
	if discriminator == "SynapseSpark" {
		compute.SynapseSpark = &SynapseSpark{}
		return json.Unmarshal(data, compute.SynapseSpark)
	}
	if discriminator == "VirtualMachine" {
		compute.VirtualMachine = &VirtualMachine{}
		return json.Unmarshal(data, compute.VirtualMachine)
	}

	// No error
	return nil
}

type AKS struct {
	// ComputeLocation: Location for the underlying compute
	ComputeLocation *string `json:"computeLocation,omitempty"`

	// ComputeType: The type of compute
	ComputeType AKS_ComputeType `json:"computeType,omitempty"`

	// Description: The description of the Machine Learning compute.
	Description *string `json:"description,omitempty"`

	// DisableLocalAuth: Opt-out of local authentication and ensure customers can use only MSI and AAD exclusively for
	// authentication.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// Properties: AKS properties
	Properties *AKS_Properties `json:"properties,omitempty"`
	ResourceId *string         `json:"resourceId,omitempty"`
}

type AmlCompute struct {
	// ComputeLocation: Location for the underlying compute
	ComputeLocation *string `json:"computeLocation,omitempty"`

	// ComputeType: The type of compute
	ComputeType AmlCompute_ComputeType `json:"computeType,omitempty"`

	// Description: The description of the Machine Learning compute.
	Description *string `json:"description,omitempty"`

	// DisableLocalAuth: Opt-out of local authentication and ensure customers can use only MSI and AAD exclusively for
	// authentication.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// Properties: Properties of AmlCompute
	Properties *AmlComputeProperties `json:"properties,omitempty"`
	ResourceId *string               `json:"resourceId,omitempty"`
}

type ComputeInstance struct {
	// ComputeLocation: Location for the underlying compute
	ComputeLocation *string `json:"computeLocation,omitempty"`

	// ComputeType: The type of compute
	ComputeType ComputeInstance_ComputeType `json:"computeType,omitempty"`

	// Description: The description of the Machine Learning compute.
	Description *string `json:"description,omitempty"`

	// DisableLocalAuth: Opt-out of local authentication and ensure customers can use only MSI and AAD exclusively for
	// authentication.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// Properties: Properties of ComputeInstance
	Properties *ComputeInstanceProperties `json:"properties,omitempty"`
	ResourceId *string                    `json:"resourceId,omitempty"`
}

type Databricks struct {
	// ComputeLocation: Location for the underlying compute
	ComputeLocation *string `json:"computeLocation,omitempty"`

	// ComputeType: The type of compute
	ComputeType Databricks_ComputeType `json:"computeType,omitempty"`

	// Description: The description of the Machine Learning compute.
	Description *string `json:"description,omitempty"`

	// DisableLocalAuth: Opt-out of local authentication and ensure customers can use only MSI and AAD exclusively for
	// authentication.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// Properties: Properties of Databricks
	Properties *DatabricksProperties `json:"properties,omitempty"`
	ResourceId *string               `json:"resourceId,omitempty"`
}

type DataFactory struct {
	// ComputeLocation: Location for the underlying compute
	ComputeLocation *string `json:"computeLocation,omitempty"`

	// ComputeType: The type of compute
	ComputeType DataFactory_ComputeType `json:"computeType,omitempty"`

	// Description: The description of the Machine Learning compute.
	Description *string `json:"description,omitempty"`

	// DisableLocalAuth: Opt-out of local authentication and ensure customers can use only MSI and AAD exclusively for
	// authentication.
	DisableLocalAuth *bool   `json:"disableLocalAuth,omitempty"`
	ResourceId       *string `json:"resourceId,omitempty"`
}

type DataLakeAnalytics struct {
	// ComputeLocation: Location for the underlying compute
	ComputeLocation *string `json:"computeLocation,omitempty"`

	// ComputeType: The type of compute
	ComputeType DataLakeAnalytics_ComputeType `json:"computeType,omitempty"`

	// Description: The description of the Machine Learning compute.
	Description *string `json:"description,omitempty"`

	// DisableLocalAuth: Opt-out of local authentication and ensure customers can use only MSI and AAD exclusively for
	// authentication.
	DisableLocalAuth *bool                         `json:"disableLocalAuth,omitempty"`
	Properties       *DataLakeAnalytics_Properties `json:"properties,omitempty"`
	ResourceId       *string                       `json:"resourceId,omitempty"`
}

type HDInsight struct {
	// ComputeLocation: Location for the underlying compute
	ComputeLocation *string `json:"computeLocation,omitempty"`

	// ComputeType: The type of compute
	ComputeType HDInsight_ComputeType `json:"computeType,omitempty"`

	// Description: The description of the Machine Learning compute.
	Description *string `json:"description,omitempty"`

	// DisableLocalAuth: Opt-out of local authentication and ensure customers can use only MSI and AAD exclusively for
	// authentication.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// Properties: HDInsight compute properties
	Properties *HDInsightProperties `json:"properties,omitempty"`
	ResourceId *string              `json:"resourceId,omitempty"`
}

type Kubernetes struct {
	// ComputeLocation: Location for the underlying compute
	ComputeLocation *string `json:"computeLocation,omitempty"`

	// ComputeType: The type of compute
	ComputeType Kubernetes_ComputeType `json:"computeType,omitempty"`

	// Description: The description of the Machine Learning compute.
	Description *string `json:"description,omitempty"`

	// DisableLocalAuth: Opt-out of local authentication and ensure customers can use only MSI and AAD exclusively for
	// authentication.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// Properties: Properties of Kubernetes
	Properties *KubernetesProperties `json:"properties,omitempty"`
	ResourceId *string               `json:"resourceId,omitempty"`
}

type SynapseSpark struct {
	// ComputeLocation: Location for the underlying compute
	ComputeLocation *string `json:"computeLocation,omitempty"`

	// ComputeType: The type of compute
	ComputeType SynapseSpark_ComputeType `json:"computeType,omitempty"`

	// Description: The description of the Machine Learning compute.
	Description *string `json:"description,omitempty"`

	// DisableLocalAuth: Opt-out of local authentication and ensure customers can use only MSI and AAD exclusively for
	// authentication.
	DisableLocalAuth *bool                    `json:"disableLocalAuth,omitempty"`
	Properties       *SynapseSpark_Properties `json:"properties,omitempty"`
	ResourceId       *string                  `json:"resourceId,omitempty"`
}

type VirtualMachine struct {
	// ComputeLocation: Location for the underlying compute
	ComputeLocation *string `json:"computeLocation,omitempty"`

	// ComputeType: The type of compute
	ComputeType VirtualMachine_ComputeType `json:"computeType,omitempty"`

	// Description: The description of the Machine Learning compute.
	Description *string `json:"description,omitempty"`

	// DisableLocalAuth: Opt-out of local authentication and ensure customers can use only MSI and AAD exclusively for
	// authentication.
	DisableLocalAuth *bool                      `json:"disableLocalAuth,omitempty"`
	Properties       *VirtualMachine_Properties `json:"properties,omitempty"`
	ResourceId       *string                    `json:"resourceId,omitempty"`
}

// +kubebuilder:validation:Enum={"AKS"}
type AKS_ComputeType string

const AKS_ComputeType_AKS = AKS_ComputeType("AKS")

// Mapping from string to AKS_ComputeType
var aKS_ComputeType_Values = map[string]AKS_ComputeType{
	"aks": AKS_ComputeType_AKS,
}

type AKS_Properties struct {
	// AgentCount: Number of agents
	AgentCount *int `json:"agentCount,omitempty"`

	// AgentVmSize: Agent virtual machine size
	AgentVmSize *string `json:"agentVmSize,omitempty"`

	// AksNetworkingConfiguration: AKS networking configuration for vnet
	AksNetworkingConfiguration *AksNetworkingConfiguration `json:"aksNetworkingConfiguration,omitempty"`

	// ClusterFqdn: Cluster full qualified domain name
	ClusterFqdn *string `json:"clusterFqdn,omitempty"`

	// ClusterPurpose: Intended usage of the cluster
	ClusterPurpose *AKS_Properties_ClusterPurpose `json:"clusterPurpose,omitempty"`

	// LoadBalancerSubnet: Load Balancer Subnet
	LoadBalancerSubnet *string `json:"loadBalancerSubnet,omitempty"`

	// LoadBalancerType: Load Balancer Type
	LoadBalancerType *AKS_Properties_LoadBalancerType `json:"loadBalancerType,omitempty"`

	// SslConfiguration: SSL configuration
	SslConfiguration *SslConfiguration `json:"sslConfiguration,omitempty"`
}

// +kubebuilder:validation:Enum={"AmlCompute"}
type AmlCompute_ComputeType string

const AmlCompute_ComputeType_AmlCompute = AmlCompute_ComputeType("AmlCompute")

// Mapping from string to AmlCompute_ComputeType
var amlCompute_ComputeType_Values = map[string]AmlCompute_ComputeType{
	"amlcompute": AmlCompute_ComputeType_AmlCompute,
}

// AML Compute properties
type AmlComputeProperties struct {
	// EnableNodePublicIp: Enable or disable node public IP address provisioning. Possible values are: Possible values are:
	// true - Indicates that the compute nodes will have public IPs provisioned. false - Indicates that the compute nodes will
	// have a private endpoint and no public IPs.
	EnableNodePublicIp *bool `json:"enableNodePublicIp,omitempty"`

	// IsolatedNetwork: Network is isolated or not
	IsolatedNetwork *bool `json:"isolatedNetwork,omitempty"`

	// OsType: Compute OS Type
	OsType *AmlComputeProperties_OsType `json:"osType,omitempty"`

	// RemoteLoginPortPublicAccess: State of the public SSH port. Possible values are: Disabled - Indicates that the public ssh
	// port is closed on all nodes of the cluster. Enabled - Indicates that the public ssh port is open on all nodes of the
	// cluster. NotSpecified - Indicates that the public ssh port is closed on all nodes of the cluster if VNet is defined,
	// else is open all public nodes. It can be default only during cluster creation time, after creation it will be either
	// enabled or disabled.
	RemoteLoginPortPublicAccess *AmlComputeProperties_RemoteLoginPortPublicAccess `json:"remoteLoginPortPublicAccess,omitempty"`

	// ScaleSettings: Scale settings for AML Compute
	ScaleSettings *ScaleSettings `json:"scaleSettings,omitempty"`

	// Subnet: Virtual network subnet resource ID the compute nodes belong to.
	Subnet *ResourceId `json:"subnet,omitempty"`

	// UserAccountCredentials: Credentials for an administrator user account that will be created on each compute node.
	UserAccountCredentials *UserAccountCredentials `json:"userAccountCredentials,omitempty"`

	// VirtualMachineImage: Virtual Machine image for AML Compute - windows only
	VirtualMachineImage *VirtualMachineImage `json:"virtualMachineImage,omitempty"`

	// VmPriority: Virtual Machine priority
	VmPriority *AmlComputeProperties_VmPriority `json:"vmPriority,omitempty"`

	// VmSize: Virtual Machine Size
	VmSize *string `json:"vmSize,omitempty"`
}

// +kubebuilder:validation:Enum={"ComputeInstance"}
type ComputeInstance_ComputeType string

const ComputeInstance_ComputeType_ComputeInstance = ComputeInstance_ComputeType("ComputeInstance")

// Mapping from string to ComputeInstance_ComputeType
var computeInstance_ComputeType_Values = map[string]ComputeInstance_ComputeType{
	"computeinstance": ComputeInstance_ComputeType_ComputeInstance,
}

// Compute Instance properties
type ComputeInstanceProperties struct {
	// ApplicationSharingPolicy: Policy for sharing applications on this compute instance among users of parent workspace. If
	// Personal, only the creator can access applications on this compute instance. When Shared, any workspace user can access
	// applications on this instance depending on his/her assigned role.
	ApplicationSharingPolicy *ComputeInstanceProperties_ApplicationSharingPolicy `json:"applicationSharingPolicy,omitempty"`

	// ComputeInstanceAuthorizationType: The Compute Instance Authorization type. Available values are personal (default).
	ComputeInstanceAuthorizationType *ComputeInstanceProperties_ComputeInstanceAuthorizationType `json:"computeInstanceAuthorizationType,omitempty"`

	// PersonalComputeInstanceSettings: Settings for a personal compute instance.
	PersonalComputeInstanceSettings *PersonalComputeInstanceSettings `json:"personalComputeInstanceSettings,omitempty"`

	// SetupScripts: Details of customized scripts to execute for setting up the cluster.
	SetupScripts *SetupScripts `json:"setupScripts,omitempty"`

	// SshSettings: Specifies policy and settings for SSH access.
	SshSettings *ComputeInstanceSshSettings `json:"sshSettings,omitempty"`

	// Subnet: Virtual network subnet resource ID the compute nodes belong to.
	Subnet *ResourceId `json:"subnet,omitempty"`

	// VmSize: Virtual Machine Size
	VmSize *string `json:"vmSize,omitempty"`
}

// +kubebuilder:validation:Enum={"Databricks"}
type Databricks_ComputeType string

const Databricks_ComputeType_Databricks = Databricks_ComputeType("Databricks")

// Mapping from string to Databricks_ComputeType
var databricks_ComputeType_Values = map[string]Databricks_ComputeType{
	"databricks": Databricks_ComputeType_Databricks,
}

// Properties of Databricks
type DatabricksProperties struct {
	// DatabricksAccessToken: Databricks access token
	DatabricksAccessToken *string `json:"databricksAccessToken,omitempty"`

	// WorkspaceUrl: Workspace Url
	WorkspaceUrl *string `json:"workspaceUrl,omitempty"`
}

// +kubebuilder:validation:Enum={"DataFactory"}
type DataFactory_ComputeType string

const DataFactory_ComputeType_DataFactory = DataFactory_ComputeType("DataFactory")

// Mapping from string to DataFactory_ComputeType
var dataFactory_ComputeType_Values = map[string]DataFactory_ComputeType{
	"datafactory": DataFactory_ComputeType_DataFactory,
}

// +kubebuilder:validation:Enum={"DataLakeAnalytics"}
type DataLakeAnalytics_ComputeType string

const DataLakeAnalytics_ComputeType_DataLakeAnalytics = DataLakeAnalytics_ComputeType("DataLakeAnalytics")

// Mapping from string to DataLakeAnalytics_ComputeType
var dataLakeAnalytics_ComputeType_Values = map[string]DataLakeAnalytics_ComputeType{
	"datalakeanalytics": DataLakeAnalytics_ComputeType_DataLakeAnalytics,
}

type DataLakeAnalytics_Properties struct {
	// DataLakeStoreAccountName: DataLake Store Account Name
	DataLakeStoreAccountName *string `json:"dataLakeStoreAccountName,omitempty"`
}

// +kubebuilder:validation:Enum={"HDInsight"}
type HDInsight_ComputeType string

const HDInsight_ComputeType_HDInsight = HDInsight_ComputeType("HDInsight")

// Mapping from string to HDInsight_ComputeType
var hDInsight_ComputeType_Values = map[string]HDInsight_ComputeType{
	"hdinsight": HDInsight_ComputeType_HDInsight,
}

// HDInsight compute properties
type HDInsightProperties struct {
	// Address: Public IP address of the master node of the cluster.
	Address *string `json:"address,omitempty"`

	// AdministratorAccount: Admin credentials for master node of the cluster
	AdministratorAccount *VirtualMachineSshCredentials `json:"administratorAccount,omitempty"`

	// SshPort: Port open for ssh connections on the master node of the cluster.
	SshPort *int `json:"sshPort,omitempty"`
}

// +kubebuilder:validation:Enum={"Kubernetes"}
type Kubernetes_ComputeType string

const Kubernetes_ComputeType_Kubernetes = Kubernetes_ComputeType("Kubernetes")

// Mapping from string to Kubernetes_ComputeType
var kubernetes_ComputeType_Values = map[string]Kubernetes_ComputeType{
	"kubernetes": Kubernetes_ComputeType_Kubernetes,
}

// Kubernetes properties
type KubernetesProperties struct {
	// DefaultInstanceType: Default instance type
	DefaultInstanceType *string `json:"defaultInstanceType,omitempty"`

	// ExtensionInstanceReleaseTrain: Extension instance release train.
	ExtensionInstanceReleaseTrain *string `json:"extensionInstanceReleaseTrain,omitempty"`

	// ExtensionPrincipalId: Extension principal-id.
	ExtensionPrincipalId *string `json:"extensionPrincipalId,omitempty"`

	// InstanceTypes: Instance Type Schema
	InstanceTypes map[string]InstanceTypeSchema `json:"instanceTypes,omitempty"`

	// Namespace: Compute namespace
	Namespace *string `json:"namespace,omitempty"`

	// RelayConnectionString: Relay connection string.
	RelayConnectionString *string `json:"relayConnectionString,omitempty"`

	// ServiceBusConnectionString: ServiceBus connection string.
	ServiceBusConnectionString *string `json:"serviceBusConnectionString,omitempty"`

	// VcName: VC name.
	VcName *string `json:"vcName,omitempty"`
}

// +kubebuilder:validation:Enum={"SynapseSpark"}
type SynapseSpark_ComputeType string

const SynapseSpark_ComputeType_SynapseSpark = SynapseSpark_ComputeType("SynapseSpark")

// Mapping from string to SynapseSpark_ComputeType
var synapseSpark_ComputeType_Values = map[string]SynapseSpark_ComputeType{
	"synapsespark": SynapseSpark_ComputeType_SynapseSpark,
}

type SynapseSpark_Properties struct {
	// AutoPauseProperties: Auto pause properties.
	AutoPauseProperties *AutoPauseProperties `json:"autoPauseProperties,omitempty"`

	// AutoScaleProperties: Auto scale properties.
	AutoScaleProperties *AutoScaleProperties `json:"autoScaleProperties,omitempty"`

	// NodeCount: The number of compute nodes currently assigned to the compute.
	NodeCount *int `json:"nodeCount,omitempty"`

	// NodeSize: Node size.
	NodeSize *string `json:"nodeSize,omitempty"`

	// NodeSizeFamily: Node size family.
	NodeSizeFamily *string `json:"nodeSizeFamily,omitempty"`

	// PoolName: Pool name.
	PoolName *string `json:"poolName,omitempty"`

	// ResourceGroup: Name of the resource group in which workspace is located.
	ResourceGroup *string `json:"resourceGroup,omitempty"`

	// SparkVersion: Spark version.
	SparkVersion *string `json:"sparkVersion,omitempty"`

	// SubscriptionId: Azure subscription identifier.
	SubscriptionId *string `json:"subscriptionId,omitempty"`

	// WorkspaceName: Name of Azure Machine Learning workspace.
	WorkspaceName *string `json:"workspaceName,omitempty"`
}

// +kubebuilder:validation:Enum={"VirtualMachine"}
type VirtualMachine_ComputeType string

const VirtualMachine_ComputeType_VirtualMachine = VirtualMachine_ComputeType("VirtualMachine")

// Mapping from string to VirtualMachine_ComputeType
var virtualMachine_ComputeType_Values = map[string]VirtualMachine_ComputeType{
	"virtualmachine": VirtualMachine_ComputeType_VirtualMachine,
}

type VirtualMachine_Properties struct {
	// Address: Public IP address of the virtual machine.
	Address *string `json:"address,omitempty"`

	// AdministratorAccount: Admin credentials for virtual machine
	AdministratorAccount *VirtualMachineSshCredentials `json:"administratorAccount,omitempty"`

	// IsNotebookInstanceCompute: Indicates whether this compute will be used for running notebooks.
	IsNotebookInstanceCompute *bool `json:"isNotebookInstanceCompute,omitempty"`

	// SshPort: Port open for ssh connections.
	SshPort *int `json:"sshPort,omitempty"`

	// VirtualMachineSize: Virtual Machine size
	VirtualMachineSize *string `json:"virtualMachineSize,omitempty"`
}

// +kubebuilder:validation:Enum={"DenseProd","DevTest","FastProd"}
type AKS_Properties_ClusterPurpose string

const (
	AKS_Properties_ClusterPurpose_DenseProd = AKS_Properties_ClusterPurpose("DenseProd")
	AKS_Properties_ClusterPurpose_DevTest   = AKS_Properties_ClusterPurpose("DevTest")
	AKS_Properties_ClusterPurpose_FastProd  = AKS_Properties_ClusterPurpose("FastProd")
)

// Mapping from string to AKS_Properties_ClusterPurpose
var aKS_Properties_ClusterPurpose_Values = map[string]AKS_Properties_ClusterPurpose{
	"denseprod": AKS_Properties_ClusterPurpose_DenseProd,
	"devtest":   AKS_Properties_ClusterPurpose_DevTest,
	"fastprod":  AKS_Properties_ClusterPurpose_FastProd,
}

// +kubebuilder:validation:Enum={"InternalLoadBalancer","PublicIp"}
type AKS_Properties_LoadBalancerType string

const (
	AKS_Properties_LoadBalancerType_InternalLoadBalancer = AKS_Properties_LoadBalancerType("InternalLoadBalancer")
	AKS_Properties_LoadBalancerType_PublicIp             = AKS_Properties_LoadBalancerType("PublicIp")
)

// Mapping from string to AKS_Properties_LoadBalancerType
var aKS_Properties_LoadBalancerType_Values = map[string]AKS_Properties_LoadBalancerType{
	"internalloadbalancer": AKS_Properties_LoadBalancerType_InternalLoadBalancer,
	"publicip":             AKS_Properties_LoadBalancerType_PublicIp,
}

// Advance configuration for AKS networking
type AksNetworkingConfiguration struct {
	// DnsServiceIP: An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address
	// range specified in serviceCidr.
	DnsServiceIP *string `json:"dnsServiceIP,omitempty"`

	// DockerBridgeCidr: A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP
	// ranges or the Kubernetes service address range.
	DockerBridgeCidr *string `json:"dockerBridgeCidr,omitempty"`

	// ServiceCidr: A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP
	// ranges.
	ServiceCidr *string `json:"serviceCidr,omitempty"`
	SubnetId    *string `json:"subnetId,omitempty"`
}

// +kubebuilder:validation:Enum={"Linux","Windows"}
type AmlComputeProperties_OsType string

const (
	AmlComputeProperties_OsType_Linux   = AmlComputeProperties_OsType("Linux")
	AmlComputeProperties_OsType_Windows = AmlComputeProperties_OsType("Windows")
)

// Mapping from string to AmlComputeProperties_OsType
var amlComputeProperties_OsType_Values = map[string]AmlComputeProperties_OsType{
	"linux":   AmlComputeProperties_OsType_Linux,
	"windows": AmlComputeProperties_OsType_Windows,
}

// +kubebuilder:validation:Enum={"Disabled","Enabled","NotSpecified"}
type AmlComputeProperties_RemoteLoginPortPublicAccess string

const (
	AmlComputeProperties_RemoteLoginPortPublicAccess_Disabled     = AmlComputeProperties_RemoteLoginPortPublicAccess("Disabled")
	AmlComputeProperties_RemoteLoginPortPublicAccess_Enabled      = AmlComputeProperties_RemoteLoginPortPublicAccess("Enabled")
	AmlComputeProperties_RemoteLoginPortPublicAccess_NotSpecified = AmlComputeProperties_RemoteLoginPortPublicAccess("NotSpecified")
)

// Mapping from string to AmlComputeProperties_RemoteLoginPortPublicAccess
var amlComputeProperties_RemoteLoginPortPublicAccess_Values = map[string]AmlComputeProperties_RemoteLoginPortPublicAccess{
	"disabled":     AmlComputeProperties_RemoteLoginPortPublicAccess_Disabled,
	"enabled":      AmlComputeProperties_RemoteLoginPortPublicAccess_Enabled,
	"notspecified": AmlComputeProperties_RemoteLoginPortPublicAccess_NotSpecified,
}

// +kubebuilder:validation:Enum={"Dedicated","LowPriority"}
type AmlComputeProperties_VmPriority string

const (
	AmlComputeProperties_VmPriority_Dedicated   = AmlComputeProperties_VmPriority("Dedicated")
	AmlComputeProperties_VmPriority_LowPriority = AmlComputeProperties_VmPriority("LowPriority")
)

// Mapping from string to AmlComputeProperties_VmPriority
var amlComputeProperties_VmPriority_Values = map[string]AmlComputeProperties_VmPriority{
	"dedicated":   AmlComputeProperties_VmPriority_Dedicated,
	"lowpriority": AmlComputeProperties_VmPriority_LowPriority,
}

// Auto pause properties
type AutoPauseProperties struct {
	DelayInMinutes *int  `json:"delayInMinutes,omitempty"`
	Enabled        *bool `json:"enabled,omitempty"`
}

// Auto scale properties
type AutoScaleProperties struct {
	Enabled      *bool `json:"enabled,omitempty"`
	MaxNodeCount *int  `json:"maxNodeCount,omitempty"`
	MinNodeCount *int  `json:"minNodeCount,omitempty"`
}

// +kubebuilder:validation:Enum={"Personal","Shared"}
type ComputeInstanceProperties_ApplicationSharingPolicy string

const (
	ComputeInstanceProperties_ApplicationSharingPolicy_Personal = ComputeInstanceProperties_ApplicationSharingPolicy("Personal")
	ComputeInstanceProperties_ApplicationSharingPolicy_Shared   = ComputeInstanceProperties_ApplicationSharingPolicy("Shared")
)

// Mapping from string to ComputeInstanceProperties_ApplicationSharingPolicy
var computeInstanceProperties_ApplicationSharingPolicy_Values = map[string]ComputeInstanceProperties_ApplicationSharingPolicy{
	"personal": ComputeInstanceProperties_ApplicationSharingPolicy_Personal,
	"shared":   ComputeInstanceProperties_ApplicationSharingPolicy_Shared,
}

// +kubebuilder:validation:Enum={"personal"}
type ComputeInstanceProperties_ComputeInstanceAuthorizationType string

const ComputeInstanceProperties_ComputeInstanceAuthorizationType_Personal = ComputeInstanceProperties_ComputeInstanceAuthorizationType("personal")

// Mapping from string to ComputeInstanceProperties_ComputeInstanceAuthorizationType
var computeInstanceProperties_ComputeInstanceAuthorizationType_Values = map[string]ComputeInstanceProperties_ComputeInstanceAuthorizationType{
	"personal": ComputeInstanceProperties_ComputeInstanceAuthorizationType_Personal,
}

// Specifies policy and settings for SSH access.
type ComputeInstanceSshSettings struct {
	// AdminPublicKey: Specifies the SSH rsa public key file as a string. Use "ssh-keygen -t rsa -b 2048" to generate your SSH
	// key pairs.
	AdminPublicKey *string `json:"adminPublicKey,omitempty"`

	// SshPublicAccess: State of the public SSH port. Possible values are: Disabled - Indicates that the public ssh port is
	// closed on this instance. Enabled - Indicates that the public ssh port is open and accessible according to the
	// VNet/subnet policy if applicable.
	SshPublicAccess *ComputeInstanceSshSettings_SshPublicAccess `json:"sshPublicAccess,omitempty"`
}

// Instance type schema.
type InstanceTypeSchema struct {
	// NodeSelector: Node Selector
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// Resources: Resource requests/limits for this instance type
	Resources *InstanceTypeSchema_Resources `json:"resources,omitempty"`
}

// Settings for a personal compute instance.
type PersonalComputeInstanceSettings struct {
	// AssignedUser: A user explicitly assigned to a personal compute instance.
	AssignedUser *AssignedUser `json:"assignedUser,omitempty"`
}

// Represents a resource ID. For example, for a subnet, it is the resource URL for the subnet.
type ResourceId struct {
	Id *string `json:"id,omitempty"`
}

// scale settings for AML Compute
type ScaleSettings struct {
	// MaxNodeCount: Max number of nodes to use
	MaxNodeCount *int `json:"maxNodeCount,omitempty"`

	// MinNodeCount: Min number of nodes to use
	MinNodeCount *int `json:"minNodeCount,omitempty"`

	// NodeIdleTimeBeforeScaleDown: Node Idle Time before scaling down amlCompute. This string needs to be in the RFC Format.
	NodeIdleTimeBeforeScaleDown *string `json:"nodeIdleTimeBeforeScaleDown,omitempty"`
}

// Details of customized scripts to execute for setting up the cluster.
type SetupScripts struct {
	// Scripts: Customized setup scripts
	Scripts *ScriptsToExecute `json:"scripts,omitempty"`
}

// The ssl configuration for scoring
type SslConfiguration struct {
	// Cert: Cert data
	Cert *string `json:"cert,omitempty"`

	// Cname: CNAME of the cert
	Cname *string `json:"cname,omitempty"`

	// Key: Key data
	Key *string `json:"key,omitempty"`

	// LeafDomainLabel: Leaf domain label of public endpoint
	LeafDomainLabel *string `json:"leafDomainLabel,omitempty"`

	// OverwriteExistingDomain: Indicates whether to overwrite existing domain label.
	OverwriteExistingDomain *bool `json:"overwriteExistingDomain,omitempty"`

	// Status: Enable or disable ssl for scoring
	Status *SslConfiguration_Status `json:"status,omitempty"`
}

// Settings for user account that gets created on each on the nodes of a compute.
type UserAccountCredentials struct {
	// AdminUserName: Name of the administrator user account which can be used to SSH to nodes.
	AdminUserName *string `json:"adminUserName,omitempty"`

	// AdminUserPassword: Password of the administrator user account.
	AdminUserPassword *string `json:"adminUserPassword,omitempty"`

	// AdminUserSshPublicKey: SSH public key of the administrator user account.
	AdminUserSshPublicKey *string `json:"adminUserSshPublicKey,omitempty"`
}

// Virtual Machine image for Windows AML Compute
type VirtualMachineImage struct {
	Id *string `json:"id,omitempty"`
}

// Admin credentials for virtual machine
type VirtualMachineSshCredentials struct {
	// Password: Password of admin account
	Password *string `json:"password,omitempty"`

	// PrivateKeyData: Private key data
	PrivateKeyData *string `json:"privateKeyData,omitempty"`

	// PublicKeyData: Public key data
	PublicKeyData *string `json:"publicKeyData,omitempty"`

	// Username: Username of admin account
	Username *string `json:"username,omitempty"`
}

// A user that can be assigned to a compute instance.
type AssignedUser struct {
	// ObjectId: User’s AAD Object Id.
	ObjectId *string `json:"objectId,omitempty"`

	// TenantId: User’s AAD Tenant Id.
	TenantId *string `json:"tenantId,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ComputeInstanceSshSettings_SshPublicAccess string

const (
	ComputeInstanceSshSettings_SshPublicAccess_Disabled = ComputeInstanceSshSettings_SshPublicAccess("Disabled")
	ComputeInstanceSshSettings_SshPublicAccess_Enabled  = ComputeInstanceSshSettings_SshPublicAccess("Enabled")
)

// Mapping from string to ComputeInstanceSshSettings_SshPublicAccess
var computeInstanceSshSettings_SshPublicAccess_Values = map[string]ComputeInstanceSshSettings_SshPublicAccess{
	"disabled": ComputeInstanceSshSettings_SshPublicAccess_Disabled,
	"enabled":  ComputeInstanceSshSettings_SshPublicAccess_Enabled,
}

type InstanceTypeSchema_Resources struct {
	// Limits: Resource limits for this instance type
	Limits map[string]string `json:"limits,omitempty"`

	// Requests: Resource requests for this instance type
	Requests map[string]string `json:"requests,omitempty"`
}

// Customized setup scripts
type ScriptsToExecute struct {
	// CreationScript: Script that's run only once during provision of the compute.
	CreationScript *ScriptReference `json:"creationScript,omitempty"`

	// StartupScript: Script that's run every time the machine starts.
	StartupScript *ScriptReference `json:"startupScript,omitempty"`
}

// +kubebuilder:validation:Enum={"Auto","Disabled","Enabled"}
type SslConfiguration_Status string

const (
	SslConfiguration_Status_Auto     = SslConfiguration_Status("Auto")
	SslConfiguration_Status_Disabled = SslConfiguration_Status("Disabled")
	SslConfiguration_Status_Enabled  = SslConfiguration_Status("Enabled")
)

// Mapping from string to SslConfiguration_Status
var sslConfiguration_Status_Values = map[string]SslConfiguration_Status{
	"auto":     SslConfiguration_Status_Auto,
	"disabled": SslConfiguration_Status_Disabled,
	"enabled":  SslConfiguration_Status_Enabled,
}

// Script reference
type ScriptReference struct {
	// ScriptArguments: Optional command line arguments passed to the script to run.
	ScriptArguments *string `json:"scriptArguments,omitempty"`

	// ScriptData: The location of scripts in the mounted volume.
	ScriptData *string `json:"scriptData,omitempty"`

	// ScriptSource: The storage source of the script: inline, workspace.
	ScriptSource *string `json:"scriptSource,omitempty"`

	// Timeout: Optional time period passed to timeout command.
	Timeout *string `json:"timeout,omitempty"`
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210701storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=machinelearningservices.azure.com,resources=workspacescomputes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=machinelearningservices.azure.com,resources={workspacescomputes/status,workspacescomputes/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210701.WorkspacesCompute
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/resourceDefinitions/workspaces_computes
type WorkspacesCompute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Workspaces_Compute_Spec `json:"spec,omitempty"`
	Status            ComputeResource_STATUS  `json:"status,omitempty"`
}

var _ conditions.Conditioner = &WorkspacesCompute{}

// GetConditions returns the conditions of the resource
func (compute *WorkspacesCompute) GetConditions() conditions.Conditions {
	return compute.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (compute *WorkspacesCompute) SetConditions(conditions conditions.Conditions) {
	compute.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &WorkspacesCompute{}

// AzureName returns the Azure name of the resource
func (compute *WorkspacesCompute) AzureName() string {
	return compute.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-07-01"
func (compute WorkspacesCompute) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (compute *WorkspacesCompute) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (compute *WorkspacesCompute) GetSpec() genruntime.ConvertibleSpec {
	return &compute.Spec
}

// GetStatus returns the status of this resource
func (compute *WorkspacesCompute) GetStatus() genruntime.ConvertibleStatus {
	return &compute.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.MachineLearningServices/workspaces/computes"
func (compute *WorkspacesCompute) GetType() string {
	return "Microsoft.MachineLearningServices/workspaces/computes"
}

// NewEmptyStatus returns a new empty (blank) status
func (compute *WorkspacesCompute) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ComputeResource_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (compute *WorkspacesCompute) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(compute.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  compute.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (compute *WorkspacesCompute) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ComputeResource_STATUS); ok {
		compute.Status = *st
		return nil
	}

	// Convert status to required version
	var st ComputeResource_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	compute.Status = st
	return nil
}

// Hub marks that this WorkspacesCompute is the hub type for conversion
func (compute *WorkspacesCompute) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (compute *WorkspacesCompute) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: compute.Spec.OriginalVersion,
		Kind:    "WorkspacesCompute",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210701.WorkspacesCompute
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/resourceDefinitions/workspaces_computes
type WorkspacesComputeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspacesCompute `json:"items"`
}

// Storage version of v1beta20210701.ComputeResource_STATUS
type ComputeResource_STATUS struct {
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Identity    *Identity_STATUS       `json:"identity,omitempty"`
	Location    *string                `json:"location,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Properties  *Compute_STATUS        `json:"properties,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Sku         *Sku_STATUS            `json:"sku,omitempty"`
	SystemData  *SystemData_STATUS     `json:"systemData,omitempty"`
	Tags        map[string]string      `json:"tags,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ComputeResource_STATUS{}

// ConvertStatusFrom populates our ComputeResource_STATUS from the provided source
func (resource *ComputeResource_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == resource {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(resource)
}

// ConvertStatusTo populates the provided destination from our ComputeResource_STATUS
func (resource *ComputeResource_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == resource {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(resource)
}

// Storage version of v1beta20210701.Workspaces_Compute_Spec
type Workspaces_Compute_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string    `json:"azureName,omitempty"`
	Identity        *Identity `json:"identity,omitempty"`
	Location        *string   `json:"location,omitempty"`
	OriginalVersion string    `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a machinelearningservices.azure.com/Workspace resource
	Owner       *genruntime.KnownResourceReference `group:"machinelearningservices.azure.com" json:"owner,omitempty" kind:"Workspace"`
	Properties  *Compute                           `json:"properties,omitempty"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Sku         *Sku                               `json:"sku,omitempty"`
	SystemData  *SystemData                        `json:"systemData,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Workspaces_Compute_Spec{}

// ConvertSpecFrom populates our Workspaces_Compute_Spec from the provided source
func (compute *Workspaces_Compute_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == compute {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(compute)
}

// ConvertSpecTo populates the provided destination from our Workspaces_Compute_Spec
func (compute *Workspaces_Compute_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == compute {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(compute)
}

// Storage version of v1beta20210701.Compute
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/Compute
type Compute struct {
	AKS               *Compute_AKS               `json:"aks,omitempty"`
	AmlCompute        *Compute_AmlCompute        `json:"amlCompute,omitempty"`
	ComputeInstance   *Compute_ComputeInstance   `json:"computeInstance,omitempty"`
	DataFactory       *Compute_DataFactory       `json:"dataFactory,omitempty"`
	DataLakeAnalytics *Compute_DataLakeAnalytics `json:"dataLakeAnalytics,omitempty"`
	Databricks        *Compute_Databricks        `json:"databricks,omitempty"`
	HDInsight         *Compute_HDInsight         `json:"hdInsight,omitempty"`
	PropertyBag       genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	SynapseSpark      *Compute_SynapseSpark      `json:"synapseSpark,omitempty"`
	VirtualMachine    *Compute_VirtualMachine    `json:"virtualMachine,omitempty"`
}

// Storage version of v1beta20210701.Compute_STATUS
type Compute_STATUS struct {
	ComputeLocation    *string                `json:"computeLocation,omitempty"`
	ComputeType        *string                `json:"computeType,omitempty"`
	CreatedOn          *string                `json:"createdOn,omitempty"`
	Description        *string                `json:"description,omitempty"`
	DisableLocalAuth   *bool                  `json:"disableLocalAuth,omitempty"`
	IsAttachedCompute  *bool                  `json:"isAttachedCompute,omitempty"`
	ModifiedOn         *string                `json:"modifiedOn,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningErrors []ErrorResponse_STATUS `json:"provisioningErrors,omitempty"`
	ProvisioningState  *string                `json:"provisioningState,omitempty"`
	ResourceId         *string                `json:"resourceId,omitempty"`
}

// Storage version of v1beta20210701.Compute_AKS
type Compute_AKS struct {
	ComputeLocation  *string                `json:"computeLocation,omitempty"`
	ComputeType      *string                `json:"computeType,omitempty"`
	Description      *string                `json:"description,omitempty"`
	DisableLocalAuth *bool                  `json:"disableLocalAuth,omitempty"`
	Properties       *AKSProperties         `json:"properties,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// ResourceReference: ARM resource id of the underlying compute
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

// Storage version of v1beta20210701.Compute_AmlCompute
type Compute_AmlCompute struct {
	ComputeLocation  *string                `json:"computeLocation,omitempty"`
	ComputeType      *string                `json:"computeType,omitempty"`
	Description      *string                `json:"description,omitempty"`
	DisableLocalAuth *bool                  `json:"disableLocalAuth,omitempty"`
	Properties       *AmlComputeProperties  `json:"properties,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// ResourceReference: ARM resource id of the underlying compute
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

// Storage version of v1beta20210701.Compute_ComputeInstance
type Compute_ComputeInstance struct {
	ComputeLocation  *string                    `json:"computeLocation,omitempty"`
	ComputeType      *string                    `json:"computeType,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	DisableLocalAuth *bool                      `json:"disableLocalAuth,omitempty"`
	Properties       *ComputeInstanceProperties `json:"properties,omitempty"`
	PropertyBag      genruntime.PropertyBag     `json:"$propertyBag,omitempty"`

	// ResourceReference: ARM resource id of the underlying compute
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

// Storage version of v1beta20210701.Compute_Databricks
type Compute_Databricks struct {
	ComputeLocation  *string                `json:"computeLocation,omitempty"`
	ComputeType      *string                `json:"computeType,omitempty"`
	Description      *string                `json:"description,omitempty"`
	DisableLocalAuth *bool                  `json:"disableLocalAuth,omitempty"`
	Properties       *DatabricksProperties  `json:"properties,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// ResourceReference: ARM resource id of the underlying compute
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

// Storage version of v1beta20210701.Compute_DataFactory
type Compute_DataFactory struct {
	ComputeLocation  *string                `json:"computeLocation,omitempty"`
	ComputeType      *string                `json:"computeType,omitempty"`
	Description      *string                `json:"description,omitempty"`
	DisableLocalAuth *bool                  `json:"disableLocalAuth,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// ResourceReference: ARM resource id of the underlying compute
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

// Storage version of v1beta20210701.Compute_DataLakeAnalytics
type Compute_DataLakeAnalytics struct {
	ComputeLocation  *string                      `json:"computeLocation,omitempty"`
	ComputeType      *string                      `json:"computeType,omitempty"`
	Description      *string                      `json:"description,omitempty"`
	DisableLocalAuth *bool                        `json:"disableLocalAuth,omitempty"`
	Properties       *DataLakeAnalyticsProperties `json:"properties,omitempty"`
	PropertyBag      genruntime.PropertyBag       `json:"$propertyBag,omitempty"`

	// ResourceReference: ARM resource id of the underlying compute
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

// Storage version of v1beta20210701.Compute_HDInsight
type Compute_HDInsight struct {
	ComputeLocation  *string                `json:"computeLocation,omitempty"`
	ComputeType      *string                `json:"computeType,omitempty"`
	Description      *string                `json:"description,omitempty"`
	DisableLocalAuth *bool                  `json:"disableLocalAuth,omitempty"`
	Properties       *HDInsightProperties   `json:"properties,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// ResourceReference: ARM resource id of the underlying compute
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

// Storage version of v1beta20210701.Compute_SynapseSpark
type Compute_SynapseSpark struct {
	ComputeLocation  *string                 `json:"computeLocation,omitempty"`
	ComputeType      *string                 `json:"computeType,omitempty"`
	Description      *string                 `json:"description,omitempty"`
	DisableLocalAuth *bool                   `json:"disableLocalAuth,omitempty"`
	Properties       *SynapseSparkProperties `json:"properties,omitempty"`
	PropertyBag      genruntime.PropertyBag  `json:"$propertyBag,omitempty"`

	// ResourceReference: ARM resource id of the underlying compute
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

// Storage version of v1beta20210701.Compute_VirtualMachine
type Compute_VirtualMachine struct {
	ComputeLocation  *string                   `json:"computeLocation,omitempty"`
	ComputeType      *string                   `json:"computeType,omitempty"`
	Description      *string                   `json:"description,omitempty"`
	DisableLocalAuth *bool                     `json:"disableLocalAuth,omitempty"`
	Properties       *VirtualMachineProperties `json:"properties,omitempty"`
	PropertyBag      genruntime.PropertyBag    `json:"$propertyBag,omitempty"`

	// ResourceReference: ARM resource id of the underlying compute
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

// Storage version of v1beta20210701.ErrorResponse_STATUS
type ErrorResponse_STATUS struct {
	Error       *ErrorDetail_STATUS    `json:"error,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210701.AKSProperties
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/AKSProperties
type AKSProperties struct {
	AgentCount                 *int                        `json:"agentCount,omitempty"`
	AgentVmSize                *string                     `json:"agentVmSize,omitempty"`
	AksNetworkingConfiguration *AksNetworkingConfiguration `json:"aksNetworkingConfiguration,omitempty"`
	ClusterFqdn                *string                     `json:"clusterFqdn,omitempty"`
	ClusterPurpose             *string                     `json:"clusterPurpose,omitempty"`
	LoadBalancerSubnet         *string                     `json:"loadBalancerSubnet,omitempty"`
	LoadBalancerType           *string                     `json:"loadBalancerType,omitempty"`
	PropertyBag                genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	SslConfiguration           *SslConfiguration           `json:"sslConfiguration,omitempty"`
}

// Storage version of v1beta20210701.AmlComputeProperties
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/AmlComputeProperties
type AmlComputeProperties struct {
	EnableNodePublicIp          *bool                   `json:"enableNodePublicIp,omitempty"`
	IsolatedNetwork             *bool                   `json:"isolatedNetwork,omitempty"`
	OsType                      *string                 `json:"osType,omitempty"`
	PropertyBag                 genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
	RemoteLoginPortPublicAccess *string                 `json:"remoteLoginPortPublicAccess,omitempty"`
	ScaleSettings               *ScaleSettings          `json:"scaleSettings,omitempty"`
	Subnet                      *ResourceId             `json:"subnet,omitempty"`
	UserAccountCredentials      *UserAccountCredentials `json:"userAccountCredentials,omitempty"`
	VirtualMachineImage         *VirtualMachineImage    `json:"virtualMachineImage,omitempty"`
	VmPriority                  *string                 `json:"vmPriority,omitempty"`
	VmSize                      *string                 `json:"vmSize,omitempty"`
}

// Storage version of v1beta20210701.ComputeInstanceProperties
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/ComputeInstanceProperties
type ComputeInstanceProperties struct {
	ApplicationSharingPolicy         *string                          `json:"applicationSharingPolicy,omitempty"`
	ComputeInstanceAuthorizationType *string                          `json:"computeInstanceAuthorizationType,omitempty"`
	PersonalComputeInstanceSettings  *PersonalComputeInstanceSettings `json:"personalComputeInstanceSettings,omitempty"`
	PropertyBag                      genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
	SetupScripts                     *SetupScripts                    `json:"setupScripts,omitempty"`
	SshSettings                      *ComputeInstanceSshSettings      `json:"sshSettings,omitempty"`
	Subnet                           *ResourceId                      `json:"subnet,omitempty"`
	VmSize                           *string                          `json:"vmSize,omitempty"`
}

// Storage version of v1beta20210701.DatabricksProperties
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/DatabricksProperties
type DatabricksProperties struct {
	DatabricksAccessToken *string                `json:"databricksAccessToken,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	WorkspaceUrl          *string                `json:"workspaceUrl,omitempty"`
}

// Storage version of v1beta20210701.DataLakeAnalyticsProperties
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/DataLakeAnalyticsProperties
type DataLakeAnalyticsProperties struct {
	DataLakeStoreAccountName *string                `json:"dataLakeStoreAccountName,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210701.ErrorDetail_STATUS
type ErrorDetail_STATUS struct {
	AdditionalInfo []ErrorAdditionalInfo_STATUS  `json:"additionalInfo,omitempty"`
	Code           *string                       `json:"code,omitempty"`
	Details        []ErrorDetail_STATUS_Unrolled `json:"details,omitempty"`
	Message        *string                       `json:"message,omitempty"`
	PropertyBag    genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Target         *string                       `json:"target,omitempty"`
}

// Storage version of v1beta20210701.HDInsightProperties
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/HDInsightProperties
type HDInsightProperties struct {
	Address              *string                       `json:"address,omitempty"`
	AdministratorAccount *VirtualMachineSshCredentials `json:"administratorAccount,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SshPort              *int                          `json:"sshPort,omitempty"`
}

// Storage version of v1beta20210701.SynapseSparkProperties
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/SynapseSparkProperties
type SynapseSparkProperties struct {
	AutoPauseProperties *AutoPauseProperties   `json:"autoPauseProperties,omitempty"`
	AutoScaleProperties *AutoScaleProperties   `json:"autoScaleProperties,omitempty"`
	NodeCount           *int                   `json:"nodeCount,omitempty"`
	NodeSize            *string                `json:"nodeSize,omitempty"`
	NodeSizeFamily      *string                `json:"nodeSizeFamily,omitempty"`
	PoolName            *string                `json:"poolName,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceGroup       *string                `json:"resourceGroup,omitempty"`
	SparkVersion        *string                `json:"sparkVersion,omitempty"`
	SubscriptionId      *string                `json:"subscriptionId,omitempty"`
	WorkspaceName       *string                `json:"workspaceName,omitempty"`
}

// Storage version of v1beta20210701.VirtualMachineProperties
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/VirtualMachineProperties
type VirtualMachineProperties struct {
	Address                   *string                       `json:"address,omitempty"`
	AdministratorAccount      *VirtualMachineSshCredentials `json:"administratorAccount,omitempty"`
	IsNotebookInstanceCompute *bool                         `json:"isNotebookInstanceCompute,omitempty"`
	PropertyBag               genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SshPort                   *int                          `json:"sshPort,omitempty"`
	VirtualMachineSize        *string                       `json:"virtualMachineSize,omitempty"`
}

// Storage version of v1beta20210701.AksNetworkingConfiguration
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/AksNetworkingConfiguration
type AksNetworkingConfiguration struct {
	DnsServiceIP     *string                `json:"dnsServiceIP,omitempty"`
	DockerBridgeCidr *string                `json:"dockerBridgeCidr,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ServiceCidr      *string                `json:"serviceCidr,omitempty"`

	// SubnetReference: Virtual network subnet resource ID the compute nodes belong to
	SubnetReference *genruntime.ResourceReference `armReference:"SubnetId" json:"subnetReference,omitempty"`
}

// Storage version of v1beta20210701.AutoPauseProperties
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/AutoPauseProperties
type AutoPauseProperties struct {
	DelayInMinutes *int                   `json:"delayInMinutes,omitempty"`
	Enabled        *bool                  `json:"enabled,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210701.AutoScaleProperties
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/AutoScaleProperties
type AutoScaleProperties struct {
	Enabled      *bool                  `json:"enabled,omitempty"`
	MaxNodeCount *int                   `json:"maxNodeCount,omitempty"`
	MinNodeCount *int                   `json:"minNodeCount,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210701.ComputeInstanceSshSettings
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/ComputeInstanceSshSettings
type ComputeInstanceSshSettings struct {
	AdminPublicKey  *string                `json:"adminPublicKey,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SshPublicAccess *string                `json:"sshPublicAccess,omitempty"`
}

// Storage version of v1beta20210701.ErrorAdditionalInfo_STATUS
type ErrorAdditionalInfo_STATUS struct {
	Info        map[string]v1.JSON     `json:"info,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1beta20210701.ErrorDetail_STATUS_Unrolled
type ErrorDetail_STATUS_Unrolled struct {
	AdditionalInfo []ErrorAdditionalInfo_STATUS `json:"additionalInfo,omitempty"`
	Code           *string                      `json:"code,omitempty"`
	Message        *string                      `json:"message,omitempty"`
	PropertyBag    genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Target         *string                      `json:"target,omitempty"`
}

// Storage version of v1beta20210701.PersonalComputeInstanceSettings
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/PersonalComputeInstanceSettings
type PersonalComputeInstanceSettings struct {
	AssignedUser *AssignedUser          `json:"assignedUser,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210701.ResourceId
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/ResourceId
type ResourceId struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	// Reference: The ID of the resource
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1beta20210701.ScaleSettings
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/ScaleSettings
type ScaleSettings struct {
	MaxNodeCount                *int                   `json:"maxNodeCount,omitempty"`
	MinNodeCount                *int                   `json:"minNodeCount,omitempty"`
	NodeIdleTimeBeforeScaleDown *string                `json:"nodeIdleTimeBeforeScaleDown,omitempty"`
	PropertyBag                 genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210701.SetupScripts
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/SetupScripts
type SetupScripts struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Scripts     *ScriptsToExecute      `json:"scripts,omitempty"`
}

// Storage version of v1beta20210701.SslConfiguration
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/SslConfiguration
type SslConfiguration struct {
	Cert                    *string                `json:"cert,omitempty"`
	Cname                   *string                `json:"cname,omitempty"`
	Key                     *string                `json:"key,omitempty"`
	LeafDomainLabel         *string                `json:"leafDomainLabel,omitempty"`
	OverwriteExistingDomain *bool                  `json:"overwriteExistingDomain,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Status                  *string                `json:"status,omitempty"`
}

// Storage version of v1beta20210701.UserAccountCredentials
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/UserAccountCredentials
type UserAccountCredentials struct {
	AdminUserName         *string                     `json:"adminUserName,omitempty"`
	AdminUserPassword     *genruntime.SecretReference `json:"adminUserPassword,omitempty"`
	AdminUserSshPublicKey *genruntime.SecretReference `json:"adminUserSshPublicKey,omitempty"`
	PropertyBag           genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210701.VirtualMachineImage
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/VirtualMachineImage
type VirtualMachineImage struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	// Reference: Virtual Machine image path
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1beta20210701.VirtualMachineSshCredentials
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/VirtualMachineSshCredentials
type VirtualMachineSshCredentials struct {
	Password       *genruntime.SecretReference `json:"password,omitempty"`
	PrivateKeyData *string                     `json:"privateKeyData,omitempty"`
	PropertyBag    genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	PublicKeyData  *string                     `json:"publicKeyData,omitempty"`
	Username       *string                     `json:"username,omitempty"`
}

// Storage version of v1beta20210701.AssignedUser
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/AssignedUser
type AssignedUser struct {
	ObjectId    *string                `json:"objectId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TenantId    *string                `json:"tenantId,omitempty"`
}

// Storage version of v1beta20210701.ScriptsToExecute
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/ScriptsToExecute
type ScriptsToExecute struct {
	CreationScript *ScriptReference       `json:"creationScript,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartupScript  *ScriptReference       `json:"startupScript,omitempty"`
}

// Storage version of v1beta20210701.ScriptReference
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/definitions/ScriptReference
type ScriptReference struct {
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ScriptArguments *string                `json:"scriptArguments,omitempty"`
	ScriptData      *string                `json:"scriptData,omitempty"`
	ScriptSource    *string                `json:"scriptSource,omitempty"`
	Timeout         *string                `json:"timeout,omitempty"`
}

func init() {
	SchemeBuilder.Register(&WorkspacesCompute{}, &WorkspacesComputeList{})
}

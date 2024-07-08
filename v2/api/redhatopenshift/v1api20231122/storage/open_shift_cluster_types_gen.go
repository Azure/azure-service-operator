// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=redhatopenshift.azure.com,resources=openshiftclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=redhatopenshift.azure.com,resources={openshiftclusters/status,openshiftclusters/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20231122.OpenShiftCluster
// Generator information:
// - Generated from: /redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/redhatopenshift.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{resourceName}
type OpenShiftCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpenShiftCluster_Spec   `json:"spec,omitempty"`
	Status            OpenShiftCluster_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &OpenShiftCluster{}

// GetConditions returns the conditions of the resource
func (cluster *OpenShiftCluster) GetConditions() conditions.Conditions {
	return cluster.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (cluster *OpenShiftCluster) SetConditions(conditions conditions.Conditions) {
	cluster.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &OpenShiftCluster{}

// AzureName returns the Azure name of the resource
func (cluster *OpenShiftCluster) AzureName() string {
	return cluster.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-11-22"
func (cluster OpenShiftCluster) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (cluster *OpenShiftCluster) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (cluster *OpenShiftCluster) GetSpec() genruntime.ConvertibleSpec {
	return &cluster.Spec
}

// GetStatus returns the status of this resource
func (cluster *OpenShiftCluster) GetStatus() genruntime.ConvertibleStatus {
	return &cluster.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (cluster *OpenShiftCluster) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.RedHatOpenShift/openShiftClusters"
func (cluster *OpenShiftCluster) GetType() string {
	return "Microsoft.RedHatOpenShift/openShiftClusters"
}

// NewEmptyStatus returns a new empty (blank) status
func (cluster *OpenShiftCluster) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &OpenShiftCluster_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (cluster *OpenShiftCluster) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(cluster.Spec)
	return cluster.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (cluster *OpenShiftCluster) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*OpenShiftCluster_STATUS); ok {
		cluster.Status = *st
		return nil
	}

	// Convert status to required version
	var st OpenShiftCluster_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	cluster.Status = st
	return nil
}

// Hub marks that this OpenShiftCluster is the hub type for conversion
func (cluster *OpenShiftCluster) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (cluster *OpenShiftCluster) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: cluster.Spec.OriginalVersion,
		Kind:    "OpenShiftCluster",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20231122.OpenShiftCluster
// Generator information:
// - Generated from: /redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/redhatopenshift.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{resourceName}
type OpenShiftClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenShiftCluster `json:"items"`
}

// Storage version of v1api20231122.APIVersion
// +kubebuilder:validation:Enum={"2023-11-22"}
type APIVersion string

const APIVersion_Value = APIVersion("2023-11-22")

// Storage version of v1api20231122.OpenShiftCluster_Spec
type OpenShiftCluster_Spec struct {
	ApiserverProfile *APIServerProfile `json:"apiserverProfile,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string           `json:"azureName,omitempty"`
	ClusterProfile  *ClusterProfile  `json:"clusterProfile,omitempty"`
	IngressProfiles []IngressProfile `json:"ingressProfiles,omitempty"`
	Location        *string          `json:"location,omitempty"`
	MasterProfile   *MasterProfile   `json:"masterProfile,omitempty"`
	NetworkProfile  *NetworkProfile  `json:"networkProfile,omitempty"`
	OriginalVersion string           `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner                   *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag             genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ProvisioningState       *string                            `json:"provisioningState,omitempty"`
	ServicePrincipalProfile *ServicePrincipalProfile           `json:"servicePrincipalProfile,omitempty"`
	Tags                    map[string]string                  `json:"tags,omitempty"`
	WorkerProfiles          []WorkerProfile                    `json:"workerProfiles,omitempty"`
}

var _ genruntime.ConvertibleSpec = &OpenShiftCluster_Spec{}

// ConvertSpecFrom populates our OpenShiftCluster_Spec from the provided source
func (cluster *OpenShiftCluster_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == cluster {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(cluster)
}

// ConvertSpecTo populates the provided destination from our OpenShiftCluster_Spec
func (cluster *OpenShiftCluster_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == cluster {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(cluster)
}

// Storage version of v1api20231122.OpenShiftCluster_STATUS
// OpenShiftCluster represents an Azure Red Hat OpenShift cluster.
type OpenShiftCluster_STATUS struct {
	ApiserverProfile        *APIServerProfile_STATUS        `json:"apiserverProfile,omitempty"`
	ClusterProfile          *ClusterProfile_STATUS          `json:"clusterProfile,omitempty"`
	Conditions              []conditions.Condition          `json:"conditions,omitempty"`
	ConsoleProfile          *ConsoleProfile_STATUS          `json:"consoleProfile,omitempty"`
	Id                      *string                         `json:"id,omitempty"`
	IngressProfiles         []IngressProfile_STATUS         `json:"ingressProfiles,omitempty"`
	Location                *string                         `json:"location,omitempty"`
	MasterProfile           *MasterProfile_STATUS           `json:"masterProfile,omitempty"`
	Name                    *string                         `json:"name,omitempty"`
	NetworkProfile          *NetworkProfile_STATUS          `json:"networkProfile,omitempty"`
	PropertyBag             genruntime.PropertyBag          `json:"$propertyBag,omitempty"`
	ProvisioningState       *string                         `json:"provisioningState,omitempty"`
	ServicePrincipalProfile *ServicePrincipalProfile_STATUS `json:"servicePrincipalProfile,omitempty"`
	SystemData              *SystemData_STATUS              `json:"systemData,omitempty"`
	Tags                    map[string]string               `json:"tags,omitempty"`
	Type                    *string                         `json:"type,omitempty"`
	WorkerProfiles          []WorkerProfile_STATUS          `json:"workerProfiles,omitempty"`
	WorkerProfilesStatus    []WorkerProfile_STATUS          `json:"workerProfilesStatus,omitempty"`
}

var _ genruntime.ConvertibleStatus = &OpenShiftCluster_STATUS{}

// ConvertStatusFrom populates our OpenShiftCluster_STATUS from the provided source
func (cluster *OpenShiftCluster_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == cluster {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(cluster)
}

// ConvertStatusTo populates the provided destination from our OpenShiftCluster_STATUS
func (cluster *OpenShiftCluster_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == cluster {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(cluster)
}

// Storage version of v1api20231122.APIServerProfile
// APIServerProfile represents an API server profile.
type APIServerProfile struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Visibility  *string                `json:"visibility,omitempty"`
}

// Storage version of v1api20231122.APIServerProfile_STATUS
// APIServerProfile represents an API server profile.
type APIServerProfile_STATUS struct {
	Ip          *string                `json:"ip,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Url         *string                `json:"url,omitempty"`
	Visibility  *string                `json:"visibility,omitempty"`
}

// Storage version of v1api20231122.ClusterProfile
// ClusterProfile represents a cluster profile.
type ClusterProfile struct {
	Domain               *string                     `json:"domain,omitempty"`
	FipsValidatedModules *string                     `json:"fipsValidatedModules,omitempty"`
	PropertyBag          genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	PullSecret           *genruntime.SecretReference `json:"pullSecret,omitempty"`
	ResourceGroupId      *string                     `json:"resourceGroupId,omitempty"`
	Version              *string                     `json:"version,omitempty"`
}

// Storage version of v1api20231122.ClusterProfile_STATUS
// ClusterProfile represents a cluster profile.
type ClusterProfile_STATUS struct {
	Domain               *string                `json:"domain,omitempty"`
	FipsValidatedModules *string                `json:"fipsValidatedModules,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceGroupId      *string                `json:"resourceGroupId,omitempty"`
	Version              *string                `json:"version,omitempty"`
}

// Storage version of v1api20231122.ConsoleProfile_STATUS
// ConsoleProfile represents a console profile.
type ConsoleProfile_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Url         *string                `json:"url,omitempty"`
}

// Storage version of v1api20231122.IngressProfile
// IngressProfile represents an ingress profile.
type IngressProfile struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Visibility  *string                `json:"visibility,omitempty"`
}

// Storage version of v1api20231122.IngressProfile_STATUS
// IngressProfile represents an ingress profile.
type IngressProfile_STATUS struct {
	Ip          *string                `json:"ip,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Visibility  *string                `json:"visibility,omitempty"`
}

// Storage version of v1api20231122.MasterProfile
// MasterProfile represents a master profile.
type MasterProfile struct {
	// DiskEncryptionSetReference: The resource ID of an associated DiskEncryptionSet, if applicable.
	DiskEncryptionSetReference *genruntime.ResourceReference `armReference:"DiskEncryptionSetId" json:"diskEncryptionSetReference,omitempty"`
	EncryptionAtHost           *string                       `json:"encryptionAtHost,omitempty"`
	PropertyBag                genruntime.PropertyBag        `json:"$propertyBag,omitempty"`

	// SubnetReference: The Azure resource ID of the master subnet.
	SubnetReference *genruntime.ResourceReference `armReference:"SubnetId" json:"subnetReference,omitempty"`
	VmSize          *string                       `json:"vmSize,omitempty"`
}

// Storage version of v1api20231122.MasterProfile_STATUS
// MasterProfile represents a master profile.
type MasterProfile_STATUS struct {
	DiskEncryptionSetId *string                `json:"diskEncryptionSetId,omitempty"`
	EncryptionAtHost    *string                `json:"encryptionAtHost,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SubnetId            *string                `json:"subnetId,omitempty"`
	VmSize              *string                `json:"vmSize,omitempty"`
}

// Storage version of v1api20231122.NetworkProfile
// NetworkProfile represents a network profile.
type NetworkProfile struct {
	LoadBalancerProfile *LoadBalancerProfile   `json:"loadBalancerProfile,omitempty"`
	OutboundType        *string                `json:"outboundType,omitempty"`
	PodCidr             *string                `json:"podCidr,omitempty"`
	PreconfiguredNSG    *string                `json:"preconfiguredNSG,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ServiceCidr         *string                `json:"serviceCidr,omitempty"`
}

// Storage version of v1api20231122.NetworkProfile_STATUS
// NetworkProfile represents a network profile.
type NetworkProfile_STATUS struct {
	LoadBalancerProfile *LoadBalancerProfile_STATUS `json:"loadBalancerProfile,omitempty"`
	OutboundType        *string                     `json:"outboundType,omitempty"`
	PodCidr             *string                     `json:"podCidr,omitempty"`
	PreconfiguredNSG    *string                     `json:"preconfiguredNSG,omitempty"`
	PropertyBag         genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	ServiceCidr         *string                     `json:"serviceCidr,omitempty"`
}

// Storage version of v1api20231122.ServicePrincipalProfile
// ServicePrincipalProfile represents a service principal profile.
type ServicePrincipalProfile struct {
	ClientId     *string                     `json:"clientId,omitempty"`
	ClientSecret *genruntime.SecretReference `json:"clientSecret,omitempty"`
	PropertyBag  genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231122.ServicePrincipalProfile_STATUS
// ServicePrincipalProfile represents a service principal profile.
type ServicePrincipalProfile_STATUS struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231122.SystemData_STATUS
// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231122.WorkerProfile
// WorkerProfile represents a worker profile.
type WorkerProfile struct {
	Count *int `json:"count,omitempty"`

	// DiskEncryptionSetReference: The resource ID of an associated DiskEncryptionSet, if applicable.
	DiskEncryptionSetReference *genruntime.ResourceReference `armReference:"DiskEncryptionSetId" json:"diskEncryptionSetReference,omitempty"`
	DiskSizeGB                 *int                          `json:"diskSizeGB,omitempty"`
	EncryptionAtHost           *string                       `json:"encryptionAtHost,omitempty"`
	Name                       *string                       `json:"name,omitempty"`
	PropertyBag                genruntime.PropertyBag        `json:"$propertyBag,omitempty"`

	// SubnetReference: The Azure resource ID of the worker subnet.
	SubnetReference *genruntime.ResourceReference `armReference:"SubnetId" json:"subnetReference,omitempty"`
	VmSize          *string                       `json:"vmSize,omitempty"`
}

// Storage version of v1api20231122.WorkerProfile_STATUS
// WorkerProfile represents a worker profile.
type WorkerProfile_STATUS struct {
	Count               *int                   `json:"count,omitempty"`
	DiskEncryptionSetId *string                `json:"diskEncryptionSetId,omitempty"`
	DiskSizeGB          *int                   `json:"diskSizeGB,omitempty"`
	EncryptionAtHost    *string                `json:"encryptionAtHost,omitempty"`
	Name                *string                `json:"name,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SubnetId            *string                `json:"subnetId,omitempty"`
	VmSize              *string                `json:"vmSize,omitempty"`
}

// Storage version of v1api20231122.LoadBalancerProfile
// LoadBalancerProfile represents the profile of the cluster public load balancer.
type LoadBalancerProfile struct {
	ManagedOutboundIps *ManagedOutboundIPs    `json:"managedOutboundIps,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231122.LoadBalancerProfile_STATUS
// LoadBalancerProfile represents the profile of the cluster public load balancer.
type LoadBalancerProfile_STATUS struct {
	EffectiveOutboundIps []EffectiveOutboundIP_STATUS `json:"effectiveOutboundIps,omitempty"`
	ManagedOutboundIps   *ManagedOutboundIPs_STATUS   `json:"managedOutboundIps,omitempty"`
	PropertyBag          genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231122.EffectiveOutboundIP_STATUS
// EffectiveOutboundIP represents an effective outbound IP resource of the cluster public load balancer.
type EffectiveOutboundIP_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231122.ManagedOutboundIPs
// ManagedOutboundIPs represents the desired managed outbound IPs for the cluster public load balancer.
type ManagedOutboundIPs struct {
	Count       *int                   `json:"count,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231122.ManagedOutboundIPs_STATUS
// ManagedOutboundIPs represents the desired managed outbound IPs for the cluster public load balancer.
type ManagedOutboundIPs_STATUS struct {
	Count       *int                   `json:"count,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&OpenShiftCluster{}, &OpenShiftClusterList{})
}

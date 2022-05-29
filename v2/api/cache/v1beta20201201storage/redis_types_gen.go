// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201storage

import (
	v20210301s "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20210301storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=cache.azure.com,resources=redis,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.azure.com,resources={redis/status,redis/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20201201.Redis
// Generator information:
// - Generated from: /redis/resource-manager/Microsoft.Cache/stable/2020-12-01/redis.json
type Redis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Redis_Spec   `json:"spec,omitempty"`
	Status            Redis_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Redis{}

// GetConditions returns the conditions of the resource
func (redis *Redis) GetConditions() conditions.Conditions {
	return redis.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (redis *Redis) SetConditions(conditions conditions.Conditions) {
	redis.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Redis{}

// AzureName returns the Azure name of the resource
func (redis *Redis) AzureName() string {
	return redis.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "20201201"
func (redis Redis) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceKind returns the kind of the resource
func (redis *Redis) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (redis *Redis) GetSpec() genruntime.ConvertibleSpec {
	return &redis.Spec
}

// GetStatus returns the status of this resource
func (redis *Redis) GetStatus() genruntime.ConvertibleStatus {
	return &redis.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (redis *Redis) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (redis *Redis) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Redis_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (redis *Redis) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(redis.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  redis.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (redis *Redis) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Redis_STATUS); ok {
		redis.Status = *st
		return nil
	}

	// Convert status to required version
	var st Redis_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	redis.Status = st
	return nil
}

// Hub marks that this Redis is the hub type for conversion
func (redis *Redis) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (redis *Redis) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: redis.Spec.OriginalVersion,
		Kind:    "Redis",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20201201.Redis
// Generator information:
// - Generated from: /redis/resource-manager/Microsoft.Cache/stable/2020-12-01/redis.json
type RedisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Redis `json:"items"`
}

// Storage version of v1beta20201201.APIVersion
// +kubebuilder:validation:Enum={"20201201"}
type APIVersion string

const APIVersion_Value = APIVersion("20201201")

// Storage version of v1beta20201201.Redis_STATUS
type Redis_STATUS struct {
	Conditions          []conditions.Condition                           `json:"conditions,omitempty"`
	EnableNonSslPort    *bool                                            `json:"enableNonSslPort,omitempty"`
	Id                  *string                                          `json:"id,omitempty"`
	Location            *string                                          `json:"location,omitempty"`
	MinimumTlsVersion   *string                                          `json:"minimumTlsVersion,omitempty"`
	PropertyBag         genruntime.PropertyBag                           `json:"$propertyBag,omitempty"`
	PublicNetworkAccess *string                                          `json:"publicNetworkAccess,omitempty"`
	RedisConfiguration  *RedisCreateProperties_RedisConfiguration_STATUS `json:"redisConfiguration,omitempty"`
	RedisVersion        *string                                          `json:"redisVersion,omitempty"`
	ReplicasPerMaster   *int                                             `json:"replicasPerMaster,omitempty"`
	ReplicasPerPrimary  *int                                             `json:"replicasPerPrimary,omitempty"`
	ShardCount          *int                                             `json:"shardCount,omitempty"`
	Sku                 *Sku_STATUS                                      `json:"sku,omitempty"`
	StaticIP            *string                                          `json:"staticIP,omitempty"`
	SubnetId            *string                                          `json:"subnetId,omitempty"`
	Tags                map[string]string                                `json:"tags,omitempty"`
	TenantSettings      map[string]string                                `json:"tenantSettings,omitempty"`
	Zones               []string                                         `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Redis_STATUS{}

// ConvertStatusFrom populates our Redis_STATUS from the provided source
func (redis *Redis_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == redis {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(redis)
}

// ConvertStatusTo populates the provided destination from our Redis_STATUS
func (redis *Redis_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == redis {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(redis)
}

// Storage version of v1beta20201201.Redis_Spec
type Redis_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName         string             `json:"azureName,omitempty"`
	EnableNonSslPort  *bool              `json:"enableNonSslPort,omitempty"`
	Location          *string            `json:"location,omitempty"`
	MinimumTlsVersion *string            `json:"minimumTlsVersion,omitempty"`
	OperatorSpec      *RedisOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion   string             `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner               *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag         genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PublicNetworkAccess *string                            `json:"publicNetworkAccess,omitempty"`
	RedisConfiguration  map[string]string                  `json:"redisConfiguration,omitempty"`
	RedisVersion        *string                            `json:"redisVersion,omitempty"`
	ReplicasPerMaster   *int                               `json:"replicasPerMaster,omitempty"`
	ReplicasPerPrimary  *int                               `json:"replicasPerPrimary,omitempty"`
	ShardCount          *int                               `json:"shardCount,omitempty"`
	Sku                 *Sku                               `json:"sku,omitempty"`
	StaticIP            *string                            `json:"staticIP,omitempty"`

	// SubnetReference: The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
	SubnetReference *genruntime.ResourceReference `armReference:"SubnetId" json:"subnetReference,omitempty"`
	Tags            map[string]string             `json:"tags,omitempty"`
	TenantSettings  map[string]string             `json:"tenantSettings,omitempty"`
	Zones           []string                      `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Redis_Spec{}

// ConvertSpecFrom populates our Redis_Spec from the provided source
func (redis *Redis_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == redis {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(redis)
}

// ConvertSpecTo populates the provided destination from our Redis_Spec
func (redis *Redis_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == redis {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(redis)
}

// Storage version of v1beta20201201.RedisCreateProperties_RedisConfiguration_STATUS
type RedisCreateProperties_RedisConfiguration_STATUS struct {
	AdditionalProperties           map[string]string      `json:"additionalProperties,omitempty"`
	AofStorageConnectionString0    *string                `json:"aof-storage-connection-string-0,omitempty"`
	AofStorageConnectionString1    *string                `json:"aof-storage-connection-string-1,omitempty"`
	Maxclients                     *string                `json:"maxclients,omitempty"`
	MaxfragmentationmemoryReserved *string                `json:"maxfragmentationmemory-reserved,omitempty"`
	MaxmemoryDelta                 *string                `json:"maxmemory-delta,omitempty"`
	MaxmemoryPolicy                *string                `json:"maxmemory-policy,omitempty"`
	MaxmemoryReserved              *string                `json:"maxmemory-reserved,omitempty"`
	PropertyBag                    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RdbBackupEnabled               *string                `json:"rdb-backup-enabled,omitempty"`
	RdbBackupFrequency             *string                `json:"rdb-backup-frequency,omitempty"`
	RdbBackupMaxSnapshotCount      *string                `json:"rdb-backup-max-snapshot-count,omitempty"`
	RdbStorageConnectionString     *string                `json:"rdb-storage-connection-string,omitempty"`
	ZonalConfiguration             *string                `json:"zonal-configuration,omitempty"`
}

// Storage version of v1beta20201201.RedisOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type RedisOperatorSpec struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Secrets     *RedisOperatorSecrets  `json:"secrets,omitempty"`
}

// Storage version of v1beta20201201.Sku
type Sku struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Family      *string                `json:"family,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignPropertiesFromSku populates our Sku from the provided source Sku
func (sku *Sku) AssignPropertiesFromSku(source *v20210301s.Sku) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Capacity
	sku.Capacity = genruntime.ClonePointerToInt(source.Capacity)

	// Family
	if propertyBag.Contains("Family") {
		var family string
		err := propertyBag.Pull("Family", &family)
		if err != nil {
			return errors.Wrap(err, "pulling 'Family' from propertyBag")
		}

		sku.Family = &family
	} else {
		sku.Family = nil
	}

	// Name
	sku.Name = genruntime.ClonePointerToString(source.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		sku.PropertyBag = propertyBag
	} else {
		sku.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSku populates the provided destination Sku from our Sku
func (sku *Sku) AssignPropertiesToSku(destination *v20210301s.Sku) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(sku.PropertyBag)

	// Capacity
	destination.Capacity = genruntime.ClonePointerToInt(sku.Capacity)

	// Family
	if sku.Family != nil {
		propertyBag.Add("Family", *sku.Family)
	}

	// Name
	destination.Name = genruntime.ClonePointerToString(sku.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1beta20201201.Sku_STATUS
type Sku_STATUS struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Family      *string                `json:"family,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignPropertiesFromSku_STATUS populates our Sku_STATUS from the provided source Sku_STATUS
func (sku *Sku_STATUS) AssignPropertiesFromSku_STATUS(source *v20210301s.Sku_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Capacity
	sku.Capacity = genruntime.ClonePointerToInt(source.Capacity)

	// Family
	if propertyBag.Contains("Family") {
		var family string
		err := propertyBag.Pull("Family", &family)
		if err != nil {
			return errors.Wrap(err, "pulling 'Family' from propertyBag")
		}

		sku.Family = &family
	} else {
		sku.Family = nil
	}

	// Name
	sku.Name = genruntime.ClonePointerToString(source.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		sku.PropertyBag = propertyBag
	} else {
		sku.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSku_STATUS populates the provided destination Sku_STATUS from our Sku_STATUS
func (sku *Sku_STATUS) AssignPropertiesToSku_STATUS(destination *v20210301s.Sku_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(sku.PropertyBag)

	// Capacity
	destination.Capacity = genruntime.ClonePointerToInt(sku.Capacity)

	// Family
	if sku.Family != nil {
		propertyBag.Add("Family", *sku.Family)
	}

	// Name
	destination.Name = genruntime.ClonePointerToString(sku.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1beta20201201.RedisOperatorSecrets
type RedisOperatorSecrets struct {
	HostName     *genruntime.SecretDestination `json:"hostName,omitempty"`
	Port         *genruntime.SecretDestination `json:"port,omitempty"`
	PrimaryKey   *genruntime.SecretDestination `json:"primaryKey,omitempty"`
	PropertyBag  genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SSLPort      *genruntime.SecretDestination `json:"sslPort,omitempty"`
	SecondaryKey *genruntime.SecretDestination `json:"secondaryKey,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Redis{}, &RedisList{})
}

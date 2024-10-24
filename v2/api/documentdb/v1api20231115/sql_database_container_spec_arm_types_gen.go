// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20231115

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type SqlDatabaseContainer_Spec_ARM struct {
	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties to create and update Azure Cosmos DB container.
	Properties *SqlContainerCreateUpdateProperties_ARM `json:"properties,omitempty"`
	Tags       map[string]string                       `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &SqlDatabaseContainer_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-11-15"
func (container SqlDatabaseContainer_Spec_ARM) GetAPIVersion() string {
	return "2023-11-15"
}

// GetName returns the Name of the resource
func (container *SqlDatabaseContainer_Spec_ARM) GetName() string {
	return container.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers"
func (container *SqlDatabaseContainer_Spec_ARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers"
}

// Properties to create and update Azure Cosmos DB container.
type SqlContainerCreateUpdateProperties_ARM struct {
	// Options: A key-value pair of options to be applied for the request. This corresponds to the headers sent with the
	// request.
	Options *CreateUpdateOptions_ARM `json:"options,omitempty"`

	// Resource: The standard JSON format of a container
	Resource *SqlContainerResource_ARM `json:"resource,omitempty"`
}

// Cosmos DB SQL container resource object
type SqlContainerResource_ARM struct {
	// AnalyticalStorageTtl: Analytical TTL.
	AnalyticalStorageTtl *int `json:"analyticalStorageTtl,omitempty"`

	// ClientEncryptionPolicy: The client encryption policy for the container.
	ClientEncryptionPolicy *ClientEncryptionPolicy_ARM `json:"clientEncryptionPolicy,omitempty"`

	// ComputedProperties: List of computed properties
	ComputedProperties []ComputedProperty_ARM `json:"computedProperties,omitempty"`

	// ConflictResolutionPolicy: The conflict resolution policy for the container.
	ConflictResolutionPolicy *ConflictResolutionPolicy_ARM `json:"conflictResolutionPolicy,omitempty"`

	// CreateMode: Enum to indicate the mode of resource creation.
	CreateMode *CreateMode_ARM `json:"createMode,omitempty"`

	// DefaultTtl: Default time to live
	DefaultTtl *int `json:"defaultTtl,omitempty"`

	// Id: Name of the Cosmos DB SQL container
	Id *string `json:"id,omitempty"`

	// IndexingPolicy: The configuration of the indexing policy. By default, the indexing is automatic for all document paths
	// within the container
	IndexingPolicy *IndexingPolicy_ARM `json:"indexingPolicy,omitempty"`

	// PartitionKey: The configuration of the partition key to be used for partitioning data into multiple partitions
	PartitionKey *ContainerPartitionKey_ARM `json:"partitionKey,omitempty"`

	// RestoreParameters: Parameters to indicate the information about the restore
	RestoreParameters *RestoreParametersBase_ARM `json:"restoreParameters,omitempty"`

	// UniqueKeyPolicy: The unique key policy configuration for specifying uniqueness constraints on documents in the
	// collection in the Azure Cosmos DB service.
	UniqueKeyPolicy *UniqueKeyPolicy_ARM `json:"uniqueKeyPolicy,omitempty"`
}

// Cosmos DB client encryption policy.
type ClientEncryptionPolicy_ARM struct {
	// IncludedPaths: Paths of the item that need encryption along with path-specific settings.
	IncludedPaths []ClientEncryptionIncludedPath_ARM `json:"includedPaths,omitempty"`

	// PolicyFormatVersion: Version of the client encryption policy definition. Supported versions are 1 and 2. Version 2
	// supports id and partition key path encryption.
	PolicyFormatVersion *int `json:"policyFormatVersion,omitempty"`
}

// The definition of a computed property
type ComputedProperty_ARM struct {
	// Name: The name of a computed property, for example - "cp_lowerName"
	Name *string `json:"name,omitempty"`

	// Query: The query that evaluates the value for computed property, for example - "SELECT VALUE LOWER(c.name) FROM c"
	Query *string `json:"query,omitempty"`
}

// The conflict resolution policy for the container.
type ConflictResolutionPolicy_ARM struct {
	// ConflictResolutionPath: The conflict resolution path in the case of LastWriterWins mode.
	ConflictResolutionPath *string `json:"conflictResolutionPath,omitempty"`

	// ConflictResolutionProcedure: The procedure to resolve conflicts in the case of custom mode.
	ConflictResolutionProcedure *string `json:"conflictResolutionProcedure,omitempty"`

	// Mode: Indicates the conflict resolution mode.
	Mode *ConflictResolutionPolicy_Mode_ARM `json:"mode,omitempty"`
}

// The configuration of the partition key to be used for partitioning data into multiple partitions
type ContainerPartitionKey_ARM struct {
	// Kind: Indicates the kind of algorithm used for partitioning. For MultiHash, multiple partition keys (upto three maximum)
	// are supported for container create
	Kind *ContainerPartitionKey_Kind_ARM `json:"kind,omitempty"`

	// Paths: List of paths using which data within the container can be partitioned
	Paths []string `json:"paths,omitempty"`

	// Version: Indicates the version of the partition key definition
	Version *int `json:"version,omitempty"`
}

// Cosmos DB indexing policy
type IndexingPolicy_ARM struct {
	// Automatic: Indicates if the indexing policy is automatic
	Automatic *bool `json:"automatic,omitempty"`

	// CompositeIndexes: List of composite path list
	CompositeIndexes [][]CompositePath_ARM `json:"compositeIndexes,omitempty"`

	// ExcludedPaths: List of paths to exclude from indexing
	ExcludedPaths []ExcludedPath_ARM `json:"excludedPaths,omitempty"`

	// IncludedPaths: List of paths to include in the indexing
	IncludedPaths []IncludedPath_ARM `json:"includedPaths,omitempty"`

	// IndexingMode: Indicates the indexing mode.
	IndexingMode *IndexingPolicy_IndexingMode_ARM `json:"indexingMode,omitempty"`

	// SpatialIndexes: List of spatial specifics
	SpatialIndexes []SpatialSpec_ARM `json:"spatialIndexes,omitempty"`
}

// The unique key policy configuration for specifying uniqueness constraints on documents in the collection in the Azure
// Cosmos DB service.
type UniqueKeyPolicy_ARM struct {
	// UniqueKeys: List of unique keys on that enforces uniqueness constraint on documents in the collection in the Azure
	// Cosmos DB service.
	UniqueKeys []UniqueKey_ARM `json:"uniqueKeys,omitempty"`
}

// .
type ClientEncryptionIncludedPath_ARM struct {
	// ClientEncryptionKeyId: The identifier of the Client Encryption Key to be used to encrypt the path.
	ClientEncryptionKeyId *string `json:"clientEncryptionKeyId,omitempty"`

	// EncryptionAlgorithm: The encryption algorithm which will be used. Eg - AEAD_AES_256_CBC_HMAC_SHA256.
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty"`

	// EncryptionType: The type of encryption to be performed. Eg - Deterministic, Randomized.
	EncryptionType *string `json:"encryptionType,omitempty"`

	// Path: Path that needs to be encrypted.
	Path *string `json:"path,omitempty"`
}

type CompositePath_ARM struct {
	// Order: Sort order for composite paths.
	Order *CompositePath_Order_ARM `json:"order,omitempty"`

	// Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
	// (/path/*)
	Path *string `json:"path,omitempty"`
}

// +kubebuilder:validation:Enum={"Custom","LastWriterWins"}
type ConflictResolutionPolicy_Mode_ARM string

const (
	ConflictResolutionPolicy_Mode_ARM_Custom         = ConflictResolutionPolicy_Mode_ARM("Custom")
	ConflictResolutionPolicy_Mode_ARM_LastWriterWins = ConflictResolutionPolicy_Mode_ARM("LastWriterWins")
)

// Mapping from string to ConflictResolutionPolicy_Mode_ARM
var conflictResolutionPolicy_Mode_ARM_Values = map[string]ConflictResolutionPolicy_Mode_ARM{
	"custom":         ConflictResolutionPolicy_Mode_ARM_Custom,
	"lastwriterwins": ConflictResolutionPolicy_Mode_ARM_LastWriterWins,
}

// +kubebuilder:validation:Enum={"Hash","MultiHash","Range"}
type ContainerPartitionKey_Kind_ARM string

const (
	ContainerPartitionKey_Kind_ARM_Hash      = ContainerPartitionKey_Kind_ARM("Hash")
	ContainerPartitionKey_Kind_ARM_MultiHash = ContainerPartitionKey_Kind_ARM("MultiHash")
	ContainerPartitionKey_Kind_ARM_Range     = ContainerPartitionKey_Kind_ARM("Range")
)

// Mapping from string to ContainerPartitionKey_Kind_ARM
var containerPartitionKey_Kind_ARM_Values = map[string]ContainerPartitionKey_Kind_ARM{
	"hash":      ContainerPartitionKey_Kind_ARM_Hash,
	"multihash": ContainerPartitionKey_Kind_ARM_MultiHash,
	"range":     ContainerPartitionKey_Kind_ARM_Range,
}

type ExcludedPath_ARM struct {
	// Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
	// (/path/*)
	Path *string `json:"path,omitempty"`
}

// The paths that are included in indexing
type IncludedPath_ARM struct {
	// Indexes: List of indexes for this path
	Indexes []Indexes_ARM `json:"indexes,omitempty"`

	// Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
	// (/path/*)
	Path *string `json:"path,omitempty"`
}

// +kubebuilder:validation:Enum={"consistent","lazy","none"}
type IndexingPolicy_IndexingMode_ARM string

const (
	IndexingPolicy_IndexingMode_ARM_Consistent = IndexingPolicy_IndexingMode_ARM("consistent")
	IndexingPolicy_IndexingMode_ARM_Lazy       = IndexingPolicy_IndexingMode_ARM("lazy")
	IndexingPolicy_IndexingMode_ARM_None       = IndexingPolicy_IndexingMode_ARM("none")
)

// Mapping from string to IndexingPolicy_IndexingMode_ARM
var indexingPolicy_IndexingMode_ARM_Values = map[string]IndexingPolicy_IndexingMode_ARM{
	"consistent": IndexingPolicy_IndexingMode_ARM_Consistent,
	"lazy":       IndexingPolicy_IndexingMode_ARM_Lazy,
	"none":       IndexingPolicy_IndexingMode_ARM_None,
}

type SpatialSpec_ARM struct {
	// Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
	// (/path/*)
	Path *string `json:"path,omitempty"`

	// Types: List of path's spatial type
	Types []SpatialType_ARM `json:"types,omitempty"`
}

// The unique key on that enforces uniqueness constraint on documents in the collection in the Azure Cosmos DB service.
type UniqueKey_ARM struct {
	// Paths: List of paths must be unique for each document in the Azure Cosmos DB service
	Paths []string `json:"paths,omitempty"`
}

// +kubebuilder:validation:Enum={"ascending","descending"}
type CompositePath_Order_ARM string

const (
	CompositePath_Order_ARM_Ascending  = CompositePath_Order_ARM("ascending")
	CompositePath_Order_ARM_Descending = CompositePath_Order_ARM("descending")
)

// Mapping from string to CompositePath_Order_ARM
var compositePath_Order_ARM_Values = map[string]CompositePath_Order_ARM{
	"ascending":  CompositePath_Order_ARM_Ascending,
	"descending": CompositePath_Order_ARM_Descending,
}

// The indexes for the path.
type Indexes_ARM struct {
	// DataType: The datatype for which the indexing behavior is applied to.
	DataType *Indexes_DataType_ARM `json:"dataType,omitempty"`

	// Kind: Indicates the type of index.
	Kind *Indexes_Kind_ARM `json:"kind,omitempty"`

	// Precision: The precision of the index. -1 is maximum precision.
	Precision *int `json:"precision,omitempty"`
}

// Indicates the spatial type of index.
// +kubebuilder:validation:Enum={"LineString","MultiPolygon","Point","Polygon"}
type SpatialType_ARM string

const (
	SpatialType_ARM_LineString   = SpatialType_ARM("LineString")
	SpatialType_ARM_MultiPolygon = SpatialType_ARM("MultiPolygon")
	SpatialType_ARM_Point        = SpatialType_ARM("Point")
	SpatialType_ARM_Polygon      = SpatialType_ARM("Polygon")
)

// Mapping from string to SpatialType_ARM
var spatialType_ARM_Values = map[string]SpatialType_ARM{
	"linestring":   SpatialType_ARM_LineString,
	"multipolygon": SpatialType_ARM_MultiPolygon,
	"point":        SpatialType_ARM_Point,
	"polygon":      SpatialType_ARM_Polygon,
}

// +kubebuilder:validation:Enum={"LineString","MultiPolygon","Number","Point","Polygon","String"}
type Indexes_DataType_ARM string

const (
	Indexes_DataType_ARM_LineString   = Indexes_DataType_ARM("LineString")
	Indexes_DataType_ARM_MultiPolygon = Indexes_DataType_ARM("MultiPolygon")
	Indexes_DataType_ARM_Number       = Indexes_DataType_ARM("Number")
	Indexes_DataType_ARM_Point        = Indexes_DataType_ARM("Point")
	Indexes_DataType_ARM_Polygon      = Indexes_DataType_ARM("Polygon")
	Indexes_DataType_ARM_String       = Indexes_DataType_ARM("String")
)

// Mapping from string to Indexes_DataType_ARM
var indexes_DataType_ARM_Values = map[string]Indexes_DataType_ARM{
	"linestring":   Indexes_DataType_ARM_LineString,
	"multipolygon": Indexes_DataType_ARM_MultiPolygon,
	"number":       Indexes_DataType_ARM_Number,
	"point":        Indexes_DataType_ARM_Point,
	"polygon":      Indexes_DataType_ARM_Polygon,
	"string":       Indexes_DataType_ARM_String,
}

// +kubebuilder:validation:Enum={"Hash","Range","Spatial"}
type Indexes_Kind_ARM string

const (
	Indexes_Kind_ARM_Hash    = Indexes_Kind_ARM("Hash")
	Indexes_Kind_ARM_Range   = Indexes_Kind_ARM("Range")
	Indexes_Kind_ARM_Spatial = Indexes_Kind_ARM("Spatial")
)

// Mapping from string to Indexes_Kind_ARM
var indexes_Kind_ARM_Values = map[string]Indexes_Kind_ARM{
	"hash":    Indexes_Kind_ARM_Hash,
	"range":   Indexes_Kind_ARM_Range,
	"spatial": Indexes_Kind_ARM_Spatial,
}

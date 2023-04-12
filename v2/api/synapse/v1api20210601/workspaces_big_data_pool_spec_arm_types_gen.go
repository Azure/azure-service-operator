// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Workspaces_BigDataPool_Spec_ARM struct {
	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Big Data pool properties
	Properties *BigDataPoolResourceProperties_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Workspaces_BigDataPool_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (pool Workspaces_BigDataPool_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (pool *Workspaces_BigDataPool_Spec_ARM) GetName() string {
	return pool.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Synapse/workspaces/bigDataPools"
func (pool *Workspaces_BigDataPool_Spec_ARM) GetType() string {
	return "Microsoft.Synapse/workspaces/bigDataPools"
}

// Properties of a Big Data pool powered by Apache Spark
type BigDataPoolResourceProperties_ARM struct {
	// AutoPause: Auto-pausing properties
	AutoPause *AutoPauseProperties_ARM `json:"autoPause,omitempty"`

	// AutoScale: Auto-scaling properties
	AutoScale *AutoScaleProperties_ARM `json:"autoScale,omitempty"`

	// CustomLibraries: List of custom libraries/packages associated with the spark pool.
	CustomLibraries []LibraryInfo_ARM `json:"customLibraries,omitempty"`

	// DefaultSparkLogFolder: The default folder where Spark logs will be written.
	DefaultSparkLogFolder *string `json:"defaultSparkLogFolder,omitempty"`

	// DynamicExecutorAllocation: Dynamic Executor Allocation
	DynamicExecutorAllocation *DynamicExecutorAllocation_ARM `json:"dynamicExecutorAllocation,omitempty"`

	// IsAutotuneEnabled: Whether autotune is required or not.
	IsAutotuneEnabled *bool `json:"isAutotuneEnabled,omitempty"`

	// IsComputeIsolationEnabled: Whether compute isolation is required or not.
	IsComputeIsolationEnabled *bool `json:"isComputeIsolationEnabled,omitempty"`

	// LibraryRequirements: Library version requirements
	LibraryRequirements *LibraryRequirements_ARM `json:"libraryRequirements,omitempty"`

	// NodeCount: The number of nodes in the Big Data pool.
	NodeCount *int `json:"nodeCount,omitempty"`

	// NodeSize: The level of compute power that each node in the Big Data pool has.
	NodeSize *BigDataPoolResourceProperties_NodeSize `json:"nodeSize,omitempty"`

	// NodeSizeFamily: The kind of nodes that the Big Data pool provides.
	NodeSizeFamily *BigDataPoolResourceProperties_NodeSizeFamily `json:"nodeSizeFamily,omitempty"`

	// ProvisioningState: The state of the Big Data pool.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// SessionLevelPackagesEnabled: Whether session level packages enabled.
	SessionLevelPackagesEnabled *bool `json:"sessionLevelPackagesEnabled,omitempty"`

	// SparkConfigProperties: Spark configuration file to specify additional properties
	SparkConfigProperties *SparkConfigProperties_ARM `json:"sparkConfigProperties,omitempty"`

	// SparkEventsFolder: The Spark events folder
	SparkEventsFolder *string `json:"sparkEventsFolder,omitempty"`

	// SparkVersion: The Apache Spark version.
	SparkVersion *string `json:"sparkVersion,omitempty"`
}

// Auto-pausing properties of a Big Data pool powered by Apache Spark
type AutoPauseProperties_ARM struct {
	// DelayInMinutes: Number of minutes of idle time before the Big Data pool is automatically paused.
	DelayInMinutes *int `json:"delayInMinutes,omitempty"`

	// Enabled: Whether auto-pausing is enabled for the Big Data pool.
	Enabled *bool `json:"enabled,omitempty"`
}

// Auto-scaling properties of a Big Data pool powered by Apache Spark
type AutoScaleProperties_ARM struct {
	// Enabled: Whether automatic scaling is enabled for the Big Data pool.
	Enabled *bool `json:"enabled,omitempty"`

	// MaxNodeCount: The maximum number of nodes the Big Data pool can support.
	MaxNodeCount *int `json:"maxNodeCount,omitempty"`

	// MinNodeCount: The minimum number of nodes the Big Data pool can support.
	MinNodeCount *int `json:"minNodeCount,omitempty"`
}

// Dynamic Executor Allocation Properties
type DynamicExecutorAllocation_ARM struct {
	// Enabled: Indicates whether Dynamic Executor Allocation is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`

	// MaxExecutors: The maximum number of executors alloted
	MaxExecutors *int `json:"maxExecutors,omitempty"`

	// MinExecutors: The minimum number of executors alloted
	MinExecutors *int `json:"minExecutors,omitempty"`
}

// Library/package information of a Big Data pool powered by Apache Spark
type LibraryInfo_ARM struct {
	// ContainerName: Storage blob container name.
	ContainerName *string `json:"containerName,omitempty"`

	// Name: Name of the library.
	Name *string `json:"name,omitempty"`

	// Path: Storage blob path of library.
	Path *string `json:"path,omitempty"`

	// Type: Type of the library.
	Type *string `json:"type,omitempty"`
}

// Library requirements for a Big Data pool powered by Apache Spark
type LibraryRequirements_ARM struct {
	// Content: The library requirements.
	Content *string `json:"content,omitempty"`

	// Filename: The filename of the library requirements file.
	Filename *string `json:"filename,omitempty"`
}

// SparkConfig Properties for a Big Data pool powered by Apache Spark
type SparkConfigProperties_ARM struct {
	// ConfigurationType: The type of the spark config properties file.
	ConfigurationType *SparkConfigProperties_ConfigurationType `json:"configurationType,omitempty"`

	// Content: The spark config properties.
	Content *string `json:"content,omitempty"`

	// Filename: The filename of the spark config properties file.
	Filename *string `json:"filename,omitempty"`
}

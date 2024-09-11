// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230315preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Fleets_UpdateRun_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: The resource-specific properties for this resource.
	Properties *UpdateRunProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Fleets_UpdateRun_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-03-15-preview"
func (updateRun Fleets_UpdateRun_Spec_ARM) GetAPIVersion() string {
	return "2023-03-15-preview"
}

// GetName returns the Name of the resource
func (updateRun *Fleets_UpdateRun_Spec_ARM) GetName() string {
	return updateRun.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerService/fleets/updateRuns"
func (updateRun *Fleets_UpdateRun_Spec_ARM) GetType() string {
	return "Microsoft.ContainerService/fleets/updateRuns"
}

// The properties of the UpdateRun.
type UpdateRunProperties_ARM struct {
	// ManagedClusterUpdate: The update to be applied to all clusters in the UpdateRun. The managedClusterUpdate can be
	// modified until the run is started.
	ManagedClusterUpdate *ManagedClusterUpdate_ARM `json:"managedClusterUpdate,omitempty"`

	// Strategy: The strategy defines the order in which the clusters will be updated.
	// If not set, all members will be updated sequentially. The UpdateRun status will show a single UpdateStage and a single
	// UpdateGroup targeting all members.
	// The strategy of the UpdateRun can be modified until the run is started.
	Strategy *UpdateRunStrategy_ARM `json:"strategy,omitempty"`
}

// The update to be applied to the ManagedClusters.
type ManagedClusterUpdate_ARM struct {
	// Upgrade: The upgrade to apply to the ManagedClusters.
	Upgrade *ManagedClusterUpgradeSpec_ARM `json:"upgrade,omitempty"`
}

// Defines the update sequence of the clusters via stages and groups.
// Stages within a run are executed sequentially one
// after another.
// Groups within a stage are executed in parallel.
// Member clusters within a group are updated sequentially
// one after another.
// A valid strategy contains no duplicate groups within or across stages.
type UpdateRunStrategy_ARM struct {
	// Stages: The list of stages that compose this update run. Min size: 1.
	Stages []UpdateStage_ARM `json:"stages"`
}

// The upgrade to apply to a ManagedCluster.
type ManagedClusterUpgradeSpec_ARM struct {
	// KubernetesVersion: The Kubernetes version to upgrade the member clusters to.
	KubernetesVersion *string `json:"kubernetesVersion,omitempty"`

	// Type: ManagedClusterUpgradeType is the type of upgrade to be applied.
	Type *ManagedClusterUpgradeType_ARM `json:"type,omitempty"`
}

// Defines a stage which contains the groups to update and the steps to take (e.g., wait for a time period) before starting
// the next stage.
type UpdateStage_ARM struct {
	// AfterStageWaitInSeconds: The time in seconds to wait at the end of this stage before starting the next one. Defaults to
	// 0 seconds if unspecified.
	AfterStageWaitInSeconds *int `json:"afterStageWaitInSeconds,omitempty"`

	// Groups: Defines the groups to be executed in parallel in this stage. Duplicate groups are not allowed. Min size: 1.
	Groups []UpdateGroup_ARM `json:"groups"`

	// Name: The name of the stage. Must be unique within the UpdateRun.
	Name *string `json:"name,omitempty"`
}

// The type of upgrade to perform when targeting ManagedClusters.
// +kubebuilder:validation:Enum={"Full","NodeImageOnly"}
type ManagedClusterUpgradeType_ARM string

const (
	ManagedClusterUpgradeType_ARM_Full          = ManagedClusterUpgradeType_ARM("Full")
	ManagedClusterUpgradeType_ARM_NodeImageOnly = ManagedClusterUpgradeType_ARM("NodeImageOnly")
)

// Mapping from string to ManagedClusterUpgradeType_ARM
var managedClusterUpgradeType_ARM_Values = map[string]ManagedClusterUpgradeType_ARM{
	"full":          ManagedClusterUpgradeType_ARM_Full,
	"nodeimageonly": ManagedClusterUpgradeType_ARM_NodeImageOnly,
}

// A group to be updated.
type UpdateGroup_ARM struct {
	// Name: Name of the group.
	// It must match a group name of an existing fleet member.
	Name *string `json:"name,omitempty"`
}

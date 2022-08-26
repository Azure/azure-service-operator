// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210501

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_ManagedClustersAgentPool_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<< HEAD:v2/api/containerservice/v1beta20210501/managed_clusters_agent_pool_status_arm_types_gen_test.go
		"Round trip of ManagedClustersAgentPool_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedClustersAgentPool_STATUSARM, ManagedClustersAgentPool_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedClustersAgentPool_STATUSARM runs a test to see if a specific instance of ManagedClustersAgentPool_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedClustersAgentPool_STATUSARM(subject ManagedClustersAgentPool_STATUSARM) string {
=======
		"Round trip of AgentPool_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPool_STATUSARM, AgentPool_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPool_STATUSARM runs a test to see if a specific instance of AgentPool_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPool_STATUSARM(subject AgentPool_STATUSARM) string {
>>>>>>> main:v2/api/containerservice/v1beta20210501/agent_pool_status_arm_types_gen_test.go
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedClustersAgentPool_STATUSARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

<<<<<<< HEAD:v2/api/containerservice/v1beta20210501/managed_clusters_agent_pool_status_arm_types_gen_test.go
// Generator of ManagedClustersAgentPool_STATUSARM instances for property testing - lazily instantiated by
// ManagedClustersAgentPool_STATUSARMGenerator()
var managedClustersAgentPool_STATUSARMGenerator gopter.Gen

// ManagedClustersAgentPool_STATUSARMGenerator returns a generator of ManagedClustersAgentPool_STATUSARM instances for property testing.
// We first initialize managedClustersAgentPool_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedClustersAgentPool_STATUSARMGenerator() gopter.Gen {
	if managedClustersAgentPool_STATUSARMGenerator != nil {
		return managedClustersAgentPool_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClustersAgentPool_STATUSARM(generators)
	managedClustersAgentPool_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ManagedClustersAgentPool_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClustersAgentPool_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForManagedClustersAgentPool_STATUSARM(generators)
	managedClustersAgentPool_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ManagedClustersAgentPool_STATUSARM{}), generators)

	return managedClustersAgentPool_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForManagedClustersAgentPool_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedClustersAgentPool_STATUSARM(gens map[string]gopter.Gen) {
=======
// Generator of AgentPool_STATUSARM instances for property testing - lazily instantiated by
// AgentPool_STATUSARMGenerator()
var agentPool_STATUSARMGenerator gopter.Gen

// AgentPool_STATUSARMGenerator returns a generator of AgentPool_STATUSARM instances for property testing.
// We first initialize agentPool_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AgentPool_STATUSARMGenerator() gopter.Gen {
	if agentPool_STATUSARMGenerator != nil {
		return agentPool_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPool_STATUSARM(generators)
	agentPool_STATUSARMGenerator = gen.Struct(reflect.TypeOf(AgentPool_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPool_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForAgentPool_STATUSARM(generators)
	agentPool_STATUSARMGenerator = gen.Struct(reflect.TypeOf(AgentPool_STATUSARM{}), generators)

	return agentPool_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForAgentPool_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPool_STATUSARM(gens map[string]gopter.Gen) {
>>>>>>> main:v2/api/containerservice/v1beta20210501/agent_pool_status_arm_types_gen_test.go
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

<<<<<<< HEAD:v2/api/containerservice/v1beta20210501/managed_clusters_agent_pool_status_arm_types_gen_test.go
// AddRelatedPropertyGeneratorsForManagedClustersAgentPool_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedClustersAgentPool_STATUSARM(gens map[string]gopter.Gen) {
=======
// AddRelatedPropertyGeneratorsForAgentPool_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAgentPool_STATUSARM(gens map[string]gopter.Gen) {
>>>>>>> main:v2/api/containerservice/v1beta20210501/agent_pool_status_arm_types_gen_test.go
	gens["Properties"] = gen.PtrOf(ManagedClusterAgentPoolProfileProperties_STATUSARMGenerator())
}

func Test_ManagedClusterAgentPoolProfileProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedClusterAgentPoolProfileProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedClusterAgentPoolProfileProperties_STATUSARM, ManagedClusterAgentPoolProfileProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedClusterAgentPoolProfileProperties_STATUSARM runs a test to see if a specific instance of ManagedClusterAgentPoolProfileProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedClusterAgentPoolProfileProperties_STATUSARM(subject ManagedClusterAgentPoolProfileProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedClusterAgentPoolProfileProperties_STATUSARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ManagedClusterAgentPoolProfileProperties_STATUSARM instances for property testing - lazily instantiated
// by ManagedClusterAgentPoolProfileProperties_STATUSARMGenerator()
var managedClusterAgentPoolProfileProperties_STATUSARMGenerator gopter.Gen

// ManagedClusterAgentPoolProfileProperties_STATUSARMGenerator returns a generator of ManagedClusterAgentPoolProfileProperties_STATUSARM instances for property testing.
// We first initialize managedClusterAgentPoolProfileProperties_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedClusterAgentPoolProfileProperties_STATUSARMGenerator() gopter.Gen {
	if managedClusterAgentPoolProfileProperties_STATUSARMGenerator != nil {
		return managedClusterAgentPoolProfileProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_STATUSARM(generators)
	managedClusterAgentPoolProfileProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ManagedClusterAgentPoolProfileProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_STATUSARM(generators)
	managedClusterAgentPoolProfileProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ManagedClusterAgentPoolProfileProperties_STATUSARM{}), generators)

	return managedClusterAgentPoolProfileProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["AvailabilityZones"] = gen.SliceOf(gen.AlphaString())
	gens["Count"] = gen.PtrOf(gen.Int())
	gens["EnableAutoScaling"] = gen.PtrOf(gen.Bool())
	gens["EnableEncryptionAtHost"] = gen.PtrOf(gen.Bool())
	gens["EnableFIPS"] = gen.PtrOf(gen.Bool())
	gens["EnableNodePublicIP"] = gen.PtrOf(gen.Bool())
	gens["EnableUltraSSD"] = gen.PtrOf(gen.Bool())
	gens["GpuInstanceProfile"] = gen.PtrOf(gen.OneConstOf(
		GPUInstanceProfile_MIG1G_STATUS,
		GPUInstanceProfile_MIG2G_STATUS,
		GPUInstanceProfile_MIG3G_STATUS,
		GPUInstanceProfile_MIG4G_STATUS,
		GPUInstanceProfile_MIG7G_STATUS))
	gens["KubeletDiskType"] = gen.PtrOf(gen.OneConstOf(KubeletDiskType_OS_STATUS, KubeletDiskType_Temporary_STATUS))
	gens["MaxCount"] = gen.PtrOf(gen.Int())
	gens["MaxPods"] = gen.PtrOf(gen.Int())
	gens["MinCount"] = gen.PtrOf(gen.Int())
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(AgentPoolMode_System_STATUS, AgentPoolMode_User_STATUS))
	gens["NodeImageVersion"] = gen.PtrOf(gen.AlphaString())
	gens["NodeLabels"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["NodePublicIPPrefixID"] = gen.PtrOf(gen.AlphaString())
	gens["NodeTaints"] = gen.SliceOf(gen.AlphaString())
	gens["OrchestratorVersion"] = gen.PtrOf(gen.AlphaString())
	gens["OsDiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["OsDiskType"] = gen.PtrOf(gen.OneConstOf(OSDiskType_Ephemeral_STATUS, OSDiskType_Managed_STATUS))
	gens["OsSKU"] = gen.PtrOf(gen.OneConstOf(OSSKU_CBLMariner_STATUS, OSSKU_Ubuntu_STATUS))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(OSType_Linux_STATUS, OSType_Windows_STATUS))
	gens["PodSubnetID"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ProximityPlacementGroupID"] = gen.PtrOf(gen.AlphaString())
	gens["ScaleSetEvictionPolicy"] = gen.PtrOf(gen.OneConstOf(ScaleSetEvictionPolicy_Deallocate_STATUS, ScaleSetEvictionPolicy_Delete_STATUS))
	gens["ScaleSetPriority"] = gen.PtrOf(gen.OneConstOf(ScaleSetPriority_Regular_STATUS, ScaleSetPriority_Spot_STATUS))
	gens["SpotMaxPrice"] = gen.PtrOf(gen.Float64())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(AgentPoolType_AvailabilitySet_STATUS, AgentPoolType_VirtualMachineScaleSets_STATUS))
	gens["VmSize"] = gen.PtrOf(gen.AlphaString())
	gens["VnetSubnetID"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["KubeletConfig"] = gen.PtrOf(KubeletConfig_STATUSARMGenerator())
	gens["LinuxOSConfig"] = gen.PtrOf(LinuxOSConfig_STATUSARMGenerator())
	gens["PowerState"] = gen.PtrOf(PowerState_STATUSARMGenerator())
	gens["UpgradeSettings"] = gen.PtrOf(AgentPoolUpgradeSettings_STATUSARMGenerator())
}

func Test_AgentPoolUpgradeSettings_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AgentPoolUpgradeSettings_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPoolUpgradeSettings_STATUSARM, AgentPoolUpgradeSettings_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPoolUpgradeSettings_STATUSARM runs a test to see if a specific instance of AgentPoolUpgradeSettings_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPoolUpgradeSettings_STATUSARM(subject AgentPoolUpgradeSettings_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AgentPoolUpgradeSettings_STATUSARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AgentPoolUpgradeSettings_STATUSARM instances for property testing - lazily instantiated by
// AgentPoolUpgradeSettings_STATUSARMGenerator()
var agentPoolUpgradeSettings_STATUSARMGenerator gopter.Gen

// AgentPoolUpgradeSettings_STATUSARMGenerator returns a generator of AgentPoolUpgradeSettings_STATUSARM instances for property testing.
func AgentPoolUpgradeSettings_STATUSARMGenerator() gopter.Gen {
	if agentPoolUpgradeSettings_STATUSARMGenerator != nil {
		return agentPoolUpgradeSettings_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings_STATUSARM(generators)
	agentPoolUpgradeSettings_STATUSARMGenerator = gen.Struct(reflect.TypeOf(AgentPoolUpgradeSettings_STATUSARM{}), generators)

	return agentPoolUpgradeSettings_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings_STATUSARM(gens map[string]gopter.Gen) {
	gens["MaxSurge"] = gen.PtrOf(gen.AlphaString())
}

func Test_KubeletConfig_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KubeletConfig_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKubeletConfig_STATUSARM, KubeletConfig_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKubeletConfig_STATUSARM runs a test to see if a specific instance of KubeletConfig_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKubeletConfig_STATUSARM(subject KubeletConfig_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KubeletConfig_STATUSARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of KubeletConfig_STATUSARM instances for property testing - lazily instantiated by
// KubeletConfig_STATUSARMGenerator()
var kubeletConfig_STATUSARMGenerator gopter.Gen

// KubeletConfig_STATUSARMGenerator returns a generator of KubeletConfig_STATUSARM instances for property testing.
func KubeletConfig_STATUSARMGenerator() gopter.Gen {
	if kubeletConfig_STATUSARMGenerator != nil {
		return kubeletConfig_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKubeletConfig_STATUSARM(generators)
	kubeletConfig_STATUSARMGenerator = gen.Struct(reflect.TypeOf(KubeletConfig_STATUSARM{}), generators)

	return kubeletConfig_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForKubeletConfig_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKubeletConfig_STATUSARM(gens map[string]gopter.Gen) {
	gens["AllowedUnsafeSysctls"] = gen.SliceOf(gen.AlphaString())
	gens["ContainerLogMaxFiles"] = gen.PtrOf(gen.Int())
	gens["ContainerLogMaxSizeMB"] = gen.PtrOf(gen.Int())
	gens["CpuCfsQuota"] = gen.PtrOf(gen.Bool())
	gens["CpuCfsQuotaPeriod"] = gen.PtrOf(gen.AlphaString())
	gens["CpuManagerPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["FailSwapOn"] = gen.PtrOf(gen.Bool())
	gens["ImageGcHighThreshold"] = gen.PtrOf(gen.Int())
	gens["ImageGcLowThreshold"] = gen.PtrOf(gen.Int())
	gens["PodMaxPids"] = gen.PtrOf(gen.Int())
	gens["TopologyManagerPolicy"] = gen.PtrOf(gen.AlphaString())
}

func Test_LinuxOSConfig_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LinuxOSConfig_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLinuxOSConfig_STATUSARM, LinuxOSConfig_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLinuxOSConfig_STATUSARM runs a test to see if a specific instance of LinuxOSConfig_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLinuxOSConfig_STATUSARM(subject LinuxOSConfig_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LinuxOSConfig_STATUSARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of LinuxOSConfig_STATUSARM instances for property testing - lazily instantiated by
// LinuxOSConfig_STATUSARMGenerator()
var linuxOSConfig_STATUSARMGenerator gopter.Gen

// LinuxOSConfig_STATUSARMGenerator returns a generator of LinuxOSConfig_STATUSARM instances for property testing.
// We first initialize linuxOSConfig_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LinuxOSConfig_STATUSARMGenerator() gopter.Gen {
	if linuxOSConfig_STATUSARMGenerator != nil {
		return linuxOSConfig_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfig_STATUSARM(generators)
	linuxOSConfig_STATUSARMGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfig_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfig_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForLinuxOSConfig_STATUSARM(generators)
	linuxOSConfig_STATUSARMGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfig_STATUSARM{}), generators)

	return linuxOSConfig_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForLinuxOSConfig_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLinuxOSConfig_STATUSARM(gens map[string]gopter.Gen) {
	gens["SwapFileSizeMB"] = gen.PtrOf(gen.Int())
	gens["TransparentHugePageDefrag"] = gen.PtrOf(gen.AlphaString())
	gens["TransparentHugePageEnabled"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForLinuxOSConfig_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLinuxOSConfig_STATUSARM(gens map[string]gopter.Gen) {
	gens["Sysctls"] = gen.PtrOf(SysctlConfig_STATUSARMGenerator())
<<<<<<< HEAD:v2/api/containerservice/v1beta20210501/managed_clusters_agent_pool_status_arm_types_gen_test.go
=======
}

func Test_PowerState_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PowerState_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPowerState_STATUSARM, PowerState_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPowerState_STATUSARM runs a test to see if a specific instance of PowerState_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPowerState_STATUSARM(subject PowerState_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PowerState_STATUSARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of PowerState_STATUSARM instances for property testing - lazily instantiated by
// PowerState_STATUSARMGenerator()
var powerState_STATUSARMGenerator gopter.Gen

// PowerState_STATUSARMGenerator returns a generator of PowerState_STATUSARM instances for property testing.
func PowerState_STATUSARMGenerator() gopter.Gen {
	if powerState_STATUSARMGenerator != nil {
		return powerState_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPowerState_STATUSARM(generators)
	powerState_STATUSARMGenerator = gen.Struct(reflect.TypeOf(PowerState_STATUSARM{}), generators)

	return powerState_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForPowerState_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPowerState_STATUSARM(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.OneConstOf(PowerState_STATUS_Code_Running, PowerState_STATUS_Code_Stopped))
>>>>>>> main:v2/api/containerservice/v1beta20210501/agent_pool_status_arm_types_gen_test.go
}

func Test_SysctlConfig_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SysctlConfig_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSysctlConfig_STATUSARM, SysctlConfig_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSysctlConfig_STATUSARM runs a test to see if a specific instance of SysctlConfig_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSysctlConfig_STATUSARM(subject SysctlConfig_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SysctlConfig_STATUSARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SysctlConfig_STATUSARM instances for property testing - lazily instantiated by
// SysctlConfig_STATUSARMGenerator()
var sysctlConfig_STATUSARMGenerator gopter.Gen

// SysctlConfig_STATUSARMGenerator returns a generator of SysctlConfig_STATUSARM instances for property testing.
func SysctlConfig_STATUSARMGenerator() gopter.Gen {
	if sysctlConfig_STATUSARMGenerator != nil {
		return sysctlConfig_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSysctlConfig_STATUSARM(generators)
	sysctlConfig_STATUSARMGenerator = gen.Struct(reflect.TypeOf(SysctlConfig_STATUSARM{}), generators)

	return sysctlConfig_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSysctlConfig_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSysctlConfig_STATUSARM(gens map[string]gopter.Gen) {
	gens["FsAioMaxNr"] = gen.PtrOf(gen.Int())
	gens["FsFileMax"] = gen.PtrOf(gen.Int())
	gens["FsInotifyMaxUserWatches"] = gen.PtrOf(gen.Int())
	gens["FsNrOpen"] = gen.PtrOf(gen.Int())
	gens["KernelThreadsMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreNetdevMaxBacklog"] = gen.PtrOf(gen.Int())
	gens["NetCoreOptmemMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreRmemDefault"] = gen.PtrOf(gen.Int())
	gens["NetCoreRmemMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreSomaxconn"] = gen.PtrOf(gen.Int())
	gens["NetCoreWmemDefault"] = gen.PtrOf(gen.Int())
	gens["NetCoreWmemMax"] = gen.PtrOf(gen.Int())
	gens["NetIpv4IpLocalPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["NetIpv4NeighDefaultGcThresh1"] = gen.PtrOf(gen.Int())
	gens["NetIpv4NeighDefaultGcThresh2"] = gen.PtrOf(gen.Int())
	gens["NetIpv4NeighDefaultGcThresh3"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpFinTimeout"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpKeepaliveProbes"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpKeepaliveTime"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpMaxSynBacklog"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpMaxTwBuckets"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpTwReuse"] = gen.PtrOf(gen.Bool())
	gens["NetIpv4TcpkeepaliveIntvl"] = gen.PtrOf(gen.Int())
	gens["NetNetfilterNfConntrackBuckets"] = gen.PtrOf(gen.Int())
	gens["NetNetfilterNfConntrackMax"] = gen.PtrOf(gen.Int())
	gens["VmMaxMapCount"] = gen.PtrOf(gen.Int())
	gens["VmSwappiness"] = gen.PtrOf(gen.Int())
	gens["VmVfsCachePressure"] = gen.PtrOf(gen.Int())
}

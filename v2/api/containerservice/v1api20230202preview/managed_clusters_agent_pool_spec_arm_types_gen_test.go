// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230202preview

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

func Test_ManagedClusters_AgentPool_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedClusters_AgentPool_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedClusters_AgentPool_Spec_ARM, ManagedClusters_AgentPool_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedClusters_AgentPool_Spec_ARM runs a test to see if a specific instance of ManagedClusters_AgentPool_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedClusters_AgentPool_Spec_ARM(subject ManagedClusters_AgentPool_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedClusters_AgentPool_Spec_ARM
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

// Generator of ManagedClusters_AgentPool_Spec_ARM instances for property testing - lazily instantiated by
// ManagedClusters_AgentPool_Spec_ARMGenerator()
var managedClusters_AgentPool_Spec_ARMGenerator gopter.Gen

// ManagedClusters_AgentPool_Spec_ARMGenerator returns a generator of ManagedClusters_AgentPool_Spec_ARM instances for property testing.
// We first initialize managedClusters_AgentPool_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedClusters_AgentPool_Spec_ARMGenerator() gopter.Gen {
	if managedClusters_AgentPool_Spec_ARMGenerator != nil {
		return managedClusters_AgentPool_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusters_AgentPool_Spec_ARM(generators)
	managedClusters_AgentPool_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(ManagedClusters_AgentPool_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusters_AgentPool_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForManagedClusters_AgentPool_Spec_ARM(generators)
	managedClusters_AgentPool_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(ManagedClusters_AgentPool_Spec_ARM{}), generators)

	return managedClusters_AgentPool_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForManagedClusters_AgentPool_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedClusters_AgentPool_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForManagedClusters_AgentPool_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedClusters_AgentPool_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ManagedClusterAgentPoolProfileProperties_ARMGenerator())
}

func Test_ManagedClusterAgentPoolProfileProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedClusterAgentPoolProfileProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedClusterAgentPoolProfileProperties_ARM, ManagedClusterAgentPoolProfileProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedClusterAgentPoolProfileProperties_ARM runs a test to see if a specific instance of ManagedClusterAgentPoolProfileProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedClusterAgentPoolProfileProperties_ARM(subject ManagedClusterAgentPoolProfileProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedClusterAgentPoolProfileProperties_ARM
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

// Generator of ManagedClusterAgentPoolProfileProperties_ARM instances for property testing - lazily instantiated by
// ManagedClusterAgentPoolProfileProperties_ARMGenerator()
var managedClusterAgentPoolProfileProperties_ARMGenerator gopter.Gen

// ManagedClusterAgentPoolProfileProperties_ARMGenerator returns a generator of ManagedClusterAgentPoolProfileProperties_ARM instances for property testing.
// We first initialize managedClusterAgentPoolProfileProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedClusterAgentPoolProfileProperties_ARMGenerator() gopter.Gen {
	if managedClusterAgentPoolProfileProperties_ARMGenerator != nil {
		return managedClusterAgentPoolProfileProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_ARM(generators)
	managedClusterAgentPoolProfileProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ManagedClusterAgentPoolProfileProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_ARM(generators)
	managedClusterAgentPoolProfileProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ManagedClusterAgentPoolProfileProperties_ARM{}), generators)

	return managedClusterAgentPoolProfileProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_ARM(gens map[string]gopter.Gen) {
	gens["AvailabilityZones"] = gen.SliceOf(gen.AlphaString())
	gens["CapacityReservationGroupID"] = gen.PtrOf(gen.AlphaString())
	gens["Count"] = gen.PtrOf(gen.Int())
	gens["EnableAutoScaling"] = gen.PtrOf(gen.Bool())
	gens["EnableCustomCATrust"] = gen.PtrOf(gen.Bool())
	gens["EnableEncryptionAtHost"] = gen.PtrOf(gen.Bool())
	gens["EnableFIPS"] = gen.PtrOf(gen.Bool())
	gens["EnableNodePublicIP"] = gen.PtrOf(gen.Bool())
	gens["EnableUltraSSD"] = gen.PtrOf(gen.Bool())
	gens["GpuInstanceProfile"] = gen.PtrOf(gen.OneConstOf(
		GPUInstanceProfile_MIG1G,
		GPUInstanceProfile_MIG2G,
		GPUInstanceProfile_MIG3G,
		GPUInstanceProfile_MIG4G,
		GPUInstanceProfile_MIG7G))
	gens["HostGroupID"] = gen.PtrOf(gen.AlphaString())
	gens["KubeletDiskType"] = gen.PtrOf(gen.OneConstOf(KubeletDiskType_OS, KubeletDiskType_Temporary))
	gens["MaxCount"] = gen.PtrOf(gen.Int())
	gens["MaxPods"] = gen.PtrOf(gen.Int())
	gens["MessageOfTheDay"] = gen.PtrOf(gen.AlphaString())
	gens["MinCount"] = gen.PtrOf(gen.Int())
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(AgentPoolMode_System, AgentPoolMode_User))
	gens["NodeLabels"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["NodePublicIPPrefixID"] = gen.PtrOf(gen.AlphaString())
	gens["NodeTaints"] = gen.SliceOf(gen.AlphaString())
	gens["OrchestratorVersion"] = gen.PtrOf(gen.AlphaString())
	gens["OsDiskSizeGB"] = gen.PtrOf(gen.Int().Map(func(it int) ContainerServiceOSDisk {
		return ContainerServiceOSDisk(it)
	}))
	gens["OsDiskType"] = gen.PtrOf(gen.OneConstOf(OSDiskType_Ephemeral, OSDiskType_Managed))
	gens["OsSKU"] = gen.PtrOf(gen.OneConstOf(
		OSSKU_CBLMariner,
		OSSKU_Mariner,
		OSSKU_Ubuntu,
		OSSKU_Windows2019,
		OSSKU_Windows2022))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(OSType_Linux, OSType_Windows))
	gens["PodSubnetID"] = gen.PtrOf(gen.AlphaString())
	gens["ProximityPlacementGroupID"] = gen.PtrOf(gen.AlphaString())
	gens["ScaleDownMode"] = gen.PtrOf(gen.OneConstOf(ScaleDownMode_Deallocate, ScaleDownMode_Delete))
	gens["ScaleSetEvictionPolicy"] = gen.PtrOf(gen.OneConstOf(ScaleSetEvictionPolicy_Deallocate, ScaleSetEvictionPolicy_Delete))
	gens["ScaleSetPriority"] = gen.PtrOf(gen.OneConstOf(ScaleSetPriority_Regular, ScaleSetPriority_Spot))
	gens["SpotMaxPrice"] = gen.PtrOf(gen.Float64())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(AgentPoolType_AvailabilitySet, AgentPoolType_VirtualMachineScaleSets))
	gens["VmSize"] = gen.PtrOf(gen.AlphaString())
	gens["VnetSubnetID"] = gen.PtrOf(gen.AlphaString())
	gens["WorkloadRuntime"] = gen.PtrOf(gen.OneConstOf(WorkloadRuntime_KataMshvVmIsolation, WorkloadRuntime_OCIContainer, WorkloadRuntime_WasmWasi))
}

// AddRelatedPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_ARM(gens map[string]gopter.Gen) {
	gens["CreationData"] = gen.PtrOf(CreationData_ARMGenerator())
	gens["KubeletConfig"] = gen.PtrOf(KubeletConfig_ARMGenerator())
	gens["LinuxOSConfig"] = gen.PtrOf(LinuxOSConfig_ARMGenerator())
	gens["NetworkProfile"] = gen.PtrOf(AgentPoolNetworkProfile_ARMGenerator())
	gens["PowerState"] = gen.PtrOf(PowerState_ARMGenerator())
	gens["UpgradeSettings"] = gen.PtrOf(AgentPoolUpgradeSettings_ARMGenerator())
	gens["WindowsProfile"] = gen.PtrOf(AgentPoolWindowsProfile_ARMGenerator())
}

func Test_AgentPoolNetworkProfile_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AgentPoolNetworkProfile_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPoolNetworkProfile_ARM, AgentPoolNetworkProfile_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPoolNetworkProfile_ARM runs a test to see if a specific instance of AgentPoolNetworkProfile_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPoolNetworkProfile_ARM(subject AgentPoolNetworkProfile_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AgentPoolNetworkProfile_ARM
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

// Generator of AgentPoolNetworkProfile_ARM instances for property testing - lazily instantiated by
// AgentPoolNetworkProfile_ARMGenerator()
var agentPoolNetworkProfile_ARMGenerator gopter.Gen

// AgentPoolNetworkProfile_ARMGenerator returns a generator of AgentPoolNetworkProfile_ARM instances for property testing.
// We first initialize agentPoolNetworkProfile_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AgentPoolNetworkProfile_ARMGenerator() gopter.Gen {
	if agentPoolNetworkProfile_ARMGenerator != nil {
		return agentPoolNetworkProfile_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolNetworkProfile_ARM(generators)
	agentPoolNetworkProfile_ARMGenerator = gen.Struct(reflect.TypeOf(AgentPoolNetworkProfile_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolNetworkProfile_ARM(generators)
	AddRelatedPropertyGeneratorsForAgentPoolNetworkProfile_ARM(generators)
	agentPoolNetworkProfile_ARMGenerator = gen.Struct(reflect.TypeOf(AgentPoolNetworkProfile_ARM{}), generators)

	return agentPoolNetworkProfile_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAgentPoolNetworkProfile_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPoolNetworkProfile_ARM(gens map[string]gopter.Gen) {
	gens["ApplicationSecurityGroups"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAgentPoolNetworkProfile_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAgentPoolNetworkProfile_ARM(gens map[string]gopter.Gen) {
	gens["AllowedHostPorts"] = gen.SliceOf(PortRange_ARMGenerator())
	gens["NodePublicIPTags"] = gen.SliceOf(IPTag_ARMGenerator())
}

func Test_AgentPoolUpgradeSettings_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AgentPoolUpgradeSettings_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPoolUpgradeSettings_ARM, AgentPoolUpgradeSettings_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPoolUpgradeSettings_ARM runs a test to see if a specific instance of AgentPoolUpgradeSettings_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPoolUpgradeSettings_ARM(subject AgentPoolUpgradeSettings_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AgentPoolUpgradeSettings_ARM
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

// Generator of AgentPoolUpgradeSettings_ARM instances for property testing - lazily instantiated by
// AgentPoolUpgradeSettings_ARMGenerator()
var agentPoolUpgradeSettings_ARMGenerator gopter.Gen

// AgentPoolUpgradeSettings_ARMGenerator returns a generator of AgentPoolUpgradeSettings_ARM instances for property testing.
func AgentPoolUpgradeSettings_ARMGenerator() gopter.Gen {
	if agentPoolUpgradeSettings_ARMGenerator != nil {
		return agentPoolUpgradeSettings_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings_ARM(generators)
	agentPoolUpgradeSettings_ARMGenerator = gen.Struct(reflect.TypeOf(AgentPoolUpgradeSettings_ARM{}), generators)

	return agentPoolUpgradeSettings_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings_ARM(gens map[string]gopter.Gen) {
	gens["MaxSurge"] = gen.PtrOf(gen.AlphaString())
}

func Test_AgentPoolWindowsProfile_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AgentPoolWindowsProfile_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPoolWindowsProfile_ARM, AgentPoolWindowsProfile_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPoolWindowsProfile_ARM runs a test to see if a specific instance of AgentPoolWindowsProfile_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPoolWindowsProfile_ARM(subject AgentPoolWindowsProfile_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AgentPoolWindowsProfile_ARM
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

// Generator of AgentPoolWindowsProfile_ARM instances for property testing - lazily instantiated by
// AgentPoolWindowsProfile_ARMGenerator()
var agentPoolWindowsProfile_ARMGenerator gopter.Gen

// AgentPoolWindowsProfile_ARMGenerator returns a generator of AgentPoolWindowsProfile_ARM instances for property testing.
func AgentPoolWindowsProfile_ARMGenerator() gopter.Gen {
	if agentPoolWindowsProfile_ARMGenerator != nil {
		return agentPoolWindowsProfile_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolWindowsProfile_ARM(generators)
	agentPoolWindowsProfile_ARMGenerator = gen.Struct(reflect.TypeOf(AgentPoolWindowsProfile_ARM{}), generators)

	return agentPoolWindowsProfile_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAgentPoolWindowsProfile_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPoolWindowsProfile_ARM(gens map[string]gopter.Gen) {
	gens["DisableOutboundNat"] = gen.PtrOf(gen.Bool())
}

func Test_KubeletConfig_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KubeletConfig_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKubeletConfig_ARM, KubeletConfig_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKubeletConfig_ARM runs a test to see if a specific instance of KubeletConfig_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKubeletConfig_ARM(subject KubeletConfig_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KubeletConfig_ARM
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

// Generator of KubeletConfig_ARM instances for property testing - lazily instantiated by KubeletConfig_ARMGenerator()
var kubeletConfig_ARMGenerator gopter.Gen

// KubeletConfig_ARMGenerator returns a generator of KubeletConfig_ARM instances for property testing.
func KubeletConfig_ARMGenerator() gopter.Gen {
	if kubeletConfig_ARMGenerator != nil {
		return kubeletConfig_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKubeletConfig_ARM(generators)
	kubeletConfig_ARMGenerator = gen.Struct(reflect.TypeOf(KubeletConfig_ARM{}), generators)

	return kubeletConfig_ARMGenerator
}

// AddIndependentPropertyGeneratorsForKubeletConfig_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKubeletConfig_ARM(gens map[string]gopter.Gen) {
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

func Test_LinuxOSConfig_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LinuxOSConfig_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLinuxOSConfig_ARM, LinuxOSConfig_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLinuxOSConfig_ARM runs a test to see if a specific instance of LinuxOSConfig_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLinuxOSConfig_ARM(subject LinuxOSConfig_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LinuxOSConfig_ARM
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

// Generator of LinuxOSConfig_ARM instances for property testing - lazily instantiated by LinuxOSConfig_ARMGenerator()
var linuxOSConfig_ARMGenerator gopter.Gen

// LinuxOSConfig_ARMGenerator returns a generator of LinuxOSConfig_ARM instances for property testing.
// We first initialize linuxOSConfig_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LinuxOSConfig_ARMGenerator() gopter.Gen {
	if linuxOSConfig_ARMGenerator != nil {
		return linuxOSConfig_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfig_ARM(generators)
	linuxOSConfig_ARMGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfig_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfig_ARM(generators)
	AddRelatedPropertyGeneratorsForLinuxOSConfig_ARM(generators)
	linuxOSConfig_ARMGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfig_ARM{}), generators)

	return linuxOSConfig_ARMGenerator
}

// AddIndependentPropertyGeneratorsForLinuxOSConfig_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLinuxOSConfig_ARM(gens map[string]gopter.Gen) {
	gens["SwapFileSizeMB"] = gen.PtrOf(gen.Int())
	gens["TransparentHugePageDefrag"] = gen.PtrOf(gen.AlphaString())
	gens["TransparentHugePageEnabled"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForLinuxOSConfig_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLinuxOSConfig_ARM(gens map[string]gopter.Gen) {
	gens["Sysctls"] = gen.PtrOf(SysctlConfig_ARMGenerator())
}

func Test_PowerState_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PowerState_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPowerState_ARM, PowerState_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPowerState_ARM runs a test to see if a specific instance of PowerState_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPowerState_ARM(subject PowerState_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PowerState_ARM
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

// Generator of PowerState_ARM instances for property testing - lazily instantiated by PowerState_ARMGenerator()
var powerState_ARMGenerator gopter.Gen

// PowerState_ARMGenerator returns a generator of PowerState_ARM instances for property testing.
func PowerState_ARMGenerator() gopter.Gen {
	if powerState_ARMGenerator != nil {
		return powerState_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPowerState_ARM(generators)
	powerState_ARMGenerator = gen.Struct(reflect.TypeOf(PowerState_ARM{}), generators)

	return powerState_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPowerState_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPowerState_ARM(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.OneConstOf(PowerState_Code_Running, PowerState_Code_Stopped))
}

func Test_IPTag_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IPTag_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIPTag_ARM, IPTag_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIPTag_ARM runs a test to see if a specific instance of IPTag_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIPTag_ARM(subject IPTag_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IPTag_ARM
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

// Generator of IPTag_ARM instances for property testing - lazily instantiated by IPTag_ARMGenerator()
var ipTag_ARMGenerator gopter.Gen

// IPTag_ARMGenerator returns a generator of IPTag_ARM instances for property testing.
func IPTag_ARMGenerator() gopter.Gen {
	if ipTag_ARMGenerator != nil {
		return ipTag_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIPTag_ARM(generators)
	ipTag_ARMGenerator = gen.Struct(reflect.TypeOf(IPTag_ARM{}), generators)

	return ipTag_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIPTag_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIPTag_ARM(gens map[string]gopter.Gen) {
	gens["IpTagType"] = gen.PtrOf(gen.AlphaString())
	gens["Tag"] = gen.PtrOf(gen.AlphaString())
}

func Test_PortRange_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PortRange_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPortRange_ARM, PortRange_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPortRange_ARM runs a test to see if a specific instance of PortRange_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPortRange_ARM(subject PortRange_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PortRange_ARM
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

// Generator of PortRange_ARM instances for property testing - lazily instantiated by PortRange_ARMGenerator()
var portRange_ARMGenerator gopter.Gen

// PortRange_ARMGenerator returns a generator of PortRange_ARM instances for property testing.
func PortRange_ARMGenerator() gopter.Gen {
	if portRange_ARMGenerator != nil {
		return portRange_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPortRange_ARM(generators)
	portRange_ARMGenerator = gen.Struct(reflect.TypeOf(PortRange_ARM{}), generators)

	return portRange_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPortRange_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPortRange_ARM(gens map[string]gopter.Gen) {
	gens["PortEnd"] = gen.PtrOf(gen.Int())
	gens["PortStart"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(PortRange_Protocol_TCP, PortRange_Protocol_UDP))
}

func Test_SysctlConfig_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SysctlConfig_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSysctlConfig_ARM, SysctlConfig_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSysctlConfig_ARM runs a test to see if a specific instance of SysctlConfig_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSysctlConfig_ARM(subject SysctlConfig_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SysctlConfig_ARM
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

// Generator of SysctlConfig_ARM instances for property testing - lazily instantiated by SysctlConfig_ARMGenerator()
var sysctlConfig_ARMGenerator gopter.Gen

// SysctlConfig_ARMGenerator returns a generator of SysctlConfig_ARM instances for property testing.
func SysctlConfig_ARMGenerator() gopter.Gen {
	if sysctlConfig_ARMGenerator != nil {
		return sysctlConfig_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSysctlConfig_ARM(generators)
	sysctlConfig_ARMGenerator = gen.Struct(reflect.TypeOf(SysctlConfig_ARM{}), generators)

	return sysctlConfig_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSysctlConfig_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSysctlConfig_ARM(gens map[string]gopter.Gen) {
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

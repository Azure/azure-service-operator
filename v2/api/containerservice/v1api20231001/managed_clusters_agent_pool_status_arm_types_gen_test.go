// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20231001

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

func Test_AgentPoolNetworkProfile_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AgentPoolNetworkProfile_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPoolNetworkProfile_STATUS_ARM, AgentPoolNetworkProfile_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPoolNetworkProfile_STATUS_ARM runs a test to see if a specific instance of AgentPoolNetworkProfile_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPoolNetworkProfile_STATUS_ARM(subject AgentPoolNetworkProfile_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AgentPoolNetworkProfile_STATUS_ARM
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

// Generator of AgentPoolNetworkProfile_STATUS_ARM instances for property testing - lazily instantiated by
// AgentPoolNetworkProfile_STATUS_ARMGenerator()
var agentPoolNetworkProfile_STATUS_ARMGenerator gopter.Gen

// AgentPoolNetworkProfile_STATUS_ARMGenerator returns a generator of AgentPoolNetworkProfile_STATUS_ARM instances for property testing.
// We first initialize agentPoolNetworkProfile_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AgentPoolNetworkProfile_STATUS_ARMGenerator() gopter.Gen {
	if agentPoolNetworkProfile_STATUS_ARMGenerator != nil {
		return agentPoolNetworkProfile_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolNetworkProfile_STATUS_ARM(generators)
	agentPoolNetworkProfile_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AgentPoolNetworkProfile_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolNetworkProfile_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForAgentPoolNetworkProfile_STATUS_ARM(generators)
	agentPoolNetworkProfile_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AgentPoolNetworkProfile_STATUS_ARM{}), generators)

	return agentPoolNetworkProfile_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAgentPoolNetworkProfile_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPoolNetworkProfile_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ApplicationSecurityGroups"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAgentPoolNetworkProfile_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAgentPoolNetworkProfile_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AllowedHostPorts"] = gen.SliceOf(PortRange_STATUS_ARMGenerator())
	gens["NodePublicIPTags"] = gen.SliceOf(IPTag_STATUS_ARMGenerator())
}

func Test_AgentPoolUpgradeSettings_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AgentPoolUpgradeSettings_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPoolUpgradeSettings_STATUS_ARM, AgentPoolUpgradeSettings_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPoolUpgradeSettings_STATUS_ARM runs a test to see if a specific instance of AgentPoolUpgradeSettings_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPoolUpgradeSettings_STATUS_ARM(subject AgentPoolUpgradeSettings_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AgentPoolUpgradeSettings_STATUS_ARM
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

// Generator of AgentPoolUpgradeSettings_STATUS_ARM instances for property testing - lazily instantiated by
// AgentPoolUpgradeSettings_STATUS_ARMGenerator()
var agentPoolUpgradeSettings_STATUS_ARMGenerator gopter.Gen

// AgentPoolUpgradeSettings_STATUS_ARMGenerator returns a generator of AgentPoolUpgradeSettings_STATUS_ARM instances for property testing.
func AgentPoolUpgradeSettings_STATUS_ARMGenerator() gopter.Gen {
	if agentPoolUpgradeSettings_STATUS_ARMGenerator != nil {
		return agentPoolUpgradeSettings_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings_STATUS_ARM(generators)
	agentPoolUpgradeSettings_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AgentPoolUpgradeSettings_STATUS_ARM{}), generators)

	return agentPoolUpgradeSettings_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DrainTimeoutInMinutes"] = gen.PtrOf(gen.Int())
	gens["MaxSurge"] = gen.PtrOf(gen.AlphaString())
}

func Test_CreationData_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CreationData_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCreationData_STATUS_ARM, CreationData_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCreationData_STATUS_ARM runs a test to see if a specific instance of CreationData_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCreationData_STATUS_ARM(subject CreationData_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CreationData_STATUS_ARM
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

// Generator of CreationData_STATUS_ARM instances for property testing - lazily instantiated by
// CreationData_STATUS_ARMGenerator()
var creationData_STATUS_ARMGenerator gopter.Gen

// CreationData_STATUS_ARMGenerator returns a generator of CreationData_STATUS_ARM instances for property testing.
func CreationData_STATUS_ARMGenerator() gopter.Gen {
	if creationData_STATUS_ARMGenerator != nil {
		return creationData_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreationData_STATUS_ARM(generators)
	creationData_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(CreationData_STATUS_ARM{}), generators)

	return creationData_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCreationData_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCreationData_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["SourceResourceId"] = gen.PtrOf(gen.AlphaString())
}

func Test_IPTag_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IPTag_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIPTag_STATUS_ARM, IPTag_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIPTag_STATUS_ARM runs a test to see if a specific instance of IPTag_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIPTag_STATUS_ARM(subject IPTag_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IPTag_STATUS_ARM
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

// Generator of IPTag_STATUS_ARM instances for property testing - lazily instantiated by IPTag_STATUS_ARMGenerator()
var ipTag_STATUS_ARMGenerator gopter.Gen

// IPTag_STATUS_ARMGenerator returns a generator of IPTag_STATUS_ARM instances for property testing.
func IPTag_STATUS_ARMGenerator() gopter.Gen {
	if ipTag_STATUS_ARMGenerator != nil {
		return ipTag_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIPTag_STATUS_ARM(generators)
	ipTag_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(IPTag_STATUS_ARM{}), generators)

	return ipTag_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIPTag_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIPTag_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["IpTagType"] = gen.PtrOf(gen.AlphaString())
	gens["Tag"] = gen.PtrOf(gen.AlphaString())
}

func Test_KubeletConfig_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KubeletConfig_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKubeletConfig_STATUS_ARM, KubeletConfig_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKubeletConfig_STATUS_ARM runs a test to see if a specific instance of KubeletConfig_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKubeletConfig_STATUS_ARM(subject KubeletConfig_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KubeletConfig_STATUS_ARM
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

// Generator of KubeletConfig_STATUS_ARM instances for property testing - lazily instantiated by
// KubeletConfig_STATUS_ARMGenerator()
var kubeletConfig_STATUS_ARMGenerator gopter.Gen

// KubeletConfig_STATUS_ARMGenerator returns a generator of KubeletConfig_STATUS_ARM instances for property testing.
func KubeletConfig_STATUS_ARMGenerator() gopter.Gen {
	if kubeletConfig_STATUS_ARMGenerator != nil {
		return kubeletConfig_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKubeletConfig_STATUS_ARM(generators)
	kubeletConfig_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(KubeletConfig_STATUS_ARM{}), generators)

	return kubeletConfig_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForKubeletConfig_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKubeletConfig_STATUS_ARM(gens map[string]gopter.Gen) {
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

func Test_LinuxOSConfig_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LinuxOSConfig_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLinuxOSConfig_STATUS_ARM, LinuxOSConfig_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLinuxOSConfig_STATUS_ARM runs a test to see if a specific instance of LinuxOSConfig_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLinuxOSConfig_STATUS_ARM(subject LinuxOSConfig_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LinuxOSConfig_STATUS_ARM
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

// Generator of LinuxOSConfig_STATUS_ARM instances for property testing - lazily instantiated by
// LinuxOSConfig_STATUS_ARMGenerator()
var linuxOSConfig_STATUS_ARMGenerator gopter.Gen

// LinuxOSConfig_STATUS_ARMGenerator returns a generator of LinuxOSConfig_STATUS_ARM instances for property testing.
// We first initialize linuxOSConfig_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LinuxOSConfig_STATUS_ARMGenerator() gopter.Gen {
	if linuxOSConfig_STATUS_ARMGenerator != nil {
		return linuxOSConfig_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfig_STATUS_ARM(generators)
	linuxOSConfig_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfig_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfig_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForLinuxOSConfig_STATUS_ARM(generators)
	linuxOSConfig_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfig_STATUS_ARM{}), generators)

	return linuxOSConfig_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForLinuxOSConfig_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLinuxOSConfig_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["SwapFileSizeMB"] = gen.PtrOf(gen.Int())
	gens["TransparentHugePageDefrag"] = gen.PtrOf(gen.AlphaString())
	gens["TransparentHugePageEnabled"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForLinuxOSConfig_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLinuxOSConfig_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Sysctls"] = gen.PtrOf(SysctlConfig_STATUS_ARMGenerator())
}

func Test_ManagedClusterAgentPoolProfileProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedClusterAgentPoolProfileProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedClusterAgentPoolProfileProperties_STATUS_ARM, ManagedClusterAgentPoolProfileProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedClusterAgentPoolProfileProperties_STATUS_ARM runs a test to see if a specific instance of ManagedClusterAgentPoolProfileProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedClusterAgentPoolProfileProperties_STATUS_ARM(subject ManagedClusterAgentPoolProfileProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedClusterAgentPoolProfileProperties_STATUS_ARM
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

// Generator of ManagedClusterAgentPoolProfileProperties_STATUS_ARM instances for property testing - lazily instantiated
// by ManagedClusterAgentPoolProfileProperties_STATUS_ARMGenerator()
var managedClusterAgentPoolProfileProperties_STATUS_ARMGenerator gopter.Gen

// ManagedClusterAgentPoolProfileProperties_STATUS_ARMGenerator returns a generator of ManagedClusterAgentPoolProfileProperties_STATUS_ARM instances for property testing.
// We first initialize managedClusterAgentPoolProfileProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedClusterAgentPoolProfileProperties_STATUS_ARMGenerator() gopter.Gen {
	if managedClusterAgentPoolProfileProperties_STATUS_ARMGenerator != nil {
		return managedClusterAgentPoolProfileProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_STATUS_ARM(generators)
	managedClusterAgentPoolProfileProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagedClusterAgentPoolProfileProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_STATUS_ARM(generators)
	managedClusterAgentPoolProfileProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagedClusterAgentPoolProfileProperties_STATUS_ARM{}), generators)

	return managedClusterAgentPoolProfileProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AvailabilityZones"] = gen.SliceOf(gen.AlphaString())
	gens["CapacityReservationGroupID"] = gen.PtrOf(gen.AlphaString())
	gens["Count"] = gen.PtrOf(gen.Int())
	gens["CurrentOrchestratorVersion"] = gen.PtrOf(gen.AlphaString())
	gens["EnableAutoScaling"] = gen.PtrOf(gen.Bool())
	gens["EnableEncryptionAtHost"] = gen.PtrOf(gen.Bool())
	gens["EnableFIPS"] = gen.PtrOf(gen.Bool())
	gens["EnableNodePublicIP"] = gen.PtrOf(gen.Bool())
	gens["EnableUltraSSD"] = gen.PtrOf(gen.Bool())
	gens["GpuInstanceProfile"] = gen.PtrOf(gen.OneConstOf(
		GPUInstanceProfile_STATUS_MIG1G,
		GPUInstanceProfile_STATUS_MIG2G,
		GPUInstanceProfile_STATUS_MIG3G,
		GPUInstanceProfile_STATUS_MIG4G,
		GPUInstanceProfile_STATUS_MIG7G))
	gens["HostGroupID"] = gen.PtrOf(gen.AlphaString())
	gens["KubeletDiskType"] = gen.PtrOf(gen.OneConstOf(KubeletDiskType_STATUS_OS, KubeletDiskType_STATUS_Temporary))
	gens["MaxCount"] = gen.PtrOf(gen.Int())
	gens["MaxPods"] = gen.PtrOf(gen.Int())
	gens["MinCount"] = gen.PtrOf(gen.Int())
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(AgentPoolMode_STATUS_System, AgentPoolMode_STATUS_User))
	gens["NodeImageVersion"] = gen.PtrOf(gen.AlphaString())
	gens["NodeLabels"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["NodePublicIPPrefixID"] = gen.PtrOf(gen.AlphaString())
	gens["NodeTaints"] = gen.SliceOf(gen.AlphaString())
	gens["OrchestratorVersion"] = gen.PtrOf(gen.AlphaString())
	gens["OsDiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["OsDiskType"] = gen.PtrOf(gen.OneConstOf(OSDiskType_STATUS_Ephemeral, OSDiskType_STATUS_Managed))
	gens["OsSKU"] = gen.PtrOf(gen.OneConstOf(
		OSSKU_STATUS_AzureLinux,
		OSSKU_STATUS_CBLMariner,
		OSSKU_STATUS_Ubuntu,
		OSSKU_STATUS_Windows2019,
		OSSKU_STATUS_Windows2022))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(OSType_STATUS_Linux, OSType_STATUS_Windows))
	gens["PodSubnetID"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ProximityPlacementGroupID"] = gen.PtrOf(gen.AlphaString())
	gens["ScaleDownMode"] = gen.PtrOf(gen.OneConstOf(ScaleDownMode_STATUS_Deallocate, ScaleDownMode_STATUS_Delete))
	gens["ScaleSetEvictionPolicy"] = gen.PtrOf(gen.OneConstOf(ScaleSetEvictionPolicy_STATUS_Deallocate, ScaleSetEvictionPolicy_STATUS_Delete))
	gens["ScaleSetPriority"] = gen.PtrOf(gen.OneConstOf(ScaleSetPriority_STATUS_Regular, ScaleSetPriority_STATUS_Spot))
	gens["SpotMaxPrice"] = gen.PtrOf(gen.Float64())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(AgentPoolType_STATUS_AvailabilitySet, AgentPoolType_STATUS_VirtualMachineScaleSets))
	gens["VmSize"] = gen.PtrOf(gen.AlphaString())
	gens["VnetSubnetID"] = gen.PtrOf(gen.AlphaString())
	gens["WorkloadRuntime"] = gen.PtrOf(gen.OneConstOf(WorkloadRuntime_STATUS_OCIContainer, WorkloadRuntime_STATUS_WasmWasi))
}

// AddRelatedPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CreationData"] = gen.PtrOf(CreationData_STATUS_ARMGenerator())
	gens["KubeletConfig"] = gen.PtrOf(KubeletConfig_STATUS_ARMGenerator())
	gens["LinuxOSConfig"] = gen.PtrOf(LinuxOSConfig_STATUS_ARMGenerator())
	gens["NetworkProfile"] = gen.PtrOf(AgentPoolNetworkProfile_STATUS_ARMGenerator())
	gens["PowerState"] = gen.PtrOf(PowerState_STATUS_ARMGenerator())
	gens["UpgradeSettings"] = gen.PtrOf(AgentPoolUpgradeSettings_STATUS_ARMGenerator())
}

func Test_ManagedClusters_AgentPool_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedClusters_AgentPool_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedClusters_AgentPool_STATUS_ARM, ManagedClusters_AgentPool_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedClusters_AgentPool_STATUS_ARM runs a test to see if a specific instance of ManagedClusters_AgentPool_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedClusters_AgentPool_STATUS_ARM(subject ManagedClusters_AgentPool_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedClusters_AgentPool_STATUS_ARM
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

// Generator of ManagedClusters_AgentPool_STATUS_ARM instances for property testing - lazily instantiated by
// ManagedClusters_AgentPool_STATUS_ARMGenerator()
var managedClusters_AgentPool_STATUS_ARMGenerator gopter.Gen

// ManagedClusters_AgentPool_STATUS_ARMGenerator returns a generator of ManagedClusters_AgentPool_STATUS_ARM instances for property testing.
// We first initialize managedClusters_AgentPool_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedClusters_AgentPool_STATUS_ARMGenerator() gopter.Gen {
	if managedClusters_AgentPool_STATUS_ARMGenerator != nil {
		return managedClusters_AgentPool_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusters_AgentPool_STATUS_ARM(generators)
	managedClusters_AgentPool_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagedClusters_AgentPool_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusters_AgentPool_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForManagedClusters_AgentPool_STATUS_ARM(generators)
	managedClusters_AgentPool_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagedClusters_AgentPool_STATUS_ARM{}), generators)

	return managedClusters_AgentPool_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForManagedClusters_AgentPool_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedClusters_AgentPool_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForManagedClusters_AgentPool_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedClusters_AgentPool_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ManagedClusterAgentPoolProfileProperties_STATUS_ARMGenerator())
}

func Test_PortRange_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PortRange_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPortRange_STATUS_ARM, PortRange_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPortRange_STATUS_ARM runs a test to see if a specific instance of PortRange_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPortRange_STATUS_ARM(subject PortRange_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PortRange_STATUS_ARM
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

// Generator of PortRange_STATUS_ARM instances for property testing - lazily instantiated by
// PortRange_STATUS_ARMGenerator()
var portRange_STATUS_ARMGenerator gopter.Gen

// PortRange_STATUS_ARMGenerator returns a generator of PortRange_STATUS_ARM instances for property testing.
func PortRange_STATUS_ARMGenerator() gopter.Gen {
	if portRange_STATUS_ARMGenerator != nil {
		return portRange_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPortRange_STATUS_ARM(generators)
	portRange_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(PortRange_STATUS_ARM{}), generators)

	return portRange_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPortRange_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPortRange_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["PortEnd"] = gen.PtrOf(gen.Int())
	gens["PortStart"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(PortRange_Protocol_STATUS_TCP, PortRange_Protocol_STATUS_UDP))
}

func Test_SysctlConfig_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SysctlConfig_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSysctlConfig_STATUS_ARM, SysctlConfig_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSysctlConfig_STATUS_ARM runs a test to see if a specific instance of SysctlConfig_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSysctlConfig_STATUS_ARM(subject SysctlConfig_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SysctlConfig_STATUS_ARM
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

// Generator of SysctlConfig_STATUS_ARM instances for property testing - lazily instantiated by
// SysctlConfig_STATUS_ARMGenerator()
var sysctlConfig_STATUS_ARMGenerator gopter.Gen

// SysctlConfig_STATUS_ARMGenerator returns a generator of SysctlConfig_STATUS_ARM instances for property testing.
func SysctlConfig_STATUS_ARMGenerator() gopter.Gen {
	if sysctlConfig_STATUS_ARMGenerator != nil {
		return sysctlConfig_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSysctlConfig_STATUS_ARM(generators)
	sysctlConfig_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SysctlConfig_STATUS_ARM{}), generators)

	return sysctlConfig_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSysctlConfig_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSysctlConfig_STATUS_ARM(gens map[string]gopter.Gen) {
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

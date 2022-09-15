// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import (
	"encoding/json"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101storage"
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

func Test_NetworkSecurityGroup_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroup to hub returns original",
		prop.ForAll(RunResourceConversionTestForNetworkSecurityGroup, NetworkSecurityGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNetworkSecurityGroup tests if a specific instance of NetworkSecurityGroup round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNetworkSecurityGroup(subject NetworkSecurityGroup) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20201101s.NetworkSecurityGroup
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NetworkSecurityGroup
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NetworkSecurityGroup_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroup to NetworkSecurityGroup via AssignProperties_To_NetworkSecurityGroup & AssignProperties_From_NetworkSecurityGroup returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkSecurityGroup, NetworkSecurityGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkSecurityGroup tests if a specific instance of NetworkSecurityGroup can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkSecurityGroup(subject NetworkSecurityGroup) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.NetworkSecurityGroup
	err := copied.AssignProperties_To_NetworkSecurityGroup(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkSecurityGroup
	err = actual.AssignProperties_From_NetworkSecurityGroup(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NetworkSecurityGroup_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroup via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroup, NetworkSecurityGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroup runs a test to see if a specific instance of NetworkSecurityGroup round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroup(subject NetworkSecurityGroup) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroup
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

// Generator of NetworkSecurityGroup instances for property testing - lazily instantiated by
// NetworkSecurityGroupGenerator()
var networkSecurityGroupGenerator gopter.Gen

// NetworkSecurityGroupGenerator returns a generator of NetworkSecurityGroup instances for property testing.
func NetworkSecurityGroupGenerator() gopter.Gen {
	if networkSecurityGroupGenerator != nil {
		return networkSecurityGroupGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroup(generators)
	networkSecurityGroupGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup{}), generators)

	return networkSecurityGroupGenerator
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroup is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroup(gens map[string]gopter.Gen) {
	gens["Spec"] = NetworkSecurityGroup_SpecGenerator()
	gens["Status"] = NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()
}

func Test_NetworkSecurityGroup_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroup_Spec to NetworkSecurityGroup_Spec via AssignProperties_To_NetworkSecurityGroup_Spec & AssignProperties_From_NetworkSecurityGroup_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkSecurityGroup_Spec, NetworkSecurityGroup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkSecurityGroup_Spec tests if a specific instance of NetworkSecurityGroup_Spec can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkSecurityGroup_Spec(subject NetworkSecurityGroup_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.NetworkSecurityGroup_Spec
	err := copied.AssignProperties_To_NetworkSecurityGroup_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkSecurityGroup_Spec
	err = actual.AssignProperties_From_NetworkSecurityGroup_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NetworkSecurityGroup_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroup_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroup_Spec, NetworkSecurityGroup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroup_Spec runs a test to see if a specific instance of NetworkSecurityGroup_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroup_Spec(subject NetworkSecurityGroup_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroup_Spec
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

// Generator of NetworkSecurityGroup_Spec instances for property testing - lazily instantiated by
// NetworkSecurityGroup_SpecGenerator()
var networkSecurityGroup_SpecGenerator gopter.Gen

// NetworkSecurityGroup_SpecGenerator returns a generator of NetworkSecurityGroup_Spec instances for property testing.
func NetworkSecurityGroup_SpecGenerator() gopter.Gen {
	if networkSecurityGroup_SpecGenerator != nil {
		return networkSecurityGroup_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroup_Spec(generators)
	networkSecurityGroup_SpecGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_Spec{}), generators)

	return networkSecurityGroup_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroup_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroup_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded to NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded via AssignProperties_To_NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded & AssignProperties_From_NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded, NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded tests if a specific instance of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded(subject NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded
	err := copied.AssignProperties_To_NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded
	err = actual.AssignProperties_From_NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded, NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded runs a test to see if a specific instance of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded(subject NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded
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

// Generator of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded instances for property testing -
// lazily instantiated by NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()
var networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator gopter.Gen

// NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator returns a generator of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded instances for property testing.
// We first initialize networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator() gopter.Gen {
	if networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator != nil {
		return networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded(generators)
	networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded(generators)
	networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded{}), generators)

	return networkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["DefaultSecurityRules"] = gen.SliceOf(SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator())
	gens["FlowLogs"] = gen.SliceOf(FlowLog_STATUS_SubResourceEmbeddedGenerator())
	gens["NetworkInterfaces"] = gen.SliceOf(NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator())
	gens["SecurityRules"] = gen.SliceOf(SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator())
	gens["Subnets"] = gen.SliceOf(Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator())
}

func Test_FlowLog_STATUS_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlowLog_STATUS_SubResourceEmbedded to FlowLog_STATUS_SubResourceEmbedded via AssignProperties_To_FlowLog_STATUS_SubResourceEmbedded & AssignProperties_From_FlowLog_STATUS_SubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlowLog_STATUS_SubResourceEmbedded, FlowLog_STATUS_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlowLog_STATUS_SubResourceEmbedded tests if a specific instance of FlowLog_STATUS_SubResourceEmbedded can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForFlowLog_STATUS_SubResourceEmbedded(subject FlowLog_STATUS_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.FlowLog_STATUS_SubResourceEmbedded
	err := copied.AssignProperties_To_FlowLog_STATUS_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlowLog_STATUS_SubResourceEmbedded
	err = actual.AssignProperties_From_FlowLog_STATUS_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_FlowLog_STATUS_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlowLog_STATUS_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlowLog_STATUS_SubResourceEmbedded, FlowLog_STATUS_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlowLog_STATUS_SubResourceEmbedded runs a test to see if a specific instance of FlowLog_STATUS_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForFlowLog_STATUS_SubResourceEmbedded(subject FlowLog_STATUS_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlowLog_STATUS_SubResourceEmbedded
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

// Generator of FlowLog_STATUS_SubResourceEmbedded instances for property testing - lazily instantiated by
// FlowLog_STATUS_SubResourceEmbeddedGenerator()
var flowLog_STATUS_SubResourceEmbeddedGenerator gopter.Gen

// FlowLog_STATUS_SubResourceEmbeddedGenerator returns a generator of FlowLog_STATUS_SubResourceEmbedded instances for property testing.
func FlowLog_STATUS_SubResourceEmbeddedGenerator() gopter.Gen {
	if flowLog_STATUS_SubResourceEmbeddedGenerator != nil {
		return flowLog_STATUS_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlowLog_STATUS_SubResourceEmbedded(generators)
	flowLog_STATUS_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(FlowLog_STATUS_SubResourceEmbedded{}), generators)

	return flowLog_STATUS_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForFlowLog_STATUS_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlowLog_STATUS_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded to NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded via AssignProperties_To_NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded & AssignProperties_From_NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded, NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded tests if a specific instance of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded(subject NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded
	err := copied.AssignProperties_To_NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded
	err = actual.AssignProperties_From_NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded, NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded runs a test to see if a specific instance of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded(subject NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded
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

// Generator of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded instances for property testing - lazily
// instantiated by NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()
var networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator gopter.Gen

// NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator returns a generator of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded instances for property testing.
// We first initialize networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator() gopter.Gen {
	if networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator != nil {
		return networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded(generators)
	networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded(generators)
	AddRelatedPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded(generators)
	networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded{}), generators)

	return networkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_STATUSGenerator())
}

func Test_SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded to SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded via AssignProperties_To_SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded & AssignProperties_From_SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForSecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded, SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded tests if a specific instance of SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForSecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded(subject SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded
	err := copied.AssignProperties_To_SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded
	err = actual.AssignProperties_From_SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded, SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded runs a test to see if a specific instance of SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded(subject SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded
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

// Generator of SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded instances for property testing - lazily
// instantiated by SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()
var securityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator gopter.Gen

// SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator returns a generator of SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded instances for property testing.
func SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator() gopter.Gen {
	if securityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator != nil {
		return securityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded(generators)
	securityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded{}), generators)

	return securityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForSecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded to Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded via AssignProperties_To_Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded & AssignProperties_From_Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded, Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded tests if a specific instance of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded(subject Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded
	err := copied.AssignProperties_To_Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded
	err = actual.AssignProperties_From_Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded, Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded runs a test to see if a specific instance of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded(subject Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded
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

// Generator of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded instances for property testing - lazily
// instantiated by Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator()
var subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator gopter.Gen

// Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator returns a generator of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded instances for property testing.
func Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator() gopter.Gen {
	if subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator != nil {
		return subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded(generators)
	subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded{}), generators)

	return subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubnet_STATUS_NetworkSecurityGroup_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

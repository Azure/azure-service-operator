// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"encoding/json"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101storage"
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

func Test_ServersDatabasesAdvancedThreatProtectionSetting_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersDatabasesAdvancedThreatProtectionSetting to hub returns original",
		prop.ForAll(RunResourceConversionTestForServersDatabasesAdvancedThreatProtectionSetting, ServersDatabasesAdvancedThreatProtectionSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForServersDatabasesAdvancedThreatProtectionSetting tests if a specific instance of ServersDatabasesAdvancedThreatProtectionSetting round trips to the hub storage version and back losslessly
func RunResourceConversionTestForServersDatabasesAdvancedThreatProtectionSetting(subject ServersDatabasesAdvancedThreatProtectionSetting) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20211101s.ServersDatabasesAdvancedThreatProtectionSetting
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual ServersDatabasesAdvancedThreatProtectionSetting
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

func Test_ServersDatabasesAdvancedThreatProtectionSetting_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersDatabasesAdvancedThreatProtectionSetting to ServersDatabasesAdvancedThreatProtectionSetting via AssignProperties_To_ServersDatabasesAdvancedThreatProtectionSetting & AssignProperties_From_ServersDatabasesAdvancedThreatProtectionSetting returns original",
		prop.ForAll(RunPropertyAssignmentTestForServersDatabasesAdvancedThreatProtectionSetting, ServersDatabasesAdvancedThreatProtectionSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServersDatabasesAdvancedThreatProtectionSetting tests if a specific instance of ServersDatabasesAdvancedThreatProtectionSetting can be assigned to v1api20211101storage and back losslessly
func RunPropertyAssignmentTestForServersDatabasesAdvancedThreatProtectionSetting(subject ServersDatabasesAdvancedThreatProtectionSetting) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.ServersDatabasesAdvancedThreatProtectionSetting
	err := copied.AssignProperties_To_ServersDatabasesAdvancedThreatProtectionSetting(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ServersDatabasesAdvancedThreatProtectionSetting
	err = actual.AssignProperties_From_ServersDatabasesAdvancedThreatProtectionSetting(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ServersDatabasesAdvancedThreatProtectionSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersDatabasesAdvancedThreatProtectionSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersDatabasesAdvancedThreatProtectionSetting, ServersDatabasesAdvancedThreatProtectionSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersDatabasesAdvancedThreatProtectionSetting runs a test to see if a specific instance of ServersDatabasesAdvancedThreatProtectionSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForServersDatabasesAdvancedThreatProtectionSetting(subject ServersDatabasesAdvancedThreatProtectionSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersDatabasesAdvancedThreatProtectionSetting
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

// Generator of ServersDatabasesAdvancedThreatProtectionSetting instances for property testing - lazily instantiated by
// ServersDatabasesAdvancedThreatProtectionSettingGenerator()
var serversDatabasesAdvancedThreatProtectionSettingGenerator gopter.Gen

// ServersDatabasesAdvancedThreatProtectionSettingGenerator returns a generator of ServersDatabasesAdvancedThreatProtectionSetting instances for property testing.
func ServersDatabasesAdvancedThreatProtectionSettingGenerator() gopter.Gen {
	if serversDatabasesAdvancedThreatProtectionSettingGenerator != nil {
		return serversDatabasesAdvancedThreatProtectionSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersDatabasesAdvancedThreatProtectionSetting(generators)
	serversDatabasesAdvancedThreatProtectionSettingGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesAdvancedThreatProtectionSetting{}), generators)

	return serversDatabasesAdvancedThreatProtectionSettingGenerator
}

// AddRelatedPropertyGeneratorsForServersDatabasesAdvancedThreatProtectionSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersDatabasesAdvancedThreatProtectionSetting(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_Databases_AdvancedThreatProtectionSetting_SpecGenerator()
	gens["Status"] = Servers_Databases_AdvancedThreatProtectionSetting_STATUSGenerator()
}

func Test_Servers_Databases_AdvancedThreatProtectionSetting_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_Databases_AdvancedThreatProtectionSetting_Spec to Servers_Databases_AdvancedThreatProtectionSetting_Spec via AssignProperties_To_Servers_Databases_AdvancedThreatProtectionSetting_Spec & AssignProperties_From_Servers_Databases_AdvancedThreatProtectionSetting_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_Databases_AdvancedThreatProtectionSetting_Spec, Servers_Databases_AdvancedThreatProtectionSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_Databases_AdvancedThreatProtectionSetting_Spec tests if a specific instance of Servers_Databases_AdvancedThreatProtectionSetting_Spec can be assigned to v1api20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_Databases_AdvancedThreatProtectionSetting_Spec(subject Servers_Databases_AdvancedThreatProtectionSetting_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_Databases_AdvancedThreatProtectionSetting_Spec
	err := copied.AssignProperties_To_Servers_Databases_AdvancedThreatProtectionSetting_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_Databases_AdvancedThreatProtectionSetting_Spec
	err = actual.AssignProperties_From_Servers_Databases_AdvancedThreatProtectionSetting_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Servers_Databases_AdvancedThreatProtectionSetting_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Databases_AdvancedThreatProtectionSetting_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Databases_AdvancedThreatProtectionSetting_Spec, Servers_Databases_AdvancedThreatProtectionSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Databases_AdvancedThreatProtectionSetting_Spec runs a test to see if a specific instance of Servers_Databases_AdvancedThreatProtectionSetting_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Databases_AdvancedThreatProtectionSetting_Spec(subject Servers_Databases_AdvancedThreatProtectionSetting_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Databases_AdvancedThreatProtectionSetting_Spec
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

// Generator of Servers_Databases_AdvancedThreatProtectionSetting_Spec instances for property testing - lazily
// instantiated by Servers_Databases_AdvancedThreatProtectionSetting_SpecGenerator()
var servers_Databases_AdvancedThreatProtectionSetting_SpecGenerator gopter.Gen

// Servers_Databases_AdvancedThreatProtectionSetting_SpecGenerator returns a generator of Servers_Databases_AdvancedThreatProtectionSetting_Spec instances for property testing.
func Servers_Databases_AdvancedThreatProtectionSetting_SpecGenerator() gopter.Gen {
	if servers_Databases_AdvancedThreatProtectionSetting_SpecGenerator != nil {
		return servers_Databases_AdvancedThreatProtectionSetting_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Databases_AdvancedThreatProtectionSetting_Spec(generators)
	servers_Databases_AdvancedThreatProtectionSetting_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_Databases_AdvancedThreatProtectionSetting_Spec{}), generators)

	return servers_Databases_AdvancedThreatProtectionSetting_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_Databases_AdvancedThreatProtectionSetting_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Databases_AdvancedThreatProtectionSetting_Spec(gens map[string]gopter.Gen) {
	gens["State"] = gen.PtrOf(gen.OneConstOf(AdvancedThreatProtectionProperties_State_Disabled, AdvancedThreatProtectionProperties_State_Enabled, AdvancedThreatProtectionProperties_State_New))
}

func Test_Servers_Databases_AdvancedThreatProtectionSetting_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_Databases_AdvancedThreatProtectionSetting_STATUS to Servers_Databases_AdvancedThreatProtectionSetting_STATUS via AssignProperties_To_Servers_Databases_AdvancedThreatProtectionSetting_STATUS & AssignProperties_From_Servers_Databases_AdvancedThreatProtectionSetting_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_Databases_AdvancedThreatProtectionSetting_STATUS, Servers_Databases_AdvancedThreatProtectionSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_Databases_AdvancedThreatProtectionSetting_STATUS tests if a specific instance of Servers_Databases_AdvancedThreatProtectionSetting_STATUS can be assigned to v1api20211101storage and back losslessly
func RunPropertyAssignmentTestForServers_Databases_AdvancedThreatProtectionSetting_STATUS(subject Servers_Databases_AdvancedThreatProtectionSetting_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_Databases_AdvancedThreatProtectionSetting_STATUS
	err := copied.AssignProperties_To_Servers_Databases_AdvancedThreatProtectionSetting_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_Databases_AdvancedThreatProtectionSetting_STATUS
	err = actual.AssignProperties_From_Servers_Databases_AdvancedThreatProtectionSetting_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Servers_Databases_AdvancedThreatProtectionSetting_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Databases_AdvancedThreatProtectionSetting_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Databases_AdvancedThreatProtectionSetting_STATUS, Servers_Databases_AdvancedThreatProtectionSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Databases_AdvancedThreatProtectionSetting_STATUS runs a test to see if a specific instance of Servers_Databases_AdvancedThreatProtectionSetting_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Databases_AdvancedThreatProtectionSetting_STATUS(subject Servers_Databases_AdvancedThreatProtectionSetting_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Databases_AdvancedThreatProtectionSetting_STATUS
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

// Generator of Servers_Databases_AdvancedThreatProtectionSetting_STATUS instances for property testing - lazily
// instantiated by Servers_Databases_AdvancedThreatProtectionSetting_STATUSGenerator()
var servers_Databases_AdvancedThreatProtectionSetting_STATUSGenerator gopter.Gen

// Servers_Databases_AdvancedThreatProtectionSetting_STATUSGenerator returns a generator of Servers_Databases_AdvancedThreatProtectionSetting_STATUS instances for property testing.
// We first initialize servers_Databases_AdvancedThreatProtectionSetting_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Servers_Databases_AdvancedThreatProtectionSetting_STATUSGenerator() gopter.Gen {
	if servers_Databases_AdvancedThreatProtectionSetting_STATUSGenerator != nil {
		return servers_Databases_AdvancedThreatProtectionSetting_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Databases_AdvancedThreatProtectionSetting_STATUS(generators)
	servers_Databases_AdvancedThreatProtectionSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_Databases_AdvancedThreatProtectionSetting_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Databases_AdvancedThreatProtectionSetting_STATUS(generators)
	AddRelatedPropertyGeneratorsForServers_Databases_AdvancedThreatProtectionSetting_STATUS(generators)
	servers_Databases_AdvancedThreatProtectionSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_Databases_AdvancedThreatProtectionSetting_STATUS{}), generators)

	return servers_Databases_AdvancedThreatProtectionSetting_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_Databases_AdvancedThreatProtectionSetting_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Databases_AdvancedThreatProtectionSetting_STATUS(gens map[string]gopter.Gen) {
	gens["CreationTime"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(AdvancedThreatProtectionProperties_State_STATUS_Disabled, AdvancedThreatProtectionProperties_State_STATUS_Enabled, AdvancedThreatProtectionProperties_State_STATUS_New))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServers_Databases_AdvancedThreatProtectionSetting_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServers_Databases_AdvancedThreatProtectionSetting_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

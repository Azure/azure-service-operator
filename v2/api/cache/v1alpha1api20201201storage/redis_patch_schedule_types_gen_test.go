// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201storage

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/cache/v1beta20201201storage"
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

func Test_RedisPatchSchedule_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisPatchSchedule to hub returns original",
		prop.ForAll(RunResourceConversionTestForRedisPatchSchedule, RedisPatchScheduleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForRedisPatchSchedule tests if a specific instance of RedisPatchSchedule round trips to the hub storage version and back losslessly
func RunResourceConversionTestForRedisPatchSchedule(subject RedisPatchSchedule) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1beta20201201storage.RedisPatchSchedule
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual RedisPatchSchedule
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

func Test_RedisPatchSchedule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisPatchSchedule to RedisPatchSchedule via AssignPropertiesToRedisPatchSchedule & AssignPropertiesFromRedisPatchSchedule returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisPatchSchedule, RedisPatchScheduleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisPatchSchedule tests if a specific instance of RedisPatchSchedule can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForRedisPatchSchedule(subject RedisPatchSchedule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1beta20201201storage.RedisPatchSchedule
	err := copied.AssignPropertiesToRedisPatchSchedule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisPatchSchedule
	err = actual.AssignPropertiesFromRedisPatchSchedule(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RedisPatchSchedule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisPatchSchedule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisPatchSchedule, RedisPatchScheduleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisPatchSchedule runs a test to see if a specific instance of RedisPatchSchedule round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisPatchSchedule(subject RedisPatchSchedule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisPatchSchedule
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

// Generator of RedisPatchSchedule instances for property testing - lazily instantiated by RedisPatchScheduleGenerator()
var redisPatchScheduleGenerator gopter.Gen

// RedisPatchScheduleGenerator returns a generator of RedisPatchSchedule instances for property testing.
func RedisPatchScheduleGenerator() gopter.Gen {
	if redisPatchScheduleGenerator != nil {
		return redisPatchScheduleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedisPatchSchedule(generators)
	redisPatchScheduleGenerator = gen.Struct(reflect.TypeOf(RedisPatchSchedule{}), generators)

	return redisPatchScheduleGenerator
}

// AddRelatedPropertyGeneratorsForRedisPatchSchedule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisPatchSchedule(gens map[string]gopter.Gen) {
	gens["Spec"] = RedisPatchSchedulesSpecGenerator()
	gens["Status"] = RedisPatchScheduleStatusGenerator()
}

func Test_RedisPatchSchedule_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisPatchSchedule_Status to RedisPatchSchedule_Status via AssignPropertiesToRedisPatchScheduleStatus & AssignPropertiesFromRedisPatchScheduleStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisPatchScheduleStatus, RedisPatchScheduleStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisPatchScheduleStatus tests if a specific instance of RedisPatchSchedule_Status can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForRedisPatchScheduleStatus(subject RedisPatchSchedule_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1beta20201201storage.RedisPatchSchedule_Status
	err := copied.AssignPropertiesToRedisPatchScheduleStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisPatchSchedule_Status
	err = actual.AssignPropertiesFromRedisPatchScheduleStatus(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RedisPatchSchedule_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisPatchSchedule_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisPatchScheduleStatus, RedisPatchScheduleStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisPatchScheduleStatus runs a test to see if a specific instance of RedisPatchSchedule_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisPatchScheduleStatus(subject RedisPatchSchedule_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisPatchSchedule_Status
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

// Generator of RedisPatchSchedule_Status instances for property testing - lazily instantiated by
//RedisPatchScheduleStatusGenerator()
var redisPatchScheduleStatusGenerator gopter.Gen

// RedisPatchScheduleStatusGenerator returns a generator of RedisPatchSchedule_Status instances for property testing.
// We first initialize redisPatchScheduleStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisPatchScheduleStatusGenerator() gopter.Gen {
	if redisPatchScheduleStatusGenerator != nil {
		return redisPatchScheduleStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPatchScheduleStatus(generators)
	redisPatchScheduleStatusGenerator = gen.Struct(reflect.TypeOf(RedisPatchSchedule_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPatchScheduleStatus(generators)
	AddRelatedPropertyGeneratorsForRedisPatchScheduleStatus(generators)
	redisPatchScheduleStatusGenerator = gen.Struct(reflect.TypeOf(RedisPatchSchedule_Status{}), generators)

	return redisPatchScheduleStatusGenerator
}

// AddIndependentPropertyGeneratorsForRedisPatchScheduleStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisPatchScheduleStatus(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisPatchScheduleStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisPatchScheduleStatus(gens map[string]gopter.Gen) {
	gens["ScheduleEntries"] = gen.SliceOf(ScheduleEntryStatusGenerator())
}

func Test_RedisPatchSchedules_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisPatchSchedules_Spec to RedisPatchSchedules_Spec via AssignPropertiesToRedisPatchSchedulesSpec & AssignPropertiesFromRedisPatchSchedulesSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisPatchSchedulesSpec, RedisPatchSchedulesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisPatchSchedulesSpec tests if a specific instance of RedisPatchSchedules_Spec can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForRedisPatchSchedulesSpec(subject RedisPatchSchedules_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1beta20201201storage.RedisPatchSchedules_Spec
	err := copied.AssignPropertiesToRedisPatchSchedulesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisPatchSchedules_Spec
	err = actual.AssignPropertiesFromRedisPatchSchedulesSpec(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RedisPatchSchedules_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisPatchSchedules_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisPatchSchedulesSpec, RedisPatchSchedulesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisPatchSchedulesSpec runs a test to see if a specific instance of RedisPatchSchedules_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisPatchSchedulesSpec(subject RedisPatchSchedules_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisPatchSchedules_Spec
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

// Generator of RedisPatchSchedules_Spec instances for property testing - lazily instantiated by
//RedisPatchSchedulesSpecGenerator()
var redisPatchSchedulesSpecGenerator gopter.Gen

// RedisPatchSchedulesSpecGenerator returns a generator of RedisPatchSchedules_Spec instances for property testing.
// We first initialize redisPatchSchedulesSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisPatchSchedulesSpecGenerator() gopter.Gen {
	if redisPatchSchedulesSpecGenerator != nil {
		return redisPatchSchedulesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPatchSchedulesSpec(generators)
	redisPatchSchedulesSpecGenerator = gen.Struct(reflect.TypeOf(RedisPatchSchedules_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPatchSchedulesSpec(generators)
	AddRelatedPropertyGeneratorsForRedisPatchSchedulesSpec(generators)
	redisPatchSchedulesSpecGenerator = gen.Struct(reflect.TypeOf(RedisPatchSchedules_Spec{}), generators)

	return redisPatchSchedulesSpecGenerator
}

// AddIndependentPropertyGeneratorsForRedisPatchSchedulesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisPatchSchedulesSpec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisPatchSchedulesSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisPatchSchedulesSpec(gens map[string]gopter.Gen) {
	gens["ScheduleEntries"] = gen.SliceOf(ScheduleEntryGenerator())
}

func Test_ScheduleEntry_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ScheduleEntry to ScheduleEntry via AssignPropertiesToScheduleEntry & AssignPropertiesFromScheduleEntry returns original",
		prop.ForAll(RunPropertyAssignmentTestForScheduleEntry, ScheduleEntryGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForScheduleEntry tests if a specific instance of ScheduleEntry can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForScheduleEntry(subject ScheduleEntry) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1beta20201201storage.ScheduleEntry
	err := copied.AssignPropertiesToScheduleEntry(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ScheduleEntry
	err = actual.AssignPropertiesFromScheduleEntry(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ScheduleEntry_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduleEntry via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduleEntry, ScheduleEntryGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduleEntry runs a test to see if a specific instance of ScheduleEntry round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduleEntry(subject ScheduleEntry) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduleEntry
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

// Generator of ScheduleEntry instances for property testing - lazily instantiated by ScheduleEntryGenerator()
var scheduleEntryGenerator gopter.Gen

// ScheduleEntryGenerator returns a generator of ScheduleEntry instances for property testing.
func ScheduleEntryGenerator() gopter.Gen {
	if scheduleEntryGenerator != nil {
		return scheduleEntryGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduleEntry(generators)
	scheduleEntryGenerator = gen.Struct(reflect.TypeOf(ScheduleEntry{}), generators)

	return scheduleEntryGenerator
}

// AddIndependentPropertyGeneratorsForScheduleEntry is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScheduleEntry(gens map[string]gopter.Gen) {
	gens["DayOfWeek"] = gen.PtrOf(gen.AlphaString())
	gens["MaintenanceWindow"] = gen.PtrOf(gen.AlphaString())
	gens["StartHourUtc"] = gen.PtrOf(gen.Int())
}

func Test_ScheduleEntry_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ScheduleEntry_Status to ScheduleEntry_Status via AssignPropertiesToScheduleEntryStatus & AssignPropertiesFromScheduleEntryStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForScheduleEntryStatus, ScheduleEntryStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForScheduleEntryStatus tests if a specific instance of ScheduleEntry_Status can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForScheduleEntryStatus(subject ScheduleEntry_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1beta20201201storage.ScheduleEntry_Status
	err := copied.AssignPropertiesToScheduleEntryStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ScheduleEntry_Status
	err = actual.AssignPropertiesFromScheduleEntryStatus(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ScheduleEntry_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduleEntry_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduleEntryStatus, ScheduleEntryStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduleEntryStatus runs a test to see if a specific instance of ScheduleEntry_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduleEntryStatus(subject ScheduleEntry_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduleEntry_Status
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

// Generator of ScheduleEntry_Status instances for property testing - lazily instantiated by
//ScheduleEntryStatusGenerator()
var scheduleEntryStatusGenerator gopter.Gen

// ScheduleEntryStatusGenerator returns a generator of ScheduleEntry_Status instances for property testing.
func ScheduleEntryStatusGenerator() gopter.Gen {
	if scheduleEntryStatusGenerator != nil {
		return scheduleEntryStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduleEntryStatus(generators)
	scheduleEntryStatusGenerator = gen.Struct(reflect.TypeOf(ScheduleEntry_Status{}), generators)

	return scheduleEntryStatusGenerator
}

// AddIndependentPropertyGeneratorsForScheduleEntryStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScheduleEntryStatus(gens map[string]gopter.Gen) {
	gens["DayOfWeek"] = gen.PtrOf(gen.AlphaString())
	gens["MaintenanceWindow"] = gen.PtrOf(gen.AlphaString())
	gens["StartHourUtc"] = gen.PtrOf(gen.Int())
}

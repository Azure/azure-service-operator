// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_AbsoluteMonthlySchedule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AbsoluteMonthlySchedule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAbsoluteMonthlySchedule_STATUS, AbsoluteMonthlySchedule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAbsoluteMonthlySchedule_STATUS runs a test to see if a specific instance of AbsoluteMonthlySchedule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAbsoluteMonthlySchedule_STATUS(subject AbsoluteMonthlySchedule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AbsoluteMonthlySchedule_STATUS
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

// Generator of AbsoluteMonthlySchedule_STATUS instances for property testing - lazily instantiated by
// AbsoluteMonthlySchedule_STATUSGenerator()
var absoluteMonthlySchedule_STATUSGenerator gopter.Gen

// AbsoluteMonthlySchedule_STATUSGenerator returns a generator of AbsoluteMonthlySchedule_STATUS instances for property testing.
func AbsoluteMonthlySchedule_STATUSGenerator() gopter.Gen {
	if absoluteMonthlySchedule_STATUSGenerator != nil {
		return absoluteMonthlySchedule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAbsoluteMonthlySchedule_STATUS(generators)
	absoluteMonthlySchedule_STATUSGenerator = gen.Struct(reflect.TypeOf(AbsoluteMonthlySchedule_STATUS{}), generators)

	return absoluteMonthlySchedule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAbsoluteMonthlySchedule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAbsoluteMonthlySchedule_STATUS(gens map[string]gopter.Gen) {
	gens["DayOfMonth"] = gen.PtrOf(gen.Int())
	gens["IntervalMonths"] = gen.PtrOf(gen.Int())
}

func Test_DailySchedule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DailySchedule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDailySchedule_STATUS, DailySchedule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDailySchedule_STATUS runs a test to see if a specific instance of DailySchedule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDailySchedule_STATUS(subject DailySchedule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DailySchedule_STATUS
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

// Generator of DailySchedule_STATUS instances for property testing - lazily instantiated by
// DailySchedule_STATUSGenerator()
var dailySchedule_STATUSGenerator gopter.Gen

// DailySchedule_STATUSGenerator returns a generator of DailySchedule_STATUS instances for property testing.
func DailySchedule_STATUSGenerator() gopter.Gen {
	if dailySchedule_STATUSGenerator != nil {
		return dailySchedule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDailySchedule_STATUS(generators)
	dailySchedule_STATUSGenerator = gen.Struct(reflect.TypeOf(DailySchedule_STATUS{}), generators)

	return dailySchedule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDailySchedule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDailySchedule_STATUS(gens map[string]gopter.Gen) {
	gens["IntervalDays"] = gen.PtrOf(gen.Int())
}

func Test_DateSpan_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DateSpan_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDateSpan_STATUS, DateSpan_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDateSpan_STATUS runs a test to see if a specific instance of DateSpan_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDateSpan_STATUS(subject DateSpan_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DateSpan_STATUS
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

// Generator of DateSpan_STATUS instances for property testing - lazily instantiated by DateSpan_STATUSGenerator()
var dateSpan_STATUSGenerator gopter.Gen

// DateSpan_STATUSGenerator returns a generator of DateSpan_STATUS instances for property testing.
func DateSpan_STATUSGenerator() gopter.Gen {
	if dateSpan_STATUSGenerator != nil {
		return dateSpan_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDateSpan_STATUS(generators)
	dateSpan_STATUSGenerator = gen.Struct(reflect.TypeOf(DateSpan_STATUS{}), generators)

	return dateSpan_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDateSpan_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDateSpan_STATUS(gens map[string]gopter.Gen) {
	gens["End"] = gen.PtrOf(gen.AlphaString())
	gens["Start"] = gen.PtrOf(gen.AlphaString())
}

func Test_MaintenanceConfigurationProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MaintenanceConfigurationProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMaintenanceConfigurationProperties_STATUS, MaintenanceConfigurationProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMaintenanceConfigurationProperties_STATUS runs a test to see if a specific instance of MaintenanceConfigurationProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMaintenanceConfigurationProperties_STATUS(subject MaintenanceConfigurationProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MaintenanceConfigurationProperties_STATUS
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

// Generator of MaintenanceConfigurationProperties_STATUS instances for property testing - lazily instantiated by
// MaintenanceConfigurationProperties_STATUSGenerator()
var maintenanceConfigurationProperties_STATUSGenerator gopter.Gen

// MaintenanceConfigurationProperties_STATUSGenerator returns a generator of MaintenanceConfigurationProperties_STATUS instances for property testing.
func MaintenanceConfigurationProperties_STATUSGenerator() gopter.Gen {
	if maintenanceConfigurationProperties_STATUSGenerator != nil {
		return maintenanceConfigurationProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMaintenanceConfigurationProperties_STATUS(generators)
	maintenanceConfigurationProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(MaintenanceConfigurationProperties_STATUS{}), generators)

	return maintenanceConfigurationProperties_STATUSGenerator
}

// AddRelatedPropertyGeneratorsForMaintenanceConfigurationProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMaintenanceConfigurationProperties_STATUS(gens map[string]gopter.Gen) {
	gens["MaintenanceWindow"] = gen.PtrOf(MaintenanceWindow_STATUSGenerator())
	gens["NotAllowedTime"] = gen.SliceOf(TimeSpan_STATUSGenerator())
	gens["TimeInWeek"] = gen.SliceOf(TimeInWeek_STATUSGenerator())
}

func Test_MaintenanceConfiguration_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MaintenanceConfiguration_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMaintenanceConfiguration_STATUS, MaintenanceConfiguration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMaintenanceConfiguration_STATUS runs a test to see if a specific instance of MaintenanceConfiguration_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMaintenanceConfiguration_STATUS(subject MaintenanceConfiguration_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MaintenanceConfiguration_STATUS
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

// Generator of MaintenanceConfiguration_STATUS instances for property testing - lazily instantiated by
// MaintenanceConfiguration_STATUSGenerator()
var maintenanceConfiguration_STATUSGenerator gopter.Gen

// MaintenanceConfiguration_STATUSGenerator returns a generator of MaintenanceConfiguration_STATUS instances for property testing.
// We first initialize maintenanceConfiguration_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MaintenanceConfiguration_STATUSGenerator() gopter.Gen {
	if maintenanceConfiguration_STATUSGenerator != nil {
		return maintenanceConfiguration_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMaintenanceConfiguration_STATUS(generators)
	maintenanceConfiguration_STATUSGenerator = gen.Struct(reflect.TypeOf(MaintenanceConfiguration_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMaintenanceConfiguration_STATUS(generators)
	AddRelatedPropertyGeneratorsForMaintenanceConfiguration_STATUS(generators)
	maintenanceConfiguration_STATUSGenerator = gen.Struct(reflect.TypeOf(MaintenanceConfiguration_STATUS{}), generators)

	return maintenanceConfiguration_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForMaintenanceConfiguration_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMaintenanceConfiguration_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMaintenanceConfiguration_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMaintenanceConfiguration_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(MaintenanceConfigurationProperties_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_MaintenanceWindow_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MaintenanceWindow_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMaintenanceWindow_STATUS, MaintenanceWindow_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMaintenanceWindow_STATUS runs a test to see if a specific instance of MaintenanceWindow_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMaintenanceWindow_STATUS(subject MaintenanceWindow_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MaintenanceWindow_STATUS
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

// Generator of MaintenanceWindow_STATUS instances for property testing - lazily instantiated by
// MaintenanceWindow_STATUSGenerator()
var maintenanceWindow_STATUSGenerator gopter.Gen

// MaintenanceWindow_STATUSGenerator returns a generator of MaintenanceWindow_STATUS instances for property testing.
// We first initialize maintenanceWindow_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MaintenanceWindow_STATUSGenerator() gopter.Gen {
	if maintenanceWindow_STATUSGenerator != nil {
		return maintenanceWindow_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMaintenanceWindow_STATUS(generators)
	maintenanceWindow_STATUSGenerator = gen.Struct(reflect.TypeOf(MaintenanceWindow_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMaintenanceWindow_STATUS(generators)
	AddRelatedPropertyGeneratorsForMaintenanceWindow_STATUS(generators)
	maintenanceWindow_STATUSGenerator = gen.Struct(reflect.TypeOf(MaintenanceWindow_STATUS{}), generators)

	return maintenanceWindow_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForMaintenanceWindow_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMaintenanceWindow_STATUS(gens map[string]gopter.Gen) {
	gens["DurationHours"] = gen.PtrOf(gen.Int())
	gens["StartDate"] = gen.PtrOf(gen.AlphaString())
	gens["StartTime"] = gen.PtrOf(gen.AlphaString())
	gens["UtcOffset"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMaintenanceWindow_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMaintenanceWindow_STATUS(gens map[string]gopter.Gen) {
	gens["NotAllowedDates"] = gen.SliceOf(DateSpan_STATUSGenerator())
	gens["Schedule"] = gen.PtrOf(Schedule_STATUSGenerator())
}

func Test_RelativeMonthlySchedule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RelativeMonthlySchedule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRelativeMonthlySchedule_STATUS, RelativeMonthlySchedule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRelativeMonthlySchedule_STATUS runs a test to see if a specific instance of RelativeMonthlySchedule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRelativeMonthlySchedule_STATUS(subject RelativeMonthlySchedule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RelativeMonthlySchedule_STATUS
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

// Generator of RelativeMonthlySchedule_STATUS instances for property testing - lazily instantiated by
// RelativeMonthlySchedule_STATUSGenerator()
var relativeMonthlySchedule_STATUSGenerator gopter.Gen

// RelativeMonthlySchedule_STATUSGenerator returns a generator of RelativeMonthlySchedule_STATUS instances for property testing.
func RelativeMonthlySchedule_STATUSGenerator() gopter.Gen {
	if relativeMonthlySchedule_STATUSGenerator != nil {
		return relativeMonthlySchedule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRelativeMonthlySchedule_STATUS(generators)
	relativeMonthlySchedule_STATUSGenerator = gen.Struct(reflect.TypeOf(RelativeMonthlySchedule_STATUS{}), generators)

	return relativeMonthlySchedule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRelativeMonthlySchedule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRelativeMonthlySchedule_STATUS(gens map[string]gopter.Gen) {
	gens["DayOfWeek"] = gen.PtrOf(gen.OneConstOf(
		WeekDay_STATUS_Friday,
		WeekDay_STATUS_Monday,
		WeekDay_STATUS_Saturday,
		WeekDay_STATUS_Sunday,
		WeekDay_STATUS_Thursday,
		WeekDay_STATUS_Tuesday,
		WeekDay_STATUS_Wednesday))
	gens["IntervalMonths"] = gen.PtrOf(gen.Int())
	gens["WeekIndex"] = gen.PtrOf(gen.OneConstOf(
		RelativeMonthlySchedule_WeekIndex_STATUS_First,
		RelativeMonthlySchedule_WeekIndex_STATUS_Fourth,
		RelativeMonthlySchedule_WeekIndex_STATUS_Last,
		RelativeMonthlySchedule_WeekIndex_STATUS_Second,
		RelativeMonthlySchedule_WeekIndex_STATUS_Third))
}

func Test_Schedule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Schedule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSchedule_STATUS, Schedule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSchedule_STATUS runs a test to see if a specific instance of Schedule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSchedule_STATUS(subject Schedule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Schedule_STATUS
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

// Generator of Schedule_STATUS instances for property testing - lazily instantiated by Schedule_STATUSGenerator()
var schedule_STATUSGenerator gopter.Gen

// Schedule_STATUSGenerator returns a generator of Schedule_STATUS instances for property testing.
func Schedule_STATUSGenerator() gopter.Gen {
	if schedule_STATUSGenerator != nil {
		return schedule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSchedule_STATUS(generators)
	schedule_STATUSGenerator = gen.Struct(reflect.TypeOf(Schedule_STATUS{}), generators)

	return schedule_STATUSGenerator
}

// AddRelatedPropertyGeneratorsForSchedule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSchedule_STATUS(gens map[string]gopter.Gen) {
	gens["AbsoluteMonthly"] = gen.PtrOf(AbsoluteMonthlySchedule_STATUSGenerator())
	gens["Daily"] = gen.PtrOf(DailySchedule_STATUSGenerator())
	gens["RelativeMonthly"] = gen.PtrOf(RelativeMonthlySchedule_STATUSGenerator())
	gens["Weekly"] = gen.PtrOf(WeeklySchedule_STATUSGenerator())
}

func Test_SystemData_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUS, SystemData_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUS runs a test to see if a specific instance of SystemData_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUS(subject SystemData_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUS
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

// Generator of SystemData_STATUS instances for property testing - lazily instantiated by SystemData_STATUSGenerator()
var systemData_STATUSGenerator gopter.Gen

// SystemData_STATUSGenerator returns a generator of SystemData_STATUS instances for property testing.
func SystemData_STATUSGenerator() gopter.Gen {
	if systemData_STATUSGenerator != nil {
		return systemData_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUS(generators)
	systemData_STATUSGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUS{}), generators)

	return systemData_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUS(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_CreatedByType_STATUS_Application,
		SystemData_CreatedByType_STATUS_Key,
		SystemData_CreatedByType_STATUS_ManagedIdentity,
		SystemData_CreatedByType_STATUS_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_LastModifiedByType_STATUS_Application,
		SystemData_LastModifiedByType_STATUS_Key,
		SystemData_LastModifiedByType_STATUS_ManagedIdentity,
		SystemData_LastModifiedByType_STATUS_User))
}

func Test_TimeInWeek_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TimeInWeek_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTimeInWeek_STATUS, TimeInWeek_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTimeInWeek_STATUS runs a test to see if a specific instance of TimeInWeek_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTimeInWeek_STATUS(subject TimeInWeek_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TimeInWeek_STATUS
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

// Generator of TimeInWeek_STATUS instances for property testing - lazily instantiated by TimeInWeek_STATUSGenerator()
var timeInWeek_STATUSGenerator gopter.Gen

// TimeInWeek_STATUSGenerator returns a generator of TimeInWeek_STATUS instances for property testing.
func TimeInWeek_STATUSGenerator() gopter.Gen {
	if timeInWeek_STATUSGenerator != nil {
		return timeInWeek_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTimeInWeek_STATUS(generators)
	timeInWeek_STATUSGenerator = gen.Struct(reflect.TypeOf(TimeInWeek_STATUS{}), generators)

	return timeInWeek_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTimeInWeek_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTimeInWeek_STATUS(gens map[string]gopter.Gen) {
	gens["Day"] = gen.PtrOf(gen.OneConstOf(
		WeekDay_STATUS_Friday,
		WeekDay_STATUS_Monday,
		WeekDay_STATUS_Saturday,
		WeekDay_STATUS_Sunday,
		WeekDay_STATUS_Thursday,
		WeekDay_STATUS_Tuesday,
		WeekDay_STATUS_Wednesday))
	gens["HourSlots"] = gen.SliceOf(gen.Int())
}

func Test_TimeSpan_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TimeSpan_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTimeSpan_STATUS, TimeSpan_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTimeSpan_STATUS runs a test to see if a specific instance of TimeSpan_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTimeSpan_STATUS(subject TimeSpan_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TimeSpan_STATUS
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

// Generator of TimeSpan_STATUS instances for property testing - lazily instantiated by TimeSpan_STATUSGenerator()
var timeSpan_STATUSGenerator gopter.Gen

// TimeSpan_STATUSGenerator returns a generator of TimeSpan_STATUS instances for property testing.
func TimeSpan_STATUSGenerator() gopter.Gen {
	if timeSpan_STATUSGenerator != nil {
		return timeSpan_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTimeSpan_STATUS(generators)
	timeSpan_STATUSGenerator = gen.Struct(reflect.TypeOf(TimeSpan_STATUS{}), generators)

	return timeSpan_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTimeSpan_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTimeSpan_STATUS(gens map[string]gopter.Gen) {
	gens["End"] = gen.PtrOf(gen.AlphaString())
	gens["Start"] = gen.PtrOf(gen.AlphaString())
}

func Test_WeeklySchedule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WeeklySchedule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWeeklySchedule_STATUS, WeeklySchedule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWeeklySchedule_STATUS runs a test to see if a specific instance of WeeklySchedule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForWeeklySchedule_STATUS(subject WeeklySchedule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WeeklySchedule_STATUS
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

// Generator of WeeklySchedule_STATUS instances for property testing - lazily instantiated by
// WeeklySchedule_STATUSGenerator()
var weeklySchedule_STATUSGenerator gopter.Gen

// WeeklySchedule_STATUSGenerator returns a generator of WeeklySchedule_STATUS instances for property testing.
func WeeklySchedule_STATUSGenerator() gopter.Gen {
	if weeklySchedule_STATUSGenerator != nil {
		return weeklySchedule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWeeklySchedule_STATUS(generators)
	weeklySchedule_STATUSGenerator = gen.Struct(reflect.TypeOf(WeeklySchedule_STATUS{}), generators)

	return weeklySchedule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForWeeklySchedule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWeeklySchedule_STATUS(gens map[string]gopter.Gen) {
	gens["DayOfWeek"] = gen.PtrOf(gen.OneConstOf(
		WeekDay_STATUS_Friday,
		WeekDay_STATUS_Monday,
		WeekDay_STATUS_Saturday,
		WeekDay_STATUS_Sunday,
		WeekDay_STATUS_Thursday,
		WeekDay_STATUS_Tuesday,
		WeekDay_STATUS_Wednesday))
	gens["IntervalWeeks"] = gen.PtrOf(gen.Int())
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201201storage

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

func Test_RedisPatchSchedule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
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
	gens["Spec"] = Redis_PatchSchedule_SpecGenerator()
	gens["Status"] = Redis_PatchSchedule_STATUSGenerator()
}

func Test_Redis_PatchSchedule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis_PatchSchedule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis_PatchSchedule_Spec, Redis_PatchSchedule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis_PatchSchedule_Spec runs a test to see if a specific instance of Redis_PatchSchedule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis_PatchSchedule_Spec(subject Redis_PatchSchedule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis_PatchSchedule_Spec
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

// Generator of Redis_PatchSchedule_Spec instances for property testing - lazily instantiated by
// Redis_PatchSchedule_SpecGenerator()
var redis_PatchSchedule_SpecGenerator gopter.Gen

// Redis_PatchSchedule_SpecGenerator returns a generator of Redis_PatchSchedule_Spec instances for property testing.
// We first initialize redis_PatchSchedule_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Redis_PatchSchedule_SpecGenerator() gopter.Gen {
	if redis_PatchSchedule_SpecGenerator != nil {
		return redis_PatchSchedule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_PatchSchedule_Spec(generators)
	redis_PatchSchedule_SpecGenerator = gen.Struct(reflect.TypeOf(Redis_PatchSchedule_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_PatchSchedule_Spec(generators)
	AddRelatedPropertyGeneratorsForRedis_PatchSchedule_Spec(generators)
	redis_PatchSchedule_SpecGenerator = gen.Struct(reflect.TypeOf(Redis_PatchSchedule_Spec{}), generators)

	return redis_PatchSchedule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRedis_PatchSchedule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedis_PatchSchedule_Spec(gens map[string]gopter.Gen) {
	gens["OriginalVersion"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForRedis_PatchSchedule_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedis_PatchSchedule_Spec(gens map[string]gopter.Gen) {
	gens["ScheduleEntries"] = gen.SliceOf(ScheduleEntryGenerator())
}

func Test_Redis_PatchSchedule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis_PatchSchedule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis_PatchSchedule_STATUS, Redis_PatchSchedule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis_PatchSchedule_STATUS runs a test to see if a specific instance of Redis_PatchSchedule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis_PatchSchedule_STATUS(subject Redis_PatchSchedule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis_PatchSchedule_STATUS
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

// Generator of Redis_PatchSchedule_STATUS instances for property testing - lazily instantiated by
// Redis_PatchSchedule_STATUSGenerator()
var redis_PatchSchedule_STATUSGenerator gopter.Gen

// Redis_PatchSchedule_STATUSGenerator returns a generator of Redis_PatchSchedule_STATUS instances for property testing.
// We first initialize redis_PatchSchedule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Redis_PatchSchedule_STATUSGenerator() gopter.Gen {
	if redis_PatchSchedule_STATUSGenerator != nil {
		return redis_PatchSchedule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_PatchSchedule_STATUS(generators)
	redis_PatchSchedule_STATUSGenerator = gen.Struct(reflect.TypeOf(Redis_PatchSchedule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_PatchSchedule_STATUS(generators)
	AddRelatedPropertyGeneratorsForRedis_PatchSchedule_STATUS(generators)
	redis_PatchSchedule_STATUSGenerator = gen.Struct(reflect.TypeOf(Redis_PatchSchedule_STATUS{}), generators)

	return redis_PatchSchedule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedis_PatchSchedule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedis_PatchSchedule_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedis_PatchSchedule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedis_PatchSchedule_STATUS(gens map[string]gopter.Gen) {
	gens["ScheduleEntries"] = gen.SliceOf(ScheduleEntry_STATUSGenerator())
}

func Test_ScheduleEntry_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
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

func Test_ScheduleEntry_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduleEntry_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduleEntry_STATUS, ScheduleEntry_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduleEntry_STATUS runs a test to see if a specific instance of ScheduleEntry_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduleEntry_STATUS(subject ScheduleEntry_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduleEntry_STATUS
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

// Generator of ScheduleEntry_STATUS instances for property testing - lazily instantiated by
// ScheduleEntry_STATUSGenerator()
var scheduleEntry_STATUSGenerator gopter.Gen

// ScheduleEntry_STATUSGenerator returns a generator of ScheduleEntry_STATUS instances for property testing.
func ScheduleEntry_STATUSGenerator() gopter.Gen {
	if scheduleEntry_STATUSGenerator != nil {
		return scheduleEntry_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduleEntry_STATUS(generators)
	scheduleEntry_STATUSGenerator = gen.Struct(reflect.TypeOf(ScheduleEntry_STATUS{}), generators)

	return scheduleEntry_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForScheduleEntry_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScheduleEntry_STATUS(gens map[string]gopter.Gen) {
	gens["DayOfWeek"] = gen.PtrOf(gen.AlphaString())
	gens["MaintenanceWindow"] = gen.PtrOf(gen.AlphaString())
	gens["StartHourUtc"] = gen.PtrOf(gen.Int())
}

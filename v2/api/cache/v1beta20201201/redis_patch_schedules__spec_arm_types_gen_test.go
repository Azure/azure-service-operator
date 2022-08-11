// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201

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

func Test_RedisPatchSchedules_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisPatchSchedules_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisPatchSchedulesSpecARM, RedisPatchSchedulesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisPatchSchedulesSpecARM runs a test to see if a specific instance of RedisPatchSchedules_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisPatchSchedulesSpecARM(subject RedisPatchSchedules_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisPatchSchedules_SpecARM
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

// Generator of RedisPatchSchedules_SpecARM instances for property testing - lazily instantiated by
// RedisPatchSchedulesSpecARMGenerator()
var redisPatchSchedulesSpecARMGenerator gopter.Gen

// RedisPatchSchedulesSpecARMGenerator returns a generator of RedisPatchSchedules_SpecARM instances for property testing.
// We first initialize redisPatchSchedulesSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisPatchSchedulesSpecARMGenerator() gopter.Gen {
	if redisPatchSchedulesSpecARMGenerator != nil {
		return redisPatchSchedulesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPatchSchedulesSpecARM(generators)
	redisPatchSchedulesSpecARMGenerator = gen.Struct(reflect.TypeOf(RedisPatchSchedules_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPatchSchedulesSpecARM(generators)
	AddRelatedPropertyGeneratorsForRedisPatchSchedulesSpecARM(generators)
	redisPatchSchedulesSpecARMGenerator = gen.Struct(reflect.TypeOf(RedisPatchSchedules_SpecARM{}), generators)

	return redisPatchSchedulesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisPatchSchedulesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisPatchSchedulesSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisPatchSchedulesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisPatchSchedulesSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ScheduleEntriesARMGenerator())
}

func Test_ScheduleEntriesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduleEntriesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduleEntriesARM, ScheduleEntriesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduleEntriesARM runs a test to see if a specific instance of ScheduleEntriesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduleEntriesARM(subject ScheduleEntriesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduleEntriesARM
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

// Generator of ScheduleEntriesARM instances for property testing - lazily instantiated by ScheduleEntriesARMGenerator()
var scheduleEntriesARMGenerator gopter.Gen

// ScheduleEntriesARMGenerator returns a generator of ScheduleEntriesARM instances for property testing.
func ScheduleEntriesARMGenerator() gopter.Gen {
	if scheduleEntriesARMGenerator != nil {
		return scheduleEntriesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForScheduleEntriesARM(generators)
	scheduleEntriesARMGenerator = gen.Struct(reflect.TypeOf(ScheduleEntriesARM{}), generators)

	return scheduleEntriesARMGenerator
}

// AddRelatedPropertyGeneratorsForScheduleEntriesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScheduleEntriesARM(gens map[string]gopter.Gen) {
	gens["ScheduleEntries"] = gen.SliceOf(ScheduleEntryARMGenerator())
}

func Test_ScheduleEntryARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduleEntryARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduleEntryARM, ScheduleEntryARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduleEntryARM runs a test to see if a specific instance of ScheduleEntryARM round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduleEntryARM(subject ScheduleEntryARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduleEntryARM
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

// Generator of ScheduleEntryARM instances for property testing - lazily instantiated by ScheduleEntryARMGenerator()
var scheduleEntryARMGenerator gopter.Gen

// ScheduleEntryARMGenerator returns a generator of ScheduleEntryARM instances for property testing.
func ScheduleEntryARMGenerator() gopter.Gen {
	if scheduleEntryARMGenerator != nil {
		return scheduleEntryARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduleEntryARM(generators)
	scheduleEntryARMGenerator = gen.Struct(reflect.TypeOf(ScheduleEntryARM{}), generators)

	return scheduleEntryARMGenerator
}

// AddIndependentPropertyGeneratorsForScheduleEntryARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScheduleEntryARM(gens map[string]gopter.Gen) {
	gens["DayOfWeek"] = gen.PtrOf(gen.OneConstOf(
		ScheduleEntryDayOfWeek_Everyday,
		ScheduleEntryDayOfWeek_Friday,
		ScheduleEntryDayOfWeek_Monday,
		ScheduleEntryDayOfWeek_Saturday,
		ScheduleEntryDayOfWeek_Sunday,
		ScheduleEntryDayOfWeek_Thursday,
		ScheduleEntryDayOfWeek_Tuesday,
		ScheduleEntryDayOfWeek_Wednesday,
		ScheduleEntryDayOfWeek_Weekend))
	gens["MaintenanceWindow"] = gen.PtrOf(gen.AlphaString())
	gens["StartHourUtc"] = gen.PtrOf(gen.Int())
}

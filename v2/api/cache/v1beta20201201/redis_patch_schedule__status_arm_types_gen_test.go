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

func Test_RedisPatchSchedule_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisPatchSchedule_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisPatchScheduleStatusARM, RedisPatchScheduleStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisPatchScheduleStatusARM runs a test to see if a specific instance of RedisPatchSchedule_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisPatchScheduleStatusARM(subject RedisPatchSchedule_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisPatchSchedule_StatusARM
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

// Generator of RedisPatchSchedule_StatusARM instances for property testing - lazily instantiated by
// RedisPatchScheduleStatusARMGenerator()
var redisPatchScheduleStatusARMGenerator gopter.Gen

// RedisPatchScheduleStatusARMGenerator returns a generator of RedisPatchSchedule_StatusARM instances for property testing.
// We first initialize redisPatchScheduleStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisPatchScheduleStatusARMGenerator() gopter.Gen {
	if redisPatchScheduleStatusARMGenerator != nil {
		return redisPatchScheduleStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPatchScheduleStatusARM(generators)
	redisPatchScheduleStatusARMGenerator = gen.Struct(reflect.TypeOf(RedisPatchSchedule_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPatchScheduleStatusARM(generators)
	AddRelatedPropertyGeneratorsForRedisPatchScheduleStatusARM(generators)
	redisPatchScheduleStatusARMGenerator = gen.Struct(reflect.TypeOf(RedisPatchSchedule_StatusARM{}), generators)

	return redisPatchScheduleStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisPatchScheduleStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisPatchScheduleStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisPatchScheduleStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisPatchScheduleStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ScheduleEntriesStatusARMGenerator())
}

func Test_ScheduleEntries_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduleEntries_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduleEntriesStatusARM, ScheduleEntriesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduleEntriesStatusARM runs a test to see if a specific instance of ScheduleEntries_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduleEntriesStatusARM(subject ScheduleEntries_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduleEntries_StatusARM
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

// Generator of ScheduleEntries_StatusARM instances for property testing - lazily instantiated by
// ScheduleEntriesStatusARMGenerator()
var scheduleEntriesStatusARMGenerator gopter.Gen

// ScheduleEntriesStatusARMGenerator returns a generator of ScheduleEntries_StatusARM instances for property testing.
func ScheduleEntriesStatusARMGenerator() gopter.Gen {
	if scheduleEntriesStatusARMGenerator != nil {
		return scheduleEntriesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForScheduleEntriesStatusARM(generators)
	scheduleEntriesStatusARMGenerator = gen.Struct(reflect.TypeOf(ScheduleEntries_StatusARM{}), generators)

	return scheduleEntriesStatusARMGenerator
}

// AddRelatedPropertyGeneratorsForScheduleEntriesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScheduleEntriesStatusARM(gens map[string]gopter.Gen) {
	gens["ScheduleEntries"] = gen.SliceOf(ScheduleEntryStatusARMGenerator())
}

func Test_ScheduleEntry_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduleEntry_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduleEntryStatusARM, ScheduleEntryStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduleEntryStatusARM runs a test to see if a specific instance of ScheduleEntry_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduleEntryStatusARM(subject ScheduleEntry_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduleEntry_StatusARM
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

// Generator of ScheduleEntry_StatusARM instances for property testing - lazily instantiated by
// ScheduleEntryStatusARMGenerator()
var scheduleEntryStatusARMGenerator gopter.Gen

// ScheduleEntryStatusARMGenerator returns a generator of ScheduleEntry_StatusARM instances for property testing.
func ScheduleEntryStatusARMGenerator() gopter.Gen {
	if scheduleEntryStatusARMGenerator != nil {
		return scheduleEntryStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduleEntryStatusARM(generators)
	scheduleEntryStatusARMGenerator = gen.Struct(reflect.TypeOf(ScheduleEntry_StatusARM{}), generators)

	return scheduleEntryStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForScheduleEntryStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScheduleEntryStatusARM(gens map[string]gopter.Gen) {
	gens["DayOfWeek"] = gen.PtrOf(gen.OneConstOf(
		ScheduleEntryStatusDayOfWeek_Everyday,
		ScheduleEntryStatusDayOfWeek_Friday,
		ScheduleEntryStatusDayOfWeek_Monday,
		ScheduleEntryStatusDayOfWeek_Saturday,
		ScheduleEntryStatusDayOfWeek_Sunday,
		ScheduleEntryStatusDayOfWeek_Thursday,
		ScheduleEntryStatusDayOfWeek_Tuesday,
		ScheduleEntryStatusDayOfWeek_Wednesday,
		ScheduleEntryStatusDayOfWeek_Weekend))
	gens["MaintenanceWindow"] = gen.PtrOf(gen.AlphaString())
	gens["StartHourUtc"] = gen.PtrOf(gen.Int())
}

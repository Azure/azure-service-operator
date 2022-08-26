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

func Test_RedisPatchSchedule_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisPatchSchedule_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisPatchSchedule_STATUSARM, RedisPatchSchedule_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisPatchSchedule_STATUSARM runs a test to see if a specific instance of RedisPatchSchedule_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisPatchSchedule_STATUSARM(subject RedisPatchSchedule_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisPatchSchedule_STATUSARM
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

// Generator of RedisPatchSchedule_STATUSARM instances for property testing - lazily instantiated by
// RedisPatchSchedule_STATUSARMGenerator()
var redisPatchSchedule_STATUSARMGenerator gopter.Gen

// RedisPatchSchedule_STATUSARMGenerator returns a generator of RedisPatchSchedule_STATUSARM instances for property testing.
// We first initialize redisPatchSchedule_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisPatchSchedule_STATUSARMGenerator() gopter.Gen {
	if redisPatchSchedule_STATUSARMGenerator != nil {
		return redisPatchSchedule_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPatchSchedule_STATUSARM(generators)
	redisPatchSchedule_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisPatchSchedule_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPatchSchedule_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForRedisPatchSchedule_STATUSARM(generators)
	redisPatchSchedule_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisPatchSchedule_STATUSARM{}), generators)

	return redisPatchSchedule_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisPatchSchedule_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisPatchSchedule_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisPatchSchedule_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisPatchSchedule_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ScheduleEntries_STATUSARMGenerator())
}

func Test_ScheduleEntries_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduleEntries_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduleEntries_STATUSARM, ScheduleEntries_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduleEntries_STATUSARM runs a test to see if a specific instance of ScheduleEntries_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduleEntries_STATUSARM(subject ScheduleEntries_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduleEntries_STATUSARM
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

// Generator of ScheduleEntries_STATUSARM instances for property testing - lazily instantiated by
// ScheduleEntries_STATUSARMGenerator()
var scheduleEntries_STATUSARMGenerator gopter.Gen

// ScheduleEntries_STATUSARMGenerator returns a generator of ScheduleEntries_STATUSARM instances for property testing.
func ScheduleEntries_STATUSARMGenerator() gopter.Gen {
	if scheduleEntries_STATUSARMGenerator != nil {
		return scheduleEntries_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForScheduleEntries_STATUSARM(generators)
	scheduleEntries_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ScheduleEntries_STATUSARM{}), generators)

	return scheduleEntries_STATUSARMGenerator
}

// AddRelatedPropertyGeneratorsForScheduleEntries_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScheduleEntries_STATUSARM(gens map[string]gopter.Gen) {
	gens["ScheduleEntries"] = gen.SliceOf(ScheduleEntry_STATUSARMGenerator())
}

func Test_ScheduleEntry_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduleEntry_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduleEntry_STATUSARM, ScheduleEntry_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduleEntry_STATUSARM runs a test to see if a specific instance of ScheduleEntry_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduleEntry_STATUSARM(subject ScheduleEntry_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduleEntry_STATUSARM
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

// Generator of ScheduleEntry_STATUSARM instances for property testing - lazily instantiated by
// ScheduleEntry_STATUSARMGenerator()
var scheduleEntry_STATUSARMGenerator gopter.Gen

// ScheduleEntry_STATUSARMGenerator returns a generator of ScheduleEntry_STATUSARM instances for property testing.
func ScheduleEntry_STATUSARMGenerator() gopter.Gen {
	if scheduleEntry_STATUSARMGenerator != nil {
		return scheduleEntry_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduleEntry_STATUSARM(generators)
	scheduleEntry_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ScheduleEntry_STATUSARM{}), generators)

	return scheduleEntry_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForScheduleEntry_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScheduleEntry_STATUSARM(gens map[string]gopter.Gen) {
	gens["DayOfWeek"] = gen.PtrOf(gen.OneConstOf(
<<<<<<< HEAD
		ScheduleEntry_DayOfWeek_Everyday_STATUS,
		ScheduleEntry_DayOfWeek_Friday_STATUS,
		ScheduleEntry_DayOfWeek_Monday_STATUS,
		ScheduleEntry_DayOfWeek_Saturday_STATUS,
		ScheduleEntry_DayOfWeek_Sunday_STATUS,
		ScheduleEntry_DayOfWeek_Thursday_STATUS,
		ScheduleEntry_DayOfWeek_Tuesday_STATUS,
		ScheduleEntry_DayOfWeek_Wednesday_STATUS,
		ScheduleEntry_DayOfWeek_Weekend_STATUS))
=======
		ScheduleEntry_STATUS_DayOfWeek_Everyday,
		ScheduleEntry_STATUS_DayOfWeek_Friday,
		ScheduleEntry_STATUS_DayOfWeek_Monday,
		ScheduleEntry_STATUS_DayOfWeek_Saturday,
		ScheduleEntry_STATUS_DayOfWeek_Sunday,
		ScheduleEntry_STATUS_DayOfWeek_Thursday,
		ScheduleEntry_STATUS_DayOfWeek_Tuesday,
		ScheduleEntry_STATUS_DayOfWeek_Wednesday,
		ScheduleEntry_STATUS_DayOfWeek_Weekend))
>>>>>>> main
	gens["MaintenanceWindow"] = gen.PtrOf(gen.AlphaString())
	gens["StartHourUtc"] = gen.PtrOf(gen.Int())
}

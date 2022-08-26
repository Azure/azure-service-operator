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

<<<<<<<< HEAD:v2/api/cache/v1beta20201201/redis_patch_schedule__spec_arm_types_gen_test.go
func Test_RedisPatchSchedule_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
========
func Test_Redis_PatchSchedules_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
>>>>>>>> main:v2/api/cache/v1beta20201201/redis_patch_schedules_spec_arm_types_gen_test.go
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<<< HEAD:v2/api/cache/v1beta20201201/redis_patch_schedule__spec_arm_types_gen_test.go
		"Round trip of RedisPatchSchedule_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisPatchSchedule_SpecARM, RedisPatchSchedule_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisPatchSchedule_SpecARM runs a test to see if a specific instance of RedisPatchSchedule_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisPatchSchedule_SpecARM(subject RedisPatchSchedule_SpecARM) string {
========
		"Round trip of Redis_PatchSchedules_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis_PatchSchedules_SpecARM, Redis_PatchSchedules_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis_PatchSchedules_SpecARM runs a test to see if a specific instance of Redis_PatchSchedules_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis_PatchSchedules_SpecARM(subject Redis_PatchSchedules_SpecARM) string {
>>>>>>>> main:v2/api/cache/v1beta20201201/redis_patch_schedules_spec_arm_types_gen_test.go
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
<<<<<<<< HEAD:v2/api/cache/v1beta20201201/redis_patch_schedule__spec_arm_types_gen_test.go
	var actual RedisPatchSchedule_SpecARM
========
	var actual Redis_PatchSchedules_SpecARM
>>>>>>>> main:v2/api/cache/v1beta20201201/redis_patch_schedules_spec_arm_types_gen_test.go
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

<<<<<<<< HEAD:v2/api/cache/v1beta20201201/redis_patch_schedule__spec_arm_types_gen_test.go
// Generator of RedisPatchSchedule_SpecARM instances for property testing - lazily instantiated by
// RedisPatchSchedule_SpecARMGenerator()
var redisPatchSchedule_SpecARMGenerator gopter.Gen

// RedisPatchSchedule_SpecARMGenerator returns a generator of RedisPatchSchedule_SpecARM instances for property testing.
// We first initialize redisPatchSchedule_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisPatchSchedule_SpecARMGenerator() gopter.Gen {
	if redisPatchSchedule_SpecARMGenerator != nil {
		return redisPatchSchedule_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPatchSchedule_SpecARM(generators)
	redisPatchSchedule_SpecARMGenerator = gen.Struct(reflect.TypeOf(RedisPatchSchedule_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPatchSchedule_SpecARM(generators)
	AddRelatedPropertyGeneratorsForRedisPatchSchedule_SpecARM(generators)
	redisPatchSchedule_SpecARMGenerator = gen.Struct(reflect.TypeOf(RedisPatchSchedule_SpecARM{}), generators)

	return redisPatchSchedule_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisPatchSchedule_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisPatchSchedule_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
========
// Generator of Redis_PatchSchedules_SpecARM instances for property testing - lazily instantiated by
// Redis_PatchSchedules_SpecARMGenerator()
var redis_PatchSchedules_SpecARMGenerator gopter.Gen

// Redis_PatchSchedules_SpecARMGenerator returns a generator of Redis_PatchSchedules_SpecARM instances for property testing.
// We first initialize redis_PatchSchedules_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Redis_PatchSchedules_SpecARMGenerator() gopter.Gen {
	if redis_PatchSchedules_SpecARMGenerator != nil {
		return redis_PatchSchedules_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_PatchSchedules_SpecARM(generators)
	redis_PatchSchedules_SpecARMGenerator = gen.Struct(reflect.TypeOf(Redis_PatchSchedules_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_PatchSchedules_SpecARM(generators)
	AddRelatedPropertyGeneratorsForRedis_PatchSchedules_SpecARM(generators)
	redis_PatchSchedules_SpecARMGenerator = gen.Struct(reflect.TypeOf(Redis_PatchSchedules_SpecARM{}), generators)

	return redis_PatchSchedules_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForRedis_PatchSchedules_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedis_PatchSchedules_SpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
>>>>>>>> main:v2/api/cache/v1beta20201201/redis_patch_schedules_spec_arm_types_gen_test.go
	gens["Name"] = gen.AlphaString()
}

<<<<<<<< HEAD:v2/api/cache/v1beta20201201/redis_patch_schedule__spec_arm_types_gen_test.go
// AddRelatedPropertyGeneratorsForRedisPatchSchedule_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisPatchSchedule_SpecARM(gens map[string]gopter.Gen) {
========
// AddRelatedPropertyGeneratorsForRedis_PatchSchedules_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedis_PatchSchedules_SpecARM(gens map[string]gopter.Gen) {
>>>>>>>> main:v2/api/cache/v1beta20201201/redis_patch_schedules_spec_arm_types_gen_test.go
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
		ScheduleEntry_DayOfWeek_Everyday,
		ScheduleEntry_DayOfWeek_Friday,
		ScheduleEntry_DayOfWeek_Monday,
		ScheduleEntry_DayOfWeek_Saturday,
		ScheduleEntry_DayOfWeek_Sunday,
		ScheduleEntry_DayOfWeek_Thursday,
		ScheduleEntry_DayOfWeek_Tuesday,
		ScheduleEntry_DayOfWeek_Wednesday,
		ScheduleEntry_DayOfWeek_Weekend))
	gens["MaintenanceWindow"] = gen.PtrOf(gen.AlphaString())
	gens["StartHourUtc"] = gen.PtrOf(gen.Int())
}

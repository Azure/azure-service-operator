// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201201

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

func Test_RedisLinkedServerCreateProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServerCreateProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServerCreateProperties_ARM, RedisLinkedServerCreateProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServerCreateProperties_ARM runs a test to see if a specific instance of RedisLinkedServerCreateProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServerCreateProperties_ARM(subject RedisLinkedServerCreateProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServerCreateProperties_ARM
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

// Generator of RedisLinkedServerCreateProperties_ARM instances for property testing - lazily instantiated by
// RedisLinkedServerCreateProperties_ARMGenerator()
var redisLinkedServerCreateProperties_ARMGenerator gopter.Gen

// RedisLinkedServerCreateProperties_ARMGenerator returns a generator of RedisLinkedServerCreateProperties_ARM instances for property testing.
func RedisLinkedServerCreateProperties_ARMGenerator() gopter.Gen {
	if redisLinkedServerCreateProperties_ARMGenerator != nil {
		return redisLinkedServerCreateProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServerCreateProperties_ARM(generators)
	redisLinkedServerCreateProperties_ARMGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServerCreateProperties_ARM{}), generators)

	return redisLinkedServerCreateProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServerCreateProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServerCreateProperties_ARM(gens map[string]gopter.Gen) {
	gens["LinkedRedisCacheId"] = gen.PtrOf(gen.AlphaString())
	gens["LinkedRedisCacheLocation"] = gen.PtrOf(gen.AlphaString())
	gens["ServerRole"] = gen.PtrOf(gen.OneConstOf(RedisLinkedServerCreateProperties_ServerRole_Primary, RedisLinkedServerCreateProperties_ServerRole_Secondary))
}

func Test_Redis_LinkedServer_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis_LinkedServer_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis_LinkedServer_Spec_ARM, Redis_LinkedServer_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis_LinkedServer_Spec_ARM runs a test to see if a specific instance of Redis_LinkedServer_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis_LinkedServer_Spec_ARM(subject Redis_LinkedServer_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis_LinkedServer_Spec_ARM
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

// Generator of Redis_LinkedServer_Spec_ARM instances for property testing - lazily instantiated by
// Redis_LinkedServer_Spec_ARMGenerator()
var redis_LinkedServer_Spec_ARMGenerator gopter.Gen

// Redis_LinkedServer_Spec_ARMGenerator returns a generator of Redis_LinkedServer_Spec_ARM instances for property testing.
// We first initialize redis_LinkedServer_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Redis_LinkedServer_Spec_ARMGenerator() gopter.Gen {
	if redis_LinkedServer_Spec_ARMGenerator != nil {
		return redis_LinkedServer_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_LinkedServer_Spec_ARM(generators)
	redis_LinkedServer_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Redis_LinkedServer_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_LinkedServer_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForRedis_LinkedServer_Spec_ARM(generators)
	redis_LinkedServer_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Redis_LinkedServer_Spec_ARM{}), generators)

	return redis_LinkedServer_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRedis_LinkedServer_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedis_LinkedServer_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForRedis_LinkedServer_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedis_LinkedServer_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RedisLinkedServerCreateProperties_ARMGenerator())
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

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

func Test_RedisLinkedServerWithProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServerWithProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServerWithProperties_STATUSARM, RedisLinkedServerWithProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServerWithProperties_STATUSARM runs a test to see if a specific instance of RedisLinkedServerWithProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServerWithProperties_STATUSARM(subject RedisLinkedServerWithProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServerWithProperties_STATUSARM
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

// Generator of RedisLinkedServerWithProperties_STATUSARM instances for property testing - lazily instantiated by
// RedisLinkedServerWithProperties_STATUSARMGenerator()
var redisLinkedServerWithProperties_STATUSARMGenerator gopter.Gen

// RedisLinkedServerWithProperties_STATUSARMGenerator returns a generator of RedisLinkedServerWithProperties_STATUSARM instances for property testing.
// We first initialize redisLinkedServerWithProperties_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisLinkedServerWithProperties_STATUSARMGenerator() gopter.Gen {
	if redisLinkedServerWithProperties_STATUSARMGenerator != nil {
		return redisLinkedServerWithProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServerWithProperties_STATUSARM(generators)
	redisLinkedServerWithProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServerWithProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServerWithProperties_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForRedisLinkedServerWithProperties_STATUSARM(generators)
	redisLinkedServerWithProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServerWithProperties_STATUSARM{}), generators)

	return redisLinkedServerWithProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServerWithProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServerWithProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisLinkedServerWithProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisLinkedServerWithProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RedisLinkedServerProperties_STATUSARMGenerator())
}

func Test_RedisLinkedServerProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServerProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServerProperties_STATUSARM, RedisLinkedServerProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServerProperties_STATUSARM runs a test to see if a specific instance of RedisLinkedServerProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServerProperties_STATUSARM(subject RedisLinkedServerProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServerProperties_STATUSARM
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

// Generator of RedisLinkedServerProperties_STATUSARM instances for property testing - lazily instantiated by
// RedisLinkedServerProperties_STATUSARMGenerator()
var redisLinkedServerProperties_STATUSARMGenerator gopter.Gen

// RedisLinkedServerProperties_STATUSARMGenerator returns a generator of RedisLinkedServerProperties_STATUSARM instances for property testing.
func RedisLinkedServerProperties_STATUSARMGenerator() gopter.Gen {
	if redisLinkedServerProperties_STATUSARMGenerator != nil {
		return redisLinkedServerProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServerProperties_STATUSARM(generators)
	redisLinkedServerProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServerProperties_STATUSARM{}), generators)

	return redisLinkedServerProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServerProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServerProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["LinkedRedisCacheId"] = gen.PtrOf(gen.AlphaString())
	gens["LinkedRedisCacheLocation"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
<<<<<<< HEAD
	gens["ServerRole"] = gen.PtrOf(gen.OneConstOf(RedisLinkedServerProperties_ServerRole_Primary_STATUS, RedisLinkedServerProperties_ServerRole_Secondary_STATUS))
=======
	gens["ServerRole"] = gen.PtrOf(gen.OneConstOf(RedisLinkedServerProperties_STATUS_ServerRole_Primary, RedisLinkedServerProperties_STATUS_ServerRole_Secondary))
>>>>>>> main
}

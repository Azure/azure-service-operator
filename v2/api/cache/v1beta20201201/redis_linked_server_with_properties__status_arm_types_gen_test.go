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

func Test_RedisLinkedServerWithProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServerWithProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServerWithPropertiesStatusARM, RedisLinkedServerWithPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServerWithPropertiesStatusARM runs a test to see if a specific instance of RedisLinkedServerWithProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServerWithPropertiesStatusARM(subject RedisLinkedServerWithProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServerWithProperties_StatusARM
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

// Generator of RedisLinkedServerWithProperties_StatusARM instances for property testing - lazily instantiated by
// RedisLinkedServerWithPropertiesStatusARMGenerator()
var redisLinkedServerWithPropertiesStatusARMGenerator gopter.Gen

// RedisLinkedServerWithPropertiesStatusARMGenerator returns a generator of RedisLinkedServerWithProperties_StatusARM instances for property testing.
// We first initialize redisLinkedServerWithPropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisLinkedServerWithPropertiesStatusARMGenerator() gopter.Gen {
	if redisLinkedServerWithPropertiesStatusARMGenerator != nil {
		return redisLinkedServerWithPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServerWithPropertiesStatusARM(generators)
	redisLinkedServerWithPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServerWithProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServerWithPropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForRedisLinkedServerWithPropertiesStatusARM(generators)
	redisLinkedServerWithPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServerWithProperties_StatusARM{}), generators)

	return redisLinkedServerWithPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServerWithPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServerWithPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisLinkedServerWithPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisLinkedServerWithPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RedisLinkedServerPropertiesStatusARMGenerator())
}

func Test_RedisLinkedServerProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServerProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServerPropertiesStatusARM, RedisLinkedServerPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServerPropertiesStatusARM runs a test to see if a specific instance of RedisLinkedServerProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServerPropertiesStatusARM(subject RedisLinkedServerProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServerProperties_StatusARM
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

// Generator of RedisLinkedServerProperties_StatusARM instances for property testing - lazily instantiated by
// RedisLinkedServerPropertiesStatusARMGenerator()
var redisLinkedServerPropertiesStatusARMGenerator gopter.Gen

// RedisLinkedServerPropertiesStatusARMGenerator returns a generator of RedisLinkedServerProperties_StatusARM instances for property testing.
func RedisLinkedServerPropertiesStatusARMGenerator() gopter.Gen {
	if redisLinkedServerPropertiesStatusARMGenerator != nil {
		return redisLinkedServerPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServerPropertiesStatusARM(generators)
	redisLinkedServerPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServerProperties_StatusARM{}), generators)

	return redisLinkedServerPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServerPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServerPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["LinkedRedisCacheId"] = gen.PtrOf(gen.AlphaString())
	gens["LinkedRedisCacheLocation"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ServerRole"] = gen.PtrOf(gen.OneConstOf(RedisLinkedServerPropertiesStatusServerRolePrimary, RedisLinkedServerPropertiesStatusServerRoleSecondary))
}

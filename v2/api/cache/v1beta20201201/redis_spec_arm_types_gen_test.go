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

func Test_Redis_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis_Spec_ARM, Redis_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis_Spec_ARM runs a test to see if a specific instance of Redis_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis_Spec_ARM(subject Redis_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis_Spec_ARM
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

// Generator of Redis_Spec_ARM instances for property testing - lazily instantiated by Redis_Spec_ARMGenerator()
var redis_Spec_ARMGenerator gopter.Gen

// Redis_Spec_ARMGenerator returns a generator of Redis_Spec_ARM instances for property testing.
// We first initialize redis_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Redis_Spec_ARMGenerator() gopter.Gen {
	if redis_Spec_ARMGenerator != nil {
		return redis_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_Spec_ARM(generators)
	redis_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Redis_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForRedis_Spec_ARM(generators)
	redis_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Redis_Spec_ARM{}), generators)

	return redis_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRedis_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedis_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedis_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedis_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RedisCreateProperties_ARMGenerator())
}

func Test_RedisCreateProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisCreateProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisCreateProperties_ARM, RedisCreateProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisCreateProperties_ARM runs a test to see if a specific instance of RedisCreateProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisCreateProperties_ARM(subject RedisCreateProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisCreateProperties_ARM
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

// Generator of RedisCreateProperties_ARM instances for property testing - lazily instantiated by
// RedisCreateProperties_ARMGenerator()
var redisCreateProperties_ARMGenerator gopter.Gen

// RedisCreateProperties_ARMGenerator returns a generator of RedisCreateProperties_ARM instances for property testing.
// We first initialize redisCreateProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisCreateProperties_ARMGenerator() gopter.Gen {
	if redisCreateProperties_ARMGenerator != nil {
		return redisCreateProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisCreateProperties_ARM(generators)
	redisCreateProperties_ARMGenerator = gen.Struct(reflect.TypeOf(RedisCreateProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisCreateProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForRedisCreateProperties_ARM(generators)
	redisCreateProperties_ARMGenerator = gen.Struct(reflect.TypeOf(RedisCreateProperties_ARM{}), generators)

	return redisCreateProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisCreateProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisCreateProperties_ARM(gens map[string]gopter.Gen) {
	gens["EnableNonSslPort"] = gen.PtrOf(gen.Bool())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.OneConstOf(RedisCreateProperties_MinimumTlsVersion_10, RedisCreateProperties_MinimumTlsVersion_11, RedisCreateProperties_MinimumTlsVersion_12))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(RedisCreateProperties_PublicNetworkAccess_Disabled, RedisCreateProperties_PublicNetworkAccess_Enabled))
	gens["RedisConfiguration"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["RedisVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ReplicasPerMaster"] = gen.PtrOf(gen.Int())
	gens["ReplicasPerPrimary"] = gen.PtrOf(gen.Int())
	gens["ShardCount"] = gen.PtrOf(gen.Int())
	gens["StaticIP"] = gen.PtrOf(gen.AlphaString())
	gens["SubnetId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantSettings"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisCreateProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisCreateProperties_ARM(gens map[string]gopter.Gen) {
	gens["Sku"] = gen.PtrOf(Sku_ARMGenerator())
}

func Test_Sku_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_ARM, Sku_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_ARM runs a test to see if a specific instance of Sku_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_ARM(subject Sku_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_ARM
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

// Generator of Sku_ARM instances for property testing - lazily instantiated by Sku_ARMGenerator()
var sku_ARMGenerator gopter.Gen

// Sku_ARMGenerator returns a generator of Sku_ARM instances for property testing.
func Sku_ARMGenerator() gopter.Gen {
	if sku_ARMGenerator != nil {
		return sku_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_ARM(generators)
	sku_ARMGenerator = gen.Struct(reflect.TypeOf(Sku_ARM{}), generators)

	return sku_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSku_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_ARM(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Family"] = gen.PtrOf(gen.OneConstOf(Sku_Family_C, Sku_Family_P))
	gens["Name"] = gen.PtrOf(gen.OneConstOf(Sku_Name_Basic, Sku_Name_Premium, Sku_Name_Standard))
}

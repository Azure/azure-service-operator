// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230401

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

func Test_RedisFirewallRuleProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisFirewallRuleProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisFirewallRuleProperties_STATUS_ARM, RedisFirewallRuleProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisFirewallRuleProperties_STATUS_ARM runs a test to see if a specific instance of RedisFirewallRuleProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisFirewallRuleProperties_STATUS_ARM(subject RedisFirewallRuleProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisFirewallRuleProperties_STATUS_ARM
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

// Generator of RedisFirewallRuleProperties_STATUS_ARM instances for property testing - lazily instantiated by
// RedisFirewallRuleProperties_STATUS_ARMGenerator()
var redisFirewallRuleProperties_STATUS_ARMGenerator gopter.Gen

// RedisFirewallRuleProperties_STATUS_ARMGenerator returns a generator of RedisFirewallRuleProperties_STATUS_ARM instances for property testing.
func RedisFirewallRuleProperties_STATUS_ARMGenerator() gopter.Gen {
	if redisFirewallRuleProperties_STATUS_ARMGenerator != nil {
		return redisFirewallRuleProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisFirewallRuleProperties_STATUS_ARM(generators)
	redisFirewallRuleProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RedisFirewallRuleProperties_STATUS_ARM{}), generators)

	return redisFirewallRuleProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisFirewallRuleProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisFirewallRuleProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["EndIP"] = gen.PtrOf(gen.AlphaString())
	gens["StartIP"] = gen.PtrOf(gen.AlphaString())
}

func Test_Redis_FirewallRule_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis_FirewallRule_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis_FirewallRule_STATUS_ARM, Redis_FirewallRule_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis_FirewallRule_STATUS_ARM runs a test to see if a specific instance of Redis_FirewallRule_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis_FirewallRule_STATUS_ARM(subject Redis_FirewallRule_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis_FirewallRule_STATUS_ARM
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

// Generator of Redis_FirewallRule_STATUS_ARM instances for property testing - lazily instantiated by
// Redis_FirewallRule_STATUS_ARMGenerator()
var redis_FirewallRule_STATUS_ARMGenerator gopter.Gen

// Redis_FirewallRule_STATUS_ARMGenerator returns a generator of Redis_FirewallRule_STATUS_ARM instances for property testing.
// We first initialize redis_FirewallRule_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Redis_FirewallRule_STATUS_ARMGenerator() gopter.Gen {
	if redis_FirewallRule_STATUS_ARMGenerator != nil {
		return redis_FirewallRule_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_FirewallRule_STATUS_ARM(generators)
	redis_FirewallRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Redis_FirewallRule_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_FirewallRule_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForRedis_FirewallRule_STATUS_ARM(generators)
	redis_FirewallRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Redis_FirewallRule_STATUS_ARM{}), generators)

	return redis_FirewallRule_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRedis_FirewallRule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedis_FirewallRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedis_FirewallRule_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedis_FirewallRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RedisFirewallRuleProperties_STATUS_ARMGenerator())
}

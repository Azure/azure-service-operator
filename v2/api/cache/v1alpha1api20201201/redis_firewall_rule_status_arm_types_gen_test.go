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

func Test_RedisFirewallRule_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisFirewallRule_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisFirewallRule_STATUSARM, RedisFirewallRule_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisFirewallRule_STATUSARM runs a test to see if a specific instance of RedisFirewallRule_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisFirewallRule_STATUSARM(subject RedisFirewallRule_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisFirewallRule_STATUSARM
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

// Generator of RedisFirewallRule_STATUSARM instances for property testing - lazily instantiated by
// RedisFirewallRule_STATUSARMGenerator()
var redisFirewallRule_STATUSARMGenerator gopter.Gen

// RedisFirewallRule_STATUSARMGenerator returns a generator of RedisFirewallRule_STATUSARM instances for property testing.
// We first initialize redisFirewallRule_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisFirewallRule_STATUSARMGenerator() gopter.Gen {
	if redisFirewallRule_STATUSARMGenerator != nil {
		return redisFirewallRule_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisFirewallRule_STATUSARM(generators)
	redisFirewallRule_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisFirewallRule_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisFirewallRule_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForRedisFirewallRule_STATUSARM(generators)
	redisFirewallRule_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisFirewallRule_STATUSARM{}), generators)

	return redisFirewallRule_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisFirewallRule_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisFirewallRule_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisFirewallRule_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisFirewallRule_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RedisFirewallRuleProperties_STATUSARMGenerator())
}

func Test_RedisFirewallRuleProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisFirewallRuleProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisFirewallRuleProperties_STATUSARM, RedisFirewallRuleProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisFirewallRuleProperties_STATUSARM runs a test to see if a specific instance of RedisFirewallRuleProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisFirewallRuleProperties_STATUSARM(subject RedisFirewallRuleProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisFirewallRuleProperties_STATUSARM
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

// Generator of RedisFirewallRuleProperties_STATUSARM instances for property testing - lazily instantiated by
// RedisFirewallRuleProperties_STATUSARMGenerator()
var redisFirewallRuleProperties_STATUSARMGenerator gopter.Gen

// RedisFirewallRuleProperties_STATUSARMGenerator returns a generator of RedisFirewallRuleProperties_STATUSARM instances for property testing.
func RedisFirewallRuleProperties_STATUSARMGenerator() gopter.Gen {
	if redisFirewallRuleProperties_STATUSARMGenerator != nil {
		return redisFirewallRuleProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisFirewallRuleProperties_STATUSARM(generators)
	redisFirewallRuleProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RedisFirewallRuleProperties_STATUSARM{}), generators)

	return redisFirewallRuleProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisFirewallRuleProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisFirewallRuleProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["EndIP"] = gen.PtrOf(gen.AlphaString())
	gens["StartIP"] = gen.PtrOf(gen.AlphaString())
}

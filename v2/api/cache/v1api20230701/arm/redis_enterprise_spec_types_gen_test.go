// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_ClusterProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ClusterProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForClusterProperties, ClusterPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForClusterProperties runs a test to see if a specific instance of ClusterProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForClusterProperties(subject ClusterProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ClusterProperties
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

// Generator of ClusterProperties instances for property testing - lazily instantiated by ClusterPropertiesGenerator()
var clusterPropertiesGenerator gopter.Gen

// ClusterPropertiesGenerator returns a generator of ClusterProperties instances for property testing.
func ClusterPropertiesGenerator() gopter.Gen {
	if clusterPropertiesGenerator != nil {
		return clusterPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForClusterProperties(generators)
	clusterPropertiesGenerator = gen.Struct(reflect.TypeOf(ClusterProperties{}), generators)

	return clusterPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForClusterProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForClusterProperties(gens map[string]gopter.Gen) {
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.OneConstOf(ClusterProperties_MinimumTlsVersion_10, ClusterProperties_MinimumTlsVersion_11, ClusterProperties_MinimumTlsVersion_12))
}

func Test_RedisEnterprise_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterprise_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterprise_Spec, RedisEnterprise_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterprise_Spec runs a test to see if a specific instance of RedisEnterprise_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterprise_Spec(subject RedisEnterprise_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterprise_Spec
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

// Generator of RedisEnterprise_Spec instances for property testing - lazily instantiated by
// RedisEnterprise_SpecGenerator()
var redisEnterprise_SpecGenerator gopter.Gen

// RedisEnterprise_SpecGenerator returns a generator of RedisEnterprise_Spec instances for property testing.
// We first initialize redisEnterprise_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisEnterprise_SpecGenerator() gopter.Gen {
	if redisEnterprise_SpecGenerator != nil {
		return redisEnterprise_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterprise_Spec(generators)
	redisEnterprise_SpecGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterprise_Spec(generators)
	AddRelatedPropertyGeneratorsForRedisEnterprise_Spec(generators)
	redisEnterprise_SpecGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise_Spec{}), generators)

	return redisEnterprise_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRedisEnterprise_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisEnterprise_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisEnterprise_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterprise_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ClusterPropertiesGenerator())
	gens["Sku"] = gen.PtrOf(SkuGenerator())
}

func Test_Sku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku, SkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku runs a test to see if a specific instance of Sku round trips to JSON and back losslessly
func RunJSONSerializationTestForSku(subject Sku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku
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

// Generator of Sku instances for property testing - lazily instantiated by SkuGenerator()
var skuGenerator gopter.Gen

// SkuGenerator returns a generator of Sku instances for property testing.
func SkuGenerator() gopter.Gen {
	if skuGenerator != nil {
		return skuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku(generators)
	skuGenerator = gen.Struct(reflect.TypeOf(Sku{}), generators)

	return skuGenerator
}

// AddIndependentPropertyGeneratorsForSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.OneConstOf(
		Sku_Name_EnterpriseFlash_F1500,
		Sku_Name_EnterpriseFlash_F300,
		Sku_Name_EnterpriseFlash_F700,
		Sku_Name_Enterprise_E10,
		Sku_Name_Enterprise_E100,
		Sku_Name_Enterprise_E20,
		Sku_Name_Enterprise_E50))
}

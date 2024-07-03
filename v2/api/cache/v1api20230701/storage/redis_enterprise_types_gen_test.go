// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

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

func Test_PrivateEndpointConnection_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_STATUS, PrivateEndpointConnection_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_STATUS runs a test to see if a specific instance of PrivateEndpointConnection_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_STATUS(subject PrivateEndpointConnection_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_STATUS
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

// Generator of PrivateEndpointConnection_STATUS instances for property testing - lazily instantiated by
// PrivateEndpointConnection_STATUSGenerator()
var privateEndpointConnection_STATUSGenerator gopter.Gen

// PrivateEndpointConnection_STATUSGenerator returns a generator of PrivateEndpointConnection_STATUS instances for property testing.
func PrivateEndpointConnection_STATUSGenerator() gopter.Gen {
	if privateEndpointConnection_STATUSGenerator != nil {
		return privateEndpointConnection_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS(generators)
	privateEndpointConnection_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS{}), generators)

	return privateEndpointConnection_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisEnterprise_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterprise via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterprise, RedisEnterpriseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterprise runs a test to see if a specific instance of RedisEnterprise round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterprise(subject RedisEnterprise) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterprise
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

// Generator of RedisEnterprise instances for property testing - lazily instantiated by RedisEnterpriseGenerator()
var redisEnterpriseGenerator gopter.Gen

// RedisEnterpriseGenerator returns a generator of RedisEnterprise instances for property testing.
func RedisEnterpriseGenerator() gopter.Gen {
	if redisEnterpriseGenerator != nil {
		return redisEnterpriseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedisEnterprise(generators)
	redisEnterpriseGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise{}), generators)

	return redisEnterpriseGenerator
}

// AddRelatedPropertyGeneratorsForRedisEnterprise is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterprise(gens map[string]gopter.Gen) {
	gens["Spec"] = RedisEnterprise_SpecGenerator()
	gens["Status"] = RedisEnterprise_STATUSGenerator()
}

func Test_RedisEnterprise_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterprise_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterprise_STATUS, RedisEnterprise_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterprise_STATUS runs a test to see if a specific instance of RedisEnterprise_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterprise_STATUS(subject RedisEnterprise_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterprise_STATUS
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

// Generator of RedisEnterprise_STATUS instances for property testing - lazily instantiated by
// RedisEnterprise_STATUSGenerator()
var redisEnterprise_STATUSGenerator gopter.Gen

// RedisEnterprise_STATUSGenerator returns a generator of RedisEnterprise_STATUS instances for property testing.
// We first initialize redisEnterprise_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisEnterprise_STATUSGenerator() gopter.Gen {
	if redisEnterprise_STATUSGenerator != nil {
		return redisEnterprise_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterprise_STATUS(generators)
	redisEnterprise_STATUSGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterprise_STATUS(generators)
	AddRelatedPropertyGeneratorsForRedisEnterprise_STATUS(generators)
	redisEnterprise_STATUSGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise_STATUS{}), generators)

	return redisEnterprise_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedisEnterprise_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisEnterprise_STATUS(gens map[string]gopter.Gen) {
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["RedisVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceState"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisEnterprise_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterprise_STATUS(gens map[string]gopter.Gen) {
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_STATUSGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSGenerator())
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
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisEnterprise_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterprise_Spec(gens map[string]gopter.Gen) {
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
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_Sku_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUS, Sku_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUS runs a test to see if a specific instance of Sku_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUS(subject Sku_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUS
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

// Generator of Sku_STATUS instances for property testing - lazily instantiated by Sku_STATUSGenerator()
var sku_STATUSGenerator gopter.Gen

// Sku_STATUSGenerator returns a generator of Sku_STATUS instances for property testing.
func Sku_STATUSGenerator() gopter.Gen {
	if sku_STATUSGenerator != nil {
		return sku_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUS(generators)
	sku_STATUSGenerator = gen.Struct(reflect.TypeOf(Sku_STATUS{}), generators)

	return sku_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUS(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

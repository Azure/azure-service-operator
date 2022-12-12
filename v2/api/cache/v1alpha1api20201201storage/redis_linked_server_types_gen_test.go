// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201storage

import (
	"encoding/json"
	v20201201s "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20201201storage"
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

func Test_RedisLinkedServer_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisLinkedServer to hub returns original",
		prop.ForAll(RunResourceConversionTestForRedisLinkedServer, RedisLinkedServerGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForRedisLinkedServer tests if a specific instance of RedisLinkedServer round trips to the hub storage version and back losslessly
func RunResourceConversionTestForRedisLinkedServer(subject RedisLinkedServer) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20201201s.RedisLinkedServer
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual RedisLinkedServer
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RedisLinkedServer_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisLinkedServer to RedisLinkedServer via AssignProperties_To_RedisLinkedServer & AssignProperties_From_RedisLinkedServer returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisLinkedServer, RedisLinkedServerGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisLinkedServer tests if a specific instance of RedisLinkedServer can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForRedisLinkedServer(subject RedisLinkedServer) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.RedisLinkedServer
	err := copied.AssignProperties_To_RedisLinkedServer(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisLinkedServer
	err = actual.AssignProperties_From_RedisLinkedServer(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RedisLinkedServer_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServer via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServer, RedisLinkedServerGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServer runs a test to see if a specific instance of RedisLinkedServer round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServer(subject RedisLinkedServer) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServer
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

// Generator of RedisLinkedServer instances for property testing - lazily instantiated by RedisLinkedServerGenerator()
var redisLinkedServerGenerator gopter.Gen

// RedisLinkedServerGenerator returns a generator of RedisLinkedServer instances for property testing.
func RedisLinkedServerGenerator() gopter.Gen {
	if redisLinkedServerGenerator != nil {
		return redisLinkedServerGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedisLinkedServer(generators)
	redisLinkedServerGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServer{}), generators)

	return redisLinkedServerGenerator
}

// AddRelatedPropertyGeneratorsForRedisLinkedServer is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisLinkedServer(gens map[string]gopter.Gen) {
	gens["Spec"] = Redis_LinkedServer_SpecGenerator()
	gens["Status"] = Redis_LinkedServer_STATUSGenerator()
}

func Test_Redis_LinkedServer_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Redis_LinkedServer_Spec to Redis_LinkedServer_Spec via AssignProperties_To_Redis_LinkedServer_Spec & AssignProperties_From_Redis_LinkedServer_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedis_LinkedServer_Spec, Redis_LinkedServer_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedis_LinkedServer_Spec tests if a specific instance of Redis_LinkedServer_Spec can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForRedis_LinkedServer_Spec(subject Redis_LinkedServer_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.Redis_LinkedServer_Spec
	err := copied.AssignProperties_To_Redis_LinkedServer_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Redis_LinkedServer_Spec
	err = actual.AssignProperties_From_Redis_LinkedServer_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Redis_LinkedServer_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis_LinkedServer_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis_LinkedServer_Spec, Redis_LinkedServer_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis_LinkedServer_Spec runs a test to see if a specific instance of Redis_LinkedServer_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis_LinkedServer_Spec(subject Redis_LinkedServer_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis_LinkedServer_Spec
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

// Generator of Redis_LinkedServer_Spec instances for property testing - lazily instantiated by
// Redis_LinkedServer_SpecGenerator()
var redis_LinkedServer_SpecGenerator gopter.Gen

// Redis_LinkedServer_SpecGenerator returns a generator of Redis_LinkedServer_Spec instances for property testing.
func Redis_LinkedServer_SpecGenerator() gopter.Gen {
	if redis_LinkedServer_SpecGenerator != nil {
		return redis_LinkedServer_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_LinkedServer_Spec(generators)
	redis_LinkedServer_SpecGenerator = gen.Struct(reflect.TypeOf(Redis_LinkedServer_Spec{}), generators)

	return redis_LinkedServer_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRedis_LinkedServer_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedis_LinkedServer_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["LinkedRedisCacheLocation"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["ServerRole"] = gen.PtrOf(gen.AlphaString())
}

func Test_Redis_LinkedServer_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Redis_LinkedServer_STATUS to Redis_LinkedServer_STATUS via AssignProperties_To_Redis_LinkedServer_STATUS & AssignProperties_From_Redis_LinkedServer_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedis_LinkedServer_STATUS, Redis_LinkedServer_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedis_LinkedServer_STATUS tests if a specific instance of Redis_LinkedServer_STATUS can be assigned to v1beta20201201storage and back losslessly
func RunPropertyAssignmentTestForRedis_LinkedServer_STATUS(subject Redis_LinkedServer_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201201s.Redis_LinkedServer_STATUS
	err := copied.AssignProperties_To_Redis_LinkedServer_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Redis_LinkedServer_STATUS
	err = actual.AssignProperties_From_Redis_LinkedServer_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Redis_LinkedServer_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis_LinkedServer_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis_LinkedServer_STATUS, Redis_LinkedServer_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis_LinkedServer_STATUS runs a test to see if a specific instance of Redis_LinkedServer_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis_LinkedServer_STATUS(subject Redis_LinkedServer_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis_LinkedServer_STATUS
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

// Generator of Redis_LinkedServer_STATUS instances for property testing - lazily instantiated by
// Redis_LinkedServer_STATUSGenerator()
var redis_LinkedServer_STATUSGenerator gopter.Gen

// Redis_LinkedServer_STATUSGenerator returns a generator of Redis_LinkedServer_STATUS instances for property testing.
func Redis_LinkedServer_STATUSGenerator() gopter.Gen {
	if redis_LinkedServer_STATUSGenerator != nil {
		return redis_LinkedServer_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_LinkedServer_STATUS(generators)
	redis_LinkedServer_STATUSGenerator = gen.Struct(reflect.TypeOf(Redis_LinkedServer_STATUS{}), generators)

	return redis_LinkedServer_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedis_LinkedServer_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedis_LinkedServer_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["LinkedRedisCacheId"] = gen.PtrOf(gen.AlphaString())
	gens["LinkedRedisCacheLocation"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ServerRole"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

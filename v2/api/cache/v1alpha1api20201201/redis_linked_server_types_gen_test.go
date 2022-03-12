// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/cache/v1alpha1api20201201storage"
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
	var hub v1alpha1api20201201storage.RedisLinkedServer
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
		"Round trip from RedisLinkedServer to RedisLinkedServer via AssignPropertiesToRedisLinkedServer & AssignPropertiesFromRedisLinkedServer returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisLinkedServer, RedisLinkedServerGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisLinkedServer tests if a specific instance of RedisLinkedServer can be assigned to v1alpha1api20201201storage and back losslessly
func RunPropertyAssignmentTestForRedisLinkedServer(subject RedisLinkedServer) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20201201storage.RedisLinkedServer
	err := copied.AssignPropertiesToRedisLinkedServer(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisLinkedServer
	err = actual.AssignPropertiesFromRedisLinkedServer(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
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
	parameters.MaxSize = 10
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
	gens["Spec"] = RedisLinkedServersSpecGenerator()
	gens["Status"] = RedisLinkedServerWithPropertiesStatusGenerator()
}

func Test_RedisLinkedServerWithProperties_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisLinkedServerWithProperties_Status to RedisLinkedServerWithProperties_Status via AssignPropertiesToRedisLinkedServerWithPropertiesStatus & AssignPropertiesFromRedisLinkedServerWithPropertiesStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisLinkedServerWithPropertiesStatus, RedisLinkedServerWithPropertiesStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisLinkedServerWithPropertiesStatus tests if a specific instance of RedisLinkedServerWithProperties_Status can be assigned to v1alpha1api20201201storage and back losslessly
func RunPropertyAssignmentTestForRedisLinkedServerWithPropertiesStatus(subject RedisLinkedServerWithProperties_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20201201storage.RedisLinkedServerWithProperties_Status
	err := copied.AssignPropertiesToRedisLinkedServerWithPropertiesStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisLinkedServerWithProperties_Status
	err = actual.AssignPropertiesFromRedisLinkedServerWithPropertiesStatus(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RedisLinkedServerWithProperties_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServerWithProperties_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServerWithPropertiesStatus, RedisLinkedServerWithPropertiesStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServerWithPropertiesStatus runs a test to see if a specific instance of RedisLinkedServerWithProperties_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServerWithPropertiesStatus(subject RedisLinkedServerWithProperties_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServerWithProperties_Status
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

// Generator of RedisLinkedServerWithProperties_Status instances for property testing - lazily instantiated by
//RedisLinkedServerWithPropertiesStatusGenerator()
var redisLinkedServerWithPropertiesStatusGenerator gopter.Gen

// RedisLinkedServerWithPropertiesStatusGenerator returns a generator of RedisLinkedServerWithProperties_Status instances for property testing.
func RedisLinkedServerWithPropertiesStatusGenerator() gopter.Gen {
	if redisLinkedServerWithPropertiesStatusGenerator != nil {
		return redisLinkedServerWithPropertiesStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServerWithPropertiesStatus(generators)
	redisLinkedServerWithPropertiesStatusGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServerWithProperties_Status{}), generators)

	return redisLinkedServerWithPropertiesStatusGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServerWithPropertiesStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServerWithPropertiesStatus(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["LinkedRedisCacheId"] = gen.PtrOf(gen.AlphaString())
	gens["LinkedRedisCacheLocation"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ServerRole"] = gen.PtrOf(gen.OneConstOf(RedisLinkedServerPropertiesStatusServerRolePrimary, RedisLinkedServerPropertiesStatusServerRoleSecondary))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisLinkedServers_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisLinkedServers_Spec to RedisLinkedServers_Spec via AssignPropertiesToRedisLinkedServersSpec & AssignPropertiesFromRedisLinkedServersSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisLinkedServersSpec, RedisLinkedServersSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisLinkedServersSpec tests if a specific instance of RedisLinkedServers_Spec can be assigned to v1alpha1api20201201storage and back losslessly
func RunPropertyAssignmentTestForRedisLinkedServersSpec(subject RedisLinkedServers_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20201201storage.RedisLinkedServers_Spec
	err := copied.AssignPropertiesToRedisLinkedServersSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisLinkedServers_Spec
	err = actual.AssignPropertiesFromRedisLinkedServersSpec(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RedisLinkedServers_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServers_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServersSpec, RedisLinkedServersSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServersSpec runs a test to see if a specific instance of RedisLinkedServers_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServersSpec(subject RedisLinkedServers_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServers_Spec
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

// Generator of RedisLinkedServers_Spec instances for property testing - lazily instantiated by
//RedisLinkedServersSpecGenerator()
var redisLinkedServersSpecGenerator gopter.Gen

// RedisLinkedServersSpecGenerator returns a generator of RedisLinkedServers_Spec instances for property testing.
func RedisLinkedServersSpecGenerator() gopter.Gen {
	if redisLinkedServersSpecGenerator != nil {
		return redisLinkedServersSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServersSpec(generators)
	redisLinkedServersSpecGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServers_Spec{}), generators)

	return redisLinkedServersSpecGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServersSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServersSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["LinkedRedisCacheLocation"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["ServerRole"] = gen.PtrOf(gen.OneConstOf(RedisLinkedServerCreatePropertiesServerRolePrimary, RedisLinkedServerCreatePropertiesServerRoleSecondary))
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

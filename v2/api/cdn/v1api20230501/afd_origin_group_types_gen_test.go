// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/storage"
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

func Test_AfdOriginGroup_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AfdOriginGroup to hub returns original",
		prop.ForAll(RunResourceConversionTestForAfdOriginGroup, AfdOriginGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForAfdOriginGroup tests if a specific instance of AfdOriginGroup round trips to the hub storage version and back losslessly
func RunResourceConversionTestForAfdOriginGroup(subject AfdOriginGroup) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.AfdOriginGroup
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual AfdOriginGroup
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

func Test_AfdOriginGroup_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AfdOriginGroup to AfdOriginGroup via AssignProperties_To_AfdOriginGroup & AssignProperties_From_AfdOriginGroup returns original",
		prop.ForAll(RunPropertyAssignmentTestForAfdOriginGroup, AfdOriginGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAfdOriginGroup tests if a specific instance of AfdOriginGroup can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAfdOriginGroup(subject AfdOriginGroup) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.AfdOriginGroup
	err := copied.AssignProperties_To_AfdOriginGroup(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AfdOriginGroup
	err = actual.AssignProperties_From_AfdOriginGroup(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AfdOriginGroup_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AfdOriginGroup via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAfdOriginGroup, AfdOriginGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAfdOriginGroup runs a test to see if a specific instance of AfdOriginGroup round trips to JSON and back losslessly
func RunJSONSerializationTestForAfdOriginGroup(subject AfdOriginGroup) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AfdOriginGroup
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

// Generator of AfdOriginGroup instances for property testing - lazily instantiated by AfdOriginGroupGenerator()
var afdOriginGroupGenerator gopter.Gen

// AfdOriginGroupGenerator returns a generator of AfdOriginGroup instances for property testing.
func AfdOriginGroupGenerator() gopter.Gen {
	if afdOriginGroupGenerator != nil {
		return afdOriginGroupGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAfdOriginGroup(generators)
	afdOriginGroupGenerator = gen.Struct(reflect.TypeOf(AfdOriginGroup{}), generators)

	return afdOriginGroupGenerator
}

// AddRelatedPropertyGeneratorsForAfdOriginGroup is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAfdOriginGroup(gens map[string]gopter.Gen) {
	gens["Spec"] = Profiles_OriginGroup_SpecGenerator()
	gens["Status"] = Profiles_OriginGroup_STATUSGenerator()
}

func Test_HealthProbeParameters_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from HealthProbeParameters to HealthProbeParameters via AssignProperties_To_HealthProbeParameters & AssignProperties_From_HealthProbeParameters returns original",
		prop.ForAll(RunPropertyAssignmentTestForHealthProbeParameters, HealthProbeParametersGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForHealthProbeParameters tests if a specific instance of HealthProbeParameters can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForHealthProbeParameters(subject HealthProbeParameters) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.HealthProbeParameters
	err := copied.AssignProperties_To_HealthProbeParameters(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual HealthProbeParameters
	err = actual.AssignProperties_From_HealthProbeParameters(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_HealthProbeParameters_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HealthProbeParameters via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHealthProbeParameters, HealthProbeParametersGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHealthProbeParameters runs a test to see if a specific instance of HealthProbeParameters round trips to JSON and back losslessly
func RunJSONSerializationTestForHealthProbeParameters(subject HealthProbeParameters) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HealthProbeParameters
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

// Generator of HealthProbeParameters instances for property testing - lazily instantiated by
// HealthProbeParametersGenerator()
var healthProbeParametersGenerator gopter.Gen

// HealthProbeParametersGenerator returns a generator of HealthProbeParameters instances for property testing.
func HealthProbeParametersGenerator() gopter.Gen {
	if healthProbeParametersGenerator != nil {
		return healthProbeParametersGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHealthProbeParameters(generators)
	healthProbeParametersGenerator = gen.Struct(reflect.TypeOf(HealthProbeParameters{}), generators)

	return healthProbeParametersGenerator
}

// AddIndependentPropertyGeneratorsForHealthProbeParameters is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHealthProbeParameters(gens map[string]gopter.Gen) {
	gens["ProbeIntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["ProbePath"] = gen.PtrOf(gen.AlphaString())
	gens["ProbeProtocol"] = gen.PtrOf(gen.OneConstOf(HealthProbeParameters_ProbeProtocol_Http, HealthProbeParameters_ProbeProtocol_Https, HealthProbeParameters_ProbeProtocol_NotSet))
	gens["ProbeRequestType"] = gen.PtrOf(gen.OneConstOf(HealthProbeParameters_ProbeRequestType_GET, HealthProbeParameters_ProbeRequestType_HEAD, HealthProbeParameters_ProbeRequestType_NotSet))
}

func Test_HealthProbeParameters_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from HealthProbeParameters_STATUS to HealthProbeParameters_STATUS via AssignProperties_To_HealthProbeParameters_STATUS & AssignProperties_From_HealthProbeParameters_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForHealthProbeParameters_STATUS, HealthProbeParameters_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForHealthProbeParameters_STATUS tests if a specific instance of HealthProbeParameters_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForHealthProbeParameters_STATUS(subject HealthProbeParameters_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.HealthProbeParameters_STATUS
	err := copied.AssignProperties_To_HealthProbeParameters_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual HealthProbeParameters_STATUS
	err = actual.AssignProperties_From_HealthProbeParameters_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_HealthProbeParameters_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HealthProbeParameters_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHealthProbeParameters_STATUS, HealthProbeParameters_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHealthProbeParameters_STATUS runs a test to see if a specific instance of HealthProbeParameters_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForHealthProbeParameters_STATUS(subject HealthProbeParameters_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HealthProbeParameters_STATUS
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

// Generator of HealthProbeParameters_STATUS instances for property testing - lazily instantiated by
// HealthProbeParameters_STATUSGenerator()
var healthProbeParameters_STATUSGenerator gopter.Gen

// HealthProbeParameters_STATUSGenerator returns a generator of HealthProbeParameters_STATUS instances for property testing.
func HealthProbeParameters_STATUSGenerator() gopter.Gen {
	if healthProbeParameters_STATUSGenerator != nil {
		return healthProbeParameters_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHealthProbeParameters_STATUS(generators)
	healthProbeParameters_STATUSGenerator = gen.Struct(reflect.TypeOf(HealthProbeParameters_STATUS{}), generators)

	return healthProbeParameters_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForHealthProbeParameters_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHealthProbeParameters_STATUS(gens map[string]gopter.Gen) {
	gens["ProbeIntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["ProbePath"] = gen.PtrOf(gen.AlphaString())
	gens["ProbeProtocol"] = gen.PtrOf(gen.OneConstOf(HealthProbeParameters_ProbeProtocol_STATUS_Http, HealthProbeParameters_ProbeProtocol_STATUS_Https, HealthProbeParameters_ProbeProtocol_STATUS_NotSet))
	gens["ProbeRequestType"] = gen.PtrOf(gen.OneConstOf(HealthProbeParameters_ProbeRequestType_STATUS_GET, HealthProbeParameters_ProbeRequestType_STATUS_HEAD, HealthProbeParameters_ProbeRequestType_STATUS_NotSet))
}

func Test_LoadBalancingSettingsParameters_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from LoadBalancingSettingsParameters to LoadBalancingSettingsParameters via AssignProperties_To_LoadBalancingSettingsParameters & AssignProperties_From_LoadBalancingSettingsParameters returns original",
		prop.ForAll(RunPropertyAssignmentTestForLoadBalancingSettingsParameters, LoadBalancingSettingsParametersGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForLoadBalancingSettingsParameters tests if a specific instance of LoadBalancingSettingsParameters can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForLoadBalancingSettingsParameters(subject LoadBalancingSettingsParameters) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.LoadBalancingSettingsParameters
	err := copied.AssignProperties_To_LoadBalancingSettingsParameters(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual LoadBalancingSettingsParameters
	err = actual.AssignProperties_From_LoadBalancingSettingsParameters(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_LoadBalancingSettingsParameters_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LoadBalancingSettingsParameters via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLoadBalancingSettingsParameters, LoadBalancingSettingsParametersGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLoadBalancingSettingsParameters runs a test to see if a specific instance of LoadBalancingSettingsParameters round trips to JSON and back losslessly
func RunJSONSerializationTestForLoadBalancingSettingsParameters(subject LoadBalancingSettingsParameters) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LoadBalancingSettingsParameters
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

// Generator of LoadBalancingSettingsParameters instances for property testing - lazily instantiated by
// LoadBalancingSettingsParametersGenerator()
var loadBalancingSettingsParametersGenerator gopter.Gen

// LoadBalancingSettingsParametersGenerator returns a generator of LoadBalancingSettingsParameters instances for property testing.
func LoadBalancingSettingsParametersGenerator() gopter.Gen {
	if loadBalancingSettingsParametersGenerator != nil {
		return loadBalancingSettingsParametersGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLoadBalancingSettingsParameters(generators)
	loadBalancingSettingsParametersGenerator = gen.Struct(reflect.TypeOf(LoadBalancingSettingsParameters{}), generators)

	return loadBalancingSettingsParametersGenerator
}

// AddIndependentPropertyGeneratorsForLoadBalancingSettingsParameters is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLoadBalancingSettingsParameters(gens map[string]gopter.Gen) {
	gens["AdditionalLatencyInMilliseconds"] = gen.PtrOf(gen.Int())
	gens["SampleSize"] = gen.PtrOf(gen.Int())
	gens["SuccessfulSamplesRequired"] = gen.PtrOf(gen.Int())
}

func Test_LoadBalancingSettingsParameters_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from LoadBalancingSettingsParameters_STATUS to LoadBalancingSettingsParameters_STATUS via AssignProperties_To_LoadBalancingSettingsParameters_STATUS & AssignProperties_From_LoadBalancingSettingsParameters_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForLoadBalancingSettingsParameters_STATUS, LoadBalancingSettingsParameters_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForLoadBalancingSettingsParameters_STATUS tests if a specific instance of LoadBalancingSettingsParameters_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForLoadBalancingSettingsParameters_STATUS(subject LoadBalancingSettingsParameters_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.LoadBalancingSettingsParameters_STATUS
	err := copied.AssignProperties_To_LoadBalancingSettingsParameters_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual LoadBalancingSettingsParameters_STATUS
	err = actual.AssignProperties_From_LoadBalancingSettingsParameters_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_LoadBalancingSettingsParameters_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LoadBalancingSettingsParameters_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLoadBalancingSettingsParameters_STATUS, LoadBalancingSettingsParameters_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLoadBalancingSettingsParameters_STATUS runs a test to see if a specific instance of LoadBalancingSettingsParameters_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForLoadBalancingSettingsParameters_STATUS(subject LoadBalancingSettingsParameters_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LoadBalancingSettingsParameters_STATUS
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

// Generator of LoadBalancingSettingsParameters_STATUS instances for property testing - lazily instantiated by
// LoadBalancingSettingsParameters_STATUSGenerator()
var loadBalancingSettingsParameters_STATUSGenerator gopter.Gen

// LoadBalancingSettingsParameters_STATUSGenerator returns a generator of LoadBalancingSettingsParameters_STATUS instances for property testing.
func LoadBalancingSettingsParameters_STATUSGenerator() gopter.Gen {
	if loadBalancingSettingsParameters_STATUSGenerator != nil {
		return loadBalancingSettingsParameters_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLoadBalancingSettingsParameters_STATUS(generators)
	loadBalancingSettingsParameters_STATUSGenerator = gen.Struct(reflect.TypeOf(LoadBalancingSettingsParameters_STATUS{}), generators)

	return loadBalancingSettingsParameters_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForLoadBalancingSettingsParameters_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLoadBalancingSettingsParameters_STATUS(gens map[string]gopter.Gen) {
	gens["AdditionalLatencyInMilliseconds"] = gen.PtrOf(gen.Int())
	gens["SampleSize"] = gen.PtrOf(gen.Int())
	gens["SuccessfulSamplesRequired"] = gen.PtrOf(gen.Int())
}

func Test_Profiles_OriginGroup_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Profiles_OriginGroup_STATUS to Profiles_OriginGroup_STATUS via AssignProperties_To_Profiles_OriginGroup_STATUS & AssignProperties_From_Profiles_OriginGroup_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForProfiles_OriginGroup_STATUS, Profiles_OriginGroup_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForProfiles_OriginGroup_STATUS tests if a specific instance of Profiles_OriginGroup_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForProfiles_OriginGroup_STATUS(subject Profiles_OriginGroup_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.Profiles_OriginGroup_STATUS
	err := copied.AssignProperties_To_Profiles_OriginGroup_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Profiles_OriginGroup_STATUS
	err = actual.AssignProperties_From_Profiles_OriginGroup_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Profiles_OriginGroup_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profiles_OriginGroup_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfiles_OriginGroup_STATUS, Profiles_OriginGroup_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfiles_OriginGroup_STATUS runs a test to see if a specific instance of Profiles_OriginGroup_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForProfiles_OriginGroup_STATUS(subject Profiles_OriginGroup_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profiles_OriginGroup_STATUS
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

// Generator of Profiles_OriginGroup_STATUS instances for property testing - lazily instantiated by
// Profiles_OriginGroup_STATUSGenerator()
var profiles_OriginGroup_STATUSGenerator gopter.Gen

// Profiles_OriginGroup_STATUSGenerator returns a generator of Profiles_OriginGroup_STATUS instances for property testing.
// We first initialize profiles_OriginGroup_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Profiles_OriginGroup_STATUSGenerator() gopter.Gen {
	if profiles_OriginGroup_STATUSGenerator != nil {
		return profiles_OriginGroup_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_OriginGroup_STATUS(generators)
	profiles_OriginGroup_STATUSGenerator = gen.Struct(reflect.TypeOf(Profiles_OriginGroup_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_OriginGroup_STATUS(generators)
	AddRelatedPropertyGeneratorsForProfiles_OriginGroup_STATUS(generators)
	profiles_OriginGroup_STATUSGenerator = gen.Struct(reflect.TypeOf(Profiles_OriginGroup_STATUS{}), generators)

	return profiles_OriginGroup_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForProfiles_OriginGroup_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfiles_OriginGroup_STATUS(gens map[string]gopter.Gen) {
	gens["DeploymentStatus"] = gen.PtrOf(gen.OneConstOf(
		AFDOriginGroupProperties_DeploymentStatus_STATUS_Failed,
		AFDOriginGroupProperties_DeploymentStatus_STATUS_InProgress,
		AFDOriginGroupProperties_DeploymentStatus_STATUS_NotStarted,
		AFDOriginGroupProperties_DeploymentStatus_STATUS_Succeeded))
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProfileName"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		AFDOriginGroupProperties_ProvisioningState_STATUS_Creating,
		AFDOriginGroupProperties_ProvisioningState_STATUS_Deleting,
		AFDOriginGroupProperties_ProvisioningState_STATUS_Failed,
		AFDOriginGroupProperties_ProvisioningState_STATUS_Succeeded,
		AFDOriginGroupProperties_ProvisioningState_STATUS_Updating))
	gens["SessionAffinityState"] = gen.PtrOf(gen.OneConstOf(AFDOriginGroupProperties_SessionAffinityState_STATUS_Disabled, AFDOriginGroupProperties_SessionAffinityState_STATUS_Enabled))
	gens["TrafficRestorationTimeToHealedOrNewEndpointsInMinutes"] = gen.PtrOf(gen.Int())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProfiles_OriginGroup_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfiles_OriginGroup_STATUS(gens map[string]gopter.Gen) {
	gens["HealthProbeSettings"] = gen.PtrOf(HealthProbeParameters_STATUSGenerator())
	gens["LoadBalancingSettings"] = gen.PtrOf(LoadBalancingSettingsParameters_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_Profiles_OriginGroup_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Profiles_OriginGroup_Spec to Profiles_OriginGroup_Spec via AssignProperties_To_Profiles_OriginGroup_Spec & AssignProperties_From_Profiles_OriginGroup_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForProfiles_OriginGroup_Spec, Profiles_OriginGroup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForProfiles_OriginGroup_Spec tests if a specific instance of Profiles_OriginGroup_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForProfiles_OriginGroup_Spec(subject Profiles_OriginGroup_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.Profiles_OriginGroup_Spec
	err := copied.AssignProperties_To_Profiles_OriginGroup_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Profiles_OriginGroup_Spec
	err = actual.AssignProperties_From_Profiles_OriginGroup_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Profiles_OriginGroup_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profiles_OriginGroup_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfiles_OriginGroup_Spec, Profiles_OriginGroup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfiles_OriginGroup_Spec runs a test to see if a specific instance of Profiles_OriginGroup_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForProfiles_OriginGroup_Spec(subject Profiles_OriginGroup_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profiles_OriginGroup_Spec
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

// Generator of Profiles_OriginGroup_Spec instances for property testing - lazily instantiated by
// Profiles_OriginGroup_SpecGenerator()
var profiles_OriginGroup_SpecGenerator gopter.Gen

// Profiles_OriginGroup_SpecGenerator returns a generator of Profiles_OriginGroup_Spec instances for property testing.
// We first initialize profiles_OriginGroup_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Profiles_OriginGroup_SpecGenerator() gopter.Gen {
	if profiles_OriginGroup_SpecGenerator != nil {
		return profiles_OriginGroup_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_OriginGroup_Spec(generators)
	profiles_OriginGroup_SpecGenerator = gen.Struct(reflect.TypeOf(Profiles_OriginGroup_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_OriginGroup_Spec(generators)
	AddRelatedPropertyGeneratorsForProfiles_OriginGroup_Spec(generators)
	profiles_OriginGroup_SpecGenerator = gen.Struct(reflect.TypeOf(Profiles_OriginGroup_Spec{}), generators)

	return profiles_OriginGroup_SpecGenerator
}

// AddIndependentPropertyGeneratorsForProfiles_OriginGroup_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfiles_OriginGroup_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["SessionAffinityState"] = gen.PtrOf(gen.OneConstOf(AFDOriginGroupProperties_SessionAffinityState_Disabled, AFDOriginGroupProperties_SessionAffinityState_Enabled))
	gens["TrafficRestorationTimeToHealedOrNewEndpointsInMinutes"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForProfiles_OriginGroup_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfiles_OriginGroup_Spec(gens map[string]gopter.Gen) {
	gens["HealthProbeSettings"] = gen.PtrOf(HealthProbeParametersGenerator())
	gens["LoadBalancingSettings"] = gen.PtrOf(LoadBalancingSettingsParametersGenerator())
}

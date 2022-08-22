// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180501preview

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

func Test_WebTest_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTest_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestSTATUSARM, WebTestSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestSTATUSARM runs a test to see if a specific instance of WebTest_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestSTATUSARM(subject WebTest_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTest_STATUSARM
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

// Generator of WebTest_STATUSARM instances for property testing - lazily instantiated by WebTestSTATUSARMGenerator()
var webTestSTATUSARMGenerator gopter.Gen

// WebTestSTATUSARMGenerator returns a generator of WebTest_STATUSARM instances for property testing.
// We first initialize webTestSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WebTestSTATUSARMGenerator() gopter.Gen {
	if webTestSTATUSARMGenerator != nil {
		return webTestSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestSTATUSARM(generators)
	webTestSTATUSARMGenerator = gen.Struct(reflect.TypeOf(WebTest_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForWebTestSTATUSARM(generators)
	webTestSTATUSARMGenerator = gen.Struct(reflect.TypeOf(WebTest_STATUSARM{}), generators)

	return webTestSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestSTATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWebTestSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWebTestSTATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(WebTestPropertiesSTATUSARMGenerator())
}

func Test_WebTestProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestPropertiesSTATUSARM, WebTestPropertiesSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestPropertiesSTATUSARM runs a test to see if a specific instance of WebTestProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestPropertiesSTATUSARM(subject WebTestProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestProperties_STATUSARM
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

// Generator of WebTestProperties_STATUSARM instances for property testing - lazily instantiated by
// WebTestPropertiesSTATUSARMGenerator()
var webTestPropertiesSTATUSARMGenerator gopter.Gen

// WebTestPropertiesSTATUSARMGenerator returns a generator of WebTestProperties_STATUSARM instances for property testing.
// We first initialize webTestPropertiesSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WebTestPropertiesSTATUSARMGenerator() gopter.Gen {
	if webTestPropertiesSTATUSARMGenerator != nil {
		return webTestPropertiesSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestPropertiesSTATUSARM(generators)
	webTestPropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestPropertiesSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForWebTestPropertiesSTATUSARM(generators)
	webTestPropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_STATUSARM{}), generators)

	return webTestPropertiesSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestPropertiesSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestPropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Frequency"] = gen.PtrOf(gen.Int())
	gens["Kind"] = gen.PtrOf(gen.OneConstOf(
		WebTestPropertiesSTATUSKind_Basic,
		WebTestPropertiesSTATUSKind_Multistep,
		WebTestPropertiesSTATUSKind_Ping,
		WebTestPropertiesSTATUSKind_Standard))
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["RetryEnabled"] = gen.PtrOf(gen.Bool())
	gens["SyntheticMonitorId"] = gen.PtrOf(gen.AlphaString())
	gens["Timeout"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForWebTestPropertiesSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWebTestPropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["Configuration"] = gen.PtrOf(WebTestPropertiesSTATUSConfigurationARMGenerator())
	gens["Locations"] = gen.SliceOf(WebTestGeolocationSTATUSARMGenerator())
	gens["Request"] = gen.PtrOf(WebTestPropertiesSTATUSRequestARMGenerator())
	gens["ValidationRules"] = gen.PtrOf(WebTestPropertiesSTATUSValidationRulesARMGenerator())
}

func Test_WebTestGeolocation_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestGeolocation_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestGeolocationSTATUSARM, WebTestGeolocationSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestGeolocationSTATUSARM runs a test to see if a specific instance of WebTestGeolocation_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestGeolocationSTATUSARM(subject WebTestGeolocation_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestGeolocation_STATUSARM
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

// Generator of WebTestGeolocation_STATUSARM instances for property testing - lazily instantiated by
// WebTestGeolocationSTATUSARMGenerator()
var webTestGeolocationSTATUSARMGenerator gopter.Gen

// WebTestGeolocationSTATUSARMGenerator returns a generator of WebTestGeolocation_STATUSARM instances for property testing.
func WebTestGeolocationSTATUSARMGenerator() gopter.Gen {
	if webTestGeolocationSTATUSARMGenerator != nil {
		return webTestGeolocationSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestGeolocationSTATUSARM(generators)
	webTestGeolocationSTATUSARMGenerator = gen.Struct(reflect.TypeOf(WebTestGeolocation_STATUSARM{}), generators)

	return webTestGeolocationSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestGeolocationSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestGeolocationSTATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_WebTestProperties_STATUS_ConfigurationARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestProperties_STATUS_ConfigurationARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestPropertiesSTATUSConfigurationARM, WebTestPropertiesSTATUSConfigurationARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestPropertiesSTATUSConfigurationARM runs a test to see if a specific instance of WebTestProperties_STATUS_ConfigurationARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestPropertiesSTATUSConfigurationARM(subject WebTestProperties_STATUS_ConfigurationARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestProperties_STATUS_ConfigurationARM
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

// Generator of WebTestProperties_STATUS_ConfigurationARM instances for property testing - lazily instantiated by
// WebTestPropertiesSTATUSConfigurationARMGenerator()
var webTestPropertiesSTATUSConfigurationARMGenerator gopter.Gen

// WebTestPropertiesSTATUSConfigurationARMGenerator returns a generator of WebTestProperties_STATUS_ConfigurationARM instances for property testing.
func WebTestPropertiesSTATUSConfigurationARMGenerator() gopter.Gen {
	if webTestPropertiesSTATUSConfigurationARMGenerator != nil {
		return webTestPropertiesSTATUSConfigurationARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestPropertiesSTATUSConfigurationARM(generators)
	webTestPropertiesSTATUSConfigurationARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_STATUS_ConfigurationARM{}), generators)

	return webTestPropertiesSTATUSConfigurationARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestPropertiesSTATUSConfigurationARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestPropertiesSTATUSConfigurationARM(gens map[string]gopter.Gen) {
	gens["WebTest"] = gen.PtrOf(gen.AlphaString())
}

func Test_WebTestProperties_STATUS_RequestARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestProperties_STATUS_RequestARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestPropertiesSTATUSRequestARM, WebTestPropertiesSTATUSRequestARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestPropertiesSTATUSRequestARM runs a test to see if a specific instance of WebTestProperties_STATUS_RequestARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestPropertiesSTATUSRequestARM(subject WebTestProperties_STATUS_RequestARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestProperties_STATUS_RequestARM
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

// Generator of WebTestProperties_STATUS_RequestARM instances for property testing - lazily instantiated by
// WebTestPropertiesSTATUSRequestARMGenerator()
var webTestPropertiesSTATUSRequestARMGenerator gopter.Gen

// WebTestPropertiesSTATUSRequestARMGenerator returns a generator of WebTestProperties_STATUS_RequestARM instances for property testing.
// We first initialize webTestPropertiesSTATUSRequestARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WebTestPropertiesSTATUSRequestARMGenerator() gopter.Gen {
	if webTestPropertiesSTATUSRequestARMGenerator != nil {
		return webTestPropertiesSTATUSRequestARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestPropertiesSTATUSRequestARM(generators)
	webTestPropertiesSTATUSRequestARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_STATUS_RequestARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestPropertiesSTATUSRequestARM(generators)
	AddRelatedPropertyGeneratorsForWebTestPropertiesSTATUSRequestARM(generators)
	webTestPropertiesSTATUSRequestARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_STATUS_RequestARM{}), generators)

	return webTestPropertiesSTATUSRequestARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestPropertiesSTATUSRequestARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestPropertiesSTATUSRequestARM(gens map[string]gopter.Gen) {
	gens["FollowRedirects"] = gen.PtrOf(gen.Bool())
	gens["HttpVerb"] = gen.PtrOf(gen.AlphaString())
	gens["ParseDependentRequests"] = gen.PtrOf(gen.Bool())
	gens["RequestBody"] = gen.PtrOf(gen.AlphaString())
	gens["RequestUrl"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWebTestPropertiesSTATUSRequestARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWebTestPropertiesSTATUSRequestARM(gens map[string]gopter.Gen) {
	gens["Headers"] = gen.SliceOf(HeaderFieldSTATUSARMGenerator())
}

func Test_WebTestProperties_STATUS_ValidationRulesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestProperties_STATUS_ValidationRulesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestPropertiesSTATUSValidationRulesARM, WebTestPropertiesSTATUSValidationRulesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestPropertiesSTATUSValidationRulesARM runs a test to see if a specific instance of WebTestProperties_STATUS_ValidationRulesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestPropertiesSTATUSValidationRulesARM(subject WebTestProperties_STATUS_ValidationRulesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestProperties_STATUS_ValidationRulesARM
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

// Generator of WebTestProperties_STATUS_ValidationRulesARM instances for property testing - lazily instantiated by
// WebTestPropertiesSTATUSValidationRulesARMGenerator()
var webTestPropertiesSTATUSValidationRulesARMGenerator gopter.Gen

// WebTestPropertiesSTATUSValidationRulesARMGenerator returns a generator of WebTestProperties_STATUS_ValidationRulesARM instances for property testing.
// We first initialize webTestPropertiesSTATUSValidationRulesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WebTestPropertiesSTATUSValidationRulesARMGenerator() gopter.Gen {
	if webTestPropertiesSTATUSValidationRulesARMGenerator != nil {
		return webTestPropertiesSTATUSValidationRulesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestPropertiesSTATUSValidationRulesARM(generators)
	webTestPropertiesSTATUSValidationRulesARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_STATUS_ValidationRulesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestPropertiesSTATUSValidationRulesARM(generators)
	AddRelatedPropertyGeneratorsForWebTestPropertiesSTATUSValidationRulesARM(generators)
	webTestPropertiesSTATUSValidationRulesARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_STATUS_ValidationRulesARM{}), generators)

	return webTestPropertiesSTATUSValidationRulesARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestPropertiesSTATUSValidationRulesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestPropertiesSTATUSValidationRulesARM(gens map[string]gopter.Gen) {
	gens["ExpectedHttpStatusCode"] = gen.PtrOf(gen.Int())
	gens["IgnoreHttpsStatusCode"] = gen.PtrOf(gen.Bool())
	gens["SSLCertRemainingLifetimeCheck"] = gen.PtrOf(gen.Int())
	gens["SSLCheck"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForWebTestPropertiesSTATUSValidationRulesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWebTestPropertiesSTATUSValidationRulesARM(gens map[string]gopter.Gen) {
	gens["ContentValidation"] = gen.PtrOf(WebTestPropertiesSTATUSValidationRulesContentValidationARMGenerator())
}

func Test_HeaderField_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HeaderField_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHeaderFieldSTATUSARM, HeaderFieldSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHeaderFieldSTATUSARM runs a test to see if a specific instance of HeaderField_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForHeaderFieldSTATUSARM(subject HeaderField_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HeaderField_STATUSARM
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

// Generator of HeaderField_STATUSARM instances for property testing - lazily instantiated by
// HeaderFieldSTATUSARMGenerator()
var headerFieldSTATUSARMGenerator gopter.Gen

// HeaderFieldSTATUSARMGenerator returns a generator of HeaderField_STATUSARM instances for property testing.
func HeaderFieldSTATUSARMGenerator() gopter.Gen {
	if headerFieldSTATUSARMGenerator != nil {
		return headerFieldSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHeaderFieldSTATUSARM(generators)
	headerFieldSTATUSARMGenerator = gen.Struct(reflect.TypeOf(HeaderField_STATUSARM{}), generators)

	return headerFieldSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForHeaderFieldSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHeaderFieldSTATUSARM(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_WebTestProperties_STATUS_ValidationRules_ContentValidationARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestProperties_STATUS_ValidationRules_ContentValidationARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestPropertiesSTATUSValidationRulesContentValidationARM, WebTestPropertiesSTATUSValidationRulesContentValidationARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestPropertiesSTATUSValidationRulesContentValidationARM runs a test to see if a specific instance of WebTestProperties_STATUS_ValidationRules_ContentValidationARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestPropertiesSTATUSValidationRulesContentValidationARM(subject WebTestProperties_STATUS_ValidationRules_ContentValidationARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestProperties_STATUS_ValidationRules_ContentValidationARM
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

// Generator of WebTestProperties_STATUS_ValidationRules_ContentValidationARM instances for property testing - lazily
// instantiated by WebTestPropertiesSTATUSValidationRulesContentValidationARMGenerator()
var webTestPropertiesSTATUSValidationRulesContentValidationARMGenerator gopter.Gen

// WebTestPropertiesSTATUSValidationRulesContentValidationARMGenerator returns a generator of WebTestProperties_STATUS_ValidationRules_ContentValidationARM instances for property testing.
func WebTestPropertiesSTATUSValidationRulesContentValidationARMGenerator() gopter.Gen {
	if webTestPropertiesSTATUSValidationRulesContentValidationARMGenerator != nil {
		return webTestPropertiesSTATUSValidationRulesContentValidationARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestPropertiesSTATUSValidationRulesContentValidationARM(generators)
	webTestPropertiesSTATUSValidationRulesContentValidationARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_STATUS_ValidationRules_ContentValidationARM{}), generators)

	return webTestPropertiesSTATUSValidationRulesContentValidationARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestPropertiesSTATUSValidationRulesContentValidationARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestPropertiesSTATUSValidationRulesContentValidationARM(gens map[string]gopter.Gen) {
	gens["ContentMatch"] = gen.PtrOf(gen.AlphaString())
	gens["IgnoreCase"] = gen.PtrOf(gen.Bool())
	gens["PassIfTextFound"] = gen.PtrOf(gen.Bool())
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220615

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

func Test_Webtest_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Webtest_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebtest_STATUS_ARM, Webtest_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebtest_STATUS_ARM runs a test to see if a specific instance of Webtest_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebtest_STATUS_ARM(subject Webtest_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Webtest_STATUS_ARM
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

// Generator of Webtest_STATUS_ARM instances for property testing - lazily instantiated by Webtest_STATUS_ARMGenerator()
var webtest_STATUS_ARMGenerator gopter.Gen

// Webtest_STATUS_ARMGenerator returns a generator of Webtest_STATUS_ARM instances for property testing.
// We first initialize webtest_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Webtest_STATUS_ARMGenerator() gopter.Gen {
	if webtest_STATUS_ARMGenerator != nil {
		return webtest_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebtest_STATUS_ARM(generators)
	webtest_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Webtest_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebtest_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForWebtest_STATUS_ARM(generators)
	webtest_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Webtest_STATUS_ARM{}), generators)

	return webtest_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWebtest_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebtest_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWebtest_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWebtest_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(WebTestProperties_STATUS_ARMGenerator())
}

func Test_WebTestProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestProperties_STATUS_ARM, WebTestProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestProperties_STATUS_ARM runs a test to see if a specific instance of WebTestProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestProperties_STATUS_ARM(subject WebTestProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestProperties_STATUS_ARM
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

// Generator of WebTestProperties_STATUS_ARM instances for property testing - lazily instantiated by
// WebTestProperties_STATUS_ARMGenerator()
var webTestProperties_STATUS_ARMGenerator gopter.Gen

// WebTestProperties_STATUS_ARMGenerator returns a generator of WebTestProperties_STATUS_ARM instances for property testing.
// We first initialize webTestProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WebTestProperties_STATUS_ARMGenerator() gopter.Gen {
	if webTestProperties_STATUS_ARMGenerator != nil {
		return webTestProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestProperties_STATUS_ARM(generators)
	webTestProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForWebTestProperties_STATUS_ARM(generators)
	webTestProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_STATUS_ARM{}), generators)

	return webTestProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Frequency"] = gen.PtrOf(gen.Int())
	gens["Kind"] = gen.PtrOf(gen.OneConstOf(WebTestProperties_Kind_STATUS_Multistep, WebTestProperties_Kind_STATUS_Ping, WebTestProperties_Kind_STATUS_Standard))
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["RetryEnabled"] = gen.PtrOf(gen.Bool())
	gens["SyntheticMonitorId"] = gen.PtrOf(gen.AlphaString())
	gens["Timeout"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForWebTestProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWebTestProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Configuration"] = gen.PtrOf(WebTestProperties_Configuration_STATUS_ARMGenerator())
	gens["Locations"] = gen.SliceOf(WebTestGeolocation_STATUS_ARMGenerator())
	gens["Request"] = gen.PtrOf(WebTestProperties_Request_STATUS_ARMGenerator())
	gens["ValidationRules"] = gen.PtrOf(WebTestProperties_ValidationRules_STATUS_ARMGenerator())
}

func Test_WebTestGeolocation_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestGeolocation_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestGeolocation_STATUS_ARM, WebTestGeolocation_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestGeolocation_STATUS_ARM runs a test to see if a specific instance of WebTestGeolocation_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestGeolocation_STATUS_ARM(subject WebTestGeolocation_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestGeolocation_STATUS_ARM
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

// Generator of WebTestGeolocation_STATUS_ARM instances for property testing - lazily instantiated by
// WebTestGeolocation_STATUS_ARMGenerator()
var webTestGeolocation_STATUS_ARMGenerator gopter.Gen

// WebTestGeolocation_STATUS_ARMGenerator returns a generator of WebTestGeolocation_STATUS_ARM instances for property testing.
func WebTestGeolocation_STATUS_ARMGenerator() gopter.Gen {
	if webTestGeolocation_STATUS_ARMGenerator != nil {
		return webTestGeolocation_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestGeolocation_STATUS_ARM(generators)
	webTestGeolocation_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(WebTestGeolocation_STATUS_ARM{}), generators)

	return webTestGeolocation_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestGeolocation_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestGeolocation_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_WebTestProperties_Configuration_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestProperties_Configuration_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestProperties_Configuration_STATUS_ARM, WebTestProperties_Configuration_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestProperties_Configuration_STATUS_ARM runs a test to see if a specific instance of WebTestProperties_Configuration_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestProperties_Configuration_STATUS_ARM(subject WebTestProperties_Configuration_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestProperties_Configuration_STATUS_ARM
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

// Generator of WebTestProperties_Configuration_STATUS_ARM instances for property testing - lazily instantiated by
// WebTestProperties_Configuration_STATUS_ARMGenerator()
var webTestProperties_Configuration_STATUS_ARMGenerator gopter.Gen

// WebTestProperties_Configuration_STATUS_ARMGenerator returns a generator of WebTestProperties_Configuration_STATUS_ARM instances for property testing.
func WebTestProperties_Configuration_STATUS_ARMGenerator() gopter.Gen {
	if webTestProperties_Configuration_STATUS_ARMGenerator != nil {
		return webTestProperties_Configuration_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestProperties_Configuration_STATUS_ARM(generators)
	webTestProperties_Configuration_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_Configuration_STATUS_ARM{}), generators)

	return webTestProperties_Configuration_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestProperties_Configuration_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestProperties_Configuration_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["WebTest"] = gen.PtrOf(gen.AlphaString())
}

func Test_WebTestProperties_Request_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestProperties_Request_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestProperties_Request_STATUS_ARM, WebTestProperties_Request_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestProperties_Request_STATUS_ARM runs a test to see if a specific instance of WebTestProperties_Request_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestProperties_Request_STATUS_ARM(subject WebTestProperties_Request_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestProperties_Request_STATUS_ARM
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

// Generator of WebTestProperties_Request_STATUS_ARM instances for property testing - lazily instantiated by
// WebTestProperties_Request_STATUS_ARMGenerator()
var webTestProperties_Request_STATUS_ARMGenerator gopter.Gen

// WebTestProperties_Request_STATUS_ARMGenerator returns a generator of WebTestProperties_Request_STATUS_ARM instances for property testing.
// We first initialize webTestProperties_Request_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WebTestProperties_Request_STATUS_ARMGenerator() gopter.Gen {
	if webTestProperties_Request_STATUS_ARMGenerator != nil {
		return webTestProperties_Request_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestProperties_Request_STATUS_ARM(generators)
	webTestProperties_Request_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_Request_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestProperties_Request_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForWebTestProperties_Request_STATUS_ARM(generators)
	webTestProperties_Request_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_Request_STATUS_ARM{}), generators)

	return webTestProperties_Request_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestProperties_Request_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestProperties_Request_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["FollowRedirects"] = gen.PtrOf(gen.Bool())
	gens["HttpVerb"] = gen.PtrOf(gen.AlphaString())
	gens["ParseDependentRequests"] = gen.PtrOf(gen.Bool())
	gens["RequestBody"] = gen.PtrOf(gen.AlphaString())
	gens["RequestUrl"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWebTestProperties_Request_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWebTestProperties_Request_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Headers"] = gen.SliceOf(HeaderField_STATUS_ARMGenerator())
}

func Test_WebTestProperties_ValidationRules_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestProperties_ValidationRules_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestProperties_ValidationRules_STATUS_ARM, WebTestProperties_ValidationRules_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestProperties_ValidationRules_STATUS_ARM runs a test to see if a specific instance of WebTestProperties_ValidationRules_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestProperties_ValidationRules_STATUS_ARM(subject WebTestProperties_ValidationRules_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestProperties_ValidationRules_STATUS_ARM
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

// Generator of WebTestProperties_ValidationRules_STATUS_ARM instances for property testing - lazily instantiated by
// WebTestProperties_ValidationRules_STATUS_ARMGenerator()
var webTestProperties_ValidationRules_STATUS_ARMGenerator gopter.Gen

// WebTestProperties_ValidationRules_STATUS_ARMGenerator returns a generator of WebTestProperties_ValidationRules_STATUS_ARM instances for property testing.
// We first initialize webTestProperties_ValidationRules_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WebTestProperties_ValidationRules_STATUS_ARMGenerator() gopter.Gen {
	if webTestProperties_ValidationRules_STATUS_ARMGenerator != nil {
		return webTestProperties_ValidationRules_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestProperties_ValidationRules_STATUS_ARM(generators)
	webTestProperties_ValidationRules_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_ValidationRules_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestProperties_ValidationRules_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForWebTestProperties_ValidationRules_STATUS_ARM(generators)
	webTestProperties_ValidationRules_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_ValidationRules_STATUS_ARM{}), generators)

	return webTestProperties_ValidationRules_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestProperties_ValidationRules_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestProperties_ValidationRules_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ExpectedHttpStatusCode"] = gen.PtrOf(gen.Int())
	gens["IgnoreHttpStatusCode"] = gen.PtrOf(gen.Bool())
	gens["SSLCertRemainingLifetimeCheck"] = gen.PtrOf(gen.Int())
	gens["SSLCheck"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForWebTestProperties_ValidationRules_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWebTestProperties_ValidationRules_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ContentValidation"] = gen.PtrOf(WebTestProperties_ValidationRules_ContentValidation_STATUS_ARMGenerator())
}

func Test_HeaderField_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HeaderField_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHeaderField_STATUS_ARM, HeaderField_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHeaderField_STATUS_ARM runs a test to see if a specific instance of HeaderField_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForHeaderField_STATUS_ARM(subject HeaderField_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HeaderField_STATUS_ARM
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

// Generator of HeaderField_STATUS_ARM instances for property testing - lazily instantiated by
// HeaderField_STATUS_ARMGenerator()
var headerField_STATUS_ARMGenerator gopter.Gen

// HeaderField_STATUS_ARMGenerator returns a generator of HeaderField_STATUS_ARM instances for property testing.
func HeaderField_STATUS_ARMGenerator() gopter.Gen {
	if headerField_STATUS_ARMGenerator != nil {
		return headerField_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHeaderField_STATUS_ARM(generators)
	headerField_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(HeaderField_STATUS_ARM{}), generators)

	return headerField_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForHeaderField_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHeaderField_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_WebTestProperties_ValidationRules_ContentValidation_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebTestProperties_ValidationRules_ContentValidation_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebTestProperties_ValidationRules_ContentValidation_STATUS_ARM, WebTestProperties_ValidationRules_ContentValidation_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebTestProperties_ValidationRules_ContentValidation_STATUS_ARM runs a test to see if a specific instance of WebTestProperties_ValidationRules_ContentValidation_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebTestProperties_ValidationRules_ContentValidation_STATUS_ARM(subject WebTestProperties_ValidationRules_ContentValidation_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebTestProperties_ValidationRules_ContentValidation_STATUS_ARM
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

// Generator of WebTestProperties_ValidationRules_ContentValidation_STATUS_ARM instances for property testing - lazily
// instantiated by WebTestProperties_ValidationRules_ContentValidation_STATUS_ARMGenerator()
var webTestProperties_ValidationRules_ContentValidation_STATUS_ARMGenerator gopter.Gen

// WebTestProperties_ValidationRules_ContentValidation_STATUS_ARMGenerator returns a generator of WebTestProperties_ValidationRules_ContentValidation_STATUS_ARM instances for property testing.
func WebTestProperties_ValidationRules_ContentValidation_STATUS_ARMGenerator() gopter.Gen {
	if webTestProperties_ValidationRules_ContentValidation_STATUS_ARMGenerator != nil {
		return webTestProperties_ValidationRules_ContentValidation_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebTestProperties_ValidationRules_ContentValidation_STATUS_ARM(generators)
	webTestProperties_ValidationRules_ContentValidation_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(WebTestProperties_ValidationRules_ContentValidation_STATUS_ARM{}), generators)

	return webTestProperties_ValidationRules_ContentValidation_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWebTestProperties_ValidationRules_ContentValidation_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebTestProperties_ValidationRules_ContentValidation_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ContentMatch"] = gen.PtrOf(gen.AlphaString())
	gens["IgnoreCase"] = gen.PtrOf(gen.Bool())
	gens["PassIfTextFound"] = gen.PtrOf(gen.Bool())
}

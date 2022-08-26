// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601storage

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

func Test_DomainsTopic_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DomainsTopic via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainsTopic, DomainsTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainsTopic runs a test to see if a specific instance of DomainsTopic round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainsTopic(subject DomainsTopic) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DomainsTopic
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

// Generator of DomainsTopic instances for property testing - lazily instantiated by DomainsTopicGenerator()
var domainsTopicGenerator gopter.Gen

// DomainsTopicGenerator returns a generator of DomainsTopic instances for property testing.
func DomainsTopicGenerator() gopter.Gen {
	if domainsTopicGenerator != nil {
		return domainsTopicGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDomainsTopic(generators)
	domainsTopicGenerator = gen.Struct(reflect.TypeOf(DomainsTopic{}), generators)

	return domainsTopicGenerator
}

// AddRelatedPropertyGeneratorsForDomainsTopic is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainsTopic(gens map[string]gopter.Gen) {
<<<<<<< HEAD
	gens["Spec"] = DomainsTopic_SpecGenerator()
	gens["Status"] = DomainsTopic_STATUSGenerator()
}

func Test_DomainsTopic_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
=======
	gens["Spec"] = Domains_Topics_SpecGenerator()
	gens["Status"] = DomainTopic_STATUSGenerator()
}

func Test_Domains_Topics_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
>>>>>>> main
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<< HEAD
		"Round trip of DomainsTopic_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainsTopic_Spec, DomainsTopic_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainsTopic_Spec runs a test to see if a specific instance of DomainsTopic_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainsTopic_Spec(subject DomainsTopic_Spec) string {
=======
		"Round trip of Domains_Topics_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomains_Topics_Spec, Domains_Topics_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomains_Topics_Spec runs a test to see if a specific instance of Domains_Topics_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDomains_Topics_Spec(subject Domains_Topics_Spec) string {
>>>>>>> main
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
<<<<<<< HEAD
	var actual DomainsTopic_Spec
=======
	var actual Domains_Topics_Spec
>>>>>>> main
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

<<<<<<< HEAD
// Generator of DomainsTopic_Spec instances for property testing - lazily instantiated by DomainsTopic_SpecGenerator()
var domainsTopic_SpecGenerator gopter.Gen

// DomainsTopic_SpecGenerator returns a generator of DomainsTopic_Spec instances for property testing.
func DomainsTopic_SpecGenerator() gopter.Gen {
	if domainsTopic_SpecGenerator != nil {
		return domainsTopic_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainsTopic_Spec(generators)
	domainsTopic_SpecGenerator = gen.Struct(reflect.TypeOf(DomainsTopic_Spec{}), generators)

	return domainsTopic_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDomainsTopic_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainsTopic_Spec(gens map[string]gopter.Gen) {
=======
// Generator of Domains_Topics_Spec instances for property testing - lazily instantiated by
// Domains_Topics_SpecGenerator()
var domains_Topics_SpecGenerator gopter.Gen

// Domains_Topics_SpecGenerator returns a generator of Domains_Topics_Spec instances for property testing.
func Domains_Topics_SpecGenerator() gopter.Gen {
	if domains_Topics_SpecGenerator != nil {
		return domains_Topics_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomains_Topics_Spec(generators)
	domains_Topics_SpecGenerator = gen.Struct(reflect.TypeOf(Domains_Topics_Spec{}), generators)

	return domains_Topics_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDomains_Topics_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomains_Topics_Spec(gens map[string]gopter.Gen) {
>>>>>>> main
	gens["AzureName"] = gen.AlphaString()
	gens["OriginalVersion"] = gen.AlphaString()
}

func Test_DomainsTopic_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
<<<<<<< HEAD
		"Round trip of DomainsTopic_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainsTopic_STATUS, DomainsTopic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainsTopic_STATUS runs a test to see if a specific instance of DomainsTopic_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainsTopic_STATUS(subject DomainsTopic_STATUS) string {
=======
		"Round trip of DomainTopic_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainTopic_STATUS, DomainTopic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainTopic_STATUS runs a test to see if a specific instance of DomainTopic_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainTopic_STATUS(subject DomainTopic_STATUS) string {
>>>>>>> main
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DomainsTopic_STATUS
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

<<<<<<< HEAD
// Generator of DomainsTopic_STATUS instances for property testing - lazily instantiated by
// DomainsTopic_STATUSGenerator()
var domainsTopic_STATUSGenerator gopter.Gen

// DomainsTopic_STATUSGenerator returns a generator of DomainsTopic_STATUS instances for property testing.
// We first initialize domainsTopic_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DomainsTopic_STATUSGenerator() gopter.Gen {
	if domainsTopic_STATUSGenerator != nil {
		return domainsTopic_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainsTopic_STATUS(generators)
	domainsTopic_STATUSGenerator = gen.Struct(reflect.TypeOf(DomainsTopic_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainsTopic_STATUS(generators)
	AddRelatedPropertyGeneratorsForDomainsTopic_STATUS(generators)
	domainsTopic_STATUSGenerator = gen.Struct(reflect.TypeOf(DomainsTopic_STATUS{}), generators)

	return domainsTopic_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDomainsTopic_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainsTopic_STATUS(gens map[string]gopter.Gen) {
=======
// Generator of DomainTopic_STATUS instances for property testing - lazily instantiated by DomainTopic_STATUSGenerator()
var domainTopic_STATUSGenerator gopter.Gen

// DomainTopic_STATUSGenerator returns a generator of DomainTopic_STATUS instances for property testing.
// We first initialize domainTopic_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DomainTopic_STATUSGenerator() gopter.Gen {
	if domainTopic_STATUSGenerator != nil {
		return domainTopic_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainTopic_STATUS(generators)
	domainTopic_STATUSGenerator = gen.Struct(reflect.TypeOf(DomainTopic_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainTopic_STATUS(generators)
	AddRelatedPropertyGeneratorsForDomainTopic_STATUS(generators)
	domainTopic_STATUSGenerator = gen.Struct(reflect.TypeOf(DomainTopic_STATUS{}), generators)

	return domainTopic_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDomainTopic_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainTopic_STATUS(gens map[string]gopter.Gen) {
>>>>>>> main
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

<<<<<<< HEAD
// AddRelatedPropertyGeneratorsForDomainsTopic_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainsTopic_STATUS(gens map[string]gopter.Gen) {
=======
// AddRelatedPropertyGeneratorsForDomainTopic_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainTopic_STATUS(gens map[string]gopter.Gen) {
>>>>>>> main
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

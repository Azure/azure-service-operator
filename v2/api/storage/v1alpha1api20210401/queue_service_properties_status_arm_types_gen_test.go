// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

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

func Test_QueueServiceProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of QueueServiceProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForQueueServicePropertiesSTATUSARM, QueueServicePropertiesSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForQueueServicePropertiesSTATUSARM runs a test to see if a specific instance of QueueServiceProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForQueueServicePropertiesSTATUSARM(subject QueueServiceProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual QueueServiceProperties_STATUSARM
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

// Generator of QueueServiceProperties_STATUSARM instances for property testing - lazily instantiated by
// QueueServicePropertiesSTATUSARMGenerator()
var queueServicePropertiesSTATUSARMGenerator gopter.Gen

// QueueServicePropertiesSTATUSARMGenerator returns a generator of QueueServiceProperties_STATUSARM instances for property testing.
// We first initialize queueServicePropertiesSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func QueueServicePropertiesSTATUSARMGenerator() gopter.Gen {
	if queueServicePropertiesSTATUSARMGenerator != nil {
		return queueServicePropertiesSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForQueueServicePropertiesSTATUSARM(generators)
	queueServicePropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(QueueServiceProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForQueueServicePropertiesSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForQueueServicePropertiesSTATUSARM(generators)
	queueServicePropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(QueueServiceProperties_STATUSARM{}), generators)

	return queueServicePropertiesSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForQueueServicePropertiesSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForQueueServicePropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForQueueServicePropertiesSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForQueueServicePropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(QueueServicePropertiesSTATUSPropertiesARMGenerator())
}

func Test_QueueServiceProperties_STATUS_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of QueueServiceProperties_STATUS_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForQueueServicePropertiesSTATUSPropertiesARM, QueueServicePropertiesSTATUSPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForQueueServicePropertiesSTATUSPropertiesARM runs a test to see if a specific instance of QueueServiceProperties_STATUS_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForQueueServicePropertiesSTATUSPropertiesARM(subject QueueServiceProperties_STATUS_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual QueueServiceProperties_STATUS_PropertiesARM
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

// Generator of QueueServiceProperties_STATUS_PropertiesARM instances for property testing - lazily instantiated by
// QueueServicePropertiesSTATUSPropertiesARMGenerator()
var queueServicePropertiesSTATUSPropertiesARMGenerator gopter.Gen

// QueueServicePropertiesSTATUSPropertiesARMGenerator returns a generator of QueueServiceProperties_STATUS_PropertiesARM instances for property testing.
func QueueServicePropertiesSTATUSPropertiesARMGenerator() gopter.Gen {
	if queueServicePropertiesSTATUSPropertiesARMGenerator != nil {
		return queueServicePropertiesSTATUSPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForQueueServicePropertiesSTATUSPropertiesARM(generators)
	queueServicePropertiesSTATUSPropertiesARMGenerator = gen.Struct(reflect.TypeOf(QueueServiceProperties_STATUS_PropertiesARM{}), generators)

	return queueServicePropertiesSTATUSPropertiesARMGenerator
}

// AddRelatedPropertyGeneratorsForQueueServicePropertiesSTATUSPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForQueueServicePropertiesSTATUSPropertiesARM(gens map[string]gopter.Gen) {
	gens["Cors"] = gen.PtrOf(CorsRulesSTATUSARMGenerator())
}

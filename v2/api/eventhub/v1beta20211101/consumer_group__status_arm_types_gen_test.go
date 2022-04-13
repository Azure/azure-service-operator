// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

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

func Test_ConsumerGroup_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConsumerGroup_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConsumerGroupStatusARM, ConsumerGroupStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConsumerGroupStatusARM runs a test to see if a specific instance of ConsumerGroup_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConsumerGroupStatusARM(subject ConsumerGroup_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConsumerGroup_StatusARM
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

// Generator of ConsumerGroup_StatusARM instances for property testing - lazily instantiated by
//ConsumerGroupStatusARMGenerator()
var consumerGroupStatusARMGenerator gopter.Gen

// ConsumerGroupStatusARMGenerator returns a generator of ConsumerGroup_StatusARM instances for property testing.
// We first initialize consumerGroupStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ConsumerGroupStatusARMGenerator() gopter.Gen {
	if consumerGroupStatusARMGenerator != nil {
		return consumerGroupStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConsumerGroupStatusARM(generators)
	consumerGroupStatusARMGenerator = gen.Struct(reflect.TypeOf(ConsumerGroup_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConsumerGroupStatusARM(generators)
	AddRelatedPropertyGeneratorsForConsumerGroupStatusARM(generators)
	consumerGroupStatusARMGenerator = gen.Struct(reflect.TypeOf(ConsumerGroup_StatusARM{}), generators)

	return consumerGroupStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForConsumerGroupStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConsumerGroupStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForConsumerGroupStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForConsumerGroupStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ConsumerGroupStatusPropertiesARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusARMGenerator())
}

func Test_ConsumerGroup_Status_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConsumerGroup_Status_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConsumerGroupStatusPropertiesARM, ConsumerGroupStatusPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConsumerGroupStatusPropertiesARM runs a test to see if a specific instance of ConsumerGroup_Status_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConsumerGroupStatusPropertiesARM(subject ConsumerGroup_Status_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConsumerGroup_Status_PropertiesARM
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

// Generator of ConsumerGroup_Status_PropertiesARM instances for property testing - lazily instantiated by
//ConsumerGroupStatusPropertiesARMGenerator()
var consumerGroupStatusPropertiesARMGenerator gopter.Gen

// ConsumerGroupStatusPropertiesARMGenerator returns a generator of ConsumerGroup_Status_PropertiesARM instances for property testing.
func ConsumerGroupStatusPropertiesARMGenerator() gopter.Gen {
	if consumerGroupStatusPropertiesARMGenerator != nil {
		return consumerGroupStatusPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConsumerGroupStatusPropertiesARM(generators)
	consumerGroupStatusPropertiesARMGenerator = gen.Struct(reflect.TypeOf(ConsumerGroup_Status_PropertiesARM{}), generators)

	return consumerGroupStatusPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForConsumerGroupStatusPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConsumerGroupStatusPropertiesARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["UserMetadata"] = gen.PtrOf(gen.AlphaString())
}

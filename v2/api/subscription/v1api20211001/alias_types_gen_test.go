// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211001

import (
	"encoding/json"
	v1api20211001s "github.com/Azure/azure-service-operator/v2/api/subscription/v1api20211001storage"
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

func Test_Alias_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Alias to hub returns original",
		prop.ForAll(RunResourceConversionTestForAlias, AliasGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForAlias tests if a specific instance of Alias round trips to the hub storage version and back losslessly
func RunResourceConversionTestForAlias(subject Alias) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1api20211001s.Alias
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual Alias
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

func Test_Alias_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Alias to Alias via AssignProperties_To_Alias & AssignProperties_From_Alias returns original",
		prop.ForAll(RunPropertyAssignmentTestForAlias, AliasGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAlias tests if a specific instance of Alias can be assigned to v1api20211001storage and back losslessly
func RunPropertyAssignmentTestForAlias(subject Alias) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20211001s.Alias
	err := copied.AssignProperties_To_Alias(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Alias
	err = actual.AssignProperties_From_Alias(&other)
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

func Test_Alias_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Alias via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAlias, AliasGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAlias runs a test to see if a specific instance of Alias round trips to JSON and back losslessly
func RunJSONSerializationTestForAlias(subject Alias) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Alias
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

// Generator of Alias instances for property testing - lazily instantiated by AliasGenerator()
var aliasGenerator gopter.Gen

// AliasGenerator returns a generator of Alias instances for property testing.
func AliasGenerator() gopter.Gen {
	if aliasGenerator != nil {
		return aliasGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAlias(generators)
	aliasGenerator = gen.Struct(reflect.TypeOf(Alias{}), generators)

	return aliasGenerator
}

// AddRelatedPropertyGeneratorsForAlias is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAlias(gens map[string]gopter.Gen) {
	gens["Spec"] = Alias_SpecGenerator()
	gens["Status"] = Alias_STATUSGenerator()
}

func Test_Alias_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Alias_Spec to Alias_Spec via AssignProperties_To_Alias_Spec & AssignProperties_From_Alias_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForAlias_Spec, Alias_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAlias_Spec tests if a specific instance of Alias_Spec can be assigned to v1api20211001storage and back losslessly
func RunPropertyAssignmentTestForAlias_Spec(subject Alias_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20211001s.Alias_Spec
	err := copied.AssignProperties_To_Alias_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Alias_Spec
	err = actual.AssignProperties_From_Alias_Spec(&other)
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

func Test_Alias_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Alias_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAlias_Spec, Alias_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAlias_Spec runs a test to see if a specific instance of Alias_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForAlias_Spec(subject Alias_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Alias_Spec
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

// Generator of Alias_Spec instances for property testing - lazily instantiated by Alias_SpecGenerator()
var alias_SpecGenerator gopter.Gen

// Alias_SpecGenerator returns a generator of Alias_Spec instances for property testing.
// We first initialize alias_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Alias_SpecGenerator() gopter.Gen {
	if alias_SpecGenerator != nil {
		return alias_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAlias_Spec(generators)
	alias_SpecGenerator = gen.Struct(reflect.TypeOf(Alias_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAlias_Spec(generators)
	AddRelatedPropertyGeneratorsForAlias_Spec(generators)
	alias_SpecGenerator = gen.Struct(reflect.TypeOf(Alias_Spec{}), generators)

	return alias_SpecGenerator
}

// AddIndependentPropertyGeneratorsForAlias_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAlias_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForAlias_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAlias_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PutAliasRequestPropertiesGenerator())
}

func Test_Alias_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Alias_STATUS to Alias_STATUS via AssignProperties_To_Alias_STATUS & AssignProperties_From_Alias_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForAlias_STATUS, Alias_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAlias_STATUS tests if a specific instance of Alias_STATUS can be assigned to v1api20211001storage and back losslessly
func RunPropertyAssignmentTestForAlias_STATUS(subject Alias_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20211001s.Alias_STATUS
	err := copied.AssignProperties_To_Alias_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Alias_STATUS
	err = actual.AssignProperties_From_Alias_STATUS(&other)
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

func Test_Alias_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Alias_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAlias_STATUS, Alias_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAlias_STATUS runs a test to see if a specific instance of Alias_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAlias_STATUS(subject Alias_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Alias_STATUS
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

// Generator of Alias_STATUS instances for property testing - lazily instantiated by Alias_STATUSGenerator()
var alias_STATUSGenerator gopter.Gen

// Alias_STATUSGenerator returns a generator of Alias_STATUS instances for property testing.
// We first initialize alias_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Alias_STATUSGenerator() gopter.Gen {
	if alias_STATUSGenerator != nil {
		return alias_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAlias_STATUS(generators)
	alias_STATUSGenerator = gen.Struct(reflect.TypeOf(Alias_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAlias_STATUS(generators)
	AddRelatedPropertyGeneratorsForAlias_STATUS(generators)
	alias_STATUSGenerator = gen.Struct(reflect.TypeOf(Alias_STATUS{}), generators)

	return alias_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAlias_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAlias_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAlias_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAlias_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SubscriptionAliasResponseProperties_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_PutAliasRequestProperties_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PutAliasRequestProperties to PutAliasRequestProperties via AssignProperties_To_PutAliasRequestProperties & AssignProperties_From_PutAliasRequestProperties returns original",
		prop.ForAll(RunPropertyAssignmentTestForPutAliasRequestProperties, PutAliasRequestPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPutAliasRequestProperties tests if a specific instance of PutAliasRequestProperties can be assigned to v1api20211001storage and back losslessly
func RunPropertyAssignmentTestForPutAliasRequestProperties(subject PutAliasRequestProperties) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20211001s.PutAliasRequestProperties
	err := copied.AssignProperties_To_PutAliasRequestProperties(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PutAliasRequestProperties
	err = actual.AssignProperties_From_PutAliasRequestProperties(&other)
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

func Test_PutAliasRequestProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PutAliasRequestProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPutAliasRequestProperties, PutAliasRequestPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPutAliasRequestProperties runs a test to see if a specific instance of PutAliasRequestProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForPutAliasRequestProperties(subject PutAliasRequestProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PutAliasRequestProperties
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

// Generator of PutAliasRequestProperties instances for property testing - lazily instantiated by
// PutAliasRequestPropertiesGenerator()
var putAliasRequestPropertiesGenerator gopter.Gen

// PutAliasRequestPropertiesGenerator returns a generator of PutAliasRequestProperties instances for property testing.
// We first initialize putAliasRequestPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PutAliasRequestPropertiesGenerator() gopter.Gen {
	if putAliasRequestPropertiesGenerator != nil {
		return putAliasRequestPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPutAliasRequestProperties(generators)
	putAliasRequestPropertiesGenerator = gen.Struct(reflect.TypeOf(PutAliasRequestProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPutAliasRequestProperties(generators)
	AddRelatedPropertyGeneratorsForPutAliasRequestProperties(generators)
	putAliasRequestPropertiesGenerator = gen.Struct(reflect.TypeOf(PutAliasRequestProperties{}), generators)

	return putAliasRequestPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForPutAliasRequestProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPutAliasRequestProperties(gens map[string]gopter.Gen) {
	gens["BillingScope"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["ResellerId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["Workload"] = gen.PtrOf(gen.OneConstOf(Workload_DevTest, Workload_Production))
}

// AddRelatedPropertyGeneratorsForPutAliasRequestProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPutAliasRequestProperties(gens map[string]gopter.Gen) {
	gens["AdditionalProperties"] = gen.PtrOf(PutAliasRequestAdditionalPropertiesGenerator())
}

func Test_SubscriptionAliasResponseProperties_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SubscriptionAliasResponseProperties_STATUS to SubscriptionAliasResponseProperties_STATUS via AssignProperties_To_SubscriptionAliasResponseProperties_STATUS & AssignProperties_From_SubscriptionAliasResponseProperties_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSubscriptionAliasResponseProperties_STATUS, SubscriptionAliasResponseProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSubscriptionAliasResponseProperties_STATUS tests if a specific instance of SubscriptionAliasResponseProperties_STATUS can be assigned to v1api20211001storage and back losslessly
func RunPropertyAssignmentTestForSubscriptionAliasResponseProperties_STATUS(subject SubscriptionAliasResponseProperties_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20211001s.SubscriptionAliasResponseProperties_STATUS
	err := copied.AssignProperties_To_SubscriptionAliasResponseProperties_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SubscriptionAliasResponseProperties_STATUS
	err = actual.AssignProperties_From_SubscriptionAliasResponseProperties_STATUS(&other)
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

func Test_SubscriptionAliasResponseProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubscriptionAliasResponseProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubscriptionAliasResponseProperties_STATUS, SubscriptionAliasResponseProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubscriptionAliasResponseProperties_STATUS runs a test to see if a specific instance of SubscriptionAliasResponseProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSubscriptionAliasResponseProperties_STATUS(subject SubscriptionAliasResponseProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubscriptionAliasResponseProperties_STATUS
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

// Generator of SubscriptionAliasResponseProperties_STATUS instances for property testing - lazily instantiated by
// SubscriptionAliasResponseProperties_STATUSGenerator()
var subscriptionAliasResponseProperties_STATUSGenerator gopter.Gen

// SubscriptionAliasResponseProperties_STATUSGenerator returns a generator of SubscriptionAliasResponseProperties_STATUS instances for property testing.
func SubscriptionAliasResponseProperties_STATUSGenerator() gopter.Gen {
	if subscriptionAliasResponseProperties_STATUSGenerator != nil {
		return subscriptionAliasResponseProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubscriptionAliasResponseProperties_STATUS(generators)
	subscriptionAliasResponseProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(SubscriptionAliasResponseProperties_STATUS{}), generators)

	return subscriptionAliasResponseProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSubscriptionAliasResponseProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubscriptionAliasResponseProperties_STATUS(gens map[string]gopter.Gen) {
	gens["AcceptOwnershipState"] = gen.PtrOf(gen.OneConstOf(AcceptOwnershipState_STATUS_Completed, AcceptOwnershipState_STATUS_Expired, AcceptOwnershipState_STATUS_Pending))
	gens["AcceptOwnershipUrl"] = gen.PtrOf(gen.AlphaString())
	gens["BillingScope"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedTime"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["ManagementGroupId"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(SubscriptionAliasResponseProperties_ProvisioningState_STATUS_Accepted, SubscriptionAliasResponseProperties_ProvisioningState_STATUS_Failed, SubscriptionAliasResponseProperties_ProvisioningState_STATUS_Succeeded))
	gens["ResellerId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionOwnerId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Workload"] = gen.PtrOf(gen.OneConstOf(Workload_STATUS_DevTest, Workload_STATUS_Production))
}

func Test_SystemData_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SystemData_STATUS to SystemData_STATUS via AssignProperties_To_SystemData_STATUS & AssignProperties_From_SystemData_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSystemData_STATUS, SystemData_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSystemData_STATUS tests if a specific instance of SystemData_STATUS can be assigned to v1api20211001storage and back losslessly
func RunPropertyAssignmentTestForSystemData_STATUS(subject SystemData_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20211001s.SystemData_STATUS
	err := copied.AssignProperties_To_SystemData_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SystemData_STATUS
	err = actual.AssignProperties_From_SystemData_STATUS(&other)
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

func Test_SystemData_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUS, SystemData_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUS runs a test to see if a specific instance of SystemData_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUS(subject SystemData_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUS
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

// Generator of SystemData_STATUS instances for property testing - lazily instantiated by SystemData_STATUSGenerator()
var systemData_STATUSGenerator gopter.Gen

// SystemData_STATUSGenerator returns a generator of SystemData_STATUS instances for property testing.
func SystemData_STATUSGenerator() gopter.Gen {
	if systemData_STATUSGenerator != nil {
		return systemData_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUS(generators)
	systemData_STATUSGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUS{}), generators)

	return systemData_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUS(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_CreatedByType_STATUS_Application,
		SystemData_CreatedByType_STATUS_Key,
		SystemData_CreatedByType_STATUS_ManagedIdentity,
		SystemData_CreatedByType_STATUS_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_LastModifiedByType_STATUS_Application,
		SystemData_LastModifiedByType_STATUS_Key,
		SystemData_LastModifiedByType_STATUS_ManagedIdentity,
		SystemData_LastModifiedByType_STATUS_User))
}

func Test_PutAliasRequestAdditionalProperties_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PutAliasRequestAdditionalProperties to PutAliasRequestAdditionalProperties via AssignProperties_To_PutAliasRequestAdditionalProperties & AssignProperties_From_PutAliasRequestAdditionalProperties returns original",
		prop.ForAll(RunPropertyAssignmentTestForPutAliasRequestAdditionalProperties, PutAliasRequestAdditionalPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPutAliasRequestAdditionalProperties tests if a specific instance of PutAliasRequestAdditionalProperties can be assigned to v1api20211001storage and back losslessly
func RunPropertyAssignmentTestForPutAliasRequestAdditionalProperties(subject PutAliasRequestAdditionalProperties) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20211001s.PutAliasRequestAdditionalProperties
	err := copied.AssignProperties_To_PutAliasRequestAdditionalProperties(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PutAliasRequestAdditionalProperties
	err = actual.AssignProperties_From_PutAliasRequestAdditionalProperties(&other)
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

func Test_PutAliasRequestAdditionalProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PutAliasRequestAdditionalProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPutAliasRequestAdditionalProperties, PutAliasRequestAdditionalPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPutAliasRequestAdditionalProperties runs a test to see if a specific instance of PutAliasRequestAdditionalProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForPutAliasRequestAdditionalProperties(subject PutAliasRequestAdditionalProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PutAliasRequestAdditionalProperties
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

// Generator of PutAliasRequestAdditionalProperties instances for property testing - lazily instantiated by
// PutAliasRequestAdditionalPropertiesGenerator()
var putAliasRequestAdditionalPropertiesGenerator gopter.Gen

// PutAliasRequestAdditionalPropertiesGenerator returns a generator of PutAliasRequestAdditionalProperties instances for property testing.
func PutAliasRequestAdditionalPropertiesGenerator() gopter.Gen {
	if putAliasRequestAdditionalPropertiesGenerator != nil {
		return putAliasRequestAdditionalPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPutAliasRequestAdditionalProperties(generators)
	putAliasRequestAdditionalPropertiesGenerator = gen.Struct(reflect.TypeOf(PutAliasRequestAdditionalProperties{}), generators)

	return putAliasRequestAdditionalPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForPutAliasRequestAdditionalProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPutAliasRequestAdditionalProperties(gens map[string]gopter.Gen) {
	gens["ManagementGroupId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionOwnerId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionTenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

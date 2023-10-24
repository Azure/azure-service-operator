// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901

import (
	"encoding/json"
	v20220901s "github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901storage"
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

func Test_StorageAccountsTableService_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsTableService to hub returns original",
		prop.ForAll(RunResourceConversionTestForStorageAccountsTableService, StorageAccountsTableServiceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForStorageAccountsTableService tests if a specific instance of StorageAccountsTableService round trips to the hub storage version and back losslessly
func RunResourceConversionTestForStorageAccountsTableService(subject StorageAccountsTableService) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20220901s.StorageAccountsTableService
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual StorageAccountsTableService
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

func Test_StorageAccountsTableService_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsTableService to StorageAccountsTableService via AssignProperties_To_StorageAccountsTableService & AssignProperties_From_StorageAccountsTableService returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsTableService, StorageAccountsTableServiceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsTableService tests if a specific instance of StorageAccountsTableService can be assigned to v1api20220901storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsTableService(subject StorageAccountsTableService) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.StorageAccountsTableService
	err := copied.AssignProperties_To_StorageAccountsTableService(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsTableService
	err = actual.AssignProperties_From_StorageAccountsTableService(&other)
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

func Test_StorageAccountsTableService_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsTableService via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsTableService, StorageAccountsTableServiceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsTableService runs a test to see if a specific instance of StorageAccountsTableService round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsTableService(subject StorageAccountsTableService) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsTableService
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

// Generator of StorageAccountsTableService instances for property testing - lazily instantiated by
// StorageAccountsTableServiceGenerator()
var storageAccountsTableServiceGenerator gopter.Gen

// StorageAccountsTableServiceGenerator returns a generator of StorageAccountsTableService instances for property testing.
func StorageAccountsTableServiceGenerator() gopter.Gen {
	if storageAccountsTableServiceGenerator != nil {
		return storageAccountsTableServiceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccountsTableService(generators)
	storageAccountsTableServiceGenerator = gen.Struct(reflect.TypeOf(StorageAccountsTableService{}), generators)

	return storageAccountsTableServiceGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccountsTableService is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsTableService(gens map[string]gopter.Gen) {
	gens["Spec"] = StorageAccounts_TableService_SpecGenerator()
	gens["Status"] = StorageAccounts_TableService_STATUSGenerator()
}

func Test_StorageAccounts_TableService_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccounts_TableService_Spec to StorageAccounts_TableService_Spec via AssignProperties_To_StorageAccounts_TableService_Spec & AssignProperties_From_StorageAccounts_TableService_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccounts_TableService_Spec, StorageAccounts_TableService_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccounts_TableService_Spec tests if a specific instance of StorageAccounts_TableService_Spec can be assigned to v1api20220901storage and back losslessly
func RunPropertyAssignmentTestForStorageAccounts_TableService_Spec(subject StorageAccounts_TableService_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.StorageAccounts_TableService_Spec
	err := copied.AssignProperties_To_StorageAccounts_TableService_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccounts_TableService_Spec
	err = actual.AssignProperties_From_StorageAccounts_TableService_Spec(&other)
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

func Test_StorageAccounts_TableService_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_TableService_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_TableService_Spec, StorageAccounts_TableService_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_TableService_Spec runs a test to see if a specific instance of StorageAccounts_TableService_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_TableService_Spec(subject StorageAccounts_TableService_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_TableService_Spec
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

// Generator of StorageAccounts_TableService_Spec instances for property testing - lazily instantiated by
// StorageAccounts_TableService_SpecGenerator()
var storageAccounts_TableService_SpecGenerator gopter.Gen

// StorageAccounts_TableService_SpecGenerator returns a generator of StorageAccounts_TableService_Spec instances for property testing.
func StorageAccounts_TableService_SpecGenerator() gopter.Gen {
	if storageAccounts_TableService_SpecGenerator != nil {
		return storageAccounts_TableService_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccounts_TableService_Spec(generators)
	storageAccounts_TableService_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_TableService_Spec{}), generators)

	return storageAccounts_TableService_SpecGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccounts_TableService_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_TableService_Spec(gens map[string]gopter.Gen) {
	gens["Cors"] = gen.PtrOf(CorsRulesGenerator())
}

func Test_StorageAccounts_TableService_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccounts_TableService_STATUS to StorageAccounts_TableService_STATUS via AssignProperties_To_StorageAccounts_TableService_STATUS & AssignProperties_From_StorageAccounts_TableService_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccounts_TableService_STATUS, StorageAccounts_TableService_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccounts_TableService_STATUS tests if a specific instance of StorageAccounts_TableService_STATUS can be assigned to v1api20220901storage and back losslessly
func RunPropertyAssignmentTestForStorageAccounts_TableService_STATUS(subject StorageAccounts_TableService_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.StorageAccounts_TableService_STATUS
	err := copied.AssignProperties_To_StorageAccounts_TableService_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccounts_TableService_STATUS
	err = actual.AssignProperties_From_StorageAccounts_TableService_STATUS(&other)
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

func Test_StorageAccounts_TableService_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_TableService_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_TableService_STATUS, StorageAccounts_TableService_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_TableService_STATUS runs a test to see if a specific instance of StorageAccounts_TableService_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_TableService_STATUS(subject StorageAccounts_TableService_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_TableService_STATUS
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

// Generator of StorageAccounts_TableService_STATUS instances for property testing - lazily instantiated by
// StorageAccounts_TableService_STATUSGenerator()
var storageAccounts_TableService_STATUSGenerator gopter.Gen

// StorageAccounts_TableService_STATUSGenerator returns a generator of StorageAccounts_TableService_STATUS instances for property testing.
// We first initialize storageAccounts_TableService_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccounts_TableService_STATUSGenerator() gopter.Gen {
	if storageAccounts_TableService_STATUSGenerator != nil {
		return storageAccounts_TableService_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_TableService_STATUS(generators)
	storageAccounts_TableService_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_TableService_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_TableService_STATUS(generators)
	AddRelatedPropertyGeneratorsForStorageAccounts_TableService_STATUS(generators)
	storageAccounts_TableService_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_TableService_STATUS{}), generators)

	return storageAccounts_TableService_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccounts_TableService_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccounts_TableService_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccounts_TableService_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_TableService_STATUS(gens map[string]gopter.Gen) {
	gens["Cors"] = gen.PtrOf(CorsRules_STATUSGenerator())
}

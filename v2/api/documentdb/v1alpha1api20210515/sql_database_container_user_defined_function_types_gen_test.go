// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"encoding/json"
	alpha20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1alpha1api20210515storage"
	v20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515storage"
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

func Test_SqlDatabaseContainerUserDefinedFunction_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerUserDefinedFunction to hub returns original",
		prop.ForAll(RunResourceConversionTestForSqlDatabaseContainerUserDefinedFunction, SqlDatabaseContainerUserDefinedFunctionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForSqlDatabaseContainerUserDefinedFunction tests if a specific instance of SqlDatabaseContainerUserDefinedFunction round trips to the hub storage version and back losslessly
func RunResourceConversionTestForSqlDatabaseContainerUserDefinedFunction(subject SqlDatabaseContainerUserDefinedFunction) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20210515s.SqlDatabaseContainerUserDefinedFunction
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual SqlDatabaseContainerUserDefinedFunction
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

func Test_SqlDatabaseContainerUserDefinedFunction_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerUserDefinedFunction to SqlDatabaseContainerUserDefinedFunction via AssignPropertiesToSqlDatabaseContainerUserDefinedFunction & AssignPropertiesFromSqlDatabaseContainerUserDefinedFunction returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseContainerUserDefinedFunction, SqlDatabaseContainerUserDefinedFunctionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseContainerUserDefinedFunction tests if a specific instance of SqlDatabaseContainerUserDefinedFunction can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseContainerUserDefinedFunction(subject SqlDatabaseContainerUserDefinedFunction) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20210515s.SqlDatabaseContainerUserDefinedFunction
	err := copied.AssignPropertiesToSqlDatabaseContainerUserDefinedFunction(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseContainerUserDefinedFunction
	err = actual.AssignPropertiesFromSqlDatabaseContainerUserDefinedFunction(&other)
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

func Test_SqlDatabaseContainerUserDefinedFunction_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerUserDefinedFunction via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerUserDefinedFunction, SqlDatabaseContainerUserDefinedFunctionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerUserDefinedFunction runs a test to see if a specific instance of SqlDatabaseContainerUserDefinedFunction round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerUserDefinedFunction(subject SqlDatabaseContainerUserDefinedFunction) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerUserDefinedFunction
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

// Generator of SqlDatabaseContainerUserDefinedFunction instances for property testing - lazily instantiated by
// SqlDatabaseContainerUserDefinedFunctionGenerator()
var sqlDatabaseContainerUserDefinedFunctionGenerator gopter.Gen

// SqlDatabaseContainerUserDefinedFunctionGenerator returns a generator of SqlDatabaseContainerUserDefinedFunction instances for property testing.
func SqlDatabaseContainerUserDefinedFunctionGenerator() gopter.Gen {
	if sqlDatabaseContainerUserDefinedFunctionGenerator != nil {
		return sqlDatabaseContainerUserDefinedFunctionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerUserDefinedFunction(generators)
	sqlDatabaseContainerUserDefinedFunctionGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerUserDefinedFunction{}), generators)

	return sqlDatabaseContainerUserDefinedFunctionGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerUserDefinedFunction is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerUserDefinedFunction(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecGenerator()
	gens["Status"] = SqlUserDefinedFunctionGetResultsStatusGenerator()
}

func Test_DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec to DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec via AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec & AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec, DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec tests if a specific instance of DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec(subject DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec
	err := copied.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec
	err = actual.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec(&other)
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

func Test_DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec, DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec(subject DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec
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

// Generator of DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec instances for property testing - lazily
// instantiated by DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecGenerator()
var databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecGenerator != nil {
		return databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec(generators)
	databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec(generators)
	databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec{}), generators)

	return databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(SqlUserDefinedFunctionResourceGenerator())
}

func Test_SqlUserDefinedFunctionGetResults_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlUserDefinedFunctionGetResults_Status to SqlUserDefinedFunctionGetResults_Status via AssignPropertiesToSqlUserDefinedFunctionGetResultsStatus & AssignPropertiesFromSqlUserDefinedFunctionGetResultsStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlUserDefinedFunctionGetResultsStatus, SqlUserDefinedFunctionGetResultsStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlUserDefinedFunctionGetResultsStatus tests if a specific instance of SqlUserDefinedFunctionGetResults_Status can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlUserDefinedFunctionGetResultsStatus(subject SqlUserDefinedFunctionGetResults_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20210515s.SqlUserDefinedFunctionGetResults_Status
	err := copied.AssignPropertiesToSqlUserDefinedFunctionGetResultsStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlUserDefinedFunctionGetResults_Status
	err = actual.AssignPropertiesFromSqlUserDefinedFunctionGetResultsStatus(&other)
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

func Test_SqlUserDefinedFunctionGetResults_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionGetResults_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionGetResultsStatus, SqlUserDefinedFunctionGetResultsStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionGetResultsStatus runs a test to see if a specific instance of SqlUserDefinedFunctionGetResults_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionGetResultsStatus(subject SqlUserDefinedFunctionGetResults_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionGetResults_Status
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

// Generator of SqlUserDefinedFunctionGetResults_Status instances for property testing - lazily instantiated by
// SqlUserDefinedFunctionGetResultsStatusGenerator()
var sqlUserDefinedFunctionGetResultsStatusGenerator gopter.Gen

// SqlUserDefinedFunctionGetResultsStatusGenerator returns a generator of SqlUserDefinedFunctionGetResults_Status instances for property testing.
// We first initialize sqlUserDefinedFunctionGetResultsStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlUserDefinedFunctionGetResultsStatusGenerator() gopter.Gen {
	if sqlUserDefinedFunctionGetResultsStatusGenerator != nil {
		return sqlUserDefinedFunctionGetResultsStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetResultsStatus(generators)
	sqlUserDefinedFunctionGetResultsStatusGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionGetResults_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetResultsStatus(generators)
	AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionGetResultsStatus(generators)
	sqlUserDefinedFunctionGetResultsStatusGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionGetResults_Status{}), generators)

	return sqlUserDefinedFunctionGetResultsStatusGenerator
}

// AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetResultsStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetResultsStatus(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionGetResultsStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionGetResultsStatus(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(SqlUserDefinedFunctionGetPropertiesStatusResourceGenerator())
}

func Test_SqlUserDefinedFunctionGetProperties_Status_Resource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlUserDefinedFunctionGetProperties_Status_Resource to SqlUserDefinedFunctionGetProperties_Status_Resource via AssignPropertiesToSqlUserDefinedFunctionGetPropertiesStatusResource & AssignPropertiesFromSqlUserDefinedFunctionGetPropertiesStatusResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlUserDefinedFunctionGetPropertiesStatusResource, SqlUserDefinedFunctionGetPropertiesStatusResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlUserDefinedFunctionGetPropertiesStatusResource tests if a specific instance of SqlUserDefinedFunctionGetProperties_Status_Resource can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlUserDefinedFunctionGetPropertiesStatusResource(subject SqlUserDefinedFunctionGetProperties_Status_Resource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20210515s.SqlUserDefinedFunctionGetProperties_Status_Resource
	err := copied.AssignPropertiesToSqlUserDefinedFunctionGetPropertiesStatusResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlUserDefinedFunctionGetProperties_Status_Resource
	err = actual.AssignPropertiesFromSqlUserDefinedFunctionGetPropertiesStatusResource(&other)
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

func Test_SqlUserDefinedFunctionGetProperties_Status_Resource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionGetProperties_Status_Resource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionGetPropertiesStatusResource, SqlUserDefinedFunctionGetPropertiesStatusResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionGetPropertiesStatusResource runs a test to see if a specific instance of SqlUserDefinedFunctionGetProperties_Status_Resource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionGetPropertiesStatusResource(subject SqlUserDefinedFunctionGetProperties_Status_Resource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionGetProperties_Status_Resource
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

// Generator of SqlUserDefinedFunctionGetProperties_Status_Resource instances for property testing - lazily instantiated
// by SqlUserDefinedFunctionGetPropertiesStatusResourceGenerator()
var sqlUserDefinedFunctionGetPropertiesStatusResourceGenerator gopter.Gen

// SqlUserDefinedFunctionGetPropertiesStatusResourceGenerator returns a generator of SqlUserDefinedFunctionGetProperties_Status_Resource instances for property testing.
func SqlUserDefinedFunctionGetPropertiesStatusResourceGenerator() gopter.Gen {
	if sqlUserDefinedFunctionGetPropertiesStatusResourceGenerator != nil {
		return sqlUserDefinedFunctionGetPropertiesStatusResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetPropertiesStatusResource(generators)
	sqlUserDefinedFunctionGetPropertiesStatusResourceGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionGetProperties_Status_Resource{}), generators)

	return sqlUserDefinedFunctionGetPropertiesStatusResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetPropertiesStatusResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionGetPropertiesStatusResource(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

func Test_SqlUserDefinedFunctionResource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlUserDefinedFunctionResource to SqlUserDefinedFunctionResource via AssignPropertiesToSqlUserDefinedFunctionResource & AssignPropertiesFromSqlUserDefinedFunctionResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlUserDefinedFunctionResource, SqlUserDefinedFunctionResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlUserDefinedFunctionResource tests if a specific instance of SqlUserDefinedFunctionResource can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlUserDefinedFunctionResource(subject SqlUserDefinedFunctionResource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20210515s.SqlUserDefinedFunctionResource
	err := copied.AssignPropertiesToSqlUserDefinedFunctionResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlUserDefinedFunctionResource
	err = actual.AssignPropertiesFromSqlUserDefinedFunctionResource(&other)
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

func Test_SqlUserDefinedFunctionResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionResource, SqlUserDefinedFunctionResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionResource runs a test to see if a specific instance of SqlUserDefinedFunctionResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionResource(subject SqlUserDefinedFunctionResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionResource
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

// Generator of SqlUserDefinedFunctionResource instances for property testing - lazily instantiated by
// SqlUserDefinedFunctionResourceGenerator()
var sqlUserDefinedFunctionResourceGenerator gopter.Gen

// SqlUserDefinedFunctionResourceGenerator returns a generator of SqlUserDefinedFunctionResource instances for property testing.
func SqlUserDefinedFunctionResourceGenerator() gopter.Gen {
	if sqlUserDefinedFunctionResourceGenerator != nil {
		return sqlUserDefinedFunctionResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResource(generators)
	sqlUserDefinedFunctionResourceGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionResource{}), generators)

	return sqlUserDefinedFunctionResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResource(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

import (
	"encoding/json"
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

func Test_SqlDatabaseContainerStoredProcedure_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerStoredProcedure to hub returns original",
		prop.ForAll(RunResourceConversionTestForSqlDatabaseContainerStoredProcedure, SqlDatabaseContainerStoredProcedureGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForSqlDatabaseContainerStoredProcedure tests if a specific instance of SqlDatabaseContainerStoredProcedure round trips to the hub storage version and back losslessly
func RunResourceConversionTestForSqlDatabaseContainerStoredProcedure(subject SqlDatabaseContainerStoredProcedure) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20210515s.SqlDatabaseContainerStoredProcedure
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual SqlDatabaseContainerStoredProcedure
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

func Test_SqlDatabaseContainerStoredProcedure_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerStoredProcedure to SqlDatabaseContainerStoredProcedure via AssignProperties_To_SqlDatabaseContainerStoredProcedure & AssignProperties_From_SqlDatabaseContainerStoredProcedure returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedure, SqlDatabaseContainerStoredProcedureGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedure tests if a specific instance of SqlDatabaseContainerStoredProcedure can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedure(subject SqlDatabaseContainerStoredProcedure) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlDatabaseContainerStoredProcedure
	err := copied.AssignProperties_To_SqlDatabaseContainerStoredProcedure(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseContainerStoredProcedure
	err = actual.AssignProperties_From_SqlDatabaseContainerStoredProcedure(&other)
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

func Test_SqlDatabaseContainerStoredProcedure_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerStoredProcedure via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerStoredProcedure, SqlDatabaseContainerStoredProcedureGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerStoredProcedure runs a test to see if a specific instance of SqlDatabaseContainerStoredProcedure round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerStoredProcedure(subject SqlDatabaseContainerStoredProcedure) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerStoredProcedure
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

// Generator of SqlDatabaseContainerStoredProcedure instances for property testing - lazily instantiated by
// SqlDatabaseContainerStoredProcedureGenerator()
var sqlDatabaseContainerStoredProcedureGenerator gopter.Gen

// SqlDatabaseContainerStoredProcedureGenerator returns a generator of SqlDatabaseContainerStoredProcedure instances for property testing.
func SqlDatabaseContainerStoredProcedureGenerator() gopter.Gen {
	if sqlDatabaseContainerStoredProcedureGenerator != nil {
		return sqlDatabaseContainerStoredProcedureGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerStoredProcedure(generators)
	sqlDatabaseContainerStoredProcedureGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerStoredProcedure{}), generators)

	return sqlDatabaseContainerStoredProcedureGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerStoredProcedure is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerStoredProcedure(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_SpecGenerator()
	gens["Status"] = SqlStoredProcedureGetResults_STATUSGenerator()
}

func Test_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec to DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec via AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec & AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec, DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec tests if a specific instance of DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(subject DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec
	err := copied.AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec
	err = actual.AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(&other)
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

func Test_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec, DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec runs a test to see if a specific instance of DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(subject DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec
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

// Generator of DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec instances for property testing - lazily
// instantiated by DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_SpecGenerator()
var databaseAccounts_SqlDatabases_Containers_StoredProcedure_SpecGenerator gopter.Gen

// DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_SpecGenerator returns a generator of DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec instances for property testing.
// We first initialize databaseAccounts_SqlDatabases_Containers_StoredProcedure_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_SpecGenerator() gopter.Gen {
	if databaseAccounts_SqlDatabases_Containers_StoredProcedure_SpecGenerator != nil {
		return databaseAccounts_SqlDatabases_Containers_StoredProcedure_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(generators)
	databaseAccounts_SqlDatabases_Containers_StoredProcedure_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(generators)
	databaseAccounts_SqlDatabases_Containers_StoredProcedure_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec{}), generators)

	return databaseAccounts_SqlDatabases_Containers_StoredProcedure_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(SqlStoredProcedureResourceGenerator())
}

func Test_SqlStoredProcedureGetResults_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlStoredProcedureGetResults_STATUS to SqlStoredProcedureGetResults_STATUS via AssignProperties_To_SqlStoredProcedureGetResults_STATUS & AssignProperties_From_SqlStoredProcedureGetResults_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlStoredProcedureGetResults_STATUS, SqlStoredProcedureGetResults_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlStoredProcedureGetResults_STATUS tests if a specific instance of SqlStoredProcedureGetResults_STATUS can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlStoredProcedureGetResults_STATUS(subject SqlStoredProcedureGetResults_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlStoredProcedureGetResults_STATUS
	err := copied.AssignProperties_To_SqlStoredProcedureGetResults_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlStoredProcedureGetResults_STATUS
	err = actual.AssignProperties_From_SqlStoredProcedureGetResults_STATUS(&other)
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

func Test_SqlStoredProcedureGetResults_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureGetResults_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureGetResults_STATUS, SqlStoredProcedureGetResults_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureGetResults_STATUS runs a test to see if a specific instance of SqlStoredProcedureGetResults_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureGetResults_STATUS(subject SqlStoredProcedureGetResults_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureGetResults_STATUS
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

// Generator of SqlStoredProcedureGetResults_STATUS instances for property testing - lazily instantiated by
// SqlStoredProcedureGetResults_STATUSGenerator()
var sqlStoredProcedureGetResults_STATUSGenerator gopter.Gen

// SqlStoredProcedureGetResults_STATUSGenerator returns a generator of SqlStoredProcedureGetResults_STATUS instances for property testing.
// We first initialize sqlStoredProcedureGetResults_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlStoredProcedureGetResults_STATUSGenerator() gopter.Gen {
	if sqlStoredProcedureGetResults_STATUSGenerator != nil {
		return sqlStoredProcedureGetResults_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureGetResults_STATUS(generators)
	sqlStoredProcedureGetResults_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureGetResults_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureGetResults_STATUS(generators)
	AddRelatedPropertyGeneratorsForSqlStoredProcedureGetResults_STATUS(generators)
	sqlStoredProcedureGetResults_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureGetResults_STATUS{}), generators)

	return sqlStoredProcedureGetResults_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureGetResults_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureGetResults_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlStoredProcedureGetResults_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlStoredProcedureGetResults_STATUS(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(SqlStoredProcedureGetProperties_STATUS_ResourceGenerator())
}

func Test_SqlStoredProcedureGetProperties_STATUS_Resource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlStoredProcedureGetProperties_STATUS_Resource to SqlStoredProcedureGetProperties_STATUS_Resource via AssignProperties_To_SqlStoredProcedureGetProperties_STATUS_Resource & AssignProperties_From_SqlStoredProcedureGetProperties_STATUS_Resource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlStoredProcedureGetProperties_STATUS_Resource, SqlStoredProcedureGetProperties_STATUS_ResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlStoredProcedureGetProperties_STATUS_Resource tests if a specific instance of SqlStoredProcedureGetProperties_STATUS_Resource can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlStoredProcedureGetProperties_STATUS_Resource(subject SqlStoredProcedureGetProperties_STATUS_Resource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlStoredProcedureGetProperties_STATUS_Resource
	err := copied.AssignProperties_To_SqlStoredProcedureGetProperties_STATUS_Resource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlStoredProcedureGetProperties_STATUS_Resource
	err = actual.AssignProperties_From_SqlStoredProcedureGetProperties_STATUS_Resource(&other)
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

func Test_SqlStoredProcedureGetProperties_STATUS_Resource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureGetProperties_STATUS_Resource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureGetProperties_STATUS_Resource, SqlStoredProcedureGetProperties_STATUS_ResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureGetProperties_STATUS_Resource runs a test to see if a specific instance of SqlStoredProcedureGetProperties_STATUS_Resource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureGetProperties_STATUS_Resource(subject SqlStoredProcedureGetProperties_STATUS_Resource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureGetProperties_STATUS_Resource
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

// Generator of SqlStoredProcedureGetProperties_STATUS_Resource instances for property testing - lazily instantiated by
// SqlStoredProcedureGetProperties_STATUS_ResourceGenerator()
var sqlStoredProcedureGetProperties_STATUS_ResourceGenerator gopter.Gen

// SqlStoredProcedureGetProperties_STATUS_ResourceGenerator returns a generator of SqlStoredProcedureGetProperties_STATUS_Resource instances for property testing.
func SqlStoredProcedureGetProperties_STATUS_ResourceGenerator() gopter.Gen {
	if sqlStoredProcedureGetProperties_STATUS_ResourceGenerator != nil {
		return sqlStoredProcedureGetProperties_STATUS_ResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureGetProperties_STATUS_Resource(generators)
	sqlStoredProcedureGetProperties_STATUS_ResourceGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureGetProperties_STATUS_Resource{}), generators)

	return sqlStoredProcedureGetProperties_STATUS_ResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureGetProperties_STATUS_Resource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureGetProperties_STATUS_Resource(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

func Test_SqlStoredProcedureResource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlStoredProcedureResource to SqlStoredProcedureResource via AssignProperties_To_SqlStoredProcedureResource & AssignProperties_From_SqlStoredProcedureResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlStoredProcedureResource, SqlStoredProcedureResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlStoredProcedureResource tests if a specific instance of SqlStoredProcedureResource can be assigned to v1beta20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlStoredProcedureResource(subject SqlStoredProcedureResource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlStoredProcedureResource
	err := copied.AssignProperties_To_SqlStoredProcedureResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlStoredProcedureResource
	err = actual.AssignProperties_From_SqlStoredProcedureResource(&other)
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

func Test_SqlStoredProcedureResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureResource, SqlStoredProcedureResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureResource runs a test to see if a specific instance of SqlStoredProcedureResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureResource(subject SqlStoredProcedureResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureResource
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

// Generator of SqlStoredProcedureResource instances for property testing - lazily instantiated by
// SqlStoredProcedureResourceGenerator()
var sqlStoredProcedureResourceGenerator gopter.Gen

// SqlStoredProcedureResourceGenerator returns a generator of SqlStoredProcedureResource instances for property testing.
func SqlStoredProcedureResourceGenerator() gopter.Gen {
	if sqlStoredProcedureResourceGenerator != nil {
		return sqlStoredProcedureResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureResource(generators)
	sqlStoredProcedureResourceGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureResource{}), generators)

	return sqlStoredProcedureResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureResource(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230315preview

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

func Test_ScheduledQueryRule_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduledQueryRule_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduledQueryRule_STATUS_ARM, ScheduledQueryRule_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduledQueryRule_STATUS_ARM runs a test to see if a specific instance of ScheduledQueryRule_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduledQueryRule_STATUS_ARM(subject ScheduledQueryRule_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduledQueryRule_STATUS_ARM
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

// Generator of ScheduledQueryRule_STATUS_ARM instances for property testing - lazily instantiated by
// ScheduledQueryRule_STATUS_ARMGenerator()
var scheduledQueryRule_STATUS_ARMGenerator gopter.Gen

// ScheduledQueryRule_STATUS_ARMGenerator returns a generator of ScheduledQueryRule_STATUS_ARM instances for property testing.
// We first initialize scheduledQueryRule_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ScheduledQueryRule_STATUS_ARMGenerator() gopter.Gen {
	if scheduledQueryRule_STATUS_ARMGenerator != nil {
		return scheduledQueryRule_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduledQueryRule_STATUS_ARM(generators)
	scheduledQueryRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRule_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduledQueryRule_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForScheduledQueryRule_STATUS_ARM(generators)
	scheduledQueryRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRule_STATUS_ARM{}), generators)

	return scheduledQueryRule_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForScheduledQueryRule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScheduledQueryRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.OneConstOf(ScheduledQueryRule_Kind_STATUS_LogAlert, ScheduledQueryRule_Kind_STATUS_LogToMetric))
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForScheduledQueryRule_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScheduledQueryRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(Identity_STATUS_ARMGenerator())
	gens["Properties"] = gen.PtrOf(ScheduledQueryRuleProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_Identity_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentity_STATUS_ARM, Identity_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentity_STATUS_ARM runs a test to see if a specific instance of Identity_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentity_STATUS_ARM(subject Identity_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity_STATUS_ARM
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

// Generator of Identity_STATUS_ARM instances for property testing - lazily instantiated by
// Identity_STATUS_ARMGenerator()
var identity_STATUS_ARMGenerator gopter.Gen

// Identity_STATUS_ARMGenerator returns a generator of Identity_STATUS_ARM instances for property testing.
// We first initialize identity_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Identity_STATUS_ARMGenerator() gopter.Gen {
	if identity_STATUS_ARMGenerator != nil {
		return identity_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity_STATUS_ARM(generators)
	identity_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Identity_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForIdentity_STATUS_ARM(generators)
	identity_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Identity_STATUS_ARM{}), generators)

	return identity_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentity_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentity_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(Identity_Type_STATUS_None, Identity_Type_STATUS_SystemAssigned, Identity_Type_STATUS_UserAssigned))
}

// AddRelatedPropertyGeneratorsForIdentity_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIdentity_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(gen.AlphaString(), UserIdentityProperties_STATUS_ARMGenerator())
}

func Test_ScheduledQueryRuleProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduledQueryRuleProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduledQueryRuleProperties_STATUS_ARM, ScheduledQueryRuleProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduledQueryRuleProperties_STATUS_ARM runs a test to see if a specific instance of ScheduledQueryRuleProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduledQueryRuleProperties_STATUS_ARM(subject ScheduledQueryRuleProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduledQueryRuleProperties_STATUS_ARM
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

// Generator of ScheduledQueryRuleProperties_STATUS_ARM instances for property testing - lazily instantiated by
// ScheduledQueryRuleProperties_STATUS_ARMGenerator()
var scheduledQueryRuleProperties_STATUS_ARMGenerator gopter.Gen

// ScheduledQueryRuleProperties_STATUS_ARMGenerator returns a generator of ScheduledQueryRuleProperties_STATUS_ARM instances for property testing.
// We first initialize scheduledQueryRuleProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ScheduledQueryRuleProperties_STATUS_ARMGenerator() gopter.Gen {
	if scheduledQueryRuleProperties_STATUS_ARMGenerator != nil {
		return scheduledQueryRuleProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduledQueryRuleProperties_STATUS_ARM(generators)
	scheduledQueryRuleProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRuleProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduledQueryRuleProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForScheduledQueryRuleProperties_STATUS_ARM(generators)
	scheduledQueryRuleProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRuleProperties_STATUS_ARM{}), generators)

	return scheduledQueryRuleProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForScheduledQueryRuleProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScheduledQueryRuleProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AutoMitigate"] = gen.PtrOf(gen.Bool())
	gens["CheckWorkspaceAlertsStorageConfigured"] = gen.PtrOf(gen.Bool())
	gens["CreatedWithApiVersion"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["EvaluationFrequency"] = gen.PtrOf(gen.AlphaString())
	gens["IsLegacyLogAnalyticsRule"] = gen.PtrOf(gen.Bool())
	gens["IsWorkspaceAlertsStorageConfigured"] = gen.PtrOf(gen.Bool())
	gens["MuteActionsDuration"] = gen.PtrOf(gen.AlphaString())
	gens["OverrideQueryTimeRange"] = gen.PtrOf(gen.AlphaString())
	gens["Scopes"] = gen.SliceOf(gen.AlphaString())
	gens["Severity"] = gen.PtrOf(gen.OneConstOf(
		ScheduledQueryRuleProperties_Severity_STATUS_0,
		ScheduledQueryRuleProperties_Severity_STATUS_1,
		ScheduledQueryRuleProperties_Severity_STATUS_2,
		ScheduledQueryRuleProperties_Severity_STATUS_3,
		ScheduledQueryRuleProperties_Severity_STATUS_4))
	gens["SkipQueryValidation"] = gen.PtrOf(gen.Bool())
	gens["TargetResourceTypes"] = gen.SliceOf(gen.AlphaString())
	gens["WindowSize"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForScheduledQueryRuleProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScheduledQueryRuleProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Actions"] = gen.PtrOf(Actions_STATUS_ARMGenerator())
	gens["Criteria"] = gen.PtrOf(ScheduledQueryRuleCriteria_STATUS_ARMGenerator())
	gens["RuleResolveConfiguration"] = gen.PtrOf(RuleResolveConfiguration_STATUS_ARMGenerator())
}

func Test_SystemData_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUS_ARM, SystemData_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUS_ARM runs a test to see if a specific instance of SystemData_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUS_ARM(subject SystemData_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUS_ARM
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

// Generator of SystemData_STATUS_ARM instances for property testing - lazily instantiated by
// SystemData_STATUS_ARMGenerator()
var systemData_STATUS_ARMGenerator gopter.Gen

// SystemData_STATUS_ARMGenerator returns a generator of SystemData_STATUS_ARM instances for property testing.
func SystemData_STATUS_ARMGenerator() gopter.Gen {
	if systemData_STATUS_ARMGenerator != nil {
		return systemData_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM(generators)
	systemData_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUS_ARM{}), generators)

	return systemData_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUS_ARM(gens map[string]gopter.Gen) {
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

func Test_Actions_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Actions_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForActions_STATUS_ARM, Actions_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForActions_STATUS_ARM runs a test to see if a specific instance of Actions_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForActions_STATUS_ARM(subject Actions_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Actions_STATUS_ARM
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

// Generator of Actions_STATUS_ARM instances for property testing - lazily instantiated by Actions_STATUS_ARMGenerator()
var actions_STATUS_ARMGenerator gopter.Gen

// Actions_STATUS_ARMGenerator returns a generator of Actions_STATUS_ARM instances for property testing.
func Actions_STATUS_ARMGenerator() gopter.Gen {
	if actions_STATUS_ARMGenerator != nil {
		return actions_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForActions_STATUS_ARM(generators)
	actions_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Actions_STATUS_ARM{}), generators)

	return actions_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForActions_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForActions_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ActionGroups"] = gen.SliceOf(gen.AlphaString())
	gens["ActionProperties"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["CustomProperties"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_RuleResolveConfiguration_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RuleResolveConfiguration_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRuleResolveConfiguration_STATUS_ARM, RuleResolveConfiguration_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRuleResolveConfiguration_STATUS_ARM runs a test to see if a specific instance of RuleResolveConfiguration_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRuleResolveConfiguration_STATUS_ARM(subject RuleResolveConfiguration_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RuleResolveConfiguration_STATUS_ARM
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

// Generator of RuleResolveConfiguration_STATUS_ARM instances for property testing - lazily instantiated by
// RuleResolveConfiguration_STATUS_ARMGenerator()
var ruleResolveConfiguration_STATUS_ARMGenerator gopter.Gen

// RuleResolveConfiguration_STATUS_ARMGenerator returns a generator of RuleResolveConfiguration_STATUS_ARM instances for property testing.
func RuleResolveConfiguration_STATUS_ARMGenerator() gopter.Gen {
	if ruleResolveConfiguration_STATUS_ARMGenerator != nil {
		return ruleResolveConfiguration_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRuleResolveConfiguration_STATUS_ARM(generators)
	ruleResolveConfiguration_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RuleResolveConfiguration_STATUS_ARM{}), generators)

	return ruleResolveConfiguration_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRuleResolveConfiguration_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRuleResolveConfiguration_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AutoResolved"] = gen.PtrOf(gen.Bool())
	gens["TimeToResolve"] = gen.PtrOf(gen.AlphaString())
}

func Test_ScheduledQueryRuleCriteria_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduledQueryRuleCriteria_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduledQueryRuleCriteria_STATUS_ARM, ScheduledQueryRuleCriteria_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduledQueryRuleCriteria_STATUS_ARM runs a test to see if a specific instance of ScheduledQueryRuleCriteria_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduledQueryRuleCriteria_STATUS_ARM(subject ScheduledQueryRuleCriteria_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduledQueryRuleCriteria_STATUS_ARM
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

// Generator of ScheduledQueryRuleCriteria_STATUS_ARM instances for property testing - lazily instantiated by
// ScheduledQueryRuleCriteria_STATUS_ARMGenerator()
var scheduledQueryRuleCriteria_STATUS_ARMGenerator gopter.Gen

// ScheduledQueryRuleCriteria_STATUS_ARMGenerator returns a generator of ScheduledQueryRuleCriteria_STATUS_ARM instances for property testing.
func ScheduledQueryRuleCriteria_STATUS_ARMGenerator() gopter.Gen {
	if scheduledQueryRuleCriteria_STATUS_ARMGenerator != nil {
		return scheduledQueryRuleCriteria_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForScheduledQueryRuleCriteria_STATUS_ARM(generators)
	scheduledQueryRuleCriteria_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ScheduledQueryRuleCriteria_STATUS_ARM{}), generators)

	return scheduledQueryRuleCriteria_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForScheduledQueryRuleCriteria_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForScheduledQueryRuleCriteria_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AllOf"] = gen.SliceOf(Condition_STATUS_ARMGenerator())
}

func Test_UserIdentityProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserIdentityProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserIdentityProperties_STATUS_ARM, UserIdentityProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserIdentityProperties_STATUS_ARM runs a test to see if a specific instance of UserIdentityProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserIdentityProperties_STATUS_ARM(subject UserIdentityProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserIdentityProperties_STATUS_ARM
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

// Generator of UserIdentityProperties_STATUS_ARM instances for property testing - lazily instantiated by
// UserIdentityProperties_STATUS_ARMGenerator()
var userIdentityProperties_STATUS_ARMGenerator gopter.Gen

// UserIdentityProperties_STATUS_ARMGenerator returns a generator of UserIdentityProperties_STATUS_ARM instances for property testing.
func UserIdentityProperties_STATUS_ARMGenerator() gopter.Gen {
	if userIdentityProperties_STATUS_ARMGenerator != nil {
		return userIdentityProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserIdentityProperties_STATUS_ARM(generators)
	userIdentityProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(UserIdentityProperties_STATUS_ARM{}), generators)

	return userIdentityProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForUserIdentityProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserIdentityProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
}

func Test_Condition_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Condition_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCondition_STATUS_ARM, Condition_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCondition_STATUS_ARM runs a test to see if a specific instance of Condition_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCondition_STATUS_ARM(subject Condition_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Condition_STATUS_ARM
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

// Generator of Condition_STATUS_ARM instances for property testing - lazily instantiated by
// Condition_STATUS_ARMGenerator()
var condition_STATUS_ARMGenerator gopter.Gen

// Condition_STATUS_ARMGenerator returns a generator of Condition_STATUS_ARM instances for property testing.
// We first initialize condition_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Condition_STATUS_ARMGenerator() gopter.Gen {
	if condition_STATUS_ARMGenerator != nil {
		return condition_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCondition_STATUS_ARM(generators)
	condition_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Condition_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCondition_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForCondition_STATUS_ARM(generators)
	condition_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Condition_STATUS_ARM{}), generators)

	return condition_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCondition_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCondition_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["MetricMeasureColumn"] = gen.PtrOf(gen.AlphaString())
	gens["MetricName"] = gen.PtrOf(gen.AlphaString())
	gens["Operator"] = gen.PtrOf(gen.OneConstOf(
		Condition_Operator_STATUS_Equals,
		Condition_Operator_STATUS_GreaterThan,
		Condition_Operator_STATUS_GreaterThanOrEqual,
		Condition_Operator_STATUS_LessThan,
		Condition_Operator_STATUS_LessThanOrEqual))
	gens["Query"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceIdColumn"] = gen.PtrOf(gen.AlphaString())
	gens["Threshold"] = gen.PtrOf(gen.Float64())
	gens["TimeAggregation"] = gen.PtrOf(gen.OneConstOf(
		Condition_TimeAggregation_STATUS_Average,
		Condition_TimeAggregation_STATUS_Count,
		Condition_TimeAggregation_STATUS_Maximum,
		Condition_TimeAggregation_STATUS_Minimum,
		Condition_TimeAggregation_STATUS_Total))
}

// AddRelatedPropertyGeneratorsForCondition_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCondition_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Dimensions"] = gen.SliceOf(Dimension_STATUS_ARMGenerator())
	gens["FailingPeriods"] = gen.PtrOf(Condition_FailingPeriods_STATUS_ARMGenerator())
}

func Test_Condition_FailingPeriods_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Condition_FailingPeriods_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCondition_FailingPeriods_STATUS_ARM, Condition_FailingPeriods_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCondition_FailingPeriods_STATUS_ARM runs a test to see if a specific instance of Condition_FailingPeriods_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCondition_FailingPeriods_STATUS_ARM(subject Condition_FailingPeriods_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Condition_FailingPeriods_STATUS_ARM
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

// Generator of Condition_FailingPeriods_STATUS_ARM instances for property testing - lazily instantiated by
// Condition_FailingPeriods_STATUS_ARMGenerator()
var condition_FailingPeriods_STATUS_ARMGenerator gopter.Gen

// Condition_FailingPeriods_STATUS_ARMGenerator returns a generator of Condition_FailingPeriods_STATUS_ARM instances for property testing.
func Condition_FailingPeriods_STATUS_ARMGenerator() gopter.Gen {
	if condition_FailingPeriods_STATUS_ARMGenerator != nil {
		return condition_FailingPeriods_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCondition_FailingPeriods_STATUS_ARM(generators)
	condition_FailingPeriods_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Condition_FailingPeriods_STATUS_ARM{}), generators)

	return condition_FailingPeriods_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForCondition_FailingPeriods_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCondition_FailingPeriods_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["MinFailingPeriodsToAlert"] = gen.PtrOf(gen.Int())
	gens["NumberOfEvaluationPeriods"] = gen.PtrOf(gen.Int())
}

func Test_Dimension_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Dimension_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDimension_STATUS_ARM, Dimension_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDimension_STATUS_ARM runs a test to see if a specific instance of Dimension_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDimension_STATUS_ARM(subject Dimension_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Dimension_STATUS_ARM
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

// Generator of Dimension_STATUS_ARM instances for property testing - lazily instantiated by
// Dimension_STATUS_ARMGenerator()
var dimension_STATUS_ARMGenerator gopter.Gen

// Dimension_STATUS_ARMGenerator returns a generator of Dimension_STATUS_ARM instances for property testing.
func Dimension_STATUS_ARMGenerator() gopter.Gen {
	if dimension_STATUS_ARMGenerator != nil {
		return dimension_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDimension_STATUS_ARM(generators)
	dimension_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Dimension_STATUS_ARM{}), generators)

	return dimension_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDimension_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDimension_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Operator"] = gen.PtrOf(gen.OneConstOf(Dimension_Operator_STATUS_Exclude, Dimension_Operator_STATUS_Include))
	gens["Values"] = gen.SliceOf(gen.AlphaString())
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230101

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

func Test_ActionGroupResource_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ActionGroupResource_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForActionGroupResource_STATUS_ARM, ActionGroupResource_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForActionGroupResource_STATUS_ARM runs a test to see if a specific instance of ActionGroupResource_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForActionGroupResource_STATUS_ARM(subject ActionGroupResource_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ActionGroupResource_STATUS_ARM
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

// Generator of ActionGroupResource_STATUS_ARM instances for property testing - lazily instantiated by
// ActionGroupResource_STATUS_ARMGenerator()
var actionGroupResource_STATUS_ARMGenerator gopter.Gen

// ActionGroupResource_STATUS_ARMGenerator returns a generator of ActionGroupResource_STATUS_ARM instances for property testing.
// We first initialize actionGroupResource_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ActionGroupResource_STATUS_ARMGenerator() gopter.Gen {
	if actionGroupResource_STATUS_ARMGenerator != nil {
		return actionGroupResource_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForActionGroupResource_STATUS_ARM(generators)
	actionGroupResource_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ActionGroupResource_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForActionGroupResource_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForActionGroupResource_STATUS_ARM(generators)
	actionGroupResource_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ActionGroupResource_STATUS_ARM{}), generators)

	return actionGroupResource_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForActionGroupResource_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForActionGroupResource_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForActionGroupResource_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForActionGroupResource_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ActionGroup_STATUS_ARMGenerator())
}

func Test_ActionGroup_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ActionGroup_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForActionGroup_STATUS_ARM, ActionGroup_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForActionGroup_STATUS_ARM runs a test to see if a specific instance of ActionGroup_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForActionGroup_STATUS_ARM(subject ActionGroup_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ActionGroup_STATUS_ARM
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

// Generator of ActionGroup_STATUS_ARM instances for property testing - lazily instantiated by
// ActionGroup_STATUS_ARMGenerator()
var actionGroup_STATUS_ARMGenerator gopter.Gen

// ActionGroup_STATUS_ARMGenerator returns a generator of ActionGroup_STATUS_ARM instances for property testing.
// We first initialize actionGroup_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ActionGroup_STATUS_ARMGenerator() gopter.Gen {
	if actionGroup_STATUS_ARMGenerator != nil {
		return actionGroup_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForActionGroup_STATUS_ARM(generators)
	actionGroup_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ActionGroup_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForActionGroup_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForActionGroup_STATUS_ARM(generators)
	actionGroup_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ActionGroup_STATUS_ARM{}), generators)

	return actionGroup_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForActionGroup_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForActionGroup_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["GroupShortName"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForActionGroup_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForActionGroup_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ArmRoleReceivers"] = gen.SliceOf(ArmRoleReceiver_STATUS_ARMGenerator())
	gens["AutomationRunbookReceivers"] = gen.SliceOf(AutomationRunbookReceiver_STATUS_ARMGenerator())
	gens["AzureAppPushReceivers"] = gen.SliceOf(AzureAppPushReceiver_STATUS_ARMGenerator())
	gens["AzureFunctionReceivers"] = gen.SliceOf(AzureFunctionReceiver_STATUS_ARMGenerator())
	gens["EmailReceivers"] = gen.SliceOf(EmailReceiver_STATUS_ARMGenerator())
	gens["EventHubReceivers"] = gen.SliceOf(EventHubReceiver_STATUS_ARMGenerator())
	gens["ItsmReceivers"] = gen.SliceOf(ItsmReceiver_STATUS_ARMGenerator())
	gens["LogicAppReceivers"] = gen.SliceOf(LogicAppReceiver_STATUS_ARMGenerator())
	gens["SmsReceivers"] = gen.SliceOf(SmsReceiver_STATUS_ARMGenerator())
	gens["VoiceReceivers"] = gen.SliceOf(VoiceReceiver_STATUS_ARMGenerator())
	gens["WebhookReceivers"] = gen.SliceOf(WebhookReceiver_STATUS_ARMGenerator())
}

func Test_ArmRoleReceiver_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ArmRoleReceiver_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForArmRoleReceiver_STATUS_ARM, ArmRoleReceiver_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForArmRoleReceiver_STATUS_ARM runs a test to see if a specific instance of ArmRoleReceiver_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForArmRoleReceiver_STATUS_ARM(subject ArmRoleReceiver_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ArmRoleReceiver_STATUS_ARM
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

// Generator of ArmRoleReceiver_STATUS_ARM instances for property testing - lazily instantiated by
// ArmRoleReceiver_STATUS_ARMGenerator()
var armRoleReceiver_STATUS_ARMGenerator gopter.Gen

// ArmRoleReceiver_STATUS_ARMGenerator returns a generator of ArmRoleReceiver_STATUS_ARM instances for property testing.
func ArmRoleReceiver_STATUS_ARMGenerator() gopter.Gen {
	if armRoleReceiver_STATUS_ARMGenerator != nil {
		return armRoleReceiver_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForArmRoleReceiver_STATUS_ARM(generators)
	armRoleReceiver_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ArmRoleReceiver_STATUS_ARM{}), generators)

	return armRoleReceiver_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForArmRoleReceiver_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForArmRoleReceiver_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["RoleId"] = gen.PtrOf(gen.AlphaString())
	gens["UseCommonAlertSchema"] = gen.PtrOf(gen.Bool())
}

func Test_AutomationRunbookReceiver_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutomationRunbookReceiver_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutomationRunbookReceiver_STATUS_ARM, AutomationRunbookReceiver_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutomationRunbookReceiver_STATUS_ARM runs a test to see if a specific instance of AutomationRunbookReceiver_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutomationRunbookReceiver_STATUS_ARM(subject AutomationRunbookReceiver_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutomationRunbookReceiver_STATUS_ARM
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

// Generator of AutomationRunbookReceiver_STATUS_ARM instances for property testing - lazily instantiated by
// AutomationRunbookReceiver_STATUS_ARMGenerator()
var automationRunbookReceiver_STATUS_ARMGenerator gopter.Gen

// AutomationRunbookReceiver_STATUS_ARMGenerator returns a generator of AutomationRunbookReceiver_STATUS_ARM instances for property testing.
func AutomationRunbookReceiver_STATUS_ARMGenerator() gopter.Gen {
	if automationRunbookReceiver_STATUS_ARMGenerator != nil {
		return automationRunbookReceiver_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutomationRunbookReceiver_STATUS_ARM(generators)
	automationRunbookReceiver_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AutomationRunbookReceiver_STATUS_ARM{}), generators)

	return automationRunbookReceiver_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAutomationRunbookReceiver_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutomationRunbookReceiver_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AutomationAccountId"] = gen.PtrOf(gen.AlphaString())
	gens["IsGlobalRunbook"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["RunbookName"] = gen.PtrOf(gen.AlphaString())
	gens["ServiceUri"] = gen.PtrOf(gen.AlphaString())
	gens["UseCommonAlertSchema"] = gen.PtrOf(gen.Bool())
	gens["WebhookResourceId"] = gen.PtrOf(gen.AlphaString())
}

func Test_AzureAppPushReceiver_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AzureAppPushReceiver_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAzureAppPushReceiver_STATUS_ARM, AzureAppPushReceiver_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAzureAppPushReceiver_STATUS_ARM runs a test to see if a specific instance of AzureAppPushReceiver_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAzureAppPushReceiver_STATUS_ARM(subject AzureAppPushReceiver_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AzureAppPushReceiver_STATUS_ARM
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

// Generator of AzureAppPushReceiver_STATUS_ARM instances for property testing - lazily instantiated by
// AzureAppPushReceiver_STATUS_ARMGenerator()
var azureAppPushReceiver_STATUS_ARMGenerator gopter.Gen

// AzureAppPushReceiver_STATUS_ARMGenerator returns a generator of AzureAppPushReceiver_STATUS_ARM instances for property testing.
func AzureAppPushReceiver_STATUS_ARMGenerator() gopter.Gen {
	if azureAppPushReceiver_STATUS_ARMGenerator != nil {
		return azureAppPushReceiver_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAzureAppPushReceiver_STATUS_ARM(generators)
	azureAppPushReceiver_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AzureAppPushReceiver_STATUS_ARM{}), generators)

	return azureAppPushReceiver_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAzureAppPushReceiver_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAzureAppPushReceiver_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["EmailAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_AzureFunctionReceiver_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AzureFunctionReceiver_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAzureFunctionReceiver_STATUS_ARM, AzureFunctionReceiver_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAzureFunctionReceiver_STATUS_ARM runs a test to see if a specific instance of AzureFunctionReceiver_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAzureFunctionReceiver_STATUS_ARM(subject AzureFunctionReceiver_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AzureFunctionReceiver_STATUS_ARM
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

// Generator of AzureFunctionReceiver_STATUS_ARM instances for property testing - lazily instantiated by
// AzureFunctionReceiver_STATUS_ARMGenerator()
var azureFunctionReceiver_STATUS_ARMGenerator gopter.Gen

// AzureFunctionReceiver_STATUS_ARMGenerator returns a generator of AzureFunctionReceiver_STATUS_ARM instances for property testing.
func AzureFunctionReceiver_STATUS_ARMGenerator() gopter.Gen {
	if azureFunctionReceiver_STATUS_ARMGenerator != nil {
		return azureFunctionReceiver_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAzureFunctionReceiver_STATUS_ARM(generators)
	azureFunctionReceiver_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(AzureFunctionReceiver_STATUS_ARM{}), generators)

	return azureFunctionReceiver_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAzureFunctionReceiver_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAzureFunctionReceiver_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["FunctionAppResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["FunctionName"] = gen.PtrOf(gen.AlphaString())
	gens["HttpTriggerUrl"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["UseCommonAlertSchema"] = gen.PtrOf(gen.Bool())
}

func Test_EmailReceiver_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EmailReceiver_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEmailReceiver_STATUS_ARM, EmailReceiver_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEmailReceiver_STATUS_ARM runs a test to see if a specific instance of EmailReceiver_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEmailReceiver_STATUS_ARM(subject EmailReceiver_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EmailReceiver_STATUS_ARM
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

// Generator of EmailReceiver_STATUS_ARM instances for property testing - lazily instantiated by
// EmailReceiver_STATUS_ARMGenerator()
var emailReceiver_STATUS_ARMGenerator gopter.Gen

// EmailReceiver_STATUS_ARMGenerator returns a generator of EmailReceiver_STATUS_ARM instances for property testing.
func EmailReceiver_STATUS_ARMGenerator() gopter.Gen {
	if emailReceiver_STATUS_ARMGenerator != nil {
		return emailReceiver_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEmailReceiver_STATUS_ARM(generators)
	emailReceiver_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(EmailReceiver_STATUS_ARM{}), generators)

	return emailReceiver_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForEmailReceiver_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEmailReceiver_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["EmailAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(ReceiverStatus_STATUS_ARM_Disabled, ReceiverStatus_STATUS_ARM_Enabled, ReceiverStatus_STATUS_ARM_NotSpecified))
	gens["UseCommonAlertSchema"] = gen.PtrOf(gen.Bool())
}

func Test_EventHubReceiver_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EventHubReceiver_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventHubReceiver_STATUS_ARM, EventHubReceiver_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventHubReceiver_STATUS_ARM runs a test to see if a specific instance of EventHubReceiver_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEventHubReceiver_STATUS_ARM(subject EventHubReceiver_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EventHubReceiver_STATUS_ARM
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

// Generator of EventHubReceiver_STATUS_ARM instances for property testing - lazily instantiated by
// EventHubReceiver_STATUS_ARMGenerator()
var eventHubReceiver_STATUS_ARMGenerator gopter.Gen

// EventHubReceiver_STATUS_ARMGenerator returns a generator of EventHubReceiver_STATUS_ARM instances for property testing.
func EventHubReceiver_STATUS_ARMGenerator() gopter.Gen {
	if eventHubReceiver_STATUS_ARMGenerator != nil {
		return eventHubReceiver_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventHubReceiver_STATUS_ARM(generators)
	eventHubReceiver_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(EventHubReceiver_STATUS_ARM{}), generators)

	return eventHubReceiver_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForEventHubReceiver_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventHubReceiver_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["EventHubName"] = gen.PtrOf(gen.AlphaString())
	gens["EventHubNameSpace"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["UseCommonAlertSchema"] = gen.PtrOf(gen.Bool())
}

func Test_ItsmReceiver_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ItsmReceiver_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForItsmReceiver_STATUS_ARM, ItsmReceiver_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForItsmReceiver_STATUS_ARM runs a test to see if a specific instance of ItsmReceiver_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForItsmReceiver_STATUS_ARM(subject ItsmReceiver_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ItsmReceiver_STATUS_ARM
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

// Generator of ItsmReceiver_STATUS_ARM instances for property testing - lazily instantiated by
// ItsmReceiver_STATUS_ARMGenerator()
var itsmReceiver_STATUS_ARMGenerator gopter.Gen

// ItsmReceiver_STATUS_ARMGenerator returns a generator of ItsmReceiver_STATUS_ARM instances for property testing.
func ItsmReceiver_STATUS_ARMGenerator() gopter.Gen {
	if itsmReceiver_STATUS_ARMGenerator != nil {
		return itsmReceiver_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForItsmReceiver_STATUS_ARM(generators)
	itsmReceiver_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ItsmReceiver_STATUS_ARM{}), generators)

	return itsmReceiver_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForItsmReceiver_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForItsmReceiver_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ConnectionId"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Region"] = gen.PtrOf(gen.AlphaString())
	gens["TicketConfiguration"] = gen.PtrOf(gen.AlphaString())
	gens["WorkspaceId"] = gen.PtrOf(gen.AlphaString())
}

func Test_LogicAppReceiver_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LogicAppReceiver_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLogicAppReceiver_STATUS_ARM, LogicAppReceiver_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLogicAppReceiver_STATUS_ARM runs a test to see if a specific instance of LogicAppReceiver_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLogicAppReceiver_STATUS_ARM(subject LogicAppReceiver_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LogicAppReceiver_STATUS_ARM
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

// Generator of LogicAppReceiver_STATUS_ARM instances for property testing - lazily instantiated by
// LogicAppReceiver_STATUS_ARMGenerator()
var logicAppReceiver_STATUS_ARMGenerator gopter.Gen

// LogicAppReceiver_STATUS_ARMGenerator returns a generator of LogicAppReceiver_STATUS_ARM instances for property testing.
func LogicAppReceiver_STATUS_ARMGenerator() gopter.Gen {
	if logicAppReceiver_STATUS_ARMGenerator != nil {
		return logicAppReceiver_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLogicAppReceiver_STATUS_ARM(generators)
	logicAppReceiver_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(LogicAppReceiver_STATUS_ARM{}), generators)

	return logicAppReceiver_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForLogicAppReceiver_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLogicAppReceiver_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CallbackUrl"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["UseCommonAlertSchema"] = gen.PtrOf(gen.Bool())
}

func Test_SmsReceiver_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SmsReceiver_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSmsReceiver_STATUS_ARM, SmsReceiver_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSmsReceiver_STATUS_ARM runs a test to see if a specific instance of SmsReceiver_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSmsReceiver_STATUS_ARM(subject SmsReceiver_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SmsReceiver_STATUS_ARM
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

// Generator of SmsReceiver_STATUS_ARM instances for property testing - lazily instantiated by
// SmsReceiver_STATUS_ARMGenerator()
var smsReceiver_STATUS_ARMGenerator gopter.Gen

// SmsReceiver_STATUS_ARMGenerator returns a generator of SmsReceiver_STATUS_ARM instances for property testing.
func SmsReceiver_STATUS_ARMGenerator() gopter.Gen {
	if smsReceiver_STATUS_ARMGenerator != nil {
		return smsReceiver_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSmsReceiver_STATUS_ARM(generators)
	smsReceiver_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SmsReceiver_STATUS_ARM{}), generators)

	return smsReceiver_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSmsReceiver_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSmsReceiver_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CountryCode"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PhoneNumber"] = gen.PtrOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(ReceiverStatus_STATUS_ARM_Disabled, ReceiverStatus_STATUS_ARM_Enabled, ReceiverStatus_STATUS_ARM_NotSpecified))
}

func Test_VoiceReceiver_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VoiceReceiver_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVoiceReceiver_STATUS_ARM, VoiceReceiver_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVoiceReceiver_STATUS_ARM runs a test to see if a specific instance of VoiceReceiver_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVoiceReceiver_STATUS_ARM(subject VoiceReceiver_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VoiceReceiver_STATUS_ARM
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

// Generator of VoiceReceiver_STATUS_ARM instances for property testing - lazily instantiated by
// VoiceReceiver_STATUS_ARMGenerator()
var voiceReceiver_STATUS_ARMGenerator gopter.Gen

// VoiceReceiver_STATUS_ARMGenerator returns a generator of VoiceReceiver_STATUS_ARM instances for property testing.
func VoiceReceiver_STATUS_ARMGenerator() gopter.Gen {
	if voiceReceiver_STATUS_ARMGenerator != nil {
		return voiceReceiver_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVoiceReceiver_STATUS_ARM(generators)
	voiceReceiver_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VoiceReceiver_STATUS_ARM{}), generators)

	return voiceReceiver_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVoiceReceiver_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVoiceReceiver_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CountryCode"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PhoneNumber"] = gen.PtrOf(gen.AlphaString())
}

func Test_WebhookReceiver_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WebhookReceiver_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWebhookReceiver_STATUS_ARM, WebhookReceiver_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWebhookReceiver_STATUS_ARM runs a test to see if a specific instance of WebhookReceiver_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWebhookReceiver_STATUS_ARM(subject WebhookReceiver_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WebhookReceiver_STATUS_ARM
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

// Generator of WebhookReceiver_STATUS_ARM instances for property testing - lazily instantiated by
// WebhookReceiver_STATUS_ARMGenerator()
var webhookReceiver_STATUS_ARMGenerator gopter.Gen

// WebhookReceiver_STATUS_ARMGenerator returns a generator of WebhookReceiver_STATUS_ARM instances for property testing.
func WebhookReceiver_STATUS_ARMGenerator() gopter.Gen {
	if webhookReceiver_STATUS_ARMGenerator != nil {
		return webhookReceiver_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWebhookReceiver_STATUS_ARM(generators)
	webhookReceiver_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(WebhookReceiver_STATUS_ARM{}), generators)

	return webhookReceiver_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForWebhookReceiver_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWebhookReceiver_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["IdentifierUri"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ObjectId"] = gen.PtrOf(gen.AlphaString())
	gens["ServiceUri"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["UseAadAuth"] = gen.PtrOf(gen.Bool())
	gens["UseCommonAlertSchema"] = gen.PtrOf(gen.Bool())
}

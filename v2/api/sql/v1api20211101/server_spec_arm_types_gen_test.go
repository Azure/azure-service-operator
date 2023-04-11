// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

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

func Test_Server_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Server_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServer_Spec_ARM, Server_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServer_Spec_ARM runs a test to see if a specific instance of Server_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServer_Spec_ARM(subject Server_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Server_Spec_ARM
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

// Generator of Server_Spec_ARM instances for property testing - lazily instantiated by Server_Spec_ARMGenerator()
var server_Spec_ARMGenerator gopter.Gen

// Server_Spec_ARMGenerator returns a generator of Server_Spec_ARM instances for property testing.
// We first initialize server_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Server_Spec_ARMGenerator() gopter.Gen {
	if server_Spec_ARMGenerator != nil {
		return server_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServer_Spec_ARM(generators)
	server_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Server_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServer_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForServer_Spec_ARM(generators)
	server_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Server_Spec_ARM{}), generators)

	return server_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServer_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServer_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServer_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServer_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(ResourceIdentity_ARMGenerator())
	gens["Properties"] = gen.PtrOf(ServerProperties_ARMGenerator())
}

func Test_ResourceIdentity_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ResourceIdentity_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForResourceIdentity_ARM, ResourceIdentity_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForResourceIdentity_ARM runs a test to see if a specific instance of ResourceIdentity_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForResourceIdentity_ARM(subject ResourceIdentity_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ResourceIdentity_ARM
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

// Generator of ResourceIdentity_ARM instances for property testing - lazily instantiated by
// ResourceIdentity_ARMGenerator()
var resourceIdentity_ARMGenerator gopter.Gen

// ResourceIdentity_ARMGenerator returns a generator of ResourceIdentity_ARM instances for property testing.
func ResourceIdentity_ARMGenerator() gopter.Gen {
	if resourceIdentity_ARMGenerator != nil {
		return resourceIdentity_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForResourceIdentity_ARM(generators)
	resourceIdentity_ARMGenerator = gen.Struct(reflect.TypeOf(ResourceIdentity_ARM{}), generators)

	return resourceIdentity_ARMGenerator
}

// AddIndependentPropertyGeneratorsForResourceIdentity_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForResourceIdentity_ARM(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		ResourceIdentity_Type_None,
		ResourceIdentity_Type_SystemAssigned,
		ResourceIdentity_Type_SystemAssignedUserAssigned,
		ResourceIdentity_Type_UserAssigned))
}

func Test_ServerProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerProperties_ARM, ServerProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerProperties_ARM runs a test to see if a specific instance of ServerProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerProperties_ARM(subject ServerProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerProperties_ARM
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

// Generator of ServerProperties_ARM instances for property testing - lazily instantiated by
// ServerProperties_ARMGenerator()
var serverProperties_ARMGenerator gopter.Gen

// ServerProperties_ARMGenerator returns a generator of ServerProperties_ARM instances for property testing.
// We first initialize serverProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerProperties_ARMGenerator() gopter.Gen {
	if serverProperties_ARMGenerator != nil {
		return serverProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerProperties_ARM(generators)
	serverProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ServerProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForServerProperties_ARM(generators)
	serverProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ServerProperties_ARM{}), generators)

	return serverProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServerProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerProperties_ARM(gens map[string]gopter.Gen) {
	gens["AdministratorLogin"] = gen.PtrOf(gen.AlphaString())
	gens["AdministratorLoginPassword"] = gen.PtrOf(gen.AlphaString())
	gens["FederatedClientId"] = gen.PtrOf(gen.AlphaString())
	gens["KeyId"] = gen.PtrOf(gen.AlphaString())
	gens["MinimalTlsVersion"] = gen.PtrOf(gen.AlphaString())
	gens["PrimaryUserAssignedIdentityId"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(ServerProperties_PublicNetworkAccess_Disabled, ServerProperties_PublicNetworkAccess_Enabled))
	gens["RestrictOutboundNetworkAccess"] = gen.PtrOf(gen.OneConstOf(ServerProperties_RestrictOutboundNetworkAccess_Disabled, ServerProperties_RestrictOutboundNetworkAccess_Enabled))
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServerProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerProperties_ARM(gens map[string]gopter.Gen) {
	gens["Administrators"] = gen.PtrOf(ServerExternalAdministrator_ARMGenerator())
}

func Test_ServerExternalAdministrator_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerExternalAdministrator_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerExternalAdministrator_ARM, ServerExternalAdministrator_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerExternalAdministrator_ARM runs a test to see if a specific instance of ServerExternalAdministrator_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerExternalAdministrator_ARM(subject ServerExternalAdministrator_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerExternalAdministrator_ARM
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

// Generator of ServerExternalAdministrator_ARM instances for property testing - lazily instantiated by
// ServerExternalAdministrator_ARMGenerator()
var serverExternalAdministrator_ARMGenerator gopter.Gen

// ServerExternalAdministrator_ARMGenerator returns a generator of ServerExternalAdministrator_ARM instances for property testing.
func ServerExternalAdministrator_ARMGenerator() gopter.Gen {
	if serverExternalAdministrator_ARMGenerator != nil {
		return serverExternalAdministrator_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerExternalAdministrator_ARM(generators)
	serverExternalAdministrator_ARMGenerator = gen.Struct(reflect.TypeOf(ServerExternalAdministrator_ARM{}), generators)

	return serverExternalAdministrator_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServerExternalAdministrator_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerExternalAdministrator_ARM(gens map[string]gopter.Gen) {
	gens["AdministratorType"] = gen.PtrOf(gen.OneConstOf(ServerExternalAdministrator_AdministratorType_ActiveDirectory))
	gens["AzureADOnlyAuthentication"] = gen.PtrOf(gen.Bool())
	gens["Login"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalType"] = gen.PtrOf(gen.OneConstOf(ServerExternalAdministrator_PrincipalType_Application, ServerExternalAdministrator_PrincipalType_Group, ServerExternalAdministrator_PrincipalType_User))
	gens["Sid"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20181130

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

func Test_UserAssignedIdentities_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentities_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentitiesSpecARM, UserAssignedIdentitiesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentitiesSpecARM runs a test to see if a specific instance of UserAssignedIdentities_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentitiesSpecARM(subject UserAssignedIdentities_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentities_SpecARM
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

// Generator of UserAssignedIdentities_SpecARM instances for property testing - lazily instantiated by
// UserAssignedIdentitiesSpecARMGenerator()
var userAssignedIdentitiesSpecARMGenerator gopter.Gen

// UserAssignedIdentitiesSpecARMGenerator returns a generator of UserAssignedIdentities_SpecARM instances for property testing.
func UserAssignedIdentitiesSpecARMGenerator() gopter.Gen {
	if userAssignedIdentitiesSpecARMGenerator != nil {
		return userAssignedIdentitiesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentitiesSpecARM(generators)
	userAssignedIdentitiesSpecARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentities_SpecARM{}), generators)

	return userAssignedIdentitiesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentitiesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentitiesSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_FlexibleServersBackup_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersBackup_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersBackup_Spec, FlexibleServersBackup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersBackup_Spec runs a test to see if a specific instance of FlexibleServersBackup_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersBackup_Spec(subject FlexibleServersBackup_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersBackup_Spec
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

// Generator of FlexibleServersBackup_Spec instances for property testing - lazily instantiated by
// FlexibleServersBackup_SpecGenerator()
var flexibleServersBackup_SpecGenerator gopter.Gen

// FlexibleServersBackup_SpecGenerator returns a generator of FlexibleServersBackup_Spec instances for property testing.
func FlexibleServersBackup_SpecGenerator() gopter.Gen {
	if flexibleServersBackup_SpecGenerator != nil {
		return flexibleServersBackup_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersBackup_Spec(generators)
	flexibleServersBackup_SpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServersBackup_Spec{}), generators)

	return flexibleServersBackup_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersBackup_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersBackup_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

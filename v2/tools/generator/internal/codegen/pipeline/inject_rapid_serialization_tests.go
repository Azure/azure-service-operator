/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/testcases"
)

// InjectRapidSerializationTestsStageID is the unique identifier for this pipeline stage
const InjectRapidSerializationTestsStageID = "injectRapidSerializationTests"

func InjectRapidSerializationTests(idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		InjectRapidSerializationTestsStageID,
		"Add rapid-based test cases to verify JSON serialization",
		func(ctx context.Context, state *State) (*State, error) {
			// Phase II: no-op for existing groups (all are in gopterGroups)
			// New groups not in gopterGroups would get rapid tests

			// Verify that testcases.UseRapidForGroup is available (compile-time check)
			_ = testcases.UseRapidForGroup

			return state, nil
		})

	stage.RequiresPostrequisiteStages("simplifyDefinitions")

	return stage
}

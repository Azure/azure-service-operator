/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/rotisserie/eris"
)

// TransformCrossResourceReferencesToStringStageID is the unique identifier for this pipeline stage
const TransformCrossResourceReferencesToStringStageID = "transformCrossResourceReferencesToString"

// TransformCrossResourceReferencesToString replaces cross resource references with string.
func TransformCrossResourceReferencesToString() *Stage {
	return NewStage(
		TransformCrossResourceReferencesToStringStageID,
		"Replace cross-resource references with string",
		func(ctx context.Context, state *State) (*State, error) {
			updatedDefs, err := stripARMIDPrimitiveTypes(state.Definitions())
			if err != nil {
				return nil, eris.Wrap(err, "failed to strip ARM ID primitive types")
			}

			return state.WithDefinitions(updatedDefs), nil
		})
}

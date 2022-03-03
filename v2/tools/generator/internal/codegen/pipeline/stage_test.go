/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/pkg/errors"
)

// RunTestPipeline is used to run a sequence of stages as a part of a unit test
// Typically the earlier stages will be used to set up the required preconditions for the final stage under test
func RunTestPipeline(state *State, stages ...*Stage) (*State, error) {
	resultState := state
	for _, stage := range stages {
		s, err := stage.Run(context.TODO(), resultState)
		if err != nil {
			return nil, errors.Wrapf(err, "running stage %q", stage.id)
		}

		resultState = s
	}

	return resultState, nil
}

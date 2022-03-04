/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
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

/*
 * checkPreconditions Tests
 */

func TestStagePreconditions_GivenNoPrerequisites_ReturnsNoError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	testStage := NewFakeStage("test")
	state := NewState()
	g.Expect(testStage.checkPreconditions(state)).To(BeNil())
}

const (
	firstStageId = "firstStage"
	testStageId  = "testStage"
	lastStageId  = "lastStage"
)

func TestStagePreconditions_GivenSatisfiedPrerequisites_ReturnsNoError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	testStage := NewFakeStage(testStageId)
	testStage.RequiresPrerequisiteStages(firstStageId)

	state := NewState().WithSeenStage(firstStageId)
	g.Expect(testStage.checkPreconditions(state)).To(BeNil())
}

func TestStagePreconditions_GivenUnsatisfiedPrerequisites_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	testStage := NewFakeStage(testStageId)
	testStage.RequiresPrerequisiteStages(firstStageId)

	state := NewState()
	err := testStage.checkPreconditions(state)

	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(ContainSubstring(testStageId))
	g.Expect(err.Error()).To(ContainSubstring(firstStageId))
}

func TestStagePreconditions_WhenSatisfiedTooEarly_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	testStage := NewFakeStage(testStageId)
	testStage.RequiresPostrequisiteStages(lastStageId)

	state := NewState().WithSeenStage(lastStageId)
	err := testStage.checkPreconditions(state)

	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(ContainSubstring(testStageId))
	g.Expect(err.Error()).To(ContainSubstring(lastStageId))
}

func NewFakeStage(id string) *Stage {
	return NewStage(
		id,
		"Stage "+id,
		func(ctx context.Context, state *State) (*State, error) {
			return state, nil
		})
}

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"
)

/*
 * CheckFinalState tests
 */

func TestStateCheckFinalState_WhenNoExpectations_ReturnsNoError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	state := NewState()
	g.Expect(state.CheckFinalState()).To(BeNil())
}

func TestStateCheckFinalState_WhenExpectationSatisfied_ReturnsNoError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	state := NewState().
		WithExpectation(firstStageId, lastStageId).
		WithSeenStage(lastStageId)

	g.Expect(state.CheckFinalState()).To(BeNil())
}

func TestStateCheckFinalState_WhenExpectationNotSatisfied_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	state := NewState().
		WithExpectation(firstStageId, lastStageId)

	err := state.CheckFinalState()
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(ContainSubstring(firstStageId))
	g.Expect(err.Error()).To(ContainSubstring(lastStageId))
}

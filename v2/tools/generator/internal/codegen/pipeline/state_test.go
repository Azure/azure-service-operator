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

func TestStateInfo_WhenValueStored_CanBeRetrieved(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	state := NewState()
	state = StateWithInfo(state, ConversionGraphInfo, 42)

	value, ok := GetStateInfo[int](state, ConversionGraphInfo)
	g.Expect(ok).To(BeTrue())
	g.Expect(value).To(Equal(42))
}

func TestStateInfo_WhenValueStored_CannotBeRetrievedWithWrongType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	state := NewState()
	state = StateWithInfo(state, ConversionGraphInfo, 42)

	_, ok := GetStateInfo[string](state, ConversionGraphInfo)
	g.Expect(ok).To(BeFalse())
}

func TestStateInfo_CanStoreAndRecallMultipleValues(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	state := NewState()
	state = StateWithInfo(state, ConversionGraphInfo, 42)
	state = StateWithInfo(state, ExportedConfigMaps, "hello")

	graphInfo, ok := GetStateInfo[int](state, ConversionGraphInfo)
	g.Expect(ok).To(BeTrue())
	g.Expect(graphInfo).To(Equal(42))

	configMaps, ok := GetStateInfo[string](state, ExportedConfigMaps)
	g.Expect(ok).To(BeTrue())
	g.Expect(configMaps).To(Equal("hello"))
}

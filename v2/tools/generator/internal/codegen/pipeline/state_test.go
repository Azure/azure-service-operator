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
		WithExpectation(firstStageID, lastStageID).
		WithSeenStage(lastStageID)

	g.Expect(state.CheckFinalState()).To(BeNil())
}

func TestStateCheckFinalState_WhenExpectationNotSatisfied_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	state := NewState().
		WithExpectation(firstStageID, lastStageID)

	err := state.CheckFinalState()
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(ContainSubstring(firstStageID))
	g.Expect(err.Error()).To(ContainSubstring(lastStageID))
}

func TestStateInfo_WhenValueStored_CanBeRetrieved(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	state := NewState()
	state = StateWithData(state, ConversionGraphInfo, 42)

	value, err := GetStateData[int](state, ConversionGraphInfo)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(value).To(Equal(42))
}

func TestStateInfo_WhenValueStored_CannotBeRetrievedWithWrongType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	state := NewState()
	state = StateWithData(state, ConversionGraphInfo, 42)

	_, err := GetStateData[string](state, ConversionGraphInfo)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("expected string"))
}

func TestStateInfo_CanStoreAndRecallMultipleValues(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	state := NewState()
	state = StateWithData(state, ConversionGraphInfo, 42)
	state = StateWithData(state, ExportedConfigMaps, "hello")

	graphInfo, err := GetStateData[int](state, ConversionGraphInfo)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(graphInfo).To(Equal(42))

	configMaps, err := GetStateData[string](state, ExportedConfigMaps)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(configMaps).To(Equal("hello"))
}

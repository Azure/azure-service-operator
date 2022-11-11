/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package conditions_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

var _ conditions.Conditioner = &TestConditioner{}

type TestConditioner struct {
	Conditions []conditions.Condition
}

func (t *TestConditioner) GetConditions() conditions.Conditions {
	return t.Conditions
}

func (t *TestConditioner) SetConditions(conditions conditions.Conditions) {
	t.Conditions = conditions
}

func Test_SetCondition_AddsCondition(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	o := &TestConditioner{}
	clock := newMockClock()
	builder := conditions.NewPositiveConditionBuilder(clock)

	newCondition := builder.MakeTrueCondition(conditions.ConditionTypeReady, 0)
	conditions.SetCondition(o, newCondition)
	g.Expect(o.Conditions).To(HaveLen(1))
	g.Expect(o.Conditions[0]).To(Equal(newCondition))
}

func Test_SetCondition_ReadyTrueToReadyFalse_UpdatesConditionAndChangesTimestamp(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	o := &TestConditioner{}
	clock := newMockClock()
	builder := conditions.NewPositiveConditionBuilder(clock)

	initialCondition := builder.MakeTrueCondition(conditions.ConditionTypeReady, 0)
	conditions.SetCondition(o, initialCondition)

	clock.Add(1 * time.Second)

	// Now update
	updatedCondition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityError,
		0,
		"MyReason",
		"a message")
	conditions.SetCondition(o, updatedCondition)

	g.Expect(o.Conditions).To(HaveLen(1))
	g.Expect(o.Conditions[0]).To(Equal(updatedCondition))
	g.Expect(o.Conditions[0].LastTransitionTime).ToNot(Equal(initialCondition.LastTransitionTime))
}

func Test_SetCondition_ChangeReason_TimestampChanged(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	o := &TestConditioner{}
	clk := newMockClock()
	builder := conditions.NewPositiveConditionBuilder(clk)

	initialCondition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityError,
		0,
		"MyReason",
		"a message")
	conditions.SetCondition(o, initialCondition)

	clk.Add(1 * time.Second)

	// Now update
	updatedCondition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityError,
		0,
		"MyNewReason",
		"a message")
	conditions.SetCondition(o, updatedCondition)

	// Set the expected condition to the updated condition
	g.Expect(o.Conditions).To(HaveLen(1))
	g.Expect(o.Conditions[0]).To(Equal(updatedCondition))
}

func Test_SetCondition_SameConditionTimestampUnchanged(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	o := &TestConditioner{}
	clk := newMockClock()
	builder := conditions.NewPositiveConditionBuilder(clk)

	initialCondition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityError,
		0,
		"MyReason",
		"a message")
	conditions.SetCondition(o, initialCondition)

	clk.Add(1 * time.Second)

	// Now update
	updatedCondition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityError,
		0,
		"MyReason",
		"a message")
	conditions.SetCondition(o, updatedCondition)

	// Set the expected condition to the updated condition
	g.Expect(o.Conditions).To(HaveLen(1))
	g.Expect(o.Conditions[0]).To(Equal(initialCondition))
}

func Test_SetCondition_OverwritesAsExpected(t *testing.T) {
	t.Parallel()
	clk := newMockClock()
	builder := conditions.NewPositiveConditionBuilder(clk)

	infoGeneration1Condition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityInfo,
		1,
		"InfoReason",
		"a message")
	differentInfoGeneration1Condition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityInfo,
		1,
		"ADifferentInfoReason",
		"a message")
	infoGeneration2Condition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityInfo,
		2,
		"InfoOtherReason",
		"a message")
	warningGeneration1Condition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityWarning,
		1,
		"WarningReason",
		"a message")
	differentWarningGeneration1Condition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityWarning,
		1,
		"ADifferentWarningReason",
		"a message")
	warningGeneration2Condition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityWarning,
		2,
		"WarningOtherReason",
		"a message")
	errorGeneration1Condition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityError,
		1,
		"MyReason",
		"a message")
	differentErrorGeneration1Condition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityError,
		1,
		"ADifferentErrorReason",
		"a message")
	errorGeneration2Condition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityError,
		2,
		"MyOtherReason",
		"a message")
	trueGeneration1Condition := builder.MakeTrueCondition(conditions.ConditionTypeReady, 1)
	trueGeneration2Condition := builder.MakeTrueCondition(conditions.ConditionTypeReady, 2)

	unknownGeneration1Condition := builder.MakeUnknownCondition(
		conditions.ConditionTypeReady,
		1,
		"UnknownReason",
		"a message")
	unknownGeneration2Condition := builder.MakeUnknownCondition(
		conditions.ConditionTypeReady,
		2,
		"UnknownOtherReason",
		"a message")

	gen1List := []conditions.Condition{
		trueGeneration1Condition,
		infoGeneration1Condition,
		warningGeneration1Condition,
		errorGeneration1Condition,
		unknownGeneration1Condition,
	}

	gen2List := []conditions.Condition{
		trueGeneration2Condition,
		infoGeneration2Condition,
		warningGeneration2Condition,
		errorGeneration2Condition,
		unknownGeneration2Condition,
	}

	type testStruct struct {
		name              string
		initial           *conditions.Condition
		new               conditions.Condition
		expectedOverwrite bool
	}

	tests := []testStruct{
		// Something overwrites nothing
		{name: "True overwrites empty", initial: nil, new: trueGeneration1Condition, expectedOverwrite: true},
		{name: "Info overwrites empty", initial: nil, new: infoGeneration1Condition, expectedOverwrite: true},
		{name: "Warning overwrites empty", initial: nil, new: warningGeneration1Condition, expectedOverwrite: true},
		{name: "Error overwrites empty", initial: nil, new: errorGeneration1Condition, expectedOverwrite: true},

		// Test overwriting within the same generation (positive test cases)
		{name: "True overwrites same generation Info", initial: &infoGeneration1Condition, new: trueGeneration1Condition, expectedOverwrite: true},
		{name: "True overwrites same generation Warning", initial: &warningGeneration1Condition, new: trueGeneration1Condition, expectedOverwrite: true},
		{name: "True overwrites same generation Error", initial: &errorGeneration1Condition, new: trueGeneration1Condition, expectedOverwrite: true},
		{name: "Info overwrites same generation Info", initial: &infoGeneration1Condition, new: differentInfoGeneration1Condition, expectedOverwrite: true},
		{name: "Warning overwrites same generation Info", initial: &infoGeneration1Condition, new: warningGeneration1Condition, expectedOverwrite: true},
		{name: "Warning overwrites same generation Warning", initial: &warningGeneration1Condition, new: differentWarningGeneration1Condition, expectedOverwrite: true},
		{name: "Warning overwrites same generation Error", initial: &errorGeneration1Condition, new: warningGeneration1Condition, expectedOverwrite: true},
		{name: "Warning overwrites same generation True", initial: &trueGeneration1Condition, new: warningGeneration1Condition, expectedOverwrite: true},
		{name: "Error overwrites same generation Info", initial: &infoGeneration1Condition, new: errorGeneration1Condition, expectedOverwrite: true},
		{name: "Error overwrites same generation Warning", initial: &warningGeneration1Condition, new: errorGeneration1Condition, expectedOverwrite: true},
		{name: "Error overwrites same generation Error", initial: &errorGeneration1Condition, new: differentErrorGeneration1Condition, expectedOverwrite: true},
		{name: "Error overwrites same generation True", initial: &trueGeneration1Condition, new: errorGeneration1Condition, expectedOverwrite: true},

		// Test overwriting within the same generation (negative test cases)
		{name: "Info does NOT overwrite same generation Warning", initial: &warningGeneration1Condition, new: infoGeneration1Condition, expectedOverwrite: false},
		{name: "Info does NOT overwrite same generation Error", initial: &errorGeneration1Condition, new: infoGeneration1Condition, expectedOverwrite: false},
		{name: "Info does NOT overwrite same generation True", initial: &trueGeneration1Condition, new: infoGeneration1Condition, expectedOverwrite: false},
		{name: "Info does NOT overwrite same generation Warning", initial: &warningGeneration1Condition, new: unknownGeneration1Condition, expectedOverwrite: false},
		{name: "Info does NOT overwrite same generation Error", initial: &errorGeneration1Condition, new: unknownGeneration1Condition, expectedOverwrite: false},
		{name: "Info does NOT overwrite same generation True", initial: &trueGeneration1Condition, new: unknownGeneration1Condition, expectedOverwrite: false},
		{name: "Info does NOT overwrite same generation Info", initial: &infoGeneration1Condition, new: unknownGeneration1Condition, expectedOverwrite: false},
	}

	// Append all the combinations of gen2 overwrites gen1
	for _, gen1 := range gen1List {
		gen1 := gen1
		for _, gen2 := range gen2List {
			tests = append(
				tests,
				testStruct{
					name:              fmt.Sprintf("Newer %s overwrites older %s", makeFriendlyString(gen2), makeFriendlyString(gen1)),
					initial:           &gen1,
					new:               gen2,
					expectedOverwrite: true,
				})
		}
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			o := &TestConditioner{}
			if tt.initial != nil {
				conditions.SetCondition(o, *tt.initial)
			}

			conditions.SetCondition(o, tt.new)
			if tt.expectedOverwrite {
				g.Expect(o.Conditions[0]).To(Equal(tt.new))
			} else {
				g.Expect(o.Conditions[0]).To(Equal(*tt.initial))
			}
		})
	}
}

func Test_SetConditionReasonAware_OverwritesAsExpected(t *testing.T) {
	t.Parallel()
	clk := newMockClock()
	builder := conditions.NewPositiveConditionBuilder(clk)

	reconcilingCondition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityInfo,
		1,
		conditions.ReasonReconciling.Name,
		"a message")
	referenceNotFoundCondition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityWarning,
		1,
		conditions.ReasonReferenceNotFound.Name,
		"a message")
	secretNotFoundCondition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityWarning,
		1,
		conditions.ReasonSecretNotFound.Name,
		"a message")
	arbitraryInfoCondition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityInfo,
		1,
		"InfoReason",
		"a message")
	arbitraryWarningCondition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityWarning,
		1,
		"WarningReason",
		"a message")
	arbitraryErrorCondition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityError,
		1,
		"ErrorReason",
		"a message")
	waitingForOwnerWarningCondition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityWarning,
		1,
		conditions.ReasonWaitingForOwner.Name,
		"a message")
	successCondition := builder.MakeTrueCondition(conditions.ConditionTypeReady, 1)

	type testStruct struct {
		name              string
		initial           *conditions.Condition
		new               conditions.Condition
		expectedOverwrite bool
	}

	tests := []testStruct{
		// Test overwriting within the same generation (positive test cases)
		{name: "Reconciling overwrites same generation ReferenceNotFound", initial: &referenceNotFoundCondition, new: reconcilingCondition, expectedOverwrite: true},
		{name: "Reconciling overwrites same generation SecretNotFound", initial: &secretNotFoundCondition, new: reconcilingCondition, expectedOverwrite: true},
		{name: "Reconciling overwrites same generation ReferenceNotFound", initial: &referenceNotFoundCondition, new: reconcilingCondition, expectedOverwrite: true},
		{name: "Reconciling overwrites same generation Info", initial: &arbitraryInfoCondition, new: reconcilingCondition, expectedOverwrite: true},
		{name: "Reconciling overwrites same generation WaitingForOwner", initial: &waitingForOwnerWarningCondition, new: reconcilingCondition, expectedOverwrite: true},

		{name: "ReferenceNotFound overwrites same generation Error", initial: &arbitraryErrorCondition, new: referenceNotFoundCondition, expectedOverwrite: true},
		{name: "ReferenceNotFound overwrites same generation Warning", initial: &arbitraryWarningCondition, new: referenceNotFoundCondition, expectedOverwrite: true},
		{name: "ReferenceNotFound overwrites same generation Info", initial: &arbitraryInfoCondition, new: referenceNotFoundCondition, expectedOverwrite: true},
		{name: "ReferenceNotFound overwrites same generation Reconciling", initial: &reconcilingCondition, new: referenceNotFoundCondition, expectedOverwrite: true},

		{name: "SecretNotFound overwrites same generation Error", initial: &arbitraryErrorCondition, new: secretNotFoundCondition, expectedOverwrite: true},
		{name: "SecretNotFound overwrites same generation Warning", initial: &arbitraryWarningCondition, new: secretNotFoundCondition, expectedOverwrite: true},
		{name: "SecretNotFound overwrites same generation Info", initial: &arbitraryInfoCondition, new: secretNotFoundCondition, expectedOverwrite: true},
		{name: "SecretNotFound overwrites same generation Reconciling", initial: &reconcilingCondition, new: secretNotFoundCondition, expectedOverwrite: true},

		// Test overwriting within the same generation (negative test cases)
		{name: "Reconciling does NOT overwrite same generation Reconciling", initial: &reconcilingCondition, new: reconcilingCondition, expectedOverwrite: false},
		{name: "Reconciling does NOT overwrite same generation Warning", initial: &arbitraryWarningCondition, new: reconcilingCondition, expectedOverwrite: false},
		{name: "Reconciling does NOT overwrite same generation Error", initial: &arbitraryErrorCondition, new: reconcilingCondition, expectedOverwrite: false},
		{name: "Reconciling does NOT overwrite same generation Success", initial: &successCondition, new: reconcilingCondition, expectedOverwrite: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			o := &TestConditioner{}
			if tt.initial != nil {
				conditions.SetCondition(o, *tt.initial)
			}

			conditions.SetConditionReasonAware(o, tt.new)
			if tt.expectedOverwrite {
				g.Expect(o.Conditions[0]).To(Equal(tt.new))
			} else {
				g.Expect(o.Conditions[0]).To(Equal(*tt.initial))
			}
		})
	}
}

func makeFriendlyString(condition conditions.Condition) string {
	result := string(condition.Severity)
	if condition.Status == metav1.ConditionTrue {
		result = "True"
	} else if condition.Status == metav1.ConditionUnknown {
		result = "Unknown"
	}

	return result
}

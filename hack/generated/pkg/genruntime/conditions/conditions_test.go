/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package conditions_test

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/conditions"
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
	g := NewWithT(t)
	o := &TestConditioner{}
	clock := newMockClock()
	builder := conditions.NewPositiveConditionBuilder(clock)

	newCondition := builder.MakeTrueCondition(conditions.ConditionTypeReady)
	conditions.SetCondition(o, newCondition)
	g.Expect(o.Conditions).To(HaveLen(1))
	g.Expect(o.Conditions[0]).To(Equal(newCondition))
}

func Test_SetCondition_ReadyTrueToReadyFalse_UpdatesConditionAndChangesTimestamp(t *testing.T) {
	g := NewWithT(t)
	o := &TestConditioner{}
	clock := newMockClock()
	builder := conditions.NewPositiveConditionBuilder(clock)

	initialCondition := builder.MakeTrueCondition(conditions.ConditionTypeReady)
	conditions.SetCondition(o, initialCondition)

	clock.Add(1 * time.Second)

	// Now update
	updatedCondition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityError,
		"MyReason",
		"a message")
	conditions.SetCondition(o, updatedCondition)

	g.Expect(o.Conditions).To(HaveLen(1))
	g.Expect(o.Conditions[0]).To(Equal(updatedCondition))
	g.Expect(o.Conditions[0].LastTransitionTime).ToNot(Equal(initialCondition.LastTransitionTime))
}

func Test_SetCondition_ChangeReason_TimestampChanged(t *testing.T) {
	g := NewWithT(t)
	o := &TestConditioner{}
	clk := newMockClock()
	builder := conditions.NewPositiveConditionBuilder(clk)

	initialCondition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityError,
		"MyReason",
		"a message")
	conditions.SetCondition(o, initialCondition)

	clk.Add(1 * time.Second)

	// Now update
	updatedCondition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityError,
		"MyNewReason",
		"a message")
	conditions.SetCondition(o, updatedCondition)

	// Set the expected condition to the updated condition
	g.Expect(o.Conditions).To(HaveLen(1))
	g.Expect(o.Conditions[0]).To(Equal(updatedCondition))
}

func Test_SetCondition_SameConditionTimestampUnchanged(t *testing.T) {
	g := NewWithT(t)
	o := &TestConditioner{}
	clk := newMockClock()
	builder := conditions.NewPositiveConditionBuilder(clk)

	initialCondition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityError,
		"MyReason",
		"a message")
	conditions.SetCondition(o, initialCondition)

	clk.Add(1 * time.Second)

	// Now update
	updatedCondition := builder.MakeFalseCondition(
		conditions.ConditionTypeReady,
		conditions.ConditionSeverityError,
		"MyReason",
		"a message")
	conditions.SetCondition(o, updatedCondition)

	// Set the expected condition to the updated condition
	g.Expect(o.Conditions).To(HaveLen(1))
	g.Expect(o.Conditions[0]).To(Equal(initialCondition))
}

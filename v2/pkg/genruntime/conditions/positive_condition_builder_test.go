/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package conditions_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

func newMockClock() *clock.Mock {
	result := clock.NewMock()
	result.Set(time.Date(2021, 7, 7, 0, 0, 0, 0, time.UTC))

	return result
}

func Test_PositiveConditionBuilder_MakeTrueCondition(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	clk := newMockClock()

	var observedGeneration int64 = 0
	builder := conditions.NewPositiveConditionBuilder(clk)
	condition := builder.MakeTrueCondition(conditions.ConditionTypeReady, observedGeneration)

	g.Expect(condition.Type).To(Equal(conditions.ConditionType("Ready")))
	g.Expect(condition.Severity).To(Equal(conditions.ConditionSeverity("")))
	g.Expect(condition.Status).To(Equal(metav1.ConditionTrue))
	g.Expect(condition.ObservedGeneration).To(Equal(observedGeneration))
	g.Expect(condition.Reason).To(Equal("Succeeded"))
	g.Expect(condition.Message).To(Equal(""))
	g.Expect(condition.LastTransitionTime).To(Equal(metav1.NewTime(clk.Now())))
}

func Test_PositiveConditionBuilder_MakeFalseCondition(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	clk := newMockClock()

	builder := conditions.NewPositiveConditionBuilder(clk)
	reason := "MyReason"
	var observedGeneration int64 = 0
	message := fmt.Sprintf("This is a message %s", "yay")
	severity := conditions.ConditionSeverityError
	condition := builder.MakeFalseCondition(conditions.ConditionTypeReady, severity, observedGeneration, reason, message)

	g.Expect(condition.Type).To(Equal(conditions.ConditionType("Ready")))
	g.Expect(condition.Severity).To(Equal(conditions.ConditionSeverity("Error")))
	g.Expect(condition.Status).To(Equal(metav1.ConditionFalse))
	g.Expect(condition.ObservedGeneration).To(Equal(observedGeneration))
	g.Expect(condition.Reason).To(Equal(reason))
	g.Expect(condition.Message).To(Equal(message))
	g.Expect(condition.LastTransitionTime).To(Equal(metav1.NewTime(clk.Now())))
}

func Test_PositiveConditionBuilder_MakeUnknownCondition(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	clk := newMockClock()

	builder := conditions.NewPositiveConditionBuilder(clk)
	reason := "MyReason"
	var observedGeneration int64 = 0
	message := fmt.Sprintf("This is a message %s", "yay")
	condition := builder.MakeUnknownCondition(conditions.ConditionTypeReady, observedGeneration, reason, message)

	g.Expect(condition.Type).To(Equal(conditions.ConditionType("Ready")))
	g.Expect(condition.Severity).To(Equal(conditions.ConditionSeverity("")))
	g.Expect(condition.Status).To(Equal(metav1.ConditionUnknown))
	g.Expect(condition.ObservedGeneration).To(Equal(observedGeneration))
	g.Expect(condition.Reason).To(Equal(reason))
	g.Expect(condition.Message).To(Equal(message))
	g.Expect(condition.LastTransitionTime).To(Equal(metav1.NewTime(clk.Now())))
}

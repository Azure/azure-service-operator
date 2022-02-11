/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"
	"fmt"

	gomegaformat "github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

func actualAsConditioner(actual interface{}) (conditions.Conditioner, error) {
	c, ok := actual.(conditions.Conditioner)
	if !ok {
		return nil, errors.Errorf("expected conditions.Conditioner, was: %T", actual)
	}

	return c, nil
}

type DesiredStateMatcher struct {
	verify *Verify
	ctx    context.Context

	readyGoalStatus    metav1.ConditionStatus
	readyGoalSeverity  conditions.ConditionSeverity
	originalGeneration int64
}

var _ types.GomegaMatcher = &DesiredStateMatcher{}

func (m *DesiredStateMatcher) Match(actual interface{}) (bool, error) {
	if actual == nil {
		return false, nil
	}

	obj, err := actualAsObj(actual)
	if err != nil {
		return false, err
	}

	return m.verify.HasState(m.ctx, obj, m.readyGoalStatus, m.readyGoalSeverity, m.originalGeneration)
}

func (m *DesiredStateMatcher) FailureMessage(actual interface{}) string {
	conditioner, err := actualAsConditioner(actual)
	if err != nil {
		// Gomegas contract is that it won't call one of the message functions
		// if Match returned an error. If we make it here somehow that contract
		// has been violated
		panic(err)
	}

	ready, ok := conditions.GetCondition(conditioner, conditions.ConditionTypeReady)

	if !ok {
		return gomegaformat.Message(
			ready,
			fmt.Sprintf("%q condition to exist", conditions.ConditionTypeReady))
	}

	return gomegaformat.Message(
		ready,
		fmt.Sprintf("status to be %q, severity to be %q.", string(m.readyGoalStatus), string(m.readyGoalSeverity)))
}

func (m *DesiredStateMatcher) NegatedFailureMessage(actual interface{}) string {
	conditioner, err := actualAsConditioner(actual)
	if err != nil {
		// Gomegas contract is that it won't call one of the message functions
		// if Match returned an error. If we make it here somehow that contract
		// has been violated
		panic(err)
	}

	ready, ok := conditions.GetCondition(conditioner, conditions.ConditionTypeReady)

	if !ok {
		return gomegaformat.Message(
			ready,
			fmt.Sprintf("%q condition to exist", conditions.ConditionTypeReady))
	}

	return gomegaformat.Message(
		ready,
		fmt.Sprintf("status not to be %q, severity not to be %q.", string(m.readyGoalStatus), string(m.readyGoalSeverity)))
}

// MatchMayChangeInTheFuture implements OracleMatcher which of course isn't exported so we can't type-assert we implement it
func (m *DesiredStateMatcher) MatchMayChangeInTheFuture(actual interface{}) bool {
	if actual == nil {
		return false
	}

	conditioner, err := actualAsConditioner(actual)
	if err != nil {
		panic(err)
	}

	ready, ok := conditions.GetCondition(conditioner, conditions.ConditionTypeReady)
	if !ok {
		// If the condition doesn't exist yet then future change is possible
		return true
	}
	isTerminalState := ready.Status == metav1.ConditionTrue || ready.Severity == conditions.ConditionSeverityError
	hasChanged := m.originalGeneration != ready.ObservedGeneration
	return !isTerminalState || !hasChanged
}

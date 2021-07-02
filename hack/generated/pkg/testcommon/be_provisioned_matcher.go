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

	"github.com/Azure/azure-service-operator/hack/generated/pkg/armclient"
)

type DesiredStateMatcher struct {
	ensure *Ensure
	ctx    context.Context

	goalState     armclient.ProvisioningState
	previousState *armclient.ProvisioningState
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

	return m.ensure.HasState(m.ctx, obj, m.goalState)
}

func (m *DesiredStateMatcher) FailureMessage(actual interface{}) string {
	obj, err := actualAsObj(actual)
	if err != nil {
		// Gomegas contract is that it won't call one of the message functions
		// if Match returned an error. If we make it here somehow that contract
		// has been violated
		panic(err)
	}

	state := obj.GetAnnotations()[m.ensure.stateAnnotation]
	errorString := obj.GetAnnotations()[m.ensure.errorAnnotation]

	return gomegaformat.Message(
		state,
		fmt.Sprintf("state to be %s. Error is: %s", string(m.goalState), errorString))
}

func (m *DesiredStateMatcher) NegatedFailureMessage(actual interface{}) string {
	obj, err := actualAsObj(actual)
	if err != nil {
		// Gomegas contract is that it won't call one of the message functions
		// if Match returned an error. If we make it here somehow that contract
		// has been violated
		panic(err)
	}

	state := obj.GetAnnotations()[m.ensure.stateAnnotation]

	return gomegaformat.Message(state, fmt.Sprintf("state not to be %s", string(m.goalState)))
}

// MatchMayChangeInTheFuture implements OracleMatcher which of course isn't exported so we can't type-assert we implement it
func (m *DesiredStateMatcher) MatchMayChangeInTheFuture(actual interface{}) bool {
	if actual == nil {
		return false
	}

	obj, err := actualAsObj(actual)
	if err != nil {
		panic(err)
	}
	state := obj.GetAnnotations()[m.ensure.stateAnnotation]
	provisioningState := armclient.ProvisioningState(state)

	return !armclient.IsTerminalProvisioningState(provisioningState) || m.previousState != nil && *m.previousState == provisioningState
}

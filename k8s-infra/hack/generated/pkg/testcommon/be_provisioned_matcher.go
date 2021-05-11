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

	"github.com/Azure/k8s-infra/hack/generated/pkg/armclient"
)

type BeProvisionedMatcher struct {
	ensure *Ensure
	ctx    context.Context
}

var _ types.GomegaMatcher = &BeProvisionedMatcher{}

func (m *BeProvisionedMatcher) Match(actual interface{}) (bool, error) {

	if actual == nil {
		return false, nil
	}

	obj, err := actualAsObj(actual)
	if err != nil {
		return false, err
	}

	return m.ensure.Provisioned(m.ctx, obj)
}

func (m *BeProvisionedMatcher) FailureMessage(actual interface{}) string {
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
		fmt.Sprintf("state to be %s. Error is: %s", string(armclient.SucceededProvisioningState), errorString))
}

func (m *BeProvisionedMatcher) NegatedFailureMessage(actual interface{}) string {
	obj, err := actualAsObj(actual)
	if err != nil {
		// Gomegas contract is that it won't call one of the message functions
		// if Match returned an error. If we make it here somehow that contract
		// has been violated
		panic(err)
	}

	state := obj.GetAnnotations()[m.ensure.stateAnnotation]

	return gomegaformat.Message(state, fmt.Sprintf("state not to be %s", string(armclient.SucceededProvisioningState)))
}

// MatchMayChangeInTheFuture implements OracleMatcher which of course isn't exported so we can't type-assert we implement it
func (m *BeProvisionedMatcher) MatchMayChangeInTheFuture(actual interface{}) bool {
	if actual == nil {
		return false
	}

	obj, err := actualAsObj(actual)
	if err != nil {
		panic(err)
	}
	state := obj.GetAnnotations()[m.ensure.stateAnnotation]

	return !armclient.IsTerminalProvisioningState(armclient.ProvisioningState(state))
}

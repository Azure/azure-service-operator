/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"
	"fmt"
	"log"

	gomegaformat "github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	"github.com/pkg/errors"

	"github.com/Azure/k8s-infra/hack/generated/pkg/armclient"
)

type AzureBeProvisionedMatcher struct {
	azureClient armclient.Applier
	ctx         context.Context
}

var _ types.GomegaMatcher = &AzureBeProvisionedMatcher{}

func (m *AzureBeProvisionedMatcher) typedActual(actual interface{}) (*armclient.Deployment, error) {
	deployment, ok := actual.(*armclient.Deployment)
	if !ok {
		return nil, errors.Errorf("Expected actual of type *armclient.Deployment, instead %T", actual)
	}

	return deployment, nil
}

func (m *AzureBeProvisionedMatcher) Match(actual interface{}) (bool, error) {

	if actual == nil {
		return false, nil
	}

	deployment, err := m.typedActual(actual)
	if err != nil {
		return false, err
	}

	updatedDeployment, _, err := m.azureClient.GetDeployment(m.ctx, deployment.ID)
	if err != nil {
		return false, err
	}

	*deployment = *updatedDeployment

	log.Printf(
		"Ongoing deployment %s is in state: %s\n",
		deployment.ID,
		deployment.Properties.ProvisioningState)

	return deployment.IsSuccessful(), nil
}

func (m *AzureBeProvisionedMatcher) message(actual interface{}, expectedMatch bool) string {
	deployment, err := m.typedActual(actual)
	if err != nil {
		// Gomegas contract is that it won't call one of the message functions
		// if Match returned an error. If we make it here somehow that contract
		// has been violated
		panic(err)
	}

	stateStr := "<nil>"
	if deployment.Properties != nil {
		stateStr = string(deployment.Properties.ProvisioningState)
	}

	notStr := ""
	if !expectedMatch {
		notStr = "not "
	}

	return gomegaformat.Message(
		stateStr,
		fmt.Sprintf("deployment %s state to %sbe %s.", deployment.ID, notStr, string(armclient.SucceededProvisioningState)))
}

func (m *AzureBeProvisionedMatcher) FailureMessage(actual interface{}) string {
	return m.message(actual, true)
}

func (m *AzureBeProvisionedMatcher) NegatedFailureMessage(actual interface{}) string {
	return m.message(actual, false)
}

// MatchMayChangeInTheFuture implements OracleMatcher which of course isn't exported so we can't type-assert we implement it
func (m *AzureBeProvisionedMatcher) MatchMayChangeInTheFuture(actual interface{}) bool {
	if actual == nil {
		return false
	}

	deployment, err := m.typedActual(actual)
	if err != nil {
		// Gomegas contract is that it won't call one of the message functions
		// if Match returned an error. If we make it here somehow that contract
		// has been violated
		panic(err)
	}

	return !deployment.IsTerminalProvisioningState()
}

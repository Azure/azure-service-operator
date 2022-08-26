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

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
)

type AzureBeProvisionedMatcher struct {
	azureClient *genericarmclient.GenericClient
	ctx         context.Context
	err         error
}

var _ types.GomegaMatcher = &AzureBeProvisionedMatcher{}

func (m *AzureBeProvisionedMatcher) typedActual(actual interface{}) (*genericarmclient.PollerResponse[genericarmclient.GenericResource], error) {
	poller, ok := actual.(*genericarmclient.PollerResponse[genericarmclient.GenericResource])
	if !ok {
		return nil, errors.Errorf("Expected actual of type *genericarmclient.PollerResponse, instead %T", actual)
	}

	return poller, nil
}

func (m *AzureBeProvisionedMatcher) Match(actual interface{}) (bool, error) {
	if actual == nil {
		return false, nil
	}

	poller, err := m.typedActual(actual)
	if err != nil {
		return false, err
	}

	// The linter doesn't realize that we don't need to close the resp body because it's already done by the poller.
	// Suppressing it as it is a false positive.
	// nolint:bodyclose
	_, err = poller.Poller.Poll(m.ctx)
	if err != nil {
		m.err = err
		return false, err
	}

	// Poller.Poll will return an error if the operation fails, so if we make it here we're successful
	return poller.Poller.Done(), nil
}

func (m *AzureBeProvisionedMatcher) message(actual interface{}, expectedMatch bool) string {
	_, err := m.typedActual(actual)
	if err != nil {
		// Gomegas contract is that it won't call one of the message functions
		// if Match returned an error. If we make it here somehow that contract
		// has been violated
		panic(err)
	}

	notStr := ""
	if !expectedMatch {
		notStr = "not "
	}

	return gomegaformat.Message(
		m.err,
		fmt.Sprintf("long running operation %sbe successful.", notStr))
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

	poller, err := m.typedActual(actual)
	if err != nil {
		// Gomegas contract is that it won't call one of the message functions
		// if Match returned an error. If we make it here somehow that contract
		// has been violated
		panic(err)
	}

	return !poller.Poller.Done()
}

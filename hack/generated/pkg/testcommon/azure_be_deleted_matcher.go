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

	"github.com/Azure/k8s-infra/hack/generated/pkg/armclient"
)

type AzureBeDeletedMatcher struct {
	azureClient armclient.Applier
	ctx         context.Context
}

var _ types.GomegaMatcher = &AzureBeDeletedMatcher{}

func (m *AzureBeDeletedMatcher) typedActual(actual interface{}) ([]string, error) {
	args, ok := actual.([]string)
	if !ok {
		return nil, errors.Errorf("Expected actual of type []string, instead %T", actual)
	}

	if len(args) != 2 {
		return nil, errors.Errorf("Expected args of length 2, actually length: %d", len(args))
	}

	return args, nil
}

func (m *AzureBeDeletedMatcher) Match(actual interface{}) (bool, error) {

	if actual == nil {
		return false, nil
	}

	args, err := m.typedActual(actual)
	if err != nil {
		return false, err
	}

	exists, _, err := m.azureClient.HeadResource(m.ctx, args[0], args[1])
	if err != nil {
		return false, err
	}

	return !exists, nil
}

func (m *AzureBeDeletedMatcher) message(actual interface{}, expectedMatch bool) string {
	args, err := m.typedActual(actual)
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

	return gomegaformat.Message(args[0], fmt.Sprintf("%sto be deleted", notStr))
}

func (m *AzureBeDeletedMatcher) FailureMessage(actual interface{}) string {
	return m.message(actual, true)
}

func (m *AzureBeDeletedMatcher) NegatedFailureMessage(actual interface{}) string {
	return m.message(actual, false)
}

// MatchMayChangeInTheFuture implements OracleMatcher which of course isn't exported so we can't type-assert we implement it
func (m *AzureBeDeletedMatcher) MatchMayChangeInTheFuture(actual interface{}) bool {
	// TODO: Good way to generically check state of resource, look for deleting or something
	// TODO: which is common across all Azure resources?
	// For now, we assume things can always change unless actual == nil
	return actual != nil
}

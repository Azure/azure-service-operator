/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package matchers

import (
	"context"
	"fmt"

	gomegaformat "github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
)

func actualAsStr(actual interface{}) (string, error) {
	str, ok := actual.(string)
	strPtr, okPtr := actual.(*string)

	if !ok && !okPtr {
		return "", eris.Errorf("expected string or *string, was: %T", actual)
	}

	if okPtr {
		return *strPtr, nil
	}

	return str, nil
}

type BeDeletedInAzureMatcher struct {
	ctx        context.Context
	apiVersion string
	client     *genericarmclient.GenericClient
}

var _ types.GomegaMatcher = &BeDeletedInAzureMatcher{}

func (m *BeDeletedInAzureMatcher) Match(actual interface{}) (bool, error) {
	if actual == nil {
		return false, nil
	}

	id, err := actualAsStr(actual)
	if err != nil {
		return false, err
	}

	exists, _, err := m.client.CheckExistenceWithGetByID(
		m.ctx,
		id,
		m.apiVersion)
	if err != nil {
		return false, err
	}

	return !exists, nil
}

func (m *BeDeletedInAzureMatcher) message(actual interface{}, negate bool) string {
	id, err := actualAsStr(actual)
	if err != nil {
		// Gomegas contract is that it won't call one of the message functions
		// if Match returned an error. If we make it here somehow that contract
		// has been violated
		panic(err)
	}

	notStr := ""
	if negate {
		notStr = "not "
	}

	return gomegaformat.Message(id, fmt.Sprintf("%sto be deleted", notStr))
}

func (m *BeDeletedInAzureMatcher) FailureMessage(actual interface{}) string {
	return m.message(actual, false)
}

func (m *BeDeletedInAzureMatcher) NegatedFailureMessage(actual interface{}) string {
	return m.message(actual, true)
}

func BeDeletedInAzure(ctx context.Context, client *genericarmclient.GenericClient, apiVersion string) *BeDeletedInAzureMatcher {
	return &BeDeletedInAzureMatcher{}
}

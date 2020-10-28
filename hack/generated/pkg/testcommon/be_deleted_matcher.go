/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"

	gomegaformat "github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func actualAsObj(actual interface{}) (controllerutil.Object, error) {
	obj, ok := actual.(controllerutil.Object)
	if !ok {
		return nil, errors.Errorf("expected controllerutil.Object, was: %T", actual)
	}

	return obj, nil
}

type BeDeletedMatcher struct {
	ensure *Ensure
	ctx    context.Context

	subsequentMissingDeleteTimestamps int
}

var _ types.GomegaMatcher = &BeDeletedMatcher{}

func (m *BeDeletedMatcher) Match(actual interface{}) (bool, error) {

	if actual == nil {
		return false, nil
	}

	obj, err := actualAsObj(actual)
	if err != nil {
		return false, err
	}

	return m.ensure.Deleted(m.ctx, obj)
}

func (m *BeDeletedMatcher) FailureMessage(actual interface{}) string {
	obj, err := actualAsObj(actual)
	if err != nil {
		// Gomegas contract is that it won't call one of the message functions
		// if Match returned an error. If we make it here somehow that contract
		// has been violated
		panic(err)
	}

	return gomegaformat.Message(obj, "to be deleted")
}

func (m *BeDeletedMatcher) NegatedFailureMessage(actual interface{}) string {
	obj, err := actualAsObj(actual)
	if err != nil {
		// Gomegas contract is that it won't call one of the message functions
		// if Match returned an error. If we make it here somehow that contract
		// has been violated
		panic(err)
	}

	return gomegaformat.Message(obj, "not to be deleted")
}

// MatchMayChangeInTheFuture implements OracleMatcher which of course isn't exported so we can't type-assert we implement it
func (m *BeDeletedMatcher) MatchMayChangeInTheFuture(actual interface{}) bool {
	if actual == nil {
		return false
	}

	obj, err := actualAsObj(actual)
	if err != nil {
		panic(err)
	}

	// Initial object may not have a deletion timestamp set yet so look instead
	// for subsequent calls that all don't have timestamp
	deletionTimestamp := obj.GetDeletionTimestamp()
	if deletionTimestamp == nil {
		m.subsequentMissingDeleteTimestamps++
	}
	return m.subsequentMissingDeleteTimestamps <= 1
}

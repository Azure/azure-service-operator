/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package secrets

import (
	"testing"

	. "github.com/onsi/gomega"

	asocel "github.com/Azure/azure-service-operator/v2/internal/util/cel"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

type Self struct{}

// Can't run in parallel because it's mocking a package scope variable
//
//nolint:paralleltest
func Test_ValidateConfigMapDestination_MissingKey(t *testing.T) {
	g := NewGomegaWithT(t)

	evaluator, err := asocel.NewExpressionEvaluator()
	g.Expect(err).ToNot(HaveOccurred())
	expressionEvaluator = func() asocel.ExpressionEvaluator {
		return evaluator
	}
	defer func() {
		expressionEvaluator = asocel.Evaluator
	}()

	warnings, err := ValidateDestinations(
		&Self{},
		nil,
		[]*core.DestinationExpression{
			{
				Name:  "my-configmap",
				Value: `"hello"`,
			},
		})
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(MatchError(ContainSubstring("CEL expression with output type string must specify destination 'key'")))
}

// Can't run in parallel because it's mocking a package scope variable
//
//nolint:paralleltest
func Test_ValidateConfigMapDestination_UnneededKey(t *testing.T) {
	g := NewGomegaWithT(t)

	evaluator, err := asocel.NewExpressionEvaluator()
	g.Expect(err).ToNot(HaveOccurred())
	expressionEvaluator = func() asocel.ExpressionEvaluator {
		return evaluator
	}
	defer func() {
		expressionEvaluator = asocel.Evaluator
	}()

	warnings, err := ValidateDestinations(
		&Self{},
		nil,
		[]*core.DestinationExpression{
			{
				Name:  "my-configmap",
				Key:   "my-key",
				Value: `{"test": "test"}`,
			},
		})
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(MatchError(ContainSubstring("CEL expression with output type map[string]string must not specify destination 'key'")))
}

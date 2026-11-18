/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cel

import (
	"context"
	"errors"
	"math"
	"reflect"
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	celgo "cel.dev/cel-go/cel"
	"cel.dev/cel-go/common/types"
	"cel.dev/cel-go/common/types/ref"
)

type ExecutionResource struct {
	Text  string   `json:"text,omitempty"`
	Items []string `json:"items,omitempty"`
}

func Test_ExpressionExecutionLimits(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name        string
		expression  string
		text        string
		budget      uint64
		errStr      string
		expectedStr string
		expectedMap map[string]string
	}{
		{
			name:        "ordinary v3 formatting succeeds",
			expression:  `"%s".format([["abc", 3.14]])`,
			budget:      aggregateCostBudget,
			expectedStr: `["abc", 3.140000]`,
		},
		{
			name:        "float at precision boundary succeeds",
			expression:  `"%.100f".format([1.0])`,
			budget:      aggregateCostBudget,
			expectedStr: "1." + strings.Repeat("0", 100),
		},
		{
			name:       "constant excessive precision fails",
			expression: `"%.101f".format([1.0])`,
			budget:     aggregateCostBudget,
			errStr:     "precision 101 exceeds maximum allowed precision 100",
		},
		{
			name:       "dynamic excessive precision fails",
			expression: `self.text.format([1.0])`,
			text:       "%.101f",
			budget:     aggregateCostBudget,
			errStr:     "precision 101 exceeds maximum allowed precision 100",
		},
		{
			name:       "per expression cost exceeds budget",
			expression: `self.text.replace("a", "aa")`,
			text:       strings.Repeat("a", 600000),
			budget:     aggregateCostBudget,
			errStr:     "actual cost limit exceeded",
		},
		{
			name: "large string output within cost budget succeeds",
			// self.text doesn't consume much CEL evaluation budget because self.text is a simple field access.
			// The fact the field being accessed is very large doesn't play into the cost calculation.
			// With that said, the practical upper bound on the size of self.text is 1MB and you can already make ASO read that into memory
			// without CEL.
			expression:  `self.text`,
			text:        strings.Repeat("x", 1<<20+1), // roughly 1mb of string
			budget:      aggregateCostBudget,
			expectedStr: strings.Repeat("x", 1<<20+1),
		},
		{
			name:       "large map output within cost budget succeeds",
			expression: `{"a": self.text, "b": self.text}`,
			text:       strings.Repeat("x", 1<<19+1),
			budget:     aggregateCostBudget,
			expectedMap: map[string]string{
				"a": strings.Repeat("x", 1<<19+1), // roughly 500kb of string
				"b": strings.Repeat("x", 1<<19+1), // roughly 500kb of string
			},
		},
		{
			name:       "aggregate budget exceeded before expression evaluation",
			expression: `"a" + "b"`,
			budget:     0,
			errStr:     "aggregate cost budget exhausted",
		},
		{
			name:       "aggregate budget exceeded after expression evaluation",
			expression: `self.text.replace("a", "aa")`,
			text:       "aaaa",
			budget:     5,
			errStr:     "aggregate cost budget exceeded",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			evaluator, err := NewExpressionEvaluator()
			g.Expect(err).ToNot(HaveOccurred())
			result, err := evaluator.CompileAndRun(
				context.Background(),
				&EvaluationBudget{remaining: c.budget},
				c.expression,
				&ExecutionResource{Text: c.text},
				nil,
			)
			if c.errStr != "" {
				g.Expect(err).To(MatchError(ContainSubstring(c.errStr)))
				g.Expect(result).To(BeNil())
			} else {
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(result).ToNot(BeNil())
				g.Expect(result.Value).To(Equal(c.expectedStr))
				g.Expect(result.Values).To(Equal(c.expectedMap))
			}
		})
	}
}

func Test_ExpressionExecution_ChargesBudget(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	evaluator, err := NewExpressionEvaluator()
	g.Expect(err).ToNot(HaveOccurred())
	self := &ExecutionResource{}
	budget := &EvaluationBudget{remaining: 10}
	result, err := evaluator.CompileAndRun(t.Context(), budget, `"a" + "b"`, self, nil)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.Value).To(Equal("ab"))
	g.Expect(budget.remaining).To(Equal(uint64(9)))
}

func Test_ExpressionExecution_CancelsComprehension(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	calls := 0

	// Cancel on the 100th predicate evaluation to verify that comprehension evaluation stops.
	newEnv := func(resource reflect.Type) (*celgo.Env, error) {
		env, err := NewEnv(resource)
		if err != nil {
			return nil, err
		}
		return env.Extend(celgo.Function("cancelEvaluation",
			celgo.Overload("cancel_evaluation", nil, celgo.BoolType,
				celgo.FunctionBinding(func(_ ...ref.Val) ref.Val {
					calls++
					cancel()
					return types.True
				}))))
	}
	evaluator, err := NewExpressionEvaluator(Cache(NewUnCache(newEnv, Compile)))
	g.Expect(err).ToNot(HaveOccurred())
	self := &ExecutionResource{Items: make([]string, 1_000)}
	result, err := evaluator.CompileAndRun(ctx, NewEvaluationBudget(),
		`self.items.all(x, cancelEvaluation()) ? "done" : "failed"`, self, nil)
	g.Expect(errors.Is(err, context.Canceled)).To(BeTrue())
	g.Expect(result).To(BeNil())
	g.Expect(calls).To(Equal(100))
}

func Test_ExpressionBudget_Overflow(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	budget := &EvaluationBudget{remaining: 10}
	g.Expect(budget.charge(math.MaxUint64)).To(HaveOccurred())
	g.Expect(budget.check()).To(HaveOccurred())
}

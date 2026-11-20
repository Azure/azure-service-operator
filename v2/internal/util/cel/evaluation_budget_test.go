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

func Test_ExpressionExecution_CPULimits(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name        string
		expression  string
		text        string
		cpuBudget   uint64
		errStr      string
		expectedStr string
		expectedMap map[string]string
	}{
		{
			name:        "ordinary v3 formatting succeeds",
			expression:  `"%s".format([["abc", 3.14]])`,
			cpuBudget:   aggregateCPUBudget,
			expectedStr: `["abc", 3.140000]`,
		},
		{
			name:        "float at precision boundary succeeds",
			expression:  `"%.100f".format([1.0])`,
			cpuBudget:   aggregateCPUBudget,
			expectedStr: "1." + strings.Repeat("0", 100),
		},
		{
			name:       "constant excessive precision fails",
			expression: `"%.101f".format([1.0])`,
			cpuBudget:  aggregateCPUBudget,
			errStr:     "precision 101 exceeds maximum allowed precision 100",
		},
		{
			name:       "dynamic excessive precision fails",
			expression: `self.text.format([1.0])`,
			text:       "%.101f",
			cpuBudget:  aggregateCPUBudget,
			errStr:     "precision 101 exceeds maximum allowed precision 100",
		},
		{
			name:       "per expression cost exceeds budget",
			expression: `string(self.text.indexOf(self.text))`, // Search costs 4000 * 4000 * 0.1 = 1.6M units, exceeding the 1M CPU limit.
			text:       strings.Repeat("a", 4000),
			cpuBudget:  aggregateCPUBudget,
			errStr:     "actual cost limit exceeded",
		},
		{
			name:       "one expensive map value within CPU budget succeeds",
			expression: `{"a": string(self.text.indexOf(self.text))}`, // Each search costs ~625k units; its result is just "0".
			text:       strings.Repeat("a", 2500),
			cpuBudget:  aggregateCPUBudget,
			expectedMap: map[string]string{
				"a": "0",
			},
		},
		{
			name:       "two expensive map values exceed CPU budget",
			expression: `{"a": string(self.text.indexOf(self.text)), "b": string(self.text.indexOf(self.text))}`, // Two searches cost ~1.25M units, exceeding the 1M limit.
			text:       strings.Repeat("a", 2500),
			cpuBudget:  aggregateCPUBudget,
			errStr:     "actual cost limit exceeded",
		},
		{
			name:       "aggregate budget exceeded before expression evaluation",
			expression: `"a" + "b"`,
			cpuBudget:  0,
			errStr:     "aggregate CPU budget exhausted",
		},
		{
			name:       "aggregate budget exceeded after expression evaluation",
			expression: `self.text.replace("a", "aa")`,
			text:       "aaaa",
			cpuBudget:  5,
			errStr:     "aggregate CPU budget exceeded",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			evaluator, err := NewExpressionEvaluator()
			g.Expect(err).ToNot(HaveOccurred())
			budget := NewEvaluationBudget()
			budget.remainingCPU = c.cpuBudget
			result, err := evaluator.CompileAndRun(
				t.Context(),
				budget,
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

func Test_ExpressionExecution_MemoryLimits(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name         string
		expression   string
		text         string
		memoryBudget uint32
		errStr       string
		expectedStr  string
		expectedMap  map[string]string
	}{
		{
			name:         "string output at memory limit succeeds",
			expression:   `self.text`,
			text:         strings.Repeat("x", 1_000_000),
			memoryBudget: aggregateMemoryBudget,
			expectedStr:  strings.Repeat("x", 1_000_000), // roughly 1mb of string
		},
		{
			name:         "map output within memory budget succeeds",
			expression:   `{"a": self.text, "b": self.text}`,
			text:         strings.Repeat("x", 400_000),
			memoryBudget: aggregateMemoryBudget,
			expectedMap: map[string]string{
				"a": strings.Repeat("x", 400_000), // roughly 400 kb of string
				"b": strings.Repeat("x", 400_000), // roughly 400 kb of string
			},
		},
		{
			name:         "per expression memory limit exceeded",
			expression:   `self.text`,
			text:         strings.Repeat("x", 1_000_001),
			memoryBudget: aggregateMemoryBudget,
			errStr:       "memory limit exceeded",
		},
		{
			name:         "oversized intermediate rejected despite small output",
			expression:   `"%s%s".format([self.text, self.text]).substring(0, 1)`,
			text:         strings.Repeat("x", 500_001),
			memoryBudget: aggregateMemoryBudget,
			errStr:       "memory limit exceeded",
		},
		{
			name: "repeated formatting exceeds memory limit through shared string references",
			// After two 10x expansions, s is 100,100 bytes. The third argument list
			// counts ten references as over 100,000 memory units.
			expression: `[self.text]` +
				`.map(s, "%s%s%s%s%s%s%s%s%s%s".format([s,s,s,s,s,s,s,s,s,s]))` +
				`.map(s, "%s%s%s%s%s%s%s%s%s%s".format([s,s,s,s,s,s,s,s,s,s]))` +
				`.map(s, "%s%s%s%s%s%s%s%s%s%s".format([s,s,s,s,s,s,s,s,s,s]))[0]`,
			text:         strings.Repeat("x", 1001),
			memoryBudget: aggregateMemoryBudget,
			errStr:       "memory limit exceeded",
		},
		{
			name:         "aggregate budget exceeded before expression evaluation",
			expression:   `"x"`,
			memoryBudget: 0,
			errStr:       "CEL aggregate memory budget exhausted",
		},
		{
			name:         "aggregate budget exceeded after expression evaluation",
			expression:   `self.text`,
			text:         strings.Repeat("x", 21),
			memoryBudget: 2,
			errStr:       "CEL aggregate memory budget exceeded",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			evaluator, err := NewExpressionEvaluator()
			g.Expect(err).ToNot(HaveOccurred())
			budget := NewEvaluationBudget()
			budget.remainingMemory = c.memoryBudget
			result, err := evaluator.CompileAndRun(
				t.Context(),
				budget,
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
	budget := NewEvaluationBudget()
	budget.remainingCPU = 10
	result, err := evaluator.CompileAndRun(t.Context(), budget, `"a" + "b"`, self, nil)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.Value).To(Equal("ab"))
	g.Expect(budget.remainingCPU).To(Equal(uint64(9)))
	g.Expect(budget.remainingMemory).To(Equal(aggregateMemoryBudget - 1))
}

func Test_ExpressionExecution_ChargesResultMemory(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		expression string
		text       string
		memory     uint32
		remaining  uint32
	}{
		{
			name:       "charges only result rather than intermediate peak",
			expression: `self.text.substring(0, 1)`,
			text:       strings.Repeat("x", 10_000),
			memory:     2,
			remaining:  1,
		},
		{
			name:       "counts UTF-8 bytes with default CEL scaling",
			expression: `self.text`,
			text:       strings.Repeat("é", 6), // 12 UTF-8 bytes cost ceil(12/10) = 2 memory units, leaving 1 of the initial 3.
			memory:     3,
			remaining:  1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			evaluator, err := NewExpressionEvaluator()
			g.Expect(err).ToNot(HaveOccurred())
			budget := NewEvaluationBudget()
			budget.remainingMemory = c.memory
			result, err := evaluator.CompileAndRun(t.Context(), budget, c.expression, &ExecutionResource{Text: c.text}, nil)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(result).ToNot(BeNil())
			g.Expect(budget.remainingMemory).To(Equal(c.remaining))
		})
	}
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
	budget := NewEvaluationBudget()
	g.Expect(budget.chargeCPU(math.MaxUint64)).To(HaveOccurred())
	g.Expect(budget.check()).To(HaveOccurred())
	budget = NewEvaluationBudget()
	g.Expect(budget.chargeMemory(math.MaxUint32)).To(HaveOccurred())
	g.Expect(budget.check()).To(MatchError("CEL aggregate memory budget exhausted"))
}

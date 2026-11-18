/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cel

import "github.com/rotisserie/eris"

const (
	singleExpressionCostLimit = uint64(1_000_000)
	aggregateCostBudget       = uint64(10_000_000)
	interruptFrequency        = uint(100)
)

// EvaluationBudget accounts for all expressions in one resource export attempt.
// It is local to that attempt and must not be shared between concurrent evaluations.
type EvaluationBudget struct {
	remaining uint64
}

// NewEvaluationBudget creates a fresh aggregate budget.
func NewEvaluationBudget() *EvaluationBudget {
	return &EvaluationBudget{
		remaining: aggregateCostBudget,
	}
}

// check reports whether the budget has any remaining cost allowance.
func (b *EvaluationBudget) check() error {
	if b.remaining == 0 {
		return eris.New("CEL aggregate cost budget exhausted")
	}
	return nil
}

// charge deducts an evaluation's cost, exhausting the budget if the cost exceeds it.
func (b *EvaluationBudget) charge(cost uint64) error {
	if cost > b.remaining {
		b.remaining = 0
		return eris.New("CEL aggregate cost budget exceeded")
	}
	b.remaining -= cost
	return nil
}

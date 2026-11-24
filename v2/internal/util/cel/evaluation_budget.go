/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cel

import (
	"cel.dev/cel-go/common/types"
	"github.com/rotisserie/eris"
)

const (
	singleExpressionCPUBudget    = uint64(1_000_000)
	aggregateCPUBudget           = uint64(10_000_000)
	singleExpressionMemoryBudget = uint32(100_000)
	// Default CEL sizing counts ten string/bytes bytes as one memory unit.
	aggregateMemoryBudget = uint32(100_000)
	interruptFrequency    = uint(100)
)

// newMemorySizeCalculator retains CEL's default units and traversal count, while
// allowing chains deeper than CEL's default depth of five.
// This is useful for things like ResourceID parent chains.
func newMemorySizeCalculator() *types.SizeCalculator {
	return types.NewSizeCalculator(types.SizeCalculatorMaxDepth(64))
}

// EvaluationBudget accounts for all expressions in one resource export attempt.
// It is local to that attempt and must not be shared between concurrent evaluations.
type EvaluationBudget struct {
	remainingCPU    uint64
	remainingMemory uint32
}

// NewEvaluationBudget creates a fresh aggregate budget.
func NewEvaluationBudget() *EvaluationBudget {
	return &EvaluationBudget{
		remainingCPU:    aggregateCPUBudget,
		remainingMemory: aggregateMemoryBudget,
	}
}

// check reports whether the budget has remaining CPU and result-memory allowance.
func (b *EvaluationBudget) check() error {
	if b.remainingCPU == 0 {
		return eris.New("CEL aggregate CPU budget exhausted")
	}
	if b.remainingMemory == 0 {
		return eris.New("CEL aggregate memory budget exhausted")
	}
	return nil
}

// chargeCPU deducts evaluation cost, exhausting the CPU budget if the cost exceeds it.
func (b *EvaluationBudget) chargeCPU(cost uint64) error {
	if cost > b.remainingCPU {
		b.remainingCPU = 0
		return eris.New("CEL aggregate CPU budget exceeded")
	}
	b.remainingCPU -= cost
	return nil
}

// chargeMemory deducts the completed result's estimated size in CEL memory units,
// exhausting the memory budget if the size exceeds the remaining allowance.
func (b *EvaluationBudget) chargeMemory(size uint32) error {
	if size > b.remainingMemory {
		b.remainingMemory = 0
		return eris.New("CEL aggregate memory budget exceeded")
	}
	b.remainingMemory -= size
	return nil
}

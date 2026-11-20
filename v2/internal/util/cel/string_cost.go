/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cel

import (
	celcost "cel.dev/cel-go/common/cost"
	"cel.dev/cel-go/common/types/ref"
)

// stringCostEstimator supplies v5 string costs without changing v3 formatting.
// Formatting retains CEL's core estimate; this is not an allocation budget.
type stringCostEstimator struct{}

var _ celcost.ActualCostEstimator = stringCostEstimator{}

func (stringCostEstimator) CallCost(_ string, overloadID string, args []ref.Val, result ref.Val) *uint64 {
	switch overloadID {
	case "string_char_at_int":
		return trackStringCharAtCost(args, result)
	case "string_lower_ascii", "string_upper_ascii", "string_substring_int",
		"string_substring_int_int", "string_trim", "string_reverse":
		return trackStringTransformCost(args, result)
	case "string_index_of_string", "string_index_of_string_int",
		"string_last_index_of_string", "string_last_index_of_string_int":
		return trackStringSearchCost(args, result)
	case "string_replace_string_string", "string_replace_string_string_int":
		return trackStringReplaceCost(args, result)
	case "string_split_string", "string_split_string_int":
		return trackStringSplitCost(args, result)
	case "list_join", "list_join_string":
		return trackStringJoinCost(args, result)
	default:
		return nil
	}
}

// trackStringCharAtCost tracks runtime cost for O(n) string operations.
// Adapted to exported cost helpers from https://github.com/cel-expr/cel-go/blob/v0.31.0/ext/strings.go#L1032-L1036
//
//nolint:unparam // Keep the upstream cost tracker signature.
func trackStringCharAtCost(args []ref.Val, result ref.Val) *uint64 {
	size := celcost.SafeMultiplyByFactor(celcost.ActualSize(args[0]), celcost.StringTraversalCostFactor)
	cost := celcost.SafeAdd(celcost.CallCost, size, 1)
	return &cost
}

// trackStringTransformCost tracks runtime cost for O(n) string operations.
// Adapted to exported cost helpers from https://github.com/cel-expr/cel-go/blob/v0.31.0/ext/strings.go#L1039-L1044
func trackStringTransformCost(args []ref.Val, result ref.Val) *uint64 {
	transformCost := celcost.SafeMultiplyByFactor(celcost.ActualSize(args[0]), celcost.StringTraversalCostFactor)
	resultSize := celcost.ActualSize(result)
	cost := celcost.SafeAdd(celcost.CallCost, transformCost, resultSize)
	return &cost
}

// trackStringSearchCost tracks runtime cost for O(n*m) string search operations.
// Adapted to exported cost helpers from https://github.com/cel-expr/cel-go/blob/v0.31.0/ext/strings.go#L1047-L1051
func trackStringSearchCost(args []ref.Val, _ ref.Val) *uint64 {
	searchSize := celcost.SafeMultiply(celcost.ActualSize(args[0]), celcost.ActualSize(args[1]))
	searchCost := celcost.SafeMultiplyByFactor(searchSize, celcost.StringTraversalCostFactor)
	cost := celcost.SafeAdd(searchCost, celcost.CallCost)
	return &cost
}

// trackStringReplaceCost tracks runtime cost for string replace operations,
// accounting for search cost and the size of the result.
// Adapted to exported cost helpers from https://github.com/cel-expr/cel-go/blob/v0.31.0/ext/strings.go#L1055-L1067
func trackStringReplaceCost(args []ref.Val, result ref.Val) *uint64 {
	targetSize := celcost.ActualSize(args[0])
	if targetSize == 0 {
		targetSize = 1
	}
	needleSize := celcost.ActualSize(args[1])
	if needleSize == 0 {
		needleSize = 1
	}
	searchSize := celcost.SafeMultiply(targetSize, needleSize)
	searchCost := celcost.SafeMultiplyByFactor(searchSize, celcost.StringTraversalCostFactor)
	cost := celcost.SafeAdd(celcost.CallCost, searchCost, celcost.ActualSize(result))
	return &cost
}

// trackStringSplitCost tracks runtime cost for string split operations,
// accounting for traversal and list allocation.
// Adapted to exported cost helpers from https://github.com/cel-expr/cel-go/blob/v0.31.0/ext/strings.go#L1071-L1076
func trackStringSplitCost(args []ref.Val, result ref.Val) *uint64 {
	traversalSize := celcost.SafeAdd(celcost.ActualSize(args[0]), 1)
	traversalCost := celcost.SafeMultiplyByFactor(traversalSize, celcost.StringTraversalCostFactor)
	resultSize := celcost.ActualSize(result)
	cost := celcost.SafeAdd(celcost.CallCost, traversalCost, resultSize, celcost.ListCreateBaseCost)
	return &cost
}

// trackStringJoinCost tracks runtime cost for string join operations,
// accounting for traversal and the size of the result.
// Adapted to exported cost helpers from https://github.com/cel-expr/cel-go/blob/v0.31.0/ext/strings.go#L1080-L1084
func trackStringJoinCost(args []ref.Val, result ref.Val) *uint64 {
	traversalSize := celcost.SafeAdd(celcost.ActualSize(args[0]), 1)
	traversalCost := celcost.SafeMultiplyByFactor(traversalSize, celcost.StringTraversalCostFactor)
	cost := celcost.SafeAdd(celcost.CallCost, traversalCost, celcost.ActualSize(result))
	return &cost
}

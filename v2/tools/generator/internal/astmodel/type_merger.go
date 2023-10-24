/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"reflect"

	"github.com/pkg/errors"
)

// TypeMerger is like a visitor for 2 types.
//
// Conceptually it takes (via Add) a list of functions of the form:
//
//	func ([ctx interface{},] left {some Type}, right {some Type}) (Type, error)
//
// where `left` and `right` can be concrete types that implement the `Type` interface.
//
// When `TypeMerger.Merge(Type, Type)` is invoked, it will iterate through the
// provided functions and invoke the first one that matches the concrete types
// passed. If none match then the fallback provided to `NewTypeMerger` will be invoked.
//
// The `ctx` argument can optionally be used to “smuggle” additional data down the call-chain.
type TypeMerger struct {
	mergers  []mergerRegistration
	fallback MergerFunc
}

type mergerRegistration struct {
	left  reflect.Type
	right reflect.Type
	merge MergerFunc
}

type MergerFunc func(ctx interface{}, left, right Type) (Type, error)

func NewTypeMerger(fallback MergerFunc) *TypeMerger {
	return &TypeMerger{fallback: fallback}
}

var (
	typeInterface  reflect.Type = reflect.TypeOf((*Type)(nil)).Elem() // yuck
	errorInterface reflect.Type = reflect.TypeOf((*error)(nil)).Elem()
	mergerFuncType reflect.Type = reflect.TypeOf((*MergerFunc)(nil)).Elem()
)

type validatedMerger struct {
	merger                    reflect.Value
	leftArgType, rightArgType reflect.Type
	needsCtx                  bool
}

func validateMerger(merger interface{}) validatedMerger {
	it := reflect.ValueOf(merger)
	if it.Kind() != reflect.Func {
		panic("merger must be a function")
	}

	mergerType := it.Type()

	badArgumentsMsg := "merger must take take arguments of type (left [some Type], right [some Type]) or (ctx X, left [some Type], right [some Type])"
	if mergerType.NumIn() < 2 || mergerType.NumIn() > 3 {
		panic(badArgumentsMsg)
	}

	argOffset := mergerType.NumIn() - 2

	needsCtx := argOffset != 0
	leftArg := mergerType.In(argOffset + 0)
	rightArg := mergerType.In(argOffset + 1)

	if !leftArg.AssignableTo(typeInterface) || !rightArg.AssignableTo(typeInterface) {
		panic(badArgumentsMsg)
	}

	if mergerType.NumOut() != 2 ||
		mergerType.Out(0) != typeInterface ||
		mergerType.Out(1) != errorInterface {
		panic("merger must return (Type, error)")
	}

	return validatedMerger{
		merger:       it,
		leftArgType:  leftArg,
		rightArgType: rightArg,
		needsCtx:     needsCtx,
	}
}

func buildMergerRegistration(v validatedMerger, flip bool) mergerRegistration {
	leftArgType := v.leftArgType
	rightArgType := v.rightArgType
	if flip {
		leftArgType, rightArgType = rightArgType, leftArgType
	}

	return mergerRegistration{
		left:  leftArgType,
		right: rightArgType,
		merge: reflect.MakeFunc(mergerFuncType, func(args []reflect.Value) []reflect.Value {
			// we dereference the Type here to the underlying value so that
			// the merger can take either Type or a specific type.
			// if it takes Type then the compiler/runtime will convert the value back to a Type
			ctxValue := args[0].Elem()
			leftValue := args[1].Elem()
			rightValue := args[2].Elem()
			if flip {
				leftValue, rightValue = rightValue, leftValue
			}

			if v.needsCtx {
				return v.merger.Call([]reflect.Value{ctxValue, leftValue, rightValue})
			} else {
				return v.merger.Call([]reflect.Value{leftValue, rightValue})
			}
		}).Interface().(MergerFunc),
	}
}

// Add adds a handler function to be invoked if applicable. See the docs on
// TypeMerger above for a full explanation.
func (m *TypeMerger) Add(mergeFunc interface{}) {
	v := validateMerger(mergeFunc)
	m.mergers = append(m.mergers, buildMergerRegistration(v, false))
}

// AddUnordered adds a handler function that doesn’t care what order
// the two type parameters are in. e.g. if it has type `(A, B) -> (Type, error)`,
// it can match either `(A, B)` or `(B, A)`. This is useful when the merger
// is symmetric and handles two different types.
func (m *TypeMerger) AddUnordered(mergeFunc interface{}) {
	v := validateMerger(mergeFunc)
	m.mergers = append(m.mergers, buildMergerRegistration(v, false), buildMergerRegistration(v, true))
}

// Merge merges the two types according to the provided mergers and fallback, with nil context
func (m *TypeMerger) Merge(left, right Type) (Type, error) {
	return m.MergeWithContext(nil, left, right)
}

// MergeWithContext merges the two types according to the provided mergers and fallback, with the provided context
func (m *TypeMerger) MergeWithContext(ctx interface{}, left, right Type) (Type, error) {
	if left == nil {
		return right, nil
	}

	if right == nil {
		return left, nil
	}

	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)

	for _, merger := range m.mergers {
		leftTypeMatches := merger.left == leftType ||
			(merger.left.Kind() == reflect.Interface && leftType.Implements(merger.left))
		rightTypeMatches := merger.right == rightType ||
			(merger.right.Kind() == reflect.Interface && rightType.Implements(merger.right))

		if leftTypeMatches && rightTypeMatches {
			result, err := merger.merge(ctx, left, right)
			if (result == nil && err == nil) || errors.Is(err, ContinueMerge) {
				// these conditions indicate that the merger was not actually applicable,
				// despite having a type that matched
				continue
			}

			return result, err
		}
	}

	return m.fallback(ctx, left, right)
}

var ContinueMerge error = errors.New("special error that indicates that the merger was not applicable")

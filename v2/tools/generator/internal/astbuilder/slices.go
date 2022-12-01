/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"go/token"

	"github.com/dave/dst"
)

// MakeSlice returns the call expression for making a slice
func MakeSlice(listType dst.Expr, len dst.Expr) *dst.CallExpr {
	/*
	 * Sample output:
	 *
	 * make([]<value>)
	 *
	 */
	return &dst.CallExpr{
		Fun: dst.NewIdent("make"),
		Args: []dst.Expr{
			listType,
			len,
		},
	}
}

// AppendItemToSlice returns a statement to append a single item to a slice
func AppendItemToSlice(lhs dst.Expr, rhs dst.Expr) dst.Stmt {
	/*
	 * Sample output:
	 *
	 * <lhs> = append(<lhs>, <rhs>)
	 *
	 */
	return SimpleAssignment(
		dst.Clone(lhs).(dst.Expr),
		CallFunc("append", dst.Clone(lhs).(dst.Expr), dst.Clone(rhs).(dst.Expr)))
}

// AppendItemsToSlice returns a statement to append many individual items to a slice
func AppendItemsToSlice(lhs dst.Expr, rhs ...dst.Expr) dst.Stmt {
	/*
	 * Sample output:
	 *
	 * <lhs> = append(<lhs>, <rhs>, <rhs>, <rhs>, ...)
	 *
	 */
	args := make([]dst.Expr, 0, len(rhs)+1)
	args = append(args, lhs)
	for _, arg := range rhs {
		args = append(args, dst.Clone(arg).(dst.Expr))
	}

	return SimpleAssignment(
		dst.Clone(lhs).(dst.Expr),
		CallFunc("append", args...))
}

// AppendSliceToSlice returns a statement to append a slice to another slice
func AppendSliceToSlice(lhs dst.Expr, rhs dst.Expr) dst.Stmt {
	/*
	 * Sample output:
	 *
	 * <lhs> = append(<lhs>, <rhs>...)
	 *
	 */
	f := CallFunc("append", dst.Clone(lhs).(dst.Expr), dst.Clone(rhs).(dst.Expr))
	f.Ellipsis = true
	return SimpleAssignment(
		dst.Clone(lhs).(dst.Expr),
		f)
}

// IterateOverSlice creates a statement to iterate over the content of a list using the specified
// identifier for each element in the list
func IterateOverSlice(item string, list dst.Expr, statements ...dst.Stmt) *dst.RangeStmt {
	/*
	 * Sample output:
	 *
	 * for _, <item> := range <list> {
	 *     <statements>
	 * }
	 *
	 */
	return &dst.RangeStmt{
		Key:   dst.NewIdent("_"),
		Value: dst.NewIdent(item),
		Tok:   token.DEFINE,
		X:     list,
		Body:  StatementBlock(statements...),
	}
}

// IterateOverSliceWithIndex creates a statement to iterate over the content of a list using the specified
// identifiers for each index and element in the list.
func IterateOverSliceWithIndex(index string, item string, list dst.Expr, statements ...dst.Stmt) *dst.RangeStmt {
	/*
	 * Sample output:
	 *
	 * for <index>, <item> := range <list> {
	 *     <statements>
	 * }
	 *
	 */
	return &dst.RangeStmt{
		Key:   dst.NewIdent(index),
		Value: dst.NewIdent(item),
		Tok:   token.DEFINE,
		X:     list,
		Body:  StatementBlock(statements...),
	}
}

// SliceLiteral creates a slice literal
func SliceLiteral(arrayType dst.Expr, items ...dst.Expr) dst.Expr {
	/*
	 * Sample output:
	 *
	 * []<arrayType>{<items...>}
	 *
	 */
	result := &dst.CompositeLit{
		Type: &dst.ArrayType{
			Elt: arrayType,
		},
		Elts: items,
	}
	return result
}

// SliceLiteralBuilder captures the information required to generate code for an inline slice initialization
type SliceLiteralBuilder struct {
	sliceType      dst.Expr
	elts           []dst.Expr
	linePerElement bool
}

// NewSliceLiteralBuilder creates a new instance for initialization of the specified slice
// The linePerElt parameter determines if the slice literal should have a single element per line
func NewSliceLiteralBuilder(sliceType dst.Expr, linePerElement bool) *SliceLiteralBuilder {
	return &SliceLiteralBuilder{
		sliceType:      sliceType,
		linePerElement: linePerElement,
	}
}

// AddElement adds an element to the slice literal
// Returns the receiver to allow method chaining when desired
func (b *SliceLiteralBuilder) AddElement(element dst.Expr) *SliceLiteralBuilder {
	expr := dst.Clone(element).(dst.Expr)

	if compositeLit, ok := expr.(*dst.CompositeLit); ok {
		compositeLit.Type = nil // Remove the Type, as we don't actually need it in an array context
	}

	if b.linePerElement {
		expr.Decorations().Before = dst.NewLine
		expr.Decorations().After = dst.NewLine
	}

	b.elts = append(b.elts, expr)

	return b
}

// Build constructs the actual dst.CompositeLit describing the slice literal
func (b *SliceLiteralBuilder) Build() dst.Expr {
	return SliceLiteral(b.sliceType, b.elts...)
}

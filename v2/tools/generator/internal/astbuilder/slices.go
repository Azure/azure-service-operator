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
//
// make([]<value>)
//
func MakeSlice(listType dst.Expr, len dst.Expr) *dst.CallExpr {
	return &dst.CallExpr{
		Fun: dst.NewIdent("make"),
		Args: []dst.Expr{
			listType,
			len,
		},
	}
}

// AppendSlice returns a statement for a slice append
//
// <lhs> = append(<lhs>, <rhs>)
//
func AppendSlice(lhs dst.Expr, rhs dst.Expr) dst.Stmt {
	return SimpleAssignment(
		dst.Clone(lhs).(dst.Expr),
		CallFunc("append", dst.Clone(lhs).(dst.Expr), dst.Clone(rhs).(dst.Expr)))
}

// IterateOverSlice creates a statement to iterate over the content of a list using the specified
// identifier for each element in the list
//
// for _, <item> := range <list> {
//     <statements>
// }
//
func IterateOverSlice(item string, list dst.Expr, statements ...dst.Stmt) *dst.RangeStmt {
	return &dst.RangeStmt{
		Key:   dst.NewIdent("_"),
		Value: dst.NewIdent(item),
		Tok:   token.DEFINE,
		X:     list,
		Body:  StatementBlock(statements...),
	}
}

// IterateOverSliceWithIndex creates a statement to iterate over the content of a list using the specified
// identifiers for each index and element in the list
//
// for <index>, <item> := range <list> {
//     <statements>
// }
//
func IterateOverSliceWithIndex(index string, item string, list dst.Expr, statements ...dst.Stmt) *dst.RangeStmt {
	return &dst.RangeStmt{
		Key:   dst.NewIdent(index),
		Value: dst.NewIdent(item),
		Tok:   token.DEFINE,
		X:     list,
		Body:  StatementBlock(statements...),
	}
}

// SliceLiteral creates a slice literal
//
// []<arrayType>{<items...>}
func SliceLiteral(arrayType dst.Expr, items ...dst.Expr) dst.Expr {
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

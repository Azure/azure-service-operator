/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"github.com/dave/dst"
)

// CompositeLiteralDetails captures the information required to generate code for an inline struct initialization
type CompositeLiteralDetails struct {
	structType dst.Expr
	elts       []dst.Expr
}

// NewCompositeLiteralDetails creates a new instance for initialization of the specified struct
// structType is an expression to handle both structs from the current package and imported ones requiring qualification
func NewCompositeLiteralDetails(structType dst.Expr) *CompositeLiteralDetails {
	return &CompositeLiteralDetails{
		structType: structType,
	}
}

// AddField adds initialization of another field
// Returns the receiver to allow method chaining when desired
func (details *CompositeLiteralDetails) AddField(name string, value dst.Expr) {
	expr := &dst.KeyValueExpr{
		Key:   dst.NewIdent(name),
		Value: dst.Clone(value).(dst.Expr),
	}

	expr.Decs.Before = dst.NewLine
	expr.Decs.After = dst.NewLine

	details.elts = append(details.elts, expr)
}

// Build constructs the actual dst.CompositeLit that's required
func (details CompositeLiteralDetails) Build() *dst.CompositeLit {
	return &dst.CompositeLit{
		Type: details.structType,
		Elts: details.elts,
	}
}

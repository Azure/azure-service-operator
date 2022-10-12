/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"github.com/dave/dst"
)

// CompositeLiteralBuilder captures the information required to generate code for an inline struct initialization
type CompositeLiteralBuilder struct {
	structType dst.Expr
	elts       []dst.Expr
	newLines   bool
}

// NewCompositeLiteralBuilder creates a new instance for initialization of the specified struct
// structType is an expression to handle both structs from the current package and imported ones requiring qualification
func NewCompositeLiteralBuilder(structType dst.Expr) *CompositeLiteralBuilder {
	return &CompositeLiteralBuilder{
		structType: structType,
		newLines:   true,
	}
}

// WithoutNewLines returns the CompositeLiteralBuilder without NewLines enabled
func (b *CompositeLiteralBuilder) WithoutNewLines() *CompositeLiteralBuilder {
	b.newLines = false
	return b
}

// AddField adds initialization of another field
// Returns the receiver to allow method chaining when desired
func (b *CompositeLiteralBuilder) AddField(name string, value dst.Expr) *CompositeLiteralBuilder {
	expr := &dst.KeyValueExpr{
		Key:   dst.NewIdent(name),
		Value: dst.Clone(value).(dst.Expr),
	}
	expr.Decs.Before = dst.NewLine
	expr.Decs.After = dst.NewLine

	b.elts = append(b.elts, expr)
	return b
}

// Build constructs the actual dst.CompositeLit that's required
func (b *CompositeLiteralBuilder) Build() *dst.CompositeLit {
	// If we only have a single element, remove the NewLine directives to force it onto a single line.
	// While there are places where this doesn't look amazing, on balance it is cleaner than having them
	// on multiple lines. This is especially true when the CompositeLit is used in a deeply nested context.
	if len(b.elts) == 1 {
		for _, elt := range b.elts {
			kvExpr := elt.(*dst.KeyValueExpr) // This is safe because it's the only type we put into elts
			kvExpr.Decs.Before = dst.None
			kvExpr.Decs.After = dst.None
		}
	}

	return &dst.CompositeLit{
		Type: b.structType,
		Elts: b.elts,
	}
}

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"go/token"
	"testing"

	"github.com/dave/dst"
	. "github.com/onsi/gomega"
)

func asplode(l, r dst.Expr) dst.Expr {
	panic("asplode was invoked")
}

func TestReduceZero(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect(func() { Reduce(asplode, dst.NewLine) }).To(PanicWith("must provide at least one expression to reduce"))
}

func TestReduceOne(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expr := &dst.BadExpr{}
	g.Expect(Reduce(asplode, dst.None, expr)).To(BeIdenticalTo(expr))
}

func TestJoinOr(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expr1 := &dst.CallExpr{}
	expr2 := &dst.BadExpr{}

	expected := &dst.BinaryExpr{X: expr1, Op: token.LOR, Y: expr2}

	g.Expect(JoinOr(expr1, expr2)).To(Equal(expected))
}

func TestJoinAnd(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expr1 := &dst.CallExpr{}
	expr2 := &dst.BadExpr{}

	expected := &dst.BinaryExpr{X: expr1, Op: token.LAND, Y: expr2}

	g.Expect(JoinAnd(expr1, expr2)).To(Equal(expected))
}

func TestSelectorWithMoreNames(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	selector := Selector(
		dst.NewIdent("chair"),
		"table",
		"miracle",
		"ecstasy",
	)

	expectedSelector := &dst.SelectorExpr{
		X: &dst.SelectorExpr{
			X: &dst.SelectorExpr{
				X:   dst.NewIdent("chair"),
				Sel: dst.NewIdent("table"),
			},
			Sel: dst.NewIdent("miracle"),
		},
		Sel: dst.NewIdent("ecstasy"),
	}

	g.Expect(selector).To(Equal(expectedSelector))
}

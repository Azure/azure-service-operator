/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"bytes"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/sebdah/goldie/v2"
)

// Assert the given expression generates the expected code, comparing it to a golden file
func AssertExprExpected(t *testing.T, expr dst.Expr) {
	t.Helper()
	g := NewGomegaWithT(t)
	gold := goldie.New(t)

	stmt := &dst.ExprStmt{
		X: expr,
		Decs: dst.ExprStmtDecorations{
			NodeDecs: dst.NodeDecs{
				Start:  []string{"// Generated by " + t.Name()},
				Before: dst.NewLine,
				After:  dst.NewLine,
			},
		},
	}

	block := &dst.BlockStmt{
		List: Statements(stmt),
	}

	fn := &dst.FuncDecl{
		Name: dst.NewIdent("TestCase"),
		Type: &dst.FuncType{},
		Body: block,
		Decs: dst.FuncDeclDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.EmptyLine,
			},
		},
	}

	file := &dst.File{
		Decs: dst.FileDecorations{
			NodeDecs: dst.NodeDecs{
				After: dst.EmptyLine,
			},
		},
		Name:  dst.NewIdent("TestCase"),
		Decls: []dst.Decl{fn},
	}

	var buffer bytes.Buffer
	err := decorator.Fprint(&buffer, file)
	g.Expect(err).To(Succeed())

	gold.Assert(t, t.Name(), buffer.Bytes())
}
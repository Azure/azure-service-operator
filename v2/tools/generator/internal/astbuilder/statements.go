/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"fmt"

	"github.com/dave/dst"
)

// Statements creates a slice of dst.Stmt by combining the given statements.
// Pass any combination of dst.Stmt and []dst.Stmt as arguments; anything else will
// result in a runtime panic.
func Statements(statements ...any) []dst.Stmt {
	// Calculate the final size required
	size := 0
	for _, s := range statements {
		switch s := s.(type) {
		case nil:
			// Skip nils
			continue
		case dst.Stmt:
			size++
		case []dst.Stmt:
			size += len(s)
		case dst.Decl:
			size++
		default:
			panic(fmt.Sprintf("expected dst.Stmt, []dst.Stmt, or dst.Decl but found %T", s))
		}
	}

	// Flatten the statements into a single slice
	stmts := make([]dst.Stmt, 0, size)
	for _, s := range statements {
		switch s := s.(type) {
		case nil:
			// Skip nils
			continue
		case dst.Stmt:
			// Add a single statement
			stmts = append(stmts, s)
		case []dst.Stmt:
			// Add many statements
			stmts = append(stmts, s...)
		case dst.Decl:
			// Convert declaration to statement
			stmt := &dst.DeclStmt{Decl: s}
			stmts = append(stmts, stmt)
		default:
			panic(fmt.Sprintf("expected dst.Stmt, []dst.Stmt, or dst.Decl but found %T", s))
		}
	}

	// Clone everything to avoid sharing nodes
	result := make([]dst.Stmt, 0, len(stmts))
	for _, st := range stmts {
		if st == nil {
			// Skip nils
			continue
		}

		result = append(result, dst.Clone(st).(dst.Stmt))
	}

	return result
}

// StatementBlock generates a block containing the supplied statements
// If we're given a single statement that's already a block, we won't double wrap it
func StatementBlock(statements ...dst.Stmt) *dst.BlockStmt {
	stmts := Statements(statements)

	if len(stmts) == 1 {
		if block, ok := stmts[0].(*dst.BlockStmt); ok {
			return block
		}
	}

	return &dst.BlockStmt{
		List: stmts,
	}
}

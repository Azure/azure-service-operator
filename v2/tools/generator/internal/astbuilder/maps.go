/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"go/token"

	"github.com/dave/dst"
)

// MakeMap returns the call expression for making a map
//
//	make(map[<key>]<value>)
func MakeMap(key dst.Expr, value dst.Expr) *dst.CallExpr {
	return &dst.CallExpr{
		Fun: dst.NewIdent("make"),
		Args: []dst.Expr{
			&dst.MapType{
				Key:   dst.Clone(key).(dst.Expr),
				Value: dst.Clone(value).(dst.Expr),
			},
		},
	}
}

// MakeMapWithCapacity returns the call expression for making a map with a predefined capacity
//
//	make(map[<key>]<value>, <capacity>)
func MakeMapWithCapacity(key dst.Expr, value dst.Expr, capacity dst.Expr) *dst.CallExpr {
	return &dst.CallExpr{
		Fun: dst.NewIdent("make"),
		Args: []dst.Expr{
			&dst.MapType{
				Key:   dst.Clone(key).(dst.Expr),
				Value: dst.Clone(value).(dst.Expr),
			},
			dst.Clone(capacity).(dst.Expr),
		},
	}
}

// InsertMap returns an assignment statement for inserting an item into a map
//
//	<mapExpr>[<key>] = <rhs>
func InsertMap(mapExpr dst.Expr, key dst.Expr, rhs dst.Expr) *dst.AssignStmt {
	return SimpleAssignment(
		&dst.IndexExpr{
			X:     dst.Clone(mapExpr).(dst.Expr),
			Index: dst.Clone(key).(dst.Expr),
		},
		dst.Clone(rhs).(dst.Expr))
}

// IterateOverMapWithValue creates a statement to iterate over the content of a map using the
// specified identifiers for each key and value found.
//
//	for <key>, <item> := range <mapExpr> {
//	    <statements>
//	}
func IterateOverMapWithValue(key string, item string, mapExpr dst.Expr, statements ...dst.Stmt) *dst.RangeStmt {
	return &dst.RangeStmt{
		Key:   dst.NewIdent(key),
		Value: dst.NewIdent(item),
		Tok:   token.DEFINE,
		X:     mapExpr,
		Body:  StatementBlock(statements...),
	}
}

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import "github.com/dave/dst"

type mapPair struct {
	Key   dst.Expr
	Value dst.Expr
}

type MapLiteral struct {
	mapPair
	content []mapPair
}

func NewMapLiteral(key dst.Expr, value dst.Expr) *MapLiteral {
	return &MapLiteral{
		mapPair: mapPair{
			Key:   key,
			Value: value,
		},
	}
}

func (m *MapLiteral) Add(key dst.Expr, value dst.Expr) {
	m.content = append(m.content, mapPair{
		Key:   key,
		Value: value,
	})
}

func (m *MapLiteral) AsExpr() dst.Expr {
	elements := make([]dst.Expr, 0, len(m.content))
	for _, p := range m.content {
		element := &dst.KeyValueExpr{
			Key:   p.Key,
			Value: p.Value,
			Decs: dst.KeyValueExprDecorations{
				NodeDecs: dst.NodeDecs{
					Before: dst.NewLine,
					After:  dst.NewLine,
				},
			},
		}

		elements = append(elements, element)
	}

	return &dst.CompositeLit{
		Type: &dst.MapType{
			Key:   m.Key,
			Value: m.Value,
		},
		Elts: elements,
	}
}

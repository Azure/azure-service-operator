/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"fmt"

	"github.com/dave/dst"
)

func Declarations(declarations ...any) []dst.Decl {
	decls := make([]dst.Decl, 0, len(declarations))
	for _, d := range declarations {
		switch d := d.(type) {
		case nil:
			// Skip nils
			continue
		case dst.Decl:
			// Add a single declaration
			decls = append(decls, d)
		case []dst.Decl:
			// Add many declarations
			decls = append(decls, d...)
		default:
			panic(fmt.Sprintf("expected dst.Decl or []dst.Decl, but found %T", d))
		}
	}

	result := make([]dst.Decl, 0, len(decls))
	for _, d := range decls {
		result = append(result, dst.Clone(d).(dst.Decl))
	}

	return result
}

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"fmt"

	"github.com/dave/dst"
)

// Declarations constructs a slice of dst.Decl by combining the given declarations.
// Pass any combination of dst.Decl and []dst.Decl as arguments; anything else will
// result in a runtime panic.
func Declarations(declarations ...any) []dst.Decl {
	// Calculate the final size required
	size := 0
	for _, d := range declarations {
		switch d := d.(type) {
		case nil:
			// Skip nils
			continue
		case dst.Decl:
			size++
		case []dst.Decl:
			size += len(d)
		default:
			panic(fmt.Sprintf("expected dst.Decl or []dst.Decl, but found %T", d))
		}
	}

	// Flatten the declarations into a single slice
	decls := make([]dst.Decl, 0, size)
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

	// Clone everything to avoid sharing nodes
	result := make([]dst.Decl, 0, len(decls))
	for _, d := range decls {
		result = append(result, dst.Clone(d).(dst.Decl))
	}

	return result
}

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"log"
	"path/filepath"
)

type PackageDefinition struct {
	PackageReference

	definitions []Definition
}

func NewPackageDefinition(reference PackageReference) *PackageDefinition {
	return &PackageDefinition{reference, nil}
}

func (pkgDef *PackageDefinition) AddDefinition(def Definition) {
	pkgDef.definitions = append(pkgDef.definitions, def)
}

func (pkgDef *PackageDefinition) EmitDefinitions(outputDir string) {

	// emit each definition
	for _, def := range pkgDef.definitions {
		genFile := NewFileDefinition(pkgDef.PackageReference, def)
		outputFile := filepath.Join(outputDir, def.FileNameHint()+"_types.go")
		log.Printf("Writing '%s'\n", outputFile)
		genFile.SaveTo(outputFile)
	}
}

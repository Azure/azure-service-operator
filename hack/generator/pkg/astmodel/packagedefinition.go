/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"
	"text/template"
)

// PackageDefinition is the definiton of a package
type PackageDefinition struct {
	PackageReference

	definitions []Definition
}

// NewPackageDefinition creates a new PackageDefinition
func NewPackageDefinition(reference PackageReference) *PackageDefinition {
	return &PackageDefinition{reference, nil}
}

// AddDefinition adds a Definition to the PackageDefinition
func (pkgDef *PackageDefinition) AddDefinition(def Definition) {
	pkgDef.definitions = append(pkgDef.definitions, def)
}

// EmitDefinitions emits the PackageDefinition to an output directory
func (pkgDef *PackageDefinition) EmitDefinitions(outputDir string) {

	defs := filterOutResources(pkgDef.definitions)

	// initialize with 1 resource per file
	filesToGenerate := make(map[string][]Definition)
	for _, resource := range defs.resources {
		filesToGenerate[resource.FileNameHint()] = []Definition{resource}
	}

	allocateTypesToFiles(defs.notResources, filesToGenerate)
	emitFiles(filesToGenerate, outputDir)
	emitGroupVersionFile(pkgDef, outputDir)
}

func emitFiles(filesToGenerate map[string][]Definition, outputDir string) {
	for fileName, defs := range filesToGenerate {
		genFile := NewFileDefinition(defs[0].Reference().PackageReference, defs...)
		outputFile := filepath.Join(outputDir, fileName+"_types.go")
		log.Printf("Writing '%s'\n", outputFile)
		genFile.SaveTo(outputFile)
	}
}

func anyReferences(defs []Definition, t Type) bool {
	for _, def := range defs {
		if def.Type().References(t) {
			return true
		}
	}

	return false
}

type resourcesAndNotResources struct {
	resources    []*StructDefinition
	notResources []Definition
}

func filterOutResources(definitions []Definition) resourcesAndNotResources {

	var resources []*StructDefinition
	var notResources []Definition

	for _, def := range definitions {
		if structDef, ok := def.(*StructDefinition); ok && structDef.IsResource() {
			resources = append(resources, structDef)
		} else {
			notResources = append(notResources, def)
		}
	}

	return resourcesAndNotResources{resources, notResources}
}

func allocateTypesToFiles(typesToAllocate []Definition, filesToGenerate map[string][]Definition) {
	for len(typesToAllocate) > 0 {
		// dequeue!
		typeToAllocate := typesToAllocate[0]
		typesToAllocate = typesToAllocate[1:]

		allocateToFile := allocateTypeToFile(typeToAllocate, filesToGenerate)

		if allocateToFile == "" {
			// couldn't find a file to put it in
			// see if any other types will reference it on a future round
			if !anyReferences(typesToAllocate, typeToAllocate.Reference()) {
				// couldn't find any references, put it in its own file
				allocateToFile = typeToAllocate.FileNameHint()
			}
		}

		if allocateToFile != "" {
			filesToGenerate[allocateToFile] = append(filesToGenerate[allocateToFile], typeToAllocate)
		} else {
			// re-queue it for later, it will eventually be allocated
			typesToAllocate = append(typesToAllocate, typeToAllocate)
		}
	}
}

func allocateTypeToFile(def Definition, filesToGenerate map[string][]Definition) string {

	var allocatedToFile string
	for fileName, fileDefs := range filesToGenerate {
		if anyReferences(fileDefs, def.Reference()) {
			if allocatedToFile == "" {
				allocatedToFile = fileName
			} else if allocatedToFile != fileName {
				// more than one owner... put it in its own file
				return def.FileNameHint()
			}
		}
	}

	return allocatedToFile
}

var groupVersionFileTemplate = template.Must(template.New("groupVersionFile").Parse(`
/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Package {{.PackageName}} contains API Schema definitions for the {{.GroupName}} {{.PackageName}} API group
// +kubebuilder:object:generate=true
// +groupName={{.GroupName}}.infra.azure.com
package {{.PackageName}}

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "{{.GroupName}}.infra.azure.com", Version: "{{.PackageName}}"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme

	localSchemeBuilder = SchemeBuilder.SchemeBuilder
)`))

func emitGroupVersionFile(pkgDef *PackageDefinition, outputDir string) {
	buf := &bytes.Buffer{}
	groupVersionFileTemplate.Execute(buf, pkgDef)

	gvFile := filepath.Join(outputDir, "groupversion_info.go")
	// TODO[dj]: handle this error
	ioutil.WriteFile(gvFile, buf.Bytes(), 0700)
}

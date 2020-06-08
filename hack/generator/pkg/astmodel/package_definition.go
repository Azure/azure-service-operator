/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"text/template"

	"k8s.io/klog/v2"
)

// PackageDefinition is the definition of a package
type PackageDefinition struct {
	GroupName   string
	PackageName string

	definitions []TypeDefiner
}

// NewPackageDefinition constructs a new package definition
func NewPackageDefinition(groupName string, packageName string) *PackageDefinition {
	return &PackageDefinition{groupName, packageName, nil}
}

// AddDefinition adds a Definition to the PackageDefinition
func (pkgDef *PackageDefinition) AddDefinition(def TypeDefiner) {
	pkgDef.definitions = append(pkgDef.definitions, def)
}

// EmitDefinitions emits the PackageDefinition to an output directory
func (pkgDef *PackageDefinition) EmitDefinitions(outputDir string) (int, error) {
	resources, otherDefinitions := partitionDefinitions(pkgDef.definitions)

	// initialize with 1 resource per file
	filesToGenerate := make(map[string][]TypeDefiner)
	for _, resource := range resources {
		filesToGenerate[FileNameHint(resource)] = []TypeDefiner{resource}
	}

	allocateTypesToFiles(otherDefinitions, filesToGenerate)
	err := emitFiles(filesToGenerate, outputDir)
	if err != nil {
		return 0, err
	}

	err = emitGroupVersionFile(pkgDef, outputDir)
	if err != nil {
		return 0, err
	}

	return len(filesToGenerate), nil
}

// DefinitionCount returns the count of definitions that have been sorted into this package
func (pkgDef *PackageDefinition) DefinitionCount() int {
	return len(pkgDef.definitions)
}

func emitFiles(filesToGenerate map[string][]TypeDefiner, outputDir string) error {
	for fileName, defs := range filesToGenerate {
		fullFileName := fileName + "_types" + CodeGeneratedFileSuffix
		genFile := NewFileDefinition(&defs[0].Name().PackageReference, defs...)
		outputFile := filepath.Join(outputDir, fullFileName)

		klog.V(5).Infof("Writing '%s'\n", outputFile)

		err := genFile.SaveToFile(outputFile)
		if err != nil {
			return fmt.Errorf("error saving definitions to file '%v'(%w)", outputFile, err)
		}
	}

	return nil
}

func anyReferences(defs []TypeDefiner, defName *TypeName) bool {
	for _, def := range defs {
		if def.Type().References(defName) {
			return true
		}
	}

	return false
}

func partitionDefinitions(definitions []TypeDefiner) (resourceStructs []*StructDefinition, otherDefinitions []TypeDefiner) {

	var resources []*StructDefinition
	var notResources []TypeDefiner

	for _, def := range definitions {
		if structDef, ok := def.(*StructDefinition); ok && structDef.IsResource() {
			resources = append(resources, structDef)
		} else {
			notResources = append(notResources, def)
		}
	}

	return resources, notResources
}

func allocateTypesToFiles(typesToAllocate []TypeDefiner, filesToGenerate map[string][]TypeDefiner) {
	// Keep track of how many types we've checked since we last allocated one
	// This lets us detect if/when we stall during allocation
	skippedTypes := 0

	// Limit to how many types we skip before taking remedial action
	// len(typesToAllocate) changes during execution, so we cache it
	skipLimit := len(typesToAllocate)

	for len(typesToAllocate) > 0 {
		// dequeue!
		typeToAllocate := typesToAllocate[0]
		typesToAllocate = typesToAllocate[1:]

		filesReferencingType := findFilesReferencingType(typeToAllocate, filesToGenerate)
		hasPendingReferences := anyReferences(typesToAllocate, typeToAllocate.Name())

		allocatedFileName := allocateFileName(filesReferencingType, typeToAllocate, hasPendingReferences)
		if allocatedFileName == "" && skippedTypes > skipLimit {
			// We've stalled during allocation of types, implying we have a cycle of mutually referencing types
			allocatedFileName = allocateFilenameWhenStalled(typesToAllocate, filesToGenerate, len(filesReferencingType), typeToAllocate)
		}

		if allocatedFileName != "" {
			filesToGenerate[allocatedFileName] = append(filesToGenerate[allocatedFileName], typeToAllocate)
			skippedTypes = 0
			continue
		}

		// re-queue it for later, it will eventually be allocated
		typesToAllocate = append(typesToAllocate, typeToAllocate)
		skippedTypes++

	}
}

func allocateFileName(filesReferencingType []string, typeToAllocate TypeDefiner, pendingReferences bool) string {
	if len(filesReferencingType) > 1 {
		// Type is referenced by more than one file, put it in its own file
		return FileNameHint(typeToAllocate)
	}

	if len(filesReferencingType) == 1 && !pendingReferences {
		// Type is only referenced from one file, and is not refeferenced anywhere else, put it in that file
		return filesReferencingType[0]
	}

	if len(filesReferencingType) == 0 && !pendingReferences {
		// Type is not referenced by any file and is not referenced elsewhere, put it in its own file
		return FileNameHint(typeToAllocate)
	}

	// Not allocating to a file at this time
	return ""
}

func allocateFilenameWhenStalled(typesToAllocate []TypeDefiner, filesToGenerate map[string][]TypeDefiner, filesReferencingTypeToAllocate int, typeToAllocate TypeDefiner) string {
	// We've processed the entire queue without allocating any files, so we have a cycle of related types to allocate
	// None of these types are referenced by multiple files (they would already be allocated, per rule above)
	// So either they're referenced by one file, or none at all
	// Breaking the cycle requires allocating one of these types; we need to do this deterministicly
	// So we prefer allocating to an existing file if we can, and we prefer names earlier in the alphabet

	for _, t := range typesToAllocate {
		refs := findFilesReferencingType(t, filesToGenerate)
		if len(refs) > filesReferencingTypeToAllocate {
			// Type 't' would go into an existing file and is a better choice
			// Don't allocate a file name to 'typeToAllocate'
			return ""
		}

		if FileNameHint(t) < FileNameHint(typeToAllocate) {
			// Type `t` is closer to the start of the alphabet and is a better choice
			// Don't allocate a file name to 'typeToAllocate'
			return ""
		}
	}

	// This is the best candidate for breaking the cycle, allocate it to a file of its own
	return FileNameHint(typeToAllocate)
}

func findFilesReferencingType(def TypeDefiner, filesToGenerate map[string][]TypeDefiner) []string {

	var result []string
	for fileName, fileDefs := range filesToGenerate {
		if anyReferences(fileDefs, def.Name()) {
			result = append(result, fileName)
		}
	}

	return result
}

var groupVersionFileTemplate = template.Must(template.New("groupVersionFile").Parse(`
/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by k8s-infra-gen. DO NOT EDIT.

// Package {{.PackageName}} contains API Schema definitions for the {{.GroupName}} {{.PackageName}} API group
// +kubebuilder:object:generate=true
// All object properties are optional by default, this will be overridden when needed:
// +kubebuilder:validation:Optional
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

func emitGroupVersionFile(pkgDef *PackageDefinition, outputDir string) error {
	buf := &bytes.Buffer{}
	err := groupVersionFileTemplate.Execute(buf, pkgDef)
	if err != nil {
		return err
	}

	gvFile := filepath.Join(outputDir, "groupversion_info"+CodeGeneratedFileSuffix)

	err = ioutil.WriteFile(gvFile, buf.Bytes(), 0700)
	if err != nil {
		return fmt.Errorf("error writing group version file '%v'(%w)", gvFile, err)
	}

	return nil
}

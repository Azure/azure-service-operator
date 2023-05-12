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
	"strings"
	"text/template"

	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/pkg/errors"
)

// PackageDefinition is the definition of a package
type PackageDefinition struct {
	GroupName   string
	PackageName string
	definitions TypeDefinitionSet
}

// NewPackageDefinition constructs a new package definition
func NewPackageDefinition(groupName string, packageName string) *PackageDefinition {
	return &PackageDefinition{groupName, packageName, make(TypeDefinitionSet)}
}

func (p *PackageDefinition) Definitions() TypeDefinitionSet {
	return p.definitions
}

func (p *PackageDefinition) GetDefinition(typeName TypeName) (TypeDefinition, error) {
	for _, def := range p.definitions {
		if TypeEquals(def.Name(), typeName) {
			return def, nil
		}
	}

	return TypeDefinition{}, errors.Errorf("no type with name %s found", typeName)
}

// AddDefinition adds a Definition to the PackageDefinition
func (p *PackageDefinition) AddDefinition(def TypeDefinition) {
	p.definitions.Add(def)
}

// EmitDefinitions emits the PackageDefinition to an output directory
func (p *PackageDefinition) EmitDefinitions(outputDir string, generatedPackages map[PackageReference]*PackageDefinition, emitDocFiles bool) (int, error) {
	filesToGenerate := allocateTypesToFiles(p.definitions)

	err := p.emitFiles(filesToGenerate, outputDir, generatedPackages)
	if err != nil {
		return 0, err
	}

	// Check if current package contains resources
	if resources := FindResourceDefinitions(p.definitions); len(resources) > 0 {

		// If package contains resources, then we generate GroupVersion file for the package
		err = emitGroupVersionFile(p, outputDir)
		if err != nil {
			return 0, err
		}

		// If emitDocFiles is true from config and is not Storage package, then we generate doc file for the package
		if !strings.HasSuffix(p.PackageName, StoragePackageSuffix) && emitDocFiles {
			if err = emitDocFile(p, outputDir); err != nil {
				return 0, err
			}
		}
	}

	return len(filesToGenerate), nil
}

// DefinitionCount returns the count of definitions that have been sorted into this package
func (p *PackageDefinition) DefinitionCount() int {
	return len(p.definitions)
}

func (p *PackageDefinition) emitFiles(filesToGenerate map[string][]TypeDefinition, outputDir string, generatedPackages map[PackageReference]*PackageDefinition) error {
	var errs []error

	for fileName, defs := range filesToGenerate {
		codeFilePath := filepath.Join(
			outputDir,
			fmt.Sprintf("%s_types%s.go", fileName, CodeGeneratedFileSuffix))

		err := p.writeCodeFile(codeFilePath, defs, generatedPackages)
		if err != nil {
			errs = append(errs, err)
		}

		testFilePath := filepath.Join(
			outputDir,
			fmt.Sprintf("%s_types%s_test.go", fileName, CodeGeneratedFileSuffix))

		err = p.writeTestFile(testFilePath, defs, generatedPackages)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return kerrors.NewAggregate(errs)
	}

	return nil
}

func (p *PackageDefinition) writeCodeFile(
	outputFile string,
	defs []TypeDefinition,
	packages map[PackageReference]*PackageDefinition,
) error {
	ref := defs[0].Name().PackageReference
	genFile := NewFileDefinition(ref, defs, packages)

	fileWriter := NewGoSourceFileWriter(genFile)
	err := fileWriter.SaveToFile(outputFile)
	if err != nil {
		return errors.Wrapf(err, "saving definitions to file %q", outputFile)
	}

	return nil
}

func (p *PackageDefinition) writeTestFile(
	outputFile string,
	defs []TypeDefinition,
	packages map[PackageReference]*PackageDefinition,
) error {
	// First check to see if we have test cases to write
	haveTestCases := false
	for _, def := range defs {
		if def.HasTestCases() {
			haveTestCases = true
			break
		}
	}

	if !haveTestCases {
		// No test
		return nil
	}

	ref := defs[0].Name().PackageReference
	genFile := NewTestFileDefinition(ref, defs, packages)

	fileWriter := NewGoSourceFileWriter(genFile)
	err := fileWriter.SaveToFile(outputFile)
	if err != nil {
		return errors.Wrapf(err, "writing test cases to file %q", outputFile)
	}

	return nil
}

func allocateTypesToFiles(definitions TypeDefinitionSet) map[string][]TypeDefinition {
	graph := MakeReferenceGraphWithResourcesAsRoots(definitions)

	type Root struct {
		depth int
		name  TypeName
	}

	// rootFor maps a type name to the best root for it
	rootFor := make(map[TypeName]Root, len(graph.roots))

	for root := range graph.roots {
		collected := make(ReachableTypes)
		graph.collectTypes(0, root, collected)
		for node, depth := range collected {
			if currentRoot, ok := rootFor[node]; ok {
				// use depth and then root name as a tiebreaker
				if depth < currentRoot.depth ||
					(depth == currentRoot.depth && root.Name() < currentRoot.name.Name()) {
					rootFor[node] = Root{depth, root}
				}
			} else {
				rootFor[node] = Root{depth, root}
			}
		}
	}

	filesToGenerate := make(map[string][]TypeDefinition, len(definitions))

	for _, def := range definitions {
		var fileName string
		if root, ok := rootFor[def.name]; ok {
			fileName = FileNameHint(root.name)
		} else {
			fileName = FileNameHint(def.name)
		}

		filesToGenerate[fileName] = append(filesToGenerate[fileName], def)
	}

	return filesToGenerate
}

var groupVersionFileTemplate = template.Must(template.New("groupVersionFile").Parse(`/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by azure-service-operator-codegen. DO NOT EDIT.

// Package {{.PackageName}} contains API Schema definitions for the {{.GroupName}} {{.PackageName}} API group
// +kubebuilder:object:generate=true
// All object properties are optional by default, this will be overridden when needed:
// +kubebuilder:validation:Optional
// +groupName={{.GroupName}}.azure.com
package {{.PackageName}}

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "{{.GroupName}}.azure.com", Version: "{{.PackageName}}"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme

	localSchemeBuilder = SchemeBuilder.SchemeBuilder
)
`))

func emitGroupVersionFile(pkgDef *PackageDefinition, outputDir string) error {
	gvFile := filepath.Join(outputDir, "groupversion_info"+CodeGeneratedFileSuffix+".go")
	return emitTemplateFile(pkgDef, groupVersionFileTemplate, gvFile)
}

var docFileTemplate = template.Must(template.New("docFile").Parse(`/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by azure-service-operator-codegen. DO NOT EDIT.

// Package {{.PackageName}} contains API Schema definitions for the {{.GroupName}} {{.PackageName}} API group
// +groupName={{.GroupName}}.azure.com
package {{.PackageName}}
`))

func emitDocFile(pkgDef *PackageDefinition, outputDir string) error {
	docFile := filepath.Join(outputDir, "doc.go")
	return emitTemplateFile(pkgDef, docFileTemplate, docFile)
}

func emitTemplateFile(pkgDef *PackageDefinition, template *template.Template, fileRef string) error {
	buf := &bytes.Buffer{}
	err := template.Execute(buf, pkgDef)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileRef, buf.Bytes(), 0o600)
	if err != nil {
		return errors.Wrapf(err, "error writing file %q", fileRef)
	}

	return nil
}

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rotisserie/eris"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

// PackageDefinition is the definition of a package
type PackageDefinition struct {
	PackageName string                   // Name  of the package
	GroupName   string                   // Group to which the package belongs
	Version     string                   // Kubernetes version of this package
	Path        string                   // relative Path to the package
	definitions TypeDefinitionSet        // set of definitions in this package
	ref         InternalPackageReference // Internal package reference used to reference this package
}

// NewPackageDefinition constructs a new package definition
func NewPackageDefinition(ref InternalPackageReference) *PackageDefinition {
	result := &PackageDefinition{
		PackageName: ref.PackageName(),
		GroupName:   ref.Group(),
		Path:        ref.FolderPath(),
		definitions: make(TypeDefinitionSet),
		ref:         ref,
	}

	result.Version = result.createVersion(ref)

	return result
}

func (p *PackageDefinition) Definitions() TypeDefinitionSet {
	return p.definitions
}

func (p *PackageDefinition) GetDefinition(typeName InternalTypeName) (TypeDefinition, error) {
	if def, ok := p.definitions[typeName]; ok {
		return def, nil
	}

	return TypeDefinition{}, eris.Errorf("no type with name %s found", typeName)
}

// AddDefinition adds a Definition to the PackageDefinition
func (p *PackageDefinition) AddDefinition(def TypeDefinition) {
	p.definitions.Add(def)
}

// EmitDefinitions emits the PackageDefinition to an output directory
func (p *PackageDefinition) EmitDefinitions(
	outputDir string,
	generatedPackages map[InternalPackageReference]*PackageDefinition,
	emitDocFiles bool,
) (int, error) {
	filesToGenerate := allocateTypesToFiles(p.definitions)

	err := p.emitFiles(filesToGenerate, outputDir, generatedPackages)
	if err != nil {
		return 0, err
	}

	packageContainsResources := p.containsResources()
	isStoragePackage := IsStoragePackageReference(p.ref)
	isCompatPackage := IsCompatPackageReference(p.ref)

	// Check if current package contains resources
	if packageContainsResources {
		// generate GroupVersion file for the package
		err = p.emitGroupVersionFile(outputDir)
		if err != nil {
			return 0, err
		}
	} else if isCompatPackage {
		// If it is a compat package, then we generate a stub GroupVersion file
		err = p.emitSubpackageStub(outputDir)
		if err != nil {
			return 0, err
		}
	}

	if emitDocFiles && packageContainsResources && !isStoragePackage {
		// If emitDocFiles is true from config and is not Storage package, then we generate doc file for the package
		if !strings.HasSuffix(p.PackageName, StoragePackageName) && emitDocFiles {
			if err = p.emitDocFile(outputDir); err != nil {
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

func (p *PackageDefinition) emitFiles(
	filesToGenerate map[string][]TypeDefinition,
	outputDir string,
	generatedPackages map[InternalPackageReference]*PackageDefinition,
) error {
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
	packages map[InternalPackageReference]*PackageDefinition,
) error {
	ref := defs[0].Name().InternalPackageReference()
	genFile := NewFileDefinition(ref, defs, packages)

	fileWriter := NewGoSourceFileWriter(genFile)
	err := fileWriter.SaveToFile(outputFile)
	if err != nil {
		return eris.Wrapf(err, "saving definitions to file %q", outputFile)
	}

	return nil
}

func (p *PackageDefinition) writeTestFile(
	outputFile string,
	defs []TypeDefinition,
	packages map[InternalPackageReference]*PackageDefinition,
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

	ref := defs[0].Name().InternalPackageReference()
	genFile := NewTestFileDefinition(ref, defs, packages)

	fileWriter := NewGoSourceFileWriter(genFile)
	err := fileWriter.SaveToFile(outputFile)
	if err != nil {
		return eris.Wrapf(err, "writing test cases to file %q", outputFile)
	}

	return nil
}

func (p *PackageDefinition) createVersion(ref InternalPackageReference) string {
	switch r := ref.(type) {
	case LocalPackageReference:
		return r.Version()
	case DerivedPackageReference:
		return p.createVersion(r.Base()) + ref.PackageName()
	}

	return ""
}

// containsResources returns true if this package contains any resources
func (p *PackageDefinition) containsResources() bool {
	for range p.definitions.AllResources() {
		return true
	}

	return false
}

// emitGroupVersionFile writes a `groupversion_info.go` file for the package
func (p *PackageDefinition) emitGroupVersionFile(outputDir string) error {
	gvFile := filepath.Join(outputDir, "groupversion_info"+CodeGeneratedFileSuffix+".go")
	return p.emitTemplateFile(groupVersionFileTemplate, gvFile)
}

// emitSubpackageStub writes a `subpackage_info_gen.go` stub for subpackages
func (p *PackageDefinition) emitSubpackageStub(outputDir string) error {
	gvFile := filepath.Join(outputDir, "subpackage_info"+CodeGeneratedFileSuffix+".go")
	return p.emitTemplateFile(subpackageStubTemplate, gvFile)
}

func (p *PackageDefinition) emitDocFile(outputDir string) error {
	docFile := filepath.Join(outputDir, "doc.go")
	return p.emitTemplateFile(docFileTemplate, docFile)
}

func (pkgDef *PackageDefinition) emitTemplateFile(template *template.Template, fileRef string) error {
	buf := &bytes.Buffer{}
	err := template.Execute(buf, pkgDef)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileRef, buf.Bytes(), 0o600)
	if err != nil {
		return eris.Wrapf(err, "error writing file %q", fileRef)
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

// groupVersionFileTemplate is a template for the `groupversion_info_gen.go` file generated
// for each API package.
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
// +versionName={{.Version}}
package {{.PackageName}}

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "{{.GroupName}}.azure.com", Version: "{{.Version}}"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme

	localSchemeBuilder = SchemeBuilder.SchemeBuilder
)
`))

// subpackageStubTemplate is a template for the `groupversion_info.go` file generated for
// subpackages within each API package.
var subpackageStubTemplate = template.Must(template.New("groupVersionFile").Parse(`/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by azure-service-operator-codegen. DO NOT EDIT.

// +kubebuilder:object:generate=true
package {{.PackageName}}
`))

// docFileTemplate is a template for the `doc.go` file we generate for each API package.
var docFileTemplate = template.Must(template.New("docFile").Parse(`/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by azure-service-operator-codegen. DO NOT EDIT.

// Package {{.PackageName}} contains API Schema definitions for the {{.GroupName}} {{.Version}} API group
// +groupName={{.GroupName}}.azure.com
package {{.PackageName}}
`))

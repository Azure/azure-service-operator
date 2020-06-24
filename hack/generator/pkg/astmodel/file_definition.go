/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"bufio"
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"sort"
	"strings"

	"k8s.io/klog/v2"
)

// FileDefinition is the content of a file we're generating
type FileDefinition struct {
	// the package this file is in
	packageReference *PackageReference
	// definitions to include in this file
	definitions []TypeDefiner
}

// NewFileDefinition creates a file definition containing specified definitions
func NewFileDefinition(packageRef *PackageReference, definitions ...TypeDefiner) *FileDefinition {

	sort.Slice(definitions, func(i, j int) bool {
		return definitions[i].Name().name < definitions[j].Name().name
	})

	// TODO: check that all definitions are from same package
	return &FileDefinition{packageRef, definitions}
}

// generateImports products the definitive set of imports for use in this file and
// disambiguates any conflicts
func (file *FileDefinition) generateImports() map[PackageImport]struct{} {

	metav1Import := NewPackageImport(
		*NewPackageReference("k8s.io/apimachinery/pkg/apis/meta/v1")).WithName("metav1")

	var requiredImports = make(map[PackageImport]struct{}) // fake set type
	requiredImports[*metav1Import] = struct{}{}

	for _, s := range file.definitions {
		for _, requiredImport := range s.Type().RequiredImports() {
			// no need to import the current package
			if !requiredImport.Equals(file.packageReference) {
				newImport := NewPackageImport(*requiredImport)
				requiredImports[*newImport] = struct{}{}
			}
		}
	}

	// TODO: Do something about conflicting imports

	// Determine if there are any conflicting imports -- these are imports with the same "name"
	// but a different package path
	for imp := range requiredImports {
		for otherImp := range requiredImports {
			if !imp.Equals(&otherImp) && imp.PackageName() == otherImp.PackageName() {
				klog.Warningf(
					"File %v: import %v (named %v) and import %v (named %v) conflict",
					file.packageReference.PackagePath(),
					imp.PackageReference.PackagePath(),
					imp.PackageName(),
					otherImp.PackageReference.PackagePath(),
					otherImp.PackageName())
			}
		}
	}

	return requiredImports
}

func (file *FileDefinition) generateImportSpecs(references map[PackageImport]struct{}) []ast.Spec {
	var importSpecs []ast.Spec
	for requiredImport := range references {
		importSpecs = append(importSpecs, requiredImport.AsImportSpec())
	}

	return importSpecs
}

// AsAst generates an AST node representing this file
func (file *FileDefinition) AsAst() ast.Node {

	var decls []ast.Decl

	// Determine imports
	packageReferences := file.generateImports()

	// Create context from imports
	codeGenContext := NewCodeGenerationContext(file.packageReference, packageReferences)

	// Create import header:
	decls = append(decls, &ast.GenDecl{Tok: token.IMPORT, Specs: file.generateImportSpecs(packageReferences)})

	// Emit all definitions:
	for _, s := range file.definitions {
		decls = append(decls, s.AsDeclarations(codeGenContext)...)
	}

	// Emit struct registration for each resource:
	var exprs []ast.Expr
	for _, defn := range file.definitions {
		if structDefn, ok := defn.(*StructDefinition); ok && structDefn.IsResource() {
			exprs = append(exprs, &ast.UnaryExpr{
				Op: token.AND,
				X:  &ast.CompositeLit{Type: structDefn.Name().AsType(codeGenContext)},
			})
		}
	}

	if len(exprs) > 0 {
		decls = append(decls,
			&ast.FuncDecl{
				Type: &ast.FuncType{Params: &ast.FieldList{}},
				Name: ast.NewIdent("init"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ExprStmt{
							X: &ast.CallExpr{
								Fun:  ast.NewIdent("SchemeBuilder.Register"), // HACK
								Args: exprs,
							},
						},
					},
				},
			})
	}

	header, headerLen := createComments(
		"Copyright (c) Microsoft Corporation.",
		"Licensed under the MIT license.",
		CodeGenerationComment)

	// We set Package (the offset of the package keyword) so that it follows the header comments
	result := &ast.File{
		Doc: &ast.CommentGroup{
			List: header,
		},
		Name:    ast.NewIdent(file.packageReference.PackageName()),
		Decls:   decls,
		Package: token.Pos(headerLen),
	}

	return result
}

// createComments converts a series of strings into a series of comments,
// returning both the comments and their text length
func createComments(lines ...string) ([]*ast.Comment, int) {
	var result []*ast.Comment
	length := 0
	for _, l := range lines {
		line := &ast.Comment{Text: "// " + l + "\n"}
		length += len(line.Text)
		result = append(result, line)
	}

	return result, length
}

// SaveToWriter writes the file to the specifier io.Writer
func (file FileDefinition) SaveToWriter(filename string, dst io.Writer) error {
	original := file.AsAst()

	// Write generated source into a memory buffer
	fset := token.NewFileSet()
	fset.AddFile(filename, 1, 102400)

	var unformattedBuffer bytes.Buffer
	err := format.Node(&unformattedBuffer, fset, original)
	if err != nil {
		return err
	}

	// This is a nasty technique with only one redeeming characteristic: It works
	reformattedBuffer := file.addBlankLinesBeforeComments(unformattedBuffer)

	// Read the source from the memory buffer (has the effect similar to 'go fmt')
	var cleanAst ast.Node
	cleanAst, err = parser.ParseFile(fset, filename, &reformattedBuffer, parser.ParseComments)
	if err != nil {
		klog.Errorf("Failed to reformat code (%s); keeping code as is.", err)
		cleanAst = original
	}

	return format.Node(dst, fset, cleanAst)
}

// SaveToFile writes this generated file to disk
func (file FileDefinition) SaveToFile(filePath string) error {

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer f.Close()

	return file.SaveToWriter(filePath, f)
}

// addBlankLinesBeforeComments reads the source in the passed buffer and injects a blank line just
// before each '//' style comment so that the comments are nicely spaced out in the generated code.
func (file FileDefinition) addBlankLinesBeforeComments(buffer bytes.Buffer) bytes.Buffer {
	// Read all the lines from the buffer
	var lines []string
	reader := bufio.NewReader(&buffer)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	isComment := func(s string) bool {
		return strings.HasPrefix(strings.TrimSpace(s), "//")
	}

	var result bytes.Buffer
	lastLineWasComment := false
	for _, l := range lines {
		// Add blank line prior to each comment block
		if !lastLineWasComment && isComment(l) {
			result.WriteString("\n")
		}

		result.WriteString(l)
		result.WriteString("\n")

		lastLineWasComment = isComment(l)
	}

	return result
}

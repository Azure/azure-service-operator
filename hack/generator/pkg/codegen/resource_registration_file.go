/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"go/token"
	"sort"
	"strings"

	"github.com/dave/dst"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// ResourceRegistrationFile is a file containing functions that assist in registering resources
// with a Kubernetes scheme.
type ResourceRegistrationFile struct {
	resources []astmodel.TypeName
}

var _ astmodel.GoSourceFile = &ResourceRegistrationFile{}

// NewResourceRegistrationFile returns a ResourceRegistrationFile for registering all of the specified resources
// with a controller
func NewResourceRegistrationFile(resources []astmodel.TypeName) *ResourceRegistrationFile {
	return &ResourceRegistrationFile{
		resources: resources,
	}
}

// AsAst returns this file as a Go AST.
func (r *ResourceRegistrationFile) AsAst() (*dst.File, error) {
	var decls []dst.Decl

	// Determine imports
	packageReferences, err := r.generateImports()
	if err != nil {
		return nil, err
	}

	codeGenContext := astmodel.NewCodeGenerationContext(
		astmodel.MakeExternalPackageReference("controllers"), // TODO: This some come from a config
		packageReferences,
		nil)

	// Create import header if needed
	if packageReferences.Length() > 0 {
		decls = append(decls, &dst.GenDecl{
			Decs: dst.GenDeclDecorations{
				NodeDecs: dst.NodeDecs{
					After: dst.EmptyLine,
				},
			},
			Tok:   token.IMPORT,
			Specs: packageReferences.AsImportSpecs(),
		})
	}

	// Create KnownTypes collection
	knownTypes, err := r.createGetKnownTypesFunc(codeGenContext)
	if err != nil {
		return nil, err
	}
	decls = append(decls, knownTypes)

	createSchemeFunc, err := r.createCreateSchemeFunc(codeGenContext)
	if err != nil {
		return nil, err
	}
	decls = append(decls, createSchemeFunc)

	// TODO: Common func?
	var header []string
	header = append(header, astmodel.CodeGenerationComments...)
	header = append(header,
		"// Copyright (c) Microsoft Corporation.",
		"// Licensed under the MIT license.")

	result := &dst.File{
		Decs: dst.FileDecorations{
			NodeDecs: dst.NodeDecs{
				Start: header,
				After: dst.EmptyLine,
			},
		},
		Name:  dst.NewIdent("controllers"), // TODO: This probably needs to come from config
		Decls: decls,
	}

	return result, nil
}

// generateImports generates the a PackageImportSet containing the imports required for the resources
// in the ResourceRegistrationFile.
func (r *ResourceRegistrationFile) generateImports() (*astmodel.PackageImportSet, error) {
	requiredImports := astmodel.NewPackageImportSet()

	for _, typeName := range r.resources {
		// Because versions always end in a date specifier such as "v20180801"
		// they are not unique. We must build a unique name for each package so that
		// there are no import conflicts.
		// This also helps improve code readability.

		localPkg, err := astmodel.PackageAsLocalPackage(typeName.PackageReference)
		if err != nil {
			return nil, err
		}

		imp := astmodel.NewPackageImport(typeName.PackageReference).WithName(r.makeUniquePackageName(localPkg.Group(), localPkg.Version()))
		requiredImports.AddImport(imp)
	}

	// We require these imports
	clientGoSchemePkgRef := makeClientGoKubernetesSchemePackageReference()
	clientGoSchemeImport := astmodel.NewPackageImport(clientGoSchemePkgRef).WithName("clientgoscheme")
	requiredImports.AddImport(clientGoSchemeImport)

	runtimeImport := astmodel.NewPackageImport(makeApiMachineryRuntimePackageReference())
	requiredImports.AddImport(runtimeImport)

	return requiredImports, nil
}

func (r *ResourceRegistrationFile) makeUniquePackageName(group string, version string) string {
	// Strip "microsoft" prefix
	result := strings.TrimPrefix(group, "microsoft")
	result = strings.ReplaceAll(result, ".", "")
	result = result + version

	return result
}

func makeApiMachineryRuntimePackageReference() astmodel.PackageReference {
	return astmodel.MakeExternalPackageReference("k8s.io/apimachinery/pkg/runtime")
}

func makeClientGoKubernetesSchemePackageReference() astmodel.PackageReference {
	return astmodel.MakeExternalPackageReference("k8s.io/client-go/kubernetes/scheme")
}

// createGetKnownTypesFunc creates a getKnownTypes function:
//		func getKnownTypes() []runtime.Object {
//			var result []runtime.Object
//			result = append(result, new(<package>.<resource>))
//			result = append(result, new(<package>.<resource>))
//			result = append(result, new(<package>.<resource>))
//			...
//			return result
//		}
func (r *ResourceRegistrationFile) createGetKnownTypesFunc(codeGenerationContext *astmodel.CodeGenerationContext) (dst.Decl, error) {
	runtime, err := codeGenerationContext.GetImportedPackageName(makeApiMachineryRuntimePackageReference())
	if err != nil {
		return nil, err
	}

	resultIdent := dst.NewIdent("result")
	resultVar := astbuilder.LocalVariableDeclaration(
		resultIdent.String(),
		&dst.ArrayType{
			Elt: &dst.SelectorExpr{
				X:   dst.NewIdent(runtime),
				Sel: dst.NewIdent("Object"),
			},
		},
		"")

	// Sort the resources for a deterministic file layout
	sort.Slice(r.resources, func(i, j int) bool {
		iVal := r.resources[i]
		jVal := r.resources[j]

		iPkgName, err := codeGenerationContext.GetImportedPackageName(iVal.PackageReference)
		if err != nil {
			panic(err)
		}
		jPkgName, err := codeGenerationContext.GetImportedPackageName(jVal.PackageReference)
		if err != nil {
			panic(err)
		}

		return iPkgName < jPkgName || iPkgName == jPkgName && iVal.Name() < jVal.Name()
	})

	var resourceAppendStatements []dst.Stmt
	for _, typeName := range r.resources {
		appendStmt := astbuilder.AppendList(
			resultIdent,
			&dst.CallExpr{
				Fun: dst.NewIdent("new"),
				Args: []dst.Expr{
					typeName.AsType(codeGenerationContext),
				},
			})
		resourceAppendStatements = append(resourceAppendStatements, appendStmt)
	}

	returnStmt := &dst.ReturnStmt{
		Results: []dst.Expr{
			resultIdent,
		},
	}

	var body []dst.Stmt
	body = append(body, resultVar)
	body = append(body, resourceAppendStatements...)
	body = append(body, returnStmt)

	f := &astbuilder.FuncDetails{
		Name:   "getKnownTypes",
		Body:   body,
		Params: []*dst.Field{},
		Returns: []*dst.Field{
			{
				Type: &dst.ArrayType{
					Elt: &dst.SelectorExpr{
						X:   dst.NewIdent(runtime),
						Sel: dst.NewIdent("Object"),
					},
				},
			},
		},
	}
	f.AddComments("returns the list of known types which can be reconciled")

	return f.DefineFunc(), nil
}

// createCreateSchemeFunc creates a createScheme() function like:
//		func createScheme() *runtime.Scheme {
//			scheme := runtime.NewScheme()
//			_ = clientgoscheme.AddToScheme(scheme)
//			_ = batchv20170901.AddToScheme(scheme)
//			_ = documentdbv20150408.AddToScheme(scheme)
//			_ = storagev20190401.AddToScheme(scheme)
//			return scheme
//		}
func (r *ResourceRegistrationFile) createCreateSchemeFunc(codeGenerationContext *astmodel.CodeGenerationContext) (dst.Decl, error) {
	runtime, err := codeGenerationContext.GetImportedPackageName(makeApiMachineryRuntimePackageReference())
	if err != nil {
		return nil, err
	}

	clientGoScheme, err := codeGenerationContext.GetImportedPackageName(makeClientGoKubernetesSchemePackageReference())
	if err != nil {
		return nil, err
	}

	scheme := "scheme"
	ignore := "_"
	addToScheme := "AddToScheme"

	initSchemeVar := astbuilder.SimpleAssignment(
		dst.NewIdent(scheme),
		token.DEFINE,
		&dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   dst.NewIdent(runtime),
				Sel: dst.NewIdent("NewScheme"),
			},
			Args: []dst.Expr{},
		})

	clientGoSchemeAssign := astbuilder.SimpleAssignment(
		dst.NewIdent(ignore),
		token.ASSIGN,
		&dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   dst.NewIdent(clientGoScheme),
				Sel: dst.NewIdent(addToScheme),
			},
			Args: []dst.Expr{
				dst.NewIdent(scheme),
			},
		})

	var importedPackageNames []string
	for pkg := range r.getImportedPackages() {
		packageName, err := codeGenerationContext.GetImportedPackageName(pkg)
		if err != nil {
			return nil, err
		}
		importedPackageNames = append(importedPackageNames, packageName)
	}

	// Sort the slice for reproducibility
	sort.Slice(importedPackageNames, func(left int, right int) bool {
		return importedPackageNames[left] < importedPackageNames[right]
	})

	var groupVersionAssignments []dst.Stmt
	for _, group := range importedPackageNames {
		groupSchemeAssign := astbuilder.SimpleAssignment(
			dst.NewIdent(ignore),
			token.ASSIGN,
			&dst.CallExpr{
				Fun: &dst.SelectorExpr{
					X:   dst.NewIdent(group),
					Sel: dst.NewIdent(addToScheme),
				},
				Args: []dst.Expr{
					dst.NewIdent(scheme),
				},
			})
		groupVersionAssignments = append(groupVersionAssignments, groupSchemeAssign)
	}

	returnStmt := &dst.ReturnStmt{
		Results: []dst.Expr{
			dst.NewIdent(scheme),
		},
	}

	var body []dst.Stmt
	body = append(body, initSchemeVar)
	body = append(body, clientGoSchemeAssign)
	body = append(body, groupVersionAssignments...)
	body = append(body, returnStmt)

	f := &astbuilder.FuncDetails{
		Name:   "createScheme",
		Body:   body,
		Params: []*dst.Field{},
		Returns: []*dst.Field{
			{
				Type: &dst.UnaryExpr{
					Op: token.MUL,
					X: &dst.SelectorExpr{
						X:   dst.NewIdent(runtime),
						Sel: dst.NewIdent("Scheme"),
					},
				},
			},
		},
	}
	f.AddComments("creates a Scheme containing the clientgo types and all of the custom types returned by getKnownTypes")

	return f.DefineFunc(), nil
}

func (r *ResourceRegistrationFile) getImportedPackages() map[astmodel.PackageReference]struct{} {
	result := make(map[astmodel.PackageReference]struct{})
	for _, typeName := range r.resources {
		result[typeName.PackageReference] = struct{}{}
	}

	return result
}

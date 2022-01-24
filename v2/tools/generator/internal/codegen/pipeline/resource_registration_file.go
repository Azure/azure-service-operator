/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"go/token"
	"sort"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// ResourceRegistrationFile is a file containing functions that assist in registering resources
// with a Kubernetes scheme.
type ResourceRegistrationFile struct {
	resources               []astmodel.TypeName
	storageVersionResources []astmodel.TypeName
}

var _ astmodel.GoSourceFile = &ResourceRegistrationFile{}

// NewResourceRegistrationFile returns a ResourceRegistrationFile for registering all of the specified resources
// with a controller
func NewResourceRegistrationFile(resources []astmodel.TypeName, storageVersionResources []astmodel.TypeName) *ResourceRegistrationFile {
	return &ResourceRegistrationFile{
		resources:               resources,
		storageVersionResources: storageVersionResources,
	}
}

// AsAst returns this file as a Go AST.
func (r *ResourceRegistrationFile) AsAst() (*dst.File, error) {
	var decls []dst.Decl

	// Determine imports
	packageReferences := r.generateImports()

	codeGenContext := astmodel.NewCodeGenerationContext(
		astmodel.MakeExternalPackageReference("controllers"), // TODO: This should come from a config
		packageReferences,
		nil)

	// Create import header if needed
	thing := packageReferences.AsImportSpecs()
	if packageReferences.Length() > 0 {
		decls = append(decls, &dst.GenDecl{
			Decs: dst.GenDeclDecorations{
				NodeDecs: dst.NodeDecs{
					After: dst.EmptyLine,
				},
			},
			Tok:   token.IMPORT,
			Specs: thing,
		})
	}

	knownStorageTypes, err := r.createGetKnownStorageTypesFunc(codeGenContext)
	if err != nil {
		return nil, err
	}
	decls = append(decls, knownStorageTypes)

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
func (r *ResourceRegistrationFile) generateImports() *astmodel.PackageImportSet {
	requiredImports := astmodel.NewPackageImportSet()

	for _, typeName := range append(r.resources, r.storageVersionResources...) {
		// Because versions always end in a date specifier such as "v20180801"
		// they are not unique. We must build a unique name for each package so that
		// there are no import conflicts.
		imp := astmodel.NewPackageImport(typeName.PackageReference)
		imp = imp.WithName(imp.VersionedNameForImport())
		requiredImports.AddImport(imp)
	}

	// We require these imports
	clientGoSchemeImport := astmodel.NewPackageImport(astmodel.ClientGoSchemeReference).WithName("clientgoscheme")
	requiredImports.AddImport(clientGoSchemeImport)

	runtimeImport := astmodel.NewPackageImport(astmodel.APIMachineryRuntimeReference)
	requiredImports.AddImport(runtimeImport)

	clientImport := astmodel.NewPackageImport(astmodel.ControllerRuntimeClient)
	requiredImports.AddImport(clientImport)

	return requiredImports
}

// createGetKnownStorageTypesFunc creates a getKnownStorageTypes function that returns all storage types:
//		func getKnownStorageTypes() []client.Object {
//			var result []client.Object
//			result = append(result, new(<package>.<resource>))
//			result = append(result, new(<package>.<resource>))
//			result = append(result, new(<package>.<resource>))
//			...
//			return result
//		}
func (r *ResourceRegistrationFile) createGetKnownStorageTypesFunc(codeGenerationContext *astmodel.CodeGenerationContext) (dst.Decl, error) {
	return createKnownTypesFuncImpl(
		codeGenerationContext,
		r.storageVersionResources,
		"getKnownStorageTypes",
		"returns the list of storage types which can be reconciled.")
}

// createGetKnownTypesFunc creates a getKnownTypes function that returns all known types:
//		func getKnownTypes() []client.Object {
//			var result []client.Object
//			result = append(result, new(<package>.<resource>))
//			result = append(result, new(<package>.<resource>))
//			result = append(result, new(<package>.<resource>))
//			...
//			return result
//		}
func (r *ResourceRegistrationFile) createGetKnownTypesFunc(codeGenerationContext *astmodel.CodeGenerationContext) (dst.Decl, error) {
	return createKnownTypesFuncImpl(
		codeGenerationContext,
		r.resources,
		"getKnownTypes",
		"returns the list of all types.")
}

func createKnownTypesFuncImpl(codeGenerationContext *astmodel.CodeGenerationContext, resources []astmodel.TypeName, funcName string, funcComment string) (dst.Decl, error) {
	client, err := codeGenerationContext.GetImportedPackageName(astmodel.ControllerRuntimeClient)
	if err != nil {
		return nil, err
	}

	resultIdent := dst.NewIdent("result")
	resultVar := astbuilder.LocalVariableDeclaration(
		resultIdent.String(),
		&dst.ArrayType{
			Elt: &dst.SelectorExpr{
				X:   dst.NewIdent(client),
				Sel: dst.NewIdent("Object"),
			},
		},
		"")

	// Sort the resources for a deterministic file layout
	sort.Slice(resources, func(i, j int) bool {
		iVal := resources[i]
		jVal := resources[j]

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
	for _, typeName := range resources {
		appendStmt := astbuilder.AppendSlice(
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
		Name:   funcName,
		Body:   body,
		Params: []*dst.Field{},
		Returns: []*dst.Field{
			{
				Type: &dst.ArrayType{
					Elt: &dst.SelectorExpr{
						X:   dst.NewIdent(client),
						Sel: dst.NewIdent("Object"),
					},
				},
			},
		},
	}
	f.AddComments(funcComment)

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
	runtime, err := codeGenerationContext.GetImportedPackageName(astmodel.APIMachineryRuntimeReference)
	if err != nil {
		return nil, err
	}

	clientGoScheme, err := codeGenerationContext.GetImportedPackageName(astmodel.ClientGoSchemeReference)
	if err != nil {
		return nil, err
	}

	scheme := "scheme"
	ignore := "_"
	addToScheme := "AddToScheme"

	initSchemeVar := astbuilder.ShortDeclaration(scheme, astbuilder.CallQualifiedFunc(runtime, "NewScheme"))

	clientGoSchemeAssign := astbuilder.SimpleAssignment(
		dst.NewIdent(ignore),
		astbuilder.CallQualifiedFunc(clientGoScheme, addToScheme, dst.NewIdent(scheme)))

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
			astbuilder.CallQualifiedFunc(group, addToScheme, dst.NewIdent(scheme)))

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
		Name: "createScheme",
		Body: body,
	}

	f.AddReturn(
		astbuilder.Dereference(
			astbuilder.Selector(dst.NewIdent(runtime), "Scheme")))

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

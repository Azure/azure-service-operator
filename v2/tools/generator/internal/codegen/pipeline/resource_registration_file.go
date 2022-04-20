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
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// ResourceRegistrationFile is a file containing functions that assist in registering resources
// with a Kubernetes scheme.
type ResourceRegistrationFile struct {
	resources               []astmodel.TypeName
	storageVersionResources []astmodel.TypeName
	resourceExtensions      []astmodel.TypeName
	indexFunctions          map[astmodel.TypeName][]*functions.IndexRegistrationFunction
	secretPropertyKeys      map[astmodel.TypeName][]string
}

var _ astmodel.GoSourceFile = &ResourceRegistrationFile{}

// NewResourceRegistrationFile returns a ResourceRegistrationFile for registering all of the specified resources
// with a controller
func NewResourceRegistrationFile(
	resources []astmodel.TypeName,
	storageVersionResources []astmodel.TypeName,
	indexFunctions map[astmodel.TypeName][]*functions.IndexRegistrationFunction,
	secretPropertyKeys map[astmodel.TypeName][]string,
	resourceExtensions []astmodel.TypeName) *ResourceRegistrationFile {

	return &ResourceRegistrationFile{
		resources:               resources,
		storageVersionResources: storageVersionResources,
		indexFunctions:          indexFunctions,
		secretPropertyKeys:      secretPropertyKeys,
		resourceExtensions:      resourceExtensions,
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

	// getKnownStorageTypes() function
	knownStorageTypes := createGetKnownStorageTypesFunc(
		codeGenContext,
		r.storageVersionResources,
		r.indexFunctions,
		r.secretPropertyKeys)
	decls = append(decls, knownStorageTypes)

	// getKnownTypes() function
	knownTypes, err := createGetKnownTypesFunc(codeGenContext, r.resources)
	if err != nil {
		return nil, err
	}
	decls = append(decls, knownTypes)

	// createScheme() function
	createSchemeFunc, err := r.createCreateSchemeFunc(codeGenContext)
	if err != nil {
		return nil, err
	}
	decls = append(decls, createSchemeFunc)

	// Create Resource Extensions
	resourceExtensionTypes := r.createGetResourceExtensions(codeGenContext)
	if err != nil {
		return nil, err
	}
	decls = append(decls, resourceExtensionTypes)

	// All the index functions
	indexFunctionDecls := r.defineIndexFunctions(codeGenContext)
	decls = append(decls, indexFunctionDecls...)

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
	typeSet := append(r.resources, r.storageVersionResources...)
	typeSet = append(typeSet, r.resourceExtensions...)
	for _, typeName := range typeSet {
		requiredImports.AddImportOfReference(typeName.PackageReference)
	}

	// We require these imports
	requiredImports.AddImport(astmodel.NewPackageImport(astmodel.GenRuntimeReference))
	requiredImports.AddImport(astmodel.NewPackageImport(astmodel.ClientGoSchemeReference).WithName("clientgoscheme"))
	requiredImports.AddImport(astmodel.NewPackageImport(astmodel.APIMachineryRuntimeReference))
	requiredImports.AddImport(astmodel.NewPackageImport(astmodel.ControllerRuntimeClient))
	requiredImports.AddImport(astmodel.NewPackageImport(astmodel.ControllerRuntimeSource))
	requiredImports.AddImport(astmodel.NewPackageImport(astmodel.GenRuntimeRegistrationReference))
	requiredImports.AddImport(astmodel.NewPackageImport(astmodel.CoreV1Reference))

	return requiredImports
}

func orderByImportedTypeName(codeGenerationContext *astmodel.CodeGenerationContext, resources []astmodel.TypeName) func(i, j int) bool {
	return func(i, j int) bool {
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
	}
}

func orderByFunctionName(functions []*functions.IndexRegistrationFunction) func(i, j int) bool {
	return func(i, j int) bool {
		iVal := functions[i]
		jVal := functions[j]

		return iVal.Name() < jVal.Name()
	}
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
func createGetKnownTypesFunc(codeGenerationContext *astmodel.CodeGenerationContext, resources []astmodel.TypeName) (dst.Decl, error) {
	funcName := "getKnownTypes"
	funcComment := "returns the list of all types."

	client, err := codeGenerationContext.GetImportedPackageName(astmodel.ControllerRuntimeClient)
	if err != nil {
		return nil, err
	}

	resultIdent := dst.NewIdent("result")
	resultVar := astbuilder.LocalVariableDeclaration(
		resultIdent.String(),
		&dst.ArrayType{
			Elt: astbuilder.Selector(dst.NewIdent(client), "Object"),
		},
		"")

	// Sort the resources for a deterministic file layout
	sort.Slice(resources, orderByImportedTypeName(codeGenerationContext, resources))

	var resourceAppendStatements []dst.Stmt
	for _, typeName := range resources {
		appendStmt := astbuilder.AppendSlice(
			resultIdent,
			astbuilder.CallFunc("new", typeName.AsType(codeGenerationContext)))
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
					Elt: astbuilder.Selector(dst.NewIdent(client), "Object"),
				},
			},
		},
	}
	f.AddComments(funcComment)

	return f.DefineFunc(), nil
}

// createGetKnownStorageTypesFunc creates a getKnownStorageTypes function that returns all storage types:
//		func getKnownStorageTypes() []registration.StorageType {
//			var result []*registration.StorageType
//			result = append(result, &registration.StorageType{
//				Obj: new(<package>.<resource>),
//				Indexes: []registration.Index{
//					{
//						Key: <key>,
//						Func: <func>,
//					},
//				},
//				Watches: []registration.Watch{
//					{
//						Src: <source> (usually corev1.Secret{}),
//					},
//				},
//			})
//			result = append(result, registration.StorageType{Obj: new(<package>.<resource>)})
//			result = append(result, registration.StorageType{Obj: new(<package>.<resource>)})
//			...
//			return result
//		}
func createGetKnownStorageTypesFunc(
	codeGenerationContext *astmodel.CodeGenerationContext,
	resources []astmodel.TypeName,
	indexFunctions map[astmodel.TypeName][]*functions.IndexRegistrationFunction,
	secretPropertyKeys map[astmodel.TypeName][]string) dst.Decl {

	funcName := "getKnownStorageTypes"
	funcComment := "returns the list of storage types which can be reconciled."

	resultIdent := dst.NewIdent("result")
	resultVar := astbuilder.LocalVariableDeclaration(
		resultIdent.String(),
		astmodel.NewArrayType(astmodel.NewOptionalType(astmodel.StorageTypeRegistrationType)).AsType(codeGenerationContext),
		"")

	sort.Slice(resources, orderByImportedTypeName(codeGenerationContext, resources))

	var resourceAppendStatements []dst.Stmt
	for _, typeName := range resources {
		newStorageTypeBuilder := astbuilder.NewCompositeLiteralBuilder(astmodel.StorageTypeRegistrationType.AsType(codeGenerationContext))
		newStorageTypeBuilder.AddField("Obj", astbuilder.CallFunc("new", typeName.AsType(codeGenerationContext)))

		// Register index functions (if needed):
		// Indexes: []registration.Index{
		//		{
		//			Key: <key>,
		//			Func: <func>,
		//		}
		//	}
		if indexFuncs, ok := indexFunctions[typeName]; ok {
			sliceBuilder := astbuilder.NewSliceLiteralBuilder(astmodel.IndexRegistrationType.AsType(codeGenerationContext), true)

			for _, indexFunc := range indexFuncs {
				newIndexFunctionBuilder := astbuilder.NewCompositeLiteralBuilder(astmodel.IndexRegistrationType.AsType(codeGenerationContext))
				newIndexFunctionBuilder.AddField("Key", astbuilder.StringLiteral(indexFunc.IndexKey()))
				newIndexFunctionBuilder.AddField("Func", dst.NewIdent(indexFunc.Name()))
				sliceBuilder.AddElement(newIndexFunctionBuilder.Build())
			}
			// astbuilder.SliceLiteral(astmodel.IndexRegistrationType.AsType(codeGenerationContext), indexRegistrations...)
			newStorageTypeBuilder.AddField("Indexes", sliceBuilder.Build())
		}

		// Register additional watches (if needed):
		// Watches: []registration.Watch{
		//		registration.Watch{
		//			Src: <event source>,
		//			MakeEventHandler: <func>,
		//		}
		//	}
		if secretKeys, ok := secretPropertyKeys[typeName]; ok {
			newWatchBuilder := astbuilder.NewCompositeLiteralBuilder(astmodel.WatchRegistrationType.AsType(codeGenerationContext))
			// Not using astbuilder here because we want this to go onto a single line
			source := &dst.CompositeLit{
				Type: astmodel.ControllerRuntimeSourceKindType.AsType(codeGenerationContext),
				Elts: []dst.Expr{
					&dst.KeyValueExpr{
						Key: dst.NewIdent("Type"),
						Value: astbuilder.AddrOf(
							&dst.CompositeLit{
								Type: astmodel.SecretType.AsType(codeGenerationContext),
							}),
					},
				},
			}
			newWatchBuilder.AddField("Src", astbuilder.AddrOf(source))

			// This is so messy and could be done much cleaner with select, if there was one
			var secretKeyParams []dst.Expr
			for _, secretKey := range secretKeys {
				secretKeyParams = append(secretKeyParams, astbuilder.StringLiteral(secretKey))
			}

			// This is a bit hacky
			listTypeName := typeName.WithName(typeName.Name() + "List").AsType(codeGenerationContext)

			eventHandler := astbuilder.CallFunc(
				"watchSecretsFactory",
				astbuilder.SliceLiteral(dst.NewIdent("string"), secretKeyParams...),
				astbuilder.AddrOf(astbuilder.NewCompositeLiteralBuilder(listTypeName).Build()))
			newWatchBuilder.AddField("MakeEventHandler", eventHandler)

			sliceBuilder := astbuilder.NewSliceLiteralBuilder(astmodel.WatchRegistrationType.AsType(codeGenerationContext), true)
			if len(secretKeys) > 0 {
				sliceBuilder.AddElement(newWatchBuilder.Build())
			}

			newStorageTypeBuilder.AddField("Watches", sliceBuilder.Build())
		}

		appendStmt := astbuilder.AppendSlice(
			resultIdent,
			astbuilder.AddrOf(newStorageTypeBuilder.Build()))
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
				Type: astmodel.NewArrayType(
					astmodel.NewOptionalType(astmodel.StorageTypeRegistrationType)).AsType(codeGenerationContext),
			},
		},
	}
	f.AddComments(funcComment)

	return f.DefineFunc()
}

func (r *ResourceRegistrationFile) createGetResourceExtensions(context *astmodel.CodeGenerationContext) dst.Decl {

	funcName := "getResourceExtensions"
	funcComment := "returns a list of resource extensions"

	sort.Slice(r.resourceExtensions, orderByImportedTypeName(context, r.resourceExtensions))
	resultIdent := dst.NewIdent("result")
	resultVar := astbuilder.LocalVariableDeclaration(
		resultIdent.String(),
		astmodel.NewArrayType(astmodel.ResourceExtensionType).AsType(context),
		"")

	var resourceAppendStatements []dst.Stmt
	for _, typeName := range r.resourceExtensions {
		appendStmt := astbuilder.AppendSlice(
			resultIdent,
			astbuilder.AddrOf(astbuilder.NewCompositeLiteralBuilder(typeName.AsType(context)).Build()),
		)
		resourceAppendStatements = append(resourceAppendStatements, appendStmt)
	}

	returnStmt := astbuilder.Returns(resultIdent)

	var body []dst.Stmt
	body = append(body, resultVar)
	body = append(body, resourceAppendStatements...)
	body = append(body, returnStmt)

	f := &astbuilder.FuncDetails{
		Name: funcName,
		Body: body,
	}

	f.AddReturn(astmodel.NewArrayType(astmodel.ResourceExtensionType).AsType(context))
	f.AddComments(funcComment)

	return f.DefineFunc()
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

func (r *ResourceRegistrationFile) defineIndexFunctions(codeGenerationContext *astmodel.CodeGenerationContext) []dst.Decl {
	var result []dst.Decl
	var indexFunctions []*functions.IndexRegistrationFunction

	for _, funcs := range r.indexFunctions {
		indexFunctions = append(indexFunctions, funcs...)
	}
	sort.Slice(indexFunctions, orderByFunctionName(indexFunctions))

	for _, f := range indexFunctions {
		result = append(result, f.AsFunc(codeGenerationContext, astmodel.TypeName{}))
	}

	return result
}

func (r *ResourceRegistrationFile) getImportedPackages() map[astmodel.PackageReference]struct{} {
	result := make(map[astmodel.PackageReference]struct{})
	for _, typeName := range r.resources {
		result[typeName.PackageReference] = struct{}{}
	}

	return result
}

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
	configMapPropertyKeys   map[astmodel.TypeName][]string
}

var _ astmodel.GoSourceFile = &ResourceRegistrationFile{}

// NewResourceRegistrationFile returns a ResourceRegistrationFile for registering all of the specified resources
// with a controller
func NewResourceRegistrationFile(
	resources []astmodel.TypeName,
	storageVersionResources []astmodel.TypeName,
	indexFunctions map[astmodel.TypeName][]*functions.IndexRegistrationFunction,
	secretPropertyKeys map[astmodel.TypeName][]string,
	configMapPropertyKeys map[astmodel.TypeName][]string,
	resourceExtensions []astmodel.TypeName,
) *ResourceRegistrationFile {
	return &ResourceRegistrationFile{
		resources:               resources,
		storageVersionResources: storageVersionResources,
		indexFunctions:          indexFunctions,
		secretPropertyKeys:      secretPropertyKeys,
		configMapPropertyKeys:   configMapPropertyKeys,
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
	knownStorageTypes := r.createGetKnownStorageTypesFunc(codeGenContext)
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

// generateImports generates the PackageImportSet containing the imports required for the resources
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
//
//	func getKnownTypes() []client.Object {
//		var result []client.Object
//		result = append(result, new(<package>.<resource>))
//		result = append(result, new(<package>.<resource>))
//		result = append(result, new(<package>.<resource>))
//		...
//		return result
//	}
func createGetKnownTypesFunc(codeGenerationContext *astmodel.CodeGenerationContext, resources []astmodel.TypeName) (dst.Decl, error) {
	funcName := "getKnownTypes"
	funcComment := "returns the list of all types."

	client, err := codeGenerationContext.GetImportedPackageName(astmodel.ControllerRuntimeClient)
	if err != nil {
		return nil, err
	}

	resultIdent := dst.NewIdent("result")
	resultType := &dst.ArrayType{
		Elt: astbuilder.Selector(dst.NewIdent(client), "Object"),
	}
	resultVar := astbuilder.LocalVariableDeclaration(resultIdent.String(), resultType, "")

	// Sort the resources for a deterministic file layout
	sort.Slice(resources, orderByImportedTypeName(codeGenerationContext, resources))

	resourceAppendStatements := make([]dst.Stmt, 0, len(resources))
	batch := make([]dst.Expr, 0, 10)
	var lastPkg astmodel.PackageReference
	for _, typeName := range resources {
		if len(batch) > 0 && typeName.PackageReference != lastPkg {
			appendStmt := astbuilder.AppendItemsToSlice(resultIdent, batch...)
			resourceAppendStatements = append(resourceAppendStatements, appendStmt)
			batch = batch[:0]
		}

		batch = append(batch, astbuilder.CallFunc("new", typeName.AsType(codeGenerationContext)))
		lastPkg = typeName.PackageReference
	}

	if len(batch) > 0 {
		appendStmt := astbuilder.AppendItemsToSlice(resultIdent, batch...)
		resourceAppendStatements = append(resourceAppendStatements, appendStmt)
	}

	returnStmt := astbuilder.Returns(resultIdent)

	body := astbuilder.Statements(resultVar, resourceAppendStatements, returnStmt)

	f := &astbuilder.FuncDetails{
		Name: funcName,
		Body: body,
	}

	f.AddReturn(resultType)
	f.AddComments(funcComment)

	return f.DefineFunc(), nil
}

// createGetKnownStorageTypesFunc creates a getKnownStorageTypes function that returns all storage types:
//
//	func getKnownStorageTypes() []registration.StorageType {
//		var result []*registration.StorageType
//		result = append(result, &registration.StorageType{
//			Obj: new(<package>.<resource>),
//			Indexes: []registration.Index{
//				{
//					Key: <key>,
//					Func: <func>,
//				},
//			},
//			Watches: []registration.Watch{
//				{
//					Src: <source> (usually corev1.Secret{}),
//				},
//			},
//		})
//		result = append(result, registration.StorageType{Obj: new(<package>.<resource>)})
//		result = append(result, registration.StorageType{Obj: new(<package>.<resource>)})
//		...
//		return result
//	}
func (r *ResourceRegistrationFile) createGetKnownStorageTypesFunc(
	codeGenerationContext *astmodel.CodeGenerationContext,
) dst.Decl {

	funcName := "getKnownStorageTypes"
	funcComment := "returns the list of storage types which can be reconciled."

	resultIdent := dst.NewIdent("result")
	resultVar := astbuilder.LocalVariableDeclaration(
		resultIdent.String(),
		astmodel.NewArrayType(astmodel.NewOptionalType(astmodel.StorageTypeRegistrationType)).AsType(codeGenerationContext),
		"")

	sort.Slice(r.storageVersionResources, orderByImportedTypeName(codeGenerationContext, r.storageVersionResources))

	resourceAppendStatements := make([]dst.Stmt, 0, len(r.storageVersionResources))
	for _, typeName := range r.storageVersionResources {
		newStorageTypeBuilder := astbuilder.NewCompositeLiteralBuilder(astmodel.StorageTypeRegistrationType.AsType(codeGenerationContext))
		newStorageTypeBuilder.AddField("Obj", astbuilder.CallFunc("new", typeName.AsType(codeGenerationContext)))

		// Register index functions (if needed):
		// Indexes: []registration.Index{
		//		{
		//			Key: <key>,
		//			Func: <func>,
		//		}
		//	}
		if indexFuncs, ok := r.indexFunctions[typeName]; ok {
			sliceBuilder := astbuilder.NewSliceLiteralBuilder(astmodel.IndexRegistrationType.AsType(codeGenerationContext), true)
			sort.Slice(indexFuncs, orderByFunctionName(indexFuncs))
			if len(indexFuncs) > 0 {
				for _, indexFunc := range indexFuncs {
					newIndexFunctionBuilder := astbuilder.NewCompositeLiteralBuilder(astmodel.IndexRegistrationType.AsType(codeGenerationContext))
					newIndexFunctionBuilder.AddField("Key", astbuilder.StringLiteral(indexFunc.IndexKey()))
					newIndexFunctionBuilder.AddField("Func", dst.NewIdent(indexFunc.Name()))
					sliceBuilder.AddElement(newIndexFunctionBuilder.Build())
				}

				newStorageTypeBuilder.AddField("Indexes", sliceBuilder.Build())
			}
		}

		// Register additional watches (if needed):
		// Watches: []registration.Watch{
		//		registration.Watch{
		//			Src: <event source>,
		//			MakeEventHandler: <func>,
		//		}
		//	}
		watches := r.makeWatchesExpr(typeName, codeGenerationContext)
		if watches != nil {
			newStorageTypeBuilder.AddField("Watches", watches)
		}

		appendStmt := astbuilder.AppendItemToSlice(
			resultIdent,
			astbuilder.AddrOf(newStorageTypeBuilder.Build()))
		resourceAppendStatements = append(resourceAppendStatements, appendStmt)
	}

	returnStmt := &dst.ReturnStmt{
		Results: []dst.Expr{
			resultIdent,
		},
	}

	body := astbuilder.Statements(resultVar, resourceAppendStatements, returnStmt)

	f := &astbuilder.FuncDetails{
		Name: funcName,
		Body: body,
	}

	f.AddReturn(
		astmodel.NewArrayType(
			astmodel.NewOptionalType(astmodel.StorageTypeRegistrationType)).
			AsType(codeGenerationContext))
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

	resourceAppendStatements := make([]dst.Stmt, 0, len(r.resourceExtensions))
	for _, typeName := range r.resourceExtensions {
		appendStmt := astbuilder.AppendItemToSlice(
			resultIdent,
			astbuilder.AddrOf(astbuilder.NewCompositeLiteralBuilder(typeName.AsType(context)).Build()),
		)
		resourceAppendStatements = append(resourceAppendStatements, appendStmt)
	}

	returnStmt := astbuilder.Returns(resultIdent)

	body := astbuilder.Statements(resultVar, resourceAppendStatements, returnStmt)

	f := &astbuilder.FuncDetails{
		Name: funcName,
		Body: body,
	}

	f.AddReturn(astmodel.NewArrayType(astmodel.ResourceExtensionType).AsType(context))
	f.AddComments(funcComment)

	return f.DefineFunc()
}

// createCreateSchemeFunc creates a createScheme() function like:
//
//	func createScheme() *runtime.Scheme {
//		scheme := runtime.NewScheme()
//		_ = clientgoscheme.AddToScheme(scheme)
//		_ = batchv20170901.AddToScheme(scheme)
//		_ = documentdbv20150408.AddToScheme(scheme)
//		_ = storagev20190401.AddToScheme(scheme)
//		return scheme
//	}
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

	importedPackages := r.getImportedPackages()
	importedPackageNames := make([]string, 0, len(importedPackages))
	for pkg := range importedPackages {
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

	groupVersionAssignments := make([]dst.Stmt, 0, len(importedPackageNames))
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

	body := astbuilder.Statements(initSchemeVar, clientGoSchemeAssign, groupVersionAssignments, returnStmt)

	f := &astbuilder.FuncDetails{
		Name: "createScheme",
		Body: body,
	}

	f.AddReturn(astbuilder.Dereference(astbuilder.Selector(dst.NewIdent(runtime), "Scheme")))
	f.AddComments("creates a Scheme containing the clientgo types and all of the custom types returned by getKnownTypes")

	return f.DefineFunc(), nil
}

func (r *ResourceRegistrationFile) defineIndexFunctions(codeGenerationContext *astmodel.CodeGenerationContext) []dst.Decl {
	var indexFunctions []*functions.IndexRegistrationFunction
	for _, funcs := range r.indexFunctions {
		indexFunctions = append(indexFunctions, funcs...)
	}
	sort.Slice(indexFunctions, orderByFunctionName(indexFunctions))

	result := make([]dst.Decl, 0, len(indexFunctions))
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

func (r *ResourceRegistrationFile) makeWatchesExpr(typeName astmodel.TypeName, codeGenerationContext *astmodel.CodeGenerationContext) dst.Expr {
	secretWatchesExpr := r.makeSimpleWatchesExpr(
		typeName,
		astmodel.SecretType,
		"watchSecretsFactory",
		r.secretPropertyKeys,
		codeGenerationContext)
	configMapWatchesExpr := r.makeSimpleWatchesExpr(
		typeName,
		astmodel.ConfigMapType,
		"watchConfigMapsFactory",
		r.configMapPropertyKeys,
		codeGenerationContext)

	if secretWatchesExpr == nil && configMapWatchesExpr == nil {
		return nil
	}

	sliceBuilder := astbuilder.NewSliceLiteralBuilder(astmodel.WatchRegistrationType.AsType(codeGenerationContext), true)
	if secretWatchesExpr != nil {
		sliceBuilder.AddElement(secretWatchesExpr)
	}
	if configMapWatchesExpr != nil {
		sliceBuilder.AddElement(configMapWatchesExpr)
	}
	return sliceBuilder.Build()
}

// makeSimpleWatchesExpr generates code for a Watches expression:
//
//	{
//		Src:              &source.Kind{Type: &<fieldType>{}},
//		MakeEventHandler: <watchHelperFuncName>([]string{<typeNameKeys[typeName]>}, &<typeName>{}),
//	}
func (r *ResourceRegistrationFile) makeSimpleWatchesExpr(
	typeName astmodel.TypeName,
	fieldType astmodel.TypeName,
	watchHelperFuncName string,
	typeNameKeys map[astmodel.TypeName][]string,
	codeGenerationContext *astmodel.CodeGenerationContext) dst.Expr {
	keys, ok := typeNameKeys[typeName]
	if !ok {
		return nil
	}

	if len(keys) == 0 {
		return nil
	}

	newWatchBuilder := astbuilder.NewCompositeLiteralBuilder(astmodel.WatchRegistrationType.AsType(codeGenerationContext))
	sort.Strings(keys)

	builder := astbuilder.NewCompositeLiteralBuilder(
		astmodel.ControllerRuntimeSourceKindType.AsType(codeGenerationContext)).WithoutNewLines()
	builder.AddField(
		"Type",
		astbuilder.AddrOf(
			&dst.CompositeLit{
				Type: fieldType.AsType(codeGenerationContext),
			}))
	source := builder.Build()
	newWatchBuilder.AddField("Src", astbuilder.AddrOf(source))

	keyParams := make([]dst.Expr, 0, len(keys))
	for _, key := range keys {
		keyParams = append(keyParams, astbuilder.StringLiteral(key))
	}

	listTypeName := typeName.WithName(typeName.Name() + "List").AsType(codeGenerationContext)

	eventHandler := astbuilder.CallFunc(
		watchHelperFuncName,
		astbuilder.SliceLiteral(dst.NewIdent("string"), keyParams...),
		astbuilder.AddrOf(astbuilder.NewCompositeLiteralBuilder(listTypeName).Build()))
	newWatchBuilder.AddField("MakeEventHandler", eventHandler)

	return newWatchBuilder.Build()
}

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"go/token"
	"sort"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"
	"golang.org/x/exp/maps"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// ResourceRegistrationFile is a file containing functions that assist in registering resources
// with a Kubernetes scheme.
type ResourceRegistrationFile struct {
	resources               []astmodel.InternalTypeName
	storageVersionResources []astmodel.InternalTypeName
	resourceExtensions      []astmodel.InternalTypeName
	indexFunctions          map[astmodel.InternalTypeName][]*functions.IndexRegistrationFunction
	secretPropertyKeys      map[astmodel.InternalTypeName][]string
	configMapPropertyKeys   map[astmodel.InternalTypeName][]string
	validatingWebhooks      map[astmodel.InternalTypeName]astmodel.InternalTypeName
	mutatingWebhooks        map[astmodel.InternalTypeName]astmodel.InternalTypeName

	storageVersionToVersionMap map[astmodel.InternalTypeName][]astmodel.InternalTypeName
}

var _ astmodel.GoSourceFile = &ResourceRegistrationFile{}

// NewResourceRegistrationFile returns a ResourceRegistrationFile for registering the specified resources
// with a controller
func NewResourceRegistrationFile(
	resources []astmodel.InternalTypeName,
	storageVersionResources []astmodel.InternalTypeName,
	indexFunctions map[astmodel.InternalTypeName][]*functions.IndexRegistrationFunction,
	secretPropertyKeys map[astmodel.InternalTypeName][]string,
	configMapPropertyKeys map[astmodel.InternalTypeName][]string,
	resourceExtensions []astmodel.InternalTypeName,
	validatingWebhooks map[astmodel.InternalTypeName]astmodel.InternalTypeName,
	mutatingWebhooks map[astmodel.InternalTypeName]astmodel.InternalTypeName,
) *ResourceRegistrationFile {
	return &ResourceRegistrationFile{
		resources:               resources,
		storageVersionResources: storageVersionResources,
		indexFunctions:          indexFunctions,
		secretPropertyKeys:      secretPropertyKeys,
		configMapPropertyKeys:   configMapPropertyKeys,
		resourceExtensions:      resourceExtensions,
		validatingWebhooks:      validatingWebhooks,
		mutatingWebhooks:        mutatingWebhooks,
	}
}

// AsAst returns this file as a Go AST.
func (r *ResourceRegistrationFile) AsAst() (*dst.File, error) {
	var decls []dst.Decl

	// Determine imports
	packageReferences := r.generateImports()

	codeGenContext := astmodel.NewCodeGenerationContext(
		// This is a little nasty, given we're generating into a package with a fixed name.
		// Do we need a specific PackageReference type for this case?
		astmodel.MakeLocalPackageReference("", "controllers", "", ""), // TODO: This should come from a config
		packageReferences,
		nil)

	r.storageVersionToVersionMap = r.getStorageVersionToVersionsMap(codeGenContext)

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
	knownStorageTypes, err := r.createGetKnownStorageTypesFunc(codeGenContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating %s function", "getKnownStorageTypes")
	}

	decls = append(decls, knownStorageTypes)

	// getKnownTypes() function
	knownTypes, err := r.createGetKnownTypesFunc(codeGenContext, r.resources)
	if err != nil {
		return nil, err
	}
	decls = append(decls, knownTypes)

	// createScheme() function
	createSchemeFunc, err := r.createCreateSchemeFunc(codeGenContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating getKnownResourceExtensions function")
	}

	decls = append(decls, createSchemeFunc)

	// Create Resource Extensions
	resourceExtensionTypes, err := r.createGetResourceExtensions(codeGenContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating getKnownResourceExtensions function")
	}

	decls = append(decls, resourceExtensionTypes)

	// All the index functions
	indexFunctionDecls, err := r.defineIndexFunctions(codeGenContext)
	if err != nil {
		return nil, eris.Wrap(err, "defining index functions")
	}

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

func (r *ResourceRegistrationFile) getStorageVersionToVersionsMap(codeGenerationContext *astmodel.CodeGenerationContext) map[astmodel.InternalTypeName][]astmodel.InternalTypeName {
	result := make(map[astmodel.InternalTypeName][]astmodel.InternalTypeName, len(r.storageVersionResources))

	for _, storageVersion := range r.storageVersionResources {
		result[storageVersion] = []astmodel.InternalTypeName{}

		group := storageVersion.InternalPackageReference().Group()

		for _, version := range r.resources {
			// Dont include storage versions here
			if astmodel.IsStoragePackageReference(version.InternalPackageReference()) {
				continue
			}

			ref := version.InternalPackageReference()
			// Only match types whose group is the same and whose name is the same
			if ref.Group() != group || storageVersion.Name() != version.Name() {
				continue
			}

			result[storageVersion] = append(result[storageVersion], version)
		}

		// Sort the versions for a deterministic file layout
		sort.Slice(result[storageVersion], orderByImportedTypeName(codeGenerationContext, result[storageVersion]))
	}

	return result
}

// generateImports generates the PackageImportSet containing the imports required for the resources
// in the ResourceRegistrationFile.
func (r *ResourceRegistrationFile) generateImports() *astmodel.PackageImportSet {
	requiredImports := astmodel.NewPackageImportSet()
	typeSet := append(r.resources, r.storageVersionResources...)
	typeSet = append(typeSet, r.resourceExtensions...)
	typeSet = append(typeSet, maps.Values(r.mutatingWebhooks)...)
	typeSet = append(typeSet, maps.Values(r.validatingWebhooks)...)
	for _, typeName := range typeSet {
		requiredImports.AddImportOfReference(typeName.PackageReference())
	}

	// We require these imports
	requiredImports.AddImport(astmodel.NewPackageImport(astmodel.GenRuntimeReference))
	requiredImports.AddImport(astmodel.NewPackageImport(astmodel.ClientGoSchemeReference).WithName("clientgoscheme"))
	requiredImports.AddImport(astmodel.NewPackageImport(astmodel.APIMachineryRuntimeReference))
	requiredImports.AddImport(astmodel.NewPackageImport(astmodel.ControllerRuntimeClient))
	requiredImports.AddImport(astmodel.NewPackageImport(astmodel.GenRuntimeRegistrationReference))
	requiredImports.AddImport(astmodel.NewPackageImport(astmodel.CoreV1Reference))

	return requiredImports
}

func orderByImportedTypeName(
	codeGenerationContext *astmodel.CodeGenerationContext,
	resources []astmodel.InternalTypeName,
) func(i, j int) bool {
	return func(i, j int) bool {
		iVal := resources[i]
		jVal := resources[j]

		iPkgName, err := codeGenerationContext.GetImportedPackageName(iVal.PackageReference())
		if err != nil {
			panic(err)
		}
		jPkgName, err := codeGenerationContext.GetImportedPackageName(jVal.PackageReference())
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

// createGetKnownTypesFunc creates a getKnownTypes function that returns all known types.
// The registration.KnownType object also contains details about that objects (== GVK) webooks.
//
//	func getKnownTypes() []&registration.KnownType {
//		var result []*registration.KnownType
//		result = append(result, &registration.KnownType{Obj: new(<package>.<resource>)})
//		result = append(result, &registration.KnownType{Obj: new(<package>.<resource>)})
//		result = append(result, &registration.KnownType{Obj: new(<package>.<resource>)})
//		...
//		return result
//	}
func (r *ResourceRegistrationFile) createGetKnownTypesFunc(
	codeGenerationContext *astmodel.CodeGenerationContext,
	resources []astmodel.InternalTypeName,
) (dst.Decl, error) {
	funcName := "getKnownTypes"
	funcComment := "returns the list of all types and their webhooks"

	webhookRegistrationType, err := astmodel.KnownTypeRegistrationType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating type expression for WebhookRegistrationType")
	}

	resultIdent := dst.NewIdent("result")
	resultType := &dst.ArrayType{
		Elt: astbuilder.PointerTo(webhookRegistrationType),
	}
	resultVar := astbuilder.LocalVariableDeclaration(resultIdent.String(), resultType, "")

	// Sort the resources for a deterministic file layout
	sort.Slice(resources, orderByImportedTypeName(codeGenerationContext, resources))

	resourceAppendStatements := make([]dst.Stmt, 0, len(resources))
	batch := make([]dst.Expr, 0, 10)
	var lastPkg astmodel.PackageReference
	for _, typeName := range resources {
		if len(batch) > 0 && typeName.PackageReference() != lastPkg {
			appendStmt := astbuilder.AppendItemsToSlice(resultIdent, batch...)
			resourceAppendStatements = append(resourceAppendStatements, appendStmt)
			batch = batch[:0]
		}

		typeNameExpr, err := typeName.AsTypeExpr(codeGenerationContext)
		if err != nil {
			return nil, eris.Wrapf(err, "creating type expression for %s", typeName.Name())
		}

		litBuilder := astbuilder.NewCompositeLiteralBuilder(webhookRegistrationType)
		litBuilder.AddField("Obj", astbuilder.CallFunc("new", typeNameExpr))

		mutatingWebhook, hasMutatingWebhook := r.mutatingWebhooks[typeName]
		validatingWebhook, hasValidatingWebhook := r.validatingWebhooks[typeName]

		var mutatingWebhookExpr dst.Expr
		var validatingWebhookExpr dst.Expr

		if hasMutatingWebhook {
			mutatingWebhookExpr, err = mutatingWebhook.AsTypeExpr(codeGenerationContext)
			if err != nil {
				return nil, eris.Wrapf(err, "creating mutating webhook expression for %s", mutatingWebhook)
			}

			litBuilder.AddField("Defaulter", astbuilder.AddrOf(&dst.CompositeLit{Type: mutatingWebhookExpr}))
		}
		if hasValidatingWebhook {
			validatingWebhookExpr, err = validatingWebhook.AsTypeExpr(codeGenerationContext)
			if err != nil {
				return nil, eris.Wrapf(err, "creating validating webhook expression for %s", validatingWebhook)
			}

			litBuilder.AddField("Validator", astbuilder.AddrOf(&dst.CompositeLit{Type: validatingWebhookExpr}))
		}

		batch = append(batch, astbuilder.AddrOf(litBuilder.Build()))
		lastPkg = typeName.PackageReference()
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
//					Type: &corev1.Secret{},
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
) (dst.Decl, error) {
	funcName := "getKnownStorageTypes"
	funcComment := "returns the list of storage types which can be reconciled."

	resultIdent := dst.NewIdent("result")
	resultType := astmodel.NewArrayType(
		astmodel.NewOptionalType(
			astmodel.StorageTypeRegistrationType))
	resultTypeExpr, err := resultType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating type expression for StorageTypeRegistrationType")
	}

	resultVar := astbuilder.LocalVariableDeclaration(
		resultIdent.String(),
		resultTypeExpr,
		"")

	sort.Slice(r.storageVersionResources, orderByImportedTypeName(codeGenerationContext, r.storageVersionResources))

	resourceAppendStatements := make([]dst.Stmt, 0, len(r.storageVersionResources))
	for _, typeName := range r.storageVersionResources {
		typeNameExpr, err := typeName.AsTypeExpr(codeGenerationContext)
		if err != nil {
			return nil, eris.Wrapf(err, "creating type expression for %s", typeName.Name())
		}

		registrationTypeExpr, err := astmodel.StorageTypeRegistrationType.AsTypeExpr(codeGenerationContext)
		if err != nil {
			return nil, eris.Wrap(err, "creating type expression for StorageTypeRegistrationType")
		}

		newStorageTypeBuilder := astbuilder.NewCompositeLiteralBuilder(registrationTypeExpr)
		newStorageTypeBuilder.AddField("Obj", astbuilder.CallFunc("new", typeNameExpr))

		// Register index functions (if needed):
		// Indexes: []registration.Index{
		//		{
		//			Key: <key>,
		//			Func: <func>,
		//		}
		//	}
		if indexFuncs, ok := r.indexFunctions[typeName]; ok {
			var indexRegistrationTypeExpr dst.Expr
			indexRegistrationTypeExpr, err = astmodel.IndexRegistrationType.AsTypeExpr(codeGenerationContext)
			if err != nil {
				return nil, eris.Wrap(err, "creating type expression for IndexRegistrationType")
			}

			sliceBuilder := astbuilder.NewSliceLiteralBuilder(indexRegistrationTypeExpr, true)
			sort.Slice(indexFuncs, orderByFunctionName(indexFuncs))
			if len(indexFuncs) > 0 {
				for _, indexFunc := range indexFuncs {
					newIndexFunctionBuilder := astbuilder.NewCompositeLiteralBuilder(indexRegistrationTypeExpr)
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
		//			Type: <field type>,
		//			MakeEventHandler: <func>,
		//		}
		//	}
		watches, err := r.makeWatchesExpr(typeName, codeGenerationContext)
		if err != nil {
			return nil, eris.Wrapf(err, "creating watches expression for %s", typeName.Name())
		}
		if watches != nil {
			newStorageTypeBuilder.AddField("Watches", watches)
		}

		appendStmt := astbuilder.AppendItemToSlice(
			resultIdent,
			astbuilder.AddrOf(newStorageTypeBuilder.Build()))
		resourceAppendStatements = append(resourceAppendStatements, appendStmt)
	}

	returnStmt := astbuilder.Returns(resultIdent)

	body := astbuilder.Statements(resultVar, resourceAppendStatements, returnStmt)

	f := &astbuilder.FuncDetails{
		Name: funcName,
		Body: body,
	}

	f.AddReturn(
		resultTypeExpr)
	f.AddComments(funcComment)

	return f.DefineFunc(), nil
}

func (r *ResourceRegistrationFile) createGetResourceExtensions(
	context *astmodel.CodeGenerationContext,
) (dst.Decl, error) {
	funcName := "getResourceExtensions"
	funcComment := "returns a list of resource extensions"

	sort.Slice(r.resourceExtensions, orderByImportedTypeName(context, r.resourceExtensions))
	resultIdent := dst.NewIdent("result")
	resultType := astmodel.NewArrayType(astmodel.ResourceExtensionType)
	resultTypeExpr, err := resultType.AsTypeExpr(context)
	if err != nil {
		return nil, eris.Wrap(err, "creating type expression for ResourceExtensionType")
	}

	resultVar := astbuilder.LocalVariableDeclaration(
		resultIdent.String(),
		resultTypeExpr,
		"")

	resourceAppendStatements := make([]dst.Stmt, 0, len(r.resourceExtensions))
	for _, typeName := range r.resourceExtensions {
		typeNameExpr, err := typeName.AsTypeExpr(context)
		if err != nil {
			return nil, eris.Wrapf(err, "creating type expression for %s", typeName.Name())
		}

		literalExpr := astbuilder.AddrOf(astbuilder.NewCompositeLiteralBuilder(typeNameExpr).Build())
		appendStmt := astbuilder.AppendItemToSlice(
			resultIdent,
			literalExpr,
		)
		resourceAppendStatements = append(resourceAppendStatements, appendStmt)
	}

	returnStmt := astbuilder.Returns(resultIdent)

	body := astbuilder.Statements(resultVar, resourceAppendStatements, returnStmt)

	f := &astbuilder.FuncDetails{
		Name: funcName,
		Body: body,
	}

	f.AddReturn(resultTypeExpr)
	f.AddComments(funcComment)

	return f.DefineFunc(), nil
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

	returnStmt := astbuilder.Returns(dst.NewIdent(scheme))

	body := astbuilder.Statements(initSchemeVar, clientGoSchemeAssign, groupVersionAssignments, returnStmt)

	f := &astbuilder.FuncDetails{
		Name: "createScheme",
		Body: body,
	}

	f.AddReturn(astbuilder.Dereference(astbuilder.Selector(dst.NewIdent(runtime), "Scheme")))
	f.AddComments("creates a Scheme containing the clientgo types and all of the custom types returned by getKnownTypes")

	return f.DefineFunc(), nil
}

func (r *ResourceRegistrationFile) defineIndexFunctions(
	codeGenerationContext *astmodel.CodeGenerationContext,
) ([]dst.Decl, error) {
	var indexFunctions []*functions.IndexRegistrationFunction
	for _, funcs := range r.indexFunctions {
		indexFunctions = append(indexFunctions, funcs...)
	}
	sort.Slice(indexFunctions, orderByFunctionName(indexFunctions))

	result := make([]dst.Decl, 0, len(indexFunctions))
	var errs []error
	for _, f := range indexFunctions {
		decl, err := f.AsFunc(codeGenerationContext, astmodel.InternalTypeName{})
		if err != nil {
			errs = append(errs, err)
			continue
		}

		result = append(result, decl)
	}

	return result, kerrors.NewAggregate(errs)
}

func (r *ResourceRegistrationFile) getImportedPackages() map[astmodel.PackageReference]struct{} {
	result := make(map[astmodel.PackageReference]struct{})
	for _, typeName := range r.resources {
		result[typeName.PackageReference()] = struct{}{}
	}

	return result
}

func (r *ResourceRegistrationFile) makeWatchesExpr(
	typeName astmodel.InternalTypeName,
	codeGenerationContext *astmodel.CodeGenerationContext,
) (dst.Expr, error) {
	secretWatchesExpr, err := r.makeSimpleWatchesExpr(
		typeName,
		astmodel.SecretType,
		"watchSecretsFactory",
		r.secretPropertyKeys,
		codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating watches expression for secrets")
	}

	configMapWatchesExpr, err := r.makeSimpleWatchesExpr(
		typeName,
		astmodel.ConfigMapType,
		"watchConfigMapsFactory",
		r.configMapPropertyKeys,
		codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating watches expression for configmaps")
	}

	if secretWatchesExpr == nil && configMapWatchesExpr == nil {
		return nil, nil
	}

	watchRegistrationTypeExpr, err := astmodel.WatchRegistrationType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating type expression for WatchRegistrationType")
	}

	sliceBuilder := astbuilder.NewSliceLiteralBuilder(watchRegistrationTypeExpr, true)
	if secretWatchesExpr != nil {
		sliceBuilder.AddElement(secretWatchesExpr)
	}
	if configMapWatchesExpr != nil {
		sliceBuilder.AddElement(configMapWatchesExpr)
	}

	return sliceBuilder.Build(), nil
}

// makeSimpleWatchesExpr generates code for a Watches expression:
//
//	{
//		Type:              &&<fieldType>{},
//		MakeEventHandler: <watchHelperFuncName>([]string{<typeNameKeys[typeName]>}, &<typeName>{}),
//	}
func (r *ResourceRegistrationFile) makeSimpleWatchesExpr(
	typeName astmodel.InternalTypeName,
	fieldType astmodel.TypeName,
	watchHelperFuncName string,
	typeNameKeys map[astmodel.InternalTypeName][]string,
	codeGenerationContext *astmodel.CodeGenerationContext,
) (dst.Expr, error) {
	keys, ok := typeNameKeys[typeName]
	if !ok {
		return nil, nil
	}

	if len(keys) == 0 {
		return nil, nil
	}

	watchRegistrationTypeExpr, err := astmodel.WatchRegistrationType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating type expression for WatchRegistrationType")
	}

	newWatchBuilder := astbuilder.NewCompositeLiteralBuilder(watchRegistrationTypeExpr)
	sort.Strings(keys)

	fieldTypeExpr, err := fieldType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", fieldType.Name())
	}

	objectType := astbuilder.AddrOf(&dst.CompositeLit{
		Type: fieldTypeExpr,
	})
	newWatchBuilder.AddField("Type", objectType)

	keyParams := make([]dst.Expr, 0, len(keys))
	for _, key := range keys {
		k := astbuilder.StringLiteral(key)

		k.Decorations().Before = dst.NewLine
		k.Decorations().After = dst.NewLine

		keyParams = append(keyParams, k)
	}

	listType := typeName.WithName(typeName.Name() + "List")
	listTypeExpr, err := listType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", listType.Name())
	}

	slice := astbuilder.SliceLiteral(dst.NewIdent("string"), keyParams...)
	slice.Decorations().Before = dst.NewLine
	slice.Decorations().After = dst.NewLine

	eventHandler := astbuilder.CallFunc(
		watchHelperFuncName,
		slice,
		astbuilder.AddrOf(astbuilder.NewCompositeLiteralBuilder(listTypeExpr).Build()))
	newWatchBuilder.AddField("MakeEventHandler", eventHandler)

	return newWatchBuilder.Build(), nil
}

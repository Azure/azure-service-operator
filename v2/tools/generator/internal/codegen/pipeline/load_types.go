/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"

	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/jsonast"
)

const LoadTypesStageID = "loadTypes"

/* LoadTypes creates a PipelineStage to load Swagger data.

This information is derived from the Azure Swagger specifications. We parse the Swagger specs and look for
any actions that appear to be ARM resources (have PUT methods with types we can use and appropriate names in the
action path). Then for each resource, we use the existing JSON AST parser to extract the status type
(the type-definition part of swagger is the same as JSON Schema). */
func LoadTypes(idFactory astmodel.IdentifierFactory, config *config.Configuration) *Stage {
	return NewLegacyStage(
		LoadTypesStageID,
		"Load all types from Swagger files",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			klog.V(1).Infof("Loading Swagger data from %q", config.SchemaRoot)

			swaggerTypes, err := loadSwaggerData(ctx, idFactory, config)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to load Swagger data")
			}

			klog.V(1).Infof("Loaded Swagger data (%d resources, %d other definitions)", len(swaggerTypes.ResourceDefinitions), len(swaggerTypes.OtherDefinitions))

			if len(swaggerTypes.ResourceDefinitions) == 0 || len(swaggerTypes.OtherDefinitions) == 0 {
				return nil, errors.Errorf("Failed to load swagger information")
			}

			resourceToStatus, statusTypes, err := generateStatusTypes(swaggerTypes)
			if err != nil {
				return nil, err
			}

			resourceToSpec, specTypes, err := generateSpecTypes(swaggerTypes)
			if err != nil {
				return nil, err
			}

			// put all definitions into a new set
			defs := make(astmodel.TypeDefinitionSet)
			defs.AddTypes(statusTypes)
			defs.AddTypes(specTypes)

			for resourceName, resourceInfo := range swaggerTypes.ResourceDefinitions {
				spec := resourceToSpec.MustGetDefinition(resourceName)
				status := resourceToStatus.MustGetDefinition(resourceName)

				specType := addRequiredSpecFields(spec.Type())
				statusType := status.Type()

				resourceType := astmodel.NewResourceType(specType, statusType)

				// add on ARM Type & URI
				resourceType = resourceType.WithARMType(resourceInfo.ARMType).WithARMURI(resourceInfo.ARMURI)

				scope := categorizeResourceScope(resourceInfo.ARMURI)
				resourceType = resourceType.WithScope(scope)
				resourceDefinition := astmodel.MakeTypeDefinition(resourceName, resourceType)

				// document origin of resource
				sourceFile := strings.TrimPrefix(resourceInfo.SourceFile, config.SchemaRoot)
				resourceDefinition = resourceDefinition.
					WithDescription(([]string{
						"Generator information:",
						fmt.Sprintf(" - Generated from: %s", filepath.ToSlash(sourceFile)),
						fmt.Sprintf(" - ARM URI: %s", resourceInfo.ARMURI),
					}))

				err := defs.AddAllowDuplicates(resourceDefinition)
				if err != nil {
					return nil, err
				}
			}

			klog.V(1).Infof("Input %d definitions, output %d definitions", len(definitions), len(defs))

			return defs, nil
		})
}

var resourceGroupScopeRegex = regexp.MustCompile(`(?i)^/subscriptions/[^/]+/resourcegroups/.*`)

func categorizeResourceScope(armURI string) astmodel.ResourceScope {
	// this is a bit of a hack, eventually we should have better scope support.
	// at the moment we assume that a resource is an extension if it can be applied to
	// any scope:
	if strings.HasPrefix(armURI, "/{scope}/") {
		return astmodel.ResourceScopeExtension
	}

	if resourceGroupScopeRegex.MatchString(armURI) {
		return astmodel.ResourceScopeResourceGroup
	}

	// TODO: Not currently possible to generate a resource with scope Location, we should fix that
	return astmodel.ResourceScopeTenant
}

var requiredSpecFields = astmodel.NewObjectType().WithProperties(
	astmodel.NewPropertyDefinition(astmodel.AzureNameProperty, "azureName", astmodel.StringType),
	astmodel.NewPropertyDefinition(astmodel.NameProperty, "name", astmodel.StringType))

func addRequiredSpecFields(t astmodel.Type) astmodel.Type {
	return astmodel.BuildAllOfType(t, requiredSpecFields)
}

type statusTypes struct {
	// resourceTypes maps Spec name to corresponding Status type
	// the typeName is lowercased to be case-insensitive
	resourceTypes resourceLookup

	// otherDefinitions has all other Status types renamed to avoid clashes with Spec types
	otherDefinitions astmodel.TypeDefinitionSet
}

func (st statusTypes) findResourceType(typeName astmodel.TypeName) (astmodel.Type, bool) {
	if statusDef, ok := st.resourceTypes.tryFind(typeName); ok {
		klog.V(4).Infof("Swagger information found for %s", typeName)
		return statusDef, true
	} else {
		klog.V(3).Infof("Swagger information missing for %s", typeName)
		// add a warning that the status is missing
		// this will be reported if the type is not pruned
		return astmodel.NewErroredType(nil, []string{fmt.Sprintf("missing status information for %s", typeName)}, nil), false
	}
}

type resourceLookup map[astmodel.TypeName]astmodel.Type

func asKey(name astmodel.TypeName) astmodel.TypeName {
	n := strings.ToLower(name.Name())
	n = strings.ReplaceAll(n, "_", "")

	return astmodel.MakeTypeName(name.PackageReference, n)
}

func (resourceLookup resourceLookup) tryFind(name astmodel.TypeName) (astmodel.Type, bool) {
	result, ok := resourceLookup[asKey(name)]
	return result, ok
}

func (resourceLookup resourceLookup) add(name astmodel.TypeName, theType astmodel.Type) {
	lower := asKey(name)
	if _, ok := resourceLookup[lower]; ok {
		panic(fmt.Sprintf("lowercase name collision: %s", name))
	}

	resourceLookup[lower] = theType
}

// statusTypeRenamer appends our standard StatusSuffix '_Status` to all types
var statusTypeRenamer = astmodel.NewRenamingVisitorFromLambda(appendStatusSuffix)

func appendStatusSuffix(typeName astmodel.TypeName) astmodel.TypeName {
	return astmodel.MakeTypeName(typeName.PackageReference, typeName.Name()+astmodel.StatusSuffix)
}

// generateStatusTypes returns the statusTypes for the input Swagger types
// all types (apart from Resources) are renamed to have "_STATUS" as a
// suffix, to avoid name clashes.
func generateStatusTypes(swaggerTypes jsonast.SwaggerTypes) (astmodel.TypeDefinitionSet, astmodel.TypeDefinitionSet, error) {
	resourceLookup, otherTypes, err := renamed(swaggerTypes, true, astmodel.StatusNameSuffix)
	if err != nil {
		return nil, nil, err
	}

	newResources := make(astmodel.TypeDefinitionSet)
	// often the top-level type in Swagger has a bad name like "CreateParametersX"
	// we'll try to substitute that with a better name here
	for resourceName, resourceDef := range resourceLookup {
		statusTypeName := resourceDef.Type().(astmodel.TypeName) // always a TypeName, see 'renamed' comment
		desiredStatusName := resourceName.WithName(resourceName.Name() + astmodel.StatusNameSuffix)

		if statusTypeName == desiredStatusName {
			newResources.Add(resourceDef)
			continue // nothing to do
		}

		targetType, err := otherTypes.FullyResolve(statusTypeName)
		if err != nil {
			panic("didn't find type in set after renaming; shouldn't be possible")
		}

		err = otherTypes.AddAllowDuplicates(astmodel.MakeTypeDefinition(desiredStatusName, targetType))
		if err != nil {
			// unable to use desiredName
			newResources.Add(resourceDef)
		} else {
			newResources.Add(astmodel.MakeTypeDefinition(resourceName, desiredStatusName))
		}
	}

	return newResources, otherTypes, nil
}

// Note that the first result is for mapping resource names → types, so it is always TypeName→TypeName.
// The second contains all the renamed types.
func renamed(swaggerTypes jsonast.SwaggerTypes, status bool, suffix string) (astmodel.TypeDefinitionSet, astmodel.TypeDefinitionSet, error) {
	renamer := astmodel.NewRenamingVisitorFromLambda(func(typeName astmodel.TypeName) astmodel.TypeName {
		return typeName.WithName(typeName.Name() + suffix)
	})

	var errs []error
	otherTypes := make(astmodel.TypeDefinitionSet)
	for _, typeDef := range swaggerTypes.OtherDefinitions {
		renamedDef, err := renamer.RenameDefinition(typeDef)
		if err != nil {
			errs = append(errs, err)
		} else {
			otherTypes.Add(renamedDef)
		}
	}

	resources := make(astmodel.TypeDefinitionSet)
	for resourceName, resourceDef := range swaggerTypes.ResourceDefinitions {
		// resourceName is not renamed as this is a lookup for the resource type
		typeToRename := resourceDef.SpecType
		if status {
			typeToRename = resourceDef.StatusType
		}
		renamedType, err := renamer.Rename(typeToRename)
		if err != nil {
			errs = append(errs, err)
		} else {
			resources.Add(astmodel.MakeTypeDefinition(resourceName, renamedType))
		}
	}

	err := kerrors.NewAggregate(errs)
	if err != nil {
		return nil, nil, err
	}

	return resources, otherTypes, nil
}

func generateSpecTypes(swaggerTypes jsonast.SwaggerTypes) (astmodel.TypeDefinitionSet, astmodel.TypeDefinitionSet, error) {
	// TODO: I think this should be renamed to "_Spec" for consistency, but will do in a separate PR for cleanliness #Naming
	// the alternative is that we place them in their own package
	resources, otherTypes, err := renamed(swaggerTypes, false, "")
	if err != nil {
		return nil, nil, err
	}

	// Fix-up: rename top-level resource types so that a Resource
	// always points to a _Spec type.
	// TODO: remove once preceding TODO is resolved and everything is consistently named _Spec #Naming
	{
		renames := make(map[astmodel.TypeName]astmodel.TypeName)
		for typeName := range otherTypes {
			if _, ok := resources[typeName]; ok {
				// would be a clash with resource name
				renames[typeName] = typeName.WithName(typeName.Name() + "Spec")
			}
		}

		renamer := astmodel.NewRenamingVisitor(renames)
		newOtherTypes, renameErr := renamer.RenameAll(otherTypes)
		if renameErr != nil {
			panic(renameErr)
		}

		otherTypes = newOtherTypes

		// for resources we must only rename the Type not the Name,
		// since this is used as a lookup:
		newResources := make(astmodel.TypeDefinitionSet)
		for rName, rType := range resources {
			newType, renameErr := renamer.Rename(rType.Type())
			if renameErr != nil {
				panic(renameErr)
			}
			newResources.Add(astmodel.MakeTypeDefinition(rName, newType))
		}

		rewriter := astmodel.TypeVisitorBuilder{
			VisitObjectType: func(this *astmodel.TypeVisitor, it *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
				// strip all readonly props
				var propsToRemove []astmodel.PropertyName
				it.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
					if prop.ReadOnly() {
						propsToRemove = append(propsToRemove, prop.PropertyName())
					}
				})

				it = it.WithoutSpecificProperties(propsToRemove...)

				return astmodel.IdentityVisitOfObjectType(this, it, ctx)
			},
		}.Build()

		otherTypes, err = rewriter.VisitDefinitions(otherTypes, nil)
		if err != nil {
			return nil, nil, err
		}

		resources = newResources
	}

	return resources, otherTypes, nil
}

func loadSwaggerData(ctx context.Context, idFactory astmodel.IdentifierFactory, config *config.Configuration) (jsonast.SwaggerTypes, error) {
	schemas, err := loadAllSchemas(ctx, config.SchemaRoot, config.LocalPathPrefix(), idFactory, config.Status.Overrides)
	if err != nil {
		return jsonast.SwaggerTypes{}, err
	}

	loader := jsonast.NewCachingFileLoader(schemas)

	typesByGroup := make(map[astmodel.LocalPackageReference][]typesFromFile)
	for schemaPath, schema := range schemas {

		extractor := jsonast.NewSwaggerTypeExtractor(
			config,
			idFactory,
			schema.Swagger,
			schemaPath,
			*schema.Package, // always set during generation
			loader)

		types, err := extractor.ExtractTypes(ctx)
		if err != nil {
			return jsonast.SwaggerTypes{}, errors.Wrapf(err, "error processing %q", schemaPath)
		}

		typesByGroup[*schema.Package] = append(typesByGroup[*schema.Package], typesFromFile{types, schemaPath})
	}

	return mergeSwaggerTypesByGroup(idFactory, typesByGroup)
}

func mergeSwaggerTypesByGroup(idFactory astmodel.IdentifierFactory, m map[astmodel.LocalPackageReference][]typesFromFile) (jsonast.SwaggerTypes, error) {
	klog.V(3).Infof("Merging types for %d groups/versions", len(m))

	result := jsonast.SwaggerTypes{
		ResourceDefinitions: make(jsonast.ResourceDefinitionSet),
		OtherDefinitions:    make(astmodel.TypeDefinitionSet),
	}

	for pkg, group := range m {
		klog.V(3).Infof("Merging types for %s", pkg)
		merged := mergeTypesForPackage(idFactory, group)
		for rn, rt := range merged.ResourceDefinitions {
			if _, ok := result.ResourceDefinitions[rn]; ok {
				panic("duplicate resource generated")
			}

			result.ResourceDefinitions[rn] = rt
		}

		err := result.OtherDefinitions.AddTypesAllowDuplicates(merged.OtherDefinitions)
		if err != nil {
			return result, errors.Wrapf(err, "when combining swagger types for %s", pkg)
		}
	}

	return result, nil
}

type typesFromFile struct {
	jsonast.SwaggerTypes
	filePath string
}

type typesFromFilesSorter struct{ x []typesFromFile }

var _ sort.Interface = typesFromFilesSorter{}

func (s typesFromFilesSorter) Len() int           { return len(s.x) }
func (s typesFromFilesSorter) Swap(i, j int)      { s.x[i], s.x[j] = s.x[j], s.x[i] }
func (s typesFromFilesSorter) Less(i, j int) bool { return s.x[i].filePath < s.x[j].filePath }

// mergeTypesForPackage merges the types for a single package from multiple files
func mergeTypesForPackage(idFactory astmodel.IdentifierFactory, typesFromFiles []typesFromFile) jsonast.SwaggerTypes {
	sort.Sort(typesFromFilesSorter{typesFromFiles})

	typeNameCounts := make(map[astmodel.TypeName]int)
	for _, typesFromFile := range typesFromFiles {
		for name := range typesFromFile.OtherDefinitions {
			typeNameCounts[name] += 1
		}
	}

	// a set of renamings, one per file
	renames := make([]map[astmodel.TypeName]astmodel.TypeName, len(typesFromFiles))
	for ix := range typesFromFiles {
		renames[ix] = make(map[astmodel.TypeName]astmodel.TypeName)
	}

	for name, count := range typeNameCounts {
		colliding := findCollidingTypeNames(typesFromFiles, name, count)
		if colliding == nil {
			continue
		}

		names := make([]string, len(colliding))
		for ix, ttc := range colliding {
			newName := generateRenaming(idFactory, name, typesFromFiles[ttc.typesFromFileIx].filePath, typeNameCounts)
			names[ix] = newName.Name()
			renames[ttc.typesFromFileIx][name] = newName
		}

		klog.V(3).Infof("Conflicting definitions for %s, renaming to: %s", name, strings.Join(names, ", "))
	}

	for ix := range typesFromFiles {
		renamesForFile := renames[ix]
		if len(renamesForFile) > 0 {
			typesFromFiles[ix] = applyRenames(renamesForFile, typesFromFiles[ix])
		}
	}

	mergedResult := jsonast.SwaggerTypes{
		ResourceDefinitions: make(jsonast.ResourceDefinitionSet),
		OtherDefinitions:    make(astmodel.TypeDefinitionSet),
	}

	for _, typesFromFile := range typesFromFiles {
		for _, t := range typesFromFile.OtherDefinitions {
			// for consistent results we always sort typesFromFiles first (at top of this function)
			// so that we always pick the same one when there are multiple
			_ = mergedResult.OtherDefinitions.AddAllowDuplicates(t)
			// errors ignored since we already checked for structural equality
			// it’s possible for types to refer to different typenames in which case they are not TypeEquals Equal
			// but they might be structurally equal
		}

		for rn, rt := range typesFromFile.ResourceDefinitions {
			if foundRT, ok := mergedResult.ResourceDefinitions[rn]; ok &&
				!(astmodel.TypeEquals(foundRT.SpecType, rt.SpecType) && astmodel.TypeEquals(foundRT.StatusType, rt.StatusType)) {
				panic(fmt.Sprintf("While merging file %s: duplicate resource types generated", typesFromFile.filePath))
			}

			mergedResult.ResourceDefinitions[rn] = rt
		}
	}

	return mergedResult
}

type typeAndSource struct {
	def             astmodel.TypeDefinition
	typesFromFileIx int
}

// findCollidingTypeNames finds any types with the given name that collide, and returns
// the definition as well as the index of the file it was found in
func findCollidingTypeNames(typesFromFiles []typesFromFile, name astmodel.TypeName, duplicateCount int) []typeAndSource {
	if duplicateCount == 1 {
		// cannot collide
		return nil
	}

	typesToCheck := make([]typeAndSource, 0, duplicateCount)
	for typesIx, types := range typesFromFiles {
		if def, ok := types.OtherDefinitions[name]; ok {
			typesToCheck = append(typesToCheck, typeAndSource{def: def, typesFromFileIx: typesIx})
			// short-circuit
			if len(typesToCheck) == duplicateCount {
				break
			}
		}
	}

	first := typesToCheck[0]
	for _, other := range typesToCheck[1:] {
		if !structurallyIdentical(
			first.def.Type(),
			typesFromFiles[first.typesFromFileIx].OtherDefinitions,
			other.def.Type(),
			typesFromFiles[other.typesFromFileIx].OtherDefinitions,
		) {
			if name != first.def.Name() || name != other.def.Name() {
				panic("assert")
			}

			_, firstOk := first.def.Type().(*astmodel.ResourceType)
			_, otherOk := other.def.Type().(*astmodel.ResourceType)
			if firstOk || otherOk {
				panic("cannot rename resources")
			}

			return typesToCheck
		}
		// Else: they are structurally identical and it is okay to pick one.
		// Note that when types are structurally identical they aren’t necessarily type.Equals equal;
		// for this reason we must ignore errors when adding to the overall result set using AddAllowDuplicates.
	}

	return nil
}

// generateRenaming finds a new name for a type based upon the file it is in
func generateRenaming(
	idFactory astmodel.IdentifierFactory,
	original astmodel.TypeName,
	filePath string,
	typeNames map[astmodel.TypeName]int,
) astmodel.TypeName {
	name := filepath.Base(filePath)
	name = strings.TrimSuffix(name, filepath.Ext(name))

	// Prefix the typename with the filename
	result := astmodel.MakeTypeName(
		original.PackageReference,
		idFactory.CreateIdentifier(name+original.Name(), astmodel.Exported))

	// see if there are any collisions: add Xs until there are no collisions
	// TODO: this might result in non-determinism depending on iteration order
	// in the calling method
	for _, ok := typeNames[result]; ok; _, ok = typeNames[result] {
		result = astmodel.MakeTypeName(
			result.PackageReference,
			result.Name()+"X",
		)
	}

	return result
}

func applyRenames(renames map[astmodel.TypeName]astmodel.TypeName, typesFromFile typesFromFile) typesFromFile {
	visitor := astmodel.NewRenamingVisitor(renames)

	// visit all other types
	newOtherTypes, err := visitor.RenameAll(typesFromFile.OtherDefinitions)
	if err != nil {
		panic(err)
	}

	// visit all resource types
	newResourceTypes := make(jsonast.ResourceDefinitionSet)
	for rn, rt := range typesFromFile.ResourceDefinitions {
		newSpecType, err := visitor.Rename(rt.SpecType)
		if err != nil {
			panic(err)
		}

		newStatusType, err := visitor.Rename(rt.StatusType)
		if err != nil {
			panic(err)
		}

		newResourceTypes[rn] = jsonast.ResourceDefinition{
			SpecType:   newSpecType,
			StatusType: newStatusType,
			SourceFile: rt.SourceFile,
		}
	}

	typesFromFile.OtherDefinitions = newOtherTypes
	typesFromFile.ResourceDefinitions = newResourceTypes
	return typesFromFile
}

// structurallyIdentical checks if the two provided types are structurally identical
// all the way to their leaf nodes (recursing into TypeNames)
func structurallyIdentical(
	leftType astmodel.Type,
	leftDefinitions astmodel.TypeDefinitionSet,
	rightType astmodel.Type,
	rightDefinitions astmodel.TypeDefinitionSet,
) bool {
	// we cannot simply recurse when we hit TypeNames as there can be cycles in types.
	// instead we store all TypeNames that need to be checked in here, and
	// check them one at a time until there is nothing left to be checked:
	type pair struct{ left, right astmodel.TypeName }
	toCheck := []pair{}            // queue of pairs to check
	checked := map[pair]struct{}{} // set of pairs that have been enqueued

	override := astmodel.EqualityOverrides{}
	override.TypeName = func(left, right astmodel.TypeName) bool {
		// note that this relies on Equals implementations preserving the left/right order
		p := pair{left, right}
		if _, ok := checked[p]; !ok {
			checked[p] = struct{}{}
			toCheck = append(toCheck, p)
		}

		// conditionally true, as long as all pairs are equal
		return true
	}

	// check the provided types
	if !astmodel.TypeEquals(leftType, rightType, override) {
		return false
	}

	// check all TypeName pairs until there are none left to check
	for len(toCheck) > 0 {
		next := toCheck[0]
		toCheck = toCheck[1:]
		if !astmodel.TypeEquals(
			leftDefinitions[next.left].Type(),
			rightDefinitions[next.right].Type(),
			override) {
			return false
		}
	}

	// if we didn’t find anything that didn’t match then they are identical
	return true
}

// TODO: is there, perhaps, a way to detect these without hardcoding these paths?
var skipDirectories = []string{
	"/examples/",
	"/quickstart-templates/",
	"/control-plane/",
	"/data-plane/",
}

func shouldSkipDir(filePath string) bool {
	p := filepath.ToSlash(filePath)

	for _, skipDir := range skipDirectories {
		if strings.Contains(p, skipDir) {
			return true
		}
	}

	return false
}

// loadAllSchemas walks all .json files in the given rootPath in directories
// of the form "Microsoft.GroupName/…/2000-01-01/…" (excluding those matching
// shouldSkipDir), and returns those files in a map of path→swagger spec.
func loadAllSchemas(
	ctx context.Context,
	rootPath string,
	localPathPrefix string,
	idFactory astmodel.IdentifierFactory,
	overrides []config.SchemaOverride,
) (map[string]jsonast.PackageAndSwagger, error) {
	var mutex sync.Mutex
	schemas := make(map[string]jsonast.PackageAndSwagger)

	var eg errgroup.Group
	err := filepath.Walk(rootPath, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}

		if shouldSkipDir(filePath) {
			return filepath.SkipDir // this is a magic error
		}

		if !fileInfo.IsDir() &&
			filepath.Ext(filePath) == ".json" {

			group := groupFromPath(filePath, rootPath, overrides)
			version := versionFromPath(filePath, rootPath)
			if group == "" || version == "" {
				return nil
			}

			pkg := astmodel.MakeLocalPackageReference(
				localPathPrefix,
				idFactory.CreateGroupName(group),
				astmodel.GeneratorVersion,
				version)

			// all files are loaded in parallel to speed this up
			eg.Go(func() error {
				var swagger spec.Swagger

				klog.V(3).Infof("Loading file %q", filePath)

				fileContent, err := ioutil.ReadFile(filePath)
				if err != nil {
					return errors.Wrapf(err, "unable to read swagger file %q", filePath)
				}

				err = swagger.UnmarshalJSON(fileContent)
				if err != nil {
					return errors.Wrapf(err, "unable to parse swagger file %q", filePath)
				}

				mutex.Lock()
				schemas[filePath] = jsonast.PackageAndSwagger{Package: &pkg, Swagger: swagger}
				mutex.Unlock()

				return nil
			})
		}

		return nil
	})

	egErr := eg.Wait() // for files to finish loading

	if err != nil {
		return nil, err
	}

	if egErr != nil {
		return nil, egErr
	}

	return schemas, nil
}

func groupFromPath(filePath string, rootPath string, overrides []config.SchemaOverride) string {
	filePath = filepath.ToSlash(filePath)
	group := jsonast.SwaggerGroupRegex.FindString(filePath)

	// see if there is a config override for this file
	for _, schemaOverride := range overrides {
		configSchemaPath := filepath.ToSlash(path.Join(rootPath, schemaOverride.BasePath))
		if strings.HasPrefix(filePath, configSchemaPath) {
			// a forced namespace: use it
			if schemaOverride.Namespace != "" {
				klog.V(1).Infof("Overriding namespace to %s for file %s", schemaOverride.Namespace, filePath)
				return schemaOverride.Namespace
			}

			// found a suffix override: apply it
			if schemaOverride.Suffix != "" {
				group = group + "." + schemaOverride.Suffix
				klog.V(1).Infof("Overriding namespace to %s for file %s", group, filePath)
				return group
			}
		}
	}

	return group
}

// supports date-based versions or v1, v2 (as used by common types)
var swaggerVersionRegex = regexp.MustCompile(`/(\d{4}-\d{2}-\d{2}(-preview|-privatepreview)?)|(v\d+)|(\d+\.\d+)/`)

func versionFromPath(filePath string, rootPath string) string {
	// we want to ignore anything in the root path, since, e.g.
	// the specs can be nested inside a directory that matches the swaggerVersionRegex
	// (and indeed this is the case with the /v2/ package)
	filePath = strings.TrimPrefix(filePath, rootPath)
	// must trim leading & trailing '/' as golang does not support lookaround
	fp := filepath.ToSlash(filePath)
	return strings.Trim(swaggerVersionRegex.FindString(fp), "/")
}

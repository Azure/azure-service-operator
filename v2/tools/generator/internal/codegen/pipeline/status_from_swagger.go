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

/* addStatusFromSwagger creates a PipelineStage to add status information into the generated resources.

This information is derived from the Azure Swagger specifications. We parse the Swagger specs and look for
any actions that appear to be ARM resources (have PUT methods with types we can use and appropriate names in the
action path). Then for each resource, we use the existing JSON AST parser to extract the status type
(the type-definition part of swagger is the same as JSON Schema).

Next, we walk over all the resources we are currently generating CRDs for and attempt to locate
a match for the resource in the status information we have parsed. If we locate a match, it is
added to the Status field of the Resource type, after we have renamed all the status types to
avoid any conflicts with existing Spec types that have already been defined.

*/
func AddStatusFromSwagger(idFactory astmodel.IdentifierFactory, config *config.Configuration) *Stage {
	return NewLegacyStage(
		"addStatusFromSwagger",
		"Add information from Swagger specs for 'status' fields",
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			if config.Status.SchemaRoot == "" {
				klog.Warningf("No status schema root specified, will not generate status types")
				return definitions, nil
			}

			klog.V(1).Infof("Loading Swagger data from %q", config.Status.SchemaRoot)

			swaggerTypes, err := loadSwaggerData(ctx, idFactory, config)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to load Swagger data")
			}

			klog.V(1).Infof("Loaded Swagger data (%d resources, %d other definitions)", len(swaggerTypes.ResourceDefinitions), len(swaggerTypes.OtherDefinitions))

			if len(swaggerTypes.ResourceDefinitions) == 0 || len(swaggerTypes.OtherDefinitions) == 0 {
				return nil, errors.Errorf("Failed to load swagger information")
			}

			statusTypes, err := generateStatusTypes(swaggerTypes)
			if err != nil {
				return nil, err
			}

			// put all definitions into a new set
			defs := make(astmodel.TypeDefinitionSet)

			// all non-resources from Swagger are added regardless of whether they are used
			// if they are not used they will be pruned off by a later pipeline stage
			// (there will be no name clashes here due to suffixing with "_Status")
			defs.AddTypes(statusTypes.otherDefinitions)

			matchedResources := 0
			// find any resources and update them with status info
			for typeName, typeDef := range definitions {
				if resource, ok := typeDef.Type().(*astmodel.ResourceType); ok {
					// find the status type (= Swagger resource type)
					newStatus, located := statusTypes.findResourceType(typeName)
					if located {
						matchedResources++
					}

					defs.Add(typeDef.WithType(resource.WithStatus(newStatus)))
				} else {
					defs.Add(typeDef)
				}
			}

			klog.V(1).Infof("Found status information for %d resources", matchedResources)
			klog.V(1).Infof("Input %d definitions, output %d definitions", len(definitions), len(defs))

			return defs, nil
		})
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

func lowerCase(name astmodel.TypeName) astmodel.TypeName {
	return astmodel.MakeTypeName(name.PackageReference, strings.ToLower(name.Name()))
}

func (resourceLookup resourceLookup) tryFind(name astmodel.TypeName) (astmodel.Type, bool) {
	result, ok := resourceLookup[lowerCase(name)]
	return result, ok
}

func (resourceLookup resourceLookup) add(name astmodel.TypeName, theType astmodel.Type) {
	lower := lowerCase(name)
	if _, ok := resourceLookup[lower]; ok {
		panic(fmt.Sprintf("lowercase name collision: %s", name))
	}

	resourceLookup[lower] = theType
}

// statusTypeRenamer appends "_Status" to all types
var statusTypeRenamer = astmodel.NewRenamingVisitorFromLambda(appendStatusSuffix)

func appendStatusSuffix(typeName astmodel.TypeName) astmodel.TypeName {
	return astmodel.MakeTypeName(typeName.PackageReference, typeName.Name()+"_Status")
}

// generateStatusTypes returns the statusTypes for the input Swagger types
// all types (apart from Resources) are renamed to have "_Status" as a
// suffix, to avoid name clashes.
func generateStatusTypes(swaggerTypes jsonast.SwaggerTypes) (statusTypes, error) {
	otherTypes, err := statusTypeRenamer.RenameAll(swaggerTypes.OtherDefinitions)
	if err != nil {
		return statusTypes{}, err
	}

	resources := make(resourceLookup)
	var errs []error
	for resourceName, resourceDef := range swaggerTypes.ResourceDefinitions {
		// resourceName is not renamed as this is a lookup for the Spec type
		var renamedType astmodel.Type
		renamedType, err = statusTypeRenamer.Rename(resourceDef.Type())
		if err != nil {
			errs = append(errs, err)
		} else {
			resources.add(resourceName, renamedType)
		}
	}

	err = kerrors.NewAggregate(errs)
	if err != nil {
		return statusTypes{}, err
	}

	return statusTypes{resources, otherTypes}, nil
}

func loadSwaggerData(ctx context.Context, idFactory astmodel.IdentifierFactory, config *config.Configuration) (jsonast.SwaggerTypes, error) {
	schemas, err := loadAllSchemas(ctx, config.Status.SchemaRoot, config.LocalPathPrefix(), idFactory, config.Status.Overrides)
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
		ResourceDefinitions: make(astmodel.TypeDefinitionSet),
		OtherDefinitions:    make(astmodel.TypeDefinitionSet),
	}

	for pkg, group := range m {
		klog.V(3).Infof("Merging types for %s", pkg)
		merged := mergeTypesForPackage(idFactory, group)
		result.ResourceDefinitions.AddTypes(merged.ResourceDefinitions)
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

	mergedResult := jsonast.SwaggerTypes{ResourceDefinitions: make(astmodel.TypeDefinitionSet), OtherDefinitions: make(astmodel.TypeDefinitionSet)}
	for _, typesFromFile := range typesFromFiles {
		for _, t := range typesFromFile.OtherDefinitions {
			// for consistent results we always sort typesFromFiles first (at top of this function)
			// so that we always pick the same one when there are multiple
			_ = mergedResult.OtherDefinitions.AddAllowDuplicates(t)
			// errors ignored since we already checked for structural equality
			// it’s possible for types to refer to different typenames in which case they are not TypeEquals Equal
			// but they might be structurally equal
		}

		err := mergedResult.ResourceDefinitions.AddTypesAllowDuplicates(typesFromFile.ResourceDefinitions)
		if err != nil {
			panic(fmt.Sprintf("merging file %s: %s", typesFromFile.filePath, err.Error()))
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
	typeNames map[astmodel.TypeName]int) astmodel.TypeName {
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
	newResourceTypes, err := visitor.RenameAll(typesFromFile.ResourceDefinitions)
	if err != nil {
		panic(err)
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
	rightDefinitions astmodel.TypeDefinitionSet) bool {

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
	overrides []config.SchemaOverride) (map[string]jsonast.PackageAndSwagger, error) {

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
				astmodel.GeneratorVersionPrefix,
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

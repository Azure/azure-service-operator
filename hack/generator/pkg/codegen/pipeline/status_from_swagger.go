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
	"strings"
	"sync"

	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/config"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/jsonast"
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
func AddStatusFromSwagger(idFactory astmodel.IdentifierFactory, config *config.Configuration) Stage {
	return MakeLegacyStage(
		"addStatusFromSwagger",
		"Add information from Swagger specs for 'status' fields",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			if config.Status.SchemaRoot == "" {
				klog.Warningf("No status schema root specified, will not generate status types")
				return types, nil
			}

			klog.V(1).Infof("Loading Swagger data from %q", config.Status.SchemaRoot)

			swaggerTypes, err := loadSwaggerData(ctx, idFactory, config)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to load Swagger data")
			}

			klog.V(1).Infof("Loaded Swagger data (%d resources, %d other types)", len(swaggerTypes.ResourceTypes), len(swaggerTypes.OtherTypes))

			statusTypes, err := generateStatusTypes(swaggerTypes)
			if err != nil {
				return nil, err
			}

			// put all types into a new set
			newTypes := make(astmodel.Types)
			// all non-resources from Swagger are added regardless of whether they are used
			// if they are not used they will be pruned off by a later pipeline stage
			// (there will be no name clashes here due to suffixing with "_Status")
			newTypes.AddTypes(statusTypes.otherTypes)

			matchedResources := 0
			// find any resources and update them with status info
			for typeName, typeDef := range types {
				if resource, ok := typeDef.Type().(*astmodel.ResourceType); ok {
					// find the status type (= Swagger resource type)
					newStatus, located := statusTypes.findResourceType(typeName)
					if located {
						matchedResources++
					}

					newTypes.Add(typeDef.WithType(resource.WithStatus(newStatus)))
				} else {
					newTypes.Add(typeDef)
				}
			}

			klog.V(1).Infof("Found status information for %d resources", matchedResources)
			klog.V(1).Infof("Input %d types, output %d types", len(types), len(newTypes))

			return newTypes, nil
		})
}

type statusTypes struct {
	// resourceTypes maps Spec name to corresponding Status type
	// the typeName is lowercased to be case-insensitive
	resourceTypes resourceLookup

	// otherTypes has all other Status types renamed to avoid clashes with Spec Types
	otherTypes astmodel.Types
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
var statusTypeRenamer astmodel.TypeVisitor = makeRenamingVisitor(appendStatusSuffix)

func appendStatusSuffix(typeName astmodel.TypeName) astmodel.TypeName {
	return astmodel.MakeTypeName(typeName.PackageReference, typeName.Name()+"_Status")
}

// generateStatusTypes returns the statusTypes for the input Swagger types
// all types (apart from Resources) are renamed to have "_Status" as a
// suffix, to avoid name clashes.
func generateStatusTypes(swaggerTypes jsonast.SwaggerTypes) (statusTypes, error) {
	var errs []error
	otherTypes := make(astmodel.Types)
	for _, typeDef := range swaggerTypes.OtherTypes {
		renamedDef, err := statusTypeRenamer.VisitDefinition(typeDef, nil)
		if err != nil {
			errs = append(errs, err)
		} else {
			otherTypes.Add(renamedDef)
		}
	}

	resources := make(resourceLookup)
	for resourceName, resourceDef := range swaggerTypes.ResourceTypes {
		// resourceName is not renamed as this is a lookup for the Spec type
		renamedDef, err := statusTypeRenamer.Visit(resourceDef.Type(), nil)
		if err != nil {
			errs = append(errs, err)
		} else {
			resources.add(resourceName, renamedDef)
		}
	}

	if len(errs) > 0 {
		return statusTypes{}, kerrors.NewAggregate(errs)
	}

	return statusTypes{resources, otherTypes}, nil
}

func makeRenamingVisitor(rename func(astmodel.TypeName) astmodel.TypeName) astmodel.TypeVisitor {
	return astmodel.TypeVisitorBuilder{
		VisitTypeName: func(this *astmodel.TypeVisitor, it astmodel.TypeName, ctx interface{}) (astmodel.Type, error) {
			return rename(it), nil
		},
	}.Build()
}

var swaggerVersionRegex = regexp.MustCompile(`/\d{4}-\d{2}-\d{2}(-preview|-privatepreview)?/`)

func loadSwaggerData(ctx context.Context, idFactory astmodel.IdentifierFactory, config *config.Configuration) (jsonast.SwaggerTypes, error) {
	schemas, err := loadAllSchemas(ctx, config.LocalPathPrefix(), config.TypeFilters, config.ExportFilters, config.Status.SchemaRoot)
	if err != nil {
		return jsonast.SwaggerTypes{}, err
	}

	cache := jsonast.NewOpenAPISchemaCache(schemas)

	typesByGroup := make(map[astmodel.LocalPackageReference][]typesFromFile)
	for schemaPath, schema := range schemas {
		// these have already been tested in the loadAllSchemas function so are guaranteed to match
		outputGroup := jsonast.SwaggerGroupRegex.FindString(schemaPath)
		// need to trim leading/trailing '/' which are matched in regex (go doesn’t support lookaround)
		outputVersion := strings.Trim(swaggerVersionRegex.FindString(schemaPath), "/")

		pkg := config.MakeLocalPackageReference(idFactory.CreateGroupName(outputGroup), astmodel.CreateLocalPackageNameFromVersion(outputVersion))

		// see if there is a config override for this file
		for _, schemaOverride := range config.Status.Overrides {
			configSchemaPath := path.Join(config.Status.SchemaRoot, schemaOverride.BasePath)
			if strings.HasPrefix(schemaPath, configSchemaPath) {
				// found an override: apply it
				if schemaOverride.Suffix != "" {
					outputGroup += "." + schemaOverride.Suffix
				}

				break
			}
		}

		extractor := jsonast.NewSwaggerTypeExtractor(
			config,
			idFactory,
			schema,
			schemaPath,
			outputGroup,
			outputVersion,
			cache)

		types, err := extractor.ExtractTypes(ctx)
		if err != nil {
			return jsonast.SwaggerTypes{}, errors.Wrapf(err, "error processing %q", schemaPath)
		}

		typesByGroup[pkg] = append(typesByGroup[pkg], typesFromFile{types, schemaPath})
	}

	return mergeSwaggerTypesByGroup(idFactory, typesByGroup), nil
}

func mergeSwaggerTypesByGroup(idFactory astmodel.IdentifierFactory, m map[astmodel.LocalPackageReference][]typesFromFile) jsonast.SwaggerTypes {
	klog.V(3).Infof("Merging types for %d groups/versions", len(m))

	result := jsonast.SwaggerTypes{
		ResourceTypes: make(astmodel.Types),
		OtherTypes:    make(astmodel.Types),
	}

	for pkg, group := range m {
		klog.V(3).Infof("Merging types for %s", pkg)
		merged := mergeTypesForPackage(idFactory, group)
		result.ResourceTypes.AddTypes(merged.ResourceTypes)
		result.OtherTypes.AddTypes(merged.OtherTypes)
	}

	return result
}

type typesFromFile struct {
	jsonast.SwaggerTypes
	filePath string
}

// mergeTypesForPackage merges the types for a single package from multiple files
func mergeTypesForPackage(idFactory astmodel.IdentifierFactory, typesFromFiles []typesFromFile) jsonast.SwaggerTypes {
	resourceNameCount := make(map[astmodel.TypeName]int)
	typeNameCount := make(map[astmodel.TypeName]int)
	for _, typesFromFile := range typesFromFiles {
		for name := range typesFromFile.ResourceTypes {
			resourceNameCount[name] += 1
		}

		for name := range typesFromFile.OtherTypes {
			typeNameCount[name] += 1
		}
	}

	// the results in the renames map match the order of the Types in `typesFromFiles`
	renames := make(map[astmodel.TypeName][]astmodel.TypeName)

	for name, count := range typeNameCount {
		if count == 1 {
			continue // not duplicate
		}

		type typeAndSource struct {
			def     astmodel.TypeDefinition
			typesIx int
		}

		typesToCheck := make([]typeAndSource, 0, count)
		for typesIx, types := range typesFromFiles {
			if def, ok := types.OtherTypes[name]; ok {
				typesToCheck = append(typesToCheck, typeAndSource{def: def, typesIx: typesIx})
				// short-circuit
				if len(typesToCheck) == count {
					break
				}
			}
		}

		first := typesToCheck[0]
		for _, other := range typesToCheck[1:] {
			if !structurallyIdentical(
				first.def.Type(),
				typesFromFiles[first.typesIx].OtherTypes,
				other.def.Type(),
				typesFromFiles[other.typesIx].OtherTypes,
			) {
				if name != first.def.Name() || name != other.def.Name() {
					panic("assert")
				}

				_, firstOk := first.def.Type().(*astmodel.ResourceType)
				_, otherOk := other.def.Type().(*astmodel.ResourceType)
				if firstOk || otherOk {
					panic("cannot rename resources")
				}

				renamed := make([]astmodel.TypeName, len(typesFromFiles))

				names := make([]string, len(typesToCheck))
				for ix, ttc := range typesToCheck {
					newName := generateRenaming(idFactory, first.def.Name(), typesFromFiles[ttc.typesIx].filePath, typeNameCount)
					names[ix] = newName.Name()
					renamed[ttc.typesIx] = newName
				}

				klog.V(3).Infof("Conflicting definitions for %s, renaming to: %s", first.def.Name(), strings.Join(names, ", "))

				renames[name] = renamed
				break
			} else {
				// klog.V(3).Infof("Conflicting definitions for %s were identical", first.def.Name())
				// should we need to collapse these in favour of one?
			}
		}
	}

	if len(renames) == 0 {
		// nothing needed renaming, we’re free to splat them all
		mergedResult := jsonast.SwaggerTypes{ResourceTypes: make(astmodel.Types), OtherTypes: make(astmodel.Types)}
		for _, typesFromFile := range typesFromFiles {
			for _, t := range typesFromFile.OtherTypes {
				_ = mergedResult.OtherTypes.AddAllowDuplicates(t)
				// errors ignored since we already checked for structural equality
				// it’s possible for types to refer to different typenames in which case they are not TypeEquals Equal
				// but they might be structurally equal
			}

			err := mergedResult.ResourceTypes.AddTypesAllowDuplicates(typesFromFile.ResourceTypes)
			if err != nil {
				panic(err)
			}
		}

		return mergedResult
	}

	for typeIx, typesFromFile := range typesFromFiles {
		visitor := makeRenamingVisitor(func(name astmodel.TypeName) astmodel.TypeName {
			if newNames, ok := renames[name]; ok {
				return newNames[typeIx]
			}

			return name
		})

		newOtherTypes := make(astmodel.Types)
		for _, def := range typesFromFile.OtherTypes {
			newDef, err := visitor.VisitDefinition(def, nil)
			if err != nil {
				panic(err)
			}

			newOtherTypes.Add(newDef)
		}

		newResourceTypes := make(astmodel.Types)
		for _, resourceDef := range typesFromFile.ResourceTypes {
			// must not rename the resource name, only within it
			newType, err := visitor.Visit(resourceDef.Type(), nil)
			if err != nil {
				panic(err)
			}

			newResourceTypes.Add(resourceDef.WithType(newType))
		}

		typesFromFiles[typeIx].OtherTypes = newOtherTypes
		typesFromFiles[typeIx].ResourceTypes = newResourceTypes
	}

	// try again!
	klog.V(3).Infof("Trying again…")
	return mergeTypesForPackage(idFactory, typesFromFiles)
}

func generateRenaming(
	idFactory astmodel.IdentifierFactory,
	original astmodel.TypeName,
	filePath string,
	typeNames map[astmodel.TypeName]int) astmodel.TypeName {
	name := filepath.Base(filePath)
	name = strings.TrimSuffix(name, filepath.Ext(name))

	result := astmodel.MakeTypeName(
		original.PackageReference,
		idFactory.CreateIdentifier(name+original.Name(), astmodel.Exported))

	for _, ok := typeNames[result]; ok; _, ok = typeNames[result] {
		result = astmodel.MakeTypeName(
			result.PackageReference,
			result.Name()+"X",
		)
	}

	return result
}

func structurallyIdentical(
	leftType astmodel.Type,
	leftTypes astmodel.Types,
	rightType astmodel.Type,
	rightTypes astmodel.Types) bool {

	override := astmodel.EqualityOverrides{}

	type pair struct{ left, right astmodel.TypeName }

	toCheck := []pair{}
	checked := map[pair]struct{}{}

	// note that this relies on Equals implementations preserving the left/right order
	override.TypeName = func(left, right astmodel.TypeName) bool {
		p := pair{left, right}
		if _, ok := checked[p]; !ok {
			checked[p] = struct{}{}
			toCheck = append(toCheck, p)
		}

		// conditionally true, as long as all pairs are equal
		return true
	}

	if !leftType.Equals(rightType, override) {
		return false
	}

	for len(toCheck) > 0 {
		next := toCheck[0]
		toCheck = toCheck[1:]
		if !leftTypes[next.left].Type().Equals(rightTypes[next.right].Type(), override) {
			return false
		}
	}

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
	packagePrefix string,
	filters []*config.TypeFilter,
	exportFilters []*config.ExportFilter,
	rootPath string) (map[string]spec.Swagger, error) {

	var eg errgroup.Group

	var mutex sync.Mutex
	schemas := make(map[string]spec.Swagger)

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
			filepath.Ext(filePath) == ".json" &&
			jsonast.SwaggerGroupRegex.MatchString(filePath) &&
			swaggerVersionRegex.MatchString(filePath) {

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
				schemas[filePath] = swagger
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

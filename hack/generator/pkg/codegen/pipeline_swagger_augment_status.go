/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

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

	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/config"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/jsonast"
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type statusTypes struct {
	// resourceTypes maps Spec name to corresponding Status type
	// the typeName is lowercased to be case-insensitive
	resourceTypes resourceLookup

	// otherTypes has all other Status types renamed to avoid clashes with Spec Types
	otherTypes astmodel.Types
}

func (st statusTypes) findResourceType(typeName astmodel.TypeName) (astmodel.Type, bool) {
	if statusDef, ok := st.resourceTypes.tryFind(typeName); ok {
		klog.V(4).Infof("Swagger information found for %v", typeName)
		return statusDef, true
	} else {
		klog.V(4).Infof("Swagger information missing for %v", typeName)
		// add a warning that the status is missing
		// this will be reported if the type is not pruned
		return astmodel.NewErroredType(nil, nil, []string{fmt.Sprintf("missing status information for %v", typeName)}), false
	}
}

/*
func (st statusTypes) findNonResourceType(typeName astmodel.TypeName) (astmodel.Type, bool) {
	typeName = appendStatusSuffix(typeName) // we must look up the "_Status" version of the type
	if statusDef, ok := st.otherTypes[typeName]; ok {
		return statusDef.Type(), true
	} else {
		return nil, false
	}
}
*/

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
		panic(fmt.Sprintf("lowercase name collision: %v", name))
	}

	resourceLookup[lower] = theType
}

// statusTypeRenamer appends "_Status" to all types
var statusTypeRenamer astmodel.TypeVisitor = makeRenamingVisitor(appendStatusSuffix)

func appendStatusSuffix(typeName astmodel.TypeName) astmodel.TypeName {
	return astmodel.MakeTypeName(typeName.PackageReference, typeName.Name()+"_Status")
}

// generateStatusTypes returns the statusTypes for the input swaggerTypes
// all types (apart from Resources) are renamed to have "_Status" as a
// suffix, to avoid name clashes.
func generateStatusTypes(swaggerTypes swaggerTypes) (statusTypes, error) {

	var errs []error
	otherTypes := make(astmodel.Types)
	for _, typeDef := range swaggerTypes.otherTypes {
		renamedDef, err := statusTypeRenamer.VisitDefinition(typeDef, nil)
		if err != nil {
			errs = append(errs, err)
		} else {
			otherTypes.Add(renamedDef)
		}
	}

	resources := make(resourceLookup)
	for resourceName, resourceDef := range swaggerTypes.resources {
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

var swaggerVersionRegex = regexp.MustCompile(`\d{4}-\d{2}-\d{2}(-preview)?`)

type swaggerTypes struct {
	resources  astmodel.Types
	otherTypes astmodel.Types
}

func loadSwaggerData(ctx context.Context, idFactory astmodel.IdentifierFactory, config *config.Configuration) (swaggerTypes, error) {

	result := swaggerTypes{
		resources:  make(astmodel.Types),
		otherTypes: make(astmodel.Types),
	}

	schemas, err := loadAllSchemas(ctx, config.Status.SchemaRoot)
	if err != nil {
		return swaggerTypes{}, err
	}

	cache := jsonast.NewOpenAPISchemaCache(schemas)

	for schemaPath, schema := range schemas {
		// these have already been tested in the loadAllSchemas function so are guaranteed to match
		outputGroup := jsonast.SwaggerGroupRegex.FindString(schemaPath)
		outputVersion := swaggerVersionRegex.FindString(schemaPath)

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
			outputGroup,
			outputVersion,
			cache)

		err := extractor.ExtractTypes(ctx, schemaPath, schema, result.resources, result.otherTypes)
		if err != nil {
			return swaggerTypes{}, errors.Wrapf(err, "error processing %q", schemaPath)
		}
	}

	return result, nil
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

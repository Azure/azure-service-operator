/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/config"
	"github.com/Azure/k8s-infra/hack/generator/pkg/jsonast"
	"github.com/xeipuuv/gojsonreference"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"

	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"
)

// CodeGenerator is a generator of code
type CodeGenerator struct {
	configuration *config.Configuration
}

// NewCodeGenerator produces a new Generator with the given configuration
func NewCodeGenerator(configurationFile string) (*CodeGenerator, error) {
	config, err := loadConfiguration(configurationFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration file '%v' (%w)", configurationFile, err)
	}

	err = config.Initialize()
	if err != nil {
		return nil, fmt.Errorf("configuration loaded from '%v' is invalid (%w)", configurationFile, err)
	}

	result := &CodeGenerator{configuration: config}

	return result, nil
}

// Generate produces the Go code corresponding to the configured JSON schema in the given output folder
func (generator *CodeGenerator) Generate(ctx context.Context, outputFolder string) error {
	klog.V(0).Infof("Loading JSON schema %v", generator.configuration.SchemaURL)
	schema, err := loadSchema(ctx, generator.configuration.SchemaURL)
	if err != nil {
		return fmt.Errorf("error loading schema from '%v' (%w)", generator.configuration.SchemaURL, err)
	}

	klog.V(0).Infof("Cleaning output folder '%v'", outputFolder)
	err = deleteGeneratedCodeFromFolder(ctx, outputFolder)
	if err != nil {
		return fmt.Errorf("error cleaning output folder '%v' (%w)", outputFolder, err)
	}

	scanner := jsonast.NewSchemaScanner(astmodel.NewIdentifierFactory(), generator.configuration)

	klog.V(0).Infof("Walking JSON schema")

	defs, err := scanner.GenerateDefinitions(ctx, schema.Root())
	if err != nil {
		return fmt.Errorf("failed to walk JSON schema (%w)", err)
	}

	defs, err = generator.FilterDefinitions(defs)
	if err != nil {
		return fmt.Errorf("failed to filter generated definitions (%w)", err)
	}

	packages, err := generator.CreatePackagesForDefinitions(defs)
	if err != nil {
		return fmt.Errorf("failed to assign generated definitions to packages (%w)", err)
	}

	packages, err = generator.MarkLatestResourceVersionsForStorage(packages)
	if err != nil {
		return fmt.Errorf("unable to mark latest resource versions for as storage versions (%w)", err)
	}

	fileCount := 0
	definitionCount := 0

	// emit each package
	klog.V(0).Infof("Writing output files into %v", outputFolder)
	for _, pkg := range packages {
		if ctx.Err() != nil { // check for cancellation
			return ctx.Err()
		}

		// create directory if not already there
		outputDir := filepath.Join(outputFolder, pkg.GroupName, pkg.PackageName)
		if _, err := os.Stat(outputDir); os.IsNotExist(err) {
			klog.V(5).Infof("Creating directory '%s'\n", outputDir)
			err = os.MkdirAll(outputDir, 0700)
			if err != nil {
				klog.Fatalf("Unable to create directory '%s'", outputDir)
			}
		}

		count, err := pkg.EmitDefinitions(outputDir)
		if err != nil {
			return fmt.Errorf("error writing definitions into '%v' (%w)", outputDir, err)
		}

		fileCount += count
		definitionCount += pkg.DefinitionCount()
	}

	klog.V(0).Infof("Completed writing %v files containing %v definitions", fileCount, definitionCount)

	return nil
}

// MarkLatestResourceVersionsForStorage marks the latest version of each resource as the storage version
func (generator *CodeGenerator) MarkLatestResourceVersionsForStorage(
	pkgs []*astmodel.PackageDefinition) ([]*astmodel.PackageDefinition, error) {

	var result []*astmodel.PackageDefinition

	resourceLookup, err := groupResourcesByVersion(pkgs)
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgs {

		resultPkg := astmodel.NewPackageDefinition(pkg.GroupName, pkg.PackageName)
		for _, def := range pkg.Definitions() {
			// see if it is a resource (only struct definitions can be resources)
			if structDef, ok := def.(*astmodel.StructDefinition); ok && structDef.IsResource() {

				unversionedName, err := getUnversionedName(structDef.TypeName)
				if err != nil {
					// should never happen as all resources have versioned names
					return nil, err
				}

				allVersionsOfResource := resourceLookup[unversionedName]
				latestVersionOfResource := allVersionsOfResource[len(allVersionsOfResource)-1]

				thisPackagePath := structDef.Name().PackageReference.PackagePath()
				latestPackagePath := latestVersionOfResource.Name().PackageReference.PackagePath()

				// mark as storage version if it's the latest version
				isLatestVersion := thisPackagePath == latestPackagePath
				structDef = structDef.WithIsStorageVersion(isLatestVersion)

				resultPkg.AddDefinition(structDef)
			} else {
				// otherwise simply add it
				resultPkg.AddDefinition(def)
			}
		}

		result = append(result, resultPkg)
	}

	return result, nil
}

func getUnversionedName(name *astmodel.TypeName) (unversionedName, error) {
	group, _, err := name.PackageReference.GroupAndPackage()
	if err != nil {
		return unversionedName{}, err
	}

	return unversionedName{group, name.Name()}, nil
}

type unversionedName struct {
	group string
	name  string
}

func groupResourcesByVersion(
	pkgs []*astmodel.PackageDefinition) (map[unversionedName][]*astmodel.StructDefinition, error) {

	result := make(map[unversionedName][]*astmodel.StructDefinition)

	for _, pkg := range pkgs {
		for _, def := range pkg.Definitions() {
			if structDef, ok := def.(*astmodel.StructDefinition); ok && structDef.IsResource() {
				name, err := getUnversionedName(structDef.TypeName)
				if err != nil {
					// this should never happen as resources will all have versioned names
					return nil, fmt.Errorf("Unable to extract unversioned name in groupResources: %w", err)
				}

				result[name] = append(result[name], structDef)
			}
		}
	}

	// order each set of resources by package name (== by version as these are sortable dates)
	for _, slice := range result {
		sort.Slice(slice, func(i, j int) bool {
			return slice[i].TypeName.PackageReference.PackageName() < slice[j].TypeName.PackageReference.PackageName()
		})
	}

	return result, nil
}

// FilterDefinitions applies the configuration include/exclude filters to the generated definitions
func (generator *CodeGenerator) FilterDefinitions(
	definitions []astmodel.TypeDefiner) ([]astmodel.TypeDefiner, error) {

	var newDefinitions []astmodel.TypeDefiner

	for _, def := range definitions {
		defName := def.Name()
		shouldExport, reason := generator.configuration.ShouldExport(defName)

		switch shouldExport {
		case config.Skip:
			klog.V(2).Infof("Skipping %s because %s", defName, reason)

		case config.Export:
			if reason == "" {
				klog.V(3).Infof("Exporting %s", defName)
			} else {
				klog.V(2).Infof("Exporting %s because %s", defName, reason)
			}

			newDefinitions = append(newDefinitions, def)
		}
	}

	return newDefinitions, nil
}

// CreatePackagesForDefinitions groups type definitions into packages
func (generator *CodeGenerator) CreatePackagesForDefinitions(
	definitions []astmodel.TypeDefiner) ([]*astmodel.PackageDefinition, error) {

	packages := make(map[astmodel.PackageReference]*astmodel.PackageDefinition)
	for _, def := range definitions {
		defName := def.Name()
		groupName, pkgName, err := defName.PackageReference.GroupAndPackage()
		if err != nil {
			return nil, err
		}

		pkgRef := defName.PackageReference
		if pkg, ok := packages[pkgRef]; ok {
			pkg.AddDefinition(def)
		} else {
			pkg = astmodel.NewPackageDefinition(groupName, pkgName)
			pkg.AddDefinition(def)
			packages[pkgRef] = pkg
		}
	}

	var pkgs []*astmodel.PackageDefinition
	for _, pkg := range packages {
		pkgs = append(pkgs, pkg)
	}

	return pkgs, nil
}

func loadConfiguration(configurationFile string) (*config.Configuration, error) {
	data, err := ioutil.ReadFile(configurationFile)
	if err != nil {
		return nil, err
	}

	result := config.NewConfiguration()

	err = yaml.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type cancellableFileSystem struct {
	ctx context.Context
}

var _ http.FileSystem = &cancellableFileSystem{} // interface assertion

func (fs *cancellableFileSystem) Open(source string) (http.File, error) {
	if fs.ctx.Err() != nil { // check for cancellation
		return nil, fs.ctx.Err()
	}

	return os.Open(source)
}

type cancellableJSONLoaderFactory struct {
	ctx   context.Context
	inner gojsonschema.JSONLoaderFactory
}

var _ gojsonschema.JSONLoaderFactory = &cancellableJSONLoaderFactory{}

func (factory *cancellableJSONLoaderFactory) New(source string) gojsonschema.JSONLoader {
	return &cancellableJSONLoader{factory.ctx, factory.inner.New(source)}
}

type cancellableJSONLoader struct {
	ctx   context.Context
	inner gojsonschema.JSONLoader
}

var _ gojsonschema.JSONLoader = &cancellableJSONLoader{}

func (loader *cancellableJSONLoader) LoadJSON() (interface{}, error) {
	if loader.ctx.Err() != nil { // check for cancellation
		return nil, loader.ctx.Err()
	}

	return loader.inner.LoadJSON()
}

func (loader *cancellableJSONLoader) JsonSource() interface{} {
	return loader.inner.JsonSource()
}

func (loader *cancellableJSONLoader) JsonReference() (gojsonreference.JsonReference, error) {
	if loader.ctx.Err() != nil { // check for cancellation
		return gojsonreference.JsonReference{}, loader.ctx.Err()
	}

	return loader.inner.JsonReference()
}

func (loader *cancellableJSONLoader) LoaderFactory() gojsonschema.JSONLoaderFactory {
	return &cancellableJSONLoaderFactory{loader.ctx, loader.inner.LoaderFactory()}
}

func loadSchema(ctx context.Context, source string) (*gojsonschema.Schema, error) {
	sl := gojsonschema.NewSchemaLoader()
	loader := &cancellableJSONLoader{
		ctx,
		gojsonschema.NewReferenceLoaderFileSystem(source, &cancellableFileSystem{ctx}),
	}

	schema, err := sl.Compile(loader)
	if err != nil {
		return nil, fmt.Errorf("error loading schema from '%v' (%w)", source, err)
	}

	return schema, nil
}

func deleteGeneratedCodeFromFolder(ctx context.Context, outputFolder string) error {
	globPattern := path.Join(outputFolder, "**", "*", "*"+astmodel.CodeGeneratedFileSuffix) + "*"

	files, err := filepath.Glob(globPattern)
	if err != nil {
		return fmt.Errorf("error globbing files with pattern '%s' (%w)", globPattern, err)
	}

	var errs []error

	for _, file := range files {
		if ctx.Err() != nil { // check for cancellation
			return ctx.Err()
		}

		isGenerated, err := isFileGenerated(file)

		if err != nil {
			errs = append(errs, fmt.Errorf("error determining if file was generated (%w)", err))
		}

		if isGenerated {
			err := os.Remove(file)
			if err != nil {
				errs = append(errs, fmt.Errorf("error removing file '%v' (%w)", file, err))
			}
		}
	}

	err = deleteEmptyDirectories(ctx, outputFolder)
	if err != nil {
		errs = append(errs, err)
	}

	return kerrors.NewAggregate(errs)
}

func isFileGenerated(filename string) (bool, error) {
	// Technically, the code generated message could be on any line according to
	// the specification at https://github.com/golang/go/issues/13560 but
	// for our purposes checking the first few lines is plenty
	maxLinesToCheck := 20

	f, err := os.Open(filename)
	if err != nil {
		return false, err
	}

	reader := bufio.NewReader(f)
	for i := 0; i < maxLinesToCheck; i++ {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return false, nil
		}
		if err != nil {
			return false, err
		}

		if strings.Contains(line, astmodel.CodeGenerationComment) {
			return true, nil
		}
	}
	defer f.Close()

	return false, nil
}

func deleteEmptyDirectories(ctx context.Context, path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}

	// TODO: There has to be a better way to do this?
	var dirs []string

	// Second pass to clean up empty directories
	walkFunction := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if ctx.Err() != nil { // check for cancellation
			return ctx.Err()
		}

		if info.IsDir() {
			dirs = append(dirs, path)
		}

		return nil
	}
	err := filepath.Walk(path, walkFunction)
	if err != nil {
		return err
	}

	// Now order the directories by deepest first - we have to do this because otherwise a directory
	// isn't empty because it has a bunch of empty directories inside of it
	sortFunction := func(i int, j int) bool {
		// Comparing by length is sufficient here because a nested directory path
		// will always be longer than just the parent directory path
		return len(dirs[i]) > len(dirs[j])
	}
	sort.Slice(dirs, sortFunction)

	var errs []error

	// Now clean things up
	for _, dir := range dirs {
		if ctx.Err() != nil { // check for cancellation
			return ctx.Err()
		}

		files, err := ioutil.ReadDir(dir)
		if err != nil {
			errs = append(errs, fmt.Errorf("error reading directory '%v' (%w)", dir, err))
		}

		if len(files) == 0 {
			// Directory is empty now, we can delete it
			err := os.Remove(dir)
			if err != nil {
				errs = append(errs, fmt.Errorf("error removing dir '%v' (%w)", dir, err))
			}
		}
	}

	return kerrors.NewAggregate(errs)
}

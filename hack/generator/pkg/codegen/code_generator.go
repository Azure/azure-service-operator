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
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/jsonast"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"

	"github.com/hashicorp/go-multierror"
	"k8s.io/klog/v2"
)

type CodeGenerator struct {
	configuration *Configuration
}

func NewCodeGenerator(configurationFile string) (*CodeGenerator, error) {
	config, err := loadConfiguration(configurationFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration file '%v' (%w)", configurationFile, err)
	}

	err = config.Validate()
	if err != nil {
		return nil, fmt.Errorf("configuration loaded from '%v' is invalid (%w)", configurationFile, err)
	}

	result := &CodeGenerator{configuration: config}

	return result, nil
}

func (generator *CodeGenerator) Generate(ctx context.Context, outputFolder string) error {

	klog.V(0).Infof("Loading JSON schema %v", generator.configuration.SchemaURL)
	schema, err := loadSchema(generator.configuration.SchemaURL)
	if err != nil {
		return fmt.Errorf("error loading schema from '%v' (%w)", generator.configuration.SchemaURL, err)
	}

	klog.V(0).Infof("Cleaning output folder '%v'", outputFolder)
	err = deleteGeneratedCodeFromFolder(outputFolder)
	if err != nil {
		return fmt.Errorf("error cleaning output folder '%v' (%w)", outputFolder, err)
	}

	scanner := jsonast.NewSchemaScanner(astmodel.NewIdentifierFactory())

	klog.V(0).Infof("Walking JSON schema")
	defs, err := scanner.GenerateDefinitions(ctx, schema.Root())
	if err != nil {
		return fmt.Errorf("failed to walk JSON schema (%w)", err)
	}

	packages, err := generator.CreatePackagesForDefinitions(defs)
	if err != nil {
		return fmt.Errorf("failed to assign generated definitions to packages (%w)", err)
	}

	fileCount := 0
	definitionCount := 0

	// emit each package
	klog.V(0).Infof("Writing output files into %v", outputFolder)
	for _, pkg := range packages {

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

func (generator *CodeGenerator) CreatePackagesForDefinitions(definitions []astmodel.TypeDefiner) ([]*astmodel.PackageDefinition, error) {
	packages := make(map[astmodel.PackageReference]*astmodel.PackageDefinition)
	for _, def := range definitions {

		shouldExport, reason := generator.configuration.ShouldExport(def)
		defName := def.Name()
		groupName, pkgName, err := defName.PackageReference.GroupAndPackage()
		if err != nil {
			return nil, err
		}

		switch shouldExport {
		case Skip:
			klog.V(2).Infof("Skipping %s/%s because %s", groupName, pkgName, reason)

		case Export:
			if reason == "" {
				klog.V(3).Infof("Exporting %s/%s", groupName, pkgName)
			} else {
				klog.V(2).Infof("Exporting %s/%s because %s", groupName, pkgName, reason)
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
	}

	var pkgs []*astmodel.PackageDefinition
	for _, pkg := range packages {
		pkgs = append(pkgs, pkg)
	}

	return pkgs, nil
}

func loadConfiguration(configurationFile string) (*Configuration, error) {
	data, err := ioutil.ReadFile(configurationFile)
	if err != nil {
		return nil, err
	}

	result := NewConfiguration()

	err = yaml.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func loadSchema(source string) (*gojsonschema.Schema, error) {
	sl := gojsonschema.NewSchemaLoader()
	schema, err := sl.Compile(gojsonschema.NewReferenceLoader(source))
	if err != nil {
		return nil, fmt.Errorf("error loading schema from '%v' (%w)", source, err)
	}

	return schema, nil
}

func deleteGeneratedCodeFromFolder(outputFolder string) error {
	globPattern := path.Join(outputFolder, "**", "*", "*"+astmodel.CodeGeneratedFileSuffix) + "*"

	files, err := filepath.Glob(globPattern)
	if err != nil {
		return fmt.Errorf("error globbing files with pattern '%s' (%w)", globPattern, err)
	}

	var result *multierror.Error

	for _, file := range files {
		isGenerated, err := isFileGenerated(file)

		if err != nil {
			result = multierror.Append(result, fmt.Errorf("error determining if file was generated (%w)", err))
		}

		if isGenerated {
			err := os.Remove(file)
			if err != nil {
				result = multierror.Append(result, fmt.Errorf("error removing file '%v' (%w)", file, err))
			}
		}
	}

	err = deleteEmptyDirectories(outputFolder)
	if err != nil {
		result = multierror.Append(result, err)
	}

	return result.ErrorOrNil()
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

func deleteEmptyDirectories(path string) error {
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

	var result *multierror.Error

	// Now clean things up
	for _, dir := range dirs {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			result = multierror.Append(result, fmt.Errorf("error reading directory '%v' (%w)", dir, err))
		}

		if len(files) == 0 {
			// Directory is empty now, we can delete it
			err := os.Remove(dir)
			if err != nil {
				result = multierror.Append(result, fmt.Errorf("error removing dir '%v' (%w)", dir, err))
			}
		}
	}

	return result.ErrorOrNil()
}

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
	"path/filepath"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/jsonast"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"

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
	err = cleanFolder(outputFolder)
	if err != nil {
		return fmt.Errorf("error cleaning output folder '%v' (%w)", generator.configuration.SchemaURL, err)
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
		groupName, pkgName, err := defName.GroupAndPackage()
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

//TODO: Only clean generated files
func cleanFolder(outputFolder string) error {
	err := os.RemoveAll(outputFolder)
	if err != nil {
		return fmt.Errorf("error removing output folder '%v' (%w)", outputFolder, err)
	}

	err = os.Mkdir(outputFolder, 0700)
	if err != nil {
		return fmt.Errorf("error creating output folder '%v' (%w)", outputFolder, err)
	}

	return nil
}

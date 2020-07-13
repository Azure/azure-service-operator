/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/config"
	"github.com/Azure/k8s-infra/hack/generator/pkg/jsonast"
	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonreference"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"k8s.io/klog/v2"
	"net/http"
	"os"
)

// CodeGenerator is a generator of code
type CodeGenerator struct {
	configuration *config.Configuration
}

// NewCodeGenerator produces a new Generator with the given configuration
func NewCodeGenerator(configurationFile string) (*CodeGenerator, error) {
	configuration, err := loadConfiguration(configurationFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load configuration file %q", configurationFile)
	}

	err = configuration.Initialize()
	if err != nil {
		return nil, errors.Wrapf(err, "configuration loaded from %q is invalid", configurationFile)
	}

	result := &CodeGenerator{configuration: configuration}

	return result, nil
}

// Generate produces the Go code corresponding to the configured JSON schema in the given output folder
func (generator *CodeGenerator) Generate(ctx context.Context) error {
	klog.V(1).Infof("Generator version: %v", combinedVersion())
	klog.V(0).Infof("Loading JSON schema %v", generator.configuration.SchemaURL)
	schema, err := loadSchema(ctx, generator.configuration.SchemaURL)
	if err != nil {
		return errors.Wrapf(err, "error loading schema from %q", generator.configuration.SchemaURL)
	}

	scanner := jsonast.NewSchemaScanner(astmodel.NewIdentifierFactory(), generator.configuration)

	klog.V(0).Infof("Walking JSON schema")

	defs, err := scanner.GenerateDefinitions(ctx, schema.Root())
	if err != nil {
		return errors.Wrapf(err, "failed to walk JSON schema")
	}

	pipeline := []PipelineStage{
		applyExportFilters(generator.configuration),
		stripUnreferencedTypeDefinitions(),
		deleteGeneratedCode(generator.configuration.OutputPath),
		exportPackages(generator.configuration.OutputPath),
	}

	for i, stage := range pipeline {
		klog.V(0).Infof("Pipeline stage %d/%d: %s", i+1, len(pipeline), stage.Name)
		defs, err = stage.Action(ctx, defs)
		if err != nil {
			return errors.Wrapf(err, "Failed during pipeline stage %d/%d: %s", i+1, len(pipeline), stage.Name)
		}
	}

	return nil
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
		return nil, errors.Wrapf(err, "error loading schema from %q", source)
	}

	return schema, nil
}

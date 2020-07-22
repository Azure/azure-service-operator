/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/config"
	"github.com/Azure/k8s-infra/hack/generator/pkg/jsonast"
	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonreference"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
	"k8s.io/klog/v2"
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

	idFactory := astmodel.NewIdentifierFactory()

	pipeline := []PipelineStage{
		loadSchema(ctx, idFactory, generator.configuration),
		nameTypesForCRD(idFactory),
		applyExportFilters(generator.configuration),
		stripUnreferencedTypeDefinitions(),
		deleteGeneratedCode(generator.configuration.OutputPath),
		exportPackages(generator.configuration.OutputPath),
	}

	defs := make(Types)
	var err error

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

func loadSchema(ctx context.Context, idFactory astmodel.IdentifierFactory, configuration *config.Configuration) PipelineStage {
	source := configuration.SchemaURL

	return PipelineStage{
		Name: fmt.Sprintf("Load and walk schema from %q", source),
		Action: func(ctx context.Context, types Types) (Types, error) {
			klog.V(0).Infof("Loading JSON schema %q", source)

			sl := gojsonschema.NewSchemaLoader()
			loader := &cancellableJSONLoader{
				ctx,
				gojsonschema.NewReferenceLoaderFileSystem(source, &cancellableFileSystem{ctx}),
			}

			schema, err := sl.Compile(loader)
			if err != nil {
				return nil, errors.Wrapf(err, "error loading schema from %q", source)
			}

			scanner := jsonast.NewSchemaScanner(idFactory, configuration)

			klog.V(0).Infof("Walking JSON schema")

			defs, err := scanner.GenerateDefinitions(ctx, schema.Root())
			if err != nil {
				return nil, errors.Wrapf(err, "failed to walk JSON schema")
			}

			return defs, nil
		},
	}
}

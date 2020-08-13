/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"net/http"
	"os"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/config"
	"github.com/Azure/k8s-infra/hack/generator/pkg/jsonast"
	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonreference"
	"github.com/xeipuuv/gojsonschema"

	"k8s.io/klog/v2"
)

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

type schemaLoader func(ctx context.Context, source string) (*gojsonschema.Schema, error)

func defaultSchemaLoader(ctx context.Context, source string) (*gojsonschema.Schema, error) {
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

func loadSchemaIntoTypes(
	idFactory astmodel.IdentifierFactory,
	configuration *config.Configuration,
	schemaLoader schemaLoader) PipelineStage {
	source := configuration.SchemaURL

	return MakePipelineStage(
		"loadSchema",
		"Load and walk schema",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			klog.V(0).Infof("Loading JSON schema %q", source)

			schema, err := schemaLoader(ctx, source)
			if err != nil {
				return nil, err
			}

			scanner := jsonast.NewSchemaScanner(idFactory, configuration)

			klog.V(0).Infof("Walking JSON schema")

			defs, err := scanner.GenerateDefinitions(ctx, schema.Root())
			if err != nil {
				return nil, errors.Wrapf(err, "failed to walk JSON schema")
			}

			return defs, nil
		})
}

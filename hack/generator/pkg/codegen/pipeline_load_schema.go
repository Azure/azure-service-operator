/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/config"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/jsonast"
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

type rewritingJSONLoader struct {
	from, to string
	inner    gojsonschema.JSONLoader
}

var _ gojsonschema.JSONLoader = &rewritingJSONLoader{}

func (loader *rewritingJSONLoader) LoadJSON() (interface{}, error) {
	return loader.inner.LoadJSON()
}

func (loader *rewritingJSONLoader) JsonSource() interface{} {
	return loader.inner.JsonSource()
}

func (loader *rewritingJSONLoader) JsonReference() (gojsonreference.JsonReference, error) {
	return loader.inner.JsonReference()
}

func (loader *rewritingJSONLoader) LoaderFactory() gojsonschema.JSONLoaderFactory {
	return &rewritingJSONLoaderFactory{loader.from, loader.to, loader.inner.LoaderFactory()}
}

type rewritingJSONLoaderFactory struct {
	from, to string
	inner    gojsonschema.JSONLoaderFactory
}

var _ gojsonschema.JSONLoaderFactory = &rewritingJSONLoaderFactory{}

func (factory *rewritingJSONLoaderFactory) New(source string) gojsonschema.JSONLoader {
	return factory.inner.New(strings.Replace(source, factory.from, factory.to, 1))
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

type schemaLoader func(ctx context.Context, rewrite *config.RewriteRule, source string) (*gojsonschema.Schema, error)

func defaultSchemaLoader(ctx context.Context, rewrite *config.RewriteRule, source string) (*gojsonschema.Schema, error) {
	sl := gojsonschema.NewSchemaLoader()
	var loader gojsonschema.JSONLoader = &cancellableJSONLoader{ctx, gojsonschema.NewReferenceLoaderFileSystem(source, &cancellableFileSystem{ctx})}

	if rewrite != nil {
		loader = &rewritingJSONLoader{
			from:  rewrite.From,
			to:    rewrite.To,
			inner: loader,
		}
	}

	schema, err := sl.Compile(loader)
	if err != nil {
		return nil, errors.Wrapf(err, "error loading schema from root %q (error might be in another schema)", source)
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

			schema, err := schemaLoader(ctx, configuration.SchemaURLRewrite, source)
			if err != nil {
				return nil, err
			}

			scanner := jsonast.NewSchemaScanner(idFactory, configuration)

			klog.V(0).Infof("Walking deployment template")

			defs, err := scanner.GenerateDefinitionsFromDeploymentTemplate(ctx, jsonast.MakeGoJSONSchema(schema.Root()))
			if err != nil {
				return nil, errors.Wrapf(err, "failed to walk JSON schema")
			}

			return defs, nil
		})
}

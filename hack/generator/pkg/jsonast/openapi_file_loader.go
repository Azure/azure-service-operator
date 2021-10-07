/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

type OpenAPIFileLoader interface {
	loadFile(absPath string) (PackageAndSwagger, error)

	knownFiles() []string
}

type PackageAndSwagger struct {
	Package *astmodel.LocalPackageReference
	Swagger spec.Swagger
}

// CachingFileLoader is a cache of schema that have been loaded,
// identified by file path
type CachingFileLoader struct {
	files map[string]PackageAndSwagger
}

var _ OpenAPIFileLoader = CachingFileLoader{}

// NewCachingFileLoader creates an OpenAPISchemaCache with the initial
// file path → spec mapping
func NewCachingFileLoader(specs map[string]PackageAndSwagger) CachingFileLoader {
	files := make(map[string]PackageAndSwagger)
	for specPath, spec := range specs {
		files[filepath.ToSlash(specPath)] = spec
	}

	return CachingFileLoader{files}
}

func (fileCache CachingFileLoader) knownFiles() []string {
	result := make([]string, 0, len(fileCache.files))
	for k := range fileCache.files {
		result = append(result, k)
	}

	return result
}

// fetchFileAbsolute fetches the schema for the absolute path specified
func (fileCache CachingFileLoader) loadFile(absPath string) (PackageAndSwagger, error) {
	if !filepath.IsAbs(absPath) {
		panic(fmt.Sprintf("filePath %s must be absolute", absPath)) // assertion, not error
	}

	key := filepath.ToSlash(absPath)
	if swagger, ok := fileCache.files[key]; ok {
		return swagger, nil
	}

	// here the package will be unpopulated,
	// which indicates to the caller to reuse the existing package for definitions
	result := PackageAndSwagger{}

	klog.V(3).Infof("Loading file into cache %q", absPath)

	fileContent, err := ioutil.ReadFile(absPath)
	if err != nil {
		return result, errors.Wrapf(err, "unable to read swagger file %q", absPath)
	}

	err = result.Swagger.UnmarshalJSON(fileContent)
	if err != nil {
		return result, errors.Wrapf(err, "unable to parse swagger file %q", absPath)
	}

	fileCache.files[key] = result

	return result, err
}

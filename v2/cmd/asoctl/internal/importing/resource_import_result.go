/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
	"sigs.k8s.io/yaml"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// ResourceImportResult represents the result of an import operation
type ResourceImportResult struct {
	resources []genruntime.MetaObject
}

// Count returns the number of successfully imported resources.
func (r *ResourceImportResult) Count() int {
	return len(r.resources)
}

func (r *ResourceImportResult) SaveToWriter(destination io.Writer) error {
	return r.writeTo(r.resources, destination)
}

func (r *ResourceImportResult) SaveToSingleFile(filepath string) error {
	return r.saveTo(r.resources, filepath)
}

func (r *ResourceImportResult) SaveToIndividualFilesInFolder(folder string) error {
	// We name the files after the resource type and name
	// We allocate resources to files using a map, just in case we have a naming collision
	// (If that happens, all the similarly named resources will be in the same file, which is not ideal,
	// but better than dropping one or more)
	fileMap := make(map[string][]genruntime.MetaObject, len(r.resources))
	for _, resource := range r.resources {
		resourceName := resource.GetName()
		typeName := resource.GetObjectKind().GroupVersionKind().Kind
		fileName := fmt.Sprintf("%s-%s.yaml", typeName, resourceName)
		fileMap[fileName] = append(fileMap[fileName], resource)
	}

	for fileName, resources := range fileMap {
		path := filepath.Join(folder, fileName)
		err := r.saveTo(resources, path)
		if err != nil {
			return errors.Wrapf(err, "unable to save to file %s", path)
		}
	}

	return nil
}

func (r *ResourceImportResult) saveTo(resources []genruntime.MetaObject, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.Wrapf(err, "unable to create file %s", path)
	}

	defer func() {
		file.Close()

		// if we are panicking, the file will be in a broken
		// state, so remove it
		if r := recover(); r != nil {
			os.Remove(path)
			panic(r)
		}
	}()

	err = r.writeTo(resources, file)
	if err != nil {
		// cleanup in case of errors
		file.Close()
		os.Remove(path)
	}

	return errors.Wrapf(err, "unable to save to file %s", path)
}

func (*ResourceImportResult) writeTo(resources []genruntime.MetaObject, destination io.Writer) error {
	buf := bufio.NewWriter(destination)
	defer func(buf *bufio.Writer) {
		_ = buf.Flush()
	}(buf)

	_, err := buf.WriteString("---\n")
	if err != nil {
		return errors.Wrap(err, "unable to save to writer")
	}

	// Sort objects into a deterministic order
	slices.SortFunc(resources, func(left genruntime.MetaObject, right genruntime.MetaObject) bool {
		leftGVK := left.GetObjectKind().GroupVersionKind()
		rightGVK := right.GetObjectKind().GroupVersionKind()

		if leftGVK.Group != rightGVK.Group {
			return leftGVK.Group < rightGVK.Group
		}

		if leftGVK.Version != rightGVK.Version {
			return leftGVK.Version < rightGVK.Version
		}

		if leftGVK.Kind != rightGVK.Kind {
			return leftGVK.Kind < rightGVK.Kind
		}

		return left.GetName() < right.GetName()
	})

	for _, resource := range resources {
		data, err := yaml.Marshal(resource)
		if err != nil {
			return errors.Wrap(err, "unable to save to writer")
		}

		_, err = buf.Write(data)
		if err != nil {
			return errors.Wrap(err, "unable to save to writer")
		}

		_, err = buf.WriteString("---\n")
		if err != nil {
			return errors.Wrap(err, "unable to save to writer")
		}
	}

	return nil
}

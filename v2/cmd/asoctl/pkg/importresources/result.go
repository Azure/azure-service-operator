/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importresources

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/rotisserie/eris"
	"golang.org/x/exp/slices"
	"sigs.k8s.io/yaml"

	"github.com/Azure/azure-service-operator/v2/internal/annotations"
	"github.com/Azure/azure-service-operator/v2/internal/labels"
)

// Result represents the result of an import operation
type Result struct {
	imported []ImportedResource
}

// Count returns the number of successfully imported resources.
func (r *Result) Count() int {
	return len(r.imported)
}

func (r *Result) SaveToWriter(destination io.Writer) error {
	return r.writeTo(r.imported, destination)
}

func (r *Result) SaveToSingleFile(filepath string) error {
	return r.saveTo(r.imported, filepath)
}

// AddAnnotations adds the given annotations to all the resources
func (r *Result) AddAnnotations(toAdd []string) error {
	// pre-parse the annotations
	parsed, err := annotations.ParseAll(toAdd)
	if err != nil {
		return err
	}

	for _, imported := range r.imported {
		resource := imported.Resource()
		resourceAnnotations := resource.GetAnnotations()
		if resourceAnnotations == nil {
			resourceAnnotations = make(map[string]string, len(toAdd))
		}
		for _, annotation := range parsed {
			resourceAnnotations[annotation.Key] = annotation.Value
		}
		resource.SetAnnotations(resourceAnnotations)
	}

	return nil
}

// AddLabels adds the given labels to all the resources
func (r *Result) AddLabels(toAdd []string) error {
	// pre-parse the labels
	parsed, err := labels.ParseAll(toAdd)
	if err != nil {
		return err
	}

	for _, imported := range r.imported {
		resource := imported.Resource()
		resourceLabels := resource.GetLabels()
		if resourceLabels == nil {
			resourceLabels = make(map[string]string, len(toAdd))
		}
		for _, label := range parsed {
			resourceLabels[label.Key] = label.Value
		}
		resource.SetLabels(resourceLabels)
	}

	return nil
}

// SetNamespace sets the namespace for all the resources
func (r *Result) SetNamespace(namespace string) {
	for _, imported := range r.imported {
		resource := imported.Resource()
		resource.SetNamespace(namespace)
	}
}

func (r *Result) SaveToIndividualFilesInFolder(folder string) error {
	// We name the files after the resource type and name
	// We allocate resources to files using a map, just in case we have a naming collision
	// (If that happens, all the similarly named resources will be in the same file, which is not ideal,
	// but better than dropping one or more)
	fileMap := make(map[string][]ImportedResource, len(r.imported))
	for _, imported := range r.imported {
		resource := imported.Resource()
		resourceName := resource.GetName()
		typeName := resource.GetObjectKind().GroupVersionKind().Kind
		fileName := fmt.Sprintf("%s-%s.yaml", typeName, resourceName)
		fileMap[fileName] = append(fileMap[fileName], imported)
	}

	for fileName, resources := range fileMap {
		path := filepath.Join(folder, fileName)
		err := r.saveTo(resources, path)
		if err != nil {
			return eris.Wrapf(err, "unable to save to file %s", path)
		}
	}

	return nil
}

func (r *Result) saveTo(resources []ImportedResource, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return eris.Wrapf(err, "unable to create file %s", path)
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

	return eris.Wrapf(err, "unable to save to file %s", path)
}

func (*Result) writeTo(resources []ImportedResource, destination io.Writer) error {
	buf := bufio.NewWriter(destination)
	defer func(buf *bufio.Writer) {
		_ = buf.Flush()
	}(buf)

	_, err := buf.WriteString("---\n")
	if err != nil {
		return eris.Wrap(err, "unable to save to writer")
	}

	// Sort objects into a deterministic order
	slices.SortFunc(
		resources,
		func(left ImportedResource, right ImportedResource) int {
			leftGVK := left.Resource().GetObjectKind().GroupVersionKind()
			rightGVK := right.Resource().GetObjectKind().GroupVersionKind()

			if leftGVK.Group < rightGVK.Group {
				return -1
			} else if leftGVK.Group > rightGVK.Group {
				return 1
			}

			if leftGVK.Version < rightGVK.Version {
				return -1
			} else if leftGVK.Version > rightGVK.Version {
				return 1
			}

			if leftGVK.Kind < rightGVK.Kind {
				return -1
			} else if leftGVK.Kind > rightGVK.Kind {
				return 1
			}

			return strings.Compare(left.Name(), right.Name())
		})

	for _, resource := range resources {
		data, err := yaml.Marshal(resource.Resource())
		if err != nil {
			return eris.Wrap(err, "unable to save to writer")
		}

		data = redact(data)

		_, err = buf.Write(data)
		if err != nil {
			return eris.Wrap(err, "unable to save to writer")
		}

		_, err = buf.WriteString("---\n")
		if err != nil {
			return eris.Wrap(err, "unable to save to writer")
		}
	}

	return nil
}

// redact removes any selected information that shouldn't be included,
// starting with empty `status { }` blocks from the yaml.
func redact(data []byte) []byte {
	content := string(data)
	content = strings.ReplaceAll(content, "status: {}", "")
	content = strings.TrimSuffix(content, "\n")
	return []byte(content)
}

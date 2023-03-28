/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"bufio"
	"io"
	"os"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
	"sigs.k8s.io/yaml"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// ResourceImportResult represents the result of an import operation
type ResourceImportResult struct {
	resources []genruntime.MetaObject
}

func (r *ResourceImportResult) SaveToWriter(destination io.Writer) error {
	buf := bufio.NewWriter(destination)
	defer func(buf *bufio.Writer) {
		_ = buf.Flush()
	}(buf)

	_, err := buf.WriteString("---\n")
	if err != nil {
		return errors.Wrap(err, "unable to save to writer")
	}

	// Sort objects into a deterministic order
	slices.SortFunc(r.resources, func(left genruntime.MetaObject, right genruntime.MetaObject) bool {
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

	for _, resource := range r.resources {
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

func (r *ResourceImportResult) SaveToFile(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return errors.Wrapf(err, "unable to create file %s", filepath)
	}

	defer func() {
		file.Close()

		// if we are panicking, the file will be in a broken
		// state, so remove it
		if r := recover(); r != nil {
			os.Remove(filepath)
			panic(r)
		}
	}()

	err = r.SaveToWriter(file)
	if err != nil {
		// cleanup in case of errors
		file.Close()
		os.Remove(filepath)
	}

	return errors.Wrapf(err, "unable to save to file %s", filepath)
}

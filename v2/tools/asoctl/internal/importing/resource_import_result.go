/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"bufio"
	"io"
	"os"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/pkg/errors"
	"sigs.k8s.io/yaml"
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

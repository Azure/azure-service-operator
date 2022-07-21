/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/runtime"
	"os"
	"path"
)

// ExportAsSample exports the provided resource to a YAML file as the basis for a sample
// This is used to generate the samples for the repo
func ExportAsSample(tc *testcommon.KubePerTestContext, resource runtime.Object) {
	tc.T.Helper()

	copy := resource.DeepCopyObject()
	gvk := copy.GetObjectKind().GroupVersionKind()
	filename := fmt.Sprintf("%s_%s_%s.yaml", gvk.Kind, gvk.Group, gvk.Version)
	filepath := path.Join(os.TempDir(), filename)

	err := ExportAsYAML(copy, filepath)
	if err != nil {
		tc.T.Fatalf("failed to export resource: %s", err)
	}

	tc.T.Logf("Exported resource to %s", filepath)
}

// ExportAsYAML exports the provided resource to a YAML file
func ExportAsYAML(resource runtime.Object, filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return errors.Wrapf(err, "failed to open file %s", filename)
	}

	defer file.Close()

	enc := yaml.NewEncoder(file)
	err = enc.Encode(resource)
	if err != nil {
		return errors.Wrapf(err, "failed to encode resource for writing to %s", filename)
	}

	return nil
}

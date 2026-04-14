/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"fmt"
	"os"
	"strings"

	"github.com/Azure/azure-service-operator/v2/internal/crdmanagement"
)

// CheckBundledCRDsDirectory checks that the given path looks like a valid bundled CRDs directory,
// i.e. that the task target `bundle-crds` has been run. It verifies the directory exists and
// contains a reasonable number of CRD YAML files with the expected naming convention.
func CheckBundledCRDsDirectory(crdPath string) error {
	entries, err := os.ReadDir(crdPath)
	if err != nil {
		return fmt.Errorf(
			"CRD path %q does not exist or is not readable (has the task target `bundle-crds` been run?): %w",
			crdPath,
			err)
	}

	crdCount := 0
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.HasPrefix(entry.Name(), crdmanagement.CRDFilePrefix) && strings.HasSuffix(entry.Name(), ".yaml") {
			crdCount++
		}
	}

	if crdCount < 10 {
		return fmt.Errorf(
			"CRD path %q contains only %d CRD files; expected many more (has the task target `bundle-crds` been run?)",
			crdPath,
			crdCount)
	}

	return nil
}

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package versions

import (
	"strings"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// Compare compares two versions, returning true if left is less than right.
func Compare(left string, right string) bool {
	return astmodel.ComparePathAndVersion(left, right)
}

// IsPreview returns true if the version is a preview version, false otherwise
func IsPreview(version string) bool {
	// Don't want to get confused by the `v1beta` prefix, so we strip it off
	v := version
	if strings.HasPrefix(v, astmodel.GeneratorVersion) {
		v = v[len(astmodel.GeneratorVersion):]
	}

	return astmodel.ContainsPreviewVersionLabel(v)
}

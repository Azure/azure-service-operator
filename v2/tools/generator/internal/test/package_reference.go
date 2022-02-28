/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package test

import (
	"strings"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

var GoModulePrefix = "github.com/Azure/azure-service-operator/testing"

// MakeLocalPackageReference makes a local package reference for testing purposes
func MakeLocalPackageReference(group string, version string) astmodel.LocalPackageReference {
	// For test purposes, we use a constant custom prefix so that our tests don't need updating when when change
	// our active version prefix.
	// For convenience, we tolerate the prefix already being present
	if strings.HasPrefix(version, "v") {
		version = version[1:]
	}

	return astmodel.MakeLocalPackageReference(GoModulePrefix, group, "v", version)
}

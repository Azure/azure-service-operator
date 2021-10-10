/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package test

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

var GoModulePrefix = "github.com/Azure/azure-service-operator/testing"

// MakeLocalPackageReference makes a local package reference for testing purposes
func MakeLocalPackageReference(group string, version string) astmodel.LocalPackageReference {
	return astmodel.MakeLocalPackageReference(GoModulePrefix, group, version)
}

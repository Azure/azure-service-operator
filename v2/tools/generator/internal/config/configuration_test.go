/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config_test

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// TODO: in a common test function rather than in a bunch of test modules?
func makeTestLocalPackageReference(group string, version string) astmodel.LocalPackageReference {
	return astmodel.MakeLocalPackageReference("github.com/Azure/azure-service-operator/v2", group, version)
}

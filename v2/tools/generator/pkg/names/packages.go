/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package names

import (
	"strings"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

func IsStorageVersion(version string) bool {
	return strings.HasSuffix(version, astmodel.StoragePackageSuffix)
}

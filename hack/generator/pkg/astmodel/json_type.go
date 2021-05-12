/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

var (
	// APIExtensionsPackage contains the type we use to represent
	// arbitrary JSON fields.
	APIExtensionsPackage = MakeExternalPackageReference("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1")

	// JSONType is used for fields that need to store arbitrary JSON
	// structures.
	JSONType = MakeTypeName(APIExtensionsPackage, "JSON")
)

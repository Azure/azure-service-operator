/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"
	"strings"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// NewGetAPIVersionFunction returns a function that returns a static API version string
func NewGetAPIVersionFunction(
	apiVersionTypeName astmodel.TypeName,
	apiVersionEnumValue astmodel.EnumValue,
	idFactory astmodel.IdentifierFactory) astmodel.Function {

	comment := fmt.Sprintf("returns the ARM API version of the resource. This is always %q", strings.Trim(apiVersionEnumValue.Value, "\""))
	result := NewObjectFunction("Get"+astmodel.APIVersionProperty, idFactory, newStaticStringReturnFunctionBody(apiVersionEnumValue.Value, comment, ReceiverTypeStruct)) // TODO: We should use the enum ID here
	result.AddPackageReference(astmodel.GenRuntimeReference)
	result.AddReferencedTypes(apiVersionTypeName)

	return result
}

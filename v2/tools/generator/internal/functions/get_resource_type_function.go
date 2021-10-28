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

type ReceiverType string

const (
	ReceiverTypePtr    = ReceiverType("ptr")
	ReceiverTypeStruct = ReceiverType("struct")
)

// NewGetTypeFunction returns a function "GetType" that returns a static string value for the ResourceType of a resource
func NewGetTypeFunction(
	armType string,
	idFactory astmodel.IdentifierFactory,
	receiverType ReceiverType) astmodel.Function {

	// Trim any "'s around armType
	armType = strings.Trim(armType, "\"")

	comment := fmt.Sprintf("returns the ARM Type of the resource. This is always %q", armType)
	result := NewObjectFunction("Get"+astmodel.TypeProperty, idFactory, createBodyReturningLiteralString(armType, comment, receiverType))
	result.AddPackageReference(astmodel.GenRuntimeReference)
	return result
}

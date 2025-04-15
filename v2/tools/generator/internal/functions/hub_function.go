/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// NewHubFunction creates an empty Hub() function that satisfies the Hub interface required by the controller
// See https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/conversion#Hub
func NewHubFunction(idFactory astmodel.IdentifierFactory) astmodel.Function {
	result := NewObjectFunction(
		"Hub",
		idFactory,
		createHubFunctionBody)
	return result
}

func createHubFunctionBody(
	fn *ObjectFunction,
	genContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	// Create a sensible name for our receiver
	receiverName := fn.IDFactory().CreateReceiver(receiver.Name())

	// We always use a pointer receiver
	receiverType := astmodel.NewOptionalType(receiver)
	receiverTypeExpr, err := receiverType.AsTypeExpr(genContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	details := astbuilder.FuncDetails{
		ReceiverIdent: receiverName,
		ReceiverType:  receiverTypeExpr,
		Name:          methodName,
		Body:          []dst.Stmt{}, // empty body
	}

	details.AddComments(fmt.Sprintf("marks that this %s is the hub type for conversion", receiver.Name()))

	return details.DefineFunc(), nil
}

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// HubFunction generates an empty Hub() function that satisfies the Hub interface required by the controller\
// See https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/conversion#Hub
type HubFunction struct {
	idFactory astmodel.IdentifierFactory
}

// Ensure HubFunction properly implements Function
var _ astmodel.Function = &HubFunction{}

// NewHubFunction creates a new instance
func NewHubFunction(idFactory astmodel.IdentifierFactory) *HubFunction {
	return &HubFunction{
		idFactory: idFactory,
	}
}

// Name returns the hard coded name of the function
func (h HubFunction) Name() string {
	return "Hub"
}

// RequiredPackageReferences indicates that this function has no required packages
func (h HubFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return astmodel.NewPackageReferenceSet()
}

// References indicates that this function references no types
func (h HubFunction) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet()
}

// AsFunc generates the required code
func (h HubFunction) AsFunc(generationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *dst.FuncDecl {
	// Create a sensible name for our receiver
	receiverName := h.idFactory.CreateIdentifier(receiver.Name(), astmodel.NotExported)

	// We always use a pointer receiver so we can modify it
	receiverType := astmodel.NewOptionalType(receiver).AsType(generationContext)

	details := astbuilder.FuncDetails{
		ReceiverIdent: receiverName,
		ReceiverType:  receiverType,
		Name:          h.Name(),
		Body:          []dst.Stmt{}, // empty body
	}

	details.AddComments(fmt.Sprintf("marks that this %s is the hub type for conversion", receiver.Name()))

	return details.DefineFunc()
}

// Equals shows that any hub function is equal to any other
func (h HubFunction) Equals(f astmodel.Function) bool {
	_, ok := f.(*HubFunction)
	return ok
}

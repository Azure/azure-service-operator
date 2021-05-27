/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
)

// HubFunction generates an empty Hub() function that satisfies the Hub interface required by the controller
type HubFunction struct {
	idFactory IdentifierFactory
}

// Ensure HubFunction properly implements Function
var _ Function = &HubFunction{}

// NewHubFunction creates a new instance
func NewHubFunction(idFactory IdentifierFactory) *HubFunction {
	return &HubFunction{
		idFactory: idFactory,
	}
}

// Name returns the hard coded name of the function
func (h HubFunction) Name() string {
	return "Hub"
}

// RequiredPackageReferences indicates that this function has no required packages
func (h HubFunction) RequiredPackageReferences() *PackageReferenceSet {
	return NewPackageReferenceSet()
}

// References indicates that this function references no types
func (h HubFunction) References() TypeNameSet {
	return NewTypeNameSet()
}

// AsFunc generates the required code
func (h HubFunction) AsFunc(generationContext *CodeGenerationContext, receiver TypeName) *dst.FuncDecl {
	// Create a sensible name for our receiver
	receiverName := h.idFactory.CreateIdentifier(receiver.Name(), NotExported)

	// We always use a pointer receiver so we can modify it
	receiverType := NewOptionalType(receiver).AsType(generationContext)

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
func (h HubFunction) Equals(f Function) bool {
	_, ok := f.(*HubFunction)
	return ok
}

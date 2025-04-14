/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"
	"strings"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// PropertyExporter is a helper that helps build functions that just export a struct property.
// For example a function GetFoo that just returned s.foo.
type PropertyExporter struct {
	resourceName astmodel.InternalTypeName
	resource     *astmodel.ResourceType
	idFactory    astmodel.IdentifierFactory

	nilChecks     [][]string
	propertyPath  []string
	functionName  string
	exportType    astmodel.Type
	interfaceName astmodel.ExternalTypeName
}

func newPropertyExporterInterface(
	resourceName astmodel.InternalTypeName,
	resource *astmodel.ResourceType,
	idFactory astmodel.IdentifierFactory,
	nilChecks [][]string,
	propertyPath []string,
	functionName string,
	exportType astmodel.Type,
	interfaceName astmodel.ExternalTypeName,
) *PropertyExporter {
	return &PropertyExporter{
		resourceName: resourceName,
		resource:     resource,
		idFactory:    idFactory,

		nilChecks:     nilChecks,
		propertyPath:  propertyPath,
		functionName:  functionName,
		exportType:    exportType,
		interfaceName: interfaceName,
	}
}

func NewConfigMapExporterInterface(
	resourceName astmodel.InternalTypeName,
	resource *astmodel.ResourceType,
	idFactory astmodel.IdentifierFactory,
) *PropertyExporter {
	return newPropertyExporterInterface(
		resourceName,
		resource,
		idFactory,
		[][]string{{"Spec", astmodel.OperatorSpecProperty}},
		[]string{"Spec", astmodel.OperatorSpecProperty, astmodel.OperatorSpecConfigMapExpressionsProperty},
		"ConfigMapDestinationExpressions",
		astmodel.DestinationExpressionCollectionType,
		astmodel.ConfigMapExporterType)
}

func NewSecretsExporterInterface(
	resourceName astmodel.InternalTypeName,
	resource *astmodel.ResourceType,
	idFactory astmodel.IdentifierFactory,
) *PropertyExporter {
	return newPropertyExporterInterface(
		resourceName,
		resource,
		idFactory,
		[][]string{{"Spec", astmodel.OperatorSpecProperty}},
		[]string{"Spec", astmodel.OperatorSpecProperty, astmodel.OperatorSpecSecretExpressionsProperty},
		"SecretDestinationExpressions",
		astmodel.DestinationExpressionCollectionType,
		astmodel.SecretExporterType)
}

func (d *PropertyExporter) ToInterfaceImplementation() *astmodel.InterfaceImplementation {
	funcs := []astmodel.Function{
		NewResourceFunction(
			d.functionName,
			d.resource,
			d.idFactory,
			d.getPropertyFunction),
	}

	return astmodel.NewInterfaceImplementation(
		d.interfaceName,
		funcs...)
}

// getPropertyFunction returns a function declaration for getting the property.
//
//	func (r *<receiver>) <funcName>() <exportType> {
//		<nilChecks>
//
//	    return <propertyPath>
//	}
func (d *PropertyExporter) getPropertyFunction(
	f *ResourceFunction,
	genContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	_ string,
) (*dst.FuncDecl, error) {
	receiverIdent := f.IDFactory().CreateReceiver(receiver.Name())
	receiverType := astmodel.NewOptionalType(receiver)
	receiverTypeExpr, err := receiverType.AsTypeExpr(genContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	nilChecks := make([]dst.Stmt, 0, len(d.nilChecks))
	for _, check := range d.nilChecks {
		nilChecks = append(
			nilChecks,
			astbuilder.ReturnIfNil(
				astbuilder.Selector(
					dst.NewIdent(receiverIdent), check...), astbuilder.Nil(),
			),
		)
	}

	ret := astbuilder.Returns(astbuilder.Selector(dst.NewIdent(receiverIdent), d.propertyPath...))

	fn := &astbuilder.FuncDetails{
		ReceiverIdent: receiverIdent,
		ReceiverType:  receiverTypeExpr,
		Name:          d.functionName,
		Body: astbuilder.Statements(
			nilChecks,
			ret),
	}

	exportTypeExpr, err := d.exportType.AsTypeExpr(genContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating return type expression")
	}

	fn.AddReturn(exportTypeExpr)
	fn.AddComments(fmt.Sprintf("returns the %s property", strings.Join(d.propertyPath, ".")))

	return fn.DefineFunc(), nil
}

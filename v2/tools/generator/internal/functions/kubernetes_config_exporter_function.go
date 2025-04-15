/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"go/token"
	"sort"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// KubernetesConfigExporterBuilder is a builder for creating genruntime.KubernetesConfigExporter interface implementations.
type KubernetesConfigExporterBuilder struct {
	resourceName astmodel.TypeName
	resource     *astmodel.ResourceType
	idFactory    astmodel.IdentifierFactory

	mappings map[string][]*astmodel.PropertyDefinition
}

// NewKubernetesConfigExporterBuilder creates a new KubernetesConfigExporterBuilder for the specified resource
func NewKubernetesConfigExporterBuilder(
	resourceName astmodel.TypeName,
	resource *astmodel.ResourceType,
	idFactory astmodel.IdentifierFactory,
	mappings map[string][]*astmodel.PropertyDefinition,
) *KubernetesConfigExporterBuilder {
	return &KubernetesConfigExporterBuilder{
		resourceName: resourceName,
		resource:     resource,
		idFactory:    idFactory,
		mappings:     mappings,
	}
}

// ToInterfaceImplementation creates an InterfaceImplementation from the KubernetesConfigExporterBuilder
func (d *KubernetesConfigExporterBuilder) ToInterfaceImplementation() *astmodel.InterfaceImplementation {
	funcs := []astmodel.Function{
		NewResourceFunction(
			"ExportKubernetesConfigMaps",
			d.resource,
			d.idFactory,
			d.exportKubernetesConfigMaps,
			astmodel.GenRuntimeReference,
			astmodel.GenRuntimeConfigMapsReference,
			astmodel.GenericARMClientReference,
			astmodel.LogrReference,
			astmodel.ControllerRuntimeClient,
			astmodel.ContextReference),
	}
	return astmodel.NewInterfaceImplementation(
		astmodel.KuberentesConfigExporterType,
		funcs...)
}

// exportKubernetesResource returns the body of the ExportKubernetesResources function, which implements
// the genruntime.KubernetesExporter interface.
// Generates code like:
//
//	func (<receiver> *<receiverType>) ExportKubernetesConfigMaps(ctx context.Context, obj MetaObject, armClient *genericarmclient.GenericClient, log logr.Logger) ([]client.Object, error) {
//		collector := configmaps.NewCollector(<receiver>.Namespace)
//		if <receiver>.Spec.OperatorSpec != nil && <receiver>.Spec.OperatorSpec.ConfigMaps != nil {
//			if <receiver>.<propertyPath> != nil {
//				collector.AddValue(<receiver>.Spec.OperatorSpec.ConfigMaps.<configMapProperty>, *<receiver>.<propertyPath>)
//			}
//		}
//		...
//		result, err := collector.Values()
//		if err != nil {
//			return nil, err
//		}
//		return configmaps.SliceToClientObjectSlice(result), nil
//	}
func (d *KubernetesConfigExporterBuilder) exportKubernetesConfigMaps(
	k *ResourceFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string,
) (*dst.FuncDecl, error) {
	receiverIdent := k.IDFactory().CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating receiver type expression")
	}

	configMapsReference := codeGenerationContext.MustGetImportedPackageName(astmodel.GenRuntimeConfigMapsReference)

	// collector := configmaps.NewCollector(<receiver>.Namespace)
	collectorIdent := "collector"
	collectorCreationStmt := astbuilder.AssignmentStatement(
		dst.NewIdent(collectorIdent),
		token.DEFINE,
		astbuilder.CallQualifiedFunc(configMapsReference, "NewCollector", astbuilder.Selector(dst.NewIdent(receiverIdent), "Namespace")))

	operatorSpecSelector := astbuilder.Selector(dst.NewIdent(receiverIdent), "Spec", astmodel.OperatorSpecProperty)
	operatorSpecConfigMapsSelector := astbuilder.Selector(operatorSpecSelector, astmodel.OperatorSpecConfigMapsProperty)

	// Iterate through the mappings in alphabetical order to ensure consistent generation
	keys := make([]string, 0, len(d.mappings))
	for key := range d.mappings {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	collectStmts := make([]dst.Stmt, 0, len(keys))
	for _, operatorSpecPropertyName := range keys {
		propertyPath := d.mappings[operatorSpecPropertyName]

		var ifBlock dst.Stmt
		var propertyNames []string
		for _, prop := range propertyPath {
			propertyNames = append(propertyNames, prop.PropertyName().String())
		}

		for i := len(propertyPath) - 1; i >= 0; i -= 1 {
			propType := propertyPath[i].PropertyType()
			if _, ok := astmodel.AsMapType(propType); ok {
				return nil, eris.Errorf(
					"Exporting Map elements as configmaps is not supported currently. Property %s has type %s",
					propertyNames[i],
					propType.String())
			}

			if _, ok := astmodel.AsArrayType(propType); ok {
				return nil, eris.Errorf(
					"Exporting Slice elements as configmaps is not supported currently. Property %s has type %s",
					propertyNames[i],
					propType.String())
			}

			if _, ok := astmodel.AsOptionalType(propType); !ok {
				// If the property in question isn't actually optional, we don't need to check it so we can just continue
				continue
			}

			var inner dst.Stmt
			inner = ifBlock
			if ifBlock == nil {
				inner = d.addCollectorStmt(
					receiverIdent,
					collectorIdent,
					propertyPath,
					propertyNames,
					operatorSpecPropertyName)
			}

			ifBlock = astbuilder.IfNotNil(
				astbuilder.Selector(dst.NewIdent(receiverIdent), propertyNames[:i+1]...),
				inner)
		}

		//	if <receiver>.Spec.OperatorSpec != nil && <receiver>.Spec.OperatorSpec.ConfigMaps != nil {
		//		if <receiver>.<path> != nil {
		//			collector.AddValue(<receiver>.Spec.OperatorSpec.ConfigMaps.<operatorSpecPropertyName>, <receiver>.<path>)
		//		}
		//	}
		condition := astbuilder.JoinAnd(
			astbuilder.NotNil(operatorSpecSelector),
			astbuilder.NotNil(operatorSpecConfigMapsSelector))

		if ifBlock == nil {
			ifBlock = d.addCollectorStmt(
				receiverIdent,
				collectorIdent,
				propertyPath,
				propertyNames,
				operatorSpecPropertyName)
		}

		collectStmts = append(
			collectStmts,
			astbuilder.SimpleIf(condition, ifBlock))
	}

	//	result, err := collector.Values()
	//	if err != nil {
	//		return nil, err
	//	}
	//	return configmaps.SliceToClientObjectSlice(result), nil
	collectorValues := astbuilder.SimpleAssignmentWithErr(
		dst.NewIdent("result"),
		token.DEFINE,
		astbuilder.CallQualifiedFunc(collectorIdent, "Values"))
	returnIfErrNotNil := astbuilder.CheckErrorAndReturn(astbuilder.Nil())
	sliceToClientObjectSlice := astbuilder.Returns(
		astbuilder.CallQualifiedFunc(
			configMapsReference,
			"SliceToClientObjectSlice",
			dst.NewIdent("result")),
		astbuilder.Nil())

	fn := &astbuilder.FuncDetails{
		Name:          methodName,
		ReceiverIdent: receiverIdent,
		ReceiverType:  astbuilder.PointerTo(receiverExpr),
		Body: astbuilder.Statements(
			collectorCreationStmt,
			collectStmts,
			collectorValues,
			returnIfErrNotNil,
			sliceToClientObjectSlice),
	}

	contextTypeExpr, err := astmodel.ContextType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating context type expression")
	}
	fn.AddParameter("_", contextTypeExpr)

	metaObjectTypeExpr, err := astmodel.GenRuntimeMetaObjectType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating meta object type expression")
	}
	fn.AddParameter("_", metaObjectTypeExpr)

	clientTypeExpr, err := astmodel.NewOptionalType(astmodel.GenericClientType).AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating client type expression")
	}
	fn.AddParameter("_", clientTypeExpr)

	logrTypeExpr, err := astmodel.LogrType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating logr type expression")
	}
	fn.AddParameter("_", logrTypeExpr)

	objectArrayType, err := astmodel.NewArrayType(astmodel.ControllerRuntimeObjectType).
		AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating object array type expression")
	}

	fn.AddReturn(objectArrayType)
	fn.AddReturn(dst.NewIdent("error"))

	fn.AddComments("defines a resource which can create ConfigMaps in Kubernetes.")
	return fn.DefineFunc(), nil
}

func (d *KubernetesConfigExporterBuilder) addCollectorStmt(
	receiverIdent string,
	collectorIdent string,
	propertyPath []*astmodel.PropertyDefinition,
	propertyNames []string,
	operatorSpecPropertyName string,
) dst.Stmt {
	operatorSpecSelector := astbuilder.Selector(dst.NewIdent(receiverIdent), "Spec", astmodel.OperatorSpecProperty)
	operatorSpecConfigMapsSelector := astbuilder.Selector(operatorSpecSelector, astmodel.OperatorSpecConfigMapsProperty)

	var valueExpr dst.Expr = astbuilder.Selector(dst.NewIdent(receiverIdent), propertyNames...)
	if _, ok := astmodel.AsOptionalType(propertyPath[len(propertyPath)-1].PropertyType()); ok {
		valueExpr = astbuilder.Dereference(valueExpr)
	}
	return astbuilder.CallQualifiedFuncAsStmt(
		collectorIdent,
		"AddValue",
		astbuilder.Selector(operatorSpecConfigMapsSelector, d.idFactory.CreateIdentifier(operatorSpecPropertyName, astmodel.Exported)),
		valueExpr)
}

package conversions

import (
	"fmt"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/dave/dst"
)

// HubConversionFunction implements conversions to/from our hub type
// Existing PropertyAssignment functions are used to implement stepwise conversion
type HubConversionFunction struct {
	// name of this conversion function
	name string
	// hub is the TypeName of the canonical hub type, the final target or original source for conversion
	hub astmodel.TypeName
	// direction specifies whether we are converting to the hub type, or from it
	direction Direction
	// propertyFunction is the name of the function we call to copy properties across
	propertyFunctionName string
	// intermediateType is the TypeName of an intermediate type we use as part of a multiple step conversion
	// If nil, we are able to convert directly to/from the hub type
	intermediateType *astmodel.TypeName
}

// Ensure we properly implement the function interface
var _ astmodel.Function = &HubConversionFunction{}

// CreateConversionToHubFunction creates a conversion function that populates our hub type from the current instance
func CreateConversionToHubFunction(hub astmodel.TypeDefinition, idFactory astmodel.IdentifierFactory) *HubConversionFunction {
	name := "ConvertTo" + idFactory.CreateIdentifier(hub.Name().Name(), astmodel.Exported)
	propertyFunction := mustFindPropertyAssignmentFunction(hub, ConvertTo)

	result := &HubConversionFunction{
		name:                 name,
		hub:                  hub.Name(),
		direction:            ConvertTo,
		propertyFunctionName: propertyFunction.Name(),
	}

	if !hub.Equals(propertyFunction.otherDefinition) {
		intermediateType := propertyFunction.otherDefinition.Name()
		result.intermediateType = &intermediateType
	}

	return result
}

// CreateConversionFromHubFunction creates a conversion function that populates the current instance from our hub type
func CreateConversionFromHubFunction(hub astmodel.TypeDefinition, idFactory astmodel.IdentifierFactory) *HubConversionFunction {
	name := "ConvertFrom" + idFactory.CreateIdentifier(hub.Name().Name(), astmodel.Exported)
	propertyFunction := mustFindPropertyAssignmentFunction(hub, ConvertFrom)

	result := &HubConversionFunction{
		name:          name,
		hub: hub.Name(),
		direction:     ConvertFrom,
		propertyFunctionName: propertyFunction.Name(),
	}

	if !hub.Equals(propertyFunction.otherDefinition) {
		intermediateType := propertyFunction.otherDefinition.Name()
		result.intermediateType = &intermediateType
	}

	return result
}

func (fn *HubConversionFunction) Name() string {
	return fn.name
}

func (fn *HubConversionFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	panic("implement me")
}

func (fn *HubConversionFunction) References() astmodel.TypeNameSet {
	panic("implement me")
}

func (fn *HubConversionFunction) AsFunc(codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *dst.FuncDecl {
	panic("implement me")
}

func (fn *HubConversionFunction) Equals(otherFn astmodel.Function) bool {
	hcf, ok := otherFn.(*HubConversionFunction)
	if !ok {
		return false
	}

	//
	return fn.name == hcf.name &&
		fn.direction == hcf.direction &&
		fn.hub.Equals(hcf.hub)
}

// mustFindPropertyAssignmentFunction searches for a property assignment function on the supplied TypeDefinition and returns
// it, if found, or nil otherwise
func mustFindPropertyAssignmentFunction(
	other astmodel.TypeDefinition,
	direction Direction) *PropertyAssignmentFunction {
	var functionContainer astmodel.FunctionContainer
	if resource, ok := astmodel.AsResourceType(other.Type()); ok {
		functionContainer = resource
	} else if objectType, ok := astmodel.AsObjectType(other.Type()); ok {
		functionContainer = objectType
	}

	if functionContainer != nil {
		for _, f := range functionContainer.Functions() {
			propertyFn, ok := f.(*PropertyAssignmentFunction)
			if !ok {
				continue
			}

			if propertyFn.HasDirection(direction) {
				return propertyFn
			}
		}
	}

	panic(fmt.Sprintf("Failed to find a PropertyAssignmentFunction on %q", other.Name()))
}

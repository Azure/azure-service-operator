/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const MakeStatusPropertiesOptionalStageID = "makeStatusPropertiesOptional"

// MakeStatusPropertiesOptional makes all top level Status properties optional. This is required because Status itself
// is actually optional (it's not set initially) and if any top level properties are not optional they end up
// always being returned, which makes for an awkward GET response initially (before the Status has been set). It also
// has implications for updating or patching the CRD resource as patching something in the spec or changing an annotation
// will deserialize the response from apiserver into the object passed in. If there are any spurious empty properties in Status
// included they will end up getting overwritten (possibly before the client.Status().Update() call can be made).
func MakeStatusPropertiesOptional() *Stage {
	return NewStage(
		MakeStatusPropertiesOptionalStageID,
		"Force all status properties to be optional",
		func(ctx context.Context, state *State) (*State, error) {
			statusDefs := astmodel.FindStatusDefinitions(state.Definitions())
			var errs []error

			result := make(astmodel.TypeDefinitionSet)
			for _, def := range statusDefs {
				modifiedType, err := makeStatusPropertiesOptional(def)
				if err != nil {
					errs = append(errs, err)
				}

				result.Add(def.WithType(modifiedType))
			}

			err := kerrors.NewAggregate(errs)
			if err != nil {
				return nil, err
			}

			remaining := state.Definitions().Except(result)
			result.AddTypes(remaining)

			return state.WithDefinitions(result), nil
		})
}

// makeStatusPropertiesOptional makes all properties optional on top level Status types
func makeStatusPropertiesOptional(statusDef astmodel.TypeDefinition) (astmodel.Type, error) {
	visitor := astmodel.TypeVisitorBuilder{
		VisitObjectType: makeObjectPropertiesOptional,
	}.Build()

	return visitor.Visit(statusDef.Type(), statusDef.Name())
}

// makeObjectPropertiesOptional makes properties optional for the object
func makeObjectPropertiesOptional(this *astmodel.TypeVisitor, ot *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
	typeName := ctx.(astmodel.TypeName)
	for _, property := range ot.Properties() {
		if property.HasKubebuilderRequiredValidation() {
			klog.V(4).Infof("\"%s.%s\" was required, changing it to optional", typeName.String(), property.PropertyName())
		}
		ot = ot.WithProperty(property.MakeOptional())
	}

	return astmodel.IdentityVisitOfObjectType(this, ot, ctx)
}

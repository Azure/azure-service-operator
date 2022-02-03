/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// AddCrossResourceReferencesStageID is the unique identifier for this pipeline stage
const AddCrossResourceReferencesStageID = "addCrossResourceReferences"

var armIDDescriptionRegex = regexp.MustCompile("(?i)(.*/subscriptions/.*?/resourceGroups/.*|ARM ID|Resource ID|resourceId)")

// TODO: For now not supporting array or map of references. Unsure if it actually ever happens in practice.

// AddCrossResourceReferences replaces cross resource references with genruntime.ResourceReference.
func AddCrossResourceReferences(configuration *config.Configuration, idFactory astmodel.IdentifierFactory) Stage {
	return MakeLegacyStage(
		AddCrossResourceReferencesStageID,
		"Replace cross-resource references with genruntime.ResourceReference",
		func(ctx context.Context, definitions astmodel.Types) (astmodel.Types, error) {
			result := make(astmodel.Types)

			var crossResourceReferenceErrs []error

			isCrossResourceReference := func(typeName astmodel.TypeName, prop *astmodel.PropertyDefinition) bool {
				isReference, err := configuration.ARMReference(typeName, prop.PropertyName())
				if DoesPropertyLookLikeARMReference(prop) && err != nil {
					if config.IsNotConfiguredError(err) {
						// This is an error for now to ensure that we don't accidentally miss adding references.
						// If/when we move to using an upstream marker for cross resource refs, we can remove this and just
						// trust the Swagger.
						crossResourceReferenceErrs = append(
							crossResourceReferenceErrs,
							errors.Wrapf(
								err,
								"%s.%s looks like a resource reference but was not labelled as one; You may need to add it to the 'objectModelConfiguration' section of the config file",
								typeName,
								prop.PropertyName()))
					} else {
						// Something else went wrong checking our configuration
						crossResourceReferenceErrs = append(
							crossResourceReferenceErrs,
							errors.Wrapf(
								err,
								"%s.%s looks like a resource reference but something went wrong checking for configuration",
								typeName,
								prop.PropertyName()))
					}
				}

				return isReference
			}

			visitor := MakeCrossResourceReferenceTypeVisitor(idFactory, isCrossResourceReference)
			for _, def := range definitions {
				// Skip Status types
				// TODO: we need flags
				if strings.Contains(def.Name().Name(), "_Status") {
					result.Add(def)
					continue
				}

				t, err := visitor.Visit(def.Type(), def.Name())
				if err != nil {
					return nil, errors.Wrapf(err, "visiting %q", def.Name())
				}
				result.Add(def.WithType(t))

				// TODO: Remove types that have only a single field ID and pull things up a level? Will need to wait for George's
				// TODO: Properties collapsing work for this.
			}

			var err error = kerrors.NewAggregate(crossResourceReferenceErrs)
			if err != nil {
				return nil, err
			}

			err = configuration.VerifyARMReferencesConsumed()
			if err != nil {
				klog.Error(err)

				return nil, errors.Wrap(
					err,
					"Found unused $armReference configurations; these need to be fixed or removed.")
			}

			return result, nil
		})
}

type crossResourceReferenceChecker func(typeName astmodel.TypeName, prop *astmodel.PropertyDefinition) bool

type CrossResourceReferenceTypeVisitor struct {
	astmodel.TypeVisitor
	// referenceChecker is a function describing what a cross resource reference looks like. It is overridable so that
	// we can use a more simplistic criteria for tests.
	isPropertyAnARMReference crossResourceReferenceChecker
}

func MakeCrossResourceReferenceTypeVisitor(idFactory astmodel.IdentifierFactory, referenceChecker crossResourceReferenceChecker) CrossResourceReferenceTypeVisitor {
	visitor := CrossResourceReferenceTypeVisitor{
		isPropertyAnARMReference: referenceChecker,
	}

	transformResourceReferenceProperties := func(_ *astmodel.TypeVisitor, it *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
		typeName := ctx.(astmodel.TypeName)

		var newProps []*astmodel.PropertyDefinition
		for _, prop := range it.Properties() {
			if visitor.isPropertyAnARMReference(typeName, prop) {
				klog.V(4).Infof("Transforming \"%s.%s\" field into genruntime.ResourceReference", typeName, prop.PropertyName())
				originalName := string(prop.PropertyName())
				prop = makeResourceReferenceProperty(idFactory, prop)

				// TODO: We could pass this information forward some other way?
				// Add tag so that we remember what this field was before
				prop = prop.WithTag(astmodel.ARMReferenceTag, originalName)
			}

			newProps = append(newProps, prop)
		}

		it = it.WithoutProperties()
		result := it.WithProperties(newProps...)
		return result, nil
	}

	visitor.TypeVisitor = astmodel.TypeVisitorBuilder{
		VisitObjectType: transformResourceReferenceProperties,
	}.Build()

	return visitor
}

// DoesPropertyLookLikeARMReference uses a simple heuristic to determine if a property looks like it might be an ARM reference.
// This can be used for logging/reporting purposes to discover references which we missed.
func DoesPropertyLookLikeARMReference(prop *astmodel.PropertyDefinition) bool {
	// The property must be a string or optional string
	isString := astmodel.TypeEquals(prop.PropertyType(), astmodel.StringType)
	isOptionalString := astmodel.TypeEquals(prop.PropertyType(), astmodel.NewOptionalType(astmodel.StringType))
	if !isString && !isOptionalString {
		return false
	}

	hasMatchingDescription := armIDDescriptionRegex.MatchString(prop.Description())
	namedID := prop.HasName("Id")
	if hasMatchingDescription || namedID {
		return true
	}

	return false
}

func makeResourceReferenceProperty(idFactory astmodel.IdentifierFactory, existing *astmodel.PropertyDefinition) *astmodel.PropertyDefinition {
	var referencePropertyName string
	// Special case for "Id" and properties that end in "Id", which are quite common in the specs. This is primarily
	// because it's awkward to have a field called "Id" not just be a string and instead but a complex type describing
	// a reference.
	if existing.PropertyName() == "Id" {
		referencePropertyName = "Reference"
	} else if strings.HasSuffix(string(existing.PropertyName()), "Id") {
		referencePropertyName = strings.TrimSuffix(string(existing.PropertyName()), "Id") + "Reference"
	} else {
		referencePropertyName = string(existing.PropertyName()) + "Reference"
	}

	newProp := astmodel.NewPropertyDefinition(
		idFactory.CreatePropertyName(referencePropertyName, astmodel.Exported),
		idFactory.CreateIdentifier(referencePropertyName, astmodel.NotExported),
		astmodel.ResourceReferenceType)

	newProp = newProp.WithDescription(existing.Description())
	if existing.HasKubebuilderRequiredValidation() {
		newProp = newProp.MakeRequired()
	} else {
		newProp = newProp.MakeOptional()
	}

	return newProp
}

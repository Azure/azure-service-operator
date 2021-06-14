/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/config"
)

var armIDDescriptionRegex = regexp.MustCompile("(?i).*/subscriptions/.*?/resourceGroups/.*")

// TODO: For now not supporting array or map of references. Unsure if it actually ever happens in practice.

// addCrossResourceReferences replaces cross resource references with genruntime.ResourceReference.
func addCrossResourceReferences(configuration *config.Configuration, idFactory astmodel.IdentifierFactory) PipelineStage {
	return MakePipelineStage(
		"addCrossResourceReferences",
		"Replaces cross resource references with genruntime.ResourceReference",
		func(ctx context.Context, definitions astmodel.Types) (astmodel.Types, error) {

			result := make(astmodel.Types)
			knownReferences := newKnownReferencesMap(configuration)
			isCrossResourceReference := func(typeName astmodel.TypeName, prop *astmodel.PropertyDefinition) bool {
				ref := referencePair{
					typeName: typeName,
					propName: prop.PropertyName(),
				}
				_, isReference := knownReferences[ref]

				if doesPropertyLookLikeARMReference(prop) && !isReference {
					klog.V(0).Infof("\"%s.%s\" looks like a resource reference but was not labelled as one", typeName, prop.PropertyName())
				}

				return isReference
			}

			visitor := makeCrossResourceReferenceTypeVisitor(idFactory, isCrossResourceReference)

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

			return result, nil
		})
}

type referencePair struct {
	typeName astmodel.TypeName
	propName astmodel.PropertyName
}

type crossResourceReferenceChecker func(typeName astmodel.TypeName, prop *astmodel.PropertyDefinition) bool

type crossResourceReferenceTypeVisitor struct {
	astmodel.TypeVisitor
	// referenceChecker is a function describing what a cross resource reference looks like. It is overridable so that
	// we can use a more simplistic criteria for tests.
	isPropertyAnARMReference crossResourceReferenceChecker
}

func makeCrossResourceReferenceTypeVisitor(idFactory astmodel.IdentifierFactory, referenceChecker crossResourceReferenceChecker) crossResourceReferenceTypeVisitor {
	visitor := crossResourceReferenceTypeVisitor{
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

// doesPropertyLookLikeARMReference uses a simple heuristic to determine if a property looks like it might be an ARM reference.
// This can be used for logging/reporting purposes to discover references which we missed.
func doesPropertyLookLikeARMReference(prop *astmodel.PropertyDefinition) bool {
	// The property must be a string or optional string
	isString := prop.PropertyType().Equals(astmodel.StringType)
	isOptionalString := prop.PropertyType().Equals(astmodel.NewOptionalType(astmodel.StringType))
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
		astmodel.ResourceReferenceTypeName)

	newProp = newProp.WithDescription(existing.Description())
	if existing.HasKubebuilderRequiredValidation() {
		newProp = newProp.MakeRequired()
	} else {
		newProp = newProp.MakeOptional()
	}

	return newProp
}

// TODO: This will go away in favor of a cleaner solution in the future, as obviously this isn't great
func newKnownReferencesMap(configuration *config.Configuration) map[referencePair]struct{} {
	return map[referencePair]struct{}{
		{
			typeName: astmodel.MakeTypeName(configuration.MakeLocalPackageReference("microsoft.batch", "v1alpha1api20210101"), "KeyVaultReference"),
			propName: "Id",
		}: {},
		{
			typeName: astmodel.MakeTypeName(configuration.MakeLocalPackageReference("microsoft.documentdb", "v1alpha1api20210515"), "VirtualNetworkRule"),
			propName: "Id",
		}: {},
		{
			typeName: astmodel.MakeTypeName(configuration.MakeLocalPackageReference("microsoft.storage", "v1alpha1api20210401"), "VirtualNetworkRule"),
			propName: "Id",
		}: {},
		{
			typeName: astmodel.MakeTypeName(configuration.MakeLocalPackageReference("microsoft.network", "v1alpha1api20201101"), "SubResource"),
			propName: "Id",
		}: {},
	}
}

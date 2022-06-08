/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/internal/set"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// ExportControllerResourceRegistrations creates a Stage to generate type registrations
// for resources.
func ExportControllerResourceRegistrations(idFactory astmodel.IdentifierFactory, outputPath string) *Stage {
	return NewLegacyStage(
		"exportControllerResourceRegistrations",
		fmt.Sprintf("Export resource registrations to %q", outputPath),
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			// If the configuration doesn't specify an output destination for us, just do nothing
			if outputPath == "" {
				return definitions, nil
			}

			var resources []astmodel.TypeName
			var storageVersionResources []astmodel.TypeName
			var resourceExtensions []astmodel.TypeName
			indexFunctions := make(map[astmodel.TypeName][]*functions.IndexRegistrationFunction)
			secretPropertyKeys := make(map[astmodel.TypeName][]string)

			// We need to register each version
			for _, def := range definitions {
				if resource, ok := astmodel.AsResourceType(def.Type()); ok {

					if resource.IsStorageVersion() {
						storageVersionResources = append(storageVersionResources, def.Name())

						chains, err := catalogSecretPropertyChains(def, definitions)
						if err != nil {
							return nil, errors.Wrapf(err, "failed to catalog %s property chains", def.Name())
						}

						resourceIndexFunctions, resourceSecretPropertyKeys := handleSecretPropertyChains(chains, idFactory, def)
						indexFunctions[def.Name()] = resourceIndexFunctions
						secretPropertyKeys[def.Name()] = resourceSecretPropertyKeys
					}

					resources = append(resources, def.Name())
				} else if object, ok := astmodel.AsObjectType(def.Type()); ok {
					if object.HasFunctionWithName(functions.ExtendedResourcesFunctionName) {
						resourceExtensions = append(resourceExtensions, def.Name())
					}
				}
			}
			file := NewResourceRegistrationFile(resources, storageVersionResources, indexFunctions, secretPropertyKeys, resourceExtensions)
			fileWriter := astmodel.NewGoSourceFileWriter(file)

			err := fileWriter.SaveToFile(outputPath)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to write controller type registration file to %q", outputPath)
			}

			return definitions, nil
		})
}

func handleSecretPropertyChains(
	chains []propertyChain,
	idFactory astmodel.IdentifierFactory,
	def astmodel.TypeDefinition,
) ([]*functions.IndexRegistrationFunction, []string) {
	indexFunctions := make([]*functions.IndexRegistrationFunction, 0, len(chains))
	secretPropertyKeys := make([]string, 0, len(chains))

	ensureIndexPropertyPathsUnique(chains)

	for _, chain := range chains {
		secretPropertyKey := chain.indexPropertyKey()
		indexFunction := functions.NewIndexRegistrationFunction(
			idFactory,
			chain.indexMethodName(idFactory, def.Name()),
			def.Name(),
			secretPropertyKey,
			chain.props)
		indexFunctions = append(indexFunctions, indexFunction)
		secretPropertyKeys = append(secretPropertyKeys, secretPropertyKey)
	}

	return indexFunctions, secretPropertyKeys
}

func catalogSecretPropertyChains(def astmodel.TypeDefinition, definitions astmodel.TypeDefinitionSet) ([]propertyChain, error) {
	indexBuilder := &indexFunctionBuilder{}

	visitor := astmodel.TypeVisitorBuilder{
		VisitObjectType: indexBuilder.catalogSecretProperties,
	}.Build()

	walker := astmodel.NewTypeWalker(definitions, visitor)
	walker.MakeContext = func(it astmodel.TypeName, ctx interface{}) (interface{}, error) {
		if ctx != nil {
			return ctx, nil
		}

		return propertyChain{}, nil
	}

	_, err := walker.Walk(def)
	if err != nil {
		return nil, errors.Wrapf(err, "error cataloging secret properties")
	}

	return indexBuilder.propChains, nil
}

type indexFunctionBuilder struct {
	propChains []propertyChain
}

type propertyChain struct {
	props []*astmodel.PropertyDefinition
}

func (chain propertyChain) clone() propertyChain {
	duplicate := append([]*astmodel.PropertyDefinition(nil), chain.props...)
	return propertyChain{props: duplicate}
}

// add returns a new chain that includes the given property at the end of the chain.
func (chain propertyChain) add(prop *astmodel.PropertyDefinition) propertyChain {
	newChain := chain.clone()
	newChain.props = append(newChain.props, prop)
	return newChain
}

func preservePropertyChain(_ *astmodel.ObjectType, prop *astmodel.PropertyDefinition, ctx interface{}) (interface{}, error) {
	chain := ctx.(propertyChain)
	return chain.add(prop), nil
}

func (b *indexFunctionBuilder) catalogSecretProperties(this *astmodel.TypeVisitor, it *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
	chain := ctx.(propertyChain)

	it.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
		if prop.IsSecret() {
			chain := chain.add(prop)
			b.propChains = append(b.propChains, chain)
		}
	})

	identityVisit := astmodel.MakeIdentityVisitOfObjectType(preservePropertyChain)
	return identityVisit(this, it, ctx)
}

func (chain *propertyChain) indexMethodName(
	idFactory astmodel.IdentifierFactory,
	resourceTypeName astmodel.TypeName,
	propertyChain propertyChain,
) string {
	// TODO: Technically speaking it's still possible to generate names that clash here, although it's pretty
	// TODO: unlikely. Do we need to do more?

	lastProp := propertyChain.props[len(propertyChain.props)-1]

	group, _ := resourceTypeName.PackageReference.GroupVersion()
	return fmt.Sprintf("index%s%s%s",
		idFactory.CreateIdentifier(group, astmodel.Exported),
		resourceTypeName.Name(),
		chain.indexPropertyPath())
}

func makeIndexPropertyKey(propertyChain propertyChain) string {
	values := []string{
		".spec",
	}
	for _, prop := range propertyChain.props {
		name, ok := prop.JSONName()
		if !ok {
			panic(fmt.Sprintf("property %s has no JSON name", prop.PropertyName()))
		}

		values = append(values, name)
	}

	return strings.Join(values, ".")
}

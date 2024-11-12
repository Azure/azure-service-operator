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

			var resources []astmodel.InternalTypeName
			var storageVersionResources []astmodel.InternalTypeName
			var resourceExtensions []astmodel.InternalTypeName
			indexFunctions := make(map[astmodel.InternalTypeName][]*functions.IndexRegistrationFunction)
			secretPropertyKeys := make(map[astmodel.InternalTypeName][]string)
			configMapPropertyKeys := make(map[astmodel.InternalTypeName][]string)

			// We need to register each version
			for _, def := range definitions {
				if resource, ok := astmodel.AsResourceType(def.Type()); ok {

					if resource.IsStorageVersion() {
						storageVersionResources = append(storageVersionResources, def.Name())

						secretChains, err := catalogSecretPropertyChains(def, definitions)
						if err != nil {
							return nil, errors.Wrapf(err, "failed to catalog %s secret property chains", def.Name())
						}

						configMapChains, err := catalogConfigMapPropertyChains(def, definitions)
						if err != nil {
							return nil, errors.Wrapf(err, "failed to catalog %s configmap property chains", def.Name())
						}

						resourceSecretIndexFunctions, resourceSecretPropertyKeys := transformChainsToIndexFunctionsAndKeys(secretChains, idFactory, def)
						resourceConfigMapIndexFunctions, resourceConfigMapPropertyKeys := transformChainsToIndexFunctionsAndKeys(configMapChains, idFactory, def)

						indexFunctions[def.Name()] = append(resourceSecretIndexFunctions, resourceConfigMapIndexFunctions...)
						secretPropertyKeys[def.Name()] = resourceSecretPropertyKeys
						configMapPropertyKeys[def.Name()] = resourceConfigMapPropertyKeys
					}

					resources = append(resources, def.Name())
				} else if object, ok := astmodel.AsObjectType(def.Type()); ok {
					if object.HasFunctionWithName(functions.ExtendedResourcesFunctionName) {
						resourceExtensions = append(resourceExtensions, def.Name())
					}
				}
			}
			file := NewResourceRegistrationFile(
				resources,
				storageVersionResources,
				indexFunctions,
				secretPropertyKeys,
				configMapPropertyKeys,
				resourceExtensions)
			fileWriter := astmodel.NewGoSourceFileWriter(file)

			err := fileWriter.SaveToFile(outputPath)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to write controller type registration file to %q", outputPath)
			}

			return definitions, nil
		})
}

func transformChainsToIndexFunctionsAndKeys(
	chains []*propertyChain,
	idFactory astmodel.IdentifierFactory,
	def astmodel.TypeDefinition,
) ([]*functions.IndexRegistrationFunction, []string) {
	indexFunctions := make([]*functions.IndexRegistrationFunction, 0, len(chains))
	propertyKeys := make([]string, 0, len(chains))

	ensureIndexPropertyPathsUnique(chains)

	for _, chain := range chains {
		propertyKey := chain.indexPropertyKey()
		indexFunction := functions.NewIndexRegistrationFunction(
			idFactory,
			chain.indexMethodName(idFactory, def.Name()),
			def.Name(),
			propertyKey,
			chain.properties())
		indexFunctions = append(indexFunctions, indexFunction)
		propertyKeys = append(propertyKeys, propertyKey)
	}

	return indexFunctions, propertyKeys
}

// ensureIndexPropertyPathsUnique looks for conflicting index property paths (which would lead to conflicting index
// method names) ensures they all produce unique names.
func ensureIndexPropertyPathsUnique(chains []*propertyChain) {
	// First mark all the properties at the end of each chain as required
	for _, chain := range chains {
		chain.requiredForPropertyPath = true
	}

	// Loop until either we have no collisions, or we can't resolve them
	for {
		// Look for collisions
		chainsByName := make(map[string][]*propertyChain)
		for _, chain := range chains {
			methodName := chain.indexPropertyPath()
			chainsByName[methodName] = append(chainsByName[methodName], chain)
		}

		// For any collision we find (where two or more chains share the same method name), try to resolve it
		pathsChanged := false
		for _, collidingChains := range chainsByName {
			if len(collidingChains) > 1 && tryResolvePropertyPathCollision(collidingChains) {
				pathsChanged = true
			}
		}

		if !pathsChanged {
			break
		}
	}
}

// tryResolvePropertyPathCollision tries to resolve a collision between multiple chains, returning true if it was able
// to make a change (this allows us to terminate if no change is made, ensuring we don't end up in an infinite loop).
// If the parents of our colliding properties have different names, we can disambiguate by including the parent name
// in the property path. If all the parents have the same name, we recursively look at their parents until we find
// either a different name, or we run out of parents.
func tryResolvePropertyPathCollision(chains []*propertyChain) bool {
	// Isolate all unique parents
	parents := set.Make[*propertyChain]()
	for _, chain := range chains {
		if chain.root != nil {
			parents.Add(chain.root)
		}
	}

	if len(parents) == 0 {
		// No parents, nothing to do
		return false
	}

	// Check for parents with different names
	names := set.Make[string]()
	for _, parent := range parents.Values() {
		name := string(parent.prop.PropertyName())
		names.Add(name)
	}

	if len(names) == 1 {
		// All parents have the same name, try resolving with their parents instead
		return tryResolvePropertyPathCollision(parents.Values())
	}

	// We have parents and their names differ, use those names in the property path
	for _, parent := range parents.Values() {
		parent.requiredForPropertyPath = true
	}

	return true
}

func catalogSecretPropertyChains(def astmodel.TypeDefinition, definitions astmodel.TypeDefinitionSet) ([]*propertyChain, error) {
	indexBuilder := &indexFunctionBuilder{}

	visitor := astmodel.TypeVisitorBuilder[*propertyChain]{
		VisitObjectType: indexBuilder.catalogSecretProperties,
	}.Build()

	return catalogPropertyChains(def, definitions, indexBuilder, visitor)
}

func catalogConfigMapPropertyChains(def astmodel.TypeDefinition, definitions astmodel.TypeDefinitionSet) ([]*propertyChain, error) {
	indexBuilder := &indexFunctionBuilder{}

	visitor := astmodel.TypeVisitorBuilder[*propertyChain]{
		VisitObjectType: indexBuilder.catalogConfigMapProperties,
	}.Build()

	return catalogPropertyChains(def, definitions, indexBuilder, visitor)
}

func catalogPropertyChains(
	def astmodel.TypeDefinition,
	definitions astmodel.TypeDefinitionSet,
	indexBuilder *indexFunctionBuilder,
	visitor astmodel.TypeVisitor[*propertyChain],
) ([]*propertyChain, error) {
	walker := astmodel.NewTypeWalker(definitions, visitor)
	walker.MakeContext = func(it astmodel.InternalTypeName, ctx *propertyChain) (*propertyChain, error) {
		if ctx != nil {
			return ctx, nil
		}

		return newPropertyChain(), nil
	}

	_, err := walker.Walk(def)
	if err != nil {
		return nil, errors.Wrapf(err, "error cataloging secret properties")
	}

	return indexBuilder.propChains, nil
}

type indexFunctionBuilder struct {
	propChains []*propertyChain
}

var identityVisitOfObjectTypeWithPropertyChain = astmodel.MakeIdentityVisitOfObjectType(preservePropertyChain)

func (b *indexFunctionBuilder) catalogSecretProperties(
	this *astmodel.TypeVisitor[*propertyChain],
	it *astmodel.ObjectType,
	ctx *propertyChain,
) (astmodel.Type, error) {
	it.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
		if prop.IsSecret() {
			b.propChains = append(b.propChains, ctx.add(prop))
		}
	})

	return identityVisitOfObjectTypeWithPropertyChain(this, it, ctx)
}

func (b *indexFunctionBuilder) catalogConfigMapProperties(
	this *astmodel.TypeVisitor[*propertyChain],
	it *astmodel.ObjectType,
	ctx *propertyChain,
) (astmodel.Type, error) {
	it.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
		typeName, ok := astmodel.AsTypeName(prop.PropertyType())
		if !ok {
			return
		}

		isConfigMapReference := typeName.Equals(astmodel.ConfigMapReferenceType, astmodel.EqualityOverrides{})
		if !isConfigMapReference {
			return
		}

		b.propChains = append(b.propChains, ctx.add(prop))
	})

	return identityVisitOfObjectTypeWithPropertyChain(this, it, ctx)
}

// propertyChain represents a chain of properties that can be used to index a property on a resource. Each chain is made
// up of a leaf property and a reference to a (potentially shared) parent chain. Sharing these parents keeps memory
// consumption down, while also allowing us to include properties partway along the path to resolve ambiguities when
// generating method names.
type propertyChain struct {
	root                    *propertyChain
	prop                    *astmodel.PropertyDefinition
	requiredForPropertyPath bool
}

// newPropertyChain returns a new chain with no properties.
func newPropertyChain() *propertyChain {
	return &propertyChain{
		root: nil,
		prop: nil,
	}
}

// add returns a new chain that includes the given property at the end of the chain.
func (chain *propertyChain) add(prop *astmodel.PropertyDefinition) *propertyChain {
	return &propertyChain{
		root: chain,
		prop: prop,
	}
}

// properties returns the properties in the chain in a new slice.
func (chain *propertyChain) properties() []*astmodel.PropertyDefinition {
	var result []*astmodel.PropertyDefinition
	if chain.root != nil {
		result = chain.root.properties()
	}

	if chain.prop != nil {
		result = append(result, chain.prop)
	}

	return result
}

func preservePropertyChain(
	_ *astmodel.ObjectType,
	prop *astmodel.PropertyDefinition,
	ctx *propertyChain,
) (*propertyChain, error) {
	return ctx.add(prop), nil
}

func (chain *propertyChain) indexMethodName(
	idFactory astmodel.IdentifierFactory,
	resourceTypeName astmodel.InternalTypeName,
) string {
	group := resourceTypeName.InternalPackageReference().Group()
	return fmt.Sprintf("index%s%s%s",
		idFactory.CreateIdentifier(group, astmodel.Exported),
		resourceTypeName.Name(),
		chain.indexPropertyPath())
}

// indexPropertyPath returns the path of the property in the chain, using only those properties that have been flagged
func (chain *propertyChain) indexPropertyPath() string {
	var result string
	if chain.root != nil {
		result = chain.root.indexPropertyPath()
	}

	if chain.requiredForPropertyPath {
		result += chain.prop.PropertyName().String()
	}

	return result
}

// indexPropertyKey makes an indexable key for this property chain. Note that this key is just a string. The fact
// that it looks like a jsonpath expression is purely coincidental. The key may refer to a property that is actually
// a member of a collection, such as .spec.secretsCollection.password. This is OK because the key is just a string
// and all that string is doing is uniquely representing this field.
func (chain *propertyChain) indexPropertyKey() string {
	values := []string{
		".spec",
	}
	for _, prop := range chain.properties() {
		name, ok := prop.JSONName()
		if !ok {
			panic(fmt.Sprintf("property %s has no JSON name", prop.PropertyName()))
		}

		values = append(values, name)
	}

	return strings.Join(values, ".")
}

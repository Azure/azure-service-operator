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
	chains []*propertyChain,
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
			chain.properties())
		indexFunctions = append(indexFunctions, indexFunction)
		secretPropertyKeys = append(secretPropertyKeys, secretPropertyKey)
	}

	return indexFunctions, secretPropertyKeys
}

// ensureIndexPropertyPathsUnique looks for conflicting index property paths (which would lead to conflicting index
// method names) ensures they all produce unique names.
func ensureIndexPropertyPathsUnique(chains []*propertyChain) {
	// First mark all the properties at the end of each chain as required
	for _, chain := range chains {
		chain.requiredForPropertyPath = true
	}

	// Look until either we have no collisions, or we can't resolve them
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
	propChains []*propertyChain
}

// propertyChain represents an chain of properties that can be used to index a secret on a resource. Each chain is made
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

type propertyChain struct {
	props []*astmodel.PropertyDefinition
}

func (ctx propertyChain) clone() propertyChain {
	duplicate := append([]*astmodel.PropertyDefinition(nil), ctx.props...)
	return propertyChain{props: duplicate}
}

func preservePropertyChain(_ *astmodel.ObjectType, prop *astmodel.PropertyDefinition, ctx interface{}) (interface{}, error) {
	chain := ctx.(propertyChain)
	newChain := chain.clone()
	newChain.props = append(newChain.props, prop)
	return newChain, nil
}

func (b *indexFunctionBuilder) catalogSecretProperties(this *astmodel.TypeVisitor, it *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
	chain := ctx.(propertyChain)

	it.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
		if prop.IsSecret() {
			newCtx := chain.clone()
			newCtx.props = append(newCtx.props, prop)
			b.propChains = append(b.propChains, newCtx.props)
		}
	})

	identityVisit := astmodel.MakeIdentityVisitOfObjectType(preservePropertyChain)
	return identityVisit(this, it, ctx)
}

func (chain *propertyChain) indexMethodName(
	idFactory astmodel.IdentifierFactory,
	resourceTypeName astmodel.TypeName,
) string {
	group, _ := resourceTypeName.PackageReference.GroupVersion()
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

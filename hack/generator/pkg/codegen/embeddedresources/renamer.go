/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package embeddedresources

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

type renamer struct {
	types astmodel.Types
}

type renameAction func(original astmodel.TypeName, associatedNames astmodel.TypeNameSet) (astmodel.TypeAssociation, error)

func (r renamer) simplifyEmbeddedNameToOriginalName(original astmodel.TypeName, associatedNames astmodel.TypeNameSet) (astmodel.TypeAssociation, error) {
	_, originalExists := r.types[original]
	if originalExists || len(associatedNames) != 1 {
		return nil, nil
	}

	klog.V(4).Infof("There are no usages of %q. Collapsing %q into the original for simplicity.", original, associatedNames.Single())
	renames := make(astmodel.TypeAssociation)
	renames[associatedNames.Single()] = original

	return renames, nil
}

func (r renamer) simplifyEmbeddedNameRemoveContextAndCount(_ astmodel.TypeName, associatedNames astmodel.TypeNameSet) (astmodel.TypeAssociation, error) {
	if len(associatedNames) != 1 {
		return nil, nil
	}

	renames := make(astmodel.TypeAssociation)
	associated := associatedNames.Single()
	embeddedName, err := parseContextualTypeName(associated)
	if err != nil {
		return nil, err
	}
	embeddedName.context = ""
	embeddedName.count = 0
	renames[associated] = embeddedName.ToSimplifiedTypeName()
	klog.V(4).Infof("There is only a single context %q is used in. Renaming it to %q for simplicity.", associated, renames[associated])

	return renames, nil
}

func (r renamer) simplifyEmbeddedNameRemoveContext(_ astmodel.TypeName, associatedNames astmodel.TypeNameSet) (astmodel.TypeAssociation, error) {
	// Gather information about the associated types
	associatedCountPerContext := make(map[string]int)
	for associated := range associatedNames {
		embeddedName, err := parseContextualTypeName(associated)
		if err != nil {
			return nil, err
		}
		associatedCountPerContext[embeddedName.context] = associatedCountPerContext[embeddedName.context] + 1
	}

	if len(associatedCountPerContext) != 1 {
		return nil, nil
	}

	// If all updated names share the same context, the context is not adding any disambiguation value so we can remove it
	renames := make(astmodel.TypeAssociation)
	for associated := range associatedNames {
		embeddedName, err := parseContextualTypeName(associated)
		if err != nil {
			return nil, err
		}
		embeddedName.context = ""
		renames[associated] = embeddedName.ToSimplifiedTypeName()
	}

	return renames, nil
}

func (r renamer) simplifyEmbeddedName(_ astmodel.TypeName, associatedNames astmodel.TypeNameSet) (astmodel.TypeAssociation, error) {
	// remove _0, which especially for the cases where there's only a single
	// kind of usage will make the type name much clearer
	renames := make(astmodel.TypeAssociation)
	for associated := range associatedNames {
		embeddedName, err := parseContextualTypeName(associated)
		if err != nil {
			return nil, err
		}

		possibleRename := embeddedName.ToSimplifiedTypeName()
		if !possibleRename.Equals(associated) {
			renames[associated] = possibleRename
		}
	}

	return renames, nil
}

func (r renamer) performRenames(
	renames astmodel.TypeAssociation,
	flag astmodel.TypeFlag) (astmodel.Types, error) {

	result := make(astmodel.Types)

	renamingVisitor := astmodel.MakeTypeVisitor()
	renamingVisitor.VisitTypeName = func(this *astmodel.TypeVisitor, it astmodel.TypeName, ctx interface{}) (astmodel.Type, error) {
		if newName, ok := renames[it]; ok {
			return astmodel.IdentityVisitOfTypeName(this, newName, ctx)
		}
		return astmodel.IdentityVisitOfTypeName(this, it, ctx)
	}

	for _, def := range r.types {
		updatedDef, err := renamingVisitor.VisitDefinition(def, nil)
		if err != nil {
			return nil, err
		}
		// TODO: If we don't remove this, something is causing these types to not be emitted.
		// TODO: Unsure what that is... should track it down
		updatedType, err := flag.RemoveFrom(updatedDef.Type())
		if err != nil {
			return nil, err
		}
		result.Add(updatedDef.WithType(updatedType))
	}

	return result, nil
}

// simplifyTypeNames simplifies contextual type names if possible.
func simplifyTypeNames(types astmodel.Types, flag astmodel.TypeFlag) (astmodel.Types, error) {
	// Find all of the type names that have the flag we're interested in
	updatedNames := make(map[astmodel.TypeName]astmodel.TypeNameSet)
	for _, def := range types {
		if flag.IsOn(def.Type()) {
			embeddedName, err := parseContextualTypeName(def.Name())
			if err != nil {
				return nil, err
			}

			updatedNames[embeddedName.original] = updatedNames[embeddedName.original].Add(def.Name())
		}
	}

	renamer := renamer{types: types}
	renameActions := []renameAction{
		renamer.simplifyEmbeddedNameToOriginalName,
		renamer.simplifyEmbeddedNameRemoveContextAndCount,
		renamer.simplifyEmbeddedNameRemoveContext,
		renamer.simplifyEmbeddedName,
	}

	renames := make(astmodel.TypeAssociation)
	for original, associatedNames := range updatedNames {
		for _, action := range renameActions {
			result, err := action(original, associatedNames)
			if err != nil {
				return nil, err
			}

			if result != nil {
				// Add renames
				for oldName, newName := range result {
					renames[oldName] = newName
				}
				break
			}
		}
	}

	return renamer.performRenames(renames, flag)
}

type embeddedResourceTypeName struct {
	original astmodel.TypeName
	context  string
	suffix   string
	count    int
}

func (e embeddedResourceTypeName) ToTypeName() astmodel.TypeName {
	if e.context == "" {
		panic("context cannot be empty when making embedded resource type name")
	}

	if e.suffix == "" {
		panic("suffix cannot be empty when making embedded resource type name")
	}

	nameContext := "_" + e.context
	suffix := "_" + e.suffix
	countString := fmt.Sprintf("_%d", e.count)
	return makeContextualTypeName(e.original, nameContext, suffix, countString)
}

func (e embeddedResourceTypeName) ToSimplifiedTypeName() astmodel.TypeName {
	nameContext := ""
	if e.context != "" {
		nameContext = "_" + e.context
	}
	countString := ""
	if e.count > 0 {
		countString = fmt.Sprintf("_%d", e.count)
	}
	suffix := ""
	if e.suffix != "" {
		suffix = "_" + e.suffix
	}
	return makeContextualTypeName(e.original, nameContext, suffix, countString)
}

func makeContextualTypeName(original astmodel.TypeName, context string, suffix string, count string) astmodel.TypeName {
	return astmodel.MakeTypeName(original.PackageReference, original.Name()+context+suffix+count)
}

func parseContextualTypeName(name astmodel.TypeName) (embeddedResourceTypeName, error) {
	split := strings.Split(name.Name(), "_")
	if len(split) < 4 {
		return embeddedResourceTypeName{}, errors.Errorf("can't split embedded resource type name: %q didn't have 4 sections", name)
	}

	original := strings.Join(split[:len(split)-3], "_")
	resource := split[len(split)-3]
	suffix := split[len(split)-2]
	count, err := strconv.Atoi(split[len(split)-1])
	if err != nil {
		return embeddedResourceTypeName{}, err
	}

	return embeddedResourceTypeName{
		original: astmodel.MakeTypeName(name.PackageReference, original),
		context:  resource,
		suffix:   suffix,
		count:    count,
	}, nil
}

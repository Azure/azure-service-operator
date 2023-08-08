/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package embeddedresources

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

type renamer struct {
	definitions astmodel.TypeDefinitionSet
}

type renameAction func(
	original astmodel.InternalTypeName, // Original name of the resource type
	associatedNames astmodel.TypeNameSet[astmodel.InternalTypeName], // all the subresource names that copies have acquired
	originalNames map[astmodel.TypeName]embeddedResourceTypeName, // map from the new names back to the original
	log logr.Logger, // Log to use for reporting
) (astmodel.TypeAssociation, error)

func (r renamer) simplifyEmbeddedNameToOriginalName(
	original astmodel.InternalTypeName,
	associatedNames astmodel.TypeNameSet[astmodel.InternalTypeName],
	_ map[astmodel.TypeName]embeddedResourceTypeName,
	log logr.Logger,
) (astmodel.TypeAssociation, error) {
	_, originalExists := r.definitions[original]
	if originalExists || len(associatedNames) != 1 {
		return nil, nil
	}

	renames := make(astmodel.TypeAssociation)
	renames[associatedNames.Single()] = original

	log.V(2).Info(
		"Type not used, renaming for simplicity",
		"originalName", associatedNames.Single(),
		"newName", original)

	return renames, nil
}

func (r renamer) simplifyEmbeddedNameRemoveContextAndCount(
	_ astmodel.InternalTypeName,
	associatedNames astmodel.TypeNameSet[astmodel.InternalTypeName],
	originalNames map[astmodel.TypeName]embeddedResourceTypeName,
	log logr.Logger,
) (astmodel.TypeAssociation, error) {
	if len(associatedNames) != 1 {
		return nil, nil
	}

	renames := make(astmodel.TypeAssociation)
	associated := associatedNames.Single()

	embeddedName, ok := originalNames[associated]
	if !ok {
		return nil, errors.Errorf("could not find original name for %q", associated)
	}

	embeddedName.context = ""
	embeddedName.count = 0
	renames[associated] = embeddedName.ToSimplifiedTypeName()

	log.V(2).Info(
		"Type used in single context, renaming for simplicity",
		"originalName", associated,
		"newName", renames[associated])

	return renames, nil
}

func (r renamer) simplifyEmbeddedNameRemoveContext(
	_ astmodel.InternalTypeName,
	associatedNames astmodel.TypeNameSet[astmodel.InternalTypeName],
	originalNames map[astmodel.TypeName]embeddedResourceTypeName,
	log logr.Logger,
) (astmodel.TypeAssociation, error) {
	// Gather information about the associated definitions
	associatedCountPerContext := make(map[string]int, len(associatedNames))
	for associated := range associatedNames {
		embeddedName, ok := originalNames[associated]
		if !ok {
			return nil, errors.Errorf("could not find original name for %q", associated)
		}
		associatedCountPerContext[embeddedName.context] = associatedCountPerContext[embeddedName.context] + 1
	}

	if len(associatedCountPerContext) != 1 {
		return nil, nil
	}

	// If all updated names share the same context, the context is not adding any disambiguation value so we can remove it
	renames := make(astmodel.TypeAssociation)
	for associated := range associatedNames {
		embeddedName, ok := originalNames[associated]
		if !ok {
			return nil, errors.Errorf("could not find original name for %q", associated)
		}
		embeddedName.context = ""
		renames[associated] = embeddedName.ToSimplifiedTypeName()

		log.V(2).Info(
			"Type used in single context, removing context from name",
			"originalName", associated,
			"newName", renames[associated])
	}

	return renames, nil
}

func (r renamer) simplifyEmbeddedName(
	_ astmodel.InternalTypeName,
	associatedNames astmodel.TypeNameSet[astmodel.InternalTypeName],
	originalNames map[astmodel.TypeName]embeddedResourceTypeName,
	log logr.Logger,
) (astmodel.TypeAssociation, error) {
	// remove _0, which especially for the cases where there's only a single
	// kind of usage will make the type name much clearer
	renames := make(astmodel.TypeAssociation)
	for associated := range associatedNames {
		embeddedName, ok := originalNames[associated]
		if !ok {
			return nil, errors.Errorf("could not find original name for %q", associated)
		}

		possibleRename := embeddedName.ToSimplifiedTypeName()
		if !astmodel.TypeEquals(possibleRename, associated) {
			renames[associated] = possibleRename
		}
	}

	return renames, nil
}

func (r renamer) performRenames(
	renames astmodel.TypeAssociation,
	flag astmodel.TypeFlag,
) (astmodel.TypeDefinitionSet, error) {
	result := make(astmodel.TypeDefinitionSet)

	renamingVisitor := astmodel.TypeVisitorBuilder{
		VisitTypeName: func(this *astmodel.TypeVisitor, it astmodel.TypeName, ctx interface{}) (astmodel.Type, error) {
			if newName, ok := renames[it]; ok {
				return astmodel.IdentityVisitOfTypeName(this, newName, ctx)
			}
			return astmodel.IdentityVisitOfTypeName(this, it, ctx)
		},
	}.Build()

	for _, def := range r.definitions {
		updatedDef, err := renamingVisitor.VisitDefinition(def, nil)
		if err != nil {
			return nil, err
		}
		// TODO: If we don't remove this, something is causing these definitions to not be emitted.
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
func simplifyTypeNames(
	definitions astmodel.TypeDefinitionSet,
	flag astmodel.TypeFlag,
	originalNames map[astmodel.TypeName]embeddedResourceTypeName,
	log logr.Logger,
) (astmodel.TypeDefinitionSet, error) {
	// Find all of the type names that have the flag we're interested in
	updatedNames := make(map[astmodel.InternalTypeName]astmodel.TypeNameSet[astmodel.InternalTypeName])
	for _, def := range definitions {
		if flag.IsOn(def.Type()) {
			en, ok := originalNames[def.Name()]
			if !ok {
				return nil, errors.Errorf("failed to find original name for renamed type %s", def.Name())
			}

			if updatedNames[en.original] == nil {
				updatedNames[en.original] = astmodel.NewTypeNameSet(def.Name())
			} else {
				updatedNames[en.original].Add(def.Name())
			}
		}
	}

	r := renamer{definitions: definitions}
	renameActions := []renameAction{
		r.simplifyEmbeddedNameToOriginalName,
		r.simplifyEmbeddedNameRemoveContextAndCount,
		r.simplifyEmbeddedNameRemoveContext,
		r.simplifyEmbeddedName,
	}

	renames := make(astmodel.TypeAssociation)
	for original, associatedNames := range updatedNames {
		for _, action := range renameActions {
			result, err := action(original, associatedNames, originalNames, log)
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

	return r.performRenames(renames, flag)
}

type embeddedResourceTypeName struct {
	original astmodel.TypeName
	context  string
	suffix   string
	count    int
}

func (e embeddedResourceTypeName) ToTypeName() astmodel.InternalTypeName {
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

func makeContextualTypeName(
	original astmodel.TypeName,
	context string,
	suffix string,
	count string,
) astmodel.InternalTypeName {
	return astmodel.MakeInternalTypeName(original.PackageReference(), original.Name()+context+suffix+count)
}

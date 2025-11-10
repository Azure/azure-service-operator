/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/rotisserie/eris"
)

const MoveTypesForVersionMigrationStageID = "moveTypesForVersionMigration"

func MoveTypesForVersionMigration(
	configuration *config.ObjectModelConfiguration,
) *Stage {
	stage := NewStage(
		MoveTypesForVersionMigrationStageID,
		"Move types into new packages to support version migration",
		func(ctx context.Context, state *State) (*State, error) {
			processor := newVersionMigrationFactory(configuration, state.Definitions())
			newDefinitions, err := processor.Process(ctx)
			if err != nil {
				return nil, err
			}

			return state.WithDefinitions(newDefinitions), nil
		})

	stage.RequiresPostrequisiteStages(CreateStorageTypesStageID)

	return stage
}

type versionMigrationFactory struct {
	configuration     *config.ObjectModelConfiguration
	definitions       astmodel.TypeDefinitionSet
	lastLegacyVersion string
}

func newVersionMigrationFactory(
	configuration *config.ObjectModelConfiguration,
	definitions astmodel.TypeDefinitionSet,
) *versionMigrationFactory {
	return &versionMigrationFactory{
		configuration:     configuration,
		definitions:       definitions,
		lastLegacyVersion: "v2.16.0", //!! Change to v2.16.0 once hybrid mode is in place
	}
}

func (p *versionMigrationFactory) Process(ctx context.Context) (astmodel.TypeDefinitionSet, error) {
	// Find all the resources currently using legacy mode
	// versioning that need to be moved to use new style versioning
	toMove := p.findLegacyModeResourcesToMove()

	moved, err := p.moveResources(toMove, "v")
	if err != nil {
		return nil, eris.Wrap(err, "moving resources for version migration")
	}

	toCopy := p.findHybridModeResourcesToCopy()

	copied, err := p.moveResources(toCopy, "v1api")
	if err != nil {
		return nil, eris.Wrap(err, "copying resources for version migration")
	}

	result := p.definitions.Except(toMove).OverlayWith(moved).OverlayWith(copied)
	return result, nil
}

// findLegacyModeResourcesToMove identifies resources using legacy mode versioning that need to be
// moved to new packages using new (simpler) versioning.
// We scan all definitions looking for resources in groups configured for Legacy mode versioning
// that are noted in configuration as being introduced in ASO version 2.17 or later.
func (p *versionMigrationFactory) findLegacyModeResourcesToMove() astmodel.TypeDefinitionSet {
	result := make(astmodel.TypeDefinitionSet)

	for _, def := range p.definitions {
		_, ok := astmodel.AsResourceType(def.Type())
		if !ok {
			continue
		}

		pkg := def.Name().InternalPackageReference()
		group := pkg.Group()

		mode := astmodel.VersionMigrationModeForGroup(group)
		if mode != astmodel.VersionMigrationModeLegacy {
			continue
		}

		introducedIn, ok := p.configuration.SupportedFrom.Lookup(def.Name())
		if !ok {
			continue
		}

		if astmodel.ComparePathAndVersion(introducedIn, p.lastLegacyVersion) > 0 {
			result.Add(def)
		}
	}

	return result
}

// findHybridModeResourcesToCopy identifies resources using hybrid mode versioning that need to be
// copied to old packages using legacy versioning.
// We scan all definitions looking for resources in groups configured for Hybrid mode versioning
// that are noted in configuration as being introduced in ASO version 2.16 or earlier.
func (p *versionMigrationFactory) findHybridModeResourcesToCopy() astmodel.TypeDefinitionSet {
	result := make(astmodel.TypeDefinitionSet)

	for _, def := range p.definitions {
		_, ok := astmodel.AsResourceType(def.Type())
		if !ok {
			continue
		}

		pkg := def.Name().InternalPackageReference()
		group := pkg.Group()

		mode := astmodel.VersionMigrationModeForGroup(group)
		if mode != astmodel.VersionMigrationModeHybrid {
			continue
		}

		introducedIn, ok := p.configuration.SupportedFrom.Lookup(def.Name())
		if !ok {
			continue
		}

		if astmodel.ComparePathAndVersion(introducedIn, p.lastLegacyVersion) <= 0 {
			result.Add(def)
		}
	}

	return result
}

func (p *versionMigrationFactory) moveResources(
	definitions astmodel.TypeDefinitionSet,
	prefix string,
) (astmodel.TypeDefinitionSet, error) {
	defs, err := p.collectRelatedDefinitions(definitions)
	if err != nil {
		return nil, eris.Wrapf(err, "moving resources to packages with version prefix %s", prefix)
	}

	renames := p.createRenameMap(defs, prefix)

	visitor := astmodel.NewRenamingVisitor(renames)
	renamed, err := visitor.RenameAll(defs)
	if err != nil {
		return nil, eris.Wrap(err, "creating types for backward compatibility")
	}

	return renamed, nil
}

// collectRelatedDefinitions finds all definitions transitively connected to the given definitions.
func (p *versionMigrationFactory) collectRelatedDefinitions(
	defs astmodel.TypeDefinitionSet,
) (astmodel.TypeDefinitionSet, error) {
	connected, err := astmodel.FindConnectedDefinitions(p.definitions, defs)
	if err != nil {
		return nil, eris.Wrap(
			err,
			"finding types connected to resources requiring version migration")
	}

	return connected, nil
}

func (p *versionMigrationFactory) createRenameMap(
	set astmodel.TypeDefinitionSet,
	versionPrefix string,
) astmodel.TypeAssociation {
	result := make(astmodel.TypeAssociation)

	for name := range set {
		if _, ok := result[name]; !ok {
			newName := p.moveTypeName(name, versionPrefix)
			result[name] = newName
		}
	}

	return result
}

func (p *versionMigrationFactory) moveTypeName(
	typeName astmodel.InternalTypeName,
	versionPrefix string,
) astmodel.InternalTypeName {
	pkg := p.movePackageReference(typeName.InternalPackageReference(), versionPrefix)
	return typeName.WithPackageReference(pkg)
}

func (p *versionMigrationFactory) movePackageReference(
	pkg astmodel.InternalPackageReference,
	versionPrefix string,
) astmodel.InternalPackageReference {
	switch r := pkg.(type) {
	case astmodel.LocalPackageReference:
		return r.WithVersionPrefix(versionPrefix)
	case astmodel.SubPackageReference:
		newParent := p.movePackageReference(r.Parent(), versionPrefix)
		return astmodel.MakeSubPackageReference(r.PackageName(), newParent)
	default:
		panic(fmt.Sprintf("unexpected package reference type %T", r))
	}
}

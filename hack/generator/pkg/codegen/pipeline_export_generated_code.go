/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"
)

// exportPackages creates a PipelineStage to export our generated code as a set of packages
func exportPackages(outputPath string) PipelineStage {
	description := fmt.Sprintf("Export packages to %q", outputPath)
	return MakePipelineStage(
		"exportPackages",
		description,
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			packages, err := CreatePackagesForDefinitions(types)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to assign generated definitions to packages")
			}

			packages, err = MarkLatestResourceVersionsForStorage(packages)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to mark latest resource versions for as storage versions")
			}

			err = writeFiles(ctx, packages, outputPath)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to write files into %q", outputPath)
			}

			return types, nil
		})
}

// CreatePackagesForDefinitions groups type definitions into packages
func CreatePackagesForDefinitions(definitions astmodel.Types) (map[astmodel.PackageReference]*astmodel.PackageDefinition, error) {

	genVersion := combinedVersion()
	packages := make(map[astmodel.PackageReference]*astmodel.PackageDefinition)
	for _, def := range definitions {
		defName := def.Name()
		groupName, pkgName, err := defName.PackageReference.GroupAndPackage()
		if err != nil {
			return nil, err
		}

		pkgRef := defName.PackageReference
		if pkg, ok := packages[pkgRef]; ok {
			pkg.AddDefinition(def)
		} else {
			pkg = astmodel.NewPackageDefinition(groupName, pkgName, genVersion)
			pkg.AddDefinition(def)
			packages[pkgRef] = pkg
		}
	}

	return packages, nil
}

// MarkLatestResourceVersionsForStorage marks the latest version of each resource as the storage version
func MarkLatestResourceVersionsForStorage(
	packages map[astmodel.PackageReference]*astmodel.PackageDefinition) (map[astmodel.PackageReference]*astmodel.PackageDefinition, error) {

	result := make(map[astmodel.PackageReference]*astmodel.PackageDefinition)

	resourceLookup, err := groupResourcesByVersion(packages)
	if err != nil {
		return nil, err
	}

	for pkgRef, pkg := range packages {

		resultPkg := astmodel.NewPackageDefinition(pkg.GroupName, pkg.PackageName, pkg.GeneratorVersion)
		for _, def := range pkg.Definitions() {
			// see if it is a resource
			if resourceType, ok := def.Type().(*astmodel.ResourceType); ok {

				unversionedName, err := getUnversionedName(def.Name())
				if err != nil {
					// should never happen as all resources have versioned names
					return nil, err
				}

				allVersionsOfResource := resourceLookup[unversionedName]
				latestVersionOfResource := allVersionsOfResource[len(allVersionsOfResource)-1]

				thisPackagePath := def.Name().PackageReference.PackagePath()
				latestPackagePath := latestVersionOfResource.Name().PackageReference.PackagePath()

				// mark as storage version if it's the latest version
				isLatestVersion := thisPackagePath == latestPackagePath
				if isLatestVersion {
					def = astmodel.MakeTypeDefinition(def.Name(), resourceType.MarkAsStorageVersion()).
						WithDescription(def.Description())
				}

				resultPkg.AddDefinition(def)
			} else {
				// otherwise simply add it
				resultPkg.AddDefinition(def)
			}
		}

		result[pkgRef] = resultPkg
	}

	return result, nil
}

func writeFiles(ctx context.Context, packages map[astmodel.PackageReference]*astmodel.PackageDefinition, outputPath string) error {
	fileCount := 0
	definitionCount := 0

	var pkgs []*astmodel.PackageDefinition
	for _, pkg := range packages {
		pkgs = append(pkgs, pkg)
	}

	// Sort the list of packages to ensure we always write them to disk in the same sequence
	sort.Slice(pkgs, func(i int, j int) bool {
		iPkg := pkgs[i]
		jPkg := pkgs[j]
		return iPkg.GroupName < jPkg.GroupName ||
			(iPkg.GroupName == jPkg.GroupName && iPkg.PackageName < jPkg.PackageName)
	})

	// emit each package
	klog.V(0).Infof("Writing output files into %v", outputPath)
	for _, pkg := range pkgs {
		if ctx.Err() != nil { // check for cancellation
			return ctx.Err()
		}

		// create directory if not already there
		outputDir := filepath.Join(outputPath, pkg.GroupName, pkg.PackageName)
		if _, err := os.Stat(outputDir); os.IsNotExist(err) {
			klog.V(5).Infof("Creating directory %q\n", outputDir)
			err = os.MkdirAll(outputDir, 0700)
			if err != nil {
				klog.Fatalf("Unable to create directory %q", outputDir)
			}
		}

		count, err := pkg.EmitDefinitions(outputDir, packages)
		if err != nil {
			return errors.Wrapf(err, "error writing definitions into %q", outputDir)
		}

		fileCount += count
		definitionCount += pkg.DefinitionCount()
	}

	klog.V(0).Infof("Completed writing %v files containing %v definitions", fileCount, definitionCount)
	return nil
}

func groupResourcesByVersion(packages map[astmodel.PackageReference]*astmodel.PackageDefinition) (map[unversionedName][]astmodel.TypeDefinition, error) {

	result := make(map[unversionedName][]astmodel.TypeDefinition)

	for _, pkg := range packages {
		for _, def := range pkg.Definitions() {
			if _, ok := def.Type().(*astmodel.ResourceType); ok {
				name, err := getUnversionedName(def.Name())
				if err != nil {
					// this should never happen as resources will all have versioned names
					return nil, errors.Wrapf(err, "Unable to extract unversioned name in groupResources")
				}

				result[name] = append(result[name], def)
			}
		}
	}

	// order each set of resources by package name (== by version as these are sortable dates)
	for _, slice := range result {
		sort.Slice(slice, func(i, j int) bool {
			return slice[i].Name().PackageReference.PackageName() < slice[j].Name().PackageReference.PackageName()
		})
	}

	return result, nil
}

func getUnversionedName(name astmodel.TypeName) (unversionedName, error) {
	group, _, err := name.PackageReference.GroupAndPackage()
	if err != nil {
		return unversionedName{}, err
	}

	return unversionedName{group, name.Name()}, nil
}

type unversionedName struct {
	group string
	name  string
}

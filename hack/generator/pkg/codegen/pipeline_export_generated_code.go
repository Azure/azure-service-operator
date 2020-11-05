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
	"time"

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
		pkgRef, ok := defName.PackageReference.AsLocalPackage()
		if !ok {
			klog.Errorf("Definition %v from external package %v skipped", defName.Name(), defName.PackageReference)
			continue
		}

		groupName := pkgRef.Group()
		pkgName := pkgRef.PackageName()

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

func writeFiles(ctx context.Context, packages map[astmodel.PackageReference]*astmodel.PackageDefinition, outputPath string) error {
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
	klog.V(0).Infof("Writing %d packages into %q", len(pkgs), outputPath)

	globalProgress := newProgressMeter()
	groupProgress := newProgressMeter()

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

		globalProgress.LogProgress("", pkg.DefinitionCount(), count)
		groupProgress.LogProgress(pkg.GroupName, pkg.DefinitionCount(), count)
	}

	globalProgress.Log()

	return nil
}

func newProgressMeter() *progressMeter {
	return &progressMeter{
		resetAt: time.Now(),
	}
}

// progressMeter is a utility struct used to improve our reporting of progress while exporting files
type progressMeter struct {
	label       string
	definitions int
	files       int
	resetAt     time.Time
}

// Log() writes a log message for our progress to this point
func (export *progressMeter) Log() {
	started := export.resetAt
	export.resetAt = time.Now()

	if export.definitions == 0 && export.files == 0 {
		return
	}

	elapsed := time.Since(started).Round(time.Millisecond)
	if export.label != "" {
		klog.V(2).Infof("Wrote %d files containing %d definitions for %v in %v", export.files, export.definitions, export.label, elapsed)
	} else {
		klog.V(2).Infof("Wrote %d files containing %d definitions in %v", export.files, export.definitions, time.Since(started))
	}

	export.resetAt = time.Now()
}

// LogProgress() accumulates totals until a new label is supplied, when it will write a log message
func (export *progressMeter) LogProgress(label string, definitions int, files int) {
	if export.label != label {
		// New group, output our current totals and reset
		export.Log()
		export.definitions = 0
		export.files = 0
		export.resetAt = time.Now()
	}

	export.label = label
	export.definitions += definitions
	export.files += files
}

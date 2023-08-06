/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// ExportPackagesStageID is the unique identifier for this pipeline stage
const ExportPackagesStageID = "exportPackages"

// ExportPackages creates a Stage to export our generated code as a set of packages
func ExportPackages(
	outputPath string,
	emitDocFiles bool,
	log logr.Logger,
) *Stage {
	description := fmt.Sprintf("Export packages to %q", outputPath)
	stage := NewLegacyStage(
		ExportPackagesStageID,
		description,
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			packages, err := CreatePackagesForDefinitions(definitions)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to assign generated definitions to packages")
			}

			err = writeFiles(packages, outputPath, emitDocFiles, log)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to write files into %q", outputPath)
			}

			return definitions, nil
		})

	stage.RequiresPrerequisiteStages(DeleteGeneratedCodeStageID)

	return stage
}

// CreatePackagesForDefinitions groups type definitions into packages
func CreatePackagesForDefinitions(definitions astmodel.TypeDefinitionSet) (map[astmodel.PackageReference]*astmodel.PackageDefinition, error) {
	packages := make(map[astmodel.PackageReference]*astmodel.PackageDefinition)
	for _, def := range definitions {
		name := def.Name()
		ref := name.PackageReference()
		if pkg, ok := packages[ref]; ok {
			pkg.AddDefinition(def)
		} else {
			pkg = astmodel.NewPackageDefinition(ref)
			pkg.AddDefinition(def)
			packages[ref] = pkg
		}
	}

	return packages, nil
}

func writeFiles(
	packages map[astmodel.PackageReference]*astmodel.PackageDefinition,
	outputPath string,
	emitDocFiles bool,
	log logr.Logger,
) error {
	pkgs := make([]*astmodel.PackageDefinition, 0, len(packages))
	for _, pkg := range packages {
		pkgs = append(pkgs, pkg)
	}

	// Sort the list of packages to ensure we always write them to disk in the same sequence
	slices.SortFunc(pkgs, func(left *astmodel.PackageDefinition, right *astmodel.PackageDefinition) bool {
		return left.Path < right.Path
	})

	// emit each package
	log.Info(
		"Writing packages",
		"count", len(pkgs),
		"outputPath", outputPath)

	globalProgress := newProgressMeter()
	groupProgress := newProgressMeter()

	var eg errgroup.Group
	eg.SetLimit(8)

	for _, pkg := range pkgs {
		pkg := pkg
		eg.Go(func() error {
			// create directory if not already there
			outputDir := filepath.Join(outputPath, pkg.Path)
			if _, err := os.Stat(outputDir); os.IsNotExist(err) {
				err = os.MkdirAll(outputDir, 0o700)
				if err != nil {
					return errors.Wrapf(err, "unable to create directory %q", outputDir)
				}
			}

			count, err := pkg.EmitDefinitions(outputDir, packages, emitDocFiles)
			if err != nil {
				return errors.Wrapf(err, "error writing definitions into %q", outputDir)
			}

			globalProgress.LogProgress("", pkg.DefinitionCount(), count, log)
			groupProgress.LogProgress(pkg.Path, pkg.DefinitionCount(), count, log)
			return nil
		})
	}

	err := eg.Wait()
	if err != nil {
		return err
	}

	// log anything leftover
	globalProgress.mutex.Lock()
	defer globalProgress.mutex.Unlock()
	globalProgress.Log(log)

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

	mutex sync.Mutex
}

// Log writes a log message for our progress to this point
func (export *progressMeter) Log(log logr.Logger) {
	started := export.resetAt
	export.resetAt = time.Now()

	if export.definitions == 0 && export.files == 0 {
		return
	}

	elapsed := time.Since(started).Round(time.Millisecond)
	if export.label != "" {
		log.V(1).Info(
			"Wrote files",
			"label", export.label,
			"files", export.files,
			"types", export.definitions,
			"elapsed", elapsed)
	} else {
		log.V(1).Info(
			"Wrote files",
			"files", export.files,
			"types", export.definitions,
			"elapsed", elapsed)
	}

	export.resetAt = time.Now()
}

// LogProgress accumulates totals until a new label is supplied, when it will write a log message
func (export *progressMeter) LogProgress(
	label string,
	definitions int,
	files int,
	log logr.Logger,
) {
	export.mutex.Lock()
	defer export.mutex.Unlock()

	if export.label != label {
		// New group, output our current totals and reset
		export.Log(log)
		export.definitions = 0
		export.files = 0
		export.resetAt = time.Now()
	}

	export.label = label
	export.definitions += definitions
	export.files += files
}

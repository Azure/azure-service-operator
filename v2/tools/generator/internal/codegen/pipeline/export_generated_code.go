/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

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

			err = writeFiles(ctx, packages, outputPath, emitDocFiles, log)
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
		ref := name.PackageReference
		group, version := ref.GroupVersion()
		if pkg, ok := packages[ref]; ok {
			pkg.AddDefinition(def)
		} else {
			pkg = astmodel.NewPackageDefinition(group, version)
			pkg.AddDefinition(def)
			packages[ref] = pkg
		}
	}

	return packages, nil
}

func writeFiles(
	ctx context.Context,
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
	sort.Slice(pkgs, func(i int, j int) bool {
		iPkg := pkgs[i]
		jPkg := pkgs[j]
		return iPkg.GroupName < jPkg.GroupName ||
			(iPkg.GroupName == jPkg.GroupName && iPkg.PackageName < jPkg.PackageName)
	})

	// emit each package
	log.Info(
		"Writing packages",
		"count", len(pkgs),
		"outputPath", outputPath)

	globalProgress := newProgressMeter()
	groupProgress := newProgressMeter()

	var wg sync.WaitGroup

	pkgQueue := make(chan *astmodel.PackageDefinition, 100)
	errs := make(chan error, 10) // we will buffer up to 10 errors and ignore any leftovers

	// write outputs with 8 workers
	// this is parallelized mostly due to 'dst' conversion being slow, see: https://github.com/Azure/k8s-infra/pull/376
	// potentially we could contribute improvements upstream
	for c := 0; c < 8; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for pkg := range pkgQueue {
				if ctx.Err() != nil { // check for cancellation
					return
				}

				// create directory if not already there
				outputDir := filepath.Join(outputPath, pkg.GroupName, pkg.PackageName)
				if _, err := os.Stat(outputDir); os.IsNotExist(err) {
					err = os.MkdirAll(outputDir, 0o700)
					if err != nil {
						select { // try to write to errs, ignore if buffer full
						case errs <- errors.Wrapf(err, "unable to create directory %q", outputDir):
						default:
						}
						return
					}
				}

				count, err := pkg.EmitDefinitions(outputDir, packages, emitDocFiles)
				if err != nil {
					select { // try to write to errs, ignore if buffer full
					case errs <- errors.Wrapf(err, "error writing definitions into %q", outputDir):
					default:
					}
					return
				} else {
					globalProgress.LogProgress("", pkg.DefinitionCount(), count, log)
					groupProgress.LogProgress(pkg.GroupName, pkg.DefinitionCount(), count, log)
				}
			}
		}()
	}

	// send to workers
	// and wait for them to finish
	for _, pkg := range pkgs {
		pkgQueue <- pkg
	}
	close(pkgQueue)
	wg.Wait()

	// collect all errors, if any
	close(errs)
	totalErrs := make([]error, 0, len(errs))
	for err := range errs {
		totalErrs = append(totalErrs, err)
	}

	err := kerrors.NewAggregate(totalErrs)
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

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"bufio"
	"context"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/bmatcuk/doublestar"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// DeleteGeneratedCodeStageID is the unique identifier of this stage
const DeleteGeneratedCodeStageID = "deleteGenerated"

// DeleteGeneratedCode creates a pipeline stage for cleanup of our output folder prior to generating files
func DeleteGeneratedCode(outputFolder string) *Stage {
	return NewLegacyStage(
		DeleteGeneratedCodeStageID,
		"Delete generated code from "+outputFolder,
		func(ctx context.Context, definitions astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
			err := deleteGeneratedCodeFromFolder(ctx, outputFolder)
			if err != nil {
				return nil, err
			}

			return definitions, nil
		})
}

func deleteGeneratedCodeFromFolder(ctx context.Context, outputFolder string) error {
	genPattern := filepath.Join(outputFolder, "**", "*", "*"+astmodel.CodeGeneratedFileSuffix+"*.go")
	err := deleteGeneratedCodeByPattern(ctx, genPattern)
	if err != nil {
		return err
	}

	docPattern := filepath.Join(outputFolder, "**", "*", "doc.go")
	err = deleteGeneratedCodeByPattern(ctx, docPattern)
	if err != nil {
		return err
	}

	structurePattern := filepath.Join(outputFolder, "**", "*", "*.txt")
	err = deleteGeneratedCodeByPattern(ctx, structurePattern)
	if err != nil {
		return err
	}

	return deleteEmptyDirectories(ctx, outputFolder)
}

func deleteGeneratedCodeByPattern(ctx context.Context, globPattern string) error {
	// We use doublestar here rather than filepath.Glob because filepath.Glob doesn't support **
	files, err := doublestar.Glob(globPattern)
	if err != nil {
		return errors.Wrapf(err, "error globbing files with pattern %q", globPattern)
	}

	// We treat files as a queue of files needing deletion
	// If we have an error deleting a file, it might be due to a transient lock on the file (e.g. from an editor or
	// security software); if this happens, we requeue the file so we retry the deletion later
	// To avoid getting stuck in an infinite loop, we keep track of the number of consequetive deletion failures; if
	// this exceeds the size of the queue, we give up and return an error

	consecutiveDeleteFailures := 0
	errorsSeen := make(map[string]error, len(files))
	for len(files) > 0 {
		file := files[0]
		if ctx.Err() != nil { // check for cancellation
			return ctx.Err()
		}

		isGenerated, err := isFileGenerated(file)
		if err != nil {
			errorsSeen[file] = errors.Wrapf(err, "error determining if file was generated")
			consecutiveDeleteFailures++
			files = append(files, file) // requeue the file
		}

		if isGenerated {
			err := os.Remove(file)
			if err != nil {
				errorsSeen[file] = errors.Wrapf(err, "error determining if file was generated")
				consecutiveDeleteFailures++
				files = append(files, file) // requeue the file
			} else {
				consecutiveDeleteFailures = 0
				delete(errorsSeen, file) // remove from errors if we previously failed
			}
		}

		files = files[1:]

		if consecutiveDeleteFailures > len(files) {
			break
		}
	}

	if len(errorsSeen) > 0 {
		errs := make([]error, 0, len(errorsSeen))
		for _, err := range errorsSeen {
			errs = append(errs, err)
		}

		return kerrors.NewAggregate(errs)
	}

	return nil
}

func isFileGenerated(filename string) (bool, error) {
	// Technically, the code generated message could be on any line according to
	// the specification at https://github.com/golang/go/issues/13560 but
	// for our purposes checking the first few lines is plenty
	maxLinesToCheck := 20

	f, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for i := 0; i < maxLinesToCheck; i++ {
		line, err := reader.ReadString('\n')
		if errors.Is(err, io.EOF) {
			return false, nil
		}

		if err != nil {
			return false, err
		}

		for _, codeGenComment := range astmodel.CodeGenerationComments {
			if strings.Contains(line, codeGenComment) {
				return true, nil
			}
		}
	}

	return false, nil
}

func deleteEmptyDirectories(ctx context.Context, path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}

	// TODO: There has to be a better way to do this?
	var dirs []string

	// Second pass to clean up empty directories
	walkFunction := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if ctx.Err() != nil { // check for cancellation
			return ctx.Err()
		}

		if info.IsDir() {
			dirs = append(dirs, path)
		}

		return nil
	}
	err := filepath.Walk(path, walkFunction)
	if err != nil {
		return err
	}

	// Now order the directories by deepest first - we have to do this because otherwise a directory
	// isn't empty because it has a bunch of empty directories inside of it
	sortFunction := func(i int, j int) bool {
		// Comparing by length is sufficient here because a nested directory path
		// will always be longer than just the parent directory path
		return len(dirs[i]) > len(dirs[j])
	}
	sort.Slice(dirs, sortFunction)

	var errs []error

	// Now clean things up
	for _, dir := range dirs {
		if ctx.Err() != nil { // check for cancellation
			return ctx.Err()
		}

		files, err := ioutil.ReadDir(dir)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "error reading directory %q", dir))
		}

		if len(files) == 0 {
			// Directory is empty now, we can delete it
			err := os.Remove(dir)
			if err != nil {
				errs = append(errs, errors.Wrapf(err, "error removing dir %q", dir))
			}
		}
	}

	return kerrors.NewAggregate(errs)
}

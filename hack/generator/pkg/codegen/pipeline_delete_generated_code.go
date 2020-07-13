/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"bufio"
	"context"
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/bmatcuk/doublestar"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

// deleteGeneratedCode creates a pipeline stage for cleanup of our output folder prior to generating files
func deleteGeneratedCode(outputFolder string) PipelineStage {
	return PipelineStage{
		"Delete generated code from " + outputFolder,
		func(ctx context.Context, types map[astmodel.TypeName]astmodel.TypeDefiner) (map[astmodel.TypeName]astmodel.TypeDefiner, error) {
			err := deleteGeneratedCodeFromFolder(ctx, outputFolder)
			if err != nil {
				return nil, err
			}

			return types, nil
		}}
}

func deleteGeneratedCodeFromFolder(ctx context.Context, outputFolder string) error {
	// We use doublestar here rather than filepath.Glob because filepath.Glob doesn't support **
	globPattern := path.Join(outputFolder, "**", "*", "*"+astmodel.CodeGeneratedFileSuffix)

	files, err := doublestar.Glob(globPattern)
	if err != nil {
		return errors.Wrapf(err, "error globbing files with pattern %q", globPattern)
	}

	var errs []error

	for _, file := range files {
		if ctx.Err() != nil { // check for cancellation
			return ctx.Err()
		}

		isGenerated, err := isFileGenerated(file)

		if err != nil {
			errs = append(errs, errors.Wrapf(err, "error determining if file was generated"))
		}

		if isGenerated {
			err := os.Remove(file)
			if err != nil {
				errs = append(errs, errors.Wrapf(err, "error removing file %q", file))
			}
		}
	}

	err = deleteEmptyDirectories(ctx, outputFolder)
	if err != nil {
		errs = append(errs, err)
	}

	return kerrors.NewAggregate(errs)
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
		if err == io.EOF {
			return false, nil
		}
		if err != nil {
			return false, err
		}

		if strings.Contains(line, astmodel.CodeGenerationComment) {
			return true, nil
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

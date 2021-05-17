/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"bufio"
	"io"
	"os"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

// GoSourceFile represents a file of Go code
type GoSourceFile interface {
	// AsAst transforms the file into a dst.File (you can think of this as the AST representation of the file)
	AsAst() (*dst.File, error)
}

type GoSourceFileWriter struct {
	file GoSourceFile
}

// NewGoSourceFileWriter creates a new writer for writing the given file
func NewGoSourceFileWriter(file GoSourceFile) *GoSourceFileWriter {
	return &GoSourceFileWriter{
		file: file,
	}
}

// SaveToWriter writes the given FileAst to the destination
func (w *GoSourceFileWriter) SaveToWriter(destination io.Writer) error {
	content, err := w.file.AsAst()
	if err != nil {
		return err
	}

	buf := bufio.NewWriter(destination)
	defer buf.Flush()

	return decorator.Fprint(buf, content)
}

// SaveToFile writes the given FileAst to the specified file path
func (w *GoSourceFileWriter) SaveToFile(filePath string) error {

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer func() {
		file.Close()

		// if we are panicking, the file will be in a broken
		// state, so remove it
		if r := recover(); r != nil {
			os.Remove(filePath)
			panic(r)
		}
	}()

	err = w.SaveToWriter(file)
	if err != nil {
		// cleanup in case of errors
		file.Close()
		os.Remove(filePath)
	}

	return err
}

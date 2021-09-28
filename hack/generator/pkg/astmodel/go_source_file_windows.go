//go:build windows
// +build windows

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"bufio"
	"io"

	"github.com/dave/dst/decorator"
	"golang.org/x/text/transform"
)

// SaveToWriter writes the given FileAst to the destination
func (w *GoSourceFileWriter) SaveToWriter(destination io.Writer) error {
	content, err := w.file.AsAst()
	if err != nil {
		return err
	}

	buf := bufio.NewWriter(destination)
	defer buf.Flush()

	writer := transform.NewWriter(buf, &crlfTransformer{})

	return decorator.Fprint(writer, content)
}

type crlfTransformer struct {
	lastChar byte
}

var _ transform.Transformer = &crlfTransformer{}

func (t *crlfTransformer) Transform(
	destination []byte,
	source []byte,
	atEOF bool) (countWritten int, countRead int, err error) {

	destinationIndex := 0
	sourceIndex := 0
	for destinationIndex < len(destination) && sourceIndex < len(source) {
		c := source[sourceIndex]
		if c == '\n' {
			if t.lastChar != '\r' {
				// Convert this EoLn to windows by injecting an extra \r
				destination[destinationIndex] = '\r'
				destinationIndex++
			}
		}

		destination[destinationIndex] = c
		destinationIndex++
		sourceIndex++

		t.lastChar = c
	}

	return destinationIndex, sourceIndex, nil
}

func (t *crlfTransformer) Reset() {
	// Nothing
}

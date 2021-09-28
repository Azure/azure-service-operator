//go:build !windows
// +build !windows

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"bufio"
	"io"

	"github.com/dave/dst/decorator"
)

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

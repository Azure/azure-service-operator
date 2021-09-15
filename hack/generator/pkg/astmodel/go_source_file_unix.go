//+build !windows

package astmodel

import (
	"bufio"

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

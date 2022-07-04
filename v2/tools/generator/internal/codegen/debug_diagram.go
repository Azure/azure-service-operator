/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/pipeline"
)

type debugDiagram struct {
	debugDir   string                       // The directory to write the diagram to
	stageIds   map[*pipeline.Stage]string   // Map of stages to their unique IDs
	stageNames map[string][]*pipeline.Stage // Map of stage names to stages
}

// newDebugDiagram creates a new debugDiagram to write into the specified directory.
func newDebugDiagram(debugDir string) *debugDiagram {
	return &debugDiagram{
		debugDir:   debugDir,
		stageIds:   make(map[*pipeline.Stage]string, 70),
		stageNames: make(map[string][]*pipeline.Stage, 70),
	}
}

// writeDiagram writes a diagram of the pipeline to pipeline.dot
func (diagram *debugDiagram) writeDiagram(stages []*pipeline.Stage) error {
	dotsrc := diagram.createDiagram(stages)
	filename := filepath.Join(diagram.debugDir, "pipeline.dot")
	err := ioutil.WriteFile(filename, dotsrc, 0600)
	klog.V(2).Infof("Wrote diagram for pipeline to %s", filename)
	return errors.Wrapf(err, "failed to write diagram to %s", filename)
}

// createDiagram creates a dot file for the pipeline
func (diagram *debugDiagram) createDiagram(stages []*pipeline.Stage) []byte {
	var b strings.Builder

	b.WriteString("// Render with\n")
	b.WriteString("// dot -Tpng -o pipeline.png .\\pipeline.dot\n")
	b.WriteString("// \n")
	b.WriteString("digraph pipeline {\n\n")
	b.WriteString("    node [shape=box];\n")

	// blockSize is the number of nodes or edges to display in each block
	// Dividing nodes and edges into blocks makes the graph source code easier to read
	const blockSize = 10

	// Create nodes
	for index, stage := range stages {
		b.WriteString(
			fmt.Sprintf("    %s [label=\"%s\"];\n",
				diagram.idFor(stage),
				diagram.safeDescription(stage.Description())))

		// If we've reached the end of the block, add a newline
		if index < len(stages)-1 && index%(blockSize+1) == blockSize {
			b.WriteString("\n")
		}
	}

	// Create edges
	// We construct edges in different directions so the nodes zigzag on the diagram,
	// ensuring it doesn't end up looking weirdly tall and thin
	forward := true // true if we're drawing edges forward, false if backwards
	b.WriteString("\n    edge [arrowhead=normal]\n")
	for index, stage := range stages {
		if index < len(stages)-1 {
			nextStage := stages[index+1]
			if forward {
				b.WriteString(
					fmt.Sprintf("    %s -> %s [dir=forward];\n",
						diagram.idFor(stage),
						diagram.idFor(nextStage)))
			} else {
				b.WriteString(
					fmt.Sprintf("    %s -> %s [dir=back];\n",
						diagram.idFor(nextStage),
						diagram.idFor(stage)))
			}
		}

		// If we've reached the end of the block, add a newline
		if index < len(stages)-1 && index%(blockSize+1) == blockSize {
			forward = !forward
			b.WriteString("\n")
		}
	}

	b.WriteString("}\n\n")

	return []byte(b.String())
}

func (diagram *debugDiagram) idFor(stage *pipeline.Stage) string {
	if id, ok := diagram.stageIds[stage]; ok {
		return id
	}

	// Generate a new ID for this stage
	// (We need to do this because the stages are sometimes reused)
	id := diagram.safeId(stage.Id())
	clashes := diagram.stageNames[id]

	if len(clashes) > 0 {
		id = fmt.Sprintf("%s%d", id, len(clashes))
	}

	// Store our id
	diagram.stageIds[stage] = id // Quick lookup if we need it again

	// Keep track of how many times we've seen a particular stage, so we can disambiguate references
	diagram.stageNames[stage.Id()] = append(clashes, stage)

	return id
}

// safeId returns a string containing only alphanumeric characters
func (diagram *debugDiagram) safeId(id string) string {
	var b strings.Builder
	for _, r := range id {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// safeDescription returns a string that can be used as a graphviz label
func (diagram *debugDiagram) safeDescription(description string) string {
	var b strings.Builder
	for _, r := range description {
		if r == '"' {
			r = '\''
		}

		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsPunct(r) || r == ' ' {
			b.WriteRune(r)
		}
	}

	lines := astbuilder.WordWrap(b.String(), 30)

	return strings.Join(lines, "\\n")
}

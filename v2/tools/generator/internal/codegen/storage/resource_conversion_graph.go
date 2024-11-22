/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"fmt"
	"io"

	"github.com/rotisserie/eris"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// ResourceConversionGraph represents the directed graph of conversions between versions for a single resource/type
type ResourceConversionGraph struct {
	name  string                   // Name of the resource needing conversions
	links astmodel.TypeAssociation // All the directed links in our conversion graph
}

// LookupTransition accepts a type name and looks up the transition to the next version in the graph
// Returns the next version, or an empty type name if not.
func (graph *ResourceConversionGraph) LookupTransition(name astmodel.InternalTypeName) astmodel.InternalTypeName {
	next := graph.links[name]
	return next
}

// TransitionCount returns the number of transitions in the graph
func (graph *ResourceConversionGraph) TransitionCount() int {
	return len(graph.links)
}

// WriteTo gives a debug dump of the conversion graph for a particular type name
func (graph *ResourceConversionGraph) WriteTo(writer io.Writer) error {
	linkStarts := maps.Keys(graph.links)
	slices.SortFunc(linkStarts, func(left astmodel.InternalTypeName, right astmodel.InternalTypeName) int {
		l := left.InternalPackageReference().FolderPath()
		r := right.InternalPackageReference().FolderPath()
		if l < r {
			return -1
		}

		if l > r {
			return 1
		}

		return 0
	})

	apiVersions := set.Make[string]()
	storageVersions := set.Make[string]()

	// API versions can be identified because they are only ever the start of a link (never the end)
	// This initial list will have too many items in it, but we'll reduce it down later
	for from := range graph.links {
		ver := from.InternalPackageReference().ImportAlias(astmodel.VersionOnly)
		apiVersions.Add(ver)
	}

	// Everything else is a storage version
	for _, to := range graph.links {
		ver := to.InternalPackageReference().ImportAlias(astmodel.VersionOnly)
		storageVersions.Add(ver)
		apiVersions.Remove(ver)
	}

	// Write out a graphviz file
	err := graph.WriteLines(
		writer,
		"# Render with Graphviz using dot",
		"# or try https://dreampuf.github.io/GraphvizOnline/",
		"graph G {",
		"    rankdir=LR;",
		"",
		"    subgraph apiversions {",
		"        rank=same;",
		"        node [shape=ellipse, style=dashed, penwidth=1, rankType=min, group=storage];",
	)
	if err != nil {
		return eris.Wrapf(err, "writing graph header")
	}

	err = graph.WriteVersions(writer, apiVersions)
	if err != nil {
		return eris.Wrapf(err, "writing graph apiversions")
	}

	err = graph.WriteLines(
		writer,
		"    }",
		"",
		"    subgraph storageversions {",
		"        rank=same;",
		"        node [shape=ellipse, style=dashed, penwidth=1, rankType=min, group=storage];",
	)
	if err != nil {
		return eris.Wrapf(err, "writing graph midsection")
	}

	err = graph.WriteVersions(writer, storageVersions)
	if err != nil {
		return eris.Wrapf(err, "writing graph storageversions")
	}

	err = graph.WriteLines(
		writer,
		"    }",
		"",
		"    edge [arrowhead=vee, arrowtail=vee, dir=forward];",
	)
	if err != nil {
		return eris.Wrapf(err, "writing graph endsection")
	}

	err = graph.writeLinks(writer, linkStarts)
	if err != nil {
		return eris.Wrapf(err, "writing graph links")
	}

	err = graph.WriteLines(
		writer,
		"}",
	)
	if err != nil {
		return eris.Wrapf(err, "writing graph end")
	}

	return nil
}

func (graph *ResourceConversionGraph) writeLinks(writer io.Writer, linkStarts []astmodel.InternalTypeName) error {
	lines := make([]string, 0, len(linkStarts))
	for _, from := range linkStarts {
		to := graph.links[from]

		f := from.InternalPackageReference().ImportAlias(astmodel.VersionOnly)
		t := to.InternalPackageReference().ImportAlias(astmodel.VersionOnly)

		line := fmt.Sprintf("    %s -- %s", f, t)
		lines = append(lines, line)
	}

	return graph.WriteLines(writer, lines...)
}

func (graph *ResourceConversionGraph) WriteLines(writer io.Writer, lines ...string) error {
	for _, line := range lines {
		_, err := io.WriteString(writer, line+"\n")
		if err != nil {
			return eris.Wrapf(err, "writing %q to writer", line)
		}
	}

	return nil
}

func (graph *ResourceConversionGraph) WriteVersions(writer io.Writer, versions set.Set[string]) error {
	lines := versions.Values()
	slices.SortFunc(lines, astmodel.ComparePathAndVersion)

	for i, line := range lines {
		lines[i] = fmt.Sprintf("        %q;", line)
	}

	return graph.WriteLines(writer, lines...)
}

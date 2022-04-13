/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import (
	"strings"
	"testing"

	"github.com/sebdah/goldie/v2"
)

func TestMarkdownTable_GivesExpectedResults(t *testing.T) {
	t.Parallel()
	g := goldie.New(t)

	table := NewMarkdownTable("Name", "Description")
	table.AddRow("Alpha", "The first, be cautious")
	table.AddRow("Beta", "The second, should be stable")
	table.AddRow("Gamma", "The third, general availability")

	var buff strings.Builder
	table.WriteTo(&buff)

	g.Assert(t, t.Name(), []byte(buff.String()))
}

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import (
	"fmt"
	"github.com/sebdah/goldie/v2"
	"strings"
	"testing"
)

func TestTable_StepByStep_GivesExpectedResults(t *testing.T) {
	steps := []struct {
		row  string
		col  string
		cell string
	}{
		{"1", "prime", "(yes)"},
		{"1", "square", "yes"},
		{"2", "prime", "yes"},
		{"3", "prime", "yes"},
		{"3", "triangle", "yes"},
		{"4", "square", "yes"},
		{"5", "prime", "yes"},
		{"6", "triangle", "yes"},
		{"7", "prime", "yes"},
		{"9", "square", "yes"},
		{"10", "triangle", "yes"},
	}

	table := NewTable()
	g := goldie.New(t)
	for i, s := range steps {
		table.SetCell(s.row, s.col, s.cell)

		var buff strings.Builder
		table.WriteTo(&buff)

		testName := fmt.Sprintf("%s_step_%d", t.Name(), i)
		g.Assert(t, testName, []byte(buff.String()))
	}
}

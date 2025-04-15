/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Azure/azure-service-operator/v2/internal/util/match"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

type DebugSettings struct {
	outputFolder  string
	groupSelector match.StringMatcher
}

// NewDebugSettings creates a new DebugSettings.
func NewDebugSettings(groupSelector string, outputFolder string) *DebugSettings {
	return &DebugSettings{
		groupSelector: match.NewStringMatcher(groupSelector),
		outputFolder:  outputFolder,
	}
}

var dashMatcher = regexp.MustCompile("-+")

// createFileName creates the filename for the debug report from the name of the stage, filtering out any characters
// that are unsafe in filenames
func (ds *DebugSettings) CreateFileName(name string) string {
	// filter symbols and other unsafe characters from the stage name to generate a safe filename for the debug log
	n := strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
			return r
		}

		return '-'
	}, name)

	// Replace any sequence of dashes with a single one using a regular expression
	n = dashMatcher.ReplaceAllString(n, "-")

	return filepath.Join(ds.outputFolder, n)
}

// MatchesGroup returns true if the InternalPackageReference matches the groupSelector.
func (ds *DebugSettings) MatchesGroup(ref astmodel.InternalPackageReference) bool {
	grp, ver := ref.GroupVersion()
	return ds.groupSelector.Matches(grp).Matched || ds.groupSelector.Matches(grp+"/"+ver).Matched
}

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	. "github.com/onsi/gomega"
	"runtime"
	"testing"
)

func Test_ShouldSkipDir_GivenPath_HasExpectedResult(t *testing.T) {
	linuxCases := []struct {
		name       string
		path       string
		shouldSkip bool
	}{
		// Simple paths
		{"Root", "/", false},
		{"Top level", "/foo/", false},
		{"Nested", "/foo/bar/", false},
		// Paths to skip
		{"Skip top level", "/examples/", true},
		{"Skip nested", "/foo/examples/", true},
		{"Skip nested, trailing directory", "/foo/examples/bar/", true},
	}

	windowsCases := []struct {
		name       string
		path       string
		shouldSkip bool
	}{
		// Simple paths
		{"Drive", "D:\\", false},
		{"Top level, Windows", "D:\\foo\\", false},
		{"Nested, Windows", "D:\\foo\\bar\\", false},
		// Paths to skip
		{"Skip top level, Windows", "D:\\examples\\", true},
		{"Skip nested, Windows", "D:\\foo\\examples\\", true},
		{"Skip nested, trailing directory, Windows", "D:\\foo\\examples\\bar\\", true},
	}

	cases := linuxCases

	// If testing on Windows, also test Windows paths
	// Can't test Windows paths on Linux because *reasons*
	if runtime.GOOS == "windows" {
		cases = append(cases, windowsCases...)
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			skipped := shouldSkipDir(c.path)

			g.Expect(skipped).To(Equal(c.shouldSkip))
		})
	}
}

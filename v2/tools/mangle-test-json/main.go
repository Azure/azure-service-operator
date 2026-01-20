/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/go-logr/logr"
)

type JSONFormat struct {
	Time    time.Time `json:"Time"`
	Action  string    `json:"Action"`
	Package string    `json:"Package"`
	Test    string    `json:"Test"`
	Output  string    `json:"Output"`
}

func main() {
	log := CreateLogger()

	if len(os.Args) <= 1 {
		log.Info("No log file specified on command line.")
		return
	}

	for _, testOutputFile := range os.Args[1:] {
		log.Info(
			"Parsing",
			"file", testOutputFile)
		fmt.Printf("# `%s`\n\n", testOutputFile)

		byPackage := loadJSON(testOutputFile, log)

		packages := []string{}
		for k, v := range byPackage {
			packages = append(packages, k)
			sort.Slice(v, func(i, j int) bool {
				return v[i].Test.Value() < v[j].Test.Value()
			})
		}

		sort.Strings(packages)

		printSummary(packages, byPackage)
		printDetails(packages, byPackage)
		printSlowTests(byPackage)
	}

	log.Info("Complete")
}

func min(i, j int) int {
	if i <= j {
		return i
	}

	return j
}

func loadJSON(
	testOutputFile string,
	log logr.Logger,
) map[string][]TestRun {
	file, err := os.Open(testOutputFile)
	if err != nil {
		log.Error(
			err,
			"Unable to open file",
			"file", testOutputFile)
		return make(map[string][]TestRun)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	factory := newTestRunFactory()
	errCount := 0
	row := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			// Skip empty lines
			continue
		}

		var d JSONFormat
		err := json.Unmarshal(line, &d)
		if err != nil {
			// Write the line to the log so we don't lose the content
			text := string(line)
			log.Info(
				"Unable to parse",
				"line", text)

			if text != "" && !strings.HasPrefix(text, "FAIL") {
				// It's a parse failure we care about, write details
				logError(log, err, row, text)
				errCount++
			}

			continue
		}

		factory.apply(d)
		row++
	}

	if err := scanner.Err(); err != nil {
		log.Error(
			err,
			"Error reading file",
			"file", testOutputFile)
	}

	if errCount > 0 {
		log.Info(
			"Errors parsing JSON",
			"count", errCount)
	}

	// package → list of tests
	byPackage := make(map[string][]TestRun)
	for pkg, pkgRuns := range factory.testRuns {
		var runs []TestRun
		for _, r := range pkgRuns {
			runs = append(runs, *r)
		}

		byPackage[pkg] = runs
	}

	return byPackage
}

func sensitiveRound(d time.Duration) time.Duration {
	if d > time.Minute {
		return d.Round(time.Second)
	}

	return d.Round(time.Millisecond)
}

func printSummary(packages []string, byPackage map[string][]TestRun) {
	fmt.Printf("## Package Summary\n\n")

	commonPrefix := findCommonPackagePrefix(packages)

	fmt.Printf("| Result | %s | Time |\n", commonPrefix)
	fmt.Printf("|--------|:---|-----:|\n")

	// output table-of-contents
	for _, pkg := range packages {
		tests := byPackage[pkg]
		if len(tests) == 1 {
			// package-only
			continue
		}

		totalRuntime := time.Duration(0)
		for _, t := range tests[1:] {
			totalRuntime += t.RunTime
		}

		overallOutcome := tests[0].actionSymbol()
		shortPkgName := displayNameForPackage(pkg, commonPrefix)
		totalRuntime = sensitiveRound(totalRuntime)

		fmt.Printf("| %s | %s | %s |\n", overallOutcome, shortPkgName, totalRuntime)
	}

	fmt.Println()
}

var maxOutputLines = 300

func printDetails(packages []string, byPackage map[string][]TestRun) {
	anyFailed := false

	for _, pkg := range packages {
		tests := byPackage[pkg]
		// check package-level indicator, which will be first ("" test name):
		if tests[0].Action != Failed {
			continue // no failed tests, skip
		} else {
			anyFailed = true
		}

		// package name as header
		fmt.Printf("### Package `%s` failed tests\n\n", pkg)

		// Output info on stderr
		fmt.Fprintf(os.Stderr, "Package failed: %s\n", pkg)

		{ // print package-level logs, if any
			packageLevel := tests[0]
			if len(packageLevel.Output) > 0 {
				trimmedOutput, output := escapeOutput(packageLevel.Output)
				summary := "Package-level output"
				if trimmedOutput {
					summary += fmt.Sprintf(" (trimmed to last %d lines) — full details available in log", maxOutputLines)
				}
				fmt.Printf("<details><summary>%s</summary><pre>%s</pre></details>\n\n", summary, output)

				// Output info on stderr, so that package failure isn’t silent on console
				// when running `task ci`, and that full logs are available if they get trimmed
				fmt.Fprintln(os.Stderr, "=== PACKAGE OUTPUT ===")
				for _, line := range packageLevel.Output {
					fmt.Fprint(os.Stderr, line)
				}
				fmt.Fprintln(os.Stderr, "=== END PACKAGE OUTPUT ===")
			}
		}

		for _, test := range tests[1:] {
			// We only want to include "interesting" tests in the output
			// "interesting" tests are ones that failed, or had a panic, or we don't know what status they have
			if !test.IsInteresting() {
				continue
			}

			fmt.Printf("#### Test `%s`\n", test.Test.Value())

			if test.Action == Failed {
				fmt.Printf("Failed in %s:\n", test.RunTime)
			} else {
				fmt.Printf("Elapsed %s:\n", test.RunTime)
				fmt.Printf("Final action %s:\n", test.Action)
			}

			trimmedOutput, output := escapeOutput(test.Output)
			summary := "Test output"
			if trimmedOutput {
				summary += fmt.Sprintf(" (trimmed to last %d lines) — full details available in log", maxOutputLines)
			}

			fmt.Printf("<details><summary>%s</summary><pre>%s</pre></details>\n\n", summary, output)

			// Output info on stderr, so that test failure isn’t silent on console
			// when running `task ci`, and that full logs are available if they get trimmed
			fmt.Fprintf(os.Stderr, "- Test failed: %s\n", test.Test.Value())
			fmt.Fprintln(os.Stderr, "=== TEST OUTPUT ===")
			for _, outputLine := range test.Output {
				fmt.Fprint(os.Stderr, outputLine) // note that line already has newline attached
			}
			fmt.Fprintln(os.Stderr, "=== END TEST OUTPUT ===")
		}

		fmt.Println()
	}

	if !anyFailed {
		fmt.Printf("**🎉 All tests passed. 🎉**\n\n")
	}
}

func escapeOutput(outputs []string) (bool, string) {
	trimmed := false
	if len(outputs) > maxOutputLines {
		outputs = outputs[len(outputs)-maxOutputLines:]
		trimmed = true
	}

	result := strings.Builder{}
	for _, output := range outputs {
		s := output
		s = strings.ReplaceAll(s, "&", "&amp;")
		s = strings.ReplaceAll(s, "<", "&lt;")
		s = strings.ReplaceAll(s, "\n", "<br>")
		result.WriteString(s)
	}

	return trimmed, result.String()
}

func printSlowTests(byPackage map[string][]TestRun) {
	fmt.Printf("## Longest-running tests\n\n")

	allTests := []TestRun{}
	for _, v := range byPackage {
		allTests = append(allTests, v[1:]...)
	}

	sort.Slice(allTests, func(i, j int) bool {
		return allTests[i].RunTime > allTests[j].RunTime
	})

	pkgPrefix := allTests[0].Package.Value()
	for _, test := range allTests {
		pkgPrefix = commonPrefix(pkgPrefix, test.Package.Value())
	}

	pkgPrefix = strings.TrimRight(pkgPrefix, "/")

	fmt.Printf("| %s | Name | Time |\n", pkgPrefix)
	fmt.Println("|---------|------|-----:|")
	for i := 0; i < min(10, len(allTests)); i += 1 {
		test := allTests[i]
		pkg := displayNameForPackage(test.Package.Value(), pkgPrefix)
		fmt.Printf("| `%s` | `%s` | %s |\n", pkg, test.Test.Value(), test.RunTime)
	}
}

func logError(
	log logr.Logger,
	err error,
	row int,
	line string,
) {
	// If syntax error, log it and return
	var syntaxError *json.SyntaxError
	if errors.As(err, &syntaxError) {
		log.Error(
			err,
			"Syntax error parsing JSON",
			"row", row,
			"column", syntaxError.Offset,
			"line", line,
		)

		return
	}

	// If unmarshal type error, log it and return
	var unmarshalError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalError) {
		log.Error(
			err,
			"Unmarshal type error parsing JSON",
			"row", row,
			"column", unmarshalError.Offset,
			"line", line,
		)

		return
	}

	// Otherwise unknown error, log it and return
	log.Error(
		err,
		"Unexpected error parsing JSON",
		"row", row,
		"line", line,
	)
}

func commonPrefix(
	left string,
	right string,
) string {
	minLen := min(len(left), len(right))
	i := 0
	for i < minLen && left[i] == right[i] {
		i++
	}

	return left[:i]
}

// findCommonPackagePrefix finds the common prefix of all package names, to shorten display
func findCommonPackagePrefix(packages []string) string {
	prefix := packages[0]
	for _, pkg := range packages {
		prefix = commonPrefix(prefix, pkg)
	}
	prefix = strings.TrimRight(prefix, "/")
	return prefix
}

func displayNameForPackage(
	pkg string,
	commonPrefix string,
) string {
	name := strings.TrimPrefix(pkg, commonPrefix)
	name = strings.TrimLeft(name, "/")
	return name
}

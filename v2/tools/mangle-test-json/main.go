/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
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
				return v[i].Test < v[j].Test
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

func actionSymbol(d TestRun) string {
	switch d.Action {
	case "pass":
		return "✅"
	case "fail":
		return "❌"
	case "skip":
		return "⏭️"
	default:
		panic(fmt.Sprintf("unhandled action: %s", d.Action))
	}
}

func loadJSON(
	testOutputFile string,
	log logr.Logger,
) map[string][]TestRun {
	content, err := os.ReadFile(testOutputFile)
	if err != nil {
		log.Error(
			err,
			"Unable to read file",
			"file", testOutputFile)
	}

	// Break into individual lines to make error reporting easier
	lines := strings.Split(string(content), "\n")

	data := make([]JSONFormat, 0, len(lines))
	errCount := 0
	for row, line := range lines {
		if len(line) == 0 {
			// Skip empty lines
			continue
		}

		var d JSONFormat
		err := json.Unmarshal([]byte(line), &d)
		if err != nil {
			// Write the line to the log so we don't lose the content
			log.Info(
				"Unable to parse",
				"line", line)

			if line != "" && !strings.HasPrefix(line, "FAIL") {
				// It's a parse failure we care about, write details
				logError(log, err, row, line)
				errCount++
			}

			continue
		}

		data = append(data, d)
	}

	if errCount > 0 {
		log.Info(
			"Errors parsing JSON",
			"count", errCount)
	}

	// Track all the test runs
	testRuns := make(map[string]*TestRun)
	for _, d := range data {

		// Find (or create) the test run for this item of data
		testrun, found := testRuns[d.key()]
		if !found {
			testrun = &TestRun{
				Package: d.Package,
				Test:    d.Test,
			}

			testRuns[d.key()] = testrun
		}

		switch d.Action {
		case "run":
			testrun.run(d.Time)

		case "pause":
			testrun.pause(d.Time)

		case "cont":
			testrun.resume(d.Time)

		case "output":
			testrun.output(d.Output)

		case "pass", "fail", "skip":
			testrun.complete(d.Action, d.Time)
		}
	}

	// package → list of tests
	byPackage := make(map[string][]TestRun, len(data))
	for _, v := range testRuns {
		byPackage[v.Package] = append(byPackage[v.Package], *v)
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

		overallOutcome := actionSymbol(tests[0])
		fmt.Printf("* %s `%s` (runtime %s)\n", overallOutcome, pkg, totalRuntime)
	}

	fmt.Println()
}

var maxOutputLines = 300

func printDetails(packages []string, byPackage map[string][]TestRun) {
	anyFailed := false

	for _, pkg := range packages {
		tests := byPackage[pkg]
		// check package-level indicator, which will be first ("" test name):
		if tests[0].Action != "fail" {
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

			fmt.Printf("#### Test `%s`\n", test.Test)

			if test.Action == "fail" {
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
			fmt.Fprintf(os.Stderr, "- Test failed: %s\n", test.Test)
			fmt.Fprintln(os.Stderr, "=== TEST OUTPUT ===")
			for _, outputLine := range test.Output {
				fmt.Fprint(os.Stderr, outputLine) // note that line already has newline attached
			}
			fmt.Fprintln(os.Stderr, "=== END TEST OUTPUT ===")
		}

		fmt.Println()
	}

	if !anyFailed {
		fmt.Println("**🎉 All tests passed. 🎉**")
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

	fmt.Println("| Package | Name | Time |")
	fmt.Println("|---------|------|-----:|")
	for i := 0; i < min(10, len(allTests)); i += 1 {
		test := allTests[i]
		fmt.Printf("| `%s` | `%s` | %s |\n", test.Package, test.Test, test.RunTime)
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

// key returns a unique key for a test run
func (d JSONFormat) key() string {
	return d.Package + "/" + d.Test
}

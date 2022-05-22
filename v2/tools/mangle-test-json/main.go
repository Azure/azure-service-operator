/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type JSONFormat struct {
	Time    time.Time
	Action  string
	Package string
	Test    string
	Output  string
}

type TestRun struct {
	Action  string
	Package string
	Test    string
	Output  []string
	RunTime time.Duration
}

func main() {
	for _, testOutputFile := range os.Args[1:] {
		log.Printf("Parsing  %s\n\n", testOutputFile)
		fmt.Printf("# `%s`\n\n", testOutputFile)

		byPackage := loadJSON(testOutputFile)

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
	log.Println("Complete.")
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

func loadJSON(testOutputFile string) map[string][]TestRun {
	content, err := ioutil.ReadFile(testOutputFile)
	if err != nil {
		log.Fatalf("Unable to read file: %e", err)
	}

	// Break into individual lines to make error reporting easier
	lines := strings.Split(string(content), "\n")

	var data []JSONFormat
	errCount := 0
	for row, line := range lines {
		var d JSONFormat
		err := json.Unmarshal([]byte(line), &d)
		if err != nil {
			// Write the line to the log so we don't lose the content
			log.Println(line)
			if line != "" && !strings.HasPrefix(line, "FAIL") {
				// It's a parse failure we care about, write details
				logError(err, row, line)
				errCount++
			}

			continue
		}

		data = append(data, d)
	}

	if errCount > 0 {
		log.Fatalf("%d fatal error(s) parsing JSON", errCount)
	}

	// track when each test started running
	startTimes := make(map[string]time.Time)
	runTimes := make(map[string]time.Duration)
	outputs := make(map[string][]string)
	key := func(d JSONFormat) string {
		return d.Package + "/" + d.Test
	}

	// package → list of tests
	byPackage := make(map[string][]TestRun, len(data))
	for _, d := range data {
		switch d.Action {
		case "run":
			if startTimes[key(d)] != (time.Time{}) {
				panic("run while already running")
			}
			startTimes[key(d)] = d.Time
		case "pause":
			if startTimes[key(d)] == (time.Time{}) {
				panic("pause while not running")
			}
			runTimes[key(d)] += d.Time.Sub(startTimes[key(d)])
			startTimes[key(d)] = time.Time{}
		case "cont":
			// cont while still in running state happens sometimes (???)
			// so don't check
			startTimes[key(d)] = d.Time
		case "output":
			outputs[key(d)] = append(outputs[key(d)], d.Output)
		case "pass", "fail", "skip":
			if d.Test != "" && startTimes[key(d)] == (time.Time{}) {
				panic("finished when not running")
			}

			runTimes[key(d)] += d.Time.Sub(startTimes[key(d)])

			byPackage[d.Package] = append(byPackage[d.Package], TestRun{
				Action:  d.Action,
				Package: d.Package,
				Test:    d.Test,
				Output:  outputs[key(d)],

				// round all runtimes to ms to avoid excessive decimal places
				RunTime: sensitiveRound(runTimes[key(d)]),
			})
		}
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
			// only printing failed tests
			if test.Action == "fail" {

				fmt.Printf("#### Test `%s`\n", test.Test)
				fmt.Printf("Failed in %s:\n", test.RunTime)

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
		for _, t := range v[1:] { // skip "" package test
			allTests = append(allTests, t)
		}
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

func logError(err error, row int, line string) {
	var syntaxError *json.SyntaxError
	if errors.As(err, &syntaxError) {
		log.Printf(
			"Syntax error parsing JSON on line %d at column %d: %s (line: %q)",
			row,
			syntaxError.Offset,
			syntaxError.Error(),
			line)
	}

	var unmarshalError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalError) {
		log.Printf(
			"Unmarshal type error parsing JSON on line %d at column %d: %s (near: %q)",
			row,
			unmarshalError.Offset,
			unmarshalError.Error(),
			line)
	}

	log.Printf("Unexpected error parsing JSON: %s", err)
}

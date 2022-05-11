/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"encoding/json"
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
		log.Fatalf("%e", err)
	}

	// make test output into valid JSON
	jsonData := "[" + strings.Join(strings.Split(strings.Trim(string(content), " \n\r"), "\n"), ",") + "]"

	data := []JSONFormat{}
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Fatalf("%e", err)
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
			if startTimes[key(d)] != (time.Time{}) {
				panic("cont while still running")
			}
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
				RunTime: runTimes[key(d)],
			})
		}
	}

	return byPackage
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

		overallOutcome := actionSymbol(tests[0])
		fmt.Printf("* %s `%s`\n", overallOutcome, pkg)
	}

	fmt.Println()
}

func printDetails(packages []string, byPackage map[string][]TestRun) {
	fmt.Printf("## Details\n\n")

	for _, pkg := range packages {
		tests := byPackage[pkg]
		if len(tests) == 1 {
			// skip package-only output
			continue
		}

		// first test is the "" test for overall package outcome
		overallOutcome := actionSymbol(tests[0])
		fmt.Printf("### %s `%s`\n\n", overallOutcome, pkg)

		fmt.Println("| Outcome | Name | Time | Output |")
		fmt.Println("|---------|------|-----:|--------|")
		for _, test := range tests[1:] {
			output := `<details><pre>` + escapeOutput(test.Output) + `</pre></details>`
			fmt.Printf("| %s | `%s` | %s | %s |\n", actionSymbol(test), test.Test, test.RunTime, output)
		}

		fmt.Println()
	}
}

func escapeOutput(outputs []string) string {
	result := strings.Builder{}
	for _, output := range outputs {
		s := output
		s = strings.ReplaceAll(s, "&", "&amp;")
		s = strings.ReplaceAll(s, "<", "&lt;")
		s = strings.ReplaceAll(s, "\n", "<br>")
		result.WriteString(s)
	}

	return result.String()
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

package test

import (
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// AssertionOption represents options that can be used to customize the behavour of our custom render methods that
// generate golden file output.
// Diff* options highlight differences between reference and actual using traditional +/- notations on each line.
// Exclude* options pare down the objects being rendered to reduce the amount of noise present in the golden file.
type AssertionOption interface {
	// configure applies this option to the passed typeAsserter
	configure(ta *typeAsserter)
}

/*
 * diffOption
 */

// DiffWith specifies the types being tested should be diff'd with the provided types to highlight differences.
// Types are matched by fully qualified name.
func DiffWith(defs ...astmodel.TypeDefinition) AssertionOption {
	return &diffOption{
		references: defs,
	}
}

// diffOption captures a GoSourceFile which acts as a base for comparison
type diffOption struct {
	references []astmodel.TypeDefinition
}

var _ AssertionOption = &diffOption{}

func (d *diffOption) configure(ta *typeAsserter) {
	ta.addReferences(d.references...)
}

/*
 * includeTestFiles
 */

func IncludeTestFiles() AssertionOption {
	return &includeTestFiles{}
}

type includeTestFiles struct {
}

var _ AssertionOption = &includeTestFiles{}

func (i *includeTestFiles) configure(ta *typeAsserter) {
	ta.writeTests = true
}

/*
 * excludeCodeFiles
 */

func ExcludeCodeFiles() AssertionOption {
	return &excludeCodeFiles{}
}

type excludeCodeFiles struct {
}

var _ AssertionOption = &excludeCodeFiles{}

func (i *excludeCodeFiles) configure(ta *typeAsserter) {
	ta.writeCode = false
}

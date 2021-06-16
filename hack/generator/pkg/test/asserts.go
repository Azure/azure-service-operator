package test

import (
	"bytes"
	"testing"

	"github.com/sebdah/goldie/v2"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// AssertFileGeneratesExpectedCode serialises the given FileDefintion as a golden file test, checking that the expected
// results are generated
func AssertFileGeneratesExpectedCode(t *testing.T, fileDef *astmodel.FileDefinition, testName string) {
	g := goldie.New(t)

	buf := &bytes.Buffer{}
	fileWriter := astmodel.NewGoSourceFileWriter(fileDef)
	err := fileWriter.SaveToWriter(buf)
	if err != nil {
		t.Fatalf("could not generate file: %v", err)
	}

	g.Assert(t, testName, buf.Bytes())
}

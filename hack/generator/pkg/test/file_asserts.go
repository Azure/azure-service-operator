package test

import (
	"bytes"
	"testing"

	"github.com/sebdah/goldie/v2"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// AssertFileGeneratesExpectedCode serialises the given FileDefinition as a golden file test, checking that the expected
// results are generated
func AssertFileGeneratesExpectedCode(
	t *testing.T,
	fileDef astmodel.GoSourceFile,
	fileName string,
	options ...AssertFileOption) {
	g := goldie.New(t)
	err := g.WithTestNameForDir(true)
	if err != nil {
		t.Fatalf("Unable to configure goldie output folder %s", err)
	}

	buf := &bytes.Buffer{}
	fileWriter := astmodel.NewGoSourceFileWriter(fileDef)
	err = fileWriter.SaveToWriter(buf)
	if err != nil {
		t.Fatalf("could not generate file: %s", err)
	}

	g.Assert(t, fileName, buf.Bytes())
}

type AssertFileOption interface {
	option()
}


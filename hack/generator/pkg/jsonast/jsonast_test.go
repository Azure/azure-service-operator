/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/sebdah/goldie/v2"
	"github.com/xeipuuv/gojsonschema"
)

func runGoldenTest(t *testing.T, path string) {
	testName := strings.TrimPrefix(t.Name(), "TestGolden/")

	g := goldie.New(t)
	inputFile, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(fmt.Errorf("Cannot read golden test input file: %w", err))
	}

	loader := gojsonschema.NewSchemaLoader()
	schema, err := loader.Compile(gojsonschema.NewBytesLoader(inputFile))

	if err != nil {
		t.Fatal(fmt.Errorf("could not compile input: %w", err))
	}

	scanner := NewSchemaScanner(astmodel.NewIdentifierFactory())
	defs, err := scanner.GenerateDefinitions(context.TODO(), schema.Root())
	if err != nil {
		t.Fatal(fmt.Errorf("could not produce nodes from scanner: %w", err))
	}

	// put all definitions in one file, regardless
	// the package reference isn't really used here
	fileDef := astmodel.NewFileDefinition(defs[0].Reference().PackageReference, defs...)

	buf := &bytes.Buffer{}
	err = fileDef.SaveToWriter(path, buf)
	if err != nil {
		t.Fatal(fmt.Errorf("could not generate file: %w", err))
	}

	g.Assert(t, testName, buf.Bytes())
}

func TestGolden(t *testing.T) {

	type Test struct {
		name string
		path string
	}

	testGroups := make(map[string][]Test)

	// find all input .json files
	err := filepath.Walk("testdata", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			groupName := filepath.Base(filepath.Dir(path))
			testName := strings.TrimSuffix(filepath.Base(path), ".json")
			testGroups[groupName] = append(testGroups[groupName], Test{testName, path})
		}

		return nil
	})

	if err != nil {
		t.Fatal(fmt.Errorf("Error enumerating files: %w", err))
	}

	// run all tests
	for groupName, fs := range testGroups {
		t.Run(groupName, func(t *testing.T) {
			for _, f := range fs {
				t.Run(f.name, func(t *testing.T) {
					runGoldenTest(t, f.path)
				})
			}
		})
	}
}

/*
func TestToNodes(t *testing.T) {
	type args struct {
		resourcesSchema *gojsonschema.SubSchema
		opts            []BuilderOption
	}
	tests := []struct {
		name        string
		argsFactory func(*testing.T) *args
		want        []*ast.Package
		wantErr     bool
	}{
		{
			name: "WithSchema",
			argsFactory: func(t *testing.T) *args {
				schema, err := getDefaultSchema()
				if err != nil {
					t.Error(err)
				}

				return &args{
					resourcesSchema: schema,
				}
			},
			want:    nil,
			wantErr: false,
		},
	}
	scanner := NewSchemaScanner()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg := tt.argsFactory(t)
			got, err := scanner.ToNodes(context.TODO(), arg.resourcesSchema.ItemsChildren[0], arg.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToNodes() got = %+v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectWithNoType(t *testing.T) {
	schema := `
{
  "title": "bar",
  "definitions": {
    "ApplicationSecurityGroupPropertiesFormat": {
      "description": "Application security group properties."
    }
  },

  "type": "object",
  "properties": {
    "foo": {
      "$ref": "#/definitions/ApplicationSecurityGroupPropertiesFormat"
    }
  }
}
`
	idFactory := astmodel.NewIdentifierFactory()
	scanner := NewSchemaScanner(idFactory)
	g := NewGomegaWithT(t)
	sl := gojsonschema.NewSchemaLoader()
	loader := gojsonschema.NewBytesLoader([]byte(schema))
	sb, err := sl.Compile(loader)
	g.Expect(err).To(BeNil())

	definition, err := scanner.ToNodes(context.TODO(), sb.Root())

	g.Expect(err).To(BeNil())
	g.Expect(definition).ToNot(BeNil())
	g.Expect(scanner.Structs).To(HaveLen(1))

	// Has no version!!
	structDefinition := scanner.Structs["bar/"]
	g.Expect(structDefinition).ToNot(BeNil())
	g.Expect(structDefinition.FieldCount()).To(Equal(1))

	propertyField := structDefinition.Field(0)
	g.Expect(propertyField.FieldName()).To(Equal("Foo"))

	g.Expect(propertyField.FieldType()).To(Equal(astmodel.AnyType))
}

func XTestAnyOfWithMultipleComplexObjects(t *testing.T) {
	schema := `
{
  "definitions": {
    "genericExtension": {
      "type": "object",
      "properties": {
        "publisher": {
          "type": "string",
          "minLength": 1,
          "description": "Microsoft.Compute/extensions - Publisher"
        }
      }
    },
    "iaaSDiagnostics": {
      "type": "object",
      "properties": {
        "publisher": {
          "enum": [
            "Microsoft.Azure.Diagnostics"
          ]
        }
      }
    }
  },

  "type": "object",
  "properties": {
	"properties": {
	  "anyOf": [
		{
		  "$ref": "#/definitions/genericExtension"
		},
		{
		  "$ref": "#/definitions/iaaSDiagnostics"
		}
	  ]
	}
  },
  "description": "Microsoft.Compute/virtualMachines/extensions"
}
`
	scanner := &SchemaScanner{}
	g := NewGomegaWithT(t)
	sl := gojsonschema.NewSchemaLoader()
	loader := gojsonschema.NewBytesLoader([]byte(schema))
	sb, err := sl.Compile(loader)
	g.Expect(err).To(BeNil())
	nodes, err := scanner.ToNodes(context.TODO(), sb.Root())
	g.Expect(err).To(BeNil())
	g.Expect(nodes).To(HaveLen(1))
	structType, ok := nodes[0].(*ast.StructType)
	g.Expect(ok).To(BeTrue())
	g.Expect(structType.Fields.List).To(HaveLen(1))
	propertiesField := structType.Fields.List[0]
	g.Expect(propertiesField.Names[0]).To(Equal(ast.NewIdent("properties")))
	g.Expect(propertiesField.Type).To(Equal(ast.NewIdent("interface{}")))
}

func TestOneOfWithPropertySibling(t *testing.T) {
	schema := `
{
  "def": {
    "type": "object",
    "oneOf": [
      {
        "properties": {
          "ruleSetType": {
            "oneOf": [
              {
                "type": "string",
                "enum": [
                  "AzureManagedRuleSet"
                ]
              },
              {
                "$ref": "https://schema.management.azure.com/schemas/common/definitions.json#/definitions/expression"
              }
            ]
          }
        }
      }
    ],
    "properties": {
      "ruleSetType": {
        "type": "string"
      }
    },
    "required": [
      "ruleSetType"
    ],
    "description": "Describes azure managed provider."
  },
  "type": "object",
  "properties": {
    "ruleSets": {
      "oneOf": [
        {
          "type": "array",
          "items": {
            "$ref": "#/def"
          }
        },
        {
          "$ref": "https://schema.management.azure.com/schemas/common/definitions.json#/definitions/expression"
        }
      ],
      "description": "List of rules"
    }
  },
  "description": "Defines ManagedRuleSets - array of managedRuleSet"
}
`
	scanner := &SchemaScanner{}
	g := NewGomegaWithT(t)
	sl := gojsonschema.NewSchemaLoader()
	loader := gojsonschema.NewBytesLoader([]byte(schema))
	sb, err := sl.Compile(loader)
	g.Expect(err).To(BeNil())
	nodes, err := scanner.ToNodes(context.TODO(), sb.Root())
	g.Expect(err).To(BeNil())
	g.Expect(nodes).To(HaveLen(1))
	structType, ok := nodes[0].(*ast.StructType)
	g.Expect(ok).To(BeTrue())
	g.Expect(structType.Fields.List).To(HaveLen(1))
}

func TestAllOfUnion(t *testing.T) {
	schema := `{
  "definitions": {
    "address": {
      "type": "object",
      "properties": {
        "street_address": { "type": "string" },
        "city":           { "type": "string" },
        "state":          { "type": "string" }
      },
      "required": ["street_address", "city", "state"]
    }
  },

  "allOf": [
    { "$ref": "#/definitions/address" },
    { "properties": {
        "type": { "enum": [ "residential", "business" ] }
      }
    }
  ]
}`
	scanner := &SchemaScanner{}
	g := NewGomegaWithT(t)
	sl := gojsonschema.NewSchemaLoader()
	loader := gojsonschema.NewBytesLoader([]byte(schema))
	sb, err := sl.Compile(loader)
	g.Expect(err).To(BeNil())
	nodes, err := scanner.ToNodes(context.TODO(), sb.Root())
	g.Expect(err).To(BeNil())
	g.Expect(nodes).To(HaveLen(1))
	structType, ok := nodes[0].(*ast.StructType)
	g.Expect(ok).To(BeTrue())
	g.Expect(structType.Fields.List).To(HaveLen(4))
}

func TestAnyOfLocation(t *testing.T) {
	schema := `
{
"anyOf": [
	{
	  "type": "string"
	},
	{
	  "enum": [
		"East Asia",
		"Southeast Asia",
		"Central US"
	  ]
	}
  ]
}`
	scanner := &SchemaScanner{}
	g := NewGomegaWithT(t)
	sl := gojsonschema.NewSchemaLoader()
	loader := gojsonschema.NewBytesLoader([]byte(schema))
	sb, err := sl.Compile(loader)
	g.Expect(err).To(BeNil())
	nodes, err := scanner.ToNodes(context.TODO(), sb.Root())
	g.Expect(err).To(BeNil())
	g.Expect(nodes).To(HaveLen(1))
	field, ok := nodes[0].(*ast.Field)
	g.Expect(ok).To(BeTrue())
	g.Expect(field.Names[0].Name).To(Equal("anyOf"))
	g.Expect(field.Type.(*ast.Ident)).To(Equal(ast.NewIdent("string")))
}

func getDefaultSchema() (*gojsonschema.SubSchema, error) {
	sl := gojsonschema.NewSchemaLoader()
	ref := "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json"
	schema, err := sl.Compile(gojsonschema.NewReferenceLoader(ref))
	if err != nil {
		return nil, err
	}

	root := schema.Root()
	for _, child := range root.PropertiesChildren {
		if child.Property == "resources" {
			return child, nil
		}
	}
	return nil, errors.New("couldn't find resources in the schema")
}

*/

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"testing"

	"github.com/go-openapi/spec"
	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

func Test_CanExtractTypeNameFromSameFile(t *testing.T) {
	g := NewGomegaWithT(t)

	schemaPath := "/dev/null/path/to/schema.json"
	schemaPackage := astmodel.MakeLocalPackageReference(
		"github.com/Azure/azure-service-operator/v2/api",
		"Microsoft.Test",
		"v1")

	loader := NewCachingFileLoader(map[string]PackageAndSwagger{
		schemaPath: {
			Package: &schemaPackage,
			Swagger: spec.Swagger{
				SwaggerProps: spec.SwaggerProps{
					Definitions: spec.Definitions{
						"TheDefinition": spec.Schema{ /* contents not used */ },
					},
				},
			},
		},
	})

	schema := spec.Schema{
		SchemaProps: spec.SchemaProps{
			Ref: spec.MustCreateRef("#/definitions/TheDefinition"),
		},
	}

	wrappedSchema := MakeOpenAPISchema(
		schema,
		schemaPath,
		schemaPackage,
		astmodel.NewIdentifierFactory(),
		loader)

	typeName, err := wrappedSchema.refTypeName()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(typeName).To(Equal(astmodel.MakeTypeName(schemaPackage, "TheDefinition")))
}

func Test_CanExtractTypeNameFromDifferentFile_AndInheritPackage(t *testing.T) {
	g := NewGomegaWithT(t)

	schemaPath := "/dev/null/path/to/schema.json"
	schemaPackage := astmodel.MakeLocalPackageReference(
		"github.com/Azure/azure-service-operator/v2/api",
		"Microsoft.Test",
		"v1")

	externalSchemaPath := "/dev/null/path/to/other.json"

	loader := NewCachingFileLoader(map[string]PackageAndSwagger{
		schemaPath: {
			Package: &schemaPackage,
			Swagger: spec.Swagger{
				SwaggerProps: spec.SwaggerProps{
					Definitions: spec.Definitions{
						"TheDefinition": spec.Schema{ /* contents not used */ },
					},
				},
			},
		},

		externalSchemaPath: {
			Package: nil, // no package set; this means it will inherit use-site’s package
			Swagger: spec.Swagger{
				SwaggerProps: spec.SwaggerProps{
					Definitions: spec.Definitions{
						"ExternalDefinition": spec.Schema{ /* contents not used */ },
					},
				},
			},
		},
	})

	schema := spec.Schema{
		SchemaProps: spec.SchemaProps{
			Ref: spec.MustCreateRef(externalSchemaPath + "#/definitions/ExternalDefinition"),
		},
	}

	wrappedSchema := MakeOpenAPISchema(
		schema,
		schemaPath,
		schemaPackage,
		astmodel.NewIdentifierFactory(),
		loader)

	typeName, err := wrappedSchema.refTypeName()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(typeName).To(Equal(astmodel.MakeTypeName(schemaPackage, "ExternalDefinition")))
}

func Test_CanExtractTypeNameFromDifferentFile_AndUsePresetPackage(t *testing.T) {
	g := NewGomegaWithT(t)

	schemaPath := "/dev/null/path/to/schema.json"
	schemaPackage := astmodel.MakeLocalPackageReference(
		"github.com/Azure/azure-service-operator/v2/api",
		"Microsoft.Test",
		"v1")

	externalSchemaPath := "/dev/null/path/to/other.json"
	externalPackage := astmodel.MakeLocalPackageReference(
		"github.com/Azure/azure-service-operator/v2/api",
		"Microsoft.Common",
		"v1")

	loader := NewCachingFileLoader(map[string]PackageAndSwagger{
		schemaPath: {
			Package: &schemaPackage,
			Swagger: spec.Swagger{
				SwaggerProps: spec.SwaggerProps{
					Definitions: spec.Definitions{
						"TheDefinition": spec.Schema{ /* contents not used */ },
					},
				},
			},
		},

		externalSchemaPath: {
			Package: &externalPackage,
			Swagger: spec.Swagger{
				SwaggerProps: spec.SwaggerProps{
					Definitions: spec.Definitions{
						"ExternalDefinition": spec.Schema{ /* contents not used */ },
					},
				},
			},
		},
	})

	schema := spec.Schema{
		SchemaProps: spec.SchemaProps{
			Ref: spec.MustCreateRef(externalSchemaPath + "#/definitions/ExternalDefinition"),
		},
	}

	wrappedSchema := MakeOpenAPISchema(
		schema,
		schemaPath,
		schemaPackage,
		astmodel.NewIdentifierFactory(),
		loader)

	typeName, err := wrappedSchema.refTypeName()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(typeName).To(Equal(astmodel.MakeTypeName(externalPackage, "ExternalDefinition")))
}

func Test_GeneratingCollidingTypeNamesReturnsError(t *testing.T) {
	g := NewGomegaWithT(t)

	schemaPath := "/dev/null/path/to/schema.json"
	schemaPackage := astmodel.MakeLocalPackageReference(
		"github.com/Azure/azure-service-operator/v2/api",
		"Microsoft.Test",
		"v1")

	externalSchemaPath := "/dev/null/path/to/other.json"

	loader := NewCachingFileLoader(map[string]PackageAndSwagger{
		schemaPath: {
			Package: &schemaPackage,
			Swagger: spec.Swagger{
				SwaggerProps: spec.SwaggerProps{
					Definitions: spec.Definitions{
						"TheDefinition": spec.Schema{ /* contents not used */ },
					},
				},
			},
		},

		externalSchemaPath: {
			Package: nil,
			Swagger: spec.Swagger{
				SwaggerProps: spec.SwaggerProps{
					Definitions: spec.Definitions{
						"TheDefinition": spec.Schema{ /* contents not used */ },
					},
				},
			},
		},
	})

	schema := spec.Schema{
		SchemaProps: spec.SchemaProps{
			Ref: spec.MustCreateRef(externalSchemaPath + "#/definitions/TheDefinition"),
		},
	}

	wrappedSchema := MakeOpenAPISchema(
		schema,
		schemaPath,
		schemaPackage,
		astmodel.NewIdentifierFactory(),
		loader)

	_, err := wrappedSchema.refTypeName()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err).To(MatchError("importing type TheDefinition from file /dev/null/path/to/other.json into package github.com/Azure/azure-service-operator/v2/api/Microsoft.Test/v1 could generate collision with type in /dev/null/path/to/schema.json"))
}

func Test_GeneratingCollidingTypeNamesWithSiblingFilesReturnsError(t *testing.T) {
	g := NewGomegaWithT(t)

	schemaPath := "/dev/null/path/to/schema.json"
	schemaPackage := astmodel.MakeLocalPackageReference(
		"github.com/Azure/azure-service-operator/v2/api",
		"Microsoft.Test",
		"v1")

	externalSchemaPath := "/dev/null/path/to/other.json"
	siblingSchemaPath := "/dev/null/path/to/sibling.json"

	loader := NewCachingFileLoader(map[string]PackageAndSwagger{
		schemaPath: {
			Package: &schemaPackage,
			Swagger: spec.Swagger{
				SwaggerProps: spec.SwaggerProps{
					Definitions: spec.Definitions{
						"TheDefinition": spec.Schema{ /* contents not used */ },
					},
				},
			},
		},

		externalSchemaPath: {
			Package: nil,
			Swagger: spec.Swagger{
				SwaggerProps: spec.SwaggerProps{
					Definitions: spec.Definitions{
						"ExternalDefinition": spec.Schema{ /* contents not used */ },
					},
				},
			},
		},

		siblingSchemaPath: {
			Package: nil,
			Swagger: spec.Swagger{
				SwaggerProps: spec.SwaggerProps{
					Definitions: spec.Definitions{
						"ExternalDefinition": spec.Schema{ /* contents not used */ },
					},
				},
			},
		},
	})

	schema := spec.Schema{
		SchemaProps: spec.SchemaProps{
			Ref: spec.MustCreateRef(externalSchemaPath + "#/definitions/ExternalDefinition"),
		},
	}

	wrappedSchema := MakeOpenAPISchema(
		schema,
		schemaPath,
		schemaPackage,
		astmodel.NewIdentifierFactory(),
		loader)

	_, err := wrappedSchema.refTypeName()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err).To(MatchError("importing type ExternalDefinition from file /dev/null/path/to/other.json into package github.com/Azure/azure-service-operator/v2/api/Microsoft.Test/v1 could generate collision with type in /dev/null/path/to/sibling.json"))
}

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/go-openapi/spec"
	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

var (
	schemaPath         = "/dev/null/path/to/schema.json"
	externalSchemaPath = "/dev/null/path/to/other.json"
	siblingSchemaPath  = "/dev/null/path/to/sibling.json"
)

func Test_CanExtractTypeNameFromSameFile(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	schemaPackage := test.MakeLocalPackageReference("Microsoft.Test", "v1")

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

	wrappedSchema := MakeOpenAPISchema("name", schema, schemaPath, schemaPackage, astmodel.NewIdentifierFactory(), loader)

	typeName, err := wrappedSchema.refTypeName()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(typeName).To(Equal(astmodel.MakeTypeName(schemaPackage, "TheDefinition")))
}

func Test_CanExtractTypeNameFromDifferentFile_AndInheritPackage(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	schemaPackage := test.MakeLocalPackageReference("Microsoft.Test", "v1")

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
			Ref: spec.MustCreateRef(filepath.Base(externalSchemaPath) + "#/definitions/ExternalDefinition"),
		},
	}

	wrappedSchema := MakeOpenAPISchema("name", schema, schemaPath, schemaPackage, astmodel.NewIdentifierFactory(), loader)

	typeName, err := wrappedSchema.refTypeName()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(typeName).To(Equal(astmodel.MakeTypeName(schemaPackage, "ExternalDefinition")))
}

func Test_CanExtractTypeNameFromDifferentFile_AndUsePresetPackage(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	schemaPackage := test.MakeLocalPackageReference("Microsoft.Test", "v1")

	externalPackage := test.MakeLocalPackageReference("Microsoft.Common", "v1")

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
			Ref: spec.MustCreateRef(filepath.Base(externalSchemaPath) + "#/definitions/ExternalDefinition"),
		},
	}

	wrappedSchema := MakeOpenAPISchema("name", schema, schemaPath, schemaPackage, astmodel.NewIdentifierFactory(), loader)

	typeName, err := wrappedSchema.refTypeName()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(typeName).To(Equal(astmodel.MakeTypeName(externalPackage, "ExternalDefinition")))
}

func Test_GeneratingCollidingTypeNamesReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	schemaPackage := test.MakeLocalPackageReference("Microsoft.Test", "v1")

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
			Ref: spec.MustCreateRef(filepath.Base(externalSchemaPath) + "#/definitions/TheDefinition"),
		},
	}

	wrappedSchema := MakeOpenAPISchema("name", schema, schemaPath, schemaPackage, astmodel.NewIdentifierFactory(), loader)

	_, err := wrappedSchema.refTypeName()
	g.Expect(err).To(HaveOccurred())

	// Check the error message contains the bits we want
	//(checking for parts allows the message to be modified without breaking the test)
	// We don't want to fuss with / vs \ so we normalise
	msg := filepath.ToSlash(err.Error())
	g.Expect(msg).To(ContainSubstring(filepath.ToSlash(schemaPath)))
	g.Expect(msg).To(ContainSubstring(filepath.ToSlash(externalSchemaPath)))
	g.Expect(msg).To(ContainSubstring(schemaPackage.String()))
	g.Expect(msg).To(ContainSubstring("could generate collision"))
}

func Test_GeneratingCollidingTypeNamesWithSiblingFilesReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	schemaPackage := test.MakeLocalPackageReference("Microsoft.Test", "v1")

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
			Ref: spec.MustCreateRef(filepath.Base(externalSchemaPath) + "#/definitions/ExternalDefinition"),
		},
	}

	wrappedSchema := MakeOpenAPISchema("name", schema, schemaPath, schemaPackage, astmodel.NewIdentifierFactory(), loader)

	_, err := wrappedSchema.refTypeName()
	g.Expect(err).To(HaveOccurred())

	// Check the error message contains the bits we want
	//(checking for parts allows the message to be modified without breaking the test)
	// We don't want to fuss with / vs \ so we normalise
	msg := filepath.ToSlash(err.Error())
	g.Expect(msg).To(ContainSubstring(filepath.ToSlash(externalSchemaPath)))
	g.Expect(msg).To(ContainSubstring(filepath.ToSlash(siblingSchemaPath)))
	g.Expect(msg).To(ContainSubstring(schemaPackage.String()))
	g.Expect(msg).To(ContainSubstring("could generate collision"))
}

func init() {
	if runtime.GOOS == "windows" {
		schemaPath = "C:" + filepath.FromSlash(schemaPath)
		externalSchemaPath = "C:" + filepath.FromSlash(externalSchemaPath)
		siblingSchemaPath = "C:" + filepath.FromSlash(siblingSchemaPath)
	}
}

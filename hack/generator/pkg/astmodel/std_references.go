/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

var (
	// References to standard Go Libraries
	FmtReference     = MakeExternalPackageReference("fmt")
	JsonReference    = MakeExternalPackageReference("encoding/json")
	ReflectReference = MakeExternalPackageReference("reflect")
	TestingReference = MakeExternalPackageReference("testing")

	// References to our Libraries
	GenRuntimeReference PackageReference = MakeExternalPackageReference(genRuntimePathPrefix)

	// References to other libraries
	ErrorsReference            = MakeExternalPackageReference("github.com/pkg/errors")
	ApiExtensionsReference     = MakeExternalPackageReference("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1")
	ApiExtensionsJsonReference = MakeExternalPackageReference("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1/JSON")

	// References to libraries used for testing
	CmpReference        = MakeExternalPackageReference("github.com/google/go-cmp/cmp")
	CmpOptsReference    = MakeExternalPackageReference("github.com/google/go-cmp/cmp/cmpopts")
	DiffReference       = MakeExternalPackageReference("github.com/kylelemons/godebug/diff")
	GopterReference     = MakeExternalPackageReference("github.com/leanovate/gopter")
	GopterGenReference  = MakeExternalPackageReference("github.com/leanovate/gopter/gen")
	GopterPropReference = MakeExternalPackageReference("github.com/leanovate/gopter/prop")
	GomegaReference     = MakeExternalPackageReference("github.com/onsi/gomega")
	PrettyReference     = MakeExternalPackageReference("github.com/kr/pretty")

	// Imports with specified names
	GomegaImport = NewPackageImport(GomegaReference).WithName(".")

	// Type names
	ResourceReferenceTypeName = MakeTypeName(GenRuntimeReference, "ResourceReference")
)

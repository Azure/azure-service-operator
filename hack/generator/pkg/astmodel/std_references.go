/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

var (
	// References to standard Go Libraries
	FmtReference     PackageReference = MakeExternalPackageReference("fmt")
	JsonReference    PackageReference = MakeExternalPackageReference("encoding/json")
	ReflectReference PackageReference = MakeExternalPackageReference("reflect")
	TestingReference PackageReference = MakeExternalPackageReference("testing")

	// References to our Libraries
	GenRuntimeReference PackageReference = MakeExternalPackageReference(genRuntimePathPrefix)

	// References to other libraries
	ApiExtensionsReference     = MakeExternalPackageReference("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1")
	ApiExtensionsJsonReference = MakeExternalPackageReference("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1/JSON")

	// References to libraries used for testing
	CmpReference        PackageReference = MakeExternalPackageReference("github.com/google/go-cmp/cmp")
	CmpOptsReference    PackageReference = MakeExternalPackageReference("github.com/google/go-cmp/cmp/cmpopts")
	DiffReference       PackageReference = MakeExternalPackageReference("github.com/kylelemons/godebug/diff")
	GopterReference     PackageReference = MakeExternalPackageReference("github.com/leanovate/gopter")
	GopterGenReference  PackageReference = MakeExternalPackageReference("github.com/leanovate/gopter/gen")
	GopterPropReference PackageReference = MakeExternalPackageReference("github.com/leanovate/gopter/prop")
	GomegaReference     PackageReference = MakeExternalPackageReference("github.com/onsi/gomega")
	PrettyReference     PackageReference = MakeExternalPackageReference("github.com/kr/pretty")

	// Imports with specified names
	GomegaImport PackageImport = NewPackageImport(GomegaReference).WithName(".")
)

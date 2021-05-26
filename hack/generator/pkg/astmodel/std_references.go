/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

var (
	// References to standard Go Libraries
	ErrorsReference  PackageReference = MakeExternalPackageReference("errors")
	FmtReference     PackageReference = MakeExternalPackageReference("fmt")
	JsonReference    PackageReference = MakeExternalPackageReference("encoding/json")
	ReflectReference PackageReference = MakeExternalPackageReference("reflect")
	TestingReference PackageReference = MakeExternalPackageReference("testing")

	// References to our Libraries
	GenRuntimeReference PackageReference = MakeExternalPackageReference(genRuntimePathPrefix)

	// References to other libraries
	APIExtensionsReference       = MakeExternalPackageReference("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1")
	APIExtensionsJSONReference   = MakeExternalPackageReference("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1/JSON")
	APIMachineryErrorsReference  = MakeExternalPackageReference("k8s.io/apimachinery/pkg/util/errors")
	APIMachineryRuntimeReference = MakeExternalPackageReference("k8s.io/apimachinery/pkg/runtime")
	ClientGoSchemeReference      = MakeExternalPackageReference("k8s.io/client-go/kubernetes/scheme")
	ControllerRuntimeAdmission   = MakeExternalPackageReference("sigs.k8s.io/controller-runtime/pkg/webhook/admission")

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

	// Type names
	ResourceReferenceTypeName = MakeTypeName(GenRuntimeReference, "ResourceReference")
)

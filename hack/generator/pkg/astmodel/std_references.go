/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

const (
	reflectHelpersPath = "github.com/Azure/azure-service-operator/hack/generated/pkg/reflecthelpers"
)

var (
	// References to standard Go Libraries
	ErrorsReference  = MakeExternalPackageReference("errors")
	FmtReference     = MakeExternalPackageReference("fmt")
	JsonReference    = MakeExternalPackageReference("encoding/json")
	ReflectReference = MakeExternalPackageReference("reflect")
	TestingReference = MakeExternalPackageReference("testing")

	// References to our Libraries
	GenRuntimeReference     = MakeExternalPackageReference(genRuntimePathPrefix)
	ReflectHelpersReference = MakeExternalPackageReference(reflectHelpersPath)

	// References to other libraries
	APIExtensionsReference       = MakeExternalPackageReference("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1")
	APIExtensionsJSONReference   = MakeExternalPackageReference("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1/JSON")
	APIMachineryErrorsReference  = MakeExternalPackageReference("k8s.io/apimachinery/pkg/util/errors")
	APIMachineryRuntimeReference = MakeExternalPackageReference("k8s.io/apimachinery/pkg/runtime")
	APIMachinerySchemaReference  = MakeExternalPackageReference("k8s.io/apimachinery/pkg/runtime/schema")

	ClientGoSchemeReference     = MakeExternalPackageReference("k8s.io/client-go/kubernetes/scheme")
	ControllerRuntimeAdmission  = MakeExternalPackageReference("sigs.k8s.io/controller-runtime/pkg/webhook/admission")
	ControllerRuntimeConversion = MakeExternalPackageReference("sigs.k8s.io/controller-runtime/pkg/conversion")
	ControllerSchemeReference   = MakeExternalPackageReference("sigs.k8s.io/controller-runtime/pkg/scheme")
	GitHubErrorsReference       = MakeExternalPackageReference("github.com/pkg/errors")

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

	// Type names - GenRuntime
	ResourceReferenceTypeName      = MakeTypeName(GenRuntimeReference, "ResourceReference")
	KnownResourceReferenceTypeName = MakeTypeName(GenRuntimeReference, "KnownResourceReference")
	ToARMConverterInterfaceType    = MakeTypeName(GenRuntimeReference, "ToARMConverter")

	// Type names - API Machinery
	GroupVersionKindTypeName = MakeTypeName(APIMachinerySchemaReference, "GroupVersionKind")
	SchemeType               = MakeTypeName(APIMachineryRuntimeReference, "Scheme")
	JSONTypeName             = MakeTypeName(APIExtensionsReference, "JSON")

	// Type names - Controller Runtime
	ConvertibleInterface = MakeTypeName(ControllerRuntimeConversion, "Convertible")
	HubInterface         = MakeTypeName(ControllerRuntimeConversion, "Hub")
)

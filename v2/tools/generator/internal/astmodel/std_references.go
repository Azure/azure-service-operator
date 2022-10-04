/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

const (
	reflectHelpersPath   = "github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	genericARMClientPath = "github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
)

var (
	// References to standard Go Libraries
	ErrorsReference  = MakeExternalPackageReference("errors")
	FmtReference     = MakeExternalPackageReference("fmt")
	JsonReference    = MakeExternalPackageReference("encoding/json")
	OSReference      = MakeExternalPackageReference("os")
	ReflectReference = MakeExternalPackageReference("reflect")
	TestingReference = MakeExternalPackageReference("testing")
	ContextReference = MakeExternalPackageReference("context")

	// References to our Libraries
	GenRuntimeReference             = MakeExternalPackageReference(genRuntimePathPrefix)
	GenRuntimeConditionsReference   = MakeExternalPackageReference(genRuntimePathPrefix + "/conditions")
	GenRuntimeRegistrationReference = MakeExternalPackageReference(genRuntimePathPrefix + "/registration")
	ReflectHelpersReference         = MakeExternalPackageReference(reflectHelpersPath)
	GenRuntimeConfigMapsReference   = MakeExternalPackageReference(genRuntimePathPrefix + "/configmaps")
	GenericARMClientReference       = MakeExternalPackageReference(genericARMClientPath)

	// References to other libraries
	APIExtensionsReference       = MakeExternalPackageReference("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1")
	APIExtensionsJSONReference   = MakeExternalPackageReference("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1/JSON")
	APIMachineryErrorsReference  = MakeExternalPackageReference("k8s.io/apimachinery/pkg/util/errors")
	APIMachineryRuntimeReference = MakeExternalPackageReference("k8s.io/apimachinery/pkg/runtime")
	APIMachinerySchemaReference  = MakeExternalPackageReference("k8s.io/apimachinery/pkg/runtime/schema")
	MetaV1Reference              = MakeExternalPackageReference("k8s.io/apimachinery/pkg/apis/meta/v1")
	CoreV1Reference              = MakeExternalPackageReference("k8s.io/api/core/v1")
	LogrReference                = MakeExternalPackageReference("github.com/go-logr/logr")

	ClientGoSchemeReference     = MakeExternalPackageReference("k8s.io/client-go/kubernetes/scheme")
	ControllerRuntimeAdmission  = MakeExternalPackageReference("sigs.k8s.io/controller-runtime/pkg/webhook/admission")
	ControllerRuntimeConversion = MakeExternalPackageReference("sigs.k8s.io/controller-runtime/pkg/conversion")
	ControllerSchemeReference   = MakeExternalPackageReference("sigs.k8s.io/controller-runtime/pkg/scheme")
	ControllerRuntimeClient     = MakeExternalPackageReference("sigs.k8s.io/controller-runtime/pkg/client")
	ControllerRuntimeSource     = MakeExternalPackageReference("sigs.k8s.io/controller-runtime/pkg/source")
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
	KubernetesResourceType           = MakeTypeName(GenRuntimeReference, "KubernetesResource")
	KubernetesExporterType           = MakeTypeName(GenRuntimeReference, "KubernetesExporter")
	TenantResourceType               = MakeTypeName(GenRuntimeReference, "TenantResource")
	ConvertibleSpecInterfaceType     = MakeTypeName(GenRuntimeReference, "ConvertibleSpec")
	ConvertibleStatusInterfaceType   = MakeTypeName(GenRuntimeReference, "ConvertibleStatus")
	ResourceReferenceType            = MakeTypeName(GenRuntimeReference, "ResourceReference")
	ArbitraryOwnerReference          = MakeTypeName(GenRuntimeReference, "ArbitraryOwnerReference")
	KnownResourceReferenceType       = MakeTypeName(GenRuntimeReference, "KnownResourceReference")
	PropertyBagType                  = MakeTypeName(GenRuntimeReference, "PropertyBag")
	ToARMConverterInterfaceType      = MakeTypeName(GenRuntimeReference, "ToARMConverter")
	ARMResourceSpecType              = MakeTypeName(GenRuntimeReference, "ARMResourceSpec")
	ARMResourceStatusType            = MakeTypeName(GenRuntimeReference, "ARMResourceStatus")
	ResourceScopeType                = MakeTypeName(GenRuntimeReference, "ResourceScope")
	ConvertToARMResolvedDetailsType  = MakeTypeName(GenRuntimeReference, "ConvertToARMResolvedDetails")
	SecretReferenceType              = MakeTypeName(GenRuntimeReference, "SecretReference")
	ResourceExtensionType            = MakeTypeName(GenRuntimeReference, "ResourceExtension")
	SecretDestinationType            = MakeTypeName(GenRuntimeReference, "SecretDestination")
	ConfigMapDestinationType         = MakeTypeName(GenRuntimeReference, "ConfigMapDestination")
	ConfigMapReferenceType           = MakeTypeName(GenRuntimeReference, "ConfigMapReference")
	GenRuntimeDefaulterInterfaceName = MakeTypeName(GenRuntimeReference, "Defaulter")
	GenRuntimeValidatorInterfaceName = MakeTypeName(GenRuntimeReference, "Validator")
	GenRuntimeMetaObjectType         = MakeTypeName(GenRuntimeReference, "MetaObject")

	// Type names - Generic ARM client
	GenericClientType = MakeTypeName(GenericARMClientReference, "GenericClient")

	// Type names - Registration
	StorageTypeRegistrationType = MakeTypeName(GenRuntimeRegistrationReference, "StorageType")
	IndexRegistrationType       = MakeTypeName(GenRuntimeRegistrationReference, "Index")
	WatchRegistrationType       = MakeTypeName(GenRuntimeRegistrationReference, "Watch")

	ConditionType   = MakeTypeName(GenRuntimeConditionsReference, "Condition")
	ConditionsType  = MakeTypeName(GenRuntimeConditionsReference, "Conditions")
	ConditionerType = MakeTypeName(GenRuntimeConditionsReference, "Conditioner")

	// Type names - API Machinery
	GroupVersionKindType = MakeTypeName(APIMachinerySchemaReference, "GroupVersionKind")
	SchemeType           = MakeTypeName(APIMachineryRuntimeReference, "Scheme")
	JSONType             = MakeTypeName(APIExtensionsReference, "JSON")
	ObjectMetaType       = MakeTypeName(MetaV1Reference, "ObjectMeta")

	// Type names - Controller Runtime
	ConvertibleInterface            = MakeTypeName(ControllerRuntimeConversion, "Convertible")
	HubInterface                    = MakeTypeName(ControllerRuntimeConversion, "Hub")
	ControllerRuntimeObjectType     = MakeTypeName(ControllerRuntimeClient, "Object")
	ControllerRuntimeSourceKindType = MakeTypeName(ControllerRuntimeSource, "Kind")
	DefaulterInterfaceName          = MakeTypeName(ControllerRuntimeAdmission, "Defaulter")
	ValidatorInterfaceName          = MakeTypeName(ControllerRuntimeAdmission, "Validator")

	// Type names - Core types
	SecretType    = MakeTypeName(CoreV1Reference, "Secret")
	ConfigMapType = MakeTypeName(CoreV1Reference, "ConfigMap")

	// Type names - stdlib types
	ContextType = MakeTypeName(ContextReference, "Context")

	// Type names - Logr types
	LogrType = MakeTypeName(LogrReference, "Logger")
)

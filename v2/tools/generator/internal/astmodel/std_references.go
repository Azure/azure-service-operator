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
	KubernetesResourceType           = MakeExternalTypeName(GenRuntimeReference, "KubernetesResource")
	KubernetesExporterType           = MakeExternalTypeName(GenRuntimeReference, "KubernetesExporter")
	TenantResourceType               = MakeExternalTypeName(GenRuntimeReference, "TenantResource")
	ConvertibleSpecInterfaceType     = MakeExternalTypeName(GenRuntimeReference, "ConvertibleSpec")
	ConvertibleStatusInterfaceType   = MakeExternalTypeName(GenRuntimeReference, "ConvertibleStatus")
	ResourceReferenceType            = MakeExternalTypeName(GenRuntimeReference, "ResourceReference")
	ArbitraryOwnerReference          = MakeExternalTypeName(GenRuntimeReference, "ArbitraryOwnerReference")
	KnownResourceReferenceType       = MakeExternalTypeName(GenRuntimeReference, "KnownResourceReference")
	PropertyBagType                  = MakeExternalTypeName(GenRuntimeReference, "PropertyBag")
	ToARMConverterInterfaceType      = MakeExternalTypeName(GenRuntimeReference, "ToARMConverter")
	ARMResourceSpecType              = MakeExternalTypeName(GenRuntimeReference, "ARMResourceSpec")
	ARMResourceStatusType            = MakeExternalTypeName(GenRuntimeReference, "ARMResourceStatus")
	ResourceScopeType                = MakeExternalTypeName(GenRuntimeReference, "ResourceScope")
	ConvertToARMResolvedDetailsType  = MakeExternalTypeName(GenRuntimeReference, "ConvertToARMResolvedDetails")
	SecretReferenceType              = MakeExternalTypeName(GenRuntimeReference, "SecretReference")
	ResourceExtensionType            = MakeExternalTypeName(GenRuntimeReference, "ResourceExtension")
	SecretDestinationType            = MakeExternalTypeName(GenRuntimeReference, "SecretDestination")
	ConfigMapDestinationType         = MakeExternalTypeName(GenRuntimeReference, "ConfigMapDestination")
	ConfigMapReferenceType           = MakeExternalTypeName(GenRuntimeReference, "ConfigMapReference")
	GenRuntimeDefaulterInterfaceName = MakeExternalTypeName(GenRuntimeReference, "Defaulter")
	GenRuntimeValidatorInterfaceName = MakeExternalTypeName(GenRuntimeReference, "Validator")
	GenRuntimeMetaObjectType         = MakeExternalTypeName(GenRuntimeReference, "MetaObject")
	LocatableResourceInterfaceName   = MakeExternalTypeName(GenRuntimeReference, "LocatableResource")
	ImportableResourceType           = MakeExternalTypeName(GenRuntimeReference, "ImportableResource")

	// Optional types - GenRuntime
	OptionalConfigMapReferenceType     = NewOptionalType(ConfigMapReferenceType)
	OptionalKnownResourceReferenceType = NewOptionalType(KnownResourceReferenceType)
	OptionalResourceReferenceType      = NewOptionalType(ResourceReferenceType)
	OptionalSecretReferenceType        = NewOptionalType(SecretReferenceType)

	// Type names - Generic ARM client
	GenericClientType = MakeExternalTypeName(GenericARMClientReference, "GenericClient")

	// Type names - Registration
	StorageTypeRegistrationType = MakeExternalTypeName(GenRuntimeRegistrationReference, "StorageType")
	IndexRegistrationType       = MakeExternalTypeName(GenRuntimeRegistrationReference, "Index")
	WatchRegistrationType       = MakeExternalTypeName(GenRuntimeRegistrationReference, "Watch")

	ConditionType   = MakeExternalTypeName(GenRuntimeConditionsReference, "Condition")
	ConditionsType  = MakeExternalTypeName(GenRuntimeConditionsReference, "Conditions")
	ConditionerType = MakeExternalTypeName(GenRuntimeConditionsReference, "Conditioner")

	// Type names - API Machinery
	GroupVersionKindType = MakeExternalTypeName(APIMachinerySchemaReference, "GroupVersionKind")
	SchemeType           = MakeExternalTypeName(APIMachineryRuntimeReference, "Scheme")
	JSONType             = MakeExternalTypeName(APIExtensionsReference, "JSON")
	ObjectMetaType       = MakeExternalTypeName(MetaV1Reference, "ObjectMeta")

	// Type names - Controller Runtime
	ConvertibleInterface        = MakeExternalTypeName(ControllerRuntimeConversion, "Convertible")
	HubInterface                = MakeExternalTypeName(ControllerRuntimeConversion, "Hub")
	ControllerRuntimeObjectType = MakeExternalTypeName(ControllerRuntimeClient, "Object")
	DefaulterInterfaceName      = MakeExternalTypeName(ControllerRuntimeAdmission, "Defaulter")
	ValidatorInterfaceName      = MakeExternalTypeName(ControllerRuntimeAdmission, "Validator")

	// Type names - Core types
	SecretType    = MakeExternalTypeName(CoreV1Reference, "Secret")
	ConfigMapType = MakeExternalTypeName(CoreV1Reference, "ConfigMap")

	// Type names - stdlib types
	ContextType = MakeExternalTypeName(ContextReference, "Context")

	// Type names - Logr types
	LogrType = MakeExternalTypeName(LogrReference, "Logger")
)

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/mod/modfile"
	"gopkg.in/yaml.v3"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

type GenerationPipeline string

const (
	GenerationPipelineAzure      = GenerationPipeline("azure")
	GenerationPipelineCrossplane = GenerationPipeline("crossplane")
)

// Configuration is used to control which types get generated
type Configuration struct {
	// Where to load Swagger schemas from
	SchemaRoot string `yaml:"schemaRoot"`
	// Information about where to locate status (Swagger) files
	Status StatusConfiguration `yaml:"status"`
	// The pipeline that should be used for code generation
	Pipeline GenerationPipeline `yaml:"pipeline"`
	// The path to the go.mod file where the code will be generated
	DestinationGoModuleFile string `yaml:"destinationGoModuleFile"`
	// The folder relative to the go.mod file path where the code should be generated
	TypesOutputPath string `yaml:"typesOutputPath"`
	// The file relative to the go.mod file path where registration of the Go types should be generated. If omitted, this step is skipped.
	TypeRegistrationOutputFile string `yaml:"typeRegistrationOutputFile"`
	// AnyTypePackages lists packages which we expect to generate
	// interface{} fields.
	AnyTypePackages []string `yaml:"anyTypePackages"`
	// Filters used to control which types are created from the JSON schema
	TypeFilters []*TypeFilter `yaml:"typeFilters"`
	// Transformers used to remap types
	Transformers []*TypeTransformer `yaml:"typeTransformers"`
	// RootURL is the root URL for ASOv2 repo, paths are appended to this to generate resource links.
	RootURL string `yaml:"rootUrl"`
	// SamplesPath is the Path the samples are accessible at. This is used to walk through the samples directory and generate sample links.
	SamplesPath string `yaml:"samplesPath"`
	// EmitDocFiles is used as a signal to create doc.go files for packages. If omitted, default is false.
	EmitDocFiles bool `yaml:"emitDocFiles"`
	// Destination file and additional information for our supported resources report
	SupportedResourcesReport *SupportedResourcesReport `yaml:"supportedResourcesReport"`
	// Additional information about our object model
	ObjectModelConfiguration *ObjectModelConfiguration `yaml:"objectModelConfiguration"`

	goModulePath string

	// after init TypeTransformers is split into property and non-property transformers
	typeTransformers     []*TypeTransformer
	propertyTransformers []*TypeTransformer
}

type RewriteRule struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
}

func (config *Configuration) LocalPathPrefix() string {
	return path.Join(config.goModulePath, config.TypesOutputPath)
}

func (config *Configuration) FullTypesOutputPath() string {
	return filepath.Join(
		filepath.Dir(config.DestinationGoModuleFile),
		config.TypesOutputPath)
}

func (config *Configuration) FullTypesRegistrationOutputFilePath() string {
	if config.TypeRegistrationOutputFile == "" {
		return ""
	}

	return filepath.Join(
		filepath.Dir(config.DestinationGoModuleFile),
		config.TypeRegistrationOutputFile)
}

func (config *Configuration) FullSamplesPath() string {
	if filepath.IsAbs(config.SamplesPath) {
		return config.SamplesPath
	}

	if config.DestinationGoModuleFile != "" {
		return filepath.Join(
			filepath.Dir(config.DestinationGoModuleFile),
			config.SamplesPath)
	}

	result, err := filepath.Abs(config.SamplesPath)
	if err != nil {
		panic(fmt.Sprintf("unable to make %q absolute: %v", result, err))
	}

	return result
}

func (config *Configuration) GetTypeFiltersError() error {
	for _, filter := range config.TypeFilters {
		if err := filter.RequiredTypesWereMatched(); err != nil {
			return errors.Wrapf(err, "type filter action: %q", filter.Action)
		}
	}

	return nil
}

func (config *Configuration) GetTypeTransformersError() error {
	for _, filter := range config.typeTransformers {
		if err := filter.RequiredTypesWereMatched(); err != nil {
			return errors.Wrap(err, "type transformer")
		}
	}

	return nil
}

func (config *Configuration) GetPropertyTransformersError() error {
	var errs []error
	for _, filter := range config.propertyTransformers {
		if err := filter.RequiredTypesWereMatched(); err != nil {
			errs = append(errs, err)
			continue
		}

		if err := filter.RequiredPropertiesWereMatched(); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Wrap(
		kerrors.NewAggregate(errs),
		"type transformer target")
}

func (config *Configuration) SetGoModulePath(path string) {
	config.goModulePath = path
}

// NewConfiguration returns a new empty Configuration
func NewConfiguration() *Configuration {
	result := &Configuration{
		ObjectModelConfiguration: NewObjectModelConfiguration(),
	}

	result.SupportedResourcesReport = NewSupportedResourcesReport(result)

	return result
}

// LoadConfiguration loads a `Configuration` from the specified file
func LoadConfiguration(configurationFile string) (*Configuration, error) {
	f, err := os.Open(configurationFile)
	if err != nil {
		return nil, err
	}

	// MUST use this to ensure that ObjectModelConfiguration is instantiated correctly
	// TODO: split Configuration struct so that domain model is not used for serialization!
	result := NewConfiguration()

	decoder := yaml.NewDecoder(f)
	decoder.KnownFields(true) // Error on unknown fields

	err = decoder.Decode(result)
	if err != nil {
		return nil, errors.Wrapf(err, "configuration file loaded from %q is not valid YAML", configurationFile)
	}

	err = result.initialize(configurationFile)
	if err != nil {
		return nil, errors.Wrapf(err, "configuration file loaded from %q is invalid", configurationFile)
	}

	return result, nil
}

// ShouldPruneResult is returned by ShouldPrune to indicate whether the supplied type should be exported
type ShouldPruneResult string

const (
	// Include indicates the specified type should be included in the type graph
	Include ShouldPruneResult = "include"
	// Prune indicates the type (and all types only referenced by it) should be pruned from the type graph
	Prune ShouldPruneResult = "prune"
)

// TypeRename looks up a type-rename for the specified type, returning the new name and true if found, or empty string
// and false if not.
func (config *Configuration) TypeRename(name astmodel.TypeName) (string, error) {
	if config.ObjectModelConfiguration == nil {
		return "", errors.Errorf("no configuration: no rename available for %s", name)
	}

	return config.ObjectModelConfiguration.LookupNameInNextVersion(name)
}

// ARMReference looks up a property to determine whether it may be an ARM reference or not.
func (config *Configuration) ARMReference(name astmodel.TypeName, property astmodel.PropertyName) (bool, error) {
	return config.ObjectModelConfiguration.ARMReference(name, property)
}

// VerifyARMReferencesConsumed returns an error if any configured ARM References were not consumed
func (config *Configuration) VerifyARMReferencesConsumed() error {
	return config.ObjectModelConfiguration.VerifyARMReferencesConsumed()
}

// IsSecret looks up a property to determine whether it is a secret.
func (config *Configuration) IsSecret(name astmodel.TypeName, property astmodel.PropertyName) (bool, error) {
	return config.ObjectModelConfiguration.IsSecret(name, property)
}

// VerifyIsSecretConsumed returns an error if any configured Secret References were not consumed
func (config *Configuration) VerifyIsSecretConsumed() error {
	return config.ObjectModelConfiguration.VerifyIsSecretConsumed()
}

// ResourceLifecycleOwnedByParent looks up a property to determine if represents a subresource whose lifecycle is owned
// by the parent resource.
func (config *Configuration) ResourceLifecycleOwnedByParent(name astmodel.TypeName, property astmodel.PropertyName) (string, error) {
	return config.ObjectModelConfiguration.ResourceLifecycleOwnedByParent(name, property)
}

// MarkResourceLifecycleOwnedByParentUnconsumed marks all ResourceLifecycleOwnedByParent as unconsumed
func (config *Configuration) MarkResourceLifecycleOwnedByParentUnconsumed() error {
	return config.ObjectModelConfiguration.MarkResourceLifecycleOwnedByParentUnconsumed()
}

// VerifyResourceLifecycleOwnedByParentConsumed returns an error if any ResourceLifecycleOwnedByParent flag is not consumed
func (config *Configuration) VerifyResourceLifecycleOwnedByParentConsumed() error {
	return config.ObjectModelConfiguration.VerifyResourceLifecycleOwnedByParentConsumed()
}

// ImportConfigMapMode looks up a property to determine its import configMap mode.
func (config *Configuration) ImportConfigMapMode(name astmodel.TypeName, property astmodel.PropertyName) (ImportConfigMapMode, error) {
	return config.ObjectModelConfiguration.ImportConfigMapMode(name, property)
}

// VerifyImportConfigMapModeConsumed returns an error if any configured ImportConfigMapMode values were not consumed
func (config *Configuration) VerifyImportConfigMapModeConsumed() error {
	return config.ObjectModelConfiguration.VerifyImportConfigMapModeConsumed()
}

// initialize checks for common errors and initializes structures inside the configuration
// which need additional setup after json deserialization
func (config *Configuration) initialize(configPath string) error {
	if config.SchemaRoot == "" {
		return errors.New("SchemaRoot missing")
	}

	absConfigLocation, err := filepath.Abs(configPath)
	if err != nil {
		return errors.Wrapf(err, "unable to find absolute config file location")
	}

	configDirectory := filepath.Dir(absConfigLocation)

	// resolve SchemaRoot relative to config file directory
	config.SchemaRoot = filepath.Join(configDirectory, config.SchemaRoot)

	if config.TypesOutputPath == "" {
		// Default to an apis folder if not specified
		config.TypesOutputPath = "apis"
	}

	// Trim any trailing slashes on the output path
	config.TypesOutputPath = strings.TrimSuffix(config.TypesOutputPath, "/")

	var errs []error

	if config.Pipeline == "" {
		// Default to the standard Azure pipeline
		config.Pipeline = GenerationPipelineAzure
	} else {
		switch pipeline := strings.ToLower(string(config.Pipeline)); pipeline {
		case string(GenerationPipelineAzure):
			config.Pipeline = GenerationPipelineAzure
		case string(GenerationPipelineCrossplane):
			config.Pipeline = GenerationPipelineCrossplane
		default:
			errs = append(errs, errors.Errorf("unknown pipeline kind %s", config.Pipeline))
		}
	}

	if config.DestinationGoModuleFile == "" {
		errs = append(errs, errors.Errorf("destination Go module must be specified"))
	}

	// Ensure config.DestinationGoModuleFile is a fully qualified path
	if !filepath.IsAbs(config.DestinationGoModuleFile) {
		config.DestinationGoModuleFile = filepath.Join(configDirectory, config.DestinationGoModuleFile)
	}

	modPath, err := getModulePathFromModFile(config.DestinationGoModuleFile)
	if err != nil {
		errs = append(errs, err)
	} else {
		config.goModulePath = modPath
	}

	for _, filter := range config.TypeFilters {
		err := filter.Initialize()
		if err != nil {
			errs = append(errs, err)
		}
	}

	// split Transformers into two sets
	var typeTransformers []*TypeTransformer
	var propertyTransformers []*TypeTransformer
	for _, transformer := range config.Transformers {
		err := transformer.Initialize(config.MakeLocalPackageReference)
		if err != nil {
			errs = append(errs, err)
		}

		if transformer.Property.IsRestrictive() {
			propertyTransformers = append(propertyTransformers, transformer)
		} else {
			typeTransformers = append(typeTransformers, transformer)
		}
	}

	config.Transformers = nil
	config.typeTransformers = typeTransformers
	config.propertyTransformers = propertyTransformers

	return kerrors.NewAggregate(errs)
}

func absDirectoryPathToURL(path string) *url.URL {
	if strings.Contains(path, "\\") {
		// assume it's a Windows path:
		// fixup  to work in URI
		path = "/" + strings.ReplaceAll(path, "\\", "/")
	}

	result, err := url.Parse("file://" + path + "/")
	if err != nil {
		panic(err)
	}

	return result
}

// ShouldPrune tests for whether a given type should be extracted from the JSON schema or pruned
func (config *Configuration) ShouldPrune(typeName astmodel.TypeName) (result ShouldPruneResult, because string) {
	for _, f := range config.TypeFilters {
		if f.AppliesToType(typeName) {
			switch f.Action {
			case TypeFilterPrune:
				return Prune, f.Because
			case TypeFilterInclude:
				return Include, f.Because
			default:
				panic(errors.Errorf("unknown typefilter directive: %s", f.Action))
			}
		}
	}

	// If the type comes from a group that we don't expect, prune it.
	// We don't also check for whether the version is expected because it's common for types to be shared
	// between versions of an API. While end up pulling them into the package alongside the resource, at this
	// point we haven't done that yet, so it's premature to filter by version.
	// Sometimes in testing, configuration will be empty, and we don't want to do any filtering when that's the case
	if !config.ObjectModelConfiguration.IsEmpty() &&
		!config.ObjectModelConfiguration.IsGroupConfigured(typeName.PackageReference) {
		return Prune, fmt.Sprintf("No resources configured for export from %s", typeName.PackageReference.PackagePath())
	}

	// By default, we include all types
	return Include, ""
}

// TransformType uses the configured type transformers to transform a type name (reference) to a different type.
// If no transformation is performed, nil is returned
func (config *Configuration) TransformType(name astmodel.TypeName) (astmodel.Type, string) {
	for _, transformer := range config.typeTransformers {
		result := transformer.TransformTypeName(name)
		if result != nil {
			return result, transformer.Because
		}
	}

	// No matches, return nil
	return nil, ""
}

// TransformTypeProperties applies any property transformers to the type
func (config *Configuration) TransformTypeProperties(name astmodel.TypeName, objectType *astmodel.ObjectType) []*PropertyTransformResult {
	var results []*PropertyTransformResult
	toTransform := objectType

	for _, transformer := range config.propertyTransformers {
		result := transformer.TransformProperty(name, toTransform)
		if result != nil {
			toTransform = result.NewType
			results = append(results, result)
		}
	}

	return results
}

// MakeLocalPackageReference creates a local package reference based on the configured destination location
func (config *Configuration) MakeLocalPackageReference(group string, version string) astmodel.LocalPackageReference {
	return astmodel.MakeLocalPackageReference(config.LocalPathPrefix(), group, astmodel.GeneratorVersion, version)
}

func getModulePathFromModFile(modFilePath string) (string, error) {
	// Calculate the actual destination
	modFileData, err := os.ReadFile(modFilePath)
	if err != nil {
		return "", err
	}

	modPath := modfile.ModulePath(modFileData)
	if modPath == "" {
		return "", errors.Errorf("couldn't find module path in mod file %s", modFilePath)
	}

	return modPath, nil
}

// StatusConfiguration provides configuration options for the
// status parts of resources, which are generated from the
// Azure Swagger specs.
type StatusConfiguration struct {
	// Custom per-group configuration
	Overrides []SchemaOverride `yaml:"overrides"`
}

// SchemaOverride provides configuration to override namespaces (groups)
// this is used (for example) to distinguish Microsoft.Network.Frontdoor
// from Microsoft.Network, even though both use Microsoft.Network in
// their Swagger specs.
type SchemaOverride struct {
	// The root for this group (relative to SchemaRoot)
	BasePath string `yaml:"basePath"`

	// A specific namespace (group name, in our domain language)
	Namespace string `yaml:"namespace"`

	// A suffix to add on to the group name
	Suffix string `yaml:"suffix"`

	// We don't use this now
	ResourceConfig []ResourceConfig `yaml:"resourceConfig"`

	// We don't use this now
	PostProcessor string `yaml:"postProcessor"`
}

type ResourceConfig struct {
	Type string `yaml:"type"`

	// TODO: Not sure that this datatype should be string, but we don't use it right now so keeping it as
	// TODO: string for simplicity
	Scopes string `yaml:"scopes"`
}

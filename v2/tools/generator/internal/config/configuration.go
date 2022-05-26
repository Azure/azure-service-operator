/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
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
	// Base URL for the JSON schema to generate
	SchemaURL string `yaml:"schemaUrl"`
	// Part of the schema URL to rewrite, allows repointing to local files
	SchemaURLRewrite *RewriteRule `yaml:"schemaUrlRewrite"`
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
	// SamplesURL is the URL the samples are accessible at. Paths will be appended to the end of this to
	// build full sample links. If this is not specified, no samples links are generated.
	SamplesURL string `yaml:"samplesUrl"`
	// EmitDocFiles is used as a signal to create doc.go files for packages. If omitted, default is false.
	EmitDocFiles bool `yaml:"emitDocFiles"`

	// Additional information about our object model
	ObjectModelConfiguration *ObjectModelConfiguration `yaml:"objectModelConfiguration"`

	GoModulePath string // TODO: Since this isn't yaml annotated it can't be set, right?

	// after init TypeTransformers is split into property and non-property transformers
	typeTransformers     []*TypeTransformer
	propertyTransformers []*TypeTransformer
}

type RewriteRule struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
}

func (config *Configuration) LocalPathPrefix() string {
	return path.Join(config.GoModulePath, config.TypesOutputPath)
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
	for _, filter := range config.propertyTransformers {
		if err := filter.RequiredTypesWereMatched(); err != nil {
			return errors.Wrap(err, "type transformer target")
		}

		if err := filter.RequiredPropertiesWereMatched(); err != nil {
			return errors.Wrapf(err, "type transformer target")
		}
	}

	return nil
}

// NewConfiguration returns a new empty Configuration
func NewConfiguration() *Configuration {
	return &Configuration{
		ObjectModelConfiguration: NewObjectModelConfiguration(),
	}
}

// LoadConfiguration loads a `Configuration` from the specified file
func LoadConfiguration(configurationFile string) (*Configuration, error) {
	data, err := os.ReadFile(configurationFile)
	if err != nil {
		return nil, err
	}

	// MUST use this to ensure that ObjectModelConfiguration is instantiated correctly
	// TODO: split Configuration struct so that domain model is not used for serialization!
	result := NewConfiguration()

	err = yaml.Unmarshal(data, result)
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

// TypeRename looks up a rename for the specified type, returning the new name and true if found, or empty string
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

// IsResourceLifecycleOwnedByParent looks up a property to determine if represents a subresource whose lifecycle is owned
// by the parent resource.
func (config *Configuration) IsResourceLifecycleOwnedByParent(name astmodel.TypeName, property astmodel.PropertyName) (bool, error) {
	return config.ObjectModelConfiguration.IsResourceLifecycleOwnedByParent(name, property)
}

// MarkIsResourceLifecycleOwnedByParentUnconsumed marks all IsResourceLifecycleOwnedByParent as unconsumed
func (config *Configuration) MarkIsResourceLifecycleOwnedByParentUnconsumed() error {
	return config.ObjectModelConfiguration.MarkIsResourceLifecycleOwnedByParentUnconsumed()
}

// VerifyIsResourceLifecycleOwnedByParentConsumed returns an error if any IsResourceLifecycleOwnedByParent flag is not consumed
func (config *Configuration) VerifyIsResourceLifecycleOwnedByParentConsumed() error {
	return config.ObjectModelConfiguration.VerifyIsResourceLifecycleOwnedByParentConsumed()
}

// initialize checks for common errors and initializes structures inside the configuration
// which need additional setup after json deserialization
func (config *Configuration) initialize(configPath string) error {
	if config.SchemaURL == "" {
		return errors.New("SchemaURL missing")
	}

	absConfigLocation, err := filepath.Abs(configPath)
	if err != nil {
		return errors.Wrapf(err, "unable to find absolute config file location")
	}

	configDirectory := filepath.Dir(absConfigLocation)

	schemaURL, err := url.Parse(config.SchemaURL)
	if err != nil {
		return errors.Wrapf(err, "SchemaURL invalid")
	}

	configDirectoryURL := absDirectoryPathToURL(configDirectory)

	// resolve URLs relative to config directory (if needed)
	config.SchemaURL = configDirectoryURL.ResolveReference(schemaURL).String()

	if config.SchemaURLRewrite != nil {
		rewrite := config.SchemaURLRewrite

		var toURL *url.URL
		toURL, err = url.Parse(rewrite.To)
		if err != nil {
			return errors.Wrapf(err, "unable to parse rewriteSchemaUrl.to as URL")
		}

		rewrite.To = configDirectoryURL.ResolveReference(toURL).String()
	}

	// resolve Status.SchemaRoot relative to config file directory
	config.Status.SchemaRoot = filepath.Join(configDirectory, config.Status.SchemaRoot)

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

	modPath, err := getModulePathFromModFile(config.DestinationGoModuleFile)
	if err != nil {
		errs = append(errs, err)
	} else {
		config.GoModulePath = modPath
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

		if transformer.Property.String() != "" {
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
	// The root URL of the status (Swagger) files (relative to this file)
	SchemaRoot string `yaml:"schemaRoot"`

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
}

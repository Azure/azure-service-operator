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

	"github.com/rotisserie/eris"
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
	// Renamers used to resolve naming collisions during loading
	TypeLoaderRenames []*TypeLoaderRename `yaml:"typeLoaderRenames"`
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
		panic(fmt.Sprintf("unable to make %q absolute: %s", result, err))
	}

	return result
}

func (config *Configuration) GetTypeFiltersError() error {
	for _, filter := range config.TypeFilters {
		if err := filter.RequiredTypesWereMatched(); err != nil {
			return eris.Wrapf(err, "type filter action: %q", filter.Action)
		}
	}

	return nil
}

func (config *Configuration) GetTransformersError() error {
	for _, filter := range config.Transformers {
		if err := filter.RequiredTypesWereMatched(); err != nil {
			return eris.Wrap(err, "type transformer")
		}

		if filter.Property.IsRestrictive() {
			if err := filter.RequiredTypesWereMatched(); err != nil {
				return eris.Wrap(err, "type transformer property")
			}
		}
	}

	return nil
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
		return nil, eris.Wrapf(err, "configuration file loaded from %q is not valid YAML", configurationFile)
	}

	err = result.initialize(configurationFile)
	if err != nil {
		return nil, eris.Wrapf(err, "configuration file loaded from %q is invalid", configurationFile)
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

// initialize checks for common errors and initializes structures inside the configuration
// which need additional setup after json deserialization
func (config *Configuration) initialize(configPath string) error {
	if config.SchemaRoot == "" {
		return eris.New("SchemaRoot missing")
	}

	absConfigLocation, err := filepath.Abs(configPath)
	if err != nil {
		return eris.Wrapf(err, "unable to find absolute config file location")
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
			errs = append(errs, eris.Errorf("unknown pipeline kind %s", config.Pipeline))
		}
	}

	if config.DestinationGoModuleFile == "" {
		errs = append(errs, eris.Errorf("destination Go module must be specified"))
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
func (config *Configuration) ShouldPrune(typeName astmodel.InternalTypeName) (result ShouldPruneResult, because string) {
	for _, f := range config.TypeFilters {
		if f.AppliesToType(typeName) {
			switch f.Action {
			case TypeFilterPrune:
				return Prune, f.Because
			case TypeFilterInclude:
				return Include, f.Because
			default:
				panic(eris.Errorf("unknown typefilter directive: %s", f.Action))
			}
		}
	}

	// If the type comes from a group that we don't expect, prune it.
	// We don't also check for whether the version is expected because it's common for types to be shared
	// between versions of an API. While end up pulling them into the package alongside the resource, at this
	// point we haven't done that yet, so it's premature to filter by version.
	// Sometimes in testing, configuration will be empty, and we don't want to do any filtering when that's the case
	if !config.ObjectModelConfiguration.IsEmpty() &&
		!config.ObjectModelConfiguration.IsGroupConfigured(typeName.InternalPackageReference()) {
		return Prune, fmt.Sprintf(
			"No resources configured for export from %s", typeName.InternalPackageReference().PackagePath())
	}

	// By default, we include all types
	return Include, ""
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
		return "", eris.Errorf("couldn't find module path in mod file %s", modFilePath)
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

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

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
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
	// Filters used to control which types are exported
	ExportFilters []*ExportFilter `yaml:"exportFilters"`
	// Filters used to control which types are created from the JSON schema
	TypeFilters []*TypeFilter `yaml:"typeFilters"`
	// Transformers used to remap types
	Transformers []*TypeTransformer `yaml:"typeTransformers"`
	GoModulePath string             // TODO: Since this isn't yaml annotated it can't be set, right?

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
		if !filter.MatchedRequiredTypes() {
			return errors.Errorf("Type filter action: %q, target: %q matched no types", filter.Action, filter.String())
		}
	}

	return nil
}

func (config *Configuration) GetTypeTransformersError() error {
	for _, filter := range config.typeTransformers {
		if !filter.MatchedRequiredTypes() {
			return errors.Errorf("Type transformer target: %q matched no types", filter.String())
		}
	}

	return nil
}

func (config *Configuration) GetPropertyTransformersError() error {
	for _, filter := range config.propertyTransformers {
		if !filter.MatchedRequiredTypes() {
			return errors.Errorf("Type transformer target: %q for property %q matched no types", filter.String(), filter.Property)
		}
	}

	return nil
}

func (config *Configuration) GetExportFiltersError() error {
	for _, filter := range config.ExportFilters {
		if !filter.MatchedRequiredTypes() {
			return errors.Errorf("Export filter action: %q, target: %q matched no types", filter.Action, filter.String())
		}
	}

	return nil
}

// NewConfiguration returns a new empty Configuration
func NewConfiguration() *Configuration {
	return &Configuration{}
}

// LoadConfiguration loads a `Configuration` from the specified file
func LoadConfiguration(configurationFile string) (*Configuration, error) {
	data, err := os.ReadFile(configurationFile)
	if err != nil {
		return nil, err
	}

	result := &Configuration{}

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

// ShouldExportResult is returned by ShouldExport to indicate whether the supplied type should be exported
type ShouldExportResult string

const (
	// Export indicates the specified type should be exported to disk
	Export ShouldExportResult = "export"
	// Skip indicates the specified type should be skipped and not exported
	Skip ShouldExportResult = "skip"
)

// ShouldPruneResult is returned by ShouldPrune to indicate whether the supplied type should be exported
type ShouldPruneResult string

const (
	// Include indicates the specified type should be included in the type graph
	Include ShouldPruneResult = "include"
	// Prune indicates the type (and all types only referenced by it) should be pruned from the type graph
	Prune ShouldPruneResult = "prune"
)

// WithExportFilters adds the provided ExportFilters to the configurations collection of ExportFilters
func (config *Configuration) WithExportFilters(filters ...*ExportFilter) *Configuration {
	result := *config
	result.ExportFilters = append(result.ExportFilters, filters...)

	return &result
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

		toURL, err := url.Parse(rewrite.To)
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

	if modPath, err := getModulePathFromModFile(config.DestinationGoModuleFile); err != nil {
		errs = append(errs, err)
	} else {
		config.GoModulePath = modPath
	}

	for _, filter := range config.ExportFilters {
		err := filter.Initialize()
		if err != nil {
			errs = append(errs, err)
		}
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

		if transformer.Property != "" {
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

type ExportFilterFunc func(astmodel.TypeName) (result ShouldExportResult, because string)

func buildExportFilterFunc(f *ExportFilter, allTypes astmodel.Types) ExportFilterFunc {
	switch f.Action {
	case ExportFilterExclude:
		return func(typeName astmodel.TypeName) (ShouldExportResult, string) {
			if f.AppliesToType(typeName) {
				return Skip, f.Because
			}

			return "", ""
		}

	case ExportFilterInclude:
		return func(typeName astmodel.TypeName) (ShouldExportResult, string) {
			if f.AppliesToType(typeName) {
				return Export, f.Because
			}

			return "", ""
		}

	case ExportFilterIncludeTransitive:
		applicableTypes := astmodel.NewTypeNameSet()
		for tn := range allTypes {
			if f.AppliesToType(tn) {
				collectAllReferencedTypes(allTypes, tn, applicableTypes)
			}
		}

		return func(typeName astmodel.TypeName) (ShouldExportResult, string) {
			if applicableTypes.Contains(typeName) {
				return Export, f.Because
			}

			return "", ""
		}

	default:
		panic(errors.Errorf("unknown exportfilter directive: %s", f.Action))
	}
}

func collectAllReferencedTypes(allTypes astmodel.Types, root astmodel.TypeName, output astmodel.TypeNameSet) {
	output.Add(root)
	for referenced := range allTypes[root].Type().References() {
		if !output.Contains(referenced) {
			collectAllReferencedTypes(allTypes, referenced, output)
		}
	}
}

// BuildExportFilterer tests for whether a given type should be exported as Go code
// Returns a result indicating whether export should occur as well as a reason for logging
func (config *Configuration) BuildExportFilterer(allTypes astmodel.Types) ExportFilterFunc {
	var filters []ExportFilterFunc
	for _, f := range config.ExportFilters {
		filter := buildExportFilterFunc(f, allTypes)
		filters = append(filters, filter)
	}

	return func(typeName astmodel.TypeName) (result ShouldExportResult, because string) {
		for _, filter := range filters {
			result, because := filter(typeName)
			if result != "" {
				return result, because
			}
		}

		// By default we export all types
		return Export, ""
	}
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

	// By default we include all types
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
	return astmodel.MakeLocalPackageReference(config.LocalPathPrefix(), group, version)
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

	// A suffix to add on to the group name
	Suffix string `yaml:"suffix"`
}

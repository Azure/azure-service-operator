/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

const AddOperatorSpecStageID = "addOperatorSpec"

func AddOperatorSpec(configuration *config.Configuration, idFactory astmodel.IdentifierFactory) *Stage {
	return NewStage(
		AddOperatorSpecStageID,
		"Adds the property 'OperatorSpec' to all Spec types that require it",
		func(ctx context.Context, state *State) (*State, error) {
			defs := state.Definitions()
			result := make(astmodel.TypeDefinitionSet)

			// ConfigMappings
			exportedTypeNameConfigMaps := NewExportedTypeNameProperties()

			for _, resource := range astmodel.FindResourceDefinitions(defs) {
				newDefs, exportedConfigMaps, err := createOperatorSpecIfNeeded(defs, configuration, idFactory, resource)
				if err != nil {
					return nil, err
				}
				result.AddTypes(newDefs)
				exportedTypeNameConfigMaps.Add(resource.Name(), exportedConfigMaps)
			}

			// confirm that all the Azure generated secrets were used. Note that this also indirectly confirms that
			// this property was only used on resources, since that's the only place we try to check it from. If it's
			// set on anything else it will be labeled unconsumed.
			err := configuration.ObjectModelConfiguration.AzureGeneratedSecrets.VerifyConsumed()
			if err != nil {
				return nil, err
			}
			err = configuration.ObjectModelConfiguration.GeneratedConfigs.VerifyConsumed()
			if err != nil {
				return nil, err
			}
			err = configuration.ObjectModelConfiguration.ManualConfigs.VerifyConsumed()
			if err != nil {
				return nil, err
			}

			result = defs.OverlayWith(result)

			return state.WithDefinitions(result).WithGeneratedConfigMaps(exportedTypeNameConfigMaps), nil
		})
}

func createOperatorSpecIfNeeded(
	defs astmodel.TypeDefinitionSet,
	configuration *config.Configuration,
	idFactory astmodel.IdentifierFactory,
	resource astmodel.TypeDefinition) (astmodel.TypeDefinitionSet, ExportedProperties, error) {

	resolved, err := defs.ResolveResourceSpecAndStatus(resource)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "resolving resource spec and status for %s", resource.Name())
	}

	hasSecrets := false
	secrets, err := configuration.ObjectModelConfiguration.AzureGeneratedSecrets.Lookup(resolved.ResourceDef.Name())
	if err == nil {
		hasSecrets = true
	} else if err != nil {
		// If error is just that there's no configured secrets, proceed
		if !config.IsNotConfiguredError(err) {
			return nil, nil, errors.Wrapf(err, "reading azureGeneratedSecrets for %s", resolved.ResourceDef.Name())
		}
	}

	configs, exportedProperties, err := getConfigMapProperties(defs, configuration, resource)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "finding properties allowed to export as config maps")
	}

	hasConfigMapProperties := len(configs) != 0

	if !hasSecrets && !hasConfigMapProperties {
		// We don't need to make an OperatorSpec type
		return nil, nil, nil
	}

	builder := newOperatorSpecBuilder(configuration, idFactory, resolved.ResourceDef)
	builder.addSecretsToOperatorSpec(secrets)
	builder.addConfigs(configs)

	operatorSpec := builder.build()

	propInjector := astmodel.NewPropertyInjector()
	updatedDef, err := propInjector.Inject(resolved.SpecDef, builder.newOperatorSpecProperty(operatorSpec))
	if err != nil {
		return nil, nil, errors.Wrapf(err, "couldn't add OperatorSpec to spec %q", resolved.SpecDef.Name())
	}

	result := make(astmodel.TypeDefinitionSet)

	result.Add(updatedDef)
	result.Add(operatorSpec)
	result.AddTypes(builder.definitions) // Add any other types that were needed as well

	return result, exportedProperties, nil
}

type configMapContext struct {
	path     []*astmodel.PropertyDefinition
	typeName astmodel.TypeName
}

func (ctx configMapContext) withPathElement(prop *astmodel.PropertyDefinition) configMapContext {
	return configMapContext{
		typeName: ctx.typeName,
		path:     append(ctx.path, prop),
	}
}

func (ctx configMapContext) withTypeName(typeName astmodel.TypeName) configMapContext {
	return configMapContext{
		typeName: typeName,
		path:     ctx.path,
	}
}

var identityConfigMapObjectTypeVisit = astmodel.MakeIdentityVisitOfObjectType(
	func(ot *astmodel.ObjectType, prop *astmodel.PropertyDefinition, ctx configMapContext) (configMapContext, error) {
		return ctx.withPathElement(prop), nil
	})

type configMapTypeWalker struct {
	configuration        *config.Configuration
	configuredProperties map[string]string
	exportedProperties   ExportedProperties
	walker               *astmodel.TypeWalker[configMapContext]
}

func newConfigMapTypeWalker(
	defs astmodel.TypeDefinitionSet,
	paths map[string]string,
) *configMapTypeWalker {

	result := &configMapTypeWalker{
		configuredProperties: paths,
		exportedProperties:   make(ExportedProperties),
	}

	visitor := astmodel.TypeVisitorBuilder[configMapContext]{
		VisitObjectType:   result.catalogObjectConfigMapProperties,
		VisitResourceType: result.includeSpecStatus,
	}.Build()
	walker := astmodel.NewTypeWalker(defs, visitor)
	walker.MakeContext = func(it astmodel.TypeName, ctx configMapContext) (configMapContext, error) {
		if ctx.typeName == nil {
			return configMapContext{
				typeName: it,
				path:     nil,
			}, nil
		}

		return ctx.withTypeName(it), nil
	}
	result.walker = walker

	return result
}

func (w *configMapTypeWalker) includeSpecStatus(
	this *astmodel.TypeVisitor[configMapContext],
	it *astmodel.ResourceType,
	ctx configMapContext,
) (astmodel.Type, error) {
	specProp, ok := it.Property("Spec")
	if !ok {
		return nil, errors.Errorf("couldn't find resource spec")
	}
	_, err := this.Visit(it.SpecType(), ctx.withPathElement(specProp))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to visit resource spec type %q", it.SpecType())
	}

	statusProp, ok := it.Property("Status")
	if !ok {
		return nil, errors.Errorf("couldn't find resource status")
	}
	_, err = this.Visit(it.StatusType(), ctx.withPathElement(statusProp))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to visit resource status type %q", it.StatusType())
	}

	// We're not planning on actually modifying any types here, so we can just return the type we started with
	return it, nil
}

func (w *configMapTypeWalker) catalogObjectConfigMapProperties(
	this *astmodel.TypeVisitor[configMapContext],
	ot *astmodel.ObjectType,
	ctx configMapContext,
) (astmodel.Type, error) {
	var errs []error
	ot.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
		path := append(ctx.path, prop)

		// Transform the path into a string and check if we have that path configured
		pathStr := makeJSONPathFromProps(path)
		for propName, propPath := range w.configuredProperties {
			if propPath == pathStr {
				w.exportedProperties[propName] = path
			}
		}
	})

	// Propagate errors
	err := kerrors.NewAggregate(errs)
	if err != nil {
		return nil, err
	}

	// Now run the identity
	return identityConfigMapObjectTypeVisit(this, ot, ctx)
}

func (w *configMapTypeWalker) Walk(def astmodel.TypeDefinition) (ExportedProperties, error) {
	_, err := w.walker.Walk(def)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to walk definition %s", def.Name())
	}

	return w.exportedProperties, nil
}

// TODO: Consider actually using a real JSONPath library here...?
func makeJSONPathFromProps(props []*astmodel.PropertyDefinition) string {
	builder := strings.Builder{}
	builder.WriteString("$") // Always starts with a "$" which is the root
	for _, prop := range props {
		builder.WriteString(".") // Always starts with a dot
		builder.WriteString(prop.PropertyName().String())
	}

	return builder.String()
}

func getConfigMapProperties(
	defs astmodel.TypeDefinitionSet,
	configuration *config.Configuration,
	resource astmodel.TypeDefinition) ([]string, ExportedProperties, error) {

	configMapPaths, err := configuration.ObjectModelConfiguration.GeneratedConfigs.Lookup(resource.Name())
	if err != nil {
		// If error is just that there's no configured secrets, proceed
		if !config.IsNotConfiguredError(err) {
			return nil, nil, errors.Wrapf(err, "reading generatedConfigs for %s", resource.Name())
		}
	}

	additionalConfigMaps, err := configuration.ObjectModelConfiguration.ManualConfigs.Lookup(resource.Name())
	if err != nil {
		// If error is just that there's no configured secrets, proceed
		if !config.IsNotConfiguredError(err) {
			return nil, nil, errors.Wrapf(err, "reading manualConfigs for %s", resource.Name())
		}
	}

	// Fast out if we don't have anything configured
	if len(configMapPaths) == 0 && len(additionalConfigMaps) == 0 {
		return nil, nil, nil
	}

	walker := newConfigMapTypeWalker(defs, configMapPaths)
	exportedConfigMapProperties, err := walker.Walk(resource)
	if err != nil {
		return nil, nil, err
	}

	// There should be an exported configMap property for every configured configMapPath
	for name, path := range configMapPaths {
		if _, ok := exportedConfigMapProperties[name]; !ok {
			return nil, nil, errors.Errorf("$generatedConfigs property %q not found at path %q", name, path)
		}
	}

	result := make([]string, 0, len(exportedConfigMapProperties)+len(additionalConfigMaps))
	for key := range exportedConfigMapProperties {
		result = append(result, key)
	}
	result = append(result, additionalConfigMaps...)

	return result, exportedConfigMapProperties, nil
}

type operatorSpecBuilder struct {
	idFactory     astmodel.IdentifierFactory
	configuration *config.Configuration
	resource      astmodel.TypeDefinition
	definitions   astmodel.TypeDefinitionSet
	operatorSpec  astmodel.TypeDefinition
}

func newOperatorSpecBuilder(
	configuration *config.Configuration,
	idFactory astmodel.IdentifierFactory,
	resource astmodel.TypeDefinition) *operatorSpecBuilder {
	result := &operatorSpecBuilder{
		idFactory:     idFactory,
		configuration: configuration,
		resource:      resource,
		definitions:   make(astmodel.TypeDefinitionSet),
	}

	result.operatorSpec = result.newEmptyOperatorSpec()

	return result
}

func (b *operatorSpecBuilder) newEmptyOperatorSpec() astmodel.TypeDefinition {
	name := b.idFactory.CreateIdentifier(b.resource.Name().Name()+"OperatorSpec", astmodel.Exported)
	operatorSpecTypeName := b.resource.Name().WithName(name).(astmodel.InternalTypeName)
	operatorSpec := astmodel.NewObjectType()

	operatorSpecDefinition := astmodel.MakeTypeDefinition(operatorSpecTypeName, operatorSpec)
	description := "Details for configuring operator behavior. Fields in this struct are " +
		"interpreted by the operator directly rather than being passed to Azure"
	operatorSpecDefinition = operatorSpecDefinition.WithDescription(description)

	return operatorSpecDefinition
}

func (b *operatorSpecBuilder) newOperatorSpecProperty(operatorSpec astmodel.TypeDefinition) *astmodel.PropertyDefinition {
	prop := astmodel.NewPropertyDefinition(
		astmodel.OperatorSpecProperty,
		b.idFactory.CreateStringIdentifier(astmodel.OperatorSpecProperty, astmodel.NotExported),
		operatorSpec.Name()).MakeTypeOptional()
	desc := "The specification for configuring operator behavior. " +
		"This field is interpreted by the operator and not passed directly to Azure"
	prop = prop.WithDescription(desc)

	return prop
}

func (b *operatorSpecBuilder) newSecretsProperty(secretsTypeName astmodel.TypeName) *astmodel.PropertyDefinition {
	secretProp := astmodel.NewPropertyDefinition(
		b.idFactory.CreatePropertyName(astmodel.OperatorSpecSecretsProperty, astmodel.Exported),
		b.idFactory.CreateStringIdentifier(astmodel.OperatorSpecSecretsProperty, astmodel.NotExported),
		secretsTypeName)
	secretProp = secretProp.WithDescription("configures where to place Azure generated secrets.")
	secretProp = secretProp.MakeTypeOptional()

	return secretProp
}

func (b *operatorSpecBuilder) newConfigMapProperty(configMapTypeName astmodel.TypeName) *astmodel.PropertyDefinition {
	configMapProp := astmodel.NewPropertyDefinition(
		b.idFactory.CreatePropertyName(astmodel.OperatorSpecConfigMapsProperty, astmodel.Exported),
		b.idFactory.CreateStringIdentifier(astmodel.OperatorSpecConfigMapsProperty, astmodel.NotExported),
		configMapTypeName)
	configMapProp = configMapProp.WithDescription("configures where to place operator written ConfigMaps.")
	configMapProp = configMapProp.MakeTypeOptional()

	return configMapProp
}

func (b *operatorSpecBuilder) addSecretsToOperatorSpec(
	azureGeneratedSecrets []string) {

	if len(azureGeneratedSecrets) == 0 {
		return // Nothing to do
	}

	operatorSpec, ok := astmodel.AsObjectType(b.operatorSpec.Type())
	if !ok {
		panic(fmt.Sprintf("OperatorSpec %q was not an ObjectType, which is impossible", b.operatorSpec.Name()))
	}

	// Create a new "secrets" type to hold the secrets
	resourceName := b.resource.Name()
	secretsTypeName := resourceName.WithName(
		b.idFactory.CreateIdentifier(
			resourceName.Name()+"OperatorSecrets",
			astmodel.Exported)).(astmodel.InternalTypeName)
	secretsType := astmodel.NewObjectType()

	// Add the "secrets" property to the operator spec
	secretProp := b.newSecretsProperty(secretsTypeName)
	operatorSpec = operatorSpec.WithProperty(secretProp)

	for _, secret := range azureGeneratedSecrets {
		prop := astmodel.NewPropertyDefinition(
			b.idFactory.CreatePropertyName(secret, astmodel.Exported),
			b.idFactory.CreateStringIdentifier(secret, astmodel.NotExported),
			astmodel.SecretDestinationType).MakeTypeOptional()
		desc := fmt.Sprintf(
			"indicates where the %s secret should be placed. If omitted, the secret will not be retrieved from Azure.",
			secret)
		prop = prop.WithDescription(desc)
		prop = prop.MakeOptional()
		secretsType = secretsType.WithProperty(prop)
	}

	secretsTypeDef := astmodel.MakeTypeDefinition(secretsTypeName, secretsType)
	b.definitions.Add(secretsTypeDef)

	b.operatorSpec = b.operatorSpec.WithType(operatorSpec)
}

func (b *operatorSpecBuilder) addConfigs(
	exportedConfigs []string) {

	if len(exportedConfigs) == 0 {
		return // Nothing to do
	}

	operatorSpec, ok := astmodel.AsObjectType(b.operatorSpec.Type())
	if !ok {
		panic(fmt.Sprintf("OperatorSpec %q was not an ObjectType, which is impossible", b.operatorSpec.Name()))
	}

	// Create a new "ConfigMaps" type to hold the config map values
	resourceName := b.resource.Name()
	configMapTypeName := resourceName.WithName(
		b.idFactory.CreateIdentifier(
			resourceName.Name()+"OperatorConfigMaps",
			astmodel.Exported)).(astmodel.InternalTypeName)
	configMapsType := astmodel.NewObjectType()

	// Add the "configMaps" property to the operator spec
	configMapProp := b.newConfigMapProperty(configMapTypeName)
	operatorSpec = operatorSpec.WithProperty(configMapProp)

	for _, exportedConfig := range exportedConfigs {
		prop := astmodel.NewPropertyDefinition(
			b.idFactory.CreatePropertyName(exportedConfig, astmodel.Exported),
			b.idFactory.CreateStringIdentifier(exportedConfig, astmodel.NotExported),
			astmodel.ConfigMapDestinationType).MakeTypeOptional()
		desc := fmt.Sprintf(
			"indicates where the %s config map should be placed. If omitted, no config map will be created.",
			exportedConfig)
		prop = prop.WithDescription(desc)
		prop = prop.MakeOptional()
		configMapsType = configMapsType.WithProperty(prop)
	}

	configMapTypeDef := astmodel.MakeTypeDefinition(configMapTypeName, configMapsType)
	b.definitions.Add(configMapTypeDef)

	b.operatorSpec = b.operatorSpec.WithType(operatorSpec)
}

func (b *operatorSpecBuilder) build() astmodel.TypeDefinition {
	return b.operatorSpec
}

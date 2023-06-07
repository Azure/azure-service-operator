/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// configurationVisitor is used to facilitate easy walking of the ObjectModelConfiguration hierarchy, abstracting
// away traversal logic so that new uses of the hierarchy can concentrate on their specific functionality.
// By default, will traverse the entire configuration but may optionally be constrained to just a specific type by
// construction with a typeName, or to a property by also providing the name of the property.
// Only one handler should be present, as we don't do any traversal below an invoked handler (but a handler is free to
// do independent visiting with a different instance if it chooses)
type configurationVisitor struct {
	ref            astmodel.PackageReference                         // Optional Package reference used to constrain the walk
	typeName       string                                            // Optional TypeName used to constrain the walk
	property       *astmodel.PropertyName                            // Optional PropertyName used to constrain the walk
	handleGroup    func(groupConfig *GroupConfiguration) error       // Optional handler for visiting a group
	handleVersion  func(versionConfig *VersionConfiguration) error   // Optional handler for visiting a version
	handleType     func(typeConfig *TypeConfiguration) error         // Optional handler for visiting a Type
	handleProperty func(propertyConfig *PropertyConfiguration) error // Optional handler for visiting a property
}

// newSinglePropertyConfigurationVisitor creates a ConfigurationVisitor to apply an action to the property specified.
// typeName is the fully qualified name of the type expected to contain the property.
// property is the name of the property to visit.
// action is the action to apply to that property.
// Returns (true, nil) if the property is found and the action successfully applied, (true, error) if the action returns
// an error, and (false, nil) if the type or property does not exist.
func newSinglePropertyConfigurationVisitor(
	typeName astmodel.TypeName,
	property astmodel.PropertyName,
	action func(configuration *PropertyConfiguration) error) *configurationVisitor {
	return &configurationVisitor{
		ref:            typeName.PackageReference,
		typeName:       typeName.Name(),
		property:       &property,
		handleProperty: action,
	}
}

// newEveryPropertyConfigurationVisitor creates a ConfigurationVisitor to apply an action to every property
// configuration we have.
// action is the action to apply to each property.
// Returns nil if every call to action was successful (returned nil); otherwise returns an aggregated error containing
// all the errors returned.
func newEveryPropertyConfigurationVisitor(
	action func(configuration *PropertyConfiguration) error) *configurationVisitor {
	return &configurationVisitor{
		handleProperty: action,
	}
}

// newSingleTypeConfigurationVisitor creates a ConfigurationVisitor to apply an action to the type specified.
// typeName is the fully qualified name of the type expected.
// action is the action to apply to that type.
// Returns (true, nil) if the type is found and the action successfully applied, (true, error) if the action returns
// an error, and (false, nil) if the type does not exist.
func newSingleTypeConfigurationVisitor(
	typeName astmodel.TypeName,
	action func(configuration *TypeConfiguration) error) *configurationVisitor {
	return &configurationVisitor{
		ref:        typeName.PackageReference,
		typeName:   typeName.Name(),
		handleType: action,
	}
}

// newEveryTypeConfigurationVisitor creates a ConfigurationVisitor to apply an action to every type configuration
// specified.
// action is the action to apply to each type.
// Returns nil if every call to action returned nil; otherwise returns an aggregated error containing all the errors returned.
func newEveryTypeConfigurationVisitor(
	action func(configuration *TypeConfiguration) error) *configurationVisitor {
	return &configurationVisitor{
		handleType: action,
	}
}

// newSingleVersionConfigurationVisitor creates a ConfigurationVisitor to apply an action to the version specified
// typeName is the fully qualified name of the type expected.
// action is the action to apply to that type.
// Returns (true, nil) if the type is found and the action successfully applied, (true, error) if the action returns
// an error, and (false, nil) if the type does not exist.
func newSingleVersionConfigurationVisitor(
	ref astmodel.PackageReference,
	action func(configuration *VersionConfiguration) error) *configurationVisitor {
	return &configurationVisitor{
		ref:           ref,
		handleVersion: action,
	}
}

// newEveryVersionConfigurationVisitor creates a ConfigurationVisitor to apply an action to every version configuration.
// action is the action to apply to each version.
func newEveryVersionConfigurationVisitor(
	action func(configuration *VersionConfiguration) error) *configurationVisitor {
	return &configurationVisitor{
		handleVersion: action,
	}
}

// newEveryGroupConfigurationVisitor creates a ConfigurationVisitor to apply an action to every group configuration
// specified.
// action is the action to apply to each group.
// Returns nil if every call to action returned nil; otherwise returns an aggregated error containing all the errors returned.
func newEveryGroupConfigurationVisitor(
	action func(configuration *GroupConfiguration) error) *configurationVisitor {
	return &configurationVisitor{
		handleGroup: action,
	}
}

// newSingleGroupConfigurationVisitor creates a ConfigurationVisitor to apply an action to the group specified
// ref is the package reference of the group expected.
// action is the action to apply to that group.
// Returns (true, nil) if the group is found and the action successfully applied, (true, error) if the action returns
// an error, and (false, nil) if the group does not exist.
func newSingleGroupConfigurationVisitor(
	ref astmodel.PackageReference,
	action func(configuration *GroupConfiguration) error) *configurationVisitor {
	return &configurationVisitor{
		ref:         ref,
		handleGroup: action,
	}
}

// newEveryGroupConfigurationVisitor creates a ConfigurationVisitor to apply an action to every group configuration.
// action is the action to apply to each group.
func newEveryGroupConfigurationVisitor(
	action func(configuration *GroupConfiguration) error) *configurationVisitor {
	return &configurationVisitor{
		handleGroup: action,
	}
}

// Visit visits the specified ObjectModelConfiguration.
func (v *configurationVisitor) Visit(omc *ObjectModelConfiguration) error {
	if v.ref != nil {
		return omc.visitGroup(v.ref, v)
	}

	return omc.visitGroups(v)
}

// visitGroup visits the specified group configuration.
// If a group handler is present, it's called. Otherwise, if we're interested in precisely one nested version, we visit
// that. Otherwise, we visit all nested versions.
func (v *configurationVisitor) visitGroup(groupConfig *GroupConfiguration) error {
	if v.handleGroup != nil {
		return v.handleGroup(groupConfig)
	}

	if v.ref != nil {
		return groupConfig.visitVersion(v.ref, v)
	}

	return groupConfig.visitVersions(v)
}

// visitVersion visits the specified version configuration.
// If a version handler is present, it's called. Otherwise, if we're interested in precisely one nested type, we visit
// that. Otherwise, we visit all nested types.
func (v *configurationVisitor) visitVersion(versionConfig *VersionConfiguration) error {
	if v.handleVersion != nil {
		return v.handleVersion(versionConfig)
	}

	if v.typeName != "" {
		return versionConfig.visitType(v.typeName, v)
	}

	return versionConfig.visitTypes(v)
}

// visitType visits the specified type configuration.
// If a type handler is present, it's called. Otherwise, if we're interested in precisely one property, we visit that.
// Otherwise, we visit all nested properties.
func (v *configurationVisitor) visitType(typeConfig *TypeConfiguration) error {
	if v.handleType != nil {
		return v.handleType(typeConfig)
	}

	if v.property != nil {
		return typeConfig.visitProperty(*v.property, v)
	}

	return typeConfig.visitProperties(v)
}

// visitProperty visits the specified property configuration. If a property handler is present, it's called.
func (v *configurationVisitor) visitProperty(propertyConfig *PropertyConfiguration) error {
	if v.handleProperty != nil {
		return v.handleProperty(propertyConfig)
	}

	return nil
}

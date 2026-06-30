/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cel

import (
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
)

// ResourceType represents an Azure resource type, for use in CEL expressions.
type ResourceType struct {
	Namespace string `json:"namespace,omitempty"`
	Type      string `json:"type,omitempty"`
}

// ResourceID represents a parsed Azure resource ID, for use in CEL expressions.
type ResourceID struct {
	Parent            *ResourceID  `json:"parent,omitempty"`
	SubscriptionID    string       `json:"subscriptionId,omitempty"`
	ResourceGroupName string       `json:"resourceGroupName,omitempty"`
	Provider          string       `json:"provider,omitempty"`
	Location          string       `json:"location,omitempty"`
	ResourceType      ResourceType `json:"resourceType,omitempty"`
	Name              string       `json:"name,omitempty"`
}

// toResourceID converts an arm.ResourceID to a ResourceID.
func toResourceID(id *arm.ResourceID) *ResourceID {
	if id == nil {
		return nil
	}

	return &ResourceID{
		Parent:            toResourceID(id.Parent),
		SubscriptionID:    id.SubscriptionID,
		ResourceGroupName: id.ResourceGroupName,
		Provider:          id.Provider,
		Location:          id.Location,
		ResourceType: ResourceType{
			Namespace: id.ResourceType.Namespace,
			Type:      id.ResourceType.Type,
		},
		Name: id.Name,
	}
}

// resourceIDNativeTypes returns the NativeTypesOptions needed to register CELResourceID and CELResourceType
// with the CEL type system. These should be passed to the main ext.NativeTypes call.
func resourceIDNativeTypes() []any {
	return []any{
		reflect.TypeOf(ResourceID{}),
		reflect.TypeOf(ResourceType{}),
	}
}

// parseResourceIDFunction returns a cel.EnvOption that adds the aso.parseResourceId(string) function
// to the CEL environment. It must be added after ext.NativeTypes and newJSONProvider so that
// the environment's type adapter can handle the ResourceID return type.
func parseResourceIDFunction() cel.EnvOption {
	return func(env *cel.Env) (*cel.Env, error) {
		adapter := env.CELTypeAdapter()
		fn := cel.Function(
			"aso.parseResourceId",
			cel.Overload(
				"aso_parseResourceId_string",
				[]*cel.Type{cel.StringType},
				cel.ObjectType("cel.ResourceID"),
				cel.UnaryBinding(func(val ref.Val) ref.Val {
					id, ok := val.Value().(string)
					if !ok {
						return types.NewErr("aso.parseResourceId: expected string argument")
					}

					parsed, err := arm.ParseResourceID(id)
					if err != nil {
						return types.NewErr("aso.parseResourceId: %s", err.Error())
					}

					return adapter.NativeToValue(toResourceID(parsed))
				}),
			),
		)
		return fn(env)
	}
}

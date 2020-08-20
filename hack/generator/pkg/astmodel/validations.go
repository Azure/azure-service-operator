/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"reflect"
	"strings"
)

// Validation represents some kind of data validation on a property
type Validation struct {
	name string

	// value is a bool, an int, or a string, or a list of those
	value interface{}
}

// GenerateKubebuilderComment converts the given validation to
// a Kubebuilder magic comment
func GenerateKubebuilderComment(validation Validation) string {
	const prefix = "// +kubebuilder:validation:"

	if validation.value != nil {
		value := reflect.ValueOf(validation.value)

		if value.Kind() == reflect.Slice {
			// handle slice values which should look like "{"x","y","z"}
			var values []string
			for i := 0; i < value.Len(); i++ {
				values = append(values, fmt.Sprintf("%v", value.Index(i)))
			}

			return fmt.Sprintf("%s%s={%s}", prefix, validation.name, strings.Join(values, ","))
		}

		// everything else
		return fmt.Sprintf("%s%s=%v", prefix, validation.name, validation.value)
	}

	// validation without argument
	return fmt.Sprintf("%s%s", prefix, validation.name)
}

func (v Validation) HasName(name string) bool {
	return v.name == name
}

func (v Validation) Equals(other Validation) bool {
	return v.name == other.name
}

/*
 * Constants for names of validation
 */

const (
	EnumValidationName     string = "Enum"
	RequiredValidationName string = "Required"
)

/*
 * Factory methods for Validation instances
 */

// ValidateEnum returns a Validation that requires the value be one of the
// passed 'permittedValues'
func ValidateEnum(permittedValues []interface{}) Validation {
	return Validation{EnumValidationName, permittedValues}
}

// ValidateRequired returns a Validation that requires a value be present
func ValidateRequired() Validation {
	return Validation{RequiredValidationName, nil}
}

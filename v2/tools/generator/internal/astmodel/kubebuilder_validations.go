/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"math/big"
	"reflect"
	"regexp"
	"strings"

	"github.com/dave/dst"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
)

// KubeBuilderValidation represents some kind of data validation on a property
type KubeBuilderValidation struct {
	name string

	// value is a bool, an int, or a string, or a list of those
	value interface{}
}

// GenerateKubebuilderComment converts the given validation to
// a Kubebuilder magic comment
func GenerateKubebuilderComment(validation KubeBuilderValidation) string {
	const prefix = "// +kubebuilder:validation:"

	if validation.value != nil {
		value := reflect.ValueOf(validation.value)

		if value.Kind() == reflect.Slice {
			// handle slice values which should look like "{"x","y","z"}
			var values []string
			for i := 0; i < value.Len(); i++ {
				values = append(values, valueAsString(value.Index(i)))
			}

			return fmt.Sprintf("%s%s={%s}", prefix, validation.name, strings.Join(values, ","))
		}

		// everything else
		return fmt.Sprintf("%s%s=%s", prefix, validation.name, valueAsString(value))
	}

	// validation without argument
	return fmt.Sprintf("%s%s", prefix, validation.name)
}

func valueAsString(value reflect.Value) string {
	switch v := value.Interface().(type) {
	case int, int16, int32, int64:
		return fmt.Sprintf("%d", v)
	case bool:
		return fmt.Sprintf("%t", v)
	case string:
		return v
	default:
		panic(fmt.Sprintf("unexpected value for kubebuilder comment - %s", value.Kind()))
	}
}

func AddValidationComments(commentList *dst.Decorations, validations []KubeBuilderValidation) {
	// generate validation comments:
	for _, validation := range validations {
		// these are not doc comments but they must go here to be emitted before the property
		astbuilder.AddComment(commentList, GenerateKubebuilderComment(validation))
	}
}

func (v KubeBuilderValidation) HasName(name string) bool {
	return v.name == name
}

func (v KubeBuilderValidation) Equals(other KubeBuilderValidation) bool {
	return v.name == other.name
}

var _ fmt.Stringer = KubeBuilderValidation{}

func (v KubeBuilderValidation) String() string {
	value := reflect.ValueOf(v.value)
	return fmt.Sprintf("%s: %s", v.name, valueAsString(value))
}

/*
 * Constants for names of validation
 */

// see: https://book.kubebuilder.io/reference/markers/crd-validation.html
const (
	EnumValidationName     string = "Enum"
	RequiredValidationName string = "Required"

	// Strings:
	MinLengthValidationName string = "MinLength"
	MaxLengthValidationName string = "MaxLength"
	PatternValidationName   string = "Pattern"

	// Arrays:
	MinItemsValidationName    string = "MinItems"
	MaxItemsValidationName    string = "MaxItems"
	UniqueItemsValidationName string = "UniqueItems"

	// Numbers:
	MaximumValidationName          string = "Maximum"
	MinimumValidationName          string = "Minimum"
	ExclusiveMaximumValidationName string = "ExclusiveMaximum"
	ExclusiveMinimumValidationName string = "ExclusiveMinimum"
	MultipleOfValidationName       string = "MultipleOf"
)

/*
 * Factory methods for Validation instances
 */

// MakeEnumValidation returns a Validation that requires the value be one of the
// passed 'permittedValues'
func MakeEnumValidation(permittedValues []interface{}) KubeBuilderValidation {
	return KubeBuilderValidation{EnumValidationName, permittedValues}
}

// MakeRequiredValidation returns a Validation that requires a value be present
func MakeRequiredValidation() KubeBuilderValidation {
	return KubeBuilderValidation{RequiredValidationName, nil}
}

func MakeMinLengthValidation(length int64) KubeBuilderValidation {
	return KubeBuilderValidation{MinLengthValidationName, length}
}

func MakeMaxLengthValidation(length int64) KubeBuilderValidation {
	return KubeBuilderValidation{MaxLengthValidationName, length}
}

func MakePatternValidation(pattern *regexp.Regexp) KubeBuilderValidation {
	return KubeBuilderValidation{PatternValidationName, fmt.Sprintf("%q", pattern.String())}
}

func MakeMinItemsValidation(length int64) KubeBuilderValidation {
	return KubeBuilderValidation{MinItemsValidationName, length}
}

func MakeMaxItemsValidation(length int64) KubeBuilderValidation {
	return KubeBuilderValidation{MaxItemsValidationName, length}
}

func MakeUniqueItemsValidation() KubeBuilderValidation {
	return KubeBuilderValidation{UniqueItemsValidationName, true}
}

func MakeMaximumValidation(value *big.Rat) KubeBuilderValidation {
	if value.IsInt() {
		return KubeBuilderValidation{MaximumValidationName, value.RatString()}
	} else {
		floatValue, ok := value.Float64()
		if !ok {
			klog.Warningf("inexact maximum: %s ⇒ %g", value.String(), floatValue)
		}

		return KubeBuilderValidation{MaximumValidationName, floatValue}
	}
}

func MaxMinimumValidation(value *big.Rat) KubeBuilderValidation {
	if value.IsInt() {
		return KubeBuilderValidation{MinimumValidationName, value.RatString()}
	} else {
		floatValue, ok := value.Float64()
		if !ok {
			klog.Warningf("inexact minimum: %s ⇒ %g", value.String(), floatValue)
		}

		return KubeBuilderValidation{MinimumValidationName, floatValue}
	}
}

func MakeExclusiveMaxiumValidation() KubeBuilderValidation {
	return KubeBuilderValidation{ExclusiveMaximumValidationName, true}
}

func MakeExclusiveMinimumValidation() KubeBuilderValidation {
	return KubeBuilderValidation{ExclusiveMinimumValidationName, true}
}

func MakeMultipleOfValidation(value *big.Rat) KubeBuilderValidation {
	if value.IsInt() {
		return KubeBuilderValidation{MultipleOfValidationName, value.RatString()}
	} else {
		floatValue, ok := value.Float64()
		if !ok {
			klog.Warningf("inexact multiple-of: %s ⇒ %g", value.String(), floatValue)
		}

		return KubeBuilderValidation{MultipleOfValidationName, floatValue}
	}
}

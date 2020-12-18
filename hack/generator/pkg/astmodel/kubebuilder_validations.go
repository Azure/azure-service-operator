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

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	ast "github.com/dave/dst"
	"k8s.io/klog/v2"
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

func AddValidationComments(commentList *ast.Decorations, validations []KubeBuilderValidation) {
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

// ValidateEnum returns a Validation that requires the value be one of the
// passed 'permittedValues'
func ValidateEnum(permittedValues []interface{}) KubeBuilderValidation {
	return KubeBuilderValidation{EnumValidationName, permittedValues}
}

// ValidateRequired returns a Validation that requires a value be present
func ValidateRequired() KubeBuilderValidation {
	return KubeBuilderValidation{RequiredValidationName, nil}
}

func ValidateMinLength(length int64) KubeBuilderValidation {
	return KubeBuilderValidation{MinLengthValidationName, length}
}

func ValidateMaxLength(length int64) KubeBuilderValidation {
	return KubeBuilderValidation{MaxLengthValidationName, length}
}

func ValidatePattern(pattern regexp.Regexp) KubeBuilderValidation {
	return KubeBuilderValidation{PatternValidationName, fmt.Sprintf("%q", pattern.String())}
}

func ValidateMinItems(length int64) KubeBuilderValidation {
	return KubeBuilderValidation{MinItemsValidationName, length}
}

func ValidateMaxItems(length int64) KubeBuilderValidation {
	return KubeBuilderValidation{MaxItemsValidationName, length}
}

func ValidateUniqueItems() KubeBuilderValidation {
	return KubeBuilderValidation{UniqueItemsValidationName, true}
}

func ValidateMaximum(value *big.Rat) KubeBuilderValidation {
	if value.IsInt() {
		return KubeBuilderValidation{MaximumValidationName, value.RatString()}
	} else {
		floatValue, ok := value.Float64()
		if !ok {
			klog.Warningf("inexact maximum: %s ⇒ %v", value.String(), floatValue)
		}

		return KubeBuilderValidation{MaximumValidationName, floatValue}
	}
}

func ValidateMinimum(value *big.Rat) KubeBuilderValidation {
	if value.IsInt() {
		return KubeBuilderValidation{MinimumValidationName, value.RatString()}
	} else {
		floatValue, ok := value.Float64()
		if !ok {
			klog.Warningf("inexact minimum: %s ⇒ %v", value.String(), floatValue)
		}

		return KubeBuilderValidation{MinimumValidationName, floatValue}
	}
}

func ValidateExclusiveMaximum() KubeBuilderValidation {
	return KubeBuilderValidation{ExclusiveMaximumValidationName, true}
}

func ValidateExclusiveMinimum() KubeBuilderValidation {
	return KubeBuilderValidation{ExclusiveMinimumValidationName, true}
}

func ValidateMultipleOf(value *big.Rat) KubeBuilderValidation {
	if value.IsInt() {
		return KubeBuilderValidation{MultipleOfValidationName, value.RatString()}
	} else {
		floatValue, ok := value.Float64()
		if !ok {
			klog.Warningf("inexact multiple-of: %s ⇒ %v", value.String(), floatValue)
		}

		return KubeBuilderValidation{MultipleOfValidationName, floatValue}
	}
}

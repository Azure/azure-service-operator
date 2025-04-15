/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/token"
	"strconv"
	"strings"
	"unicode"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"
	"golang.org/x/exp/slices"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
)

// EnumType represents a set of mutually exclusive predefined options
type EnumType struct {
	// BaseType is the underlying type used to define the values
	baseType *PrimitiveType
	// Options is the set of all unique values
	options []EnumValue

	// emitValidation determines if the enum includes kubebuilder validation when emitted.
	// Statuses may use an enum type but without validation.
	emitValidation bool
}

// EnumType must implement the Type interface correctly
var _ Type = &EnumType{}

// NewEnumType defines a new enumeration including the legal values
func NewEnumType(baseType *PrimitiveType, options ...EnumValue) *EnumType {
	if baseType == nil {
		panic("baseType must be provided")
	}

	slices.SortFunc(
		options,
		func(left EnumValue, right EnumValue) int {
			if baseType == IntType {
				// Compare integers as numbers if we can
				l, lOk := strconv.Atoi(left.Identifier)
				r, rOk := strconv.Atoi(right.Identifier)
				if lOk == nil && rOk == nil {
					return l - r
				}

				if lOk == nil && rOk != nil {
					return -1 // put the int first
				}

				if lOk != nil && rOk == nil {
					return 1 // put the int first
				}
			}

			return strings.Compare(left.Identifier, right.Identifier)
		})

	return &EnumType{
		baseType:       baseType,
		options:        options,
		emitValidation: true,
	}
}

// WithoutValidation returns a copy of this enum, without associated Kubebuilder annotations.
func (enum *EnumType) WithoutValidation() *EnumType {
	result := enum.clone()
	result.emitValidation = false
	return result
}

// WithValidation returns a copy of this enum, with associated Kubebuilder annotations.
func (enum *EnumType) WithValidation() *EnumType {
	result := enum.clone()
	result.emitValidation = true
	return result
}

// AsDeclarations converts the EnumType to a series of Go AST Decls
func (enum *EnumType) AsDeclarations(
	codeGenerationContext *CodeGenerationContext,
	declContext DeclarationContext,
) ([]dst.Decl, error) {
	declareEnum := enum.createBaseDeclaration(
		codeGenerationContext,
		declContext.Name,
		declContext.Description,
		declContext.Validations)

	valuesDeclaration := enum.createValuesDeclaration(declContext)
	mapperDeclaration, err := enum.createMappingDeclaration(declContext.Name, codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating mapping declaration")
	}

	return astbuilder.Declarations(
		declareEnum,
		valuesDeclaration,
		mapperDeclaration,
	), nil
}

func (enum *EnumType) createBaseDeclaration(
	codeGenerationContext *CodeGenerationContext,
	name TypeName,
	description []string,
	validations []KubeBuilderValidation,
) dst.Decl {
	baseTypeExpr, _ := enum.baseType.AsTypeExpr(codeGenerationContext)
	typeSpecification := &dst.TypeSpec{
		Name: dst.NewIdent(name.Name()),
		Type: baseTypeExpr,
	}

	declaration := &dst.GenDecl{
		Decs: dst.GenDeclDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.EmptyLine,
			},
		},
		Tok: token.TYPE,
		Specs: []dst.Spec{
			typeSpecification,
		},
	}

	astbuilder.AddWrappedComments(&declaration.Decs.Start, description)
	AddValidationComments(&declaration.Decs.Start, validations)

	if enum.emitValidation {
		validationComment := GenerateKubebuilderComment(enum.CreateValidation())
		astbuilder.AddComment(&declaration.Decs.Start, validationComment)
	}

	return declaration
}

func (enum *EnumType) createValuesDeclaration(
	declContext DeclarationContext,
) dst.Decl {
	values := make([]dst.Spec, 0, len(enum.options))
	for _, v := range enum.options {
		value := enum.createValueDeclaration(declContext.Name, v)
		values = append(values, value)
	}

	if len(values) == 0 {
		return nil
	}

	return &dst.GenDecl{
		Tok:   token.CONST,
		Specs: values,
	}
}

func (enum *EnumType) createMappingDeclaration(
	name InternalTypeName,
	codeGenerationContext *CodeGenerationContext,
) (dst.Decl, error) {
	if !enum.NeedsMappingConversion(name) {
		// We don't need a mapping conversion map for this enum
		return nil, nil
	}

	baseTypeExpr, err := enum.baseType.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating base type expression")
	}

	nameExpr, err := name.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating name expression")
	}

	literal := astbuilder.NewMapLiteral(
		baseTypeExpr,
		nameExpr)

	for _, v := range enum.options {
		key := astbuilder.TextLiteral(strings.ToLower(v.Value))
		value := dst.NewIdent(GetEnumValueID(name.Name(), v))
		literal.Add(key, value)
	}

	decl := astbuilder.NewVariableAssignment(
		enum.MapperVariableName(name),
		literal.AsExpr(),
	)

	decl.Decorations().Before = dst.EmptyLine
	decl.Decorations().Start = append(
		decl.Decorations().Start,
		fmt.Sprintf("// Mapping from string to %s", name.Name()))

	return decl, nil
}

func (enum *EnumType) createValueDeclaration(name TypeName, value EnumValue) dst.Spec {
	valueSpec := &dst.ValueSpec{
		Names: []*dst.Ident{dst.NewIdent(GetEnumValueID(name.Name(), value))},
		Values: []dst.Expr{
			astbuilder.CallFunc(name.Name(), astbuilder.TextLiteral(value.Value)),
		},
	}

	return valueSpec
}

// AsType implements Type for EnumType
func (enum *EnumType) AsTypeExpr(codeGenerationContext *CodeGenerationContext) (dst.Expr, error) {
	// this should "never" happen as we name all enums; panic if it does
	panic(eris.New("Emitting unnamed enum, something’s awry"))
}

// AsZero renders an expression for the "zero" value of the type,
// based on the underlying type of the enumeration
func (enum *EnumType) AsZero(definitions TypeDefinitionSet, ctx *CodeGenerationContext) dst.Expr {
	return enum.baseType.AsZero(definitions, ctx)
}

// References returns any types the underlying type refers to.
func (enum *EnumType) References() TypeNameSet {
	return enum.baseType.References()
}

// Equals will return true if the supplied type has the same base type and options
func (enum *EnumType) Equals(t Type, overrides EqualityOverrides) bool {
	if enum == t {
		return true // short-circuit
	}

	if e, ok := t.(*EnumType); ok {
		if !enum.baseType.Equals(e.baseType, overrides) {
			return false
		}

		if enum.emitValidation != e.emitValidation {
			return false
		}

		if len(enum.options) != len(e.options) {
			// Different number of properties, not equal
			return false
		}

		for i := range enum.options {
			if !enum.options[i].Equals(&e.options[i]) {
				return false
			}
		}

		// All options match, equal
		return true
	}

	return false
}

// RequiredPackageReferences indicates that Enums never need additional imports
func (enum *EnumType) RequiredPackageReferences() *PackageReferenceSet {
	return NewPackageReferenceSet()
}

// Options returns all the enum options
// A copy of the slice is returned to preserve immutability
func (enum *EnumType) Options() []EnumValue {
	return slices.Clone(enum.options)
}

// CreateValidation creates the validation annotation for this Enum
func (enum *EnumType) CreateValidation() KubeBuilderValidation {
	values := make([]interface{}, 0, len(enum.options))
	for _, opt := range enum.options {
		values = append(values, opt.Value)
	}

	return MakeEnumValidation(values)
}

// BaseType returns the base type of the enum
func (enum *EnumType) BaseType() *PrimitiveType {
	return enum.baseType
}

func (enum *EnumType) clone() *EnumType {
	result := *enum
	result.emitValidation = enum.emitValidation
	result.options = slices.Clone(enum.options)
	result.baseType = enum.baseType

	return &result
}

func GetEnumValueID(name string, value EnumValue) string {
	return name + "_" + value.Identifier
}

// String implements fmt.Stringer
func (enum *EnumType) String() string {
	return fmt.Sprintf("(enum: %s)", enum.baseType.String())
}

// WriteDebugDescription adds a description of the current enum type, including option names, to the
// passed builder.
// builder receives the full description.
// currentPackage is the package that the enum is being written into.
func (enum *EnumType) WriteDebugDescription(builder *strings.Builder, currentPackage InternalPackageReference) {
	if enum == nil {
		builder.WriteString("<nilEnum>")
		return
	}

	builder.WriteString("enum:")
	enum.baseType.WriteDebugDescription(builder, currentPackage)
	if len(enum.options) > 0 {
		builder.WriteString("[")
		for i, v := range enum.options {
			if i > 0 {
				builder.WriteString("|")
			}

			builder.WriteString(v.Identifier)
		}

		builder.WriteString("]")
	}
}

var availableMapperSuffixes = []string{"Values", "Cache", "Mapping", "Mapper"}

// MapperVariableName returns the name of the variable used to map enum values to strings.
// We check the values of the enum to ensure we don't have a name collision.
func (enum *EnumType) MapperVariableName(name InternalTypeName) string {
	used := set.Make[string]()
	for _, v := range enum.options {
		used.Add(v.Identifier)
	}

	for _, suffix := range availableMapperSuffixes {
		if !used.Contains(suffix) {
			name := fmt.Sprintf("%s_%s", name.Name(), suffix)
			// Force the name to an internal one
			runes := []rune(name)
			runes[0] = unicode.ToLower(runes[0])
			return string(runes)
		}
	}

	panic("No available suffix for enum mapper variable name")
}

// NeedsMappingConversion returns true if the enum needs a mapping conversion.
// A mapping conversion is needed if the base type is a string and the enum is
// *not* APIVersion (that enum isn't actually used on Spec or Status types).
func (enum *EnumType) NeedsMappingConversion(name InternalTypeName) bool {
	return enum.baseType == StringType &&
		name.Name() != "APIVersion"
}

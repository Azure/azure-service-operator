/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "fmt"

// TypeVisitorBuilder provides a flexible way to create a TypeVisitor. Fields should be initialized
// with funcs matching one of the following forms:
//
// func(this *TypeVisitor, it <sometype>, ctx interface{}) (Type, error)
// func(it <sometype>) (Type, error)
// func(it <sometype>) Type
//
// o  Must always return Type, and optionally an error
// o  <sometype> must match the type for the field being initialized
//
// Some examples:
//
// VisitTypeName = func(it TypeName) Type      // Works
// VisitTypeName = func(it *ObjectType) Type   // Fails - parameter is not a TypeName
// VisitTypeName = func(it TypeName) TypeName  // Fails - return type is not Type
//
type TypeVisitorBuilder struct {
	VisitTypeName      interface{}
	VisitOneOfType     interface{}
	VisitAllOfType     interface{}
	VisitArrayType     interface{}
	VisitPrimitive     interface{}
	VisitObjectType    interface{}
	VisitMapType       interface{}
	VisitOptionalType  interface{}
	VisitEnumType      interface{}
	VisitResourceType  interface{}
	VisitFlaggedType   interface{}
	VisitValidatedType interface{}
	VisitErroredType   interface{}
}

func (b TypeVisitorBuilder) Build() TypeVisitor {
	return TypeVisitor{
		visitTypeName:      b.buildVisitTypeName(),
		visitOneOfType:     b.buildVisitOneOfType(),
		visitAllOfType:     b.buildVisitAllOfType(),
		visitArrayType:     b.buildVisitArrayType(),
		visitPrimitive:     b.buildVisitPrimitive(),
		visitObjectType:    b.buildVisitObjectType(),
		visitMapType:       b.buildVisitMapType(),
		visitOptionalType:  b.buildVisitOptionalType(),
		visitEnumType:      b.buildVisitEnumType(),
		visitResourceType:  b.buildVisitResourceType(),
		visitFlaggedType:   b.buildVisitFlaggedType(),
		visitValidatedType: b.buildVisitValidatedType(),
		visitErroredType:   b.buildVisitErroredType(),
	}
}

// buildVisitTypeName returns a function to use in the TypeVisitor
// If the field VisitTypeName is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder) buildVisitTypeName() func(*TypeVisitor, TypeName, interface{}) (Type, error) {
	if b.VisitTypeName == nil {
		return IdentityVisitOfTypeName
	}

	switch v := b.VisitTypeName.(type) {
	case func(*TypeVisitor, TypeName, interface{}) (Type, error):
		return v
	case func(TypeName) (Type, error):
		return func(_ *TypeVisitor, it TypeName, _ interface{}) (Type, error) {
			return v(it)
		}
	case func(TypeName) Type:
		return func(_ *TypeVisitor, it TypeName, _ interface{}) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected TypeName func %#v", b.VisitTypeName))
}

// buildVisitOneOfType returns a function to use in the TypeVisitor
// If the field VisitOneOfType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder) buildVisitOneOfType() func(*TypeVisitor, *OneOfType, interface{}) (Type, error) {
	if b.VisitOneOfType == nil {
		return IdentityVisitOfOneOfType
	}

	switch v := b.VisitOneOfType.(type) {
	case func(*TypeVisitor, *OneOfType, interface{}) (Type, error):
		return v
	case func(*OneOfType) (Type, error):
		return func(_ *TypeVisitor, it *OneOfType, _ interface{}) (Type, error) {
			return v(it)
		}
	case func(*OneOfType) Type:
		return func(_ *TypeVisitor, it *OneOfType, _ interface{}) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected OneOfType func %#v", b.VisitOneOfType))
}

// buildVisitAllOfType returns a function to use in the TypeVisitor
// If the field VisitAllOfType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder) buildVisitAllOfType() func(*TypeVisitor, *AllOfType, interface{}) (Type, error) {
	if b.VisitAllOfType == nil {
		return IdentityVisitOfAllOfType
	}

	switch v := b.VisitAllOfType.(type) {
	case func(*TypeVisitor, *AllOfType, interface{}) (Type, error):
		return v
	case func(*AllOfType) (Type, error):
		return func(_ *TypeVisitor, it *AllOfType, _ interface{}) (Type, error) {
			return v(it)
		}
	case func(*AllOfType) Type:
		return func(_ *TypeVisitor, it *AllOfType, _ interface{}) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected AllOfType func %#v", b.VisitAllOfType))
}

// buildVisitArrayType returns a function to use in the TypeVisitor
// If the field VisitArrayType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder) buildVisitArrayType() func(*TypeVisitor, *ArrayType, interface{}) (Type, error) {
	if b.VisitArrayType == nil {
		return IdentityVisitOfArrayType
	}

	switch v := b.VisitArrayType.(type) {
	case func(*TypeVisitor, *ArrayType, interface{}) (Type, error):
		return v
	case func(*ArrayType) (Type, error):
		return func(_ *TypeVisitor, it *ArrayType, _ interface{}) (Type, error) {
			return v(it)
		}
	case func(*ArrayType) Type:
		return func(_ *TypeVisitor, it *ArrayType, _ interface{}) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected ArrayType func %#v", b.VisitArrayType))
}

// buildVisitPrimitive returns a function to use in the TypeVisitor
// If the field VisitPrimitive is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder) buildVisitPrimitive() func(*TypeVisitor, *PrimitiveType, interface{}) (Type, error) {
	if b.VisitPrimitive == nil {
		return IdentityVisitOfPrimitiveType
	}

	switch v := b.VisitPrimitive.(type) {
	case func(*TypeVisitor, *PrimitiveType, interface{}) (Type, error):
		return v
	case func(*PrimitiveType) (Type, error):
		return func(_ *TypeVisitor, it *PrimitiveType, _ interface{}) (Type, error) {
			return v(it)
		}
	case func(*PrimitiveType) Type:
		return func(_ *TypeVisitor, it *PrimitiveType, _ interface{}) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected Primitive func %#v", b.VisitPrimitive))
}

// buildVisitObjectType returns a function to use in the TypeVisitor
// If the field VisitObjectType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder) buildVisitObjectType() func(*TypeVisitor, *ObjectType, interface{}) (Type, error) {
	if b.VisitObjectType == nil {
		return IdentityVisitOfObjectType
	}

	switch v := b.VisitObjectType.(type) {
	case func(*TypeVisitor, *ObjectType, interface{}) (Type, error):
		return v
	case func(*ObjectType) (Type, error):
		return func(_ *TypeVisitor, it *ObjectType, _ interface{}) (Type, error) {
			return v(it)
		}
	case func(*ObjectType) Type:
		return func(_ *TypeVisitor, it *ObjectType, _ interface{}) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected ObjectType func %#v", b.VisitObjectType))
}

// buildVisitMapType returns a function to use in the TypeVisitor
// If the field VisitMapType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder) buildVisitMapType() func(*TypeVisitor, *MapType, interface{}) (Type, error) {
	if b.VisitMapType == nil {
		return IdentityVisitOfMapType
	}

	switch v := b.VisitMapType.(type) {
	case func(*TypeVisitor, *MapType, interface{}) (Type, error):
		return v
	case func(*MapType) (Type, error):
		return func(_ *TypeVisitor, it *MapType, _ interface{}) (Type, error) {
			return v(it)
		}
	case func(*MapType) Type:
		return func(_ *TypeVisitor, it *MapType, _ interface{}) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected MapType func %#v", b.VisitMapType))
}

// buildVisitOptionalType returns a function to use in the TypeVisitor
// If the field VisitOptionalType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder) buildVisitOptionalType() func(*TypeVisitor, *OptionalType, interface{}) (Type, error) {
	if b.VisitOptionalType == nil {
		return IdentityVisitOfOptionalType
	}

	switch v := b.VisitOptionalType.(type) {
	case func(*TypeVisitor, *OptionalType, interface{}) (Type, error):
		return v
	case func(*OptionalType) (Type, error):
		return func(_ *TypeVisitor, it *OptionalType, _ interface{}) (Type, error) {
			return v(it)
		}
	case func(*OptionalType) Type:
		return func(_ *TypeVisitor, it *OptionalType, _ interface{}) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected OptionalType func %#v", b.VisitOptionalType))
}

// buildVisitEnumType returns a function to use in the TypeVisitor
// If the field VisitEnumType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder) buildVisitEnumType() func(*TypeVisitor, *EnumType, interface{}) (Type, error) {
	if b.VisitEnumType == nil {
		return IdentityVisitOfEnumType
	}

	switch v := b.VisitEnumType.(type) {
	case func(*TypeVisitor, *EnumType, interface{}) (Type, error):
		return v
	case func(*EnumType) (Type, error):
		return func(_ *TypeVisitor, it *EnumType, _ interface{}) (Type, error) {
			return v(it)
		}
	case func(*EnumType) Type:
		return func(_ *TypeVisitor, it *EnumType, _ interface{}) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected EnumType func %#v", b.VisitEnumType))
}

// buildVisitResourceType returns a function to use in the TypeVisitor
// If the field VisitResourceType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder) buildVisitResourceType() func(*TypeVisitor, *ResourceType, interface{}) (Type, error) {
	if b.VisitResourceType == nil {
		return IdentityVisitOfResourceType
	}

	switch v := b.VisitResourceType.(type) {
	case func(*TypeVisitor, *ResourceType, interface{}) (Type, error):
		return v
	case func(*ResourceType) (Type, error):
		return func(_ *TypeVisitor, it *ResourceType, _ interface{}) (Type, error) {
			return v(it)
		}
	case func(*ResourceType) Type:
		return func(_ *TypeVisitor, it *ResourceType, _ interface{}) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected ResourceType func %#v", b.VisitResourceType))
}

// buildVisitFlaggedType returns a function to use in the TypeVisitor
// If the field VisitFlaggedType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder) buildVisitFlaggedType() func(*TypeVisitor, *FlaggedType, interface{}) (Type, error) {
	if b.VisitFlaggedType == nil {
		return IdentityVisitOfFlaggedType
	}

	switch v := b.VisitFlaggedType.(type) {
	case func(*TypeVisitor, *FlaggedType, interface{}) (Type, error):
		return v
	case func(*FlaggedType) (Type, error):
		return func(_ *TypeVisitor, it *FlaggedType, _ interface{}) (Type, error) {
			return v(it)
		}
	case func(*FlaggedType) Type:
		return func(_ *TypeVisitor, it *FlaggedType, _ interface{}) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected FlaggedType func %#v", b.VisitFlaggedType))
}

// buildVisitValidatedType returns a function to use in the TypeVisitor
// If the field VisitValidatedType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder) buildVisitValidatedType() func(*TypeVisitor, *ValidatedType, interface{}) (Type, error) {
	if b.VisitValidatedType == nil {
		return IdentityVisitOfValidatedType
	}

	switch v := b.VisitValidatedType.(type) {
	case func(*TypeVisitor, *ValidatedType, interface{}) (Type, error):
		return v
	case func(*ValidatedType) (Type, error):
		return func(_ *TypeVisitor, it *ValidatedType, _ interface{}) (Type, error) {
			return v(it)
		}
	case func(*ValidatedType) Type:
		return func(_ *TypeVisitor, it *ValidatedType, _ interface{}) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected ValidatedType func %#v", b.VisitValidatedType))
}

// buildVisitErroredType returns a function to use in the TypeVisitor
// If the field VisitErroredType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder) buildVisitErroredType() func(*TypeVisitor, *ErroredType, interface{}) (Type, error) {
	if b.VisitErroredType == nil {
		return IdentityVisitOfErroredType
	}

	switch v := b.VisitErroredType.(type) {
	case func(*TypeVisitor, *ErroredType, interface{}) (Type, error):
		return v
	case func(*ErroredType) (Type, error):
		return func(_ *TypeVisitor, it *ErroredType, _ interface{}) (Type, error) {
			return v(it)
		}
	case func(*ErroredType) Type:
		return func(_ *TypeVisitor, it *ErroredType, _ interface{}) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected ErroredType func %#v", b.VisitErroredType))
}

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "fmt"

// TypeVisitorBuilder provides a flexible way to create a TypeVisitor.
// C is the type to use for the context parameter in each visit method.
// Fields should be initialized with funcs matching one of the following forms:
//
// func(this *TypeVisitor, it <sometype>, ctx any) (Type, error)
// func(it <sometype>) (Type, error)
// func(it <sometype>) Type
//
// o  Must always return Type, and optionally an error
// o  <sometype> must match the type for the field being initialized
//
// Some examples:
//
// VisitInternalTypeName = func(it InternalTypeName) Type                                   // Works
// VisitInternalTypeName = func(this TypeVisitor, it InternalTypeName, ctx C) (Type, error) // Works
// VisitInternalTypeName = func(it *ObjectType) Type                                        // Fails - parameter is not a TypeName
// VisitInternalTypeName = func(it TypeName) TypeName                                       // Fails - return type is not Type
//
// VisitObjectType = func(it *ObjectType) Type                                      // Works
// VisitObjectType = func(this TypeVisitor, it *ObjectType, ctx C) (Type, error)    // Works
// VisitObjectType = func(it TypeName) Type                                         // Fails - parameter is not an *ObjectType
// VisitObjectType = func(this TypeVisitor, it TypeName, ctx C) (ObjectType, error) // Fails -return is not Type
type TypeVisitorBuilder[C any] struct {
	VisitInternalTypeName any
	VisitExternalTypeName any
	VisitOneOfType        any
	VisitAllOfType        any
	VisitArrayType        any
	VisitPrimitive        any
	VisitObjectType       any
	VisitMapType          any
	VisitOptionalType     any
	VisitEnumType         any
	VisitResourceType     any
	VisitFlaggedType      any
	VisitValidatedType    any
	VisitErroredType      any
	VisitInterfaceType    any
}

func (b TypeVisitorBuilder[C]) Build() TypeVisitor[C] {
	return TypeVisitor[C]{
		visitInternalTypeName: b.buildVisitInternalTypeName(),
		visitExternalTypeName: b.buildVisitExternalTypeName(),
		visitOneOfType:        b.buildVisitOneOfType(),
		visitAllOfType:        b.buildVisitAllOfType(),
		visitArrayType:        b.buildVisitArrayType(),
		visitPrimitive:        b.buildVisitPrimitive(),
		visitObjectType:       b.buildVisitObjectType(),
		visitMapType:          b.buildVisitMapType(),
		visitOptionalType:     b.buildVisitOptionalType(),
		visitEnumType:         b.buildVisitEnumType(),
		visitResourceType:     b.buildVisitResourceType(),
		visitFlaggedType:      b.buildVisitFlaggedType(),
		visitValidatedType:    b.buildVisitValidatedType(),
		visitErroredType:      b.buildVisitErroredType(),
		visitInterfaceType:    b.buildVisitInterfaceType(),
	}
}

// buildVisitInternalTypeName returns a function to use in the TypeVisitor
// If the field VisitInternalTypeName is nil, we return nil, unlike other build functions.
// We do this to avoid the boxing inherent with passing an InternalTypeName if there's nothing to be done.
// Otherwise, we attempt to convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder[C]) buildVisitInternalTypeName() func(*TypeVisitor[C], InternalTypeName, C) (Type, error) {
	if b.VisitInternalTypeName == nil {
		return nil
	}

	switch v := b.VisitInternalTypeName.(type) {
	case func(*TypeVisitor[C], InternalTypeName, C) (Type, error):
		return v
	case func(InternalTypeName) (Type, error):
		return func(_ *TypeVisitor[C], it InternalTypeName, _ C) (Type, error) {
			return v(it)
		}
	case func(InternalTypeName) Type:
		return func(_ *TypeVisitor[C], it InternalTypeName, _ C) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected InternalTypeName func %#v", b.VisitInternalTypeName))
}

// buildVisitExternalTypeName returns a function to use in the TypeVisitor
// If the field VisitExternalTypeName is nil, we return nil, unlike other build functions.
// We do this to avoid the boxing inherent with passing an ExternalTypeName if there's nothing to be done.
// Otherwise, we attempt to convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder[C]) buildVisitExternalTypeName() func(*TypeVisitor[C], ExternalTypeName, C) (Type, error) {
	if b.VisitExternalTypeName == nil {
		return nil
	}

	switch v := b.VisitExternalTypeName.(type) {
	case func(*TypeVisitor[C], ExternalTypeName, C) (Type, error):
		return v
	case func(ExternalTypeName) (Type, error):
		return func(_ *TypeVisitor[C], it ExternalTypeName, _ C) (Type, error) {
			return v(it)
		}
	case func(ExternalTypeName) Type:
		return func(_ *TypeVisitor[C], it ExternalTypeName, _ C) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected ExternalTypeName func %#v", b.VisitExternalTypeName))
}

// buildVisitOneOfType returns a function to use in the TypeVisitor
// If the field VisitOneOfType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder[C]) buildVisitOneOfType() func(*TypeVisitor[C], *OneOfType, C) (Type, error) {
	if b.VisitOneOfType == nil {
		return IdentityVisitOfOneOfType[C]
	}

	switch v := b.VisitOneOfType.(type) {
	case func(*TypeVisitor[C], *OneOfType, C) (Type, error):
		return v
	case func(*OneOfType) (Type, error):
		return func(_ *TypeVisitor[C], it *OneOfType, _ C) (Type, error) {
			return v(it)
		}
	case func(*OneOfType) Type:
		return func(_ *TypeVisitor[C], it *OneOfType, _ C) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected OneOfType func %#v", b.VisitOneOfType))
}

// buildVisitAllOfType returns a function to use in the TypeVisitor
// If the field VisitAllOfType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder[C]) buildVisitAllOfType() func(*TypeVisitor[C], *AllOfType, C) (Type, error) {
	if b.VisitAllOfType == nil {
		return IdentityVisitOfAllOfType[C]
	}

	switch v := b.VisitAllOfType.(type) {
	case func(*TypeVisitor[C], *AllOfType, C) (Type, error):
		return v
	case func(*AllOfType) (Type, error):
		return func(_ *TypeVisitor[C], it *AllOfType, _ C) (Type, error) {
			return v(it)
		}
	case func(*AllOfType) Type:
		return func(_ *TypeVisitor[C], it *AllOfType, _ C) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected AllOfType func %#v", b.VisitAllOfType))
}

// buildVisitArrayType returns a function to use in the TypeVisitor
// If the field VisitArrayType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder[C]) buildVisitArrayType() func(*TypeVisitor[C], *ArrayType, C) (Type, error) {
	if b.VisitArrayType == nil {
		return IdentityVisitOfArrayType[C]
	}

	switch v := b.VisitArrayType.(type) {
	case func(*TypeVisitor[C], *ArrayType, C) (Type, error):
		return v
	case func(*ArrayType) (Type, error):
		return func(_ *TypeVisitor[C], it *ArrayType, _ C) (Type, error) {
			return v(it)
		}
	case func(*ArrayType) Type:
		return func(_ *TypeVisitor[C], it *ArrayType, _ C) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected ArrayType func %#v", b.VisitArrayType))
}

// buildVisitPrimitive returns a function to use in the TypeVisitor
// If the field VisitPrimitive is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder[C]) buildVisitPrimitive() func(*TypeVisitor[C], *PrimitiveType, C) (Type, error) {
	if b.VisitPrimitive == nil {
		return IdentityVisitOfPrimitiveType[C]
	}

	switch v := b.VisitPrimitive.(type) {
	case func(*TypeVisitor[C], *PrimitiveType, C) (Type, error):
		return v
	case func(*PrimitiveType) (Type, error):
		return func(_ *TypeVisitor[C], it *PrimitiveType, _ C) (Type, error) {
			return v(it)
		}
	case func(*PrimitiveType) Type:
		return func(_ *TypeVisitor[C], it *PrimitiveType, _ C) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected Primitive func %#v", b.VisitPrimitive))
}

// buildVisitObjectType returns a function to use in the TypeVisitor
// If the field VisitObjectType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder[C]) buildVisitObjectType() func(*TypeVisitor[C], *ObjectType, C) (Type, error) {
	if b.VisitObjectType == nil {
		return IdentityVisitOfObjectType[C]
	}

	switch v := b.VisitObjectType.(type) {
	case func(*TypeVisitor[C], *ObjectType, C) (Type, error):
		return v
	case func(*ObjectType) (Type, error):
		return func(_ *TypeVisitor[C], it *ObjectType, _ C) (Type, error) {
			return v(it)
		}
	case func(*ObjectType) Type:
		return func(_ *TypeVisitor[C], it *ObjectType, _ C) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected ObjectType func %#v", b.VisitObjectType))
}

// buildVisitMapType returns a function to use in the TypeVisitor
// If the field VisitMapType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder[C]) buildVisitMapType() func(*TypeVisitor[C], *MapType, C) (Type, error) {
	if b.VisitMapType == nil {
		return IdentityVisitOfMapType[C]
	}

	switch v := b.VisitMapType.(type) {
	case func(*TypeVisitor[C], *MapType, C) (Type, error):
		return v
	case func(*MapType) (Type, error):
		return func(_ *TypeVisitor[C], it *MapType, _ C) (Type, error) {
			return v(it)
		}
	case func(*MapType) Type:
		return func(_ *TypeVisitor[C], it *MapType, _ C) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected MapType func %#v", b.VisitMapType))
}

// buildVisitOptionalType returns a function to use in the TypeVisitor
// If the field VisitOptionalType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder[C]) buildVisitOptionalType() func(*TypeVisitor[C], *OptionalType, C) (Type, error) {
	if b.VisitOptionalType == nil {
		return IdentityVisitOfOptionalType[C]
	}

	switch v := b.VisitOptionalType.(type) {
	case func(*TypeVisitor[C], *OptionalType, C) (Type, error):
		return v
	case func(*OptionalType) (Type, error):
		return func(_ *TypeVisitor[C], it *OptionalType, _ C) (Type, error) {
			return v(it)
		}
	case func(*OptionalType) Type:
		return func(_ *TypeVisitor[C], it *OptionalType, _ C) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected OptionalType func %#v", b.VisitOptionalType))
}

// buildVisitEnumType returns a function to use in the TypeVisitor
// If the field VisitEnumType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder[C]) buildVisitEnumType() func(*TypeVisitor[C], *EnumType, C) (Type, error) {
	if b.VisitEnumType == nil {
		return IdentityVisitOfEnumType[C]
	}

	switch v := b.VisitEnumType.(type) {
	case func(*TypeVisitor[C], *EnumType, C) (Type, error):
		return v
	case func(*EnumType) (Type, error):
		return func(_ *TypeVisitor[C], it *EnumType, _ C) (Type, error) {
			return v(it)
		}
	case func(*EnumType) Type:
		return func(_ *TypeVisitor[C], it *EnumType, _ C) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected EnumType func %#v", b.VisitEnumType))
}

// buildVisitResourceType returns a function to use in the TypeVisitor
// If the field VisitResourceType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder[C]) buildVisitResourceType() func(*TypeVisitor[C], *ResourceType, C) (Type, error) {
	if b.VisitResourceType == nil {
		return IdentityVisitOfResourceType[C]
	}

	switch v := b.VisitResourceType.(type) {
	case func(*TypeVisitor[C], *ResourceType, C) (Type, error):
		return v
	case func(*ResourceType) (Type, error):
		return func(_ *TypeVisitor[C], it *ResourceType, _ C) (Type, error) {
			return v(it)
		}
	case func(*ResourceType) Type:
		return func(_ *TypeVisitor[C], it *ResourceType, _ C) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected ResourceType func %#v", b.VisitResourceType))
}

// buildVisitFlaggedType returns a function to use in the TypeVisitor
// If the field VisitFlaggedType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder[C]) buildVisitFlaggedType() func(*TypeVisitor[C], *FlaggedType, C) (Type, error) {
	if b.VisitFlaggedType == nil {
		return IdentityVisitOfFlaggedType[C]
	}

	switch v := b.VisitFlaggedType.(type) {
	case func(*TypeVisitor[C], *FlaggedType, C) (Type, error):
		return v
	case func(*FlaggedType) (Type, error):
		return func(_ *TypeVisitor[C], it *FlaggedType, _ C) (Type, error) {
			return v(it)
		}
	case func(*FlaggedType) Type:
		return func(_ *TypeVisitor[C], it *FlaggedType, _ C) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected FlaggedType func %#v", b.VisitFlaggedType))
}

// buildVisitValidatedType returns a function to use in the TypeVisitor
// If the field VisitValidatedType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder[C]) buildVisitValidatedType() func(*TypeVisitor[C], *ValidatedType, C) (Type, error) {
	if b.VisitValidatedType == nil {
		return IdentityVisitOfValidatedType[C]
	}

	switch v := b.VisitValidatedType.(type) {
	case func(*TypeVisitor[C], *ValidatedType, C) (Type, error):
		return v
	case func(*ValidatedType) (Type, error):
		return func(_ *TypeVisitor[C], it *ValidatedType, _ C) (Type, error) {
			return v(it)
		}
	case func(*ValidatedType) Type:
		return func(_ *TypeVisitor[C], it *ValidatedType, _ C) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected ValidatedType func %#v", b.VisitValidatedType))
}

// buildVisitErroredType returns a function to use in the TypeVisitor
// If the field VisitErroredType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder[C]) buildVisitErroredType() func(*TypeVisitor[C], *ErroredType, C) (Type, error) {
	if b.VisitErroredType == nil {
		return IdentityVisitOfErroredType[C]
	}

	switch v := b.VisitErroredType.(type) {
	case func(*TypeVisitor[C], *ErroredType, C) (Type, error):
		return v
	case func(*ErroredType) (Type, error):
		return func(_ *TypeVisitor[C], it *ErroredType, _ C) (Type, error) {
			return v(it)
		}
	case func(*ErroredType) Type:
		return func(_ *TypeVisitor[C], it *ErroredType, _ C) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected ErroredType func %#v", b.VisitErroredType))
}

// buildVisitInterfaceType returns a function to use in the TypeVisitor
// If the field VisitInterfaceType is nil, we return an identity visitor. Otherwise we attempt to
// convert the func found in the field, triggering a panic if no suitable func is found.
func (b *TypeVisitorBuilder[C]) buildVisitInterfaceType() func(*TypeVisitor[C], *InterfaceType, C) (Type, error) {
	if b.VisitInterfaceType == nil {
		return IdentityVisitOfInterfaceType[C]
	}

	switch v := b.VisitInterfaceType.(type) {
	case func(*TypeVisitor[C], *InterfaceType, C) (Type, error):
		return v
	case func(*InterfaceType) (Type, error):
		return func(_ *TypeVisitor[C], it *InterfaceType, _ C) (Type, error) {
			return v(it)
		}
	case func(*InterfaceType) Type:
		return func(_ *TypeVisitor[C], it *InterfaceType, _ C) (Type, error) {
			return v(it), nil
		}
	}

	panic(fmt.Sprintf("unexpected InterfaceType func %#v", b.VisitInterfaceType))
}

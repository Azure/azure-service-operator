/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

const (
	group   = "grp"
	version = "v1"
)

var (
	rootTypeName  = MakeInternalTypeName(makeTestLocalPackageReference(group, version), "Root")
	leftTypeName  = MakeInternalTypeName(makeTestLocalPackageReference(group, version), "Left")
	rightTypeName = MakeInternalTypeName(makeTestLocalPackageReference(group, version), "Right")
)

func makeSimpleTestTypeGraph() TypeDefinitionSet {
	result := make(TypeDefinitionSet)

	leftChildType := NewObjectType().WithProperty(
		NewPropertyDefinition("SimpleString", "simpleString", StringType))
	leftChildDef := MakeTypeDefinition(
		leftTypeName,
		leftChildType)
	rightChildType := NewObjectType().WithProperty(
		NewPropertyDefinition("SimpleInt", "simpleInt", IntType))
	rightChildDef := MakeTypeDefinition(
		rightTypeName,
		rightChildType)

	rootType := NewObjectType().WithProperty(
		NewPropertyDefinition("Left", "left", leftChildDef.Name())).WithProperty(
		NewPropertyDefinition("Right", "right", rightChildDef.Name()))
	rootDef := MakeTypeDefinition(
		rootTypeName,
		rootType)

	result.Add(leftChildDef)
	result.Add(rightChildDef)
	result.Add(rootDef)

	return result
}

func makeDuplicateReferencesTypeGraph() TypeDefinitionSet {
	result := make(TypeDefinitionSet)

	childType := NewObjectType().WithProperty(
		NewPropertyDefinition("SimpleString", "simpleString", StringType))
	childDef := MakeTypeDefinition(
		leftTypeName,
		childType)

	rootType := NewObjectType().WithProperty(
		NewPropertyDefinition("Left", "left", childDef.Name())).WithProperty(
		NewPropertyDefinition("Right", "right", childDef.Name()))
	rootDef := MakeTypeDefinition(
		rootTypeName,
		rootType)

	result.Add(childDef)
	result.Add(rootDef)

	return result
}

func makeCycleTypeGraph() TypeDefinitionSet {
	result := make(TypeDefinitionSet)

	leftChildType := NewObjectType().WithProperty(
		NewPropertyDefinition("Root", "root", rootTypeName))
	leftChildDef := MakeTypeDefinition(
		leftTypeName,
		leftChildType)
	rightChildType := NewObjectType().WithProperty(
		NewPropertyDefinition("SimpleInt", "simpleInt", IntType))
	rightChildDef := MakeTypeDefinition(
		rightTypeName,
		rightChildType)

	rootType := NewObjectType().WithProperty(
		NewPropertyDefinition("Left", "left", leftChildDef.Name())).WithProperty(
		NewPropertyDefinition("Right", "right", rightChildDef.Name()))
	rootDef := MakeTypeDefinition(
		rootTypeName,
		rootType)

	result.Add(leftChildDef)
	result.Add(rightChildDef)
	result.Add(rootDef)

	return result
}

func TestTypeWalker_IdentityWalkReturnsIdenticalTypes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	types := makeSimpleTestTypeGraph()
	visitor := TypeVisitorBuilder[any]{}.Build()
	walker := NewTypeWalker(types, visitor)

	var walked []string

	walker.AfterVisit = func(original TypeDefinition, updated TypeDefinition, ctx any) (TypeDefinition, error) {
		walked = append(walked, original.Name().Name())
		return IdentityAfterVisit(original, updated, ctx)
	}

	rootDef := types[rootTypeName]
	updatedTypes, err := walker.Walk(rootDef)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(walked)).To(Equal(3))
	g.Expect(len(updatedTypes)).To(Equal(3))

	g.Expect(walked).To(ContainElement("Left"))
	g.Expect(walked).To(ContainElement("Right"))
	g.Expect(walked).To(ContainElement("Root"))

	for _, updated := range updatedTypes {
		g.Expect(TypeEquals(updated.Type(), types[updated.Name()].Type())).To(BeTrue())
	}
}

func TestTypeWalker_DuplicateTypesAreWalkedOnceEach_ReturnedOnce(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	types := makeDuplicateReferencesTypeGraph()
	visitor := TypeVisitorBuilder[any]{}.Build()
	walker := NewTypeWalker(types, visitor)

	var walked []TypeDefinition

	walker.AfterVisit = func(original TypeDefinition, updated TypeDefinition, ctx any) (TypeDefinition, error) {
		walked = append(walked, original)
		return IdentityAfterVisit(original, updated, ctx)
	}

	rootDef := types[rootTypeName]
	updatedTypes, err := walker.Walk(rootDef)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(walked)).To(Equal(3))
	g.Expect(len(updatedTypes)).To(Equal(2))

	g.Expect(walked[0].Name().Name()).To(Equal("Left"))
	g.Expect(walked[1].Name().Name()).To(Equal("Left"))
	g.Expect(walked[2].Name().Name()).To(Equal("Root"))

	for _, updated := range updatedTypes {
		g.Expect(TypeEquals(updated.Type(), types[updated.Name()].Type())).To(BeTrue())
	}
}

func TestTypeWalker_CyclesAllowed_AreNotWalked(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	types := makeCycleTypeGraph()
	visitor := TypeVisitorBuilder[any]{}.Build()
	walker := NewTypeWalker(types, visitor)

	var walked []string

	walker.AfterVisit = func(original TypeDefinition, updated TypeDefinition, ctx any) (TypeDefinition, error) {
		walked = append(walked, original.Name().Name())
		return IdentityAfterVisit(original, updated, ctx)
	}

	rootDef := types[rootTypeName]
	updatedTypes, err := walker.Walk(rootDef)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(walked)).To(Equal(3))
	g.Expect(len(updatedTypes)).To(Equal(3))

	g.Expect(walked).To(ContainElement("Left"))
	g.Expect(walked).To(ContainElement("Right"))
	g.Expect(walked).To(ContainElement("Root"))

	for _, updated := range updatedTypes {
		g.Expect(TypeEquals(updated.Type(), types[updated.Name()].Type())).To(BeTrue())
	}
}

func TestTypeWalker_CanPruneCycles(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	types := makeCycleTypeGraph()
	visitor := TypeVisitorBuilder[any]{}.Build()
	walker := NewTypeWalker(types, visitor)

	var walked []string

	walker.AfterVisit = func(original TypeDefinition, updated TypeDefinition, ctx any) (TypeDefinition, error) {
		walked = append(walked, original.Name().Name())
		return IdentityAfterVisit(original, updated, ctx)
	}

	walker.ShouldRemoveCycle = func(original TypeDefinition, ctx any) (bool, error) {
		// Prune all cycles
		return true, nil
	}

	rootDef := types[rootTypeName]
	updatedTypes, err := walker.Walk(rootDef)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(walked)).To(Equal(3))
	g.Expect(len(updatedTypes)).To(Equal(3))

	g.Expect(walked).To(ContainElement("Left"))
	g.Expect(walked).To(ContainElement("Right"))
	g.Expect(walked).To(ContainElement("Root"))

	for _, updated := range updatedTypes {
		// Expect only left to be updated - that is only type with cycle property
		if TypeEquals(updated.Name(), leftTypeName) {
			g.Expect(TypeEquals(updated.Type(), types[updated.Name()].Type())).To(BeFalse())
			obj, ok := updated.Type().(*ObjectType)
			g.Expect(ok).To(BeTrue())
			g.Expect(obj.Properties().Len()).To(Equal(0))
		} else {
			g.Expect(TypeEquals(updated.Type(), types[updated.Name()].Type())).To(BeTrue())
		}
	}
}

func TestTypeWalker_ContextPropagated(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	types := makeSimpleTestTypeGraph()
	visitor := TypeVisitorBuilder[any]{}.Build()
	walker := NewTypeWalker(types, visitor)

	walked := make(map[TypeName]int)

	walker.AfterVisit = func(original TypeDefinition, updated TypeDefinition, ctx any) (TypeDefinition, error) {
		typedCtx := ctx.(int)
		walked[updated.Name()] = typedCtx
		return IdentityAfterVisit(original, updated, ctx)
	}
	walker.MakeContext = func(_ TypeName, ctx any) (any, error) {
		if ctx == nil {
			return 0, nil
		}
		typedCtx := ctx.(int)
		typedCtx += 1

		return typedCtx, nil
	}

	rootDef := types[rootTypeName]
	updatedTypes, err := walker.Walk(rootDef)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(walked)).To(Equal(3))
	g.Expect(len(updatedTypes)).To(Equal(3))

	g.Expect(walked[rootTypeName]).To(Equal(0))
	g.Expect(walked[leftTypeName]).To(Equal(1))
	g.Expect(walked[rightTypeName]).To(Equal(1))

	for _, updated := range updatedTypes {
		g.Expect(TypeEquals(updated.Type(), types[updated.Name()].Type())).To(BeTrue())
	}
}

func TestTypeWalker_VisitorApplied(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	types := makeSimpleTestTypeGraph()
	visitor := TypeVisitorBuilder[any]{
		VisitObjectType: func(this *TypeVisitor[any], it *ObjectType, ctx any) (Type, error) {
			_ = ctx.(int) // Ensure context is the right shape

			// Find any properties of type string and remove them
			it.Properties().ForEach(func(prop *PropertyDefinition) {
				if prop.PropertyType() == StringType {
					it = it.WithoutProperty(prop.PropertyName())
				}
			})

			return IdentityVisitOfObjectType(this, it, ctx)
		},
	}.Build()

	walker := NewTypeWalker(types, visitor)
	walker.MakeContext = func(_ TypeName, _ any) (any, error) {
		return 0, nil
	}

	rootDef := types[rootTypeName]
	updatedTypes, err := walker.Walk(rootDef)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(updatedTypes)).To(Equal(3))

	for _, updated := range updatedTypes {
		// Expect only left to be updated - that is only type with string property
		if TypeEquals(updated.Name(), leftTypeName) {
			g.Expect(TypeEquals(updated.Type(), types[updated.Name()].Type())).To(BeFalse())
			obj, ok := updated.Type().(*ObjectType)
			g.Expect(ok).To(BeTrue())
			g.Expect(obj.Properties().Len()).To(Equal(0))
		} else {
			g.Expect(TypeEquals(updated.Type(), types[updated.Name()].Type())).To(BeTrue())
		}
	}
}

func TestTypeWalker_CanChangeNameInOnlyCertainPlaces(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	types := makeDuplicateReferencesTypeGraph()
	visitor := TypeVisitorBuilder[any]{}.Build()
	walker := NewTypeWalker(types, visitor)

	left2TypeName := MakeInternalTypeName(leftTypeName.PackageReference(), "Left2")

	changed := false
	walker.AfterVisit = func(original TypeDefinition, updated TypeDefinition, ctx any) (TypeDefinition, error) {
		if updated.Name().Name() == "Left" && !changed {
			changed = true
			updated = updated.WithName(left2TypeName)
		}
		return IdentityAfterVisit(original, updated, ctx)
	}

	rootDef := types[rootTypeName]
	updatedTypes, err := walker.Walk(rootDef)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(updatedTypes)).To(Equal(3))
	g.Expect(updatedTypes).To(HaveKey(left2TypeName))
	g.Expect(updatedTypes).To(HaveKey(leftTypeName))
	g.Expect(TypeEquals(updatedTypes[left2TypeName].Type(), updatedTypes[leftTypeName].Type())).To(BeTrue())
}

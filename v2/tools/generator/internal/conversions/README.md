# Property Conversions

One of our challenges is generating the code to copy losslessly between different versions of each resource.

## How it works

The primary entry point is [`CreateTypeConversion`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/conversions/property_conversions.go#L150):

``` go
func CreateTypeConversion(
    sourceEndpoint *TypedConversionEndpoint,
    destinationEndpoint *TypedConversionEndpoint,
    conversionContext *PropertyConversionContext) (PropertyConversion, error)
```

* `sourceEndpoint` captures the name and type of the source value that needs to be transferred;
* `destinationEndpoint` similarly captures the name and type of the destination that needs to be populated;
* `conversionContext` contains other information about the conversion being generated.

The core algorithm works by scanning through the handlers listed in [`propertyConversionFactories`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/conversions/property_conversions.go#L51) to find one that can handle the current conversion. Many of those conversions recursively call `CreateTypeConversion` with a simpler conversion, breaking the problem down into simpler steps.

All of the handlers follow a similar structure:

* Check prerequisites
* Preparation
* Generation

In the examples below we'll see how these play out.

## Simple properties

Support for simple assignments, where the value is a simple atomic value (such as an **int** or **string** value) that needs to be copied across, is provided by `assignPrimitiveFromPrimitive()`.

All handlers have the same function signature:

``` go
func assignPrimitiveFromPrimitive(
    sourceEndpoint *TypedConversionEndpoint,
    destinationEndpoint *TypedConversionEndpoint,
    _ *PropertyConversionContext) (PropertyConversion, error) {
```

The function starts with checking prerequisites - the conditions under which this handler will trigger. Each bails out early if we find that the prerequisite is not satisfied.

The type of an endpoint is typically a composite type - for example, an `*int` is represented by an [`OptionalType`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/astmodel/object_type.go#L20) wrapping the underlying primitive type. Our handler prerequisites therefore need to check for these wrapper types to avoid triggering inappropriately.

``` go
    // Require both source and destination to not be bag items
    if sourceEndpoint.IsBagItem() || destinationEndpoint.IsBagItem() {
        return nil, nil
    }

    // Require both source and destination to be non-optional
    if sourceEndpoint.IsOptional() || destinationEndpoint.IsOptional() {
        return nil, nil
    }
```

Only primitive types can be simply copied across, so we check that both endpoints are the same primitive type.

``` go
    // Require source to be a primitive type
    sourcePrimitive, sourceIsPrimitive := astmodel.AsPrimitiveType(sourceEndpoint.Type())
    if !sourceIsPrimitive {
        return nil, nil
    }

    // Require destination to be a primitive type
    destinationPrimitive, destinationIsPrimitive := astmodel.AsPrimitiveType(destinationEndpoint.Type())
    if !destinationIsPrimitive {
        return nil, nil
    }

    // Require both properties to have the same primitive type
    if !astmodel.TypeEquals(sourcePrimitive, destinationPrimitive) {
        return nil, nil
    }
```

We finish by returning a function that creates the required code.

``` go
    return func(
        reader dst.Expr, 
        writer func(dst.Expr) []dst.Stmt, 
        knownLocals *astmodel.KnownLocalsSet, 
        generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
        return writer(reader)
    }, nil
}
```

* `reader` is an expression giving access to the original value
* `writer` is a function takes an expression and writes that to the final destination
* `knownLocals` is a set that contains all local variables declared so far (not used in this handler)
* `generationContext` is context for code generation (including which packages are available for import)

For a simple property assignment, we don't need any intermediate transformation, so we can just write the value directly.

## Optional Values

When one or the other values is optional, we can't just copy the pointer across - we'd end up sharing the contents between the source and destination. This may not be safe as changes made to one object would be visible to another.

A pair of handlers - `assignFromOptional` and `assignToOptional` - work together to handle these cases.

Let's look at how `assignFromOptional` works. We start with the standard function signature required of all handlers:

``` go
func assignFromOptional(
    sourceEndpoint *TypedConversionEndpoint,
    destinationEndpoint *TypedConversionEndpoint,
    conversionContext *PropertyConversionContext) (PropertyConversion, error) {
```

For our prerequisites, we require that we aren't reading from a property bag, and that the source must be optional.

``` go
    // Require source to not be a bag item
    if sourceEndpoint.IsBagItem() {
        return nil, nil
    }

    // Require source to be optional
    sourceOptional, sourceIsOptional := astmodel.AsOptionalType(sourceEndpoint.Type())
    if !sourceIsOptional {
        return nil, nil
    }
```

We need to be able to convert between the non-optional source and our destination.

To work out whether this is supported, we make a recursive call to `CreateTypeConversion`. If that doesn't work, we bail out.

``` go
    // Require a conversion between the unwrapped type and our source
    unwrappedEndpoint := sourceEndpoint.
        WithType(sourceOptional.Element())
    conversion, err := CreateTypeConversion(
        unwrappedEndpoint,
        destinationEndpoint,
        conversionContext)
    if err != nil {
        return nil, err
    }
    if conversion == nil {
        return nil, nil
    }
```

Since we have the required conversion, we return our generator function.

``` go
    return func(
        reader dst.Expr, 
        writer func(dst.Expr) []dst.Stmt, 
        knownLocals *astmodel.KnownLocalsSet, 
        generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
```

We value the quality of our generated code, so we check what kind of `reader` we are given. If it's a local variable or a field selector, we can use it directly and skip creating a local variable.

``` go
        var cacheOriginal dst.Stmt
        var actualReader dst.Expr

        switch reader.(type) {
        case *dst.Ident, *dst.SelectorExpr:
            // reading a local variable or a field
            cacheOriginal = nil
            actualReader = reader
        default:
            // Something else, so we cache the original
            local := knownLocals.CreateSingularLocal(sourceEndpoint.Name(), "", "AsRead")
            cacheOriginal = astbuilder.ShortDeclaration(local, reader)
            actualReader = dst.NewIdent(local)
        }
```

* `cacheOriginal` is an optional statement that reads our value and stashes it in a local variable.
* `actualReader` is an expression for use later in the function. It reads from our local variable (if used), or directly from our source (if not).

``` go
        checkForNil := astbuilder.AreNotEqual(actualReader, astbuilder.Nil())
```

If our source value is nil, we are going to write a zero value (to ensure we always assign *something*); otherwise we use the nested conversion we found earlier to write the value.

``` go
        writeActualValue := conversion(
            astbuilder.Dereference(actualReader),
            writer,
            knownLocals.Clone(),
            generationContext)

        writeZeroValue := writer(
            destinationEndpoint.Type().
                AsZero(conversionContext.Types(), generationContext))

        stmt := astbuilder.SimpleIfElse(
            checkForNil,
            writeActualValue,
            writeZeroValue)
```

We use a clone of `knownLocals` as any local variables created by the nested conversion will be scoped within our **if** statement and won't be visible to later code.

From the `astbuilder` package we use `Statements()` to stitch together our fragments into a sequence of statements that we return. This helper is smart enough to omit any missing fragments (as when we don't need `cacheOriginal`) and to flatten any passed statement slices.

``` go
        return astbuilder.Statements(cacheOriginal, stmt)
    }, nil
}
```

## Container types

If the property is a container (a list or dictionary), we need to copy all the individual items across. The [`assignArrayFromArray`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/conversions/property_conversions.go#L863) and [`assignMapFromMap`](https://github.com/Azure/azure-service-operator/blob/main/v2/tools/generator/internal/conversions/property_conversions.go#L990) handlers do this.

The `assignArrayFromArray` handler is representative. As usual, we start with the standard function signature.

``` go 
func assignArrayFromArray(
    sourceEndpoint *TypedConversionEndpoint,
    destinationEndpoint *TypedConversionEndpoint,
    conversionContext *PropertyConversionContext) (PropertyConversion, error) {
```
We don't support bag items here.

``` go
    // Require both source and destination to not be bag items
    if sourceEndpoint.IsBagItem() || destinationEndpoint.IsBagItem() {
        return nil, nil
    }
```

Both the source and destination must be array types.
``` go
    // Require source to be an array type
    sourceArray, sourceIsArray := 
        astmodel.AsArrayType(sourceEndpoint.Type())
    if !sourceIsArray {
        return nil, nil
    }

    // Require destination to be an array type
    destinationArray, destinationIsArray := 
        astmodel.AsArrayType(destinationEndpoint.Type())
    if !destinationIsArray {
        return nil, nil
    }
```

We must have a way to convert each item from the source array into an item for the destination array. 

To find such a conversion we create a pair of endpoints by unwrapping the array definitions, and then recursively call `CreateTypeConversion`.

``` go
    // Require a conversion between the array definitions
    unwrappedSourceEndpoint := sourceEndpoint.
        WithType(sourceArray.Element())
    unwrappedDestinationEndpoint := destinationEndpoint.
        WithType(destinationArray.Element())
    conversion, err := CreateTypeConversion(
        unwrappedSourceEndpoint,
        unwrappedDestinationEndpoint,
        conversionContext)
```

Returning a high quality error message is critical to anyone trying to debug conversion failures.

``` go
    if err != nil {
        return nil, errors.Wrapf(
            err,
            "finding array conversion from %s to %s",
            astmodel.DebugDescription(
                sourceEndpoint.Type(), 
                conversionContext.Types()),
            astmodel.DebugDescription(
                destinationEndpoint.Type(), 
                conversionContext.Types()))
    }
    if conversion == nil {
        return nil, nil
    }
```

With all our prerequisites satisfied, we return our conversion function.

``` go 
    return func(
        reader dst.Expr, 
        writer func(dst.Expr) []dst.Stmt, 
        knownLocals *astmodel.KnownLocalsSet, 
        generationContext *astmodel.CodeGenerationContext) []dst.Stmt {
``` 

As we did above in `assignFromOptional`, we want to keep the quality of our generated code high by avoiding a local variable if we don't need one.

``` go
        var cacheOriginal dst.Stmt
        var actualReader dst.Expr

        switch reader.(type) {
        case *dst.Ident, *dst.SelectorExpr:
            // reading a local variable or a field
            cacheOriginal = nil
            actualReader = reader
        default:
            // Something else, so we cache the original
            local := knownLocals.CreateSingularLocal(
                sourceEndpoint.Name(), 
                "", 
                "Cache")
            cacheOriginal = astbuilder.ShortDeclaration(
                local, 
                reader)
            actualReader = dst.NewIdent(local)
        }
```

* `cacheOriginal` is an optional statement that reads our value and stashes it in a local variable.
* `actualReader` is an expression for use later in the function. It reads from our local variable (if used), or directly from our source (if not).

Our array may be nil, so we need to check and avoid dereferencing it if so.

``` go
        checkForNil := astbuilder.AreNotEqual(actualReader, astbuilder.Nil())
```

Using `knownLocals`, we create three obviously related identifiers to use for the array conversion. Each uses the name of our source as a base, with a different suffix.

* `List` is the final list we're creating
* `Item` is an item from the source list
* `Index` is the index of the item we're converting

Clones of `knownLocals` are used for each nested scope so that definitions don't leak out.

We use different suffixes here to those used in `assignMapFromMap` in order to keep the code easy to read if the two are nested.

``` go
        branchLocals := knownLocals.Clone()
        tempId := branchLocals.
            CreateSingularLocal(sourceEndpoint.Name(), "List")

        loopLocals := branchLocals.Clone()
        itemId := loopLocals.
            CreateSingularLocal(sourceEndpoint.Name(), "Item")
        indexId := loopLocals.
            CreateSingularLocal(sourceEndpoint.Name(), "Index")
```

Our declaration intializes a new list of the correct final size.

``` go
        declaration := astbuilder.ShortDeclaration(
            tempId,
            astbuilder.MakeSlice(
                destinationArray.AsType(generationContext), 
                astbuilder.CallFunc("len", actualReader)))
```

We create a custom writer to assign values into our list by index.

``` go
        writeToElement := func(expr dst.Expr) []dst.Stmt {
            return []dst.Stmt{
                astbuilder.SimpleAssignment(
                    &dst.IndexExpr{
                        X:     dst.NewIdent(tempId),
                        Index: dst.NewIdent(indexId),
                    },
                    expr),
            }
        }
```

Within our loop we create a copy of the item variable, to avoid aliasing between loops. We don't know what our inner conversions are going to do, so this is a safety precaution.

``` go
        avoidAliasing := astbuilder.ShortDeclaration(itemId, dst.NewIdent(itemId))
        avoidAliasing.Decs.Start.
            Append("// Shadow the loop variable to avoid aliasing")
        avoidAliasing.Decs.Before = dst.NewLine
```

The loop body copies one item across to the final array.

``` go
        loopBody := astbuilder.Statements(
            avoidAliasing,
            conversion(
                dst.NewIdent(itemId), 
                writeToElement, 
                loopLocals, 
                generationContext))
```

After iterating through all the items, we assign the final list to the destination by calling `writer`.

``` go
        loop := astbuilder.IterateOverSliceWithIndex(indexId, itemId, reader, loopBody...)
        assignValue := writer(dst.NewIdent(tempId))
```

Remember our test for **nil** earlier? If the source array exists, we use the loop to copy the items across. If not, we assign an explicit **nil**.

``` go
        trueBranch := astbuilder.Statements(
            declaration, 
            loop, 
            assignValue)

        assignZero := writer(astbuilder.Nil())

        return astbuilder.Statements(
            cacheOriginal,
            astbuilder.SimpleIfElse(
                checkForNil, 
                trueBranch, 
                assignZero))
    }, nil
}
```

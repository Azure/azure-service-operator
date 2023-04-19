---
title: Chained Storage Versions
---

This case study explores the alternative solution of using a *chained storage versions*. We update the storage schema of each resource each release of the service operator. We'll keep the storage version up to date with the latest GA release of each resource. Older storage versions are retained, both as intermediate steps in the hub-and-spoke conversions, and to allow upgrades.

For the purposes of discussion, we'll be following the version by version evolution of a theoretical ARM service that provides customer resource management (CRM) services. Synthetic examples are used to allow focus on specific scenarios one by one, providing motivation for specific features.

Examples shown are deliberately simplified in order to focus, and therefore minutiae should be considered motivational, not binding. Reference the formal specification for precise details.

# Version 2011-01-01 - Initial Release

The initial release of the CRM includes a simple definition to capture information about a particular person:

``` go
package v20110101

type Person struct {
    Id        Guid
    FirstName string
    LastName  string
}
```

We're not reusing the API version directly as our storage version. Instead, we define a separate (independent) type with a similar structure:

``` go
package v20110101storage

type Person struct {
    PropertyBag
    FirstName   *string
    Id          *Guid
    LastName    *string
}

// Hub marks this type as a conversion hub.
func (*Person) Hub() {}
```

Every property is marked as optional. Optionality doesn't matter at this point, as we currently have only single version of the API. However, as we'll see with later versions, forward and backward compatibility issues would arise if they were not optional.

The `PropertyBag` type provides storage for other properties, plus helper methods. It is always included in storage versions, but in this case will be unused. The method `Hub()` marks this version as the storage schema.

## Storage Conversion

We need to implement the [Convertible](https://book.kubebuilder.io/multiversion-tutorial/conversion.html) interface to allow conversion to and from the storage version:

``` go
package v20110101

import storage "v20110101storage"

// ConvertTo converts this Person to the Hub storage version.
func (person *Person) ConvertTo(raw conversion.Hub) error {
    p := raw.(*storage.Person)
    return ConvertToStorage(p)
}

// ConvertToStorage converts this Person to a storage version
func (person *Person) ConvertToStorage(dest storage.Person) error {
    // Copy simple properties across
    dest.FirstName = person.FirstName
    dest.Id = person.Id
    dest.LastName = person.LastName

    return nil
}

// ConvertFrom converts from the Hub storage version
func (person *Person) ConvertFrom(raw conversion.Hub) error {
    p := raw.(*storage.Person)
    return ConvertFromStorage(p)
}

// ConvertFrom converts from a storage version to this version.
func (person *Person) ConvertFromStorage(source storage.Person) error {
    // Copy simple properties across
    person.FirstName = source.FirstName
    person.Id = source.Id
    person.LastName = source.LastName

    return nil
}

```

Conversion is separated into two methods (e.g. `ConvertFrom()` and `ConvertFromStorage()`) to allow for reuse of the `ConvertFromStorage()` methods for conversion of nested complex properties, as we'll see later on.

These methods will be automatically generated in order to handle the majority of the required conversions. Since they never change, the `ConvertTo()` and `ConvertFrom()` methods are omitted from the following discussion.

## Version Map

With only two classes, our version map is simple and straightforward.

![](2011-01-01.png)


# Version 2012-02-02 - No Change

In this release of the CRM service, there are no changes made to the structure of `Person`:

``` go
package v20120202

type Person struct {
    Id        Guid
    FirstName string
    LastName  string
}
```

## Storage Conversion

The existing conversion between the `v20110101` API version and `v20110101storage` version is retained, preserving in-place a conversion that's already known to be reliable.

The new API version `20120202` has a matching storage version `v20120202storage` which becomes the authoratative storage version for the CRD. This conversion is identical to the earlier version.

An additional bidirectional conversion between `v20110101storage` and `v20120202storage` is also generated. Since both versions have the same structure, this is also trivial.

## Version Map

Our version map diagram is becoming useful for seeing the relationship between versions:

![](2012-02-02.png)

Observe that the prior storage version is still shown, with a bidirectional conversion with the current storage version. Existing users who upgrade their service operator will have their storage upgraded using this conversion. The conversion between storage versions will be generated with the same approach, and with the same structure, as all our other conversions.

# Version 2013-03-03 - New Property

In response to customer feedback, this release of the CRM adds a new property to `Person` to allow a persons middle name to be stored:

``` go
package v20130303

type Person struct {
    Id         Guid
    FirstName  string
    MiddleName string // *** New ***
    LastName   string
}
```

The new storage version, based on this version, is what you'd expect:

``` go
package v20130303storage

type Person struct {
    PropertyBag
    Id          *Guid
    FirstName   *string
    MiddleName  *string // *** New storage ***
    LastName    *string
}

// Hub marks this type as a conversion hub.
func (*Person) Hub() {}
```

## Storage Conversions

Conversions to and from earlier versions of Person are unchanged, as those versions do not support `MiddleName`. For the new version of `Person`, the new property will be included in the generated methods:

``` go
package v20130303

import storage "v20130303storage"

// ConvertTo converts this Person to the Hub version.
func (person *Person) ConvertToStorage(dest storage.Person) error {
    dest.FirstName = person.FirstName
    dest.Id = person.Id
    dest.LastName = person.LastName
    dest.MiddleName = person.MiddleName // *** New property copied too ***

    return nil
}

// ConvertFrom converts from the Hub version to this version.
func (person *Person) ConvertFromStorage(source storage.Person) error {
    person.FirstName = source.FirstName
    person.Id = source.Id
    person.LastName = source.LastName
    person.MiddleName = source.MiddleName // *** New property copied too ***

    return nil
}
```

The new property is shown at the end of the list not because it is new, but because values are copied across in alphabetical order. This is to guarantee that code generation is deterministic and generates the same result each time.

Conversion methods for earlier API versions of `Person` are unchanged, as they still convert to the same storage versions. 

A new bidirectional conversion between `v20120202storage` and `v20130303storage` versions is introduced. When down-converting to `v20120202storage`, the `MiddleName` property is stashed in the property bag; when up-converting to `v20130303storage`, the PropertyBag is checked to see if it contains `MiddleName`:

``` go
package v20120202storage

import vnext "v20130303storage"

// ConvertTo converts this Person to the storage Hub version.
func (person *Person) ConvertToStorage(dest vnext.Person) error {
    dest.FirstName = person.FirstName
    dest.Id = person.Id
    dest.LastName = person.LastName

    if middleName, ok := PropertyBag.ReadString("MiddleName"); ok {
        dest.MiddleName = middleName // *** New property copied too ***
    }

    return nil
}

// ConvertFrom converts from the Hub version to this version.
func (person *Person) ConvertFromStorage(vnext storage.Person) error {
    person.FirstName = source.FirstName
    person.Id = source.Id
    person.LastName = source.LastName

    person.WriteString("MiddleName", source.MiddleName)

    return nil
}
```

## Version Map

A graph of our conversions now starts to show the chaining between storage versions that gives the name to this approach. Bidirectional conversions to and from earlier versions of storage allow conversion between any pairs of API versions.

![](2013-03-03.png)

## How often are new properties added?

At the time of writing, there were **381** version-to-version changes where the only change between versions was solely the addition of new properties. Of those, **249** were adding just a single property, and **71** added two properties. 

# Version 2014-04-04 Preview - Schema Change

To allow the CRM to better support cultures that have differing ideas about how names are written, a preview release of the service modifies the schema considerably:

``` go
package v20140404preview

type Person struct {
    Id         Guid   // ** Only Id is unchanged ***
    FullName   string
    FamilyName string
    KnownAs    string
}
```

This is a preview version, but it still gets a dedicated storage version, `v20140404previewStorage`. **The official hub version is left unchanged as `v20130303storage`**.

## Storage Conversion

The new properties don't exist on prior storage versions, so the generated `ConvertToStorage()` and `ConvertFromStorage()` methods used to convert between `v20130303storage` and `v20140404previewStorage` must use the `PropertyBag` to carry the properties:

``` go
package v20140404previewStorage

import vprior "v20130303storage"

// ConvertTo converts this Person to the Hub version.
func (person *Person) ConvertToStorage(dest vprior.Person) error {
    dest.Id = person.Id

    // *** Store in the property bag ***
    dest.WriteString("FamilyName", person.FamilyName)
    dest.WriteString("FullName", person.FullName)
    dest.WriteString("KnownAs", person.KnownAs)

    return nil
}

// ConvertFrom converts from the Hub version to this version.
func (person *Person) ConvertFromStorage(source vprior.Person) error {
    person.Id = source.Id

    // *** Read from the property bag ***

    if familyName, ok := source.ReadString("FamilyName");ok {
        person.FamilyName = familyName
    }

    if fullName, ok := source.ReadString("FullName"); ok {
        person.FullName = fullName
    }

    if knownAs, ok := source.ReadString("KnownAs"); ok {
        person.KnownAs = knownAs
    }

    return nil
}
```

In the example above, we show first copying all the directly supported properties, then using the property bag. We may not separate these steps in the generated code.

These methods are always generated on the storage versions furthest from the hub version, converting towards that version. In the usual case we'll use the import name `vnext` (or equivalent) but in this case, given we have a preview version, we'll use `vprior` to emphasize the direction of conversion.

This provides round-trip support for the preview release, but does not provide backward compatibility with prior official releases.

The storage version of a `Person` written by the preview release will have no values for `FirstName`, `LastName`, and `MiddleName`. Similarly, an older version won't have `FamilyName`, `FullName` nor `KnownAs`.

These kinds of cross-version conversions cannot be automatically generated as they require more understanding the semantic changes between versions. 

To allow injection of manual conversion steps, interfaces will be generated as follows:

``` go
package v20130303storage

// AssignableToPersonV20130303 provides methods to augment conversion to storage
type AssignableToPersonV20130303 interface {
    AssignToV20130303(person Person) error
}

// AssignableFromPersonV20130303 provides methods to augment conversion from storage
type AssignableFromPersonV20130303 interface {
    AssignFromV20130303(person Person) error
}
```

This interface can be optionally implemented by API versions (spoke types) to augment the generated conversion.

----

**Outstanding Issue**: The interfaces and methods shown above include the version number of the target in order to disambiguate between versions. This is necessitated by having multiple storage versions in flight at the same time, and needing to avoid name collisions. Contrast this with the *rolling storage version* case study where there's only one active storage version at a time. 

Is there a way we could structure this approach to avoid the need for version numbers in method names?

----

The generated `ConvertToStorage()` and `ConvertFromStorage()` methods will test for the presence of this interface and will call it if available:

``` go
package v20140404preview

import storage "v20130303storage"

// ConvertTo converts this Person to the Hub version.
func (person *Person) ConvertToStorage(dest storage.Person) error {
    // … property copying and property bag use elided …

    // *** Check for the interface and use it if found ***
    if assignable, ok := person.(AssignableToPersonV20130303); ok {
        assignable.AssignToV20130303(dest)
    }

    return nil
}

// ConvertFrom converts from the Hub version to this version.
func (person *Person) ConvertFromStorage(source storage.Person) error {
    // … property copying and property bag use elided …

    // *** Check for the interface and use it if found ***
    if assignable, ok := person.(AssignableFromPersonV20130303); ok {
        assignable.AssignFromV20130303(source)
    }

    return nil
}
```

## Version Map
 
Preview releases, by definition, include unstable changes that may differ once the feature reaches general availability.

We don't want to make changes to our storage versions based on these speculative changes, so we handle persistence of the preview release with the existing storage version, by way of a down-conversion to `v20130303storage`:

![](2014-04-04-preview.png)


# Version 2014-04-04 - Schema Change

Based on feedback generated by the preview release, the CRM schema changes have gone ahead with a few minor changes:

``` go
package v20140404

type Person struct {
    Id         Guid
    LegalName  string // *** Was FullName in preview ***
    FamilyName string
    KnownAs    string
    AlphaKey   string // *** Added after preview ***
}
```

As usual, a custom storage version is generated:

``` go
package v20140404storage

type Person struct {
    PropertyBag
    AlphaKey    *string
    FamilyName  *string
    LegalName   *string
    Id          *Guid
    KnownAs     *string
}

// Hub marks this type as a conversion hub.
func (*Person) Hub() {}
```

## Storage Conversion

The `ConvertToStorage()` and `ConvertFromStorage()` methods between the API version `v20140404` and the storage version `v20140404storage` are trivial and not shown.

For conversions between storage versions, the preview storage version is not considered - it's out of the main line of processing. Instead, we have a bidirectional conversion between `v20130303storage` and `v20140404storage`. As usual, the conversion is implemented further away from the (new) hub version, on `v20130303storage`.

With a large difference in structure between the two versions, the PropertyBag gets a workout:

``` go
package v20130303storage

import vnext "v20140404storage"

// ConvertTo converts this Person to the Hub version.
func (person *Person) ConvertToStorage(dest vnext.Person) error {
    dest.Id = person.Id

    dest.WriteString("FirstName", person.FirstName)
    dest.WriteString("LastName",  person.LastName)
    dest.WriteString("MiddleName",  person.MiddleName)

    if assignable, ok := person.(AssignableToPerson); ok {
        assignable.AssignTo(dest)
    }

    return nil
}

// ConvertFrom converts from the Hub version to this version.
func (person *Person) ConvertFromStorage(source vnext.Person) error {
    person.Id = source.Id

    if firstName, ok := source.ReadString("FirstName"); ok {
        person.FirstName = firstName
    }

    if middleName, ok := source.ReadString("MiddleName"); ok {
        person.MiddleName = middleName
    }

    if lastName, ok := source.ReadString("LastName"); ok {
        person.LastName = lastName
    }

    // *** Check for the interface and use it if found ***
    if assignable, ok := person.(AssignableFromPersonV20140404); ok {
        assignable.AssignFromV20140404(source)
    }

    return nil
}
```

To interoperate between different versions of `Person`, we need to add manual conversions between the storage versions where the schema change occcurs.

When we are converting from `v20130303storage` up to `v20140404storage`, we need to use `FirstName`, `LastName` and `MiddleName` to populate `AlphaKey`, `FamilyName`, `KnownAs` and `LegalName`.

Conversely, When we are converting from `v20140404storage` down to `v20130303storage`, we need to use `AlphaKey`, `FamilyName`, `KnownAs` and `LegalName` to populate `FirstName`, `LastName` and `MiddleName`.

These conversions occur *in addition* to use of the PropertyBag to store those same properties.

``` go
package v20130303storage

import vnext "v20140404storage"

func (person *Person) AssignToV20140404(dest vnext.Person) error {
    if dest.KnownAs == "" {
        dest.KnownAs = person.FirstName
    }

    if dest.FamilyName == "" {
        dest.FamilyName = person.LastName
    }

    if dest.LegalName == "" {
        dest.LegalName = person.FirstName +" "+ person.MiddleName + " " + person.LastName
    }

    if dest.AlphaKey == "" {
        dest.AlphaKey = person.lastName
    }
}

func (person *Person) AssignFrom(source vNext.Person) error {
    if person.FirstName == "" {
        person.FirstName = source.KnownAs
    }

    if person.LastName == "" {
        person.LastName = source.FamilyName
    }

    if person.MiddleName == "" {
        person.MiddleName = // ... elided ...
    }
}
```

For each property we need to consider that it might have already been populated with a more accurate value from the PropertyBag, so we only synthesize values when needed.

## Version Map

We can see in our version map that the preview release is still supported, but the associated storage version is not in the main chain of interconvertible versions.

![](2014-04-04.png)

# Version 2015-05-05 - Property Rename

The term `AlphaKey` was found to be confusing to users, so in this release of the API it is renamed to `SortKey`. This better reflects its purpose of sorting names together (e.g. so that the family name *McDonald* gets sorted as though spelt *MacDonald*).

``` go
package v20150505

type Person struct {
    Id         Guid
    LegalName  string
    FamilyName string
    KnownAs    string
    SortKey    string // *** Used to be AlphaKey ***
}
```

As expected, a matching storage version is also generated:

``` go
package v20150505storage

type Person struct {
    PropertyBag
    Id          *Guid
    LegalName   *string
    FamilyName  *string
    KnownAs     *string
    SortKey     *string
}

// Hub marks this type as a conversion hub.
func (*Person) Hub() {}
```

## Storage Conversion

By documenting the renames in the configuration of our code generator, this rename will be automatically handled within the `ConvertTo()` and `ConvertFrom()` methods that are generated between the `v20140404storage` and `v20150505storage` versions:

``` go
package v20140404

import vNext "v20150505storage"

// ConvertTo converts this Person to the Hub version.
func (person *Person) ConvertToStorage(dest vNext.Person) error {
    dest.FamilyName = person.FamilyName
    dest.Id = person.Id
    dest.KnownAs = person.KnownAs
    dest.LegalName = person.LegalName
    dest.SortKey = person.AlphaKey // *** Rename is automatically handled ***

    if assignable, ok := person.(AssignableToPerson); ok {
        assignable.AssignTo(dest)
    }

    return nil
}

// ConvertFrom converts from the Hub version to this version.
func (person *Person) ConvertFromStorage(source vNext.Person) error {
    person.AlphaKey = source.SortKey // *** Rename is automatically handled ***
    person.FamilyName = source.FamilyName
    person.Id = source.Id
    person.KnownAs = source.KnownAs
    person.LegalName = source.LegalName

    if assignable, ok := person.(AssignableFromPerson); ok {
        assignable.AssignFrom(source)
    }

    return nil
}
```

While `SortKey` appears at the end of the list of assignments in the first method, the mirror assignment of `AlphaKey` appears at the start of the list in the second method. In both cases the properties are shown in alphabetical order.

## Version Map

Here we see our horizon policy coming into effect, with support for version 2011-01-01 being dropped in this release:

![](2015-05-05.png)

For users staying up to date with releases of the service operator, this will likely have no effect - but users still using the original release (storage version `v2011-01-01storage`) will need to update to an intermediate release before adopting this version.

An alternative approach would be to always support conversion from every storage version, even if the related API version has been dropped:

![](2015-05-05-alternate.png)

This would allow users to upgrade from almost any older version of the service operator. ("Almost" because we would still have older versions drop off when they are retired by ARM.)

## How often do property renames happen?

At the time of writing, there were nearly **60** cases of properties being renamed between versions; **17** of these involved changes to letter case alone. (Count is somewhat inexact because renaming was manually inferred from the similarity of names.)

# Version 2016-06-06 - Complex Properties

With some customers expressing a desire to send physical mail to their customers, this release extends the API with mailing address for each person.

``` go
package v20160606

type Address struct {
    Street string
    City   string
}

type Person struct {
    Id             Guid
    LegalName      string
    FamilyName     string
    KnownAs        string
    SortKey        string
    MailingAddress Address
}
```

We now have two structs that make up our storage version:

``` go
package v20160606storage

type Person struct {
    PropertyBag
    Id             *Guid
    LegalName      *string
    FamilyName     *string
    KnownAs        *string
    MailingAddress *Address // *** New ***
    SortKey        *string
}

type Address struct {
    PropertyBag
    City        *string
    Street      *string
}

// Hub marks this type of Person as a conversion hub.
func (*Person) Hub() {}
```

## Storage Conversion

The required `ConvertToStorage()` and `ConvertFromStorage()` methods between the API version `v20160606` and the storage version `v201606061` get generated in the expected way:

``` go
package v20160606

import storage "v20160606storage"

// ConvertTo converts this Person to the Storage version.
func (person *Person) ConvertToStorage(dest storage.Person) error {
    dest.FamilyName = person.FamilyName
    dest.Id = person.Id
    dest.KnownAs = person.KnownAs
    dest.LegalName = person.LegalName
    dest.SortKey = person.AlphaKey

    // *** Copy the mailing address over too ***
    address := &storage.Address{}
    err := person.MailingAddress.ConvertToStorage(address)
    if err != nil {
        return err
    }

    dest.MailingAddress = address

    if assignable, ok := person.(AssignableToPerson); ok {
        err := assignable.AssignTo(dest)
        if err != nill {
            return err
        }
    }

    return nil
}

// ConvertToStorage converts this Address to the hub storage version
func (address *Address) ConvertToStorage(dest storage.Address) error {
    dest.City = address.City
    dest.Street = address.Street

    if assignable, ok := person.(AssignableToAddress); ok {
        err := assignable.AssignTo(dest)
        if err != nill {
            return err
        }
    }

    return nil
}

// ConvertFrom converts from the Hub version to this version.
func (person *Person) ConvertFromStorage(source storage.Person) error {
    person.AlphaKey = source.SortKey
    person.FamilyName = source.FamilyName
    person.Id = source.Id
    person.KnownAs = source.KnownAs
    person.LegalName = source.LegalName

    // *** Copy the mailing address over too ***
    if storage.MailingAddress != nil {
        address := &Address{}
        err := address.ConvertFromStorage(storage.Address)
        person.MailingAddress = address
    }

    if assignable, ok := person.(AssignableFromPerson); ok {
        err := assignable.AssignFrom(source)
        if err != nill {
            return err
        }
    }

    return nil
}

// ConvertFromStorage converts from the hub storage version to this version
func (address *Address) ConvertFromStorage(source storage.Address) error {
    address.Street = source.Street
    address.City = source.City

    if assignable, ok := person.(AssignableFromAddress); ok {
        err := assignable.AssignFrom(source)
        if err != nill {
            return err
        }
    }

    return nil
}
```

We're recursively applying the same conversion pattern to `Address` as we have already been using for `Person`. This scales to any level of nesting without the code becoming unweildy.

## Version Map

Again we see the oldest version drop out, allowing users of the three prior versions of the service operator to upgrade cleanly:

![](2016-06-06.png)

# Version 2017-07-07 - Optionality changes

In the `2016-06-06` version of the API, the `MailingAddress` property was mandatory. Since not everyone has a mailing address (some people receive no physical mail), this is now being made optional.

The change to the API declarations is simple:

``` go
package v20170707

type Address struct {
    Street string
    City   string
}

type Person struct {
    Id             Guid
    LegalName      string
    FamilyName     string
    KnownAs        string
    SortKey        string
    MailingAddress *Address // *** Was mandatory, now optional ***
}
```

## Storage Conversion

The storage versions are identical to those used previously and are not shown here.

What does change is the `ConvertToStorage()` method, which now needs to handle the case where the `MailingAddress` has not been included:

``` go
package v20170707

import storage "v20170707storage"

// ConvertTo converts this Person to the Hub version.
func (person *Person) ConvertToStorage(dest storage.Person) error {
    dest.SortKey = person.AlphaKey
    dest.FamilyName = person.FamilyName
    dest.Id = person.Id
    dest.KnownAs = person.KnownAs
    dest.LegalName = person.LegalName

    // *** Need to check whether we have a mailing address to copy ***
    if person.MailingAddress != nil {
        address := &storage.Address{}
        err := person.MailingAddress.ConvertToStorage(address)
        if err != nil {
            return err
        }

        dest.MailingAddress = address
    }

    if assignable, ok := person.(AssignableToPerson); ok {
        err := assignable.AssignTo(dest)
        if err != nill {
            return err
        }
    }

    return nil
}
```

If we instead had an _optional_ field that became _required_ in a later version of the API, the generated code for `ConvertToStorage()` would become simpler as the check for **nil** would not be needed.

## Version Map

![](2017-07-07.png)

## How often does optionality change?

At the time of writing, there are **100** version-to-version changes where fields became **optional** in the later version of the API, and **99** version-to-version changes where fields became **required**.

# Version 2018-08-08 - Extending nested properties

Defining an address simply as `Street` and `City` has been found to be overly simplistic, so this release makes changes to allow a more flexible approach.

``` go
package v20180808

type Address struct {
    // FullAddress shows the entire address as should be used on postage
    FullAddress  string
    City         string
    Country      string
    PostCode     string
}
```

As before, the storage version is generated to match, with prior conversions using the property bag to store additional properties:

``` go
package v20180808storage

type Address struct {
    PropertyBag
    City         *string
    Country      *string
    FullAddress  *string
    PostCode     *string
}
```

These changes are entirely similar to those previously covered in version 2014-04-04, above.

## Version Map

In this release, we see that support for both `2014-04-04` and the preview version `2014-04-04preview` has been dropped:

![](2018-08-08.png)

Users still running earlier releases of the service operator that are using `2014-04-04` or earlier will need to install an intermediate release in order to upgrade to this one.

# Version 2019-09-09 - Changing types

Realizing that some people get deliveries to places that don't appear in any formal database of addresses, in this release the name of the type changes to `Location` and location coordinates are added:

``` go
package v20190909

type Location struct {
    FullAddress string
    City        string
    Country     string
    PostCode    string
    Latitude    double
    Longitude   double
}
```

The storage version gets generated in a straightforward way:

``` go
package v20190909storage

type Location struct {
    PropertyBag
    City        *string
    Country     *string
    FullAddress *string
    Latitude    *double
    Longitude   *double
    PostCode    *string
}
```

## Storage Conversion

The conversion methods need to change as well. If we configure metadata detailing the rename of the type (as we did for properties in version 2015-05-05), we can generate the required conversions automatically:

``` go
package v20180808storage

// *** Updated storage version ***
import vNext "v20190909storage"

// ConvertTo converts this Person to the Hub version.
func (person *Person) ConvertToStorage(dest vNext.Person) error {
    // ... elided properties ...

    if person.MailingAddress != nil {
        address := &vNext.Location{} // ** New Type ***
        err := person.MailingAddress.ConvertToStorage(address)
        if err != nil {
            return err
        }

        dest.MailingAddress = address
    }

    if assignable, ok := person.(AssignableToPerson); ok {
        err := assignable.AssignTo(dest)
        if err != nill {
            return err
        }
    }

    return nil
}

// ConvertToStorage converts this Address to the hub storage version
// ** Different parameter type for dest *** 
func (address *Address) ConvertToStorage(dest vNext.Location) error {
    dest.Street = address.Street
    dest.City = address.City

    // *** Interface has been renamed too **
    if assignable, ok := person.(AssignableToLocation); ok {
        err := assignable.AssignTo(dest)
        if err != nill {
            return err
        }
    }

    return nil
}

```

If we don't include metadata to capture type renames, the conversion can be manually injected by implementing the `AssignableToLocation` interface.

## Version Map

![](2019-09-09.png)

## How often do properties change their type?

At the time of writing, there are **160** version-to-version changes where the type of the property changes. This count excludes cases where an optional property become mandatory, or vice versa.

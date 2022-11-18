/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package test

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

/*
 * Shared building blocks for testing
 */

var (
	// Common groups for testing
	Group                   = "person"
	BatchGroup              = "batch"
	maxRestrictedNameLength = int64(25)

	// Reusable Properties - any package version

	// Properties to simplate an actual ARM resource
	NameProperty = astmodel.NewPropertyDefinition("Name", "name", astmodel.StringType).
			WithDescription("The name of the resource")

	FullNameProperty = astmodel.NewPropertyDefinition("FullName", "fullName", astmodel.StringType).
				WithDescription("As would be used to address mail")

	FamilyNameProperty = astmodel.NewPropertyDefinition("FamilyName", "familyName", astmodel.StringType).
				WithDescription("Shared name of the family")

	KnownAsProperty = astmodel.NewPropertyDefinition("KnownAs", "knownAs", astmodel.StringType).
			WithDescription("How the person is generally known")

	RestrictedNameProperty = astmodel.NewPropertyDefinition("RestrictedName", "restrictedName",
		astmodel.NewValidatedType(astmodel.StringType, astmodel.StringValidations{MaxLength: &maxRestrictedNameLength})).
		WithDescription("The name of the resource, restricted to max 25 character length")

	FullAddressProperty = astmodel.NewPropertyDefinition("FullAddress", "fullAddress", astmodel.StringType).
				WithDescription("Full written address for map or postal use")

	CityProperty = astmodel.NewPropertyDefinition("City", "city", astmodel.StringType).
			WithDescription("City or town (or nearest)")

	StatusProperty = astmodel.NewPropertyDefinition("Status", "status", astmodel.StringType).
			WithDescription("Current status")

	OptionalStringProperty = astmodel.NewPropertyDefinition("OptionalString", "optionalString", astmodel.OptionalStringType).
				WithDescription("An optional string")

	PropertyBagProperty = astmodel.NewPropertyDefinition("PropertyBag", "$propertyBag", astmodel.PropertyBagType).
				WithDescription("Stash for extra properties")

	// Reference Packages for 2020
	Pkg2020  = MakeLocalPackageReference(Group, "v20200101")
	Pkg2020s = astmodel.MakeStoragePackageReference(Pkg2020)

	// Reference Package 2021 preview
	Pkg2021Preview        = MakeLocalPackageReference(Group, "v20211231preview")
	Pkg2021PreviewStorage = astmodel.MakeStoragePackageReference(Pkg2021Preview)

	// Reference Packages for 2021
	Pkg2021  = MakeLocalPackageReference(Group, "v20211231")
	Pkg2021s = astmodel.MakeStoragePackageReference(Pkg2021)

	// Reference Packages for 2022
	Pkg2022  = MakeLocalPackageReference(Group, "v20220630")
	Pkg2022s = astmodel.MakeStoragePackageReference(Pkg2022)

	// Reference Package Batch 2020
	BatchPkg2020  = MakeLocalPackageReference(BatchGroup, "v20200101")
	BatchPkg2020s = astmodel.MakeStoragePackageReference(BatchPkg2020)

	// Reference Package Batch 2021
	BatchPkg2021  = MakeLocalPackageReference(BatchGroup, "v20210630")
	BatchPkg2021s = astmodel.MakeStoragePackageReference(BatchPkg2021)

	// Reference Package Batch 2021 - beta
	BatchPkgBeta2021  = MakeLocalPackageReference(BatchGroup, "v1beta20210101")
	BatchPkgBeta2021s = astmodel.MakeStoragePackageReference(BatchPkg2021)

	// Objects in Pkg2020
	Pkg2020APIVersion = astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(Pkg2020, "APIVersion"),
		astmodel.NewEnumType(astmodel.StringType, astmodel.EnumValue{Identifier: "v2020", Value: "v2020"}))

	// Objects in Pkg2021

	Address2021 = CreateObjectDefinition(Pkg2021, "Address", FullAddressProperty, CityProperty)

	// Reusable Properties - only in Pkg2021

	ResidentialAddress2021 = astmodel.NewPropertyDefinition("ResidentialAddress", "residentialAddress", Address2021.Name())

	PostalAddress2021 = astmodel.NewPropertyDefinition("PostalAddress", "postalAddress", Address2021.Name())
)

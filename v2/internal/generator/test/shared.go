/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package test

import (
	"github.com/Azure/azure-service-operator/v2/internal/generator/astmodel"
)

/*
 * Shared building blocks for testing
 */

var (
	// Common groups for testing
	Group      = "microsoft.person"
	BatchGroup = "microsoft.batch"

	// Reusable Properties - any package version

	FullNameProperty = astmodel.NewPropertyDefinition("FullName", "fullName", astmodel.StringType).
				WithDescription("As would be used to address mail")

	FamilyNameProperty = astmodel.NewPropertyDefinition("FamilyName", "familyName", astmodel.StringType).
				WithDescription("Shared name of the family")

	KnownAsProperty = astmodel.NewPropertyDefinition("KnownAs", "knownAs", astmodel.StringType).
			WithDescription("How the person is generally known")

	FullAddressProperty = astmodel.NewPropertyDefinition("FullAddress", "fullAddress", astmodel.StringType).
				WithDescription("Full written address for map or postal use")

	CityProperty = astmodel.NewPropertyDefinition("City", "city", astmodel.StringType).
			WithDescription("City or town (or nearest)")

	StatusProperty = astmodel.NewPropertyDefinition("Status", "status", astmodel.StringType).
			WithDescription("Current status")

	PropertyBagProperty = astmodel.NewPropertyDefinition("PropertyBag", "$propertyBag", astmodel.PropertyBagType).
				WithDescription("Stash for extra properties")

	// Reference Packages for 2020
	Pkg2020  = MakeLocalPackageReference(Group, "v20200101")
	Pkg2020s = astmodel.MakeStoragePackageReference(Pkg2020)

	// Reference Package 2021 preview
	Pkg2021Preview = MakeLocalPackageReference(Group, "v20211231preview")

	// Reference Packages for 2021
	Pkg2021  = MakeLocalPackageReference(Group, "v20211231")
	Pkg2021s = astmodel.MakeStoragePackageReference(Pkg2021)

	// Reference Packages for 2022
	Pkg2022  = MakeLocalPackageReference(Group, "v20220630")
	Pkg2022s = astmodel.MakeStoragePackageReference(Pkg2022)

	// Reference Package Batch 2020
	BatchPkg2020 = MakeLocalPackageReference(BatchGroup, "v20200101")

	// Reference Package Batch 2021
	BatchPkg2021 = MakeLocalPackageReference(BatchGroup, "v20210630")

	// Objects in Pkg2021

	Address2021 = CreateObjectDefinition(Pkg2021, "Address", FullAddressProperty, CityProperty)

	// Reusable Properties - only in Pkg2021

	ResidentialAddress2021 = astmodel.NewPropertyDefinition("ResidentialAddress", "residentialAddress", Address2021.Name())

	PostalAddress2021 = astmodel.NewPropertyDefinition("PostalAddress", "postalAddress", Address2021.Name())
)

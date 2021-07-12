/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package test

import (
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

/*
 * Shared building blocks for testing
 */

var (
	// Common group for testing
	Group = "microsoft.person"

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

	// Reference Package 2020
	Pkg2020 = MakeLocalPackageReference(Group, "v20200101")

	// Reference Package 2021 preview
	Pkg2021Preview = MakeLocalPackageReference(Group, "v20211231preview")

	// Reference Package 2021
	Pkg2021 = MakeLocalPackageReference(Group, "v20211231")

	// Reference Package 2022
	Pkg2022 = MakeLocalPackageReference(Group, "v20220630")

	// Objects in Pkg2021

	Address2021 = CreateObjectDefinition(Pkg2021, "Address", FullAddressProperty, CityProperty)

	// Reusable Properties - only in Pkg2021

	ResidentialAddress2021 = astmodel.NewPropertyDefinition("ResidentialAddress", "residentialAddress", Address2021.Name())

	PostalAddress2021 = astmodel.NewPropertyDefinition("PostalAddress", "postalAddress", Address2021.Name())
)

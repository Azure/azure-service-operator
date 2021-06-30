/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

/*
 * Shared building blocks for testing
 */

var (
	// Common group for testing
	testGroup = "microsoft.person"

	// Reusable Properties - any package version

	fullNameProperty = astmodel.NewPropertyDefinition("FullName", "fullName", astmodel.StringType).
				WithDescription("As would be used to address mail")

	familyNameProperty = astmodel.NewPropertyDefinition("FamilyName", "familyName", astmodel.StringType).
				WithDescription("Shared name of the family")

	knownAsProperty = astmodel.NewPropertyDefinition("KnownAs", "knownAs", astmodel.StringType).
			WithDescription("How the person is generally known")

	fullAddressProperty = astmodel.NewPropertyDefinition("FullAddress", "fullAddress", astmodel.StringType).
				WithDescription("Full written address for map or postal use")

	cityProperty = astmodel.NewPropertyDefinition("City", "city", astmodel.StringType).
			WithDescription("City or town (or nearest)")

	// Reference Package 2020
	pkg2020 = test.MakeLocalPackageReference(testGroup, "v20200101")

	// Reference Package 2021
	pkg2021 = test.MakeLocalPackageReference(testGroup, "v20211231")

	// Objects in pkg2021

	address2021 = test.CreateObjectDefinition(pkg2021, "Address", fullAddressProperty, cityProperty)

	// Reusable Properties - only in pkg2021

	residentialAddress2021 = astmodel.NewPropertyDefinition("ResidentialAddress", "residentialAddress", address2021.Name())

	postalAddress2021 = astmodel.NewPropertyDefinition("PostalAddress", "postalAddress", address2021.Name())
)

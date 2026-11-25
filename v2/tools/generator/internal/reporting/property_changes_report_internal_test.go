/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestFormatAddedProperties_ReturnsAlphabeticalPrefixedNames(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	rows := []propertyChangeRow{
		{
			nextName: "KnownAs",
			statuses: []changeStatus{statusNew},
		},
		{
			thisName: "Unchanged",
			nextName: "Unchanged",
		},
		{
			nextName: "FullName",
			statuses: []changeStatus{statusNew},
		},
		{
			thisName: "Retired",
			statuses: []changeStatus{statusRetired},
		},
	}

	g.Expect(formatAddedProperties(rows)).To(gomega.Equal("+FullName +KnownAs"))
}

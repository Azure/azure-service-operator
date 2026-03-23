/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v4

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestUrlPath_GivenURLWithQueryParams_ReturnsPathOnly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect(urlPath("https://management.azure.com/subscriptions/sub1/resourceGroups/rg1?api-version=2021-04-01")).
		To(Equal("/subscriptions/sub1/resourceGroups/rg1"))
}

func TestUrlPath_GivenURLWithoutQueryParams_ReturnsPathOnly(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect(urlPath("https://management.azure.com/subscriptions/sub1/resourceGroups/rg1")).
		To(Equal("/subscriptions/sub1/resourceGroups/rg1"))
}

func TestUrlPath_GivenRelativePath_ReturnsPath(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect(urlPath("/subscriptions/sub1/resourceGroups/rg1?api-version=2021-04-01")).
		To(Equal("/subscriptions/sub1/resourceGroups/rg1"))
}

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
)

func Test_IsRetryableDNSZoneRecordError_NotFound(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	err := genericarmclient.NewTestCloudError("BadRequest", "The referenced resource '/subscriptions/xxx/resourceGroups/xxx/providers/Microsoft.Network/trafficManagerProfiles/test' was not found")
	g.Expect(isRetryableDNSZoneRecordError(err)).To(BeTrue())
}

func Test_IsRetryableDNSZoneRecordError_BadRequest(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	err := genericarmclient.NewTestCloudError("BadRequest", "Invalid value specified for field")
	g.Expect(isRetryableDNSZoneRecordError(err)).To(BeFalse())
}

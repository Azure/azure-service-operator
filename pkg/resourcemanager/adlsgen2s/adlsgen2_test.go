package adlsgen2s

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"context"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/adlsgen2s"
	"github.com/Azure/azure-service-operator/pkg/helpers"

)

var _ = Describe("ADLS Gen2", func() {

	const timeout = time.Second * 180

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	Context("No error on Create Data Lake Instances", func() {
		It("should create and delete data lake in azure", func() {
			name := "tdevadls" + helpers.RandomString(10)
			adlsgen2Manager := tc.AdlsGen2Managers

			var err error

			_, err = adlsgen2Manager.AdlsGen2s.Create()

			Expect(err).NotTo(HaveOccurred())

			
		})
	})
})

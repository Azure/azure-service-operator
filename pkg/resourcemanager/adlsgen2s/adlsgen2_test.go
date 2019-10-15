package adlsgen2s

import (
	"context"
	"fmt"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	// "github.com/Azure/go-autorest/autorest/to"
	"time"
	"net/http"

	// "github.com/Azure/azure-service-operator/pkg/resourcemanager/adlsgen2s"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ADLS Gen2", func() {

	const timeout = time.Second * 180

	datalakeName := "tdevadls" + helpers.RandomString(10)

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
		// TODO: take this out
		_, _ = tc.DataLakeManagers.Storage.DeleteAdlsGen2(context.Background(), tc.ResourceGroupName, datalakeName)
	})

	Context("Create and Delete Azure Datalake Instances", func() {
		It("should create and delete a datalake in azure", func() {

			// Create a data lake enbaled storage account
			adlsLocation := config.DefaultLocation()
			_, err := tc.DataLakeManagers.Storage.CreateAdlsGen2(context.Background(), tc.ResourceGroupName, datalakeName, adlsLocation, azurev1alpha1.StorageSku{
				Name: "Standard_LRS",
			}, "StorageV2", map[string]*string{}, "", nil)

			if err != nil {
				fmt.Println("data lake wasn't created")
			}

			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := tc.DataLakeManagers.Storage.GetAdlsGen2(context.Background(), tc.ResourceGroupName, datalakeName)
				return result.Response.StatusCode == http.StatusOK
			}, timeout,
			).Should(BeTrue())

			_, err = tc.DataLakeManagers.Storage.DeleteAdlsGen2(context.Background(), tc.ResourceGroupName, datalakeName)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := tc.DataLakeManagers.Storage.GetAdlsGen2(context.Background(), tc.ResourceGroupName, datalakeName)
				return result.Response.StatusCode == http.StatusNotFound
			}, timeout,
			).Should(BeTrue())
		})
	})
})

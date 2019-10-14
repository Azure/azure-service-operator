package adlsgen2s

import (
	"time"

	"context"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/go-autorest/autorest/to"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	// "github.com/Azure/azure-service-operator/pkg/resourcemanager/adlsgen2s"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ADLS Gen2", func() {

	const timeout = time.Second * 180

	adlsName := "tdevadls" + helpers.RandomString(10)

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		// Create a data lake enbaled storage account
		adlsLocation := config.DefaultLocation()
		_, _ = tc.StorageManagers.Storage.CreateStorage(context.Background(), tc.ResourceGroupName, adlsName, adlsLocation, azurev1alpha1.StorageSku{
			Name: "Standard_LRS",
		}, "Storage", map[string]*string{}, "", nil,  to.BoolPtr(true))
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
		_, _ = tc.StorageManagers.Storage.DeleteStorage(context.Background(), tc.ResourceGroupName, adlsName)
	})

	Context("No error on Create File System Instances", func() {
		It("should create and delete filesystems in azure data lake", func() {
			adlsLocation := config.DefaultLocation()
			adlsgen2Manager := tc.FileSystemManager

			var err error

			_, err = adlsgen2Manager.AdlsGen2s.Create(context.Background(), tc.ResourceGroupName, adlsName, adlsLocation, )

			Expect(err).NotTo(HaveOccurred())

		})
	})
})

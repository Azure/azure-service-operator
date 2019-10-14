package adlsgen2s

import (
	"time"

	"context"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	// "github.com/Azure/go-autorest/autorest/to"

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
		_, _ = tc.DataLakeManagers.Storage.CreateAdlsGen2(context.Background(), tc.ResourceGroupName, adlsName, adlsLocation, azurev1alpha1.StorageSku{
			Name: "Standard_LRS",
		}, "Storage", map[string]*string{}, "", nil)
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
		_, _ = tc.DataLakeManagers.Storage.DeleteAdlsGen2(context.Background(), tc.ResourceGroupName, adlsName)
	})

	Context("No error on Create File System Instances", func() {
		It("should create and delete filesystems in azure data lake", func() {
			fileSystemManager := tc.DataLakeManagers.FileSystem
			fileSystemName := "tfilesystem" + helpers.RandomString(5)
			xMsProperties := ""
			xMsClientRequestID := ""
			xMsDate := ""

			var err error

			_, err = fileSystemManager.CreateFileSystem(context.Background(), fileSystemName, xMsProperties, xMsClientRequestID, nil, xMsDate, adlsName)

			Expect(err).NotTo(HaveOccurred())

		})
	})
})

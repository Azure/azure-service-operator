package adlsgen2s

import (
	"context"
	"fmt"
	// "net/http"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/go-autorest/autorest/to"
	"time"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("File System", func() {

	const timeout = time.Second * 180

	adlsName := "tdevadls" + helpers.RandomString(10)

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		// Create a data lake enbaled storage account
		adlsLocation := config.DefaultLocation()
		_, err := tc.DataLakeManagers.Storage.CreateAdlsGen2(context.Background(), tc.ResourceGroupName, adlsName, adlsLocation, azurev1alpha1.StorageSku{
			Name: "Standard_LRS",
		}, map[string]*string{}, "", nil)

		if err != nil {
			fmt.Println("data lake wasn't created")
		}
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
			xMsDate := time.Now().String()

			var err error

			_, err = fileSystemManager.CreateFileSystem(context.Background(), tc.ResourceGroupName, fileSystemName, xMsProperties, xMsClientRequestID, to.Int32Ptr(20), xMsDate, adlsName)
			Expect(err).NotTo(HaveOccurred())

			// Eventually(func() bool {
			// 	result, _ := fileSystemManager.GetFileSystem(context.Background(), fileSystemName, xMsClientRequestID, xMsDate, adlsName)
			// 	return result.Response.StatusCode == http.StatusOK
			// }, timeout,
			// ).Should(BeTrue())

			// _, err = fileSystemManager.DeleteFileSystem(context.Background(), fileSystemName, xMsClientRequestID, xMsDate, adlsName)
			// Expect(err).NotTo(HaveOccurred())

			// Eventually(func() bool {
			// 	result, _ := fileSystemManager.GetFileSystem(context.Background(), fileSystemName, xMsClientRequestID, xMsDate, adlsName)
			// 	return result.Response.StatusCode == http.StatusNotFound
			// }, timeout,
			// ).Should(BeTrue())

		})
	})
})

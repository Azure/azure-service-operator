package adlsgen2s

import (
	"context"
	"fmt"
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
			xMsClientRequestID := "false"
			t := time.Now()
			xMsDate := t.String()

			var err error

			_, err = fileSystemManager.CreateFileSystem(context.Background(), fileSystemName, xMsProperties, xMsClientRequestID, to.Int32Ptr(100), xMsDate, adlsName)
			fmt.Println(err.Error())
			Expect(err).NotTo(HaveOccurred())

		})
	})
})

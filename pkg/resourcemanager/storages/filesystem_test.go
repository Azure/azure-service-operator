package storages

import (
	"context"
	"net/http"

	"time"

	s "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("File System", func() {

	const timeout = time.Second * 180

	adlsName := "tdevadls" + helpers.RandomString(10)

	ctx := context.Background()

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		// Create a data lake enbaled storage account
		adlsLocation := config.DefaultLocation()
		_, _ = tc.StorageManagers.Storage.CreateStorage(ctx, tc.ResourceGroupName, adlsName, adlsLocation, azurev1alpha1.StorageSku{
			Name: "Standard_LRS",
		}, "StorageV2", map[string]*string{}, "", nil, to.BoolPtr(true))

		Eventually(func() s.ProvisioningState {
			result, _ := tc.StorageManagers.Storage.GetStorage(ctx, tc.ResourceGroupName, adlsName)
			return result.ProvisioningState
		}, tc.timeout, tc.retryInterval,
		).Should(Equal(s.Succeeded))
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
		_, err := tc.StorageManagers.Storage.DeleteStorage(context.Background(), tc.ResourceGroupName, adlsName)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("Create and Delete File System Instances", func() {
		It("should create and delete filesystems in azure data lake", func() {

			fileSystemManager := tc.StorageManagers.FileSystem
			fileSystemName := "tfilesystem" + helpers.RandomString(5)
			requestTimeout := to.Int32Ptr(20)
			xMsDate := time.Now().String()

			var err error

			_, err = fileSystemManager.CreateFileSystem(context.Background(), tc.ResourceGroupName, fileSystemName, requestTimeout, xMsDate, adlsName)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := fileSystemManager.GetFileSystem(context.Background(), tc.ResourceGroupName, fileSystemName, requestTimeout, xMsDate, adlsName)
				return result.Response.StatusCode == http.StatusOK
			}, timeout,
			).Should(BeTrue())

			_, err = fileSystemManager.DeleteFileSystem(context.Background(), tc.ResourceGroupName, fileSystemName, requestTimeout, xMsDate, adlsName)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := fileSystemManager.GetFileSystem(context.Background(), tc.ResourceGroupName, fileSystemName, requestTimeout, xMsDate, adlsName)
				return result.Response.StatusCode == http.StatusNotFound
			}, timeout,
			).Should(BeTrue())

		})
	})
})

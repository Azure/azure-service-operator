package keyvault_test

import (
	"testing"

	resourceconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func TestKeyvault(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keyvault Suite")
}

var _ = BeforeSuite(func(done Done) {
	logf.SetLogger(zap.LoggerTo(GinkgoWriter, true))

	By("BeforeSuite - KeyVault Suite")

	resourceconfig.ParseEnvironment()

	close(done)
}, 60)

var _ = AfterSuite(func() {
	//clean up the resources created for test

	By("AfterSuite - KeyVault Suite")

})

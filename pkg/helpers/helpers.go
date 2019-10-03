package helpers

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	log = ctrl.Log.WithName("helpers")
)

// KubernetesResourceName returns the resource name for other components
func KubernetesResourceName(name string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9_-]+")
	return reg.ReplaceAllString(name, "-")
}

func AzrueResourceGroupName(subscriptionID, clusterName, resourceType, name, namespace string) string {
	nameParts := []string{subscriptionID, clusterName, resourceType, name, namespace}
	nameString := strings.Join(nameParts, "-")
	log.V(1).Info("Getting Azure Resource Group Name", "nameString", nameString)
	hash := md5.New()
	io.WriteString(hash, nameString)
	return fmt.Sprintf("aso-%x", hash.Sum(nil))
}

func IgnoreKubernetesResourceNotFound(err error) error {
	return client.IgnoreNotFound(err)
}

func IgnoreAzureResourceNotFound(err error) error {
	if err.(autorest.DetailedError).StatusCode.(int) == http.StatusNotFound {
		return nil
	}
	return err
}

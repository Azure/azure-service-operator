package helpers

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/mitchellh/hashstructure"

	"github.com/Azure/go-autorest/autorest"
	"github.com/sethvargo/go-password/password"
	"k8s.io/apimachinery/pkg/runtime"
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

// GenerateRandomUsername - helper function to generate random username for sql server
func GenerateRandomUsername(n int, numOfDigits int) (string, error) {

	// Generate a username that is n characters long, with n/2 digits and 0 symbols (not allowed),
	// allowing only lower case letters (upper case not allowed), and disallowing repeat characters.
	res, err := password.Generate(n, numOfDigits, 0, true, false)
	if err != nil {
		return "", err
	}

	return res, nil
}

// GenerateRandomPassword - helper function to generate random password for sql server
func GenerateRandomPassword(n int) (string, error) {

	// Math - Generate a password where: 1/3 of the # of chars are digits, 1/3 of the # of chars are symbols,
	// and the remaining 1/3 is a mix of upper- and lower-case letters
	digits := n / 3
	symbols := n / 3

	// Generate a password that is n characters long, with # of digits and symbols described above,
	// allowing upper and lower case letters, and disallowing repeat characters.
	res, err := password.Generate(n, digits, symbols, false, false)
	if err != nil {
		return "", err
	}

	return res, nil
}

// GetKubernetesObjectHash - helper function that generates a unique hash for a Kubernetes runtime.Object resource
func GetKubernetesObjectHash(obj runtime.Object) (uint64, error) {
	unstructured, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return 0, err
	}

	hash, err := hashstructure.Hash(unstructured["spec"].(map[string]interface{}), nil)
	if err != nil {
		return 0, err
	}

	return hash, nil
}

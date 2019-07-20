package helpers

import (
	"crypto/md5"
	"fmt"
	"io"
	"regexp"
	"strings"

	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	log = ctrl.Log.WithName("helpers")
)

func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func RemoveString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

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

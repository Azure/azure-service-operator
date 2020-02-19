package helpers

import (
	"strconv"
	"strings"

	"github.com/Azure/azure-service-operator/pkg/config"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateExternalNameService will create a Kubernetes Servic Using ExternalName types
func CreateExternalNameService(resource interface{}, svcName string, svcNamespace string, externalNameTemplate string, svcPortTemplate string) string {
	externalName, err := Templatize(externalNameTemplate, Data{Obj: resource})
	if err != nil {
		log.Error(err, "error parsing external name template")
		return ""
	}

	svcPortString, err := Templatize(svcPortTemplate, Data{Obj: resource})
	if err != nil {
		log.Error(err, "error parsing service port template")
		return ""
	}

	svcPortStripSlash := strings.Replace(svcPortString, "\\", "", -1)

	svcPortInt64, err := strconv.ParseInt(svcPortStripSlash, 0, 16)
	if err != nil {
		log.Error(err, "error converting service port template string to int")
		return ""
	}

	// ParseInt only returns an int64, must convert to int32 for azurev1alpha1.ServicePort field
	svcPort := int32(svcPortInt64)

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: KubernetesResourceName(svcName),
		},
		Spec: corev1.ServiceSpec{
			Type:         corev1.ServiceTypeExternalName,
			ExternalName: externalName,
			Ports: []corev1.ServicePort{
				corev1.ServicePort{
					Port: svcPort,
				},
			},
		},
	}

	newService, err := config.Instance.KubeClientset.CoreV1().Services(svcNamespace).Create(service)
	if err != nil {
		log.Error(err, "error creating service")
	}
	return newService.Name
}

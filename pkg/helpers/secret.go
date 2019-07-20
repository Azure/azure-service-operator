package helpers

import (
	"github.com/Azure/azure-service-operator/pkg/config"
	apiv1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateSecret(resource interface{}, svcName, svcNamespace string, secretTemplate map[string]string) string {
	data := map[string]string{}
	for key, value := range secretTemplate {
		tempValue, err := Templatize(value, Data{Obj: resource})
		if err != nil {
			log.Error(err, "error parsing config map template")
			return ""
		}
		data[key] = tempValue
	}

	secretName := KubernetesResourceName(svcName)
	secretObj := &apiv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: svcNamespace,
		},
		StringData: data,
	}

	_, err := config.Instance.KubeClientset.CoreV1().Secrets(svcNamespace).Get(secretName, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		log.Info("Creating Secret", "Secret.Name", secretName)
		_, err := config.Instance.KubeClientset.CoreV1().Secrets(svcNamespace).Create(secretObj)
		if err != nil {
			log.Error(err, "error creating Secret")
		}
	} else {
		log.Info("Updating Secret", "Secret.Name", secretName)
		_, err := config.Instance.KubeClientset.CoreV1().Secrets(svcNamespace).Update(secretObj)
		if err != nil {
			log.Error(err, "error updating Secret")
		}
	}

	return secretName
}

func DeleteSecret(svcName, svcNamespace string) error {
	secretName := KubernetesResourceName(svcName)
	return config.Instance.KubeClientset.CoreV1().Secrets(svcNamespace).Delete(secretName, &metav1.DeleteOptions{})
}

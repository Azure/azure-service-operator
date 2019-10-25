package eventhub

import (
	"context"
	"github.com/Azure/azure-service-operator/controller_refactor"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	v1 "k8s.io/api/core/v1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type secretsWriter struct {
	*controller_refactor.GenericController
	eventHubManager eventhubs.EventHubManager
}

func (writer *secretsWriter) Run(ctx context.Context, r runtime.Object) error {
	instance, err := convertInstance(r)
	if err != nil {
		return err
	}

	eventhubName := instance.ObjectMeta.Name

	spec := instance.Spec
	eventhubNamespace := instance.Spec.Namespace

	secretName := instance.Spec.SecretName
	if len(secretName) == 0 {
		secretName = eventhubName
	}

	authorizationRuleName := spec.AuthorizationRule.Name

	manager := writer.eventHubManager
	result, err := manager.ListKeys(ctx, spec.ResourceGroup, spec.Namespace, eventhubName, authorizationRuleName)
	if err != nil {
		//log error and kill it
		writer.Recorder.Event(instance, "Warning", "Failed", "Unable to list keys")
	}

	csecret := &v1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "apps/v1beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: instance.Namespace,
		},
		Data: map[string][]byte{
			"primaryconnectionstring":   []byte(*result.PrimaryConnectionString),
			"secondaryconnectionstring": []byte(*result.SecondaryConnectionString),
			"primaryKey":                []byte(*result.PrimaryKey),
			"secondaryKey":              []byte(*result.SecondaryKey),
			"sharedaccesskey":           []byte(authorizationRuleName),
			"eventhubnamespace":         []byte(eventhubNamespace),
			"eventhubName":              []byte(eventhubName),
		},
		Type: "Opaque",
	}

	_, err = controllerutil.CreateOrUpdate(context.Background(), writer.KubeClient, csecret, func() error {
		writer.Log.Info("mutating secret bundle", "Secret", csecret.Name)
		innerErr := controllerutil.SetControllerReference(instance, csecret, writer.Scheme)
		if innerErr != nil {
			return innerErr
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

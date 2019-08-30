package config

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest/azure"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var Instance *Config

type Config struct {
	KubeClientset     kubernetes.Interface
	Resources         map[string]bool `json:"resources"`
	ClusterName       string          `json:"clusterName"`
	CloudName         string          `json:"cloudName"`
	TenantID          string          `json:"tenantID"`
	SubscriptionID    string          `json:"subscriptionID"`
	ClientID          string          `json:"clientID"`
	ClientSecret      string          `json:"clientSecret"`
	UseAADPodIdentity bool            `json:"useAADPodIdentity"`
}

func getKubeconfig(masterURL, kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	}
	return rest.InClusterConfig()
}

func CreateKubeClientset(masterURL, kubeconfig string) (kubernetes.Interface, error) {
	config, err := getKubeconfig(masterURL, kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("failed to get k8s config. %+v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to get k8s client. %+v", err)
	}

	return clientset, nil
}

// Environment() returns an `azure.Environment{...}` for the current cloud.
func Environment() azure.Environment {
	cloudName := Instance.CloudName
	env, err := azure.EnvironmentFromName(cloudName)
	if err != nil {
		panic(fmt.Sprintf(
			"invalid cloud name '%s' specified, cannot continue\n", cloudName))
	}
	return env
}

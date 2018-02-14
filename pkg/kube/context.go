package kube

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	pf "k8s.io/helm/pkg/helm/portforwarder"
)

// getKubeClient creates a Kubernetes config and client for a given kubeconfig context.
func getKubeClient(context string) (kubernetes.Interface, *rest.Config, error) {
	_, config, err := configForContext(context)
	if err != nil {
		return nil, nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, fmt.Errorf("could not get Kubernetes client: %s", err)
	}
	return client, config, nil
}

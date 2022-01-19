package util

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type PrimaryClientSet struct {
	restConfig *rest.Config
}

func NewPrimaryClientSet(config *rest.Config) (*kubernetes.Clientset, error) {
	primaryClientSet := PrimaryClientSet{restConfig: config}
	return NewPrimaryClientSetFromConfig(primaryClientSet)
}

func NewPrimaryClientSetFromConfig(c PrimaryClientSet) (*kubernetes.Clientset, error) {
	clientset, err := kubernetes.NewForConfig(c.restConfig)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

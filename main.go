package main

import (
	"flag"
	"k8s-client-demo/demo"
	"os"
	"path/filepath"

	"github.com/emqx/emqx-operator/api/v1beta1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubectl/pkg/scheme"
)

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}

func init() {
	v1beta1.AddToScheme(scheme.Scheme)
}

func main() {
	var kubeconfig *string

	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// demo.DemoForNamespace(config)

	demo.DemoForEmqxBroker(config, "default")

}

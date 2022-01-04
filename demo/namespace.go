package demo

import (
	"context"
	"fmt"
	"k8s-client-demo/resource"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func DemoForNamespace(config *rest.Config) {
	// Create namespace client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	namespacesClient := clientset.CoreV1().Namespaces()

	// >>>>>> Create namespace: test
	fmt.Println("[> create namespace: test")

	nsForCreate, err := namespacesClient.Create(context.TODO(), resource.GenerateNamespace("test"), metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("create namespace: %+v\n", nsForCreate)

	// >>>>> Get namespace instance by name
	Prompt()
	fmt.Println("[> get namespace: test")
	ns, err := namespacesClient.Get(context.TODO(), "test", metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("ns found: %+v\n", ns)

	// Get namepace list
	Prompt()
	fmt.Println("[> get namespace list")
	nsForList, err := namespacesClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("found namespace list :%+v\n", nsForList)

	// Delete namesapce by name: test
	Prompt()
	fmt.Println("[> delete namespace: test")
	err = namespacesClient.Delete(context.TODO(), "test", metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Delete namespace successfully")
}

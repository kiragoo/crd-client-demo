package demo

import (
	"context"
	"fmt"
	"k8s-client-demo/pkg"
	"k8s-client-demo/resource"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func DemoForEmqxBroker(config *rest.Config, ns string) {
	// Create emqxbroker restclient
	clientset, err := pkg.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	emqxbrokerClient := clientset.EmqxBrokersV1Beta2().EmqxBrokers(ns)

	// Create emqxbroker instance
	Prompt()
	fmt.Println("[> create emqxbroker")
	emqxbroker, err := emqxbrokerClient.Create(context.TODO(), resource.GenerateEmqxbroker(ns), metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("create emqxbroker: %+v\n", emqxbroker)

	// Get emqxbroker instance
	Prompt()
	eb, err := emqxbrokerClient.Get(context.TODO(), "emqx", metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("emqxbroker found: %+v\n", eb)

	// Get emqxbroker list
	Prompt()
	eblist, err := emqxbrokerClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("emqxbroker list: %+v\n", eblist)

	// Delete emqxbroker instance
	Prompt()
	err = emqxbrokerClient.Delete(context.TODO(), "emqx", metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Delete emqxbroker successfully")
}

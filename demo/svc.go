package demo

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func DemoForSvc(config *rest.Config) {
	// Create namespace client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	svcsClient := clientset.CoreV1().Services("emqx-1")

	fmt.Println("[>Get svc: emqx-1-emqx-ee]")
	Prompt()

	svc, err := svcsClient.Get(context.TODO(), "emqx-1-emqx-ee", metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("svc found: %+v\n", svc.Spec)

}

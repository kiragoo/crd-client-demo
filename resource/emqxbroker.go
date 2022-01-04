package resource

import (
	"github.com/emqx/emqx-operator/api/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GenerateEmqxbroker(ns string) *v1beta1.EmqxBroker {
	return &v1beta1.EmqxBroker{
		ObjectMeta: metav1.ObjectMeta{
			Name: "emqx",
		},
		Spec: v1beta1.EmqxBrokerSpec{
			ServiceAccountName: "emqx",
			Image:              "emqx/emqx:4.3.10",
			Replicas:           Int32Ptr(3),
			Labels: map[string]string{
				"cluster": "emqx",
			},
			// Listener: v1beta1.Listener{
			// 	Type: corev1.ServiceTypeLoadBalancer,
			// 	Ports: v1beta1.Ports{
			// 		MQTT:      1883,
			// 		MQTTS:     8883,
			// 		WS:        8083,
			// 		WSS:       8084,
			// 		Dashboard: 18083,
			// 		API:       8081,
			// 	},
			// },
		},
	}
}

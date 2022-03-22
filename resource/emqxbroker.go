package resource

import (
	"github.com/emqx/emqx-operator/apis/apps/v1beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GenerateEmqxbroker(ns string) *v1beta2.EmqxBroker {
	return &v1beta2.EmqxBroker{
		ObjectMeta: metav1.ObjectMeta{
			Name: "emqx",
		},
		Spec: v1beta2.EmqxBrokerSpec{
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

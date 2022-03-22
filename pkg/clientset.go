package pkg

import (
	"fmt"
	v1beta2 "k8s-client-demo/pkg/clientset/v1beta2"
	"net/http"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	EmqxBrokerV1Beta2Interface() v1beta2.EmqxBrokerInterface
}

type Clientset struct {
	emqxbrokerv1beta2 *v1beta2.EmqxBrokerV1Beta1Client
}

func (c *Clientset) EmqxBrokersV1Beta2() v1beta2.EmqxBrokerV1Beta1Interface {
	return c.emqxbrokerv1beta2
}

func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.emqxbrokerv1beta2, err = v1beta2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

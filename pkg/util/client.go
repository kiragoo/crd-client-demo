package util

import "k8s.io/client-go/rest"

// create the admin serviceaccount
// kubectl create sa root
// kubectl -n default create clusterrolebinding root --clusterrole cluster-admin --serviceaccount=default:root

// the token get in secret should be decode by base64
// Usage:
//      kubectl get secret xxx -o jsonpath={.data.token}|base64 -d
func NewRestConfigFromToken(url, token string) *rest.Config {
	return &rest.Config{
		Host: url,
		ContentConfig: rest.ContentConfig{
			AcceptContentTypes: "application/json",
			ContentType:        "application/json",
		},
		BearerToken: token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
}

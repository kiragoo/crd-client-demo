package main

import (
	"k8s-client-demo/demo"
	"k8s-client-demo/pkg/util"

	"github.com/emqx/emqx-operator/apis/apps/v1beta2"
	"k8s.io/kubectl/pkg/scheme"
)

func init() {
	v1beta2.AddToScheme(scheme.Scheme)
}

func main() {

	config := util.NewRestConfigFromToken(
		"https://127.0.0.1:61290",
		"eyJhbGciOiJSUzI1NiIsImtpZCI6Im12MmVoU0hpLVdyeVJMX2RhR3ItMGxOelpRS0NkTW5sbGY0T2JJeEgxTzAifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6InJvb3QtdG9rZW4tdG56dGoiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoicm9vdCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjI0OTZmODRhLWQ3N2UtNDA1MS05MWNhLWE5ZTYxMjIzMGVmMSIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZWZhdWx0OnJvb3QifQ.bKusaPJJG0vRwSa38ulIiv-D0uchTSYhwpbIseQPW1vJtzJAfEWMzq6YHh_Bg1EBQT2OxBGh73fUhK_w_iBAr5sL2sJcrgWOXyvqNMqulP86bZX4yO6HaQqvMExIPU5Qkj3JOjC6rwAUzLhazmZwsfjDgO_rpup3QQxQFU9ElPf6uBDnlZvDCndRx7SJFDrj-kvDE-D_Dwdlq0Q3hnj0yASM5i_odLvkiYxqP7LwJrohcja2fZgJDgtyT3XqSfBjfDW-UMPyF6JFuiiSNur--XRhtudvCjPwrhRJcI6zjzNljdXbQ4YdDoNfv0QHb3AwCuANZVlJI8FMMyYF2e5S5w",
	)
	// config := config.NewConfigFromKubeConfig()
	// demo.DemoForNamespace(config)

	// demo.DemoForEmqxBroker(config, "default")
	demo.DemoForSvc(config)

}

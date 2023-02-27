package main

import (
	"context"
	"flag"
	"path/filepath"

	policyv1alpha1 "github.com/linkerd/linkerd2/controller/gen/apis/policy/v1alpha1"
	linkerdclientset "github.com/linkerd/linkerd2/controller/gen/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	gatewayv1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

func main() {
	clientset := mustSetupClient()

	ap := &policyv1alpha1.AuthorizationPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "example",
			Namespace: "default",
		},
		Spec: policyv1alpha1.AuthorizationPolicySpec{
			TargetRef: gatewayv1alpha2.PolicyTargetReference{
				Kind: "Namespace",
				Name: "default",
			},
		},
	}

	_, err := clientset.PolicyV1alpha1().AuthorizationPolicies(ap.Namespace).Create(context.Background(), ap, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
}

func mustSetupClient() *linkerdclientset.Clientset {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	return linkerdclientset.NewForConfigOrDie(config)
}

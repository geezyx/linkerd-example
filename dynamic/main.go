package main

import (
	"context"
	"flag"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	dynamicClient := mustSetupClient()

	gvr := schema.GroupVersionResource{
		Group:    "policy.linkerd.io",
		Version:  "v1alpha1",
		Resource: "authorizationpolicies",
	}

	gvk := schema.GroupVersionKind{
		Group:   "policy.linkerd.io",
		Version: "v1alpha1",
		Kind:    "AuthorizationPolicy",
	}

	ap := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"spec": map[string]interface{}{
				"requiredAuthenticationRefs": []interface{}{},
				"targetRef": map[string]interface{}{
					"kind": "Namespace",
					"name": "default",
				},
			},
		},
	}
	ap.SetName("example")
	ap.SetGroupVersionKind(gvk)

	_, err := dynamicClient.Resource(gvr).Namespace("default").Create(context.Background(), ap, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
}

func mustSetupClient() dynamic.Interface {
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
	return dynamic.NewForConfigOrDie(config)
}

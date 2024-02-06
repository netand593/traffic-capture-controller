package main

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"

	trafficcaptureclient "github.com/netand593/traffic-capture-controller/pkg/controller"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	trafficCaptureClient, err := trafficcaptureclient.NewForConfig(config)
	if err != nil {
		klog.Fatalf("Error building traffic capture clientset: %s", err.Error())
	}

	controller := NewController(clientset, trafficCaptureClient)

	stopCh := make(chan struct{})
	defer close(stopCh)

	go controller.Run(stopCh)

	// Wait forever
	select {}
}


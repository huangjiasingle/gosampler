package client

import (
	"fmt"
	"log"
	"net"
	"os"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var Client *Clientset

func init() {
	Client = NewClient()
}

func GetClient() *Clientset {
	return Client
}

type Clientset struct {
	Clientset kubernetes.Interface
	Dynamic   dynamic.Interface
	Config    *rest.Config
}

func NewClient() *Clientset {
	//create k8s clientset and config
	client, config := mustNewKubeClientAndConfig()
	//create k8s dynamic client
	dynamicInterface, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalf("init k8s client err: %v", err)
	}
	return &Clientset{
		Clientset: client,
		Dynamic:   dynamicInterface,
		Config:    config,
	}
}

func mustNewKubeClientAndConfig() (kubernetes.Interface, *rest.Config) {
	var cfg *rest.Config
	var err error
	if os.Getenv("KUBE-CONFIG") != "" {
		cfg, err = outOfClusterConfig()
	} else if lookupConfig() {
		cfg, err = outOfClusterConfig()
	} else {
		cfg, err = inClusterConfig()
	}
	if err != nil {
		panic(err)
	}
	return kubernetes.NewForConfigOrDie(cfg), cfg
}

func inClusterConfig() (*rest.Config, error) {
	if len(os.Getenv("KUBERNETES_SERVICE_HOST")) == 0 {
		addrs, err := net.LookupHost("kubernetes.default.svc")
		if err != nil {
			return nil, err
		}
		os.Setenv("KUBERNETES_SERVICE_HOST", addrs[0])
	}
	if len(os.Getenv("KUBERNETES_SERVICE_PORT")) == 0 {
		os.Setenv("KUBERNETES_SERVICE_PORT", "443")
	}
	return rest.InClusterConfig()
}

func outOfClusterConfig() (*rest.Config, error) {
	kubeconfig := os.Getenv("KUBE-CONFIG")
	if kubeconfig == "" {
		if os.Getenv("HOME") != "" {
			kubeconfig = fmt.Sprintf("%v/.kube/config", os.Getenv("HOME"))
		} else {
			kubeconfig = fmt.Sprintf("%v/.kube/config", os.Getenv("USERPROFILE"))
		}
		_, err := os.Stat(kubeconfig)
		if err != nil {
			return nil, err
		}
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	return config, err
}

func lookupConfig() bool {
	var kubeconfig string
	if os.Getenv("HOME") != "" {
		kubeconfig = fmt.Sprintf("%v/.kube/config", os.Getenv("HOME"))
	} else {
		kubeconfig = fmt.Sprintf("%v/.kube/config", os.Getenv("USERPROFILE"))
	}
	_, err := os.Stat(kubeconfig)
	if err != nil {
		return false
	}
	return true
}

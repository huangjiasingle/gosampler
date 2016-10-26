package main

import (
	"flag"
	"fmt"
	// "time"

	// "common/jsonx"
	"k8s.io/client-go/1.4/kubernetes"
	"k8s.io/client-go/1.4/pkg/api"
	// "k8s.io/client-go/1.4/pkg/api/resource"
	// "k8s.io/client-go/1.4/pkg/api/unversioned"
	// "k8s.io/client-go/1.4/pkg/api/v1"
	"k8s.io/client-go/1.4/tools/clientcmd"
)

var (
	kubeconfig = flag.String("kubeconfig", "./config", "absolute path to the kubeconfig file")
)

func main() {
	flag.Parse()
	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("183.131.19.231:8080", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	/*//创建replicationController
	rcTypeMeta := unversioned.TypeMeta{Kind: "ReplicationController", APIVersion: "v1"}

	rcObjectMeta := v1.ObjectMeta{
		Name: "redis",
		Labels: map[string]string{
			"name": "redis",
		},
	}

	instanceCnt := int32(2)
	rcSpec := v1.ReplicationControllerSpec{
		Replicas: &instanceCnt,
		Selector: map[string]string{
			"name": "redis",
		},
		Template: &v1.PodTemplateSpec{
			v1.ObjectMeta{
				Labels: map[string]string{
					"name": "redis",
				},
			},
			v1.PodSpec{
				Containers: []v1.Container{
					v1.Container{
						Name:  "redis",
						Image: "redis",
						Ports: []v1.ContainerPort{
							v1.ContainerPort{
								ContainerPort: 6379,
								Protocol:      v1.ProtocolTCP,
							},
						},
						Resources: v1.ResourceRequirements{
							Requests: v1.ResourceList{
								v1.ResourceCPU:    resource.MustParse("100m"),
								v1.ResourceMemory: resource.MustParse("100Mi"),
							},
						},
					},
				},
			},
		},
	}

	rc := new(v1.ReplicationController)
	rc.TypeMeta = rcTypeMeta
	rc.ObjectMeta = rcObjectMeta
	rc.Spec = rcSpec

	//创建
	result, err := clientset.Core().ReplicationControllers("default").Create(rc)
	if err != nil {
		fmt.Printf("deploy application failed ,the reason is %s", err.Error())
		fmt.Println()
	}

	str, _ := jsonx.ToJson(result)

	fmt.Printf("the result is %v", str)
	fmt.Println()*/

	// err = clientset.Core().ReplicationControllers("default").Delete("redis", &api.DeleteOptions{
	// 	TypeMeta: unversioned.TypeMeta{
	// 		Kind:       "ReplicationController",
	// 		APIVersion: "v1",
	// 	},
	// 	OrphanDependents: new(bool),
	// })

	// if err != nil {
	// 	fmt.Printf("err is %v", err.Error())
	// }
	// for {
	pods, err := clientset.Core().Pods("default").List(api.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < len(pods.Items); i++ {

	}

	fmt.Printf("There are %#v pods in the cluster\n", pods)

	// 	str, _ := jsonx.ToJson(pods.Items)
	// 	fmt.Printf("pods in the cluster  is %v \n", str)

	// 	time.Sleep(10 * time.Second)
	// }
}

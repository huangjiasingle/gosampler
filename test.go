/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	// "encoding/json"
	"fmt"
	"time"

	// "k8s.io/client-go/1.4/kubernetes"
	// "k8s.io/client-go/1.4/pkg/api"
	// "k8s.io/client-go/1.4/pkg/api/resource"
	// // "k8s.io/client-go/1.4/pkg/api/unversioned"
	// "k8s.io/client-go/1.4/pkg/api/v1"
	"math/rand"
	// "k8s.io/client-go/1.4/rest"
)

func main() {
	/*//-----------------------------------------------------------------------------
	//创建pod
	typeMete := unversioned.TypeMeta{Kind: "pods", APIVersion: "v1"}

	objectMeta := v1.ObjectMeta{
		Name: "redis-master",
		Labels: map[string]string{
			"name": "redis-master",
		},
	}

	podSpec := v1.PodSpec{
		RestartPolicy: v1.RestartPolicyAlways,
		NodeSelector: map[string]string{
			"name": "redis-master",
		},
		Containers: []v1.Container{
			v1.Container{
				Name:    "hello",
				Image:   "ubuntu:14.04",
				Command: []string{"/bin/echo", "hello", "world"},
			},
		},
	}

	pod := new(v1.Pod)

	pod.TypeMeta = typeMete

	pod.ObjectMeta = objectMeta

	pod.Spec = podSpec

	by, _ := json.MarshalIndent(pod, " ", "  ")
	fmt.Printf("pod = %s", string(by))

	//----------------------------------------------------------------------------------
	//创建replicationController
	rcTypeMeta := unversioned.TypeMeta{Kind: "ReplicationController", APIVersion: "v1"}

	rcObjectMeta := v1.ObjectMeta{
		Name: "redis-master",
		Labels: map[string]string{
			"name": "redis-master",
		},
	}

	rcSpec := v1.ReplicationControllerSpec{
		Replicas: 1,
		Selector: map[string]string{
			"name": "redis-master",
		},
		Template: v1.PodTemplate{
			v1.ObjectMeta{
				Name: "redis-master",
				Labels: map[string]string{
					"name": "redis-master",
				},
			},
			v1.PodSpec{
				RestartPolicy: v1.RestartPolicyAlways,
				NodeSelector: map[string]string{
					"name": "redis-master",
				},
				Containers: []v1.Container{
					v1.Container{
						Name:  "master",
						Image: "redis",
						Ports: []v1.ContainerPort{
							ContainerPort: 6379,
							Protocol:      v1.ProtocolTCP,
						},
					},
				},
			},
		},
	}

	replicationController := new(v1.ReplicationController)*/

	// q := resource.MustParse("100m")
	// fmt.Printf("%#v", q)
	// fmt.Println()

	// fmt.Printf("%#v", v1.ResourceList{
	// 	v1.ResourceCPU:    resource.MustParse("100"),
	// 	v1.ResourceMemory: resource.MustParse("256"),
	// })

	rand.Seed(time.Now().UnixNano())

	x := rand.Intn(100)
	fmt.Println(x)
	a := new(int32)
	*a = 1
	fmt.Println(*a)
}

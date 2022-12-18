package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	v1 "k8s.io/api/apps/v1"
	//corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
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
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

//---------------------------------------------------------------

	b, err := ioutil.ReadFile("yamls/nginx.yaml")
	nginxDep := &v1.Deployment{}
	nginxJson, _ := yaml.ToJSON(b)
	if err = json.Unmarshal(nginxJson, nginxDep); err != nil {
		return
	}
	// Create Deployment
	fmt.Println("Creating deployment nginx...")
	deploymentList, err := clientset.AppsV1().Deployments("webserver").Create(context.Background(), nginxDep, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", deploymentList.GetObjectMeta().GetName())

//---------------------------------------------------------------

	//namespace := "webserver"
	//var replicas int32 = 1
	//deployment := &v1.Deployment{
	//	ObjectMeta: metav1.ObjectMeta{
	//		Name: "nginx",
	//		Labels: map[string]string{
	//			"app": "nginx",
	//			"env": "dev",
	//		},
	//	},
	//	Spec: v1.DeploymentSpec{
	//		Replicas: &replicas,
	//		Selector: &metav1.LabelSelector{
	//			MatchLabels: map[string]string{
	//				"app": "nginx",
	//				"env": "dev",
	//			},
	//		},
	//		Template: corev1.PodTemplateSpec{
	//			ObjectMeta: metav1.ObjectMeta{
	//				Name: "nginx",
	//				Labels: map[string]string{
	//					"app": "nginx",
	//					"env": "dev",
	//				},
	//			},
	//			Spec: corev1.PodSpec{
	//				Containers: []corev1.Container{
	//					{
	//						Name:  "nginx",
	//						Image: "nginx:1.16.1",
	//						Ports: []corev1.ContainerPort{
	//							{
	//								Name:          "http",
	//								Protocol:      corev1.ProtocolTCP,
	//								ContainerPort: 80,
	//							},
	//						},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	//// Create Deployment
	//fmt.Println("Creating deployment nginx...")
	//deploymentList, err := clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Created deployment %q.\n", deploymentList.GetObjectMeta().GetName())
}
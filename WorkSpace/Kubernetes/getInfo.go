package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//"k8s.io/client-go/util/homedir"
	"os/user"
)

func main() {

	home := GetHomePath()
	config, err := clientcmd.BuildConfigFromFlags("", fmt.Sprintf("%s/.kube/config", home))
	if err != nil {
		panic(err)
	}

	// 创建一个k8s客户端
	client, _ := kubernetes.NewForConfig(config)

	// 获取k8s集群版本
	k8sVersion, _ := client.ServerVersion()
	fmt.Println("Kubernetes集群版本是：",k8sVersion)

	// 查询k8s集群的节点信息，相当于命令：kubectl get nodes -o yaml, 如果没有管理员权限可能会失败，可以改成查Pods的接口
	// nodes, err := clientSet.CoreV1().Nodes().List(metav1.ListOptions{}) 老版本写法
	nodes, err := client.CoreV1().Nodes().List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		fmt.Printf("%v",err)
		return
	}

	for _,node := range nodes.Items {
		fmt.Println("Kubernetes集群节点是：",node.Name)
	}

	fmt.Printf("------------------------------------------\n")

	// 获取集群下的所有命名空间
	k8s_ns, _ := client.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})

	fmt.Printf("kubernetes集群命名空间如下：\n")
	for _, item := range k8s_ns.Items {
		fmt.Println(item.Name)

	}

	fmt.Printf("------------------------------------------\n")

	// 获取指定命名空间下的deployment信息
	deploymentList, err := client.AppsV1().Deployments("jobfun").List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _,k8s_deploy := range deploymentList.Items {
		fmt.Printf("命名空间是：%v \n deployment服务名字：%v \n 副本个数：%v \n\n",k8s_deploy.Namespace,k8s_deploy.Name,k8s_deploy.Status.Replicas)
	}

	fmt.Printf("------------------------------------------\n")

	// 获取指定命名空间下的pod信息
	pod, _ := client.CoreV1().Pods("jobfun").List(context.Background(), metav1.ListOptions{})
	for _, k8s_pod := range pod.Items {
		fmt.Printf("%v空间内的Pod有：\n %v \n\n",k8s_pod.Namespace,k8s_pod.Name)
	}
}

func GetHomePath() string {
	u , err := user.Current()
	if err == nil {
		return u.HomeDir
	}
	return ""
}

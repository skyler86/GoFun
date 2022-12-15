package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//	"k8s.io/client-go/util/homedir"
	"os/user"
)

//-----------------------------------------------------------------------------

func main() {
	// 读取配置文件
	home := GetHomePath()
	k8sConfig, err := clientcmd.BuildConfigFromFlags("", fmt.Sprintf("%s/.kube/config", home))  // 使用 kubectl 默认配置 ~/.kube/config
	if err != nil {
		fmt.Printf("%v",err)
		return
	}

	client, err := kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(client.ServerVersion())

	// 创建一个k8s客户端
	clientSet, err := kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		fmt.Printf("%v",err)
		return
	}

	// 查询k8s集群的节点信息，相当于命令：kubectl get nodes -o yaml, 如果没有管理员权限可能会失败，可以改成查Pods的接口
	// nodes, err := clientSet.CoreV1().Nodes().List(metav1.ListOptions{}) 老版本写法
	nodes, err := clientSet.CoreV1().Nodes().List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		fmt.Printf("%v",err)
		return
	}

	for _,node := range nodes.Items {
		fmt.Println(node.Name)
	}

}

func GetHomePath() string {
	u , err := user.Current()
	if err == nil {
		return u.HomeDir
	}
	return ""
}

//-----------------------------------------------------------------------------

//func main() {
//
//	home := GetHomePath()
//	config, err := clientcmd.BuildConfigFromFlags("", fmt.Sprintf("%s/.kube/config", home))
//	if err != nil {
//		panic(err)
//	}
//	client, _ := kubernetes.NewForConfig(config)
//	deploymentList, err := client.AppsV1().Deployments("kube-system").List(context.TODO(),metav1.ListOptions{})
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	for _,v := range deploymentList.Items {
//		fmt.Printf("命名空间是：%v\n deployment服务名字：%v\n 副本个数：%v\n\n",v.Namespace,v.Name,v.Status.Replicas)
//	}
//}
//
//func GetHomePath() string {
//	u , err := user.Current()
//	if err == nil {
//		return u.HomeDir
//	}
//	return ""
//}

//-----------------------------------------------------------------------------
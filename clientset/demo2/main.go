package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1、加载配置文件，生成config对象
	config, err := clientcmd.BuildConfigFromFlags("", "../../kubeconfig")
	if err != nil {
		panic(err.Error())
	}

	// 2、实例化ClientSet对象
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.
		CoreV1().                                  // 返回CoreV1Client实例
		Pods("kube-system").                       // 调用newPods函数，查询Pod列表
		List(context.TODO(), metav1.ListOptions{}) // 底层还是用restclient发起的请求
	if err != nil {
		panic(err.Error())
	}

	for _, item := range pods.Items {
		fmt.Printf("namespace: %v, name: %v\n", item.Namespace, item.Name)
	}
}

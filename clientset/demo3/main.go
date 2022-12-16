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
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 3、调用监听方法
	w, err := clientSet.AppsV1().Deployments("default").Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("start......")
	for e := range w.ResultChan() {
		fmt.Println(e.Type, e.Object)
	}
	// for {
	// 	select {
	// 	case e, _ := <-w.ResultChan():
	// 		fmt.Println(e.Type, e.Object)
	// 		// e.Type: 表示事件变化的类型，Added, Deleted, modify
	// 		// e.Object: 表示变化后的数据
	// 	}
	// }
}

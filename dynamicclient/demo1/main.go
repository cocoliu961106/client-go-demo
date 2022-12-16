package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1.加载配置文件，生成config对象
	config, err := clientcmd.BuildConfigFromFlags("", "../../kubeconfig")
	if err != nil {
		panic(err.Error())
	}

	// 2.实例化动态客户端对象
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 3.配置需要调用的GVR
	gvr := schema.GroupVersionResource{
		Group:    "", // 核心资源组不需要group
		Version:  "v1",
		Resource: "pods",
	}

	// 4. 发送请求，并得到返回结果
	unstructedObj, err := dynamicClient.
		Resource(gvr).
		Namespace("kube-system").
		List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// 5. unstructedObj转化为结构化数据
	podList := &corev1.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(
		unstructedObj.UnstructuredContent(),
		podList,
	)
	if err != nil {
		panic(err.Error())
	}
	/*
		1) Resource：基于gvr生成了一个针对于资源的客户端，也可以称之为动态资源客户端，dunamicResourceClient
		2) Namespace：指定一个可操作的命名空间，同时它是dynamicResourceClient的方法
		3) List：首先通过RESTClient调用k8s APIServer的接口返回了Pod的数据，返回的数据格式是二进制的Json格式，然后通过一些列的解析方法，转换成unstructured.UnstructuredList
	*/

	/*
		1) dynamicClient.Resource()函数用于设置请求的资源组、资源版本、资源名称。
		2) Namespace函数用于设置请求的命名空间。
		3) List函数用于获取Pod列表，得到的Pod列表为unstructured.UnstructuredList指针类型
		4) 然后通过runtime.DefaultUnstructuredConverter.FromUnstructured函数将unstructured.UnstructuredList转换为PodList类型。
	*/

	// 6. 格式化输出结果
	for _, item := range podList.Items {
		fmt.Printf("namespace: %v, name: %v, status: %v\n", item.Namespace, item.Name, item.Status.Phase)
	}

}

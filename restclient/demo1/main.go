package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1.加载配置文件，生成config对象
	config, err := clientcmd.BuildConfigFromFlags("", "../../kubeconfig")
	if err != nil {
		panic(err.Error())
	}

	// 2.配置api路径
	config.APIPath = "api" // pods, /api/v1/pods

	// 3.配置分组版本
	config.GroupVersion = &corev1.SchemeGroupVersion // 核心资源组，group: "", version: "v1"

	// 4.配置数据的编解码工具
	config.NegotiatedSerializer = scheme.Codecs

	// 5.实例化RESTClient对象
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err.Error())
	}

	// 6.定义接收返回值的变量
	result := &corev1.PodList{}

	// 7.跟apiserver交互
	err = restClient.
		Get().
		Namespace("kube-system").
		Resource("pods").
		VersionedParams(&metav1.ListOptions{}, scheme.ParameterCodec). // 参数及参数的序列化工具
		Do(context.TODO()).                                            // 触发请求
		Into(result)                                                   // 写入返回结果
	if err != nil {
		panic(err.Error())
	}

	/*
		1）Get定义请求方式，返回了一个Request结构体对象。这个Request结构体对象，就是构建访问APIServer用的。
		2）依次执行了Namespace、Resource、VersionedParams，构建与APIServer交互的参数。
		3）Do方法通过requets发起请求，然后通过transformResponse解析请求返回，并绑定到对应资源对象的结构体对象上。这里的话，表示的是corev1.PodList对象。
		4）request先是检查了有没有可用的client，在这里开始调用net/http包的功能。
	*/

	/*
		1) RESTClient发送请求的过程对Go语言标准库net/http进行了封装，由Do -> Request函数实现。
		2) 请求发送之前需要根据请求参数生成请求的RESTful URL，由r.URL().String()函数完成。 http://xxxx:x/api/v1/namespaces/kube-system/pods?limit=1
		3) 通过Go语言标准库net/http向RESTful URL(即kube-apiserver)发送请求，请求得到的结果存放在http.Response的Body对象中，fn函数（即transformResponse）将结果转化为资源对象。
	*/

	for _, item := range result.Items {
		fmt.Printf("namespace: %v, name: %v\n", item.Namespace, item.Name)
	}
}

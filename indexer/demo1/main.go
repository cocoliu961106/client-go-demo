package main

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

func main() {
	// 实例一个indexer对象
	index := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{
		"namespace": NamespaceIndexFunc,
		"nodeName":  NodeNameIndexFunc,
	})

	// 模拟数据
	pod1 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "index-pod-1",
			Namespace: "default",
		},
		Spec: v1.PodSpec{
			NodeName: "node1",
		},
	}
	pod2 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "index-pod-2",
			Namespace: "default",
		},
		Spec: v1.PodSpec{
			NodeName: "node2",
		},
	}
	pod3 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "index-pod-3",
			Namespace: "kube-system",
		},
		Spec: v1.PodSpec{
			NodeName: "node3",
		},
	}

	// 将数据写入到Indexer中
	_ = index.Add(pod1)
	_ = index.Add(pod2)
	_ = index.Add(pod3)

	// 通过索引器函数查询数据
	pods, err := index.ByIndex("namespace", "default")
	if err != nil {
		panic(err)
	}

	for _, pod := range pods {
		fmt.Println(pod.(*v1.Pod).Name)
	}

}

// 1、实现两个索引器函数，分别基于Namespace、NodeName，资源对象为pod

func NamespaceIndexFunc(obj interface{}) (result []string, err error) {
	pod, ok := obj.(*v1.Pod)
	if !ok {
		return nil, fmt.Errorf("类型错误， %v", err)
	}
	result = []string{pod.Namespace}
	return
}

func NodeNameIndexFunc(obj interface{}) (result []string, err error) {
	pod, ok := obj.(*v1.Pod)
	if !ok {
		return nil, fmt.Errorf("类型错误， %v", err)
	}
	result = []string{pod.Spec.NodeName}
	return
}

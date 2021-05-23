package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"time"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		time.Sleep(10 * time.Second)
		fmt.Printf("\n")
		namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			fmt.Printf("list namespace error: %s\n", err)
			continue
		}
		for _, namespace := range namespaces.Items {
			pods, err := clientset.CoreV1().Pods(namespace.Name).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				fmt.Printf("list pods error: %s\n", err)
				continue
			}
			for _, pod := range pods.Items {
				fmt.Printf("%s\t%s\t%s\n", namespace.Name, pod.Name, pod.Status.Phase)
			}
		}
	}
}

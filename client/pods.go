package client

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api"
)

func GetPods() {
	Pods = []Pod{}
	GetNamespaces()

	// get pods
	for _, ns := range Namespaces {
		podList, err := c.Pods(ns).List(api.ListOptions{})
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, pod := range podList.Items {
			readyContainers := 0
			for _, cs := range pod.Status.ContainerStatuses {
				if cs.Ready {
					readyContainers++
				}
			}
			p := Pod{
				Name:            pod.ObjectMeta.Name,
				Namespace:       pod.ObjectMeta.Namespace,
				NodeName:        pod.Spec.NodeName,
				Phase:           pod.Status.Phase,
				TotalContainers: len(pod.Status.ContainerStatuses),
				ReadyContainers: readyContainers,
			}
			Pods = append(Pods, p)
		}
	}
}

func ListPodsByNamespace() {
	fmt.Printf("Pods:\n\n")
	for _, v := range Pods {
		fmt.Printf("\tStatus: %s; Ready: %d of %d; %s/%s\n", v.Phase, v.ReadyContainers, v.TotalContainers, v.Namespace, v.Name)
	}
	fmt.Printf("\n")
}

func ListPodsByNodeName() {
	fmt.Printf("Pods:\n\n")
	for _, v := range Pods {
		fmt.Printf("\t%s/%s\n", v.NodeName, v.Name)
	}
	fmt.Printf("\n")
}

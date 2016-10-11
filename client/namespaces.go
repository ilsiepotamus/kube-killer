package client

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api"
)

func GetNamespaces() {
	Namespaces = []string{}

	// get namespaces
	namespaceList, err := c.Namespaces().List(api.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, namespace := range namespaceList.Items {
		Namespaces = append(Namespaces, namespace.ObjectMeta.Name)
	}
}

func ListNamespaces() {
	fmt.Printf("Namespaces:\n\n")
	for _, v := range Namespaces {
		fmt.Printf("\t%s\n", v)
	}
	fmt.Printf("\n")
}

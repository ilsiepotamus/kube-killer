package client

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api"
)

func GetNamespaces() []string {
	Namespaces = []string{}

	// get namespaces
	namespaceList, err := c.Namespaces().List(api.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return Namespaces
	}
	for _, namespace := range namespaceList.Items {
		Namespaces = append(Namespaces, namespace.ObjectMeta.Name)
	}

	return Namespaces
}

func ValidateNamespaces(namespaces ...string) []string {
	validNS := GetNamespaces()
	ns := []string{}

	for _, v := range namespaces {
		for _, c := range validNS {
			if v == c {
				ns = append(ns, v)
			}
		}
	}

	return ns
}

func ListNamespaces() {
	fmt.Printf("Namespaces:\n\n")
	for _, v := range Namespaces {
		fmt.Printf("\t%s\n", v)
	}
	fmt.Printf("\n")
}

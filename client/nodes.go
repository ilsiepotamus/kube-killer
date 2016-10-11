package client

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api"
)

func GetNodes() {
	Nodes = []string{}

	// get nodes
	nodeList, err := c.Nodes().List(api.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, node := range nodeList.Items {
		Nodes = append(Nodes, node.ObjectMeta.Name)
	}
}

func ListNodes() {
	fmt.Printf("Nodes:\n\n")
	for _, v := range Nodes {
		fmt.Printf("\t%s\n", v)
	}
	fmt.Printf("\n")
}

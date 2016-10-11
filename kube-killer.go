package main

import (
	"github.com/ilsiepotamus/kube-killer/client"
)

func main() {
	client.New()
	client.GetNodes()
	client.GetNamespaces()
	client.GetReplicationControllers()
	client.GetPods()

	client.ListNamespaces()
	client.ListNodes()
	client.ListReplicationControllers()
	client.ListPodsByNamespace()
}

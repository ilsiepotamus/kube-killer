package client

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api"
)

func GetReplicationControllers() {
	RCs = []RC{}
	GetNamespaces()

	// get replication controllers
	for _, ns := range Namespaces {
		rcList, err := c.ReplicationControllers(ns).List(api.ListOptions{})
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, rc := range rcList.Items {
			r := RC{
				Name:              rc.ObjectMeta.Name,
				Namespace:         rc.ObjectMeta.Namespace,
				Replicas:          rc.Status.Replicas,
				ReadyReplicas:     rc.Status.ReadyReplicas,
				AvailableReplicas: rc.Status.AvailableReplicas,
			}
			RCs = append(RCs, r)
		}
	}
}

func ListReplicationControllers() {
	fmt.Printf("Replication Controllers:\n\n")
	for _, v := range RCs {
		fmt.Printf("\t%s/%s\n", v.Namespace, v.Name)
	}
	fmt.Printf("\n")
}

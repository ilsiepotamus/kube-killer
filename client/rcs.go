package client

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api"
)

func GetReplicationControllers(namespaces ...string) {
	RCs = map[string]RC{}
	if len(namespaces) == 0 {
		namespaces = GetNamespaces()
	} else {
		namespaces = ValidateNamespaces(namespaces...)
	}

	// get replication controllers
	for _, ns := range namespaces {
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
				Definition:        &rc,
			}
			RCs[rc.ObjectMeta.Namespace+`-`+rc.ObjectMeta.Name] = r
			fmt.Println(rc.ObjectMeta.Namespace + `-` + rc.ObjectMeta.Name)
			fmt.Println(RCs[rc.ObjectMeta.Namespace+`-`+rc.ObjectMeta.Name])
		}
	}
	fmt.Println(namespaces)
	fmt.Println(RCs)
}

func ListReplicationControllers() {
	fmt.Printf("Replication Controllers:\n\n")
	for _, v := range RCs {
		fmt.Printf("\t%s/%s\n", v.Namespace, v.Name)
	}
	fmt.Printf("\n")
}

func (rc RC) Scale(replicas int32) {
	fmt.Println(rc.Definition.Spec)
	rc.Definition.Spec.Replicas = replicas
	newRC, err := c.ReplicationControllers(rc.Namespace).Update(rc.Definition)
	if err != nil {
		fmt.Println("Couldn't scale.")
	}
	fmt.Println(newRC.Spec)
	fmt.Println(newRC.Status)
	rc.Definition = newRC
	rc.Replicas = newRC.Status.Replicas
	rc.ReadyReplicas = newRC.Status.ReadyReplicas
	rc.AvailableReplicas = newRC.Status.AvailableReplicas
}

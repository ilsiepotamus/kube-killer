package client

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/restclient"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

var Nodes []string
var Namespaces []string
var RCs map[string]RC
var Pods []Pod
var c *client.Client

type RC struct {
	Name              string
	Namespace         string
	Replicas          int32
	ReadyReplicas     int32
	AvailableReplicas int32
	Definition        *api.ReplicationController
}

type Pod struct {
	Name            string
	Namespace       string
	NodeName        string
	Phase           api.PodPhase
	TotalContainers int
	ReadyContainers int
	Definition      *api.Pod
}

func New() *client.Client {
	config := &restclient.Config{
		Host: "http://localhost:8080",
	}

	var err error
	c, err = client.New(config)
	if err != nil {
		fmt.Println(err)
	}

	return c
}

func Snapshot() {

}

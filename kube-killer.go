package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ilsiepotamus/kube-killer/client"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("nothing is happening!")
		os.Exit(1)
	}

	client.New()
	// get Pods flag set
	getNodesCommand := flag.NewFlagSet("nodes", flag.ExitOnError)
	getNamespacesCommand := flag.NewFlagSet("ns", flag.ExitOnError)
	getRCsCommand := flag.NewFlagSet("rcs", flag.ExitOnError)
	getPodsCommand := flag.NewFlagSet("pods", flag.ExitOnError)
	getKillCommand := flag.NewFlagSet("kill", flag.ExitOnError)
	getRestoreCommand := flag.NewFlagSet("restore", flag.ExitOnError)

	podNSPtr := getPodsCommand.String("ns", "default", "comma separated namespaces to get pods from")
	rcNSPtr := getRCsCommand.String("ns", "default", "comma separated namespaces to get replication controllers from")
	killRCPtr := getKillCommand.String("rc", "", "ns/rc to kill")
	restoreRCPtr := getRestoreCommand.String("rc", "", "ns/rc to restore")

	switch os.Args[1] {
	case "nodes":
		getNodesCommand.Parse(os.Args[2:])
		client.GetNodes()
		client.ListNodes()
	case "ns":
		getNamespacesCommand.Parse(os.Args[2:])
		client.GetNamespaces()
		client.ListNamespaces()
	case "rcs":
		getRCsCommand.Parse(os.Args[2:])
		ns := strings.Split(*rcNSPtr, ",")
		client.GetReplicationControllers(ns...)
		client.ListReplicationControllers()
	case "pods":
		getPodsCommand.Parse(os.Args[2:])
		ns := strings.Split(*podNSPtr, ",")
		client.GetPods(ns...)
		client.ListPodsByNamespace()
	case "kill":
		getKillCommand.Parse(os.Args[2:])
		ns := strings.Split(*killRCPtr, "/")[0]
		fmt.Println(ns)
		client.GetReplicationControllers(ns)
		client.RCs[*killRCPtr].Scale(0)
	case "restore":
		getRestoreCommand.Parse(os.Args[2:])
		ns := strings.Split(*restoreRCPtr, "/")[0]
		rc := strings.Split(*restoreRCPtr, "/")[1]
		fmt.Println(ns)
		fmt.Println(ns + "-" + rc)
		client.GetReplicationControllers(ns)
		client.RCs[ns+"-"+rc].Scale(1)
	default:
		panic("Not a command")
	}
}

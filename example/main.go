package main

import (
	"context"
	"log"
	"os"

	"github.com/phelian/goelector"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	nodeID, ok := os.LookupEnv("HOSTNAME")
	if !ok {
		panic("missing HOSTNAME env var")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clientset, err := newClientset()
	if err != nil {
		log.Printf("failed to connect to cluster")
		return
	}

	goelector.TurnOffKlog()
	if err := goelector.Start(ctx, goelector.GetDefaultConfig(), nodeID, clientset); err != nil {
		log.Printf(err.Error())
		return
	}
}

func newClientset() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

package kuevs

import (
	"context"
	"encoding/json"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

func WatchEvent(kubeconfig, database string) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	Check(err)

	clienetset := kubernetes.NewForConfigOrDie(config)

	watcher, err := clienetset.CoreV1().Events(v1.NamespaceAll).Watch(context.Background(), metav1.ListOptions{})
	Check(err)

	db := InitDatabase(database)
	for event := range watcher.ResultChan() {
		evt := InitEvent(event)
		db.Save(evt)
	}
}

func ShowEvent(e watch.Event) {
	data, err := json.MarshalIndent(e, "", "  ")
	Check(err)
	fmt.Println(string(data))
}

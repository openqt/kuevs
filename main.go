package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	version, commit, date = "dev", "dev", ""
)

func checkerr(err error) {
	if err != nil {
		glog.Errorln(err)
	}
}

func main() {
	// fmt.Printf("version : %s\ncommit  : %s\ncompiled: %s\n", version, commit, date)

	kubeconfig := flag.String("kubeconfig", os.ExpandEnv("$HOME/.kube/config"), "absolute path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	checkerr(err)

	clienetset := kubernetes.NewForConfigOrDie(config)

	watcher, err := clienetset.CoreV1().Events("").Watch(context.Background(), metav1.ListOptions{})
	checkerr(err)

	for event := range watcher.ResultChan() {
		data, err := json.MarshalIndent(event, "", "  ")
		checkerr(err)
		fmt.Println(string(data))
	}
}

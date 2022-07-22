package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/openqt/kuevs/kuevs"
)

var (
	version, commit, date = "0.0", "dev", "0000"
)

func main() {
	kubeconfig := flag.String("kubeconfig", os.ExpandEnv("$HOME/.kube/config"), "absolute path to the kubeconfig file")
	database := flag.String("database", "ks-events.sqlite", "set database filename")
	ver := flag.Bool("version", false, "show current version")
	flag.Parse()
	defer glog.Flush()

	if *ver {
		fmt.Printf("version : %s\ncommit  : %s\ncompiled: %s\n", version, commit, date)
		os.Exit(0)
	}

	kuevs.WatchEvent(*kubeconfig, *database)
}

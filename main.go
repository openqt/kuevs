package main

import "fmt"

var (
	version, commit, date = "dev", "dev", ""
)

func main() {
	fmt.Printf("version : %s\ncommit  : %s\ncompiled: %s\n", version, commit, date)
}

package main

import "fmt"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	fmt.Printf("my app %s, commit %s, built at %s", version, commit, date)
}

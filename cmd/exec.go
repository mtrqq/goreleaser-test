package main

import (
	"fmt"

	"github.com/mtrqq/goreleaser-test/lib"
)

func main() {
	fmt.Printf("my app %s, commit %s, built at %s", lib.Version, lib.Commit, lib.Date)
}

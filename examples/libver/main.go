package main

import (
	clasp "github.com/synesissoftware/CLASP.Go"
	ver2go "github.com/synesissoftware/ver2go"

	"fmt"
)

func main() {
	fmt.Printf("CLASP.Go v%s\n", clasp.VersionString())
	fmt.Printf("ver2go v%s\n", ver2go.VersionString())
}

// examples/show_usage_and_version.go

package main

import (

	clasp "github.com/synesissoftware/CLASP.Go"

	"fmt"
	"os"
)

const (

	ProgramVersion	=	"0.0.1"
)

func main() {

	aliases	:= []clasp.Alias {

		clasp.HelpFlag(),
		clasp.VersionFlag(),
	}

	args := clasp.Parse(os.Args, clasp.ParseParams{ Aliases: aliases })

	if args.FlagIsSpecified(clasp.HelpFlag()) {

		clasp.ShowUsage(aliases, clasp.UsageParams{ Version: ProgramVersion, InfoLines: []string { "examples", "", ":version:", "" }})
	}

	if args.FlagIsSpecified(clasp.VersionFlag()) {

		clasp.ShowVersion(aliases, clasp.UsageParams{ Version: ProgramVersion })
	}

	if unused := args.GetUnusedFlagsAndOptions(); 0 != len(unused) {

		fmt.Fprintf(os.Stderr, "%s: unrecognised flag/option: %s\n", args.ProgramName, unused[0].Str())

		os.Exit(1)
	}

	fmt.Printf("no flags specified\n")
}


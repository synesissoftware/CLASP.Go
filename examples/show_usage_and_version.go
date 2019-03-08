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

	// Specify aliases, parse, and checking standard flags

	aliases	:= []clasp.Alias {

		clasp.HelpFlag(),
		clasp.VersionFlag(),
	}

	args := clasp.Parse(os.Args, clasp.ParseParams{ Aliases: aliases })

	if args.FlagIsSpecified(clasp.HelpFlag()) {

		clasp.ShowUsage(aliases, clasp.UsageParams{

			Version: ProgramVersion,
			InfoLines: []string { "CLASP.Go Examples", "", ":version:", "" }
		})
	}

	if args.FlagIsSpecified(clasp.VersionFlag()) {

		clasp.ShowVersion(aliases, clasp.UsageParams{ Version: ProgramVersion })
	}


	// Check for any unrecognised flags or options

	if unused := args.GetUnusedFlagsAndOptions(); 0 != len(unused) {

		fmt.Fprintf(os.Stderr, "%s: unrecognised flag/option: %s\n", args.ProgramName, unused[0].Str())

		os.Exit(1)
	}


	// Finish normal processing

	fmt.Printf("no flags specified\n")
}


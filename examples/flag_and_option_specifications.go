// examples/flag_and_option_specifications.go

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

	// Specify specifications, parse, and checking standard flags

	flag_Debug			:=	clasp.Flag("--debug").SetHelp("runs in Debug mode").SetAlias("-d").End()
	option_Verbosity	:=	clasp.Option("--verbosity").SetHelp("specifies the verbosity").SetAlias("-v").SetValues("terse", "quiet", "silent", "chatty").End()
	flag_Chatty			:=	clasp.AliasesFor("--verbosity=chatty", "-c")

	specifications	:= []clasp.Specification {

		flag_Debug,
		option_Verbosity,
		flag_Chatty,

		clasp.HelpFlag(),
		clasp.VersionFlag(),
	}

	args := clasp.Parse(os.Args, clasp.ParseParams{ Specifications: specifications })

	if args.FlagIsSpecified(clasp.HelpFlag()) {

		clasp.ShowUsage(specifications, clasp.UsageParams{

			Version: ProgramVersion,
			InfoLines: []string { "CLASP.Go Examples", "", ":version:", "" },
		})
	}

	if args.FlagIsSpecified(clasp.VersionFlag()) {

		clasp.ShowVersion(specifications, clasp.UsageParams{ Version: ProgramVersion })
	}


	// Program-specific processing of flags/options

	if opt, found := args.LookupOption("--verbosity"); found {

		fmt.Printf("verbosity is specified as: %s\n", opt.Value)
	}

	if args.FlagIsSpecified("--debug") {

		fmt.Printf("Debug mode is specified\n")
	}


	// Check for any unrecognised flags or options

	if unused := args.GetUnusedFlagsAndOptions(); 0 != len(unused) {

		fmt.Fprintf(os.Stderr, "%s: unrecognised flag/option: %s\n", args.ProgramName, unused[0].Str())

		os.Exit(1)
	}


	// Finish normal processing

	return
}


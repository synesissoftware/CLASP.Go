// examples/flag_and_option_aliases.go

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

	flag_Debug			:=	clasp.Alias{ clasp.Flag, "--debug", []string{ "-d" }, "runs in Debug mode", nil, 0 }
	option_Verbosity	:=	clasp.Alias{ clasp.Option, "--verbosity", []string{ "-v" }, "specifies the verbosity", []string{ "terse", "quiet", "silent", "chatty" }, 0 }
	flag_Chatty			:=	clasp.Alias{ clasp.Flag, "--verbosity=chatty", []string{ "-c" }, "", nil, 0 }

	aliases	:= []clasp.Alias {

		flag_Debug,
		option_Verbosity,
		flag_Chatty,

		clasp.HelpFlag(),
		clasp.VersionFlag(),
	}

	args := clasp.Parse(os.Args, clasp.ParseParams{ Aliases: aliases })

	if args.FlagIsSpecified(clasp.HelpFlag()) {

		clasp.ShowUsage(aliases, clasp.UsageParams{

			Version: ProgramVersion,
			InfoLines: []string { "CLASP.Go Examples", "", ":version:", "" },
		})
	}

	if args.FlagIsSpecified(clasp.VersionFlag()) {

		clasp.ShowVersion(aliases, clasp.UsageParams{ Version: ProgramVersion })
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


// examples/check_and_show_usage.go

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

		clasp.ShowUsage(aliases, clasp.UsageParams{})
	}

	if args.FlagIsSpecified(clasp.VersionFlag()) {

		clasp.ShowVersion(aliases, clasp.UsageParams{ Version: ProgramVersion })
	}

	if o_v, found := args.LookupOption("--verbosity"); found {

		fmt.Printf("verbosity is specified as: %s\n", o_v.Value)
	}

	if args.FlagIsSpecified("--debug") {

		fmt.Printf("Debug mode is specified\n")
	}

	if unused := args.GetUnusedFlagsAndOptions(); 0 != len(unused) {

		fmt.Fprintf(os.Stderr, "%s: unrecognised flag/option: %s\n", args.ProgramName, unused[0].Str())

		os.Exit(1)
	}

	fmt.Printf("no flags/options specified\n")
}


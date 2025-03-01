// examples/bit_flags.go

package main

import (
	"fmt"

	clasp "github.com/synesissoftware/CLASP.Go"

	"os"
)

const (
	ProgramVersion = "0.0.1"
)

const (
	BF_Sound int64 = 1 << iota
	BF_Vision
)

func main() {

	// Specify specifications, parse, and checking standard flags

	var flags int64 = 0

	flag_Sound := clasp.Flag("--enable-sound").SetAlias("-s").SetHelp("Enables sound").SetBitFlags64(BF_Sound, &flags)
	flag_Vision := clasp.Flag("--enable-vision").SetAlias("-v").SetHelp("Enables vision").SetBitFlags64(BF_Vision, &flags)

	specifications := []clasp.Specification{

		clasp.Section("behaviour:"),
		flag_Sound,
		flag_Vision,

		clasp.Section("standard:"),
		clasp.HelpFlag(),
		clasp.VersionFlag(),
	}

	args := clasp.Parse(os.Args, clasp.ParseParams{Specifications: specifications})

	if args.FlagIsSpecified(clasp.HelpFlag()) {

		clasp.ShowUsage(specifications, clasp.UsageParams{

			Version: ProgramVersion,
			InfoLines: []string{
				"CLASP.Go Examples",
				"Example illustrating use of bit-mask flags",
				":version:",
				"",
			},
		})
	}

	// Check for any unrecognised flags or options

	if unused := args.GetUnusedFlagsAndOptions(); 0 != len(unused) {

		fmt.Fprintf(os.Stderr, "%s: unrecognised flag/option: %s\n", args.ProgramName, unused[0].Str())

		os.Exit(1)
	}

	// Program logic

	if 0 != (BF_Sound & flags) {
		fmt.Println("running with sound")
	}
	if 0 != (BF_Vision & flags) {
		fmt.Println("running with vision")
	}
	if 0 == flags {
		fmt.Println("running in default mode")
	}
}

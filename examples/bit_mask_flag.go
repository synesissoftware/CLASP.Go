// examples/bit_mask_flag.go

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

	flag_Sound := clasp.Flag("--enable-sound").SetAlias("-s").SetHelp("Enables sound").SetBitmask64(BF_Sound, nil)
	flag_Vision := clasp.Flag("--enable-vision").SetAlias("-v").SetHelp("Enables vision").SetBitmask64(BF_Vision, nil)

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

	flags := args.CheckAllBit64Flags()

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

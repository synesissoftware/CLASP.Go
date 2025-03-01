# CLASP.Go Example - **bit_flags**

## Summary

Example illustrating use of `BitFlags()` / `BitFlags64()` for associating a given flag specification with a bit-flags value and, optionally, a bitmask variable, to be applied automatically when detected from the command-line during parsing

## Source

```Go
// examples/bit_flags.go

package main

import (
	clasp "github.com/synesissoftware/CLASP.Go"

	"os"
	"fmt"
)

const (
	ProgramVersion = "0.0.2"
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
```

## Usage

### No arguments

If executed with no arguments

```bash
go run examples/bit_flags.go
```

it gives the output:

```
running in default mode
```

### Show usage

If executed with the arguments

```bash
go run examples/bit_flags.go --help
```

it gives the output:

```
CLASP.Go Examples
Example illustrating use of bit-mask flags
bit_flags 0.0.1

USAGE: bit_flags [ ... flags and options ... ]

flags/options:

	behaviour:

	-s
	--enable-sound
		Enables sound

	-v
	--enable-vision
		Enables vision

	standard:

	--help
		Shows this help and exits

	--version
		Shows version information and exits
```

### Specify flags and options in long-form

If executed with the arguments

```bash
go run examples/bit_flags.go --enable-sound
```

it gives the output:

```
running with sound
```

### Specify flags and options in short-form

If executed with the arguments

```bash
go run examples/bit_flags.go -v
```

it gives the (same) output:

```
running with vision
```

### Specify flags and options with combined short-form

If executed with the arguments

```bash
go run examples/bit_flags.go -sv
```

it gives the (same) output:

```
running with sound
running with vision
```


<!-- ########################### end of file ########################### -->


# CLASP.Go Example - **flag_and_option_specifications**

## Summary

Example illustrating various kinds of *flag* and *option* specifications, including the combination of short-names.

## Source

```Go
// examples/flag_and_option_specifications.go

package main

import (
	clasp "github.com/synesissoftware/CLASP.Go"

	"fmt"
	"os"
)

const (
	ProgramVersion = "0.1.1"
)

func main() {

	// Specify specifications, parse, and checking standard flags

	flag_Debug := clasp.Flag("--debug").SetHelp("runs in Debug mode").SetAlias("-d").End()
	option_Verbosity := clasp.Option("--verbosity").SetHelp("specifies the verbosity").SetAlias("-v").SetValues("terse", "quiet", "silent", "chatty").End()
	flag_Chatty := clasp.AliasesFor("--verbosity=chatty", "-c")

	specifications := []clasp.Specification{

		clasp.Section("behaviour:"),
		flag_Debug,
		option_Verbosity,
		flag_Chatty,

		clasp.Section("standard:"),
		flag_Debug,
		clasp.HelpFlag(),
		clasp.VersionFlag(),
	}

	args := clasp.Parse(os.Args, clasp.ParseParams{Specifications: specifications})

	if args.FlagIsSpecified(clasp.HelpFlag()) {

		clasp.ShowUsage(specifications, clasp.UsageParams{

			Version: ProgramVersion,
			InfoLines: []string{
				"CLASP.Go Examples",
				"Example illustrating various kinds of flag and option specifications, including the combination of short-names",
				":version:",
				"",
			},
		})
	}

	if args.FlagIsSpecified(clasp.VersionFlag()) {

		clasp.ShowVersion(specifications, clasp.UsageParams{Version: ProgramVersion})
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
```

## Usage

### No arguments

If executed with no arguments

```bash
go run examples/flag_and_option_specifications.go
```

it gives the output:

```
```

### Show usage

If executed with the arguments

```bash
go run examples/flag_and_option_specifications.go --help
```

it gives the output:

```
CLASP.Go Examples
Example illustrating various kinds of flag and option specifications, including the combination of short-names
flag_and_option_specifications 0.1.1

USAGE: flag_and_option_specifications [ ... flags and options ... ]

flags/options:

	behaviour:

	-d
	--debug
		runs in Debug mode

	-c --verbosity=chatty
	-v <value>
	--verbosity=<value>
		specifies the verbosity
		where <value> one of:
			terse
			quiet
			silent
			chatty

	standard:

	-d
	--debug
		runs in Debug mode

	--help
		Shows this help and exits

	--version
		Shows version information and exits
```

### Specify flags and options in long-form

If executed with the arguments

```bash
go run examples/flag_and_option_specifications.go --debug --verbosity=silent
```

it gives the output:

```
verbosity is specified as: silent
Debug mode is specified
```

### Specify flags and options in short-form

If executed with the arguments

```bash
go run examples/flag_and_option_specifications.go -v silent -d
```

it gives the (same) output:

```
verbosity is specified as: silent
Debug mode is specified
```

### Specify flags and options in short-form, including an alias for an option-with-value

If executed with the arguments

```bash
go run examples/flag_and_option_specifications.go -c -d
```

it gives the output:

```
verbosity is specified as: chatty
Debug mode is specified
```

### Specify flags and options with combined short-form

If executed with the arguments

```bash
go run examples/flag_and_option_specifications.go -dc
```

it gives the (same) output:

```
verbosity is specified as: chatty
Debug mode is specified
```


<!-- ########################### end of file ########################### -->


# CLASP.Go Example - **flag_and_option_aliases**

## Summary

Example illustrating various kinds of *flag* and *option* aliases, including the combination of short-names.

## Source

```Go
// examples/flag_and_option_aliases.go

package main

import (

    clasp "github.com/synesissoftware/CLASP.Go"

    "fmt"
    "os"
)

const (

    ProgramVersion  =   "0.0.1"
)

func main() {

    // Specify aliases, parse, and checking standard flags

    flag_Debug          :=  clasp.Alias{ clasp.Flag, "--debug", []string{ "-d" }, "runs in Debug mode", nil, 0 }
    option_Verbosity    :=  clasp.Alias{ clasp.Option, "--verbosity", []string{ "-v" }, "specifies the verbosity", []string{ "terse", "quiet", "silent", "chatty" }, 0 }
    flag_Chatty         :=  clasp.Alias{ clasp.Flag, "--verbosity=chatty", []string{ "-c" }, "", nil, 0 }

    aliases := []clasp.Alias {

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
```

## Usage

### No arguments

If executed with no arguments

```
    go run examples/flag_and_option_aliases.go
```

it gives the output:

```
```

### Show usage

If executed with the arguments

```
    go run examples/flag_and_option_aliases.go --help
```

it gives the output:

```
CLASP.Go Examples

flag_and_option_aliases 0.0.1

USAGE: flag_and_option_aliases [ ... flags and options ... ]

flags/options:

    -d
    --debug
        runs in Debug mode

    -v <value>
    --verbosity=<value>
        specifies the verbosity

    -c
    --verbosity=chatty


    --help
        Shows this helps and exits

    --version
        Shows version information and exits
```

### Specify flags and options in long-form

If executed with the arguments

```
    go run examples/flag_and_option_aliases.go --debug --verbosity=silent
```

it gives the output:

```
verbosity is specified as: silent
Debug mode is specified
```

### Specify flags and options in short-form

If executed with the arguments

```
    go run examples/flag_and_option_aliases.go -v silent -d
```

it gives the (same) output:

```
verbosity is specified as: silent
Debug mode is specified
```

### Specify flags and options in short-form, including an alias for an option-with-value

If executed with the arguments

```
    go run examples/flag_and_option_aliases.go -c -d
```

it gives the output:

```
verbosity is specified as: chatty
Debug mode is specified
```

### Specify flags and options with combined short-form

If executed with the arguments

```
    go run examples/flag_and_option_aliases.go -dc
```

it gives the (same) output:

```
verbosity is specified as: chatty
Debug mode is specified
```


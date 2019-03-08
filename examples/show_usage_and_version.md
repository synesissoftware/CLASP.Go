# CLASP.Go Example - **show_usage_and_version**

### Example1 - show_usage_and_version.go



### Source

```Go
// examples/show_usage_and_version.go

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

    aliases := []clasp.Alias {

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
```

### Usage

#### No arguments

If executed with no arguments

```
    go run examples/show_usage_and_version.go
```

it gives the output:

```
no flags specified
```

#### Show usage

If executed with the arguments

```
    go run examples/show_usage_and_version.go --help
```

it gives the output:

```
CLASP.Go Examples

show_usage_and_version 0.0.1

USAGE: show_usage_and_version [ ... flags and options ... ]

flags/options:

    --help
        Shows this helps and exits

    --version
        Shows version information and exits
```

#### Show version

If executed with the arguments

```
    go run examples/show_usage_and_version.go --version
```

it gives the output:

```
show_usage_and_version 0.0.1
```

#### Unknown option

If executed with the arguments

```
    go run examples/show_usage_and_version.go --unknown=value
```

it gives the output (on the standard error stream):

```
show_usage_and_version: unrecognised flag/option: --unknown=value
```

with an exit code of 1

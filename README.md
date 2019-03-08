# CLASP.Go
Command-Line Argument Sorting and Parsing for Go

## Table of Contents

1. [Introduction](#introduction)
2. [Installation](#installation)
3. [Components](#components)
4. [Examples](#examples)
5. [Project Information](#project-information)

## Installation

```Go

import clasp "github.com/synesissoftware/CLASP.Go"
```

## Components

TBC

## Examples

### Example1 - show_usage_and_version.go

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

    aliases := []clasp.Alias {

        clasp.HelpFlag(),
        clasp.VersionFlag(),
    }

    args := clasp.Parse(os.Args, clasp.ParseParams{ Aliases: aliases })

    if args.FlagIsSpecified(clasp.HelpFlag()) {

        clasp.ShowUsage(aliases, clasp.UsageParams{ Version: ProgramVersion, InfoLines: []string { "examples", "", ":version:", "" }})
    }

    if args.FlagIsSpecified(clasp.VersionFlag()) {

        clasp.ShowVersion(aliases, clasp.UsageParams{ Version: ProgramVersion })
    }

    if unused := args.GetUnusedFlagsAndOptions(); 0 != len(unused) {

        fmt.Fprintf(os.Stderr, "%s: unrecognised flag/option: %s\n", args.ProgramName, unused[0].Str())

        os.Exit(1)
    }

    fmt.Printf("no flags specified\n")
}
```

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
examples

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

with an exit code of 0



## Project Information

### Where to get help

[GitHub Page](https://github.com/synesissoftware/CLASP.Gp "GitHub Page")

### Contribution guidelines

Defect reports, feature requests, and pull requests are welcome on https://github.com/synesissoftware/CLASP.Gp.

### Dependencies

### Related projects

### License

**CLASP.Gp** is released under the 3-clause BSD license. See [LICENSE](./LICENSE) for details.

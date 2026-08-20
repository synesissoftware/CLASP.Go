# CLASP.Go <!-- omit in toc -->

**C**ommand-**L**ine **A**rgument **S**orting and **P**arsing, for Go

![Language](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GitHub release](https://img.shields.io/github/v/release/synesissoftware/CLASP.Go.svg)](https://github.com/synesissoftware/CLASP.Go/releases/latest)
[![Last Commit](https://img.shields.io/github/last-commit/synesissoftware/CLASP.Go)](https://github.com/synesissoftware/CLASP.Go/commits/master)
[![Go](https://github.com/synesissoftware/CLASP.Go/actions/workflows/go.yml/badge.svg)](https://github.com/synesissoftware/CLASP.Go/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/synesissoftware/CLASP.Go.svg)](https://pkg.go.dev/github.com/synesissoftware/CLASP.Go)


## Table of Contents <!-- omit in toc -->

- [Introduction](#introduction)
- [Installation](#installation)
- [Components](#components)
  - [Command-line parsing](#command-line-parsing)
  - [Declarative specification of flags and options](#declarative-specification-of-flags-and-options)
  - [Utility functions for displaying usage and version information](#utility-functions-for-displaying-usage-and-version-information)
  - [Version](#version)
- [Examples](#examples)
- [Project Information](#project-information)
  - [Where to get help](#where-to-get-help)
  - [Contribution guidelines](#contribution-guidelines)
  - [Dependencies](#dependencies)
    - [Development/Testing Dependencies](#developmenttesting-dependencies)
  - [Related projects](#related-projects)
  - [License](#license)


## Introduction

**CLASP** stands for **C**ommand-**L**ine **A**rgument **S**orting and
**P**arsing. The first CLASP library was a C library with a C++ wrapper. There
have been [several implementations in other languages](#related-projects). **CLASP.Go** is the
Go version.


## Installation

Install via `go get`, as in:

```bash
go get "github.com/synesissoftware/CLASP.Go"
```

and then import as:

```Go
import clasp "github.com/synesissoftware/CLASP.Go"
```

or, simply, as:

```Go
import "github.com/synesissoftware/CLASP.Go"
```


## Components

### Command-line parsing

All **CLASP** libraries discriminate between *flags*, *options*, and *values*. **CLASP.Go** provides types and functions to specify arguments (`Flag()`, `Option()`, `Section()`, …), parse the command line (`Parse()`), and inspect the results (`Argument`, `Arguments`).

### Declarative specification of flags and options

`Specification` describes each command-line element (name, aliases, help, value sets, bit-flags receivers). See [EXAMPLES.md](./EXAMPLES.md) for worked examples.

### Utility functions for displaying usage and version information

`ShowUsage()` and `ShowVersion()` display help and version information and may terminate the process. Standard `--help` and `--version` specifications are available via `HelpFlag()` and `VersionFlag()`.

### Version

```Go
const (
	VersionMajor uint16 = /* ... */
	VersionMinor uint16 = /* ... */
	VersionPatch uint16 = /* ... */
	VersionAB    uint16 = /* ... */
)

func Version() uint64
func VersionString() string
```


## Examples

Examples are provided in the `examples` directory, along with a markdown description for each. A detailed list TOC of them is provided in [EXAMPLES.md](./EXAMPLES.md).


## Project Information


### Where to get help

[GitHub Page](https://github.com/synesissoftware/CLASP.Go "GitHub Page")


### Contribution guidelines

Defect reports, feature requests, and pull requests are welcome on https://github.com/synesissoftware/CLASP.Go.


### Dependencies

* [**ver2go**](https://github.com/synesissoftware/ver2go/);


#### Development/Testing Dependencies

* [**ANGoLS**](https://github.com/synesissoftware/ANGoLS/);
* [**require**](https://github.com/stretchr/testify/);
* [**STEGoL**](https://github.com/synesissoftware/STEGoL/);


### Related projects

**CLASP.Ruby** is inspired by the [C/C++ CLASP library](https://github.com/synesissoftware/CLASP), which is documented in the articles:

* _An Introduction to \CLASP_, Matthew Wilson, [CVu](http://accu.org/index.php/journals/c77/), January 2012;
* _[Anatomy of a CLI Program written in C](http://synesis.com.au/publishing/software-anatomies/anatomy-of-a-cli-program-written-in-c.html)_, Matthew Wilson, [CVu](http://accu.org/index.php/journals/c77/), September 2012; and
* _[Anatomy of a CLI Program written in C++](http://synesis.com.au/publishing/software-anatomies/anatomy-of-a-cli-program-written-in-c++.html)_, Matthew Wilson, [CVu](http://accu.org/index.php/journals/c77/), September 2015.

Other CLASP libraries include:

* [**CLASP**](https://github.com/synesissoftware/CLASP/);
* [**CLASP.js**](https://github.com/synesissoftware/CLASP.js/);
* [**CLASP.NET**](https://github.com/synesissoftware/CLASP.NET/);
* [**CLASP.Python**](https://github.com/synesissoftware/CLASP.Python/);
* [**CLASP.Ruby**](https://github.com/synesissoftware/CLASP.Ruby/);

Projects in which **CLASP.Go** is used include:

**CLASP.Go** is used in the **[libCLImate.Go](https://github.com/synesissoftware/libCLImate.Go)** library.


### License

**CLASP.Go** is released under the 3-clause BSD license. See [LICENSE](./LICENSE) for details.


<!-- ########################### end of file ########################### -->

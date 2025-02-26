# CLASP.Go <!-- omit in toc -->
**C**ommand-**L**ine **A**rgument **S**orting and **P**arsing for Go


## Introduction

**CLASP** stands for **C**ommand-**L**ine **A**rgument **S**orting and
**P**arsing. The first CLASP library was a C library with a C++ wrapper. There
have been [several implementations in other languages](#related-projects). **CLASP.Go** is the
Go version.


## Table of Contents <!-- omit in toc -->

- [Introduction](#introduction)
- [Installation](#installation)
- [Components](#components)
- [Examples](#examples)
- [Project Information](#project-information)
	- [Where to get help](#where-to-get-help)
	- [Contribution guidelines](#contribution-guidelines)
	- [Dependencies](#dependencies)
		- [Development/Testing Dependencies](#developmenttesting-dependencies)
	- [Related projects](#related-projects)
	- [License](#license)

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

T.B.C.


## Examples

Examples are provided in the ```examples``` directory, along with a markdown description for each. A detailed list TOC of them is provided in [EXAMPLES.md](./EXAMPLES.md).


## Project Information


### Where to get help

[GitHub Page](https://github.com/synesissoftware/CLASP.Go "GitHub Page")


### Contribution guidelines

Defect reports, feature requests, and pull requests are welcome on https://github.com/synesissoftware/CLASP.Go.


### Dependencies

None


#### Development/Testing Dependencies

* [**ANGoLS**](https://github.com/synesissoftware/ANGoLS/);
* [**ver2go**](https://github.com/synesissoftware/ver2go/);


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


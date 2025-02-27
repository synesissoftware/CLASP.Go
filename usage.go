// Copyright 2019-2025 Matthew Wilson and Synesis Information Systems.
// Copyright 2015-2019 Matthew Wilson. All rights reserved. Use of this
// source code is governed by a BSD-style license that can be found in the
// LICENSE file.

/*
 * Created: 4th September 2015
 * Updated: 28th February 2025
 */

package clasp

import (
	"fmt"
	"io"
	"os"
	"path"
	"reflect"
	"strings"
)

/* /////////////////////////////////////////////////////////////////////////
 * types
 */

type UsageFlag int

type Exiter interface {
	Exit(exitCode int)
}

type UsageParams struct {
	Stream        io.Writer
	ProgramName   string
	UsageFlags    UsageFlag
	ExitCode      int
	Exiter        Exiter
	Version       interface{}
	VersionPrefix string
	InfoLines     []string
	ValuesString  string
	// If the empty string is specified, then a default string is used if
	// any specifications are specified; if a whitespace-only string is specified,
	// then no flags/options element is presented
	FlagsAndOptionsString string
}

func (params UsageParams) String() string {

	return fmt.Sprintf("<%T{ Stream=%v, ProgramName=%q, UsageFlags=0x%x, ExitCode=%d, Exiter=%v, Version=%v, VersionPrefix=%q, InfoLines=%v, ValuesString=%q, FlagsAndOptionsString=%q }>", params, params.Stream, params.ProgramName, params.UsageFlags, params.ExitCode, params.Exiter, params.Version, params.VersionPrefix, params.InfoLines, params.ValuesString, params.FlagsAndOptionsString)
}

/* /////////////////////////////////////////////////////////////////////////
 * helpers
 */

const (
	SkipBlanksBetweenLines UsageFlag = 1 << iota
	DontCallExit           UsageFlag = 1 << iota
	DontCallExitIfZero     UsageFlag = 1 << iota
)

/* /////////////////////////////////////////////////////////////////////////
 * locals
 */

type default_exiter struct {
}

func (de *default_exiter) Exit(exitCode int) {

	os.Exit(exitCode)
}

func should_call_Exit(params UsageParams) bool {

	if 0 != (DontCallExit & params.UsageFlags) {

		return false
	}

	if 0 == params.ExitCode {

		if 0 != (DontCallExitIfZero & params.UsageFlags) {

			return false
		}
	}

	return true
}

func collect_array_as_strings(a []interface{}) []string {

	r := make([]string, len(a))

	for i, v := range a {

		r[i] = fmt.Sprintf("%v", v)
	}

	return r
}

func get_program_name(params UsageParams) string {

	program_name := params.ProgramName

	if 0 == len(program_name) {

		arg0 := os.Args[0]

		program_name = path.Base(arg0)
	}

	return program_name
}

func generate_version_string(params UsageParams, apiFunctionName string) string {

	program_name := get_program_name(params)
	version_prefix := params.VersionPrefix

	var version string

	switch v := params.Version.(type) {

	case string:

		version = v
	case []string:

		version = strings.Join(v, ".")
	case []int:
		as := make([]string, len(v))

		for i, n := range v {

			as[i] = fmt.Sprintf("%v", n)
		}

		version = strings.Join(as, ".")
	case []uint16:

		as := make([]string, len(v))

		for i, n := range v {

			as[i] = fmt.Sprintf("%v", n)
		}

		version = strings.Join(as, ".")
	default:

		panic(fmt.Sprintf("%v() called with UsageParams.Version of an invalid type %T, but must be instance of string, []string, []int, []uint16", v, reflect.TypeOf(v)))
	}

	return fmt.Sprintf("%s %s%s", program_name, version_prefix, version)
}

/* /////////////////////////////////////////////////////////////////////////
 * API
 */

func ShowUsage(specifications []Specification, params UsageParams) (rc int, err error) {

	for i, a := range specifications {

		switch a.Type {

		case FlagType, OptionType:

		case SectionType:

		default:

			panic(fmt.Sprintf("specification[%d] - '%v' - is an instance of type %T, but must be instance of either %T or %T!", i, a, a, FlagType, OptionType))
		}
	}

	exiter := params.Exiter

	if exiter == nil {

		exiter = new(default_exiter)
	}

	if params.Stream == nil {

		if 0 == params.ExitCode {

			params.Stream = os.Stdout
		} else {

			params.Stream = os.Stderr
		}
	}

	program_name := get_program_name(params)

	if "" == params.FlagsAndOptionsString && 0 != len(specifications) {

		params.FlagsAndOptionsString = "[ ... flags and options ... ]"
	}

	if "" != strings.TrimSpace(params.FlagsAndOptionsString) {

		params.FlagsAndOptionsString = " " + params.FlagsAndOptionsString
	} else {

		params.FlagsAndOptionsString = ""
	}

	if "" != params.ValuesString {

		params.ValuesString = " " + params.ValuesString
	}

	for _, info_line := range params.InfoLines {

		if ":version:" == info_line {

			version := generate_version_string(params, "ShowUsage")

			fmt.Fprintf(params.Stream, "%s\n", version)
		} else {

			fmt.Fprintf(params.Stream, "%s\n", info_line)
		}
	}

	fmt.Fprintf(params.Stream, "USAGE: %s%s%s\n", program_name, params.FlagsAndOptionsString, params.ValuesString)

	if 0 != len(specifications) {

		fmt.Fprintf(params.Stream, "\n")
		fmt.Fprintf(params.Stream, "flags/options:\n")
		if 0 == (SkipBlanksBetweenLines & params.UsageFlags) {

			fmt.Fprintf(params.Stream, "\n")
		}

		voas := make(map[string][]Specification)
		pure := make([]Specification, 0)

		for _, a := range specifications {

			ix_eq := strings.Index(a.Name, "=")

			if ix_eq < 0 {

				pure = append(pure, a)
			} else {

				name := a.Name[0:ix_eq]

				if _, ok := voas[name]; !ok {

					voas[name] = make([]Specification, 0)
				}

				voas[name] = append(voas[name], a)
			}
		}

		for _, a := range pure {

			switch a.Type {

			case FlagType:

				for _, b := range a.Aliases {

					fmt.Fprintf(params.Stream, "\t%v\n", b)
				}
				fmt.Fprintf(params.Stream, "\t%v\n", a.Name)

			case OptionType:

				for _, c := range voas[a.Name] {

					fmt.Fprintf(params.Stream, "\t%v %v\n", c.Aliases[0], c.Name)
				}
				for _, b := range a.Aliases {

					fmt.Fprintf(params.Stream, "\t%v <value>\n", b)
				}
				fmt.Fprintf(params.Stream, "\t%v=<value>\n", a.Name)

			case SectionType:

				if 0 != (SkipBlanksBetweenLines & params.UsageFlags) {

					fmt.Fprintf(params.Stream, "\n")
				}
				fmt.Fprintf(params.Stream, "\t%v\n\n", a.Name)

				continue
			}

			if 0 != len(a.Help) {

				fmt.Fprintf(params.Stream, "\t\t%v\n", a.Help)
			}

			if 0 != len(a.ValueSet) {

				fmt.Fprintf(params.Stream, "\t\twhere <value> one of:\n")
				for j := 0; j != len(a.ValueSet); j++ {

					fmt.Fprintf(params.Stream, "\t\t\t%v\n", a.ValueSet[j])
				}
			}

			if 0 == (SkipBlanksBetweenLines & params.UsageFlags) {

				fmt.Fprintf(params.Stream, "\n")
			}
		}
	}

	if should_call_Exit(params) {

		exiter.Exit(params.ExitCode)
	}

	return params.ExitCode, nil
}

func ShowVersion(specifications []Specification, params UsageParams) (rc int, err error) {

	exiter := params.Exiter

	if params.Stream == nil {

		if 0 == params.ExitCode {

			params.Stream = os.Stdout
		} else {

			params.Stream = os.Stderr
		}
	}

	if exiter == nil {

		exiter = new(default_exiter)
	}

	version := generate_version_string(params, "ShowVersion")

	fmt.Fprintf(params.Stream, "%s\n", version)

	if should_call_Exit(params) {

		exiter.Exit(params.ExitCode)
	}

	return
}

/* ///////////////////////////// end of file //////////////////////////// */

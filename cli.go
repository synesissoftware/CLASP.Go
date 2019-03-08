/* /////////////////////////////////////////////////////////////////////////
 * File:        clasp/cli.go
 *
 * Purpose:     CLI utilities for CLASP.Go
 *
 * Created:     4th September 2015
 * Updated:     8th March 2019
 *
 * Home:        http://synesis.com.au/software
 *
 * Copyright (c) 2015-2019, Matthew Wilson
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 * - Redistributions of source code must retain the above copyright notice,
 *   this list of conditions and the following disclaimer.
 * - Redistributions in binary form must reproduce the above copyright
 *   notice, this list of conditions and the following disclaimer in the
 *   documentation and/or other materials provided with the distribution.
 * - Neither the names of Matthew Wilson and Synesis Software nor the names
 *   of any contributors may be used to endorse or promote products derived
 *   from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS
 * IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO,
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
 * PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR
 * CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
 * PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
 * LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
 * NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 * ////////////////////////////////////////////////////////////////////// */


package clasp

import (

	"fmt"
	"io"
	"os"
	"reflect"
	"path"
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

	Stream					io.Writer
	ProgramName				string
	UsageFlags				UsageFlag
	ExitCode				int
	Exiter					Exiter
	Version					interface{}
	VersionPrefix			string
	InfoLines				[]string
	ValuesString			string
	// If the empty string is specified, then a default string is used if
	// any aliases are specified; if a whitespace-only string is specified,
	// then no flags/options element is presented
	FlagsAndOptionsString	string
}

func (params UsageParams) String() string {

	return fmt.Sprintf("<%T{ Stream=%v, ProgramName=%q, UsageFlags=0x%x, ExitCode=%d, Exiter=%v, Version=%v, VersionPrefix=%q, InfoLines=%v, ValuesString=%q, FlagsAndOptionsString=%q }>", params, params.Stream, params.ProgramName, params.UsageFlags, params.ExitCode, params.Exiter, params.Version, params.VersionPrefix, params.InfoLines, params.ValuesString, params.FlagsAndOptionsString)
}

/* /////////////////////////////////////////////////////////////////////////
 * helpers
 */

const (

	SkipBlanksBetweenLines  UsageFlag = 1 << iota
	DontCallExit            UsageFlag = 1 << iota
	DontCallExitIfZero      UsageFlag = 1 << iota
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

	r	:=	make([]string, len(a))

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

	program_name	:=	get_program_name(params)
	version_prefix	:=	params.VersionPrefix

	var version string;

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
	default:

		panic(fmt.Sprintf("%v() called with UsageParams.Version of an invalid type %T, but must be instance of string, []string, or []int", v, reflect.TypeOf(v)))
	}

	return fmt.Sprintf("%s %s%s", program_name, version_prefix, version)
}

/* /////////////////////////////////////////////////////////////////////////
 * API
 */

func ShowUsage(aliases []Alias, params UsageParams) (rc int, err error) {

	for i, a := range aliases {

		switch a.Type {

		case Flag, Option:

			;

		default:

			panic(fmt.Sprintf("alias[%d] - '%v' - is an instance of type %T, but must be instance of either %T or %T!", i, a, a, Flag, Option))
		}
	}

	if params.Stream == nil {

		if 0 == params.ExitCode {

			params.Stream = os.Stdout
		} else {

			params.Stream = os.Stderr
		}
	}

	program_name	:=	get_program_name(params)

	if "" == params.FlagsAndOptionsString && 0 != len(aliases) {

		params.FlagsAndOptionsString = "[ ... flags and options ... ]"
	}

	if "" != strings.TrimSpace(params.FlagsAndOptionsString) {

		params.FlagsAndOptionsString = " " + params.FlagsAndOptionsString
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

	if 0 != len(aliases) {

		fmt.Fprintf(params.Stream, "\n")
		fmt.Fprintf(params.Stream, "flags/options:\n")
		fmt.Fprintf(params.Stream, "\n")

		for _, a := range aliases {

			switch a.Type {

			case Flag:

				for _, b := range a.Aliases {

					fmt.Fprintf(params.Stream, "\t%v\n", b)
				}
				fmt.Fprintf(params.Stream, "\t%v\n", a.Name)

			case Option:

				for _, b := range a.Aliases {

					fmt.Fprintf(params.Stream, "\t%v <value>\n", b)
				}
				fmt.Fprintf(params.Stream, "\t%v=<value>\n", a.Name)
			}
			fmt.Fprintf(params.Stream, "\t\t%v\n", a.Help)
			if 0 == (SkipBlanksBetweenLines & params.UsageFlags) {

				fmt.Fprintf(params.Stream, "\n")
			}
		}
	}

	if should_call_Exit(params) {

		os.Exit(params.ExitCode)
	}

	return params.ExitCode, nil
}

func ShowVersion(aliases []Alias, params UsageParams) (rc int, err error) {

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



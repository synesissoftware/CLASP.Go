/* /////////////////////////////////////////////////////////////////////////
 * File:        clasp/cli.go
 *
 * Purpose:     CLASP library in Go
 *
 * Created:     4th September 2015
 * Updated:     4th September 2015
 *
 * Home:        http://synesis.com.au/software
 *
 * Copyright (c) 2015, Matthew Wilson
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

/*
	"fmt"
	"path"
	"strings"
 */
)

/* /////////////////////////////////////////////////////////////////////////
 * types
 */

type UsageFlag	int

type UsageParams struct {

	Aliases		[]Alias
	Stream		io.Writer
	ProgramName	string
	Flags		UsageFlag
	Values		string
	ExitCode	int
}

/* /////////////////////////////////////////////////////////////////////////
 * helpers
 */

const (

	SkipBlanksBetweenLines	UsageFlag	=	1 << iota
	CallExitWhenDone		UsageFlag	=	1 << iota
	CallExitIfNoneZero		UsageFlag	=	1 << iota
)

/* /////////////////////////////////////////////////////////////////////////
 * API
 */

func ShowUsage(arguments *Arguments, params UsageParams) (rc int, err error) {

	aliases := arguments.aliases_

	for i, a := range aliases {

		switch a.Type {
		case Flag, Option:
			;
		default:
			panic(fmt.Sprintf("alias[%d] - '%v' - is an instance of type %T, and must be instance of either %T or %T!", i, a, a, Flag, Option))
		}
	}

	if params.Stream == nil {

		if 0 == params.ExitCode {

			params.Stream = os.Stdout
		} else {

			params.Stream = os.Stderr
		}
	}

	if 0 == len(params.ProgramName) {

		params.ProgramName = arguments.ProgramName
	}


	fmt.Fprintf(params.Stream, "USAGE: %s [ ... flags and options ... ] %s\n", params.ProgramName, params.Values)
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
		if 0 == (SkipBlanksBetweenLines & params.Flags) {

			fmt.Fprintf(params.Stream, "\n")
		}
	}


	if 0 != (CallExitWhenDone & params.Flags) || (0 != params.ExitCode && 0 != (CallExitIfNoneZero & params.Flags)) {

		os.Exit(params.ExitCode)
	}

	return params.ExitCode, nil
}

/* ///////////////////////////// end of file //////////////////////////// */



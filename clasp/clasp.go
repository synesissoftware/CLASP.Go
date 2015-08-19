/* /////////////////////////////////////////////////////////////////////////
 * File:        clasp/clasp.go
 *
 * Purpose:     CLASP library in Go
 *
 * Created:     15th August 2015
 * Updated:     19th August 2015
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
	"path"
	"strings"
)

const (
	VersionMajor int16		=	0
	VersionMinor int16		=	3
	VersionRevision int16	=	1
	Version int64			=	(int64(VersionMajor) << 48) + (int64(VersionMinor) << 32) + (int64(VersionRevision) << 16)
)

/* /////////////////////////////////////////////////////////////////////////
 * constants
 */

type ParseFlag int

const (
	ParseTreatSingleHyphenAsValue ParseFlag = 1 << iota
	ParseDontRecogniseDoubleHyphenToStartValues ParseFlag = 1 << iota
)

/* /////////////////////////////////////////////////////////////////////////
 * types
 */

type ArgType int

const (
	Flag ArgType = 1
	Option ArgType = 2
	Value ArgType = 3
)

type Alias struct {
	Type			ArgType
	Name			string
	Aliases			[]string
	Help			string
	ValueSet		[]string
	BitFlags		int
}

type Argument struct {
	ResolvedName	string
	GivenName		string
	Value			string
	Type			ArgType
	CmdLineIndex	int
	NumGivenHyphens	int
	AliasIndex		int
	Flags			int
}

type Arguments struct {
	Arguments	[]*Argument
	Flags		[]*Argument
	Options		[]*Argument
	Values		[]*Argument
	Argv		[]string
	ProgramName	string
}

type ParseParams struct {
	Aliases		[]Alias
	Flags		ParseFlag
}

/* /////////////////////////////////////////////////////////////////////////
 * helpers
 */

func (params ParseParams) findAlias(name string) (found bool, alias Alias) {

	// Algorithm:
	//
	// 1. search for alias with that name
	// 2. search for alias with that alias
	// 3. return nil

	for _, a := range params.Aliases {
		if name == a.Name {
			return true, a
		}
	}

	for _, a := range params.Aliases {
		for _, n := range a.Aliases {
			if name == n {
				return true, a
			}
		}
	}

	var dummy Alias

	return false, dummy
}

/* /////////////////////////////////////////////////////////////////////////
 * API
 */

func Parse(argv []string, params ParseParams) *Arguments {
	args := new(Arguments)

	args.Arguments		=	make([]*Argument, 0)
	args.Flags			=	make([]*Argument, 0)
	args.Options		=	make([]*Argument, 0)
	args.Values			=	make([]*Argument, 0)
	args.Argv			=	argv
	args.ProgramName	=	path.Base(argv[0])

	treatingAsValues	:=	false
	nextIsOptValue		:=	false

	for i, s := range argv[1:] {

		if !treatingAsValues && "--" == s && (0 == (params.Flags & ParseDontRecogniseDoubleHyphenToStartValues)) {
			treatingAsValues = true
			continue
		}

		if nextIsOptValue {
			nextIsOptValue = false
			args.Arguments[len(args.Arguments) - 1].Value = s
			continue
		}

		arg := new(Argument)


		arg.CmdLineIndex	=	i + 1
		arg.Flags			=	int(params.Flags)
		arg.AliasIndex		=	-1

		numHyphens			:=	0
		isSingle			:=	false

		if !treatingAsValues {
			l				:=	len(s)
			if 1 == l && "-" == s {
				numHyphens	=	1
				isSingle	=	true
			} else if 2 == l && "--" == s {
				numHyphens	=	2
			} else {
				numHyphens	=	strings.IndexFunc(s, func(c rune) bool { return '-' != c })
			}
		}


		arg.NumGivenHyphens	=	numHyphens

		switch numHyphens {

			case 0:
				arg.Type				=	Value
				arg.Value				=	s
			default:
				nv					:=	strings.SplitN(s, "=", 2)
				if len(nv) > 1 {
					arg.Type			=	Option
					arg.GivenName		=	nv[0]
					arg.ResolvedName	=	nv[0]
					arg.Value			=	nv[1]
				} else {
					// Here we have to be flexible, and examine
					// whether the apparent flag is, in fact, a
					// an option

					resolvedName		:=	s
					argType				:=	Flag

					if found, alias := params.findAlias(s); found {
						resolvedName	=	alias.Name
						argType			=	alias.Type
					}

					arg.Type			=	argType
					arg.GivenName		=	s
					arg.ResolvedName	=	resolvedName

					if Option == argType {
						nextIsOptValue	=	true
					} else {
						if isSingle	&& (0 != (params.Flags & ParseTreatSingleHyphenAsValue)) {
							arg.Type	=	Value
							arg.Value	=	s
						}
					}
				}
		}

		args.Arguments	=	append(args.Arguments, arg)
	}

	for _, arg := range args.Arguments {

		switch(arg.Type) {
			case Flag:
				args.Flags		=	append(args.Flags, arg)
			case Option:
				args.Options	=	append(args.Options, arg)
			case Value:
				args.Values		=	append(args.Values, arg)
		}
	}

	return args
}

func (args *Arguments) FlagIsSpecified(id interface{}) bool {

	name	:=	""
	found	:=	false

	if s, is_string := id.(string); is_string {
		name	=	s
		found	=	true
	}

	if a, is_Alias := id.(Alias); is_Alias {
		switch a.Type {
			case Option:
				// TODO: issue warning
				fallthrough
			case Flag:
				name	=	a.Name
				found	=	true
			default:
				panic(fmt.Sprintf("invoked FlagIsSpecified() passing a non-Flag (and non-Option) Alias '%v'", a))
		}
	}

	if !found && nil != id {
		panic(fmt.Sprintf("invoked FlagIsSpecified() passing a value - '%v' - that is neither string nor alias", id))
	}

	for i, f := range args.Flags {
		// TODO: mark as used
		_ = i
		if name == f.ResolvedName {
			return true
		}
	}

	return false
}

func (args *Arguments) LookupOption(id interface{}) (*Argument, bool) {

	name	:=	""
	found	:=	false

	if s, is_string := id.(string); is_string {
		name	=	s
		found	=	true
	}

	if a, is_Alias := id.(Alias); is_Alias {
		switch a.Type {
			case Option:
				name	=	a.Name
				found	=	true
			default:
				panic(fmt.Sprintf("invoked LookupOption() passing a non-Option Alias '%v'", a))
		}
	}

	if !found && nil != id {
		panic(fmt.Sprintf("invoked LookupOption() passing a value - '%v' - that is neither string nor alias", id))
	}

	for i, o := range args.Options {
		// TODO: mark as used
		_ = i
		if name == o.ResolvedName {
			return o, true
		}
	}

	return nil, false
}

/* ///////////////////////////// end of file //////////////////////////// */


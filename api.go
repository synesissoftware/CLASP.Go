/* /////////////////////////////////////////////////////////////////////////
 * File:        api.go
 *
 * Purpose:     Main file for CLASP.Go
 *
 * Created:     15th August 2015
 * Updated:     28th March 2019
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
	"path"
	"strings"
)

/* /////////////////////////////////////////////////////////////////////////
 * types
 */

type ParseFlag int

/* /////////////////////////////////////////////////////////////////////////
 * constants
 */

const (

	ParseTreatSingleHyphenAsValue ParseFlag = 1 << iota
	ParseDontRecogniseDoubleHyphenToStartValues ParseFlag = 1 << iota
)

/* /////////////////////////////////////////////////////////////////////////
 * types
 */

type ArgType int

const (

	FlagType		ArgType = 1
	OptionType		ArgType = 2
	ValueType		ArgType = 3

	optionViaAlias	ArgType = -98
	int_1_			ArgType = -99
)

type Alias struct {

	Type			ArgType
	Name			string
	Aliases			[]string
	Help			string
	ValueSet		[]string
	BitFlags		int
	Extras			map[string]interface{}
}

type Argument struct {

	ResolvedName	string
	GivenName		string
	Value			string
	Type			ArgType
	CmdLineIndex	int
	NumGivenHyphens	int
	ArgumentAlias	*Alias
	Flags			int

	used_			int
}

type Arguments struct {

	Arguments	[]*Argument
	Flags		[]*Argument
	Options		[]*Argument
	Values		[]*Argument
	Argv		[]string
	ProgramName	string
	aliases_	[]*Alias
}

type ParseParams struct {

	Aliases		[]Alias
	Flags		ParseFlag
}

// Obtains, by value, an Alias containing a stock specification of a '--help' flag
func HelpFlag() Alias {

	return Alias{ FlagType, "--help", nil, "Shows this helps and exits", nil, 0, nil }
}

// Obtains, by value, an Alias containing a stock specification of a '--version' flag
func VersionFlag() Alias {

	return Alias{ FlagType, "--version", nil, "Shows version information and exits", nil, 0, nil }
}

func (at ArgType) String() string {

	switch(at) {

	case FlagType:

		return "Flag"
	case OptionType, optionViaAlias:

		return "Option"
	case ValueType:

		return "Value"
	default:

		return fmt.Sprintf("<%T %d>", at, at)
	}
}

func (alias Alias) String() string {

	return fmt.Sprintf("<%T{ Type=%v, Name=%q, Aliases=%v, Help=%q, ValueSet=%v, BitFlags=0x%x, Extras=%v }>", alias, alias.Type, alias.Name, alias.Aliases, alias.Help, alias.ValueSet, alias.BitFlags, alias.Extras)
}

func (argument Argument) String() string {

	return fmt.Sprintf("<%T{ ResolvedName=%q, GivenName=%q, Value=%q, Type=%v, CmdLineIndex=%d, NumGivenHyphens=%d, ArgumentAlias=%v, Flags=0x%x, used=%t }>", argument, argument.ResolvedName, argument.GivenName, argument.Value, argument.Type, argument.CmdLineIndex, argument.NumGivenHyphens, argument.ArgumentAlias, argument.Flags, argument.used_ != 0)
}

func (arguments Arguments) String() string {

	return fmt.Sprintf("<%T{ Arguments=%v, Flags=%v, Options=%v, Values=%v, Argv=%v, ProgramName=%q }>", arguments, arguments.Arguments, arguments.Flags, arguments.Options, arguments.Values, arguments.Argv, arguments.ProgramName)
}

func (params ParseParams) String() string {

	return fmt.Sprintf("<%T{ Aliases=%v, Flags=0x%x }>", params, params.Aliases, params.Flags)
}

/* builders */

// Creates a flag alias, with the given name
func Flag(name string) (result Alias) {

	result.Type = FlagType
	result.Name = name

	return
}

// Creates an option alias, with the given name
func Option(name string) (result Alias) {

	result.Type = OptionType
	result.Name = name

	return
}

// Creates an alias for an actual flag/option, with 1 or more aliases
func AliasesFor(actual string, alias0 string, other_aliases ...string) (result Alias) {

	result.Type = FlagType
	result.Name = actual
	result.Aliases = append([]string { alias0 }, other_aliases...)

	return
}

// Builder method to set the help for an alias
func (alias Alias) SetHelp(help string) (Alias) {

	alias.Help = help

	return alias
}

// Builder method to set the values for an option alias
func (alias Alias) SetValues(values ...string) (Alias) {

	alias.ValueSet = values

	return alias
}

// Sets the alias
func (alias Alias) SetAlias(s string) (Alias) {

	alias.Aliases = []string { s }

	return alias
}

// Sets one or more aliases
func (alias Alias) SetAliases(aliases ...string) (Alias) {

	alias.Aliases = aliases

	return alias
}

// Builder method to an Extras entry
func (alias Alias) SetExtra(key string, value interface{}) (Alias) {

	if alias.Extras == nil {

		alias.Extras = make(map[string]interface{})
	}

	alias.Extras[key] = value

	return alias
}

/* /////////////////////////////////////////////////////////////////////////
 * helpers
 */

func (params *ParseParams) findAlias(name string) (found bool, alias *Alias, aliasIndex int) {

	// Algorithm:
	//
	// 1. search for alias with that name
	// 2. search for alias with that alias
	// 3. return nil

	for i, a := range params.Aliases {

		if name == a.Name {

			return true, &a, i
		}
	}

	for i, a := range params.Aliases {

		for _, n := range a.Aliases {

			if name == n {

				return true, &a, i
			}
		}
	}

	return false, nil, -1
}

/* /////////////////////////////////////////////////////////////////////////
 * API
 */

func (arg *Argument) Use() {

	arg.used_ = 1
}

func (arg Argument) Str() string {

	switch arg.Type {

	case OptionType:

		return fmt.Sprintf("%s=%s", arg.ResolvedName, arg.Value)
	default:

		return arg.ResolvedName
	}
}

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
		arg.ArgumentAlias	=	nil

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

			arg.Type				=	ValueType
			arg.Value				=	s
		default:

			nv := strings.SplitN(s, "=", 2)
			if len(nv) > 1 {

				arg.Type			=	OptionType
				arg.GivenName		=	nv[0]
				arg.ResolvedName	=	nv[0]
				arg.Value			=	nv[1]

				// TODO: handle "-<opt-alias>=value"
			} else {

				// Here we have to be flexible, and examine
				// whether the apparent flag is, in fact, an
				// option

				resolvedName		:=	s
				argType				:=	FlagType

				if found, alias, _ := params.findAlias(s); found {

					resolvedName		=	alias.Name
					argType				=	alias.Type
					arg.ArgumentAlias	=	alias

					if ix_equals := strings.Index(resolvedName, "="); ix_equals >= 0 {

						res_nm	:=	resolvedName[:ix_equals]
						value	:=	resolvedName[ix_equals + 1:]

						argType			=	optionViaAlias
						s				=	resolvedName
						resolvedName	=	res_nm
						arg.Value		=	value
					}
				} else {

					// Now we test to see whether every character yields
					// an alias. If so, we convert all, add them in, then
					// skip to the next input

					validCompoundFlag	:=	len(s) > 1

					compoundArguments	:=	make([]*Argument, 0, len(s) - 1)

					for i, c := range s {

						if 0 == i {

							continue
						}

						testAlias	:=	fmt.Sprintf("-%c", c)

						if compoundFound, compoundAlias, _ := params.findAlias(testAlias); compoundFound && compoundAlias.Type == FlagType {

							var compoundArg Argument

							compoundArg.ResolvedName	=	compoundAlias.Name
							compoundArg.GivenName		=	s
							compoundArg.Value			=	""
							compoundArg.Type			=	FlagType
							compoundArg.CmdLineIndex	=	arg.CmdLineIndex
							compoundArg.ArgumentAlias	=	compoundAlias
							compoundArg.Flags			=	arg.Flags

							if ix_equals := strings.Index(compoundArg.ResolvedName, "="); ix_equals >= 0 {

								res_nm	:=	compoundArg.ResolvedName[:ix_equals]
								value	:=	compoundArg.ResolvedName[ix_equals + 1:]

								compoundArg.Type			=	OptionType
								s							=	compoundArg.ResolvedName
								compoundArg.ResolvedName	=	res_nm
								compoundArg.Value			=	value

								// Now need to look up the actual underlying alias

								if actualFound, actualAlias, _ := params.findAlias(res_nm); actualFound {

									compoundArg.ArgumentAlias	=	actualAlias
								}
							}

							compoundArguments			=	append(compoundArguments, &compoundArg)
						} else {

							validCompoundFlag = false
							break
						}
					}

					if validCompoundFlag {

						args.Arguments	=	append(args.Arguments, compoundArguments...)
						continue
					}
				}

				switch argType {

				case optionViaAlias:

					arg.Type	=	OptionType
				default:

					arg.Type	=	argType
				}

				arg.GivenName		=	s
				arg.ResolvedName	=	resolvedName

				if OptionType == arg.Type {

					if optionViaAlias != argType {

						nextIsOptValue	=	true
					}
				} else {

					if isSingle	&& (0 != (params.Flags & ParseTreatSingleHyphenAsValue)) {

						arg.Type	=	ValueType
						arg.Value	=	s
					}
				}
			}
		}

		args.Arguments	=	append(args.Arguments, arg)
	}

	for _, arg := range args.Arguments {

		switch(arg.Type) {

			case FlagType:

				args.Flags		=	append(args.Flags, arg)
			case OptionType:

				args.Options	=	append(args.Options, arg)
			case ValueType:

				args.Values		=	append(args.Values, arg)
		}
	}

	args.aliases_	=	make([]*Alias, len(params.Aliases))

	for i, a := range(params.Aliases) {

		var p *Alias = new(Alias)

		*p = a

		args.aliases_[i] = p
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

			case OptionType:

				// TODO: issue warning
				fallthrough
			case FlagType:

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

			f.used_ = 1
			return true
		}
	}

	return false
}

func (args *Arguments) LookupFlag(id interface{}) (*Argument, bool) {

	name	:=	""
	found	:=	false

	if s, is_string := id.(string); is_string {

		name	=	s
		found	=	true
	}

	if a, is_Alias := id.(Alias); is_Alias {

		switch a.Type {

			case FlagType:

				name	=	a.Name
				found	=	true
			default:

				panic(fmt.Sprintf("invoked LookupFlag() passing a non-Flag Alias '%v'", a))
		}
	}

	if !found && nil != id {

		panic(fmt.Sprintf("invoked LookupFlag() passing a value - '%v' - that is neither string nor alias", id))
	}

	for i, o := range args.Flags {

		// TODO: mark as used
		_ = i
		if name == o.ResolvedName {

			o.used_ = 1
			return o, true
		}
	}

	return nil, false
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

			case OptionType:

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

			o.used_ = 1
			return o, true
		}
	}

	return nil, false
}

func (args *Arguments) GetUnusedFlags() []*Argument {

	var unused []*Argument

	for _, f := range args.Flags {

		if 0 == f.used_ {

			unused = append(unused, f)
		}
	}

	return unused
}

func (args *Arguments) GetUnusedOptions() []*Argument {

	var unused []*Argument

	for _, o := range args.Options {

		if 0 == o.used_ {

			unused = append(unused, o)
		}
	}

	return unused
}

func (args *Arguments) GetUnusedFlagsAndOptions() []*Argument {

	var unused []*Argument

	for _, a := range args.Arguments {

		switch a.Type {

			case FlagType:

				fallthrough
			case OptionType:

				if 0 == a.used_ {

					unused = append(unused, a)
				}
				break;
		}
	}

	return unused
}

func check_flag_bits(args *Arguments, flags *int, only_unused bool) int {

	var dummy_ int

	if nil == flags {

		flags = &dummy_
	}

	*flags = 0

	for _, arg := range args.Flags {

		if !only_unused || 0 == arg.used_ {

			for _, al := range args.aliases_ {

				if al.Name == arg.ResolvedName {

					*flags |= al.BitFlags
					if only_unused {

						arg.used_ = 1
					}
				}
			}
		}
	}

	return *flags
}

// Examines the unused parsed flags held by the Arguments instance and
// combines the BitFlags values of their corresponding aliases.
//
// NOTE: Marks any of the flags as used
func (args *Arguments) CheckUnusedFlagBits(flags *int) int {

	return check_flag_bits(args, flags, true)
}

// Examines all parsed flags held by the Arguments instance and combines
// the BitFlags values of their corresponding aliases.
//
// NOTE: Does NOT mark any of the flags as used
func (args *Arguments) CheckAllFlagBits(flags *int) int {

	return check_flag_bits(args, flags, false)
}

func Aliases(aliases...string) []string {

	r := make([]string, len(aliases))

	for i, alias := range aliases {

		r[i] = alias
	}

	return r
}

/* ///////////////////////////// end of file //////////////////////////// */



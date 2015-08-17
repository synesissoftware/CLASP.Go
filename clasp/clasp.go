
package clasp

import (
	"fmt"
	"path"
	"strings"
)

type ParseFlag int

const (
	ParseTreatSingleHyphenAsValue ParseFlag = 1 << iota
)

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

		if !treatingAsValues && "--" == s {
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
			if 1 == len(s) && "-" == s {
				numHyphens	=	1
				isSingle	=	true
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


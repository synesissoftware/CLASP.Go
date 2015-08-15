
package clasp

import (
	"path"
	"strings"
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
	MappedArgument	string
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
	Aliases	[]Alias
	Flags	int
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

	for i, s := range argv[1:] {

		if !treatingAsValues && "--" == s {
			treatingAsValues = true
			continue
		}

		arg := new(Argument)

		arg.CmdLineIndex	=	i + 1
		arg.Flags			=	params.Flags
		arg.AliasIndex		=	-1

		numHyphens			:=	0

		if !treatingAsValues {
			numHyphens		=	strings.IndexFunc(s, func(c rune) bool { return '-' != c })
		}

		arg.NumGivenHyphens	=	numHyphens

		arg.Value			=	s
		arg.Type			=	Value

		args.Arguments	=	append(args.Arguments, arg)
	}

	args.Values			=	args.Arguments

	return args
}


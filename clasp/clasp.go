
package clasp

import (
	"path"
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

	for i, s := range argv[1:] {

		arg := new(Argument)

		arg.Value			=	s
		arg.Type			=	Value
		arg.CmdLineIndex	=	i + 1

		args.Arguments	=	append(args.Arguments, arg)
	}

	args.Values			=	args.Arguments

	return args
}


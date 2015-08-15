
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
	a := new(Arguments)

	a.Arguments		=	make([]*Argument, 0)
	a.Flags			=	make([]*Argument, 0)
	a.Options		=	make([]*Argument, 0)
	a.Values		=	make([]*Argument, 0)
	a.Argv			=	argv
	a.ProgramName	=	path.Base(argv[0])


	return a
}


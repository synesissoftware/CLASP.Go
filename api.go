// Copyright 2019-2025 Matthew Wilson and Synesis Information Systems.
// Copyright 2015-2019 Matthew Wilson. All rights reserved. Use of this
// source code is governed by a BSD-style license that can be found in the
// LICENSE file.

/*
 * Created: 15th August 2015
 * Updated: 28th February 2025
 */

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
	ParseTreatSingleHyphenAsValue               ParseFlag = 1 << iota
	ParseDontRecogniseDoubleHyphenToStartValues ParseFlag = 1 << iota
)

/* /////////////////////////////////////////////////////////////////////////
 * types
 */

type ArgType int

const (
	FlagType   ArgType = 1
	OptionType ArgType = 2
	ValueType  ArgType = 3

	SectionType ArgType = 21

	optionViaAlias ArgType = -98
	int_1_         ArgType = -99
)

type Specification struct {
	Type     ArgType
	Name     string
	Aliases  []string
	Help     string
	ValueSet []string
	BitFlags int
	Extras   map[string]interface{}
}

type Argument struct {
	ResolvedName          string
	GivenName             string
	Value                 string
	Type                  ArgType
	CmdLineIndex          int
	NumGivenHyphens       int
	ArgumentSpecification *Specification
	Flags                 int

	used_ int
}

type Arguments struct {
	Arguments       []*Argument
	Flags           []*Argument
	Options         []*Argument
	Values          []*Argument
	Argv            []string
	ProgramName     string
	specifications_ []*Specification
}

type ParseParams struct {
	Specifications []Specification
	Flags          ParseFlag
}

// Obtains, by value, a specification containing a stock specification of a '--help' flag.
func HelpFlag() Specification {

	return Specification{FlagType, "--help", nil, "Shows this help and exits", nil, 0, nil}
}

// Obtains, by value, a specification containing a stock specification of a '--version' flag.
func VersionFlag() Specification {

	return Specification{FlagType, "--version", nil, "Shows version information and exits", nil, 0, nil}
}

func (at ArgType) String() string {

	switch at {

	case FlagType:

		return "Flag"
	case OptionType, optionViaAlias:

		return "Option"
	case ValueType:

		return "Value"
	case SectionType:

		return "Section"
	default:

		return fmt.Sprintf("<%T %d>", at, at)
	}
}

func (specification Specification) String() string {

	return fmt.Sprintf("<%T{ Type=%v, Name=%q, Aliases=%v, Help=%q, ValueSet=%v, BitFlags=0x%x, Extras=%v }>", specification, specification.Type, specification.Name, specification.Aliases, specification.Help, specification.ValueSet, specification.BitFlags, specification.Extras)
}

func (argument Argument) String() string {

	return fmt.Sprintf("<%T{ ResolvedName=%q, GivenName=%q, Value=%q, Type=%v, CmdLineIndex=%d, NumGivenHyphens=%d, ArgumentSpecification=%v, Flags=0x%x, used=%t }>", argument, argument.ResolvedName, argument.GivenName, argument.Value, argument.Type, argument.CmdLineIndex, argument.NumGivenHyphens, argument.ArgumentSpecification, argument.Flags, argument.used_ != 0)
}

func (arguments Arguments) String() string {

	return fmt.Sprintf("<%T{ Arguments=%v, Flags=%v, Options=%v, Values=%v, Argv=%v, ProgramName=%q }>", arguments, arguments.Arguments, arguments.Flags, arguments.Options, arguments.Values, arguments.Argv, arguments.ProgramName)
}

func (params ParseParams) String() string {

	return fmt.Sprintf("<%T{ Specifications=%v, Flags=0x%x }>", params, params.Specifications, params.Flags)
}

/* builders */

// Creates a flag specification, with the given name.
func Flag(name string) (result Specification) {

	result.Type = FlagType
	result.Name = name

	return
}

// Creates an option specification, with the given name.
func Option(name string) (result Specification) {

	result.Type = OptionType
	result.Name = name

	return
}

// Creates a flag specification, with the given name.
func Section(name string) (result Specification) {

	result.Type = SectionType
	result.Name = name

	return
}

// Creates a specification for an actual flag/option, with 1 or more aliases.
func AliasesFor(actual string, alias0 string, other_aliases ...string) (result Specification) {

	result.Type = FlagType
	result.Name = actual
	result.Aliases = append([]string{alias0}, other_aliases...)

	return
}

// Builder method to set the help for a specification.
func (specification Specification) SetHelp(help string) Specification {

	specification.Help = help

	return specification
}

// Builder method to set the values for an option specification.
func (specification Specification) SetValues(values ...string) Specification {

	specification.ValueSet = values

	return specification
}

// Builder method that sets the alias.
func (specification Specification) SetAlias(s string) Specification {

	specification.Aliases = []string{s}

	return specification
}

// Builder method that sets one or more aliases.
func (specification Specification) SetAliases(aliases ...string) Specification {

	specification.Aliases = aliases

	return specification
}

// Builder method to set an Extras entry.
func (specification Specification) SetExtra(key string, value interface{}) Specification {

	if specification.Extras == nil {

		specification.Extras = make(map[string]interface{})
	}

	specification.Extras[key] = value

	return specification
}

// Builder method to denote the end of building. All uses of building should
// be using this method, as future versions may change the types/semantics
// of the other builder methods.
func (specification Specification) End() Specification {

	return specification
}

/* /////////////////////////////////////////////////////////////////////////
 * helpers
 */

func (params *ParseParams) findSpecification(name string) (found bool, specification *Specification, specificationIndex int) {

	// Algorithm:
	//
	// 1. search for specification with that name
	// 2. search for specification with that alias
	// 3. return nil

	for i, spec := range params.Specifications {

		if name == spec.Name {

			return true, &spec, i
		}
	}

	for i, spec := range params.Specifications {

		for _, n := range spec.Aliases {

			if name == n {

				return true, &spec, i
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

	args.Arguments = make([]*Argument, 0)
	args.Flags = make([]*Argument, 0)
	args.Options = make([]*Argument, 0)
	args.Values = make([]*Argument, 0)
	args.Argv = argv
	args.ProgramName = path.Base(argv[0])

	treatingAsValues := false
	nextIsOptValue := false

	for i, s := range argv[1:] {

		if !treatingAsValues && "--" == s && (0 == (params.Flags & ParseDontRecogniseDoubleHyphenToStartValues)) {

			treatingAsValues = true
			continue
		}

		if nextIsOptValue {

			nextIsOptValue = false
			args.Arguments[len(args.Arguments)-1].Value = s
			continue
		}

		arg := new(Argument)

		arg.CmdLineIndex = i + 1
		arg.Flags = int(params.Flags)
		arg.ArgumentSpecification = nil

		numHyphens := 0
		isSingle := false

		if !treatingAsValues {

			l := len(s)
			if 1 == l && "-" == s {

				numHyphens = 1
				isSingle = true
			} else if 2 == l && "--" == s {

				numHyphens = 2
			} else {

				numHyphens = strings.IndexFunc(s, func(c rune) bool { return '-' != c })
			}
		}

		arg.NumGivenHyphens = numHyphens

		switch numHyphens {

		case 0:

			arg.Type = ValueType
			arg.Value = s
		default:

			nv := strings.SplitN(s, "=", 2)
			if len(nv) > 1 {

				arg.Type = OptionType
				arg.GivenName = nv[0]
				arg.ResolvedName = nv[0]
				arg.Value = nv[1]

				if found, specification, _ := params.findSpecification(arg.ResolvedName); found {

					arg.ResolvedName = specification.Name
					arg.ArgumentSpecification = specification
				} else {

				}
			} else {

				// Here we have to be flexible, and examine
				// whether the apparent flag is, in fact, an
				// option

				resolvedName := s
				argType := FlagType

				if found, specification, _ := params.findSpecification(s); found {

					resolvedName = specification.Name
					argType = specification.Type
					arg.ArgumentSpecification = specification

					if ix_equals := strings.Index(resolvedName, "="); ix_equals >= 0 {

						res_nm := resolvedName[:ix_equals]
						value := resolvedName[ix_equals+1:]

						argType = optionViaAlias
						s = resolvedName
						resolvedName = res_nm
						arg.Value = value

						// Now need to look up the actual underlying specification

						if actualFound, actualSpecification, _ := params.findSpecification(res_nm); actualFound {

							arg.ArgumentSpecification = actualSpecification
						}
					}
				} else {

					// Now we test to see whether every character yields
					// a specification. If so, we convert all, add them in, then
					// skip to the next input

					validCompoundFlag := len(s) > 1

					compoundArguments := make([]*Argument, 0, len(s)-1)

					for j, c := range s {

						if 0 == j {

							continue
						}

						testAlias := fmt.Sprintf("-%c", c)

						if compoundFound, compoundSpec, _ := params.findSpecification(testAlias); compoundFound && compoundSpec.Type == FlagType {

							var compoundArg Argument

							compoundArg.ResolvedName = compoundSpec.Name
							compoundArg.GivenName = s
							compoundArg.Value = ""
							compoundArg.Type = FlagType
							compoundArg.CmdLineIndex = arg.CmdLineIndex
							compoundArg.ArgumentSpecification = compoundSpec
							compoundArg.Flags = arg.Flags

							if ix_equals := strings.Index(compoundArg.ResolvedName, "="); ix_equals >= 0 {

								res_nm := compoundArg.ResolvedName[:ix_equals]
								value := compoundArg.ResolvedName[ix_equals+1:]

								compoundArg.Type = OptionType
								s = compoundArg.ResolvedName
								compoundArg.ResolvedName = res_nm
								compoundArg.Value = value

								// Now need to look up the actual underlying specification

								if actualFound, actualSpecification, _ := params.findSpecification(res_nm); actualFound {

									compoundArg.ArgumentSpecification = actualSpecification
								}
							}

							compoundArguments = append(compoundArguments, &compoundArg)
						} else {

							validCompoundFlag = false
							break
						}
					}

					if validCompoundFlag {

						args.Arguments = append(args.Arguments, compoundArguments...)
						continue
					}
				}

				switch argType {

				case optionViaAlias:

					arg.Type = OptionType
				default:

					arg.Type = argType
				}

				arg.GivenName = s
				arg.ResolvedName = resolvedName

				if OptionType == arg.Type {

					if optionViaAlias != argType {

						nextIsOptValue = true
					}
				} else {

					if isSingle && (0 != (params.Flags & ParseTreatSingleHyphenAsValue)) {

						arg.Type = ValueType
						arg.Value = s
					}
				}
			}
		}

		args.Arguments = append(args.Arguments, arg)
	}

	for _, arg := range args.Arguments {

		switch arg.Type {

		case FlagType:

			args.Flags = append(args.Flags, arg)
		case OptionType:

			args.Options = append(args.Options, arg)
		case ValueType:

			args.Values = append(args.Values, arg)
		}
	}

	args.specifications_ = make([]*Specification, len(params.Specifications))

	for i, spec := range params.Specifications {

		var p *Specification = new(Specification)

		*p = spec

		args.specifications_[i] = p
	}

	return args
}

func (args *Arguments) FlagIsSpecified(id interface{}) bool {

	name := ""
	found := false

	if s, is_string := id.(string); is_string {

		name = s
		found = true
	}

	if spec, is_Specification := id.(Specification); is_Specification {

		switch spec.Type {

		case OptionType:

			// TODO: issue warning
			fallthrough
		case FlagType:

			name = spec.Name
			found = true
		default:

			panic(fmt.Sprintf("invoked FlagIsSpecified() passing a non-Flag (and non-Option) Specification '%v'", spec))
		}
	}

	if !found && nil != id {

		panic(fmt.Sprintf("invoked FlagIsSpecified() passing a value - '%v' - that is neither string nor specification", id))
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

	name := ""
	found := false

	if s, is_string := id.(string); is_string {

		name = s
		found = true
	}

	if spec, is_Specification := id.(Specification); is_Specification {

		switch spec.Type {

		case FlagType:

			name = spec.Name
			found = true
		default:

			panic(fmt.Sprintf("invoked LookupFlag() passing a non-Flag Specification '%v'", spec))
		}
	}

	if !found && nil != id {

		panic(fmt.Sprintf("invoked LookupFlag() passing a value - '%v' - that is neither string nor specification", id))
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

	name := ""
	found := false

	if s, is_string := id.(string); is_string {

		name = s
		found = true
	}

	if spec, is_Specification := id.(Specification); is_Specification {

		switch spec.Type {

		case OptionType:

			name = spec.Name
			found = true
		default:

			panic(fmt.Sprintf("invoked LookupOption() passing a non-Option Specification '%v'", spec))
		}
	}

	if !found && nil != id {

		panic(fmt.Sprintf("invoked LookupOption() passing a value - '%v' - that is neither string nor specification", id))
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

	for _, spec := range args.Arguments {

		switch spec.Type {

		case FlagType:

			fallthrough
		case OptionType:

			if 0 == spec.used_ {

				unused = append(unused, spec)
			}
			break
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

			for _, al := range args.specifications_ {

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
// NOTE: Marks any of the flags as used.
func (args *Arguments) CheckUnusedFlagBits(flags *int) int {

	return check_flag_bits(args, flags, true)
}

// Examines all parsed flags held by the Arguments instance and combines
// the BitFlags values of their corresponding aliases.
//
// NOTE: Does NOT mark any of the flags as used.
func (args *Arguments) CheckAllFlagBits(flags *int) int {

	return check_flag_bits(args, flags, false)
}

func Aliases(aliases ...string) []string {

	r := make([]string, len(aliases))

	for i, alias := range aliases {

		r[i] = alias
	}

	return r
}

/* ///////////////////////////// end of file //////////////////////////// */

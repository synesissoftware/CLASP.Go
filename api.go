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
	Parse_None ParseFlag = 0
)

const (
	Parse_TreatSingleHyphenAsValue                    ParseFlag = 1 << iota // T.B.C.
	Parse_DontRecogniseDoubleHyphenToStartValues                            // T.B.C.
	Parse_DontMergeBitFlagsIntoBitFlags64                                   // Suppresses the default behaviour to mix into the `int64` result matched `int` bitFlagss (see [Specification.SetBitFlags]) when no matched `int64` bitFlagss (see [Specification.SetBitFlags64]) are specified.
	Parse_DontMarkUsedDuringParseWhenMatchingBitFlags                       // Suppresses the default behaviour to mark as used (see [Argument.Use]) flags that have been provided receiver variables in [Specification.SetBitFlags] or [Specification.SetBitFlags64].
)

const (
	ParseTreatSingleHyphenAsValue               = Parse_TreatSingleHyphenAsValue               // Deprecated: Instead use [Parse_TreatSingleHyphenAsValue].
	ParseDontRecogniseDoubleHyphenToStartValues = Parse_DontRecogniseDoubleHyphenToStartValues // Deprecated: Instead use [Parse_DontRecogniseDoubleHyphenToStartValues].
)

/* /////////////////////////////////////////////////////////////////////////
 * types
 */

// T.B.C.
type ArgType int

const (
	FlagType   ArgType = 1 // T.B.C.
	OptionType ArgType = 2 // T.B.C.
	ValueType  ArgType = 3 // T.B.C.

	SectionType ArgType = 21 // T.B.C.

	optionViaAlias ArgType = -98
	int_1_         ArgType = -99
)

// T.B.C.
type Specification struct {
	Type       ArgType
	Name       string
	Aliases    []string
	Help       string
	ValueSet   []string
	BitFlags   int
	BitFlags64 int64
	Extras     map[string]interface{}

	flags_receiver   *int
	flags64_receiver *int64
}

// T.B.C.
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

// T.B.C.
type Arguments struct {
	Arguments      []*Argument
	Flags          []*Argument
	Options        []*Argument
	Values         []*Argument
	Argv           []string
	ProgramName    string
	specifications []*Specification

	bitFlags   int
	bitFlags64 int64
}

// T.B.C.
type ParseParams struct {
	Specifications []Specification
	Flags          ParseFlag
}

// Obtains, by value, a specification containing a stock specification of a '--help' flag.
func HelpFlag() Specification {

	// TODO: reimplement in terms of [Flag] ??
	return Specification{FlagType, "--help", nil, "Shows this help and exits", nil, 0, 0, nil, nil, nil}
}

// Obtains, by value, a specification containing a stock specification of a '--version' flag.
func VersionFlag() Specification {

	// TODO: reimplement in terms of [Flag] ??
	return Specification{FlagType, "--version", nil, "Shows version information and exits", nil, 0, 0, nil, nil, nil}
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

func valueSetString(vs []string) string {

	elems := make([]string, len(vs))

	for ix, elem := range vs {
		elems[ix] = fmt.Sprintf(`"%s"`, elem)
	}

	return "[" + strings.Join(elems, ", ") + "]"
}

func (specification Specification) String() string {

	switch specification.Type {
	case FlagType:
		return fmt.Sprintf("<%T{ Type=%v, Name=%q, Aliases=%v, Help=%q, BitFlags=0x%x, flags_receiver=%p, BitFlags64=0x%x, flags64_receiver=%p, Extras=%v }>", specification, specification.Type, specification.Name, specification.Aliases, specification.Help, specification.BitFlags, specification.flags_receiver, specification.BitFlags64, specification.flags64_receiver, specification.Extras)

	case OptionType, optionViaAlias:
		return fmt.Sprintf("<%T{ Type=%v, Name=%q, Aliases=%v, Help=%q, ValueSet=%v, Extras=%v }>", specification, specification.Type, specification.Name, specification.Aliases, specification.Help, valueSetString(specification.ValueSet), specification.Extras)

	case ValueType:
		return fmt.Sprintf("<%T{ Type=%v, Extras=%v }>", specification, specification.Type, specification.Extras)

	case SectionType:
		return fmt.Sprintf("<%T{ Type=%v, Name=%q }>", specification, specification.Type, specification.Name)

	default:
		return fmt.Sprintf("<%T{ Type=%v, Name=%q, Aliases=%v, Help=%q, ValueSet=%v, BitFlags=0x%x, BitFlags64=0x%x, Extras=%v }>", specification, specification.Type, specification.Name, specification.Aliases, specification.Help, valueSetString(specification.ValueSet), specification.BitFlags, specification.BitFlags64, specification.Extras)
	}
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

// T.B.C.
func (specification Specification) SetBitFlags(bitFlags int, flags_receiver *int) (result Specification) {

	specification.BitFlags = bitFlags
	specification.flags_receiver = flags_receiver

	return specification
}

// Specifies bit flag(s) and, optionally, a flags receiver variable to be
// associated with the specification. If a flags receiver variable is given
// then a matching [Argument] will be marked as used automationally during
// parsing ([Parse]).
//
// NOTE: This is meaningful only to specifications that describes [Type] is
// [FlagType]. A future version may issue a panic if called on another
// argument type.
func (specification Specification) SetBitFlags64(bitFlags int64, flags_receiver *int64) (result Specification) {

	specification.BitFlags64 = bitFlags
	specification.flags64_receiver = flags_receiver

	return specification
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

// Marks an argument as used, such that it will not be obtained in a call to
// [Arguments.GetUnusedFlags] / [Arguments.GetUnusedOptions] /
// [Arguments.GetUnusedFlagsAndOptions].
//
// NOTE: This is meaningful only to arguments whose [Argument.Type] is
// [FlagType] or [OptionType]. A future version may issue a panic if called
// on another argument type.
func (arg *Argument) Use() {

	// TODO: switch on `FlagType` / `OptionType` and warn in other cases

	arg.used_ = 1
}

func (arg Argument) isUnused() bool {
	return 0 == arg.used_
}

// T.B.C.
func (arg Argument) Str() string {

	switch arg.Type {

	case OptionType:

		return fmt.Sprintf("%s=%s", arg.ResolvedName, arg.Value)
	default:

		return arg.ResolvedName
	}
}

// T.B.C.
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

		if !treatingAsValues && "--" == s && (0 == (params.Flags & Parse_DontRecogniseDoubleHyphenToStartValues)) {

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

	args.specifications = make([]*Specification, len(params.Specifications))

	for i, spec := range params.Specifications {

		var p *Specification = new(Specification)

		*p = spec

		args.specifications[i] = p
	}

	// now process the bit flags

	{
		for _, arg := range args.Flags {

			spec := arg.ArgumentSpecification

			if nil != spec {

				if 0 != spec.BitFlags64 {

					if nil != spec.flags64_receiver {

						*spec.flags64_receiver |= spec.BitFlags64

						if 0 == (Parse_DontMarkUsedDuringParseWhenMatchingBitFlags & params.Flags) {

							arg.Use()
						}
					}

					args.bitFlags64 |= spec.BitFlags64
				} else {
					if 0 != spec.BitFlags {

						if nil != spec.flags_receiver {

							*spec.flags_receiver |= spec.BitFlags

							if 0 == (Parse_DontMarkUsedDuringParseWhenMatchingBitFlags & params.Flags) {

								arg.Use()
							}
						}

						args.bitFlags |= spec.BitFlags

						if 0 == (Parse_DontMergeBitFlagsIntoBitFlags64 & params.Flags) {

							if nil != spec.flags64_receiver {

								*spec.flags64_receiver |= spec.BitFlags64

								if 0 == (Parse_DontMarkUsedDuringParseWhenMatchingBitFlags & params.Flags) {

									arg.Use()
								}
							}

							args.bitFlags64 |= int64(spec.BitFlags)
						}
					}
				}
			}
		}
	}

	return args
}

// T.B.C.
func (args Arguments) CheckAllBitFlags() int {

	return args.bitFlags
}

// T.B.C.
func (args Arguments) CheckAllBit64Flags() int64 {

	return args.bitFlags64
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

			f.Use()
			return true
		}
	}

	return false
}

// T.B.C.
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

			o.Use()
			return o, true
		}
	}

	return nil, false
}

// T.B.C.
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

			o.Use()
			return o, true
		}
	}

	return nil, false
}

// T.B.C.
func (args *Arguments) GetUnusedFlags() []*Argument {

	var unused []*Argument

	for _, f := range args.Flags {

		if f.isUnused() {

			unused = append(unused, f)
		}
	}

	return unused
}

// T.B.C.
func (args *Arguments) GetUnusedOptions() []*Argument {

	var unused []*Argument

	for _, o := range args.Options {

		if o.isUnused() {

			unused = append(unused, o)
		}
	}

	return unused
}

// T.B.C.
func (args *Arguments) GetUnusedFlagsAndOptions() []*Argument {

	var unused []*Argument

	for _, arg := range args.Arguments {

		switch arg.Type {

		case FlagType:

			fallthrough
		case OptionType:

			if arg.isUnused() {

				unused = append(unused, arg)
			}
			break
		}
	}

	return unused
}

// T.B.C.
func Aliases(aliases ...string) []string {

	r := make([]string, len(aliases))

	for i, alias := range aliases {

		r[i] = alias
	}

	return r
}

/* ///////////////////////////// end of file //////////////////////////// */

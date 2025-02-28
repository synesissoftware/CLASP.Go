package clasp_test

import (
	clasp "github.com/synesissoftware/CLASP.Go"
	stegol "github.com/synesissoftware/STEGoL"

	"testing"
)

// ArgType

func Test_String_of_ArgType_Flag(t *testing.T) {

	at_F := clasp.FlagType

	expected := "Flag"
	actual := at_F.String()

	if expected != actual {

		t.Errorf("expected '%s' does not equal '%s'", expected, actual)
	}
}

func Test_String_of_ArgType_Option(t *testing.T) {

	at_F := clasp.OptionType

	expected := "Option"
	actual := at_F.String()

	if expected != actual {

		t.Errorf("expected '%s' does not equal '%s'", expected, actual)
	}
}

func Test_String_of_ArgType_Value(t *testing.T) {

	at_F := clasp.ValueType

	expected := "Value"
	actual := at_F.String()

	if expected != actual {

		t.Errorf("expected '%s' does not equal '%s'", expected, actual)
	}
}

func Test_String_of_ArgType_unknown(t *testing.T) {

	var at_F clasp.ArgType = 101

	expected := "<clasp.ArgType 101>"
	actual := at_F.String()

	if expected != actual {

		t.Errorf("expected '%s' does not equal '%s'", expected, actual)
	}
}

// Specification

func Test_String_OF_UNINITIALISED_Specification(t *testing.T) {

	specification := clasp.Specification{}

	expected := "<clasp.Specification{ Type=<clasp.ArgType 0>, Name=\"\", Aliases=[], Help=\"\", ValueSet=[], BitFlags=0x0, BitFlags64=0x0, Extras=map[] }>"
	actual := specification.String()

	stegol.CheckStringEqual(t, expected, actual)
}

func Test_String_OF_Flag_Specification_1(t *testing.T) {

	specification := clasp.Specification{
		Type:     clasp.FlagType,
		Name:     "--flagpole",
		Help:     "help, plz",
		BitFlags: 0x1234,

		// following fields not used by `Option` so should not appear
		ValueSet: []string{"one", "two"},
	}

	expected := "<clasp.Specification{ Type=Flag, Name=\"--flagpole\", Aliases=[], Help=\"help, plz\", BitFlags=0x1234, flags_receiver=0x0, BitFlags64=0x0, flags64_receiver=0x0, Extras=map[] }>"
	actual := specification.String()

	stegol.CheckStringEqual(t, expected, actual)
}

func Test_String_OF_Flag_Specification_2(t *testing.T) {

	specification := clasp.Specification{
		Type:       clasp.FlagType,
		Name:       "--flagpole",
		Help:       "help, plz",
		BitFlags64: 0x5678,

		// following fields not used by `Option` so should not appear
		ValueSet: []string{"one", "two"},
	}

	expected := "<clasp.Specification{ Type=Flag, Name=\"--flagpole\", Aliases=[], Help=\"help, plz\", BitFlags=0x0, flags_receiver=0x0, BitFlags64=0x5678, flags64_receiver=0x0, Extras=map[] }>"
	actual := specification.String()

	stegol.CheckStringEqual(t, expected, actual)
}

func Test_String_OF_Option_Specification_1(t *testing.T) {

	specification := clasp.Specification{
		Type: clasp.OptionType,
		Name: "--count",
		Help: "number of things",
		ValueSet: []string{
			"one", "two",
		},

		// following fields not used by `Option` so should not appear
		BitFlags:   0x1234,
		BitFlags64: 0x5678,
	}

	expected := "<clasp.Specification{ Type=Option, Name=\"--count\", Aliases=[], Help=\"number of things\", ValueSet=[\"one\", \"two\"], Extras=map[] }>"
	actual := specification.String()

	stegol.CheckStringEqual(t, expected, actual)
}

func Test_String_OF_Section_Specification_1(t *testing.T) {

	specification := clasp.Specification{
		Type: clasp.SectionType,
		Name: "standard elements:",
	}

	expected := "<clasp.Specification{ Type=Section, Name=\"standard elements:\" }>"
	actual := specification.String()

	stegol.CheckStringEqual(t, expected, actual)
}

// Argument

func Test_String_of_Argument_1(t *testing.T) {

	argument := clasp.Argument{

		ResolvedName:    "--help",
		GivenName:       "--help",
		Value:           "",
		Type:            clasp.FlagType,
		CmdLineIndex:    1,
		NumGivenHyphens: 2,
		Flags:           0x1234,
	}

	expected := "<clasp.Argument{ ResolvedName=\"--help\", GivenName=\"--help\", Value=\"\", Type=Flag, CmdLineIndex=1, NumGivenHyphens=2, ArgumentSpecification=<nil>, Flags=0x1234, used=false }>"
	actual := argument.String()

	if expected != actual {

		t.Errorf("expected '%s' does not equal '%s'", expected, actual)
	}
}

// Arguments

func Test_String_of_Arguments_1(t *testing.T) {

	arguments := clasp.Arguments{}

	expected := "<clasp.Arguments{ Arguments=[], Flags=[], Options=[], Values=[], Argv=[], ProgramName=\"\" }>"
	actual := arguments.String()

	if expected != actual {

		t.Errorf("expected '%s' does not equal '%s'", expected, actual)
	}
}

// ParseParams

func Test_String_of_ParseParams_1(t *testing.T) {

	parseParams := clasp.ParseParams{}

	expected := "<clasp.ParseParams{ Specifications=[], Flags=0x0 }>"
	actual := parseParams.String()

	if expected != actual {

		t.Errorf("expected '%s' does not equal '%s'", expected, actual)
	}
}

// UsageParams

func Test_String_of_UsageParams_1(t *testing.T) {

	parseParams := clasp.UsageParams{}

	expected := "<clasp.UsageParams{ Stream=<nil>, ProgramName=\"\", UsageFlags=0x0, ExitCode=0, Exiter=<nil>, Version=<nil>, VersionPrefix=\"\", InfoLines=[], ValuesString=\"\", FlagsAndOptionsString=\"\" }>"
	actual := parseParams.String()

	if expected != actual {

		t.Errorf("expected '%s' does not equal '%s'", expected, actual)
	}
}

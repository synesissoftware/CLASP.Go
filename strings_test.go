
package clasp_test

import (

	clasp "github.com/synesissoftware/CLASP.Go"

	"testing"
)

// ArgType

func Test_String_of_ArgType_Flag(t *testing.T) {

	at_F		:=	clasp.Flag

	expected	:=	"Flag"
	actual		:=	at_F.String()

	if expected != actual {

		t.Errorf("expected '%s' does not equal '%s'", expected, actual)
	}
}

func Test_String_of_ArgType_Option(t *testing.T) {

	at_F		:=	clasp.Option

	expected	:=	"Option"
	actual		:=	at_F.String()

	if expected != actual {

		t.Errorf("expected '%s' does not equal '%s'", expected, actual)
	}
}

func Test_String_of_ArgType_Value(t *testing.T) {

	at_F		:=	clasp.Value

	expected	:=	"Value"
	actual		:=	at_F.String()

	if expected != actual {

		t.Errorf("expected '%s' does not equal '%s'", expected, actual)
	}
}

func Test_String_of_ArgType_unknown(t *testing.T) {

	var at_F clasp.ArgType = 101

	expected	:=	"<clasp.ArgType 101>"
	actual		:=	at_F.String()

	if expected != actual {

		t.Errorf("expected '%s' does not equal '%s'", expected, actual)
	}
}


// Alias

func Test_String_of_Alias_1(t *testing.T) {

	alias		:=	clasp.Alias{}

	expected	:=	"<clasp.Alias{ Type=<clasp.ArgType 0>, Name=\"\", Aliases=[], Help=\"\", ValueSet=[], BitFlags=0x0 }>"
	actual		:=	alias.String()

	if expected != actual {

		t.Errorf("expected '%s' does not equal '%s'", expected, actual)
	}
}

func Test_String_of_Alias_2(t *testing.T) {

	alias		:=	clasp.Alias{ Help: "help, plz", BitFlags: 0x1234, Name: "--flagpole", Type: clasp.Option }

	expected	:=	"<clasp.Alias{ Type=Option, Name=\"--flagpole\", Aliases=[], Help=\"help, plz\", ValueSet=[], BitFlags=0x1234 }>"
	actual		:=	alias.String()

	if expected != actual {

		t.Errorf("expected '%s' does not equal '%s'", expected, actual)
	}
}


// Argument

func Test_String_of_Argument_1(t *testing.T) {

	argument	:=	clasp.Argument{

		ResolvedName: "--help",
		GivenName: "--help",
		Value: "",
		Type: clasp.Flag,
		CmdLineIndex: 1,
		NumGivenHyphens: 2,
		Flags: 0x1234,
	}

	expected	:=	"<clasp.Argument{ ResolvedName=\"--help\", GivenName=\"--help\", Value=\"\", Type=Flag, CmdLineIndex=1, NumGivenHyphens=2, AliasIndex=0, Flags=0x1234, used=false }>"
	actual		:=	argument.String()

	if expected != actual {

		t.Errorf("expected '%s' does not equal '%s'", expected, actual)
	}
}


// Arguments

func Test_String_of_Arguments_1(t *testing.T) {

	arguments	:=	clasp.Arguments{}

	expected	:=	"<clasp.Arguments{ Arguments=[], Flags=[], Options=[], Values=[], Argv=[], ProgramName=\"\" }>"
	actual		:=	arguments.String()

	if expected != actual {

		t.Errorf("expected '%s' does not equal '%s'", expected, actual)
	}
}


// ParseParams

func Test_String_of_ParseParams_1(t *testing.T) {

	parseParams	:=	clasp.ParseParams{}

	expected	:=	"<clasp.ParseParams{ Aliases=[], Flags=0x0 }>"
	actual		:=	parseParams.String()

	if expected != actual {

		t.Errorf("expected '%s' does not equal '%s'", expected, actual)
	}
}


// UsageParams

func Test_String_of_UsageParams_1(t *testing.T) {

	parseParams	:=	clasp.UsageParams{}

	expected	:=	"<clasp.UsageParams{ Stream=<nil>, ProgramName=\"\", UsageFlags=0x0, ExitCode=0, Exiter=<nil>, Version=<nil>, VersionPrefix=\"\", InfoLines=[], ValuesString=\"\", FlagsAndOptionsString=\"\" }>"
	actual		:=	parseParams.String()

	if expected != actual {

		t.Errorf("expected '%s' does not equal '%s'", expected, actual)
	}
}


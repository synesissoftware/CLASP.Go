
package clasp

import "fmt"
import "path"
import "runtime"
import "testing"

func require(t *testing.T, cond bool, format string, args... interface{}) {
	if !cond {
		_, file, line, hasCallInfo := runtime.Caller(1)

		if hasCallInfo {
			fmt.Printf("\t%s:%d: %s\n", path.Base(file), line, fmt.Sprintf(format, args...))
		} else {
			fmt.Printf("\t%s\n", fmt.Sprintf(format, args...))
		}

		t.FailNow()
	}
}

func check(t *testing.T, cond bool, format string, args... interface{}) bool {
	if !cond {
		_, file, line, hasCallInfo := runtime.Caller(1)

		if hasCallInfo {
			fmt.Printf("\t%s:%d: %s\n", path.Base(file), line, fmt.Sprintf(format, args...))
		} else {
			fmt.Printf("\t%s\n", fmt.Sprintf(format, args...))
		}

		t.Fail()
	}
	return cond
}


func Test_no_args(t *testing.T) {

	case1_argv := []string { "path/blah" }

	args1 := Parse(case1_argv, ParseParams{})

	if expected, actual := 0, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Flags); check(t, expected == actual, "arguments object has wrong number of Flags: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 1, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {
		for i, expected := range case1_argv {
			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_single_value(t *testing.T) {

	case1_argv := []string { "path/blah", "abc" }

	args1 := Parse(case1_argv, ParseParams{})

	if expected, actual := 1, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {
		argument0 := args1.Arguments[0]
		check(t, "abc" == argument0.Value, "arguments has wrong argument")
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
	}
	if expected, actual := 0, len(args1.Flags); check(t, expected == actual, "arguments object has wrong number of Flags: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 1, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
		value0 := args1.Values[0]
		check(t, 1 == value0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Value == value0.Type, "argument has wrong type: expected: %v; received %v", Value, value0.Type)
		check(t, "abc" == value0.Value, "arguments has wrong value")
	}
	if expected, actual := 2, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {
		for i, expected := range case1_argv {
			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_several_values(t *testing.T) {

	case1_argv := []string { "path/blah", "abc", "def", "g", "hi" }

	args1 := Parse(case1_argv, ParseParams{})

	if expected, actual := 4, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {
		argument0 := args1.Arguments[0]
		argument1 := args1.Arguments[1]
		argument2 := args1.Arguments[2]
		argument3 := args1.Arguments[3]
		check(t, "abc" == argument0.Value, "arguments has wrong argument")
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
		check(t, "def" == argument1.Value, "arguments has wrong argument")
		check(t, 2 == argument1.CmdLineIndex, "arguments has wrong argument")
		check(t, "g" == argument2.Value, "arguments has wrong argument")
		check(t, 3 == argument2.CmdLineIndex, "arguments has wrong argument")
		check(t, "hi" == argument3.Value, "arguments has wrong argument")
		check(t, 4 == argument3.CmdLineIndex, "arguments has wrong argument")
	}
	if expected, actual := 0, len(args1.Flags); check(t, expected == actual, "arguments object has wrong number of Flags: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 4, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
		value0 := args1.Values[0]
		value1 := args1.Values[1]
		value2 := args1.Values[2]
		value3 := args1.Values[3]
		check(t, 1 == value0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Value == value0.Type, "argument has wrong type: expected: %v; received %v", Value, value0.Type)
		check(t, "abc" == value0.Value, "arguments has wrong value")
		check(t, 2 == value1.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Value == value1.Type, "argument has wrong type: expected: %v; received %v", Value, value1.Type)
		check(t, "def" == value1.Value, "arguments has wrong value")
		check(t, 3 == value2.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Value == value2.Type, "argument has wrong type: expected: %v; received %v", Value, value2.Type)
		check(t, "g" == value2.Value, "arguments has wrong value")
		check(t, 4 == value3.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Value == value3.Type, "argument has wrong type: expected: %v; received %v", Value, value3.Type)
		check(t, "hi" == value3.Value, "arguments has wrong value")
	}
	if expected, actual := 5, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {
		for i, expected := range case1_argv {
			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_several_values_with_double_hyphen(t *testing.T) {

	case1_argv := []string { "path/blah", "abc", "def", "--", "-g", "--hi" }

	args1 := Parse(case1_argv, ParseParams{})

	if expected, actual := 4, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {
		argument0 := args1.Arguments[0]
		argument1 := args1.Arguments[1]
		argument2 := args1.Arguments[2]
		argument3 := args1.Arguments[3]
		check(t, "abc" == argument0.Value, "arguments has wrong argument")
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
		check(t, "def" == argument1.Value, "arguments has wrong argument")
		check(t, 2 == argument1.CmdLineIndex, "arguments has wrong argument")
		check(t, "-g" == argument2.Value, "arguments has wrong argument")
		check(t, 4 == argument2.CmdLineIndex, "arguments has wrong argument")
		check(t, "--hi" == argument3.Value, "arguments has wrong argument")
		check(t, 5 == argument3.CmdLineIndex, "arguments has wrong argument")
	}
	if expected, actual := 0, len(args1.Flags); check(t, expected == actual, "arguments object has wrong number of Flags: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 4, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
		value0 := args1.Values[0]
		value1 := args1.Values[1]
		value2 := args1.Values[2]
		value3 := args1.Values[3]
		check(t, 1 == value0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Value == value0.Type, "argument has wrong type: expected: %v; received %v", Value, value0.Type)
		check(t, "abc" == value0.Value, "arguments has wrong value")
		check(t, 0 == value0.NumGivenHyphens, "arguments has wrong number of hyphens")
		check(t, 2 == value1.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Value == value1.Type, "argument has wrong type: expected: %v; received %v", Value, value1.Type)
		check(t, "def" == value1.Value, "arguments has wrong value")
		check(t, 0 == value1.NumGivenHyphens, "arguments has wrong number of hyphens")
		check(t, 4 == value2.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Value == value2.Type, "argument has wrong type: expected: %v; received %v", Value, value2.Type)
		check(t, "-g" == value2.Value, "arguments has wrong value")
		check(t, 0 == value2.NumGivenHyphens, "arguments has wrong number of hyphens")
		check(t, 5 == value3.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Value == value3.Type, "argument has wrong type: expected: %v; received %v", Value, value3.Type)
		check(t, "--hi" == value3.Value, "arguments has wrong value")
		check(t, 0 == value3.NumGivenHyphens, "arguments has wrong number of hyphens")
	}
	if expected, actual := 6, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {
		for i, expected := range case1_argv {
			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_single_flag(t *testing.T) {

	case1_argv := []string { "path/blah", "-f" }

	args1 := Parse(case1_argv, ParseParams{})

	if expected, actual := 1, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {
		argument0 := args1.Arguments[0]
		check(t, "" == argument0.Value, "arguments has wrong argument")
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
	}
	if expected := 1; check(t, expected == len(args1.Flags), "arguments object has wrong number of flags: expected %v; actual %v", expected, len(args1.Flags)) {
		flag0 := args1.Flags[0]
		check(t, 1 == flag0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Flag == flag0.Type, "argument has wrong type: expected: %v; received %v", Flag, flag0.Type)
		check(t, "-f" == flag0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-f" == flag0.GivenName, "arguments has wrong given name")
		check(t, "" == flag0.Value, "arguments has wrong value")
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 2, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {
		for i, expected := range case1_argv {
			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_several_flags(t *testing.T) {

	case1_argv := []string { "path/blah", "-f", "--flag2", "---flag3" }

	args1 := Parse(case1_argv, ParseParams{})

	if expected, actual := 3, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {
		argument0 := args1.Arguments[0]
		check(t, "" == argument0.Value, "arguments has wrong argument")
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
	}
	if expected := 3; check(t, expected == len(args1.Flags), "arguments object has wrong number of flags: expected %v; actual %v", expected, len(args1.Flags)) {
		flag0 := args1.Flags[0]
		flag1 := args1.Flags[1]
		flag2 := args1.Flags[2]
		check(t, 1 == flag0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Flag == flag0.Type, "argument has wrong type: expected: %v; received %v", Flag, flag0.Type)
		check(t, "-f" == flag0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-f" == flag0.GivenName, "arguments has wrong given name")
		check(t, "" == flag0.Value, "arguments has wrong value")
		check(t, 2 == flag1.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Flag == flag1.Type, "argument has wrong type: expected: %v; received %v", Flag, flag1.Type)
		check(t, "--flag2" == flag1.ResolvedName, "arguments has wrong resolved name")
		check(t, "--flag2" == flag1.GivenName, "arguments has wrong given name")
		check(t, "" == flag1.Value, "arguments has wrong value")
		check(t, 3 == flag2.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Flag == flag2.Type, "argument has wrong type: expected: %v; received %v", Flag, flag2.Type)
		check(t, "---flag3" == flag2.ResolvedName, "arguments has wrong resolved name")
		check(t, "---flag3" == flag2.GivenName, "arguments has wrong given name")
		check(t, "" == flag2.Value, "arguments has wrong value")
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 4, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {
		for i, expected := range case1_argv {
			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_flags_and_values(t *testing.T) {

	case1_argv := []string { "path/blah", "abc", "-f", "def", "--flag2", "--", "--value-really" }

	args1 := Parse(case1_argv, ParseParams{})

	if expected, actual := 5, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {
		argument0 := args1.Arguments[0]
		argument1 := args1.Arguments[1]
		argument2 := args1.Arguments[2]
		argument3 := args1.Arguments[3]
		argument4 := args1.Arguments[4]
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
		check(t, 2 == argument1.CmdLineIndex, "arguments has wrong argument")
		check(t, 3 == argument2.CmdLineIndex, "arguments has wrong argument")
		check(t, 4 == argument3.CmdLineIndex, "arguments has wrong argument")
		check(t, 6 == argument4.CmdLineIndex, "arguments has wrong argument")
		check(t, "abc" == argument0.Value, "arguments has wrong argument")
		check(t, "def" == argument2.Value, "arguments has wrong argument")
		check(t, "--value-really" == argument4.Value, "arguments has wrong argument")
	}
	if expected := 2; check(t, expected == len(args1.Flags), "arguments object has wrong number of flags: expected %v; actual %v", expected, len(args1.Flags)) {
		flag0 := args1.Flags[0]
		flag1 := args1.Flags[1]
		check(t, 2 == flag0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Flag == flag0.Type, "argument has wrong type: expected: %v; received %v", Flag, flag0.Type)
		check(t, "-f" == flag0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-f" == flag0.GivenName, "arguments has wrong given name")
		check(t, "" == flag0.Value, "arguments has wrong value")
		check(t, 4 == flag1.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Flag == flag1.Type, "argument has wrong type: expected: %v; received %v", Flag, flag1.Type)
		check(t, "--flag2" == flag1.ResolvedName, "arguments has wrong resolved name")
		check(t, "--flag2" == flag1.GivenName, "arguments has wrong given name")
		check(t, "" == flag1.Value, "arguments has wrong value")
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 3, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
		value0 := args1.Values[0]
		value1 := args1.Values[1]
		value2 := args1.Values[2]
		check(t, 1 == value0.CmdLineIndex, "values has wrong value")
		check(t, 3 == value1.CmdLineIndex, "values has wrong value")
		check(t, 6 == value2.CmdLineIndex, "values has wrong value")
		check(t, "abc" == value0.Value, "values has wrong value")
		check(t, "def" == value1.Value, "values has wrong value")
		check(t, "--value-really" == value2.Value, "values has wrong value")
	}
	if expected, actual := 7, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {
		for i, expected := range case1_argv {
			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_flags_values_and_option(t *testing.T) {

	case1_argv := []string { "path/blah", "abc", "-f", "def", "--opt1=value1", "--", "--value-really" }

	args1 := Parse(case1_argv, ParseParams{})

	if expected, actual := 5, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {
		argument0 := args1.Arguments[0]
		argument1 := args1.Arguments[1]
		argument2 := args1.Arguments[2]
		argument3 := args1.Arguments[3]
		argument4 := args1.Arguments[4]
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
		check(t, 2 == argument1.CmdLineIndex, "arguments has wrong argument")
		check(t, 3 == argument2.CmdLineIndex, "arguments has wrong argument")
		check(t, 4 == argument3.CmdLineIndex, "arguments has wrong argument")
		check(t, 6 == argument4.CmdLineIndex, "arguments has wrong argument")
		check(t, "abc" == argument0.Value, "arguments has wrong value")
		check(t, "def" == argument2.Value, "arguments has wrong value")
		check(t, "value1" == argument3.Value, "arguments has wrong value")
		check(t, "--value-really" == argument4.Value, "arguments has wrong value")
	}
	if expected := 1; check(t, expected == len(args1.Flags), "arguments object has wrong number of flags: expected %v; actual %v", expected, len(args1.Flags)) {
		flag0 := args1.Flags[0]
		check(t, 2 == flag0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Flag == flag0.Type, "argument has wrong type: expected: %v; received %v", Flag, flag0.Type)
		check(t, "-f" == flag0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-f" == flag0.GivenName, "arguments has wrong given name")
		check(t, "" == flag0.Value, "arguments has wrong value")
	}
	if expected, actual := 1, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
		option0 := args1.Options[0]
		check(t, 4 == option0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Option == option0.Type, "argument has wrong type: expected: %v; received %v", Option, option0.Type)
		check(t, "--opt1" == option0.ResolvedName, "arguments has wrong resolved name")
		check(t, "--opt1" == option0.GivenName, "arguments has wrong given name")
		check(t, "value1" == option0.Value, "arguments has wrong value")
	}
	if expected, actual := 3, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
		value0 := args1.Values[0]
		value1 := args1.Values[1]
		value2 := args1.Values[2]
		check(t, 1 == value0.CmdLineIndex, "values has wrong value")
		check(t, 3 == value1.CmdLineIndex, "values has wrong value")
		check(t, 6 == value2.CmdLineIndex, "values has wrong value")
		check(t, "abc" == value0.Value, "values has wrong value")
		check(t, "def" == value1.Value, "values has wrong value")
		check(t, "--value-really" == value2.Value, "values has wrong value")
	}
	if expected, actual := 7, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {
		for i, expected := range case1_argv {
			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_single_hyphen_default(t *testing.T) {

	case1_argv := []string { "path/blah", "-" }

	args1 := Parse(case1_argv, ParseParams{})

	if expected, actual := 1, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {
		argument0 := args1.Arguments[0]
		check(t, "" == argument0.Value, "arguments has wrong argument")
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
	}
	if expected := 1; check(t, expected == len(args1.Flags), "arguments object has wrong number of flags: expected %v; actual %v", expected, len(args1.Flags)) {
		flag0 := args1.Flags[0]
		check(t, 1 == flag0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Flag == flag0.Type, "argument has wrong type: expected: %v; received %v", Flag, flag0.Type)
		check(t, "-" == flag0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-" == flag0.GivenName, "arguments has wrong given name")
		check(t, "" == flag0.Value, "arguments has wrong value")
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 2, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {
		for i, expected := range case1_argv {
			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_single_hyphen_as_value(t *testing.T) {

	case1_argv := []string { "path/blah", "-" }

	args1 := Parse(case1_argv, ParseParams{ Flags : ParseTreatSingleHyphenAsValue })

	if expected, actual := 1, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {
		argument0 := args1.Arguments[0]
		check(t, "-" == argument0.Value, "arguments has wrong argument")
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
	}
	if expected := 0; check(t, expected == len(args1.Flags), "arguments object has wrong number of flags: expected %v; actual %v", expected, len(args1.Flags)) {
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 1, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
		value0 := args1.Values[0]
		check(t, 1 == value0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Value == value0.Type, "argument has wrong type: expected: %v; received %v", Value, value0.Type)
		check(t, "-" == value0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-" == value0.GivenName, "arguments has wrong given name")
		check(t, "-" == value0.Value, "arguments has wrong value")
	}
	if expected, actual := 2, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {
		for i, expected := range case1_argv {
			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_alias_0(t *testing.T) {

	aliases := []Alias {
	}
	case1_argv := []string { "path/blah", "-f", "-f2", "---flag3" }

	args1 := Parse(case1_argv, ParseParams{ Aliases: aliases })

	if expected, actual := 3, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {
		argument0 := args1.Arguments[0]
		check(t, "" == argument0.Value, "arguments has wrong argument")
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
	}
	if expected := 3; check(t, expected == len(args1.Flags), "arguments object has wrong number of flags: expected %v; actual %v", expected, len(args1.Flags)) {
		flag0 := args1.Flags[0]
		flag1 := args1.Flags[1]
		flag2 := args1.Flags[2]
		check(t, 1 == flag0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Flag == flag0.Type, "argument has wrong type: expected: %v; received %v", Flag, flag0.Type)
		check(t, "-f" == flag0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-f" == flag0.GivenName, "arguments has wrong given name")
		check(t, "" == flag0.Value, "arguments has wrong value")
		check(t, 2 == flag1.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Flag == flag1.Type, "argument has wrong type: expected: %v; received %v", Flag, flag1.Type)
		check(t, "-f2" == flag1.ResolvedName, "arguments has wrong resolved name")
		check(t, "-f2" == flag1.GivenName, "arguments has wrong given name")
		check(t, "" == flag1.Value, "arguments has wrong value")
		check(t, 3 == flag2.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Flag == flag2.Type, "argument has wrong type: expected: %v; received %v", Flag, flag2.Type)
		check(t, "---flag3" == flag2.ResolvedName, "arguments has wrong resolved name")
		check(t, "---flag3" == flag2.GivenName, "arguments has wrong given name")
		check(t, "" == flag2.Value, "arguments has wrong value")
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 4, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {
		for i, expected := range case1_argv {
			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_alias_1(t *testing.T) {

	aliases := []Alias {
		{ Flag, "--flag2", []string { "-f2" }, "second flag", nil, 0 },
//		{ clasp.Flag, "--flag2", []{ "-f2" }, "second flag", nil, 0 }
	}
	case1_argv := []string { "path/blah", "-f", "-f2", "---flag3" }

	args1 := Parse(case1_argv, ParseParams{ Aliases: aliases })

	if expected, actual := 3, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {
		argument0 := args1.Arguments[0]
		check(t, "" == argument0.Value, "arguments has wrong argument")
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
	}
	if expected := 3; check(t, expected == len(args1.Flags), "arguments object has wrong number of flags: expected %v; actual %v", expected, len(args1.Flags)) {
		flag0 := args1.Flags[0]
		flag1 := args1.Flags[1]
		flag2 := args1.Flags[2]
		check(t, 1 == flag0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Flag == flag0.Type, "argument has wrong type: expected: %v; received %v", Flag, flag0.Type)
		check(t, "-f" == flag0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-f" == flag0.GivenName, "arguments has wrong given name")
		check(t, "" == flag0.Value, "arguments has wrong value")
		check(t, 2 == flag1.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Flag == flag1.Type, "argument has wrong type: expected: %v; received %v", Flag, flag1.Type)
		check(t, "--flag2" == flag1.ResolvedName, "arguments has wrong resolved name")
		check(t, "-f2" == flag1.GivenName, "arguments has wrong given name")
		check(t, "" == flag1.Value, "arguments has wrong value")
		check(t, 3 == flag2.CmdLineIndex, "arguments has wrong command-line index")
		check(t, Flag == flag2.Type, "argument has wrong type: expected: %v; received %v", Flag, flag2.Type)
		check(t, "---flag3" == flag2.ResolvedName, "arguments has wrong resolved name")
		check(t, "---flag3" == flag2.GivenName, "arguments has wrong given name")
		check(t, "" == flag2.Value, "arguments has wrong value")
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 4, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {
		for i, expected := range case1_argv {
			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}



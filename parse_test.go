
package clasp_test

import clasp "github.com/synesissoftware/CLASP.Go"

import "fmt"
import "path"
import "runtime"
import "testing"

func equalInNonNillLhs(lhs clasp.Argument, rhs clasp.Argument) bool {

	if lhs.Type != rhs.Type {

		return false
	}

	if "" != lhs.ResolvedName && lhs.ResolvedName != rhs.ResolvedName {

		return false
	}

	if "" != lhs.GivenName && lhs.GivenName != rhs.GivenName {

		return false
	}

	if "" != lhs.Value && lhs.Value != rhs.Value {

		return false
	}

	if 0 != lhs.CmdLineIndex && lhs.CmdLineIndex != rhs.CmdLineIndex {

		return false
	}

	if 0 != lhs.NumGivenHyphens && lhs.NumGivenHyphens != rhs.NumGivenHyphens {

		return false
	}

	if 0 != lhs.Flags && lhs.Flags != rhs.Flags {

		return false
	}

	return true
}

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

	argv := []string { "path/blah" }

	args1 := clasp.Parse(argv, clasp.ParseParams{})

	if expected, actual := 0, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Flags); check(t, expected == actual, "arguments object has wrong number of Flags: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 1, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {

		for i, expected := range argv {

			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_single_value(t *testing.T) {

	argv := []string { "path/blah", "abc" }

	args1 := clasp.Parse(argv, clasp.ParseParams{})

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
		check(t, clasp.Value == value0.Type, "argument has wrong type: expected: %v; received %v", clasp.Value, value0.Type)
		check(t, "abc" == value0.Value, "arguments has wrong value")
	}
	if expected, actual := 2, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {

		for i, expected := range argv {

			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_several_values(t *testing.T) {

	argv := []string { "path/blah", "abc", "def", "g", "hi" }

	args1 := clasp.Parse(argv, clasp.ParseParams{})

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
		check(t, clasp.Value == value0.Type, "argument has wrong type: expected: %v; received %v", clasp.Value, value0.Type)
		check(t, "abc" == value0.Value, "arguments has wrong value")
		check(t, 2 == value1.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Value == value1.Type, "argument has wrong type: expected: %v; received %v", clasp.Value, value1.Type)
		check(t, "def" == value1.Value, "arguments has wrong value")
		check(t, 3 == value2.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Value == value2.Type, "argument has wrong type: expected: %v; received %v", clasp.Value, value2.Type)
		check(t, "g" == value2.Value, "arguments has wrong value")
		check(t, 4 == value3.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Value == value3.Type, "argument has wrong type: expected: %v; received %v", clasp.Value, value3.Type)
		check(t, "hi" == value3.Value, "arguments has wrong value")
	}
	if expected, actual := 5, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {

		for i, expected := range argv {

			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_several_values_with_double_hyphen(t *testing.T) {

	argv := []string { "path/blah", "abc", "def", "--", "-g", "--hi" }

	args1 := clasp.Parse(argv, clasp.ParseParams{})

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
		check(t, clasp.Value == value0.Type, "argument has wrong type: expected: %v; received %v", clasp.Value, value0.Type)
		check(t, "abc" == value0.Value, "arguments has wrong value")
		check(t, 0 == value0.NumGivenHyphens, "arguments has wrong number of hyphens")
		check(t, 2 == value1.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Value == value1.Type, "argument has wrong type: expected: %v; received %v", clasp.Value, value1.Type)
		check(t, "def" == value1.Value, "arguments has wrong value")
		check(t, 0 == value1.NumGivenHyphens, "arguments has wrong number of hyphens")
		check(t, 4 == value2.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Value == value2.Type, "argument has wrong type: expected: %v; received %v", clasp.Value, value2.Type)
		check(t, "-g" == value2.Value, "arguments has wrong value")
		check(t, 0 == value2.NumGivenHyphens, "arguments has wrong number of hyphens")
		check(t, 5 == value3.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Value == value3.Type, "argument has wrong type: expected: %v; received %v", clasp.Value, value3.Type)
		check(t, "--hi" == value3.Value, "arguments has wrong value")
		check(t, 0 == value3.NumGivenHyphens, "arguments has wrong number of hyphens")
	}
	if expected, actual := 6, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {

		for i, expected := range argv {

			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_several_values_with_double_hyphen_suppressed(t *testing.T) {

	argv := []string { "path/blah", "abc", "def", "--", "-g", "--hi" }

	args1 := clasp.Parse(argv, clasp.ParseParams{ Flags : clasp.ParseDontRecogniseDoubleHyphenToStartValues })

	if expected, actual := 5, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {

		argument0 := args1.Arguments[0]
		argument1 := args1.Arguments[1]
		argument2 := args1.Arguments[2]
		argument3 := args1.Arguments[3]
		argument4 := args1.Arguments[4]
		check(t, "abc" == argument0.Value, "arguments has wrong value: actual '%v'", argument0.Value)
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, "def" == argument1.Value, "arguments has wrong value: actual '%v'", argument1.Value)
		check(t, 2 == argument1.CmdLineIndex, "arguments has wrong command-line index")
		check(t, "--" == argument2.GivenName, "arguments has wrong given name: actual '%v'", argument2.GivenName)
		check(t, 3 == argument2.CmdLineIndex, "arguments has wrong command-line index")
		check(t, "-g" == argument3.GivenName, "arguments has wrong given name: actual '%v'", argument3.GivenName)
		check(t, 4 == argument3.CmdLineIndex, "arguments has wrong command-line index")
		check(t, "--hi" == argument4.GivenName, "arguments has wrong given name: actual '%v'", argument4.GivenName)
		check(t, 5 == argument4.CmdLineIndex, "arguments has wrong command-line index")
	}
	if expected, actual := 3, len(args1.Flags); check(t, expected == actual, "arguments object has wrong number of Flags: expected %v; actual %v", expected, actual) {

		flag0 := args1.Flags[0]
		flag1 := args1.Flags[1]
		flag2 := args1.Flags[2]
		check(t, 3 == flag0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag0.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag0.Type)
		check(t, "--" == flag0.GivenName, "arguments has wrong given name")
		check(t, 2 == flag0.NumGivenHyphens, "arguments has wrong number of hyphens: received %v", flag0.NumGivenHyphens)
		check(t, 4 == flag1.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag1.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag1.Type)
		check(t, "-g" == flag1.GivenName, "arguments has wrong given name")
		check(t, 1 == flag1.NumGivenHyphens, "arguments has wrong number of hyphens")
		check(t, 5 == flag2.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag2.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag2.Type)
		check(t, "--hi" == flag2.GivenName, "arguments has wrong given name")
		check(t, 2 == flag2.NumGivenHyphens, "arguments has wrong number of hyphens")
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 2, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {

		value0 := args1.Values[0]
		value1 := args1.Values[1]
		check(t, 1 == value0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Value == value0.Type, "argument has wrong type: expected: %v; received %v", clasp.Value, value0.Type)
		check(t, "abc" == value0.Value, "arguments has wrong value")
		check(t, 0 == value0.NumGivenHyphens, "arguments has wrong number of hyphens")
		check(t, 2 == value1.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Value == value1.Type, "argument has wrong type: expected: %v; received %v", clasp.Value, value1.Type)
		check(t, "def" == value1.Value, "arguments has wrong value")
		check(t, 0 == value1.NumGivenHyphens, "arguments has wrong number of hyphens")
	}
	if expected, actual := 6, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {

		for i, expected := range argv {

			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_single_flag(t *testing.T) {

	argv := []string { "path/blah", "-f" }

	args1 := clasp.Parse(argv, clasp.ParseParams{})

	if expected, actual := 1, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {

		argument0 := args1.Arguments[0]
		check(t, "" == argument0.Value, "arguments has wrong argument")
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
	}
	if expected := 1; check(t, expected == len(args1.Flags), "arguments object has wrong number of Flags: expected %v; actual %v", expected, len(args1.Flags)) {

		flag0 := args1.Flags[0]
		check(t, 1 == flag0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag0.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag0.Type)
		check(t, "-f" == flag0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-f" == flag0.GivenName, "arguments has wrong given name")
		check(t, "" == flag0.Value, "arguments has wrong value")
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 2, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {

		for i, expected := range argv {

			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_several_flags(t *testing.T) {

	argv := []string { "path/blah", "-f", "--flag2", "---flag3" }

	args1 := clasp.Parse(argv, clasp.ParseParams{})

	if expected, actual := 3, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {

		argument0 := args1.Arguments[0]
		check(t, "" == argument0.Value, "arguments has wrong argument")
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
	}
	if expected := 3; check(t, expected == len(args1.Flags), "arguments object has wrong number of Flags: expected %v; actual %v", expected, len(args1.Flags)) {

		flag0 := args1.Flags[0]
		flag1 := args1.Flags[1]
		flag2 := args1.Flags[2]
		check(t, 1 == flag0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag0.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag0.Type)
		check(t, "-f" == flag0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-f" == flag0.GivenName, "arguments has wrong given name")
		check(t, "" == flag0.Value, "arguments has wrong value")
		check(t, 2 == flag1.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag1.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag1.Type)
		check(t, "--flag2" == flag1.ResolvedName, "arguments has wrong resolved name")
		check(t, "--flag2" == flag1.GivenName, "arguments has wrong given name")
		check(t, "" == flag1.Value, "arguments has wrong value")
		check(t, 3 == flag2.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag2.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag2.Type)
		check(t, "---flag3" == flag2.ResolvedName, "arguments has wrong resolved name")
		check(t, "---flag3" == flag2.GivenName, "arguments has wrong given name")
		check(t, "" == flag2.Value, "arguments has wrong value")
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 4, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {

		for i, expected := range argv {

			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_flags_and_values(t *testing.T) {

	argv := []string { "path/blah", "abc", "-f", "def", "--flag2", "--", "--value-really" }

	args1 := clasp.Parse(argv, clasp.ParseParams{})

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
	if expected := 2; check(t, expected == len(args1.Flags), "arguments object has wrong number of Flags: expected %v; actual %v", expected, len(args1.Flags)) {

		flag0 := args1.Flags[0]
		flag1 := args1.Flags[1]
		check(t, 2 == flag0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag0.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag0.Type)
		check(t, "-f" == flag0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-f" == flag0.GivenName, "arguments has wrong given name")
		check(t, "" == flag0.Value, "arguments has wrong value")
		check(t, 4 == flag1.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag1.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag1.Type)
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

		for i, expected := range argv {

			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_flags_values_and_option(t *testing.T) {

	argv := []string { "path/blah", "abc", "-f", "def", "--opt1=value1", "--", "--value-really" }

	args1 := clasp.Parse(argv, clasp.ParseParams{})

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
	if expected := 1; check(t, expected == len(args1.Flags), "arguments object has wrong number of Flags: expected %v; actual %v", expected, len(args1.Flags)) {

		flag0 := args1.Flags[0]
		check(t, 2 == flag0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag0.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag0.Type)
		check(t, "-f" == flag0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-f" == flag0.GivenName, "arguments has wrong given name")
		check(t, "" == flag0.Value, "arguments has wrong value")
	}
	if expected, actual := 1, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {

		option0 := args1.Options[0]
		check(t, 4 == option0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Option == option0.Type, "argument has wrong type: expected: %v; received %v", clasp.Option, option0.Type)
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

		for i, expected := range argv {

			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_single_hyphen_default(t *testing.T) {

	argv := []string { "path/blah", "-" }

	args1 := clasp.Parse(argv, clasp.ParseParams{})

	if expected, actual := 1, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {

		argument0 := args1.Arguments[0]
		check(t, "" == argument0.Value, "arguments has wrong argument")
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
	}
	if expected := 1; check(t, expected == len(args1.Flags), "arguments object has wrong number of Flags: expected %v; actual %v", expected, len(args1.Flags)) {

		flag0 := args1.Flags[0]
		check(t, 1 == flag0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag0.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag0.Type)
		check(t, "-" == flag0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-" == flag0.GivenName, "arguments has wrong given name")
		check(t, "" == flag0.Value, "arguments has wrong value")
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 2, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {

		for i, expected := range argv {

			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_single_hyphen_as_value(t *testing.T) {

	argv := []string { "path/blah", "-" }

	args1 := clasp.Parse(argv, clasp.ParseParams{ Flags : clasp.ParseTreatSingleHyphenAsValue })

	if expected, actual := 1, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {

		argument0 := args1.Arguments[0]
		check(t, "-" == argument0.Value, "arguments has wrong argument")
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
	}
	if expected := 0; check(t, expected == len(args1.Flags), "arguments object has wrong number of Flags: expected %v; actual %v", expected, len(args1.Flags)) {
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 1, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {

		value0 := args1.Values[0]
		check(t, 1 == value0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Value == value0.Type, "argument has wrong type: expected: %v; received %v", clasp.Value, value0.Type)
		check(t, "-" == value0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-" == value0.GivenName, "arguments has wrong given name")
		check(t, "-" == value0.Value, "arguments has wrong value")
	}
	if expected, actual := 2, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {

		for i, expected := range argv {

			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_alias_0(t *testing.T) {

	aliases := []clasp.Alias {
	}
	argv := []string { "path/blah", "-f", "-f2", "---flag3" }

	args1 := clasp.Parse(argv, clasp.ParseParams{ Aliases: aliases })

	if expected, actual := 3, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {

		argument0 := args1.Arguments[0]
		check(t, "" == argument0.Value, "arguments has wrong argument")
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
	}
	if expected := 3; check(t, expected == len(args1.Flags), "arguments object has wrong number of Flags: expected %v; actual %v", expected, len(args1.Flags)) {

		flag0 := args1.Flags[0]
		flag1 := args1.Flags[1]
		flag2 := args1.Flags[2]
		check(t, 1 == flag0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag0.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag0.Type)
		check(t, "-f" == flag0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-f" == flag0.GivenName, "arguments has wrong given name")
		check(t, "" == flag0.Value, "arguments has wrong value")
		check(t, 2 == flag1.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag1.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag1.Type)
		check(t, "-f2" == flag1.ResolvedName, "arguments has wrong resolved name")
		check(t, "-f2" == flag1.GivenName, "arguments has wrong given name")
		check(t, "" == flag1.Value, "arguments has wrong value")
		check(t, 3 == flag2.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag2.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag2.Type)
		check(t, "---flag3" == flag2.ResolvedName, "arguments has wrong resolved name")
		check(t, "---flag3" == flag2.GivenName, "arguments has wrong given name")
		check(t, "" == flag2.Value, "arguments has wrong value")
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 4, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {

		for i, expected := range argv {

			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_alias_1(t *testing.T) {

	aliases := []clasp.Alias {

		{ clasp.Flag, "--flag2", []string { "-f2" }, "second flag", nil, 0 },
	}
	argv := []string { "path/blah", "-f", "-f2", "---flag3" }

	args1 := clasp.Parse(argv, clasp.ParseParams{ Aliases: aliases })

	if expected, actual := 3, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {

		argument0 := args1.Arguments[0]
		check(t, "" == argument0.Value, "arguments has wrong argument")
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
	}
	if expected := 3; check(t, expected == len(args1.Flags), "arguments object has wrong number of Flags: expected %v; actual %v", expected, len(args1.Flags)) {

		flag0 := args1.Flags[0]
		flag1 := args1.Flags[1]
		flag2 := args1.Flags[2]
		check(t, 1 == flag0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag0.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag0.Type)
		check(t, "-f" == flag0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-f" == flag0.GivenName, "arguments has wrong given name")
		check(t, "" == flag0.Value, "arguments has wrong value")
		check(t, 2 == flag1.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag1.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag1.Type)
		check(t, "--flag2" == flag1.ResolvedName, "arguments has wrong resolved name")
		check(t, "-f2" == flag1.GivenName, "arguments has wrong given name")
		check(t, "" == flag1.Value, "arguments has wrong value")
		check(t, 3 == flag2.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag2.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag2.Type)
		check(t, "---flag3" == flag2.ResolvedName, "arguments has wrong resolved name")
		check(t, "---flag3" == flag2.GivenName, "arguments has wrong given name")
		check(t, "" == flag2.Value, "arguments has wrong value")
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 4, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {

		for i, expected := range argv {

			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_alias_2(t *testing.T) {

	aliases := []clasp.Alias {

		{ clasp.Option, "--flag2", []string { "-f2" }, "f2-option", nil, 0 },
	}
	argv := []string { "path/blah", "-f", "-f2", "abc", "---flag3" }

	args1 := clasp.Parse(argv, clasp.ParseParams{ Aliases: aliases })

	if expected, actual := 3, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {

		argument0 := args1.Arguments[0]
		argument1 := args1.Arguments[1]
		argument2 := args1.Arguments[2]
		check(t, "" == argument0.Value, "arguments has wrong argument")
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
		check(t, "abc" == argument1.Value, "arguments has wrong argument")
		check(t, 2 == argument1.CmdLineIndex, "arguments has wrong argument")
		check(t, "" == argument2.Value, "arguments has wrong argument")
		check(t, 4 == argument2.CmdLineIndex, "arguments has wrong argument")
	}
	if expected := 2; check(t, expected == len(args1.Flags), "arguments object has wrong number of Flags: expected %v; actual %v", expected, len(args1.Flags)) {

		flag0 := args1.Flags[0]
		flag1 := args1.Flags[1]
		check(t, 1 == flag0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag0.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag0.Type)
		check(t, "-f" == flag0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-f" == flag0.GivenName, "arguments has wrong given name")
		check(t, "" == flag0.Value, "arguments has wrong value")
		check(t, 4 == flag1.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag1.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag1.Type)
		check(t, "---flag3" == flag1.ResolvedName, "arguments has wrong resolved name")
		check(t, "---flag3" == flag1.GivenName, "arguments has wrong given name")
		check(t, "" == flag1.Value, "arguments has wrong value")
	}
	if expected, actual := 1, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {

		option0 := args1.Options[0]
		check(t, 2 == option0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Option == option0.Type, "argument has wrong type: expected: %v; received %v", clasp.Option, option0.Type)
		check(t, "--flag2" == option0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-f2" == option0.GivenName, "arguments has wrong given name")
		check(t, "abc" == option0.Value, "arguments has wrong value")
	}
	if expected, actual := 0, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 5, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {

		for i, expected := range argv {

			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_find_unused_flags_1(t *testing.T) {

	argv := []string { "path/blah", "-f1", "-f2", "abc", "---flag3" }

	args1 := clasp.Parse(argv, clasp.ParseParams{})

	unused := args1.GetUnusedFlags()
	if expected, actual := 3, len(unused); check(t, expected == actual, "unused slice has wrong size: expected %v; actual %v", expected, actual) {

		if expected, actual := (clasp.Argument{ Type : clasp.Flag, ResolvedName : "-f1", CmdLineIndex : 1 }), *unused[0]; check(t, equalInNonNillLhs(expected, actual), "argument not equal to expected: expected '%v'; actual '%v'", expected, actual) {
		}
		if expected, actual := (clasp.Argument{ Type : clasp.Flag, ResolvedName : "-f2", CmdLineIndex : 2 }), *unused[1]; check(t, equalInNonNillLhs(expected, actual), "argument not equal to expected: expected '%v'; actual '%v'", expected, actual) {
		}
		if expected, actual := (clasp.Argument{ Type : clasp.Flag, ResolvedName : "---flag3", CmdLineIndex : 4 }), *unused[2]; check(t, equalInNonNillLhs(expected, actual), "argument not equal to expected: expected '%v'; actual '%v'", expected, actual) {
		}
	}

	if check(t, args1.FlagIsSpecified("-f2"), "flag '-f2' should be present") {

		unused = args1.GetUnusedFlags()
		if expected, actual := 2, len(unused); check(t, expected == actual, "unused slice has wrong size: expected %v; actual %v", expected, actual) {

			if expected, actual := (clasp.Argument{ Type : clasp.Flag, ResolvedName : "-f1", CmdLineIndex : 1 }), *unused[0]; check(t, equalInNonNillLhs(expected, actual), "argument not equal to expected: expected '%v'; actual '%v'", expected, actual) {
			}
			if expected, actual := (clasp.Argument{ Type : clasp.Flag, ResolvedName : "---flag3", CmdLineIndex : 4 }), *unused[1]; check(t, equalInNonNillLhs(expected, actual), "argument not equal to expected: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
}


func Test_find_unused_options_1(t *testing.T) {

	argv := []string { "path/blah", "-o1=v1", "-o2=v2", "abc", "---option3=value3" }

	args1 := clasp.Parse(argv, clasp.ParseParams{})

	unused := args1.GetUnusedOptions()
	if expected, actual := 3, len(unused); check(t, expected == actual, "unused slice has wrong size: expected %v; actual %v", expected, actual) {

		if expected, actual := (clasp.Argument{ Type : clasp.Option, ResolvedName : "-o1", CmdLineIndex : 1 }), *unused[0]; check(t, equalInNonNillLhs(expected, actual), "argument not equal to expected: expected '%v'; actual '%v'", expected, actual) {
		}
		if expected, actual := (clasp.Argument{ Type : clasp.Option, ResolvedName : "-o2", CmdLineIndex : 2 }), *unused[1]; check(t, equalInNonNillLhs(expected, actual), "argument not equal to expected: expected '%v'; actual '%v'", expected, actual) {
		}
		if expected, actual := (clasp.Argument{ Type : clasp.Option, ResolvedName : "---option3", CmdLineIndex : 4 }), *unused[2]; check(t, equalInNonNillLhs(expected, actual), "argument not equal to expected: expected '%v'; actual '%v'", expected, actual) {
		}
	}

	if _, isFound := args1.LookupOption("-o2"); check(t, isFound, "option '-o2' should be present") {

		unused = args1.GetUnusedOptions()
		if expected, actual := 2, len(unused); check(t, expected == actual, "unused slice has wrong size: expected %v; actual %v", expected, actual) {

			if expected, actual := (clasp.Argument{ Type : clasp.Option, ResolvedName : "-o1", CmdLineIndex : 1 }), *unused[0]; check(t, equalInNonNillLhs(expected, actual), "argument not equal to expected: expected '%v'; actual '%v'", expected, actual) {
			}
			if expected, actual := (clasp.Argument{ Type : clasp.Option, ResolvedName : "---option3", CmdLineIndex : 4 }), *unused[1]; check(t, equalInNonNillLhs(expected, actual), "argument not equal to expected: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
}


func Test_CheckAllFlagBits(t *testing.T) {

	aliases := []clasp.Alias {

		{ clasp.Flag, "-f1", nil, "", nil, 0x01 },
		{ clasp.Flag, "-f2", nil, "", nil, 0x02 },
		{ clasp.Flag, "-f4", nil, "", nil, 0x04 },
	}

	argv1 := []string { "path/blah", "-f1", "-f2", "-f4" }
	args1 := clasp.Parse(argv1, clasp.ParseParams{ Aliases: aliases })
	if expected, actual := 0x7, args1.CheckAllFlagBits(nil); check(t, expected == actual, "flags not as expected: expected %x; actual %x", expected, actual) {
	}

	argv2 := []string { "path/blah", "-f1", "-f4" }
	args2 := clasp.Parse(argv2, clasp.ParseParams{ Aliases: aliases })
	if expected, actual := 0x5, args2.CheckAllFlagBits(nil); check(t, expected == actual, "flags not as expected: expected %x; actual %x", expected, actual) {
	}
}


func Test_groupedFlags_1(t *testing.T) {

	aliases := []clasp.Alias {

		{ clasp.Flag, "--high", []string { "-h" }, "second flag", nil, 0 },
		{ clasp.Flag, "--mid",  []string { "-m" }, "second flag", nil, 0 },
		{ clasp.Flag, "--low",  []string { "-l" }, "second flag", nil, 0 },
	}
	argv := []string { "path/blah", "-hm", "-l" }

	args1 := clasp.Parse(argv, clasp.ParseParams{ Aliases: aliases })

	if expected, actual := 3, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {

		argument0 := args1.Arguments[0]
		check(t, "" == argument0.Value, "arguments has wrong argument")
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
	}
	if expected := 3; check(t, expected == len(args1.Flags), "arguments object has wrong number of Flags: expected %v; actual %v", expected, len(args1.Flags)) {

		flag0 := args1.Flags[0]
		flag1 := args1.Flags[1]
		flag2 := args1.Flags[2]
		check(t, 1 == flag0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag0.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag0.Type)
		check(t, "--high" == flag0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-hm" == flag0.GivenName, "arguments has wrong given name")
		check(t, "" == flag0.Value, "arguments has wrong value")
		check(t, 1 == flag1.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag1.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag1.Type)
		check(t, "--mid" == flag1.ResolvedName, "arguments has wrong resolved name")
		check(t, "-hm" == flag1.GivenName, "arguments has wrong given name")
		check(t, "" == flag1.Value, "arguments has wrong value")
		check(t, 2 == flag2.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag2.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag2.Type)
		check(t, "--low" == flag2.ResolvedName, "arguments has wrong resolved name")
		check(t, "-l" == flag2.GivenName, "arguments has wrong given name")
		check(t, "" == flag2.Value, "arguments has wrong value")
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 3, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {

		for i, expected := range argv {

			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}


func Test_groupedFlags_2(t *testing.T) {

	aliases := []clasp.Alias {

		{ clasp.Flag, "--high", []string { "-h" }, "second flag", nil, 0 },
		{ clasp.Flag, "--mid",  []string { "-m" }, "second flag", nil, 0 },
		{ clasp.Flag, "--low",  []string { "-l" }, "second flag", nil, 0 },
	}
	argv := []string { "path/blah", "-hmx", "-l" }

	args1 := clasp.Parse(argv, clasp.ParseParams{ Aliases: aliases })

	if expected, actual := 2, len(args1.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {

		argument0 := args1.Arguments[0]
		check(t, "" == argument0.Value, "arguments has wrong argument")
		check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
	}
	if expected := 2; check(t, expected == len(args1.Flags), "arguments object has wrong number of Flags: expected %v; actual %v", expected, len(args1.Flags)) {

		flag0 := args1.Flags[0]
		flag1 := args1.Flags[1]
		check(t, 1 == flag0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag0.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag0.Type)
		check(t, "-hmx" == flag0.ResolvedName, "arguments has wrong resolved name")
		check(t, "-hmx" == flag0.GivenName, "arguments has wrong given name")
		check(t, "" == flag0.Value, "arguments has wrong value")
		check(t, 2 == flag1.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Flag == flag1.Type, "argument has wrong type: expected: %v; received %v", clasp.Flag, flag1.Type)
		check(t, "--low" == flag1.ResolvedName, "arguments has wrong resolved name")
		check(t, "-l" == flag1.GivenName, "arguments has wrong given name")
		check(t, "" == flag1.Value, "arguments has wrong value")
	}
	if expected, actual := 0, len(args1.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 0, len(args1.Values); check(t, expected == actual, "arguments object has wrong number of Values: expected %v; actual %v", expected, actual) {
	}
	if expected, actual := 3, len(args1.Argv); check(t, expected == actual, "arguments object has wrong number of Argv members: expected %v; actual %v", expected, actual) {

		for i, expected := range argv {

			if actual := args1.Argv[i]; check(t, expected == actual, "arguments has wrong Argv member: expected '%v'; actual '%v'", expected, actual) {
			}
		}
	}
	if expected, actual := "blah", args1.ProgramName; check(t, expected == actual, "arguments object has wrong program name: expected '%v'; actual '%v'", expected, actual) {
	}
}

func Test_flag_alias_of_option_with_value(t *testing.T) {

	aliases	:=	[]clasp.Alias {

		{ clasp.Option, "--verbosity", nil, "Specifies the verbosity", []string{ "low", "medium", "high" }, 0 },
		{ clasp.Flag, "--verbosity=high", []string{ "-v" }, "", nil, 0 },
	}
	argv	:=	[]string{ "path/blah", "-v" }

	args	:=	clasp.Parse(argv, clasp.ParseParams{ Aliases: aliases })

	if expected, actual := 1, len(args.Arguments); check(t, expected == actual, "arguments object has wrong number of Arguments: expected %v; actual %v", expected, actual) {

		argument0 := args.Arguments[0]
		check(t, 1 == argument0.CmdLineIndex, "argument0 has wrong CmdLineIndex")
		//check(t, "high" == argument0.Value, "argument0 has wrong value")
	}

	if expected, actual := 0, len(args.Flags); check(t, expected == actual, "arguments object has wrong number of Flags: expected %v; actual %v", expected, actual) {
	}

	if expected, actual := 1, len(args.Options); check(t, expected == actual, "arguments object has wrong number of Options: expected %v; actual %v", expected, actual) {

		option0 := args.Options[0]
		check(t, 1 == option0.CmdLineIndex, "arguments has wrong command-line index")
		check(t, clasp.Option == option0.Type, "argument has wrong type: expected: %v; received %v", clasp.Option, option0.Type)
		check(t, "--verbosity" == option0.ResolvedName, "arguments has wrong resolved name")
		check(t, "--verbosity=high" == option0.GivenName, "arguments has wrong given name")
		check(t, "high" == option0.Value, "arguments has wrong value")
	}
}


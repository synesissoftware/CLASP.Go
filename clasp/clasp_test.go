
package clasp

import "fmt"
import "path"
import "runtime"
import "testing"

func check(t *testing.T, cond bool, format string, args... interface{}) bool {
	if !cond {
		_, file, line, hasCallInfo := runtime.Caller(1)

		if hasCallInfo {
			fmt.Printf("\t%s:%d: %s\n", path.Base(file), line, fmt.Sprintf(format, args...))
		} else {
			fmt.Printf("\t%s\n", fmt.Sprintf(format, args...))
		}

		t.FailNow()
	}
	return cond
}


func Test_no_args(t *testing.T) {

	case1_argv := []string { "path/blah" }

	args1 := Parse(case1_argv, ParseParams{})

	check(t, 0 == len(args1.Arguments), "arguments object has wrong number of arguments")
	check(t, 0 == len(args1.Flags), "arguments object has wrong number of flags")
	check(t, 0 == len(args1.Options), "arguments object has wrong number of options")
	check(t, 0 == len(args1.Values), "arguments object has wrong number of values")
	check(t, 1 == len(args1.Argv), "arguments object has wrong number of argv members")
	check(t, case1_argv[0] == args1.Argv[0], "arguments object has wrong executable name")
	check(t, "blah" == args1.ProgramName, "arguments object has wrong program name: has '%s' and should be '%s'", args1.ProgramName, "blah")
}


func Test_single_value(t *testing.T) {

	case1_argv := []string { "path/blah", "abc" }

	args1 := Parse(case1_argv, ParseParams{})

	check(t, 1 == len(args1.Arguments), "arguments object has wrong number of arguments")
	argument0 := args1.Arguments[0]
	check(t, "abc" == argument0.Value, "arguments has wrong argument")
	check(t, 1 == argument0.CmdLineIndex, "arguments has wrong argument")
	check(t, 0 == len(args1.Flags), "arguments object has wrong number of flags")
	check(t, 0 == len(args1.Options), "arguments object has wrong number of options")
	check(t, 1 == len(args1.Values), "arguments object has wrong number of values")
	value0 := args1.Values[0]
	check(t, "abc" == value0.Value, "arguments has wrong value")
	check(t, 1 == value0.CmdLineIndex, "arguments has wrong command-line index")
	check(t, 2 == len(args1.Argv), "arguments object has wrong number of argv members")
	check(t, case1_argv[0] == args1.Argv[0], "arguments object has wrong executable name")
	check(t, "blah" == args1.ProgramName, "arguments object has wrong program name: has '%s' and should be '%s'", args1.ProgramName, "blah")
}


func Test_several_values(t *testing.T) {

	case1_argv := []string { "path/blah", "abc", "def", "g", "hi" }

	args1 := Parse(case1_argv, ParseParams{})

	check(t, 4 == len(args1.Arguments), "arguments object has wrong number of arguments")
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

	check(t, 0 == len(args1.Flags), "arguments object has wrong number of flags")

	check(t, 0 == len(args1.Options), "arguments object has wrong number of options")

	check(t, 4 == len(args1.Values), "arguments object has wrong number of values")
	value0 := args1.Values[0]
	value1 := args1.Values[1]
	value2 := args1.Values[2]
	value3 := args1.Values[3]
	check(t, "abc" == value0.Value, "arguments has wrong value")
	check(t, 1 == value0.CmdLineIndex, "arguments has wrong command-line index")
	check(t, "def" == value1.Value, "arguments has wrong value")
	check(t, 2 == value1.CmdLineIndex, "arguments has wrong command-line index")
	check(t, "g" == value2.Value, "arguments has wrong value")
	check(t, 3 == value2.CmdLineIndex, "arguments has wrong command-line index")
	check(t, "hi" == value3.Value, "arguments has wrong value")
	check(t, 4 == value3.CmdLineIndex, "arguments has wrong command-line index")

	check(t, 5 == len(args1.Argv), "arguments object has wrong number of argv members")
	check(t, case1_argv[0] == args1.Argv[0], "arguments object has wrong executable name")
	check(t, case1_argv[1] == args1.Argv[1], "arguments object has wrong executable name")
	check(t, case1_argv[2] == args1.Argv[2], "arguments object has wrong executable name")
	check(t, case1_argv[3] == args1.Argv[3], "arguments object has wrong executable name")
	check(t, case1_argv[4] == args1.Argv[4], "arguments object has wrong executable name")

	check(t, "blah" == args1.ProgramName, "arguments object has wrong program name: has '%s' and should be '%s'", args1.ProgramName, "blah")
}


func Test_several_values_with_double_hyphen(t *testing.T) {

	case1_argv := []string { "path/blah", "abc", "def", "--", "-g", "--hi" }

	args1 := Parse(case1_argv, ParseParams{})

	check(t, 4 == len(args1.Arguments), "arguments object has wrong number of arguments: expected %d; has %d", 4, len(args1.Arguments))
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

	check(t, 0 == len(args1.Flags), "arguments object has wrong number of flags")

	check(t, 0 == len(args1.Options), "arguments object has wrong number of options")

	check(t, 4 == len(args1.Values), "arguments object has wrong number of values")
	value0 := args1.Values[0]
	value1 := args1.Values[1]
	value2 := args1.Values[2]
	value3 := args1.Values[3]
	check(t, "abc" == value0.Value, "arguments has wrong value")
	check(t, 1 == value0.CmdLineIndex, "arguments has wrong command-line index")
	check(t, 0 == value0.NumGivenHyphens, "arguments has wrong number of hyphens")
	check(t, "def" == value1.Value, "arguments has wrong value")
	check(t, 2 == value1.CmdLineIndex, "arguments has wrong command-line index")
	check(t, 0 == value1.NumGivenHyphens, "arguments has wrong number of hyphens")
	check(t, "-g" == value2.Value, "arguments has wrong value")
	check(t, 4 == value2.CmdLineIndex, "arguments has wrong command-line index")
	check(t, 0 == value2.NumGivenHyphens, "arguments has wrong number of hyphens")
	check(t, "--hi" == value3.Value, "arguments has wrong value")
	check(t, 5 == value3.CmdLineIndex, "arguments has wrong command-line index")
	check(t, 0 == value3.NumGivenHyphens, "arguments has wrong number of hyphens")

	check(t, 6 == len(args1.Argv), "arguments object has wrong number of argv members")
	check(t, case1_argv[0] == args1.Argv[0], "arguments object has wrong executable name")
	check(t, case1_argv[1] == args1.Argv[1], "arguments object has wrong executable name")
	check(t, case1_argv[2] == args1.Argv[2], "arguments object has wrong executable name")
	check(t, case1_argv[3] == args1.Argv[3], "arguments object has wrong executable name")
	check(t, case1_argv[4] == args1.Argv[4], "arguments object has wrong executable name")
	check(t, case1_argv[5] == args1.Argv[5], "arguments object has wrong executable name")

	check(t, "blah" == args1.ProgramName, "arguments object has wrong program name: has '%s' and should be '%s'", args1.ProgramName, "blah")
}


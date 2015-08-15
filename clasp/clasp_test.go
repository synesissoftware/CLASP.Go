
package clasp

import "testing"

func check(t *testing.T, cond bool, fmt string, args... interface{}) bool {
	if !cond {
		t.Errorf(fmt, args...)
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


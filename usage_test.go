
package clasp_test

import (

	clasp "github.com/synesissoftware/CLASP.Go"

	"bytes"
	"os"
	"path"
	"strings"
	"testing"
)

func filter_AS(lines []string, f func(line string) bool) (r []string) {

	for _, line := range lines {

		if f(line) {

			r = append(r, line)
		}
	}

	return r
}

func check_num_lines(t *testing.T, result []string, num_lines int) {

	if num_lines != len(result) {

		t.Errorf("expected number of lines %d does not equal actual %d", num_lines, len(result))
	}
}

func check_num_nonblank_lines(t *testing.T, result []string, num_lines int) {

	nb_lines := filter_AS(result, func(line string) bool { return 0 != len(line) })

	if num_lines != len(nb_lines) {

		t.Errorf("expected number of non-blank lines %d does not equal actual %d", num_lines, len(nb_lines))
	}
}

func check_line_equal(t *testing.T, actual, expected string) {

	if actual != expected {

		t.Errorf("expected line '%s' does not equal actual '%s'", expected, actual)
	}
}

func check_stripped_line_equal(t *testing.T, actual, expected string) {

	if strings.TrimSpace(actual) != strings.TrimSpace(expected) {

		t.Errorf("expected line '%s' does not equal, when ignoring leading/trailing space, actual '%s'", expected, actual)
	}
}

func call_ShowUsage_(t *testing.T, aliases []clasp.Alias, ups clasp.UsageParams) (result []string, err error) {

	buf			:=	new(bytes.Buffer)

	ups.Stream = buf

	var xc	int;

	xc, err		=	clasp.ShowUsage(aliases, ups)
	if err != nil {

		t.Errorf("ShowUsage() failed with exit code: %d", xc)
	} else {

		if 0 != xc {

			t.Error("return code is not 0")
		} else {

			s := buf.String()

			result = strings.Split(s, "\n")
		}
	}

	return
}

func test_ShowVersion_(t *testing.T, expected string, aliases []clasp.Alias, program_name string, version interface{}, version_prefix string) {

	buf			:=	new(bytes.Buffer)
	params		:=	clasp.UsageParams { Stream: buf, ProgramName: program_name, UsageFlags: 0, ValuesString: "", ExitCode: 0, Exiter: nil, Version: version, VersionPrefix: version_prefix }

	xc, err		:=	clasp.ShowVersion(aliases, params)
	if err != nil {

		t.Errorf("ShowVersion() failed with exit code: %d", xc)
	} else {

		if 0 != xc {

			t.Error("return code is not 0")
		} else {

			actual		:=	buf.String()

			if expected != actual {

				t.Errorf("expected '%s' does not equal actual '%s'", expected, actual)
			}
		}
	}
}

func Test_ShowVersion_with_int_array_version(t *testing.T) {

	var aliases		[]clasp.Alias

	test_ShowVersion_(t, "myprog 1.2.3\n", aliases, "myprog", []int{ 1, 2, 3 }, "")

	test_ShowVersion_(t, "myprog v1.2.3\n", aliases, "myprog", []int{ 1, 2, 3 }, "v")
}

func Test_ShowVersion_with_string_array_version(t *testing.T) {

	var aliases		[]clasp.Alias

	test_ShowVersion_(t, "myprog 0.1.9\n", aliases, "myprog", []string{ "0", "1", "9" }, "")

	test_ShowVersion_(t, "myprog v0.1.9\n", aliases, "myprog", []string{ "0", "1", "9" }, "v")
}

func Test_ShowVersion_with_string_version(t *testing.T) {

	var aliases		[]clasp.Alias

	test_ShowVersion_(t, "myprog 0.1.9\n", aliases, "myprog", "0.1.9", "")

	test_ShowVersion_(t, "myprog v0.1.9\n", aliases, "myprog", "0.1.9", "v")
}

func Test_ShowVersion_with_inferred_ProgramName(t *testing.T) {

	var aliases		[]clasp.Alias

	program_name	:=	os.Args[0]
	program_name	=	path.Base(program_name)

	test_ShowVersion_(t, program_name + " 0.1.2\n", aliases, "", "0.1.2", "")
}

func Test_ShowUsage_1(t *testing.T) {

	var	aliases					[]clasp.Alias
	values_string				:=	""
	flags_and_options_string	:=	""
	info_lines					:=	make([]string, 0)

	usage_params_base			:=	clasp.UsageParams {

		ProgramName:			"myprogram",
		VersionPrefix:			"v",
		Version:				"0.1.2",
		ExitCode:				0,
		ValuesString:			values_string,
		FlagsAndOptionsString:	flags_and_options_string,
		InfoLines:				info_lines,
	}

	result, err					:=	call_ShowUsage_(t, aliases, usage_params_base)
	if err != nil {

	} else {

		check_num_nonblank_lines(t, result, 1)
		check_num_lines(t, result, 2)

		check_line_equal(t, result[0], "USAGE: myprogram")
	}
}

func Test_ShowUsage_2(t *testing.T) {

	var	aliases					[]clasp.Alias
	values_string				:=	""
	flags_and_options_string	:=	""
	info_lines					:=	[]string {

		"CLASP.Go Test Suite",
		"",
		":version:",
		"",
	}

	usage_params_base			:=	clasp.UsageParams {

		ProgramName:			"myprogram",
		VersionPrefix:			"V",
		Version:				"0.1.2",
		ExitCode:				0,
		ValuesString:			values_string,
		FlagsAndOptionsString:	flags_and_options_string,
		InfoLines:				info_lines,
	}

	result, err					:=	call_ShowUsage_(t, aliases, usage_params_base)
	if err != nil {

	} else {

		check_num_nonblank_lines(t, result, 3)
		check_num_lines(t, result, 6)

		check_line_equal(t, result[0], "CLASP.Go Test Suite")
		check_line_equal(t, result[1], "")
		check_line_equal(t, result[2], "myprogram V0.1.2")
		check_line_equal(t, result[3], "")
		check_line_equal(t, result[4], "USAGE: myprogram")
	}
}

func Test_ShowUsage_3(t *testing.T) {

	aliases						:=	[]clasp.Alias{

		{ clasp.Flag, "--high", []string { "-h" }, "makes things high", nil, 0 },
	}
	values_string				:=	""
	flags_and_options_string	:=	""
	info_lines					:=	make([]string, 0)

	usage_params_base			:=	clasp.UsageParams {

		ProgramName:			"myprogram",
		VersionPrefix:			"v",
		Version:				"0.1.2",
		ExitCode:				0,
		ValuesString:			values_string,
		FlagsAndOptionsString:	flags_and_options_string,
		InfoLines:				info_lines,
	}

	result, err					:=	call_ShowUsage_(t, aliases, usage_params_base)
	if err != nil {

	} else {

		check_num_nonblank_lines(t, result, 5)
		check_num_lines(t, result, 9)

		check_line_equal(t, result[0], "USAGE: myprogram")
		check_line_equal(t, result[2], "flags/options:")
		check_stripped_line_equal(t, result[4], "-h")
		check_stripped_line_equal(t, result[5], "--high")
	}
}


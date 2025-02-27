package clasp_test

import (
	angols "github.com/synesissoftware/ANGoLS"
	clasp "github.com/synesissoftware/CLASP.Go"

	"bytes"
	"os"
	"path"
	"strings"
	"testing"
)

func check_num_lines(t *testing.T, result []string, num_lines int) {

	t.Helper()

	if num_lines != len(result) {

		t.Errorf("expected number of lines %d does not equal actual %d", num_lines, len(result))
	}
}

func check_num_nonblank_lines(t *testing.T, result []string, num_lines int) {

	t.Helper()

	nb_lines, _ := angols.SelectSliceOfString(result, func(index int, line string) (bool, error) {

		return 0 != len(line), nil
	})

	if num_lines != len(nb_lines) {

		t.Errorf("expected number of non-blank lines %d does not equal actual %d", num_lines, len(nb_lines))
	}
}

func check_line_equal(t *testing.T, actual, expected string) {

	t.Helper()

	if actual != expected {

		t.Errorf("expected line '%s' does not equal actual '%s'", expected, actual)
	}
}

func check_stripped_line_equal(t *testing.T, actual, expected string) {

	t.Helper()

	if strings.TrimSpace(actual) != strings.TrimSpace(expected) {

		t.Errorf("expected line '%s' does not equal, when ignoring leading/trailing space, actual '%s'", expected, actual)
	}
}

func call_ShowUsage_(t *testing.T, specifications []clasp.Specification, ups clasp.UsageParams) (result []string, err error) {

	t.Helper()

	buf := new(bytes.Buffer)

	ups.Stream = buf
	ups.UsageFlags |= clasp.DontCallExit

	var xc int

	xc, err = clasp.ShowUsage(specifications, ups)
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

func test_ShowVersion_(t *testing.T, expected string, specifications []clasp.Specification, program_name string, version interface{}, version_prefix string) {

	t.Helper()

	buf := new(bytes.Buffer)
	params := clasp.UsageParams{

		Stream:        buf,
		ProgramName:   program_name,
		UsageFlags:    clasp.DontCallExit,
		ValuesString:  "",
		ExitCode:      0,
		Exiter:        nil,
		Version:       version,
		VersionPrefix: version_prefix,
	}

	xc, err := clasp.ShowVersion(specifications, params)
	if err != nil {

		t.Errorf("ShowVersion() failed with exit code: %d", xc)
	} else {

		if 0 != xc {

			t.Error("return code is not 0")
		} else {

			actual := buf.String()

			if expected != actual {

				t.Errorf("expected '%s' does not equal actual '%s'", expected, actual)
			}
		}
	}
}

func Test_ShowVersion_with_int_array_version(t *testing.T) {

	var specifications []clasp.Specification

	test_ShowVersion_(t, "myprog 1.2.3\n", specifications, "myprog", []int{1, 2, 3}, "")

	test_ShowVersion_(t, "myprog v1.2.3\n", specifications, "myprog", []int{1, 2, 3}, "v")
}

func Test_ShowVersion_with_uint16_array_version(t *testing.T) {

	var specifications []clasp.Specification

	test_ShowVersion_(t, "myprog 1.2.3\n", specifications, "myprog", []uint16{1, 2, 3}, "")

	test_ShowVersion_(t, "myprog v1.2.3\n", specifications, "myprog", []uint16{1, 2, 3}, "v")
}

func Test_ShowVersion_with_string_array_version(t *testing.T) {

	var specifications []clasp.Specification

	test_ShowVersion_(t, "myprog 0.1.9\n", specifications, "myprog", []string{"0", "1", "9"}, "")

	test_ShowVersion_(t, "myprog v0.1.9\n", specifications, "myprog", []string{"0", "1", "9"}, "v")
}

func Test_ShowVersion_with_string_version(t *testing.T) {

	var specifications []clasp.Specification

	test_ShowVersion_(t, "myprog 0.1.9\n", specifications, "myprog", "0.1.9", "")

	test_ShowVersion_(t, "myprog v0.1.9\n", specifications, "myprog", "0.1.9", "v")
}

func Test_ShowVersion_with_inferred_ProgramName(t *testing.T) {

	var specifications []clasp.Specification

	program_name := os.Args[0]
	program_name = path.Base(program_name)

	test_ShowVersion_(t, program_name+" 0.1.2\n", specifications, "", "0.1.2", "")
}

func Test_ShowUsage_1(t *testing.T) {

	var specifications []clasp.Specification
	values_string := ""
	flags_and_options_string := ""
	info_lines := make([]string, 0)

	usage_params_base := clasp.UsageParams{

		ProgramName:           "myprogram",
		VersionPrefix:         "v",
		Version:               "0.1.2",
		ExitCode:              0,
		ValuesString:          values_string,
		FlagsAndOptionsString: flags_and_options_string,
		InfoLines:             info_lines,
	}

	result, err := call_ShowUsage_(t, specifications, usage_params_base)
	if err != nil {

		t.Fail()
	} else {

		check_num_nonblank_lines(t, result, 1)
		check_num_lines(t, result, 2)

		check_line_equal(t, result[0], "USAGE: myprogram")
	}
}

func Test_ShowUsage_2(t *testing.T) {

	var specifications []clasp.Specification
	values_string := ""
	flags_and_options_string := ""
	info_lines := []string{

		"CLASP.Go Test Suite",
		"",
		":version:",
		"",
	}

	usage_params_base := clasp.UsageParams{

		ProgramName:           "myprogram",
		VersionPrefix:         "V",
		Version:               []int{0, 1, 2},
		ExitCode:              0,
		ValuesString:          values_string,
		FlagsAndOptionsString: flags_and_options_string,
		InfoLines:             info_lines,
	}

	result, err := call_ShowUsage_(t, specifications, usage_params_base)
	if err != nil {

		t.Fail()
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

func Test_ShowUsage_3_a(t *testing.T) {

	specifications := []clasp.Specification{

		clasp.Flag("--high").SetAlias("-h").SetHelp("makes things high"),
	}
	values_string := ""
	flags_and_options_string := ""
	info_lines := make([]string, 0)

	usage_params_base := clasp.UsageParams{

		ProgramName:           "myprogram",
		VersionPrefix:         "v",
		Version:               "0.1.2",
		ExitCode:              0,
		ValuesString:          values_string,
		FlagsAndOptionsString: flags_and_options_string,
		InfoLines:             info_lines,
	}

	result, err := call_ShowUsage_(t, specifications, usage_params_base)
	if err != nil {

		t.Fail()
	} else {

		check_num_nonblank_lines(t, result, 5)
		check_num_lines(t, result, 9)

		check_line_equal(t, result[0], "USAGE: myprogram [ ... flags and options ... ]")
		check_line_equal(t, result[2], "flags/options:")
		check_stripped_line_equal(t, result[4], "-h")
		check_stripped_line_equal(t, result[5], "--high")
		check_stripped_line_equal(t, result[6], "makes things high")
	}
}

func Test_ShowUsage_3_b(t *testing.T) {

	specifications := []clasp.Specification{

		clasp.Flag("--high").SetAlias("-h").SetHelp("makes things high"),
	}
	values_string := ""
	flags_and_options_string := " "
	info_lines := make([]string, 0)

	usage_params_base := clasp.UsageParams{

		ProgramName:           "myprogram",
		VersionPrefix:         "v",
		Version:               "0.1.2",
		ExitCode:              0,
		ValuesString:          values_string,
		FlagsAndOptionsString: flags_and_options_string,
		InfoLines:             info_lines,
	}

	result, err := call_ShowUsage_(t, specifications, usage_params_base)
	if err != nil {

		t.Fail()
	} else {

		check_num_nonblank_lines(t, result, 5)
		check_num_lines(t, result, 9)

		check_line_equal(t, result[0], "USAGE: myprogram")
		check_line_equal(t, result[2], "flags/options:")
		check_stripped_line_equal(t, result[4], "-h")
		check_stripped_line_equal(t, result[5], "--high")
		check_stripped_line_equal(t, result[6], "makes things high")
	}
}

func Test_ShowUsage_3_c(t *testing.T) {

	specifications := []clasp.Specification{

		clasp.Flag("--high").SetAlias("-h").SetHelp("makes things high"),
	}
	values_string := ""
	flags_and_options_string := "[ flags/options ]"
	info_lines := make([]string, 0)

	usage_params_base := clasp.UsageParams{

		ProgramName:           "myprogram",
		VersionPrefix:         "v",
		Version:               "0.1.2",
		ExitCode:              0,
		ValuesString:          values_string,
		FlagsAndOptionsString: flags_and_options_string,
		InfoLines:             info_lines,
	}

	result, err := call_ShowUsage_(t, specifications, usage_params_base)
	if err != nil {

		t.Fail()
	} else {

		check_num_nonblank_lines(t, result, 5)
		check_num_lines(t, result, 9)

		check_line_equal(t, result[0], "USAGE: myprogram [ flags/options ]")
		check_line_equal(t, result[2], "flags/options:")
		check_stripped_line_equal(t, result[4], "-h")
		check_stripped_line_equal(t, result[5], "--high")
		check_stripped_line_equal(t, result[6], "makes things high")
	}
}

func Test_ShowUsage_4(t *testing.T) {

	specifications := []clasp.Specification{

		clasp.Option("--verbosity").SetHelp("Specifies the verbosity").SetValues("low", "medium", "high"),
		clasp.Flag("--verbosity=high").SetAlias("-v"),
	}
	values_string := ""
	flags_and_options_string := "[ flags/options ]"
	info_lines := make([]string, 0)

	usage_params_base := clasp.UsageParams{

		ProgramName:           "myprogram",
		VersionPrefix:         "v",
		Version:               []int{0, 1, 2},
		ExitCode:              0,
		ValuesString:          values_string,
		FlagsAndOptionsString: flags_and_options_string,
		InfoLines:             info_lines,
	}

	result, err := call_ShowUsage_(t, specifications, usage_params_base)
	if err != nil {

		t.Fail()
	} else {

		check_num_nonblank_lines(t, result, 9)
		check_num_lines(t, result, 13)

		check_line_equal(t, result[0], "USAGE: myprogram [ flags/options ]")
		check_line_equal(t, result[2], "flags/options:")
		check_stripped_line_equal(t, result[4], "-v --verbosity=high")
		check_stripped_line_equal(t, result[5], "--verbosity=<value>")
		check_stripped_line_equal(t, result[6], "Specifies the verbosity")
		check_stripped_line_equal(t, result[7], "where <value> one of:")
		check_stripped_line_equal(t, result[8], "low")
		check_stripped_line_equal(t, result[9], "medium")
		check_stripped_line_equal(t, result[10], "high")
	}
}

func Test_ShowUsage_5_a(t *testing.T) {

	specifications := []clasp.Specification{

		clasp.Section("verbosity:"),
		clasp.Option("--verbosity").SetHelp("Specifies the verbosity").SetValues("low", "medium", "high"),
		clasp.Flag("--verbosity=high").SetAlias("-v"),
		clasp.Flag("--debug"),
	}
	values_string := ""
	flags_and_options_string := "[ flags/options ]"
	info_lines := []string{"CLASP.Go Examples", "", ":version:", ""}

	usage_params_base := clasp.UsageParams{

		ProgramName:           "myprogram",
		VersionPrefix:         "v",
		Version:               "0.1.2",
		ExitCode:              0,
		ValuesString:          values_string,
		FlagsAndOptionsString: flags_and_options_string,
		InfoLines:             info_lines,
		UsageFlags:            0,
	}

	result, err := call_ShowUsage_(t, specifications, usage_params_base)
	if err != nil {

		t.Fail()
	} else {

		check_num_nonblank_lines(t, result, 13)
		check_num_lines(t, result, 21)

		check_line_equal(t, result[0], "CLASP.Go Examples")
		check_line_equal(t, result[1], "")
		check_line_equal(t, result[2], "myprogram v0.1.2")
		check_line_equal(t, result[3], "")
		check_line_equal(t, result[4], "USAGE: myprogram [ flags/options ]")
		check_line_equal(t, result[5], "")
		check_line_equal(t, result[6], "flags/options:")
		check_line_equal(t, result[7], "")
		check_stripped_line_equal(t, result[8], "verbosity:")
		check_line_equal(t, result[9], "")
		check_stripped_line_equal(t, result[10], "-v --verbosity=high")
		check_stripped_line_equal(t, result[11], "--verbosity=<value>")
		check_stripped_line_equal(t, result[12], "Specifies the verbosity")
		check_stripped_line_equal(t, result[13], "where <value> one of:")
		check_stripped_line_equal(t, result[14], "low")
		check_stripped_line_equal(t, result[15], "medium")
		check_stripped_line_equal(t, result[16], "high")
		check_line_equal(t, result[17], "")
		check_stripped_line_equal(t, result[18], "--debug")
		check_line_equal(t, result[19], "")
		check_line_equal(t, result[20], "")
	}
}

func Test_ShowUsage_5_b(t *testing.T) {

	specifications := []clasp.Specification{

		clasp.Section("verbosity:"),
		clasp.Option("--verbosity").SetHelp("Specifies the verbosity").SetValues("low", "medium", "high"),
		clasp.Flag("--verbosity=high").SetAlias("-v"),
		clasp.Flag("--debug"),
	}
	values_string := ""
	flags_and_options_string := "[ flags/options ]"
	info_lines := make([]string, 0)

	usage_params_base := clasp.UsageParams{

		ProgramName:           "myprogram",
		VersionPrefix:         "v",
		Version:               "0.1.2",
		ExitCode:              0,
		ValuesString:          values_string,
		FlagsAndOptionsString: flags_and_options_string,
		InfoLines:             info_lines,
		UsageFlags:            clasp.SkipBlanksBetweenLines,
	}

	result, err := call_ShowUsage_(t, specifications, usage_params_base)
	if err != nil {

		t.Fail()
	} else {

		check_num_nonblank_lines(t, result, 11)
		check_num_lines(t, result, 15)

		check_line_equal(t, result[0], "USAGE: myprogram [ flags/options ]")
		check_line_equal(t, result[2], "flags/options:")
		check_stripped_line_equal(t, result[4], "verbosity:")
		check_stripped_line_equal(t, result[6], "-v --verbosity=high")
		check_stripped_line_equal(t, result[7], "--verbosity=<value>")
		check_stripped_line_equal(t, result[8], "Specifies the verbosity")
		check_stripped_line_equal(t, result[9], "where <value> one of:")
		check_stripped_line_equal(t, result[10], "low")
		check_stripped_line_equal(t, result[11], "medium")
		check_stripped_line_equal(t, result[12], "high")
		check_stripped_line_equal(t, result[13], "--debug")
	}
}


package clasp_test

import (

	clasp "github.com/synesissoftware/CLASP.Go"

	"bytes"
	"os"
	"path"
	"testing"
)

func test_ShowVersion_(t *testing.T, expected string, aliases []clasp.Alias, program_name string, version interface{}, version_prefix string) {

	buf			:=	new(bytes.Buffer)
	params		:=	clasp.UsageParams { Stream: buf, ProgramName: program_name, Flags: 0, Values: "", ExitCode: 0, Exiter: nil, Version: version, VersionPrefix: version_prefix }

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



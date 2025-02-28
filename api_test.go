package clasp_test

import (
	"github.com/stretchr/testify/require"
	clasp "github.com/synesissoftware/CLASP.Go"

	"testing"
)

/* /////////////////////////////////////////////////////////////////////////
 * helper functions
 */

/* /////////////////////////////////////////////////////////////////////////
 * test functions
 */

func Test_PARSE_Flags_1(t *testing.T) {
	require.Equal(t, int(0), int(clasp.Parse_None))

	require.NotEqual(t, clasp.Parse_TreatSingleHyphenAsValue, clasp.Parse_DontRecogniseDoubleHyphenToStartValues)
	require.Equal(t, int(0), int(clasp.Parse_TreatSingleHyphenAsValue&clasp.Parse_DontRecogniseDoubleHyphenToStartValues))
}

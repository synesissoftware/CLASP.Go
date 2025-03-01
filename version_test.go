package clasp

import (
	"github.com/stretchr/testify/require"

	"testing"
)

const (
	Expected_VersionMajor uint16 = 0
	Expected_VersionMinor uint16 = 17
	Expected_VersionPatch uint16 = 0
	Expected_VersionAB    uint16 = 0x4002
)

func Test_Version_Elements(t *testing.T) {
	require.Equal(t, Expected_VersionMajor, VersionMajor)
	require.Equal(t, Expected_VersionMinor, VersionMinor)
	require.Equal(t, Expected_VersionPatch, VersionPatch)
	require.Equal(t, Expected_VersionAB, VersionAB)
}

func Test_Version(t *testing.T) {
	require.Equal(t, uint64(0x0000_0011_0000_4002), Version)
}

func Test_Version_String(t *testing.T) {
	require.Equal(t, "0.17.0-alpha2", VersionString())
}

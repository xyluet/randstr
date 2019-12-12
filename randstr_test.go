package randstr_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xyluet/randstr"
)

func Test_String(t *testing.T) {
	s, err := randstr.String(64)
	require.NoError(t, err)
	require.Equal(t, 64, len(s))
}

func Test_StringWithCharset(t *testing.T) {
	t.Fatal(randstr.StringWithCharset(16, "ABC"))
	s, err := randstr.StringWithCharset(64, "ABC")
	require.NoError(t, err)
	require.Equal(t, 64, len(s))
	for _, x := range s {
		require.True(t, strings.ContainsRune("ABC", rune(x)))
	}
}

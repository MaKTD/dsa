package bracketschecker

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckBrackets(t *testing.T) {
	seq := ""
	require.Equal(t, true, CheckBrackets(seq))

	seq = "("
	require.Equal(t, false, CheckBrackets(seq))

	seq = "[({"
	require.Equal(t, false, CheckBrackets(seq))

	seq = "[]"
	require.Equal(t, true, CheckBrackets(seq))

	seq = "[](){}"
	require.Equal(t, true, CheckBrackets(seq))

	seq = "[({})]"
	require.Equal(t, true, CheckBrackets(seq))

	seq = "[]([])[{}{}]([[]]{})[{{}}]"
	require.Equal(t, true, CheckBrackets(seq))

	seq = "]}"
	require.Equal(t, false, CheckBrackets(seq))

	seq = "}1234["
	require.Equal(t, false, CheckBrackets(seq))

	seq = "[fsd]"
	require.Equal(t, false, CheckBrackets(seq))

	seq = "[["
	require.Equal(t, false, CheckBrackets(seq))

}

package bracketschecker

import (
	"github.com/MaKTD/go-dsa/stack"
	"golang.org/x/exp/slices"
)

var OpenBrackets = []string{"{", "[", "("}
var ClosedToOpenBrackets = map[string]string{
	"}": "{",
	"]": "[",
	")": "(",
}

func CheckBrackets(str string) bool {
	st := stack.NewStack[string]()

	for _, char := range str {
		if slices.Contains(OpenBrackets, string(char)) {
			st.Push(string(char))
		} else if st.IsEmpty() {
			return false
		} else if expectedOpenBr, ok := ClosedToOpenBrackets[string(char)]; ok {
			openBr, _ := st.Pop()
			if openBr != expectedOpenBr {
				return false
			}
		} else {
			return false
		}
	}

	return st.IsEmpty()
}

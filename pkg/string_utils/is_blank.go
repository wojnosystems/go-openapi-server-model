package string_utils

import "strings"

const whitespaceCutset = " \t\n\r"

func IsBlank(v *string) bool {
	if v == nil {
		return true
	}
	trimmed := strings.Trim(*v, whitespaceCutset)
	return len(trimmed) == 0
}

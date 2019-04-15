package tools

import s "strings"

func ValidateString(value string) bool {
	value = s.TrimSpace(value)
	return len(s.TrimSpace(value)) > 0
}

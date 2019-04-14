package utilities

import s "strings"

func validateString(value string) bool {
	s.TrimSpace(value)
	return len(s.TrimSpace(value)) > 0
}

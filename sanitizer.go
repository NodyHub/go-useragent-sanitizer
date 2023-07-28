package sanitizer

import (
	"regexp"
	"strings"
)

const (
	validCharacters = "[A-Z]|[a-z]|[0-9]|/|\\.| |;|:|\\+|\\(|\\)|;|_|,"
)

// Replace returns a copy of userAgent, all non-valid runes are replaced
// by replaceWith.
func Replace(userAgent, replaceWith string) string {
	// prepare regex
	allowList := regexp.MustCompile(validCharacters)

	// prepare result string
	var sb strings.Builder

	// iterate over input
	for _, srcrune := range userAgent {
		if !allowList.MatchString(string(srcrune)) {
			sb.WriteString(replaceWith)
		} else {
			sb.WriteString(string(srcrune))
		}
	}

	return sb.String()
}

// Returns a copy of userAgent, without all non-valid user agent runes.
func Remove(userAgent string) string {
	return Replace(userAgent, "")
}

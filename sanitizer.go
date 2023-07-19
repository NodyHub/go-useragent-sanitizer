package sanitizer

import (
	"regexp"
)

const (
	validCharacters = "[A-Z]|[a-z]|[0-9]|/|\\.| |;|:|\\+|\\(|\\)|;|_|,"
)

func Sanitize(userAgent string, replaceWith byte) string {
	var changed bool
	work := []byte(userAgent)
	allowList := regexp.MustCompile(validCharacters)
	for idx := range work {
		val := string(work[idx])
		if !allowList.MatchString(val) {
			work[idx] = byte(replaceWith)
			changed = true
		}
	}
	if changed {
		userAgent = string(work)
	}

	return userAgent
}

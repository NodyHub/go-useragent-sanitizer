package sanitizer

import (
	"regexp"
)

const (
	validCharacters = "[A-Z]|[a-z]|[0-9]|/|\\.| |;|:|\\+|\\(|\\)|;|_|,"
)

func Replace(userAgent string, replaceWith byte) (string, bool) {
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

	return userAgent, changed
}

func Remove(userAgent string) (string, bool) {
	var changed bool
	work := []byte(userAgent)
	allowList := regexp.MustCompile(validCharacters)
	for idx := len(work) - 1; idx >= 0; idx -= 1 {
		val := string(work[idx])
		if !allowList.MatchString(val) {
			work = append(work[:idx], work[idx+1:]...)
			changed = true
		}
	}
	if changed {
		userAgent = string(work)
	}

	return userAgent, changed
}

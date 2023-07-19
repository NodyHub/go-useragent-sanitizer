package sanitizer

import (
	"regexp"
)

const (
	validCharacters = "[A-Z]|[a-z]|[0-9]|/|\\.| |;|:|\\+|\\(|\\)|;|_|,"
)

// Replace replaces all non-valid character from the provided userAgent with replaceWith and returns the the updates
// userAgent and if any changes have been done.
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

// Remove removes all non-valid character from the provided userAgent and returns the the updates
// userAgent and if any changes have been done.
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

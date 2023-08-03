package sanitizer

import (
	"regexp"
)

const (
	deniedCharacters = `[^A-Za-z0-9/\. ;:\+\(|\)_,]`
)

// Replace returns a copy of userAgent, all non-valid runes are replaced
// by replaceWith.
func Replace(userAgent, replaceWith string) string {
	denyList := regexp.MustCompile(deniedCharacters)
	return denyList.ReplaceAllString(userAgent, replaceWith)
}

// Returns a copy of userAgent, without all non-valid user agent runes.
func Remove(userAgent string) string {
	return Replace(userAgent, "")
}

package sanitizer

import (
	"regexp"
	"testing"
)

func TestReplace(t *testing.T) {
	userAgent := "Mozilla/5.0 '\"<h1>"
	userAgentWant := regexp.MustCompile(`Mozilla/5.0 ___h1_`)
	sanitized := Replace(userAgent, "_")
	if !userAgentWant.MatchString(sanitized) {
		t.Fatalf(`Replace("Mozilla/5.0 '\"<h1>",'_') = %q want match for %#q`, sanitized, userAgentWant)
	}
}

func TestRemove(t *testing.T) {
	userAgent := "Mozilla/5.0 '\"<h1>"
	userAgentWant := regexp.MustCompile(`Mozilla/5.0 h1`)
	sanitized := Remove(userAgent)
	if !userAgentWant.MatchString(sanitized) {
		t.Fatalf(`Remove("Mozilla/5.0 '\"<h1>") = %q want match for %#q`, sanitized, userAgentWant)
	}
}

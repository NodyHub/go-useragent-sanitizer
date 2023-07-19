package main

import (
	"fmt"

	sanitizer "github.com/NodyHub/go-useragent-sanitizer"
)

func main() {
	// Test data
	malformedUserAgent := "Mozilla/5.0 (platform; <h1>'\"%0D%0A"
	// Before
	fmt.Printf("Malformed UserAgent: %v\n", malformedUserAgent)
	// Remove
	if sanitized, changed := sanitizer.Remove(malformedUserAgent); changed {
		fmt.Printf("Sanitized with remove: %v\n", sanitized)
	}
	// Replace
	if sanitized, changed := sanitizer.Replace(malformedUserAgent, '_'); changed {
		fmt.Printf("Sanitized with replace: %v\n", sanitized)
	}
}

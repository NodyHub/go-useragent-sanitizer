# go-useragent-sanitizer

Sanitize user agent strings, basded on current examples from [Mozilla](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/User-Agent).  Current in-use regex is: `"[A-Z]|[a-z]|[0-9]|/|\\.| |;|:|\\+|\\(|\\)|;|_|,"`.


## Example usage

Example code:

```golang
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
```

Full example can be found [here](example/main.go).
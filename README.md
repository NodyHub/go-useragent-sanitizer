# go-useragent-sanitizer

Sanitize user agent strings, basded on current examples from [Mozilla](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/User-Agent).  Current in-use regex is: `"[A-Z]|[a-z]|[0-9]|/|\\.| |;|:|\\+|\\(|\\)|;|_|,"`.


## Instalation

```shell
$ go get github.com/NodyHub/go-useragent-sanitizer
```

## Example 

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

Output:

```shell
$ go run ./main.go
Malformed UserAgent: Mozilla/5.0 (platform; <h1>'"%0D%0A
Sanitized with remove: Mozilla/5.0 (platform; h10D0A
Sanitized with replace: Mozilla/5.0 (platform; _h1____0D_0A
```

Full example can be found [here](example/main.go).
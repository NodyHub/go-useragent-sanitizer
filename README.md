# go-useragent-sanitizer

[![Go](https://github.com/NodyHub/go-useragent-sanitizer/actions/workflows/test.yaml/badge.svg)](https://github.com/NodyHub/go-useragent-sanitizer/actions/workflows/test.yaml)

Sanitize user agent strings, basded on current examples from [Mozilla](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/User-Agent).  Current in-use regex is: `[^A-Za-z0-9/\. ;:\+\(|\)_,]`.

## Instalation

```shell
go get github.com/NodyHub/go-useragent-sanitizer
```

## Example

Example code:

```golang
 // Test data
 malformedUserAgent := "Mozilla/5.0 (platform; <h1>'\"%0D%0A"
 // Before
 fmt.Printf("Malformed UserAgent: %v\n", malformedUserAgent)
 // Remove
 fmt.Printf("Sanitized with remove: %v\n", sanitizer.Remove(malformedUserAgent))
 // Replace
 fmt.Printf("Sanitized with replace: %v\n", sanitizer.Replace(malformedUserAgent, '_'))
```

Output:

```shell
Malformed UserAgent: Mozilla/5.0 (platform; <h1>'"%0D%0A
Sanitized with remove: Mozilla/5.0 (platform; h10D0A
Sanitized with replace: Mozilla/5.0 (platform; _h1____0D_0A
```

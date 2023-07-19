# go-useragent-sanitizer
Allow listing valid common HTTP user agents and replace invalid characters.


## Example usage

Code:

```golang
package main

import (
	"fmt"

	sanitizer "github.com/NodyHub/go-useragent-sanitizer"
)

func main() {

	userAgents := []string{
		"Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:42.0) Gecko/20100101 Firefox/42.0",
		"Mozilla/5.0 (X11; Linux \"x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safari/537.36 OPR/38.0.2220.41",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safar'i/537.36 OPR/38.0.2220.41",
		"Opera/9.80 (Macintosh; Intel Mac OS X; U; en) Presto/2.2.15 Version/10.00",
		"Opera/9.60 (Windows NT 6.0; U; en) Presto/2.1.1",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.59",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.1 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		"Mozilla/5.0 (compatible; YandexAccessibilityBot/3.0; +http://yandex.com/b<ots)",
		"Mozilla/5.0 (compatible; YandexAccessibilityBot/3.0; +http://yandex.com/bots)",
	}

	fmt.Println("> Remove example")
	for _, userAgent := range userAgents {
		fmt.Printf("UserAgent:\t%v\n", userAgent)
		after, changed := sanitizer.Remove(userAgent)
		if changed {
			fmt.Printf("Sanitized:\t%v\n", after)
		}
	}

	fmt.Println()
	fmt.Println("> Replace example")
	for _, userAgent := range userAgents {
		fmt.Printf("UserAgent:\t%v\n", userAgent)
		after, changed := sanitizer.Replace(userAgent, '_')
		if changed {
			fmt.Printf("Sanitized:\t%v\n", after)
		}
	}
}

```

Run:

```shell
$  go run ./main.go
> Remove example
UserAgent:	Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion
UserAgent:	Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0
UserAgent:	Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:42.0) Gecko/20100101 Firefox/42.0
UserAgent:	Mozilla/5.0 (X11; Linux "x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36
Sanitized:	Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36
UserAgent:	Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36
UserAgent:	Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safari/537.36 OPR/38.0.2220.41
UserAgent:	Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safar'i/537.36 OPR/38.0.2220.41
Sanitized:	Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safari/537.36 OPR/38.0.2220.41
UserAgent:	Opera/9.80 (Macintosh; Intel Mac OS X; U; en) Presto/2.2.15 Version/10.00
UserAgent:	Opera/9.60 (Windows NT 6.0; U; en) Presto/2.1.1
UserAgent:	Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.59
UserAgent:	Mozilla/5.0 (iPhone; CPU iPhone OS 13_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.1 Mobile/15E148 Safari/604.1
UserAgent:	Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)
UserAgent:	Mozilla/5.0 (compatible; YandexAccessibilityBot/3.0; +http://yandex.com/b<ots)
Sanitized:	Mozilla/5.0 (compatible; YandexAccessibilityBot/3.0; +http://yandex.com/bots)
UserAgent:	Mozilla/5.0 (compatible; YandexAccessibilityBot/3.0; +http://yandex.com/bots)

> Replace example
UserAgent:	Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion
UserAgent:	Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0
UserAgent:	Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:42.0) Gecko/20100101 Firefox/42.0
UserAgent:	Mozilla/5.0 (X11; Linux "x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36
Sanitized:	Mozilla/5.0 (X11; Linux _x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36
UserAgent:	Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36
UserAgent:	Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safari/537.36 OPR/38.0.2220.41
UserAgent:	Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safar'i/537.36 OPR/38.0.2220.41
Sanitized:	Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safar_i/537.36 OPR/38.0.2220.41
UserAgent:	Opera/9.80 (Macintosh; Intel Mac OS X; U; en) Presto/2.2.15 Version/10.00
UserAgent:	Opera/9.60 (Windows NT 6.0; U; en) Presto/2.1.1
UserAgent:	Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.59
UserAgent:	Mozilla/5.0 (iPhone; CPU iPhone OS 13_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.1 Mobile/15E148 Safari/604.1
UserAgent:	Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)
UserAgent:	Mozilla/5.0 (compatible; YandexAccessibilityBot/3.0; +http://yandex.com/b<ots)
Sanitized:	Mozilla/5.0 (compatible; YandexAccessibilityBot/3.0; +http://yandex.com/b_ots)
UserAgent:	Mozilla/5.0 (compatible; YandexAccessibilityBot/3.0; +http://yandex.com/bots)
```
package sanitizer

import (
	"fmt"
	"testing"
)

func TestReplace(t *testing.T) {
	// prepare cases
	cases := []struct {
		input      string
		replacedBy string
		exp        string
	}{
		{
			input:      "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion",
			replacedBy: "_",
			exp:        "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion",
		},
		{
			input:      "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safari/537.36 OPR/38.0.2220.41",
			replacedBy: "_",
			exp:        "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safari/537.36 OPR/38.0.2220.41",
		},
		{
			input:      "Opera/9.80 (Macintosh; Intel Mac OS X; U; en) Presto/2.2.15 Version/10.00",
			replacedBy: "_",
			exp:        "Opera/9.80 (Macintosh; Intel Mac OS X; U; en) Presto/2.2.15 Version/10.00",
		},
		{
			input:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.59",
			replacedBy: "_",
			exp:        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.59",
		},

		{
			input:      "日本語",
			replacedBy: "_",
			exp:        "___",
		},
		{
			input:      "Mozilla/5.0 '\"<h1>",
			replacedBy: "_",
			exp:        "Mozilla/5.0 ___h1_",
		},
		{
			input:      "Mozi本lla/5.0 '\"<h1>",
			replacedBy: "_",
			exp:        "Mozi_lla/5.0 ___h1_",
		},
	}

	// run cases
	for i, tc := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			want := tc.exp
			got := Replace(tc.input, tc.replacedBy)
			if got != want {
				t.Errorf("expected %q with non-matchers "+
					"replaced by %s "+
					"to be %s, "+
					"got %s",
					tc.input, tc.replacedBy, want, got)
			}
		})
	}

}

func TestRemove(t *testing.T) {
	// prepare cases
	cases := []struct {
		input string
		exp   string
	}{
		{
			input: "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion",
			exp:   "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion",
		},
		{
			input: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safari/537.36 OPR/38.0.2220.41",
			exp:   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safari/537.36 OPR/38.0.2220.41",
		},
		{
			input: "Opera/9.80 (Macintosh; Intel Mac OS X; U; en) Presto/2.2.15 Version/10.00",
			exp:   "Opera/9.80 (Macintosh; Intel Mac OS X; U; en) Presto/2.2.15 Version/10.00",
		},
		{
			input: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.59",
			exp:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.59",
		},

		{
			input: "日本語",
			exp:   "",
		},
		{
			input: "Mozilla/5.0 '\"<h1>",
			exp:   "Mozilla/5.0 h1",
		},
		{
			input: "Mozi本lla/5.0 '\"<h1>",
			exp:   "Mozilla/5.0 h1",
		},
	}

	// run cases
	for i, tc := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			want := tc.exp
			got := Remove(tc.input)
			if got != want {
				t.Errorf("expected %q with non-matchers "+
					"to be %s, "+
					"got %s",
					tc.input, want, got)
			}
		})
	}
}

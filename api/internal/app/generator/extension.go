package generator

import (
	"regexp"

	"github.com/kostaspt/domane-server/internal/app/parser"
)

func CommonExtensions(text string) Results {
	extensions := TopExtensions()

	text = parser.Clean(text)

	results := make(Results, len(extensions))
	for i, e := range extensions {
		results[i] = Result{
			Domain: text + "." + e,
			Kind:   CommonExtension,
		}

		if e == "com" {
			results[i].Kind = DotComExtension
		}
	}

	return results
}

func CommonExtensionsHyphenated(text string) Results {
	extensions := TopExtensions()

	text = parser.Hyphenate(text)

	results := make(Results, len(extensions))
	for i, e := range extensions {
		results[i] = Result{
			Domain: text + "." + e,
			Kind:   CommonExtensionHyphenated,
		}

		if e == "com" {
			results[i].Kind = DotComExtension
		}
	}

	return results
}

func ShortExtensions(text string) Results {
	extensions := AllExtensions()

	text = parser.Clean(text)

	results := make(Results, 0, len(extensions))
	for _, e := range extensions {
		re := regexp.MustCompile(`(.{2})` + e + `.*?$`)

		if !re.MatchString(text) {
			continue
		}

		r := Result{
			Domain: re.ReplaceAllString(text, `$1.`+e),
			Kind:   ShortExtension,
		}

		if e == "com" {
			r.Kind = DotComExtension
		}

		results = append(results, r)
	}

	return results
}

func AllExtensions() []string {
	allExtensions := allExtensions()

	es := make([]string, 0, len(allExtensions))
	for _, tx := range allExtensions {
		es = append(es, tx)
	}

	return es
}

func TopExtensions() [31]string {
	return [31]string{
		"com",
		"xyz",
		"co",
		"app",
		"io",
		"org",
		"biz",
		"blog",
		"business",
		"church",
		"club",
		"design",
		"global",
		"guru",
		"info",
		"life",
		"me",
		"media",
		"net",
		"news",
		"online",
		"photography",
		"rocks",
		"science",
		"site",
		"solutions",
		"space",
		"tech",
		"today",
		"website",
		"world",
	}
}

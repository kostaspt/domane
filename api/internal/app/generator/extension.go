package generator

import (
	_ "embed"
	"encoding/json"
	"regexp"

	"github.com/rs/zerolog/log"

	"github.com/kostaspt/domane/api/pkg/parser"
)

//go:embed tlds.json
var tlds string

func CommonExtensions(text string) Results {
	extensions := TopExtensions()

	text = parser.Clean(text)

	results := make(Results, len(extensions))
	for i, e := range extensions {
		results[i] = Result{
			Domain: text + "." + e,
			Kind:   ExtensionCommon,
		}

		if e == "com" {
			results[i].Kind = ExtensionDotCom
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
			Kind:   ExtensionCommonHyphenated,
		}

		if e == "com" {
			results[i].Kind = ExtensionDotCom
		}
	}

	return results
}

func ShortExtensions(text string) Results {
	extensions := allExtensions()

	text = parser.Clean(text)

	results := make(Results, 0, len(extensions))
	for _, e := range extensions {
		re := regexp.MustCompile(`(.{2})` + e + `.*?$`)

		if !re.MatchString(text) {
			continue
		}

		r := Result{
			Domain: re.ReplaceAllString(text, `$1.`+e),
			Kind:   ExtensionShort,
		}

		if e == "com" {
			r.Kind = ExtensionDotCom
		}

		results = append(results, r)
	}

	return results
}

func allExtensions() []string {
	var extensions map[string]string

	err := json.Unmarshal([]byte(tlds), &extensions)
	if err != nil {
		log.Err(err).Send()
		return []string{}
	}

	es := make([]string, 0, len(extensions))
	for _, tx := range extensions {
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

package generator

import (
	"github.com/kostaspt/go-datamuse/v2"
	"github.com/rs/zerolog/log"
	"github.com/thoas/go-funk"

	"github.com/kostaspt/domane/api/pkg/parser"
)

func Synonyms(text string) Results {
	words := parser.Words(text)

	synonymWords := make([][]string, 0, len(words))

	for _, w := range words {
		dmRes, err := datamuse.New().Words().MeansLike(w).Get()
		if err != nil {
			log.Err(err).Send()
		}

		// Add the actual word, so we can match this with other words
		synonyms := []string{w}

		for _, r := range dmRes {
			if !funk.Contains(r.Tags, datamuse.RelatedSynonyms) {
				continue
			}
			synonyms = append(synonyms, r.Word)
		}

		synonymWords = append(synonymWords, synonyms)
	}

	results := make(Results, 0, len(synonymWords))
	for _, synonym := range getPermutation(synonymWords) {
		name := ""
		for _, s := range synonym {
			name += parser.Clean(s)
		}

		d := name + ".com"

		// We do not want to end up with a domain we already generated.
		if d == parser.Clean(text)+".com" {
			continue
		}

		results = append(results, Result{
			Domain: d,
			Kind:   SimilarSynonym,
		})
	}

	return results
}

func getPermutation(words [][]string) [][]string {
	var res [][]string

	if len(words) == 0 {
		return res
	}

	if len(words) == 1 {
		for _, vid := range words[0] {
			res = append(res, []string{vid})
		}
		return res
	}

	t := getPermutation(words[1:])
	for _, vid := range words[0] {
		for _, perm := range t {
			p := append([]string{vid}, perm...)
			res = append(res, p)
		}
	}

	return res
}

package generator

import (
	"testing"

	"github.com/kostaspt/go-datamuse/v2"
	"github.com/stretchr/testify/assert"
	"github.com/thoas/go-funk"
)

func TestSynonyms(t *testing.T) {
	sHello, _ := datamuse.New().Words().MeansLike("hello").Get()
	sHelloSyn := funk.Filter(sHello, func(s datamuse.Result) bool {
		return funk.Contains(s.Tags, datamuse.RelatedSynonyms)
	}).([]datamuse.Result)

	sWorld, _ := datamuse.New().Words().MeansLike("world").Get()
	sWorldSyn := funk.Filter(sWorld, func(s datamuse.Result) bool {
		return funk.Contains(s.Tags, datamuse.RelatedSynonyms)
	}).([]datamuse.Result)

	synonyms := Synonyms("hello world")

	// The +1 is needed here because we append the actual word to synonyms
	// The -1 is needed because we remove the helloworld.com as a valid result
	assert.Equal(t, (len(sHelloSyn)+1)*(len(sWorldSyn)+1)-1, len(synonyms))

	// First result should be the hello + a synonym of world
	assert.Equal(t, "hello" + sWorld[0].Word + ".com", synonyms[0].Domain)
}

func TestGetPermutations(t *testing.T) {
	words := [][]string{
		{"a", "b"},
		{"c", "d"},
	}

	res := getPermutation(words)
	assert.Len(t, res, 4)
	assert.Equal(t, []string{"a", "c"}, res[0])
	assert.Equal(t, []string{"a", "d"}, res[1])
	assert.Equal(t, []string{"b", "c"}, res[2])
	assert.Equal(t, []string{"b", "d"}, res[3])
}
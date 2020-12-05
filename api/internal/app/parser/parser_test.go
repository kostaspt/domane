package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const dirtyText = `. ! Foo   "Bar" -$`

func TestClean(t *testing.T) {
	assert.Equal(t, "foobar", Clean(dirtyText))
}

func TestPrune(t *testing.T) {
	assert.Equal(t, "Foo Bar", Prune(dirtyText))
}

func TestPruneAndLowercase(t *testing.T) {
	assert.Equal(t, "foo bar", PruneAndLowercase(dirtyText))
}

func TestHyphenate(t *testing.T) {
	assert.Equal(t, "foo-bar", Hyphenate(dirtyText))
}

func TestWords(t *testing.T) {
	assert.Equal(t, []string{"foo", "bar"}, Words(dirtyText))
}

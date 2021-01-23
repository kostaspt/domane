package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonExtensions(t *testing.T) {
	results := CommonExtensions("foo bar")

	assert.Equal(t, len(TopExtensions()), len(results))

	assert.Equal(t, "foobar.com", results[0].Domain)
	assert.Equal(t, ExtensionDotCom, results[0].Kind)

	assert.Equal(t, "foobar.xyz", results[1].Domain)
	assert.Equal(t, ExtensionCommon, results[1].Kind)
}

func TestCommonExtensionsHyphenated(t *testing.T) {
	results := CommonExtensionsHyphenated("foo bar")

	assert.Equal(t, len(TopExtensions()), len(results))

	assert.Equal(t, "foo-bar.com", results[0].Domain)
	assert.Equal(t, ExtensionDotCom, results[0].Kind)

	assert.Equal(t, "foo-bar.xyz", results[1].Domain)
	assert.Equal(t, ExtensionCommonHyphenated, results[1].Kind)
}

func TestShortExtensions(t *testing.T) {
	results := ShortExtensions("foo bar")
	assert.GreaterOrEqual(t, len(results), 1)

	for _, r := range results {
		if r.Domain == "foo.bar" && r.Kind == ExtensionShort {
			return
		}
	}

	t.Fail()
}
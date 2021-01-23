package generator

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResultsSort(t *testing.T) {
	rs := make(Results, 2)
	rs = Results{
		{Domain: "foo.xyz", Kind: ExtensionCommon},
		{Domain: "foo.com", Kind: ExtensionDotCom},
		{Domain: "foo.io", Kind: ExtensionCommon},
		{Domain: "f.io", Kind: ExtensionCommon},
	}

	sort.Sort(rs)

	assert.Equal(t, "foo.com", rs[0].Domain)
	assert.Equal(t, "f.io", rs[1].Domain)
	assert.Equal(t, "foo.io", rs[2].Domain)
	assert.Equal(t, "foo.xyz", rs[3].Domain)
}

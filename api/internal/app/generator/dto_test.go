package generator

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResultsSort(t *testing.T) {
	rs := make(Results, 2)
	rs = Results{
		{Domain: "foo.xyz", Kind: CommonExtension},
		{Domain: "foo.com", Kind: DotComExtension},
		{Domain: "foo.io", Kind: CommonExtension},
		{Domain: "f.io", Kind: CommonExtension},
	}

	sort.Sort(rs)

	assert.Equal(t, "foo.com", rs[0].Domain)
	assert.Equal(t, "f.io", rs[1].Domain)
	assert.Equal(t, "foo.io", rs[2].Domain)
	assert.Equal(t, "foo.xyz", rs[3].Domain)
}

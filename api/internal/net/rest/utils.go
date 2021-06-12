package rest

import (
	"sort"

	"github.com/kostaspt/domane/api/internal/app/generator"
)

func sortResults(res generator.Results) generator.Results {
	sort.Sort(res)

	// Add position to make sure it will be kept at the front-end
	for i := range res {
		res[i].Position = uint(i + 1)
	}

	return res
}

// Sort map by key

package datastructure

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortMapByKey(t *testing.T) {
	basket := map[string]int{
		"orange":     5,
		"apple":      7,
		"mango":      3,
		"strawberry": 9,
	}
	keys := make([]string, 0, len(basket))
	for k := range basket {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, basket[k])
	}
}

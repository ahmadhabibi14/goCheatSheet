package datastructure

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortMapByValue(t *testing.T) {
	basket := map[string]int{
		"orange":     5,
		"apple":      7,
		"mango":      3,
		"strawberry": 9,
	}
	keys := make([]string, 0, len(basket))
	for key := range basket {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return basket[keys[i]] < basket[keys[j]]
	})

	for _, k := range keys {
		fmt.Println(k, basket[k])
	}
}

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package treesort_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"gopl.io/ch4/treesort"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	fmt.Println(rand.Intn(10000))

	for i := range data {
		data[i] = rand.Intn(10000) % 50
	}
	fmt.Println(data[:0])

	fmt.Println(data)
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
	fmt.Println(data)
}

package seq

import (
	"fmt"
	"testing"
)

func TestConcat(t *testing.T) {
	origin := []int{1, 2, 3}
	list := []int{4, 5, 6}

	result := Concat(origin, list)
	fmt.Println(result) // Output: [1 2 3 4 5 6]
}

func TestSublist(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7}
	n := 3

	result, err := SubList(list, n)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result) // Output: [[1 2 3] [4 5 6] [7]]
}

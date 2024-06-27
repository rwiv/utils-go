package iter

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	//var ret []int
	//for _, elem := range []int{1, 2, 3, 4, 5} {
	//	if elem%2 == 0 {
	//		ret = append(ret, elem)
	//	}
	//}
	ret := Filter([]int{1, 2, 3, 4, 5}, func(elem int) bool {
		return elem%2 == 0
	})
	fmt.Println(ret)
}

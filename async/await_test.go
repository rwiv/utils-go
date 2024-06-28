package async

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestThread(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := AwaitAll(list, logic)
	Sort(result)
	fmt.Println(FilterSuccess(result))
}

func logic(i int) (*int, error) {
	if i == 6 {
		return nil, errors.New("error")
	}
	time.Sleep(time.Second)
	result := i * 2
	return &result, nil
}

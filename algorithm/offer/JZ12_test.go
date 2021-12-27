package offer

import (
	"fmt"
	"testing"
)

func Test_hasPath(t *testing.T) {
	fmt.Println(hasPath([][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}, "abcd"))
}
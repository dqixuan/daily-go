package double_index

import (
	"fmt"
	"testing"
)

func Test_compress(t *testing.T) {
	fmt.Println(compress([]byte{'a','a','b','b','c','c','c'}))
}
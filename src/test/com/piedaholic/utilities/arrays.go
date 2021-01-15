package testarrays

import (
	"fmt"
	"testing"
)

func testReverseArray(b *testing.B) {
	arr := [3]int{1, 2, 3}
	rev_arr := utilities.reverse_array(arr)
	fmt.Println(rev_arr)
}

package slices

import (
	"fmt"
	"testing"
)

func TestSliceMaker(t *testing.T) {

	//  Slice 의 len 과 cap을 이런식으로 생각하면 더 쉬울 것 같습니다.

	a := make([]int, 5) // [0, 0, 0, 0, 0]
	PrintSlice("a", a)

	b := make([]int, 0, 5) // [x, x, x, x, x] == []
	PrintSlice("b", b)

	c := b[:2] // [0, 0, x, x, x] == [0, 0]
	PrintSlice("c", c)

	d := c[2:5] // [x, x, 0, 0, 0] == [0, 0, 0]
	PrintSlice("d", d)

	g := c[2:3] // [x, x, 0, x, x] == [0]
	PrintSlice("g", g)

}

func PrintSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

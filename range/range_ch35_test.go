package rrange

import (
	"fmt"
	"testing"
)

func TestRangeTwo(t *testing.T) {

	pow := make([]int, 10)

	for k := range pow {
		pow[k] = 1 << uint(k)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	fmt.Println("s")
}

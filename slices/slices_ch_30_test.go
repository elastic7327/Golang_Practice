package slices

import (
	"fmt"
	"testing"
)

func TestFunction(t *testing.T) {

	p := []int{2, 3, 5, 7, 11, 13}

	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])

	}
}

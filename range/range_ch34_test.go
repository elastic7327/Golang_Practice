package rrange

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for k, v := range pow {

		fmt.Printf("2**%d = %d \n", k, v)

	}

}

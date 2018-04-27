package pointers

import (
	"fmt"
	"testing"
)

func TestNewFunc(t *testing.T) {

	v := new(Vertex)

	fmt.Println(v)

	v.X, v.Y = 11, 9

	fmt.Println("v")

}

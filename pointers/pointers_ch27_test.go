package pointers

import (
	"fmt"
	"testing"
)

type Vertex struct {
	X int
	Y int
}

func TestFunction(t *testing.T) {

	p := Vertex{1, 2}
	q := &p
	q.X = 1e9
	fmt.Println(p)
	q.X++

}

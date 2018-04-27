package pointers

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {

	var (
		p = Vertex{1, 2}
		q = &Vertex{1, 2}
		r = Vertex{X: 1}
		s = Vertex{}
	)

	fmt.Println(p, q, r, s)

}

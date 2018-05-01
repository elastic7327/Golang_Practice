package method

import (
	"fmt"
	"math"
	"testing"
)

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) AAbs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func TestMethodPointerRecevier(t *testing.T) {

	v := &Vertex{3, 4}

	v.Scale(5)

	fmt.Println(v, v.AAbs())

}

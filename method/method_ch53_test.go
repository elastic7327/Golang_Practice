package method

import (
	"fmt"
	"math"
	"testing"
)

type Abser interface {
	/* TODO: add methods */
	Abs() float64
}

type MMyFloat float64

func (f MMyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type VVertex struct {
	X, Y float64
}

func (v *VVertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func TestMethodAndInterface(t *testing.T) {

	/*
		인터페이스는 메소드의 집합으로 정의됩니다.

		그 메소드들의 구현되어 있는 타입의 값은 모두 인터페이스 타입의 값이 될 수 있습니다.
	*/

	var a Abser

	f := MMyFloat(-math.Sqrt2)

	v := VVertex{3, 4}

	a = f
	a = &v

	// a = v

	fmt.Println(a.Abs())

	g := Abser(f)

	fmt.Println(g.Abs())

}

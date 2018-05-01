package method

import (
	"fmt"
	"math"
	"testing"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func TestMethodWorld(t *testing.T) {
	/*
		고에는 클래스가 없습니다. 하지만 메소드를 구조체(struct)에 붙일 수 있습니다.

		메소드 리시버(method receiver) 는 func 키워드와 메소드의 이름 사이에 인자로 들어갑니다.
	*/

	v := &Vertex{3, 4}

	fmt.Println(v.Abs())

}

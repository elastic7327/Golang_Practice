package function

import (
	"fmt"
	"math"
	"testing"
)

func TestFunctionValues(t *testing.T) {

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
}

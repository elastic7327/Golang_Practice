package function

import (
	"fmt"
	"testing"
)

func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

func TestFunctionClosure(t *testing.T) {
	/*

		그리고 함수는 클로져(full closures) 입니다.

		코드에서 adder 함수는 클로져(closure)를 반환합니다.

		각각의 클로져는 자신만의 sum 변수를 가집니다.

	*/

	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

}

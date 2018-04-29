package function

import (
	"fmt"
	"testing"
)

func fib() func() int {

	a, b := 0, 1

	return func() int {

		a, b = b, a+b

		return a
	}
}

func TestFibonacci(t *testing.T) {
	/*
		함수를 가지고 놀아봅시다.

		fibonacci 함수를 구현합니다. 이 함수는 이어지는 피보나치 수를 반환하는 함수 (클로져)를 반환해야 합니다.
	*/
	f := fib()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

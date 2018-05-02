package channels

import (
	"fmt"
	"testing"
)

func sum(a []int, c chan int) {

	sum := 0
	for _, v := range a {
		sum += v
	}

	c <- sum // send sum to c
}

func TestChannels(t *testing.T) {
	/*
		채널은 채널 연산자 <- 를 이용해 값을 주고 받을 수 있는, 타입이 존재하는 파이프입니다.

		ch <- v    // v 를 ch로 보냅니다.
		v := <-ch  // ch로부터 값을 받아서
		           // v 로 넘깁니다.
				   (데이터가 화살표 방향에 따라 흐릅니다.)
		   맵이나 슬라이스처럼, 채널은 사용되기 전에 생성되어야 합니다:

	*/

	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c // refeice from c

	fmt.Println(x, y, x+y)

}

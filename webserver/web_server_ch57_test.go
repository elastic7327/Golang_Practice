package webserver

import (
	"fmt"
	"net/http"
	"testing"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func TestWebServer(t *testing.T) {

	t.Skip()

	/*

		Package http 는 http.Handler 를 구현한 어떠 값을 사용하여 HTTP 요청(requests)을 제공합니다.

		package http

		type Handler interface {
				ServeHTTP(w ResponseWriter, r *Request)

		}
		이 예제에서, Hello 라는 타입은 http.Handler 를 구현합니다.

		이 코드를 로컬에서 실행하고, http://localhost:4000/ 에 접속해보세요.

	*/

	var h Hello
	http.ListenAndServe("localhost:4000", h)

}

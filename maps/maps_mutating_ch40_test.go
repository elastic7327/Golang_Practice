package maps

import (
	"fmt"
	"strings"
	"testing"
)

func TestMapsMutating(t *testing.T) {
	m := make(map[string]int)

	m["Answer"] = 42

	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 32

	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")

	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]

	fmt.Println("the value:", v, "and ok", ok)

	m["Answer"] = 772

	v, ok = m["Answer"]

	fmt.Println("the value:", v, "and ok", ok)

}

func WordCount(s string) map[string]int {

	m := make(map[string]int)

	for _, v := range strings.Fields(s) {
		fmt.Println(v)
		m[v]++
	}

	return m

}

func TestMapPractice(t *testing.T) {

	gooString := "Hello World My Name is Daniel What is your Name?"

	fmt.Println(WordCount(gooString))

}

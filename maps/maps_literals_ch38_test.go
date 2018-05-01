package maps

import (
	"fmt"
	"testing"
)

func TestMapsLiterals(t *testing.T) {
	var m = map[string]Vertex{
		"Bell Labs": {
			3.14, 3.141592,
		}, "Google labs": {
			-99, -99999,
		},
	}

	var x = map[string]Vertex{
		"Gmail": {40.1234, 40.121},
		"Nmail": {410.1234, 340.121},
	}

	fmt.Println(m)

	fmt.Println(x)

}

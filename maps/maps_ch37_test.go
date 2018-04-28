package maps

import (
	"fmt"
	"testing"
)

type Vertex struct {
	Lat, Long float64
}

func TestMaps(t *testing.T) {

	var m map[string]Vertex

	var stranger map[string]int

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	stranger = make(map[string]int)
	stranger["Hello World"] = 8

	fmt.Println("{} {}", m, stranger)

}

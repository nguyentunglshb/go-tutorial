package map_package

import "fmt"

type Vertex struct {
	Lat, Long float64
}


func Map() {

	m:=map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)

}
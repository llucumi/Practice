package main

import "fmt"

const (
	rectType   = "RECT"
	circleType = "CIRCLE"
)

func newGeometry(geoType string, values ...float64) geometry {
	switch geoType {
	case rectType:
		return rect{width: values[0], height: values[1]}
	case circleType:
		return circle{radius: values[0]}
	}
	return nil
}

func main() {
	r := newGeometry(rectType, 2, 3)
	fmt.Println(r.area())
	fmt.Println(r.perim())
	c := newGeometry(circleType, 2)
	fmt.Println(c.area())
	fmt.Println(c.perim())
}

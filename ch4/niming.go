package main

import "fmt"

type Point struct {
	X, Y int
}

//type Circle struct {
//	Center Point
//	Radius int
//}
//
//type Wheel struct {
//	Circle Circle
//	Spokes int
//}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	//var w Wheel
	//w.Circle.Center.X = 8
	//w.Circle.Center.Y = 8
	//w.Circle.Radius = 5
	//w.Spokes = 20
	//
	//var w Wheel
	//w.X = 8      // equivalent to w.Circle.Point.X = 8
	//w.Y = 8      // equivalent to w.Circle.Point.Y = 8
	//w.Radius = 5 // equivalent to w.Circle.Radius = 5
	//w.Spokes = 20

	//w = Wheel{8, 8, 5, 20}                       // compile error: unknown fields
	//w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields

	w := Wheel{
		Circle: Circle{
			Point: Point{
				X: 1,
				Y: 2,
			},
			Radius: 5,
		},
		Spokes: 10,
	}
	fmt.Printf("%#v\n", w)
}

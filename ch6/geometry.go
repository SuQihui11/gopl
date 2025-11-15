package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i].Distance(path[i-1])
		}
	}
	return sum
}

//type ColorPoint struct {
//	Point Point
//	Color color.RGBA
//}

//func (p ColorPoint) Distance() float64 {
//	return p.Point.Distance(p.Point)
//}

type ColoredPoint struct {
	*Point
	Color color.RGBA
}

func main() {
	//var cp ColoredPoint
	//cp.Point.X = 1 // 会 panic：因为 cp.Point == nil，访问 cp.Point.X 会 dereference nil 指针。
	//fmt.Println(cp.Point.X, cp.Point.Y)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{&Point{1, 1}, red}
	var q = ColoredPoint{&Point{5, 4}, blue}

	fmt.Println(p.Distance(*q.Point))
}

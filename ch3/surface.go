// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func mainTest() {

	fp, err := os.OpenFile("output.svg", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "File open failed: ", err)
		return
	}
	fmt.Fprintf(fp, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	minZ, maxZ := math.Inf(-1), math.Inf(1)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)
			z := f(x, y)
			if !math.IsNaN(z) && !math.IsInf(z, 0) {
				if z < minZ {
					minZ = z
				}
				if z > maxZ {
					maxZ = z
				}
			}
		}
	}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z1, ok1 := corner(i+1, j)
			bx, by, z2, ok2 := corner(i, j)
			cx, cy, z3, ok3 := corner(i, j+1)
			dx, dy, z4, ok4 := corner(i+1, j+1)
			//只有当所有角点都有效时才绘制多边形
			if !ok1 || !ok2 || !ok3 || !ok4 {
				continue
			}
			//计算多边形的平均高度
			avgZ := (z1 + z2 + z3 + z4) / 4
			//根据高度计算颜色
			color := getColor(avgZ, maxZ, minZ)
			fmt.Fprintf(fp, "<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
				"style='fill:%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintln(fp, "</svg>")
}

func corner(i, j int) (float64, float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, 9, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	if math.IsNaN(sx) || math.IsNaN(sy) || math.IsInf(sx, 0) || math.IsInf(sy, 0) {
		return 0, 0, 0, false
	}
	return sx, sy, z, true
}

func f(x, y float64) float64 {
	// 5. 复杂波纹
	r := math.Hypot(x, y)
	return math.Cos(r) / (r/10 + 1) * 0.5
}
func getColor(z, minZ, maxZ float64) string {
	normalized := (z - minZ) / (maxZ - minZ)
	red := int(normalized * 255)
	blue := int((1 - normalized) * 255)
	return fmt.Sprintf("#%02x00%02x", red, blue)
}

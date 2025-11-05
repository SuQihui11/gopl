package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler1(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	// 获取参数
	width := getIntParam(r, "width", 600)
	height := getIntParam(r, "height", 320)
	funcType := r.URL.Query().Get("func")
	if funcType == "" {
		funcType = "sin"
	}

	// 生成 SVG
	generateSVG(w, width, height, funcType)
}

func getIntParam(r *http.Request, param string, defaultValue int) int {
	str := r.URL.Query().Get(param)
	if str == "" {
		return defaultValue
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}
	return val
}

func generateSVG(w http.ResponseWriter, width, height int, funcType string) {
	xyscale := float64(width) / 2 / xyrange
	zscale := float64(height) * 0.4

	// SVG 开始标签（只输出一次）
	_, err := fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	if err != nil {
	}

	// 找到 z 的范围用于着色
	minZ, maxZ := findZRange(funcType)

	// 绘制多边形
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z1, ok1 := cornerHttp(i+1, j, width, height, xyscale, zscale, funcType)
			bx, by, z2, ok2 := cornerHttp(i, j, width, height, xyscale, zscale, funcType)
			cx, cy, z3, ok3 := cornerHttp(i, j+1, width, height, xyscale, zscale, funcType)
			dx, dy, z4, ok4 := cornerHttp(i+1, j+1, width, height, xyscale, zscale, funcType)

			if !ok1 || !ok2 || !ok3 || !ok4 {
				continue
			}

			avgZ := (z1 + z2 + z3 + z4) / 4
			color := getColor(avgZ, minZ, maxZ)

			_, _ = fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
				"style='fill:%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}

	// SVG 结束标签（只输出一次）
	fmt.Fprintf(w, "</svg>\n")
}

func cornerHttp(i, j, width, height int, xyscale, zscale float64, funcType string) (sx, sy, z float64, ok bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z = getFunction(funcType)(x, y)

	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, 0, false
	}

	sx = float64(width)/2 + (x-y)*cos30*xyscale
	sy = float64(height)/2 + (x+y)*sin30*xyscale - z*zscale

	if math.IsInf(sx, 0) || math.IsNaN(sx) ||
		math.IsInf(sy, 0) || math.IsNaN(sy) {
		return 0, 0, 0, false
	}

	return sx, sy, z, true
}

func findZRange(funcType string) (minZ, maxZ float64) {
	minZ, maxZ = math.Inf(1), math.Inf(-1)
	f := getFunction(funcType)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)
			z := f(x, y)
			if !math.IsInf(z, 0) && !math.IsNaN(z) {
				if z < minZ {
					minZ = z
				}
				if z > maxZ {
					maxZ = z
				}
			}
		}
	}
	return minZ, maxZ
}

func getFunction(funcType string) func(float64, float64) float64 {
	switch funcType {
	case "eggbox":
		return eggBox
	case "moguls":
		return moguls
	case "saddle":
		return saddle
	case "waves":
		return waves
	default:
		return sinOverR
	}
}

func sinOverR(x, y float64) float64 {
	r := math.Hypot(x, y)
	if r == 0 {
		return 1
	}
	return math.Sin(r) / r
}

func eggBox(x, y float64) float64 {
	return (math.Sin(x) + math.Sin(y)) * 0.1
}

func moguls(x, y float64) float64 {
	return math.Sin(x) * math.Cos(y) * 0.2
}

func saddle(x, y float64) float64 {
	return (x*x - y*y) / 500
}

func waves(x, y float64) float64 {
	return math.Sin(x/2) * math.Cos(y/2) * 0.3
}

package eval

import (
	"fmt"
	"io"
	"math"
	"net/http"
)

func parseAndCheck(s string) (Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}

	expr, err := Parse(s)
	if err != nil {
		return nil, fmt.Errorf("error parsing expression '%s': %s", s, err)
	}

	vars := make(map[Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}

	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("unknown variable '%s'", v)
		}
	}
	return expr, nil
}

func Polt(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, func(x, y float64) float64 {
		// 自动计算 r (当前点到原点的距离)
		r := math.Hypot(x, y)
		// 构造环境 Env，把当前的 x, y, r 的值传进去
		// 然后调用 Eval 计算最终的高度值
		return expr.Eval(Env{"x": x, "y": y, "r": r})
	})
}

// 定义画布和绘图常量 (参考书本 3.2 节)
const (
	width, height = 600, 320            // 画布大小 (像素)
	cells         = 100                 // 网格单元数量 (100x100)
	xyrange       = 30.0                // 坐标轴范围 (-30 到 +30)
	xyscale       = width / 2 / xyrange // x 或 y 轴单位对应的像素数
	zscale        = height * 0.4        // z 轴单位对应的像素数
	angle         = math.Pi / 6         // 等距投影的角度 (30度)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// surface 将 3D 曲面的 SVG 数据写入 out
// f 是用来计算 z 高度的函数，由调用者(plot)提供
func surface(out io.Writer, f func(x, y float64) float64) {
	// 写入 SVG 头信息
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	// 遍历网格
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// 计算网格单元四个角的坐标
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)

			// 检查数值是否有效（防止无穷大或NaN导致SVG渲染错误）
			if math.IsNaN(ax) || math.IsNaN(ay) ||
				math.IsNaN(bx) || math.IsNaN(by) ||
				math.IsNaN(cx) || math.IsNaN(cy) ||
				math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}

			// 绘制多边形
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

// corner 计算网格点 (i,j) 在 2D 画布上的坐标
func corner(i, j int, f func(x, y float64) float64) (float64, float64) {
	// 1. 求出网格点对应的 x, y 坐标 (范围 -30.0 到 +30.0)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 2. 调用传入的函数 f 计算高度 z
	// 这里的 f 实际上会去调用 expr.Eval(...)
	z := f(x, y)

	// 3. 将 (x,y,z) 等距投影到 2D 画布坐标 (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

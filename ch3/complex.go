// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"time"
)

func main() {
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)               // "(-5+10i)"
	fmt.Println(real(x * y))         // "-5"
	fmt.Println(imag(x * y))         // "10"
}
func main2() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	file, err := os.Create("out.png")
	defer file.Close()
	if err != nil {
		fmt.Printf("create file failed: %v\n", err)
		return
	}
	start := time.Now()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin

			// Image point (px, py) represents complex value z.
			img.Set(px, py, superSample(x, y, width, height, xmin, xmax, ymin, ymax))
		}
	}
	png.Encode(file, img) // NOTE: ignoring errors
	fmt.Println(time.Since(start).Seconds())
}

func superSample(centerX, centerY float64, width, height int,
	xmin, xmax, ymin, ymax float64) color.Color {

	// 2x2 子像素偏移
	//offsets := []float64{-0.25, 0.25}
	offsets := []float64{-0.375, -0.125, 0.125, 0.375}

	var r, g, b, a uint32
	count := 0

	// 计算单个像素的大小
	pixelWidth := (xmax - xmin) / float64(width)
	pixelHeight := (ymax - ymin) / float64(height)

	// 采样 4 个子像素
	for _, offsetX := range offsets {
		for _, offsetY := range offsets {
			// 计算子像素坐标
			subX := centerX + offsetX*pixelWidth
			subY := centerY + offsetY*pixelHeight
			z := complex(subX, subY)
			c := mandelbrot(z)

			r1, g1, b1, a1 := c.RGBA()

			r += r1
			g += g1
			b += b1
			a += a1
			count++
		}
	}
	// 平均颜色
	return color.RGBA{
		R: uint8(r / uint32(count) >> 8),
		G: uint8(g / uint32(count) >> 8),
		B: uint8(b / uint32(count) >> 8),
		A: uint8(a / uint32(count) >> 8),
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			// 彩色平滑过渡，从红色到蓝色
			//ratio := float64(n) / iterations
			//return color.RGBA{
			//	uint8(255 * ratio),
			//	uint8(128 * (1 - ratio)),
			//	uint8(255 * (1 - ratio)),
			//	255,
			//}
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotComplex64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		// complex64 没有 cmplx.Abs，手动计算
		if real(v)*real(v)+imag(v)*imag(v) > 4 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

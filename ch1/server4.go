package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette1 = []color.Color{color.White, color.Black, color.RGBA{0, 128, 0, 255}}

const (
	whiteIndex = 0 //画板中的第一种颜色  --这里是画板底色
	blackIndex = 1 // 画板中的下一种颜色
	blueIndex  = 2
)

//
//func main() {
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		lissajousWithCycle(w, r)
//	})
//	log.Fatal(http.ListenAndServe("localhost:8000", nil))
//}

func lissajousWithCycle(out io.Writer, r *http.Request) {
	var cycleStr = r.FormValue("cycles")
	cycles, _ := strconv.Atoi(cycleStr)
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	colorIdx := 1
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette1)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(colorIdx))
		}
		colorIdx = (colorIdx % 7) + 1
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		return
	} // NOTE: ignoring encoding errors
}

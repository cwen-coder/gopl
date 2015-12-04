package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.Black}

const paletteSize = 255

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Red turning to Green
	for i := 1; i < paletteSize/3; i++ {
		offset := i * 3
		r := math.MaxUint8 - offset
		g := offset
		b := 0
		palette = append(palette, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
	}
	// Green turning to Blue
	for i := paletteSize / 3; i < 2*paletteSize/3; i++ {
		offset := (i - paletteSize/3) * 3
		r := 0
		g := math.MaxUint8 - offset
		b := offset
		palette = append(palette, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
	}
	// Blue turning to Red
	for i := 2 * paletteSize / 3; i <= paletteSize; i++ {
		offset := (i - 2*paletteSize/3) * 3
		r := offset
		g := 0
		b := math.MaxUint8 - offset
		palette = append(palette, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
			return
		}
		freq, err := strconv.ParseFloat(r.Form.Get("freq"), 64)
		if err != nil {
			freq = rand.Float64() * 3.0 // relative frequency of y oscillator
		}
		phaseStep, err := strconv.ParseFloat(r.Form.Get("phaseStep"), 64)
		if err != nil {
			phaseStep = 0.01
		}
		cycles, err := strconv.ParseFloat(r.Form.Get("cycles"), 64)
		if err != nil {
			cycles = 5
		}
		lissajous(w, freq, phaseStep, cycles)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func lissajous(out io.Writer, freq float64, phaseStep float64, cycles float64) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// Range from 1 to MaxUint8
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(t*(math.MaxUint8-1)/(cycles*2*math.Pi))+1)
		}
		phase += phaseStep
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

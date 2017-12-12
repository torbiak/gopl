// ex1.6 generates GIF animations of random Lissajous figures, with a
// gradient applied on the time dimension and the number of cycles to display
// can be supplied in the url query string when running as a web server.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

// Packages not needed by version in book.
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	var cycles = 5 // number of complete x oscillator revolutions

	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			cyclesStr := r.FormValue("cycles")
			if cyclesStr != "" {
				var err error
				cycles, err = strconv.Atoi(cyclesStr)
				if err != nil {
					fmt.Fprintf(w, "bad cycles param: %s", cyclesStr, err)
					return
				}
			}
			lissajous(cycles, w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(cycles, os.Stdout)
}

func lissajous(cycles int, out io.Writer) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	palette := make([]color.Color, 0, nframes)
	palette = append(palette, color.RGBA{0, 0, 0, 255})
	for i := 0; i < nframes; i++ {
		scale := float64(i) / float64(nframes)
		c := color.RGBA{uint8(55 + 200*scale), uint8(55 + 200*scale), uint8(55 + 200*scale), 255}
		palette = append(palette, c)
	}
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8((i%(len(palette)-1))+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

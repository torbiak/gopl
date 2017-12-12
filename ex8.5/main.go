// ex8.5 is a parallellized Mandelbrot fractal generator.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"runtime"
	"sync"
	"time"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	workers := runtime.GOMAXPROCS(-1)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	start := time.Now()
	wg := sync.WaitGroup{}
	rows := make(chan int, height)
	for row := 0; row < height; row++ {
		rows <- row
	}
	close(rows)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for py := range rows {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					// Image point (px, py) represents complex value z.
					z := complex(x, y)
					img.Set(px, py, newton(z))
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("rendered in:", time.Since(start))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		png.Encode(w, img) // NOTE: ignoring errors
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// z^4 - 1 = 0
// 4z^3
//
// z' = z - f(x)/f'(x)
// z' = z - (z^4 - 1) / 4z^3
// z' = z - z/4 - z^3/4
// z' = (3z-z^3)/4
// OR z' -= (z-z^3)/4
func newton(z complex128) color.Color {
	iterations := 37
	for n := uint8(0); int(n) < iterations; n++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(cmplx.Pow(z, 4)-1) < 1e-6 {
			return color.Gray{255 - uint8(math.Log(float64(n))/math.Log(float64(iterations+0))*255)}
		}
	}
	return color.Black
}

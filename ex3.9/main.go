// ex3.9 serves images of fractals over http.
//
// Query parameters (interpreted as floating point numbers):
//   xmin, xmax, ymin, and ymax specify the domain
//   zoom: modify the domain to zoom in (for values greater than 1) or out
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
	"strconv"
)

func main() {
	const (
		width, height = 1024, 1024
	)
	params := map[string]float64{
		"xmin": -2,
		"xmax": 2,
		"ymin": -2,
		"ymax": 2,
		"zoom": 1,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for name := range params {
			s := r.FormValue(name)
			if s == "" {
				continue
			}
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				http.Error(w, fmt.Sprintf("query param %s: %s", name, err), http.StatusBadRequest)
				return
			}
			params[name] = f
		}
		if params["xmax"] <= params["xmin"] || params["ymax"] <= params["ymin"] {
			http.Error(w, fmt.Sprintf("min coordinate greater than max"), http.StatusBadRequest)
			return
		}
		xmin := params["xmin"]
		xmax := params["xmax"]
		ymin := params["ymin"]
		ymax := params["ymax"]
		zoom := params["zoom"]

		lenX := xmax - xmin
		midX := xmin + lenX/2
		xmin = midX - lenX/2/zoom
		xmax = midX + lenX/2/zoom
		lenY := ymax - ymin
		midY := ymin + lenY/2
		ymin = midY - lenY/2/zoom
		ymax = midY + lenY/2/zoom

		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value z.
				img.Set(px, py, mandelbrot(z))
			}
		}
		err := png.Encode(w, img)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50: // dark red
				return color.RGBA{100, 0, 0, 255}
			default:
				// logarithmic blue gradient to show small differences on the
				// periphery of the fractal.
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}

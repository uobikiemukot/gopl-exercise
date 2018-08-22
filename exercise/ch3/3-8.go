// Mandelbrot emits a PNG image of the Mandelbrot fractal.
// Supersampling
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

/*
	complex64
	complex128
	big.Float
	Big.Rat
*/

func mandelbrot64(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	c := complex64(z)

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + c
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrot128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7

	for n := uint8(0); n < iterations; n++ {
		z -= (z - 1/(z*z*z)) / 4
		//fmt.Fprintln(os.Stderr, cmplx.Abs(z))
		if cmplx.Abs(complex128(z)) < 0.5 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// Decrease image width/height to 1/2
func reduce(src image.Image) image.Image {
	rect := src.Bounds()
	dst := image.NewGray(image.Rect(0, 0, rect.Dx()/2, rect.Dy()/2))

	for y := 0; y < dst.Bounds().Dy(); y++ {
		for x := 0; x < dst.Bounds().Dx(); x++ {
			p1 := uint32(src.At((x * 2), (y * 2)).(color.Gray).Y)
			p2 := uint32(src.At((x*2)+1, (y * 2)).(color.Gray).Y)
			p3 := uint32(src.At((x * 2), (y*2)+1).(color.Gray).Y)
			p4 := uint32(src.At((x*2)+1, (y*2)+1).(color.Gray).Y)
			dst.Set(x, y, color.Gray{uint8((p1 + p2 + p3 + p4) / 4)})
		}
	}
	return dst
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 4096, 4096
		vwidth, vheight        = width/2, height/2 // visible view size
		vx, vy                 = 0, height/4       // left top coordinate of visible view
	)

	img := image.NewGray(image.Rect(0, 0, vwidth, vheight))
	for py := vy; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := vx; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px-vx, py-vy, mandelbrot128(z))
		}
	}
	//dst := reduce(img)
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

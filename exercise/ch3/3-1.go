// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
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

func main() {
	if math.IsNaN(sin30) || math.IsNaN(cos30) {
		log.Fatalln("sin30 is NaN or cos30 is NaN")
	}

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			var ax, ay, bx, by, cx, cy, dx, dy float64
			var ok bool

			if ax, ay, ok = corner(i+1, j); !ok {
				fmt.Fprintln(os.Stderr, "either ax or ay is NaN")
				continue
			}
			if bx, by, ok = corner(i, j); !ok {
				fmt.Fprintln(os.Stderr, "either bx or by is NaN")
				continue
			}
			if cx, cy, ok = corner(i, j+1); !ok {
				fmt.Fprintln(os.Stderr, "either cx or cy is NaN")
				continue
			}
			if dx, dy, ok = corner(i+1, j+1); !ok {
				fmt.Fprintln(os.Stderr, "either dx or dy is NaN")
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ok := f(x, y)
	if !ok {
		fmt.Fprintf(os.Stderr, "f(%g, %g) returned NaN: ", x, y)
		return 0.0, 0.0, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) (float64, bool) {
	h := math.Hypot(x, y) // distance from (0,0)
	if math.IsNaN(h) {
		fmt.Fprintf(os.Stderr, "math.Hypot(%g, %g) returned NaN: ", x, y)
		return 0.0, false
	}

	r := math.Sin(h) / h
	if math.IsNaN(r) {
		fmt.Fprintf(os.Stderr, "(math.Sin(%g) / %g) returned NaN: ", h, h)
		return 0.0, false
	}

	return r, true
}

//!-

package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells = 100         // number of grid cells
	angle = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surface(out io.Writer, width, height float64, color_min, color_max uint32) {
	if math.IsNaN(sin30) || math.IsNaN(cos30) {
		return
	}

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			var ax, ay, bx, by, cx, cy, dx, dy, z float64
			var ok bool

			if ax, ay, _, ok = corner(i+1, j, width, height); !ok {
				continue
			}
			if bx, by, z, ok = corner(i, j, width, height); !ok {
				continue
			}
			if cx, cy, _, ok = corner(i, j+1, width, height); !ok {
				continue
			}
			if dx, dy, _, ok = corner(i+1, j+1, width, height); !ok {
				continue
			}

			z = z * 128
			z = math.Max(0.0, z)
			//fmt.Printf("z: %f\n", z)
			max := (z / 80.0) * float64(color_max)
			min := (1.0 - (z / 80)) * float64(color_min)
			/*
			max = math.Min(float64(color_max), 0xFF)
			min = math.Min(float64(color_min), 0xFF)
			max = math.Max(float64(color_max), 0.0)
			min = math.Max(float64(color_min), 0.0)
			*/
			c := fmt.Sprintf("#%.6X", uint32(min+max))

			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, c)
		}
	}

	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int, width, height float64) (float64, float64, float64, bool) {
	xyrange := 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale := width / 2 / xyrange // pixels per x or y unit
	zscale := height * 0.4         // pixels per z unit

	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ok := f(x, y)
	if !ok {
		return 0.0, 0.0, 0.0, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, true
}

func f(x, y float64) (float64, bool) {
	// half sphere
	g := func(x float64) float64 {
		const (
			diameter = 0.55
			scale    = 0.09
		)
		return math.Sqrt(math.Pow(diameter, 2) - math.Pow(x*scale, 2))
	}
	r := g(x) + g(y) - 0.5

	// egg box
	//r := (math.Sin(x) * 0.05) + (math.Cos(y) * 0.05)

	// saddles?
	//r := math.Sin(x * 0.2) * 0.2 + math.Cos(y * 0.2) * 0.15

	if math.IsNaN(r) {
		return 0.0, false
	}

	return r, true
}

func handler(w http.ResponseWriter, r *http.Request) {
	// initial values
	// canvas size in pixels
	var width float64 = 600
	var height float64 = 320
	var color_min uint32 = 0x0000FF
	var color_max uint32 = 0xFF0000

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		switch k {
		case "width":
			if f, err := strconv.ParseFloat(v[0], 64); err == nil {
				width = f
			}
		case "height":
			if f, err := strconv.ParseFloat(v[0], 64); err == nil {
				height = f
			}
		case "color_min":
			if u, err := strconv.ParseUint(v[0], 16, 32); err == nil {
				color_min = uint32(u)
			}
		case "color_max":
			if u, err := strconv.ParseUint(v[0], 16, 32); err == nil {
				color_max = uint32(u)
			}
		}
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, width, height, color_min, color_max)
}

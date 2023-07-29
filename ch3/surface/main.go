// Surface computes an SVG rendering of a 3-D surface function and writes it to a GIF file.
package main

import (
	"image"
	"image/color"
	"image/gif"
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
	anim := gif.GIF{LoopCount: -1} // create a new GIF object

	for i := 0; i < 64; i++ { // create 64 frames
		rect := image.Rect(0, 0, width, height)
		img := image.NewPaletted(rect, color.Palette{color.White, color.Black})

		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)

				// Use color.White or color.Black based on (i+j)%2
				imgColor := color.White
				if (i+j)%2 == 1 {
					imgColor = color.Black
				}

				// Set color of each pixel in the image
				img.SetColorIndex(int(ax), int(ay), uint8(img.Palette.Index(imgColor)))
				img.SetColorIndex(int(bx), int(by), uint8(img.Palette.Index(imgColor)))
				img.SetColorIndex(int(cx), int(cy), uint8(img.Palette.Index(imgColor)))
				img.SetColorIndex(int(dx), int(dy), uint8(img.Palette.Index(imgColor)))
			}
		}

		anim.Delay = append(anim.Delay, 8) // delay between frames
		anim.Image = append(anim.Image, img)
	}

	file, err := os.Create("surface.gif")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := gif.EncodeAll(file, &anim); err != nil { // write the GIF file
		panic(err)
	}
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {

	// Image
	const image_width = 1920
	const aspect_ratio = 16.0 / 10.0
	const image_height = int(image_width / aspect_ratio)

	// Camera
	viewport_height := 2.0
	viewport_width := aspect_ratio * viewport_height
	focal_length := 1.0

	origin := NewVec3(0, 0, 0)
	horizontal := NewVec3(viewport_width, 0, 0)
	vertical := NewVec3(0, viewport_height, 0)
	lower_left_corner := origin.Sub(horizontal.DivC(2)).Sub(vertical.DivC(2)).Sub(NewVec3(0, 0, focal_length))

	// Render the image
	img := image.NewNRGBA(image.Rect(0, 0, image_width, image_height))

	// Render to the image
	for j := image_height - 1; j >= 0; j-- {
		fmt.Printf("Rendering line: %d \r", image_height-j)
		for i := 0; i < image_width; i++ {
			u := float64(i) / float64(image_width-1)
			v := float64(j) / float64(image_height-1)
			r := &Ray{
				origin,
				lower_left_corner.Add(horizontal.MultC(u)).Add(vertical.MultC(v)).Sub(origin),
			}
			pixel_color := r.RayColor()

			ir := uint8(255.999 * pixel_color.x)
			ig := uint8(255.999 * pixel_color.y)
			ib := uint8(255.999 * pixel_color.z)

			img.Set(i, j, color.NRGBA{
				R: ir,
				G: ig,
				B: ib,
				A: 255,
			})
		}
	}

	fmt.Println("Done rendering!                  ")

	// Write the image to a file
	f, err := os.Create("render.png")
	if err != nil {
		log.Fatal("Error opening file.", err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal("Error writing PNG.", err)
	}

	if err := f.Close(); err != nil {
		log.Fatal("Error closing file.", err)
	}
}

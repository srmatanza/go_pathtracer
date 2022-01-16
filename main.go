package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"time"
)

func main() {

	// Image
	const image_width = 1280
	const aspect_ratio = 16.0 / 10.0
	const image_height = int(image_width / aspect_ratio)

	// World
	world := NewHittableList()

	world.Add(&Sphere{NewVec3(0, 0, -1), 0.5})
	world.Add(&Sphere{NewVec3(-1, 0.25, -1.5), 0.5})
	world.Add(&Sphere{NewVec3(1, 0.25, -1), 0.5})

	world.Add(&Sphere{NewVec3(0.25, 0, -4), 0.5})
	world.Add(&Sphere{NewVec3(0.375, 0, -9), 0.5})

	world.Add(&Sphere{NewVec3(5, 4, -15), 2.5})

	world.Add(&Sphere{NewVec3(0, -100000.5, -1), 100000})

	// Camera
	viewport_height := 1.6
	viewport_width := aspect_ratio * viewport_height
	focal_length := 1.0

	origin := NewVec3(0, 1, 1)
	horizontal := NewVec3(viewport_width, 0, 0)
	vertical := NewVec3(0, viewport_height, 0)
	lower_left_corner := origin.Sub(horizontal.DivC(2)).Sub(vertical.DivC(2)).Sub(NewVec3(0, 0, focal_length))

	// Render the image
	img := image.NewNRGBA(image.Rect(0, 0, image_width, image_height))

	start_render := time.Now()

	// Render to the image
	for j := 0; j < image_height; j++ {
		fmt.Printf("Rendering line: %d \r", image_height-j)
		for i := 0; i < image_width; i++ {
			u := float64(image_width-i-1) / float64(image_width-1)
			v := float64(image_height-j-1) / float64(image_height-1)
			r := &Ray{
				origin,
				lower_left_corner.Add(horizontal.MultC(u)).Add(vertical.MultC(v)).Sub(origin),
			}
			pixel_color := r.RayColorInWorld(world)

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

	end_render := time.Now()
	fmt.Println("Done rendering!                  ")
	fmt.Printf("Time elapsed: %s\n", end_render.Sub(start_render))

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

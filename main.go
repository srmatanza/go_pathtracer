package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"time"
)

func main() {

	// Image
	const image_width = 512
	const aspect_ratio = 16.0 / 10.0
	const image_height = int(image_width / aspect_ratio)
	const max_depth = 50
	const samples_per_pixel = 100

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
	cam := NewCamera()

	// Render the image
	img := image.NewNRGBA(image.Rect(0, 0, image_width, image_height))

	start_render := time.Now()

	// Render to the image
	for j := 0; j < image_height; j++ {
		fmt.Printf("Rendering... %d%% \r", (j*100)/image_height)
		for i := 0; i < image_width; i++ {
			pixel_color := NewVec3(0, 0, 0)

			for k := 0; k < samples_per_pixel; k++ {
				u := (float64(i) + random_float64()) / float64(image_width-1)
				v := (float64(j) + random_float64()) / float64(image_height-1)
				r := cam.GetRay(u, v)
				new_color := r.RayColorInWorld(world, max_depth)
				pixel_color = pixel_color.Add(new_color)
			}

			scale := 1.0 / samples_per_pixel
			fr := math.Sqrt(scale * pixel_color.x)
			fg := math.Sqrt(scale * pixel_color.y)
			fb := math.Sqrt(scale * pixel_color.z)

			img.Set(i, image_height-j-1, color.NRGBA{
				R: uint8(255.999 * clamp(fr, 0, 0.999)),
				G: uint8(255.999 * clamp(fg, 0, 0.999)),
				B: uint8(255.999 * clamp(fb, 0, 0.999)),
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

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"
)

func clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func random_float64() float64 {
	return rand.Float64()
}

func main() {

	// Image
	const image_width = 512
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
	cam := NewCamera()

	// Render the image
	img := image.NewNRGBA(image.Rect(0, 0, image_width, image_height))

	start_render := time.Now()

	// Render to the image
	for j := 0; j < image_height; j++ {
		fmt.Printf("Rendering... %d%% \r", (j*100)/image_height)
		for i := 0; i < image_width; i++ {

			const samples_per_pixel = 32
			pixel_color := NewVec3(0, 0, 0)

			for k := 0; k < samples_per_pixel; k++ {
				u := (float64(i) + random_float64()) / float64(image_width-1)
				v := (float64(j) + random_float64()) / float64(image_height-1)
				r := cam.GetRay(u, v)
				new_color := r.RayColorInWorld(world)
				pixel_color = pixel_color.Add(new_color)
			}

			ir := uint8(255.999 * clamp(pixel_color.x/float64(samples_per_pixel), 0, 0.999))
			ig := uint8(255.999 * clamp(pixel_color.y/float64(samples_per_pixel), 0, 0.999))
			ib := uint8(255.999 * clamp(pixel_color.z/float64(samples_per_pixel), 0, 0.999))

			img.Set(i, image_height-j-1, color.NRGBA{
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

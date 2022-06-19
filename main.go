package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

// Render Settings
const max_depth = 50

var cpuprofile = flag.String("prof", "", "Write a cpu profile to the specified file.")
var memprofile = flag.String("mem", "", "Write a memory profile to the specified file.")
var samples_per_job = flag.Int("s", 100, "Number of samples per job.")
var num_jobs = flag.Int("j", 1, "Number of jobs to run simultaneously.")

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		defer f.Close()

		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// World
	world := NewHittableList()

	lamb_red := Lambertian{NewVec3(0.7, 0.3, 0.3)}
	lamb_grey := Lambertian{NewVec3(0.8, 0.8, 0.8)}

	metal_blue := Metal{NewVec3(0.2, 0.2, 0.8), 1.0}
	metal_silver := Metal{NewVec3(0.8, 0.8, 0.8), 0.01}

	world.Add(&Sphere{NewVec3(0, 0, -1), lamb_red, 0.5})

	world.Add(&Sphere{NewVec3(-1.5, 0.25, -1.5), metal_silver, 0.75})

	world.Add(&Sphere{NewVec3(1, 0.25, -1), metal_blue, 0.5})
	world.Add(&Sphere{NewVec3(0.25, 0, -4), lamb_grey, 0.5})
	world.Add(&Sphere{NewVec3(0.375, 0, -9), lamb_grey, 0.5})

	world.Add(&Sphere{NewVec3(5, 4, -15), lamb_grey, 2.5})

	world.Add(&Sphere{NewVec3(0, -100000.5, -1), Lambertian{NewVec3(0.8, 0.8, 0.0)}, 100000})

	// Camera
	cam := NewCamera()

	// Create the Render Image
	render := NewRender(512, 16.0/10.0, *samples_per_job, *num_jobs)

	start_render := time.Now()

	// Render to the image
	img := render.RenderImage(cam, world)

	end_render := time.Now()

	// Write memory profile if necessary
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("Couldn't write memory profile: ", err)
		}
	}

	fmt.Println("Done rendering!")
	fmt.Printf("Total time elapsed: %s\n", end_render.Sub(start_render))

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

package main

import (
	"testing"
)

func BenchmarkRenderImage(b *testing.B) {
	world := NewHittableList()
	world.Add(&Sphere{NewVec3(0, 0, -1), &Lambertian{NewVec3(0.8, 0.8, 0.8)}, 0.5})
	world.Add(&Sphere{NewVec3(0, -100000.5, -1), &Lambertian{NewVec3(0.8, 0.8, 0.8)}, 100000})

	cam := NewCamera()

	render := NewRender(512, 16.0/10.0, 16, 4)

	render.RenderImage(cam, world)
}

package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"sync"
	"time"
)

type Render struct {
	image_width, image_height int
	aspect_ratio              float64
	img                       []Vec3
	samples,
	total_samples,
	parallel_renders int
	render_lock  sync.Mutex
	render_queue chan bool
	render_time,
	combine_time time.Duration
}

func NewRender(width int, aspect_ratio float64, total_samples, jobs int) *Render {
	height := int(float64(width) / aspect_ratio)
	ret := &Render{
		image_width:      width,
		aspect_ratio:     16.0 / 10.0,
		image_height:     height,
		img:              make([]Vec3, width*height),
		parallel_renders: jobs,
		total_samples:    total_samples,
	}
	ret.render_queue = make(chan bool)
	return ret
}

func (r *Render) waitForRenderToFinish() {
	for i := 0; i < r.parallel_renders; i++ {
		<-r.render_queue
	}
	log.Println("WaitForRenderToFinish is done.")
}

func (r *Render) RenderImage(cam *Camera, world *HittableList) *image.NRGBA {
	final_img := image.NewNRGBA(image.Rectangle{image.Pt(0, 0), image.Pt(r.image_width, r.image_height)})

	samples_per_job := r.total_samples / r.parallel_renders
	remaining_samples := r.total_samples % r.parallel_renders
	for i := 0; i < r.parallel_renders; i++ {
		samples := samples_per_job
		if i == 0 {
			samples = samples + remaining_samples
		}
		go r.renderImage(cam, world, samples)
	}

	r.waitForRenderToFinish()

	fmt.Printf("Time Rendering: %v\nTime Combining Images: %v\n", r.render_time, r.combine_time)

	for j := 0; j < r.image_height; j++ {
		for i := 0; i < r.image_width; i++ {

			pixel_color := r.img[i+j*r.image_width]
			fr := math.Sqrt(pixel_color.x)
			fg := math.Sqrt(pixel_color.y)
			fb := math.Sqrt(pixel_color.z)

			final_img.Set(i, r.image_height-j-1, color.NRGBA{
				R: uint8(255.999 * clamp(fr, 0, 0.999)),
				G: uint8(255.999 * clamp(fg, 0, 0.999)),
				B: uint8(255.999 * clamp(fb, 0, 0.999)),
				A: 255,
			})
		}
	}

	return final_img
}

func (r *Render) renderImage(cam *Camera, world *HittableList, render_samples int) {
	defer func() { r.render_queue <- true }()

	if render_samples == 0 {
		return
	}

	sample_img := make([]Vec3, r.image_width*r.image_height)

	log.Printf("Calling renderImage...\n")
	render_start_t := time.Now()
	for j := 0; j < r.image_height; j++ {
		// fmt.Printf("Rendering... %d%% \r", (j*100)/r.image_height)
		for i := 0; i < r.image_width; i++ {
			pixel_color := NewVec3(0, 0, 0)

			for k := 0; k < render_samples; k++ {
				u := (float64(i) + random_float64()) / float64(r.image_width-1)
				v := (float64(j) + random_float64()) / float64(r.image_height-1)
				r := cam.GetRay(u, v)
				new_color := r.RayColorInWorld(world, max_depth)
				pixel_color = pixel_color.Add(new_color)
			}

			scale := 1.0 / float64(render_samples)
			pixel_color = pixel_color.MultC(scale)
			sample_img[i+j*r.image_width] = pixel_color
		}
	}
	d_render := time.Since(render_start_t)
	log.Printf("Done rendering %d samples.\nCombining image...\n", render_samples)

	// Combine sample_img and r.img
	r.render_lock.Lock()
	combine_start_t := time.Now()
	for j := 0; j < r.image_height; j++ {
		for i := 0; i < r.image_width; i++ {
			sample_pixel := sample_img[i+j*r.image_width]
			img_pixel := r.img[i+j*r.image_width]

			new_color := sample_pixel.MultC(float64(render_samples)).Add(img_pixel.MultC(float64(r.samples)))
			new_color = new_color.DivC(float64(render_samples + r.samples))

			r.img[i+j*r.image_width] = new_color
		}
	}
	d_combine := time.Since(combine_start_t)

	log.Println("Done combining image.")

	r.render_time = r.render_time + d_render
	r.combine_time = r.combine_time + d_combine

	r.samples += render_samples
	r.render_lock.Unlock()
}

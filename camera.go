package main

type Camera struct {
	origin,
	lower_left_corner,
	horizontal,
	vertical *Vec3
}

func NewCamera() *Camera {
	aspect_ratio := 16.0 / 10.0
	viewport_height := 1.6
	viewport_width := aspect_ratio * viewport_height
	focal_length := 1.0

	o := NewVec3(0, 0, 0)
	h := NewVec3(viewport_width, 0, 0)
	v := NewVec3(0, viewport_height, 0)
	llc := o.Sub(h.DivC(2)).Sub(v.DivC(2)).Sub(NewVec3(0, 0, focal_length))

	return &Camera{
		origin:            o,
		lower_left_corner: llc,
		horizontal:        h,
		vertical:          v,
	}
}

func (c *Camera) GetRay(u, v float64) *Ray {
	return &Ray{
		c.origin,
		c.lower_left_corner.Add(c.horizontal.MultC(u)).Add(c.vertical.MultC(v)).Sub(c.origin),
	}
}

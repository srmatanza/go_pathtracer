package main

import "math"

type Sphere struct {
	center *Vec3
	radius float64
}

func (s *Sphere) Hit(r *Ray, t_min, t_max float64) (bool, *HitRecord) {
	oc := r.origin.Sub(s.center)
	a := r.direction.LengthSquared()
	half_b := oc.Dot(r.direction)
	c := oc.LengthSquared() - s.radius*s.radius
	discriminant := half_b*half_b - a*c

	if discriminant < 0 {
		return false, nil
	}

	sqrtd := math.Sqrt(discriminant)

	// Find the nearest root that lies in the t_range
	root := (-half_b - sqrtd) / a
	if root < t_min || t_max < root {
		root = (-half_b + sqrtd) / a
		if root < t_min || t_max < root {
			return false, nil
		}
	}

	t := root
	p := r.At(t)
	normal := p.Sub(s.center).DivC(s.radius)
	rec := &HitRecord{p, normal, t, true}
	rec.SetFaceNormal(r, normal)

	return true, rec
}

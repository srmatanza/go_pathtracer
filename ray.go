package main

import (
	"math"
)

const infinity = math.MaxFloat64

type Ray struct {
	origin    *Vec3
	direction *Vec3
}

func (r *Ray) At(t float64) *Vec3 {
	return r.origin.Add(r.direction.MultC(t))
}

func (r *Ray) RayColorInWorld(world *HittableList, depth int) *Vec3 {

	if depth <= 0 {
		return NewVec3(0, 0, 0)
	}

	if is_hit, hit_record := world.Hit(r, 0.00001, infinity); is_hit {
		// target := hit_record.p.Add(hit_record.normal).Add(RandomInUnitSphere())
		// target := hit_record.p.Add(hit_record.normal).Add(RandomUnitVector())
		target := hit_record.p.Add(RandomInHemisphere(hit_record.normal))
		new_ray := &Ray{hit_record.p, target.Sub(hit_record.p)}
		return new_ray.RayColorInWorld(world, depth-1).MultC(0.5)
	}

	unit_direction := r.direction.Unit()
	t := 0.5 * (unit_direction.y + 1.0)
	return NewVec3(1, 1, 1).MultC(1.0 - t).Add(NewVec3(0.5, 0.7, 1.0).MultC(t))
}

func (r *Ray) RayColor() *Vec3 {
	t := r.HitSphere(NewVec3(0, 0, -1), 0.5)
	if t > 0 {
		N := r.At(t).Sub(NewVec3(0, 0, -1)).Unit()
		return NewVec3(N.x+1, N.y+1, N.z+1).MultC(0.5)
	}

	// Render background color
	unit_direction := r.direction.Unit()
	t = 0.5 * (unit_direction.y + 1.0)
	return NewVec3(1, 1, 1).MultC(1.0 - t).Add(NewVec3(0.5, 0.7, 1.0).MultC(t))
}

func (r *Ray) HitSphere(center *Vec3, radius float64) float64 {
	oc := r.origin.Sub(center)
	a := r.direction.LengthSquared()
	half_b := oc.Dot(r.direction)
	c := oc.LengthSquared() - radius*radius
	discriminant := half_b*half_b - a*c

	if discriminant < 0 {
		return -1.0
	} else {
		return (-half_b - math.Sqrt(discriminant)) / a
	}
}

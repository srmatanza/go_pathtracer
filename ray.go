package main

type Ray struct {
	origin    *Vec3
	direction *Vec3
}

func (r *Ray) At(t float64) *Vec3 {
	return r.origin.Add(r.direction.MultC(t))
}

func (r *Ray) RayColor() *Vec3 {
	if r.HitSphere(NewVec3(0, 0, -1), 0.5) {
		return NewVec3(1, 0, 0)
	}

	// Render background color
	unit_direction := r.direction.Unit()
	t := 0.5 * (unit_direction.y + 1.0)
	return NewVec3(1, 1, 1).MultC(1.0 - t).Add(NewVec3(0.5, 0.7, 1.0).MultC(t))
}

func (r *Ray) HitSphere(center *Vec3, radius float64) bool {
	oc := r.origin.Sub(center)
	a := r.direction.Dot(r.direction)
	b := oc.Dot(r.direction) * 2.0
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c

	return discriminant > 0
}

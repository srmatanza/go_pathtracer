package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Vec3 struct {
	x, y, z float64
}

func NewVec3(x, y, z float64) *Vec3 {
	return &Vec3{x, y, z}
}

func RandomVec3(min, max float64) *Vec3 {
	return &Vec3{rand_rng_float64(min, max), rand_rng_float64(min, max), rand_rng_float64(min, max)}
}

func RandomInUnitSphere() *Vec3 {
	for {
		p := RandomVec3(-1, 1)
		if p.LengthSquared() >= 1 {
			continue
		}
		return p
	}
}

func RandomUnitVector() *Vec3 {
	return RandomInUnitSphere().Unit()
}

func RandomInHemisphere(normal *Vec3) *Vec3 {
	ius := RandomInUnitSphere()
	if ius.Dot(normal) > 0 {
		return ius
	} else {
		return ius.Neg()
	}
}

func (v Vec3) String() string {
	return fmt.Sprintf("%f %f %f", v.x, v.y, v.z)
}

func (v *Vec3) Equals(u *Vec3) bool {
	if u.x == v.x && u.y == v.y && u.z == v.z {
		return true
	}
	return false
}

func (v *Vec3) Neg() *Vec3 {
	return &Vec3{-v.x, -v.y, -v.z}
}

func (v *Vec3) Add(u *Vec3) *Vec3 {
	return &Vec3{v.x + u.x, v.y + u.y, v.z + u.z}
}

func (v *Vec3) Sub(u *Vec3) *Vec3 {
	return &Vec3{v.x - u.x, v.y - u.y, v.z - u.z}
}

func (v *Vec3) Mult(u *Vec3) *Vec3 {
	return &Vec3{v.x * u.x, v.y * u.y, v.z * u.z}
}

func (v *Vec3) MultC(c float64) *Vec3 {
	return &Vec3{v.x * c, v.y * c, v.z * c}
}

func (v *Vec3) DivC(c float64) *Vec3 {
	return &Vec3{v.x / c, v.y / c, v.z / c}
}

func (v *Vec3) Dot(u *Vec3) float64 {
	return u.x*v.x + u.y*v.y + u.z*v.z
}

func (u *Vec3) Cross(v *Vec3) *Vec3 {
	return &Vec3{
		u.y*v.z - u.z*v.y,
		u.z*v.x - u.x*v.z,
		u.x*v.y - u.y*v.x,
	}
}

func (u *Vec3) Unit() *Vec3 {
	return u.DivC(u.Length())
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v *Vec3) LengthSquared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

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

func rand_rng_float64(min, max float64) float64 {
	r := rand.Float64()
	return (r * (max - min)) + min
}

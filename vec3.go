package main

import (
	"fmt"
	"math"
)

type Vec3 struct {
	x, y, z float64
}

func NewVec3(x, y, z float64) *Vec3 {
	return &Vec3{x, y, z}
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

package main

import (
	"testing"
)

func TestVectorOps(t *testing.T) {
	origin := Vec3{0, 0, 0}
	a := Vec3{1, 1, 1}
	b := Vec3{1, 2, 3}
	c := Vec3{2, 0, 0}
	k := 2.0

	ans := origin.Add(a)
	expected := Vec3{1, 1, 1}
	if !ans.Equals(expected) {
		t.Errorf("TestVectorOps fail: %s\n\t\t\texpected %s", ans, expected)
	}

	ans = a.Neg()
	expected = Vec3{-1, -1, -1}
	if !ans.Equals(expected) {
		t.Errorf("TestVectorOps fail: %s\n\t\t\texpected %s", ans, expected)
	}

	ans = b.Sub(a)
	expected = Vec3{0, 1, 2}
	if !ans.Equals(expected) {
		t.Errorf("TestVectorOps fail: %s\n\t\t\texpected %s", ans, expected)
	}

	ans = b.Mult(b)
	expected = Vec3{1, 4, 9}
	if !ans.Equals(expected) {
		t.Errorf("TestVectorOps fail: %s\n\t\t\texpected %s", ans, expected)
	}

	ans = a.MultC(k)
	expected = Vec3{k, k, k}
	if !ans.Equals(expected) {
		t.Errorf("TestVectorOps fail: %s\n\t\t\texpected %s", ans, expected)
	}

	ans = b.DivC(k)
	expected = Vec3{1 / k, 2 / k, 3 / k}
	if !ans.Equals(expected) {
		t.Errorf("TestVectorOps fail: %s\n\t\t\texpected %s", ans, expected)
	}

	ans = c.Unit()
	expected = Vec3{1, 0, 0}
	if !ans.Equals(expected) {
		t.Errorf("TestVectorOps fail: %s\n\t\t\texpected %s", ans, expected)
	}

	ans = c.Cross(Vec3{0, 1, 0})
	expected = Vec3{0, 0, 2}
	if !ans.Equals(expected) {
		t.Errorf("TestVectorOps fail: %s\n\t\t\texpected %s", ans, expected)
	}

	if c.Dot(Vec3{0, 1, 0}) != 0.0 {
		t.Errorf("TestVectorOps fail: %s\n\t\t\texpected %f", c, 0.0)
	}

	if c.Length() != 2.0 {
		t.Errorf("TestVectorOps fail: %s\n\t\t\texpected %f", c, 2.0)
	}

	if c.LengthSquared() != 4.0 {
		t.Errorf("TestVectorOps fail: %s\n\t\t\texpected %f", c, 4.0)
	}
}

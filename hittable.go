package main

type HitRecord struct {
	p, normal  Vec3
	mat        Material
	t          float64
	front_face bool
}

func (h *HitRecord) SetFaceNormal(r Ray, outward_normal Vec3) {
	h.front_face = r.direction.Dot(outward_normal) < 0
	if h.front_face {
		h.normal = outward_normal
	} else {
		h.normal = outward_normal.Neg()
	}
}

type Hittable interface {
	Hit(r Ray, t_min, t_max float64) (bool, *HitRecord)
}

type HittableList struct {
	objects []Hittable
}

func NewHittableList() *HittableList {
	hl := &HittableList{}
	hl.Clear()
	return hl
}

func (hl *HittableList) Clear() {
	hl.objects = make([]Hittable, 0)
}

func (hl *HittableList) Add(h Hittable) {
	hl.objects = append(hl.objects, h)
}

func (hl *HittableList) Hit(r Ray, t_min, t_max float64) (hit_anything bool, ret_rec *HitRecord) {
	closest_so_far := t_max

	for _, obj := range hl.objects {
		if is_hit, rec := obj.Hit(r, t_min, closest_so_far); is_hit {
			hit_anything = true
			closest_so_far = rec.t
			ret_rec = rec
		}
	}

	return
}

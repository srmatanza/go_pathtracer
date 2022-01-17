package main

type Material interface {
	Scatter(r *Ray, rec *HitRecord) (attenuation *Vec3, scattered *Ray, success bool)
}

type Lambertian struct {
	albedo *Vec3
}

func (l Lambertian) Scatter(r *Ray, rec *HitRecord) (attenuation *Vec3, scattered *Ray, success bool) {
	scatter_direction := rec.normal.Add(RandomUnitVector())
	if scatter_direction.NearZero() {
		scatter_direction = rec.normal
	}
	scattered = &Ray{rec.p, scatter_direction}
	attenuation = l.albedo
	success = true
	return
}

type Metal struct {
	albedo *Vec3
	fuzz   float64
}

func reflect(v *Vec3, n *Vec3) *Vec3 {
	u := n.MultC(2.0 * v.Dot(n))
	return v.Sub(u)
}

func (m Metal) Scatter(r *Ray, rec *HitRecord) (attenuation *Vec3, scattered *Ray, success bool) {
	reflected := reflect(r.direction.Unit(), rec.normal)
	scattered = &Ray{rec.p, reflected.Add(RandomInUnitSphere().MultC(m.fuzz))}
	attenuation = m.albedo
	success = scattered.direction.Dot(rec.normal) > 0
	return
}

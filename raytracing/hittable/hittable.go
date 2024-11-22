package hittable

import (
  "raytracing/vec3"
  "raytracing/ray"
)

type HitRecord struct {
  P vec3.Point3
  Normal vec3.Vec3
  T float64
  FrontFace bool
}

type Hittable interface {
  Hit(r ray.Ray, tMin, tMax float64, rec *HitRecord) bool
}

func (h *HitRecord) SetFaceNormal(r ray.Ray, outwardNormal vec3.Vec3) {
  if vec3.Dot(r.Direction, outwardNormal) < 0.0 {
    h.FrontFace = false
  } else {
    h.FrontFace = true
  }
  if h.FrontFace {
    h.Normal = outwardNormal
  } else {
    h.Normal = outwardNormal.Negate()
  }
}

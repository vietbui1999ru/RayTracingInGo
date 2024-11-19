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

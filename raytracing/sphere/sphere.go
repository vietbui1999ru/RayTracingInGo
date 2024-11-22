package sphere

import (
	"math"
	"raytracing/hittable"
	"raytracing/ray"
	"raytracing/vec3"
)

type Sphere struct {

  Center vec3.Point3
  Radius float64

}

func (s Sphere) Hit(r ray.Ray, ray_tmin, ray_tmax float64, rec *hittable.HitRecord) bool {
  oc := s.Center.Subtract(r.Origin)
  a := vec3.LengthSquared(r.Direction)

  h := vec3.Dot(r.Direction, oc)
  c := vec3.LengthSquared(oc) - (s.Radius * s.Radius)

  discriminant := h*h - a*c

  if discriminant < 0 {
    return false
  }

  sqrtd := math.Sqrt(discriminant)

  // Find nearest root that lies in acceptable range

  root := (h - sqrtd) / a

  if root <= ray_tmin || root >= ray_tmax {
    root = (h + sqrtd) / a
    if root <= ray_tmin || root >= ray_tmax {
      return false
    }
  }

  rec.T = root
  rec.P = r.At(rec.T)

  outwardNormal := rec.P.Subtract(s.Center).DivByScalar(s.Radius)
  // rec.Normal = outwardNormal
  rec.SetFaceNormal(r, outwardNormal)
  
  return true


}



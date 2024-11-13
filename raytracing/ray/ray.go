package ray

import "raytracing/vec3"

// Ray : function of P(t) = A + tB
type Ray struct {
  Origin vec3.Point3 // A: Ray Origin
  Direction vec3.Vec3 // B: Ray Direction
}

// NewRay: Creates new Ray with Origin and Direction as inputs
func NewRay(origin, direction vec3.Vec3) Ray {
  return Ray{
    Origin: origin,
    Direction: direction,
  }
}

// At: returns the Ray at position t along the ray 
func (r Ray) At(t float64) vec3.Vec3 {
  return r.Origin.Add(r.Direction.MultByScalar(t))
}

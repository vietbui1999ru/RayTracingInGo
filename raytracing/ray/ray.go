package ray

import (
	"math"
	color "raytracing/colorUtil"
	"raytracing/vec3"
)

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

func RayColor(r Ray) color.Color {
    // fmt.Printf("Vector: %v\n", r.Direction) Works
  t := HitSphere(vec3.Point3{X:0,Y:0,Z:-1}, 0.5, r)
  if t > 0 {
    Pt := r.At(t)
    N := (Pt.Subtract(vec3.Vec3{X:0, Y:0, Z:-1})).UnitVector()
    return color.Color{X:N.X+1, Y:N.Y+1, Z:N.Z+1}.MultByScalar(0.5)
  }
  unitDirection := r.Direction.UnitVector()
  // fmt.Printf("UnitDirection : %v\n", unitDirection) // UnitVector not Working
  a := 0.5 * unitDirection.Y + 1.0
  // a = vec3.Clamp(a, 0, 1)
    
    // Use Add without modifying in place to ensure clean results
  return color.Color{X:1.0, Y:1.0, Z:1.0}.MultByScalar(1.0 - a).Add(color.Color{X:0.5, Y:0.7, Z:1.0}.MultByScalar(a))
}

func HitSphere(center vec3.Point3, radius float64, r Ray) float64 {
  oc := center.Subtract(r.Origin)
  a := vec3.Dot(r.Direction, r.Direction)

  h := vec3.Dot(r.Direction, oc)
  c := vec3.LengthSquared(oc) - radius * radius

  // --- OLD
  // b := -2.0 * vec3.Dot(r.Direction, oc)
  // c := vec3.Dot(oc, oc) - radius * radius
  // discriminant := b*b - 4 * a * c
  // ---


  discriminant := h*h - a*c
  // evaulate to true, means that we have 0 or 1 or 2 real solutions

  if (discriminant < 0) {
    return -1.0
  } else {
  return (h - math.Sqrt(discriminant)) / a
  }
}

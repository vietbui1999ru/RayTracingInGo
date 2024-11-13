package vec3

import (
	"fmt"
	"math"
)

// Vec3: V = [x, y, z]. 3D Vector
type Vec3 struct {
  X, Y, Z float64
}
// Geometric readability with alias
type Point3 = Vec3

// Create new 3D Vector with specified Coordinates.
// i.e: v := NewVec3(3, 6, 9)
func NewVec3(x, y, z float64) Vec3 {
  return Vec3{X: x, Y: y, Z:z}
}

// No need, can access directly
// func (v Vec3) GetX() float64 { 
//   return v.X 
// }
// func (v Vec3) GetY() float64 {
//   return v.Y 
// }
// func (v Vec3) GetZ() float64 {
//   return v.Z 
// }

// Utility Functions for Vector Manipulation

// Returns the Negation of original Vector.
// i.e: V = -V, -V = V
func (v Vec3) Negate() Vec3 {
   var a float64
   var b float64
   var c float64
   if v.X != 0 {
     a = -v.X
   } else {
     a = -v.X
   }
   if v.X != 0 {
     b = -v.Y
   } else {
     b = -v.Y
   }
   if v.Z != 0 {
     c = -v.Z
   } else {
     c = -v.Z
   }
  // return Vec3{X: -v.X, Y: -v.Y, Z:-v.Z}

  return Vec3{X: a, Y: b, Z: c}
}

// At: returns value of specified coordinate x:0, y:1, z:2
func (v Vec3) At(i int) (float64, error) {
  switch i {
  case 0:
    return v.X, nil
  case 1:
    return v.Y, nil
  case 2:
    return v.Z, nil
  default:
    return 0, fmt.Errorf("index out of range")
  }
}

// SetAt: change value at coordinates x:0, y:1, z:2
func (v *Vec3) SetAt(i int, value float64) error {

  switch i {
  case 0:
    v.X = value
  case 1:
    v.Y = value
  case 2:
    v.Z = value
  default:
    return fmt.Errorf("Index out of range")
  }
  return nil
}

// AddInPlace Vector addition, V + U, where V, U are Diff Vectors
// returning Vec3 supports chaining a + b + c + ...
func (v Vec3) Add(u Vec3) Vec3 {
  return Vec3{X: v.X + u.X, Y: v.Y + u.Y, Z: v.Z + u.Z}
}
func (v *Vec3) AddInPlace(u Vec3) *Vec3 {
  v.X += u.X
  v.Y += u.Y
  v.Z += u.Z
  return v
}

// AddByScalar: Vector addition by a scalar
func (v Vec3) AddByScalar(u float64) Vec3 {
  return Vec3{X: v.X + u, Y: v.Y + u, Z: v.Z + u}
}
func (v *Vec3) AddByScalarInPlace(u float64) *Vec3 {
  v.X += u
  v.Y += u
  v.Z += u
  return v
}

// SubtractInPlace: Vector subtraction, V - U 
// returning Vec3 supports chaining a - b - c ...
func (v Vec3) Subtract(u Vec3) Vec3 {
  b := u.Negate()
  return v.Add(b)
}

func (v *Vec3) SubtractInPlace(u Vec3) *Vec3 {
  b := u.Negate()
  return v.AddInPlace(b)
}

// AddByScalar: Vector subtraction by a scalar
func (v Vec3) SubtractByScalar(u float64) Vec3 {
  return Vec3{X: v.X - u, Y: v.Y - u, Z: v.Z - u}

}
func (v *Vec3) SubtractByScalarInPlace(u float64) *Vec3 {
  v.X -= u
  v.Y -= u
  v.Z -= u
  return v
}

// MultByScalar: Vector Multiplication with a scalar value. Modifies and Points to Original Vector Address
func (v Vec3) MultByScalar(u float64) Vec3 {
  return Vec3{X: v.X * u, Y: v.Y * u, Z: v.Z * u}
}
func (v *Vec3) MultByScalarInPlace(u float64) *Vec3 {
  v.X *= u
  v.Y *= u
  v.Z *= u
  return v
}

// MultByVector: Element-wise multiplcation of 2 Vectors (Hadamard/Schur product). Modifies and Points to Original Vector Address
func (v Vec3) MultByVector(u Vec3) Vec3 {
return Vec3{
  X: v.X * u.X,
  Y: v.Y * u.Y,
  Z: v.Z * u.Z,
  }
}

// Div: Vector Division with a scalar value. Modifies and Points to Original Vector Address
func (v *Vec3) DivByScalarInPlace(u float64) *Vec3 {
  return v.MultByScalarInPlace(1/u)
}
func (v Vec3) DivByScalar(u float64) Vec3 {
  return v.MultByScalar(1/u)
}

// Magnitude: |V| = sqrt(x^2 + y^2 + z^2)
func Magnitude(v Vec3) float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

// LengthSquared l = x^2 + y^2 + z^2
func LengthSquared(v Vec3) float64 {
  return v.X * v.X + v.Y * v.Y + v.Z * v.Z
}

// Dot: Dot product between 2 vectors. Return scalar value
func Dot(v , u Vec3) float64 {
  return v.X * u.X + v.Y * u.Y + v.Z * u.Z
}

// Cross: Cross product between 2 vectors. Return a new Vector
func Cross(v, u Vec3) Vec3 {
  return NewVec3(
    v.Y * u.Z - v.Z * u.Y,
    v.Z * u.X - v.X * u.Z,
    v.X * u.Y - v.Y * u.X,
    )
}

// UnitVector: Returns unit vector in same direction as Original Vector v.
func (v Vec3) UnitVector() Vec3 {
  m := Magnitude(v)
  // fmt.Printf("Magnitude %v\n", m)
  return Vec3{X: v.X / m,Y: v.Y / m, Z: v.Z / m }
}
func (v Vec3) UnitVectorInPlace() *Vec3 {
  return v.DivByScalarInPlace(Magnitude(v))
}
func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

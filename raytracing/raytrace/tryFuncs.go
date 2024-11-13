package main

import (
	"fmt"
	"log"
	"raytracing/pkg"
	"raytracing/ray"
	"raytracing/vec3"
)

func Try() {

  err:= pkg.RenderFirstImage("output.ppm")
  if err != nil {
    log.Fatalf("Failed to render image: %v", err)
  }

  v:= vec3.NewVec3(4, 10, 5)
  var point vec3.Point3 = vec3.NewVec3(5, 2, -5)
  fmt.Println("Point: ", point)
  var errpoint vec3.Point3
  fmt.Println("ErrorPoint will print default: \n", errpoint)
  compZ, err := v.At(1) // 0, 1, 2
  if err != nil {
    log.Fatalf("Unkown value: %s", err)
  }
  v.MultByScalar(3)
  point = v.MultByVector(point)
  v.AddInPlace(point)
  negatedV := v.Negate()
  unitVector := v.UnitVector()
  mag := vec3.Magnitude(v)
  fmt.Printf("New Vector of value : %.2f\n", v)
  fmt.Printf("New Vector negation of value : %.2f\n", negatedV)
  fmt.Printf("Vector's component z : %v\n", compZ)
  fmt.Printf("Vector's magnitude : %.2f\n", mag)
  fmt.Printf("Vector's unit vector : %.2f\n", unitVector)

  fmt.Println()
  firstRay := ray.NewRay(point, v)
  fmt.Printf("First Ray: %v\n", firstRay)
  fmt.Printf("Ray with linear interpolation: %v\n", firstRay.At(5))


}

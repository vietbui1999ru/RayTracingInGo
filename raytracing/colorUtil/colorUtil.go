package color

import (
	"fmt"
	"os"
	"raytracing/vec3"
)

type Color = vec3.Vec3

func WriteColor(pixelColor Color, file *os.File) {
  r := pixelColor.X
  g := pixelColor.Y
  b := pixelColor.Z
  rbyte := int(255.999 * r)
  gbyte := int(255.999 * g)
  bbyte := int(255.999 * b)

  // fmt.Printf("Pixel Color Components:\n %v %v %v\n", rbyte, gbyte, bbyte)
  fmt.Fprintf(file, "%d %d %d\n", rbyte, gbyte, bbyte)
}

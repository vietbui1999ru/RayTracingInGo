package main

import (
	"fmt"
	"os"
	color "raytracing/colorUtil"
	"raytracing/ray"
	"raytracing/vec3"
)

var(
  aspectRatio float64 = 16.0 / 9.0
  imageWidth int = 400
  imageHeight int = int(float64(imageWidth)/aspectRatio)
  focalLength float64 = 1.0
  viewportHeight float64 = 2.0
  viewportWidth float64 = viewportHeight * (float64(imageWidth) / float64(imageHeight))
  
  cameraCenter vec3.Point3 = vec3.Point3{X:0, Y: 0, Z: 0}
)


func main() {

  if imageHeight < 1 {
    imageHeight = 1
  }

  // Calc vectors across horizontal and down the vertical viewport edges
  viewport_u := vec3.NewVec3(viewportWidth, 0, 0)
  viewport_v := vec3.NewVec3(0, -viewportHeight, 0)

  // Calc horizontal and vertical delta vectors from pixel to pixel
  pixel_delta_u := viewport_u.DivByScalar(float64(imageWidth))
  pixel_delta_v := viewport_v.DivByScalar(float64(imageHeight))

  // Calculate the location of upper left pixel P(0, 0)
  viewportUpperLeft := cameraCenter.Subtract(vec3.NewVec3(0, 0, focalLength)).Subtract(viewport_u.DivByScalar(2)).Subtract(viewport_v.DivByScalar(2))
  pixelOriginLoc := viewportUpperLeft.AddInPlace((pixel_delta_u.Add(pixel_delta_v)).MultByScalar(0.5))

  // Render
  fmt.Printf("cameraCenter: %v\n", cameraCenter)
  fmt.Printf("viewportwidth: %v\n: ", viewportWidth)


  file, err := os.Create("output.ppm")
  if err != nil {
    fmt.Println("Error: %w", err)
    os.Exit(1)
  }
  defer file.Close()

  fmt.Fprintf(file, "P3\n%v %v\n255\n", imageWidth, imageHeight)
  for j:=0; j<imageHeight; j++ {
    // fmt.Printf("\rScanlines remaining: %d\n", imageHeight - j)

    for i:=0; i < imageWidth; i++ {
      a := pixel_delta_u.MultByScalar(float64(i))
      b := pixel_delta_v.MultByScalar(float64(j))
      c := a.Add(b)
      pixelCenter := pixelOriginLoc.Add(c)
      // fmt.Printf("Pixel Center : %v\n", pixelCenter)
      rayDirection := pixelCenter.Subtract(cameraCenter)
      // fmt.Printf("rayDirection: %v\n", rayDirection)
      r := ray.NewRay(cameraCenter, rayDirection)
      // fmt.Printf("Ray r := %v\n", r)

      pixelColor := ray.RayColor(r)
      // fmt.Printf("Pixel Color: %v\n", pixelColor)
      color.WriteColor(pixelColor, file)

    }
  }

  fmt.Println("\nDone")

}

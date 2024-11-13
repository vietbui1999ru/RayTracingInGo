package pkg

import (
	"fmt"
	"os"
  "raytracing/colorUtil"
)

var (

  imageWidth int = 256
  imageHeight int = 256

)

// RenderFirstImage : renders the image
func RenderFirstImage(filename string) error {

  fmt.Printf("P3\n image width = %d , image height = %d, \n255\n", imageWidth, imageHeight)
  file, err := os.Create(filename)
  if err != nil {
    return fmt.Errorf("Error: %w", err)
  }
  defer file.Close()

  // Write PPM header
  fmt.Fprintf(file, "P3\n%d %d\n255\n", imageWidth, imageHeight)

  // Using colorUtil
  for j:=0; j < imageHeight;j++ {
    fmt.Printf("\rScanlines remaining : %d\n", imageHeight - j)

    for i:=0; i < imageWidth;i++ {
      pixelColor := color.Color{
        X: float64(i)/(float64(imageWidth-1)), 
        Y: float64(j)/float64(imageHeight-1), 
        Z: 0}
      color.WriteColor(pixelColor, file)
    }
  }

  fmt.Print("\rDone\n")
  return nil
}

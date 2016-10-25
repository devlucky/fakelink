package images

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"math/rand"
)

func imagesAreEqual(a, b image.Image) bool {
	return a.Bounds().Eq(b.Bounds())
}

func generateRandomImage() *image.RGBA {
	return generateRandomImageWithSize(20, 20)
}

func getFixtureImage(name string) image.Image {
	data, err := ioutil.ReadFile(fmt.Sprintf("../../assets/images/%s", name))
	if err != nil {
		panic(err)
	}

	img, err := jpeg.Decode(bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	return img
}

func generateRandomImageWithSize(width, height int) *image.RGBA {
	bounds := image.Rect(0, 0, width, height)
	img := image.NewRGBA(bounds)
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X-1; x++ {
			a0 := uint8(rand.Float32() * 255)
			rgb0 := uint8(rand.Float32() * 255)
			rgb0 = rgb0 * a0
			img.SetRGBA(x, y, color.RGBA{rgb0, rgb0, rgb0, a0})
		}
	}

	return img
}

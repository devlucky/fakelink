package images

import (
	"image"
	"github.com/disintegration/imaging"
)

// TODO: Benchmark different filters
// TODO: List sizes for different bounds and make estimations

func Thumbnail(img image.Image, maxWidth, maxHeight int) (*image.NRGBA) {
	return imaging.Fit(img, maxWidth, maxHeight, imaging.Lanczos)
}

package images

import (
	"image"
	"github.com/disintegration/imaging"
)

func Thumbnail(img image.Image, maxWidth, maxHeight int) (*image.NRGBA) {
	return imaging.Fit(img, maxWidth, maxHeight, imaging.Lanczos)
}

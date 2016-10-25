package images

import (
	"github.com/disintegration/imaging"
	"image"
)

func Thumbnail(img image.Image, maxWidth, maxHeight int) *image.NRGBA {
	return imaging.Fit(img, maxWidth, maxHeight, imaging.Lanczos)
}

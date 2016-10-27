package images

import (
	"github.com/disintegration/imaging"
	"image"
)

// Thumbnail returns a thumbail of the given image resized to the given dimensions.
func Thumbnail(img image.Image, maxWidth, maxHeight int) *image.NRGBA {
	return imaging.Fit(img, maxWidth, maxHeight, imaging.Lanczos)
}
